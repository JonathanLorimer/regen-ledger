package marketplace

import (
	"context"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/golang/mock/gomock"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

var (
	any = gomock.Any()
	batchDenom = "C01-20200101-20200201-001"
	classId = "C01"
	start, end = timestamppb.Now(), timestamppb.Now()
	ask = sdk.NewInt64Coin("ufoo", 10)
	creditType = ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
)

func TestUpdateSellOrders_Quantity(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)

	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).AnyTimes()
	expiration := time.Now()
	_, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner:  s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "5.22", AskPrice: &ask, DisableAutoRetire: false, Expiration: &expiration},
			{BatchDenom: batchDenom, Quantity: "30", AskPrice: &ask, DisableAutoRetire: true, Expiration: &expiration},
		},
	})
	assert.NilError(t, err)

	balBefore, supBefore := getBalanceAndSupply(t, s.ctx, s.coreStore, 1, s.addr)

	_, err = s.k.UpdateSellOrders(s.ctx, &marketplace.MsgUpdateSellOrders{
		Owner:   s.addr.String(),
		Updates: []*marketplace.MsgUpdateSellOrders_Update{
			{SellOrderId: 1, NewQuantity: "10", DisableAutoRetire: true},
			{SellOrderId: 2, NewQuantity: "28.7232", DisableAutoRetire: false},
		},
	})
	assert.NilError(t, err)

	balAfter, supAfter := getBalanceAndSupply(t, s.ctx, s.coreStore, 1, s.addr)

	// sellOrder 1: 5.22 originally, increased by 10 = change of 4.78
	// sellOrder 2: 30 originally, decreased by 28.7232 = change of -1.2768
	// total change: 4.78 + (-1.2768) = 3.5032

	actualEscrowChange, err := math.NewDecFromString("3.5032")
	assert.NilError(t, err)

	assertCoinsEscrowed(t, balBefore, balAfter, supBefore, supAfter, actualEscrowChange)

	order1, err := s.marketStore.SellOrderTable().Get(s.ctx, 1)
	assert.NilError(t, err)
	order2, err := s.marketStore.SellOrderTable().Get(s.ctx, 2)
	assert.NilError(t, err)
	assert.Equal(t, "10", order1.Quantity)
	assert.Equal(t, "28.7232", order2.Quantity)

	assert.Equal(t, true, order1.DisableAutoRetire)
	assert.Equal(t, false, order2.DisableAutoRetire)
}

func TestUpdateSellOrders_QuantityInvalid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)

	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).AnyTimes()
	expiration := time.Now()
	_, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner:  s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "5.22", AskPrice: &ask, DisableAutoRetire: false, Expiration: &expiration},
			{BatchDenom: batchDenom, Quantity: "30", AskPrice: &ask, DisableAutoRetire: true, Expiration: &expiration},
		},
	})
	assert.NilError(t, err)

	// cannot increase sell order with more credits than in balance
	_, err = s.k.UpdateSellOrders(s.ctx, &marketplace.MsgUpdateSellOrders{
		Owner:   s.addr.String(),
		Updates: []*marketplace.MsgUpdateSellOrders_Update{
			{SellOrderId: 1, NewQuantity: "1000000000"},
		},
	})
	assert.ErrorContains(t, err, sdkerrors.ErrInsufficientFunds.Error())
}

func TestUpdateSellOrders_Unauthorized(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)
	_,_,unauthorized := testdata.KeyTestPubAddr()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).AnyTimes()
	expiration := time.Now()
	_, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner:  s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "5.22", AskPrice: &ask, DisableAutoRetire: false, Expiration: &expiration},
			{BatchDenom: batchDenom, Quantity: "30", AskPrice: &ask, DisableAutoRetire: true, Expiration: &expiration},
		},
	})
	assert.NilError(t, err)

	// cannot edit the sell order with this address
	_, err = s.k.UpdateSellOrders(s.ctx, &marketplace.MsgUpdateSellOrders{
		Owner:   unauthorized.String(),
		Updates: []*marketplace.MsgUpdateSellOrders_Update{
			{SellOrderId: 1, NewQuantity: "1"},
		},
	})
	assert.ErrorContains(t, err, sdkerrors.ErrUnauthorized.Error())
}

func TestUpdateSellOrder_AskPrice(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)

	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).AnyTimes()
	expiration := time.Now()
	_, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner:  s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "5.22", AskPrice: &ask, DisableAutoRetire: false, Expiration: &expiration},
		},
	})
	assert.NilError(t, err)

	orderBefore, err := s.marketStore.SellOrderTable().Get(s.ctx, 1)
	// can update price of same denom
	askUpdate := sdk.NewInt64Coin(ask.Denom, 25)
	_, err = s.k.UpdateSellOrders(s.ctx, &marketplace.MsgUpdateSellOrders{
		Owner:   s.addr.String(),
		Updates: []*marketplace.MsgUpdateSellOrders_Update{
			{SellOrderId: 1, NewAskPrice: &askUpdate},
		},
	})
	assert.NilError(t, err)
	orderAfter, err := s.marketStore.SellOrderTable().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, orderAfter.AskPrice, askUpdate.Amount.String())
	assert.Equal(t, orderBefore.MarketId, orderAfter.MarketId)

	orderBefore, err = s.marketStore.SellOrderTable().Get(s.ctx, 1)
	assert.NilError(t, err)

	askUpdate = sdk.NewInt64Coin("ubar", 18)
	_, err = s.k.UpdateSellOrders(s.ctx, &marketplace.MsgUpdateSellOrders{
		Owner:   s.addr.String(),
		Updates: []*marketplace.MsgUpdateSellOrders_Update{
			{SellOrderId: 1, NewAskPrice: &askUpdate},
		},
	})
	assert.NilError(t, err)

	order, err := s.marketStore.SellOrderTable().Get(s.ctx, 1)
	assert.NilError(t, err)

	assert.Equal(t, order.AskPrice, askUpdate.Amount.String())
	assert.Equal(t, order.MarketId, orderBefore.MarketId + 1)
}

func TestUpdateSellOrder_Expiration(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)

	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).AnyTimes()

	future := time.Date(2077, 1,1,1,1,1,1, time.Local)
	middle := time.Date(2022,1,1,1,1,1,1,time.Local)
	past := time.Date(1970, 1,1,1,1,1,1,time.Local)

	_, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner:  s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "5.22", AskPrice: &ask, DisableAutoRetire: false, Expiration: &past},
		},
	})
	assert.NilError(t, err)


	s.sdkCtx = s.sdkCtx.WithBlockTime(middle)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// should work with future time
	_, err = s.k.UpdateSellOrders(s.ctx, &marketplace.MsgUpdateSellOrders{
		Owner:   s.addr.String(),
		Updates: []*marketplace.MsgUpdateSellOrders_Update{
			{SellOrderId: 1, NewExpiration: &future},
		},
	})
	assert.NilError(t, err)

	// should not work with past time
	_, err = s.k.UpdateSellOrders(s.ctx, &marketplace.MsgUpdateSellOrders{
		Owner:   s.addr.String(),
		Updates: []*marketplace.MsgUpdateSellOrders_Update{
			{SellOrderId: 1, NewExpiration: &past},
		},
	})
	assert.ErrorContains(t, err, "expiration must be in the future")
}

func getBalanceAndSupply(t *testing.T, ctx context.Context, store ecocreditv1.StateStore, batchId uint64, addr sdk.AccAddress) (*ecocreditv1.BatchBalance, *ecocreditv1.BatchSupply){
	bal, err := store.BatchBalanceTable().Get(ctx, addr, batchId)
	assert.NilError(t, err)
	sup, err := store.BatchSupplyTable().Get(ctx, batchId)
	assert.NilError(t, err)
	return bal, sup
}