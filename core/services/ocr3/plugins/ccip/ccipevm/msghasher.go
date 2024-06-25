package ccipevm

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"golang.org/x/crypto/sha3"
)

var (
	LeafDomainSeparator = [1]byte{0x00}
)

type MessageHasher struct {
	metaDataHash [32]byte
}

func NewMessageHasher(metaDataHash [32]byte) *MessageHasher {
	return &MessageHasher{
		metaDataHash: metaDataHash,
	}
}

func (h *MessageHasher) Hash(_ context.Context, msg cciptypes.CCIPMsg) (cciptypes.Bytes32, error) {
	encodedTokens, err := h.abiEncode(`[{"components": [{"name":"token","type":"address"},{"name":"amount","type":"uint256"}], "type":"tuple[]"}]`, msg.TokenAmounts)
	if err != nil {
		return [32]byte{}, fmt.Errorf("abi encode token amounts: %w", err)
	}

	bytesArray, err := abi.NewType("bytes[]", "bytes[]", nil)
	if err != nil {
		return [32]byte{}, fmt.Errorf("new abi type bytes[]: %w", err)
	}

	encodedSourceTokenData, err := abi.Arguments{abi.Argument{Type: bytesArray}}.PackValues([]interface{}{msg.SourceTokenData})
	if err != nil {
		return [32]byte{}, fmt.Errorf("pack source token data: %w", err)
	}

	packedFixedSizeValues, err := h.abiEncode(
		`[
{"name": "sender", "type":"address"},
{"name": "receiver", "type":"address"},
{"name": "sequenceNumber", "type":"uint64"},
{"name": "gasLimit", "type":"uint256"},
{"name": "strict", "type":"bool"},
{"name": "nonce", "type":"uint64"},
{"name": "feeToken","type": "address"},
{"name": "feeTokenAmount","type": "uint256"}
]`,
		common.HexToAddress(string(msg.Sender)),
		common.HexToAddress(string(msg.Receiver)),
		uint64(msg.SeqNum),
		msg.ChainFeeLimit.Int,
		msg.Strict,
		msg.Nonce,
		common.HexToAddress(string(msg.FeeToken)),
		msg.FeeTokenAmount.Int,
	)
	if err != nil {
		return [32]byte{}, fmt.Errorf("abi encode fixed size values: %w", err)
	}
	fixedSizeValuesHash := h.keccak256Fixed(packedFixedSizeValues)

	packedValues, err := h.abiEncode(
		`[
{"name": "leafDomainSeparator","type":"bytes1"},
{"name": "metadataHash", "type":"bytes32"},
{"name": "fixedSizeValuesHash", "type":"bytes32"},
{"name": "dataHash", "type":"bytes32"},
{"name": "tokenAmountsHash", "type":"bytes32"},
{"name": "sourceTokenDataHash", "type":"bytes32"}
]`,
		LeafDomainSeparator,
		h.metaDataHash,
		fixedSizeValuesHash,
		h.keccak256Fixed(msg.Data),
		h.keccak256Fixed(encodedTokens),
		h.keccak256Fixed(encodedSourceTokenData),
	)
	if err != nil {
		return [32]byte{}, fmt.Errorf("abi encode packed values: %w", err)
	}

	return h.keccak256Fixed(packedValues), nil
}

func (h *MessageHasher) abiEncode(abiStr string, values ...interface{}) ([]byte, error) {
	inDef := fmt.Sprintf(`[{ "name" : "method", "type": "function", "inputs": %s}]`, abiStr)
	inAbi, err := abi.JSON(strings.NewReader(inDef))
	if err != nil {
		return nil, err
	}
	res, err := inAbi.Pack("method", values...)
	if err != nil {
		return nil, err
	}
	return res[4:], nil
}

func (h *MessageHasher) keccak256Fixed(in []byte) [32]byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(in)
	var hs [32]byte
	copy(hs[:], hash.Sum(nil))
	return hs
}

// Interface compliance check
var _ cciptypes.MessageHasher = (*MessageHasher)(nil)
