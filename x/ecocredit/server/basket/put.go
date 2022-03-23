package basket

import (
	"context"
	"time"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/orm"
	regenmath "github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

// Put deposits ecocredits into a basket, returning fungible coins to the depositor.
// NOTE: the credits MUST adhere to the following specifications set by the basket: credit type, class, and date criteria.
func (k Keeper) Put(ctx context.Context, req *baskettypes.MsgPut) (*baskettypes.MsgPutResponse, error) {
	ownerAddr, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	// get the basket
	basket, err := k.stateStore.BasketTable().GetByBasketDenom(ctx, req.BasketDenom)
	if err != nil {
		if orm.ErrNotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("basket %s not found", req.BasketDenom)
		}
		return nil, err
	}

	// keep track of the total amount of tokens to give to the depositor
	amountReceived := sdk.NewInt(0)
	sdkContext := sdk.UnwrapSDKContext(ctx)
	for _, credit := range req.Credits {
		sdkContext.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/basket/MsgPut iteration")
		// get credit batch info
		res, err := k.ecocreditKeeper.BatchInfo(ctx, &core.QueryBatchInfoRequest{BatchDenom: credit.BatchDenom})
		if err != nil {
			if orm.ErrNotFound.Is(err) {
				return nil, sdkerrors.ErrNotFound.Wrapf("%s batch not found", credit.BatchDenom)
			}
			return nil, err
		}
		batchInfo := res.Info

		// validate that the credit batch adheres to the basket's specifications
		if err := k.canBasketAcceptCredit(ctx, basket, batchInfo); err != nil {
			return nil, err
		}
		// get the amount of credits in dec
		amt, err := regenmath.NewPositiveFixedDecFromString(credit.Amount, basket.Exponent)
		if err != nil {
			return nil, err
		}
		// update the user and basket balances
		if err = k.transferToBasket(ctx, ownerAddr, amt, basket, batchInfo); err != nil {
			if sdkerrors.ErrInsufficientFunds.Is(err) {
				return nil, ErrInsufficientCredits
			}
			return nil, err
		}
		// get the amount of basket tokens to give to the depositor
		tokens, err := creditAmountToBasketCoins(amt, basket.Exponent, basket.BasketDenom)
		if err != nil {
			return nil, err
		}
		// update the total amount received so far
		amountReceived = amountReceived.Add(tokens[0].Amount)
	}

	// mint and send tokens to depositor
	coinsToSend := sdk.Coins{sdk.NewCoin(basket.BasketDenom, amountReceived)}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if err = k.bankKeeper.MintCoins(sdkCtx, baskettypes.BasketSubModuleName, coinsToSend); err != nil {
		return nil, err
	}
	if err = k.bankKeeper.SendCoinsFromModuleToAccount(sdkCtx, baskettypes.BasketSubModuleName, ownerAddr, coinsToSend); err != nil {
		return nil, err
	}

	if err = sdkCtx.EventManager().EmitTypedEvent(&baskettypes.EventPut{
		Owner:       ownerAddr.String(),
		BasketDenom: basket.BasketDenom,
		Credits:     req.Credits,
		Amount:      amountReceived.String(),
	}); err != nil {
		return nil, err
	}

	return &baskettypes.MsgPutResponse{AmountReceived: amountReceived.String()}, nil
}

// canBasketAcceptCredit checks that a credit adheres to the specifications of a basket. Specifically, it checks:
//  - batch's start time is within the basket's specified time window or min start date
//  - class is in the basket's allowed class store
//  - type matches the baskets specified credit type.
func (k Keeper) canBasketAcceptCredit(ctx context.Context, basket *api.Basket, batchInfo *core.BatchInfo) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockTime := sdkCtx.BlockTime()
	errInvalidReq := sdkerrors.ErrInvalidRequest

	if basket.DateCriteria != nil {
		// check time window match
		var minStartDate time.Time
		var criteria = basket.DateCriteria
		if criteria.MinStartDate != nil {
			minStartDate = criteria.MinStartDate.AsTime()
		} else if criteria.StartDateWindow != nil {
			window := criteria.StartDateWindow.AsDuration()
			minStartDate = blockTime.Add(-window)
		} else if criteria.YearsInThePast != 0 {
			year := blockTime.Year() - int(criteria.YearsInThePast)
			minStartDate = time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
		}

		timestamp := timestamppb.Timestamp{
			Seconds: batchInfo.StartDate.Seconds,
			Nanos:   batchInfo.StartDate.Nanos,
		}
		startDate := timestamp.AsTime()
		if startDate.Before(minStartDate) {
			return errInvalidReq.Wrapf("cannot put a credit from a batch with start date %s "+
				"into a basket that requires an earliest start date of %s", batchInfo.StartDate.String(), minStartDate.String())
		}

	}

	classId := ecocredit.GetClassIdFromBatchDenom(batchInfo.BatchDenom)

	// check credit class match
	// TODO: do we even need to check the credit type at this point? theres no way the class
	// 		would be accepted otherwise
	found, err := k.stateStore.BasketClassTable().Has(ctx, basket.Id, classId)
	if err != nil {
		return err
	}
	if !found {
		return errInvalidReq.Wrapf("credit class %s is not allowed in this basket", classId)
	}
	return nil
}

// transferToBasket updates the balance of the user in the legacy KVStore as well as the basket's balance in the ORM.
func (k Keeper) transferToBasket(ctx context.Context, sender sdk.AccAddress, amt regenmath.Dec, basket *api.Basket, batchInfo *core.BatchInfo) error {
	// update user balance, subtracting from their tradable balance
	userBal, err := k.coreStore.BatchBalanceTable().Get(ctx, sender, batchInfo.Id)
	if err != nil {
		return ecocredit.ErrInsufficientFunds.Wrapf("could not get batch %s balance for %s", batchInfo.BatchDenom, sender.String())
	}
	tradable, err := regenmath.NewPositiveDecFromString(userBal.Tradable)
	if err != nil {
		return err
	}
	newTradable, err := regenmath.SafeSubBalance(tradable, amt)
	if err != nil {
		return ecocredit.ErrInsufficientFunds.Wrapf("cannot put %v credits into the basket with a balance of %v: %s", amt, tradable, err.Error())
	}
	userBal.Tradable = newTradable.String()
	if err = k.coreStore.BatchBalanceTable().Save(ctx, userBal); err != nil {
		return err
	}

	// update basket balance with amount sent, adding to the basket's balance.
	var bal *api.BasketBalance
	bal, err = k.stateStore.BasketBalanceTable().Get(ctx, basket.Id, batchInfo.BatchDenom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			bal = &api.BasketBalance{
				BasketId:   basket.Id,
				BatchDenom: batchInfo.BatchDenom,
				Balance:    amt.String(),
				BatchStartDate: &timestamppb.Timestamp{
					Seconds: batchInfo.StartDate.Seconds,
					Nanos:   batchInfo.StartDate.Nanos,
				},
			}
		} else {
			return err
		}
	} else {
		newBalance, err := regenmath.NewPositiveFixedDecFromString(bal.Balance, basket.Exponent)
		if err != nil {
			return err
		}
		newBalance, err = newBalance.Add(amt)
		if err != nil {
			return err
		}
		bal.Balance = newBalance.String()
	}
	if err = k.stateStore.BasketBalanceTable().Save(ctx, bal); err != nil {
		return err
	}
	return nil
}

// creditAmountToBasketCoins calculates the tokens to award to the depositor
func creditAmountToBasketCoins(creditAmt regenmath.Dec, exp uint32, denom string) (sdk.Coins, error) {
	var coins sdk.Coins
	multiplier := regenmath.NewDecFinite(1, int32(exp))
	tokenAmt, err := multiplier.MulExact(creditAmt)
	if err != nil {
		return coins, err
	}

	amtInt, err := tokenAmt.BigInt()
	if err != nil {
		return coins, err
	}

	return sdk.Coins{sdk.NewCoin(denom, sdk.NewIntFromBigInt(amtInt))}, nil
}
