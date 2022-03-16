package basket_test

import (
	"context"
	"testing"
	"time"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecocreditApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	basket2 "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
)

func TestPut(t *testing.T) {
	basketDenom := "BASKET"
	basketDenom2 := "BASKET2"
	classId := "C02"
	projectId := "P01"
	startDate, err := time.Parse("2006-01-02", "2020-01-01")
	require.NoError(t, err)
	endDate, err := time.Parse("2006-01-02", "2021-01-01")
	require.NoError(t, err)
	denom, err := ecocredit.FormatDenom(classId, 1, &startDate, &endDate)
	require.NoError(t, err)
	testClassInfo := ecocredit.ClassInfo{
		ClassId:  classId,
		Admin:    "somebody",
		Issuers:  nil,
		Metadata: nil,
		CreditType: &ecocredit.CreditType{
			Name:         "carbon",
			Abbreviation: "C",
			Unit:         "many carbons",
			Precision:    6,
		},
		NumBatches: 1,
	}
	classInfoRes := ecocredit.QueryClassInfoResponse{Info: &testClassInfo}
	testProjectInfo := ecocredit.ProjectInfo{
		ProjectId:       projectId,
		ClassId:         classId,
		Issuer:          "somebody",
		ProjectLocation: "US-NY",
		Metadata:        nil,
	}
	projectInfoRes := ecocredit.QueryProjectInfoResponse{Info: &testProjectInfo}
	testBatchInfo := ecocredit.BatchInfo{
		ProjectId:       projectId,
		BatchDenom:      denom,
		TotalAmount:     "1000000000000000000000000000",
		Metadata:        nil,
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	}
	batchInfoRes := ecocredit.QueryBatchInfoResponse{Info: &testBatchInfo}

	s := setupBase(t)

	err = s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		BasketDenom:       basketDenom,
		Name:              basketDenom,
		DisableAutoRetire: true,
		CreditTypeAbbrev:  "C",
		DateCriteria:      &api.DateCriteria{MinStartDate: timestamppb.New(startDate)},
		Exponent:          6,
	})
	require.NoError(t, err)

	var dur time.Duration = 500000000000000000
	validStartDateWindow := startDate.Add(-dur)
	err = s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		BasketDenom:       basketDenom2,
		Name:              basketDenom2,
		DisableAutoRetire: true,
		CreditTypeAbbrev:  "C",
		DateCriteria:      &api.DateCriteria{StartDateWindow: durationpb.New(dur)},
		Exponent:          6,
	})
	require.NoError(t, err)
	basketDenomToId := make(map[string]uint64)
	basketDenomToId[basketDenom] = 1
	basketDenomToId[basketDenom2] = 2
	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: 1,
		ClassId:  classId,
	})
	require.NoError(t, err)
	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: 2,
		ClassId:  classId,
	})
	require.NoError(t, err)

	s.sdkCtx = s.sdkCtx.WithBlockTime(startDate)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	testCases := []struct {
		name                string
		startingBalance     string
		msg                 basket2.MsgPut
		expectedCredits     []*basket2.BasketCredit
		expectedBasketCoins string
		expectCalls         func()
		errMsg              string
	}{
		{
			name:            "valid case",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedCredits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ProjectInfo(s.ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&classInfoRes, nil)

				coinAward := sdk.NewCoins(sdk.NewCoin(basketDenom, sdk.NewInt(2_000_000)))
				s.bankKeeper.EXPECT().
					MintCoins(s.sdkCtx, basket2.BasketSubModuleName, coinAward).
					Return(nil)

				s.bankKeeper.EXPECT().
					SendCoinsFromModuleToAccount(s.sdkCtx, basket2.BasketSubModuleName, s.addr, coinAward).
					Return(nil)
			},
		},
		{
			name:            "valid case - basket with existing balance",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "1"}},
			},
			expectedCredits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "3"}},
			expectedBasketCoins: "1000000", // 1 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ProjectInfo(s.ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&classInfoRes, nil)

				coinAward := sdk.NewCoins(sdk.NewCoin(basketDenom, sdk.NewInt(1_000_000)))
				s.bankKeeper.EXPECT().
					MintCoins(s.sdkCtx, basket2.BasketSubModuleName, coinAward).
					Return(nil)

				s.bankKeeper.EXPECT().
					SendCoinsFromModuleToAccount(s.sdkCtx, basket2.BasketSubModuleName, s.addr, coinAward).
					Return(nil)
			},
		},
		{
			name:            "valid case - basket 2 with rolling window",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom2,
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedCredits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ProjectInfo(s.ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&classInfoRes, nil)

				coinAward := sdk.NewCoins(sdk.NewCoin(basketDenom2, sdk.NewInt(2_000_000)))
				s.bankKeeper.EXPECT().
					MintCoins(s.sdkCtx, basket2.BasketSubModuleName, coinAward).
					Return(nil)

				s.bankKeeper.EXPECT().
					SendCoinsFromModuleToAccount(s.sdkCtx, basket2.BasketSubModuleName, s.addr, coinAward).
					Return(nil)
			},
		},
		{
			name:            "insufficient funds",
			startingBalance: "1",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ProjectInfo(s.ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&classInfoRes, nil)

			},
			errMsg: basket.ErrInsufficientCredits.Error(),
		},
		{
			name:            "basket not found",
			startingBalance: "1",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: "FooBar",
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
			},
			errMsg: ormerrors.NotFound.Error(),
		},
		{
			name:            "batch not found",
			startingBalance: "20",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: "FooBarBaz", Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: "FooBarBaz"}).
					Return(nil, orm.ErrNotFound)
			},
			errMsg: orm.ErrNotFound.Error(),
		},
		//{
		//	name:            "class not allowed",
		//	startingBalance: "100000000",
		//	msg: basket2.MsgPut{
		//		Owner:       s.addr.String(),
		//		BasketDenom: basketDenom,
		//		Credits:     []*basket2.BasketCredit{{BatchDenom: "blah", Amount: "2"}},
		//	},
		//	expectedBasketCoins: "2000000", // 2 million
		//	expectCalls: func() {
		//		badInfo := *batchInfoRes.Info
		//		badInfo.ClassId = "blah01"
		//		ecocreditKeeper.EXPECT().
		//			BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: "blah"}).
		//			Return(&ecocredit.QueryBatchInfoResponse{Info: &badInfo}, nil)
		//	},
		//	errMsg: "credit class blah01 is not allowed in this basket",
		//},
		{
			name:            "wrong credit type",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ProjectInfo(s.ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				badClass := *classInfoRes.Info
				badClass.CreditType.Abbreviation = "FOO"
				s.ecocreditKeeper.EXPECT().
					ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&ecocredit.QueryClassInfoResponse{Info: &badClass}, nil)
			},
			errMsg: "cannot use credit of type FOO in a basket that requires credit type C",
		},
		{
			name:            "batch out of time window",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				badTime, err := time.Parse("2006-01-02", "1984-01-01")
				require.NoError(t, err)
				badTimeInfo := *batchInfoRes.Info
				badTimeInfo.StartDate = &badTime
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&ecocredit.QueryBatchInfoResponse{Info: &badTimeInfo}, nil)

			},
			errMsg: "cannot put a credit from a batch with start date",
		},
		{
			name:            "batch outside of rolling time window",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom2,
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				badTimeInfo := *batchInfoRes.Info
				bogusDur := time.Duration(999999999999999)
				badTime := validStartDateWindow.Add(-bogusDur)
				badTimeInfo.StartDate = &badTime
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&ecocredit.QueryBatchInfoResponse{Info: &badTimeInfo}, nil)

			},
			errMsg: "cannot put a credit from a batch with start date",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.expectCalls()
			legacyStore := s.sdkCtx.KVStore(s.storeKey)
			tradKey := ecocredit.TradableBalanceKey(s.addr, ecocredit.BatchDenomT(denom))
			userFunds, err := math.NewDecFromString(tc.startingBalance)
			require.NoError(t, err)
			ecocredit.SetDecimal(legacyStore, tradKey, userFunds)
			res, err := s.k.Put(s.ctx, &tc.msg)
			if tc.errMsg == "" { //  no error!
				require.NoError(t, err)
				require.Equal(t, res.AmountReceived, tc.expectedBasketCoins)
				for _, credit := range tc.msg.Credits {
					assertUserSentCredits(t, userFunds, credit.Amount, tradKey, legacyStore)
				}
				for _, credit := range tc.expectedCredits {
					assertBasketHasCredits(t, s.ctx, credit, basketDenomToId[tc.msg.BasketDenom], s.stateStore.BasketBalanceTable())
				}
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errMsg)
			}
		})
	}
}

func assertBasketHasCredits(t *testing.T, ctx context.Context, credit *basket2.BasketCredit, basketID uint64, balanceTable api.BasketBalanceTable) {
	bal, err := balanceTable.Get(ctx, basketID, credit.BatchDenom)
	require.NoError(t, err)
	require.Equal(t, bal.Balance, credit.Amount)
}

func assertUserSentCredits(t *testing.T, oldBalance math.Dec, amountSent string, balanceKey []byte, store types.KVStore) {
	amtSent, err := math.NewDecFromString(amountSent)
	require.NoError(t, err)
	currentBalance, err := ecocredit.GetDecimal(store, balanceKey)
	require.NoError(t, err)

	checkBalance, err := currentBalance.Add(amtSent)
	require.NoError(t, err)

	require.True(t, checkBalance.Equal(oldBalance))
}

type putSuite struct {
	*baseSuite
	blockTimestamp  time.Time
	basketDenom     string
	basketId        uint64
	batch20110101   *ecocreditApi.BatchInfo
	batch20110401   *ecocreditApi.BatchInfo
	batch20110701   *ecocreditApi.BatchInfo
	batch20120101   *ecocreditApi.BatchInfo
	batch20120401   *ecocreditApi.BatchInfo
	batch20120701   *ecocreditApi.BatchInfo
	batch20130101   *ecocreditApi.BatchInfo
	batch20130401   *ecocreditApi.BatchInfo
	batch20130701   *ecocreditApi.BatchInfo
	tradableCredits string
	err             error
}

func TestPutDate(t *testing.T) {
	gocuke.NewRunner(t, &putSuite{}).Path("../../features/basket/put_date.feature").Run()
}

func (s *putSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.tradableCredits = "10"
}

func (s *putSuite) ACurrentBlockTimestampOf20220401() {
	s.blockTimestamp = time.Date(2022, 04, 01, 0, 0, 0, 0, time.UTC)
}

func (s *putSuite) ABasketWithDateCriteriaYearsinthepastOf(a int64) {
	var err error

	s.basketDenom = "foo"

	s.basketId, err = s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:  s.basketDenom,
		DateCriteria: &api.DateCriteria{YearsInThePast: 10},
	})
	assert.NilError(s.t, err)
}

func (s *putSuite) AUserOwnsCreditsFromABatchWithStartDate20110101() {
	var err error

	batchDenom := "batch20110101"
	startDate := timestamppb.New(time.Date(2011, 01, 01, 0, 0, 0, 0, time.UTC))

	err = s.ecocreditStore.BatchInfoTable().Insert(s.ctx, &ecocreditApi.BatchInfo{
		BatchDenom: batchDenom,
		StartDate:  startDate,
	})
	assert.NilError(s.t, err)

	s.batch20110101, err = s.ecocreditStore.BatchInfoTable().GetByBatchDenom(s.ctx, batchDenom)

	err = s.ecocreditStore.BatchBalanceTable().Insert(s.ctx, &ecocreditApi.BatchBalance{
		Address:  s.addr,
		BatchId:  s.batch20110101.Id,
		Tradable: s.tradableCredits,
	})
	assert.NilError(s.t, err)
}

func (s *putSuite) AUserOwnsCreditsFromABatchWithStartDate20110401() {
	var err error

	batchDenom := "batch20110401"
	startDate := timestamppb.New(time.Date(2011, 01, 01, 0, 0, 0, 0, time.UTC))

	err = s.ecocreditStore.BatchInfoTable().Insert(s.ctx, &ecocreditApi.BatchInfo{
		BatchDenom: batchDenom,
		StartDate:  startDate,
	})
	assert.NilError(s.t, err)

	s.batch20110401, err = s.ecocreditStore.BatchInfoTable().GetByBatchDenom(s.ctx, batchDenom)

	err = s.ecocreditStore.BatchBalanceTable().Insert(s.ctx, &ecocreditApi.BatchBalance{
		Address:  s.addr,
		BatchId:  s.batch20110401.Id,
		Tradable: s.tradableCredits,
	})
	assert.NilError(s.t, err)
}

func (s *putSuite) AUserOwnsCreditsFromABatchWithStartDate20110701() {
	var err error

	batchDenom := "batch20110701"
	startDate := timestamppb.New(time.Date(2011, 01, 01, 0, 0, 0, 0, time.UTC))

	err = s.ecocreditStore.BatchInfoTable().Insert(s.ctx, &ecocreditApi.BatchInfo{
		BatchDenom: batchDenom,
		StartDate:  startDate,
	})
	assert.NilError(s.t, err)

	s.batch20110701, err = s.ecocreditStore.BatchInfoTable().GetByBatchDenom(s.ctx, batchDenom)

	err = s.ecocreditStore.BatchBalanceTable().Insert(s.ctx, &ecocreditApi.BatchBalance{
		Address:  s.addr,
		BatchId:  s.batch20110701.Id,
		Tradable: s.tradableCredits,
	})
	assert.NilError(s.t, err)
}

func (s *putSuite) AUserOwnsCreditsFromABatchWithStartDate20120101() {
	var err error

	batchDenom := "batch20120101"
	startDate := timestamppb.New(time.Date(2011, 01, 01, 0, 0, 0, 0, time.UTC))

	err = s.ecocreditStore.BatchInfoTable().Insert(s.ctx, &ecocreditApi.BatchInfo{
		BatchDenom: batchDenom,
		StartDate:  startDate,
	})
	assert.NilError(s.t, err)

	s.batch20120101, err = s.ecocreditStore.BatchInfoTable().GetByBatchDenom(s.ctx, batchDenom)

	err = s.ecocreditStore.BatchBalanceTable().Insert(s.ctx, &ecocreditApi.BatchBalance{
		Address:  s.addr,
		BatchId:  s.batch20120101.Id,
		Tradable: s.tradableCredits,
	})
	assert.NilError(s.t, err)
}

func (s *putSuite) AUserOwnsCreditsFromABatchWithStartDate20120401() {
	var err error

	batchDenom := "batch20120401"
	startDate := timestamppb.New(time.Date(2011, 01, 01, 0, 0, 0, 0, time.UTC))

	err = s.ecocreditStore.BatchInfoTable().Insert(s.ctx, &ecocreditApi.BatchInfo{
		BatchDenom: batchDenom,
		StartDate:  startDate,
	})
	assert.NilError(s.t, err)

	s.batch20120401, err = s.ecocreditStore.BatchInfoTable().GetByBatchDenom(s.ctx, batchDenom)

	err = s.ecocreditStore.BatchBalanceTable().Insert(s.ctx, &ecocreditApi.BatchBalance{
		Address:  s.addr,
		BatchId:  s.batch20120401.Id,
		Tradable: s.tradableCredits,
	})
	assert.NilError(s.t, err)
}

func (s *putSuite) AUserOwnsCreditsFromABatchWithStartDate20120701() {
	var err error

	batchDenom := "batch20120701"
	startDate := timestamppb.New(time.Date(2011, 01, 01, 0, 0, 0, 0, time.UTC))

	err = s.ecocreditStore.BatchInfoTable().Insert(s.ctx, &ecocreditApi.BatchInfo{
		BatchDenom: batchDenom,
		StartDate:  startDate,
	})
	assert.NilError(s.t, err)

	s.batch20120701, err = s.ecocreditStore.BatchInfoTable().GetByBatchDenom(s.ctx, batchDenom)

	err = s.ecocreditStore.BatchBalanceTable().Insert(s.ctx, &ecocreditApi.BatchBalance{
		Address:  s.addr,
		BatchId:  s.batch20120701.Id,
		Tradable: s.tradableCredits,
	})
	assert.NilError(s.t, err)
}

func (s *putSuite) AUserOwnsCreditsFromABatchWithStartDate20130101() {
	var err error

	batchDenom := "batch20130101"
	startDate := timestamppb.New(time.Date(2011, 01, 01, 0, 0, 0, 0, time.UTC))

	err = s.ecocreditStore.BatchInfoTable().Insert(s.ctx, &ecocreditApi.BatchInfo{
		BatchDenom: batchDenom,
		StartDate:  startDate,
	})
	assert.NilError(s.t, err)

	s.batch20130101, err = s.ecocreditStore.BatchInfoTable().GetByBatchDenom(s.ctx, batchDenom)

	err = s.ecocreditStore.BatchBalanceTable().Insert(s.ctx, &ecocreditApi.BatchBalance{
		Address:  s.addr,
		BatchId:  s.batch20130101.Id,
		Tradable: s.tradableCredits,
	})
	assert.NilError(s.t, err)
}

func (s *putSuite) AUserOwnsCreditsFromABatchWithStartDate20130701() {
	var err error

	batchDenom := "batch20130701"
	startDate := timestamppb.New(time.Date(2011, 01, 01, 0, 0, 0, 0, time.UTC))

	err = s.ecocreditStore.BatchInfoTable().Insert(s.ctx, &ecocreditApi.BatchInfo{
		BatchDenom: batchDenom,
		StartDate:  startDate,
	})
	assert.NilError(s.t, err)

	s.batch20130701, err = s.ecocreditStore.BatchInfoTable().GetByBatchDenom(s.ctx, batchDenom)

	err = s.ecocreditStore.BatchBalanceTable().Insert(s.ctx, &ecocreditApi.BatchBalance{
		Address:  s.addr,
		BatchId:  s.batch20130701.Id,
		Tradable: s.tradableCredits,
	})
	assert.NilError(s.t, err)
}

func (s *putSuite) AUserOwnsCreditsFromABatchWithStartDate20130401() {
	var err error

	batchDenom := "batch20130401"
	startDate := timestamppb.New(time.Date(2011, 01, 01, 0, 0, 0, 0, time.UTC))

	err = s.ecocreditStore.BatchInfoTable().Insert(s.ctx, &ecocreditApi.BatchInfo{
		BatchDenom: batchDenom,
		StartDate:  startDate,
	})
	assert.NilError(s.t, err)

	s.batch20130401, err = s.ecocreditStore.BatchInfoTable().GetByBatchDenom(s.ctx, batchDenom)

	err = s.ecocreditStore.BatchBalanceTable().Insert(s.ctx, &ecocreditApi.BatchBalance{
		Address:  s.addr,
		BatchId:  s.batch20130401.Id,
		Tradable: s.tradableCredits,
	})
	assert.NilError(s.t, err)
}

func (s *putSuite) TheUserAttemptsToPutTheCreditsIntoTheBasket() {
	s.ecocreditKeeper.EXPECT().BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: "batch20110101"})

	_, s.err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: s.basketDenom,
		Credits: []*basket.BasketCredit{
			{
				BatchDenom: "batch20110101",
				Amount:     s.tradableCredits,
			},
		},
	})
}

func (s *putSuite) TheCreditsArePutIntoTheBasket() {
	assert.ErrorIs(s.t, s.err, nil)
}

func (s *putSuite) TheCreditsAreNotPutIntoTheBasket() {
	assert.ErrorIs(s.t, s.err, sdkerrors.ErrInvalidRequest) // TODO: start date error
}
