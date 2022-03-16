package basket_test

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecocreditApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	mocks2 "github.com/regen-network/regen-ledger/x/ecocredit/mocks"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket/mocks"
)

type baseSuite struct {
	t               gocuke.TestingT
	db              ormdb.ModuleDB
	ctx             context.Context
	k               basket.Keeper
	ctrl            *gomock.Controller
	addr            sdk.AccAddress
	stateStore      api.StateStore
	ecocreditStore  ecocreditApi.StateStore
	bankKeeper      *mocks2.MockBankKeeper
	ecocreditKeeper *mocks.MockEcocreditKeeper
	storeKey        *sdk.KVStoreKey
	sdkCtx          sdk.Context
	distKeeper      *mocks2.MockDistributionKeeper
}

func setupBase(t gocuke.TestingT) *baseSuite {
	// prepare database
	s := &baseSuite{t: t}
	var err error
	s.db, err = ormdb.NewModuleDB(&server.ModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)
	s.stateStore, err = api.NewStateStore(s.db)
	assert.NilError(t, err)
	s.ecocreditStore, err = ecocreditApi.NewStateStore(s.db)
	assert.NilError(t, err)

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	s.storeKey = sdk.NewKVStoreKey("test")
	cms.MountStoreWithDB(s.storeKey, sdk.StoreTypeIAVL, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	s.sdkCtx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// setup test keeper
	s.ctrl = gomock.NewController(t)
	assert.NilError(t, err)
	s.bankKeeper = mocks2.NewMockBankKeeper(s.ctrl)
	s.ecocreditKeeper = mocks.NewMockEcocreditKeeper(s.ctrl)
	s.distKeeper = mocks2.NewMockDistributionKeeper(s.ctrl)
	s.k = basket.NewKeeper(s.db, s.ecocreditKeeper, s.bankKeeper, s.distKeeper, s.storeKey)

	_, _, s.addr = testdata.KeyTestPubAddr()

	return s
}
