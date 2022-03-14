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

var OffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"contractAggregatorV2V3Interface[]\",\"name\":\"feeds\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionDelaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokensLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionFeeLink\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDataSize\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionDelayError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"InvalidExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceFeedConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MerkleProofError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RelayReportError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lastMaxSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMinSequenceNumber\",\"type\":\"uint256\"}],\"name\":\"SequenceError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenFeedMistmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMistmatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"UnexpectedPayloadData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"CrossChainMessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"ExecutionDelaySecondsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"}],\"name\":\"ExecutionFeeLinkSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"MaxDataSizeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"minSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSequenceNumber\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCCIP.RelayReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"ReportAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"addFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"needFee\",\"type\":\"bool\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"getExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutionDelaySeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutionFeeLink\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastReport\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"minSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSequenceNumber\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.RelayReport\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxDataSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"removeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"executionDelaySeconds\",\"type\":\"uint256\"}],\"name\":\"setExecutionDelaySeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"executionFeeLink\",\"type\":\"uint256\"}],\"name\":\"setExecutionFeeLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxDataSize\",\"type\":\"uint256\"}],\"name\":\"setMaxDataSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200589438038062005894833981016040819052620000349162000645565b6000805460ff191681556001908a90899082908c908b908b903390819081620000a45760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000de57620000de816200038d565b5050506001600160a01b0382161580620000f6575080155b156200011557604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03939093169290921790915560035580518251146200015b5760405162d8548360e71b815260040160405180910390fd5b8151620001709060059060208501906200043f565b5060005b8251811015620002545760008282815181106200019557620001956200073f565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001df57620001df6200073f565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff19166001179055806200024b8162000755565b91505062000174565b50505080518251146200027a5760405163ee9d106b60e01b815260040160405180910390fd5b81516200028f9060089060208501906200043f565b5060005b82518110156200035c576000828281518110620002b457620002b46200073f565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060076000868581518110620002fe57620002fe6200073f565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b02929091169190911790555080620003538162000755565b91505062000293565b505050151560805260a09a909a5260c09890985260129190915560175550505060119290925550506013556200077f565b6001600160a01b038116331415620003e85760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016200009b565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b82805482825590600052602060002090810192821562000497579160200282015b828111156200049757825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000460565b50620004a5929150620004a9565b5090565b5b80821115620004a55760008155600101620004aa565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620005015762000501620004c0565b604052919050565b60006001600160401b03821115620005255762000525620004c0565b5060051b60200190565b6001600160a01b03811681146200054557600080fd5b50565b600082601f8301126200055a57600080fd5b81516020620005736200056d8362000509565b620004d6565b82815260059290921b840181019181810190868411156200059357600080fd5b8286015b84811015620005bb578051620005ad816200052f565b835291830191830162000597565b509695505050505050565b600082601f830112620005d857600080fd5b81516020620005eb6200056d8362000509565b82815260059290921b840181019181810190868411156200060b57600080fd5b8286015b84811015620005bb57805162000625816200052f565b83529183019183016200060f565b805162000640816200052f565b919050565b60008060008060008060008060008060006101608c8e0312156200066857600080fd5b8b5160208d015160408e0151919c509a506001600160401b038111156200068e57600080fd5b6200069c8e828f0162000548565b60608e0151909a5090506001600160401b03811115620006bb57600080fd5b620006c98e828f0162000548565b60808e015190995090506001600160401b03811115620006e857600080fd5b620006f68e828f01620005c6565b9750506200070760a08d0162000633565b955060c08c0151945060e08c015193506101008c015192506101208c015191506101408c015190509295989b509295989b9093969950565b634e487b7160e01b600052603260045260246000fd5b60006000198214156200077857634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c0516150de620007b6600039600061051b01526000818161048401526133f40152600061210d01526150de6000f3fe608060405234801561001057600080fd5b50600436106102ad5760003560e01c806381be8fa41161017b578063afcb95d7116100d8578063c15d605b1161008c578063eb511dd411610071578063eb511dd414610663578063eefa7a3e14610676578063f2fde38b146106cc57600080fd5b8063c15d605b1461063d578063e3d0e7121461065057600080fd5b8063b1dc65a4116100bd578063b1dc65a4146105de578063b6608c3b146105f1578063bbe4f6db1461060457600080fd5b8063afcb95d7146105b6578063b034909c146105d657600080fd5b806389c065681161012f5780638da5cb5b116101145780638da5cb5b14610560578063975de62b14610583578063a7206cd61461059657600080fd5b806389c065681461055057806389e19b231461055857600080fd5b80638456cb59116101605780638456cb591461050e57806385e1f4d014610516578063879b32c51461053d57600080fd5b806381be8fa4146104d657806381ff7048146104de57600080fd5b80633f4ba83a11610229578063744b92e2116101dd578063768c577b116101c2578063768c577b146104a657806379ba5097146104b957806381411834146104c157600080fd5b8063744b92e21461046c57806374be21501461047f57600080fd5b806359e96b5b1161020e57806359e96b5b146104155780635b16ebb7146104285780635c975abb1461046157600080fd5b80633f4ba83a146103fa5780635853c6271461040257600080fd5b80631a830a1e116102805780632222dd42116102655780632222dd42146103b65780632b898c25146103d45780633b8d08ef146103e757600080fd5b80631a830a1e1461037b578063219475071461038357600080fd5b806307db2b21146102b2578063108ee5fc146102c957806316b8e731146102de578063181f5a771461033c575b600080fd5b6011545b6040519081526020015b60405180910390f35b6102dc6102d73660046140f5565b6106df565b005b6103176102ec3660046140f5565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102c0565b604080518082018252600d81527f4f666652616d7020302e302e3100000000000000000000000000000000000000602082015290516102c0919061418c565b6012546102b6565b6103a661039136600461419f565b60009081526010602052604090205460ff1690565b60405190151581526020016102c0565b60025473ffffffffffffffffffffffffffffffffffffffff16610317565b6102dc6103e23660046141b8565b6107bb565b6102dc6103f536600461467e565b610b87565b6102dc6114af565b6102dc6104103660046141b8565b6114c1565b6102dc6104233660046146f6565b6116d9565b6103a66104363660046140f5565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff166103a6565b6102dc61047a3660046141b8565b611757565b6102b67f000000000000000000000000000000000000000000000000000000000000000081565b6102b66104b4366004614737565b611b48565b6102dc611cd4565b6104c9611df6565b6040516102c091906147ec565b6104c9611e65565b600b546009546040805163ffffffff808516825264010000000090940490931660208401528201526060016102c0565b6102dc611ed2565b6102b67f000000000000000000000000000000000000000000000000000000000000000081565b6102dc61054b36600461419f565b611ee2565b6104c9611f26565b6013546102b6565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610317565b6102dc61059136600461419f565b611f93565b6102b66105a436600461419f565b6000908152600f602052604090205490565b6040805160018152600060208201819052918101919091526060016102c0565b6003546102b6565b6102dc6105ec36600461484b565b611fd0565b6102dc6105ff36600461419f565b612679565b6103176106123660046140f5565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6102dc61064b36600461419f565b6126f6565b6102dc61065e3660046149bd565b612733565b6102dc6106713660046141b8565b613118565b6040805160608082018352600080835260208084018290529284015282518082018452601454808252601554828501908152601654928601928352855191825251938101939093525192820192909252016102c0565b6102dc6106da3660046140f5565b613358565b6106e761336c565b73ffffffffffffffffffffffffffffffffffffffff8116610734576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b6107c361336c565b600854806107fd576040517f2e70248b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610898576040517f3917193900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610901576040517f6c17b98700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006008610910600185614ab9565b8154811061092057610920614ad0565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600882602001516bffffffffffffffffffffffff168154811061097257610972614ad0565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1660086109a1600186614ab9565b815481106109b1576109b1614ad0565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600883602001516bffffffffffffffffffffffff1681548110610a1f57610a1f614ad0565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526007909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556008805480610ac157610ac1614aff565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600783526040808520949094558351908152908716918101919091527fa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d91015b60405180910390a15050505050565b60005460ff1615610bf9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610c68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c8c9190614b2e565b15610cc2576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b4916004808301926060929190829003018187875af1158015610d33573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d579190614b4b565b9050600354816020015142610d6c9190614ab9565b1115610da4576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610db08585611b48565b6000818152600f602052604090205490915080610dfd5784866040517f07e6809a000000000000000000000000000000000000000000000000000000008152600401610bf0929190614cbd565b4260125482610e0c9190614d1f565b10610e43576040517f15c33ba200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855160009081526010602052604090205460ff1615610e945785516040517f6a64e9610000000000000000000000000000000000000000000000000000000081526004810191909152602401610bf0565b60608601516080015173ffffffffffffffffffffffffffffffffffffffff1615801590610edf575060608601516080015173ffffffffffffffffffffffffffffffffffffffff163314155b15610f1c5785516040517fd8e90b980000000000000000000000000000000000000000000000000000000081526004810191909152602401610bf0565b610f25866133f2565b610f2e8661350b565b8551600090815260106020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055831561117a57600080876060015160000151600081518110610f8b57610f8b614ad0565b602002602001015190506000610fc68273ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff8116611015576040517f83135fec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166350d25bcd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611060573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110849190614d37565b6011546110919190614d50565b9250828960600151602001516000815181106110af576110af614ad0565b602002602001018181516110c39190614ab9565b90525073ffffffffffffffffffffffffffffffffffffffff808316600090815260046020526040902054166040517fea6192a20000000000000000000000000000000000000000000000000000000081523360048201526024810185905273ffffffffffffffffffffffffffffffffffffffff919091169063ea6192a290604401600060405180830381600087803b15801561115e57600080fd5b505af1158015611172573d6000803e3d6000fd5b505050505050505b60005b606087015151518110156113345760006111dd88606001516000015183815181106111aa576111aa614ad0565b602002602001015173ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff81166112695760608801515180518390811061121257611212614ad0565b60200260200101516040517fbf16aab6000000000000000000000000000000000000000000000000000000008152600401610bf0919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b8073ffffffffffffffffffffffffffffffffffffffff1663ea6192a28960600151606001518a606001516020015185815181106112a8576112a8614ad0565b60200260200101516040518363ffffffff1660e01b81526004016112ee92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600060405180830381600087803b15801561130857600080fd5b505af115801561131c573d6000803e3d6000fd5b5050505050808061132c90614d8d565b91505061117d565b50606080870151015173ffffffffffffffffffffffffffffffffffffffff163b156114605785606001516060015173ffffffffffffffffffffffffffffffffffffffff16639c598468876040518263ffffffff1660e01b815260040161139a9190614dc6565b600060405180830381600087803b1580156113b457600080fd5b505af19250505080156113c5575060015b611430573d8080156113f3576040519150601f19603f3d011682016040523d82523d6000602084013e6113f8565b606091505b5086516040517f6a3fd4f2000000000000000000000000000000000000000000000000000000008152610bf091908390600401614dd9565b85516040517fc51bf0f6d90b467e0849da0ad18a4d9144a4b78b9f83202e1c65cd68f72d4a3890600090a26114a7565b606086015160a0015151156114a75785516040517fe0244be30000000000000000000000000000000000000000000000000000000081526004810191909152602401610bf0565b505050505050565b6114b761336c565b6114bf6135ba565b565b6114c961336c565b73ffffffffffffffffffffffffffffffffffffffff82161580611500575073ffffffffffffffffffffffffffffffffffffffff8116155b15611537576040517fee9d106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156115d3576040517f965ffb7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600880546bffffffffffffffffffffffff908116602080870191825288861660008181526007835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821790558351908152908101919091527f037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb60130591015b60405180910390a1505050565b6116e161336c565b61170273ffffffffffffffffffffffffffffffffffffffff8416838361369b565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8906060016116cc565b61175f61336c565b60055480611799576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290611834576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff161461189d576040517fd428911900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060056118ac600185614ab9565b815481106118bc576118bc614ad0565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff168154811061190e5761190e614ad0565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16600561193d600186614ab9565b8154811061194d5761194d614ad0565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff16815481106119bb576119bb614ad0565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480611a5d57611a5d614aff565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101610b78565b600080600060f81b84604051602001611b619190614dc6565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052611b9d9291602001614df2565b60405160208183030381529060405280519060200120905060005b835151811015611ccc57600084600001518281518110611bda57611bda614ad0565b6020026020010151905060028560200151611bf59190614e69565b611c50576040517f010000000000000000000000000000000000000000000000000000000000000060208201526021810184905260418101829052606101604051602081830303815290604052805190602001209250611ca3565b6040517f0100000000000000000000000000000000000000000000000000000000000000602082015260218101829052604181018490526061016040516020818303038152906040528051906020012092505b60028560200151611cb49190614e7d565b60208601525080611cc481614d8d565b915050611bb8565b509392505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611d55576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610bf0565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600e805480602002602001604051908101604052809291908181526020018280548015611e5b57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611e30575b5050505050905090565b60606008805480602002602001604051908101604052809291908181526020018280548015611e5b5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611e30575050505050905090565b611eda61336c565b6114bf61372d565b611eea61336c565b60118190556040518181527fd7fd8173fe9fe63cbd28468ae5012d7f8db1f1550fbfe8c22a1b49e4b0b703ae906020015b60405180910390a150565b60606005805480602002602001604051908101604052809291908181526020018280548015611e5b5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611e30575050505050905090565b611f9b61336c565b60138190556040518181527f15e8871b18d0879e199f8794daad195ca64bbb7341d01951a980a1abf0800bfb90602001611f1b565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161202691849163ffffffff851691908e908e90819084018382808284376000920191909152506137ed92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600954808252600a5460ff808216602085015261010090910416928201929092529083146120fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610bf0565b6121098b8b8b8b8b8b613b5e565b60007f000000000000000000000000000000000000000000000000000000000000000015612166576002826020015183604001516121479190614e91565b6121519190614eb6565b61215c906001614e91565b60ff16905061217c565b6020820151612176906001614e91565b60ff1690505b8881146121e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610bf0565b88871461224e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610bf0565b336000908152600c602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561229157612291614ed8565b60028111156122a2576122a2614ed8565b90525090506002816020015160028111156122bf576122bf614ed8565b1480156123065750600e816000015160ff16815481106122e1576122e1614ad0565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b61236c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610bf0565b505050505060008888604051612383929190614f07565b60405190819003812061239a918c90602001614f17565b6040516020818303038152906040528051906020012090506123ba6140b4565b604080518082019091526000808252602082015260005b888110156126575760006001858884602081106123f0576123f0614ad0565b6123fd91901a601b614e91565b8d8d8681811061240f5761240f614ad0565b905060200201358c8c8781811061242857612428614ad0565b9050602002013560405160008152602001604052604051612465949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015612487573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600c602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561250757612507614ed8565b600281111561251857612518614ed8565b905250925060018360200151600281111561253557612535614ed8565b1461259c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610bf0565b8251849060ff16601f81106125b3576125b3614ad0565b60200201511561261f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610bf0565b600184846000015160ff16601f811061263a5761263a614ad0565b91151560209092020152508061264f81614d8d565b9150506123d1565b5050505063ffffffff811061266e5761266e614f33565b505050505050505050565b61268161336c565b806126b8576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c25191016107af565b6126fe61336c565b60128190556040518181527ffbb92e09fd0b7bcabafe655657e745c10bcd9812c1e32e205e477a3c5ff74fc090602001611f1b565b855185518560ff16601f8311156127a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610bf0565b60008111612810576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610bf0565b81831461289e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610bf0565b6128a9816003614d50565b8311612911576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610bf0565b61291961336c565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600d5415612b0c57600d5460009061297190600190614ab9565b90506000600d828154811061298857612988614ad0565b6000918252602082200154600e805473ffffffffffffffffffffffffffffffffffffffff909216935090849081106129c2576129c2614ad0565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600c909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600d80549192509080612a4257612a42614aff565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600e805480612aab57612aab614aff565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550612957915050565b60005b815151811015612f73576000600c600084600001518481518110612b3557612b35614ad0565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612b7f57612b7f614ed8565b14612be6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610bf0565b6040805180820190915260ff821681526001602082015282518051600c9160009185908110612c1757612c17614ad0565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612cb857612cb8614ed8565b021790555060009150612cc89050565b600c600084602001518481518110612ce257612ce2614ad0565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612d2c57612d2c614ed8565b14612d93576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610bf0565b6040805180820190915260ff82168152602081016002815250600c600084602001518481518110612dc657612dc6614ad0565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612e6757612e67614ed8565b02179055505082518051600d925083908110612e8557612e85614ad0565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600e919083908110612f0157612f01614ad0565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691909117905580612f6b81614d8d565b915050612b0f565b506040810151600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600b80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092613005928692908216911617614f62565b92506101000a81548163ffffffff021916908363ffffffff1602179055506130644630600b60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613c15565b600981905582518051600a805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600b5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598613103988b98919763ffffffff909216969095919491939192614f8a565b60405180910390a15050505050505050505050565b61312061336c565b73ffffffffffffffffffffffffffffffffffffffff82161580613157575073ffffffffffffffffffffffffffffffffffffffff8116155b1561318e576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152901561322a576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c91016116cc565b61336061336c565b61336981613cc0565b50565b600054610100900473ffffffffffffffffffffffffffffffffffffffff1633146114bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610bf0565b7f00000000000000000000000000000000000000000000000000000000000000008160200151146134575780602001516040517fd44bc9eb000000000000000000000000000000000000000000000000000000008152600401610bf091815260200190565b6017546060820151515111806134795750606081015160208101515190515114155b156134b0576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601354816060015160a0015151111561336957601354606082015160a00151516040517f8693378900000000000000000000000000000000000000000000000000000000815260048101929092526024820152604401610bf0565b606080820151015173ffffffffffffffffffffffffffffffffffffffff163014806135625750606080820151015173ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff165b156133695760608082015101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610bf0565b60005460ff16613626576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610bf0565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052613728908490613dbc565b505050565b60005460ff161561379a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610bf0565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586136713390565b60005460ff161561385a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610bf0565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b81526004016020604051808303816000875af11580156138c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138ed9190614b2e565b15613923576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b4916004808301926060929190829003018187875af1158015613994573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139b89190614b4b565b90506003548160200151426139cd9190614ab9565b1115613a05576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082806020019051810190613a1b9190614b4b565b9050806040015181602001511115613a5f576040517f67a3824c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516060810182526014548082526015546020830152601654928201929092529015613ae8576040810151613a97906001614d1f565b826020015114613ae857806040015182602001516040517fcc7f1bd0000000000000000000000000000000000000000000000000000000008152600401610bf0929190918252602082015260400190565b81516000908152600f60209081526040918290204290558351601481905581850180516015558386018051601655845192835290519282019290925290518183015290517f07d7bce06be2a7b0230e4dd0d72523c0407e82419ab1d947c5ddaf59ca36484e9181900360600190a1505050505050565b6000613b6b826020614d50565b613b76856020614d50565b613b8288610144614d1f565b613b8c9190614d1f565b613b969190614d1f565b613ba1906000614d1f565b9050368114613c0c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610bf0565b50505050505050565b6000808a8a8a8a8a8a8a8a8a604051602001613c3999989796959493929190615020565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116331415613d40576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610bf0565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000613e1e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16613ec89092919063ffffffff16565b8051909150156137285780806020019051810190613e3c9190614b2e565b613728576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610bf0565b6060613ed78484600085613ee1565b90505b9392505050565b606082471015613f73576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610bf0565b843b613fdb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610bf0565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161400491906150b5565b60006040518083038185875af1925050503d8060008114614041576040519150601f19603f3d011682016040523d82523d6000602084013e614046565b606091505b5091509150614056828286614061565b979650505050505050565b60608315614070575081613eda565b8251156140805782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bf0919061418c565b604051806103e00160405280601f906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff8116811461336957600080fd5b60006020828403121561410757600080fd5b8135613eda816140d3565b60005b8381101561412d578181015183820152602001614115565b8381111561413c576000848401525b50505050565b6000815180845261415a816020860160208601614112565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613eda6020830184614142565b6000602082840312156141b157600080fd5b5035919050565b600080604083850312156141cb57600080fd5b82356141d6816140d3565b915060208301356141e6816140d3565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff81118282101715614243576142436141f1565b60405290565b60405160e0810167ffffffffffffffff81118282101715614243576142436141f1565b6040805190810167ffffffffffffffff81118282101715614243576142436141f1565b6040516060810167ffffffffffffffff81118282101715614243576142436141f1565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156142f9576142f96141f1565b604052919050565b803561430c816140d3565b919050565b600067ffffffffffffffff82111561432b5761432b6141f1565b5060051b60200190565b600082601f83011261434657600080fd5b8135602061435b61435683614311565b6142b2565b82815260059290921b8401810191818101908684111561437a57600080fd5b8286015b8481101561439e578035614391816140d3565b835291830191830161437e565b509695505050505050565b600082601f8301126143ba57600080fd5b813560206143ca61435683614311565b82815260059290921b840181019181810190868411156143e957600080fd5b8286015b8481101561439e57803583529183019183016143ed565b600082601f83011261441557600080fd5b813567ffffffffffffffff81111561442f5761442f6141f1565b61446060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016142b2565b81815284602083860101111561447557600080fd5b816020850160208301376000918101602001919091529392505050565b6000608082840312156144a457600080fd5b6144ac614220565b9050813581526020820135602082015260408201356144ca816140d3565b6040820152606082013567ffffffffffffffff808211156144ea57600080fd5b9083019060e082860312156144fe57600080fd5b614506614249565b82358281111561451557600080fd5b61452187828601614335565b82525060208301358281111561453657600080fd5b614542878286016143a9565b6020830152506040830135604082015261455e60608401614301565b606082015261456f60808401614301565b608082015260a08301358281111561458657600080fd5b61459287828601614404565b60a08301525060c0830135828111156145aa57600080fd5b6145b687828601614404565b60c083015250606084015250909392505050565b6000604082840312156145dc57600080fd5b6145e461426c565b9050813567ffffffffffffffff8111156145fd57600080fd5b8201601f8101841361460e57600080fd5b8035602061461e61435683614311565b82815260059290921b8301810191818101908784111561463d57600080fd5b938201935b8385101561465b57843582529382019390820190614642565b85525093840135938301939093525092915050565b801515811461336957600080fd5b60008060006060848603121561469357600080fd5b833567ffffffffffffffff808211156146ab57600080fd5b6146b787838801614492565b945060208601359150808211156146cd57600080fd5b506146da868287016145ca565b92505060408401356146eb81614670565b809150509250925092565b60008060006060848603121561470b57600080fd5b8335614716816140d3565b92506020840135614726816140d3565b929592945050506040919091013590565b6000806040838503121561474a57600080fd5b823567ffffffffffffffff8082111561476257600080fd5b61476e86838701614492565b9350602085013591508082111561478457600080fd5b50614791858286016145ca565b9150509250929050565b600081518084526020808501945080840160005b838110156147e157815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016147af565b509495945050505050565b602081526000613eda602083018461479b565b60008083601f84011261481157600080fd5b50813567ffffffffffffffff81111561482957600080fd5b6020830191508360208260051b850101111561484457600080fd5b9250929050565b60008060008060008060008060e0898b03121561486757600080fd5b606089018a81111561487857600080fd5b8998503567ffffffffffffffff8082111561489257600080fd5b818b0191508b601f8301126148a657600080fd5b8135818111156148b557600080fd5b8c60208285010111156148c757600080fd5b6020830199508098505060808b01359150808211156148e557600080fd5b6148f18c838d016147ff565b909750955060a08b013591508082111561490a57600080fd5b506149178b828c016147ff565b999c989b50969995989497949560c00135949350505050565b600082601f83011261494157600080fd5b8135602061495161435683614311565b82815260059290921b8401810191818101908684111561497057600080fd5b8286015b8481101561439e578035614987816140d3565b8352918301918301614974565b803560ff8116811461430c57600080fd5b803567ffffffffffffffff8116811461430c57600080fd5b60008060008060008060c087890312156149d657600080fd5b863567ffffffffffffffff808211156149ee57600080fd5b6149fa8a838b01614930565b97506020890135915080821115614a1057600080fd5b614a1c8a838b01614930565b9650614a2a60408a01614994565b95506060890135915080821115614a4057600080fd5b614a4c8a838b01614404565b9450614a5a60808a016149a5565b935060a0890135915080821115614a7057600080fd5b50614a7d89828a01614404565b9150509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015614acb57614acb614a8a565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060208284031215614b4057600080fd5b8151613eda81614670565b600060608284031215614b5d57600080fd5b614b6561428f565b8251815260208301516020820152604083015160408201528091505092915050565b8051825260006020808301518185015273ffffffffffffffffffffffffffffffffffffffff6040840151166040850152606083015160806060860152805160e06080870152614bda61016087018261479b565b838301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80888303810160a08a0152815180845291860193506000929091908601905b80841015614c3d5784518252938601936001939093019290860190614c1d565b50604085015160c08a0152606085015173ffffffffffffffffffffffffffffffffffffffff90811660e08b01526080860151166101008a015260a085015189820383016101208b01529550614c928187614142565b95505060c084015193508088860301610140890152505050614cb48282614142565b95945050505050565b60408152600060808201845160408085015281815180845260a08601915060209350838301925060005b81811015614d0357835183529284019291840191600101614ce7565b5050828701516060860152848103838601526140568187614b87565b60008219821115614d3257614d32614a8a565b500190565b600060208284031215614d4957600080fd5b5051919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614d8857614d88614a8a565b500290565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415614dbf57614dbf614a8a565b5060010190565b602081526000613eda6020830184614b87565b828152604060208201526000613ed76040830184614142565b7fff000000000000000000000000000000000000000000000000000000000000008316815260008251614e2c816001850160208701614112565b919091016001019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082614e7857614e78614e3a565b500690565b600082614e8c57614e8c614e3a565b500490565b600060ff821660ff84168060ff03821115614eae57614eae614a8a565b019392505050565b600060ff831680614ec957614ec9614e3a565b8060ff84160491505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600063ffffffff808316818516808303821115614f8157614f81614a8a565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614fba8184018a61479b565b90508281036080840152614fce818961479b565b905060ff871660a084015282810360c0840152614feb8187614142565b905067ffffffffffffffff851660e08401528281036101008401526150108185614142565b9c9b505050505050505050505050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526150678285018b61479b565b9150838203608085015261507b828a61479b565b915060ff881660a085015283820360c08501526150988288614142565b90861660e085015283810361010085015290506150108185614142565b600082516150c7818460208701614112565b919091019291505056fea164736f6c634300080c000a",
}

var OffRampABI = OffRampMetaData.ABI

var OffRampBin = OffRampMetaData.Bin

func DeployOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, sourceTokens []common.Address, pools []common.Address, feeds []common.Address, afn common.Address, maxTimeWithoutAFNSignal *big.Int, executionDelaySeconds *big.Int, maxTokensLength *big.Int, executionFeeLink *big.Int, maxDataSize *big.Int) (common.Address, *types.Transaction, *OffRamp, error) {
	parsed, err := OffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OffRampBin), backend, sourceChainId, chainId, sourceTokens, pools, feeds, afn, maxTimeWithoutAFNSignal, executionDelaySeconds, maxTokensLength, executionFeeLink, maxDataSize)
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

func (_OffRamp *OffRampCaller) GetExecutionDelaySeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getExecutionDelaySeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRamp *OffRampSession) GetExecutionDelaySeconds() (*big.Int, error) {
	return _OffRamp.Contract.GetExecutionDelaySeconds(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetExecutionDelaySeconds() (*big.Int, error) {
	return _OffRamp.Contract.GetExecutionDelaySeconds(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) GetExecutionFeeLink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getExecutionFeeLink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRamp *OffRampSession) GetExecutionFeeLink() (*big.Int, error) {
	return _OffRamp.Contract.GetExecutionFeeLink(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetExecutionFeeLink() (*big.Int, error) {
	return _OffRamp.Contract.GetExecutionFeeLink(&_OffRamp.CallOpts)
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

func (_OffRamp *OffRampCaller) GetMaxDataSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getMaxDataSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRamp *OffRampSession) GetMaxDataSize() (*big.Int, error) {
	return _OffRamp.Contract.GetMaxDataSize(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetMaxDataSize() (*big.Int, error) {
	return _OffRamp.Contract.GetMaxDataSize(&_OffRamp.CallOpts)
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

func (_OffRamp *OffRampTransactor) SetExecutionDelaySeconds(opts *bind.TransactOpts, executionDelaySeconds *big.Int) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "setExecutionDelaySeconds", executionDelaySeconds)
}

func (_OffRamp *OffRampSession) SetExecutionDelaySeconds(executionDelaySeconds *big.Int) (*types.Transaction, error) {
	return _OffRamp.Contract.SetExecutionDelaySeconds(&_OffRamp.TransactOpts, executionDelaySeconds)
}

func (_OffRamp *OffRampTransactorSession) SetExecutionDelaySeconds(executionDelaySeconds *big.Int) (*types.Transaction, error) {
	return _OffRamp.Contract.SetExecutionDelaySeconds(&_OffRamp.TransactOpts, executionDelaySeconds)
}

func (_OffRamp *OffRampTransactor) SetExecutionFeeLink(opts *bind.TransactOpts, executionFeeLink *big.Int) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "setExecutionFeeLink", executionFeeLink)
}

func (_OffRamp *OffRampSession) SetExecutionFeeLink(executionFeeLink *big.Int) (*types.Transaction, error) {
	return _OffRamp.Contract.SetExecutionFeeLink(&_OffRamp.TransactOpts, executionFeeLink)
}

func (_OffRamp *OffRampTransactorSession) SetExecutionFeeLink(executionFeeLink *big.Int) (*types.Transaction, error) {
	return _OffRamp.Contract.SetExecutionFeeLink(&_OffRamp.TransactOpts, executionFeeLink)
}

func (_OffRamp *OffRampTransactor) SetMaxDataSize(opts *bind.TransactOpts, maxDataSize *big.Int) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "setMaxDataSize", maxDataSize)
}

func (_OffRamp *OffRampSession) SetMaxDataSize(maxDataSize *big.Int) (*types.Transaction, error) {
	return _OffRamp.Contract.SetMaxDataSize(&_OffRamp.TransactOpts, maxDataSize)
}

func (_OffRamp *OffRampTransactorSession) SetMaxDataSize(maxDataSize *big.Int) (*types.Transaction, error) {
	return _OffRamp.Contract.SetMaxDataSize(&_OffRamp.TransactOpts, maxDataSize)
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

type OffRampExecutionDelaySecondsSetIterator struct {
	Event *OffRampExecutionDelaySecondsSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampExecutionDelaySecondsSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampExecutionDelaySecondsSet)
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
		it.Event = new(OffRampExecutionDelaySecondsSet)
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

func (it *OffRampExecutionDelaySecondsSetIterator) Error() error {
	return it.fail
}

func (it *OffRampExecutionDelaySecondsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampExecutionDelaySecondsSet struct {
	Delay *big.Int
	Raw   types.Log
}

func (_OffRamp *OffRampFilterer) FilterExecutionDelaySecondsSet(opts *bind.FilterOpts) (*OffRampExecutionDelaySecondsSetIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "ExecutionDelaySecondsSet")
	if err != nil {
		return nil, err
	}
	return &OffRampExecutionDelaySecondsSetIterator{contract: _OffRamp.contract, event: "ExecutionDelaySecondsSet", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchExecutionDelaySecondsSet(opts *bind.WatchOpts, sink chan<- *OffRampExecutionDelaySecondsSet) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "ExecutionDelaySecondsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampExecutionDelaySecondsSet)
				if err := _OffRamp.contract.UnpackLog(event, "ExecutionDelaySecondsSet", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseExecutionDelaySecondsSet(log types.Log) (*OffRampExecutionDelaySecondsSet, error) {
	event := new(OffRampExecutionDelaySecondsSet)
	if err := _OffRamp.contract.UnpackLog(event, "ExecutionDelaySecondsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampExecutionFeeLinkSetIterator struct {
	Event *OffRampExecutionFeeLinkSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampExecutionFeeLinkSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampExecutionFeeLinkSet)
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
		it.Event = new(OffRampExecutionFeeLinkSet)
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

func (it *OffRampExecutionFeeLinkSetIterator) Error() error {
	return it.fail
}

func (it *OffRampExecutionFeeLinkSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampExecutionFeeLinkSet struct {
	ExecutionFee *big.Int
	Raw          types.Log
}

func (_OffRamp *OffRampFilterer) FilterExecutionFeeLinkSet(opts *bind.FilterOpts) (*OffRampExecutionFeeLinkSetIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "ExecutionFeeLinkSet")
	if err != nil {
		return nil, err
	}
	return &OffRampExecutionFeeLinkSetIterator{contract: _OffRamp.contract, event: "ExecutionFeeLinkSet", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchExecutionFeeLinkSet(opts *bind.WatchOpts, sink chan<- *OffRampExecutionFeeLinkSet) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "ExecutionFeeLinkSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampExecutionFeeLinkSet)
				if err := _OffRamp.contract.UnpackLog(event, "ExecutionFeeLinkSet", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseExecutionFeeLinkSet(log types.Log) (*OffRampExecutionFeeLinkSet, error) {
	event := new(OffRampExecutionFeeLinkSet)
	if err := _OffRamp.contract.UnpackLog(event, "ExecutionFeeLinkSet", log); err != nil {
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

type OffRampMaxDataSizeSetIterator struct {
	Event *OffRampMaxDataSizeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampMaxDataSizeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampMaxDataSizeSet)
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
		it.Event = new(OffRampMaxDataSizeSet)
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

func (it *OffRampMaxDataSizeSetIterator) Error() error {
	return it.fail
}

func (it *OffRampMaxDataSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampMaxDataSizeSet struct {
	Size *big.Int
	Raw  types.Log
}

func (_OffRamp *OffRampFilterer) FilterMaxDataSizeSet(opts *bind.FilterOpts) (*OffRampMaxDataSizeSetIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "MaxDataSizeSet")
	if err != nil {
		return nil, err
	}
	return &OffRampMaxDataSizeSetIterator{contract: _OffRamp.contract, event: "MaxDataSizeSet", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchMaxDataSizeSet(opts *bind.WatchOpts, sink chan<- *OffRampMaxDataSizeSet) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "MaxDataSizeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampMaxDataSizeSet)
				if err := _OffRamp.contract.UnpackLog(event, "MaxDataSizeSet", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseMaxDataSizeSet(log types.Log) (*OffRampMaxDataSizeSet, error) {
	event := new(OffRampMaxDataSizeSet)
	if err := _OffRamp.contract.UnpackLog(event, "MaxDataSizeSet", log); err != nil {
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
	case _OffRamp.abi.Events["ExecutionDelaySecondsSet"].ID:
		return _OffRamp.ParseExecutionDelaySecondsSet(log)
	case _OffRamp.abi.Events["ExecutionFeeLinkSet"].ID:
		return _OffRamp.ParseExecutionFeeLinkSet(log)
	case _OffRamp.abi.Events["FeedAdded"].ID:
		return _OffRamp.ParseFeedAdded(log)
	case _OffRamp.abi.Events["FeedRemoved"].ID:
		return _OffRamp.ParseFeedRemoved(log)
	case _OffRamp.abi.Events["FeesWithdrawn"].ID:
		return _OffRamp.ParseFeesWithdrawn(log)
	case _OffRamp.abi.Events["MaxDataSizeSet"].ID:
		return _OffRamp.ParseMaxDataSizeSet(log)
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

func (OffRampExecutionDelaySecondsSet) Topic() common.Hash {
	return common.HexToHash("0xfbb92e09fd0b7bcabafe655657e745c10bcd9812c1e32e205e477a3c5ff74fc0")
}

func (OffRampExecutionFeeLinkSet) Topic() common.Hash {
	return common.HexToHash("0xd7fd8173fe9fe63cbd28468ae5012d7f8db1f1550fbfe8c22a1b49e4b0b703ae")
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

func (OffRampMaxDataSizeSet) Topic() common.Hash {
	return common.HexToHash("0x15e8871b18d0879e199f8794daad195ca64bbb7341d01951a980a1abf0800bfb")
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

	FilterExecutionDelaySecondsSet(opts *bind.FilterOpts) (*OffRampExecutionDelaySecondsSetIterator, error)

	WatchExecutionDelaySecondsSet(opts *bind.WatchOpts, sink chan<- *OffRampExecutionDelaySecondsSet) (event.Subscription, error)

	ParseExecutionDelaySecondsSet(log types.Log) (*OffRampExecutionDelaySecondsSet, error)

	FilterExecutionFeeLinkSet(opts *bind.FilterOpts) (*OffRampExecutionFeeLinkSetIterator, error)

	WatchExecutionFeeLinkSet(opts *bind.WatchOpts, sink chan<- *OffRampExecutionFeeLinkSet) (event.Subscription, error)

	ParseExecutionFeeLinkSet(log types.Log) (*OffRampExecutionFeeLinkSet, error)

	FilterFeedAdded(opts *bind.FilterOpts) (*OffRampFeedAddedIterator, error)

	WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *OffRampFeedAdded) (event.Subscription, error)

	ParseFeedAdded(log types.Log) (*OffRampFeedAdded, error)

	FilterFeedRemoved(opts *bind.FilterOpts) (*OffRampFeedRemovedIterator, error)

	WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *OffRampFeedRemoved) (event.Subscription, error)

	ParseFeedRemoved(log types.Log) (*OffRampFeedRemoved, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*OffRampFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *OffRampFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*OffRampFeesWithdrawn, error)

	FilterMaxDataSizeSet(opts *bind.FilterOpts) (*OffRampMaxDataSizeSetIterator, error)

	WatchMaxDataSizeSet(opts *bind.WatchOpts, sink chan<- *OffRampMaxDataSizeSet) (event.Subscription, error)

	ParseMaxDataSizeSet(log types.Log) (*OffRampMaxDataSizeSet, error)

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
