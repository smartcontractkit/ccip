package ocrcommon

import (
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2/types"
)

func AccountToAddress(accounts []ocrtypes2.Account) (addresses []common.Address, err error) {
	for _, signer := range accounts {
		bytes, err := hex.DecodeString(strings.TrimPrefix(string(signer), "0x"))
		if err != nil {
			return []common.Address{}, errors.Wrap(err, "error decoding hex string")
		}
		if len(bytes) != 20 {
			return []common.Address{}, errors.Errorf("address is not 20 bytes %s", signer)
		}
		addresses = append(addresses, common.BytesToAddress(bytes))
	}
	return addresses, nil
}

func OnchainPublicKeyToAddress(publicKeys []ocrtypes2.OnchainPublicKey) (addresses []common.Address, err error) {
	for _, signer := range publicKeys {
		if len(signer) != 20 {
			return []common.Address{}, errors.Errorf("address is not 20 bytes %s", signer)
		}
		addresses = append(addresses, common.BytesToAddress(signer))
	}
	return addresses, nil
}
