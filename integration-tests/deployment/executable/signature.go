package executable

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/ccip-owner-contracts/gethwrappers"
)

type Signature struct {
	R common.Hash
	S common.Hash
	V uint8
}

func NewSignatureFromBytes(sig []byte) Signature {
	return Signature{
		R: common.BytesToHash(sig[:32]),
		S: common.BytesToHash(sig[32:64]),
		V: uint8(sig[64]),
	}
}

func (s Signature) ToGethSignature() gethwrappers.ManyChainMultiSigSignature {
	return gethwrappers.ManyChainMultiSigSignature{
		R: [32]byte(s.R.Bytes()),
		S: [32]byte(s.S.Bytes()),
		V: s.V,
	}
}

func (s Signature) ToBytes() []byte {
	return append(s.R.Bytes(), append(s.S.Bytes(), []byte{byte(s.V)}...)...)
}

func (s Signature) Recover(hash common.Hash) (common.Address, error) {
	return recoverAddressFromSignature(hash, s.ToBytes())
}
