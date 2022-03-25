package core

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"golang.org/x/mod/sumdb/note"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/types/structvalid"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var supportedOriginTxType = map[string]struct{}{
	"Polygon":  struct{}{},
	"Ethereum": struct{}{},
	"Verra":    struct{}{},
}

var _ legacytx.LegacyMsg = &MsgMintBatch{}

// Route implements the LegacyMsg interface.
func (m MsgMintBatch) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgMintBatch) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgMintBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgMintBatch) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Issuer); err != nil {
		return sdkerrors.Wrap(err, "issuer")
	}
	errs := structvalid.AssertStrMaxLen("note", m.Note, 255, nil)
	errs = structvalid.AssertStrMaxLen("id", m.OriginTx.Id, 255, nil)
	if err := structvalid.ErrsToError(errs); err != nil {
		return err
	}
	for _, bi := range m.Issuance {
		if err := validateBatchIssuance(bi); err != nil {
			return err
		}
	}

	return nil

}
