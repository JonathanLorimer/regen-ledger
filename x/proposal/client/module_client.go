package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
	proposalcmd "gitlab.com/regen-network/regen-ledger/x/proposal/client/cli"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	cdc *amino.Codec
}

func NewModuleClient(cdc *amino.Codec) ModuleClient {
	return ModuleClient{cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	proposalQueryCmd := &cobra.Command{
		Use:   "proposal",
		Short: "Querying commands for the proposal module",
	}

	//proposalQueryCmd.AddCommand(client.GetCommands(
	//	//proposalcmd.GetCmdGetData(mc.storeKey, mc.cdc),
	//)...)

	return proposalQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	proposalTxCmd := &cobra.Command{
		Use:   "proposal",
		Short: "Proposal transactions subcommands",
	}

	cdc := mc.cdc

	proposalTxCmd.AddCommand(client.PostCommands(
		proposalcmd.GetCmdApprove(cdc),
		proposalcmd.GetCmdUnapprove(cdc),
		proposalcmd.GetCmdTryExec(cdc),
		proposalcmd.GetCmdWithdraw(cdc),
	)...)

	return proposalTxCmd
}
