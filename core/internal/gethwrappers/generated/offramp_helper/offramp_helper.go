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
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005b5638038062005b56833981016040819052620000349162000703565b6040805160808101825260018082526001600160401b0385811660208401526103e89383019390935291831660608201526000805460ff191681558b928b928b928b928b928b928b92908790869082908990889088903390819081620000e15760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b0319909216919091179091558116156200011b576200011b8162000434565b5050506001600160a01b038216158062000133575080155b156200015257604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001985760405162d8548360e71b815260040160405180910390fd5b8151620001ad906005906020850190620004e5565b5060005b825181101562000291576000828281518110620001d257620001d2620007eb565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b0316815250600460008685815181106200021c576200021c620007eb565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff1916600117905580620002888162000801565b915050620001b1565b5050508051825114620002b75760405163ee9d106b60e01b815260040160405180910390fd5b8151620002cc906008906020850190620004e5565b5060005b825181101562000399576000828281518110620002f157620002f1620007eb565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b0316815250600760008685815181106200033b576200033b620007eb565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b02929091169190911790555080620003908162000801565b915050620002d0565b505050151560805260a09790975250505060c0929092525050805160148054602084015160408501516060909501516001600160401b03908116600160c01b026001600160c01b03968216600160801b02969096166001600160801b0392821668010000000000000000026001600160801b031990941691909516179190911716919091179190911790555062000829975050505050505050565b336001600160a01b038216036200048e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000d8565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200053d579160200282015b828111156200053d57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000506565b506200054b9291506200054f565b5090565b5b808211156200054b576000815560010162000550565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620005a757620005a762000566565b604052919050565b60006001600160401b03821115620005cb57620005cb62000566565b5060051b60200190565b6001600160a01b0381168114620005eb57600080fd5b50565b600082601f8301126200060057600080fd5b81516020620006196200061383620005af565b6200057c565b82815260059290921b840181019181810190868411156200063957600080fd5b8286015b84811015620006615780516200065381620005d5565b83529183019183016200063d565b509695505050505050565b600082601f8301126200067e57600080fd5b81516020620006916200061383620005af565b82815260059290921b84018101918181019086841115620006b157600080fd5b8286015b8481101562000661578051620006cb81620005d5565b8352918301918301620006b5565b8051620006e681620005d5565b919050565b80516001600160401b0381168114620006e657600080fd5b60008060008060008060008060006101208a8c0312156200072357600080fd5b895160208b015160408c0151919a5098506001600160401b03808211156200074a57600080fd5b620007588d838e01620005ee565b985060608c01519150808211156200076f57600080fd5b6200077d8d838e01620005ee565b975060808c01519150808211156200079457600080fd5b50620007a38c828d016200066c565b955050620007b460a08b01620006d9565b935060c08a01519250620007cb60e08b01620006eb565b9150620007dc6101008b01620006eb565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b6000600182016200082257634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c0516152f66200086060003960006104d101526000818161042c015261335c015260006120a301526152f66000f3fe608060405234801561001057600080fd5b506004361061025c5760003560e01c806381be8fa411610145578063b034909c116100bd578063bbe4f6db1161008c578063eb511dd411610071578063eb511dd4146106c9578063eefa7a3e146106dc578063f2fde38b1461073257600080fd5b8063bbe4f6db1461067d578063e3d0e712146106b657600080fd5b8063b034909c1461063c578063b1dc65a414610644578063b576716614610657578063b6608c3b1461066a57600080fd5b806389c0656811610114578063a7206cd6116100f9578063a7206cd61461051e578063a8ebd0f41461053e578063afcb95d71461061c57600080fd5b806389c06568146104f35780638da5cb5b146104fb57600080fd5b806381be8fa41461048c57806381ff7048146104945780638456cb59146104c457806385e1f4d0146104cc57600080fd5b80635853c627116101d8578063744b92e2116101a7578063768c577b1161018c578063768c577b1461045c57806379ba50971461046f578063814118341461047757600080fd5b8063744b92e21461041457806374be21501461042757600080fd5b80635853c627146103aa57806359e96b5b146103bd5780635b16ebb7146103d05780635c975abb1461040957600080fd5b80632222dd421161022f5780633b8d08ef116102145780633b8d08ef1461037c5780633f4ba83a1461038f578063461c551b1461039757600080fd5b80632222dd421461034b5780632b898c251461036957600080fd5b8063108ee5fc1461026157806316b8e73114610276578063181f5a77146102d95780632194750714610318575b600080fd5b61027461026f36600461413a565b610745565b005b6102af61028436600461413a565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b604080518082018252600d81527f4f666652616d7020302e302e3100000000000000000000000000000000000000602082015290516102d091906141d1565b61033b6103263660046141e4565b60009081526010602052604090205460ff1690565b60405190151581526020016102d0565b60025473ffffffffffffffffffffffffffffffffffffffff166102af565b6102746103773660046141fd565b610821565b61027461038a3660046146be565b610bf1565b61027461146d565b6102746103a5366004614736565b61147f565b6102746103b83660046141fd565b6114d1565b6102746103cb36600461474e565b6116e9565b61033b6103de36600461413a565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff1661033b565b6102746104223660046141fd565b611767565b61044e7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016102d0565b61044e61046a36600461478f565b611b5c565b610274611ceb565b61047f611e0d565b6040516102d09190614844565b61047f611e7c565b600b546009546040805163ffffffff808516825264010000000090940490931660208401528201526060016102d0565b610274611ee9565b61044e7f000000000000000000000000000000000000000000000000000000000000000081565b61047f611ef9565b600054610100900473ffffffffffffffffffffffffffffffffffffffff166102af565b61044e61052c3660046141e4565b6000908152600f602052604090205490565b6105d8604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260145467ffffffffffffffff808216835268010000000000000000820481166020840152700100000000000000000000000000000000820481169383019390935278010000000000000000000000000000000000000000000000009004909116606082015290565b6040516102d09190815167ffffffffffffffff9081168252602080840151821690830152604080840151821690830152606092830151169181019190915260800190565b6040805160018152600060208201819052918101919091526060016102d0565b60035461044e565b6102746106523660046148a3565b611f66565b610274610665366004614988565b61260f565b6102746106783660046141e4565b61261e565b6102af61068b36600461413a565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6102746106c4366004614a5b565b61269e565b6102746106d73660046141fd565b613083565b6040805160608082018352600080835260208084018290529284015282518082018452601154808252601254828501908152601354928601928352855191825251938101939093525192820192909252016102d0565b61027461074036600461413a565b6132c3565b61074d6132d4565b73ffffffffffffffffffffffffffffffffffffffff811661079a576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b6108296132d4565b6008546000819003610867576040517f2e70248b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610902576040517f3917193900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff161461096b576040517f6c17b98700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600861097a600185614b57565b8154811061098a5761098a614b6e565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600882602001516bffffffffffffffffffffffff16815481106109dc576109dc614b6e565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166008610a0b600186614b57565b81548110610a1b57610a1b614b6e565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600883602001516bffffffffffffffffffffffff1681548110610a8957610a89614b6e565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526007909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556008805480610b2b57610b2b614b9d565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600783526040808520949094558351908152908716918101919091527fa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d91015b60405180910390a15050505050565b60005460ff1615610c63576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610cd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cf69190614bcc565b15610d2c576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b4916004808301926060929190829003018187875af1158015610d9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc19190614be9565b9050600354816020015142610dd69190614b57565b1115610e0e576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610e1a8585611b5c565b6000818152600f6020526040812054919250819003610e695784866040517f07e6809a000000000000000000000000000000000000000000000000000000008152600401610c5a929190614d5b565b6014544290610e8e9068010000000000000000900467ffffffffffffffff1683614dbd565b10610ec5576040517f15c33ba200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855160009081526010602052604090205460ff1615610f165785516040517f6a64e9610000000000000000000000000000000000000000000000000000000081526004810191909152602401610c5a565b60608601516080015173ffffffffffffffffffffffffffffffffffffffff1615801590610f61575060608601516080015173ffffffffffffffffffffffffffffffffffffffff163314155b15610f9e5785516040517fd8e90b980000000000000000000000000000000000000000000000000000000081526004810191909152602401610c5a565b610fa78661335a565b610fb0866134d5565b8551600090815260106020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905583156111ee5760008087606001516000015160008151811061100d5761100d614b6e565b6020026020010151905060006110488273ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff8116611097576040517f83135fec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166350d25bcd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111069190614dd5565b60145461111d919067ffffffffffffffff16614dee565b925082156111ea578289606001516020015160008151811061114157611141614b6e565b602002602001018181516111559190614b57565b90525061116182613584565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081523360048201526024810185905273ffffffffffffffffffffffffffffffffffffffff919091169063ea6192a290604401600060405180830381600087803b1580156111d157600080fd5b505af11580156111e5573d6000803e3d6000fd5b505050505b5050505b60005b606087015151518110156112f257611229876060015160000151828151811061121c5761121c614b6e565b6020026020010151613584565b73ffffffffffffffffffffffffffffffffffffffff1663ea6192a2886060015160600151896060015160200151848151811061126757611267614b6e565b60200260200101516040518363ffffffff1660e01b81526004016112ad92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600060405180830381600087803b1580156112c757600080fd5b505af11580156112db573d6000803e3d6000fd5b5050505080806112ea90614e2b565b9150506111f1565b50606080870151015173ffffffffffffffffffffffffffffffffffffffff163b156113f35785606001516060015173ffffffffffffffffffffffffffffffffffffffff16639c598468876040518263ffffffff1660e01b81526004016113589190614e63565b600060405180830381600087803b15801561137257600080fd5b505af1925050508015611383575060015b6113ee573d8080156113b1576040519150601f19603f3d011682016040523d82523d6000602084013e6113b6565b606091505b5086516040517f6a3fd4f2000000000000000000000000000000000000000000000000000000008152610c5a91908390600401614e76565b61143a565b606086015160a00151511561143a5785516040517fe0244be30000000000000000000000000000000000000000000000000000000081526004810191909152602401610c5a565b85516040517fc51bf0f6d90b467e0849da0ad18a4d9144a4b78b9f83202e1c65cd68f72d4a3890600090a2505050505050565b6114756132d4565b61147d613600565b565b6114876132d4565b8060146114948282614e8f565b9050507fe4cd88b1b5e20a0b843af3207ba74d3f84af4acff45830469490ac9c6ab8a745816040516114c69190614fa0565b60405180910390a150565b6114d96132d4565b73ffffffffffffffffffffffffffffffffffffffff82161580611510575073ffffffffffffffffffffffffffffffffffffffff8116155b15611547576040517fee9d106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156115e3576040517f965ffb7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600880546bffffffffffffffffffffffff908116602080870191825288861660008181526007835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821790558351908152908101919091527f037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb60130591015b60405180910390a1505050565b6116f16132d4565b61171273ffffffffffffffffffffffffffffffffffffffff841683836136e1565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8906060016116dc565b61176f6132d4565b60055460008190036117ad576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290611848576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff16146118b1576040517fd428911900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060056118c0600185614b57565b815481106118d0576118d0614b6e565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff168154811061192257611922614b6e565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166005611951600186614b57565b8154811061196157611961614b6e565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff16815481106119cf576119cf614b6e565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480611a7157611a71614b9d565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101610be2565b600080600060f81b84604051602001611b759190614e63565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052611bb1929160200161500a565b60405160208183030381529060405280519060200120905060005b835151811015611ce357600084600001518281518110611bee57611bee614b6e565b6020026020010151905060028560200151611c099190615081565b600003611c67576040517f010000000000000000000000000000000000000000000000000000000000000060208201526021810184905260418101829052606101604051602081830303815290604052805190602001209250611cba565b6040517f0100000000000000000000000000000000000000000000000000000000000000602082015260218101829052604181018490526061016040516020818303038152906040528051906020012092505b60028560200151611ccb9190615095565b60208601525080611cdb81614e2b565b915050611bcc565b509392505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611d6c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610c5a565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600e805480602002602001604051908101604052809291908181526020018280548015611e7257602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611e47575b5050505050905090565b60606008805480602002602001604051908101604052809291908181526020018280548015611e725760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611e47575050505050905090565b611ef16132d4565b61147d613773565b60606005805480602002602001604051908101604052809291908181526020018280548015611e725760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611e47575050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c013591611fbc91849163ffffffff851691908e908e908190840183828082843760009201919091525061383392505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600954808252600a5460ff80821660208501526101009091041692820192909252908314612091576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610c5a565b61209f8b8b8b8b8b8b613ba4565b60007f0000000000000000000000000000000000000000000000000000000000000000156120fc576002826020015183604001516120dd91906150a9565b6120e791906150ce565b6120f29060016150a9565b60ff169050612112565b602082015161210c9060016150a9565b60ff1690505b88811461217b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610c5a565b8887146121e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610c5a565b336000908152600c602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115612227576122276150f0565b6002811115612238576122386150f0565b9052509050600281602001516002811115612255576122556150f0565b14801561229c5750600e816000015160ff168154811061227757612277614b6e565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b612302576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610c5a565b50505050506000888860405161231992919061511f565b604051908190038120612330918c9060200161512f565b6040516020818303038152906040528051906020012090506123506140f9565b604080518082019091526000808252602082015260005b888110156125ed57600060018588846020811061238657612386614b6e565b61239391901a601b6150a9565b8d8d868181106123a5576123a5614b6e565b905060200201358c8c878181106123be576123be614b6e565b90506020020135604051600081526020016040526040516123fb949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa15801561241d573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600c602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561249d5761249d6150f0565b60028111156124ae576124ae6150f0565b90525092506001836020015160028111156124cb576124cb6150f0565b14612532576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610c5a565b8251849060ff16601f811061254957612549614b6e565b6020020151156125b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610c5a565b600184846000015160ff16601f81106125d0576125d0614b6e565b9115156020909202015250806125e581614e2b565b915050612367565b5050505063ffffffff81106126045761260461514b565b505050505050505050565b61261b60008083613833565b50565b6126266132d4565b80600003612660576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c2519101610815565b855185518560ff16601f831115612711576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610c5a565b6000811161277b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610c5a565b818314612809576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610c5a565b612814816003614dee565b831161287c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610c5a565b6128846132d4565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600d5415612a7757600d546000906128dc90600190614b57565b90506000600d82815481106128f3576128f3614b6e565b6000918252602082200154600e805473ffffffffffffffffffffffffffffffffffffffff9092169350908490811061292d5761292d614b6e565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600c909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600d805491925090806129ad576129ad614b9d565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600e805480612a1657612a16614b9d565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055506128c2915050565b60005b815151811015612ede576000600c600084600001518481518110612aa057612aa0614b6e565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612aea57612aea6150f0565b14612b51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610c5a565b6040805180820190915260ff821681526001602082015282518051600c9160009185908110612b8257612b82614b6e565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612c2357612c236150f0565b021790555060009150612c339050565b600c600084602001518481518110612c4d57612c4d614b6e565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612c9757612c976150f0565b14612cfe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610c5a565b6040805180820190915260ff82168152602081016002815250600c600084602001518481518110612d3157612d31614b6e565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612dd257612dd26150f0565b02179055505082518051600d925083908110612df057612df0614b6e565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600e919083908110612e6c57612e6c614b6e565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691909117905580612ed681614e2b565b915050612a7a565b506040810151600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600b80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092612f7092869290821691161761517a565b92506101000a81548163ffffffff021916908363ffffffff160217905550612fcf4630600b60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613c5b565b600981905582518051600a805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600b5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e059861306e988b98919763ffffffff9092169690959194919391926151a2565b60405180910390a15050505050505050505050565b61308b6132d4565b73ffffffffffffffffffffffffffffffffffffffff821615806130c2575073ffffffffffffffffffffffffffffffffffffffff8116155b156130f9576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015613195576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c91016116dc565b6132cb6132d4565b61261b81613d06565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16331461147d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610c5a565b7f00000000000000000000000000000000000000000000000000000000000000008160200151146133bf5780602001516040517fd44bc9eb000000000000000000000000000000000000000000000000000000008152600401610c5a91815260200190565b60145460608201515151780100000000000000000000000000000000000000000000000090910467ffffffffffffffff1610806134085750606081015160208101515190515114155b1561343f576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601454606082015160a001515170010000000000000000000000000000000090910467ffffffffffffffff16101561261b57601454606082015160a00151516040517f8693378900000000000000000000000000000000000000000000000000000000815270010000000000000000000000000000000090920467ffffffffffffffff1660048301526024820152604401610c5a565b606080820151015173ffffffffffffffffffffffffffffffffffffffff1630148061352c5750606080820151015173ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff165b1561261b5760608082015101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610c5a565b73ffffffffffffffffffffffffffffffffffffffff81811660009081526004602052604090205416806135fb576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610c5a565b919050565b60005460ff1661366c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610c5a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905261376e908490613e01565b505050565b60005460ff16156137e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c5a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586136b73390565b60005460ff16156138a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c5a565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b81526004016020604051808303816000875af115801561390f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139339190614bcc565b15613969576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b4916004808301926060929190829003018187875af11580156139da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139fe9190614be9565b9050600354816020015142613a139190614b57565b1115613a4b576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082806020019051810190613a619190614be9565b9050806040015181602001511115613aa5576040517f67a3824c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516060810182526011548082526012546020830152601354928201929092529015613b2e576040810151613add906001614dbd565b826020015114613b2e57806040015182602001516040517fcc7f1bd0000000000000000000000000000000000000000000000000000000008152600401610c5a929190918252602082015260400190565b81516000908152600f60209081526040918290204290558351601181905581850180516012558386018051601355845192835290519282019290925290518183015290517f07d7bce06be2a7b0230e4dd0d72523c0407e82419ab1d947c5ddaf59ca36484e9181900360600190a1505050505050565b6000613bb1826020614dee565b613bbc856020614dee565b613bc888610144614dbd565b613bd29190614dbd565b613bdc9190614dbd565b613be7906000614dbd565b9050368114613c52576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610c5a565b50505050505050565b6000808a8a8a8a8a8a8a8a8a604051602001613c7f99989796959493929190615238565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603613d85576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610c5a565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000613e63826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16613f0d9092919063ffffffff16565b80519091501561376e5780806020019051810190613e819190614bcc565b61376e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610c5a565b6060613f1c8484600085613f26565b90505b9392505050565b606082471015613fb8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610c5a565b843b614020576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610c5a565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161404991906152cd565b60006040518083038185875af1925050503d8060008114614086576040519150601f19603f3d011682016040523d82523d6000602084013e61408b565b606091505b509150915061409b8282866140a6565b979650505050505050565b606083156140b5575081613f1f565b8251156140c55782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c5a91906141d1565b604051806103e00160405280601f906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff8116811461261b57600080fd5b60006020828403121561414c57600080fd5b8135613f1f81614118565b60005b8381101561417257818101518382015260200161415a565b83811115614181576000848401525b50505050565b6000815180845261419f816020860160208601614157565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613f1f6020830184614187565b6000602082840312156141f657600080fd5b5035919050565b6000806040838503121561421057600080fd5b823561421b81614118565b9150602083013561422b81614118565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff8111828210171561428857614288614236565b60405290565b60405160e0810167ffffffffffffffff8111828210171561428857614288614236565b6040805190810167ffffffffffffffff8111828210171561428857614288614236565b6040516060810167ffffffffffffffff8111828210171561428857614288614236565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561433e5761433e614236565b604052919050565b80356135fb81614118565b600067ffffffffffffffff82111561436b5761436b614236565b5060051b60200190565b600082601f83011261438657600080fd5b8135602061439b61439683614351565b6142f7565b82815260059290921b840181019181810190868411156143ba57600080fd5b8286015b848110156143de5780356143d181614118565b83529183019183016143be565b509695505050505050565b600082601f8301126143fa57600080fd5b8135602061440a61439683614351565b82815260059290921b8401810191818101908684111561442957600080fd5b8286015b848110156143de578035835291830191830161442d565b600082601f83011261445557600080fd5b813567ffffffffffffffff81111561446f5761446f614236565b6144a060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016142f7565b8181528460208386010111156144b557600080fd5b816020850160208301376000918101602001919091529392505050565b6000608082840312156144e457600080fd5b6144ec614265565b90508135815260208201356020820152604082013561450a81614118565b6040820152606082013567ffffffffffffffff8082111561452a57600080fd5b9083019060e0828603121561453e57600080fd5b61454661428e565b82358281111561455557600080fd5b61456187828601614375565b82525060208301358281111561457657600080fd5b614582878286016143e9565b6020830152506040830135604082015261459e60608401614346565b60608201526145af60808401614346565b608082015260a0830135828111156145c657600080fd5b6145d287828601614444565b60a08301525060c0830135828111156145ea57600080fd5b6145f687828601614444565b60c083015250606084015250909392505050565b60006040828403121561461c57600080fd5b6146246142b1565b9050813567ffffffffffffffff81111561463d57600080fd5b8201601f8101841361464e57600080fd5b8035602061465e61439683614351565b82815260059290921b8301810191818101908784111561467d57600080fd5b938201935b8385101561469b57843582529382019390820190614682565b85525093840135938301939093525092915050565b801515811461261b57600080fd5b6000806000606084860312156146d357600080fd5b833567ffffffffffffffff808211156146eb57600080fd5b6146f7878388016144d2565b9450602086013591508082111561470d57600080fd5b5061471a8682870161460a565b925050604084013561472b816146b0565b809150509250925092565b60006080828403121561474857600080fd5b50919050565b60008060006060848603121561476357600080fd5b833561476e81614118565b9250602084013561477e81614118565b929592945050506040919091013590565b600080604083850312156147a257600080fd5b823567ffffffffffffffff808211156147ba57600080fd5b6147c6868387016144d2565b935060208501359150808211156147dc57600080fd5b506147e98582860161460a565b9150509250929050565b600081518084526020808501945080840160005b8381101561483957815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101614807565b509495945050505050565b602081526000613f1f60208301846147f3565b60008083601f84011261486957600080fd5b50813567ffffffffffffffff81111561488157600080fd5b6020830191508360208260051b850101111561489c57600080fd5b9250929050565b60008060008060008060008060e0898b0312156148bf57600080fd5b606089018a8111156148d057600080fd5b8998503567ffffffffffffffff808211156148ea57600080fd5b818b0191508b601f8301126148fe57600080fd5b81358181111561490d57600080fd5b8c602082850101111561491f57600080fd5b6020830199508098505060808b013591508082111561493d57600080fd5b6149498c838d01614857565b909750955060a08b013591508082111561496257600080fd5b5061496f8b828c01614857565b999c989b50969995989497949560c00135949350505050565b60006020828403121561499a57600080fd5b813567ffffffffffffffff8111156149b157600080fd5b6149bd84828501614444565b949350505050565b600082601f8301126149d657600080fd5b813560206149e661439683614351565b82815260059290921b84018101918181019086841115614a0557600080fd5b8286015b848110156143de578035614a1c81614118565b8352918301918301614a09565b803560ff811681146135fb57600080fd5b67ffffffffffffffff8116811461261b57600080fd5b80356135fb81614a3a565b60008060008060008060c08789031215614a7457600080fd5b863567ffffffffffffffff80821115614a8c57600080fd5b614a988a838b016149c5565b97506020890135915080821115614aae57600080fd5b614aba8a838b016149c5565b9650614ac860408a01614a29565b95506060890135915080821115614ade57600080fd5b614aea8a838b01614444565b9450614af860808a01614a50565b935060a0890135915080821115614b0e57600080fd5b50614b1b89828a01614444565b9150509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015614b6957614b69614b28565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060208284031215614bde57600080fd5b8151613f1f816146b0565b600060608284031215614bfb57600080fd5b614c036142d4565b8251815260208301516020820152604083015160408201528091505092915050565b8051825260006020808301518185015273ffffffffffffffffffffffffffffffffffffffff6040840151166040850152606083015160806060860152805160e06080870152614c786101608701826147f3565b838301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80888303810160a08a0152815180845291860193506000929091908601905b80841015614cdb5784518252938601936001939093019290860190614cbb565b50604085015160c08a0152606085015173ffffffffffffffffffffffffffffffffffffffff90811660e08b01526080860151166101008a015260a085015189820383016101208b01529550614d308187614187565b95505060c084015193508088860301610140890152505050614d528282614187565b95945050505050565b60408152600060808201845160408085015281815180845260a08601915060209350838301925060005b81811015614da157835183529284019291840191600101614d85565b50508287015160608601528481038386015261409b8187614c25565b60008219821115614dd057614dd0614b28565b500190565b600060208284031215614de757600080fd5b5051919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614e2657614e26614b28565b500290565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614e5c57614e5c614b28565b5060010190565b602081526000613f1f6020830184614c25565b828152604060208201526000613f1c6040830184614187565b8135614e9a81614a3a565b67ffffffffffffffff811690508154817fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000082161783556020840135614ede81614a3a565b6fffffffffffffffff00000000000000008160401b16905080837fffffffffffffffffffffffffffffffff000000000000000000000000000000008416171784556040850135614f2d81614a3a565b77ffffffffffffffff000000000000000000000000000000008160801b1690507fffffffffffffffff000000000000000000000000000000000000000000000000818582861617841717865560608701359350614f8984614a3a565b808460c01b16858417831717865550505050505050565b608081018235614faf81614a3a565b67ffffffffffffffff9081168352602084013590614fcc82614a3a565b9081166020840152604084013590614fe382614a3a565b9081166040840152606084013590614ffa82614a3a565b8082166060850152505092915050565b7fff000000000000000000000000000000000000000000000000000000000000008316815260008251615044816001850160208701614157565b919091016001019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261509057615090615052565b500690565b6000826150a4576150a4615052565b500490565b600060ff821660ff84168060ff038211156150c6576150c6614b28565b019392505050565b600060ff8316806150e1576150e1615052565b8060ff84160491505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600063ffffffff80831681851680830382111561519957615199614b28565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526151d28184018a6147f3565b905082810360808401526151e681896147f3565b905060ff871660a084015282810360c08401526152038187614187565b905067ffffffffffffffff851660e08401528281036101008401526152288185614187565b9c9b505050505050505050505050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b16604085015281606085015261527f8285018b6147f3565b91508382036080850152615293828a6147f3565b915060ff881660a085015283820360c08501526152b08288614187565b90861660e085015283810361010085015290506152288185614187565b600082516152df818460208701614157565b919091019291505056fea164736f6c634300080d000a",
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
