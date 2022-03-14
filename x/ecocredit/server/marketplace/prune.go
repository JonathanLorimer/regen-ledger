package marketplace

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// PruneSellOrders is an EndBlock function that moves escrowed credits back into their tradable balance and deletes orders
// that have expired.
func (k Keeper) PruneSellOrders(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	min, blockTime := timestamppb.New(time.Now()), timestamppb.New(sdkCtx.BlockTime())
	fromKey, toKey := marketplacev1.SellOrderExpirationIndexKey{}.WithExpiration(min), marketplacev1.SellOrderExpirationIndexKey{}.WithExpiration(blockTime)
	it, err := k.stateStore.SellOrderStore().ListRange(ctx, fromKey, toKey)
	if err != nil {
		return err
	}
	defer it.Close()
	for it.Next() {
		sellOrder, err := it.Value()
		if err != nil {
			return err
		}
		if err = k.unescrowCredits(ctx, sellOrder.Seller, sellOrder.BatchId, sellOrder.Quantity); err != nil {
			return err
		}
	}
	return k.stateStore.SellOrderStore().DeleteRange(ctx, fromKey, toKey)
}

// unescrowCredits moves `amount` of credits from the sellerAddr's escrowed balance, into their tradable balance
// as well as from the supply's escrowed amount, into the supply's tradable amount.
func (k Keeper) unescrowCredits(ctx context.Context, sellerAddr sdk.AccAddress, batchId uint64, amount string) error {

	creditAmt, err := math.NewDecFromString(amount)
	if err != nil {
		return err
	}

	moveCredits := func(escrowed, tradable string) (string, string, error) {
		escrowedDec, err := math.NewDecFromString(escrowed)
		if err != nil {
			return "", "", err
		}
		tradableDec, err := math.NewDecFromString(tradable)
		if err != nil {
			return "", "", err
		}

		// subtract from escrow
		escrowedDec, err = math.SafeSubBalance(escrowedDec, creditAmt)
		if err != nil {
			return "", "", err
		}
		// add to tradable
		tradableDec, err = math.SafeAddBalance(tradableDec, creditAmt)
		if err != nil {
			return "", "", err
		}
		return escrowedDec.String(), tradableDec.String(), nil
	}

	bal, err := k.coreStore.BatchBalanceStore().Get(ctx, sellerAddr, batchId)
	if err != nil {
		return err
	}
	bal.Escrowed, bal.Tradable, err = moveCredits(bal.Escrowed, bal.Tradable)
	if err != nil {
		return err
	}
	if err = k.coreStore.BatchBalanceStore().Update(ctx, bal); err != nil {
		return err
	}

	// update supply
	sup, err := k.coreStore.BatchSupplyStore().Get(ctx, batchId)
	if err != nil {
		return err
	}
	sup.EscrowedAmount, sup.TradableAmount, err = moveCredits(sup.EscrowedAmount, sup.TradableAmount)
	if err != nil {
		return err
	}
	return k.coreStore.BatchSupplyStore().Update(ctx, sup)
}
