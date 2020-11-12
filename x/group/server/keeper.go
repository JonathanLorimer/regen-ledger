package server

import (
	"fmt"
	"reflect"

	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/group/types"
)

const (
	// Group Table
	GroupTablePrefix        byte = 0x0
	GroupTableSeqPrefix     byte = 0x1
	GroupByAdminIndexPrefix byte = 0x2

	// Group Member Table
	GroupMemberTablePrefix         byte = 0x10
	GroupMemberByGroupIndexPrefix  byte = 0x11
	GroupMemberByMemberIndexPrefix byte = 0x12

	// Group Account Table
	GroupAccountTablePrefix        byte = 0x20
	GroupAccountTableSeqPrefix     byte = 0x21
	GroupAccountByGroupIndexPrefix byte = 0x22
	GroupAccountByAdminIndexPrefix byte = 0x23

	// ProposalBase Table
	ProposalBaseTablePrefix               byte = 0x30
	ProposalBaseTableSeqPrefix            byte = 0x31
	ProposalBaseByGroupAccountIndexPrefix byte = 0x32
	ProposalBaseByProposerIndexPrefix     byte = 0x33

	// Vote Table
	VoteTablePrefix               byte = 0x40
	VoteByProposalBaseIndexPrefix byte = 0x41
	VoteByVoterIndexPrefix        byte = 0x42
)

type Keeper struct {
	storeKey          sdk.StoreKey
	proposalModelType reflect.Type

	// Group Table
	groupSeq          orm.Sequence
	groupTable        orm.Table
	groupByAdminIndex orm.Index

	// Group Member Table
	groupMemberTable         orm.NaturalKeyTable
	groupMemberByGroupIndex  orm.UInt64Index
	groupMemberByMemberIndex orm.Index

	// Group Account Table
	groupAccountSeq          orm.Sequence
	groupAccountTable        orm.NaturalKeyTable
	groupAccountByGroupIndex orm.UInt64Index
	groupAccountByAdminIndex orm.Index

	// ProposalBase Table
	proposalTable             orm.AutoUInt64Table
	ProposalGroupAccountIndex orm.Index
	ProposalByProposerIndex   orm.Index

	// Vote Table
	voteTable               orm.NaturalKeyTable
	voteByProposalBaseIndex orm.UInt64Index
	voteByVoterIndex        orm.Index

	paramSpace paramstypes.Subspace
	router     sdk.Router
}

func NewGroupKeeper(storeKey sdk.StoreKey, paramSpace paramstypes.Subspace, router sdk.Router, proposalModel types.ProposalI) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(paramstypes.NewKeyTable().RegisterParamSet(&types.Params{}))
	}
	if storeKey == nil {
		panic("storeKey must not be nil")
	}

	if proposalModel == nil {
		panic("proposalModel must not be nil")
	}
	if router == nil {
		panic("router must not be nil")
	}
	tp := reflect.TypeOf(proposalModel)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}

	k := Keeper{storeKey: storeKey, paramSpace: paramSpace, proposalModelType: tp, router: router}

	// Group Table
	groupTableBuilder := orm.NewTableBuilder(GroupTablePrefix, storeKey, &types.GroupMetadata{}, orm.FixLengthIndexKeys(orm.EncodedSeqLength))
	k.groupSeq = orm.NewSequence(storeKey, GroupTableSeqPrefix)
	k.groupByAdminIndex = orm.NewIndex(groupTableBuilder, GroupByAdminIndexPrefix, func(val interface{}) ([]orm.RowID, error) {
		return []orm.RowID{val.(*types.GroupMetadata).Admin.Bytes()}, nil
	})
	k.groupTable = groupTableBuilder.Build()

	// Group Member Table
	groupMemberTableBuilder := orm.NewNaturalKeyTableBuilder(GroupMemberTablePrefix, storeKey, &types.GroupMember{}, orm.Max255DynamicLengthIndexKeyCodec{})
	k.groupMemberByGroupIndex = orm.NewUInt64Index(groupMemberTableBuilder, GroupMemberByGroupIndexPrefix, func(val interface{}) ([]uint64, error) {
		group := val.(*types.GroupMember).Group
		return []uint64{uint64(group)}, nil
	})
	k.groupMemberByMemberIndex = orm.NewIndex(groupMemberTableBuilder, GroupMemberByMemberIndexPrefix, func(val interface{}) ([]orm.RowID, error) {
		member := val.(*types.GroupMember).Member
		return []orm.RowID{member.Bytes()}, nil
	})
	k.groupMemberTable = groupMemberTableBuilder.Build()

	// Group Account Table
	k.groupAccountSeq = orm.NewSequence(storeKey, GroupAccountTableSeqPrefix)
	groupAccountTableBuilder := orm.NewNaturalKeyTableBuilder(GroupAccountTablePrefix, storeKey, &types.GroupAccountMetadata{}, orm.Max255DynamicLengthIndexKeyCodec{})
	k.groupAccountByGroupIndex = orm.NewUInt64Index(groupAccountTableBuilder, GroupAccountByGroupIndexPrefix, func(value interface{}) ([]uint64, error) {
		group := value.(*types.GroupAccountMetadata).Group
		return []uint64{uint64(group)}, nil
	})
	k.groupAccountByAdminIndex = orm.NewIndex(groupAccountTableBuilder, GroupAccountByAdminIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		admin := value.(*types.GroupAccountMetadata).Admin
		return []orm.RowID{admin.Bytes()}, nil
	})
	k.groupAccountTable = groupAccountTableBuilder.Build()

	// Proposal Table
	proposalTableBuilder := orm.NewAutoUInt64TableBuilder(ProposalBaseTablePrefix, ProposalBaseTableSeqPrefix, storeKey, proposalModel)
	k.ProposalGroupAccountIndex = orm.NewIndex(proposalTableBuilder, ProposalBaseByGroupAccountIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		account := value.(types.ProposalI).GetBase().GroupAccount
		return []orm.RowID{account.Bytes()}, nil

	})
	k.ProposalByProposerIndex = orm.NewIndex(proposalTableBuilder, ProposalBaseByProposerIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		proposers := value.(types.ProposalI).GetBase().Proposers
		r := make([]orm.RowID, len(proposers))
		for i := range proposers {
			r[i] = proposers[i].Bytes()
		}
		return r, nil
	})
	k.proposalTable = proposalTableBuilder.Build()

	// Vote Table
	voteTableBuilder := orm.NewNaturalKeyTableBuilder(VoteTablePrefix, storeKey, &types.Vote{}, orm.Max255DynamicLengthIndexKeyCodec{})
	k.voteByProposalBaseIndex = orm.NewUInt64Index(voteTableBuilder, VoteByProposalBaseIndexPrefix, func(value interface{}) ([]uint64, error) {
		return []uint64{uint64(value.(*types.Vote).Proposal)}, nil
	})
	k.voteByVoterIndex = orm.NewIndex(voteTableBuilder, VoteByVoterIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		return []orm.RowID{value.(*types.Vote).Voter.Bytes()}, nil
	})
	k.voteTable = voteTableBuilder.Build()

	return k
}

// MaxCommentSize returns the maximum length of a comment
func (k Keeper) MaxCommentSize(ctx sdk.Context) int {
	var result uint32
	k.paramSpace.Get(ctx, types.ParamMaxCommentLength, &result)
	return int(result)
}

func (k Keeper) CreateGroup(ctx sdk.Context, admin sdk.AccAddress, members types.Members, comment string) (types.GroupID, error) {
	if err := members.ValidateBasic(); err != nil {
		return 0, err
	}

	maxCommentSize := k.MaxCommentSize(ctx)
	if len(comment) > maxCommentSize {
		return 0, sdkerrors.Wrap(types.ErrMaxLimit, "group comment")
	}

	totalWeight := sdk.ZeroDec()
	for i := range members {
		m := members[i]
		if len(m.Comment) > maxCommentSize {
			return 0, sdkerrors.Wrap(types.ErrMaxLimit, "member comment")
		}
		totalWeight = totalWeight.Add(m.Power)
	}

	groupID := types.GroupID(k.groupSeq.NextVal(ctx))
	err := k.groupTable.Create(ctx, groupID.Bytes(), &types.GroupMetadata{
		Group:       groupID,
		Admin:       admin,
		Comment:     comment,
		Version:     1,
		TotalWeight: totalWeight,
	})
	if err != nil {
		return 0, sdkerrors.Wrap(err, "could not create group")
	}

	for i := range members {
		m := members[i]
		err := k.groupMemberTable.Create(ctx, &types.GroupMember{
			Group:   groupID,
			Member:  m.Address,
			Weight:  m.Power,
			Comment: m.Comment,
		})
		if err != nil {
			return 0, sdkerrors.Wrapf(err, "could not store member %d", i)
		}
	}
	return groupID, nil
}

func (k Keeper) GetGroup(ctx sdk.Context, id types.GroupID) (types.GroupMetadata, error) {
	var obj types.GroupMetadata
	return obj, k.groupTable.GetOne(ctx, id.Bytes(), &obj)
}

func (k Keeper) HasGroup(ctx sdk.Context, rowID orm.RowID) bool {
	return k.groupTable.Has(ctx, rowID)
}

func (k Keeper) UpdateGroup(ctx sdk.Context, g *types.GroupMetadata) error {
	g.Version++
	return k.groupTable.Save(ctx, g.Group.Bytes(), g)
}

func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var p types.Params
	k.paramSpace.GetParamSet(ctx, &p)
	return p
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// CreateGroupAccount creates and persists a `GroupAccountMetadata`
func (k Keeper) CreateGroupAccount(ctx sdk.Context, admin sdk.AccAddress, groupID types.GroupID, policy types.DecisionPolicy, comment string) (sdk.AccAddress, error) {
	maxCommentSize := k.MaxCommentSize(ctx)
	if len(comment) > maxCommentSize {
		return nil, sdkerrors.Wrap(types.ErrMaxLimit,
			"group account comment")
	}

	g, err := k.GetGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}
	if !g.Admin.Equals(admin) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not group admin")
	}
	accountAddr := types.AccountCondition(k.groupAccountSeq.NextVal(ctx)).Address()
	groupAccount, err := types.NewGroupAccountMetadata(
		accountAddr,
		groupID,
		admin,
		comment,
		1,
		policy,
	)
	if err != nil {
		return nil, err
	}

	if err := k.groupAccountTable.Create(ctx, &groupAccount); err != nil {
		return nil, sdkerrors.Wrap(err, "could not create group account")
	}
	return accountAddr, nil
}

func (k Keeper) HasGroupAccount(ctx sdk.Context, address sdk.AccAddress) bool {
	return k.groupAccountTable.Has(ctx, address.Bytes())
}

func (k Keeper) GetGroupAccount(ctx sdk.Context, accountAddress sdk.AccAddress) (types.GroupAccountMetadata, error) {
	var obj types.GroupAccountMetadata
	return obj, k.groupAccountTable.GetOne(ctx, accountAddress.Bytes(), &obj)
}

func (k Keeper) UpdateGroupAccount(ctx sdk.Context, obj *types.GroupAccountMetadata) error {
	obj.Version++
	return k.groupAccountTable.Save(ctx, obj)
}

func (k Keeper) GetGroupByGroupAccount(ctx sdk.Context, accountAddress sdk.AccAddress) (types.GroupMetadata, error) {
	obj, err := k.GetGroupAccount(ctx, accountAddress)
	if err != nil {
		return types.GroupMetadata{}, sdkerrors.Wrap(err, "load group account")
	}
	return k.GetGroup(ctx, obj.Group)
}

func (k Keeper) GetGroupMembersByGroup(ctx sdk.Context, id types.GroupID) (orm.Iterator, error) {
	return k.groupMemberByGroupIndex.Get(ctx, id.Uint64())
}

func (k Keeper) Vote(ctx sdk.Context, id types.ProposalID, voters []sdk.AccAddress, choice types.Choice, comment string) error {
	maxCommentSize := k.MaxCommentSize(ctx)
	if len(comment) > maxCommentSize {
		return sdkerrors.Wrap(types.ErrMaxLimit, "comment")
	}
	if len(voters) == 0 {
		return sdkerrors.Wrap(types.ErrEmpty, "voters")
	}

	blockTime, err := gogotypes.TimestampProto(ctx.BlockTime())
	if err != nil {
		return err
	}
	proposal, err := k.GetProposal(ctx, id)
	if err != nil {
		return err
	}
	base := proposal.GetBase()
	if base.Status != types.ProposalStatusSubmitted {
		return sdkerrors.Wrap(types.ErrInvalid, "proposal not open")
	}
	votingPeriodEnd, err := gogotypes.TimestampFromProto(&base.Timeout)
	if err != nil {
		return err
	}
	if votingPeriodEnd.Before(ctx.BlockTime()) || votingPeriodEnd.Equal(ctx.BlockTime()) {
		return sdkerrors.Wrap(types.ErrExpired, "voting period has ended already")
	}
	var accountMetadata types.GroupAccountMetadata
	if err := k.groupAccountTable.GetOne(ctx, base.GroupAccount.Bytes(), &accountMetadata); err != nil {
		return sdkerrors.Wrap(err, "load group account")
	}
	if base.GroupAccountVersion != accountMetadata.Version {
		return sdkerrors.Wrap(types.ErrModified, "group account was modified")
	}

	electorate, err := k.GetGroup(ctx, accountMetadata.Group)
	if err != nil {
		return err
	}
	if electorate.Version != base.GroupVersion {
		return sdkerrors.Wrap(types.ErrModified, "group was modified")
	}

	// count and store votes
	for _, voterAddr := range voters {
		voter := types.GroupMember{Group: electorate.Group, Member: voterAddr}
		if err := k.groupMemberTable.GetOne(ctx, voter.NaturalKey(), &voter); err != nil {
			return sdkerrors.Wrapf(err, "address: %s", voterAddr)
		}
		newVote := types.Vote{
			Proposal:    id,
			Voter:       voterAddr,
			Choice:      choice,
			Comment:     comment,
			SubmittedAt: *blockTime,
		}
		if err := base.VoteState.Add(newVote, voter.Weight); err != nil {
			return sdkerrors.Wrap(err, "add new vote")
		}

		if err := k.voteTable.Create(ctx, &newVote); err != nil {
			return sdkerrors.Wrap(err, "store vote")
		}
	}

	// run tally with new votes to close early
	if err := doTally(ctx, &base, electorate, accountMetadata); err != nil {
		return err
	}

	proposal.SetBase(base)
	return k.proposalTable.Save(ctx, id.Uint64(), proposal)
}

func doTally(ctx sdk.Context, base *types.ProposalBase, electorate types.GroupMetadata, accountMetadata types.GroupAccountMetadata) error {
	policy := accountMetadata.GetDecisionPolicy()
	submittedAt, err := gogotypes.TimestampFromProto(&base.SubmittedAt)
	if err != nil {
		return err
	}
	switch result, err := policy.Allow(base.VoteState, electorate.TotalWeight, ctx.BlockTime().Sub(submittedAt)); {
	case err != nil:
		return sdkerrors.Wrap(err, "policy execution")
	case result == types.DecisionPolicyResult{Allow: true, Final: true}:
		base.Result = types.ProposalResultAccepted
		base.Status = types.ProposalStatusClosed
	case result == types.DecisionPolicyResult{Allow: false, Final: true}:
		base.Result = types.ProposalResultRejected
		base.Status = types.ProposalStatusClosed
	}
	return nil
}

// ExecProposal can be executed n times before the timeout. It will update the proposal status and executes the msg payload.
// There are no separate transactions for the payload messages so that it is a full atomic operation that
// would either succeed or fail.
func (k Keeper) ExecProposal(ctx sdk.Context, id types.ProposalID) error {
	proposal, err := k.GetProposal(ctx, id)
	if err != nil {
		return err
	}
	// check constraints
	base := proposal.GetBase()

	if base.Status != types.ProposalStatusSubmitted && base.Status != types.ProposalStatusClosed {
		return sdkerrors.Wrapf(types.ErrInvalid, "not possible with proposal status %s", base.Status.String())
	}

	var accountMetadata types.GroupAccountMetadata
	if err := k.groupAccountTable.GetOne(ctx, base.GroupAccount.Bytes(), &accountMetadata); err != nil {
		return sdkerrors.Wrap(err, "load group account")
	}

	storeUpdates := func() error {
		proposal.SetBase(base)
		return k.proposalTable.Save(ctx, id.Uint64(), proposal)
	}

	if base.Status == types.ProposalStatusSubmitted {
		if base.GroupAccountVersion != accountMetadata.Version {
			base.Result = types.ProposalResultUndefined
			base.Status = types.ProposalStatusAborted
			return storeUpdates()
		}

		electorate, err := k.GetGroup(ctx, accountMetadata.Group)
		if err != nil {
			return sdkerrors.Wrap(err, "load group")
		}

		if electorate.Version != base.GroupVersion {
			base.Result = types.ProposalResultUndefined
			base.Status = types.ProposalStatusAborted
			return storeUpdates()
		}
		if err := doTally(ctx, &base, electorate, accountMetadata); err != nil {
			return err
		}
	}

	// execute proposal payload
	if base.Status == types.ProposalStatusClosed && base.Result == types.ProposalResultAccepted && base.ExecutorResult != types.ProposalExecutorResultSuccess {
		logger := ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
		ctx, flush := ctx.CacheContext()
		_, err := DoExecuteMsgs(ctx, k.router, accountMetadata.GroupAccount, proposal.GetMsgs())
		if err != nil {
			base.ExecutorResult = types.ProposalExecutorResultFailure
			proposalType := reflect.TypeOf(proposal).String()
			logger.Info("proposal execution failed", "cause", err, "type", proposalType, "proposalID", id)
		} else {
			base.ExecutorResult = types.ProposalExecutorResultSuccess
			flush()
		}
	}
	return storeUpdates()
}

func (k Keeper) GetProposal(ctx sdk.Context, id types.ProposalID) (types.ProposalI, error) {
	loaded := reflect.New(k.proposalModelType).Interface().(types.ProposalI)
	if _, err := k.proposalTable.GetOne(ctx, id.Uint64(), loaded); err != nil {
		return nil, sdkerrors.Wrap(err, "load proposal")
	}
	return loaded, nil
}

func (k Keeper) CreateProposal(ctx sdk.Context, accountAddress sdk.AccAddress, comment string, proposers []sdk.AccAddress, msgs []sdk.Msg) (types.ProposalID, error) {
	maxCommentSize := k.MaxCommentSize(ctx)
	if len(comment) > maxCommentSize {
		return 0, sdkerrors.Wrap(types.ErrMaxLimit, "comment")
	}

	account, err := k.GetGroupAccount(ctx, accountAddress.Bytes())

	if err != nil {
		return 0, sdkerrors.Wrap(err, "load group account")
	}

	g, err := k.GetGroup(ctx, account.Group)
	if err != nil {
		return 0, sdkerrors.Wrap(err, "get group by account")
	}

	// only members can propose
	for i := range proposers {
		if !k.groupMemberTable.Has(ctx, types.GroupMember{Group: g.Group, Member: proposers[i]}.NaturalKey()) {
			return 0, sdkerrors.Wrapf(types.ErrUnauthorized, "not in group: %s", proposers[i])
		}
	}

	if err := ensureMsgAuthZ(msgs, account.GroupAccount); err != nil {
		return 0, err
	}

	blockTime, err := gogotypes.TimestampProto(ctx.BlockTime())
	if err != nil {
		return 0, sdkerrors.Wrap(err, "block time conversion")
	}

	policy := account.GetDecisionPolicy()

	if policy == nil {
		return 0, sdkerrors.Wrap(types.ErrEmpty, "nil policy")
	}

	// prevent proposal that can not succeed
	err = policy.Validate(g)
	if err != nil {
		return 0, err
	}

	timeout := policy.GetTimeout()
	window, err := gogotypes.DurationFromProto(&timeout)
	if err != nil {
		return 0, sdkerrors.Wrap(err, "maxVotingWindow time conversion")
	}
	endTime, err := gogotypes.TimestampProto(ctx.BlockTime().Add(window))
	if err != nil {
		return 0, sdkerrors.Wrap(err, "end time conversion")
	}

	m := reflect.New(k.proposalModelType).Interface().(types.ProposalI)
	m.SetBase(types.ProposalBase{
		GroupAccount:        accountAddress,
		Comment:             comment,
		Proposers:           proposers,
		SubmittedAt:         *blockTime,
		GroupVersion:        g.Version,
		GroupAccountVersion: account.Version,
		Result:              types.ProposalResultUndefined,
		Status:              types.ProposalStatusSubmitted,
		ExecutorResult:      types.ProposalExecutorResultNotRun,
		Timeout:             *endTime,
		VoteState: types.Tally{
			YesCount:     sdk.ZeroDec(),
			NoCount:      sdk.ZeroDec(),
			AbstainCount: sdk.ZeroDec(),
			VetoCount:    sdk.ZeroDec(),
		},
	})
	if err := m.SetMsgs(msgs); err != nil {
		return 0, sdkerrors.Wrap(err, "create proposal")
	}

	id, err := k.proposalTable.Create(ctx, m)
	if err != nil {
		return 0, sdkerrors.Wrap(err, "create proposal")
	}
	return types.ProposalID(id), nil
}

func (k Keeper) GetVote(ctx sdk.Context, id types.ProposalID, voter sdk.AccAddress) (types.Vote, error) {
	var v types.Vote
	return v, k.voteTable.GetOne(ctx, types.Vote{Proposal: id, Voter: voter}.NaturalKey(), &v)
}