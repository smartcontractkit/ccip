// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blob_verifier_helper

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

var BlobVerifierHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"contractAggregatorV2V3Interface[]\",\"name\":\"feeds\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionDelayError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"InvalidExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceFeedConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"MerkleProofError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RelayReportError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastMaxSequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newMinSequenceNumber\",\"type\":\"uint64\"}],\"name\":\"SequenceError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenFeedMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnexpectedPayloadData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"CrossChainMessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minSequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceNumber\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structCCIP.RelayReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"ReportAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"addFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagsBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"needFee\",\"type\":\"bool\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastReport\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minSequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceNumber\",\"type\":\"uint64\"}],\"internalType\":\"structCCIP.RelayReport\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRampConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagsBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"removeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"merkle\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setOffRampConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005fcb38038062005fcb833981016040819052620000349162000703565b6040805160808101825260018082526001600160401b0385811660208401526103e89383019390935291831660608201526000805460ff191681558b928b928b928b928b928b928b92908790869082908990889088903390819081620000e15760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b0319909216919091179091558116156200011b576200011b8162000434565b5050506001600160a01b038216158062000133575080155b156200015257604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001985760405162d8548360e71b815260040160405180910390fd5b8151620001ad906005906020850190620004e5565b5060005b825181101562000291576000828281518110620001d257620001d2620007eb565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b0316815250600460008685815181106200021c576200021c620007eb565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff1916600117905580620002888162000801565b915050620001b1565b5050508051825114620002b75760405163ee9d106b60e01b815260040160405180910390fd5b8151620002cc906008906020850190620004e5565b5060005b825181101562000399576000828281518110620002f157620002f1620007eb565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b0316815250600760008685815181106200033b576200033b620007eb565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b02929091169190911790555080620003908162000801565b915050620002d0565b505050151560805260a09790975250505060c0929092525050805160138054602084015160408501516060909501516001600160401b03908116600160c01b026001600160c01b03968216600160801b02969096166001600160801b0392821668010000000000000000026001600160801b031990941691909516179190911716919091179190911790555062000829975050505050505050565b336001600160a01b038216036200048e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000d8565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200053d579160200282015b828111156200053d57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000506565b506200054b9291506200054f565b5090565b5b808211156200054b576000815560010162000550565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620005a757620005a762000566565b604052919050565b60006001600160401b03821115620005cb57620005cb62000566565b5060051b60200190565b6001600160a01b0381168114620005eb57600080fd5b50565b600082601f8301126200060057600080fd5b81516020620006196200061383620005af565b6200057c565b82815260059290921b840181019181810190868411156200063957600080fd5b8286015b84811015620006615780516200065381620005d5565b83529183019183016200063d565b509695505050505050565b600082601f8301126200067e57600080fd5b81516020620006916200061383620005af565b82815260059290921b84018101918181019086841115620006b157600080fd5b8286015b8481101562000661578051620006cb81620005d5565b8352918301918301620006b5565b8051620006e681620005d5565b919050565b80516001600160401b0381168114620006e657600080fd5b60008060008060008060008060006101208a8c0312156200072357600080fd5b895160208b015160408c0151919a5098506001600160401b03808211156200074a57600080fd5b620007588d838e01620005ee565b985060608c01519150808211156200076f57600080fd5b6200077d8d838e01620005ee565b975060808c01519150808211156200079457600080fd5b50620007a38c828d016200066c565b955050620007b460a08b01620006d9565b935060c08a01519250620007cb60e08b01620006eb565b9150620007dc6101008b01620006eb565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b6000600182016200082257634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05161576b62000860600039600061053f01526000818161048e01526136910152600061235b015261576b6000f3fe608060405234801561001057600080fd5b50600436106102ad5760003560e01c806381be8fa41161017b578063b0f479a1116100d8578063c0d786551161008c578063eb511dd411610071578063eb511dd414610768578063eefa7a3e1461077b578063f2fde38b1461080b57600080fd5b8063c0d7865514610742578063e3d0e7121461075557600080fd5b8063b5767166116100bd578063b5767166146106e3578063b6608c3b146106f6578063bbe4f6db1461070957600080fd5b8063b0f479a1146106b2578063b1dc65a4146106d057600080fd5b80638da5cb5b1161012f578063a8ebd0f411610114578063a8ebd0f4146105ac578063afcb95d71461068a578063b034909c146106aa57600080fd5b80638da5cb5b14610569578063a7206cd61461058c57600080fd5b80638456cb59116101605780638456cb591461053257806385e1f4d01461053a57806389c065681461056157600080fd5b806381be8fa4146104fa57806381ff70481461050257600080fd5b8063567c814b11610229578063744b92e2116101dd57806379ba5097116101c257806379ba5097146104b057806380d9a1b7146104b857806381411834146104e557600080fd5b8063744b92e21461047657806374be21501461048957600080fd5b806359e96b5b1161020e57806359e96b5b1461041f5780635b16ebb7146104325780635c975abb1461046b57600080fd5b8063567c814b146103e95780635853c6271461040c57600080fd5b8063295938ec116102805780633dd80c70116102655780633dd80c70146103ad5780633f4ba83a146103ce578063461c551b146103d657600080fd5b8063295938ec146103875780632b898c251461039a57600080fd5b8063108ee5fc146102b257806316b8e731146102c7578063181f5a771461032a5780632222dd4214610369575b600080fd5b6102c56102c0366004614614565b61081e565b005b6103006102d5366004614614565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b604080518082018252600d81527f4f666652616d7020302e302e31000000000000000000000000000000000000006020820152905161032191906146b6565b60025473ffffffffffffffffffffffffffffffffffffffff16610300565b6102c5610395366004614b35565b6108fa565b6102c56103a8366004614b87565b6111c3565b6103c06103bb366004614bb5565b611593565b604051908152602001610321565b6102c561176d565b6102c56103e4366004614bea565b61177f565b6103fc6103f7366004614c02565b6117d1565b6040519015158152602001610321565b6102c561041a366004614b87565b611918565b6102c561042d366004614c1b565b611b30565b6103fc610440366004614614565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff166103fc565b6102c5610484366004614b87565b611bae565b6103c07f000000000000000000000000000000000000000000000000000000000000000081565b6102c5611fa3565b6103fc6104c6366004614c5c565b67ffffffffffffffff1660009081526010602052604090205460ff1690565b6104ed6120c5565b6040516103219190614cca565b6104ed612134565b600b546009546040805163ffffffff80851682526401000000009094049093166020840152820152606001610321565b6102c56121a1565b6103c07f000000000000000000000000000000000000000000000000000000000000000081565b6104ed6121b1565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610300565b6103c061059a366004614c02565b6000908152600f602052604090205490565b610646604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260135467ffffffffffffffff808216835268010000000000000000820481166020840152700100000000000000000000000000000000820481169383019390935278010000000000000000000000000000000000000000000000009004909116606082015290565b6040516103219190815167ffffffffffffffff9081168252602080840151821690830152604080840151821690830152606092830151169181019190915260800190565b604080516001815260006020820181905291810191909152606001610321565b6003546103c0565b60145473ffffffffffffffffffffffffffffffffffffffff16610300565b6102c56106de366004614d29565b61221e565b6102c56106f1366004614e0e565b6128c7565b6102c5610704366004614c02565b6128d6565b610300610717366004614614565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6102c5610750366004614614565b612956565b6102c5610763366004614eb8565b6129d1565b6102c5610776366004614b87565b6133b6565b6107d760408051606081018252600080825260208201819052918101919091525060408051606081018252601154815260125467ffffffffffffffff808216602084015268010000000000000000909104169181019190915290565b604080518251815260208084015167ffffffffffffffff908116918301919091529282015190921690820152606001610321565b6102c5610819366004614614565b6135f6565b610826613607565b73ffffffffffffffffffffffffffffffffffffffff8116610873576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b60005460ff161561096c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109fd9190614f85565b15610a33576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa158015610aa3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac79190614fa2565b9050600354816020015142610adc919061500d565b1115610b14576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60145473ffffffffffffffffffffffffffffffffffffffff16610b63576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610b6e84611593565b6000818152600f6020526040812054919250819003610bbc576040517f851bdf5300000000000000000000000000000000000000000000000000000000815260048101839052602401610963565b6013544290610be19068010000000000000000900467ffffffffffffffff1683615024565b10610c18576040517f15c33ba200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8551518110156111bb57600086600001518281518110610c3d57610c3d61503c565b6020908102919091018101518082015167ffffffffffffffff166000908152601090925260409091205490915060ff1615610cb65760208101516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610963565b610cbf8161368d565b610cc8816137f8565b60208082015167ffffffffffffffff16600090815260109091526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558515610f0c576000808260a00151600081518110610d2f57610d2f61503c565b602002602001015190506000610d6a8273ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff8116610db9576040517f83135fec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166350d25bcd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e28919061506b565b601354610e3f919067ffffffffffffffff16615084565b92508215610f0857828460c00151600081518110610e5f57610e5f61503c565b60200260200101818151610e73919061500d565b905250610e7f8261389e565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081523360048201526024810185905273ffffffffffffffffffffffffffffffffffffffff919091169063ea6192a290604401600060405180830381600087803b158015610eef57600080fd5b505af1158015610f03573d6000803e3d6000fd5b505050505b5050505b60005b8160a001515181101561100357610f428260a001518281518110610f3557610f3561503c565b602002602001015161389e565b73ffffffffffffffffffffffffffffffffffffffff1663ea6192a283606001518460c001518481518110610f7857610f7861503c565b60200260200101516040518363ffffffff1660e01b8152600401610fbe92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600060405180830381600087803b158015610fd857600080fd5b505af1158015610fec573d6000803e3d6000fd5b505050508080610ffb906150c1565b915050610f0f565b50606081015173ffffffffffffffffffffffffffffffffffffffff163b1561111e5760145460608201516040517ffd12f6e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9092169163fd12f6e591611080918590600401615232565b600060405180830381600087803b15801561109a57600080fd5b505af19250505080156110ab575060015b611119573d8080156110d9576040519150601f19603f3d011682016040523d82523d6000602084013e6110de565b606091505b508160200151816040517fa1dc8185000000000000000000000000000000000000000000000000000000008152600401610963929190615261565b61116d565b6080810151511561116d5760208101516040517fc945cae000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610963565b806020015167ffffffffffffffff167f88a2c71ae86800edd209d83430f348ce70539bd40085d045340ff569f058a2a060405160405180910390a250806111b3816150c1565b915050610c1b565b505050505050565b6111cb613607565b6008546000819003611209576040517f2e70248b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906112a4576040517f3917193900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff161461130d576040517f9403a50500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600861131c60018561500d565b8154811061132c5761132c61503c565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600882602001516bffffffffffffffffffffffff168154811061137e5761137e61503c565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1660086113ad60018661500d565b815481106113bd576113bd61503c565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600883602001516bffffffffffffffffffffffff168154811061142b5761142b61503c565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526007909152604090912080546bffffffffffffffffffffffff90921674010000000000000000000000000000000000000000029190921617905560088054806114cd576114cd615284565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600783526040808520949094558351908152908716918101919091527fa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d91015b60405180910390a15050505050565b8051516020820151516000919082906001906115af9084615024565b6115b9919061500d565b90506101008111156115ca57600080fd5b60008167ffffffffffffffff8111156115e5576115e56146c9565b60405190808252806020026020018201604052801561160e578160200160208202803683370190505b5090506000806000805b85811015611716576040890151811c6001908116146116f08161165d5760208b0151805160018601959081106116505761165061503c565b60200260200101516116a4565b88861061167b5786516001860195889181106116505761165061503c565b8a51805160018801976116a4929181106116975761169761503c565b602002602001015161391a565b8987106116cf5787516001870196899181106116c2576116c261503c565b602002602001015161398a565b8b51805160018901986116eb929181106116975761169761503c565b61398a565b8683815181106117025761170261503c565b602090810291909101015250600101611618565b508415611747578360018603815181106117325761173261503c565b60200260200101519650505050505050919050565b61176188600001516000815181106116975761169761503c565b98975050505050505050565b611775613607565b61177d613a48565b565b611787613607565b80601361179482826152b3565b9050507fe4cd88b1b5e20a0b843af3207ba74d3f84af4acff45830469490ac9c6ab8a745816040516117c691906153c4565b60405180910390a150565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa158015611841573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118659190614f85565b1580156119125750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa1580156118dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119019190614fa2565b6020015161190f908461500d565b11155b92915050565b611920613607565b73ffffffffffffffffffffffffffffffffffffffff82161580611957575073ffffffffffffffffffffffffffffffffffffffff8116155b1561198e576040517fee9d106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015611a2a576040517f965ffb7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600880546bffffffffffffffffffffffff908116602080870191825288861660008181526007835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821790558351908152908101919091527f037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb60130591015b60405180910390a1505050565b611b38613607565b611b5973ffffffffffffffffffffffffffffffffffffffff84168383613b29565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa890606001611b23565b611bb6613607565b6005546000819003611bf4576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290611c8f576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614611cf8576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006005611d0760018561500d565b81548110611d1757611d1761503c565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff1681548110611d6957611d6961503c565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166005611d9860018661500d565b81548110611da857611da861503c565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110611e1657611e1661503c565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480611eb857611eb8615284565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101611584565b60015473ffffffffffffffffffffffffffffffffffffffff163314612024576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610963565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600e80548060200260200160405190810160405280929190818152602001828054801561212a57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116120ff575b5050505050905090565b6060600880548060200260200160405190810160405280929190818152602001828054801561212a5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116120ff575050505050905090565b6121a9613607565b61177d613bbb565b6060600580548060200260200160405190810160405280929190818152602001828054801561212a5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116120ff575050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161227491849163ffffffff851691908e908e9081908401838280828437600092019190915250613c7b92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600954808252600a5460ff80821660208501526101009091041692820192909252908314612349576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610963565b6123578b8b8b8b8b8b614080565b60007f0000000000000000000000000000000000000000000000000000000000000000156123b457600282602001518360400151612395919061542e565b61239f9190615453565b6123aa90600161542e565b60ff1690506123ca565b60208201516123c490600161542e565b60ff1690505b888114612433576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610963565b88871461249c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610963565b336000908152600c602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156124df576124df61549c565b60028111156124f0576124f061549c565b905250905060028160200151600281111561250d5761250d61549c565b1480156125545750600e816000015160ff168154811061252f5761252f61503c565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b6125ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610963565b5050505050600088886040516125d19291906154cb565b6040519081900381206125e8918c906020016154db565b6040516020818303038152906040528051906020012090506126086145d3565b604080518082019091526000808252602082015260005b888110156128a557600060018588846020811061263e5761263e61503c565b61264b91901a601b61542e565b8d8d8681811061265d5761265d61503c565b905060200201358c8c878181106126765761267661503c565b90506020020135604051600081526020016040526040516126b3949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156126d5573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600c602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156127555761275561549c565b60028111156127665761276661549c565b90525092506001836020015160028111156127835761278361549c565b146127ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610963565b8251849060ff16601f81106128015761280161503c565b60200201511561286d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610963565b600184846000015160ff16601f81106128885761288861503c565b91151560209092020152508061289d816150c1565b91505061261f565b5050505063ffffffff81106128bc576128bc6154f7565b505050505050505050565b6128d360008083613c7b565b50565b6128de613607565b80600003612918576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c25191016108ee565b61295e613607565b601480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4906020016117c6565b855185518560ff16601f831115612a44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610963565b60008111612aae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610963565b818314612b3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610963565b612b47816003615084565b8311612baf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610963565b612bb7613607565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600d5415612daa57600d54600090612c0f9060019061500d565b90506000600d8281548110612c2657612c2661503c565b6000918252602082200154600e805473ffffffffffffffffffffffffffffffffffffffff90921693509084908110612c6057612c6061503c565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600c909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600d80549192509080612ce057612ce0615284565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600e805480612d4957612d49615284565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550612bf5915050565b60005b815151811015613211576000600c600084600001518481518110612dd357612dd361503c565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612e1d57612e1d61549c565b14612e84576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610963565b6040805180820190915260ff821681526001602082015282518051600c9160009185908110612eb557612eb561503c565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612f5657612f5661549c565b021790555060009150612f669050565b600c600084602001518481518110612f8057612f8061503c565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612fca57612fca61549c565b14613031576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610963565b6040805180820190915260ff82168152602081016002815250600c6000846020015184815181106130645761306461503c565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156131055761310561549c565b02179055505082518051600d9250839081106131235761312361503c565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600e91908390811061319f5761319f61503c565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691909117905580613209816150c1565b915050612dad565b506040810151600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600b80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff4381168202928317855590830481169360019390926000926132a3928692908216911617615526565b92506101000a81548163ffffffff021916908363ffffffff1602179055506133024630600b60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151614137565b600981905582518051600a805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600b5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05986133a1988b98919763ffffffff90921696909591949193919261554e565b60405180910390a15050505050505050505050565b6133be613607565b73ffffffffffffffffffffffffffffffffffffffff821615806133f5575073ffffffffffffffffffffffffffffffffffffffff8116155b1561342c576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156134c8576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c9101611b23565b6135fe613607565b6128d3816141e2565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16331461177d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610963565b80517f0000000000000000000000000000000000000000000000000000000000000000146136ed5780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610963565b60135460a082015151780100000000000000000000000000000000000000000000000090910467ffffffffffffffff16108061373357508060c00151518160a001515114155b1561376a576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60135460808201515170010000000000000000000000000000000090910467ffffffffffffffff1610156128d3576013546080820151516040517f8693378900000000000000000000000000000000000000000000000000000000815270010000000000000000000000000000000090920467ffffffffffffffff1660048301526024820152604401610963565b606081015173ffffffffffffffffffffffffffffffffffffffff163014806138495750606081015173ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff165b156128d35760608101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610963565b73ffffffffffffffffffffffffffffffffffffffff8181166000908152600460205260409020541680613915576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610963565b919050565b60008060f81b8260405160200161393191906155e4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261396d92916020016155f7565b604051602081830303815290604052805190602001209050919050565b60008183106139ec57604080517f01000000000000000000000000000000000000000000000000000000000000006020808301919091526021820185905260418083018790528351808403909101815260619092019092528051910120613a41565b604080517f010000000000000000000000000000000000000000000000000000000000000060208083019190915260218201869052604180830186905283518084039091018152606190920190925280519101205b9392505050565b60005460ff16613ab4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610963565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052613bb69084906142dd565b505050565b60005460ff1615613c28576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610963565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613aff3390565b60005460ff1615613ce8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610963565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613d55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d799190614f85565b15613daf576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa158015613e1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e439190614fa2565b9050600354816020015142613e58919061500d565b1115613e90576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082806020019051810190613ea6919061563f565b9050806040015167ffffffffffffffff16816020015167ffffffffffffffff161115613efe576040517f67a3824c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160608101825260115480825260125467ffffffffffffffff80821660208501526801000000000000000090910416928201929092529015613fbb576040810151613f4d90600161568a565b67ffffffffffffffff16826020015167ffffffffffffffff1614613fbb57604080820151602084015191517f8e8c0add00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff918216600482015291166024820152604401610963565b81516000908152600f602090815260409182902042905583516011819055818501805160128054868901805167ffffffffffffffff90811668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090931694811694909417919091179091558551938452915181169383019390935251909116918101919091527f6a2424b54e9b18f0c3d440b6003882f5ffa34272205608ad9451c921c0d99a359060600160405180910390a1505050505050565b600061408d826020615084565b614098856020615084565b6140a488610144615024565b6140ae9190615024565b6140b89190615024565b6140c3906000615024565b905036811461412e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610963565b50505050505050565b6000808a8a8a8a8a8a8a8a8a60405160200161415b999897969594939291906156ad565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603614261576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610963565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600061433f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166143e99092919063ffffffff16565b805190915015613bb6578080602001905181019061435d9190614f85565b613bb6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610963565b60606143f88484600085614400565b949350505050565b606082471015614492576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610963565b843b6144fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610963565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516145239190615742565b60006040518083038185875af1925050503d8060008114614560576040519150601f19603f3d011682016040523d82523d6000602084013e614565565b606091505b5091509150614575828286614580565b979650505050505050565b6060831561458f575081613a41565b82511561459f5782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096391906146b6565b604051806103e00160405280601f906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff811681146128d357600080fd5b60006020828403121561462657600080fd5b8135613a41816145f2565b8035613915816145f2565b60005b8381101561465757818101518382015260200161463f565b83811115614666576000848401525b50505050565b6000815180845261468481602086016020860161463c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613a41602083018461466c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561471b5761471b6146c9565b60405290565b604051610140810167ffffffffffffffff8111828210171561471b5761471b6146c9565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561478c5761478c6146c9565b604052919050565b600067ffffffffffffffff8211156147ae576147ae6146c9565b5060051b60200190565b67ffffffffffffffff811681146128d357600080fd5b8035613915816147b8565b600082601f8301126147ea57600080fd5b813567ffffffffffffffff811115614804576148046146c9565b61483560207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601614745565b81815284602083860101111561484a57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261487857600080fd5b8135602061488d61488883614794565b614745565b82815260059290921b840181019181810190868411156148ac57600080fd5b8286015b848110156148d05780356148c3816145f2565b83529183019183016148b0565b509695505050505050565b600082601f8301126148ec57600080fd5b813560206148fc61488883614794565b82815260059290921b8401810191818101908684111561491b57600080fd5b8286015b848110156148d0578035835291830191830161491f565b60006060828403121561494857600080fd5b6149506146f8565b9050813567ffffffffffffffff8082111561496a57600080fd5b818401915084601f83011261497e57600080fd5b8135602061498e61488883614794565b82815260059290921b840181019181810190888411156149ad57600080fd5b8286015b84811015614aec578035868111156149c857600080fd5b8701610140818c037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018113156149fe57600080fd5b614a06614721565b868301358152614a18604084016147ce565b87820152614a2860608401614631565b6040820152614a3960808401614631565b606082015260a083013589811115614a5057600080fd5b614a5e8e89838701016147d9565b60808301525060c083013589811115614a775760008081fd5b614a858e8983870101614867565b60a08301525060e0808401358a811115614a9f5760008081fd5b614aad8f8a838801016148db565b60c084015250610100614ac1818601614631565b91830191909152610120848101359183019190915291909201359082015283529183019183016149b1565b5086525085810135935082841115614b0357600080fd5b614b0f878588016148db565b81860152505050506040820135604082015292915050565b80151581146128d357600080fd5b60008060408385031215614b4857600080fd5b823567ffffffffffffffff811115614b5f57600080fd5b614b6b85828601614936565b9250506020830135614b7c81614b27565b809150509250929050565b60008060408385031215614b9a57600080fd5b8235614ba5816145f2565b91506020830135614b7c816145f2565b600060208284031215614bc757600080fd5b813567ffffffffffffffff811115614bde57600080fd5b6143f884828501614936565b600060808284031215614bfc57600080fd5b50919050565b600060208284031215614c1457600080fd5b5035919050565b600080600060608486031215614c3057600080fd5b8335614c3b816145f2565b92506020840135614c4b816145f2565b929592945050506040919091013590565b600060208284031215614c6e57600080fd5b8135613a41816147b8565b600081518084526020808501945080840160005b83811015614cbf57815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101614c8d565b509495945050505050565b602081526000613a416020830184614c79565b60008083601f840112614cef57600080fd5b50813567ffffffffffffffff811115614d0757600080fd5b6020830191508360208260051b8501011115614d2257600080fd5b9250929050565b60008060008060008060008060e0898b031215614d4557600080fd5b606089018a811115614d5657600080fd5b8998503567ffffffffffffffff80821115614d7057600080fd5b818b0191508b601f830112614d8457600080fd5b813581811115614d9357600080fd5b8c6020828501011115614da557600080fd5b6020830199508098505060808b0135915080821115614dc357600080fd5b614dcf8c838d01614cdd565b909750955060a08b0135915080821115614de857600080fd5b50614df58b828c01614cdd565b999c989b50969995989497949560c00135949350505050565b600060208284031215614e2057600080fd5b813567ffffffffffffffff811115614e3757600080fd5b6143f8848285016147d9565b600082601f830112614e5457600080fd5b81356020614e6461488883614794565b82815260059290921b84018101918181019086841115614e8357600080fd5b8286015b848110156148d0578035614e9a816145f2565b8352918301918301614e87565b803560ff8116811461391557600080fd5b60008060008060008060c08789031215614ed157600080fd5b863567ffffffffffffffff80821115614ee957600080fd5b614ef58a838b01614e43565b97506020890135915080821115614f0b57600080fd5b614f178a838b01614e43565b9650614f2560408a01614ea7565b95506060890135915080821115614f3b57600080fd5b614f478a838b016147d9565b9450614f5560808a016147ce565b935060a0890135915080821115614f6b57600080fd5b50614f7889828a016147d9565b9150509295509295509295565b600060208284031215614f9757600080fd5b8151613a4181614b27565b600060608284031215614fb457600080fd5b614fbc6146f8565b8251815260208301516020820152604083015160408201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561501f5761501f614fde565b500390565b6000821982111561503757615037614fde565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561507d57600080fd5b5051919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156150bc576150bc614fde565b500290565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036150f2576150f2614fde565b5060010190565b600081518084526020808501945080840160005b83811015614cbf5781518752958201959082019060010161510d565b600061014082518452602083015161514d602086018267ffffffffffffffff169052565b506040830151615175604086018273ffffffffffffffffffffffffffffffffffffffff169052565b50606083015161519d606086018273ffffffffffffffffffffffffffffffffffffffff169052565b5060808301518160808601526151b58286018261466c565b91505060a083015184820360a08601526151cf8282614c79565b91505060c083015184820360c08601526151e982826150f9565b91505060e083015161521360e086018273ffffffffffffffffffffffffffffffffffffffff169052565b5061010083810151908501526101209283015192909301919091525090565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006143f86040830184615129565b67ffffffffffffffff831681526040602082015260006143f8604083018461466c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b81356152be816147b8565b67ffffffffffffffff811690508154817fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000082161783556020840135615302816147b8565b6fffffffffffffffff00000000000000008160401b16905080837fffffffffffffffffffffffffffffffff000000000000000000000000000000008416171784556040850135615351816147b8565b77ffffffffffffffff000000000000000000000000000000008160801b1690507fffffffffffffffff0000000000000000000000000000000000000000000000008185828616178417178655606087013593506153ad846147b8565b808460c01b16858417831717865550505050505050565b6080810182356153d3816147b8565b67ffffffffffffffff90811683526020840135906153f0826147b8565b9081166020840152604084013590615407826147b8565b908116604084015260608401359061541e826147b8565b8082166060850152505092915050565b600060ff821660ff84168060ff0382111561544b5761544b614fde565b019392505050565b600060ff83168061548d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8060ff84160491505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600063ffffffff80831681851680830382111561554557615545614fde565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b1660408501525080606084015261557e8184018a614c79565b905082810360808401526155928189614c79565b905060ff871660a084015282810360c08401526155af818761466c565b905067ffffffffffffffff851660e08401528281036101008401526155d4818561466c565b9c9b505050505050505050505050565b602081526000613a416020830184615129565b7fff00000000000000000000000000000000000000000000000000000000000000831681526000825161563181600185016020870161463c565b919091016001019392505050565b60006060828403121561565157600080fd5b6156596146f8565b82518152602083015161566b816147b8565b6020820152604083015161567e816147b8565b60408201529392505050565b600067ffffffffffffffff80831681851680830382111561554557615545614fde565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526156f48285018b614c79565b91508382036080850152615708828a614c79565b915060ff881660a085015283820360c0850152615725828861466c565b90861660e085015283810361010085015290506155d4818561466c565b6000825161575481846020870161463c565b919091019291505056fea164736f6c634300080d000a",
}

var BlobVerifierHelperABI = BlobVerifierHelperMetaData.ABI

var BlobVerifierHelperBin = BlobVerifierHelperMetaData.Bin

func DeployBlobVerifierHelper(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, sourceTokens []common.Address, pools []common.Address, feeds []common.Address, afn common.Address, maxTimeWithoutAFNSignal *big.Int, executionDelaySeconds uint64, maxTokensLength uint64) (common.Address, *types.Transaction, *BlobVerifierHelper, error) {
	parsed, err := BlobVerifierHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlobVerifierHelperBin), backend, sourceChainId, chainId, sourceTokens, pools, feeds, afn, maxTimeWithoutAFNSignal, executionDelaySeconds, maxTokensLength)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlobVerifierHelper{BlobVerifierHelperCaller: BlobVerifierHelperCaller{contract: contract}, BlobVerifierHelperTransactor: BlobVerifierHelperTransactor{contract: contract}, BlobVerifierHelperFilterer: BlobVerifierHelperFilterer{contract: contract}}, nil
}

type BlobVerifierHelper struct {
	address common.Address
	abi     abi.ABI
	BlobVerifierHelperCaller
	BlobVerifierHelperTransactor
	BlobVerifierHelperFilterer
}

type BlobVerifierHelperCaller struct {
	contract *bind.BoundContract
}

type BlobVerifierHelperTransactor struct {
	contract *bind.BoundContract
}

type BlobVerifierHelperFilterer struct {
	contract *bind.BoundContract
}

type BlobVerifierHelperSession struct {
	Contract     *BlobVerifierHelper
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BlobVerifierHelperCallerSession struct {
	Contract *BlobVerifierHelperCaller
	CallOpts bind.CallOpts
}

type BlobVerifierHelperTransactorSession struct {
	Contract     *BlobVerifierHelperTransactor
	TransactOpts bind.TransactOpts
}

type BlobVerifierHelperRaw struct {
	Contract *BlobVerifierHelper
}

type BlobVerifierHelperCallerRaw struct {
	Contract *BlobVerifierHelperCaller
}

type BlobVerifierHelperTransactorRaw struct {
	Contract *BlobVerifierHelperTransactor
}

func NewBlobVerifierHelper(address common.Address, backend bind.ContractBackend) (*BlobVerifierHelper, error) {
	abi, err := abi.JSON(strings.NewReader(BlobVerifierHelperABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBlobVerifierHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelper{address: address, abi: abi, BlobVerifierHelperCaller: BlobVerifierHelperCaller{contract: contract}, BlobVerifierHelperTransactor: BlobVerifierHelperTransactor{contract: contract}, BlobVerifierHelperFilterer: BlobVerifierHelperFilterer{contract: contract}}, nil
}

func NewBlobVerifierHelperCaller(address common.Address, caller bind.ContractCaller) (*BlobVerifierHelperCaller, error) {
	contract, err := bindBlobVerifierHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperCaller{contract: contract}, nil
}

func NewBlobVerifierHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*BlobVerifierHelperTransactor, error) {
	contract, err := bindBlobVerifierHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperTransactor{contract: contract}, nil
}

func NewBlobVerifierHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*BlobVerifierHelperFilterer, error) {
	contract, err := bindBlobVerifierHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperFilterer{contract: contract}, nil
}

func bindBlobVerifierHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlobVerifierHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_BlobVerifierHelper *BlobVerifierHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlobVerifierHelper.Contract.BlobVerifierHelperCaller.contract.Call(opts, result, method, params...)
}

func (_BlobVerifierHelper *BlobVerifierHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.BlobVerifierHelperTransactor.contract.Transfer(opts)
}

func (_BlobVerifierHelper *BlobVerifierHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.BlobVerifierHelperTransactor.contract.Transact(opts, method, params...)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlobVerifierHelper.Contract.contract.Call(opts, result, method, params...)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.contract.Transfer(opts)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.contract.Transact(opts, method, params...)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) CHAINID() (*big.Int, error) {
	return _BlobVerifierHelper.Contract.CHAINID(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) CHAINID() (*big.Int, error) {
	return _BlobVerifierHelper.Contract.CHAINID(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "SOURCE_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) SOURCECHAINID() (*big.Int, error) {
	return _BlobVerifierHelper.Contract.SOURCECHAINID(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) SOURCECHAINID() (*big.Int, error) {
	return _BlobVerifierHelper.Contract.SOURCECHAINID(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) GetAFN() (common.Address, error) {
	return _BlobVerifierHelper.Contract.GetAFN(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) GetAFN() (common.Address, error) {
	return _BlobVerifierHelper.Contract.GetAFN(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) GetExecuted(opts *bind.CallOpts, sequenceNumber uint64) (bool, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "getExecuted", sequenceNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) GetExecuted(sequenceNumber uint64) (bool, error) {
	return _BlobVerifierHelper.Contract.GetExecuted(&_BlobVerifierHelper.CallOpts, sequenceNumber)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) GetExecuted(sequenceNumber uint64) (bool, error) {
	return _BlobVerifierHelper.Contract.GetExecuted(&_BlobVerifierHelper.CallOpts, sequenceNumber)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "getFeed", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) GetFeed(token common.Address) (common.Address, error) {
	return _BlobVerifierHelper.Contract.GetFeed(&_BlobVerifierHelper.CallOpts, token)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) GetFeed(token common.Address) (common.Address, error) {
	return _BlobVerifierHelper.Contract.GetFeed(&_BlobVerifierHelper.CallOpts, token)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "getFeedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) GetFeedTokens() ([]common.Address, error) {
	return _BlobVerifierHelper.Contract.GetFeedTokens(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) GetFeedTokens() ([]common.Address, error) {
	return _BlobVerifierHelper.Contract.GetFeedTokens(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) GetLastReport(opts *bind.CallOpts) (CCIPRelayReport, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "getLastReport")

	if err != nil {
		return *new(CCIPRelayReport), err
	}

	out0 := *abi.ConvertType(out[0], new(CCIPRelayReport)).(*CCIPRelayReport)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) GetLastReport() (CCIPRelayReport, error) {
	return _BlobVerifierHelper.Contract.GetLastReport(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) GetLastReport() (CCIPRelayReport, error) {
	return _BlobVerifierHelper.Contract.GetLastReport(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "getMaxSecondsWithoutAFNHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _BlobVerifierHelper.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _BlobVerifierHelper.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) GetMerkleRoot(opts *bind.CallOpts, root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "getMerkleRoot", root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) GetMerkleRoot(root [32]byte) (*big.Int, error) {
	return _BlobVerifierHelper.Contract.GetMerkleRoot(&_BlobVerifierHelper.CallOpts, root)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) GetMerkleRoot(root [32]byte) (*big.Int, error) {
	return _BlobVerifierHelper.Contract.GetMerkleRoot(&_BlobVerifierHelper.CallOpts, root)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) GetOffRampConfig(opts *bind.CallOpts) (TollOffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "getOffRampConfig")

	if err != nil {
		return *new(TollOffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(TollOffRampInterfaceOffRampConfig)).(*TollOffRampInterfaceOffRampConfig)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) GetOffRampConfig() (TollOffRampInterfaceOffRampConfig, error) {
	return _BlobVerifierHelper.Contract.GetOffRampConfig(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) GetOffRampConfig() (TollOffRampInterfaceOffRampConfig, error) {
	return _BlobVerifierHelper.Contract.GetOffRampConfig(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _BlobVerifierHelper.Contract.GetPool(&_BlobVerifierHelper.CallOpts, sourceToken)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _BlobVerifierHelper.Contract.GetPool(&_BlobVerifierHelper.CallOpts, sourceToken)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) GetPoolTokens() ([]common.Address, error) {
	return _BlobVerifierHelper.Contract.GetPoolTokens(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _BlobVerifierHelper.Contract.GetPoolTokens(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) GetRouter() (common.Address, error) {
	return _BlobVerifierHelper.Contract.GetRouter(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) GetRouter() (common.Address, error) {
	return _BlobVerifierHelper.Contract.GetRouter(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "isHealthy", timeNow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _BlobVerifierHelper.Contract.IsHealthy(&_BlobVerifierHelper.CallOpts, timeNow)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _BlobVerifierHelper.Contract.IsHealthy(&_BlobVerifierHelper.CallOpts, timeNow)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) IsPool(addr common.Address) (bool, error) {
	return _BlobVerifierHelper.Contract.IsPool(&_BlobVerifierHelper.CallOpts, addr)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) IsPool(addr common.Address) (bool, error) {
	return _BlobVerifierHelper.Contract.IsPool(&_BlobVerifierHelper.CallOpts, addr)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _BlobVerifierHelper.Contract.LatestConfigDetails(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _BlobVerifierHelper.Contract.LatestConfigDetails(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _BlobVerifierHelper.Contract.LatestConfigDigestAndEpoch(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _BlobVerifierHelper.Contract.LatestConfigDigestAndEpoch(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) MerkleRoot(opts *bind.CallOpts, report CCIPExecutionReport) ([32]byte, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "merkleRoot", report)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) MerkleRoot(report CCIPExecutionReport) ([32]byte, error) {
	return _BlobVerifierHelper.Contract.MerkleRoot(&_BlobVerifierHelper.CallOpts, report)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) MerkleRoot(report CCIPExecutionReport) ([32]byte, error) {
	return _BlobVerifierHelper.Contract.MerkleRoot(&_BlobVerifierHelper.CallOpts, report)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) Owner() (common.Address, error) {
	return _BlobVerifierHelper.Contract.Owner(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) Owner() (common.Address, error) {
	return _BlobVerifierHelper.Contract.Owner(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) Paused() (bool, error) {
	return _BlobVerifierHelper.Contract.Paused(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) Paused() (bool, error) {
	return _BlobVerifierHelper.Contract.Paused(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) Transmitters() ([]common.Address, error) {
	return _BlobVerifierHelper.Contract.Transmitters(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) Transmitters() ([]common.Address, error) {
	return _BlobVerifierHelper.Contract.Transmitters(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BlobVerifierHelper.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BlobVerifierHelper *BlobVerifierHelperSession) TypeAndVersion() (string, error) {
	return _BlobVerifierHelper.Contract.TypeAndVersion(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperCallerSession) TypeAndVersion() (string, error) {
	return _BlobVerifierHelper.Contract.TypeAndVersion(&_BlobVerifierHelper.CallOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "acceptOwnership")
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) AcceptOwnership() (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.AcceptOwnership(&_BlobVerifierHelper.TransactOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.AcceptOwnership(&_BlobVerifierHelper.TransactOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "addFeed", token, feed)
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.AddFeed(&_BlobVerifierHelper.TransactOpts, token, feed)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.AddFeed(&_BlobVerifierHelper.TransactOpts, token, feed)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "addPool", token, pool)
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.AddPool(&_BlobVerifierHelper.TransactOpts, token, pool)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.AddPool(&_BlobVerifierHelper.TransactOpts, token, pool)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) ExecuteTransaction(opts *bind.TransactOpts, report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "executeTransaction", report, needFee)
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) ExecuteTransaction(report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.ExecuteTransaction(&_BlobVerifierHelper.TransactOpts, report, needFee)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) ExecuteTransaction(report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.ExecuteTransaction(&_BlobVerifierHelper.TransactOpts, report, needFee)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "pause")
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) Pause() (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.Pause(&_BlobVerifierHelper.TransactOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) Pause() (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.Pause(&_BlobVerifierHelper.TransactOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "removeFeed", token, feed)
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.RemoveFeed(&_BlobVerifierHelper.TransactOpts, token, feed)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.RemoveFeed(&_BlobVerifierHelper.TransactOpts, token, feed)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "removePool", token, pool)
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.RemovePool(&_BlobVerifierHelper.TransactOpts, token, pool)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.RemovePool(&_BlobVerifierHelper.TransactOpts, token, pool)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) Report(opts *bind.TransactOpts, merkle []byte) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "report", merkle)
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) Report(merkle []byte) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.Report(&_BlobVerifierHelper.TransactOpts, merkle)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) Report(merkle []byte) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.Report(&_BlobVerifierHelper.TransactOpts, merkle)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "setAFN", afn)
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.SetAFN(&_BlobVerifierHelper.TransactOpts, afn)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.SetAFN(&_BlobVerifierHelper.TransactOpts, afn)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.SetConfig(&_BlobVerifierHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.SetConfig(&_BlobVerifierHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_BlobVerifierHelper.TransactOpts, newTime)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_BlobVerifierHelper.TransactOpts, newTime)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) SetOffRampConfig(opts *bind.TransactOpts, config TollOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "setOffRampConfig", config)
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) SetOffRampConfig(config TollOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.SetOffRampConfig(&_BlobVerifierHelper.TransactOpts, config)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) SetOffRampConfig(config TollOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.SetOffRampConfig(&_BlobVerifierHelper.TransactOpts, config)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "setRouter", router)
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.SetRouter(&_BlobVerifierHelper.TransactOpts, router)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.SetRouter(&_BlobVerifierHelper.TransactOpts, router)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "transferOwnership", to)
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.TransferOwnership(&_BlobVerifierHelper.TransactOpts, to)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.TransferOwnership(&_BlobVerifierHelper.TransactOpts, to)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.Transmit(&_BlobVerifierHelper.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.Transmit(&_BlobVerifierHelper.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "unpause")
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) Unpause() (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.Unpause(&_BlobVerifierHelper.TransactOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) Unpause() (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.Unpause(&_BlobVerifierHelper.TransactOpts)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactor) WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BlobVerifierHelper.contract.Transact(opts, "withdrawAccumulatedFees", feeToken, recipient, amount)
}

func (_BlobVerifierHelper *BlobVerifierHelperSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.WithdrawAccumulatedFees(&_BlobVerifierHelper.TransactOpts, feeToken, recipient, amount)
}

func (_BlobVerifierHelper *BlobVerifierHelperTransactorSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BlobVerifierHelper.Contract.WithdrawAccumulatedFees(&_BlobVerifierHelper.TransactOpts, feeToken, recipient, amount)
}

type BlobVerifierHelperAFNMaxHeartbeatTimeSetIterator struct {
	Event *BlobVerifierHelperAFNMaxHeartbeatTimeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperAFNMaxHeartbeatTimeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperAFNMaxHeartbeatTimeSet)
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
		it.Event = new(BlobVerifierHelperAFNMaxHeartbeatTimeSet)
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

func (it *BlobVerifierHelperAFNMaxHeartbeatTimeSetIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperAFNMaxHeartbeatTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperAFNMaxHeartbeatTimeSet struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*BlobVerifierHelperAFNMaxHeartbeatTimeSetIterator, error) {

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperAFNMaxHeartbeatTimeSetIterator{contract: _BlobVerifierHelper.contract, event: "AFNMaxHeartbeatTimeSet", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperAFNMaxHeartbeatTimeSet) (event.Subscription, error) {

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperAFNMaxHeartbeatTimeSet)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParseAFNMaxHeartbeatTimeSet(log types.Log) (*BlobVerifierHelperAFNMaxHeartbeatTimeSet, error) {
	event := new(BlobVerifierHelperAFNMaxHeartbeatTimeSet)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperAFNSetIterator struct {
	Event *BlobVerifierHelperAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperAFNSet)
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
		it.Event = new(BlobVerifierHelperAFNSet)
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

func (it *BlobVerifierHelperAFNSetIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterAFNSet(opts *bind.FilterOpts) (*BlobVerifierHelperAFNSetIterator, error) {

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperAFNSetIterator{contract: _BlobVerifierHelper.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperAFNSet) (event.Subscription, error) {

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperAFNSet)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParseAFNSet(log types.Log) (*BlobVerifierHelperAFNSet, error) {
	event := new(BlobVerifierHelperAFNSet)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperConfigSetIterator struct {
	Event *BlobVerifierHelperConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperConfigSet)
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
		it.Event = new(BlobVerifierHelperConfigSet)
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

func (it *BlobVerifierHelperConfigSetIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperConfigSet struct {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterConfigSet(opts *bind.FilterOpts) (*BlobVerifierHelperConfigSetIterator, error) {

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperConfigSetIterator{contract: _BlobVerifierHelper.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperConfigSet) (event.Subscription, error) {

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperConfigSet)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParseConfigSet(log types.Log) (*BlobVerifierHelperConfigSet, error) {
	event := new(BlobVerifierHelperConfigSet)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperCrossChainMessageExecutedIterator struct {
	Event *BlobVerifierHelperCrossChainMessageExecuted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperCrossChainMessageExecutedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperCrossChainMessageExecuted)
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
		it.Event = new(BlobVerifierHelperCrossChainMessageExecuted)
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

func (it *BlobVerifierHelperCrossChainMessageExecutedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperCrossChainMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperCrossChainMessageExecuted struct {
	SequenceNumber uint64
	Raw            types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterCrossChainMessageExecuted(opts *bind.FilterOpts, sequenceNumber []uint64) (*BlobVerifierHelperCrossChainMessageExecutedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "CrossChainMessageExecuted", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperCrossChainMessageExecutedIterator{contract: _BlobVerifierHelper.contract, event: "CrossChainMessageExecuted", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchCrossChainMessageExecuted(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperCrossChainMessageExecuted, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "CrossChainMessageExecuted", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperCrossChainMessageExecuted)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "CrossChainMessageExecuted", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParseCrossChainMessageExecuted(log types.Log) (*BlobVerifierHelperCrossChainMessageExecuted, error) {
	event := new(BlobVerifierHelperCrossChainMessageExecuted)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "CrossChainMessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperFeedAddedIterator struct {
	Event *BlobVerifierHelperFeedAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperFeedAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperFeedAdded)
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
		it.Event = new(BlobVerifierHelperFeedAdded)
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

func (it *BlobVerifierHelperFeedAddedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperFeedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperFeedAdded struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterFeedAdded(opts *bind.FilterOpts) (*BlobVerifierHelperFeedAddedIterator, error) {

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperFeedAddedIterator{contract: _BlobVerifierHelper.contract, event: "FeedAdded", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperFeedAdded) (event.Subscription, error) {

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperFeedAdded)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "FeedAdded", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParseFeedAdded(log types.Log) (*BlobVerifierHelperFeedAdded, error) {
	event := new(BlobVerifierHelperFeedAdded)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "FeedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperFeedRemovedIterator struct {
	Event *BlobVerifierHelperFeedRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperFeedRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperFeedRemoved)
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
		it.Event = new(BlobVerifierHelperFeedRemoved)
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

func (it *BlobVerifierHelperFeedRemovedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperFeedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperFeedRemoved struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterFeedRemoved(opts *bind.FilterOpts) (*BlobVerifierHelperFeedRemovedIterator, error) {

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperFeedRemovedIterator{contract: _BlobVerifierHelper.contract, event: "FeedRemoved", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperFeedRemoved) (event.Subscription, error) {

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperFeedRemoved)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParseFeedRemoved(log types.Log) (*BlobVerifierHelperFeedRemoved, error) {
	event := new(BlobVerifierHelperFeedRemoved)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperFeesWithdrawnIterator struct {
	Event *BlobVerifierHelperFeesWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperFeesWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperFeesWithdrawn)
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
		it.Event = new(BlobVerifierHelperFeesWithdrawn)
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

func (it *BlobVerifierHelperFeesWithdrawnIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperFeesWithdrawn struct {
	FeeToken  common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*BlobVerifierHelperFeesWithdrawnIterator, error) {

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperFeesWithdrawnIterator{contract: _BlobVerifierHelper.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperFeesWithdrawn)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParseFeesWithdrawn(log types.Log) (*BlobVerifierHelperFeesWithdrawn, error) {
	event := new(BlobVerifierHelperFeesWithdrawn)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperOffRampConfigSetIterator struct {
	Event *BlobVerifierHelperOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperOffRampConfigSet)
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
		it.Event = new(BlobVerifierHelperOffRampConfigSet)
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

func (it *BlobVerifierHelperOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperOffRampConfigSet struct {
	Config TollOffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*BlobVerifierHelperOffRampConfigSetIterator, error) {

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperOffRampConfigSetIterator{contract: _BlobVerifierHelper.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperOffRampConfigSet)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParseOffRampConfigSet(log types.Log) (*BlobVerifierHelperOffRampConfigSet, error) {
	event := new(BlobVerifierHelperOffRampConfigSet)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperOffRampRouterSetIterator struct {
	Event *BlobVerifierHelperOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperOffRampRouterSet)
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
		it.Event = new(BlobVerifierHelperOffRampRouterSet)
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

func (it *BlobVerifierHelperOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperOffRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts) (*BlobVerifierHelperOffRampRouterSetIterator, error) {

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "OffRampRouterSet")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperOffRampRouterSetIterator{contract: _BlobVerifierHelper.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperOffRampRouterSet) (event.Subscription, error) {

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "OffRampRouterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperOffRampRouterSet)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParseOffRampRouterSet(log types.Log) (*BlobVerifierHelperOffRampRouterSet, error) {
	event := new(BlobVerifierHelperOffRampRouterSet)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperOwnershipTransferRequestedIterator struct {
	Event *BlobVerifierHelperOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperOwnershipTransferRequested)
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
		it.Event = new(BlobVerifierHelperOwnershipTransferRequested)
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

func (it *BlobVerifierHelperOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BlobVerifierHelperOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperOwnershipTransferRequestedIterator{contract: _BlobVerifierHelper.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperOwnershipTransferRequested)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParseOwnershipTransferRequested(log types.Log) (*BlobVerifierHelperOwnershipTransferRequested, error) {
	event := new(BlobVerifierHelperOwnershipTransferRequested)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperOwnershipTransferredIterator struct {
	Event *BlobVerifierHelperOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperOwnershipTransferred)
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
		it.Event = new(BlobVerifierHelperOwnershipTransferred)
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

func (it *BlobVerifierHelperOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BlobVerifierHelperOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperOwnershipTransferredIterator{contract: _BlobVerifierHelper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperOwnershipTransferred)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParseOwnershipTransferred(log types.Log) (*BlobVerifierHelperOwnershipTransferred, error) {
	event := new(BlobVerifierHelperOwnershipTransferred)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperPausedIterator struct {
	Event *BlobVerifierHelperPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperPaused)
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
		it.Event = new(BlobVerifierHelperPaused)
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

func (it *BlobVerifierHelperPausedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterPaused(opts *bind.FilterOpts) (*BlobVerifierHelperPausedIterator, error) {

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperPausedIterator{contract: _BlobVerifierHelper.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperPaused) (event.Subscription, error) {

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperPaused)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParsePaused(log types.Log) (*BlobVerifierHelperPaused, error) {
	event := new(BlobVerifierHelperPaused)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperPoolAddedIterator struct {
	Event *BlobVerifierHelperPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperPoolAdded)
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
		it.Event = new(BlobVerifierHelperPoolAdded)
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

func (it *BlobVerifierHelperPoolAddedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*BlobVerifierHelperPoolAddedIterator, error) {

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperPoolAddedIterator{contract: _BlobVerifierHelper.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperPoolAdded) (event.Subscription, error) {

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperPoolAdded)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParsePoolAdded(log types.Log) (*BlobVerifierHelperPoolAdded, error) {
	event := new(BlobVerifierHelperPoolAdded)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperPoolRemovedIterator struct {
	Event *BlobVerifierHelperPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperPoolRemoved)
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
		it.Event = new(BlobVerifierHelperPoolRemoved)
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

func (it *BlobVerifierHelperPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*BlobVerifierHelperPoolRemovedIterator, error) {

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperPoolRemovedIterator{contract: _BlobVerifierHelper.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperPoolRemoved)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParsePoolRemoved(log types.Log) (*BlobVerifierHelperPoolRemoved, error) {
	event := new(BlobVerifierHelperPoolRemoved)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperReportAcceptedIterator struct {
	Event *BlobVerifierHelperReportAccepted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperReportAcceptedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperReportAccepted)
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
		it.Event = new(BlobVerifierHelperReportAccepted)
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

func (it *BlobVerifierHelperReportAcceptedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperReportAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperReportAccepted struct {
	Report CCIPRelayReport
	Raw    types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterReportAccepted(opts *bind.FilterOpts) (*BlobVerifierHelperReportAcceptedIterator, error) {

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "ReportAccepted")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperReportAcceptedIterator{contract: _BlobVerifierHelper.contract, event: "ReportAccepted", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchReportAccepted(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperReportAccepted) (event.Subscription, error) {

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "ReportAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperReportAccepted)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "ReportAccepted", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParseReportAccepted(log types.Log) (*BlobVerifierHelperReportAccepted, error) {
	event := new(BlobVerifierHelperReportAccepted)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "ReportAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperTransmittedIterator struct {
	Event *BlobVerifierHelperTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperTransmitted)
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
		it.Event = new(BlobVerifierHelperTransmitted)
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

func (it *BlobVerifierHelperTransmittedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterTransmitted(opts *bind.FilterOpts) (*BlobVerifierHelperTransmittedIterator, error) {

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperTransmittedIterator{contract: _BlobVerifierHelper.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperTransmitted) (event.Subscription, error) {

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperTransmitted)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParseTransmitted(log types.Log) (*BlobVerifierHelperTransmitted, error) {
	event := new(BlobVerifierHelperTransmitted)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BlobVerifierHelperUnpausedIterator struct {
	Event *BlobVerifierHelperUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BlobVerifierHelperUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobVerifierHelperUnpaused)
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
		it.Event = new(BlobVerifierHelperUnpaused)
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

func (it *BlobVerifierHelperUnpausedIterator) Error() error {
	return it.fail
}

func (it *BlobVerifierHelperUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BlobVerifierHelperUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BlobVerifierHelperUnpausedIterator, error) {

	logs, sub, err := _BlobVerifierHelper.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BlobVerifierHelperUnpausedIterator{contract: _BlobVerifierHelper.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperUnpaused) (event.Subscription, error) {

	logs, sub, err := _BlobVerifierHelper.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BlobVerifierHelperUnpaused)
				if err := _BlobVerifierHelper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelperFilterer) ParseUnpaused(log types.Log) (*BlobVerifierHelperUnpaused, error) {
	event := new(BlobVerifierHelperUnpaused)
	if err := _BlobVerifierHelper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_BlobVerifierHelper *BlobVerifierHelper) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BlobVerifierHelper.abi.Events["AFNMaxHeartbeatTimeSet"].ID:
		return _BlobVerifierHelper.ParseAFNMaxHeartbeatTimeSet(log)
	case _BlobVerifierHelper.abi.Events["AFNSet"].ID:
		return _BlobVerifierHelper.ParseAFNSet(log)
	case _BlobVerifierHelper.abi.Events["ConfigSet"].ID:
		return _BlobVerifierHelper.ParseConfigSet(log)
	case _BlobVerifierHelper.abi.Events["CrossChainMessageExecuted"].ID:
		return _BlobVerifierHelper.ParseCrossChainMessageExecuted(log)
	case _BlobVerifierHelper.abi.Events["FeedAdded"].ID:
		return _BlobVerifierHelper.ParseFeedAdded(log)
	case _BlobVerifierHelper.abi.Events["FeedRemoved"].ID:
		return _BlobVerifierHelper.ParseFeedRemoved(log)
	case _BlobVerifierHelper.abi.Events["FeesWithdrawn"].ID:
		return _BlobVerifierHelper.ParseFeesWithdrawn(log)
	case _BlobVerifierHelper.abi.Events["OffRampConfigSet"].ID:
		return _BlobVerifierHelper.ParseOffRampConfigSet(log)
	case _BlobVerifierHelper.abi.Events["OffRampRouterSet"].ID:
		return _BlobVerifierHelper.ParseOffRampRouterSet(log)
	case _BlobVerifierHelper.abi.Events["OwnershipTransferRequested"].ID:
		return _BlobVerifierHelper.ParseOwnershipTransferRequested(log)
	case _BlobVerifierHelper.abi.Events["OwnershipTransferred"].ID:
		return _BlobVerifierHelper.ParseOwnershipTransferred(log)
	case _BlobVerifierHelper.abi.Events["Paused"].ID:
		return _BlobVerifierHelper.ParsePaused(log)
	case _BlobVerifierHelper.abi.Events["PoolAdded"].ID:
		return _BlobVerifierHelper.ParsePoolAdded(log)
	case _BlobVerifierHelper.abi.Events["PoolRemoved"].ID:
		return _BlobVerifierHelper.ParsePoolRemoved(log)
	case _BlobVerifierHelper.abi.Events["ReportAccepted"].ID:
		return _BlobVerifierHelper.ParseReportAccepted(log)
	case _BlobVerifierHelper.abi.Events["Transmitted"].ID:
		return _BlobVerifierHelper.ParseTransmitted(log)
	case _BlobVerifierHelper.abi.Events["Unpaused"].ID:
		return _BlobVerifierHelper.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (BlobVerifierHelperAFNMaxHeartbeatTimeSet) Topic() common.Hash {
	return common.HexToHash("0x72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251")
}

func (BlobVerifierHelperAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (BlobVerifierHelperConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (BlobVerifierHelperCrossChainMessageExecuted) Topic() common.Hash {
	return common.HexToHash("0x88a2c71ae86800edd209d83430f348ce70539bd40085d045340ff569f058a2a0")
}

func (BlobVerifierHelperFeedAdded) Topic() common.Hash {
	return common.HexToHash("0x037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb601305")
}

func (BlobVerifierHelperFeedRemoved) Topic() common.Hash {
	return common.HexToHash("0xa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d")
}

func (BlobVerifierHelperFeesWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8")
}

func (BlobVerifierHelperOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xe4cd88b1b5e20a0b843af3207ba74d3f84af4acff45830469490ac9c6ab8a745")
}

func (BlobVerifierHelperOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4")
}

func (BlobVerifierHelperOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BlobVerifierHelperOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BlobVerifierHelperPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (BlobVerifierHelperPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (BlobVerifierHelperPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (BlobVerifierHelperReportAccepted) Topic() common.Hash {
	return common.HexToHash("0x6a2424b54e9b18f0c3d440b6003882f5ffa34272205608ad9451c921c0d99a35")
}

func (BlobVerifierHelperTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (BlobVerifierHelperUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_BlobVerifierHelper *BlobVerifierHelper) Address() common.Address {
	return _BlobVerifierHelper.address
}

type BlobVerifierHelperInterface interface {
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

	Report(opts *bind.TransactOpts, merkle []byte) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error)

	SetOffRampConfig(opts *bind.TransactOpts, config TollOffRampInterfaceOffRampConfig) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*BlobVerifierHelperAFNMaxHeartbeatTimeSetIterator, error)

	WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperAFNMaxHeartbeatTimeSet) (event.Subscription, error)

	ParseAFNMaxHeartbeatTimeSet(log types.Log) (*BlobVerifierHelperAFNMaxHeartbeatTimeSet, error)

	FilterAFNSet(opts *bind.FilterOpts) (*BlobVerifierHelperAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*BlobVerifierHelperAFNSet, error)

	FilterConfigSet(opts *bind.FilterOpts) (*BlobVerifierHelperConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*BlobVerifierHelperConfigSet, error)

	FilterCrossChainMessageExecuted(opts *bind.FilterOpts, sequenceNumber []uint64) (*BlobVerifierHelperCrossChainMessageExecutedIterator, error)

	WatchCrossChainMessageExecuted(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperCrossChainMessageExecuted, sequenceNumber []uint64) (event.Subscription, error)

	ParseCrossChainMessageExecuted(log types.Log) (*BlobVerifierHelperCrossChainMessageExecuted, error)

	FilterFeedAdded(opts *bind.FilterOpts) (*BlobVerifierHelperFeedAddedIterator, error)

	WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperFeedAdded) (event.Subscription, error)

	ParseFeedAdded(log types.Log) (*BlobVerifierHelperFeedAdded, error)

	FilterFeedRemoved(opts *bind.FilterOpts) (*BlobVerifierHelperFeedRemovedIterator, error)

	WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperFeedRemoved) (event.Subscription, error)

	ParseFeedRemoved(log types.Log) (*BlobVerifierHelperFeedRemoved, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*BlobVerifierHelperFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*BlobVerifierHelperFeesWithdrawn, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*BlobVerifierHelperOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*BlobVerifierHelperOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts) (*BlobVerifierHelperOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperOffRampRouterSet) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*BlobVerifierHelperOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BlobVerifierHelperOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BlobVerifierHelperOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BlobVerifierHelperOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BlobVerifierHelperOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*BlobVerifierHelperPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*BlobVerifierHelperPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*BlobVerifierHelperPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*BlobVerifierHelperPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*BlobVerifierHelperPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*BlobVerifierHelperPoolRemoved, error)

	FilterReportAccepted(opts *bind.FilterOpts) (*BlobVerifierHelperReportAcceptedIterator, error)

	WatchReportAccepted(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperReportAccepted) (event.Subscription, error)

	ParseReportAccepted(log types.Log) (*BlobVerifierHelperReportAccepted, error)

	FilterTransmitted(opts *bind.FilterOpts) (*BlobVerifierHelperTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*BlobVerifierHelperTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*BlobVerifierHelperUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BlobVerifierHelperUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*BlobVerifierHelperUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
