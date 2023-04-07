package metatx

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	forwarder_wrapper "github.com/smartcontractkit/chainlink/core/gethwrappers/generated/forwarder"
	"github.com/smartcontractkit/chainlink/core/utils"
)

const (
	MetaERC20Name    = "MetaERC20"
	MetaERC20Version = "v1"
	TypeHash         = "ForwardRequest(address from,address target,uint256 nonce,bytes data,uint256 validUntilTime)"
)

func SignMetaTransfer(
	forwarder forwarder_wrapper.Forwarder,
	ownerPrivateKey *ecdsa.PrivateKey,
	owner, sourceTokenAddress, to common.Address,
	calldataHash [32]byte,
	deadline *big.Int,
) (signature []byte, domainSeparatorHash [32]byte, typeHash [32]byte, nonce *big.Int, err error) {
	nonce, err = forwarder.GetNonce(nil, owner)
	if err != nil {
		return nil, [32]byte{}, [32]byte{}, nil, errors.Wrapf(err, "failed to get nonce of %s", owner.Hex())
	}

	domainSeparator, err := forwarder.GetDomainSeparator(nil, MetaERC20Name, MetaERC20Version)
	if err != nil {
		return nil, [32]byte{}, [32]byte{}, nil, errors.Wrap(err, "failed to get domain separator from contract")
	}
	domainSeparatorHashRaw := crypto.Keccak256(domainSeparator)
	copy(domainSeparatorHash[:], domainSeparatorHashRaw[:])

	typeHashRaw := crypto.Keccak256([]byte(TypeHash))
	copy(typeHash[:], typeHashRaw[:])
	message := []byte{0x19, 0x01} // \x19\x01
	message = append(message, domainSeparatorHashRaw[:]...)

	encodedCall, err := utils.ABIEncode(
		`
	[
			{"name": "typeHash","type":"bytes32"},
			{"name": "from","type":"address"},
			{"name": "target", "type": "address"},
			{"name": "nonce", "type": "uint256"},
			{"name": "data", "type": "bytes32"},
			{"name": "validUntilTime", "type": "uint256"}
	]
	`, typeHash, owner, sourceTokenAddress, nonce, calldataHash, deadline,
	)

	if err != nil {
		return nil, [32]byte{}, [32]byte{}, nil, errors.Wrap(err, "failed to abi encode")
	}

	encodedHash := crypto.Keccak256(encodedCall)

	message = append(message, encodedHash...)
	messageDigest := crypto.Keccak256(message)
	rawSignature, err := crypto.Sign(messageDigest, ownerPrivateKey)
	if err != nil {
		return nil, [32]byte{}, [32]byte{}, nil, errors.Wrap(err, "failed to sign message digest")
	}

	// decompose signature into v, r and s
	// the returned byte array is [R || S || V]
	if len(rawSignature) != 65 {
		panic("rawSignature should be 65 bytes long")
	}
	var (
		v uint8
		r [32]byte
		s [32]byte
	)
	rSlice := rawSignature[:32] // first 32 bytes is R
	copy(r[:], rSlice[:])
	sSlice := rawSignature[32:64] // second 32 bytes is S
	copy(s[:], sSlice[:])
	v = uint8(rawSignature[64])
	if v == 1 {
		v = 28
	} else {
		v = 27
	}

	signature = append(signature, r[:]...)
	signature = append(signature, s[:]...)
	signature = append(signature, v)

	return
}
