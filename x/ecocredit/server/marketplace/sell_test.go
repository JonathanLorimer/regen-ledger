package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

func TestSell_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom := "C01-20200101-20200201-001"
	start, end := timestamppb.Now(), timestamppb.Now()
	ask := sdk.NewInt64Coin("ufoo", 10)
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).Times(1)

	sellTime := time.Now()
	res, err := s.k.Sell(s.ctx, &v1.MsgSell{
		Owner:  s.addr.String(),
		Orders: []*v1.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10",AskPrice: &ask, DisableAutoRetire: false, Expiration: &sellTime},
			{BatchDenom: batchDenom, Quantity: "10",AskPrice: &ask, DisableAutoRetire: false, Expiration: &sellTime},
		},
	})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.SellOrderIds))

	it, err := s.marketStore.SellOrderStore().List(s.ctx, marketApi.SellOrderSellerIndexKey{})
	assert.NilError(t, err)
	count := 0
	for it.Next() {
		val, err := it.Value()
		assert.NilError(t, err)
		assert.Equal(t, "10", val.Quantity)
		assert.Equal(t, ask.Denom, val.AskPrice)
		count++
	}
	assert.Equal(t, 2, count)
}

func TestSell_Invalid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	any := gomock.Any()
	batchDenom := "C01-20200101-20200201-001"
	start, end := timestamppb.Now(), timestamppb.Now()
	ask := sdk.NewInt64Coin("ufoo", 10)
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	sellTime := time.Now()

	// invalid batch
	_, err := s.k.Sell(s.ctx, &v1.MsgSell{
		Owner:  s.addr.String(),
		Orders: []*v1.MsgSell_Order{
			{BatchDenom: "foo-bar-baz-001", Quantity: "10", AskPrice: &ask, DisableAutoRetire: true, Expiration: &sellTime},
		},
	})
	assert.ErrorContains(t, err, "batch denom foo-bar-baz-001")

	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).Times(3)

	// market not found
	invalidDenom := sdk.NewInt64Coin("baz", 50)
	_, err = s.k.Sell(s.ctx, &v1.MsgSell{
		Owner:  s.addr.String(),
		Orders: []*v1.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &invalidDenom, DisableAutoRetire: true, Expiration: &sellTime},
		},
	})
	assert.ErrorContains(t, err, "market: batch denom", batchDenom)

	// invalid balance
	_, err = s.k.Sell(s.ctx, &v1.MsgSell{
		Owner:  s.addr.String(),
		Orders: []*v1.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10000000000", AskPrice: &ask, DisableAutoRetire: true, Expiration: &sellTime},
		},
	})
	assert.ErrorContains(t, err, ecocredit.ErrInsufficientFunds.Error())

	// order expiration not in the future
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now())
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)
	invalidExpirationTime, err := time.Parse("2006-01-02", "1500-01-01")
	_, err = s.k.Sell(s.ctx, &v1.MsgSell{
		Owner:  s.addr.String(),
		Orders: []*v1.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: true, Expiration: &invalidExpirationTime},
		},
	})
	assert.ErrorContains(t, err, "expiration must be in the future")
}

func testSellSetup(t *testing.T, s *baseSuite, batchDenom, bankDenom, displayDenom, classId string, start, end *timestamppb.Timestamp, creditType ecocredit.CreditType) {
	assert.NilError(t, s.coreStore.BatchInfoStore().Insert(s.ctx, &ecocreditv1.BatchInfo{
		ProjectId:  1,
		BatchDenom: batchDenom,
		Metadata:   nil,
		StartDate:  start,
		EndDate:    end,
	}))
	assert.NilError(t, s.coreStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1.ClassInfo{
		Name:       classId,
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: creditType.Abbreviation,
	}))
	assert.NilError(t, s.marketStore.MarketStore().Insert(s.ctx, &marketApi.Market{
		CreditType:        creditType.Abbreviation,
		BankDenom:         bankDenom,
		PrecisionModifier: 0,
	}))
	assert.NilError(t, s.marketStore.AllowedDenomStore().Insert(s.ctx, &marketApi.AllowedDenom{
		BankDenom:    bankDenom,
		DisplayDenom: displayDenom,
		Exponent:     1,
	}))
	assert.NilError(t, s.k.coreStore.BatchBalanceStore().Insert(s.ctx, &ecocreditv1.BatchBalance{
		Address:  s.addr,
		BatchId:  1,
		Tradable: "100",
		Retired:  "100",
	}))
}
