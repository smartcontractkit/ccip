package config

import (
	chainselectors "github.com/smartcontractkit/chain-selectors"
)

var testSelectors = map[uint64]uint64{
	90000001: 909606746561742123,
	90000002: 5548718428018410741,
	90000003: 789068866484373046,
	90000004: 5721565186521185178,
	90000005: 964127714438319834,
	90000006: 8966794841936584464,
	90000007: 8412806778050735057,
	90000008: 4066443121807923198,
	90000009: 6747736380229414777,
	90000010: 8694984074292254623,
	90000011: 328334718812072308,
	90000012: 7715160997071429212,
	90000013: 3574539439524578558,
	90000014: 4543928599863227519,
	90000015: 6443235356619661032,
}

func AllSimulatedChainIds() []uint64 {
	var chainIds []uint64
	for k := range testSelectors {
		chainIds = append(chainIds, k)
	}
	return chainIds
}

func ChainIdFromSelector(chainSelectorId uint64) (uint64, error) {
	for k, v := range testSelectors {
		if v == chainSelectorId {
			return k, nil
		}
	}
	return chainselectors.ChainIdFromSelector(chainSelectorId)
}

func SelectorFromChainId(chainId uint64) (uint64, error) {
	if chainSelectorId, exist := testSelectors[chainId]; exist {
		return chainSelectorId, nil
	}
	return chainselectors.SelectorFromChainId(chainId)
}
