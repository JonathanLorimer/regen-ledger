package core

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// CreateBatch creates a new batch of credits.
// Credits in the batch must not have more decimal places than the credit type's specified precision.
func (k Keeper) MintBatch(ctx context.Context, req *ecocredit.MsgMintBatch) (*ecocredit.MsgMintBatchResponse, error) {
	panic("not implemented")
}
