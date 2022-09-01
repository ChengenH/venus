// FETCHED FROM LOTUS: builtin/market/state.go.template

package market

import (
	"bytes"

	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/go-bitfield"
	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"

	"github.com/filecoin-project/venus/venus-shared/actors/adt"
	types "github.com/filecoin-project/venus/venus-shared/internal"

	market8 "github.com/filecoin-project/go-state-types/builtin/v8/market"
	adt8 "github.com/filecoin-project/go-state-types/builtin/v8/util/adt"
)

var _ State = (*state8)(nil)

func load8(store adt.Store, root cid.Cid) (State, error) {
	out := state8{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func make8(store adt.Store) (State, error) {
	out := state8{store: store}

	s, err := market8.ConstructState(store)
	if err != nil {
		return nil, err
	}

	out.State = *s

	return &out, nil
}

type state8 struct {
	market8.State
	store adt.Store
}

func (s *state8) TotalLocked() (abi.TokenAmount, error) {
	fml := types.BigAdd(s.TotalClientLockedCollateral, s.TotalProviderLockedCollateral)
	fml = types.BigAdd(fml, s.TotalClientStorageFee)
	return fml, nil
}

func (s *state8) BalancesChanged(otherState State) (bool, error) {
	otherState8, ok := otherState.(*state8)
	if !ok {
		// there's no way to compare different versions of the state, so let's
		// just say that means the state of balances has changed
		return true, nil
	}
	return !s.State.EscrowTable.Equals(otherState8.State.EscrowTable) || !s.State.LockedTable.Equals(otherState8.State.LockedTable), nil
}

func (s *state8) StatesChanged(otherState State) (bool, error) {
	otherState8, ok := otherState.(*state8)
	if !ok {
		// there's no way to compare different versions of the state, so let's
		// just say that means the state of balances has changed
		return true, nil
	}
	return !s.State.States.Equals(otherState8.State.States), nil
}

func (s *state8) States() (DealStates, error) {
	stateArray, err := adt8.AsArray(s.store, s.State.States, market8.StatesAmtBitwidth)
	if err != nil {
		return nil, err
	}
	return &dealStates8{stateArray}, nil
}

func (s *state8) ProposalsChanged(otherState State) (bool, error) {
	otherState8, ok := otherState.(*state8)
	if !ok {
		// there's no way to compare different versions of the state, so let's
		// just say that means the state of balances has changed
		return true, nil
	}
	return !s.State.Proposals.Equals(otherState8.State.Proposals), nil
}

func (s *state8) Proposals() (DealProposals, error) {
	proposalArray, err := adt8.AsArray(s.store, s.State.Proposals, market8.ProposalsAmtBitwidth)
	if err != nil {
		return nil, err
	}
	return &dealProposals8{proposalArray}, nil
}

func (s *state8) EscrowTable() (BalanceTable, error) {
	bt, err := adt8.AsBalanceTable(s.store, s.State.EscrowTable)
	if err != nil {
		return nil, err
	}
	return &balanceTable8{bt}, nil
}

func (s *state8) LockedTable() (BalanceTable, error) {
	bt, err := adt8.AsBalanceTable(s.store, s.State.LockedTable)
	if err != nil {
		return nil, err
	}
	return &balanceTable8{bt}, nil
}

func (s *state8) VerifyDealsForActivation(
	minerAddr address.Address, deals []abi.DealID, currEpoch, sectorExpiry abi.ChainEpoch,
) (weight, verifiedWeight abi.DealWeight, err error) {
	w, vw, _, err := market8.ValidateDealsForActivation(&s.State, s.store, deals, minerAddr, sectorExpiry, currEpoch)
	return w, vw, err
}

func (s *state8) NextID() (abi.DealID, error) {
	return s.State.NextID, nil
}

type balanceTable8 struct {
	*adt8.BalanceTable
}

func (bt *balanceTable8) ForEach(cb func(address.Address, abi.TokenAmount) error) error {
	asMap := (*adt8.Map)(bt.BalanceTable)
	var ta abi.TokenAmount
	return asMap.ForEach(&ta, func(key string) error {
		a, err := address.NewFromBytes([]byte(key))
		if err != nil {
			return err
		}
		return cb(a, ta)
	})
}

type dealStates8 struct {
	adt.Array
}

func (s *dealStates8) Get(dealID abi.DealID) (*DealState, bool, error) {
	var deal8 market8.DealState
	found, err := s.Array.Get(uint64(dealID), &deal8)
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}
	deal := fromV8DealState(deal8)
	return &deal, true, nil
}

func (s *dealStates8) ForEach(cb func(dealID abi.DealID, ds DealState) error) error {
	var ds8 market8.DealState
	return s.Array.ForEach(&ds8, func(idx int64) error {
		return cb(abi.DealID(idx), fromV8DealState(ds8))
	})
}

func (s *dealStates8) decode(val *cbg.Deferred) (*DealState, error) {
	var ds8 market8.DealState
	if err := ds8.UnmarshalCBOR(bytes.NewReader(val.Raw)); err != nil {
		return nil, err
	}
	ds := fromV8DealState(ds8)
	return &ds, nil
}

func (s *dealStates8) array() adt.Array {
	return s.Array
}

func fromV8DealState(v8 market8.DealState) DealState {
	return (DealState)(v8)
}

type dealProposals8 struct {
	adt.Array
}

func (s *dealProposals8) Get(dealID abi.DealID) (*DealProposal, bool, error) {
	var proposal8 market8.DealProposal
	found, err := s.Array.Get(uint64(dealID), &proposal8)
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}

	proposal, err := fromV8DealProposal(proposal8)
	if err != nil {
		return nil, true, fmt.Errorf("decoding proposal: %w", err)
	}

	return &proposal, true, nil
}

func (s *dealProposals8) ForEach(cb func(dealID abi.DealID, dp DealProposal) error) error {
	var dp8 market8.DealProposal
	return s.Array.ForEach(&dp8, func(idx int64) error {
		dp, err := fromV8DealProposal(dp8)
		if err != nil {
			return fmt.Errorf("decoding proposal: %w", err)
		}

		return cb(abi.DealID(idx), dp)
	})
}

func (s *dealProposals8) decode(val *cbg.Deferred) (*DealProposal, error) {
	var dp8 market8.DealProposal
	if err := dp8.UnmarshalCBOR(bytes.NewReader(val.Raw)); err != nil {
		return nil, err
	}

	dp, err := fromV8DealProposal(dp8)
	if err != nil {
		return nil, err
	}

	return &dp, nil
}

func (s *dealProposals8) array() adt.Array {
	return s.Array
}

func fromV8DealProposal(v8 market8.DealProposal) (DealProposal, error) {

	label := v8.Label

	return DealProposal{
		PieceCID:     v8.PieceCID,
		PieceSize:    v8.PieceSize,
		VerifiedDeal: v8.VerifiedDeal,
		Client:       v8.Client,
		Provider:     v8.Provider,

		Label: label,

		StartEpoch:           v8.StartEpoch,
		EndEpoch:             v8.EndEpoch,
		StoragePricePerEpoch: v8.StoragePricePerEpoch,

		ProviderCollateral: v8.ProviderCollateral,
		ClientCollateral:   v8.ClientCollateral,
	}, nil
}

func (s *state8) GetState() interface{} {
	return &s.State
}

var _ PublishStorageDealsReturn = (*publishStorageDealsReturn8)(nil)

func decodePublishStorageDealsReturn8(b []byte) (PublishStorageDealsReturn, error) {
	var retval market8.PublishStorageDealsReturn
	if err := retval.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, fmt.Errorf("failed to unmarshal PublishStorageDealsReturn: %w", err)
	}

	return &publishStorageDealsReturn8{retval}, nil
}

type publishStorageDealsReturn8 struct {
	market8.PublishStorageDealsReturn
}

func (r *publishStorageDealsReturn8) IsDealValid(index uint64) (bool, int, error) {

	set, err := r.ValidDeals.IsSet(index)
	if err != nil || !set {
		return false, -1, err
	}
	maskBf, err := bitfield.NewFromIter(&rlepluslazy.RunSliceIterator{
		Runs: []rlepluslazy.Run{rlepluslazy.Run{Val: true, Len: index}}})
	if err != nil {
		return false, -1, err
	}
	before, err := bitfield.IntersectBitField(maskBf, r.ValidDeals)
	if err != nil {
		return false, -1, err
	}
	outIdx, err := before.Count()
	if err != nil {
		return false, -1, err
	}
	return set, int(outIdx), nil

}

func (r *publishStorageDealsReturn8) DealIDs() ([]abi.DealID, error) {
	return r.IDs, nil
}