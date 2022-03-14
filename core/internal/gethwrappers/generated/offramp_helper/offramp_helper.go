// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package offramp_helper

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

var OffRampHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"contractAggregatorV2V3Interface[]\",\"name\":\"feeds\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionDelaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokensLength\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionDelayError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"InvalidExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceFeedConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MerkleProofError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RelayReportError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lastMaxSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMinSequenceNumber\",\"type\":\"uint256\"}],\"name\":\"SequenceError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenFeedMistmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMistmatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"UnexpectedPayloadData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"CrossChainMessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"ExecutionDelaySecondsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"}],\"name\":\"ExecutionFeeLinkSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"MaxDataSizeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"minSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSequenceNumber\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCCIP.RelayReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"ReportAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"addFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"needFee\",\"type\":\"bool\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"getExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutionDelaySeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutionFeeLink\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastReport\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"minSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSequenceNumber\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.RelayReport\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxDataSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"removeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"merkle\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"executionDelaySeconds\",\"type\":\"uint256\"}],\"name\":\"setExecutionDelaySeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"executionFeeLink\",\"type\":\"uint256\"}],\"name\":\"setExecutionFeeLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxDataSize\",\"type\":\"uint256\"}],\"name\":\"setMaxDataSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620058b9380380620058b983398101604081905262000034916200067b565b88888888888888888860016103e8600189888b8b8a8a33806000806000806101000a81548160ff02191690831515021790555060006001600160a01b0316826001600160a01b03161415620000d05760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b0319909216919091179091558116156200010a576200010a81620003c3565b5050506001600160a01b038216158062000122575080155b156200014157604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001875760405162d8548360e71b815260040160405180910390fd5b81516200019c90600590602085019062000475565b5060005b825181101562000280576000828281518110620001c157620001c162000751565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b0316815250600460008685815181106200020b576200020b62000751565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff1916600117905580620002778162000767565b915050620001a0565b5050508051825114620002a65760405163ee9d106b60e01b815260040160405180910390fd5b8151620002bb90600890602085019062000475565b5060005b825181101562000388576000828281518110620002e057620002e062000751565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b0316815250600760008685815181106200032a576200032a62000751565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b029290911691909117905550806200037f8162000767565b915050620002bf565b505050151560805260a09a909a5260c09890985260129190915560175550505060119290925550506013555062000791975050505050505050565b6001600160a01b0381163314156200041e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000c7565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620004cd579160200282015b82811115620004cd57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000496565b50620004db929150620004df565b5090565b5b80821115620004db5760008155600101620004e0565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620005375762000537620004f6565b604052919050565b60006001600160401b038211156200055b576200055b620004f6565b5060051b60200190565b6001600160a01b03811681146200057b57600080fd5b50565b600082601f8301126200059057600080fd5b81516020620005a9620005a3836200053f565b6200050c565b82815260059290921b84018101918181019086841115620005c957600080fd5b8286015b84811015620005f1578051620005e38162000565565b8352918301918301620005cd565b509695505050505050565b600082601f8301126200060e57600080fd5b8151602062000621620005a3836200053f565b82815260059290921b840181019181810190868411156200064157600080fd5b8286015b84811015620005f15780516200065b8162000565565b835291830191830162000645565b8051620006768162000565565b919050565b60008060008060008060008060006101208a8c0312156200069b57600080fd5b895160208b015160408c0151919a5098506001600160401b0380821115620006c257600080fd5b620006d08d838e016200057e565b985060608c0151915080821115620006e757600080fd5b620006f58d838e016200057e565b975060808c01519150808211156200070c57600080fd5b506200071b8c828d01620005fc565b9550506200072c60a08b0162000669565b935060c08a0151925060e08a015191506101008a015190509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b60006000198214156200078a57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c0516150f1620007c8600039600061053601526000818161049f015261342e0152600061213b01526150f16000f3fe608060405234801561001057600080fd5b50600436106102c85760003560e01c806381ff70481161017b578063b034909c116100d8578063c15d605b1161008c578063eb511dd411610071578063eb511dd414610691578063eefa7a3e146106a4578063f2fde38b146106fa57600080fd5b8063c15d605b1461066b578063e3d0e7121461067e57600080fd5b8063b5767166116100bd578063b57671661461060c578063b6608c3b1461061f578063bbe4f6db1461063257600080fd5b8063b034909c146105f1578063b1dc65a4146105f957600080fd5b806389e19b231161012f578063975de62b11610114578063975de62b1461059e578063a7206cd6146105b1578063afcb95d7146105d157600080fd5b806389e19b23146105735780638da5cb5b1461057b57600080fd5b806385e1f4d01161016057806385e1f4d014610531578063879b32c51461055857806389c065681461056b57600080fd5b806381ff7048146104f95780638456cb591461052957600080fd5b80635853c6271161022957806374be2150116101dd57806379ba5097116101c257806379ba5097146104d457806381411834146104dc57806381be8fa4146104f157600080fd5b806374be21501461049a578063768c577b146104c157600080fd5b80635b16ebb71161020e5780635b16ebb7146104435780635c975abb1461047c578063744b92e21461048757600080fd5b80635853c6271461041d57806359e96b5b1461043057600080fd5b806321947507116102805780632b898c25116102655780632b898c25146103ef5780633b8d08ef146104025780633f4ba83a1461041557600080fd5b8063219475071461039e5780632222dd42146103d157600080fd5b806316b8e731116102b157806316b8e731146102f9578063181f5a77146103575780631a830a1e1461039657600080fd5b806307db2b21146102cd578063108ee5fc146102e4575b600080fd5b6011545b6040519081526020015b60405180910390f35b6102f76102f236600461412f565b61070d565b005b61033261030736600461412f565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102db565b604080518082018252600d81527f4f666652616d7020302e302e3100000000000000000000000000000000000000602082015290516102db91906141c6565b6012546102d1565b6103c16103ac3660046141d9565b60009081526010602052604090205460ff1690565b60405190151581526020016102db565b60025473ffffffffffffffffffffffffffffffffffffffff16610332565b6102f76103fd3660046141f2565b6107e9565b6102f76104103660046146b8565b610bb5565b6102f76114dd565b6102f761042b3660046141f2565b6114ef565b6102f761043e366004614730565b611707565b6103c161045136600461412f565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff166103c1565b6102f76104953660046141f2565b611785565b6102d17f000000000000000000000000000000000000000000000000000000000000000081565b6102d16104cf366004614771565b611b76565b6102f7611d02565b6104e4611e24565b6040516102db9190614826565b6104e4611e93565b600b546009546040805163ffffffff808516825264010000000090940490931660208401528201526060016102db565b6102f7611f00565b6102d17f000000000000000000000000000000000000000000000000000000000000000081565b6102f76105663660046141d9565b611f10565b6104e4611f54565b6013546102d1565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610332565b6102f76105ac3660046141d9565b611fc1565b6102d16105bf3660046141d9565b6000908152600f602052604090205490565b6040805160018152600060208201819052918101919091526060016102db565b6003546102d1565b6102f7610607366004614885565b611ffe565b6102f761061a36600461496a565b6126a7565b6102f761062d3660046141d9565b6126b6565b61033261064036600461412f565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6102f76106793660046141d9565b612733565b6102f761068c3660046149d0565b612770565b6102f761069f3660046141f2565b613155565b6040805160608082018352600080835260208084018290529284015282518082018452601454808252601554828501908152601654928601928352855191825251938101939093525192820192909252016102db565b6102f761070836600461412f565b613395565b6107156133a6565b73ffffffffffffffffffffffffffffffffffffffff8116610762576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b6107f16133a6565b6008548061082b576040517f2e70248b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906108c6576040517f3917193900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff161461092f576040517f6c17b98700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600861093e600185614acc565b8154811061094e5761094e614ae3565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600882602001516bffffffffffffffffffffffff16815481106109a0576109a0614ae3565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1660086109cf600186614acc565b815481106109df576109df614ae3565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600883602001516bffffffffffffffffffffffff1681548110610a4d57610a4d614ae3565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526007909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556008805480610aef57610aef614b12565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600783526040808520949094558351908152908716918101919091527fa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d91015b60405180910390a15050505050565b60005460ff1615610c27576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610c96573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cba9190614b41565b15610cf0576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b4916004808301926060929190829003018187875af1158015610d61573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d859190614b5e565b9050600354816020015142610d9a9190614acc565b1115610dd2576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610dde8585611b76565b6000818152600f602052604090205490915080610e2b5784866040517f07e6809a000000000000000000000000000000000000000000000000000000008152600401610c1e929190614cd0565b4260125482610e3a9190614d32565b10610e71576040517f15c33ba200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855160009081526010602052604090205460ff1615610ec25785516040517f6a64e9610000000000000000000000000000000000000000000000000000000081526004810191909152602401610c1e565b60608601516080015173ffffffffffffffffffffffffffffffffffffffff1615801590610f0d575060608601516080015173ffffffffffffffffffffffffffffffffffffffff163314155b15610f4a5785516040517fd8e90b980000000000000000000000000000000000000000000000000000000081526004810191909152602401610c1e565b610f538661342c565b610f5c86613545565b8551600090815260106020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905583156111a857600080876060015160000151600081518110610fb957610fb9614ae3565b602002602001015190506000610ff48273ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff8116611043576040517f83135fec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166350d25bcd6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561108e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110b29190614d4a565b6011546110bf9190614d63565b9250828960600151602001516000815181106110dd576110dd614ae3565b602002602001018181516110f19190614acc565b90525073ffffffffffffffffffffffffffffffffffffffff808316600090815260046020526040902054166040517fea6192a20000000000000000000000000000000000000000000000000000000081523360048201526024810185905273ffffffffffffffffffffffffffffffffffffffff919091169063ea6192a290604401600060405180830381600087803b15801561118c57600080fd5b505af11580156111a0573d6000803e3d6000fd5b505050505050505b60005b6060870151515181101561136257600061120b88606001516000015183815181106111d8576111d8614ae3565b602002602001015173ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff81166112975760608801515180518390811061124057611240614ae3565b60200260200101516040517fbf16aab6000000000000000000000000000000000000000000000000000000008152600401610c1e919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b8073ffffffffffffffffffffffffffffffffffffffff1663ea6192a28960600151606001518a606001516020015185815181106112d6576112d6614ae3565b60200260200101516040518363ffffffff1660e01b815260040161131c92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600060405180830381600087803b15801561133657600080fd5b505af115801561134a573d6000803e3d6000fd5b5050505050808061135a90614da0565b9150506111ab565b50606080870151015173ffffffffffffffffffffffffffffffffffffffff163b1561148e5785606001516060015173ffffffffffffffffffffffffffffffffffffffff16639c598468876040518263ffffffff1660e01b81526004016113c89190614dd9565b600060405180830381600087803b1580156113e257600080fd5b505af19250505080156113f3575060015b61145e573d808015611421576040519150601f19603f3d011682016040523d82523d6000602084013e611426565b606091505b5086516040517f6a3fd4f2000000000000000000000000000000000000000000000000000000008152610c1e91908390600401614dec565b85516040517fc51bf0f6d90b467e0849da0ad18a4d9144a4b78b9f83202e1c65cd68f72d4a3890600090a26114d5565b606086015160a0015151156114d55785516040517fe0244be30000000000000000000000000000000000000000000000000000000081526004810191909152602401610c1e565b505050505050565b6114e56133a6565b6114ed6135f4565b565b6114f76133a6565b73ffffffffffffffffffffffffffffffffffffffff8216158061152e575073ffffffffffffffffffffffffffffffffffffffff8116155b15611565576040517fee9d106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015611601576040517f965ffb7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600880546bffffffffffffffffffffffff908116602080870191825288861660008181526007835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821790558351908152908101919091527f037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb60130591015b60405180910390a1505050565b61170f6133a6565b61173073ffffffffffffffffffffffffffffffffffffffff841683836136d5565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8906060016116fa565b61178d6133a6565b600554806117c7576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290611862576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff16146118cb576040517fd428911900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060056118da600185614acc565b815481106118ea576118ea614ae3565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff168154811061193c5761193c614ae3565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16600561196b600186614acc565b8154811061197b5761197b614ae3565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff16815481106119e9576119e9614ae3565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480611a8b57611a8b614b12565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101610ba6565b600080600060f81b84604051602001611b8f9190614dd9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052611bcb9291602001614e05565b60405160208183030381529060405280519060200120905060005b835151811015611cfa57600084600001518281518110611c0857611c08614ae3565b6020026020010151905060028560200151611c239190614e7c565b611c7e576040517f010000000000000000000000000000000000000000000000000000000000000060208201526021810184905260418101829052606101604051602081830303815290604052805190602001209250611cd1565b6040517f0100000000000000000000000000000000000000000000000000000000000000602082015260218101829052604181018490526061016040516020818303038152906040528051906020012092505b60028560200151611ce29190614e90565b60208601525080611cf281614da0565b915050611be6565b509392505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611d83576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610c1e565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600e805480602002602001604051908101604052809291908181526020018280548015611e8957602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611e5e575b5050505050905090565b60606008805480602002602001604051908101604052809291908181526020018280548015611e895760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611e5e575050505050905090565b611f086133a6565b6114ed613767565b611f186133a6565b60118190556040518181527fd7fd8173fe9fe63cbd28468ae5012d7f8db1f1550fbfe8c22a1b49e4b0b703ae906020015b60405180910390a150565b60606005805480602002602001604051908101604052809291908181526020018280548015611e895760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611e5e575050505050905090565b611fc96133a6565b60138190556040518181527f15e8871b18d0879e199f8794daad195ca64bbb7341d01951a980a1abf0800bfb90602001611f49565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161205491849163ffffffff851691908e908e908190840183828082843760009201919091525061382792505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600954808252600a5460ff80821660208501526101009091041692820192909252908314612129576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610c1e565b6121378b8b8b8b8b8b613b98565b60007f000000000000000000000000000000000000000000000000000000000000000015612194576002826020015183604001516121759190614ea4565b61217f9190614ec9565b61218a906001614ea4565b60ff1690506121aa565b60208201516121a4906001614ea4565b60ff1690505b888114612213576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610c1e565b88871461227c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610c1e565b336000908152600c602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156122bf576122bf614eeb565b60028111156122d0576122d0614eeb565b90525090506002816020015160028111156122ed576122ed614eeb565b1480156123345750600e816000015160ff168154811061230f5761230f614ae3565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b61239a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610c1e565b5050505050600088886040516123b1929190614f1a565b6040519081900381206123c8918c90602001614f2a565b6040516020818303038152906040528051906020012090506123e86140ee565b604080518082019091526000808252602082015260005b8881101561268557600060018588846020811061241e5761241e614ae3565b61242b91901a601b614ea4565b8d8d8681811061243d5761243d614ae3565b905060200201358c8c8781811061245657612456614ae3565b9050602002013560405160008152602001604052604051612493949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156124b5573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600c602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561253557612535614eeb565b600281111561254657612546614eeb565b905250925060018360200151600281111561256357612563614eeb565b146125ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610c1e565b8251849060ff16601f81106125e1576125e1614ae3565b60200201511561264d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610c1e565b600184846000015160ff16601f811061266857612668614ae3565b91151560209092020152508061267d81614da0565b9150506123ff565b5050505063ffffffff811061269c5761269c614f46565b505050505050505050565b6126b360008083613827565b50565b6126be6133a6565b806126f5576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c25191016107dd565b61273b6133a6565b60128190556040518181527ffbb92e09fd0b7bcabafe655657e745c10bcd9812c1e32e205e477a3c5ff74fc090602001611f49565b855185518560ff16601f8311156127e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610c1e565b6000811161284d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610c1e565b8183146128db576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610c1e565b6128e6816003614d63565b831161294e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610c1e565b6129566133a6565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600d5415612b4957600d546000906129ae90600190614acc565b90506000600d82815481106129c5576129c5614ae3565b6000918252602082200154600e805473ffffffffffffffffffffffffffffffffffffffff909216935090849081106129ff576129ff614ae3565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600c909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600d80549192509080612a7f57612a7f614b12565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600e805480612ae857612ae8614b12565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550612994915050565b60005b815151811015612fb0576000600c600084600001518481518110612b7257612b72614ae3565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612bbc57612bbc614eeb565b14612c23576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610c1e565b6040805180820190915260ff821681526001602082015282518051600c9160009185908110612c5457612c54614ae3565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612cf557612cf5614eeb565b021790555060009150612d059050565b600c600084602001518481518110612d1f57612d1f614ae3565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612d6957612d69614eeb565b14612dd0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610c1e565b6040805180820190915260ff82168152602081016002815250600c600084602001518481518110612e0357612e03614ae3565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612ea457612ea4614eeb565b02179055505082518051600d925083908110612ec257612ec2614ae3565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600e919083908110612f3e57612f3e614ae3565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691909117905580612fa881614da0565b915050612b4c565b506040810151600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600b80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092613042928692908216911617614f75565b92506101000a81548163ffffffff021916908363ffffffff1602179055506130a14630600b60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613c4f565b600981905582518051600a805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600b5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598613140988b98919763ffffffff909216969095919491939192614f9d565b60405180910390a15050505050505050505050565b61315d6133a6565b73ffffffffffffffffffffffffffffffffffffffff82161580613194575073ffffffffffffffffffffffffffffffffffffffff8116155b156131cb576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015613267576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c91016116fa565b61339d6133a6565b6126b381613cfa565b600054610100900473ffffffffffffffffffffffffffffffffffffffff1633146114ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610c1e565b7f00000000000000000000000000000000000000000000000000000000000000008160200151146134915780602001516040517fd44bc9eb000000000000000000000000000000000000000000000000000000008152600401610c1e91815260200190565b6017546060820151515111806134b35750606081015160208101515190515114155b156134ea576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601354816060015160a001515111156126b357601354606082015160a00151516040517f8693378900000000000000000000000000000000000000000000000000000000815260048101929092526024820152604401610c1e565b606080820151015173ffffffffffffffffffffffffffffffffffffffff1630148061359c5750606080820151015173ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff165b156126b35760608082015101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610c1e565b60005460ff16613660576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610c1e565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052613762908490613df6565b505050565b60005460ff16156137d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c1e565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586136ab3390565b60005460ff1615613894576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c1e565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015613903573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139279190614b41565b1561395d576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b4916004808301926060929190829003018187875af11580156139ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139f29190614b5e565b9050600354816020015142613a079190614acc565b1115613a3f576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082806020019051810190613a559190614b5e565b9050806040015181602001511115613a99576040517f67a3824c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516060810182526014548082526015546020830152601654928201929092529015613b22576040810151613ad1906001614d32565b826020015114613b2257806040015182602001516040517fcc7f1bd0000000000000000000000000000000000000000000000000000000008152600401610c1e929190918252602082015260400190565b81516000908152600f60209081526040918290204290558351601481905581850180516015558386018051601655845192835290519282019290925290518183015290517f07d7bce06be2a7b0230e4dd0d72523c0407e82419ab1d947c5ddaf59ca36484e9181900360600190a1505050505050565b6000613ba5826020614d63565b613bb0856020614d63565b613bbc88610144614d32565b613bc69190614d32565b613bd09190614d32565b613bdb906000614d32565b9050368114613c46576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610c1e565b50505050505050565b6000808a8a8a8a8a8a8a8a8a604051602001613c7399989796959493929190615033565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116331415613d7a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610c1e565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000613e58826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16613f029092919063ffffffff16565b8051909150156137625780806020019051810190613e769190614b41565b613762576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610c1e565b6060613f118484600085613f1b565b90505b9392505050565b606082471015613fad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610c1e565b843b614015576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610c1e565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161403e91906150c8565b60006040518083038185875af1925050503d806000811461407b576040519150601f19603f3d011682016040523d82523d6000602084013e614080565b606091505b509150915061409082828661409b565b979650505050505050565b606083156140aa575081613f14565b8251156140ba5782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1e91906141c6565b604051806103e00160405280601f906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff811681146126b357600080fd5b60006020828403121561414157600080fd5b8135613f148161410d565b60005b8381101561416757818101518382015260200161414f565b83811115614176576000848401525b50505050565b6000815180845261419481602086016020860161414c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613f14602083018461417c565b6000602082840312156141eb57600080fd5b5035919050565b6000806040838503121561420557600080fd5b82356142108161410d565b915060208301356142208161410d565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff8111828210171561427d5761427d61422b565b60405290565b60405160e0810167ffffffffffffffff8111828210171561427d5761427d61422b565b6040805190810167ffffffffffffffff8111828210171561427d5761427d61422b565b6040516060810167ffffffffffffffff8111828210171561427d5761427d61422b565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156143335761433361422b565b604052919050565b80356143468161410d565b919050565b600067ffffffffffffffff8211156143655761436561422b565b5060051b60200190565b600082601f83011261438057600080fd5b813560206143956143908361434b565b6142ec565b82815260059290921b840181019181810190868411156143b457600080fd5b8286015b848110156143d85780356143cb8161410d565b83529183019183016143b8565b509695505050505050565b600082601f8301126143f457600080fd5b813560206144046143908361434b565b82815260059290921b8401810191818101908684111561442357600080fd5b8286015b848110156143d85780358352918301918301614427565b600082601f83011261444f57600080fd5b813567ffffffffffffffff8111156144695761446961422b565b61449a60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016142ec565b8181528460208386010111156144af57600080fd5b816020850160208301376000918101602001919091529392505050565b6000608082840312156144de57600080fd5b6144e661425a565b9050813581526020820135602082015260408201356145048161410d565b6040820152606082013567ffffffffffffffff8082111561452457600080fd5b9083019060e0828603121561453857600080fd5b614540614283565b82358281111561454f57600080fd5b61455b8782860161436f565b82525060208301358281111561457057600080fd5b61457c878286016143e3565b602083015250604083013560408201526145986060840161433b565b60608201526145a96080840161433b565b608082015260a0830135828111156145c057600080fd5b6145cc8782860161443e565b60a08301525060c0830135828111156145e457600080fd5b6145f08782860161443e565b60c083015250606084015250909392505050565b60006040828403121561461657600080fd5b61461e6142a6565b9050813567ffffffffffffffff81111561463757600080fd5b8201601f8101841361464857600080fd5b803560206146586143908361434b565b82815260059290921b8301810191818101908784111561467757600080fd5b938201935b838510156146955784358252938201939082019061467c565b85525093840135938301939093525092915050565b80151581146126b357600080fd5b6000806000606084860312156146cd57600080fd5b833567ffffffffffffffff808211156146e557600080fd5b6146f1878388016144cc565b9450602086013591508082111561470757600080fd5b5061471486828701614604565b9250506040840135614725816146aa565b809150509250925092565b60008060006060848603121561474557600080fd5b83356147508161410d565b925060208401356147608161410d565b929592945050506040919091013590565b6000806040838503121561478457600080fd5b823567ffffffffffffffff8082111561479c57600080fd5b6147a8868387016144cc565b935060208501359150808211156147be57600080fd5b506147cb85828601614604565b9150509250929050565b600081518084526020808501945080840160005b8381101561481b57815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016147e9565b509495945050505050565b602081526000613f1460208301846147d5565b60008083601f84011261484b57600080fd5b50813567ffffffffffffffff81111561486357600080fd5b6020830191508360208260051b850101111561487e57600080fd5b9250929050565b60008060008060008060008060e0898b0312156148a157600080fd5b606089018a8111156148b257600080fd5b8998503567ffffffffffffffff808211156148cc57600080fd5b818b0191508b601f8301126148e057600080fd5b8135818111156148ef57600080fd5b8c602082850101111561490157600080fd5b6020830199508098505060808b013591508082111561491f57600080fd5b61492b8c838d01614839565b909750955060a08b013591508082111561494457600080fd5b506149518b828c01614839565b999c989b50969995989497949560c00135949350505050565b60006020828403121561497c57600080fd5b813567ffffffffffffffff81111561499357600080fd5b61499f8482850161443e565b949350505050565b803560ff8116811461434657600080fd5b803567ffffffffffffffff8116811461434657600080fd5b60008060008060008060c087890312156149e957600080fd5b863567ffffffffffffffff80821115614a0157600080fd5b614a0d8a838b0161436f565b97506020890135915080821115614a2357600080fd5b614a2f8a838b0161436f565b9650614a3d60408a016149a7565b95506060890135915080821115614a5357600080fd5b614a5f8a838b0161443e565b9450614a6d60808a016149b8565b935060a0890135915080821115614a8357600080fd5b50614a9089828a0161443e565b9150509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015614ade57614ade614a9d565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060208284031215614b5357600080fd5b8151613f14816146aa565b600060608284031215614b7057600080fd5b614b786142c9565b8251815260208301516020820152604083015160408201528091505092915050565b8051825260006020808301518185015273ffffffffffffffffffffffffffffffffffffffff6040840151166040850152606083015160806060860152805160e06080870152614bed6101608701826147d5565b838301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80888303810160a08a0152815180845291860193506000929091908601905b80841015614c505784518252938601936001939093019290860190614c30565b50604085015160c08a0152606085015173ffffffffffffffffffffffffffffffffffffffff90811660e08b01526080860151166101008a015260a085015189820383016101208b01529550614ca5818761417c565b95505060c084015193508088860301610140890152505050614cc7828261417c565b95945050505050565b60408152600060808201845160408085015281815180845260a08601915060209350838301925060005b81811015614d1657835183529284019291840191600101614cfa565b5050828701516060860152848103838601526140908187614b9a565b60008219821115614d4557614d45614a9d565b500190565b600060208284031215614d5c57600080fd5b5051919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614d9b57614d9b614a9d565b500290565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415614dd257614dd2614a9d565b5060010190565b602081526000613f146020830184614b9a565b828152604060208201526000613f11604083018461417c565b7fff000000000000000000000000000000000000000000000000000000000000008316815260008251614e3f81600185016020870161414c565b919091016001019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082614e8b57614e8b614e4d565b500690565b600082614e9f57614e9f614e4d565b500490565b600060ff821660ff84168060ff03821115614ec157614ec1614a9d565b019392505050565b600060ff831680614edc57614edc614e4d565b8060ff84160491505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600063ffffffff808316818516808303821115614f9457614f94614a9d565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614fcd8184018a6147d5565b90508281036080840152614fe181896147d5565b905060ff871660a084015282810360c0840152614ffe818761417c565b905067ffffffffffffffff851660e0840152828103610100840152615023818561417c565b9c9b505050505050505050505050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b16604085015281606085015261507a8285018b6147d5565b9150838203608085015261508e828a6147d5565b915060ff881660a085015283820360c08501526150ab828861417c565b90861660e08501528381036101008501529050615023818561417c565b600082516150da81846020870161414c565b919091019291505056fea164736f6c634300080c000a",
}

var OffRampHelperABI = OffRampHelperMetaData.ABI

var OffRampHelperBin = OffRampHelperMetaData.Bin

func DeployOffRampHelper(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, sourceTokens []common.Address, pools []common.Address, feeds []common.Address, afn common.Address, maxTimeWithoutAFNSignal *big.Int, executionDelaySeconds *big.Int, maxTokensLength *big.Int) (common.Address, *types.Transaction, *OffRampHelper, error) {
	parsed, err := OffRampHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OffRampHelperBin), backend, sourceChainId, chainId, sourceTokens, pools, feeds, afn, maxTimeWithoutAFNSignal, executionDelaySeconds, maxTokensLength)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OffRampHelper{OffRampHelperCaller: OffRampHelperCaller{contract: contract}, OffRampHelperTransactor: OffRampHelperTransactor{contract: contract}, OffRampHelperFilterer: OffRampHelperFilterer{contract: contract}}, nil
}

type OffRampHelper struct {
	address common.Address
	abi     abi.ABI
	OffRampHelperCaller
	OffRampHelperTransactor
	OffRampHelperFilterer
}

type OffRampHelperCaller struct {
	contract *bind.BoundContract
}

type OffRampHelperTransactor struct {
	contract *bind.BoundContract
}

type OffRampHelperFilterer struct {
	contract *bind.BoundContract
}

type OffRampHelperSession struct {
	Contract     *OffRampHelper
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OffRampHelperCallerSession struct {
	Contract *OffRampHelperCaller
	CallOpts bind.CallOpts
}

type OffRampHelperTransactorSession struct {
	Contract     *OffRampHelperTransactor
	TransactOpts bind.TransactOpts
}

type OffRampHelperRaw struct {
	Contract *OffRampHelper
}

type OffRampHelperCallerRaw struct {
	Contract *OffRampHelperCaller
}

type OffRampHelperTransactorRaw struct {
	Contract *OffRampHelperTransactor
}

func NewOffRampHelper(address common.Address, backend bind.ContractBackend) (*OffRampHelper, error) {
	abi, err := abi.JSON(strings.NewReader(OffRampHelperABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOffRampHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OffRampHelper{address: address, abi: abi, OffRampHelperCaller: OffRampHelperCaller{contract: contract}, OffRampHelperTransactor: OffRampHelperTransactor{contract: contract}, OffRampHelperFilterer: OffRampHelperFilterer{contract: contract}}, nil
}

func NewOffRampHelperCaller(address common.Address, caller bind.ContractCaller) (*OffRampHelperCaller, error) {
	contract, err := bindOffRampHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OffRampHelperCaller{contract: contract}, nil
}

func NewOffRampHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*OffRampHelperTransactor, error) {
	contract, err := bindOffRampHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OffRampHelperTransactor{contract: contract}, nil
}

func NewOffRampHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*OffRampHelperFilterer, error) {
	contract, err := bindOffRampHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OffRampHelperFilterer{contract: contract}, nil
}

func bindOffRampHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OffRampHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_OffRampHelper *OffRampHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffRampHelper.Contract.OffRampHelperCaller.contract.Call(opts, result, method, params...)
}

func (_OffRampHelper *OffRampHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRampHelper.Contract.OffRampHelperTransactor.contract.Transfer(opts)
}

func (_OffRampHelper *OffRampHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffRampHelper.Contract.OffRampHelperTransactor.contract.Transact(opts, method, params...)
}

func (_OffRampHelper *OffRampHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffRampHelper.Contract.contract.Call(opts, result, method, params...)
}

func (_OffRampHelper *OffRampHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRampHelper.Contract.contract.Transfer(opts)
}

func (_OffRampHelper *OffRampHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffRampHelper.Contract.contract.Transact(opts, method, params...)
}

func (_OffRampHelper *OffRampHelperCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) CHAINID() (*big.Int, error) {
	return _OffRampHelper.Contract.CHAINID(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) CHAINID() (*big.Int, error) {
	return _OffRampHelper.Contract.CHAINID(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "SOURCE_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) SOURCECHAINID() (*big.Int, error) {
	return _OffRampHelper.Contract.SOURCECHAINID(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) SOURCECHAINID() (*big.Int, error) {
	return _OffRampHelper.Contract.SOURCECHAINID(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) GetAFN() (common.Address, error) {
	return _OffRampHelper.Contract.GetAFN(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) GetAFN() (common.Address, error) {
	return _OffRampHelper.Contract.GetAFN(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) GetExecuted(opts *bind.CallOpts, sequenceNumber *big.Int) (bool, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "getExecuted", sequenceNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) GetExecuted(sequenceNumber *big.Int) (bool, error) {
	return _OffRampHelper.Contract.GetExecuted(&_OffRampHelper.CallOpts, sequenceNumber)
}

func (_OffRampHelper *OffRampHelperCallerSession) GetExecuted(sequenceNumber *big.Int) (bool, error) {
	return _OffRampHelper.Contract.GetExecuted(&_OffRampHelper.CallOpts, sequenceNumber)
}

func (_OffRampHelper *OffRampHelperCaller) GetExecutionDelaySeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "getExecutionDelaySeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) GetExecutionDelaySeconds() (*big.Int, error) {
	return _OffRampHelper.Contract.GetExecutionDelaySeconds(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) GetExecutionDelaySeconds() (*big.Int, error) {
	return _OffRampHelper.Contract.GetExecutionDelaySeconds(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) GetExecutionFeeLink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "getExecutionFeeLink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) GetExecutionFeeLink() (*big.Int, error) {
	return _OffRampHelper.Contract.GetExecutionFeeLink(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) GetExecutionFeeLink() (*big.Int, error) {
	return _OffRampHelper.Contract.GetExecutionFeeLink(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "getFeed", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) GetFeed(token common.Address) (common.Address, error) {
	return _OffRampHelper.Contract.GetFeed(&_OffRampHelper.CallOpts, token)
}

func (_OffRampHelper *OffRampHelperCallerSession) GetFeed(token common.Address) (common.Address, error) {
	return _OffRampHelper.Contract.GetFeed(&_OffRampHelper.CallOpts, token)
}

func (_OffRampHelper *OffRampHelperCaller) GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "getFeedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) GetFeedTokens() ([]common.Address, error) {
	return _OffRampHelper.Contract.GetFeedTokens(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) GetFeedTokens() ([]common.Address, error) {
	return _OffRampHelper.Contract.GetFeedTokens(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) GetLastReport(opts *bind.CallOpts) (CCIPRelayReport, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "getLastReport")

	if err != nil {
		return *new(CCIPRelayReport), err
	}

	out0 := *abi.ConvertType(out[0], new(CCIPRelayReport)).(*CCIPRelayReport)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) GetLastReport() (CCIPRelayReport, error) {
	return _OffRampHelper.Contract.GetLastReport(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) GetLastReport() (CCIPRelayReport, error) {
	return _OffRampHelper.Contract.GetLastReport(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) GetMaxDataSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "getMaxDataSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) GetMaxDataSize() (*big.Int, error) {
	return _OffRampHelper.Contract.GetMaxDataSize(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) GetMaxDataSize() (*big.Int, error) {
	return _OffRampHelper.Contract.GetMaxDataSize(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "getMaxSecondsWithoutAFNHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _OffRampHelper.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _OffRampHelper.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) GetMerkleRoot(opts *bind.CallOpts, root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "getMerkleRoot", root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) GetMerkleRoot(root [32]byte) (*big.Int, error) {
	return _OffRampHelper.Contract.GetMerkleRoot(&_OffRampHelper.CallOpts, root)
}

func (_OffRampHelper *OffRampHelperCallerSession) GetMerkleRoot(root [32]byte) (*big.Int, error) {
	return _OffRampHelper.Contract.GetMerkleRoot(&_OffRampHelper.CallOpts, root)
}

func (_OffRampHelper *OffRampHelperCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _OffRampHelper.Contract.GetPool(&_OffRampHelper.CallOpts, sourceToken)
}

func (_OffRampHelper *OffRampHelperCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _OffRampHelper.Contract.GetPool(&_OffRampHelper.CallOpts, sourceToken)
}

func (_OffRampHelper *OffRampHelperCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) GetPoolTokens() ([]common.Address, error) {
	return _OffRampHelper.Contract.GetPoolTokens(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _OffRampHelper.Contract.GetPoolTokens(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) IsPool(addr common.Address) (bool, error) {
	return _OffRampHelper.Contract.IsPool(&_OffRampHelper.CallOpts, addr)
}

func (_OffRampHelper *OffRampHelperCallerSession) IsPool(addr common.Address) (bool, error) {
	return _OffRampHelper.Contract.IsPool(&_OffRampHelper.CallOpts, addr)
}

func (_OffRampHelper *OffRampHelperCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_OffRampHelper *OffRampHelperSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _OffRampHelper.Contract.LatestConfigDetails(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _OffRampHelper.Contract.LatestConfigDetails(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_OffRampHelper *OffRampHelperSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _OffRampHelper.Contract.LatestConfigDigestAndEpoch(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _OffRampHelper.Contract.LatestConfigDigestAndEpoch(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) MerkleRoot(opts *bind.CallOpts, message CCIPMessage, proof CCIPMerkleProof) ([32]byte, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "merkleRoot", message, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) MerkleRoot(message CCIPMessage, proof CCIPMerkleProof) ([32]byte, error) {
	return _OffRampHelper.Contract.MerkleRoot(&_OffRampHelper.CallOpts, message, proof)
}

func (_OffRampHelper *OffRampHelperCallerSession) MerkleRoot(message CCIPMessage, proof CCIPMerkleProof) ([32]byte, error) {
	return _OffRampHelper.Contract.MerkleRoot(&_OffRampHelper.CallOpts, message, proof)
}

func (_OffRampHelper *OffRampHelperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) Owner() (common.Address, error) {
	return _OffRampHelper.Contract.Owner(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) Owner() (common.Address, error) {
	return _OffRampHelper.Contract.Owner(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) Paused() (bool, error) {
	return _OffRampHelper.Contract.Paused(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) Paused() (bool, error) {
	return _OffRampHelper.Contract.Paused(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) Transmitters() ([]common.Address, error) {
	return _OffRampHelper.Contract.Transmitters(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) Transmitters() ([]common.Address, error) {
	return _OffRampHelper.Contract.Transmitters(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) TypeAndVersion() (string, error) {
	return _OffRampHelper.Contract.TypeAndVersion(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) TypeAndVersion() (string, error) {
	return _OffRampHelper.Contract.TypeAndVersion(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "acceptOwnership")
}

func (_OffRampHelper *OffRampHelperSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffRampHelper.Contract.AcceptOwnership(&_OffRampHelper.TransactOpts)
}

func (_OffRampHelper *OffRampHelperTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffRampHelper.Contract.AcceptOwnership(&_OffRampHelper.TransactOpts)
}

func (_OffRampHelper *OffRampHelperTransactor) AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "addFeed", token, feed)
}

func (_OffRampHelper *OffRampHelperSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OffRampHelper.Contract.AddFeed(&_OffRampHelper.TransactOpts, token, feed)
}

func (_OffRampHelper *OffRampHelperTransactorSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OffRampHelper.Contract.AddFeed(&_OffRampHelper.TransactOpts, token, feed)
}

func (_OffRampHelper *OffRampHelperTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "addPool", token, pool)
}

func (_OffRampHelper *OffRampHelperSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OffRampHelper.Contract.AddPool(&_OffRampHelper.TransactOpts, token, pool)
}

func (_OffRampHelper *OffRampHelperTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OffRampHelper.Contract.AddPool(&_OffRampHelper.TransactOpts, token, pool)
}

func (_OffRampHelper *OffRampHelperTransactor) ExecuteTransaction(opts *bind.TransactOpts, message CCIPMessage, proof CCIPMerkleProof, needFee bool) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "executeTransaction", message, proof, needFee)
}

func (_OffRampHelper *OffRampHelperSession) ExecuteTransaction(message CCIPMessage, proof CCIPMerkleProof, needFee bool) (*types.Transaction, error) {
	return _OffRampHelper.Contract.ExecuteTransaction(&_OffRampHelper.TransactOpts, message, proof, needFee)
}

func (_OffRampHelper *OffRampHelperTransactorSession) ExecuteTransaction(message CCIPMessage, proof CCIPMerkleProof, needFee bool) (*types.Transaction, error) {
	return _OffRampHelper.Contract.ExecuteTransaction(&_OffRampHelper.TransactOpts, message, proof, needFee)
}

func (_OffRampHelper *OffRampHelperTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "pause")
}

func (_OffRampHelper *OffRampHelperSession) Pause() (*types.Transaction, error) {
	return _OffRampHelper.Contract.Pause(&_OffRampHelper.TransactOpts)
}

func (_OffRampHelper *OffRampHelperTransactorSession) Pause() (*types.Transaction, error) {
	return _OffRampHelper.Contract.Pause(&_OffRampHelper.TransactOpts)
}

func (_OffRampHelper *OffRampHelperTransactor) RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "removeFeed", token, feed)
}

func (_OffRampHelper *OffRampHelperSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OffRampHelper.Contract.RemoveFeed(&_OffRampHelper.TransactOpts, token, feed)
}

func (_OffRampHelper *OffRampHelperTransactorSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OffRampHelper.Contract.RemoveFeed(&_OffRampHelper.TransactOpts, token, feed)
}

func (_OffRampHelper *OffRampHelperTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "removePool", token, pool)
}

func (_OffRampHelper *OffRampHelperSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OffRampHelper.Contract.RemovePool(&_OffRampHelper.TransactOpts, token, pool)
}

func (_OffRampHelper *OffRampHelperTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OffRampHelper.Contract.RemovePool(&_OffRampHelper.TransactOpts, token, pool)
}

func (_OffRampHelper *OffRampHelperTransactor) Report(opts *bind.TransactOpts, merkle []byte) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "report", merkle)
}

func (_OffRampHelper *OffRampHelperSession) Report(merkle []byte) (*types.Transaction, error) {
	return _OffRampHelper.Contract.Report(&_OffRampHelper.TransactOpts, merkle)
}

func (_OffRampHelper *OffRampHelperTransactorSession) Report(merkle []byte) (*types.Transaction, error) {
	return _OffRampHelper.Contract.Report(&_OffRampHelper.TransactOpts, merkle)
}

func (_OffRampHelper *OffRampHelperTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "setAFN", afn)
}

func (_OffRampHelper *OffRampHelperSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetAFN(&_OffRampHelper.TransactOpts, afn)
}

func (_OffRampHelper *OffRampHelperTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetAFN(&_OffRampHelper.TransactOpts, afn)
}

func (_OffRampHelper *OffRampHelperTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_OffRampHelper *OffRampHelperSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetConfig(&_OffRampHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_OffRampHelper *OffRampHelperTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetConfig(&_OffRampHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_OffRampHelper *OffRampHelperTransactor) SetExecutionDelaySeconds(opts *bind.TransactOpts, executionDelaySeconds *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "setExecutionDelaySeconds", executionDelaySeconds)
}

func (_OffRampHelper *OffRampHelperSession) SetExecutionDelaySeconds(executionDelaySeconds *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetExecutionDelaySeconds(&_OffRampHelper.TransactOpts, executionDelaySeconds)
}

func (_OffRampHelper *OffRampHelperTransactorSession) SetExecutionDelaySeconds(executionDelaySeconds *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetExecutionDelaySeconds(&_OffRampHelper.TransactOpts, executionDelaySeconds)
}

func (_OffRampHelper *OffRampHelperTransactor) SetExecutionFeeLink(opts *bind.TransactOpts, executionFeeLink *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "setExecutionFeeLink", executionFeeLink)
}

func (_OffRampHelper *OffRampHelperSession) SetExecutionFeeLink(executionFeeLink *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetExecutionFeeLink(&_OffRampHelper.TransactOpts, executionFeeLink)
}

func (_OffRampHelper *OffRampHelperTransactorSession) SetExecutionFeeLink(executionFeeLink *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetExecutionFeeLink(&_OffRampHelper.TransactOpts, executionFeeLink)
}

func (_OffRampHelper *OffRampHelperTransactor) SetMaxDataSize(opts *bind.TransactOpts, maxDataSize *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "setMaxDataSize", maxDataSize)
}

func (_OffRampHelper *OffRampHelperSession) SetMaxDataSize(maxDataSize *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetMaxDataSize(&_OffRampHelper.TransactOpts, maxDataSize)
}

func (_OffRampHelper *OffRampHelperTransactorSession) SetMaxDataSize(maxDataSize *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetMaxDataSize(&_OffRampHelper.TransactOpts, maxDataSize)
}

func (_OffRampHelper *OffRampHelperTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_OffRampHelper *OffRampHelperSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_OffRampHelper.TransactOpts, newTime)
}

func (_OffRampHelper *OffRampHelperTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_OffRampHelper.TransactOpts, newTime)
}

func (_OffRampHelper *OffRampHelperTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "transferOwnership", to)
}

func (_OffRampHelper *OffRampHelperSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OffRampHelper.Contract.TransferOwnership(&_OffRampHelper.TransactOpts, to)
}

func (_OffRampHelper *OffRampHelperTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OffRampHelper.Contract.TransferOwnership(&_OffRampHelper.TransactOpts, to)
}

func (_OffRampHelper *OffRampHelperTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_OffRampHelper *OffRampHelperSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffRampHelper.Contract.Transmit(&_OffRampHelper.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_OffRampHelper *OffRampHelperTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffRampHelper.Contract.Transmit(&_OffRampHelper.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_OffRampHelper *OffRampHelperTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "unpause")
}

func (_OffRampHelper *OffRampHelperSession) Unpause() (*types.Transaction, error) {
	return _OffRampHelper.Contract.Unpause(&_OffRampHelper.TransactOpts)
}

func (_OffRampHelper *OffRampHelperTransactorSession) Unpause() (*types.Transaction, error) {
	return _OffRampHelper.Contract.Unpause(&_OffRampHelper.TransactOpts)
}

func (_OffRampHelper *OffRampHelperTransactor) WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "withdrawAccumulatedFees", feeToken, recipient, amount)
}

func (_OffRampHelper *OffRampHelperSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.Contract.WithdrawAccumulatedFees(&_OffRampHelper.TransactOpts, feeToken, recipient, amount)
}

func (_OffRampHelper *OffRampHelperTransactorSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.Contract.WithdrawAccumulatedFees(&_OffRampHelper.TransactOpts, feeToken, recipient, amount)
}

type OffRampHelperAFNMaxHeartbeatTimeSetIterator struct {
	Event *OffRampHelperAFNMaxHeartbeatTimeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperAFNMaxHeartbeatTimeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperAFNMaxHeartbeatTimeSet)
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
		it.Event = new(OffRampHelperAFNMaxHeartbeatTimeSet)
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

func (it *OffRampHelperAFNMaxHeartbeatTimeSetIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperAFNMaxHeartbeatTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperAFNMaxHeartbeatTimeSet struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*OffRampHelperAFNMaxHeartbeatTimeSetIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperAFNMaxHeartbeatTimeSetIterator{contract: _OffRampHelper.contract, event: "AFNMaxHeartbeatTimeSet", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *OffRampHelperAFNMaxHeartbeatTimeSet) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperAFNMaxHeartbeatTimeSet)
				if err := _OffRampHelper.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseAFNMaxHeartbeatTimeSet(log types.Log) (*OffRampHelperAFNMaxHeartbeatTimeSet, error) {
	event := new(OffRampHelperAFNMaxHeartbeatTimeSet)
	if err := _OffRampHelper.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperAFNSetIterator struct {
	Event *OffRampHelperAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperAFNSet)
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
		it.Event = new(OffRampHelperAFNSet)
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

func (it *OffRampHelperAFNSetIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterAFNSet(opts *bind.FilterOpts) (*OffRampHelperAFNSetIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperAFNSetIterator{contract: _OffRampHelper.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *OffRampHelperAFNSet) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperAFNSet)
				if err := _OffRampHelper.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseAFNSet(log types.Log) (*OffRampHelperAFNSet, error) {
	event := new(OffRampHelperAFNSet)
	if err := _OffRampHelper.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperConfigSetIterator struct {
	Event *OffRampHelperConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperConfigSet)
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
		it.Event = new(OffRampHelperConfigSet)
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

func (it *OffRampHelperConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperConfigSet struct {
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

func (_OffRampHelper *OffRampHelperFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OffRampHelperConfigSetIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperConfigSetIterator{contract: _OffRampHelper.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampHelperConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperConfigSet)
				if err := _OffRampHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseConfigSet(log types.Log) (*OffRampHelperConfigSet, error) {
	event := new(OffRampHelperConfigSet)
	if err := _OffRampHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperCrossChainMessageExecutedIterator struct {
	Event *OffRampHelperCrossChainMessageExecuted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperCrossChainMessageExecutedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperCrossChainMessageExecuted)
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
		it.Event = new(OffRampHelperCrossChainMessageExecuted)
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

func (it *OffRampHelperCrossChainMessageExecutedIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperCrossChainMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperCrossChainMessageExecuted struct {
	SequenceNumber *big.Int
	Raw            types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterCrossChainMessageExecuted(opts *bind.FilterOpts, sequenceNumber []*big.Int) (*OffRampHelperCrossChainMessageExecutedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "CrossChainMessageExecuted", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &OffRampHelperCrossChainMessageExecutedIterator{contract: _OffRampHelper.contract, event: "CrossChainMessageExecuted", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchCrossChainMessageExecuted(opts *bind.WatchOpts, sink chan<- *OffRampHelperCrossChainMessageExecuted, sequenceNumber []*big.Int) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "CrossChainMessageExecuted", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperCrossChainMessageExecuted)
				if err := _OffRampHelper.contract.UnpackLog(event, "CrossChainMessageExecuted", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseCrossChainMessageExecuted(log types.Log) (*OffRampHelperCrossChainMessageExecuted, error) {
	event := new(OffRampHelperCrossChainMessageExecuted)
	if err := _OffRampHelper.contract.UnpackLog(event, "CrossChainMessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperExecutionDelaySecondsSetIterator struct {
	Event *OffRampHelperExecutionDelaySecondsSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperExecutionDelaySecondsSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperExecutionDelaySecondsSet)
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
		it.Event = new(OffRampHelperExecutionDelaySecondsSet)
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

func (it *OffRampHelperExecutionDelaySecondsSetIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperExecutionDelaySecondsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperExecutionDelaySecondsSet struct {
	Delay *big.Int
	Raw   types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterExecutionDelaySecondsSet(opts *bind.FilterOpts) (*OffRampHelperExecutionDelaySecondsSetIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "ExecutionDelaySecondsSet")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperExecutionDelaySecondsSetIterator{contract: _OffRampHelper.contract, event: "ExecutionDelaySecondsSet", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchExecutionDelaySecondsSet(opts *bind.WatchOpts, sink chan<- *OffRampHelperExecutionDelaySecondsSet) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "ExecutionDelaySecondsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperExecutionDelaySecondsSet)
				if err := _OffRampHelper.contract.UnpackLog(event, "ExecutionDelaySecondsSet", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseExecutionDelaySecondsSet(log types.Log) (*OffRampHelperExecutionDelaySecondsSet, error) {
	event := new(OffRampHelperExecutionDelaySecondsSet)
	if err := _OffRampHelper.contract.UnpackLog(event, "ExecutionDelaySecondsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperExecutionFeeLinkSetIterator struct {
	Event *OffRampHelperExecutionFeeLinkSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperExecutionFeeLinkSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperExecutionFeeLinkSet)
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
		it.Event = new(OffRampHelperExecutionFeeLinkSet)
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

func (it *OffRampHelperExecutionFeeLinkSetIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperExecutionFeeLinkSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperExecutionFeeLinkSet struct {
	ExecutionFee *big.Int
	Raw          types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterExecutionFeeLinkSet(opts *bind.FilterOpts) (*OffRampHelperExecutionFeeLinkSetIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "ExecutionFeeLinkSet")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperExecutionFeeLinkSetIterator{contract: _OffRampHelper.contract, event: "ExecutionFeeLinkSet", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchExecutionFeeLinkSet(opts *bind.WatchOpts, sink chan<- *OffRampHelperExecutionFeeLinkSet) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "ExecutionFeeLinkSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperExecutionFeeLinkSet)
				if err := _OffRampHelper.contract.UnpackLog(event, "ExecutionFeeLinkSet", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseExecutionFeeLinkSet(log types.Log) (*OffRampHelperExecutionFeeLinkSet, error) {
	event := new(OffRampHelperExecutionFeeLinkSet)
	if err := _OffRampHelper.contract.UnpackLog(event, "ExecutionFeeLinkSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperFeedAddedIterator struct {
	Event *OffRampHelperFeedAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperFeedAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperFeedAdded)
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
		it.Event = new(OffRampHelperFeedAdded)
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

func (it *OffRampHelperFeedAddedIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperFeedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperFeedAdded struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterFeedAdded(opts *bind.FilterOpts) (*OffRampHelperFeedAddedIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperFeedAddedIterator{contract: _OffRampHelper.contract, event: "FeedAdded", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *OffRampHelperFeedAdded) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperFeedAdded)
				if err := _OffRampHelper.contract.UnpackLog(event, "FeedAdded", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseFeedAdded(log types.Log) (*OffRampHelperFeedAdded, error) {
	event := new(OffRampHelperFeedAdded)
	if err := _OffRampHelper.contract.UnpackLog(event, "FeedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperFeedRemovedIterator struct {
	Event *OffRampHelperFeedRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperFeedRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperFeedRemoved)
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
		it.Event = new(OffRampHelperFeedRemoved)
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

func (it *OffRampHelperFeedRemovedIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperFeedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperFeedRemoved struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterFeedRemoved(opts *bind.FilterOpts) (*OffRampHelperFeedRemovedIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperFeedRemovedIterator{contract: _OffRampHelper.contract, event: "FeedRemoved", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *OffRampHelperFeedRemoved) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperFeedRemoved)
				if err := _OffRampHelper.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseFeedRemoved(log types.Log) (*OffRampHelperFeedRemoved, error) {
	event := new(OffRampHelperFeedRemoved)
	if err := _OffRampHelper.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperFeesWithdrawnIterator struct {
	Event *OffRampHelperFeesWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperFeesWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperFeesWithdrawn)
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
		it.Event = new(OffRampHelperFeesWithdrawn)
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

func (it *OffRampHelperFeesWithdrawnIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperFeesWithdrawn struct {
	FeeToken  common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*OffRampHelperFeesWithdrawnIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperFeesWithdrawnIterator{contract: _OffRampHelper.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *OffRampHelperFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperFeesWithdrawn)
				if err := _OffRampHelper.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseFeesWithdrawn(log types.Log) (*OffRampHelperFeesWithdrawn, error) {
	event := new(OffRampHelperFeesWithdrawn)
	if err := _OffRampHelper.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperMaxDataSizeSetIterator struct {
	Event *OffRampHelperMaxDataSizeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperMaxDataSizeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperMaxDataSizeSet)
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
		it.Event = new(OffRampHelperMaxDataSizeSet)
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

func (it *OffRampHelperMaxDataSizeSetIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperMaxDataSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperMaxDataSizeSet struct {
	Size *big.Int
	Raw  types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterMaxDataSizeSet(opts *bind.FilterOpts) (*OffRampHelperMaxDataSizeSetIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "MaxDataSizeSet")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperMaxDataSizeSetIterator{contract: _OffRampHelper.contract, event: "MaxDataSizeSet", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchMaxDataSizeSet(opts *bind.WatchOpts, sink chan<- *OffRampHelperMaxDataSizeSet) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "MaxDataSizeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperMaxDataSizeSet)
				if err := _OffRampHelper.contract.UnpackLog(event, "MaxDataSizeSet", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseMaxDataSizeSet(log types.Log) (*OffRampHelperMaxDataSizeSet, error) {
	event := new(OffRampHelperMaxDataSizeSet)
	if err := _OffRampHelper.contract.UnpackLog(event, "MaxDataSizeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperOwnershipTransferRequestedIterator struct {
	Event *OffRampHelperOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperOwnershipTransferRequested)
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
		it.Event = new(OffRampHelperOwnershipTransferRequested)
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

func (it *OffRampHelperOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampHelperOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffRampHelperOwnershipTransferRequestedIterator{contract: _OffRampHelper.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffRampHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperOwnershipTransferRequested)
				if err := _OffRampHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseOwnershipTransferRequested(log types.Log) (*OffRampHelperOwnershipTransferRequested, error) {
	event := new(OffRampHelperOwnershipTransferRequested)
	if err := _OffRampHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperOwnershipTransferredIterator struct {
	Event *OffRampHelperOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperOwnershipTransferred)
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
		it.Event = new(OffRampHelperOwnershipTransferred)
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

func (it *OffRampHelperOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampHelperOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffRampHelperOwnershipTransferredIterator{contract: _OffRampHelper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffRampHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperOwnershipTransferred)
				if err := _OffRampHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseOwnershipTransferred(log types.Log) (*OffRampHelperOwnershipTransferred, error) {
	event := new(OffRampHelperOwnershipTransferred)
	if err := _OffRampHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperPausedIterator struct {
	Event *OffRampHelperPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperPaused)
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
		it.Event = new(OffRampHelperPaused)
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

func (it *OffRampHelperPausedIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterPaused(opts *bind.FilterOpts) (*OffRampHelperPausedIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperPausedIterator{contract: _OffRampHelper.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OffRampHelperPaused) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperPaused)
				if err := _OffRampHelper.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParsePaused(log types.Log) (*OffRampHelperPaused, error) {
	event := new(OffRampHelperPaused)
	if err := _OffRampHelper.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperPoolAddedIterator struct {
	Event *OffRampHelperPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperPoolAdded)
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
		it.Event = new(OffRampHelperPoolAdded)
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

func (it *OffRampHelperPoolAddedIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*OffRampHelperPoolAddedIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperPoolAddedIterator{contract: _OffRampHelper.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *OffRampHelperPoolAdded) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperPoolAdded)
				if err := _OffRampHelper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParsePoolAdded(log types.Log) (*OffRampHelperPoolAdded, error) {
	event := new(OffRampHelperPoolAdded)
	if err := _OffRampHelper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperPoolRemovedIterator struct {
	Event *OffRampHelperPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperPoolRemoved)
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
		it.Event = new(OffRampHelperPoolRemoved)
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

func (it *OffRampHelperPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*OffRampHelperPoolRemovedIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperPoolRemovedIterator{contract: _OffRampHelper.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *OffRampHelperPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperPoolRemoved)
				if err := _OffRampHelper.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParsePoolRemoved(log types.Log) (*OffRampHelperPoolRemoved, error) {
	event := new(OffRampHelperPoolRemoved)
	if err := _OffRampHelper.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperReportAcceptedIterator struct {
	Event *OffRampHelperReportAccepted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperReportAcceptedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperReportAccepted)
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
		it.Event = new(OffRampHelperReportAccepted)
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

func (it *OffRampHelperReportAcceptedIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperReportAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperReportAccepted struct {
	Report CCIPRelayReport
	Raw    types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterReportAccepted(opts *bind.FilterOpts) (*OffRampHelperReportAcceptedIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "ReportAccepted")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperReportAcceptedIterator{contract: _OffRampHelper.contract, event: "ReportAccepted", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchReportAccepted(opts *bind.WatchOpts, sink chan<- *OffRampHelperReportAccepted) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "ReportAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperReportAccepted)
				if err := _OffRampHelper.contract.UnpackLog(event, "ReportAccepted", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseReportAccepted(log types.Log) (*OffRampHelperReportAccepted, error) {
	event := new(OffRampHelperReportAccepted)
	if err := _OffRampHelper.contract.UnpackLog(event, "ReportAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperTransmittedIterator struct {
	Event *OffRampHelperTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperTransmitted)
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
		it.Event = new(OffRampHelperTransmitted)
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

func (it *OffRampHelperTransmittedIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterTransmitted(opts *bind.FilterOpts) (*OffRampHelperTransmittedIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperTransmittedIterator{contract: _OffRampHelper.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OffRampHelperTransmitted) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperTransmitted)
				if err := _OffRampHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseTransmitted(log types.Log) (*OffRampHelperTransmitted, error) {
	event := new(OffRampHelperTransmitted)
	if err := _OffRampHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampHelperUnpausedIterator struct {
	Event *OffRampHelperUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperUnpaused)
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
		it.Event = new(OffRampHelperUnpaused)
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

func (it *OffRampHelperUnpausedIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OffRampHelperUnpausedIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperUnpausedIterator{contract: _OffRampHelper.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OffRampHelperUnpaused) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperUnpaused)
				if err := _OffRampHelper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseUnpaused(log types.Log) (*OffRampHelperUnpaused, error) {
	event := new(OffRampHelperUnpaused)
	if err := _OffRampHelper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_OffRampHelper *OffRampHelper) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OffRampHelper.abi.Events["AFNMaxHeartbeatTimeSet"].ID:
		return _OffRampHelper.ParseAFNMaxHeartbeatTimeSet(log)
	case _OffRampHelper.abi.Events["AFNSet"].ID:
		return _OffRampHelper.ParseAFNSet(log)
	case _OffRampHelper.abi.Events["ConfigSet"].ID:
		return _OffRampHelper.ParseConfigSet(log)
	case _OffRampHelper.abi.Events["CrossChainMessageExecuted"].ID:
		return _OffRampHelper.ParseCrossChainMessageExecuted(log)
	case _OffRampHelper.abi.Events["ExecutionDelaySecondsSet"].ID:
		return _OffRampHelper.ParseExecutionDelaySecondsSet(log)
	case _OffRampHelper.abi.Events["ExecutionFeeLinkSet"].ID:
		return _OffRampHelper.ParseExecutionFeeLinkSet(log)
	case _OffRampHelper.abi.Events["FeedAdded"].ID:
		return _OffRampHelper.ParseFeedAdded(log)
	case _OffRampHelper.abi.Events["FeedRemoved"].ID:
		return _OffRampHelper.ParseFeedRemoved(log)
	case _OffRampHelper.abi.Events["FeesWithdrawn"].ID:
		return _OffRampHelper.ParseFeesWithdrawn(log)
	case _OffRampHelper.abi.Events["MaxDataSizeSet"].ID:
		return _OffRampHelper.ParseMaxDataSizeSet(log)
	case _OffRampHelper.abi.Events["OwnershipTransferRequested"].ID:
		return _OffRampHelper.ParseOwnershipTransferRequested(log)
	case _OffRampHelper.abi.Events["OwnershipTransferred"].ID:
		return _OffRampHelper.ParseOwnershipTransferred(log)
	case _OffRampHelper.abi.Events["Paused"].ID:
		return _OffRampHelper.ParsePaused(log)
	case _OffRampHelper.abi.Events["PoolAdded"].ID:
		return _OffRampHelper.ParsePoolAdded(log)
	case _OffRampHelper.abi.Events["PoolRemoved"].ID:
		return _OffRampHelper.ParsePoolRemoved(log)
	case _OffRampHelper.abi.Events["ReportAccepted"].ID:
		return _OffRampHelper.ParseReportAccepted(log)
	case _OffRampHelper.abi.Events["Transmitted"].ID:
		return _OffRampHelper.ParseTransmitted(log)
	case _OffRampHelper.abi.Events["Unpaused"].ID:
		return _OffRampHelper.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (OffRampHelperAFNMaxHeartbeatTimeSet) Topic() common.Hash {
	return common.HexToHash("0x72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251")
}

func (OffRampHelperAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (OffRampHelperConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (OffRampHelperCrossChainMessageExecuted) Topic() common.Hash {
	return common.HexToHash("0xc51bf0f6d90b467e0849da0ad18a4d9144a4b78b9f83202e1c65cd68f72d4a38")
}

func (OffRampHelperExecutionDelaySecondsSet) Topic() common.Hash {
	return common.HexToHash("0xfbb92e09fd0b7bcabafe655657e745c10bcd9812c1e32e205e477a3c5ff74fc0")
}

func (OffRampHelperExecutionFeeLinkSet) Topic() common.Hash {
	return common.HexToHash("0xd7fd8173fe9fe63cbd28468ae5012d7f8db1f1550fbfe8c22a1b49e4b0b703ae")
}

func (OffRampHelperFeedAdded) Topic() common.Hash {
	return common.HexToHash("0x037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb601305")
}

func (OffRampHelperFeedRemoved) Topic() common.Hash {
	return common.HexToHash("0xa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d")
}

func (OffRampHelperFeesWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8")
}

func (OffRampHelperMaxDataSizeSet) Topic() common.Hash {
	return common.HexToHash("0x15e8871b18d0879e199f8794daad195ca64bbb7341d01951a980a1abf0800bfb")
}

func (OffRampHelperOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OffRampHelperOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (OffRampHelperPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (OffRampHelperPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (OffRampHelperPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (OffRampHelperReportAccepted) Topic() common.Hash {
	return common.HexToHash("0x07d7bce06be2a7b0230e4dd0d72523c0407e82419ab1d947c5ddaf59ca36484e")
}

func (OffRampHelperTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (OffRampHelperUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_OffRampHelper *OffRampHelper) Address() common.Address {
	return _OffRampHelper.address
}

type OffRampHelperInterface interface {
	CHAINID(opts *bind.CallOpts) (*big.Int, error)

	SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetExecuted(opts *bind.CallOpts, sequenceNumber *big.Int) (bool, error)

	GetExecutionDelaySeconds(opts *bind.CallOpts) (*big.Int, error)

	GetExecutionFeeLink(opts *bind.CallOpts) (*big.Int, error)

	GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error)

	GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetLastReport(opts *bind.CallOpts) (CCIPRelayReport, error)

	GetMaxDataSize(opts *bind.CallOpts) (*big.Int, error)

	GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error)

	GetMerkleRoot(opts *bind.CallOpts, root [32]byte) (*big.Int, error)

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

	Report(opts *bind.TransactOpts, merkle []byte) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	SetExecutionDelaySeconds(opts *bind.TransactOpts, executionDelaySeconds *big.Int) (*types.Transaction, error)

	SetExecutionFeeLink(opts *bind.TransactOpts, executionFeeLink *big.Int) (*types.Transaction, error)

	SetMaxDataSize(opts *bind.TransactOpts, maxDataSize *big.Int) (*types.Transaction, error)

	SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*OffRampHelperAFNMaxHeartbeatTimeSetIterator, error)

	WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *OffRampHelperAFNMaxHeartbeatTimeSet) (event.Subscription, error)

	ParseAFNMaxHeartbeatTimeSet(log types.Log) (*OffRampHelperAFNMaxHeartbeatTimeSet, error)

	FilterAFNSet(opts *bind.FilterOpts) (*OffRampHelperAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *OffRampHelperAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*OffRampHelperAFNSet, error)

	FilterConfigSet(opts *bind.FilterOpts) (*OffRampHelperConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampHelperConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*OffRampHelperConfigSet, error)

	FilterCrossChainMessageExecuted(opts *bind.FilterOpts, sequenceNumber []*big.Int) (*OffRampHelperCrossChainMessageExecutedIterator, error)

	WatchCrossChainMessageExecuted(opts *bind.WatchOpts, sink chan<- *OffRampHelperCrossChainMessageExecuted, sequenceNumber []*big.Int) (event.Subscription, error)

	ParseCrossChainMessageExecuted(log types.Log) (*OffRampHelperCrossChainMessageExecuted, error)

	FilterExecutionDelaySecondsSet(opts *bind.FilterOpts) (*OffRampHelperExecutionDelaySecondsSetIterator, error)

	WatchExecutionDelaySecondsSet(opts *bind.WatchOpts, sink chan<- *OffRampHelperExecutionDelaySecondsSet) (event.Subscription, error)

	ParseExecutionDelaySecondsSet(log types.Log) (*OffRampHelperExecutionDelaySecondsSet, error)

	FilterExecutionFeeLinkSet(opts *bind.FilterOpts) (*OffRampHelperExecutionFeeLinkSetIterator, error)

	WatchExecutionFeeLinkSet(opts *bind.WatchOpts, sink chan<- *OffRampHelperExecutionFeeLinkSet) (event.Subscription, error)

	ParseExecutionFeeLinkSet(log types.Log) (*OffRampHelperExecutionFeeLinkSet, error)

	FilterFeedAdded(opts *bind.FilterOpts) (*OffRampHelperFeedAddedIterator, error)

	WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *OffRampHelperFeedAdded) (event.Subscription, error)

	ParseFeedAdded(log types.Log) (*OffRampHelperFeedAdded, error)

	FilterFeedRemoved(opts *bind.FilterOpts) (*OffRampHelperFeedRemovedIterator, error)

	WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *OffRampHelperFeedRemoved) (event.Subscription, error)

	ParseFeedRemoved(log types.Log) (*OffRampHelperFeedRemoved, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*OffRampHelperFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *OffRampHelperFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*OffRampHelperFeesWithdrawn, error)

	FilterMaxDataSizeSet(opts *bind.FilterOpts) (*OffRampHelperMaxDataSizeSetIterator, error)

	WatchMaxDataSizeSet(opts *bind.WatchOpts, sink chan<- *OffRampHelperMaxDataSizeSet) (event.Subscription, error)

	ParseMaxDataSizeSet(log types.Log) (*OffRampHelperMaxDataSizeSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampHelperOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffRampHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OffRampHelperOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampHelperOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffRampHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OffRampHelperOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*OffRampHelperPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *OffRampHelperPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*OffRampHelperPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*OffRampHelperPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *OffRampHelperPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*OffRampHelperPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*OffRampHelperPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *OffRampHelperPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*OffRampHelperPoolRemoved, error)

	FilterReportAccepted(opts *bind.FilterOpts) (*OffRampHelperReportAcceptedIterator, error)

	WatchReportAccepted(opts *bind.WatchOpts, sink chan<- *OffRampHelperReportAccepted) (event.Subscription, error)

	ParseReportAccepted(log types.Log) (*OffRampHelperReportAccepted, error)

	FilterTransmitted(opts *bind.FilterOpts) (*OffRampHelperTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OffRampHelperTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*OffRampHelperTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*OffRampHelperUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OffRampHelperUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*OffRampHelperUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
