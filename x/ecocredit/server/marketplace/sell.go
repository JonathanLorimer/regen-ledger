package marketplace

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	marketplacev1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Sell creates new sell orders for credits
func (k Keeper) Sell(ctx context.Context, req *marketplacev1.MsgSell) (*marketplacev1.MsgSellResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	ownerAcc, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	sellOrderIds := make([]uint64, len(req.Orders))

	for i, order := range req.Orders {
		batch, err := k.coreStore.BatchInfoStore().GetByBatchDenom(ctx, order.BatchDenom)
		if err != nil {
			return nil, fmt.Errorf("batch denom %v: %v", order.BatchDenom, err)
		}
		ct, err := server.GetCreditTypeFromBatchDenom(ctx, k.coreStore, k.params, batch.BatchDenom)
		if err != nil {
			return nil, err
		}
		marketId, err := k.getOrCreateMarketId(ctx, ct.Abbreviation, order.AskPrice.Denom)
		// verify expiration is in the future
		if order.Expiration != nil && order.Expiration.Before(sdkCtx.BlockTime()) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("expiration must be in the future: %s", order.Expiration)
		}
		sellQty, err := math.NewDecFromString(order.Quantity)
		if err != nil {
			return nil, err
		}
		if err := assertHasBalance(ctx, k.coreStore, ownerAcc, batch.Id, sellQty); err != nil {
			return nil, err
		}
		has, err := isDenomAllowed(ctx, k.stateStore, order.AskPrice.Denom)
		if err != nil {
			return nil, err
		}
		if !has {
			return nil, ecocredit.ErrInvalidSellOrder.Wrapf("cannot use coin with denom %s in sell orders", order.AskPrice.Denom)
		}
		id, err := k.stateStore.SellOrderStore().InsertReturningID(ctx, &marketApi.SellOrder{
			Seller:            ownerAcc,
			BatchId:           batch.Id,
			Quantity:          order.Quantity,
			MarketId:          marketId,
			AskPrice:          order.AskPrice.Amount.String(),
			DisableAutoRetire: order.DisableAutoRetire,
			Expiration:        timestamppb.New(*order.Expiration),
			Maker:             true, // maker is always true for sell orders
		})
		if err != nil {
			return nil, err
		}
		sellOrderIds[i] = id
		sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "create sell order")
	}
	return &marketplacev1.MsgSellResponse{SellOrderIds: sellOrderIds}, nil
}

// getOrCreateMarketId attempts to get a market, creating one otherwise, and return the Id.
func (k Keeper) getOrCreateMarketId(ctx context.Context, creditTypeAbbrev, bankDenom string) (uint64, error) {
	market, err := k.stateStore.MarketStore().GetByCreditTypeBankDenom(ctx, creditTypeAbbrev, bankDenom)
	switch err {
	case nil:
		return market.Id, nil
	case ormerrors.NotFound:
		return k.stateStore.MarketStore().InsertReturningID(ctx, &marketApi.Market{
			CreditType:        creditTypeAbbrev,
			BankDenom:         bankDenom,
			PrecisionModifier: 0,
		})
	default:
		return 0, err
	}
}
