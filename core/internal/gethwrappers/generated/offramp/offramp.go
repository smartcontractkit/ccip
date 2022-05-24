// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package offramp

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

type CCIPExecutionReport struct {
	Messages       []CCIPMessage
	Proofs         [][32]byte
	ProofFlagsBits *big.Int
}

type CCIPMessage struct {
	SourceChainId  *big.Int
	SequenceNumber uint64
	Sender         common.Address
	Payload        CCIPMessagePayload
}

type CCIPMessagePayload struct {
	Tokens             []common.Address
	Amounts            []*big.Int
	DestinationChainId *big.Int
	Receiver           common.Address
	Executor           common.Address
	Data               []byte
}

type CCIPRelayReport struct {
	MerkleRoot        [32]byte
	MinSequenceNumber uint64
	MaxSequenceNumber uint64
}

type OffRampInterfaceOffRampConfig struct {
	ExecutionFeeJuels     uint64
	ExecutionDelaySeconds uint64
	MaxDataSize           uint64
	MaxTokensLength       uint64
}

var OffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"contractAggregatorV2V3Interface[]\",\"name\":\"feeds\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionDelayError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"InvalidExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceFeedConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"MerkleProofError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RelayReportError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastMaxSequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newMinSequenceNumber\",\"type\":\"uint64\"}],\"name\":\"SequenceError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenFeedMistmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMistmatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnexpectedPayloadData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"CrossChainMessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minSequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceNumber\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structCCIP.RelayReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"ReportAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"addFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagsBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"needFee\",\"type\":\"bool\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastReport\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minSequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceNumber\",\"type\":\"uint64\"}],\"internalType\":\"structCCIP.RelayReport\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRampConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagsBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"removeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setOffRampConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200602238038062006022833981016040819052620000349162000747565b6000805460ff191681556001908790869082908990889088903390819081620000a45760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000de57620000de81620003ed565b5050506001600160a01b0382161580620000f6575080155b156200011557604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03939093169290921790915560035580518251146200015b5760405162d8548360e71b815260040160405180910390fd5b8151620001709060059060208501906200049e565b5060005b8251811015620002545760008282815181106200019557620001956200081d565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001df57620001df6200081d565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff19166001179055806200024b8162000833565b91505062000174565b50505080518251146200027a5760405163ee9d106b60e01b815260040160405180910390fd5b81516200028f9060089060208501906200049e565b5060005b82518110156200035c576000828281518110620002b457620002b46200081d565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060076000868581518110620002fe57620002fe6200081d565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b02929091169190911790555080620003538162000833565b91505062000293565b505050151560805260a09790975250505060c0929092525050805160138054602084015160408501516060909501516001600160401b03908116600160c01b026001600160c01b03968216600160801b02969096166001600160801b0392821668010000000000000000026001600160801b031990941691909516179190911716919091179190911790556200085b565b336001600160a01b03821603620004475760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016200009b565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620004f6579160200282015b82811115620004f657825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620004bf565b506200050492915062000508565b5090565b5b8082111562000504576000815560010162000509565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156200056057620005606200051f565b604052919050565b60006001600160401b038211156200058457620005846200051f565b5060051b60200190565b6001600160a01b0381168114620005a457600080fd5b50565b600082601f830112620005b957600080fd5b81516020620005d2620005cc8362000568565b62000535565b82815260059290921b84018101918181019086841115620005f257600080fd5b8286015b848110156200061a5780516200060c816200058e565b8352918301918301620005f6565b509695505050505050565b600082601f8301126200063757600080fd5b815160206200064a620005cc8362000568565b82815260059290921b840181019181810190868411156200066a57600080fd5b8286015b848110156200061a57805162000684816200058e565b83529183019183016200066e565b80516200069f816200058e565b919050565b80516001600160401b03811681146200069f57600080fd5b600060808284031215620006cf57600080fd5b604051608081016001600160401b0381118282101715620006f457620006f46200051f565b6040529050806200070583620006a4565b81526200071560208401620006a4565b60208201526200072860408401620006a4565b60408201526200073b60608401620006a4565b60608201525092915050565b600080600080600080600080610160898b0312156200076557600080fd5b885160208a015160408b015191995097506001600160401b03808211156200078c57600080fd5b6200079a8c838d01620005a7565b975060608b0151915080821115620007b157600080fd5b620007bf8c838d01620005a7565b965060808b0151915080821115620007d657600080fd5b50620007e58b828c0162000625565b945050620007f660a08a0162000692565b925060c089015191506200080e8a60e08b01620006bc565b90509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b6000600182016200085457634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c051615790620008926000396000610511015260008181610452015261387d015260006123e001526157906000f3fe608060405234801561001057600080fd5b50600436106102925760003560e01c806381ff704811610160578063b034909c116100d8578063c0d786551161008c578063eb511dd411610071578063eb511dd41461073a578063eefa7a3e1461074d578063f2fde38b146107dd57600080fd5b8063c0d7865514610714578063e3d0e7121461072757600080fd5b8063b1dc65a4116100bd578063b1dc65a4146106b5578063b6608c3b146106c8578063bbe4f6db146106db57600080fd5b8063b034909c1461068f578063b0f479a11461069757600080fd5b80638da5cb5b1161012f578063a7206cd611610114578063a7206cd614610571578063a8ebd0f414610591578063afcb95d71461066f57600080fd5b80638da5cb5b1461053b578063a655e9fb1461055e57600080fd5b806381ff7048146104d45780638456cb591461050457806385e1f4d01461050c57806389c065681461053357600080fd5b806359e96b5b1161020e57806374be2150116101c257806380d9a1b7116101a757806380d9a1b71461048a57806381411834146104b757806381be8fa4146104cc57600080fd5b806374be21501461044d57806379ba50971461048257600080fd5b80635c975abb116101f35780635c975abb1461041c5780636642031d14610427578063744b92e21461043a57600080fd5b806359e96b5b146103d05780635b16ebb7146103e357600080fd5b80632b898c2511610265578063461c551b1161024a578063461c551b14610387578063567c814b1461039a5780635853c627146103bd57600080fd5b80632b898c251461036c5780633f4ba83a1461037f57600080fd5b8063108ee5fc1461029757806316b8e731146102ac578063181f5a771461030f5780632222dd421461034e575b600080fd5b6102aa6102a53660046146a1565b6107f0565b005b6102e56102ba3660046146a1565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b604080518082018252600d81527f4f666652616d7020302e302e3100000000000000000000000000000000000000602082015290516103069190614738565b60025473ffffffffffffffffffffffffffffffffffffffff166102e5565b6102aa61037a36600461474b565b6108cc565b6102aa610c9c565b6102aa610395366004614784565b610cae565b6103ad6103a836600461479c565b610d00565b6040519015158152602001610306565b6102aa6103cb36600461474b565b610e47565b6102aa6103de3660046147c0565b61105f565b6103ad6103f13660046146a1565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff166103ad565b6102aa610435366004614cb5565b6110dd565b6102aa61044836600461474b565b611a59565b6104747f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610306565b6102aa611e4e565b6103ad610498366004614cfc565b67ffffffffffffffff1660009081526010602052604090205460ff1690565b6104bf611f70565b6040516103069190614d6a565b6104bf611fdf565b600b546009546040805163ffffffff80851682526401000000009094049093166020840152820152606001610306565b6102aa61204c565b6104747f000000000000000000000000000000000000000000000000000000000000000081565b6104bf61205c565b600054610100900473ffffffffffffffffffffffffffffffffffffffff166102e5565b61047461056c366004614d7d565b6120c9565b61047461057f36600461479c565b6000908152600f602052604090205490565b61062b604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260135467ffffffffffffffff808216835268010000000000000000820481166020840152700100000000000000000000000000000000820481169383019390935278010000000000000000000000000000000000000000000000009004909116606082015290565b6040516103069190815167ffffffffffffffff9081168252602080840151821690830152604080840151821690830152606092830151169181019190915260800190565b604080516001815260006020820181905291810191909152606001610306565b600354610474565b60145473ffffffffffffffffffffffffffffffffffffffff166102e5565b6102aa6106c3366004614dfe565b6122a3565b6102aa6106d636600461479c565b61294c565b6102e56106e93660046146a1565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6102aa6107223660046146a1565b6129cc565b6102aa610735366004614ef4565b612a47565b6102aa61074836600461474b565b61342c565b6107a960408051606081018252600080825260208201819052918101919091525060408051606081018252601154815260125467ffffffffffffffff808216602084015268010000000000000000909104169181019190915290565b604080518251815260208084015167ffffffffffffffff908116918301919091529282015190921690820152606001610306565b6102aa6107eb3660046146a1565b61366c565b6107f8613680565b73ffffffffffffffffffffffffffffffffffffffff8116610845576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b6108d4613680565b6008546000819003610912576040517f2e70248b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906109ad576040517f3917193900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610a16576040517f6c17b98700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006008610a25600185614ff0565b81548110610a3557610a35615007565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600882602001516bffffffffffffffffffffffff1681548110610a8757610a87615007565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166008610ab6600186614ff0565b81548110610ac657610ac6615007565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600883602001516bffffffffffffffffffffffff1681548110610b3457610b34615007565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526007909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556008805480610bd657610bd6615036565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600783526040808520949094558351908152908716918101919091527fa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d91015b60405180910390a15050505050565b610ca4613680565b610cac613706565b565b610cb6613680565b806013610cc38282615065565b9050507fe4cd88b1b5e20a0b843af3207ba74d3f84af4acff45830469490ac9c6ab8a74581604051610cf59190615176565b60405180910390a150565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa158015610d70573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d9491906151e0565b158015610e415750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610e0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e3091906151fd565b60200151610e3e9084614ff0565b11155b92915050565b610e4f613680565b73ffffffffffffffffffffffffffffffffffffffff82161580610e86575073ffffffffffffffffffffffffffffffffffffffff8116155b15610ebd576040517fee9d106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015610f59576040517f965ffb7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600880546bffffffffffffffffffffffff908116602080870191825288861660008181526007835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821790558351908152908101919091527f037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb60130591015b60405180910390a1505050565b611067613680565b61108873ffffffffffffffffffffffffffffffffffffffff841683836137e7565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa890606001611052565b60005460ff161561114f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111e091906151e0565b15611216576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa158015611286573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112aa91906151fd565b90506003548160200151426112bf9190614ff0565b11156112f7576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60145473ffffffffffffffffffffffffffffffffffffffff16611346576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611351846120c9565b6000818152600f602052604081205491925081900361139f576040517f851bdf5300000000000000000000000000000000000000000000000000000000815260048101839052602401611146565b60135442906113c49068010000000000000000900467ffffffffffffffff1683615239565b106113fb576040517f15c33ba200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b855151811015611a515760008660000151828151811061142057611420615007565b6020908102919091018101518082015167ffffffffffffffff166000908152601090925260409091205490915060ff16156114995760208101516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611146565b60608101516080015173ffffffffffffffffffffffffffffffffffffffff16158015906114e4575060608101516080015173ffffffffffffffffffffffffffffffffffffffff163314155b1561152d5760208101516040517f0525dc3b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611146565b61153681613879565b61153f816139ef565b60208082015167ffffffffffffffff16600090815260109091526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055851561178b576000808260600151600001516000815181106115aa576115aa615007565b6020026020010151905060006115e58273ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff8116611634576040517f83135fec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166350d25bcd6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561167f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116a39190615251565b6013546116ba919067ffffffffffffffff1661526a565b9250821561178757828460600151602001516000815181106116de576116de615007565b602002602001018181516116f29190614ff0565b9052506116fe82613a9e565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081523360048201526024810185905273ffffffffffffffffffffffffffffffffffffffff919091169063ea6192a290604401600060405180830381600087803b15801561176e57600080fd5b505af1158015611782573d6000803e3d6000fd5b505050505b5050505b60005b6060820151515181101561188f576117c682606001516000015182815181106117b9576117b9615007565b6020026020010151613a9e565b73ffffffffffffffffffffffffffffffffffffffff1663ea6192a2836060015160600151846060015160200151848151811061180457611804615007565b60200260200101516040518363ffffffff1660e01b815260040161184a92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600060405180830381600087803b15801561186457600080fd5b505af1158015611878573d6000803e3d6000fd5b505050508080611887906152a7565b91505061178e565b50606080820151015173ffffffffffffffffffffffffffffffffffffffff163b156119b05760145460608083015101516040517ff3c9dd1a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9092169163f3c9dd1a91611912918590600401615401565b600060405180830381600087803b15801561192c57600080fd5b505af192505050801561193d575060015b6119ab573d80801561196b576040519150601f19603f3d011682016040523d82523d6000602084013e611970565b606091505b508160200151816040517fa1dc8185000000000000000000000000000000000000000000000000000000008152600401611146929190615430565b611a03565b606081015160a001515115611a035760208101516040517fc945cae000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611146565b806020015167ffffffffffffffff167f88a2c71ae86800edd209d83430f348ce70539bd40085d045340ff569f058a2a060405160405180910390a25080611a49816152a7565b9150506113fe565b505050505050565b611a61613680565b6005546000819003611a9f576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290611b3a576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614611ba3576040517fd428911900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006005611bb2600185614ff0565b81548110611bc257611bc2615007565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff1681548110611c1457611c14615007565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166005611c43600186614ff0565b81548110611c5357611c53615007565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110611cc157611cc1615007565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480611d6357611d63615036565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101610c8d565b60015473ffffffffffffffffffffffffffffffffffffffff163314611ecf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401611146565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600e805480602002602001604051908101604052809291908181526020018280548015611fd557602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611faa575b5050505050905090565b60606008805480602002602001604051908101604052809291908181526020018280548015611fd55760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611faa575050505050905090565b612054613680565b610cac613b1a565b60606005805480602002602001604051908101604052809291908181526020018280548015611fd55760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611faa575050505050905090565b8051516020820151516000919082906001906120e59084615239565b6120ef9190614ff0565b905061010081111561210057600080fd5b60008167ffffffffffffffff81111561211b5761211b614801565b604051908082528060200260200182016040528015612144578160200160208202803683370190505b5090506000806000805b8581101561224c576040890151811c600190811614612226816121935760208b01518051600186019590811061218657612186615007565b60200260200101516121da565b8886106121b157865160018601958891811061218657612186615007565b8a51805160018801976121da929181106121cd576121cd615007565b6020026020010151613bda565b8987106122055787516001870196899181106121f8576121f8615007565b6020026020010151613c4a565b8b5180516001890198612221929181106121cd576121cd615007565b613c4a565b86838151811061223857612238615007565b60209081029190910101525060010161214e565b50841561227d5783600186038151811061226857612268615007565b60200260200101519650505050505050919050565b61229788600001516000815181106121cd576121cd615007565b98975050505050505050565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916122f991849163ffffffff851691908e908e9081908401838280828437600092019190915250613d0892505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600954808252600a5460ff808216602085015261010090910416928201929092529083146123ce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401611146565b6123dc8b8b8b8b8b8b61410d565b60007f0000000000000000000000000000000000000000000000000000000000000000156124395760028260200151836040015161241a9190615453565b6124249190615478565b61242f906001615453565b60ff16905061244f565b6020820151612449906001615453565b60ff1690505b8881146124b8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401611146565b888714612521576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401611146565b336000908152600c602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115612564576125646154c1565b6002811115612575576125756154c1565b9052509050600281602001516002811115612592576125926154c1565b1480156125d95750600e816000015160ff16815481106125b4576125b4615007565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b61263f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401611146565b5050505050600088886040516126569291906154f0565b60405190819003812061266d918c90602001615500565b60405160208183030381529060405280519060200120905061268d614660565b604080518082019091526000808252602082015260005b8881101561292a5760006001858884602081106126c3576126c3615007565b6126d091901a601b615453565b8d8d868181106126e2576126e2615007565b905060200201358c8c878181106126fb576126fb615007565b9050602002013560405160008152602001604052604051612738949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa15801561275a573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600c602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156127da576127da6154c1565b60028111156127eb576127eb6154c1565b9052509250600183602001516002811115612808576128086154c1565b1461286f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401611146565b8251849060ff16601f811061288657612886615007565b6020020151156128f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401611146565b600184846000015160ff16601f811061290d5761290d615007565b911515602090920201525080612922816152a7565b9150506126a4565b5050505063ffffffff81106129415761294161551c565b505050505050505050565b612954613680565b8060000361298e576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c25191016108c0565b6129d4613680565b601480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490602001610cf5565b855185518560ff16601f831115612aba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401611146565b60008111612b24576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401611146565b818314612bb2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401611146565b612bbd81600361526a565b8311612c25576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401611146565b612c2d613680565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600d5415612e2057600d54600090612c8590600190614ff0565b90506000600d8281548110612c9c57612c9c615007565b6000918252602082200154600e805473ffffffffffffffffffffffffffffffffffffffff90921693509084908110612cd657612cd6615007565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600c909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600d80549192509080612d5657612d56615036565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600e805480612dbf57612dbf615036565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550612c6b915050565b60005b815151811015613287576000600c600084600001518481518110612e4957612e49615007565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612e9357612e936154c1565b14612efa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401611146565b6040805180820190915260ff821681526001602082015282518051600c9160009185908110612f2b57612f2b615007565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612fcc57612fcc6154c1565b021790555060009150612fdc9050565b600c600084602001518481518110612ff657612ff6615007565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115613040576130406154c1565b146130a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401611146565b6040805180820190915260ff82168152602081016002815250600c6000846020015184815181106130da576130da615007565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561317b5761317b6154c1565b02179055505082518051600d92508390811061319957613199615007565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600e91908390811061321557613215615007565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790558061327f816152a7565b915050612e23565b506040810151600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600b80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff43811682029283178555908304811693600193909260009261331992869290821691161761554b565b92506101000a81548163ffffffff021916908363ffffffff1602179055506133784630600b60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a001516141c4565b600981905582518051600a805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600b5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598613417988b98919763ffffffff909216969095919491939192615573565b60405180910390a15050505050505050505050565b613434613680565b73ffffffffffffffffffffffffffffffffffffffff8216158061346b575073ffffffffffffffffffffffffffffffffffffffff8116155b156134a2576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152901561353e576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c9101611052565b613674613680565b61367d8161426f565b50565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610cac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401611146565b60005460ff16613772576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401611146565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905261387490849061436a565b505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146138d95780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401611146565b60135460608201515151780100000000000000000000000000000000000000000000000090910467ffffffffffffffff1610806139225750606081015160208101515190515114155b15613959576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601354606082015160a001515170010000000000000000000000000000000090910467ffffffffffffffff16101561367d57601354606082015160a00151516040517f8693378900000000000000000000000000000000000000000000000000000000815270010000000000000000000000000000000090920467ffffffffffffffff1660048301526024820152604401611146565b606080820151015173ffffffffffffffffffffffffffffffffffffffff16301480613a465750606080820151015173ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff165b1561367d5760608082015101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401611146565b73ffffffffffffffffffffffffffffffffffffffff8181166000908152600460205260409020541680613b15576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401611146565b919050565b60005460ff1615613b87576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611146565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586137bd3390565b60008060f81b82604051602001613bf19190615609565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052613c2d929160200161561c565b604051602081830303815290604052805190602001209050919050565b6000818310613cac57604080517f01000000000000000000000000000000000000000000000000000000000000006020808301919091526021820185905260418083018790528351808403909101815260619092019092528051910120613d01565b604080517f010000000000000000000000000000000000000000000000000000000000000060208083019190915260218201869052604180830186905283518084039091018152606190920190925280519101205b9392505050565b60005460ff1615613d75576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611146565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613de2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e0691906151e0565b15613e3c576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa158015613eac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ed091906151fd565b9050600354816020015142613ee59190614ff0565b1115613f1d576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082806020019051810190613f339190615664565b9050806040015167ffffffffffffffff16816020015167ffffffffffffffff161115613f8b576040517f67a3824c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160608101825260115480825260125467ffffffffffffffff80821660208501526801000000000000000090910416928201929092529015614048576040810151613fda9060016156af565b67ffffffffffffffff16826020015167ffffffffffffffff161461404857604080820151602084015191517f8e8c0add00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff918216600482015291166024820152604401611146565b81516000908152600f602090815260409182902042905583516011819055818501805160128054868901805167ffffffffffffffff90811668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090931694811694909417919091179091558551938452915181169383019390935251909116918101919091527f6a2424b54e9b18f0c3d440b6003882f5ffa34272205608ad9451c921c0d99a359060600160405180910390a1505050505050565b600061411a82602061526a565b61412585602061526a565b61413188610144615239565b61413b9190615239565b6141459190615239565b614150906000615239565b90503681146141bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401611146565b50505050505050565b6000808a8a8a8a8a8a8a8a8a6040516020016141e8999897969594939291906156d2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036142ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401611146565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006143cc826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166144769092919063ffffffff16565b80519091501561387457808060200190518101906143ea91906151e0565b613874576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401611146565b6060614485848460008561448d565b949350505050565b60608247101561451f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401611146565b843b614587576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611146565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516145b09190615767565b60006040518083038185875af1925050503d80600081146145ed576040519150601f19603f3d011682016040523d82523d6000602084013e6145f2565b606091505b509150915061460282828661460d565b979650505050505050565b6060831561461c575081613d01565b82511561462c5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111469190614738565b604051806103e00160405280601f906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff8116811461367d57600080fd5b6000602082840312156146b357600080fd5b8135613d018161467f565b60005b838110156146d95781810151838201526020016146c1565b838111156146e8576000848401525b50505050565b600081518084526147068160208601602086016146be565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613d0160208301846146ee565b6000806040838503121561475e57600080fd5b82356147698161467f565b915060208301356147798161467f565b809150509250929050565b60006080828403121561479657600080fd5b50919050565b6000602082840312156147ae57600080fd5b5035919050565b8035613b158161467f565b6000806000606084860312156147d557600080fd5b83356147e08161467f565b925060208401356147f08161467f565b929592945050506040919091013590565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561485357614853614801565b60405290565b6040516080810167ffffffffffffffff8111828210171561485357614853614801565b60405160c0810167ffffffffffffffff8111828210171561485357614853614801565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156148e6576148e6614801565b604052919050565b600067ffffffffffffffff82111561490857614908614801565b5060051b60200190565b67ffffffffffffffff8116811461367d57600080fd5b8035613b1581614912565b600082601f83011261494457600080fd5b81356020614959614954836148ee565b61489f565b82815260059290921b8401810191818101908684111561497857600080fd5b8286015b8481101561499c57803561498f8161467f565b835291830191830161497c565b509695505050505050565b600082601f8301126149b857600080fd5b813560206149c8614954836148ee565b82815260059290921b840181019181810190868411156149e757600080fd5b8286015b8481101561499c57803583529183019183016149eb565b600082601f830112614a1357600080fd5b813567ffffffffffffffff811115614a2d57614a2d614801565b614a5e60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161489f565b818152846020838601011115614a7357600080fd5b816020850160208301376000918101602001919091529392505050565b600060608284031215614aa257600080fd5b614aaa614830565b9050813567ffffffffffffffff80821115614ac457600080fd5b818401915084601f830112614ad857600080fd5b81356020614ae8614954836148ee565b82815260059290921b84018101918181019088841115614b0757600080fd5b8286015b84811015614c6c57803586811115614b2257600080fd5b87017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06080828d0382011215614b5757600080fd5b614b5f614859565b8683013581526040830135614b7381614912565b818801526060830135614b858161467f565b6040820152608083013589811115614b9c57600080fd5b929092019160c0838e0383011215614bb357600080fd5b614bbb61487c565b91508683013589811115614bce57600080fd5b614bdc8e8983870101614933565b835250604083013589811115614bf157600080fd5b614bff8e89838701016149a7565b888401525060608301356040830152614c1a608084016147b5565b6060830152614c2b60a084016147b5565b608083015260c083013589811115614c4257600080fd5b614c508e8983870101614a02565b60a0840152506060810191909152845250918301918301614b0b565b5086525085810135935082841115614c8357600080fd5b614c8f878588016149a7565b81860152505050506040820135604082015292915050565b801515811461367d57600080fd5b60008060408385031215614cc857600080fd5b823567ffffffffffffffff811115614cdf57600080fd5b614ceb85828601614a90565b925050602083013561477981614ca7565b600060208284031215614d0e57600080fd5b8135613d0181614912565b600081518084526020808501945080840160005b83811015614d5f57815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101614d2d565b509495945050505050565b602081526000613d016020830184614d19565b600060208284031215614d8f57600080fd5b813567ffffffffffffffff811115614da657600080fd5b61448584828501614a90565b60008083601f840112614dc457600080fd5b50813567ffffffffffffffff811115614ddc57600080fd5b6020830191508360208260051b8501011115614df757600080fd5b9250929050565b60008060008060008060008060e0898b031215614e1a57600080fd5b606089018a811115614e2b57600080fd5b8998503567ffffffffffffffff80821115614e4557600080fd5b818b0191508b601f830112614e5957600080fd5b813581811115614e6857600080fd5b8c6020828501011115614e7a57600080fd5b6020830199508098505060808b0135915080821115614e9857600080fd5b614ea48c838d01614db2565b909750955060a08b0135915080821115614ebd57600080fd5b50614eca8b828c01614db2565b999c989b50969995989497949560c00135949350505050565b803560ff81168114613b1557600080fd5b60008060008060008060c08789031215614f0d57600080fd5b863567ffffffffffffffff80821115614f2557600080fd5b614f318a838b01614933565b97506020890135915080821115614f4757600080fd5b614f538a838b01614933565b9650614f6160408a01614ee3565b95506060890135915080821115614f7757600080fd5b614f838a838b01614a02565b9450614f9160808a01614928565b935060a0890135915080821115614fa757600080fd5b50614fb489828a01614a02565b9150509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561500257615002614fc1565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b813561507081614912565b67ffffffffffffffff811690508154817fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000821617835560208401356150b481614912565b6fffffffffffffffff00000000000000008160401b16905080837fffffffffffffffffffffffffffffffff00000000000000000000000000000000841617178455604085013561510381614912565b77ffffffffffffffff000000000000000000000000000000008160801b1690507fffffffffffffffff00000000000000000000000000000000000000000000000081858286161784171786556060870135935061515f84614912565b808460c01b16858417831717865550505050505050565b60808101823561518581614912565b67ffffffffffffffff90811683526020840135906151a282614912565b90811660208401526040840135906151b982614912565b90811660408401526060840135906151d082614912565b8082166060850152505092915050565b6000602082840312156151f257600080fd5b8151613d0181614ca7565b60006060828403121561520f57600080fd5b615217614830565b8251815260208301516020820152604083015160408201528091505092915050565b6000821982111561524c5761524c614fc1565b500190565b60006020828403121561526357600080fd5b5051919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156152a2576152a2614fc1565b500290565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036152d8576152d8614fc1565b5060010190565b805182526000602067ffffffffffffffff81840151168185015273ffffffffffffffffffffffffffffffffffffffff6040840151166040850152606083015160806060860152805160c0608087015261533c610140870182614d19565b838301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80888303810160a08a0152815180845291860193506000929091908601905b8084101561539f578451825293860193600193909301929086019061537f565b50604085015160c08a0152606085015173ffffffffffffffffffffffffffffffffffffffff90811660e08b01526080860151166101008a015260a085015189820383016101208b015295506153f481876146ee565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8316815260406020820152600061448560408301846152df565b67ffffffffffffffff8316815260406020820152600061448560408301846146ee565b600060ff821660ff84168060ff0382111561547057615470614fc1565b019392505050565b600060ff8316806154b2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8060ff84160491505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600063ffffffff80831681851680830382111561556a5761556a614fc1565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526155a38184018a614d19565b905082810360808401526155b78189614d19565b905060ff871660a084015282810360c08401526155d481876146ee565b905067ffffffffffffffff851660e08401528281036101008401526155f981856146ee565b9c9b505050505050505050505050565b602081526000613d0160208301846152df565b7fff0000000000000000000000000000000000000000000000000000000000000083168152600082516156568160018501602087016146be565b919091016001019392505050565b60006060828403121561567657600080fd5b61567e614830565b82518152602083015161569081614912565b602082015260408301516156a381614912565b60408201529392505050565b600067ffffffffffffffff80831681851680830382111561556a5761556a614fc1565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526157198285018b614d19565b9150838203608085015261572d828a614d19565b915060ff881660a085015283820360c085015261574a82886146ee565b90861660e085015283810361010085015290506155f981856146ee565b600082516157798184602087016146be565b919091019291505056fea164736f6c634300080d000a",
}

var OffRampABI = OffRampMetaData.ABI

var OffRampBin = OffRampMetaData.Bin

func DeployOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, sourceTokens []common.Address, pools []common.Address, feeds []common.Address, afn common.Address, maxTimeWithoutAFNSignal *big.Int, config OffRampInterfaceOffRampConfig) (common.Address, *types.Transaction, *OffRamp, error) {
	parsed, err := OffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OffRampBin), backend, sourceChainId, chainId, sourceTokens, pools, feeds, afn, maxTimeWithoutAFNSignal, config)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OffRamp{OffRampCaller: OffRampCaller{contract: contract}, OffRampTransactor: OffRampTransactor{contract: contract}, OffRampFilterer: OffRampFilterer{contract: contract}}, nil
}

type OffRamp struct {
	address common.Address
	abi     abi.ABI
	OffRampCaller
	OffRampTransactor
	OffRampFilterer
}

type OffRampCaller struct {
	contract *bind.BoundContract
}

type OffRampTransactor struct {
	contract *bind.BoundContract
}

type OffRampFilterer struct {
	contract *bind.BoundContract
}

type OffRampSession struct {
	Contract     *OffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OffRampCallerSession struct {
	Contract *OffRampCaller
	CallOpts bind.CallOpts
}

type OffRampTransactorSession struct {
	Contract     *OffRampTransactor
	TransactOpts bind.TransactOpts
}

type OffRampRaw struct {
	Contract *OffRamp
}

type OffRampCallerRaw struct {
	Contract *OffRampCaller
}

type OffRampTransactorRaw struct {
	Contract *OffRampTransactor
}

func NewOffRamp(address common.Address, backend bind.ContractBackend) (*OffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(OffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OffRamp{address: address, abi: abi, OffRampCaller: OffRampCaller{contract: contract}, OffRampTransactor: OffRampTransactor{contract: contract}, OffRampFilterer: OffRampFilterer{contract: contract}}, nil
}

func NewOffRampCaller(address common.Address, caller bind.ContractCaller) (*OffRampCaller, error) {
	contract, err := bindOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OffRampCaller{contract: contract}, nil
}

func NewOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*OffRampTransactor, error) {
	contract, err := bindOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OffRampTransactor{contract: contract}, nil
}

func NewOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*OffRampFilterer, error) {
	contract, err := bindOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OffRampFilterer{contract: contract}, nil
}

func bindOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OffRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_OffRamp *OffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffRamp.Contract.OffRampCaller.contract.Call(opts, result, method, params...)
}

func (_OffRamp *OffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRamp.Contract.OffRampTransactor.contract.Transfer(opts)
}

func (_OffRamp *OffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffRamp.Contract.OffRampTransactor.contract.Transact(opts, method, params...)
}

func (_OffRamp *OffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_OffRamp *OffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRamp.Contract.contract.Transfer(opts)
}

func (_OffRamp *OffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_OffRamp *OffRampCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRamp *OffRampSession) CHAINID() (*big.Int, error) {
	return _OffRamp.Contract.CHAINID(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) CHAINID() (*big.Int, error) {
	return _OffRamp.Contract.CHAINID(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "SOURCE_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRamp *OffRampSession) SOURCECHAINID() (*big.Int, error) {
	return _OffRamp.Contract.SOURCECHAINID(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) SOURCECHAINID() (*big.Int, error) {
	return _OffRamp.Contract.SOURCECHAINID(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRamp *OffRampSession) GetAFN() (common.Address, error) {
	return _OffRamp.Contract.GetAFN(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetAFN() (common.Address, error) {
	return _OffRamp.Contract.GetAFN(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) GetExecuted(opts *bind.CallOpts, sequenceNumber uint64) (bool, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getExecuted", sequenceNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OffRamp *OffRampSession) GetExecuted(sequenceNumber uint64) (bool, error) {
	return _OffRamp.Contract.GetExecuted(&_OffRamp.CallOpts, sequenceNumber)
}

func (_OffRamp *OffRampCallerSession) GetExecuted(sequenceNumber uint64) (bool, error) {
	return _OffRamp.Contract.GetExecuted(&_OffRamp.CallOpts, sequenceNumber)
}

func (_OffRamp *OffRampCaller) GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getFeed", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRamp *OffRampSession) GetFeed(token common.Address) (common.Address, error) {
	return _OffRamp.Contract.GetFeed(&_OffRamp.CallOpts, token)
}

func (_OffRamp *OffRampCallerSession) GetFeed(token common.Address) (common.Address, error) {
	return _OffRamp.Contract.GetFeed(&_OffRamp.CallOpts, token)
}

func (_OffRamp *OffRampCaller) GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getFeedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OffRamp *OffRampSession) GetFeedTokens() ([]common.Address, error) {
	return _OffRamp.Contract.GetFeedTokens(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetFeedTokens() ([]common.Address, error) {
	return _OffRamp.Contract.GetFeedTokens(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) GetLastReport(opts *bind.CallOpts) (CCIPRelayReport, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getLastReport")

	if err != nil {
		return *new(CCIPRelayReport), err
	}

	out0 := *abi.ConvertType(out[0], new(CCIPRelayReport)).(*CCIPRelayReport)

	return out0, err

}

func (_OffRamp *OffRampSession) GetLastReport() (CCIPRelayReport, error) {
	return _OffRamp.Contract.GetLastReport(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetLastReport() (CCIPRelayReport, error) {
	return _OffRamp.Contract.GetLastReport(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getMaxSecondsWithoutAFNHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRamp *OffRampSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _OffRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _OffRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) GetMerkleRoot(opts *bind.CallOpts, root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getMerkleRoot", root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRamp *OffRampSession) GetMerkleRoot(root [32]byte) (*big.Int, error) {
	return _OffRamp.Contract.GetMerkleRoot(&_OffRamp.CallOpts, root)
}

func (_OffRamp *OffRampCallerSession) GetMerkleRoot(root [32]byte) (*big.Int, error) {
	return _OffRamp.Contract.GetMerkleRoot(&_OffRamp.CallOpts, root)
}

func (_OffRamp *OffRampCaller) GetOffRampConfig(opts *bind.CallOpts) (OffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getOffRampConfig")

	if err != nil {
		return *new(OffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OffRampInterfaceOffRampConfig)).(*OffRampInterfaceOffRampConfig)

	return out0, err

}

func (_OffRamp *OffRampSession) GetOffRampConfig() (OffRampInterfaceOffRampConfig, error) {
	return _OffRamp.Contract.GetOffRampConfig(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetOffRampConfig() (OffRampInterfaceOffRampConfig, error) {
	return _OffRamp.Contract.GetOffRampConfig(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRamp *OffRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _OffRamp.Contract.GetPool(&_OffRamp.CallOpts, sourceToken)
}

func (_OffRamp *OffRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _OffRamp.Contract.GetPool(&_OffRamp.CallOpts, sourceToken)
}

func (_OffRamp *OffRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OffRamp *OffRampSession) GetPoolTokens() ([]common.Address, error) {
	return _OffRamp.Contract.GetPoolTokens(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _OffRamp.Contract.GetPoolTokens(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRamp *OffRampSession) GetRouter() (common.Address, error) {
	return _OffRamp.Contract.GetRouter(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetRouter() (common.Address, error) {
	return _OffRamp.Contract.GetRouter(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "isHealthy", timeNow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OffRamp *OffRampSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _OffRamp.Contract.IsHealthy(&_OffRamp.CallOpts, timeNow)
}

func (_OffRamp *OffRampCallerSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _OffRamp.Contract.IsHealthy(&_OffRamp.CallOpts, timeNow)
}

func (_OffRamp *OffRampCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OffRamp *OffRampSession) IsPool(addr common.Address) (bool, error) {
	return _OffRamp.Contract.IsPool(&_OffRamp.CallOpts, addr)
}

func (_OffRamp *OffRampCallerSession) IsPool(addr common.Address) (bool, error) {
	return _OffRamp.Contract.IsPool(&_OffRamp.CallOpts, addr)
}

func (_OffRamp *OffRampCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_OffRamp *OffRampSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _OffRamp.Contract.LatestConfigDetails(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _OffRamp.Contract.LatestConfigDetails(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_OffRamp *OffRampSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _OffRamp.Contract.LatestConfigDigestAndEpoch(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _OffRamp.Contract.LatestConfigDigestAndEpoch(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) MerkleRoot(opts *bind.CallOpts, report CCIPExecutionReport) ([32]byte, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "merkleRoot", report)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_OffRamp *OffRampSession) MerkleRoot(report CCIPExecutionReport) ([32]byte, error) {
	return _OffRamp.Contract.MerkleRoot(&_OffRamp.CallOpts, report)
}

func (_OffRamp *OffRampCallerSession) MerkleRoot(report CCIPExecutionReport) ([32]byte, error) {
	return _OffRamp.Contract.MerkleRoot(&_OffRamp.CallOpts, report)
}

func (_OffRamp *OffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRamp *OffRampSession) Owner() (common.Address, error) {
	return _OffRamp.Contract.Owner(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) Owner() (common.Address, error) {
	return _OffRamp.Contract.Owner(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OffRamp *OffRampSession) Paused() (bool, error) {
	return _OffRamp.Contract.Paused(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) Paused() (bool, error) {
	return _OffRamp.Contract.Paused(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OffRamp *OffRampSession) Transmitters() ([]common.Address, error) {
	return _OffRamp.Contract.Transmitters(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) Transmitters() ([]common.Address, error) {
	return _OffRamp.Contract.Transmitters(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OffRamp *OffRampSession) TypeAndVersion() (string, error) {
	return _OffRamp.Contract.TypeAndVersion(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) TypeAndVersion() (string, error) {
	return _OffRamp.Contract.TypeAndVersion(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_OffRamp *OffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffRamp.Contract.AcceptOwnership(&_OffRamp.TransactOpts)
}

func (_OffRamp *OffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffRamp.Contract.AcceptOwnership(&_OffRamp.TransactOpts)
}

func (_OffRamp *OffRampTransactor) AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "addFeed", token, feed)
}

func (_OffRamp *OffRampSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.AddFeed(&_OffRamp.TransactOpts, token, feed)
}

func (_OffRamp *OffRampTransactorSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.AddFeed(&_OffRamp.TransactOpts, token, feed)
}

func (_OffRamp *OffRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_OffRamp *OffRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.AddPool(&_OffRamp.TransactOpts, token, pool)
}

func (_OffRamp *OffRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.AddPool(&_OffRamp.TransactOpts, token, pool)
}

func (_OffRamp *OffRampTransactor) ExecuteTransaction(opts *bind.TransactOpts, report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "executeTransaction", report, needFee)
}

func (_OffRamp *OffRampSession) ExecuteTransaction(report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _OffRamp.Contract.ExecuteTransaction(&_OffRamp.TransactOpts, report, needFee)
}

func (_OffRamp *OffRampTransactorSession) ExecuteTransaction(report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _OffRamp.Contract.ExecuteTransaction(&_OffRamp.TransactOpts, report, needFee)
}

func (_OffRamp *OffRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "pause")
}

func (_OffRamp *OffRampSession) Pause() (*types.Transaction, error) {
	return _OffRamp.Contract.Pause(&_OffRamp.TransactOpts)
}

func (_OffRamp *OffRampTransactorSession) Pause() (*types.Transaction, error) {
	return _OffRamp.Contract.Pause(&_OffRamp.TransactOpts)
}

func (_OffRamp *OffRampTransactor) RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "removeFeed", token, feed)
}

func (_OffRamp *OffRampSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.RemoveFeed(&_OffRamp.TransactOpts, token, feed)
}

func (_OffRamp *OffRampTransactorSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.RemoveFeed(&_OffRamp.TransactOpts, token, feed)
}

func (_OffRamp *OffRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_OffRamp *OffRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.RemovePool(&_OffRamp.TransactOpts, token, pool)
}

func (_OffRamp *OffRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.RemovePool(&_OffRamp.TransactOpts, token, pool)
}

func (_OffRamp *OffRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "setAFN", afn)
}

func (_OffRamp *OffRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.SetAFN(&_OffRamp.TransactOpts, afn)
}

func (_OffRamp *OffRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.SetAFN(&_OffRamp.TransactOpts, afn)
}

func (_OffRamp *OffRampTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_OffRamp *OffRampSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _OffRamp.Contract.SetConfig(&_OffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_OffRamp *OffRampTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _OffRamp.Contract.SetConfig(&_OffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_OffRamp *OffRampTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_OffRamp *OffRampSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _OffRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_OffRamp.TransactOpts, newTime)
}

func (_OffRamp *OffRampTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _OffRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_OffRamp.TransactOpts, newTime)
}

func (_OffRamp *OffRampTransactor) SetOffRampConfig(opts *bind.TransactOpts, config OffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "setOffRampConfig", config)
}

func (_OffRamp *OffRampSession) SetOffRampConfig(config OffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _OffRamp.Contract.SetOffRampConfig(&_OffRamp.TransactOpts, config)
}

func (_OffRamp *OffRampTransactorSession) SetOffRampConfig(config OffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _OffRamp.Contract.SetOffRampConfig(&_OffRamp.TransactOpts, config)
}

func (_OffRamp *OffRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "setRouter", router)
}

func (_OffRamp *OffRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.SetRouter(&_OffRamp.TransactOpts, router)
}

func (_OffRamp *OffRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.SetRouter(&_OffRamp.TransactOpts, router)
}

func (_OffRamp *OffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_OffRamp *OffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.TransferOwnership(&_OffRamp.TransactOpts, to)
}

func (_OffRamp *OffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.TransferOwnership(&_OffRamp.TransactOpts, to)
}

func (_OffRamp *OffRampTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_OffRamp *OffRampSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffRamp.Contract.Transmit(&_OffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_OffRamp *OffRampTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffRamp.Contract.Transmit(&_OffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_OffRamp *OffRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "unpause")
}

func (_OffRamp *OffRampSession) Unpause() (*types.Transaction, error) {
	return _OffRamp.Contract.Unpause(&_OffRamp.TransactOpts)
}

func (_OffRamp *OffRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _OffRamp.Contract.Unpause(&_OffRamp.TransactOpts)
}

func (_OffRamp *OffRampTransactor) WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "withdrawAccumulatedFees", feeToken, recipient, amount)
}

func (_OffRamp *OffRampSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OffRamp.Contract.WithdrawAccumulatedFees(&_OffRamp.TransactOpts, feeToken, recipient, amount)
}

func (_OffRamp *OffRampTransactorSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OffRamp.Contract.WithdrawAccumulatedFees(&_OffRamp.TransactOpts, feeToken, recipient, amount)
}

type OffRampAFNMaxHeartbeatTimeSetIterator struct {
	Event *OffRampAFNMaxHeartbeatTimeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampAFNMaxHeartbeatTimeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampAFNMaxHeartbeatTimeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampAFNMaxHeartbeatTimeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampAFNMaxHeartbeatTimeSetIterator) Error() error {
	return it.fail
}

func (it *OffRampAFNMaxHeartbeatTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampAFNMaxHeartbeatTimeSet struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log
}

func (_OffRamp *OffRampFilterer) FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*OffRampAFNMaxHeartbeatTimeSetIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return &OffRampAFNMaxHeartbeatTimeSetIterator{contract: _OffRamp.contract, event: "AFNMaxHeartbeatTimeSet", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *OffRampAFNMaxHeartbeatTimeSet) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampAFNMaxHeartbeatTimeSet)
				if err := _OffRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParseAFNMaxHeartbeatTimeSet(log types.Log) (*OffRampAFNMaxHeartbeatTimeSet, error) {
	event := new(OffRampAFNMaxHeartbeatTimeSet)
	if err := _OffRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampAFNSetIterator struct {
	Event *OffRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampAFNSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampAFNSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *OffRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_OffRamp *OffRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*OffRampAFNSetIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &OffRampAFNSetIterator{contract: _OffRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *OffRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampAFNSet)
				if err := _OffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParseAFNSet(log types.Log) (*OffRampAFNSet, error) {
	event := new(OffRampAFNSet)
	if err := _OffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampConfigSetIterator struct {
	Event *OffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_OffRamp *OffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OffRampConfigSetIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffRampConfigSetIterator{contract: _OffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampConfigSet)
				if err := _OffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParseConfigSet(log types.Log) (*OffRampConfigSet, error) {
	event := new(OffRampConfigSet)
	if err := _OffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampCrossChainMessageExecutedIterator struct {
	Event *OffRampCrossChainMessageExecuted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampCrossChainMessageExecutedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampCrossChainMessageExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampCrossChainMessageExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampCrossChainMessageExecutedIterator) Error() error {
	return it.fail
}

func (it *OffRampCrossChainMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampCrossChainMessageExecuted struct {
	SequenceNumber uint64
	Raw            types.Log
}

func (_OffRamp *OffRampFilterer) FilterCrossChainMessageExecuted(opts *bind.FilterOpts, sequenceNumber []uint64) (*OffRampCrossChainMessageExecutedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "CrossChainMessageExecuted", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &OffRampCrossChainMessageExecutedIterator{contract: _OffRamp.contract, event: "CrossChainMessageExecuted", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchCrossChainMessageExecuted(opts *bind.WatchOpts, sink chan<- *OffRampCrossChainMessageExecuted, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "CrossChainMessageExecuted", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampCrossChainMessageExecuted)
				if err := _OffRamp.contract.UnpackLog(event, "CrossChainMessageExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParseCrossChainMessageExecuted(log types.Log) (*OffRampCrossChainMessageExecuted, error) {
	event := new(OffRampCrossChainMessageExecuted)
	if err := _OffRamp.contract.UnpackLog(event, "CrossChainMessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampFeedAddedIterator struct {
	Event *OffRampFeedAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampFeedAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampFeedAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampFeedAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampFeedAddedIterator) Error() error {
	return it.fail
}

func (it *OffRampFeedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampFeedAdded struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_OffRamp *OffRampFilterer) FilterFeedAdded(opts *bind.FilterOpts) (*OffRampFeedAddedIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return &OffRampFeedAddedIterator{contract: _OffRamp.contract, event: "FeedAdded", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *OffRampFeedAdded) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampFeedAdded)
				if err := _OffRamp.contract.UnpackLog(event, "FeedAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParseFeedAdded(log types.Log) (*OffRampFeedAdded, error) {
	event := new(OffRampFeedAdded)
	if err := _OffRamp.contract.UnpackLog(event, "FeedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampFeedRemovedIterator struct {
	Event *OffRampFeedRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampFeedRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampFeedRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampFeedRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampFeedRemovedIterator) Error() error {
	return it.fail
}

func (it *OffRampFeedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampFeedRemoved struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_OffRamp *OffRampFilterer) FilterFeedRemoved(opts *bind.FilterOpts) (*OffRampFeedRemovedIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return &OffRampFeedRemovedIterator{contract: _OffRamp.contract, event: "FeedRemoved", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *OffRampFeedRemoved) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampFeedRemoved)
				if err := _OffRamp.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParseFeedRemoved(log types.Log) (*OffRampFeedRemoved, error) {
	event := new(OffRampFeedRemoved)
	if err := _OffRamp.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampFeesWithdrawnIterator struct {
	Event *OffRampFeesWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampFeesWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampFeesWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampFeesWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampFeesWithdrawnIterator) Error() error {
	return it.fail
}

func (it *OffRampFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampFeesWithdrawn struct {
	FeeToken  common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_OffRamp *OffRampFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*OffRampFeesWithdrawnIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &OffRampFeesWithdrawnIterator{contract: _OffRamp.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *OffRampFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampFeesWithdrawn)
				if err := _OffRamp.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParseFeesWithdrawn(log types.Log) (*OffRampFeesWithdrawn, error) {
	event := new(OffRampFeesWithdrawn)
	if err := _OffRamp.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOffRampConfigSetIterator struct {
	Event *OffRampOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOffRampConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampOffRampConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOffRampConfigSet struct {
	Config OffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_OffRamp *OffRampFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*OffRampOffRampConfigSetIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffRampOffRampConfigSetIterator{contract: _OffRamp.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOffRampConfigSet)
				if err := _OffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParseOffRampConfigSet(log types.Log) (*OffRampOffRampConfigSet, error) {
	event := new(OffRampOffRampConfigSet)
	if err := _OffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOffRampRouterSetIterator struct {
	Event *OffRampOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOffRampRouterSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampOffRampRouterSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *OffRampOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOffRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_OffRamp *OffRampFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts) (*OffRampOffRampRouterSetIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "OffRampRouterSet")
	if err != nil {
		return nil, err
	}
	return &OffRampOffRampRouterSetIterator{contract: _OffRamp.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *OffRampOffRampRouterSet) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "OffRampRouterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOffRampRouterSet)
				if err := _OffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParseOffRampRouterSet(log types.Log) (*OffRampOffRampRouterSet, error) {
	event := new(OffRampOffRampRouterSet)
	if err := _OffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOwnershipTransferRequestedIterator struct {
	Event *OffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OffRamp *OffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffRampOwnershipTransferRequestedIterator{contract: _OffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOwnershipTransferRequested)
				if err := _OffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*OffRampOwnershipTransferRequested, error) {
	event := new(OffRampOwnershipTransferRequested)
	if err := _OffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOwnershipTransferredIterator struct {
	Event *OffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OffRamp *OffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffRampOwnershipTransferredIterator{contract: _OffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOwnershipTransferred)
				if err := _OffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParseOwnershipTransferred(log types.Log) (*OffRampOwnershipTransferred, error) {
	event := new(OffRampOwnershipTransferred)
	if err := _OffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampPausedIterator struct {
	Event *OffRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampPausedIterator) Error() error {
	return it.fail
}

func (it *OffRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_OffRamp *OffRampFilterer) FilterPaused(opts *bind.FilterOpts) (*OffRampPausedIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OffRampPausedIterator{contract: _OffRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OffRampPaused) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampPaused)
				if err := _OffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParsePaused(log types.Log) (*OffRampPaused, error) {
	event := new(OffRampPaused)
	if err := _OffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampPoolAddedIterator struct {
	Event *OffRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampPoolAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampPoolAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *OffRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_OffRamp *OffRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*OffRampPoolAddedIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &OffRampPoolAddedIterator{contract: _OffRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *OffRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampPoolAdded)
				if err := _OffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParsePoolAdded(log types.Log) (*OffRampPoolAdded, error) {
	event := new(OffRampPoolAdded)
	if err := _OffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampPoolRemovedIterator struct {
	Event *OffRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampPoolRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampPoolRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *OffRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_OffRamp *OffRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*OffRampPoolRemovedIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &OffRampPoolRemovedIterator{contract: _OffRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *OffRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampPoolRemoved)
				if err := _OffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParsePoolRemoved(log types.Log) (*OffRampPoolRemoved, error) {
	event := new(OffRampPoolRemoved)
	if err := _OffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampReportAcceptedIterator struct {
	Event *OffRampReportAccepted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampReportAcceptedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampReportAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampReportAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampReportAcceptedIterator) Error() error {
	return it.fail
}

func (it *OffRampReportAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampReportAccepted struct {
	Report CCIPRelayReport
	Raw    types.Log
}

func (_OffRamp *OffRampFilterer) FilterReportAccepted(opts *bind.FilterOpts) (*OffRampReportAcceptedIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "ReportAccepted")
	if err != nil {
		return nil, err
	}
	return &OffRampReportAcceptedIterator{contract: _OffRamp.contract, event: "ReportAccepted", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchReportAccepted(opts *bind.WatchOpts, sink chan<- *OffRampReportAccepted) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "ReportAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampReportAccepted)
				if err := _OffRamp.contract.UnpackLog(event, "ReportAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParseReportAccepted(log types.Log) (*OffRampReportAccepted, error) {
	event := new(OffRampReportAccepted)
	if err := _OffRamp.contract.UnpackLog(event, "ReportAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampTransmittedIterator struct {
	Event *OffRampTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampTransmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampTransmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampTransmittedIterator) Error() error {
	return it.fail
}

func (it *OffRampTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_OffRamp *OffRampFilterer) FilterTransmitted(opts *bind.FilterOpts) (*OffRampTransmittedIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &OffRampTransmittedIterator{contract: _OffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OffRampTransmitted) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampTransmitted)
				if err := _OffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParseTransmitted(log types.Log) (*OffRampTransmitted, error) {
	event := new(OffRampTransmitted)
	if err := _OffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampUnpausedIterator struct {
	Event *OffRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(OffRampUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *OffRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *OffRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_OffRamp *OffRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OffRampUnpausedIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OffRampUnpausedIterator{contract: _OffRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OffRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampUnpaused)
				if err := _OffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_OffRamp *OffRampFilterer) ParseUnpaused(log types.Log) (*OffRampUnpaused, error) {
	event := new(OffRampUnpaused)
	if err := _OffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}
type LatestConfigDigestAndEpoch struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}

func (_OffRamp *OffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OffRamp.abi.Events["AFNMaxHeartbeatTimeSet"].ID:
		return _OffRamp.ParseAFNMaxHeartbeatTimeSet(log)
	case _OffRamp.abi.Events["AFNSet"].ID:
		return _OffRamp.ParseAFNSet(log)
	case _OffRamp.abi.Events["ConfigSet"].ID:
		return _OffRamp.ParseConfigSet(log)
	case _OffRamp.abi.Events["CrossChainMessageExecuted"].ID:
		return _OffRamp.ParseCrossChainMessageExecuted(log)
	case _OffRamp.abi.Events["FeedAdded"].ID:
		return _OffRamp.ParseFeedAdded(log)
	case _OffRamp.abi.Events["FeedRemoved"].ID:
		return _OffRamp.ParseFeedRemoved(log)
	case _OffRamp.abi.Events["FeesWithdrawn"].ID:
		return _OffRamp.ParseFeesWithdrawn(log)
	case _OffRamp.abi.Events["OffRampConfigSet"].ID:
		return _OffRamp.ParseOffRampConfigSet(log)
	case _OffRamp.abi.Events["OffRampRouterSet"].ID:
		return _OffRamp.ParseOffRampRouterSet(log)
	case _OffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _OffRamp.ParseOwnershipTransferRequested(log)
	case _OffRamp.abi.Events["OwnershipTransferred"].ID:
		return _OffRamp.ParseOwnershipTransferred(log)
	case _OffRamp.abi.Events["Paused"].ID:
		return _OffRamp.ParsePaused(log)
	case _OffRamp.abi.Events["PoolAdded"].ID:
		return _OffRamp.ParsePoolAdded(log)
	case _OffRamp.abi.Events["PoolRemoved"].ID:
		return _OffRamp.ParsePoolRemoved(log)
	case _OffRamp.abi.Events["ReportAccepted"].ID:
		return _OffRamp.ParseReportAccepted(log)
	case _OffRamp.abi.Events["Transmitted"].ID:
		return _OffRamp.ParseTransmitted(log)
	case _OffRamp.abi.Events["Unpaused"].ID:
		return _OffRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (OffRampAFNMaxHeartbeatTimeSet) Topic() common.Hash {
	return common.HexToHash("0x72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251")
}

func (OffRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (OffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (OffRampCrossChainMessageExecuted) Topic() common.Hash {
	return common.HexToHash("0x88a2c71ae86800edd209d83430f348ce70539bd40085d045340ff569f058a2a0")
}

func (OffRampFeedAdded) Topic() common.Hash {
	return common.HexToHash("0x037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb601305")
}

func (OffRampFeedRemoved) Topic() common.Hash {
	return common.HexToHash("0xa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d")
}

func (OffRampFeesWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8")
}

func (OffRampOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xe4cd88b1b5e20a0b843af3207ba74d3f84af4acff45830469490ac9c6ab8a745")
}

func (OffRampOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4")
}

func (OffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (OffRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (OffRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (OffRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (OffRampReportAccepted) Topic() common.Hash {
	return common.HexToHash("0x6a2424b54e9b18f0c3d440b6003882f5ffa34272205608ad9451c921c0d99a35")
}

func (OffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (OffRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_OffRamp *OffRamp) Address() common.Address {
	return _OffRamp.address
}

type OffRampInterface interface {
	CHAINID(opts *bind.CallOpts) (*big.Int, error)

	SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetExecuted(opts *bind.CallOpts, sequenceNumber uint64) (bool, error)

	GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error)

	GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetLastReport(opts *bind.CallOpts) (CCIPRelayReport, error)

	GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error)

	GetMerkleRoot(opts *bind.CallOpts, root [32]byte) (*big.Int, error)

	GetOffRampConfig(opts *bind.CallOpts) (OffRampInterfaceOffRampConfig, error)

	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error)

	IsPool(opts *bind.CallOpts, addr common.Address) (bool, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	MerkleRoot(opts *bind.CallOpts, report CCIPExecutionReport) ([32]byte, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	ExecuteTransaction(opts *bind.TransactOpts, report CCIPExecutionReport, needFee bool) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error)

	SetOffRampConfig(opts *bind.TransactOpts, config OffRampInterfaceOffRampConfig) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*OffRampAFNMaxHeartbeatTimeSetIterator, error)

	WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *OffRampAFNMaxHeartbeatTimeSet) (event.Subscription, error)

	ParseAFNMaxHeartbeatTimeSet(log types.Log) (*OffRampAFNMaxHeartbeatTimeSet, error)

	FilterAFNSet(opts *bind.FilterOpts) (*OffRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *OffRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*OffRampAFNSet, error)

	FilterConfigSet(opts *bind.FilterOpts) (*OffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*OffRampConfigSet, error)

	FilterCrossChainMessageExecuted(opts *bind.FilterOpts, sequenceNumber []uint64) (*OffRampCrossChainMessageExecutedIterator, error)

	WatchCrossChainMessageExecuted(opts *bind.WatchOpts, sink chan<- *OffRampCrossChainMessageExecuted, sequenceNumber []uint64) (event.Subscription, error)

	ParseCrossChainMessageExecuted(log types.Log) (*OffRampCrossChainMessageExecuted, error)

	FilterFeedAdded(opts *bind.FilterOpts) (*OffRampFeedAddedIterator, error)

	WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *OffRampFeedAdded) (event.Subscription, error)

	ParseFeedAdded(log types.Log) (*OffRampFeedAdded, error)

	FilterFeedRemoved(opts *bind.FilterOpts) (*OffRampFeedRemovedIterator, error)

	WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *OffRampFeedRemoved) (event.Subscription, error)

	ParseFeedRemoved(log types.Log) (*OffRampFeedRemoved, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*OffRampFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *OffRampFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*OffRampFeesWithdrawn, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*OffRampOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*OffRampOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts) (*OffRampOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *OffRampOffRampRouterSet) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*OffRampOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OffRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*OffRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *OffRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*OffRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*OffRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *OffRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*OffRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*OffRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *OffRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*OffRampPoolRemoved, error)

	FilterReportAccepted(opts *bind.FilterOpts) (*OffRampReportAcceptedIterator, error)

	WatchReportAccepted(opts *bind.WatchOpts, sink chan<- *OffRampReportAccepted) (event.Subscription, error)

	ParseReportAccepted(log types.Log) (*OffRampReportAccepted, error)

	FilterTransmitted(opts *bind.FilterOpts) (*OffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*OffRampTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*OffRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OffRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*OffRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
