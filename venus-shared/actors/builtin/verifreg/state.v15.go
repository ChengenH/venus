// FETCHED FROM LOTUS: builtin/verifreg/state.go.template

package verifreg

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	actorstypes "github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/venus/venus-shared/actors"
	"github.com/filecoin-project/venus/venus-shared/actors/adt"

	builtin15 "github.com/filecoin-project/go-state-types/builtin"
	adt15 "github.com/filecoin-project/go-state-types/builtin/v15/util/adt"
	verifreg15 "github.com/filecoin-project/go-state-types/builtin/v15/verifreg"

	"github.com/filecoin-project/go-state-types/big"

	verifreg9 "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"
)

var _ State = (*state15)(nil)

func load15(store adt.Store, root cid.Cid) (State, error) {
	out := state15{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func make15(store adt.Store, rootKeyAddress address.Address) (State, error) {
	out := state15{store: store}

	s, err := verifreg15.ConstructState(store, rootKeyAddress)
	if err != nil {
		return nil, err
	}

	out.State = *s

	return &out, nil
}

type state15 struct {
	verifreg15.State
	store adt.Store
}

func (s *state15) RootKey() (address.Address, error) {
	return s.State.RootKey, nil
}

func (s *state15) VerifiedClientDataCap(addr address.Address) (bool, abi.StoragePower, error) {

	return false, big.Zero(), fmt.Errorf("unsupported in actors v15")

}

func (s *state15) VerifierDataCap(addr address.Address) (bool, abi.StoragePower, error) {
	return getDataCap(s.store, actors.Version15, s.verifiers, addr)
}

func (s *state15) RemoveDataCapProposalID(verifier address.Address, client address.Address) (bool, uint64, error) {
	return getRemoveDataCapProposalID(s.store, actors.Version15, s.removeDataCapProposalIDs, verifier, client)
}

func (s *state15) ForEachVerifier(cb func(addr address.Address, dcap abi.StoragePower) error) error {
	return forEachCap(s.store, actors.Version15, s.verifiers, cb)
}

func (s *state15) ForEachClient(cb func(addr address.Address, dcap abi.StoragePower) error) error {

	return fmt.Errorf("unsupported in actors v15")

}

func (s *state15) verifiedClients() (adt.Map, error) {

	return nil, fmt.Errorf("unsupported in actors v15")

}

func (s *state15) verifiers() (adt.Map, error) {
	return adt15.AsMap(s.store, s.Verifiers, builtin15.DefaultHamtBitwidth)
}

func (s *state15) removeDataCapProposalIDs() (adt.Map, error) {
	return adt15.AsMap(s.store, s.RemoveDataCapProposalIDs, builtin15.DefaultHamtBitwidth)
}

func (s *state15) GetState() interface{} {
	return &s.State
}

func (s *state15) GetAllocation(clientIdAddr address.Address, allocationId verifreg9.AllocationId) (*Allocation, bool, error) {

	alloc, ok, err := s.FindAllocation(s.store, clientIdAddr, verifreg15.AllocationId(allocationId))
	return (*Allocation)(alloc), ok, err
}

func (s *state15) GetAllocations(clientIdAddr address.Address) (map[AllocationId]Allocation, error) {

	v15Map, err := s.LoadAllocationsToMap(s.store, clientIdAddr)

	retMap := make(map[AllocationId]Allocation, len(v15Map))
	for k, v := range v15Map {
		retMap[AllocationId(k)] = Allocation(v)
	}

	return retMap, err

}

func (s *state15) GetAllAllocations() (map[AllocationId]Allocation, error) {

	v15Map, err := s.State.GetAllAllocations(s.store)

	retMap := make(map[AllocationId]Allocation, len(v15Map))
	for k, v := range v15Map {
		retMap[AllocationId(k)] = Allocation(v)
	}

	return retMap, err

}

func (s *state15) GetClaim(providerIdAddr address.Address, claimId verifreg9.ClaimId) (*Claim, bool, error) {

	claim, ok, err := s.FindClaim(s.store, providerIdAddr, verifreg15.ClaimId(claimId))
	return (*Claim)(claim), ok, err

}

func (s *state15) GetClaims(providerIdAddr address.Address) (map[ClaimId]Claim, error) {

	v15Map, err := s.LoadClaimsToMap(s.store, providerIdAddr)

	retMap := make(map[ClaimId]Claim, len(v15Map))
	for k, v := range v15Map {
		retMap[ClaimId(k)] = Claim(v)
	}

	return retMap, err

}

func (s *state15) GetAllClaims() (map[ClaimId]Claim, error) {

	v15Map, err := s.State.GetAllClaims(s.store)

	retMap := make(map[ClaimId]Claim, len(v15Map))
	for k, v := range v15Map {
		retMap[ClaimId(k)] = Claim(v)
	}

	return retMap, err

}

func (s *state15) GetClaimIdsBySector(providerIdAddr address.Address) (map[abi.SectorNumber][]ClaimId, error) {

	v15Map, err := s.LoadClaimsToMap(s.store, providerIdAddr)

	retMap := make(map[abi.SectorNumber][]ClaimId)
	for k, v := range v15Map {
		claims, ok := retMap[v.Sector]
		if !ok {
			retMap[v.Sector] = []ClaimId{ClaimId(k)}
		} else {
			retMap[v.Sector] = append(claims, ClaimId(k))
		}
	}

	return retMap, err

}

func (s *state15) ActorKey() string {
	return manifest.VerifregKey
}

func (s *state15) ActorVersion() actorstypes.Version {
	return actorstypes.Version15
}

func (s *state15) Code() cid.Cid {
	code, ok := actors.GetActorCodeID(s.ActorVersion(), s.ActorKey())
	if !ok {
		panic(fmt.Errorf("didn't find actor %v code id for actor version %d", s.ActorKey(), s.ActorVersion()))
	}

	return code
}
