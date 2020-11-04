package client_test

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	gocid "github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"github.com/regen-network/regen-ledger/app"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/grpc"
	"testing"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	from        sdk.AccAddress
	msgClient   data.MsgClient
	queryClient data.QueryClient
	ctx         context.Context
	txFactory   clienttx.Factory
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := network.DefaultConfig()
	encConfig := app.MakeEncodingConfig()
	cfg.NumValidators = 1
	cfg.InterfaceRegistry = encConfig.InterfaceRegistry
	cfg.TxConfig = encConfig.TxConfig
	cfg.Codec = encConfig.Marshaler
	cfg.LegacyAmino = encConfig.Amino
	cfg.AppConstructor = func(val network.Validator) types.Application {
		return app.NewRegenApp(
			val.Ctx.Logger, dbm.NewMemDB(), nil, true, make(map[int64]bool), val.Ctx.Config.RootDir, 0,
			encConfig,
			baseapp.SetPruning(storetypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
			baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
		)
	}
	cfg.GenesisState = app.NewDefaultGenesisState()

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	s.ctx = context.Background()
	val := s.network.Validators[0]
	s.from = val.Address
	s.queryClient = data.NewQueryClient(val.ClientCtx)
	s.txFactory = clienttx.NewFactoryCLI(val.ClientCtx, pflag.NewFlagSet("test", pflag.ContinueOnError))
	clientCtx := val.ClientCtx
	clientCtx = clientCtx.WithFrom(val.Address.String())
	fromAddr, fromName, err := client.GetFromFields(clientCtx.Keyring, clientCtx.From, false)
	clientCtx = clientCtx.
		WithFromAddress(fromAddr).
		WithFromName(fromName)
	s.txFactory = s.txFactory.WithFees(sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String())
	txConn := ClientTxConn{
		ClientContext: clientCtx,
		Factory:       s.txFactory,
	}
	s.msgClient = data.NewMsgClient(txConn)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

type ClientTxConn struct {
	ClientContext client.Context
	Factory       clienttx.Factory
}

func (c ClientTxConn) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	req, ok := args.(sdk.MsgRequest)
	if !ok {
		return fmt.Errorf("%T should implement %T", args, (*sdk.MsgRequest)(nil))
	}

	err := req.ValidateBasic()
	if err != nil {
		return err
	}

	msg := sdk.ServiceMsg{
		MethodName: method,
		Request:    req,
	}

	clientCtx := c.ClientContext
	txf := c.Factory

	txf, err = clienttx.PrepareFactory(clientCtx, txf)
	if err != nil {
		return err
	}

	_, adjusted, err := clienttx.CalculateGas(clientCtx.QueryWithData, txf, msg)
	if err != nil {
		return err
	}

	txf = txf.WithGas(adjusted)

	tx, err := clienttx.BuildUnsignedTx(txf, msg)
	if err != nil {
		return err
	}

	err = clienttx.Sign(txf, clientCtx.GetFromName(), tx)
	if err != nil {
		return err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(tx.GetTx())
	if err != nil {
		return err
	}

	clientCtx = clientCtx.WithBroadcastMode(flags.BroadcastBlock)

	// broadcast to a Tendermint node
	txRes, err := clientCtx.BroadcastTx(txBytes)
	if err != nil {
		return err
	}

	protoRes, ok := reply.(proto.Message)
	if !ok {
		return fmt.Errorf("expected proto.Message, got %T", reply)
	}

	if txRes.Code != 0 {
		return sdkerrors.ABCIError(txRes.Codespace, txRes.Code, txRes.RawLog)
	}

	return proto.Unmarshal([]byte(txRes.Data), protoRes)
}

func (c ClientTxConn) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("unsupported")
}

var _ grpc.ClientConnInterface = ClientTxConn{}

func (s *IntegrationTestSuite) TestScenario() {
	testContent := []byte("xyzabc123")
	mh, err := multihash.Sum(testContent, multihash.SHA2_256, -1)
	s.Require().NoError(err)
	cid := gocid.NewCidV1(gocid.Raw, mh)

	// anchor some data
	cidBz := cid.Bytes()
	anchorRes, err := s.msgClient.AnchorData(s.ctx, &data.MsgAnchorDataRequest{
		Sender: s.from.String(),
		Cid:    cidBz,
	})
	s.Require().NoError(err)
	s.Require().NotNil(anchorRes)

	// can't anchor same data twice
	_, err = s.msgClient.AnchorData(s.ctx, &data.MsgAnchorDataRequest{
		Sender: s.from.String(),
		Cid:    cidBz,
	})
	s.Require().Error(err)

	// can query data and get timestamp
	queryRes, err := s.queryClient.Data(s.ctx, &data.QueryDataRequest{Cid: cidBz})
	s.Require().NoError(err)
	s.Require().NotNil(queryRes)
	s.Require().Equal(anchorRes.Timestamp, queryRes.Timestamp)
	s.Require().Empty(queryRes.Signers)
	s.Require().Empty(queryRes.Content)

	// can sign data
	_, err = s.msgClient.SignData(s.ctx, &data.MsgSignDataRequest{
		Signers: []string{s.from.String()},
		Cid:     cidBz,
	})
	s.Require().NoError(err)

	// can retrieve signature, same timestamp
	// can query data and get timestamp
	queryRes, err = s.queryClient.Data(s.ctx, &data.QueryDataRequest{Cid: cidBz})
	s.Require().NoError(err)
	s.Require().NotNil(queryRes)
	s.Require().Equal(anchorRes.Timestamp, queryRes.Timestamp)
	s.Require().Equal([]string{s.from.String()}, queryRes.Signers)
	s.Require().Empty(queryRes.Content)

	// can't store bad data
	_, err = s.msgClient.StoreData(s.ctx, &data.MsgStoreDataRequest{
		Sender:  s.from.String(),
		Cid:     cidBz,
		Content: []byte("sgkjhsgouiyh"),
	})
	s.Require().Error(err)

	// can store good data
	_, err = s.msgClient.StoreData(s.ctx, &data.MsgStoreDataRequest{
		Sender:  s.from.String(),
		Cid:     cidBz,
		Content: testContent,
	})
	s.Require().NoError(err)

	// can retrieve signature, same timestamp, and data
	s.Require().NoError(err)
	s.Require().NotNil(queryRes)
	s.Require().Equal(anchorRes.Timestamp, queryRes.Timestamp)
	s.Require().Equal([]string{s.from.String()}, queryRes.Signers)
	s.Require().Equal(testContent, queryRes.Content)
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}