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

type CCIPMerkleProof struct {
	Path  [][32]byte
	Index *big.Int
}

type CCIPMessage struct {
	SequenceNumber *big.Int
	SourceChainId  *big.Int
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
	Options            []byte
}

type CCIPRelayReport struct {
	MerkleRoot        [32]byte
	MinSequenceNumber *big.Int
	MaxSequenceNumber *big.Int
}

type OffRampInterfaceOffRampConfig struct {
	ExecutionFeeJuels     uint64
	ExecutionDelaySeconds uint64
	MaxDataSize           uint64
	MaxTokensLength       uint64
}

var OffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"contractAggregatorV2V3Interface[]\",\"name\":\"feeds\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionDelayError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"InvalidExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceFeedConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MerkleProofError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RelayReportError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lastMaxSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMinSequenceNumber\",\"type\":\"uint256\"}],\"name\":\"SequenceError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenFeedMistmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMistmatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"UnexpectedPayloadData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"CrossChainMessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"minSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSequenceNumber\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCCIP.RelayReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"ReportAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"addFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"needFee\",\"type\":\"bool\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"getExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastReport\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"minSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSequenceNumber\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.RelayReport\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRampConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"removeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setOffRampConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005aad38038062005aad833981016040819052620000349162000747565b6000805460ff191681556001908790869082908990889088903390819081620000a45760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000de57620000de81620003ed565b5050506001600160a01b0382161580620000f6575080155b156200011557604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03939093169290921790915560035580518251146200015b5760405162d8548360e71b815260040160405180910390fd5b8151620001709060059060208501906200049e565b5060005b8251811015620002545760008282815181106200019557620001956200081d565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001df57620001df6200081d565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff19166001179055806200024b8162000833565b91505062000174565b50505080518251146200027a5760405163ee9d106b60e01b815260040160405180910390fd5b81516200028f9060089060208501906200049e565b5060005b82518110156200035c576000828281518110620002b457620002b46200081d565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060076000868581518110620002fe57620002fe6200081d565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b02929091169190911790555080620003538162000833565b91505062000293565b505050151560805260a09790975250505060c0929092525050805160148054602084015160408501516060909501516001600160401b03908116600160c01b026001600160c01b03968216600160801b02969096166001600160801b0392821668010000000000000000026001600160801b031990941691909516179190911716919091179190911790556200085b565b336001600160a01b03821603620004475760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016200009b565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620004f6579160200282015b82811115620004f657825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620004bf565b506200050492915062000508565b5090565b5b8082111562000504576000815560010162000509565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156200056057620005606200051f565b604052919050565b60006001600160401b038211156200058457620005846200051f565b5060051b60200190565b6001600160a01b0381168114620005a457600080fd5b50565b600082601f830112620005b957600080fd5b81516020620005d2620005cc8362000568565b62000535565b82815260059290921b84018101918181019086841115620005f257600080fd5b8286015b848110156200061a5780516200060c816200058e565b8352918301918301620005f6565b509695505050505050565b600082601f8301126200063757600080fd5b815160206200064a620005cc8362000568565b82815260059290921b840181019181810190868411156200066a57600080fd5b8286015b848110156200061a57805162000684816200058e565b83529183019183016200066e565b80516200069f816200058e565b919050565b80516001600160401b03811681146200069f57600080fd5b600060808284031215620006cf57600080fd5b604051608081016001600160401b0381118282101715620006f457620006f46200051f565b6040529050806200070583620006a4565b81526200071560208401620006a4565b60208201526200072860408401620006a4565b60408201526200073b60608401620006a4565b60608201525092915050565b600080600080600080600080610160898b0312156200076557600080fd5b885160208a015160408b015191995097506001600160401b03808211156200078c57600080fd5b6200079a8c838d01620005a7565b975060608b0151915080821115620007b157600080fd5b620007bf8c838d01620005a7565b965060808b0151915080821115620007d657600080fd5b50620007e58b828c0162000625565b945050620007f660a08a0162000692565b925060c089015191506200080e8a60e08b01620006bc565b90509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b6000600182016200085457634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05161521b6200089260003960006104b6015260008181610411015261332201526000612075015261521b6000f3fe608060405234801561001057600080fd5b50600436106102415760003560e01c80638141183411610145578063afcb95d7116100bd578063bbe4f6db1161008c578063eb511dd411610071578063eb511dd41461069b578063eefa7a3e146106ae578063f2fde38b1461070457600080fd5b8063bbe4f6db1461064f578063e3d0e7121461068857600080fd5b8063afcb95d714610601578063b034909c14610621578063b1dc65a414610629578063b6608c3b1461063c57600080fd5b806385e1f4d0116101145780638da5cb5b116100f95780638da5cb5b146104e0578063a7206cd614610503578063a8ebd0f41461052357600080fd5b806385e1f4d0146104b157806389c06568146104d857600080fd5b8063814118341461045c57806381be8fa41461047157806381ff7048146104795780638456cb59146104a957600080fd5b8063461c551b116101d85780635c975abb116101a757806374be21501161018c57806374be21501461040c578063768c577b1461044157806379ba50971461045457600080fd5b80635c975abb146103ee578063744b92e2146103f957600080fd5b8063461c551b1461037c5780635853c6271461038f57806359e96b5b146103a25780635b16ebb7146103b557600080fd5b80632222dd42116102145780632222dd42146103305780632b898c251461034e5780633b8d08ef146103615780633f4ba83a1461037457600080fd5b8063108ee5fc1461024657806316b8e7311461025b578063181f5a77146102be57806321947507146102fd575b600080fd5b610259610254366004614100565b610717565b005b610294610269366004614100565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b604080518082018252600d81527f4f666652616d7020302e302e3100000000000000000000000000000000000000602082015290516102b59190614197565b61032061030b3660046141aa565b60009081526010602052604090205460ff1690565b60405190151581526020016102b5565b60025473ffffffffffffffffffffffffffffffffffffffff16610294565b61025961035c3660046141c3565b6107f3565b61025961036f366004614684565b610bc3565b61025961143f565b61025961038a3660046146fc565b611451565b61025961039d3660046141c3565b6114a3565b6102596103b0366004614714565b6116bb565b6103206103c3366004614100565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff16610320565b6102596104073660046141c3565b611739565b6104337f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016102b5565b61043361044f366004614755565b611b2e565b610259611cbd565b610464611ddf565b6040516102b5919061480a565b610464611e4e565b600b546009546040805163ffffffff808516825264010000000090940490931660208401528201526060016102b5565b610259611ebb565b6104337f000000000000000000000000000000000000000000000000000000000000000081565b610464611ecb565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610294565b6104336105113660046141aa565b6000908152600f602052604090205490565b6105bd604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260145467ffffffffffffffff808216835268010000000000000000820481166020840152700100000000000000000000000000000000820481169383019390935278010000000000000000000000000000000000000000000000009004909116606082015290565b6040516102b59190815167ffffffffffffffff9081168252602080840151821690830152604080840151821690830152606092830151169181019190915260800190565b6040805160018152600060208201819052918101919091526060016102b5565b600354610433565b610259610637366004614869565b611f38565b61025961064a3660046141aa565b6125e1565b61029461065d366004614100565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b610259610696366004614980565b612661565b6102596106a93660046141c3565b613046565b6040805160608082018352600080835260208084018290529284015282518082018452601154808252601254828501908152601354928601928352855191825251938101939093525192820192909252016102b5565b610259610712366004614100565b613286565b61071f61329a565b73ffffffffffffffffffffffffffffffffffffffff811661076c576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b6107fb61329a565b6008546000819003610839576040517f2e70248b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906108d4576040517f3917193900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff161461093d576040517f6c17b98700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600861094c600185614a7c565b8154811061095c5761095c614a93565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600882602001516bffffffffffffffffffffffff16815481106109ae576109ae614a93565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1660086109dd600186614a7c565b815481106109ed576109ed614a93565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600883602001516bffffffffffffffffffffffff1681548110610a5b57610a5b614a93565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526007909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556008805480610afd57610afd614ac2565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600783526040808520949094558351908152908716918101919091527fa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d91015b60405180910390a15050505050565b60005460ff1615610c35576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610ca4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc89190614af1565b15610cfe576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b4916004808301926060929190829003018187875af1158015610d6f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d939190614b0e565b9050600354816020015142610da89190614a7c565b1115610de0576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610dec8585611b2e565b6000818152600f6020526040812054919250819003610e3b5784866040517f07e6809a000000000000000000000000000000000000000000000000000000008152600401610c2c929190614c80565b6014544290610e609068010000000000000000900467ffffffffffffffff1683614ce2565b10610e97576040517f15c33ba200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855160009081526010602052604090205460ff1615610ee85785516040517f6a64e9610000000000000000000000000000000000000000000000000000000081526004810191909152602401610c2c565b60608601516080015173ffffffffffffffffffffffffffffffffffffffff1615801590610f33575060608601516080015173ffffffffffffffffffffffffffffffffffffffff163314155b15610f705785516040517fd8e90b980000000000000000000000000000000000000000000000000000000081526004810191909152602401610c2c565b610f7986613320565b610f828661349b565b8551600090815260106020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905583156111c057600080876060015160000151600081518110610fdf57610fdf614a93565b60200260200101519050600061101a8273ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff8116611069576040517f83135fec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166350d25bcd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d89190614cfa565b6014546110ef919067ffffffffffffffff16614d13565b925082156111bc578289606001516020015160008151811061111357611113614a93565b602002602001018181516111279190614a7c565b9052506111338261354a565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081523360048201526024810185905273ffffffffffffffffffffffffffffffffffffffff919091169063ea6192a290604401600060405180830381600087803b1580156111a357600080fd5b505af11580156111b7573d6000803e3d6000fd5b505050505b5050505b60005b606087015151518110156112c4576111fb87606001516000015182815181106111ee576111ee614a93565b602002602001015161354a565b73ffffffffffffffffffffffffffffffffffffffff1663ea6192a2886060015160600151896060015160200151848151811061123957611239614a93565b60200260200101516040518363ffffffff1660e01b815260040161127f92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600060405180830381600087803b15801561129957600080fd5b505af11580156112ad573d6000803e3d6000fd5b5050505080806112bc90614d50565b9150506111c3565b50606080870151015173ffffffffffffffffffffffffffffffffffffffff163b156113c55785606001516060015173ffffffffffffffffffffffffffffffffffffffff16639c598468876040518263ffffffff1660e01b815260040161132a9190614d88565b600060405180830381600087803b15801561134457600080fd5b505af1925050508015611355575060015b6113c0573d808015611383576040519150601f19603f3d011682016040523d82523d6000602084013e611388565b606091505b5086516040517f6a3fd4f2000000000000000000000000000000000000000000000000000000008152610c2c91908390600401614d9b565b61140c565b606086015160a00151511561140c5785516040517fe0244be30000000000000000000000000000000000000000000000000000000081526004810191909152602401610c2c565b85516040517fc51bf0f6d90b467e0849da0ad18a4d9144a4b78b9f83202e1c65cd68f72d4a3890600090a2505050505050565b61144761329a565b61144f6135c6565b565b61145961329a565b8060146114668282614db4565b9050507fe4cd88b1b5e20a0b843af3207ba74d3f84af4acff45830469490ac9c6ab8a745816040516114989190614ec5565b60405180910390a150565b6114ab61329a565b73ffffffffffffffffffffffffffffffffffffffff821615806114e2575073ffffffffffffffffffffffffffffffffffffffff8116155b15611519576040517fee9d106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156115b5576040517f965ffb7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600880546bffffffffffffffffffffffff908116602080870191825288861660008181526007835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821790558351908152908101919091527f037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb60130591015b60405180910390a1505050565b6116c361329a565b6116e473ffffffffffffffffffffffffffffffffffffffff841683836136a7565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8906060016116ae565b61174161329a565b600554600081900361177f576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529061181a576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614611883576040517fd428911900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006005611892600185614a7c565b815481106118a2576118a2614a93565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff16815481106118f4576118f4614a93565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166005611923600186614a7c565b8154811061193357611933614a93565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff16815481106119a1576119a1614a93565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480611a4357611a43614ac2565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101610bb4565b600080600060f81b84604051602001611b479190614d88565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052611b839291602001614f2f565b60405160208183030381529060405280519060200120905060005b835151811015611cb557600084600001518281518110611bc057611bc0614a93565b6020026020010151905060028560200151611bdb9190614fa6565b600003611c39576040517f010000000000000000000000000000000000000000000000000000000000000060208201526021810184905260418101829052606101604051602081830303815290604052805190602001209250611c8c565b6040517f0100000000000000000000000000000000000000000000000000000000000000602082015260218101829052604181018490526061016040516020818303038152906040528051906020012092505b60028560200151611c9d9190614fba565b60208601525080611cad81614d50565b915050611b9e565b509392505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611d3e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610c2c565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600e805480602002602001604051908101604052809291908181526020018280548015611e4457602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611e19575b5050505050905090565b60606008805480602002602001604051908101604052809291908181526020018280548015611e445760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611e19575050505050905090565b611ec361329a565b61144f613739565b60606005805480602002602001604051908101604052809291908181526020018280548015611e445760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611e19575050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c013591611f8e91849163ffffffff851691908e908e90819084018382808284376000920191909152506137f992505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600954808252600a5460ff80821660208501526101009091041692820192909252908314612063576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610c2c565b6120718b8b8b8b8b8b613b6a565b60007f0000000000000000000000000000000000000000000000000000000000000000156120ce576002826020015183604001516120af9190614fce565b6120b99190614ff3565b6120c4906001614fce565b60ff1690506120e4565b60208201516120de906001614fce565b60ff1690505b88811461214d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610c2c565b8887146121b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610c2c565b336000908152600c602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156121f9576121f9615015565b600281111561220a5761220a615015565b905250905060028160200151600281111561222757612227615015565b14801561226e5750600e816000015160ff168154811061224957612249614a93565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b6122d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610c2c565b5050505050600088886040516122eb929190615044565b604051908190038120612302918c90602001615054565b6040516020818303038152906040528051906020012090506123226140bf565b604080518082019091526000808252602082015260005b888110156125bf57600060018588846020811061235857612358614a93565b61236591901a601b614fce565b8d8d8681811061237757612377614a93565b905060200201358c8c8781811061239057612390614a93565b90506020020135604051600081526020016040526040516123cd949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156123ef573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600c602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561246f5761246f615015565b600281111561248057612480615015565b905250925060018360200151600281111561249d5761249d615015565b14612504576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610c2c565b8251849060ff16601f811061251b5761251b614a93565b602002015115612587576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610c2c565b600184846000015160ff16601f81106125a2576125a2614a93565b9115156020909202015250806125b781614d50565b915050612339565b5050505063ffffffff81106125d6576125d6615070565b505050505050505050565b6125e961329a565b80600003612623576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c25191016107e7565b855185518560ff16601f8311156126d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610c2c565b6000811161273e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610c2c565b8183146127cc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610c2c565b6127d7816003614d13565b831161283f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610c2c565b61284761329a565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600d5415612a3a57600d5460009061289f90600190614a7c565b90506000600d82815481106128b6576128b6614a93565b6000918252602082200154600e805473ffffffffffffffffffffffffffffffffffffffff909216935090849081106128f0576128f0614a93565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600c909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600d8054919250908061297057612970614ac2565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600e8054806129d9576129d9614ac2565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550612885915050565b60005b815151811015612ea1576000600c600084600001518481518110612a6357612a63614a93565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612aad57612aad615015565b14612b14576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610c2c565b6040805180820190915260ff821681526001602082015282518051600c9160009185908110612b4557612b45614a93565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612be657612be6615015565b021790555060009150612bf69050565b600c600084602001518481518110612c1057612c10614a93565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612c5a57612c5a615015565b14612cc1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610c2c565b6040805180820190915260ff82168152602081016002815250600c600084602001518481518110612cf457612cf4614a93565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612d9557612d95615015565b02179055505082518051600d925083908110612db357612db3614a93565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600e919083908110612e2f57612e2f614a93565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691909117905580612e9981614d50565b915050612a3d565b506040810151600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600b80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092612f3392869290821691161761509f565b92506101000a81548163ffffffff021916908363ffffffff160217905550612f924630600b60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613c21565b600981905582518051600a805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600b5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598613031988b98919763ffffffff9092169690959194919391926150c7565b60405180910390a15050505050505050505050565b61304e61329a565b73ffffffffffffffffffffffffffffffffffffffff82161580613085575073ffffffffffffffffffffffffffffffffffffffff8116155b156130bc576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015613158576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c91016116ae565b61328e61329a565b61329781613ccc565b50565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16331461144f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610c2c565b7f00000000000000000000000000000000000000000000000000000000000000008160200151146133855780602001516040517fd44bc9eb000000000000000000000000000000000000000000000000000000008152600401610c2c91815260200190565b60145460608201515151780100000000000000000000000000000000000000000000000090910467ffffffffffffffff1610806133ce5750606081015160208101515190515114155b15613405576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601454606082015160a001515170010000000000000000000000000000000090910467ffffffffffffffff16101561329757601454606082015160a00151516040517f8693378900000000000000000000000000000000000000000000000000000000815270010000000000000000000000000000000090920467ffffffffffffffff1660048301526024820152604401610c2c565b606080820151015173ffffffffffffffffffffffffffffffffffffffff163014806134f25750606080820151015173ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff165b156132975760608082015101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610c2c565b73ffffffffffffffffffffffffffffffffffffffff81811660009081526004602052604090205416806135c1576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610c2c565b919050565b60005460ff16613632576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610c2c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052613734908490613dc7565b505050565b60005460ff16156137a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c2c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861367d3390565b60005460ff1615613866576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c2c565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b81526004016020604051808303816000875af11580156138d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138f99190614af1565b1561392f576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b4916004808301926060929190829003018187875af11580156139a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139c49190614b0e565b90506003548160200151426139d99190614a7c565b1115613a11576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082806020019051810190613a279190614b0e565b9050806040015181602001511115613a6b576040517f67a3824c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516060810182526011548082526012546020830152601354928201929092529015613af4576040810151613aa3906001614ce2565b826020015114613af457806040015182602001516040517fcc7f1bd0000000000000000000000000000000000000000000000000000000008152600401610c2c929190918252602082015260400190565b81516000908152600f60209081526040918290204290558351601181905581850180516012558386018051601355845192835290519282019290925290518183015290517f07d7bce06be2a7b0230e4dd0d72523c0407e82419ab1d947c5ddaf59ca36484e9181900360600190a1505050505050565b6000613b77826020614d13565b613b82856020614d13565b613b8e88610144614ce2565b613b989190614ce2565b613ba29190614ce2565b613bad906000614ce2565b9050368114613c18576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610c2c565b50505050505050565b6000808a8a8a8a8a8a8a8a8a604051602001613c459998979695949392919061515d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603613d4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610c2c565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000613e29826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16613ed39092919063ffffffff16565b8051909150156137345780806020019051810190613e479190614af1565b613734576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610c2c565b6060613ee28484600085613eec565b90505b9392505050565b606082471015613f7e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610c2c565b843b613fe6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610c2c565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161400f91906151f2565b60006040518083038185875af1925050503d806000811461404c576040519150601f19603f3d011682016040523d82523d6000602084013e614051565b606091505b509150915061406182828661406c565b979650505050505050565b6060831561407b575081613ee5565b82511561408b5782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2c9190614197565b604051806103e00160405280601f906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff8116811461329757600080fd5b60006020828403121561411257600080fd5b8135613ee5816140de565b60005b83811015614138578181015183820152602001614120565b83811115614147576000848401525b50505050565b6000815180845261416581602086016020860161411d565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613ee5602083018461414d565b6000602082840312156141bc57600080fd5b5035919050565b600080604083850312156141d657600080fd5b82356141e1816140de565b915060208301356141f1816140de565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff8111828210171561424e5761424e6141fc565b60405290565b60405160e0810167ffffffffffffffff8111828210171561424e5761424e6141fc565b6040805190810167ffffffffffffffff8111828210171561424e5761424e6141fc565b6040516060810167ffffffffffffffff8111828210171561424e5761424e6141fc565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715614304576143046141fc565b604052919050565b80356135c1816140de565b600067ffffffffffffffff821115614331576143316141fc565b5060051b60200190565b600082601f83011261434c57600080fd5b8135602061436161435c83614317565b6142bd565b82815260059290921b8401810191818101908684111561438057600080fd5b8286015b848110156143a4578035614397816140de565b8352918301918301614384565b509695505050505050565b600082601f8301126143c057600080fd5b813560206143d061435c83614317565b82815260059290921b840181019181810190868411156143ef57600080fd5b8286015b848110156143a457803583529183019183016143f3565b600082601f83011261441b57600080fd5b813567ffffffffffffffff811115614435576144356141fc565b61446660207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016142bd565b81815284602083860101111561447b57600080fd5b816020850160208301376000918101602001919091529392505050565b6000608082840312156144aa57600080fd5b6144b261422b565b9050813581526020820135602082015260408201356144d0816140de565b6040820152606082013567ffffffffffffffff808211156144f057600080fd5b9083019060e0828603121561450457600080fd5b61450c614254565b82358281111561451b57600080fd5b6145278782860161433b565b82525060208301358281111561453c57600080fd5b614548878286016143af565b602083015250604083013560408201526145646060840161430c565b60608201526145756080840161430c565b608082015260a08301358281111561458c57600080fd5b6145988782860161440a565b60a08301525060c0830135828111156145b057600080fd5b6145bc8782860161440a565b60c083015250606084015250909392505050565b6000604082840312156145e257600080fd5b6145ea614277565b9050813567ffffffffffffffff81111561460357600080fd5b8201601f8101841361461457600080fd5b8035602061462461435c83614317565b82815260059290921b8301810191818101908784111561464357600080fd5b938201935b8385101561466157843582529382019390820190614648565b85525093840135938301939093525092915050565b801515811461329757600080fd5b60008060006060848603121561469957600080fd5b833567ffffffffffffffff808211156146b157600080fd5b6146bd87838801614498565b945060208601359150808211156146d357600080fd5b506146e0868287016145d0565b92505060408401356146f181614676565b809150509250925092565b60006080828403121561470e57600080fd5b50919050565b60008060006060848603121561472957600080fd5b8335614734816140de565b92506020840135614744816140de565b929592945050506040919091013590565b6000806040838503121561476857600080fd5b823567ffffffffffffffff8082111561478057600080fd5b61478c86838701614498565b935060208501359150808211156147a257600080fd5b506147af858286016145d0565b9150509250929050565b600081518084526020808501945080840160005b838110156147ff57815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016147cd565b509495945050505050565b602081526000613ee560208301846147b9565b60008083601f84011261482f57600080fd5b50813567ffffffffffffffff81111561484757600080fd5b6020830191508360208260051b850101111561486257600080fd5b9250929050565b60008060008060008060008060e0898b03121561488557600080fd5b606089018a81111561489657600080fd5b8998503567ffffffffffffffff808211156148b057600080fd5b818b0191508b601f8301126148c457600080fd5b8135818111156148d357600080fd5b8c60208285010111156148e557600080fd5b6020830199508098505060808b013591508082111561490357600080fd5b61490f8c838d0161481d565b909750955060a08b013591508082111561492857600080fd5b506149358b828c0161481d565b999c989b50969995989497949560c00135949350505050565b803560ff811681146135c157600080fd5b67ffffffffffffffff8116811461329757600080fd5b80356135c18161495f565b60008060008060008060c0878903121561499957600080fd5b863567ffffffffffffffff808211156149b157600080fd5b6149bd8a838b0161433b565b975060208901359150808211156149d357600080fd5b6149df8a838b0161433b565b96506149ed60408a0161494e565b95506060890135915080821115614a0357600080fd5b614a0f8a838b0161440a565b9450614a1d60808a01614975565b935060a0890135915080821115614a3357600080fd5b50614a4089828a0161440a565b9150509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015614a8e57614a8e614a4d565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060208284031215614b0357600080fd5b8151613ee581614676565b600060608284031215614b2057600080fd5b614b2861429a565b8251815260208301516020820152604083015160408201528091505092915050565b8051825260006020808301518185015273ffffffffffffffffffffffffffffffffffffffff6040840151166040850152606083015160806060860152805160e06080870152614b9d6101608701826147b9565b838301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80888303810160a08a0152815180845291860193506000929091908601905b80841015614c005784518252938601936001939093019290860190614be0565b50604085015160c08a0152606085015173ffffffffffffffffffffffffffffffffffffffff90811660e08b01526080860151166101008a015260a085015189820383016101208b01529550614c55818761414d565b95505060c084015193508088860301610140890152505050614c77828261414d565b95945050505050565b60408152600060808201845160408085015281815180845260a08601915060209350838301925060005b81811015614cc657835183529284019291840191600101614caa565b5050828701516060860152848103838601526140618187614b4a565b60008219821115614cf557614cf5614a4d565b500190565b600060208284031215614d0c57600080fd5b5051919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614d4b57614d4b614a4d565b500290565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614d8157614d81614a4d565b5060010190565b602081526000613ee56020830184614b4a565b828152604060208201526000613ee2604083018461414d565b8135614dbf8161495f565b67ffffffffffffffff811690508154817fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000082161783556020840135614e038161495f565b6fffffffffffffffff00000000000000008160401b16905080837fffffffffffffffffffffffffffffffff000000000000000000000000000000008416171784556040850135614e528161495f565b77ffffffffffffffff000000000000000000000000000000008160801b1690507fffffffffffffffff000000000000000000000000000000000000000000000000818582861617841717865560608701359350614eae8461495f565b808460c01b16858417831717865550505050505050565b608081018235614ed48161495f565b67ffffffffffffffff9081168352602084013590614ef18261495f565b9081166020840152604084013590614f088261495f565b9081166040840152606084013590614f1f8261495f565b8082166060850152505092915050565b7fff000000000000000000000000000000000000000000000000000000000000008316815260008251614f6981600185016020870161411d565b919091016001019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082614fb557614fb5614f77565b500690565b600082614fc957614fc9614f77565b500490565b600060ff821660ff84168060ff03821115614feb57614feb614a4d565b019392505050565b600060ff83168061500657615006614f77565b8060ff84160491505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600063ffffffff8083168185168083038211156150be576150be614a4d565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526150f78184018a6147b9565b9050828103608084015261510b81896147b9565b905060ff871660a084015282810360c0840152615128818761414d565b905067ffffffffffffffff851660e084015282810361010084015261514d818561414d565b9c9b505050505050505050505050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526151a48285018b6147b9565b915083820360808501526151b8828a6147b9565b915060ff881660a085015283820360c08501526151d5828861414d565b90861660e0850152838103610100850152905061514d818561414d565b6000825161520481846020870161411d565b919091019291505056fea164736f6c634300080d000a",
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

func (_OffRamp *OffRampCaller) GetExecuted(opts *bind.CallOpts, sequenceNumber *big.Int) (bool, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getExecuted", sequenceNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OffRamp *OffRampSession) GetExecuted(sequenceNumber *big.Int) (bool, error) {
	return _OffRamp.Contract.GetExecuted(&_OffRamp.CallOpts, sequenceNumber)
}

func (_OffRamp *OffRampCallerSession) GetExecuted(sequenceNumber *big.Int) (bool, error) {
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

func (_OffRamp *OffRampCaller) MerkleRoot(opts *bind.CallOpts, message CCIPMessage, proof CCIPMerkleProof) ([32]byte, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "merkleRoot", message, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_OffRamp *OffRampSession) MerkleRoot(message CCIPMessage, proof CCIPMerkleProof) ([32]byte, error) {
	return _OffRamp.Contract.MerkleRoot(&_OffRamp.CallOpts, message, proof)
}

func (_OffRamp *OffRampCallerSession) MerkleRoot(message CCIPMessage, proof CCIPMerkleProof) ([32]byte, error) {
	return _OffRamp.Contract.MerkleRoot(&_OffRamp.CallOpts, message, proof)
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

func (_OffRamp *OffRampTransactor) ExecuteTransaction(opts *bind.TransactOpts, message CCIPMessage, proof CCIPMerkleProof, needFee bool) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "executeTransaction", message, proof, needFee)
}

func (_OffRamp *OffRampSession) ExecuteTransaction(message CCIPMessage, proof CCIPMerkleProof, needFee bool) (*types.Transaction, error) {
	return _OffRamp.Contract.ExecuteTransaction(&_OffRamp.TransactOpts, message, proof, needFee)
}

func (_OffRamp *OffRampTransactorSession) ExecuteTransaction(message CCIPMessage, proof CCIPMerkleProof, needFee bool) (*types.Transaction, error) {
	return _OffRamp.Contract.ExecuteTransaction(&_OffRamp.TransactOpts, message, proof, needFee)
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
	SequenceNumber *big.Int
	Raw            types.Log
}

func (_OffRamp *OffRampFilterer) FilterCrossChainMessageExecuted(opts *bind.FilterOpts, sequenceNumber []*big.Int) (*OffRampCrossChainMessageExecutedIterator, error) {

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

func (_OffRamp *OffRampFilterer) WatchCrossChainMessageExecuted(opts *bind.WatchOpts, sink chan<- *OffRampCrossChainMessageExecuted, sequenceNumber []*big.Int) (event.Subscription, error) {

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
	return common.HexToHash("0xc51bf0f6d90b467e0849da0ad18a4d9144a4b78b9f83202e1c65cd68f72d4a38")
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
	return common.HexToHash("0x07d7bce06be2a7b0230e4dd0d72523c0407e82419ab1d947c5ddaf59ca36484e")
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

	GetExecuted(opts *bind.CallOpts, sequenceNumber *big.Int) (bool, error)

	GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error)

	GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetLastReport(opts *bind.CallOpts) (CCIPRelayReport, error)

	GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error)

	GetMerkleRoot(opts *bind.CallOpts, root [32]byte) (*big.Int, error)

	GetOffRampConfig(opts *bind.CallOpts) (OffRampInterfaceOffRampConfig, error)

	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	IsPool(opts *bind.CallOpts, addr common.Address) (bool, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	MerkleRoot(opts *bind.CallOpts, message CCIPMessage, proof CCIPMerkleProof) ([32]byte, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	ExecuteTransaction(opts *bind.TransactOpts, message CCIPMessage, proof CCIPMerkleProof, needFee bool) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error)

	SetOffRampConfig(opts *bind.TransactOpts, config OffRampInterfaceOffRampConfig) (*types.Transaction, error)

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

	FilterCrossChainMessageExecuted(opts *bind.FilterOpts, sequenceNumber []*big.Int) (*OffRampCrossChainMessageExecutedIterator, error)

	WatchCrossChainMessageExecuted(opts *bind.WatchOpts, sink chan<- *OffRampCrossChainMessageExecuted, sequenceNumber []*big.Int) (event.Subscription, error)

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
