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

type OffRampInterfaceOffRampConfig struct {
	ExecutionFeeJuels     uint64
	ExecutionDelaySeconds uint64
	MaxDataSize           uint64
	MaxTokensLength       uint64
}

var OffRampHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"contractAggregatorV2V3Interface[]\",\"name\":\"feeds\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionDelayError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"InvalidExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceFeedConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MerkleProofError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RelayReportError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lastMaxSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMinSequenceNumber\",\"type\":\"uint256\"}],\"name\":\"SequenceError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenFeedMistmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMistmatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"UnexpectedPayloadData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"CrossChainMessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"minSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSequenceNumber\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCCIP.RelayReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"ReportAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"addFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"needFee\",\"type\":\"bool\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"getExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastReport\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"minSequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSequenceNumber\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.RelayReport\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRampConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"removeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"merkle\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setOffRampConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005ba238038062005ba2833981016040819052620000349162000704565b6040805160808101825260018082526001600160401b0385811660208401526103e89383019390935291831660608201526000805460ff191681558b928b928b928b928b928b928b92908790869082908990889088903390819081620000e15760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b0319909216919091179091558116156200011b576200011b8162000434565b5050506001600160a01b038216158062000133575080155b156200015257604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001985760405162d8548360e71b815260040160405180910390fd5b8151620001ad906005906020850190620004e6565b5060005b825181101562000291576000828281518110620001d257620001d2620007ec565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b0316815250600460008685815181106200021c576200021c620007ec565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff1916600117905580620002888162000802565b915050620001b1565b5050508051825114620002b75760405163ee9d106b60e01b815260040160405180910390fd5b8151620002cc906008906020850190620004e6565b5060005b825181101562000399576000828281518110620002f157620002f1620007ec565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b0316815250600760008685815181106200033b576200033b620007ec565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b02929091169190911790555080620003908162000802565b915050620002d0565b505050151560805260a09790975250505060c0929092525050805160148054602084015160408501516060909501516001600160401b03908116600160c01b026001600160c01b03968216600160801b02969096166001600160801b0392821668010000000000000000026001600160801b03199094169190951617919091171691909117919091179055506200082c975050505050505050565b6001600160a01b0381163314156200048f5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000d8565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200053e579160200282015b828111156200053e57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000507565b506200054c92915062000550565b5090565b5b808211156200054c576000815560010162000551565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620005a857620005a862000567565b604052919050565b60006001600160401b03821115620005cc57620005cc62000567565b5060051b60200190565b6001600160a01b0381168114620005ec57600080fd5b50565b600082601f8301126200060157600080fd5b815160206200061a6200061483620005b0565b6200057d565b82815260059290921b840181019181810190868411156200063a57600080fd5b8286015b84811015620006625780516200065481620005d6565b83529183019183016200063e565b509695505050505050565b600082601f8301126200067f57600080fd5b81516020620006926200061483620005b0565b82815260059290921b84018101918181019086841115620006b257600080fd5b8286015b8481101562000662578051620006cc81620005d6565b8352918301918301620006b6565b8051620006e781620005d6565b919050565b80516001600160401b0381168114620006e757600080fd5b60008060008060008060008060006101208a8c0312156200072457600080fd5b895160208b015160408c0151919a5098506001600160401b03808211156200074b57600080fd5b620007598d838e01620005ef565b985060608c01519150808211156200077057600080fd5b6200077e8d838e01620005ef565b975060808c01519150808211156200079557600080fd5b50620007a48c828d016200066d565b955050620007b560a08b01620006da565b935060c08a01519250620007cc60e08b01620006ec565b9150620007dd6101008b01620006ec565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b60006000198214156200082557634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05161533f6200086360003960006104d101526000818161042c015261341a01526000612164015261533f6000f3fe608060405234801561001057600080fd5b506004361061025c5760003560e01c806381be8fa411610145578063b034909c116100bd578063bbe4f6db1161008c578063eb511dd411610071578063eb511dd4146106c9578063eefa7a3e146106dc578063f2fde38b1461073257600080fd5b8063bbe4f6db1461067d578063e3d0e712146106b657600080fd5b8063b034909c1461063c578063b1dc65a414610644578063b576716614610657578063b6608c3b1461066a57600080fd5b806389c0656811610114578063a7206cd6116100f9578063a7206cd61461051e578063a8ebd0f41461053e578063afcb95d71461061c57600080fd5b806389c06568146104f35780638da5cb5b146104fb57600080fd5b806381be8fa41461048c57806381ff7048146104945780638456cb59146104c457806385e1f4d0146104cc57600080fd5b80635853c627116101d8578063744b92e2116101a7578063768c577b1161018c578063768c577b1461045c57806379ba50971461046f578063814118341461047757600080fd5b8063744b92e21461041457806374be21501461042757600080fd5b80635853c627146103aa57806359e96b5b146103bd5780635b16ebb7146103d05780635c975abb1461040957600080fd5b80632222dd421161022f5780633b8d08ef116102145780633b8d08ef1461037c5780633f4ba83a1461038f578063461c551b1461039757600080fd5b80632222dd421461034b5780632b898c251461036957600080fd5b8063108ee5fc1461026157806316b8e73114610276578063181f5a77146102d95780632194750714610318575b600080fd5b61027461026f36600461417d565b610745565b005b6102af61028436600461417d565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b604080518082018252600d81527f4f666652616d7020302e302e3100000000000000000000000000000000000000602082015290516102d09190614214565b61033b610326366004614227565b60009081526010602052604090205460ff1690565b60405190151581526020016102d0565b60025473ffffffffffffffffffffffffffffffffffffffff166102af565b610274610377366004614240565b610821565b61027461038a366004614706565b610bed565b610274611535565b6102746103a536600461477e565b611547565b6102746103b8366004614240565b611599565b6102746103cb366004614796565b6117b1565b61033b6103de36600461417d565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff1661033b565b610274610422366004614240565b61182f565b61044e7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016102d0565b61044e61046a3660046147d7565b611c20565b610274611dac565b61047f611ece565b6040516102d0919061488c565b61047f611f3d565b600b546009546040805163ffffffff808516825264010000000090940490931660208401528201526060016102d0565b610274611faa565b61044e7f000000000000000000000000000000000000000000000000000000000000000081565b61047f611fba565b600054610100900473ffffffffffffffffffffffffffffffffffffffff166102af565b61044e61052c366004614227565b6000908152600f602052604090205490565b6105d8604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260145467ffffffffffffffff808216835268010000000000000000820481166020840152700100000000000000000000000000000000820481169383019390935278010000000000000000000000000000000000000000000000009004909116606082015290565b6040516102d09190815167ffffffffffffffff9081168252602080840151821690830152604080840151821690830152606092830151169181019190915260800190565b6040805160018152600060208201819052918101919091526060016102d0565b60035461044e565b6102746106523660046148eb565b612027565b6102746106653660046149d0565b6126d0565b610274610678366004614227565b6126df565b6102af61068b36600461417d565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6102746106c4366004614aa3565b61275c565b6102746106d7366004614240565b613141565b6040805160608082018352600080835260208084018290529284015282518082018452601154808252601254828501908152601354928601928352855191825251938101939093525192820192909252016102d0565b61027461074036600461417d565b613381565b61074d613392565b73ffffffffffffffffffffffffffffffffffffffff811661079a576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b610829613392565b60085480610863576040517f2e70248b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906108fe576040517f3917193900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610967576040517f6c17b98700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006008610976600185614b9f565b8154811061098657610986614bb6565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600882602001516bffffffffffffffffffffffff16815481106109d8576109d8614bb6565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166008610a07600186614b9f565b81548110610a1757610a17614bb6565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600883602001516bffffffffffffffffffffffff1681548110610a8557610a85614bb6565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526007909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556008805480610b2757610b27614be5565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600783526040808520949094558351908152908716918101919091527fa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d91015b60405180910390a15050505050565b60005460ff1615610c5f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610cce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cf29190614c14565b15610d28576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b4916004808301926060929190829003018187875af1158015610d99573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dbd9190614c31565b9050600354816020015142610dd29190614b9f565b1115610e0a576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610e168585611c20565b6000818152600f602052604090205490915080610e635784866040517f07e6809a000000000000000000000000000000000000000000000000000000008152600401610c56929190614da3565b6014544290610e889068010000000000000000900467ffffffffffffffff1683614e05565b10610ebf576040517f15c33ba200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855160009081526010602052604090205460ff1615610f105785516040517f6a64e9610000000000000000000000000000000000000000000000000000000081526004810191909152602401610c56565b60608601516080015173ffffffffffffffffffffffffffffffffffffffff1615801590610f5b575060608601516080015173ffffffffffffffffffffffffffffffffffffffff163314155b15610f985785516040517fd8e90b980000000000000000000000000000000000000000000000000000000081526004810191909152602401610c56565b610fa186613418565b610faa86613593565b8551600090815260106020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905583156112005760008087606001516000015160008151811061100757611007614bb6565b6020026020010151905060006110428273ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff8116611091576040517f83135fec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166350d25bcd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111009190614e1d565b601454611117919067ffffffffffffffff16614e36565b92508289606001516020015160008151811061113557611135614bb6565b602002602001018181516111499190614b9f565b90525073ffffffffffffffffffffffffffffffffffffffff808316600090815260046020526040902054166040517fea6192a20000000000000000000000000000000000000000000000000000000081523360048201526024810185905273ffffffffffffffffffffffffffffffffffffffff919091169063ea6192a290604401600060405180830381600087803b1580156111e457600080fd5b505af11580156111f8573d6000803e3d6000fd5b505050505050505b60005b606087015151518110156113ba576000611263886060015160000151838151811061123057611230614bb6565b602002602001015173ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff81166112ef5760608801515180518390811061129857611298614bb6565b60200260200101516040517fbf16aab6000000000000000000000000000000000000000000000000000000008152600401610c56919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b8073ffffffffffffffffffffffffffffffffffffffff1663ea6192a28960600151606001518a6060015160200151858151811061132e5761132e614bb6565b60200260200101516040518363ffffffff1660e01b815260040161137492919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600060405180830381600087803b15801561138e57600080fd5b505af11580156113a2573d6000803e3d6000fd5b505050505080806113b290614e73565b915050611203565b50606080870151015173ffffffffffffffffffffffffffffffffffffffff163b156114e65785606001516060015173ffffffffffffffffffffffffffffffffffffffff16639c598468876040518263ffffffff1660e01b81526004016114209190614eac565b600060405180830381600087803b15801561143a57600080fd5b505af192505050801561144b575060015b6114b6573d808015611479576040519150601f19603f3d011682016040523d82523d6000602084013e61147e565b606091505b5086516040517f6a3fd4f2000000000000000000000000000000000000000000000000000000008152610c5691908390600401614ebf565b85516040517fc51bf0f6d90b467e0849da0ad18a4d9144a4b78b9f83202e1c65cd68f72d4a3890600090a261152d565b606086015160a00151511561152d5785516040517fe0244be30000000000000000000000000000000000000000000000000000000081526004810191909152602401610c56565b505050505050565b61153d613392565b611545613642565b565b61154f613392565b80601461155c8282614ed8565b9050507fe4cd88b1b5e20a0b843af3207ba74d3f84af4acff45830469490ac9c6ab8a7458160405161158e9190614fe9565b60405180910390a150565b6115a1613392565b73ffffffffffffffffffffffffffffffffffffffff821615806115d8575073ffffffffffffffffffffffffffffffffffffffff8116155b1561160f576040517fee9d106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156116ab576040517f965ffb7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600880546bffffffffffffffffffffffff908116602080870191825288861660008181526007835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821790558351908152908101919091527f037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb60130591015b60405180910390a1505050565b6117b9613392565b6117da73ffffffffffffffffffffffffffffffffffffffff84168383613723565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8906060016117a4565b611837613392565b60055480611871576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529061190c576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614611975576040517fd428911900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006005611984600185614b9f565b8154811061199457611994614bb6565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff16815481106119e6576119e6614bb6565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166005611a15600186614b9f565b81548110611a2557611a25614bb6565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110611a9357611a93614bb6565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480611b3557611b35614be5565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101610bde565b600080600060f81b84604051602001611c399190614eac565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052611c759291602001615053565b60405160208183030381529060405280519060200120905060005b835151811015611da457600084600001518281518110611cb257611cb2614bb6565b6020026020010151905060028560200151611ccd91906150ca565b611d28576040517f010000000000000000000000000000000000000000000000000000000000000060208201526021810184905260418101829052606101604051602081830303815290604052805190602001209250611d7b565b6040517f0100000000000000000000000000000000000000000000000000000000000000602082015260218101829052604181018490526061016040516020818303038152906040528051906020012092505b60028560200151611d8c91906150de565b60208601525080611d9c81614e73565b915050611c90565b509392505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611e2d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610c56565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600e805480602002602001604051908101604052809291908181526020018280548015611f3357602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611f08575b5050505050905090565b60606008805480602002602001604051908101604052809291908181526020018280548015611f335760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611f08575050505050905090565b611fb2613392565b6115456137b5565b60606005805480602002602001604051908101604052809291908181526020018280548015611f335760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611f08575050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161207d91849163ffffffff851691908e908e908190840183828082843760009201919091525061387592505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600954808252600a5460ff80821660208501526101009091041692820192909252908314612152576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610c56565b6121608b8b8b8b8b8b613be6565b60007f0000000000000000000000000000000000000000000000000000000000000000156121bd5760028260200151836040015161219e91906150f2565b6121a89190615117565b6121b39060016150f2565b60ff1690506121d3565b60208201516121cd9060016150f2565b60ff1690505b88811461223c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610c56565b8887146122a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610c56565b336000908152600c602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156122e8576122e8615139565b60028111156122f9576122f9615139565b905250905060028160200151600281111561231657612316615139565b14801561235d5750600e816000015160ff168154811061233857612338614bb6565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b6123c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610c56565b5050505050600088886040516123da929190615168565b6040519081900381206123f1918c90602001615178565b60405160208183030381529060405280519060200120905061241161413c565b604080518082019091526000808252602082015260005b888110156126ae57600060018588846020811061244757612447614bb6565b61245491901a601b6150f2565b8d8d8681811061246657612466614bb6565b905060200201358c8c8781811061247f5761247f614bb6565b90506020020135604051600081526020016040526040516124bc949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156124de573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600c602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561255e5761255e615139565b600281111561256f5761256f615139565b905250925060018360200151600281111561258c5761258c615139565b146125f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610c56565b8251849060ff16601f811061260a5761260a614bb6565b602002015115612676576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610c56565b600184846000015160ff16601f811061269157612691614bb6565b9115156020909202015250806126a681614e73565b915050612428565b5050505063ffffffff81106126c5576126c5615194565b505050505050505050565b6126dc60008083613875565b50565b6126e7613392565b8061271e576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c2519101610815565b855185518560ff16601f8311156127cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610c56565b60008111612839576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610c56565b8183146128c7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610c56565b6128d2816003614e36565b831161293a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610c56565b612942613392565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600d5415612b3557600d5460009061299a90600190614b9f565b90506000600d82815481106129b1576129b1614bb6565b6000918252602082200154600e805473ffffffffffffffffffffffffffffffffffffffff909216935090849081106129eb576129eb614bb6565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600c909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600d80549192509080612a6b57612a6b614be5565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600e805480612ad457612ad4614be5565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550612980915050565b60005b815151811015612f9c576000600c600084600001518481518110612b5e57612b5e614bb6565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612ba857612ba8615139565b14612c0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610c56565b6040805180820190915260ff821681526001602082015282518051600c9160009185908110612c4057612c40614bb6565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612ce157612ce1615139565b021790555060009150612cf19050565b600c600084602001518481518110612d0b57612d0b614bb6565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612d5557612d55615139565b14612dbc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610c56565b6040805180820190915260ff82168152602081016002815250600c600084602001518481518110612def57612def614bb6565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612e9057612e90615139565b02179055505082518051600d925083908110612eae57612eae614bb6565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600e919083908110612f2a57612f2a614bb6565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691909117905580612f9481614e73565b915050612b38565b506040810151600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600b80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff43811682029283178555908304811693600193909260009261302e9286929082169116176151c3565b92506101000a81548163ffffffff021916908363ffffffff16021790555061308d4630600b60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613c9d565b600981905582518051600a805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600b5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e059861312c988b98919763ffffffff9092169690959194919391926151eb565b60405180910390a15050505050505050505050565b613149613392565b73ffffffffffffffffffffffffffffffffffffffff82161580613180575073ffffffffffffffffffffffffffffffffffffffff8116155b156131b7576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015613253576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c91016117a4565b613389613392565b6126dc81613d48565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314611545576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610c56565b7f000000000000000000000000000000000000000000000000000000000000000081602001511461347d5780602001516040517fd44bc9eb000000000000000000000000000000000000000000000000000000008152600401610c5691815260200190565b60145460608201515151780100000000000000000000000000000000000000000000000090910467ffffffffffffffff1610806134c65750606081015160208101515190515114155b156134fd576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601454606082015160a001515170010000000000000000000000000000000090910467ffffffffffffffff1610156126dc57601454606082015160a00151516040517f8693378900000000000000000000000000000000000000000000000000000000815270010000000000000000000000000000000090920467ffffffffffffffff1660048301526024820152604401610c56565b606080820151015173ffffffffffffffffffffffffffffffffffffffff163014806135ea5750606080820151015173ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff165b156126dc5760608082015101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610c56565b60005460ff166136ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610c56565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526137b0908490613e44565b505050565b60005460ff1615613822576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c56565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586136f93390565b60005460ff16156138e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c56565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015613951573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139759190614c14565b156139ab576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b4916004808301926060929190829003018187875af1158015613a1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a409190614c31565b9050600354816020015142613a559190614b9f565b1115613a8d576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082806020019051810190613aa39190614c31565b9050806040015181602001511115613ae7576040517f67a3824c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516060810182526011548082526012546020830152601354928201929092529015613b70576040810151613b1f906001614e05565b826020015114613b7057806040015182602001516040517fcc7f1bd0000000000000000000000000000000000000000000000000000000008152600401610c56929190918252602082015260400190565b81516000908152600f60209081526040918290204290558351601181905581850180516012558386018051601355845192835290519282019290925290518183015290517f07d7bce06be2a7b0230e4dd0d72523c0407e82419ab1d947c5ddaf59ca36484e9181900360600190a1505050505050565b6000613bf3826020614e36565b613bfe856020614e36565b613c0a88610144614e05565b613c149190614e05565b613c1e9190614e05565b613c29906000614e05565b9050368114613c94576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610c56565b50505050505050565b6000808a8a8a8a8a8a8a8a8a604051602001613cc199989796959493929190615281565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116331415613dc8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610c56565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000613ea6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16613f509092919063ffffffff16565b8051909150156137b05780806020019051810190613ec49190614c14565b6137b0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610c56565b6060613f5f8484600085613f69565b90505b9392505050565b606082471015613ffb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610c56565b843b614063576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610c56565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161408c9190615316565b60006040518083038185875af1925050503d80600081146140c9576040519150601f19603f3d011682016040523d82523d6000602084013e6140ce565b606091505b50915091506140de8282866140e9565b979650505050505050565b606083156140f8575081613f62565b8251156141085782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c569190614214565b604051806103e00160405280601f906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff811681146126dc57600080fd5b60006020828403121561418f57600080fd5b8135613f628161415b565b60005b838110156141b557818101518382015260200161419d565b838111156141c4576000848401525b50505050565b600081518084526141e281602086016020860161419a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613f6260208301846141ca565b60006020828403121561423957600080fd5b5035919050565b6000806040838503121561425357600080fd5b823561425e8161415b565b9150602083013561426e8161415b565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff811182821017156142cb576142cb614279565b60405290565b60405160e0810167ffffffffffffffff811182821017156142cb576142cb614279565b6040805190810167ffffffffffffffff811182821017156142cb576142cb614279565b6040516060810167ffffffffffffffff811182821017156142cb576142cb614279565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561438157614381614279565b604052919050565b80356143948161415b565b919050565b600067ffffffffffffffff8211156143b3576143b3614279565b5060051b60200190565b600082601f8301126143ce57600080fd5b813560206143e36143de83614399565b61433a565b82815260059290921b8401810191818101908684111561440257600080fd5b8286015b848110156144265780356144198161415b565b8352918301918301614406565b509695505050505050565b600082601f83011261444257600080fd5b813560206144526143de83614399565b82815260059290921b8401810191818101908684111561447157600080fd5b8286015b848110156144265780358352918301918301614475565b600082601f83011261449d57600080fd5b813567ffffffffffffffff8111156144b7576144b7614279565b6144e860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161433a565b8181528460208386010111156144fd57600080fd5b816020850160208301376000918101602001919091529392505050565b60006080828403121561452c57600080fd5b6145346142a8565b9050813581526020820135602082015260408201356145528161415b565b6040820152606082013567ffffffffffffffff8082111561457257600080fd5b9083019060e0828603121561458657600080fd5b61458e6142d1565b82358281111561459d57600080fd5b6145a9878286016143bd565b8252506020830135828111156145be57600080fd5b6145ca87828601614431565b602083015250604083013560408201526145e660608401614389565b60608201526145f760808401614389565b608082015260a08301358281111561460e57600080fd5b61461a8782860161448c565b60a08301525060c08301358281111561463257600080fd5b61463e8782860161448c565b60c083015250606084015250909392505050565b60006040828403121561466457600080fd5b61466c6142f4565b9050813567ffffffffffffffff81111561468557600080fd5b8201601f8101841361469657600080fd5b803560206146a66143de83614399565b82815260059290921b830181019181810190878411156146c557600080fd5b938201935b838510156146e3578435825293820193908201906146ca565b85525093840135938301939093525092915050565b80151581146126dc57600080fd5b60008060006060848603121561471b57600080fd5b833567ffffffffffffffff8082111561473357600080fd5b61473f8783880161451a565b9450602086013591508082111561475557600080fd5b5061476286828701614652565b9250506040840135614773816146f8565b809150509250925092565b60006080828403121561479057600080fd5b50919050565b6000806000606084860312156147ab57600080fd5b83356147b68161415b565b925060208401356147c68161415b565b929592945050506040919091013590565b600080604083850312156147ea57600080fd5b823567ffffffffffffffff8082111561480257600080fd5b61480e8683870161451a565b9350602085013591508082111561482457600080fd5b5061483185828601614652565b9150509250929050565b600081518084526020808501945080840160005b8381101561488157815173ffffffffffffffffffffffffffffffffffffffff168752958201959082019060010161484f565b509495945050505050565b602081526000613f62602083018461483b565b60008083601f8401126148b157600080fd5b50813567ffffffffffffffff8111156148c957600080fd5b6020830191508360208260051b85010111156148e457600080fd5b9250929050565b60008060008060008060008060e0898b03121561490757600080fd5b606089018a81111561491857600080fd5b8998503567ffffffffffffffff8082111561493257600080fd5b818b0191508b601f83011261494657600080fd5b81358181111561495557600080fd5b8c602082850101111561496757600080fd5b6020830199508098505060808b013591508082111561498557600080fd5b6149918c838d0161489f565b909750955060a08b01359150808211156149aa57600080fd5b506149b78b828c0161489f565b999c989b50969995989497949560c00135949350505050565b6000602082840312156149e257600080fd5b813567ffffffffffffffff8111156149f957600080fd5b614a058482850161448c565b949350505050565b600082601f830112614a1e57600080fd5b81356020614a2e6143de83614399565b82815260059290921b84018101918181019086841115614a4d57600080fd5b8286015b84811015614426578035614a648161415b565b8352918301918301614a51565b803560ff8116811461439457600080fd5b67ffffffffffffffff811681146126dc57600080fd5b803561439481614a82565b60008060008060008060c08789031215614abc57600080fd5b863567ffffffffffffffff80821115614ad457600080fd5b614ae08a838b01614a0d565b97506020890135915080821115614af657600080fd5b614b028a838b01614a0d565b9650614b1060408a01614a71565b95506060890135915080821115614b2657600080fd5b614b328a838b0161448c565b9450614b4060808a01614a98565b935060a0890135915080821115614b5657600080fd5b50614b6389828a0161448c565b9150509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015614bb157614bb1614b70565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060208284031215614c2657600080fd5b8151613f62816146f8565b600060608284031215614c4357600080fd5b614c4b614317565b8251815260208301516020820152604083015160408201528091505092915050565b8051825260006020808301518185015273ffffffffffffffffffffffffffffffffffffffff6040840151166040850152606083015160806060860152805160e06080870152614cc061016087018261483b565b838301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80888303810160a08a0152815180845291860193506000929091908601905b80841015614d235784518252938601936001939093019290860190614d03565b50604085015160c08a0152606085015173ffffffffffffffffffffffffffffffffffffffff90811660e08b01526080860151166101008a015260a085015189820383016101208b01529550614d7881876141ca565b95505060c084015193508088860301610140890152505050614d9a82826141ca565b95945050505050565b60408152600060808201845160408085015281815180845260a08601915060209350838301925060005b81811015614de957835183529284019291840191600101614dcd565b5050828701516060860152848103838601526140de8187614c6d565b60008219821115614e1857614e18614b70565b500190565b600060208284031215614e2f57600080fd5b5051919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614e6e57614e6e614b70565b500290565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415614ea557614ea5614b70565b5060010190565b602081526000613f626020830184614c6d565b828152604060208201526000613f5f60408301846141ca565b8135614ee381614a82565b67ffffffffffffffff811690508154817fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000082161783556020840135614f2781614a82565b6fffffffffffffffff00000000000000008160401b16905080837fffffffffffffffffffffffffffffffff000000000000000000000000000000008416171784556040850135614f7681614a82565b77ffffffffffffffff000000000000000000000000000000008160801b1690507fffffffffffffffff000000000000000000000000000000000000000000000000818582861617841717865560608701359350614fd284614a82565b808460c01b16858417831717865550505050505050565b608081018235614ff881614a82565b67ffffffffffffffff908116835260208401359061501582614a82565b908116602084015260408401359061502c82614a82565b908116604084015260608401359061504382614a82565b8082166060850152505092915050565b7fff00000000000000000000000000000000000000000000000000000000000000831681526000825161508d81600185016020870161419a565b919091016001019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826150d9576150d961509b565b500690565b6000826150ed576150ed61509b565b500490565b600060ff821660ff84168060ff0382111561510f5761510f614b70565b019392505050565b600060ff83168061512a5761512a61509b565b8060ff84160491505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600063ffffffff8083168185168083038211156151e2576151e2614b70565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b1660408501525080606084015261521b8184018a61483b565b9050828103608084015261522f818961483b565b905060ff871660a084015282810360c084015261524c81876141ca565b905067ffffffffffffffff851660e084015282810361010084015261527181856141ca565b9c9b505050505050505050505050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526152c88285018b61483b565b915083820360808501526152dc828a61483b565b915060ff881660a085015283820360c08501526152f982886141ca565b90861660e0850152838103610100850152905061527181856141ca565b6000825161532881846020870161419a565b919091019291505056fea164736f6c634300080c000a",
}

var OffRampHelperABI = OffRampHelperMetaData.ABI

var OffRampHelperBin = OffRampHelperMetaData.Bin

func DeployOffRampHelper(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, sourceTokens []common.Address, pools []common.Address, feeds []common.Address, afn common.Address, maxTimeWithoutAFNSignal *big.Int, executionDelaySeconds uint64, maxTokensLength uint64) (common.Address, *types.Transaction, *OffRampHelper, error) {
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

func (_OffRampHelper *OffRampHelperCaller) GetOffRampConfig(opts *bind.CallOpts) (OffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _OffRampHelper.contract.Call(opts, &out, "getOffRampConfig")

	if err != nil {
		return *new(OffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OffRampInterfaceOffRampConfig)).(*OffRampInterfaceOffRampConfig)

	return out0, err

}

func (_OffRampHelper *OffRampHelperSession) GetOffRampConfig() (OffRampInterfaceOffRampConfig, error) {
	return _OffRampHelper.Contract.GetOffRampConfig(&_OffRampHelper.CallOpts)
}

func (_OffRampHelper *OffRampHelperCallerSession) GetOffRampConfig() (OffRampInterfaceOffRampConfig, error) {
	return _OffRampHelper.Contract.GetOffRampConfig(&_OffRampHelper.CallOpts)
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

func (_OffRampHelper *OffRampHelperTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_OffRampHelper *OffRampHelperSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_OffRampHelper.TransactOpts, newTime)
}

func (_OffRampHelper *OffRampHelperTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_OffRampHelper.TransactOpts, newTime)
}

func (_OffRampHelper *OffRampHelperTransactor) SetOffRampConfig(opts *bind.TransactOpts, config OffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _OffRampHelper.contract.Transact(opts, "setOffRampConfig", config)
}

func (_OffRampHelper *OffRampHelperSession) SetOffRampConfig(config OffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetOffRampConfig(&_OffRampHelper.TransactOpts, config)
}

func (_OffRampHelper *OffRampHelperTransactorSession) SetOffRampConfig(config OffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _OffRampHelper.Contract.SetOffRampConfig(&_OffRampHelper.TransactOpts, config)
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

type OffRampHelperOffRampConfigSetIterator struct {
	Event *OffRampHelperOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampHelperOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampHelperOffRampConfigSet)
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
		it.Event = new(OffRampHelperOffRampConfigSet)
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

func (it *OffRampHelperOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampHelperOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampHelperOffRampConfigSet struct {
	Config OffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_OffRampHelper *OffRampHelperFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*OffRampHelperOffRampConfigSetIterator, error) {

	logs, sub, err := _OffRampHelper.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffRampHelperOffRampConfigSetIterator{contract: _OffRampHelper.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRampHelper *OffRampHelperFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampHelperOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffRampHelper.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampHelperOffRampConfigSet)
				if err := _OffRampHelper.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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

func (_OffRampHelper *OffRampHelperFilterer) ParseOffRampConfigSet(log types.Log) (*OffRampHelperOffRampConfigSet, error) {
	event := new(OffRampHelperOffRampConfigSet)
	if err := _OffRampHelper.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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
	case _OffRampHelper.abi.Events["FeedAdded"].ID:
		return _OffRampHelper.ParseFeedAdded(log)
	case _OffRampHelper.abi.Events["FeedRemoved"].ID:
		return _OffRampHelper.ParseFeedRemoved(log)
	case _OffRampHelper.abi.Events["FeesWithdrawn"].ID:
		return _OffRampHelper.ParseFeesWithdrawn(log)
	case _OffRampHelper.abi.Events["OffRampConfigSet"].ID:
		return _OffRampHelper.ParseOffRampConfigSet(log)
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

func (OffRampHelperFeedAdded) Topic() common.Hash {
	return common.HexToHash("0x037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb601305")
}

func (OffRampHelperFeedRemoved) Topic() common.Hash {
	return common.HexToHash("0xa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d")
}

func (OffRampHelperFeesWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8")
}

func (OffRampHelperOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xe4cd88b1b5e20a0b843af3207ba74d3f84af4acff45830469490ac9c6ab8a745")
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

	Report(opts *bind.TransactOpts, merkle []byte) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error)

	SetOffRampConfig(opts *bind.TransactOpts, config OffRampInterfaceOffRampConfig) (*types.Transaction, error)

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

	FilterFeedAdded(opts *bind.FilterOpts) (*OffRampHelperFeedAddedIterator, error)

	WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *OffRampHelperFeedAdded) (event.Subscription, error)

	ParseFeedAdded(log types.Log) (*OffRampHelperFeedAdded, error)

	FilterFeedRemoved(opts *bind.FilterOpts) (*OffRampHelperFeedRemovedIterator, error)

	WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *OffRampHelperFeedRemoved) (event.Subscription, error)

	ParseFeedRemoved(log types.Log) (*OffRampHelperFeedRemoved, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*OffRampHelperFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *OffRampHelperFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*OffRampHelperFeesWithdrawn, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*OffRampHelperOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampHelperOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*OffRampHelperOffRampConfigSet, error)

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
