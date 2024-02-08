package ccipcalc

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/types/cciptypes"
)

func EvmAddrsToGeneric(evmAddrs ...common.Address) []cciptypes.Address {
	res := make([]cciptypes.Address, 0, len(evmAddrs))
	for _, addr := range evmAddrs {
		res = append(res, cciptypes.Address(addr.String()))
	}
	return res
}

func GenericAddrsToEvm(genericAddrs ...cciptypes.Address) ([]common.Address, error) {
	evmAddrs := make([]common.Address, 0, len(genericAddrs))
	for _, addr := range genericAddrs {
		if !common.IsHexAddress(string(addr)) {
			return nil, fmt.Errorf("%s not an evm address", addr)
		}
		evmAddrs = append(evmAddrs, common.HexToAddress(string(addr)))
	}
	return evmAddrs, nil
}
