package addressparser

import (
	"github.com/ethereum/go-ethereum/common"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2/types"
)

func AccountToAddress(accounts []ocrtypes2.Account) (addresses []common.Address) {
	for _, signer := range accounts {
		if len(signer) != 20 {
			panic("public key is not 20 bytes")
		}
		addresses = append(addresses, common.BytesToAddress([]byte(signer)))
	}
	return addresses
}

func OnchainPublicKeyToAddress(publicKeys []ocrtypes2.OnchainPublicKey) (addresses []common.Address) {
	for _, signer := range publicKeys {
		if len(signer) != 20 {
			panic("public key is not 20 bytes")
		}
		addresses = append(addresses, common.BytesToAddress(signer))
	}
	return addresses
}
