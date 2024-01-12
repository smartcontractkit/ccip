package ocr3impls

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
)

var _ types.OffchainConfigDigester = (*MultichainConfigDigester)(nil)

type MultichainConfigDigester struct {
	ChainID         uint64
	ContractAddress common.Address
}

func (d MultichainConfigDigester) ConfigDigest(cc types.ContractConfig) (types.ConfigDigest, error) {
	// signers := []common.Address{}
	// for i, signer := range cc.Signers {
	// 	if len(signer) != 20 {
	// 		return types.ConfigDigest{}, fmt.Errorf("%v-th evm signer should be a 20 byte address, but got %x", i, signer)
	// 	}
	// 	a := common.BytesToAddress(signer)
	// 	signers = append(signers, a)
	// }
	// transmitterGroups := [][]common.Address{}
	// for i, transmitter := range cc.Transmitters {
	// 	var group []common.Address
	// 	splits := SplitTransmitters(string(transmitter))
	// 	for _, split := range splits {
	// 		if !strings.HasPrefix(split, "0x") || len(split) != 42 || !common.IsHexAddress(split) {
	// 			return types.ConfigDigest{}, fmt.Errorf("%v-th evm transmitter should be a 42 character Ethereum address string, but got '%v'", i, transmitter)
	// 		}
	// 		a := common.HexToAddress(string(transmitter))
	// 		group = append(group, a)
	// 	}
	// 	transmitterGroups = append(transmitterGroups, group)
	// }

	// TODO: figure out how to calculate the digest without brute-force calculating
	// every possible combination of transmitters
	// assuming N chains and one transmitter per chain, there are N! possible
	// combinations of transmitters which is too many to brute force for even a small N.
	return cc.ConfigDigest, nil
}

func (d MultichainConfigDigester) ConfigDigestPrefix() (types.ConfigDigestPrefix, error) {
	return types.ConfigDigestPrefixEVM, nil
}

// func makeConfigDigestArgs() abi.Arguments {
// 	abi, err := abi.JSON(strings.NewReader(
// 		no_op_ocr3.NoOpOCR3MetaData.ABI))
// 	if err != nil {
// 		// assertion
// 		panic(fmt.Sprintf("could not parse ocr3 ABI: %s", err.Error()))
// 	}
// 	return abi.Methods["exposedConfigDigestFromConfigData"].Inputs
// }

// var configDigestArgs = makeConfigDigestArgs()

// func configDigest(
// 	chainID uint64,
// 	contractAddress common.Address,
// 	configCount uint64,
// 	oracles []common.Address,
// 	transmitters []common.Address,
// 	f uint8,
// 	onchainConfig []byte,
// 	offchainConfigVersion uint64,
// 	offchainConfig []byte,
// ) types.ConfigDigest {
// 	chainIDBig := new(big.Int)
// 	chainIDBig.SetUint64(chainID)
// 	msg, err := configDigestArgs.Pack(
// 		chainIDBig,
// 		contractAddress,
// 		configCount,
// 		oracles,
// 		transmitters,
// 		f,
// 		onchainConfig,
// 		offchainConfigVersion,
// 		offchainConfig,
// 	)
// 	if err != nil {
// 		// assertion
// 		panic(err)
// 	}
// 	rawHash := crypto.Keccak256(msg)
// 	configDigest := types.ConfigDigest{}
// 	if n := copy(configDigest[:], rawHash); n != len(configDigest) {
// 		// assertion
// 		panic("copy too little data")
// 	}
// 	if types.ConfigDigestPrefixEVM != 1 {
// 		// assertion
// 		panic("wrong ConfigDigestPrefix")
// 	}
// 	configDigest[0] = 0
// 	configDigest[1] = 1
// 	return configDigest
// }
