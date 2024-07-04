package ccipevm

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_onramp"
)

// bytes32 internal constant LEAF_DOMAIN_SEPARATOR = 0x0000000000000000000000000000000000000000000000000000000000000000;
var leafDomainSeparator = [32]byte{}

// bytes32 internal constant ANY_2_EVM_MESSAGE_HASH = keccak256("Any2EVMMessageHashV1");
var ANY_2_EVM_MESSAGE_HASH = utils.Keccak256Fixed([]byte("Any2EVMMessageHashV1"))

// MessageHasherV1 implements the MessageHasher interface.
// Compatible with:
// - "EVM2EVMMultiOnRamp 1.6.0-dev"
type MessageHasherV1 struct {
	// ABIs and types for encoding the message data similar to on-chain implementation:
	// https://github.com/smartcontractkit/ccip/blob/54ee4f13143d3e414627b6a0b9f71d5dfade76c5/contracts/src/v0.8/ccip/libraries/Internal.sol#L135
	bytesArrayType     abi.Type
	tokensAbi          abi.ABI
	fixedSizeValuesAbi abi.ABI
	packedValuesAbi    abi.ABI
	metaDataHashABI    abi.ABI

	// TODO: move these to CCIPMsg instead?
	destChainSelector cciptypes.ChainSelector
	onrampAddress     []byte
}

func NewMessageHasherV1(
	onrampAddress []byte,
	destChainSelector cciptypes.ChainSelector,
) *MessageHasherV1 {
	bytesArray, err := abi.NewType("bytes[]", "bytes[]", nil)
	if err != nil {
		panic(fmt.Sprintf("failed to create bytes[] type: %v", err))
	}

	return &MessageHasherV1{
		bytesArrayType:     bytesArray,
		tokensAbi:          mustParseInputsAbi(tokensABI),
		fixedSizeValuesAbi: mustParseInputsAbi(fixedSizeValuesABI),
		packedValuesAbi:    mustParseInputsAbi(finalHashInputABI),
		metaDataHashABI:    mustParseInputsAbi(metaDataHashABI),

		destChainSelector: destChainSelector,
		onrampAddress:     onrampAddress,
	}
}

// Hash implements the MessageHasher interface.
// It constructs all of the inputs to the final keccak256 hash in Internal._hash(Any2EVMRampMessage).
// The main structure of the hash is as follows:
/*
	keccak256(
		leafDomainSeparator,
		keccak256(any_2_evm_message_hash, header.sourceChainSelector, header.destinationChainSelector, onRamp),
		keccak256(fixedSizeMessageFields),
		keccak256(messageData),
		keccak256(encodedTokenAmounts),
		keccak256(encodedSourceTokenData),
	)
*/
func (h *MessageHasherV1) Hash(_ context.Context, msg cciptypes.CCIPMsg) (cciptypes.Bytes32, error) {
	tokenAmounts := make([]evm_2_evm_multi_onramp.ClientEVMTokenAmount, len(msg.TokenAmounts))
	for i, ta := range msg.TokenAmounts {
		tokenAmounts[i] = evm_2_evm_multi_onramp.ClientEVMTokenAmount{
			Token:  common.HexToAddress(string(ta.Token)),
			Amount: ta.Amount,
		}
	}
	encodedTokens, err := h.abiEncode(h.tokensAbi, tokenAmounts)
	if err != nil {
		return [32]byte{}, fmt.Errorf("abi encode token amounts: %w", err)
	}

	encodedSourceTokenData, err := abi.Arguments{abi.Argument{Type: h.bytesArrayType}}.
		PackValues([]interface{}{msg.SourceTokenData})
	if err != nil {
		return [32]byte{}, fmt.Errorf("pack source token data: %w", err)
	}

	metaDataHashInput, err := h.abiEncode(
		h.metaDataHashABI,
		ANY_2_EVM_MESSAGE_HASH,
		uint64(msg.SourceChain),
		uint64(h.destChainSelector),
		h.onrampAddress,
	)
	if err != nil {
		return [32]byte{}, fmt.Errorf("abi encode metadata hash input: %w", err)
	}
	metaDataHash := utils.Keccak256Fixed(metaDataHashInput)

	var msgID [32]byte
	decoded, err := hex.DecodeString(msg.ID)
	if err != nil {
		return [32]byte{}, fmt.Errorf("decode message ID: %w", err)
	}
	if len(decoded) != 32 {
		return [32]byte{}, fmt.Errorf("message ID must be 32 bytes")
	}
	copy(msgID[:], decoded)

	// NOTE: msg.Sender is not necessarily an EVM address since this is Any2EVM.
	// Accordingly, sender is defined as "bytes" in the onchain message definition
	// rather than "address".
	// However, its not clear how best to translate from Sender being a string representation
	// to bytes. For now, we assume that the string is hex encoded, but ideally Sender would
	// just be a byte array in the CCIPMsg struct that represents a sender encoded in the
	// source chain family encoding scheme.
	decodedSender, err := hex.DecodeString(
		strings.TrimPrefix(string(msg.Sender), "0x"),
	)
	if err != nil {
		return [32]byte{}, fmt.Errorf("decode sender '%s': %w", msg.Sender, err)
	}
	fixedSizeFieldsEncoded, err := h.abiEncode(
		h.fixedSizeValuesAbi,
		msgID,
		decodedSender,
		common.HexToAddress(string(msg.Receiver)),
		uint64(msg.SeqNum),
		msg.ChainFeeLimit.Int,
		msg.Nonce,
	)
	if err != nil {
		return [32]byte{}, fmt.Errorf("abi encode fixed size values: %w", err)
	}
	fixedSizeFieldsHash := utils.Keccak256Fixed(fixedSizeFieldsEncoded)

	packedValues, err := h.abiEncode(
		h.packedValuesAbi,
		leafDomainSeparator,
		metaDataHash,
		fixedSizeFieldsHash,
		utils.Keccak256Fixed(msg.Data),
		utils.Keccak256Fixed(encodedTokens),
		utils.Keccak256Fixed(encodedSourceTokenData),
	)
	if err != nil {
		return [32]byte{}, fmt.Errorf("abi encode packed values: %w", err)
	}

	return utils.Keccak256Fixed(packedValues), nil
}

func (h *MessageHasherV1) abiEncode(theAbi abi.ABI, values ...interface{}) ([]byte, error) {
	res, err := theAbi.Pack("method", values...)
	if err != nil {
		return nil, err
	}
	return res[4:], nil
}

func mustParseInputsAbi(s string) abi.ABI {
	inDef := fmt.Sprintf(`[{ "name" : "method", "type": "function", "inputs": %s}]`, s)
	inAbi, err := abi.JSON(strings.NewReader(inDef))
	if err != nil {
		panic(fmt.Errorf("failed to create %s ABI: %v", s, err))
	}
	return inAbi
}

// Interface compliance check
var _ cciptypes.MessageHasher = (*MessageHasherV1)(nil)
