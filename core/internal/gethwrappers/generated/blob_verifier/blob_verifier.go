// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blob_verifier

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

type CCIPAny2EVMTollMessage struct {
	SourceChainId  *big.Int
	SequenceNumber uint64
	Sender         common.Address
	Receiver       common.Address
	Data           []byte
	Tokens         []common.Address
	Amounts        []*big.Int
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	GasLimit       *big.Int
}

type CCIPExecutionReport struct {
	Messages       []CCIPAny2EVMTollMessage
	Proofs         [][32]byte
	ProofFlagsBits *big.Int
}

type CCIPRelayReport struct {
	MerkleRoot        [32]byte
	MinSequenceNumber uint64
	MaxSequenceNumber uint64
}

type TollOffRampInterfaceOffRampConfig struct {
	ExecutionFeeJuels     uint64
	ExecutionDelaySeconds uint64
	MaxDataSize           uint64
	MaxTokensLength       uint64
}

var BlobVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"contractAggregatorV2V3Interface[]\",\"name\":\"feeds\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionDelayError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"InvalidExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceFeedConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"MerkleProofError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RelayReportError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastMaxSequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newMinSequenceNumber\",\"type\":\"uint64\"}],\"name\":\"SequenceError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenFeedMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnexpectedPayloadData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"CrossChainMessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minSequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceNumber\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structCCIP.RelayReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"ReportAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"addFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagsBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"needFee\",\"type\":\"bool\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastReport\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minSequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceNumber\",\"type\":\"uint64\"}],\"internalType\":\"structCCIP.RelayReport\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRampConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagsBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"removeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setOffRampConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005f2a38038062005f2a833981016040819052620000349162000747565b6000805460ff191681556001908790869082908990889088903390819081620000a45760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000de57620000de81620003ed565b5050506001600160a01b0382161580620000f6575080155b156200011557604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03939093169290921790915560035580518251146200015b5760405162d8548360e71b815260040160405180910390fd5b8151620001709060059060208501906200049e565b5060005b8251811015620002545760008282815181106200019557620001956200081d565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001df57620001df6200081d565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff19166001179055806200024b8162000833565b91505062000174565b50505080518251146200027a5760405163ee9d106b60e01b815260040160405180910390fd5b81516200028f9060089060208501906200049e565b5060005b82518110156200035c576000828281518110620002b457620002b46200081d565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060076000868581518110620002fe57620002fe6200081d565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b02929091169190911790555080620003538162000833565b91505062000293565b505050151560805260a09790975250505060c0929092525050805160138054602084015160408501516060909501516001600160401b03908116600160c01b026001600160c01b03968216600160801b02969096166001600160801b0392821668010000000000000000026001600160801b031990941691909516179190911716919091179190911790556200085b565b336001600160a01b03821603620004475760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016200009b565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620004f6579160200282015b82811115620004f657825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620004bf565b506200050492915062000508565b5090565b5b8082111562000504576000815560010162000509565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156200056057620005606200051f565b604052919050565b60006001600160401b038211156200058457620005846200051f565b5060051b60200190565b6001600160a01b0381168114620005a457600080fd5b50565b600082601f830112620005b957600080fd5b81516020620005d2620005cc8362000568565b62000535565b82815260059290921b84018101918181019086841115620005f257600080fd5b8286015b848110156200061a5780516200060c816200058e565b8352918301918301620005f6565b509695505050505050565b600082601f8301126200063757600080fd5b815160206200064a620005cc8362000568565b82815260059290921b840181019181810190868411156200066a57600080fd5b8286015b848110156200061a57805162000684816200058e565b83529183019183016200066e565b80516200069f816200058e565b919050565b80516001600160401b03811681146200069f57600080fd5b600060808284031215620006cf57600080fd5b604051608081016001600160401b0381118282101715620006f457620006f46200051f565b6040529050806200070583620006a4565b81526200071560208401620006a4565b60208201526200072860408401620006a4565b60408201526200073b60608401620006a4565b60608201525092915050565b600080600080600080600080610160898b0312156200076557600080fd5b885160208a015160408b015191995097506001600160401b03808211156200078c57600080fd5b6200079a8c838d01620005a7565b975060608b0151915080821115620007b157600080fd5b620007bf8c838d01620005a7565b965060808b0151915080821115620007d657600080fd5b50620007e58b828c0162000625565b945050620007f660a08a0162000692565b925060c089015191506200080e8a60e08b01620006bc565b90509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b6000600182016200085457634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05161569862000892600039600061052401526000818161047301526136570152600061232d01526156986000f3fe608060405234801561001057600080fd5b50600436106102925760003560e01c806381be8fa411610160578063b034909c116100d8578063c0d786551161008c578063eb511dd411610071578063eb511dd41461073a578063eefa7a3e1461074d578063f2fde38b146107dd57600080fd5b8063c0d7865514610714578063e3d0e7121461072757600080fd5b8063b1dc65a4116100bd578063b1dc65a4146106b5578063b6608c3b146106c8578063bbe4f6db146106db57600080fd5b8063b034909c1461068f578063b0f479a11461069757600080fd5b806389c065681161012f578063a7206cd611610114578063a7206cd614610571578063a8ebd0f414610591578063afcb95d71461066f57600080fd5b806389c06568146105465780638da5cb5b1461054e57600080fd5b806381be8fa4146104df57806381ff7048146104e75780638456cb591461051757806385e1f4d01461051f57600080fd5b8063567c814b1161020e578063744b92e2116101c257806379ba5097116101a757806379ba50971461049557806380d9a1b71461049d57806381411834146104ca57600080fd5b8063744b92e21461045b57806374be21501461046e57600080fd5b806359e96b5b116101f357806359e96b5b146104045780635b16ebb7146104175780635c975abb1461045057600080fd5b8063567c814b146103ce5780635853c627146103f157600080fd5b8063295938ec116102655780633dd80c701161024a5780633dd80c70146103925780633f4ba83a146103b3578063461c551b146103bb57600080fd5b8063295938ec1461036c5780632b898c251461037f57600080fd5b8063108ee5fc1461029757806316b8e731146102ac578063181f5a771461030f5780632222dd421461034e575b600080fd5b6102aa6102a53660046145da565b6107f0565b005b6102e56102ba3660046145da565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b604080518082018252600d81527f4f666652616d7020302e302e310000000000000000000000000000000000000060208201529051610306919061467c565b60025473ffffffffffffffffffffffffffffffffffffffff166102e5565b6102aa61037a366004614afb565b6108cc565b6102aa61038d366004614b4d565b611195565b6103a56103a0366004614b7b565b611565565b604051908152602001610306565b6102aa61173f565b6102aa6103c9366004614bb0565b611751565b6103e16103dc366004614bc8565b6117a3565b6040519015158152602001610306565b6102aa6103ff366004614b4d565b6118ea565b6102aa610412366004614be1565b611b02565b6103e16104253660046145da565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff166103e1565b6102aa610469366004614b4d565b611b80565b6103a57f000000000000000000000000000000000000000000000000000000000000000081565b6102aa611f75565b6103e16104ab366004614c22565b67ffffffffffffffff1660009081526010602052604090205460ff1690565b6104d2612097565b6040516103069190614c90565b6104d2612106565b600b546009546040805163ffffffff80851682526401000000009094049093166020840152820152606001610306565b6102aa612173565b6103a57f000000000000000000000000000000000000000000000000000000000000000081565b6104d2612183565b600054610100900473ffffffffffffffffffffffffffffffffffffffff166102e5565b6103a561057f366004614bc8565b6000908152600f602052604090205490565b61062b604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260135467ffffffffffffffff808216835268010000000000000000820481166020840152700100000000000000000000000000000000820481169383019390935278010000000000000000000000000000000000000000000000009004909116606082015290565b6040516103069190815167ffffffffffffffff9081168252602080840151821690830152604080840151821690830152606092830151169181019190915260800190565b604080516001815260006020820181905291810191909152606001610306565b6003546103a5565b60145473ffffffffffffffffffffffffffffffffffffffff166102e5565b6102aa6106c3366004614cef565b6121f0565b6102aa6106d6366004614bc8565b612899565b6102e56106e93660046145da565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6102aa6107223660046145da565b612919565b6102aa610735366004614de5565b612994565b6102aa610748366004614b4d565b613379565b6107a960408051606081018252600080825260208201819052918101919091525060408051606081018252601154815260125467ffffffffffffffff808216602084015268010000000000000000909104169181019190915290565b604080518251815260208084015167ffffffffffffffff908116918301919091529282015190921690820152606001610306565b6102aa6107eb3660046145da565b6135b9565b6107f86135cd565b73ffffffffffffffffffffffffffffffffffffffff8116610845576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b60005460ff161561093e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109cf9190614eb2565b15610a05576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa158015610a75573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a999190614ecf565b9050600354816020015142610aae9190614f3a565b1115610ae6576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60145473ffffffffffffffffffffffffffffffffffffffff16610b35576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610b4084611565565b6000818152600f6020526040812054919250819003610b8e576040517f851bdf5300000000000000000000000000000000000000000000000000000000815260048101839052602401610935565b6013544290610bb39068010000000000000000900467ffffffffffffffff1683614f51565b10610bea576040517f15c33ba200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b85515181101561118d57600086600001518281518110610c0f57610c0f614f69565b6020908102919091018101518082015167ffffffffffffffff166000908152601090925260409091205490915060ff1615610c885760208101516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610935565b610c9181613653565b610c9a816137be565b60208082015167ffffffffffffffff16600090815260109091526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558515610ede576000808260a00151600081518110610d0157610d01614f69565b602002602001015190506000610d3c8273ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff8116610d8b576040517f83135fec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166350d25bcd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dd6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dfa9190614f98565b601354610e11919067ffffffffffffffff16614fb1565b92508215610eda57828460c00151600081518110610e3157610e31614f69565b60200260200101818151610e459190614f3a565b905250610e5182613864565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081523360048201526024810185905273ffffffffffffffffffffffffffffffffffffffff919091169063ea6192a290604401600060405180830381600087803b158015610ec157600080fd5b505af1158015610ed5573d6000803e3d6000fd5b505050505b5050505b60005b8160a0015151811015610fd557610f148260a001518281518110610f0757610f07614f69565b6020026020010151613864565b73ffffffffffffffffffffffffffffffffffffffff1663ea6192a283606001518460c001518481518110610f4a57610f4a614f69565b60200260200101516040518363ffffffff1660e01b8152600401610f9092919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600060405180830381600087803b158015610faa57600080fd5b505af1158015610fbe573d6000803e3d6000fd5b505050508080610fcd90614fee565b915050610ee1565b50606081015173ffffffffffffffffffffffffffffffffffffffff163b156110f05760145460608201516040517ffd12f6e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9092169163fd12f6e59161105291859060040161515f565b600060405180830381600087803b15801561106c57600080fd5b505af192505050801561107d575060015b6110eb573d8080156110ab576040519150601f19603f3d011682016040523d82523d6000602084013e6110b0565b606091505b508160200151816040517fa1dc818500000000000000000000000000000000000000000000000000000000815260040161093592919061518e565b61113f565b6080810151511561113f5760208101516040517fc945cae000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610935565b806020015167ffffffffffffffff167f88a2c71ae86800edd209d83430f348ce70539bd40085d045340ff569f058a2a060405160405180910390a2508061118581614fee565b915050610bed565b505050505050565b61119d6135cd565b60085460008190036111db576040517f2e70248b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290611276576040517f3917193900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff16146112df576040517f9403a50500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060086112ee600185614f3a565b815481106112fe576112fe614f69565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600882602001516bffffffffffffffffffffffff168154811061135057611350614f69565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16600861137f600186614f3a565b8154811061138f5761138f614f69565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600883602001516bffffffffffffffffffffffff16815481106113fd576113fd614f69565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526007909152604090912080546bffffffffffffffffffffffff909216740100000000000000000000000000000000000000000291909216179055600880548061149f5761149f6151b1565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600783526040808520949094558351908152908716918101919091527fa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d91015b60405180910390a15050505050565b8051516020820151516000919082906001906115819084614f51565b61158b9190614f3a565b905061010081111561159c57600080fd5b60008167ffffffffffffffff8111156115b7576115b761468f565b6040519080825280602002602001820160405280156115e0578160200160208202803683370190505b5090506000806000805b858110156116e8576040890151811c6001908116146116c28161162f5760208b01518051600186019590811061162257611622614f69565b6020026020010151611676565b88861061164d57865160018601958891811061162257611622614f69565b8a51805160018801976116769291811061166957611669614f69565b60200260200101516138e0565b8987106116a157875160018701968991811061169457611694614f69565b6020026020010151613950565b8b51805160018901986116bd9291811061166957611669614f69565b613950565b8683815181106116d4576116d4614f69565b6020908102919091010152506001016115ea565b5084156117195783600186038151811061170457611704614f69565b60200260200101519650505050505050919050565b611733886000015160008151811061166957611669614f69565b98975050505050505050565b6117476135cd565b61174f613a0e565b565b6117596135cd565b80601361176682826151e0565b9050507fe4cd88b1b5e20a0b843af3207ba74d3f84af4acff45830469490ac9c6ab8a7458160405161179891906152f1565b60405180910390a150565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa158015611813573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118379190614eb2565b1580156118e45750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa1580156118af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118d39190614ecf565b602001516118e19084614f3a565b11155b92915050565b6118f26135cd565b73ffffffffffffffffffffffffffffffffffffffff82161580611929575073ffffffffffffffffffffffffffffffffffffffff8116155b15611960576040517fee9d106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156119fc576040517f965ffb7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600880546bffffffffffffffffffffffff908116602080870191825288861660008181526007835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821790558351908152908101919091527f037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb60130591015b60405180910390a1505050565b611b0a6135cd565b611b2b73ffffffffffffffffffffffffffffffffffffffff84168383613aef565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa890606001611af5565b611b886135cd565b6005546000819003611bc6576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290611c61576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614611cca576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006005611cd9600185614f3a565b81548110611ce957611ce9614f69565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff1681548110611d3b57611d3b614f69565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166005611d6a600186614f3a565b81548110611d7a57611d7a614f69565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110611de857611de8614f69565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480611e8a57611e8a6151b1565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101611556565b60015473ffffffffffffffffffffffffffffffffffffffff163314611ff6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610935565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600e8054806020026020016040519081016040528092919081815260200182805480156120fc57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116120d1575b5050505050905090565b606060088054806020026020016040519081016040528092919081815260200182805480156120fc5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116120d1575050505050905090565b61217b6135cd565b61174f613b81565b606060058054806020026020016040519081016040528092919081815260200182805480156120fc5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116120d1575050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161224691849163ffffffff851691908e908e9081908401838280828437600092019190915250613c4192505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600954808252600a5460ff8082166020850152610100909104169282019290925290831461231b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610935565b6123298b8b8b8b8b8b614046565b60007f00000000000000000000000000000000000000000000000000000000000000001561238657600282602001518360400151612367919061535b565b6123719190615380565b61237c90600161535b565b60ff16905061239c565b602082015161239690600161535b565b60ff1690505b888114612405576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610935565b88871461246e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610935565b336000908152600c602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156124b1576124b16153c9565b60028111156124c2576124c26153c9565b90525090506002816020015160028111156124df576124df6153c9565b1480156125265750600e816000015160ff168154811061250157612501614f69565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b61258c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610935565b5050505050600088886040516125a39291906153f8565b6040519081900381206125ba918c90602001615408565b6040516020818303038152906040528051906020012090506125da614599565b604080518082019091526000808252602082015260005b8881101561287757600060018588846020811061261057612610614f69565b61261d91901a601b61535b565b8d8d8681811061262f5761262f614f69565b905060200201358c8c8781811061264857612648614f69565b9050602002013560405160008152602001604052604051612685949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156126a7573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600c602090815290849020838501909452835460ff80821685529296509294508401916101009004166002811115612727576127276153c9565b6002811115612738576127386153c9565b9052509250600183602001516002811115612755576127556153c9565b146127bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610935565b8251849060ff16601f81106127d3576127d3614f69565b60200201511561283f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610935565b600184846000015160ff16601f811061285a5761285a614f69565b91151560209092020152508061286f81614fee565b9150506125f1565b5050505063ffffffff811061288e5761288e615424565b505050505050505050565b6128a16135cd565b806000036128db576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c25191016108c0565b6129216135cd565b601480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490602001611798565b855185518560ff16601f831115612a07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610935565b60008111612a71576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610935565b818314612aff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610935565b612b0a816003614fb1565b8311612b72576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610935565b612b7a6135cd565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600d5415612d6d57600d54600090612bd290600190614f3a565b90506000600d8281548110612be957612be9614f69565b6000918252602082200154600e805473ffffffffffffffffffffffffffffffffffffffff90921693509084908110612c2357612c23614f69565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600c909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600d80549192509080612ca357612ca36151b1565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600e805480612d0c57612d0c6151b1565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550612bb8915050565b60005b8151518110156131d4576000600c600084600001518481518110612d9657612d96614f69565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612de057612de06153c9565b14612e47576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610935565b6040805180820190915260ff821681526001602082015282518051600c9160009185908110612e7857612e78614f69565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612f1957612f196153c9565b021790555060009150612f299050565b600c600084602001518481518110612f4357612f43614f69565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612f8d57612f8d6153c9565b14612ff4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610935565b6040805180820190915260ff82168152602081016002815250600c60008460200151848151811061302757613027614f69565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156130c8576130c86153c9565b02179055505082518051600d9250839081106130e6576130e6614f69565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600e91908390811061316257613162614f69565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055806131cc81614fee565b915050612d70565b506040810151600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600b80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092613266928692908216911617615453565b92506101000a81548163ffffffff021916908363ffffffff1602179055506132c54630600b60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a001516140fd565b600981905582518051600a805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600b5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598613364988b98919763ffffffff90921696909591949193919261547b565b60405180910390a15050505050505050505050565b6133816135cd565b73ffffffffffffffffffffffffffffffffffffffff821615806133b8575073ffffffffffffffffffffffffffffffffffffffff8116155b156133ef576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152901561348b576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c9101611af5565b6135c16135cd565b6135ca816141a8565b50565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16331461174f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610935565b80517f0000000000000000000000000000000000000000000000000000000000000000146136b35780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610935565b60135460a082015151780100000000000000000000000000000000000000000000000090910467ffffffffffffffff1610806136f957508060c00151518160a001515114155b15613730576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60135460808201515170010000000000000000000000000000000090910467ffffffffffffffff1610156135ca576013546080820151516040517f8693378900000000000000000000000000000000000000000000000000000000815270010000000000000000000000000000000090920467ffffffffffffffff1660048301526024820152604401610935565b606081015173ffffffffffffffffffffffffffffffffffffffff1630148061380f5750606081015173ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff165b156135ca5760608101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610935565b73ffffffffffffffffffffffffffffffffffffffff81811660009081526004602052604090205416806138db576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610935565b919050565b60008060f81b826040516020016138f79190615511565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526139339291602001615524565b604051602081830303815290604052805190602001209050919050565b60008183106139b257604080517f01000000000000000000000000000000000000000000000000000000000000006020808301919091526021820185905260418083018790528351808403909101815260619092019092528051910120613a07565b604080517f010000000000000000000000000000000000000000000000000000000000000060208083019190915260218201869052604180830186905283518084039091018152606190920190925280519101205b9392505050565b60005460ff16613a7a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610935565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052613b7c9084906142a3565b505050565b60005460ff1615613bee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610935565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613ac53390565b60005460ff1615613cae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610935565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613d1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d3f9190614eb2565b15613d75576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa158015613de5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e099190614ecf565b9050600354816020015142613e1e9190614f3a565b1115613e56576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082806020019051810190613e6c919061556c565b9050806040015167ffffffffffffffff16816020015167ffffffffffffffff161115613ec4576040517f67a3824c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160608101825260115480825260125467ffffffffffffffff80821660208501526801000000000000000090910416928201929092529015613f81576040810151613f139060016155b7565b67ffffffffffffffff16826020015167ffffffffffffffff1614613f8157604080820151602084015191517f8e8c0add00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff918216600482015291166024820152604401610935565b81516000908152600f602090815260409182902042905583516011819055818501805160128054868901805167ffffffffffffffff90811668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090931694811694909417919091179091558551938452915181169383019390935251909116918101919091527f6a2424b54e9b18f0c3d440b6003882f5ffa34272205608ad9451c921c0d99a359060600160405180910390a1505050505050565b6000614053826020614fb1565b61405e856020614fb1565b61406a88610144614f51565b6140749190614f51565b61407e9190614f51565b614089906000614f51565b90503681146140f4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610935565b50505050505050565b6000808a8a8a8a8a8a8a8a8a604051602001614121999897969594939291906155da565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603614227576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610935565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000614305826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166143af9092919063ffffffff16565b805190915015613b7c57808060200190518101906143239190614eb2565b613b7c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610935565b60606143be84846000856143c6565b949350505050565b606082471015614458576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610935565b843b6144c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610935565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516144e9919061566f565b60006040518083038185875af1925050503d8060008114614526576040519150601f19603f3d011682016040523d82523d6000602084013e61452b565b606091505b509150915061453b828286614546565b979650505050505050565b60608315614555575081613a07565b8251156145655782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610935919061467c565b604051806103e00160405280601f906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff811681146135ca57600080fd5b6000602082840312156145ec57600080fd5b8135613a07816145b8565b80356138db816145b8565b60005b8381101561461d578181015183820152602001614605565b8381111561462c576000848401525b50505050565b6000815180845261464a816020860160208601614602565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613a076020830184614632565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156146e1576146e161468f565b60405290565b604051610140810167ffffffffffffffff811182821017156146e1576146e161468f565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156147525761475261468f565b604052919050565b600067ffffffffffffffff8211156147745761477461468f565b5060051b60200190565b67ffffffffffffffff811681146135ca57600080fd5b80356138db8161477e565b600082601f8301126147b057600080fd5b813567ffffffffffffffff8111156147ca576147ca61468f565b6147fb60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161470b565b81815284602083860101111561481057600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261483e57600080fd5b8135602061485361484e8361475a565b61470b565b82815260059290921b8401810191818101908684111561487257600080fd5b8286015b84811015614896578035614889816145b8565b8352918301918301614876565b509695505050505050565b600082601f8301126148b257600080fd5b813560206148c261484e8361475a565b82815260059290921b840181019181810190868411156148e157600080fd5b8286015b8481101561489657803583529183019183016148e5565b60006060828403121561490e57600080fd5b6149166146be565b9050813567ffffffffffffffff8082111561493057600080fd5b818401915084601f83011261494457600080fd5b8135602061495461484e8361475a565b82815260059290921b8401810191818101908884111561497357600080fd5b8286015b84811015614ab25780358681111561498e57600080fd5b8701610140818c037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018113156149c457600080fd5b6149cc6146e7565b8683013581526149de60408401614794565b878201526149ee606084016145f7565b60408201526149ff608084016145f7565b606082015260a083013589811115614a1657600080fd5b614a248e898387010161479f565b60808301525060c083013589811115614a3d5760008081fd5b614a4b8e898387010161482d565b60a08301525060e0808401358a811115614a655760008081fd5b614a738f8a838801016148a1565b60c084015250610100614a878186016145f7565b9183019190915261012084810135918301919091529190920135908201528352918301918301614977565b5086525085810135935082841115614ac957600080fd5b614ad5878588016148a1565b81860152505050506040820135604082015292915050565b80151581146135ca57600080fd5b60008060408385031215614b0e57600080fd5b823567ffffffffffffffff811115614b2557600080fd5b614b31858286016148fc565b9250506020830135614b4281614aed565b809150509250929050565b60008060408385031215614b6057600080fd5b8235614b6b816145b8565b91506020830135614b42816145b8565b600060208284031215614b8d57600080fd5b813567ffffffffffffffff811115614ba457600080fd5b6143be848285016148fc565b600060808284031215614bc257600080fd5b50919050565b600060208284031215614bda57600080fd5b5035919050565b600080600060608486031215614bf657600080fd5b8335614c01816145b8565b92506020840135614c11816145b8565b929592945050506040919091013590565b600060208284031215614c3457600080fd5b8135613a078161477e565b600081518084526020808501945080840160005b83811015614c8557815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101614c53565b509495945050505050565b602081526000613a076020830184614c3f565b60008083601f840112614cb557600080fd5b50813567ffffffffffffffff811115614ccd57600080fd5b6020830191508360208260051b8501011115614ce857600080fd5b9250929050565b60008060008060008060008060e0898b031215614d0b57600080fd5b606089018a811115614d1c57600080fd5b8998503567ffffffffffffffff80821115614d3657600080fd5b818b0191508b601f830112614d4a57600080fd5b813581811115614d5957600080fd5b8c6020828501011115614d6b57600080fd5b6020830199508098505060808b0135915080821115614d8957600080fd5b614d958c838d01614ca3565b909750955060a08b0135915080821115614dae57600080fd5b50614dbb8b828c01614ca3565b999c989b50969995989497949560c00135949350505050565b803560ff811681146138db57600080fd5b60008060008060008060c08789031215614dfe57600080fd5b863567ffffffffffffffff80821115614e1657600080fd5b614e228a838b0161482d565b97506020890135915080821115614e3857600080fd5b614e448a838b0161482d565b9650614e5260408a01614dd4565b95506060890135915080821115614e6857600080fd5b614e748a838b0161479f565b9450614e8260808a01614794565b935060a0890135915080821115614e9857600080fd5b50614ea589828a0161479f565b9150509295509295509295565b600060208284031215614ec457600080fd5b8151613a0781614aed565b600060608284031215614ee157600080fd5b614ee96146be565b8251815260208301516020820152604083015160408201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015614f4c57614f4c614f0b565b500390565b60008219821115614f6457614f64614f0b565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215614faa57600080fd5b5051919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614fe957614fe9614f0b565b500290565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361501f5761501f614f0b565b5060010190565b600081518084526020808501945080840160005b83811015614c855781518752958201959082019060010161503a565b600061014082518452602083015161507a602086018267ffffffffffffffff169052565b5060408301516150a2604086018273ffffffffffffffffffffffffffffffffffffffff169052565b5060608301516150ca606086018273ffffffffffffffffffffffffffffffffffffffff169052565b5060808301518160808601526150e282860182614632565b91505060a083015184820360a08601526150fc8282614c3f565b91505060c083015184820360c08601526151168282615026565b91505060e083015161514060e086018273ffffffffffffffffffffffffffffffffffffffff169052565b5061010083810151908501526101209283015192909301919091525090565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006143be6040830184615056565b67ffffffffffffffff831681526040602082015260006143be6040830184614632565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b81356151eb8161477e565b67ffffffffffffffff811690508154817fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008216178355602084013561522f8161477e565b6fffffffffffffffff00000000000000008160401b16905080837fffffffffffffffffffffffffffffffff00000000000000000000000000000000841617178455604085013561527e8161477e565b77ffffffffffffffff000000000000000000000000000000008160801b1690507fffffffffffffffff0000000000000000000000000000000000000000000000008185828616178417178655606087013593506152da8461477e565b808460c01b16858417831717865550505050505050565b6080810182356153008161477e565b67ffffffffffffffff908116835260208401359061531d8261477e565b90811660208401526040840135906153348261477e565b908116604084015260608401359061534b8261477e565b8082166060850152505092915050565b600060ff821660ff84168060ff0382111561537857615378614f0b565b019392505050565b600060ff8316806153ba577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8060ff84160491505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600063ffffffff80831681851680830382111561547257615472614f0b565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526154ab8184018a614c3f565b905082810360808401526154bf8189614c3f565b905060ff871660a084015282810360c08401526154dc8187614632565b905067ffffffffffffffff851660e08401528281036101008401526155018185614632565b9c9b505050505050505050505050565b602081526000613a076020830184615056565b7fff00000000000000000000000000000000000000000000000000000000000000831681526000825161555e816001850160208701614602565b919091016001019392505050565b60006060828403121561557e57600080fd5b6155866146be565b8251815260208301516155988161477e565b602082015260408301516155ab8161477e565b60408201529392505050565b600067ffffffffffffffff80831681851680830382111561547257615472614f0b565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526156218285018b614c3f565b91508382036080850152615635828a614c3f565b915060ff881660a085015283820360c08501526156528288614632565b90861660e085015283810361010085015290506155018185614632565b60008251615681818460208701614602565b919091019291505056fea164736f6c634300080d000a",
}

var BlobVerifierABI = BlobVerifierMetaData.ABI

var BlobVerifierBin = BlobVerifierMetaData.Bin

func DeployBlobVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, sourceTokens []common.Address, pools []common.Address, feeds []common.Address, afn common.Address, maxTimeWithoutAFNSignal *big.Int, config TollOffRampInterfaceOffRampConfig) (common.Address, *types.Transaction, *BlobVerifier, error) {
	parsed, err := BlobVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlobVerifierBin), backend, sourceChainId, chainId, sourceTokens, pools, feeds, afn, maxTimeWithoutAFNSignal, config)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlobVerifier{BlobVerifierCaller: BlobVerifierCaller{contract: contract}, BlobVerifierTransactor: BlobVerifierTransactor{contract: contract}, BlobVerifierFilterer: BlobVerifierFilterer{contract: contract}}, nil
}

type BlobVerifier struct {
	address common.Address
	abi     abi.ABI
	BlobVerifierCaller
	BlobVerifierTransactor
	BlobVerifierFilterer
}

type BlobVerifierCaller struct {
	contract *bind.BoundContract
}

type BlobVerifierTransactor struct {
	contract *bind.BoundContract
}

type BlobVerifierFilterer struct {
	contract *bind.BoundContract
}

type BlobVerifierSession struct {
	Contract     *BlobVerifier
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BlobVerifierCallerSession struct {
	Contract *BlobVerifierCaller
	CallOpts bind.CallOpts
}

type BlobVerifierTransactorSession struct {
	Contract     *BlobVerifierTransactor
	TransactOpts bind.TransactOpts
}

type BlobVerifierRaw struct {
	Contract *BlobVerifier
}

type BlobVerifierCallerRaw struct {
	Contract *BlobVerifierCaller
}

type BlobVerifierTransactorRaw struct {
	Contract *BlobVerifierTransactor
}

func NewBlobVerifier(address common.Address, backend bind.ContractBackend) (*BlobVerifier, error) {
	abi, err := abi.JSON(strings.NewReader(BlobVerifierABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBlobVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlobVerifier{address: address, abi: abi, BlobVerifierCaller: BlobVerifierCaller{contract: contract}, BlobVerifierTransactor: BlobVerifierTransactor{contract: contract}, BlobVerifierFilterer: BlobVerifierFilterer{contract: contract}}, nil
}

func NewBlobVerifierCaller(address common.Address, caller bind.ContractCaller) (*BlobVerifierCaller, error) {
	contract, err := bindBlobVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlobVerifierCaller{contract: contract}, nil
}

func NewBlobVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*BlobVerifierTransactor, error) {
	contract, err := bindBlobVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlobVerifierTransactor{contract: contract}, nil
}

func NewBlobVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*BlobVerifierFilterer, error) {
	contract, err := bindBlobVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlobVerifierFilterer{contract: contract}, nil
}

func bindBlobVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlobVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_BlobVerifier *BlobVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlobVerifier.Contract.BlobVerifierCaller.contract.Call(opts, result, method, params...)
}

func (_BlobVerifier *BlobVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlobVerifier.Contract.BlobVerifierTransactor.contract.Transfer(opts)
}

func (_BlobVerifier *BlobVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlobVerifier.Contract.BlobVerifierTransactor.contract.Transact(opts, method, params...)
}

func (_BlobVerifier *BlobVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlobVerifier.Contract.contract.Call(opts, result, method, params...)
}

func (_BlobVerifier *BlobVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlobVerifier.Contract.contract.Transfer(opts)
}

func (_BlobVerifier *BlobVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlobVerifier.Contract.contract.Transact(opts, method, params...)
}

func (_BlobVerifier *BlobVerifierCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) CHAINID() (*big.Int, error) {
	return _BlobVerifier.Contract.CHAINID(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) CHAINID() (*big.Int, error) {
	return _BlobVerifier.Contract.CHAINID(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCaller) SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "SOURCE_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) SOURCECHAINID() (*big.Int, error) {
	return _BlobVerifier.Contract.SOURCECHAINID(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) SOURCECHAINID() (*big.Int, error) {
	return _BlobVerifier.Contract.SOURCECHAINID(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) GetAFN() (common.Address, error) {
	return _BlobVerifier.Contract.GetAFN(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) GetAFN() (common.Address, error) {
	return _BlobVerifier.Contract.GetAFN(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCaller) GetExecuted(opts *bind.CallOpts, sequenceNumber uint64) (bool, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "getExecuted", sequenceNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) GetExecuted(sequenceNumber uint64) (bool, error) {
	return _BlobVerifier.Contract.GetExecuted(&_BlobVerifier.CallOpts, sequenceNumber)
}

func (_BlobVerifier *BlobVerifierCallerSession) GetExecuted(sequenceNumber uint64) (bool, error) {
	return _BlobVerifier.Contract.GetExecuted(&_BlobVerifier.CallOpts, sequenceNumber)
}

func (_BlobVerifier *BlobVerifierCaller) GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "getFeed", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) GetFeed(token common.Address) (common.Address, error) {
	return _BlobVerifier.Contract.GetFeed(&_BlobVerifier.CallOpts, token)
}

func (_BlobVerifier *BlobVerifierCallerSession) GetFeed(token common.Address) (common.Address, error) {
	return _BlobVerifier.Contract.GetFeed(&_BlobVerifier.CallOpts, token)
}

func (_BlobVerifier *BlobVerifierCaller) GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "getFeedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) GetFeedTokens() ([]common.Address, error) {
	return _BlobVerifier.Contract.GetFeedTokens(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) GetFeedTokens() ([]common.Address, error) {
	return _BlobVerifier.Contract.GetFeedTokens(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCaller) GetLastReport(opts *bind.CallOpts) (CCIPRelayReport, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "getLastReport")

	if err != nil {
		return *new(CCIPRelayReport), err
	}

	out0 := *abi.ConvertType(out[0], new(CCIPRelayReport)).(*CCIPRelayReport)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) GetLastReport() (CCIPRelayReport, error) {
	return _BlobVerifier.Contract.GetLastReport(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) GetLastReport() (CCIPRelayReport, error) {
	return _BlobVerifier.Contract.GetLastReport(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCaller) GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "getMaxSecondsWithoutAFNHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _BlobVerifier.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _BlobVerifier.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCaller) GetMerkleRoot(opts *bind.CallOpts, root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "getMerkleRoot", root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) GetMerkleRoot(root [32]byte) (*big.Int, error) {
	return _BlobVerifier.Contract.GetMerkleRoot(&_BlobVerifier.CallOpts, root)
}

func (_BlobVerifier *BlobVerifierCallerSession) GetMerkleRoot(root [32]byte) (*big.Int, error) {
	return _BlobVerifier.Contract.GetMerkleRoot(&_BlobVerifier.CallOpts, root)
}

func (_BlobVerifier *BlobVerifierCaller) GetOffRampConfig(opts *bind.CallOpts) (TollOffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "getOffRampConfig")

	if err != nil {
		return *new(TollOffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(TollOffRampInterfaceOffRampConfig)).(*TollOffRampInterfaceOffRampConfig)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) GetOffRampConfig() (TollOffRampInterfaceOffRampConfig, error) {
	return _BlobVerifier.Contract.GetOffRampConfig(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) GetOffRampConfig() (TollOffRampInterfaceOffRampConfig, error) {
	return _BlobVerifier.Contract.GetOffRampConfig(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _BlobVerifier.Contract.GetPool(&_BlobVerifier.CallOpts, sourceToken)
}

func (_BlobVerifier *BlobVerifierCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _BlobVerifier.Contract.GetPool(&_BlobVerifier.CallOpts, sourceToken)
}

func (_BlobVerifier *BlobVerifierCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) GetPoolTokens() ([]common.Address, error) {
	return _BlobVerifier.Contract.GetPoolTokens(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _BlobVerifier.Contract.GetPoolTokens(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) GetRouter() (common.Address, error) {
	return _BlobVerifier.Contract.GetRouter(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) GetRouter() (common.Address, error) {
	return _BlobVerifier.Contract.GetRouter(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCaller) IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "isHealthy", timeNow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _BlobVerifier.Contract.IsHealthy(&_BlobVerifier.CallOpts, timeNow)
}

func (_BlobVerifier *BlobVerifierCallerSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _BlobVerifier.Contract.IsHealthy(&_BlobVerifier.CallOpts, timeNow)
}

func (_BlobVerifier *BlobVerifierCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) IsPool(addr common.Address) (bool, error) {
	return _BlobVerifier.Contract.IsPool(&_BlobVerifier.CallOpts, addr)
}

func (_BlobVerifier *BlobVerifierCallerSession) IsPool(addr common.Address) (bool, error) {
	return _BlobVerifier.Contract.IsPool(&_BlobVerifier.CallOpts, addr)
}

func (_BlobVerifier *BlobVerifierCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_BlobVerifier *BlobVerifierSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _BlobVerifier.Contract.LatestConfigDetails(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _BlobVerifier.Contract.LatestConfigDetails(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_BlobVerifier *BlobVerifierSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _BlobVerifier.Contract.LatestConfigDigestAndEpoch(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _BlobVerifier.Contract.LatestConfigDigestAndEpoch(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCaller) MerkleRoot(opts *bind.CallOpts, report CCIPExecutionReport) ([32]byte, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "merkleRoot", report)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) MerkleRoot(report CCIPExecutionReport) ([32]byte, error) {
	return _BlobVerifier.Contract.MerkleRoot(&_BlobVerifier.CallOpts, report)
}

func (_BlobVerifier *BlobVerifierCallerSession) MerkleRoot(report CCIPExecutionReport) ([32]byte, error) {
	return _BlobVerifier.Contract.MerkleRoot(&_BlobVerifier.CallOpts, report)
}

func (_BlobVerifier *BlobVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) Owner() (common.Address, error) {
	return _BlobVerifier.Contract.Owner(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) Owner() (common.Address, error) {
	return _BlobVerifier.Contract.Owner(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) Paused() (bool, error) {
	return _BlobVerifier.Contract.Paused(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) Paused() (bool, error) {
	return _BlobVerifier.Contract.Paused(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) Transmitters() ([]common.Address, error) {
	return _BlobVerifier.Contract.Transmitters(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) Transmitters() ([]common.Address, error) {
	return _BlobVerifier.Contract.Transmitters(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BlobVerifier.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BlobVerifier *BlobVerifierSession) TypeAndVersion() (string, error) {
	return _BlobVerifier.Contract.TypeAndVersion(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierCallerSession) TypeAndVersion() (string, error) {
	return _BlobVerifier.Contract.TypeAndVersion(&_BlobVerifier.CallOpts)
}

func (_BlobVerifier *BlobVerifierTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "acceptOwnership")
}

func (_BlobVerifier *BlobVerifierSession) AcceptOwnership() (*types.Transaction, error) {
	return _BlobVerifier.Contract.AcceptOwnership(&_BlobVerifier.TransactOpts)
}

func (_BlobVerifier *BlobVerifierTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BlobVerifier.Contract.AcceptOwnership(&_BlobVerifier.TransactOpts)
}

func (_BlobVerifier *BlobVerifierTransactor) AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "addFeed", token, feed)
}

func (_BlobVerifier *BlobVerifierSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _BlobVerifier.Contract.AddFeed(&_BlobVerifier.TransactOpts, token, feed)
}

func (_BlobVerifier *BlobVerifierTransactorSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _BlobVerifier.Contract.AddFeed(&_BlobVerifier.TransactOpts, token, feed)
}

func (_BlobVerifier *BlobVerifierTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "addPool", token, pool)
}

func (_BlobVerifier *BlobVerifierSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _BlobVerifier.Contract.AddPool(&_BlobVerifier.TransactOpts, token, pool)
}

func (_BlobVerifier *BlobVerifierTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _BlobVerifier.Contract.AddPool(&_BlobVerifier.TransactOpts, token, pool)
}

func (_BlobVerifier *BlobVerifierTransactor) ExecuteTransaction(opts *bind.TransactOpts, report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "executeTransaction", report, needFee)
}

func (_BlobVerifier *BlobVerifierSession) ExecuteTransaction(report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _BlobVerifier.Contract.ExecuteTransaction(&_BlobVerifier.TransactOpts, report, needFee)
}

func (_BlobVerifier *BlobVerifierTransactorSession) ExecuteTransaction(report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _BlobVerifier.Contract.ExecuteTransaction(&_BlobVerifier.TransactOpts, report, needFee)
}

func (_BlobVerifier *BlobVerifierTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "pause")
}

func (_BlobVerifier *BlobVerifierSession) Pause() (*types.Transaction, error) {
	return _BlobVerifier.Contract.Pause(&_BlobVerifier.TransactOpts)
}

func (_BlobVerifier *BlobVerifierTransactorSession) Pause() (*types.Transaction, error) {
	return _BlobVerifier.Contract.Pause(&_BlobVerifier.TransactOpts)
}

func (_BlobVerifier *BlobVerifierTransactor) RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "removeFeed", token, feed)
}

func (_BlobVerifier *BlobVerifierSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _BlobVerifier.Contract.RemoveFeed(&_BlobVerifier.TransactOpts, token, feed)
}

func (_BlobVerifier *BlobVerifierTransactorSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _BlobVerifier.Contract.RemoveFeed(&_BlobVerifier.TransactOpts, token, feed)
}

func (_BlobVerifier *BlobVerifierTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "removePool", token, pool)
}

func (_BlobVerifier *BlobVerifierSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _BlobVerifier.Contract.RemovePool(&_BlobVerifier.TransactOpts, token, pool)
}

func (_BlobVerifier *BlobVerifierTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _BlobVerifier.Contract.RemovePool(&_BlobVerifier.TransactOpts, token, pool)
}

func (_BlobVerifier *BlobVerifierTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "setAFN", afn)
}

func (_BlobVerifier *BlobVerifierSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _BlobVerifier.Contract.SetAFN(&_BlobVerifier.TransactOpts, afn)
}

func (_BlobVerifier *BlobVerifierTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _BlobVerifier.Contract.SetAFN(&_BlobVerifier.TransactOpts, afn)
}

func (_BlobVerifier *BlobVerifierTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_BlobVerifier *BlobVerifierSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _BlobVerifier.Contract.SetConfig(&_BlobVerifier.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_BlobVerifier *BlobVerifierTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _BlobVerifier.Contract.SetConfig(&_BlobVerifier.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_BlobVerifier *BlobVerifierTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_BlobVerifier *BlobVerifierSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _BlobVerifier.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_BlobVerifier.TransactOpts, newTime)
}

func (_BlobVerifier *BlobVerifierTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _BlobVerifier.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_BlobVerifier.TransactOpts, newTime)
}

func (_BlobVerifier *BlobVerifierTransactor) SetOffRampConfig(opts *bind.TransactOpts, config TollOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "setOffRampConfig", config)
}

func (_BlobVerifier *BlobVerifierSession) SetOffRampConfig(config TollOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _BlobVerifier.Contract.SetOffRampConfig(&_BlobVerifier.TransactOpts, config)
}

func (_BlobVerifier *BlobVerifierTransactorSession) SetOffRampConfig(config TollOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _BlobVerifier.Contract.SetOffRampConfig(&_BlobVerifier.TransactOpts, config)
}

func (_BlobVerifier *BlobVerifierTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "setRouter", router)
}

func (_BlobVerifier *BlobVerifierSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _BlobVerifier.Contract.SetRouter(&_BlobVerifier.TransactOpts, router)
}

func (_BlobVerifier *BlobVerifierTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _BlobVerifier.Contract.SetRouter(&_BlobVerifier.TransactOpts, router)
}

func (_BlobVerifier *BlobVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "transferOwnership", to)
}

func (_BlobVerifier *BlobVerifierSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BlobVerifier.Contract.TransferOwnership(&_BlobVerifier.TransactOpts, to)
}

func (_BlobVerifier *BlobVerifierTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BlobVerifier.Contract.TransferOwnership(&_BlobVerifier.TransactOpts, to)
}

func (_BlobVerifier *BlobVerifierTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_BlobVerifier *BlobVerifierSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _BlobVerifier.Contract.Transmit(&_BlobVerifier.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_BlobVerifier *BlobVerifierTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _BlobVerifier.Contract.Transmit(&_BlobVerifier.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_BlobVerifier *BlobVerifierTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "unpause")
}

func (_BlobVerifier *BlobVerifierSession) Unpause() (*types.Transaction, error) {
	return _BlobVerifier.Contract.Unpause(&_BlobVerifier.TransactOpts)
}

func (_BlobVerifier *BlobVerifierTransactorSession) Unpause() (*types.Transaction, error) {
	return _BlobVerifier.Contract.Unpause(&_BlobVerifier.TransactOpts)
}

func (_BlobVerifier *BlobVerifierTransactor) WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BlobVerifier.contract.Transact(opts, "withdrawAccumulatedFees", feeToken, recipient, amount)
}

func (_BlobVerifier *BlobVerifierSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BlobVerifier.Contract.WithdrawAccumulatedFees(&_BlobVerifier.TransactOpts, feeToken, recipient, amount)
}

func (_BlobVerifier *BlobVerifierTransactorSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BlobVerifier.Contract.WithdrawAccumulatedFees(&_BlobVerifier.TransactOpts, feeToken, recipient, amount)
}

type BlobVerifierAFNMaxHeartbeatTimeSetIterator struct {
	Event *BlobVerifierAFNMaxHeartbeatTimeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierAFNMaxHeartbeatTimeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierAFNMaxHeartbeatTimeSet)
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
		it.Event = new(BlobVerifierAFNMaxHeartbeatTimeSet)
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

func (it *BlobVerifierAFNMaxHeartbeatTimeSetIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierAFNMaxHeartbeatTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierAFNMaxHeartbeatTimeSet struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*BlobVerifierAFNMaxHeartbeatTimeSetIterator, error) {

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierAFNMaxHeartbeatTimeSetIterator{contract: _BlobVerifier.contract, event: "AFNMaxHeartbeatTimeSet", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierAFNMaxHeartbeatTimeSet) (event.Subscription, error) {

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierAFNMaxHeartbeatTimeSet)
				if err := _BlobVerifier.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParseAFNMaxHeartbeatTimeSet(log types.Log) (*BlobVerifierAFNMaxHeartbeatTimeSet, error) {
	event := new(BlobVerifierAFNMaxHeartbeatTimeSet)
	if err := _BlobVerifier.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierAFNSetIterator struct {
	Event *BlobVerifierAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierAFNSet)
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
		it.Event = new(BlobVerifierAFNSet)
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

func (it *BlobVerifierAFNSetIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterAFNSet(opts *bind.FilterOpts) (*BlobVerifierAFNSetIterator, error) {

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierAFNSetIterator{contract: _BlobVerifier.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierAFNSet) (event.Subscription, error) {

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierAFNSet)
				if err := _BlobVerifier.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParseAFNSet(log types.Log) (*BlobVerifierAFNSet, error) {
	event := new(BlobVerifierAFNSet)
	if err := _BlobVerifier.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierConfigSetIterator struct {
	Event *BlobVerifierConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierConfigSet)
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
		it.Event = new(BlobVerifierConfigSet)
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

func (it *BlobVerifierConfigSetIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierConfigSet struct {
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

func (_BlobVerifier *BlobVerifierFilterer) FilterConfigSet(opts *bind.FilterOpts) (*BlobVerifierConfigSetIterator, error) {

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierConfigSetIterator{contract: _BlobVerifier.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierConfigSet) (event.Subscription, error) {

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierConfigSet)
				if err := _BlobVerifier.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParseConfigSet(log types.Log) (*BlobVerifierConfigSet, error) {
	event := new(BlobVerifierConfigSet)
	if err := _BlobVerifier.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierCrossChainMessageExecutedIterator struct {
	Event *BlobVerifierCrossChainMessageExecuted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierCrossChainMessageExecutedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierCrossChainMessageExecuted)
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
		it.Event = new(BlobVerifierCrossChainMessageExecuted)
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

func (it *BlobVerifierCrossChainMessageExecutedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierCrossChainMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierCrossChainMessageExecuted struct {
	SequenceNumber uint64
	Raw            types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterCrossChainMessageExecuted(opts *bind.FilterOpts, sequenceNumber []uint64) (*BlobVerifierCrossChainMessageExecutedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "CrossChainMessageExecuted", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &BlobVerifierCrossChainMessageExecutedIterator{contract: _BlobVerifier.contract, event: "CrossChainMessageExecuted", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchCrossChainMessageExecuted(opts *bind.WatchOpts, sink chan<- *BlobVerifierCrossChainMessageExecuted, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "CrossChainMessageExecuted", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierCrossChainMessageExecuted)
				if err := _BlobVerifier.contract.UnpackLog(event, "CrossChainMessageExecuted", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParseCrossChainMessageExecuted(log types.Log) (*BlobVerifierCrossChainMessageExecuted, error) {
	event := new(BlobVerifierCrossChainMessageExecuted)
	if err := _BlobVerifier.contract.UnpackLog(event, "CrossChainMessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierFeedAddedIterator struct {
	Event *BlobVerifierFeedAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierFeedAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierFeedAdded)
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
		it.Event = new(BlobVerifierFeedAdded)
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

func (it *BlobVerifierFeedAddedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierFeedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierFeedAdded struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterFeedAdded(opts *bind.FilterOpts) (*BlobVerifierFeedAddedIterator, error) {

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierFeedAddedIterator{contract: _BlobVerifier.contract, event: "FeedAdded", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *BlobVerifierFeedAdded) (event.Subscription, error) {

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierFeedAdded)
				if err := _BlobVerifier.contract.UnpackLog(event, "FeedAdded", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParseFeedAdded(log types.Log) (*BlobVerifierFeedAdded, error) {
	event := new(BlobVerifierFeedAdded)
	if err := _BlobVerifier.contract.UnpackLog(event, "FeedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierFeedRemovedIterator struct {
	Event *BlobVerifierFeedRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierFeedRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierFeedRemoved)
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
		it.Event = new(BlobVerifierFeedRemoved)
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

func (it *BlobVerifierFeedRemovedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierFeedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierFeedRemoved struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterFeedRemoved(opts *bind.FilterOpts) (*BlobVerifierFeedRemovedIterator, error) {

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierFeedRemovedIterator{contract: _BlobVerifier.contract, event: "FeedRemoved", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *BlobVerifierFeedRemoved) (event.Subscription, error) {

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierFeedRemoved)
				if err := _BlobVerifier.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParseFeedRemoved(log types.Log) (*BlobVerifierFeedRemoved, error) {
	event := new(BlobVerifierFeedRemoved)
	if err := _BlobVerifier.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierFeesWithdrawnIterator struct {
	Event *BlobVerifierFeesWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierFeesWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierFeesWithdrawn)
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
		it.Event = new(BlobVerifierFeesWithdrawn)
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

func (it *BlobVerifierFeesWithdrawnIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierFeesWithdrawn struct {
	FeeToken  common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*BlobVerifierFeesWithdrawnIterator, error) {

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierFeesWithdrawnIterator{contract: _BlobVerifier.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *BlobVerifierFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierFeesWithdrawn)
				if err := _BlobVerifier.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParseFeesWithdrawn(log types.Log) (*BlobVerifierFeesWithdrawn, error) {
	event := new(BlobVerifierFeesWithdrawn)
	if err := _BlobVerifier.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierOffRampConfigSetIterator struct {
	Event *BlobVerifierOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierOffRampConfigSet)
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
		it.Event = new(BlobVerifierOffRampConfigSet)
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

func (it *BlobVerifierOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierOffRampConfigSet struct {
	Config TollOffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*BlobVerifierOffRampConfigSetIterator, error) {

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierOffRampConfigSetIterator{contract: _BlobVerifier.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierOffRampConfigSet)
				if err := _BlobVerifier.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParseOffRampConfigSet(log types.Log) (*BlobVerifierOffRampConfigSet, error) {
	event := new(BlobVerifierOffRampConfigSet)
	if err := _BlobVerifier.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierOffRampRouterSetIterator struct {
	Event *BlobVerifierOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierOffRampRouterSet)
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
		it.Event = new(BlobVerifierOffRampRouterSet)
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

func (it *BlobVerifierOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierOffRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts) (*BlobVerifierOffRampRouterSetIterator, error) {

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "OffRampRouterSet")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierOffRampRouterSetIterator{contract: _BlobVerifier.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierOffRampRouterSet) (event.Subscription, error) {

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "OffRampRouterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierOffRampRouterSet)
				if err := _BlobVerifier.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParseOffRampRouterSet(log types.Log) (*BlobVerifierOffRampRouterSet, error) {
	event := new(BlobVerifierOffRampRouterSet)
	if err := _BlobVerifier.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierOwnershipTransferRequestedIterator struct {
	Event *BlobVerifierOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierOwnershipTransferRequested)
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
		it.Event = new(BlobVerifierOwnershipTransferRequested)
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

func (it *BlobVerifierOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BlobVerifierOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BlobVerifierOwnershipTransferRequestedIterator{contract: _BlobVerifier.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BlobVerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierOwnershipTransferRequested)
				if err := _BlobVerifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParseOwnershipTransferRequested(log types.Log) (*BlobVerifierOwnershipTransferRequested, error) {
	event := new(BlobVerifierOwnershipTransferRequested)
	if err := _BlobVerifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierOwnershipTransferredIterator struct {
	Event *BlobVerifierOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierOwnershipTransferred)
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
		it.Event = new(BlobVerifierOwnershipTransferred)
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

func (it *BlobVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BlobVerifierOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BlobVerifierOwnershipTransferredIterator{contract: _BlobVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlobVerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierOwnershipTransferred)
				if err := _BlobVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*BlobVerifierOwnershipTransferred, error) {
	event := new(BlobVerifierOwnershipTransferred)
	if err := _BlobVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierPausedIterator struct {
	Event *BlobVerifierPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierPaused)
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
		it.Event = new(BlobVerifierPaused)
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

func (it *BlobVerifierPausedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterPaused(opts *bind.FilterOpts) (*BlobVerifierPausedIterator, error) {

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierPausedIterator{contract: _BlobVerifier.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BlobVerifierPaused) (event.Subscription, error) {

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierPaused)
				if err := _BlobVerifier.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParsePaused(log types.Log) (*BlobVerifierPaused, error) {
	event := new(BlobVerifierPaused)
	if err := _BlobVerifier.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierPoolAddedIterator struct {
	Event *BlobVerifierPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierPoolAdded)
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
		it.Event = new(BlobVerifierPoolAdded)
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

func (it *BlobVerifierPoolAddedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*BlobVerifierPoolAddedIterator, error) {

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierPoolAddedIterator{contract: _BlobVerifier.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *BlobVerifierPoolAdded) (event.Subscription, error) {

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierPoolAdded)
				if err := _BlobVerifier.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParsePoolAdded(log types.Log) (*BlobVerifierPoolAdded, error) {
	event := new(BlobVerifierPoolAdded)
	if err := _BlobVerifier.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierPoolRemovedIterator struct {
	Event *BlobVerifierPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierPoolRemoved)
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
		it.Event = new(BlobVerifierPoolRemoved)
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

func (it *BlobVerifierPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*BlobVerifierPoolRemovedIterator, error) {

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierPoolRemovedIterator{contract: _BlobVerifier.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *BlobVerifierPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierPoolRemoved)
				if err := _BlobVerifier.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParsePoolRemoved(log types.Log) (*BlobVerifierPoolRemoved, error) {
	event := new(BlobVerifierPoolRemoved)
	if err := _BlobVerifier.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierReportAcceptedIterator struct {
	Event *BlobVerifierReportAccepted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierReportAcceptedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierReportAccepted)
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
		it.Event = new(BlobVerifierReportAccepted)
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

func (it *BlobVerifierReportAcceptedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierReportAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierReportAccepted struct {
	Report CCIPRelayReport
	Raw    types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterReportAccepted(opts *bind.FilterOpts) (*BlobVerifierReportAcceptedIterator, error) {

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "ReportAccepted")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierReportAcceptedIterator{contract: _BlobVerifier.contract, event: "ReportAccepted", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchReportAccepted(opts *bind.WatchOpts, sink chan<- *BlobVerifierReportAccepted) (event.Subscription, error) {

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "ReportAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierReportAccepted)
				if err := _BlobVerifier.contract.UnpackLog(event, "ReportAccepted", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParseReportAccepted(log types.Log) (*BlobVerifierReportAccepted, error) {
	event := new(BlobVerifierReportAccepted)
	if err := _BlobVerifier.contract.UnpackLog(event, "ReportAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierTransmittedIterator struct {
	Event *BlobVerifierTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierTransmitted)
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
		it.Event = new(BlobVerifierTransmitted)
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

func (it *BlobVerifierTransmittedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterTransmitted(opts *bind.FilterOpts) (*BlobVerifierTransmittedIterator, error) {

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierTransmittedIterator{contract: _BlobVerifier.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *BlobVerifierTransmitted) (event.Subscription, error) {

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierTransmitted)
				if err := _BlobVerifier.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParseTransmitted(log types.Log) (*BlobVerifierTransmitted, error) {
	event := new(BlobVerifierTransmitted)
	if err := _BlobVerifier.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierUnpausedIterator struct {
	Event *BlobVerifierUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierUnpaused)
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
		it.Event = new(BlobVerifierUnpaused)
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

func (it *BlobVerifierUnpausedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_BlobVerifier *BlobVerifierFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BlobVerifierUnpausedIterator, error) {

	logs, sub, err := _BlobVerifier.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierUnpausedIterator{contract: _BlobVerifier.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_BlobVerifier *BlobVerifierFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BlobVerifierUnpaused) (event.Subscription, error) {

	logs, sub, err := _BlobVerifier.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierUnpaused)
				if err := _BlobVerifier.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_BlobVerifier *BlobVerifierFilterer) ParseUnpaused(log types.Log) (*BlobVerifierUnpaused, error) {
	event := new(BlobVerifierUnpaused)
	if err := _BlobVerifier.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_BlobVerifier *BlobVerifier) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BlobVerifier.abi.Events["AFNMaxHeartbeatTimeSet"].ID:
		return _BlobVerifier.ParseAFNMaxHeartbeatTimeSet(log)
	case _BlobVerifier.abi.Events["AFNSet"].ID:
		return _BlobVerifier.ParseAFNSet(log)
	case _BlobVerifier.abi.Events["ConfigSet"].ID:
		return _BlobVerifier.ParseConfigSet(log)
	case _BlobVerifier.abi.Events["CrossChainMessageExecuted"].ID:
		return _BlobVerifier.ParseCrossChainMessageExecuted(log)
	case _BlobVerifier.abi.Events["FeedAdded"].ID:
		return _BlobVerifier.ParseFeedAdded(log)
	case _BlobVerifier.abi.Events["FeedRemoved"].ID:
		return _BlobVerifier.ParseFeedRemoved(log)
	case _BlobVerifier.abi.Events["FeesWithdrawn"].ID:
		return _BlobVerifier.ParseFeesWithdrawn(log)
	case _BlobVerifier.abi.Events["OffRampConfigSet"].ID:
		return _BlobVerifier.ParseOffRampConfigSet(log)
	case _BlobVerifier.abi.Events["OffRampRouterSet"].ID:
		return _BlobVerifier.ParseOffRampRouterSet(log)
	case _BlobVerifier.abi.Events["OwnershipTransferRequested"].ID:
		return _BlobVerifier.ParseOwnershipTransferRequested(log)
	case _BlobVerifier.abi.Events["OwnershipTransferred"].ID:
		return _BlobVerifier.ParseOwnershipTransferred(log)
	case _BlobVerifier.abi.Events["Paused"].ID:
		return _BlobVerifier.ParsePaused(log)
	case _BlobVerifier.abi.Events["PoolAdded"].ID:
		return _BlobVerifier.ParsePoolAdded(log)
	case _BlobVerifier.abi.Events["PoolRemoved"].ID:
		return _BlobVerifier.ParsePoolRemoved(log)
	case _BlobVerifier.abi.Events["ReportAccepted"].ID:
		return _BlobVerifier.ParseReportAccepted(log)
	case _BlobVerifier.abi.Events["Transmitted"].ID:
		return _BlobVerifier.ParseTransmitted(log)
	case _BlobVerifier.abi.Events["Unpaused"].ID:
		return _BlobVerifier.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (BlobVerifierAFNMaxHeartbeatTimeSet) Topic() common.Hash {
	return common.HexToHash("0x72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251")
}

func (BlobVerifierAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (BlobVerifierConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (BlobVerifierCrossChainMessageExecuted) Topic() common.Hash {
	return common.HexToHash("0x88a2c71ae86800edd209d83430f348ce70539bd40085d045340ff569f058a2a0")
}

func (BlobVerifierFeedAdded) Topic() common.Hash {
	return common.HexToHash("0x037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb601305")
}

func (BlobVerifierFeedRemoved) Topic() common.Hash {
	return common.HexToHash("0xa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d")
}

func (BlobVerifierFeesWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8")
}

func (BlobVerifierOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xe4cd88b1b5e20a0b843af3207ba74d3f84af4acff45830469490ac9c6ab8a745")
}

func (BlobVerifierOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4")
}

func (BlobVerifierOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BlobVerifierOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BlobVerifierPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (BlobVerifierPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (BlobVerifierPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (BlobVerifierReportAccepted) Topic() common.Hash {
	return common.HexToHash("0x6a2424b54e9b18f0c3d440b6003882f5ffa34272205608ad9451c921c0d99a35")
}

func (BlobVerifierTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (BlobVerifierUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_BlobVerifier *BlobVerifier) Address() common.Address {
	return _BlobVerifier.address
}

type BlobVerifierInterface interface {
	CHAINID(opts *bind.CallOpts) (*big.Int, error)

	SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetExecuted(opts *bind.CallOpts, sequenceNumber uint64) (bool, error)

	GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error)

	GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetLastReport(opts *bind.CallOpts) (CCIPRelayReport, error)

	GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error)

	GetMerkleRoot(opts *bind.CallOpts, root [32]byte) (*big.Int, error)

	GetOffRampConfig(opts *bind.CallOpts) (TollOffRampInterfaceOffRampConfig, error)

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

	SetOffRampConfig(opts *bind.TransactOpts, config TollOffRampInterfaceOffRampConfig) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*BlobVerifierAFNMaxHeartbeatTimeSetIterator, error)

	WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierAFNMaxHeartbeatTimeSet) (event.Subscription, error)

	ParseAFNMaxHeartbeatTimeSet(log types.Log) (*BlobVerifierAFNMaxHeartbeatTimeSet, error)

	FilterAFNSet(opts *bind.FilterOpts) (*BlobVerifierAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*BlobVerifierAFNSet, error)

	FilterConfigSet(opts *bind.FilterOpts) (*BlobVerifierConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*BlobVerifierConfigSet, error)

	FilterCrossChainMessageExecuted(opts *bind.FilterOpts, sequenceNumber []uint64) (*BlobVerifierCrossChainMessageExecutedIterator, error)

	WatchCrossChainMessageExecuted(opts *bind.WatchOpts, sink chan<- *BlobVerifierCrossChainMessageExecuted, sequenceNumber []uint64) (event.Subscription, error)

	ParseCrossChainMessageExecuted(log types.Log) (*BlobVerifierCrossChainMessageExecuted, error)

	FilterFeedAdded(opts *bind.FilterOpts) (*BlobVerifierFeedAddedIterator, error)

	WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *BlobVerifierFeedAdded) (event.Subscription, error)

	ParseFeedAdded(log types.Log) (*BlobVerifierFeedAdded, error)

	FilterFeedRemoved(opts *bind.FilterOpts) (*BlobVerifierFeedRemovedIterator, error)

	WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *BlobVerifierFeedRemoved) (event.Subscription, error)

	ParseFeedRemoved(log types.Log) (*BlobVerifierFeedRemoved, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*BlobVerifierFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *BlobVerifierFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*BlobVerifierFeesWithdrawn, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*BlobVerifierOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*BlobVerifierOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts) (*BlobVerifierOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierOffRampRouterSet) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*BlobVerifierOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BlobVerifierOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BlobVerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BlobVerifierOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BlobVerifierOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlobVerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BlobVerifierOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*BlobVerifierPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *BlobVerifierPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*BlobVerifierPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*BlobVerifierPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *BlobVerifierPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*BlobVerifierPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*BlobVerifierPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *BlobVerifierPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*BlobVerifierPoolRemoved, error)

	FilterReportAccepted(opts *bind.FilterOpts) (*BlobVerifierReportAcceptedIterator, error)

	WatchReportAccepted(opts *bind.WatchOpts, sink chan<- *BlobVerifierReportAccepted) (event.Subscription, error)

	ParseReportAccepted(log types.Log) (*BlobVerifierReportAccepted, error)

	FilterTransmitted(opts *bind.FilterOpts) (*BlobVerifierTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *BlobVerifierTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*BlobVerifierTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*BlobVerifierUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BlobVerifierUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*BlobVerifierUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
