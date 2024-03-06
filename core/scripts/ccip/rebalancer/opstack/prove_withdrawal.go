package opstack

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip/rebalancer/multienv"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	optimism_l2_output_oracle "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/optimism_l2_output_oracle"
	optimism_l2_to_l1_message_passer "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/optimism_l2_to_l1_message_passer"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/optimism_portal"
)

var (
	MessagePassedTopic = common.HexToHash("0x02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054")
)

func ProveWithdrawal(
	env multienv.Env,
	l1ChainID,
	l2ChainID uint64,
	l1BridgeAdapterAddress,
	l2OutputOracleAddress,
	optimismPortalAddress common.Address,
	l2TxHash common.Hash,
) {
	l2Client, ok := env.Clients[l2ChainID]
	if !ok {
		panic(fmt.Sprintf("No client found for chain %d, map: %+v", l2ChainID, env.Clients))
	}
	l1Client, ok := env.Clients[l1ChainID]
	if !ok {
		panic(fmt.Sprintf("No client found for chain %d, map: %+v", l1ChainID, env.Clients))
	}
	proveMessage(env, l1Client, l2Client, l2TxHash, optimismPortalAddress, l2OutputOracleAddress)
}

type outputRootProof struct {
	Version                  [32]byte
	StateRoot                [32]byte
	MessagePasserStorageRoot [32]byte
	LatestBlockHash          [32]byte
}

type bedrockMessageProof struct {
	WithdrawalProof [][]byte
	L2OutputIndex   *big.Int
	OutputRootProof outputRootProof
}

func proveMessage(
	env multienv.Env,
	l1Client, l2Client *ethclient.Client,
	l2TxHash common.Hash,
	optimismPortalAddress common.Address,
	l2OutputOracleAddress common.Address,
) {
	lowLevelMsg, messageBedrockOutput := getBedrockMessageProof(
		l1Client, l2Client, l2TxHash, optimismPortalAddress, l2OutputOracleAddress)

	optimismPortal, err := optimism_portal.NewOptimismPortal(optimismPortalAddress, l1Client)
	helpers.PanicErr(err)

	l1ChainID, err := l1Client.ChainID(context.Background())
	helpers.PanicErr(err)

	fmt.Println("Calling proveWithdrawalTransaction on OptimismPortal, nonce:", lowLevelMsg.MessageNonce, "\n",
		"sender:", lowLevelMsg.Sender.String(), "\n",
		"target:", lowLevelMsg.Target.String(), "\n",
		"value:", lowLevelMsg.Value.String(), "\n",
		"gasLimit:", lowLevelMsg.MinGasLimit.String(), "\n",
		"data:", hexutil.Encode(lowLevelMsg.Message), "\n",
		"l2OutputIndex:", messageBedrockOutput.L2OutputIndex, "\n",
		"outputRootProof version:", hexutil.Encode(messageBedrockOutput.OutputRootProof.Version[:]), "\n",
		"outputRootProof stateRoot:", hexutil.Encode(messageBedrockOutput.OutputRootProof.StateRoot[:]), "\n",
		"outputRootProof messagePasserStorageRoot:", hexutil.Encode(messageBedrockOutput.OutputRootProof.MessagePasserStorageRoot[:]), "\n",
		"outputRootProof latestBlockHash:", hexutil.Encode(messageBedrockOutput.OutputRootProof.LatestBlockHash[:]), "\n",
		"withdrawalProof:", formatWithdrawalProof(messageBedrockOutput.WithdrawalProof))

	tx, err := optimismPortal.ProveWithdrawalTransaction(env.Transactors[l1ChainID.Uint64()],
		optimism_portal.TypesWithdrawalTransaction{
			Nonce:    lowLevelMsg.MessageNonce,
			Sender:   lowLevelMsg.Sender,
			Target:   lowLevelMsg.Target,
			Value:    lowLevelMsg.Value,
			GasLimit: lowLevelMsg.MinGasLimit,
			Data:     lowLevelMsg.Message,
		},
		messageBedrockOutput.L2OutputIndex,
		optimism_portal.TypesOutputRootProof{
			Version:                  messageBedrockOutput.OutputRootProof.Version,
			StateRoot:                messageBedrockOutput.OutputRootProof.StateRoot,
			MessagePasserStorageRoot: messageBedrockOutput.OutputRootProof.MessagePasserStorageRoot,
			LatestBlockhash:          messageBedrockOutput.OutputRootProof.LatestBlockHash,
		},
		messageBedrockOutput.WithdrawalProof,
	)
	helpers.PanicErr(err)
	helpers.ConfirmTXMined(context.Background(), l1Client, tx, int64(l1ChainID.Uint64()), "ProveWithdrawal")
}

func formatWithdrawalProof(proof [][]byte) string {
	var builder strings.Builder
	builder.WriteString("{")
	for i, p := range proof {
		builder.WriteString(hexutil.Encode(p))
		if i < len(proof)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("}")
	return builder.String()
}

func getBedrockMessageProof(
	l1Client, l2Client *ethclient.Client,
	l2TxHash common.Hash,
	optimismPortalAddress,
	l2OutputOracleAddress common.Address,
) (
	lowLevelMessage,
	bedrockMessageProof,
) {
	receipt, err := l2Client.TransactionReceipt(context.Background(), l2TxHash)
	helpers.PanicErr(err)

	messagePassedLog := getMessagePassedLog(receipt.Logs)
	if messagePassedLog == nil {
		panic(fmt.Sprintf("No message passed log found in receipt %s", receipt.TxHash.String()))
	}

	messagePassed := parseMessagePassedLog(messagePassedLog)

	lowLevelMsg := toLowLevelMessage(messagePassed)
	llmHash := hashLowLevelMessage(lowLevelMsg)
	messageSlot := hashMessageHash(llmHash)

	fmt.Println("Low level message, message nonce:", lowLevelMsg.MessageNonce.String(), "\n",
		"sender:", lowLevelMsg.Sender.String(), "\n",
		"target:", lowLevelMsg.Target.String(), "\n",
		"value:", lowLevelMsg.Value.String(), "\n",
		"min gas limit:", lowLevelMsg.MinGasLimit.String(), "\n",
		"message:", hexutil.Encode(lowLevelMsg.Message), "\n",
		"message hash:", hexutil.Encode(llmHash[:]), "\n",
		"message slot:", hexutil.Encode(messageSlot[:]))

	fmt.Println("Low level message hash:", hexutil.Encode(llmHash[:]), "\n",
		"Message slot:", hexutil.Encode(messageSlot[:]))

	messageBedrockOutput := getMessageBedrockOutput(
		optimismPortalAddress,
		l2OutputOracleAddress,
		receipt.BlockNumber,
		l1Client,
	)

	fmt.Println("got bedrock output, l1 timestamp:", messageBedrockOutput.L1Timestamp.String(), "\n",
		"l2 block number:", messageBedrockOutput.L2BlockNumber.String(), "\n",
		"l2 output index:", messageBedrockOutput.L2OutputIndex.String(), "\n",
		"output root:", hexutil.Encode(messageBedrockOutput.OutputRoot[:]))

	stateTrieProof := makeStateTrieProof(
		l2Client.Client(),
		messageBedrockOutput.L2BlockNumber,
		messagePassed.Raw.Address,
		messageSlot,
	)

	header, err := l2Client.HeaderByNumber(
		context.Background(), messageBedrockOutput.L2BlockNumber)
	helpers.PanicErr(err)

	return lowLevelMsg, bedrockMessageProof{
		OutputRootProof: outputRootProof{
			Version:                  [32]byte{},
			StateRoot:                header.Root,
			MessagePasserStorageRoot: stateTrieProof.StorageRoot,
			LatestBlockHash:          header.Hash(),
		},
		WithdrawalProof: toProofBytes(stateTrieProof.StorageProof),
		L2OutputIndex:   messageBedrockOutput.L2OutputIndex,
	}
}

func toProofBytes(proof []hexutil.Bytes) [][]byte {
	proofBytes := make([][]byte, len(proof))
	for i, p := range proof {
		proofBytes[i] = p
	}
	return proofBytes
}

// See https://eips.ethereum.org/EIPS/eip-1186#specification
type ProofNode struct {
	Key   hexutil.Bytes
	Value hexutil.Big
	Proof []hexutil.Bytes
}

type getProofResponse struct {
	AccountProof []hexutil.Bytes `json:"accountProof"`
	Balance      *hexutil.Big    `json:"balance"`
	CodeHash     common.Hash     `json:"codeHash"`
	Nonce        *hexutil.Big    `json:"nonce"`
	StorageHash  common.Hash     `json:"storageHash"`
	StorageProof []ProofNode     `json:"storageProof"`
}

type stateTrieProof struct {
	AccountProof []hexutil.Bytes
	StorageProof []hexutil.Bytes
	StorageValue *big.Int
	StorageRoot  [32]byte
}

func makeStateTrieProof(
	l2Client *rpc.Client,
	l2BlockNumber *big.Int,
	address common.Address,
	slot [32]byte,
) stateTrieProof {
	fmt.Println("calling eth_getProof with args, address:", address.String(), "slot:", hexutil.Encode(slot[:]), "block:", l2BlockNumber.String())

	var resp getProofResponse
	err := l2Client.Call(&resp, "eth_getProof",
		address, []string{hexutil.Encode(slot[:])}, hexutil.EncodeBig(l2BlockNumber))
	helpers.PanicErr(err)

	// TODO
	// resp.StorageProof[0].Proof = maybeAddProofNode(
	// 	crypto.Keccak256Hash(slot[:]), resp.StorageProof[0].Proof)

	return stateTrieProof{
		AccountProof: resp.AccountProof,
		StorageProof: resp.StorageProof[0].Proof,
		StorageValue: resp.StorageProof[0].Value.ToInt(),
		StorageRoot:  resp.StorageHash,
	}
}

// func maybeAddProofNode(
// 	key [32]byte,
// 	proof []hexutil.Bytes,
// ) []hexutil.Bytes {
// 	modifiedProof := make([]hexutil.Bytes, len(proof))
// 	copy(modifiedProof, proof)
// 	finalProofEl := modifiedProof[len(modifiedProof)-1]
// 	finalProofElDecoded := rlp.DecodeBytes(finalProofEl)
// }

func hashMessageHash(h [32]byte) [32]byte {
	var zeroHash [32]byte
	encoded, err := utils.ABIEncode(`[{"type": "bytes32"}, {"type": "bytes32"}]`, h, zeroHash)
	helpers.PanicErr(err)
	return crypto.Keccak256Hash(encoded)
}

func hashLowLevelMessage(llm lowLevelMessage) [32]byte {
	encoded, err := utils.ABIEncode(
		`[{"type": "uint256"}, {"type": "address"}, {"type": "address"}, {"type": "uint256"}, {"type": "uint256"}, {"type": "bytes"}]`,
		llm.MessageNonce,
		llm.Sender,
		llm.Target,
		llm.Value,
		llm.MinGasLimit,
		llm.Message,
	)
	helpers.PanicErr(err)
	return crypto.Keccak256Hash(encoded)
}

func toLowLevelMessage(
	messagePassed *optimism_l2_to_l1_message_passer.OptimismL2ToL1MessagePasserMessagePassed) lowLevelMessage {
	return lowLevelMessage{
		MessageNonce: messagePassed.Nonce,
		Sender:       messagePassed.Sender,
		Target:       messagePassed.Target,
		Value:        messagePassed.Value,
		Message:      messagePassed.Data,
		MinGasLimit:  messagePassed.GasLimit,
	}
}

func getMessageBedrockOutput(
	optimismPortalAddress,
	l2OutputOracleAddress common.Address,
	l2BlockNumber *big.Int,
	l1Client *ethclient.Client,
) bedrockOutput {
	if getFPAC(optimismPortalAddress, l1Client) {
		panic("fpac support not implemented")
	}

	// Try to find the output index that corresponds to the block number attached to the message.
	// We'll explicitly handle "cannot get output" errors as a null return value, but anything else
	// needs to get thrown. Might need to revisit this in the future to be a little more robust
	// when connected to RPCs that don't return nice error messages.
	l2OutputOracle, err := optimism_l2_output_oracle.NewOptimismL2OutputOracle(l2OutputOracleAddress, l1Client)
	helpers.PanicErr(err)

	l2OutputIndex, err := l2OutputOracle.GetL2OutputIndexAfter(nil, l2BlockNumber)
	helpers.PanicErr(err)

	// Now pull the proposal out given the output index. Should always work as long as the above
	// codepath completed successfully.
	proposal, err := l2OutputOracle.GetL2Output(nil, l2OutputIndex)
	helpers.PanicErr(err)

	return bedrockOutput{
		OutputRoot:    proposal.OutputRoot,
		L1Timestamp:   proposal.Timestamp,
		L2BlockNumber: proposal.L2BlockNumber,
		L2OutputIndex: l2OutputIndex,
	}
}

func getMessagePassedLog(logs []*gethtypes.Log) *gethtypes.Log {
	for _, lg := range logs {
		if lg.Topics[0] == MessagePassedTopic {
			return lg
		}
	}
	return nil
}

func getFPAC(portalAddress common.Address, l1Client *ethclient.Client) bool {
	opPortal, err := optimism_portal.NewOptimismPortal(portalAddress, l1Client)
	helpers.PanicErr(err)
	semVer, err := opPortal.Version(nil)
	helpers.PanicErr(err)

	version := semver.MustParse(semVer)
	return version.GreaterThan(semver.MustParse("3.0.0")) || version.Equal(semver.MustParse("3.0.0"))
}

type bedrockOutput struct {
	OutputRoot    [32]byte
	L1Timestamp   *big.Int
	L2BlockNumber *big.Int
	L2OutputIndex *big.Int
}

func parseMessagePassedLog(log *gethtypes.Log) *optimism_l2_to_l1_message_passer.OptimismL2ToL1MessagePasserMessagePassed {
	cdm, err := optimism_l2_to_l1_message_passer.NewOptimismL2ToL1MessagePasser(common.HexToAddress("0x0"), nil)
	helpers.PanicErr(err)
	messagePassed, err := cdm.ParseMessagePassed(*log)
	helpers.PanicErr(err)
	return messagePassed
}

type lowLevelMessage struct {
	Sender       common.Address
	Target       common.Address
	Message      []byte
	MessageNonce *big.Int
	Value        *big.Int
	MinGasLimit  *big.Int
}
