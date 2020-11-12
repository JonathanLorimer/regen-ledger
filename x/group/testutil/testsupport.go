package testutil

import (
	"github.com/cosmos/cosmos-sdk/store/rootmulti"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/group/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

func NewContext(keys ...sdk.StoreKey) sdk.Context {
	db := dbm.NewMemDB()
	cms := rootmulti.NewStore(db)
	for _, v := range keys {
		storeType := sdk.StoreTypeIAVL
		if _, ok := v.(*sdk.TransientStoreKey); ok {
			storeType = sdk.StoreTypeTransient
		}
		cms.MountStoreWithDB(v, storeType, db)
		if err := cms.LoadLatestVersion(); err != nil {
			panic(err)
		}
	}
	return sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
}

// func CreateGroupKeeper() (types.Keeper, sdk.Context) {
// 	encodingConfig := app.MakeEncodingConfig()
// 	pKey, pTKey := sdk.NewKVStoreKey(paramstypes.StoreKey), sdk.NewTransientStoreKey(paramstypes.TStoreKey)
// 	paramSpace := paramstypes.NewSubspace(encodingConfig.Marshaler, encodingConfig.Amino, pKey, pTKey, types.DefaultParamspace)

// 	groupKey := sdk.NewKVStoreKey(types.StoreKey)
// 	k := types.NewGroupKeeper(groupKey, paramSpace, baseapp.NewRouter(), &MockProposalI{})
// 	ctx := NewContext(pKey, pTKey, groupKey)
// 	k.SetParams(ctx, types.DefaultParams())
// 	return k, ctx
// }

type MockProposalI struct {
}

func (m MockProposalI) Marshal() ([]byte, error) {
	panic("implement me")
}

func (m MockProposalI) Unmarshal([]byte) error {
	panic("implement me")
}

func (m MockProposalI) GetBase() types.ProposalBase {
	panic("implement me")
}

func (m MockProposalI) SetBase(types.ProposalBase) {
	panic("implement me")
}

func (m MockProposalI) GetMsgs() []sdk.Msg {
	panic("implement me")
}

func (m MockProposalI) SetMsgs([]sdk.Msg) error {
	panic("implement me")
}

func (m MockProposalI) ValidateBasic() error {
	panic("implement me")
}