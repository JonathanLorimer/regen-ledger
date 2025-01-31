package core

import (
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"gotest.tools/v3/assert"
	"strings"
	"testing"
)

func TestQuery_BatchesByClass(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// make a class
	assert.NilError(t, s.stateStore.ClassInfoTable().Insert(s.ctx, &api.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   "foo",
		CreditType: "C",
	}))
	// make some batches under it
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectId:  1,
		BatchDenom: "C01-20200101-20200102-001",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectId:  1,
		BatchDenom: "C01-20190203-20200102-002",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectId:  1,
		BatchDenom: "C01-20500404-20900102-003",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))

	// Classes that SHOULD NOT show up from a query for "C01"
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectId:  1,
		BatchDenom: "C011-20500404-20900102-003",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectId:  1,
		BatchDenom: "BIO1-20500404-20900102-003",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))

	res, err := s.k.BatchesByClass(s.ctx, &core.QueryBatchesByClassRequest{
		ClassId:    "C01",
		Pagination: nil,
	})
	assert.NilError(t, err)
	assert.Equal(t, 3, len(res.Batches))
	for _, batch := range res.Batches {
		assert.Check(t, strings.Contains(batch.BatchDenom, "C01"))
	}
}
