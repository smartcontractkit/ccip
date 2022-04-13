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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"contractAggregatorV2V3Interface[]\",\"name\":\"feeds\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionDelayError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"InvalidExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceFeedConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MerkleProofError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RelayReportError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lastMaxSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMinSequenceNumber\",\"type\":\"uint256\"}],\"name\":\"SequenceError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenFeedMistmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMistmatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"UnexpectedPayloadData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"CrossChainMessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"minSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSequenceNumber\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCCIP.RelayReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"ReportAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"addFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"needFee\",\"type\":\"bool\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"getExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastReport\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"minSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSequenceNumber\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.RelayReport\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRampConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"removeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setOffRampConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005c2a38038062005c2a833981016040819052620000349162000747565b6000805460ff191681556001908790869082908990889088903390819081620000a45760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000de57620000de81620003ed565b5050506001600160a01b0382161580620000f6575080155b156200011557604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03939093169290921790915560035580518251146200015b5760405162d8548360e71b815260040160405180910390fd5b8151620001709060059060208501906200049e565b5060005b8251811015620002545760008282815181106200019557620001956200081d565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001df57620001df6200081d565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff19166001179055806200024b8162000833565b91505062000174565b50505080518251146200027a5760405163ee9d106b60e01b815260040160405180910390fd5b81516200028f9060089060208501906200049e565b5060005b82518110156200035c576000828281518110620002b457620002b46200081d565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060076000868581518110620002fe57620002fe6200081d565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b02929091169190911790555080620003538162000833565b91505062000293565b505050151560805260a09790975250505060c0929092525050805160148054602084015160408501516060909501516001600160401b03908116600160c01b026001600160c01b03968216600160801b02969096166001600160801b0392821668010000000000000000026001600160801b031990941691909516179190911716919091179190911790556200085b565b336001600160a01b03821603620004475760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016200009b565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620004f6579160200282015b82811115620004f657825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620004bf565b506200050492915062000508565b5090565b5b8082111562000504576000815560010162000509565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156200056057620005606200051f565b604052919050565b60006001600160401b038211156200058457620005846200051f565b5060051b60200190565b6001600160a01b0381168114620005a457600080fd5b50565b600082601f830112620005b957600080fd5b81516020620005d2620005cc8362000568565b62000535565b82815260059290921b84018101918181019086841115620005f257600080fd5b8286015b848110156200061a5780516200060c816200058e565b8352918301918301620005f6565b509695505050505050565b600082601f8301126200063757600080fd5b815160206200064a620005cc8362000568565b82815260059290921b840181019181810190868411156200066a57600080fd5b8286015b848110156200061a57805162000684816200058e565b83529183019183016200066e565b80516200069f816200058e565b919050565b80516001600160401b03811681146200069f57600080fd5b600060808284031215620006cf57600080fd5b604051608081016001600160401b0381118282101715620006f457620006f46200051f565b6040529050806200070583620006a4565b81526200071560208401620006a4565b60208201526200072860408401620006a4565b60408201526200073b60608401620006a4565b60608201525092915050565b600080600080600080600080610160898b0312156200076557600080fd5b885160208a015160408b015191995097506001600160401b03808211156200078c57600080fd5b6200079a8c838d01620005a7565b975060608b0151915080821115620007b157600080fd5b620007bf8c838d01620005a7565b965060808b0151915080821115620007d657600080fd5b50620007e58b828c0162000625565b945050620007f660a08a0162000692565b925060c089015191506200080e8a60e08b01620006bc565b90509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b6000600182016200085457634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c0516153986200089260003960006104ec01526000818161044701526134700152600061214801526153986000f3fe608060405234801561001057600080fd5b50600436106102775760003560e01c806381be8fa411610160578063b034909c116100d8578063c0d786551161008c578063eb511dd411610071578063eb511dd414610702578063eefa7a3e14610715578063f2fde38b1461076b57600080fd5b8063c0d78655146106dc578063e3d0e712146106ef57600080fd5b8063b1dc65a4116100bd578063b1dc65a41461067d578063b6608c3b14610690578063bbe4f6db146106a357600080fd5b8063b034909c14610657578063b0f479a11461065f57600080fd5b806389c065681161012f578063a7206cd611610114578063a7206cd614610539578063a8ebd0f414610559578063afcb95d71461063757600080fd5b806389c065681461050e5780638da5cb5b1461051657600080fd5b806381be8fa4146104a757806381ff7048146104af5780638456cb59146104df57806385e1f4d0146104e757600080fd5b80635853c627116101f3578063744b92e2116101c2578063768c577b116101a7578063768c577b1461047757806379ba50971461048a578063814118341461049257600080fd5b8063744b92e21461042f57806374be21501461044257600080fd5b80635853c627146103c557806359e96b5b146103d85780635b16ebb7146103eb5780635c975abb1461042457600080fd5b80632222dd421161024a5780633b8d08ef1161022f5780633b8d08ef146103975780633f4ba83a146103aa578063461c551b146103b257600080fd5b80632222dd42146103665780632b898c251461038457600080fd5b8063108ee5fc1461027c57806316b8e73114610291578063181f5a77146102f45780632194750714610333575b600080fd5b61028f61028a36600461424e565b61077e565b005b6102ca61029f36600461424e565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b604080518082018252600d81527f4f666652616d7020302e302e3100000000000000000000000000000000000000602082015290516102eb91906142e5565b6103566103413660046142f8565b60009081526010602052604090205460ff1690565b60405190151581526020016102eb565b60025473ffffffffffffffffffffffffffffffffffffffff166102ca565b61028f610392366004614311565b61085a565b61028f6103a53660046147d2565b610c2a565b61028f611512565b61028f6103c036600461484a565b611524565b61028f6103d3366004614311565b611576565b61028f6103e6366004614862565b61178e565b6103566103f936600461424e565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff16610356565b61028f61043d366004614311565b61180c565b6104697f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016102eb565b6104696104853660046148a3565b611c01565b61028f611d90565b61049a611eb2565b6040516102eb9190614958565b61049a611f21565b600b546009546040805163ffffffff808516825264010000000090940490931660208401528201526060016102eb565b61028f611f8e565b6104697f000000000000000000000000000000000000000000000000000000000000000081565b61049a611f9e565b600054610100900473ffffffffffffffffffffffffffffffffffffffff166102ca565b6104696105473660046142f8565b6000908152600f602052604090205490565b6105f3604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260145467ffffffffffffffff808216835268010000000000000000820481166020840152700100000000000000000000000000000000820481169383019390935278010000000000000000000000000000000000000000000000009004909116606082015290565b6040516102eb9190815167ffffffffffffffff9081168252602080840151821690830152604080840151821690830152606092830151169181019190915260800190565b6040805160018152600060208201819052918101919091526060016102eb565b600354610469565b60155473ffffffffffffffffffffffffffffffffffffffff166102ca565b61028f61068b3660046149b7565b61200b565b61028f61069e3660046142f8565b6126b4565b6102ca6106b136600461424e565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b61028f6106ea36600461424e565b612734565b61028f6106fd366004614ace565b6127af565b61028f610710366004614311565b613194565b6040805160608082018352600080835260208084018290529284015282518082018452601154808252601254828501908152601354928601928352855191825251938101939093525192820192909252016102eb565b61028f61077936600461424e565b6133d4565b6107866133e8565b73ffffffffffffffffffffffffffffffffffffffff81166107d3576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b6108626133e8565b60085460008190036108a0576040517f2e70248b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529061093b576040517f3917193900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff16146109a4576040517f6c17b98700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060086109b3600185614bca565b815481106109c3576109c3614be1565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600882602001516bffffffffffffffffffffffff1681548110610a1557610a15614be1565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166008610a44600186614bca565b81548110610a5457610a54614be1565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600883602001516bffffffffffffffffffffffff1681548110610ac257610ac2614be1565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526007909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556008805480610b6457610b64614c10565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600783526040808520949094558351908152908716918101919091527fa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d91015b60405180910390a15050505050565b60005460ff1615610c9c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610d0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2f9190614c3f565b15610d65576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b4916004808301926060929190829003018187875af1158015610dd6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dfa9190614c5c565b9050600354816020015142610e0f9190614bca565b1115610e47576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60155473ffffffffffffffffffffffffffffffffffffffff16610e96576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610ea28585611c01565b6000818152600f6020526040812054919250819003610ef15784866040517f07e6809a000000000000000000000000000000000000000000000000000000008152600401610c93929190614dce565b6014544290610f169068010000000000000000900467ffffffffffffffff1683614e30565b10610f4d576040517f15c33ba200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855160009081526010602052604090205460ff1615610f9e5785516040517f6a64e9610000000000000000000000000000000000000000000000000000000081526004810191909152602401610c93565b60608601516080015173ffffffffffffffffffffffffffffffffffffffff1615801590610fe9575060608601516080015173ffffffffffffffffffffffffffffffffffffffff163314155b156110265785516040517fd8e90b980000000000000000000000000000000000000000000000000000000081526004810191909152602401610c93565b61102f8661346e565b611038866135e9565b8551600090815260106020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905583156112765760008087606001516000015160008151811061109557611095614be1565b6020026020010151905060006110d08273ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff811661111f576040517f83135fec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166350d25bcd6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561116a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061118e9190614e48565b6014546111a5919067ffffffffffffffff16614e61565b9250821561127257828960600151602001516000815181106111c9576111c9614be1565b602002602001018181516111dd9190614bca565b9052506111e982613698565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081523360048201526024810185905273ffffffffffffffffffffffffffffffffffffffff919091169063ea6192a290604401600060405180830381600087803b15801561125957600080fd5b505af115801561126d573d6000803e3d6000fd5b505050505b5050505b60005b6060870151515181101561137a576112b187606001516000015182815181106112a4576112a4614be1565b6020026020010151613698565b73ffffffffffffffffffffffffffffffffffffffff1663ea6192a288606001516060015189606001516020015184815181106112ef576112ef614be1565b60200260200101516040518363ffffffff1660e01b815260040161133592919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600060405180830381600087803b15801561134f57600080fd5b505af1158015611363573d6000803e3d6000fd5b50505050808061137290614e9e565b915050611279565b50606080870151015173ffffffffffffffffffffffffffffffffffffffff163b156114985760155460608088015101516040517f5dd0851c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90921691635dd0851c916113fd918a90600401614ed6565b600060405180830381600087803b15801561141757600080fd5b505af1925050508015611428575060015b611493573d808015611456576040519150601f19603f3d011682016040523d82523d6000602084013e61145b565b606091505b5086516040517f6a3fd4f2000000000000000000000000000000000000000000000000000000008152610c9391908390600401614f05565b6114df565b606086015160a0015151156114df5785516040517fe0244be30000000000000000000000000000000000000000000000000000000081526004810191909152602401610c93565b85516040517fc51bf0f6d90b467e0849da0ad18a4d9144a4b78b9f83202e1c65cd68f72d4a3890600090a2505050505050565b61151a6133e8565b611522613714565b565b61152c6133e8565b8060146115398282614f1e565b9050507fe4cd88b1b5e20a0b843af3207ba74d3f84af4acff45830469490ac9c6ab8a7458160405161156b919061502f565b60405180910390a150565b61157e6133e8565b73ffffffffffffffffffffffffffffffffffffffff821615806115b5575073ffffffffffffffffffffffffffffffffffffffff8116155b156115ec576040517fee9d106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015611688576040517f965ffb7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600880546bffffffffffffffffffffffff908116602080870191825288861660008181526007835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821790558351908152908101919091527f037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb60130591015b60405180910390a1505050565b6117966133e8565b6117b773ffffffffffffffffffffffffffffffffffffffff841683836137f5565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa890606001611781565b6118146133e8565b6005546000819003611852576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906118ed576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614611956576040517fd428911900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006005611965600185614bca565b8154811061197557611975614be1565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff16815481106119c7576119c7614be1565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1660056119f6600186614bca565b81548110611a0657611a06614be1565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110611a7457611a74614be1565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480611b1657611b16614c10565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101610c1b565b600080600060f81b84604051602001611c1a9190615099565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052611c5692916020016150ac565b60405160208183030381529060405280519060200120905060005b835151811015611d8857600084600001518281518110611c9357611c93614be1565b6020026020010151905060028560200151611cae9190615123565b600003611d0c576040517f010000000000000000000000000000000000000000000000000000000000000060208201526021810184905260418101829052606101604051602081830303815290604052805190602001209250611d5f565b6040517f0100000000000000000000000000000000000000000000000000000000000000602082015260218101829052604181018490526061016040516020818303038152906040528051906020012092505b60028560200151611d709190615137565b60208601525080611d8081614e9e565b915050611c71565b509392505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611e11576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610c93565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600e805480602002602001604051908101604052809291908181526020018280548015611f1757602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611eec575b5050505050905090565b60606008805480602002602001604051908101604052809291908181526020018280548015611f175760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611eec575050505050905090565b611f966133e8565b611522613887565b60606005805480602002602001604051908101604052809291908181526020018280548015611f175760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611eec575050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161206191849163ffffffff851691908e908e908190840183828082843760009201919091525061394792505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600954808252600a5460ff80821660208501526101009091041692820192909252908314612136576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610c93565b6121448b8b8b8b8b8b613cb8565b60007f0000000000000000000000000000000000000000000000000000000000000000156121a157600282602001518360400151612182919061514b565b61218c9190615170565b61219790600161514b565b60ff1690506121b7565b60208201516121b190600161514b565b60ff1690505b888114612220576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610c93565b888714612289576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610c93565b336000908152600c602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156122cc576122cc615192565b60028111156122dd576122dd615192565b90525090506002816020015160028111156122fa576122fa615192565b1480156123415750600e816000015160ff168154811061231c5761231c614be1565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b6123a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610c93565b5050505050600088886040516123be9291906151c1565b6040519081900381206123d5918c906020016151d1565b6040516020818303038152906040528051906020012090506123f561420d565b604080518082019091526000808252602082015260005b8881101561269257600060018588846020811061242b5761242b614be1565b61243891901a601b61514b565b8d8d8681811061244a5761244a614be1565b905060200201358c8c8781811061246357612463614be1565b90506020020135604051600081526020016040526040516124a0949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156124c2573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600c602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561254257612542615192565b600281111561255357612553615192565b905250925060018360200151600281111561257057612570615192565b146125d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610c93565b8251849060ff16601f81106125ee576125ee614be1565b60200201511561265a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610c93565b600184846000015160ff16601f811061267557612675614be1565b91151560209092020152508061268a81614e9e565b91505061240c565b5050505063ffffffff81106126a9576126a96151ed565b505050505050505050565b6126bc6133e8565b806000036126f6576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251910161084e565b61273c6133e8565b601580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d49060200161156b565b855185518560ff16601f831115612822576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610c93565b6000811161288c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610c93565b81831461291a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610c93565b612925816003614e61565b831161298d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610c93565b6129956133e8565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600d5415612b8857600d546000906129ed90600190614bca565b90506000600d8281548110612a0457612a04614be1565b6000918252602082200154600e805473ffffffffffffffffffffffffffffffffffffffff90921693509084908110612a3e57612a3e614be1565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600c909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600d80549192509080612abe57612abe614c10565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600e805480612b2757612b27614c10565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055506129d3915050565b60005b815151811015612fef576000600c600084600001518481518110612bb157612bb1614be1565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612bfb57612bfb615192565b14612c62576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610c93565b6040805180820190915260ff821681526001602082015282518051600c9160009185908110612c9357612c93614be1565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612d3457612d34615192565b021790555060009150612d449050565b600c600084602001518481518110612d5e57612d5e614be1565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612da857612da8615192565b14612e0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610c93565b6040805180820190915260ff82168152602081016002815250600c600084602001518481518110612e4257612e42614be1565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612ee357612ee3615192565b02179055505082518051600d925083908110612f0157612f01614be1565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600e919083908110612f7d57612f7d614be1565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691909117905580612fe781614e9e565b915050612b8b565b506040810151600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600b80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff43811682029283178555908304811693600193909260009261308192869290821691161761521c565b92506101000a81548163ffffffff021916908363ffffffff1602179055506130e04630600b60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613d6f565b600981905582518051600a805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600b5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e059861317f988b98919763ffffffff909216969095919491939192615244565b60405180910390a15050505050505050505050565b61319c6133e8565b73ffffffffffffffffffffffffffffffffffffffff821615806131d3575073ffffffffffffffffffffffffffffffffffffffff8116155b1561320a576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156132a6576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c9101611781565b6133dc6133e8565b6133e581613e1a565b50565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314611522576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610c93565b7f00000000000000000000000000000000000000000000000000000000000000008160200151146134d35780602001516040517fd44bc9eb000000000000000000000000000000000000000000000000000000008152600401610c9391815260200190565b60145460608201515151780100000000000000000000000000000000000000000000000090910467ffffffffffffffff16108061351c5750606081015160208101515190515114155b15613553576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601454606082015160a001515170010000000000000000000000000000000090910467ffffffffffffffff1610156133e557601454606082015160a00151516040517f8693378900000000000000000000000000000000000000000000000000000000815270010000000000000000000000000000000090920467ffffffffffffffff1660048301526024820152604401610c93565b606080820151015173ffffffffffffffffffffffffffffffffffffffff163014806136405750606080820151015173ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff165b156133e55760608082015101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610c93565b73ffffffffffffffffffffffffffffffffffffffff818116600090815260046020526040902054168061370f576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610c93565b919050565b60005460ff16613780576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610c93565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052613882908490613f15565b505050565b60005460ff16156138f4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c93565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586137cb3390565b60005460ff16156139b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c93565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015613a23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a479190614c3f565b15613a7d576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b4916004808301926060929190829003018187875af1158015613aee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b129190614c5c565b9050600354816020015142613b279190614bca565b1115613b5f576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082806020019051810190613b759190614c5c565b9050806040015181602001511115613bb9576040517f67a3824c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516060810182526011548082526012546020830152601354928201929092529015613c42576040810151613bf1906001614e30565b826020015114613c4257806040015182602001516040517fcc7f1bd0000000000000000000000000000000000000000000000000000000008152600401610c93929190918252602082015260400190565b81516000908152600f60209081526040918290204290558351601181905581850180516012558386018051601355845192835290519282019290925290518183015290517f07d7bce06be2a7b0230e4dd0d72523c0407e82419ab1d947c5ddaf59ca36484e9181900360600190a1505050505050565b6000613cc5826020614e61565b613cd0856020614e61565b613cdc88610144614e30565b613ce69190614e30565b613cf09190614e30565b613cfb906000614e30565b9050368114613d66576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610c93565b50505050505050565b6000808a8a8a8a8a8a8a8a8a604051602001613d93999897969594939291906152da565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603613e99576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610c93565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000613f77826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166140219092919063ffffffff16565b8051909150156138825780806020019051810190613f959190614c3f565b613882576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610c93565b6060614030848460008561403a565b90505b9392505050565b6060824710156140cc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610c93565b843b614134576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610c93565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161415d919061536f565b60006040518083038185875af1925050503d806000811461419a576040519150601f19603f3d011682016040523d82523d6000602084013e61419f565b606091505b50915091506141af8282866141ba565b979650505050505050565b606083156141c9575081614033565b8251156141d95782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9391906142e5565b604051806103e00160405280601f906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff811681146133e557600080fd5b60006020828403121561426057600080fd5b81356140338161422c565b60005b8381101561428657818101518382015260200161426e565b83811115614295576000848401525b50505050565b600081518084526142b381602086016020860161426b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000614033602083018461429b565b60006020828403121561430a57600080fd5b5035919050565b6000806040838503121561432457600080fd5b823561432f8161422c565b9150602083013561433f8161422c565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff8111828210171561439c5761439c61434a565b60405290565b60405160e0810167ffffffffffffffff8111828210171561439c5761439c61434a565b6040805190810167ffffffffffffffff8111828210171561439c5761439c61434a565b6040516060810167ffffffffffffffff8111828210171561439c5761439c61434a565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156144525761445261434a565b604052919050565b803561370f8161422c565b600067ffffffffffffffff82111561447f5761447f61434a565b5060051b60200190565b600082601f83011261449a57600080fd5b813560206144af6144aa83614465565b61440b565b82815260059290921b840181019181810190868411156144ce57600080fd5b8286015b848110156144f25780356144e58161422c565b83529183019183016144d2565b509695505050505050565b600082601f83011261450e57600080fd5b8135602061451e6144aa83614465565b82815260059290921b8401810191818101908684111561453d57600080fd5b8286015b848110156144f25780358352918301918301614541565b600082601f83011261456957600080fd5b813567ffffffffffffffff8111156145835761458361434a565b6145b460207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161440b565b8181528460208386010111156145c957600080fd5b816020850160208301376000918101602001919091529392505050565b6000608082840312156145f857600080fd5b614600614379565b90508135815260208201356020820152604082013561461e8161422c565b6040820152606082013567ffffffffffffffff8082111561463e57600080fd5b9083019060e0828603121561465257600080fd5b61465a6143a2565b82358281111561466957600080fd5b61467587828601614489565b82525060208301358281111561468a57600080fd5b614696878286016144fd565b602083015250604083013560408201526146b26060840161445a565b60608201526146c36080840161445a565b608082015260a0830135828111156146da57600080fd5b6146e687828601614558565b60a08301525060c0830135828111156146fe57600080fd5b61470a87828601614558565b60c083015250606084015250909392505050565b60006040828403121561473057600080fd5b6147386143c5565b9050813567ffffffffffffffff81111561475157600080fd5b8201601f8101841361476257600080fd5b803560206147726144aa83614465565b82815260059290921b8301810191818101908784111561479157600080fd5b938201935b838510156147af57843582529382019390820190614796565b85525093840135938301939093525092915050565b80151581146133e557600080fd5b6000806000606084860312156147e757600080fd5b833567ffffffffffffffff808211156147ff57600080fd5b61480b878388016145e6565b9450602086013591508082111561482157600080fd5b5061482e8682870161471e565b925050604084013561483f816147c4565b809150509250925092565b60006080828403121561485c57600080fd5b50919050565b60008060006060848603121561487757600080fd5b83356148828161422c565b925060208401356148928161422c565b929592945050506040919091013590565b600080604083850312156148b657600080fd5b823567ffffffffffffffff808211156148ce57600080fd5b6148da868387016145e6565b935060208501359150808211156148f057600080fd5b506148fd8582860161471e565b9150509250929050565b600081518084526020808501945080840160005b8381101561494d57815173ffffffffffffffffffffffffffffffffffffffff168752958201959082019060010161491b565b509495945050505050565b6020815260006140336020830184614907565b60008083601f84011261497d57600080fd5b50813567ffffffffffffffff81111561499557600080fd5b6020830191508360208260051b85010111156149b057600080fd5b9250929050565b60008060008060008060008060e0898b0312156149d357600080fd5b606089018a8111156149e457600080fd5b8998503567ffffffffffffffff808211156149fe57600080fd5b818b0191508b601f830112614a1257600080fd5b813581811115614a2157600080fd5b8c6020828501011115614a3357600080fd5b6020830199508098505060808b0135915080821115614a5157600080fd5b614a5d8c838d0161496b565b909750955060a08b0135915080821115614a7657600080fd5b50614a838b828c0161496b565b999c989b50969995989497949560c00135949350505050565b803560ff8116811461370f57600080fd5b67ffffffffffffffff811681146133e557600080fd5b803561370f81614aad565b60008060008060008060c08789031215614ae757600080fd5b863567ffffffffffffffff80821115614aff57600080fd5b614b0b8a838b01614489565b97506020890135915080821115614b2157600080fd5b614b2d8a838b01614489565b9650614b3b60408a01614a9c565b95506060890135915080821115614b5157600080fd5b614b5d8a838b01614558565b9450614b6b60808a01614ac3565b935060a0890135915080821115614b8157600080fd5b50614b8e89828a01614558565b9150509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015614bdc57614bdc614b9b565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060208284031215614c5157600080fd5b8151614033816147c4565b600060608284031215614c6e57600080fd5b614c766143e8565b8251815260208301516020820152604083015160408201528091505092915050565b8051825260006020808301518185015273ffffffffffffffffffffffffffffffffffffffff6040840151166040850152606083015160806060860152805160e06080870152614ceb610160870182614907565b838301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80888303810160a08a0152815180845291860193506000929091908601905b80841015614d4e5784518252938601936001939093019290860190614d2e565b50604085015160c08a0152606085015173ffffffffffffffffffffffffffffffffffffffff90811660e08b01526080860151166101008a015260a085015189820383016101208b01529550614da3818761429b565b95505060c084015193508088860301610140890152505050614dc5828261429b565b95945050505050565b60408152600060808201845160408085015281815180845260a08601915060209350838301925060005b81811015614e1457835183529284019291840191600101614df8565b5050828701516060860152848103838601526141af8187614c98565b60008219821115614e4357614e43614b9b565b500190565b600060208284031215614e5a57600080fd5b5051919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614e9957614e99614b9b565b500290565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614ecf57614ecf614b9b565b5060010190565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006140306040830184614c98565b828152604060208201526000614030604083018461429b565b8135614f2981614aad565b67ffffffffffffffff811690508154817fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000082161783556020840135614f6d81614aad565b6fffffffffffffffff00000000000000008160401b16905080837fffffffffffffffffffffffffffffffff000000000000000000000000000000008416171784556040850135614fbc81614aad565b77ffffffffffffffff000000000000000000000000000000008160801b1690507fffffffffffffffff00000000000000000000000000000000000000000000000081858286161784171786556060870135935061501884614aad565b808460c01b16858417831717865550505050505050565b60808101823561503e81614aad565b67ffffffffffffffff908116835260208401359061505b82614aad565b908116602084015260408401359061507282614aad565b908116604084015260608401359061508982614aad565b8082166060850152505092915050565b6020815260006140336020830184614c98565b7fff0000000000000000000000000000000000000000000000000000000000000083168152600082516150e681600185016020870161426b565b919091016001019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082615132576151326150f4565b500690565b600082615146576151466150f4565b500490565b600060ff821660ff84168060ff0382111561516857615168614b9b565b019392505050565b600060ff831680615183576151836150f4565b8060ff84160491505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600063ffffffff80831681851680830382111561523b5761523b614b9b565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526152748184018a614907565b905082810360808401526152888189614907565b905060ff871660a084015282810360c08401526152a5818761429b565b905067ffffffffffffffff851660e08401528281036101008401526152ca818561429b565b9c9b505050505050505050505050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526153218285018b614907565b91508382036080850152615335828a614907565b915060ff881660a085015283820360c0850152615352828861429b565b90861660e085015283810361010085015290506152ca818561429b565b6000825161538181846020870161426b565b919091019291505056fea164736f6c634300080d000a",
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

	GetRouter(opts *bind.CallOpts) (common.Address, error)

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
