package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgUpdateClassAdmin{}

func (m MsgUpdateClassAdmin) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassAdmin) Type() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassAdmin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

func (m *MsgUpdateClassAdmin) ValidateBasic() error {
	if m.Admin == m.NewAdmin {
		return sdkerrors.ErrInvalidAddress.Wrap("new admin should be a different address from the signer")
	}

	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	if _, err := sdk.AccAddressFromBech32(m.NewAdmin); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	if err := ValidateClassID(m.ClassId); err != nil {
		return err
	}

	return nil
}

func (m *MsgUpdateClassAdmin) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}