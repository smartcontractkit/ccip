// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package onramp

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

type CCIPEVMToAnyTollMessage struct {
	Receiver       common.Address
	Data           []byte
	Tokens         []common.Address
	Amounts        []*big.Int
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	GasLimit       *big.Int
}

type CCIPEVMToEVMTollEvent struct {
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

type OnRampInterfaceOnRampConfig struct {
	Router           common.Address
	RelayingFeeJuels uint64
	MaxDataSize      uint64
	MaxTokensLength  uint64
}

var OnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"contractAggregatorV2V3Interface[]\",\"name\":\"feeds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceFeedConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenFeedMistmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMistmatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AllowlistEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"AllowlistSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCCIP.EVMToEVMTollEvent\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"full\",\"type\":\"bool\"}],\"name\":\"NewTokenBucketConstructed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OnRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DESTINATION_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"addFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMToAnyTollMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"getRequiredFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"removeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"setAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setAllowlistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162003f8338038062003f838339810160408190526200003491620007ed565b6000805460ff191681558790869082908990879087903390819081620000a15760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000db57620000db816200048f565b5050506001600160a01b0382161580620000f3575080155b156200011257604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001585760405162d8548360e71b815260040160405180910390fd5b81516200016d90600590602085019062000540565b5060005b825181101562000251576000828281518110620001925762000192620008eb565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001dc57620001dc620008eb565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff1916600117905580620002488162000901565b91505062000171565b5050508051825114620002775760405163ee9d106b60e01b815260040160405180910390fd5b81516200028c90600890602085019062000540565b5060005b825181101562000359576000828281518110620002b157620002b1620008eb565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060076000868581518110620002fb57620002fb620008eb565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b02929091169190911790555080620003508162000901565b91505062000290565b505050608089905260a0889052600980546001600160401b0319166001179055835115620003a557600c805460ff191660011790558351620003a390600e90602087019062000540565b505b60005b845181101562000414576001600d6000878481518110620003cd57620003cd620008eb565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055806200040b8162000901565b915050620003a8565b508051600a805460208401516001600160a01b039093166001600160e01b031990911617600160a01b6001600160401b03938416021790556040820151600b80546060909401519183166001600160801b03199094169390931768010000000000000000919092160217905550620009299650505050505050565b336001600160a01b03821603620004e95760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000098565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b82805482825590600052602060002090810192821562000598579160200282015b828111156200059857825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000561565b50620005a6929150620005aa565b5090565b5b80821115620005a65760008155600101620005ab565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620006025762000602620005c1565b604052919050565b60006001600160401b03821115620006265762000626620005c1565b5060051b60200190565b6001600160a01b03811681146200064657600080fd5b50565b600082601f8301126200065b57600080fd5b81516020620006746200066e836200060a565b620005d7565b82815260059290921b840181019181810190868411156200069457600080fd5b8286015b84811015620006bc578051620006ae8162000630565b835291830191830162000698565b509695505050505050565b600082601f830112620006d957600080fd5b81516020620006ec6200066e836200060a565b82815260059290921b840181019181810190868411156200070c57600080fd5b8286015b84811015620006bc578051620007268162000630565b835291830191830162000710565b8051620007418162000630565b919050565b80516001600160401b03811681146200074157600080fd5b6000608082840312156200077157600080fd5b604051608081016001600160401b0381118282101715620007965762000796620005c1565b80604052508091508251620007ab8162000630565b8152620007bb6020840162000746565b6020820152620007ce6040840162000746565b6040820152620007e16060840162000746565b60608201525092915050565b60008060008060008060008060006101808a8c0312156200080d57600080fd5b895160208b015160408c0151919a5098506001600160401b03808211156200083457600080fd5b620008428d838e0162000649565b985060608c01519150808211156200085957600080fd5b620008678d838e0162000649565b975060808c01519150808211156200087e57600080fd5b6200088c8d838e01620006c7565b965060a08c0151915080821115620008a357600080fd5b50620008b28c828d01620006c7565b945050620008c360c08b0162000734565b925060e08a01519150620008dc8b6101008c016200075e565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b6000600182016200092257634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05161362d62000956600039600061036a0152600081816104910152610c78015261362d6000f3fe608060405234801561001057600080fd5b50600436106102265760003560e01c8063744b92e21161012a578063b6608c3b116100bd578063d0d5de611161008c578063eb511dd411610071578063eb511dd414610664578063f2fde38b14610677578063f78faa321461068a57600080fd5b8063d0d5de611461063e578063d7644ba21461065157600080fd5b8063b6608c3b146104e6578063bbe4f6db146104f9578063c3f909d414610532578063c5eff3d01461062957600080fd5b806385e1f4d0116100f957806385e1f4d01461048c57806389c06568146104b35780638da5cb5b146104bb578063b034909c146104de57600080fd5b8063744b92e21461045457806379ba50971461046757806381be8fa41461046f5780638456cb591461048457600080fd5b80632ea02369116101bd578063567c814b1161018c57806359e96b5b1161017157806359e96b5b146103fd5780635b16ebb7146104105780635c975abb1461044957600080fd5b8063567c814b146103c75780635853c627146103ea57600080fd5b80632ea02369146103655780633f4ba83a1461039a57806342af35fd146103a2578063552b818b146103b457600080fd5b8063181f5a77116101f9578063181f5a77146102e25780632222dd42146103215780632b898c251461033f5780632df836c01461035257600080fd5b806304c2a34a1461022b57806305afe24a14610268578063108ee5fc1461029457806316b8e731146102a9575b600080fd5b61023e610239366004612afc565b610695565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61027b610276366004612d41565b6106c6565b60405167ffffffffffffffff909116815260200161025f565b6102a76102a2366004612afc565b610dc6565b005b61023e6102b7366004612afc565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b604080518082018252600c81527f4f6e52616d7020302e302e3100000000000000000000000000000000000000006020820152905161025f9190612eb9565b60025473ffffffffffffffffffffffffffffffffffffffff1661023e565b6102a761034d366004612ecc565b610ea2565b6102a7610360366004612f05565b611272565b61038c7f000000000000000000000000000000000000000000000000000000000000000081565b60405190815260200161025f565b6102a76112c4565b60095467ffffffffffffffff1661027b565b6102a76103c2366004612f1d565b6112d6565b6103da6103d5366004612f92565b6114c7565b604051901515815260200161025f565b6102a76103f8366004612ecc565b61160d565b6102a761040b366004612fab565b61181c565b6103da61041e366004612afc565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff166103da565b6102a7610462366004612ecc565b61189a565b6102a7611c8f565b610477611db1565b60405161025f919061303d565b6102a7611e20565b61038c7f000000000000000000000000000000000000000000000000000000000000000081565b610477611e30565b600054610100900473ffffffffffffffffffffffffffffffffffffffff1661023e565b60035461038c565b6102a76104f4366004612f92565b611e9d565b61023e610507366004612afc565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6105ce6040805160808101825260008082526020820181905291810182905260608101919091525060408051608081018252600a5473ffffffffffffffffffffffffffffffffffffffff8116825267ffffffffffffffff7401000000000000000000000000000000000000000090910481166020830152600b548082169383019390935268010000000000000000909204909116606082015290565b60405161025f9190815173ffffffffffffffffffffffffffffffffffffffff16815260208083015167ffffffffffffffff90811691830191909152604080840151821690830152606092830151169181019190915260800190565b610631611f1d565b60405161025f9190613050565b61038c61064c366004612afc565b611f8a565b6102a761065f3660046130b8565b6120c8565b6102a7610672366004612ecc565b61212f565b6102a7610685366004612afc565b61236f565b600c5460ff166103da565b73ffffffffffffffffffffffffffffffffffffffff8082166000908152600460205260408120549091165b92915050565b6000805460ff1615610739576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ca91906130d5565b15610800576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa158015610870573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089491906130f2565b90506003548160200151426108a9919061317d565b11156108e1576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a54339073ffffffffffffffffffffffffffffffffffffffff168114610934576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8416610981576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b5460208601515167ffffffffffffffff90911610156109e957600b546020860151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301526024820152604401610730565b600b546040860151516801000000000000000090910467ffffffffffffffff161080610a1f575084606001515185604001515114155b15610a56576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c5460ff168015610a8e575073ffffffffffffffffffffffffffffffffffffffff84166000908152600d602052604090205460ff16155b15610add576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610730565b60005b856040015151811015610c4c57600086604001518281518110610b0557610b05613194565b602002602001015190506000610b408273ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff8116610ba7576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610730565b600088606001518481518110610bbf57610bbf613194565b602002602001015190508173ffffffffffffffffffffffffffffffffffffffff1663503c2858826040518263ffffffff1660e01b8152600401610c0491815260200190565b600060405180830381600087803b158015610c1e57600080fd5b505af1158015610c32573d6000803e3d6000fd5b505050505050508080610c44906131c3565b915050610ae0565b506000600960009054906101000a900467ffffffffffffffff16905060006040518061014001604052807f000000000000000000000000000000000000000000000000000000000000000081526020018367ffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff168152602001886000015173ffffffffffffffffffffffffffffffffffffffff168152602001886020015181526020018860400151815260200188606001518152602001886080015173ffffffffffffffffffffffffffffffffffffffff1681526020018860a0015181526020018860c001518152509050816001610d4791906131fb565b600980547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff929092169190911790556040517fab2ca9da6d303be28d1a5e854e3e170be286e07696245e77f8ea11f55367d48190610db0908390613257565b60405180910390a1602001519695505050505050565b610dce612383565b73ffffffffffffffffffffffffffffffffffffffff8116610e1b576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b610eaa612383565b6008546000819003610ee8576040517f2e70248b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610f83576040517f3917193900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610fec576040517f6c17b98700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006008610ffb60018561317d565b8154811061100b5761100b613194565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600882602001516bffffffffffffffffffffffff168154811061105d5761105d613194565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16600861108c60018661317d565b8154811061109c5761109c613194565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600883602001516bffffffffffffffffffffffff168154811061110a5761110a613194565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526007909152604090912080546bffffffffffffffffffffffff90921674010000000000000000000000000000000000000000029190921617905560088054806111ac576111ac613383565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600783526040808520949094558351908152908716918101919091527fa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d91015b60405180910390a15050505050565b61127a612383565b80600a61128782826133c8565b9050507feac62265bdcb30e1e7a4822fecd5035bf208f242c899453ca9a3cdb5eb44225b816040516112b991906134e2565b60405180910390a150565b6112cc612383565b6112d4612409565b565b6112de612383565b6000600e80548060200260200160405190810160405280929190818152602001828054801561134357602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611318575b5050505050905060005b81518110156113dd576000600d600084848151811061136e5761136e613194565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055806113d5816131c3565b91505061134d565b506113ea600e8484612a2d565b5060005b82811015611488576001600d600086868581811061140e5761140e613194565b90506020020160208101906114239190612afc565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905580611480816131c3565b9150506113ee565b507f27f242de1bc4ed72c4329591ffff7d223b5f025e3514a07e05afec6d4eb889cf83836040516114ba92919061355e565b60405180910390a1505050565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa158015611537573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061155b91906130d5565b1580156106c05750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa1580156115d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115f791906130f2565b60200151611605908461317d565b111592915050565b611615612383565b73ffffffffffffffffffffffffffffffffffffffff8216158061164c575073ffffffffffffffffffffffffffffffffffffffff8116155b15611683576040517fee9d106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152901561171f576040517f965ffb7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600880546bffffffffffffffffffffffff908116602080870191825288861660008181526007835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821790558351908152908101919091527f037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb60130591016114ba565b611824612383565b61184573ffffffffffffffffffffffffffffffffffffffff841683836124ea565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8906060016114ba565b6118a2612383565b60055460008190036118e0576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529061197b576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff16146119e4576040517fd428911900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060056119f360018561317d565b81548110611a0357611a03613194565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff1681548110611a5557611a55613194565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166005611a8460018661317d565b81548110611a9457611a94613194565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110611b0257611b02613194565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480611ba457611ba4613383565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101611263565b60015473ffffffffffffffffffffffffffffffffffffffff163314611d10576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610730565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60606008805480602002602001604051908101604052809291908181526020018280548015611e1657602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611deb575b5050505050905090565b611e28612383565b6112d461257c565b60606005805480602002602001604051908101604052809291908181526020018280548015611e165760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611deb575050505050905090565b611ea5612383565b80600003611edf576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c2519101610e96565b6060600e805480602002602001604051908101604052809291908181526020018280548015611e165760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611deb575050505050905090565b600080611fbc8373ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff8116612023576040517feef7849700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610730565b8073ffffffffffffffffffffffffffffffffffffffff166350d25bcd6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561206e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061209291906135ae565b600a546120c1919074010000000000000000000000000000000000000000900467ffffffffffffffff166135c7565b9392505050565b6120d0612383565b600c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168215159081179091556040519081527fa1bf86c493917580dec207969ef59976f0c378f10ece581237f19acfbd858f1c906020016112b9565b612137612383565b73ffffffffffffffffffffffffffffffffffffffff8216158061216e575073ffffffffffffffffffffffffffffffffffffffff8116155b156121a5576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612241576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c91016114ba565b612377612383565b6123808161263c565b50565b600054610100900473ffffffffffffffffffffffffffffffffffffffff1633146112d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610730565b60005460ff16612475576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610730565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052612577908490612737565b505050565b60005460ff16156125e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610730565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586124c03390565b3373ffffffffffffffffffffffffffffffffffffffff8216036126bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610730565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000612799826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166128439092919063ffffffff16565b80519091501561257757808060200190518101906127b791906130d5565b612577576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610730565b6060612852848460008561285a565b949350505050565b6060824710156128ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610730565b843b612954576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610730565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161297d9190613604565b60006040518083038185875af1925050503d80600081146129ba576040519150601f19603f3d011682016040523d82523d6000602084013e6129bf565b606091505b50915091506129cf8282866129da565b979650505050505050565b606083156129e95750816120c1565b8251156129f95782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107309190612eb9565b828054828255906000526020600020908101928215612aa5579160200282015b82811115612aa55781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff843516178255602090920191600190910190612a4d565b50612ab1929150612ab5565b5090565b5b80821115612ab15760008155600101612ab6565b73ffffffffffffffffffffffffffffffffffffffff8116811461238057600080fd5b8035612af781612aca565b919050565b600060208284031215612b0e57600080fd5b81356120c181612aca565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715612b6b57612b6b612b19565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612bb857612bb8612b19565b604052919050565b600082601f830112612bd157600080fd5b813567ffffffffffffffff811115612beb57612beb612b19565b612c1c60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601612b71565b818152846020838601011115612c3157600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115612c6857612c68612b19565b5060051b60200190565b600082601f830112612c8357600080fd5b81356020612c98612c9383612c4e565b612b71565b82815260059290921b84018101918181019086841115612cb757600080fd5b8286015b84811015612cdb578035612cce81612aca565b8352918301918301612cbb565b509695505050505050565b600082601f830112612cf757600080fd5b81356020612d07612c9383612c4e565b82815260059290921b84018101918181019086841115612d2657600080fd5b8286015b84811015612cdb5780358352918301918301612d2a565b60008060408385031215612d5457600080fd5b823567ffffffffffffffff80821115612d6c57600080fd5b9084019060e08287031215612d8057600080fd5b612d88612b48565b612d9183612aec565b8152602083013582811115612da557600080fd5b612db188828601612bc0565b602083015250604083013582811115612dc957600080fd5b612dd588828601612c72565b604083015250606083013582811115612ded57600080fd5b612df988828601612ce6565b606083015250612e0b60808401612aec565b608082015260a083013560a082015260c083013560c0820152809450505050612e3660208401612aec565b90509250929050565b60005b83811015612e5a578181015183820152602001612e42565b83811115612e69576000848401525b50505050565b60008151808452612e87816020860160208601612e3f565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006120c16020830184612e6f565b60008060408385031215612edf57600080fd5b8235612eea81612aca565b91506020830135612efa81612aca565b809150509250929050565b600060808284031215612f1757600080fd5b50919050565b60008060208385031215612f3057600080fd5b823567ffffffffffffffff80821115612f4857600080fd5b818501915085601f830112612f5c57600080fd5b813581811115612f6b57600080fd5b8660208260051b8501011115612f8057600080fd5b60209290920196919550909350505050565b600060208284031215612fa457600080fd5b5035919050565b600080600060608486031215612fc057600080fd5b8335612fcb81612aca565b92506020840135612fdb81612aca565b929592945050506040919091013590565b600081518084526020808501945080840160005b8381101561303257815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101613000565b509495945050505050565b6020815260006120c16020830184612fec565b6020808252825182820181905260009190848201906040850190845b8181101561309e57835173ffffffffffffffffffffffffffffffffffffffff168352928401929184019160010161306c565b50909695505050505050565b801515811461238057600080fd5b6000602082840312156130ca57600080fd5b81356120c1816130aa565b6000602082840312156130e757600080fd5b81516120c1816130aa565b60006060828403121561310457600080fd5b6040516060810181811067ffffffffffffffff8211171561312757613127612b19565b80604052508251815260208301516020820152604083015160408201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561318f5761318f61314e565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036131f4576131f461314e565b5060010190565b600067ffffffffffffffff80831681851680830382111561321e5761321e61314e565b01949350505050565b600081518084526020808501945080840160005b838110156130325781518752958201959082019060010161323b565b60208152815160208201526000602083015161327f604084018267ffffffffffffffff169052565b50604083015173ffffffffffffffffffffffffffffffffffffffff8116606084015250606083015173ffffffffffffffffffffffffffffffffffffffff811660808401525060808301516101408060a08501526132e0610160850183612e6f565b915060a08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808685030160c087015261331c8483612fec565b935060c08701519150808685030160e08701525061333a8382613227565b92505060e08501516101006133668187018373ffffffffffffffffffffffffffffffffffffffff169052565b860151610120868101919091529095015193019290925250919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b67ffffffffffffffff8116811461238057600080fd5b81356133d381612aca565b73ffffffffffffffffffffffffffffffffffffffff811690508154817fffffffffffffffffffffffff000000000000000000000000000000000000000082161783556020840135613423816133b2565b7bffffffffffffffff00000000000000000000000000000000000000008160a01b16837fffffffff00000000000000000000000000000000000000000000000000000000841617178455505050600181016040830135613482816133b2565b81546060850135613492816133b2565b6fffffffffffffffff00000000000000008160401b1667ffffffffffffffff84167fffffffffffffffffffffffffffffffff00000000000000000000000000000000841617178455505050505050565b6080810182356134f181612aca565b73ffffffffffffffffffffffffffffffffffffffff1682526020830135613517816133b2565b67ffffffffffffffff9081166020840152604084013590613537826133b2565b908116604084015260608401359061354e826133b2565b8082166060850152505092915050565b60208082528181018390526000908460408401835b86811015612cdb57823561358681612aca565b73ffffffffffffffffffffffffffffffffffffffff1682529183019190830190600101613573565b6000602082840312156135c057600080fd5b5051919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156135ff576135ff61314e565b500290565b60008251613616818460208701612e3f565b919091019291505056fea164736f6c634300080d000a",
}

var OnRampABI = OnRampMetaData.ABI

var OnRampBin = OnRampMetaData.Bin

func DeployOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, chainId *big.Int, destinationChainId *big.Int, tokens []common.Address, pools []common.Address, feeds []common.Address, allowlist []common.Address, afn common.Address, maxTimeWithoutAFNSignal *big.Int, config OnRampInterfaceOnRampConfig) (common.Address, *types.Transaction, *OnRamp, error) {
	parsed, err := OnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnRampBin), backend, chainId, destinationChainId, tokens, pools, feeds, allowlist, afn, maxTimeWithoutAFNSignal, config)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnRamp{OnRampCaller: OnRampCaller{contract: contract}, OnRampTransactor: OnRampTransactor{contract: contract}, OnRampFilterer: OnRampFilterer{contract: contract}}, nil
}

type OnRamp struct {
	address common.Address
	abi     abi.ABI
	OnRampCaller
	OnRampTransactor
	OnRampFilterer
}

type OnRampCaller struct {
	contract *bind.BoundContract
}

type OnRampTransactor struct {
	contract *bind.BoundContract
}

type OnRampFilterer struct {
	contract *bind.BoundContract
}

type OnRampSession struct {
	Contract     *OnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OnRampCallerSession struct {
	Contract *OnRampCaller
	CallOpts bind.CallOpts
}

type OnRampTransactorSession struct {
	Contract     *OnRampTransactor
	TransactOpts bind.TransactOpts
}

type OnRampRaw struct {
	Contract *OnRamp
}

type OnRampCallerRaw struct {
	Contract *OnRampCaller
}

type OnRampTransactorRaw struct {
	Contract *OnRampTransactor
}

func NewOnRamp(address common.Address, backend bind.ContractBackend) (*OnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(OnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnRamp{address: address, abi: abi, OnRampCaller: OnRampCaller{contract: contract}, OnRampTransactor: OnRampTransactor{contract: contract}, OnRampFilterer: OnRampFilterer{contract: contract}}, nil
}

func NewOnRampCaller(address common.Address, caller bind.ContractCaller) (*OnRampCaller, error) {
	contract, err := bindOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampCaller{contract: contract}, nil
}

func NewOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*OnRampTransactor, error) {
	contract, err := bindOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampTransactor{contract: contract}, nil
}

func NewOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*OnRampFilterer, error) {
	contract, err := bindOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnRampFilterer{contract: contract}, nil
}

func bindOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OnRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_OnRamp *OnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRamp.Contract.OnRampCaller.contract.Call(opts, result, method, params...)
}

func (_OnRamp *OnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.Contract.OnRampTransactor.contract.Transfer(opts)
}

func (_OnRamp *OnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRamp.Contract.OnRampTransactor.contract.Transact(opts, method, params...)
}

func (_OnRamp *OnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_OnRamp *OnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.Contract.contract.Transfer(opts)
}

func (_OnRamp *OnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_OnRamp *OnRampCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OnRamp *OnRampSession) CHAINID() (*big.Int, error) {
	return _OnRamp.Contract.CHAINID(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) CHAINID() (*big.Int, error) {
	return _OnRamp.Contract.CHAINID(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "DESTINATION_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OnRamp *OnRampSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _OnRamp.Contract.DESTINATIONCHAINID(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _OnRamp.Contract.DESTINATIONCHAINID(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetAFN() (common.Address, error) {
	return _OnRamp.Contract.GetAFN(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetAFN() (common.Address, error) {
	return _OnRamp.Contract.GetAFN(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetAllowlist(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getAllowlist")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetAllowlist() ([]common.Address, error) {
	return _OnRamp.Contract.GetAllowlist(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetAllowlist() ([]common.Address, error) {
	return _OnRamp.Contract.GetAllowlist(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetAllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getAllowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OnRamp *OnRampSession) GetAllowlistEnabled() (bool, error) {
	return _OnRamp.Contract.GetAllowlistEnabled(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetAllowlistEnabled() (bool, error) {
	return _OnRamp.Contract.GetAllowlistEnabled(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetConfig(opts *bind.CallOpts) (OnRampInterfaceOnRampConfig, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(OnRampInterfaceOnRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OnRampInterfaceOnRampConfig)).(*OnRampInterfaceOnRampConfig)

	return out0, err

}

func (_OnRamp *OnRampSession) GetConfig() (OnRampInterfaceOnRampConfig, error) {
	return _OnRamp.Contract.GetConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetConfig() (OnRampInterfaceOnRampConfig, error) {
	return _OnRamp.Contract.GetConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getFeed", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetFeed(token common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetFeed(&_OnRamp.CallOpts, token)
}

func (_OnRamp *OnRampCallerSession) GetFeed(token common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetFeed(&_OnRamp.CallOpts, token)
}

func (_OnRamp *OnRampCaller) GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getFeedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetFeedTokens() ([]common.Address, error) {
	return _OnRamp.Contract.GetFeedTokens(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetFeedTokens() ([]common.Address, error) {
	return _OnRamp.Contract.GetFeedTokens(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getMaxSecondsWithoutAFNHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OnRamp *OnRampSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _OnRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _OnRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetPool(&_OnRamp.CallOpts, sourceToken)
}

func (_OnRamp *OnRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetPool(&_OnRamp.CallOpts, sourceToken)
}

func (_OnRamp *OnRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetPoolTokens() ([]common.Address, error) {
	return _OnRamp.Contract.GetPoolTokens(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _OnRamp.Contract.GetPoolTokens(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetRequiredFee(opts *bind.CallOpts, feeToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getRequiredFee", feeToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OnRamp *OnRampSession) GetRequiredFee(feeToken common.Address) (*big.Int, error) {
	return _OnRamp.Contract.GetRequiredFee(&_OnRamp.CallOpts, feeToken)
}

func (_OnRamp *OnRampCallerSession) GetRequiredFee(feeToken common.Address) (*big.Int, error) {
	return _OnRamp.Contract.GetRequiredFee(&_OnRamp.CallOpts, feeToken)
}

func (_OnRamp *OnRampCaller) GetSequenceNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getSequenceNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_OnRamp *OnRampSession) GetSequenceNumber() (uint64, error) {
	return _OnRamp.Contract.GetSequenceNumber(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetSequenceNumber() (uint64, error) {
	return _OnRamp.Contract.GetSequenceNumber(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetTokenPool(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getTokenPool", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetTokenPool(token common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetTokenPool(&_OnRamp.CallOpts, token)
}

func (_OnRamp *OnRampCallerSession) GetTokenPool(token common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetTokenPool(&_OnRamp.CallOpts, token)
}

func (_OnRamp *OnRampCaller) IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "isHealthy", timeNow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OnRamp *OnRampSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _OnRamp.Contract.IsHealthy(&_OnRamp.CallOpts, timeNow)
}

func (_OnRamp *OnRampCallerSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _OnRamp.Contract.IsHealthy(&_OnRamp.CallOpts, timeNow)
}

func (_OnRamp *OnRampCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OnRamp *OnRampSession) IsPool(addr common.Address) (bool, error) {
	return _OnRamp.Contract.IsPool(&_OnRamp.CallOpts, addr)
}

func (_OnRamp *OnRampCallerSession) IsPool(addr common.Address) (bool, error) {
	return _OnRamp.Contract.IsPool(&_OnRamp.CallOpts, addr)
}

func (_OnRamp *OnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) Owner() (common.Address, error) {
	return _OnRamp.Contract.Owner(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) Owner() (common.Address, error) {
	return _OnRamp.Contract.Owner(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OnRamp *OnRampSession) Paused() (bool, error) {
	return _OnRamp.Contract.Paused(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) Paused() (bool, error) {
	return _OnRamp.Contract.Paused(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OnRamp *OnRampSession) TypeAndVersion() (string, error) {
	return _OnRamp.Contract.TypeAndVersion(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) TypeAndVersion() (string, error) {
	return _OnRamp.Contract.TypeAndVersion(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_OnRamp *OnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRamp.Contract.AcceptOwnership(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRamp.Contract.AcceptOwnership(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactor) AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "addFeed", token, feed)
}

func (_OnRamp *OnRampSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.AddFeed(&_OnRamp.TransactOpts, token, feed)
}

func (_OnRamp *OnRampTransactorSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.AddFeed(&_OnRamp.TransactOpts, token, feed)
}

func (_OnRamp *OnRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_OnRamp *OnRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.AddPool(&_OnRamp.TransactOpts, token, pool)
}

func (_OnRamp *OnRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.AddPool(&_OnRamp.TransactOpts, token, pool)
}

func (_OnRamp *OnRampTransactor) ForwardFromRouter(opts *bind.TransactOpts, message CCIPEVMToAnyTollMessage, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "forwardFromRouter", message, originalSender)
}

func (_OnRamp *OnRampSession) ForwardFromRouter(message CCIPEVMToAnyTollMessage, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.ForwardFromRouter(&_OnRamp.TransactOpts, message, originalSender)
}

func (_OnRamp *OnRampTransactorSession) ForwardFromRouter(message CCIPEVMToAnyTollMessage, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.ForwardFromRouter(&_OnRamp.TransactOpts, message, originalSender)
}

func (_OnRamp *OnRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "pause")
}

func (_OnRamp *OnRampSession) Pause() (*types.Transaction, error) {
	return _OnRamp.Contract.Pause(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactorSession) Pause() (*types.Transaction, error) {
	return _OnRamp.Contract.Pause(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactor) RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "removeFeed", token, feed)
}

func (_OnRamp *OnRampSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.RemoveFeed(&_OnRamp.TransactOpts, token, feed)
}

func (_OnRamp *OnRampTransactorSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.RemoveFeed(&_OnRamp.TransactOpts, token, feed)
}

func (_OnRamp *OnRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_OnRamp *OnRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.RemovePool(&_OnRamp.TransactOpts, token, pool)
}

func (_OnRamp *OnRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.RemovePool(&_OnRamp.TransactOpts, token, pool)
}

func (_OnRamp *OnRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "setAFN", afn)
}

func (_OnRamp *OnRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.SetAFN(&_OnRamp.TransactOpts, afn)
}

func (_OnRamp *OnRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.SetAFN(&_OnRamp.TransactOpts, afn)
}

func (_OnRamp *OnRampTransactor) SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "setAllowlist", allowlist)
}

func (_OnRamp *OnRampSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.SetAllowlist(&_OnRamp.TransactOpts, allowlist)
}

func (_OnRamp *OnRampTransactorSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.SetAllowlist(&_OnRamp.TransactOpts, allowlist)
}

func (_OnRamp *OnRampTransactor) SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "setAllowlistEnabled", enabled)
}

func (_OnRamp *OnRampSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _OnRamp.Contract.SetAllowlistEnabled(&_OnRamp.TransactOpts, enabled)
}

func (_OnRamp *OnRampTransactorSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _OnRamp.Contract.SetAllowlistEnabled(&_OnRamp.TransactOpts, enabled)
}

func (_OnRamp *OnRampTransactor) SetConfig(opts *bind.TransactOpts, config OnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "setConfig", config)
}

func (_OnRamp *OnRampSession) SetConfig(config OnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _OnRamp.Contract.SetConfig(&_OnRamp.TransactOpts, config)
}

func (_OnRamp *OnRampTransactorSession) SetConfig(config OnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _OnRamp.Contract.SetConfig(&_OnRamp.TransactOpts, config)
}

func (_OnRamp *OnRampTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_OnRamp *OnRampSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _OnRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_OnRamp.TransactOpts, newTime)
}

func (_OnRamp *OnRampTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _OnRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_OnRamp.TransactOpts, newTime)
}

func (_OnRamp *OnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_OnRamp *OnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.TransferOwnership(&_OnRamp.TransactOpts, to)
}

func (_OnRamp *OnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.TransferOwnership(&_OnRamp.TransactOpts, to)
}

func (_OnRamp *OnRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "unpause")
}

func (_OnRamp *OnRampSession) Unpause() (*types.Transaction, error) {
	return _OnRamp.Contract.Unpause(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _OnRamp.Contract.Unpause(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactor) WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "withdrawAccumulatedFees", feeToken, recipient, amount)
}

func (_OnRamp *OnRampSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OnRamp.Contract.WithdrawAccumulatedFees(&_OnRamp.TransactOpts, feeToken, recipient, amount)
}

func (_OnRamp *OnRampTransactorSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OnRamp.Contract.WithdrawAccumulatedFees(&_OnRamp.TransactOpts, feeToken, recipient, amount)
}

type OnRampAFNMaxHeartbeatTimeSetIterator struct {
	Event *OnRampAFNMaxHeartbeatTimeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAFNMaxHeartbeatTimeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAFNMaxHeartbeatTimeSet)
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
		it.Event = new(OnRampAFNMaxHeartbeatTimeSet)
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

func (it *OnRampAFNMaxHeartbeatTimeSetIterator) Error() error {
	return it.fail
}

func (it *OnRampAFNMaxHeartbeatTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAFNMaxHeartbeatTimeSet struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log
}

func (_OnRamp *OnRampFilterer) FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*OnRampAFNMaxHeartbeatTimeSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return &OnRampAFNMaxHeartbeatTimeSetIterator{contract: _OnRamp.contract, event: "AFNMaxHeartbeatTimeSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *OnRampAFNMaxHeartbeatTimeSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAFNMaxHeartbeatTimeSet)
				if err := _OnRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAFNMaxHeartbeatTimeSet(log types.Log) (*OnRampAFNMaxHeartbeatTimeSet, error) {
	event := new(OnRampAFNMaxHeartbeatTimeSet)
	if err := _OnRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampAFNSetIterator struct {
	Event *OnRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAFNSet)
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
		it.Event = new(OnRampAFNSet)
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

func (it *OnRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *OnRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_OnRamp *OnRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*OnRampAFNSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &OnRampAFNSetIterator{contract: _OnRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *OnRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAFNSet)
				if err := _OnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAFNSet(log types.Log) (*OnRampAFNSet, error) {
	event := new(OnRampAFNSet)
	if err := _OnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampAllowlistEnabledSetIterator struct {
	Event *OnRampAllowlistEnabledSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAllowlistEnabledSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAllowlistEnabledSet)
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
		it.Event = new(OnRampAllowlistEnabledSet)
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

func (it *OnRampAllowlistEnabledSetIterator) Error() error {
	return it.fail
}

func (it *OnRampAllowlistEnabledSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAllowlistEnabledSet struct {
	Enabled bool
	Raw     types.Log
}

func (_OnRamp *OnRampFilterer) FilterAllowlistEnabledSet(opts *bind.FilterOpts) (*OnRampAllowlistEnabledSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AllowlistEnabledSet")
	if err != nil {
		return nil, err
	}
	return &OnRampAllowlistEnabledSetIterator{contract: _OnRamp.contract, event: "AllowlistEnabledSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAllowlistEnabledSet(opts *bind.WatchOpts, sink chan<- *OnRampAllowlistEnabledSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AllowlistEnabledSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAllowlistEnabledSet)
				if err := _OnRamp.contract.UnpackLog(event, "AllowlistEnabledSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAllowlistEnabledSet(log types.Log) (*OnRampAllowlistEnabledSet, error) {
	event := new(OnRampAllowlistEnabledSet)
	if err := _OnRamp.contract.UnpackLog(event, "AllowlistEnabledSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampAllowlistSetIterator struct {
	Event *OnRampAllowlistSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAllowlistSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAllowlistSet)
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
		it.Event = new(OnRampAllowlistSet)
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

func (it *OnRampAllowlistSetIterator) Error() error {
	return it.fail
}

func (it *OnRampAllowlistSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAllowlistSet struct {
	Allowlist []common.Address
	Raw       types.Log
}

func (_OnRamp *OnRampFilterer) FilterAllowlistSet(opts *bind.FilterOpts) (*OnRampAllowlistSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AllowlistSet")
	if err != nil {
		return nil, err
	}
	return &OnRampAllowlistSetIterator{contract: _OnRamp.contract, event: "AllowlistSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAllowlistSet(opts *bind.WatchOpts, sink chan<- *OnRampAllowlistSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AllowlistSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAllowlistSet)
				if err := _OnRamp.contract.UnpackLog(event, "AllowlistSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAllowlistSet(log types.Log) (*OnRampAllowlistSet, error) {
	event := new(OnRampAllowlistSet)
	if err := _OnRamp.contract.UnpackLog(event, "AllowlistSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampCCIPSendRequestedIterator struct {
	Event *OnRampCCIPSendRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampCCIPSendRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampCCIPSendRequested)
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
		it.Event = new(OnRampCCIPSendRequested)
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

func (it *OnRampCCIPSendRequestedIterator) Error() error {
	return it.fail
}

func (it *OnRampCCIPSendRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampCCIPSendRequested struct {
	Message CCIPEVMToEVMTollEvent
	Raw     types.Log
}

func (_OnRamp *OnRampFilterer) FilterCCIPSendRequested(opts *bind.FilterOpts) (*OnRampCCIPSendRequestedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return &OnRampCCIPSendRequestedIterator{contract: _OnRamp.contract, event: "CCIPSendRequested", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *OnRampCCIPSendRequested) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampCCIPSendRequested)
				if err := _OnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseCCIPSendRequested(log types.Log) (*OnRampCCIPSendRequested, error) {
	event := new(OnRampCCIPSendRequested)
	if err := _OnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampFeeChargedIterator struct {
	Event *OnRampFeeCharged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampFeeChargedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampFeeCharged)
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
		it.Event = new(OnRampFeeCharged)
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

func (it *OnRampFeeChargedIterator) Error() error {
	return it.fail
}

func (it *OnRampFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampFeeCharged struct {
	From common.Address
	To   common.Address
	Fee  *big.Int
	Raw  types.Log
}

func (_OnRamp *OnRampFilterer) FilterFeeCharged(opts *bind.FilterOpts) (*OnRampFeeChargedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return &OnRampFeeChargedIterator{contract: _OnRamp.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *OnRampFeeCharged) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampFeeCharged)
				if err := _OnRamp.contract.UnpackLog(event, "FeeCharged", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseFeeCharged(log types.Log) (*OnRampFeeCharged, error) {
	event := new(OnRampFeeCharged)
	if err := _OnRamp.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampFeedAddedIterator struct {
	Event *OnRampFeedAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampFeedAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampFeedAdded)
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
		it.Event = new(OnRampFeedAdded)
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

func (it *OnRampFeedAddedIterator) Error() error {
	return it.fail
}

func (it *OnRampFeedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampFeedAdded struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_OnRamp *OnRampFilterer) FilterFeedAdded(opts *bind.FilterOpts) (*OnRampFeedAddedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return &OnRampFeedAddedIterator{contract: _OnRamp.contract, event: "FeedAdded", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *OnRampFeedAdded) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampFeedAdded)
				if err := _OnRamp.contract.UnpackLog(event, "FeedAdded", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseFeedAdded(log types.Log) (*OnRampFeedAdded, error) {
	event := new(OnRampFeedAdded)
	if err := _OnRamp.contract.UnpackLog(event, "FeedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampFeedRemovedIterator struct {
	Event *OnRampFeedRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampFeedRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampFeedRemoved)
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
		it.Event = new(OnRampFeedRemoved)
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

func (it *OnRampFeedRemovedIterator) Error() error {
	return it.fail
}

func (it *OnRampFeedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampFeedRemoved struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_OnRamp *OnRampFilterer) FilterFeedRemoved(opts *bind.FilterOpts) (*OnRampFeedRemovedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return &OnRampFeedRemovedIterator{contract: _OnRamp.contract, event: "FeedRemoved", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *OnRampFeedRemoved) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampFeedRemoved)
				if err := _OnRamp.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseFeedRemoved(log types.Log) (*OnRampFeedRemoved, error) {
	event := new(OnRampFeedRemoved)
	if err := _OnRamp.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampFeesWithdrawnIterator struct {
	Event *OnRampFeesWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampFeesWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampFeesWithdrawn)
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
		it.Event = new(OnRampFeesWithdrawn)
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

func (it *OnRampFeesWithdrawnIterator) Error() error {
	return it.fail
}

func (it *OnRampFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampFeesWithdrawn struct {
	FeeToken  common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_OnRamp *OnRampFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*OnRampFeesWithdrawnIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &OnRampFeesWithdrawnIterator{contract: _OnRamp.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampFeesWithdrawn)
				if err := _OnRamp.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseFeesWithdrawn(log types.Log) (*OnRampFeesWithdrawn, error) {
	event := new(OnRampFeesWithdrawn)
	if err := _OnRamp.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampNewTokenBucketConstructedIterator struct {
	Event *OnRampNewTokenBucketConstructed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampNewTokenBucketConstructedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampNewTokenBucketConstructed)
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
		it.Event = new(OnRampNewTokenBucketConstructed)
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

func (it *OnRampNewTokenBucketConstructedIterator) Error() error {
	return it.fail
}

func (it *OnRampNewTokenBucketConstructedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampNewTokenBucketConstructed struct {
	Rate     *big.Int
	Capacity *big.Int
	Full     bool
	Raw      types.Log
}

func (_OnRamp *OnRampFilterer) FilterNewTokenBucketConstructed(opts *bind.FilterOpts) (*OnRampNewTokenBucketConstructedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "NewTokenBucketConstructed")
	if err != nil {
		return nil, err
	}
	return &OnRampNewTokenBucketConstructedIterator{contract: _OnRamp.contract, event: "NewTokenBucketConstructed", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchNewTokenBucketConstructed(opts *bind.WatchOpts, sink chan<- *OnRampNewTokenBucketConstructed) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "NewTokenBucketConstructed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampNewTokenBucketConstructed)
				if err := _OnRamp.contract.UnpackLog(event, "NewTokenBucketConstructed", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseNewTokenBucketConstructed(log types.Log) (*OnRampNewTokenBucketConstructed, error) {
	event := new(OnRampNewTokenBucketConstructed)
	if err := _OnRamp.contract.UnpackLog(event, "NewTokenBucketConstructed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOnRampConfigSetIterator struct {
	Event *OnRampOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOnRampConfigSet)
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
		it.Event = new(OnRampOnRampConfigSet)
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

func (it *OnRampOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *OnRampOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOnRampConfigSet struct {
	Config OnRampInterfaceOnRampConfig
	Raw    types.Log
}

func (_OnRamp *OnRampFilterer) FilterOnRampConfigSet(opts *bind.FilterOpts) (*OnRampOnRampConfigSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "OnRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &OnRampOnRampConfigSetIterator{contract: _OnRamp.contract, event: "OnRampConfigSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchOnRampConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "OnRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOnRampConfigSet)
				if err := _OnRamp.contract.UnpackLog(event, "OnRampConfigSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseOnRampConfigSet(log types.Log) (*OnRampOnRampConfigSet, error) {
	event := new(OnRampOnRampConfigSet)
	if err := _OnRamp.contract.UnpackLog(event, "OnRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOwnershipTransferRequestedIterator struct {
	Event *OnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOwnershipTransferRequested)
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
		it.Event = new(OnRampOwnershipTransferRequested)
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

func (it *OnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRamp *OnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOwnershipTransferRequestedIterator{contract: _OnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOwnershipTransferRequested)
				if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*OnRampOwnershipTransferRequested, error) {
	event := new(OnRampOwnershipTransferRequested)
	if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOwnershipTransferredIterator struct {
	Event *OnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOwnershipTransferred)
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
		it.Event = new(OnRampOwnershipTransferred)
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

func (it *OnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRamp *OnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOwnershipTransferredIterator{contract: _OnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOwnershipTransferred)
				if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseOwnershipTransferred(log types.Log) (*OnRampOwnershipTransferred, error) {
	event := new(OnRampOwnershipTransferred)
	if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampPausedIterator struct {
	Event *OnRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampPaused)
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
		it.Event = new(OnRampPaused)
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

func (it *OnRampPausedIterator) Error() error {
	return it.fail
}

func (it *OnRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_OnRamp *OnRampFilterer) FilterPaused(opts *bind.FilterOpts) (*OnRampPausedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OnRampPausedIterator{contract: _OnRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OnRampPaused) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampPaused)
				if err := _OnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParsePaused(log types.Log) (*OnRampPaused, error) {
	event := new(OnRampPaused)
	if err := _OnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampPoolAddedIterator struct {
	Event *OnRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampPoolAdded)
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
		it.Event = new(OnRampPoolAdded)
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

func (it *OnRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *OnRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_OnRamp *OnRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*OnRampPoolAddedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &OnRampPoolAddedIterator{contract: _OnRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *OnRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampPoolAdded)
				if err := _OnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParsePoolAdded(log types.Log) (*OnRampPoolAdded, error) {
	event := new(OnRampPoolAdded)
	if err := _OnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampPoolRemovedIterator struct {
	Event *OnRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampPoolRemoved)
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
		it.Event = new(OnRampPoolRemoved)
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

func (it *OnRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *OnRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_OnRamp *OnRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*OnRampPoolRemovedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &OnRampPoolRemovedIterator{contract: _OnRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *OnRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampPoolRemoved)
				if err := _OnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParsePoolRemoved(log types.Log) (*OnRampPoolRemoved, error) {
	event := new(OnRampPoolRemoved)
	if err := _OnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampRouterSetIterator struct {
	Event *OnRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampRouterSet)
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
		it.Event = new(OnRampRouterSet)
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

func (it *OnRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *OnRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_OnRamp *OnRampFilterer) FilterRouterSet(opts *bind.FilterOpts) (*OnRampRouterSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return &OnRampRouterSetIterator{contract: _OnRamp.contract, event: "RouterSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchRouterSet(opts *bind.WatchOpts, sink chan<- *OnRampRouterSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampRouterSet)
				if err := _OnRamp.contract.UnpackLog(event, "RouterSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseRouterSet(log types.Log) (*OnRampRouterSet, error) {
	event := new(OnRampRouterSet)
	if err := _OnRamp.contract.UnpackLog(event, "RouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampUnpausedIterator struct {
	Event *OnRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampUnpaused)
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
		it.Event = new(OnRampUnpaused)
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

func (it *OnRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *OnRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_OnRamp *OnRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OnRampUnpausedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OnRampUnpausedIterator{contract: _OnRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OnRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampUnpaused)
				if err := _OnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseUnpaused(log types.Log) (*OnRampUnpaused, error) {
	event := new(OnRampUnpaused)
	if err := _OnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_OnRamp *OnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OnRamp.abi.Events["AFNMaxHeartbeatTimeSet"].ID:
		return _OnRamp.ParseAFNMaxHeartbeatTimeSet(log)
	case _OnRamp.abi.Events["AFNSet"].ID:
		return _OnRamp.ParseAFNSet(log)
	case _OnRamp.abi.Events["AllowlistEnabledSet"].ID:
		return _OnRamp.ParseAllowlistEnabledSet(log)
	case _OnRamp.abi.Events["AllowlistSet"].ID:
		return _OnRamp.ParseAllowlistSet(log)
	case _OnRamp.abi.Events["CCIPSendRequested"].ID:
		return _OnRamp.ParseCCIPSendRequested(log)
	case _OnRamp.abi.Events["FeeCharged"].ID:
		return _OnRamp.ParseFeeCharged(log)
	case _OnRamp.abi.Events["FeedAdded"].ID:
		return _OnRamp.ParseFeedAdded(log)
	case _OnRamp.abi.Events["FeedRemoved"].ID:
		return _OnRamp.ParseFeedRemoved(log)
	case _OnRamp.abi.Events["FeesWithdrawn"].ID:
		return _OnRamp.ParseFeesWithdrawn(log)
	case _OnRamp.abi.Events["NewTokenBucketConstructed"].ID:
		return _OnRamp.ParseNewTokenBucketConstructed(log)
	case _OnRamp.abi.Events["OnRampConfigSet"].ID:
		return _OnRamp.ParseOnRampConfigSet(log)
	case _OnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _OnRamp.ParseOwnershipTransferRequested(log)
	case _OnRamp.abi.Events["OwnershipTransferred"].ID:
		return _OnRamp.ParseOwnershipTransferred(log)
	case _OnRamp.abi.Events["Paused"].ID:
		return _OnRamp.ParsePaused(log)
	case _OnRamp.abi.Events["PoolAdded"].ID:
		return _OnRamp.ParsePoolAdded(log)
	case _OnRamp.abi.Events["PoolRemoved"].ID:
		return _OnRamp.ParsePoolRemoved(log)
	case _OnRamp.abi.Events["RouterSet"].ID:
		return _OnRamp.ParseRouterSet(log)
	case _OnRamp.abi.Events["Unpaused"].ID:
		return _OnRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (OnRampAFNMaxHeartbeatTimeSet) Topic() common.Hash {
	return common.HexToHash("0x72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251")
}

func (OnRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (OnRampAllowlistEnabledSet) Topic() common.Hash {
	return common.HexToHash("0xa1bf86c493917580dec207969ef59976f0c378f10ece581237f19acfbd858f1c")
}

func (OnRampAllowlistSet) Topic() common.Hash {
	return common.HexToHash("0x27f242de1bc4ed72c4329591ffff7d223b5f025e3514a07e05afec6d4eb889cf")
}

func (OnRampCCIPSendRequested) Topic() common.Hash {
	return common.HexToHash("0xab2ca9da6d303be28d1a5e854e3e170be286e07696245e77f8ea11f55367d481")
}

func (OnRampFeeCharged) Topic() common.Hash {
	return common.HexToHash("0x945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d")
}

func (OnRampFeedAdded) Topic() common.Hash {
	return common.HexToHash("0x037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb601305")
}

func (OnRampFeedRemoved) Topic() common.Hash {
	return common.HexToHash("0xa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d")
}

func (OnRampFeesWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8")
}

func (OnRampNewTokenBucketConstructed) Topic() common.Hash {
	return common.HexToHash("0xfaf3310019e551542b5c6014c1ae13e2a8d3943d7611d779c4df9b36c111924f")
}

func (OnRampOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xeac62265bdcb30e1e7a4822fecd5035bf208f242c899453ca9a3cdb5eb44225b")
}

func (OnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (OnRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (OnRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (OnRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (OnRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0xc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15")
}

func (OnRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_OnRamp *OnRamp) Address() common.Address {
	return _OnRamp.address
}

type OnRampInterface interface {
	CHAINID(opts *bind.CallOpts) (*big.Int, error)

	DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetAllowlist(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowlistEnabled(opts *bind.CallOpts) (bool, error)

	GetConfig(opts *bind.CallOpts) (OnRampInterfaceOnRampConfig, error)

	GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error)

	GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error)

	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetRequiredFee(opts *bind.CallOpts, feeToken common.Address) (*big.Int, error)

	GetSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetTokenPool(opts *bind.CallOpts, token common.Address) (common.Address, error)

	IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error)

	IsPool(opts *bind.CallOpts, addr common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, message CCIPEVMToAnyTollMessage, originalSender common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error)

	SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, config OnRampInterfaceOnRampConfig) (*types.Transaction, error)

	SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*OnRampAFNMaxHeartbeatTimeSetIterator, error)

	WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *OnRampAFNMaxHeartbeatTimeSet) (event.Subscription, error)

	ParseAFNMaxHeartbeatTimeSet(log types.Log) (*OnRampAFNMaxHeartbeatTimeSet, error)

	FilterAFNSet(opts *bind.FilterOpts) (*OnRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *OnRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*OnRampAFNSet, error)

	FilterAllowlistEnabledSet(opts *bind.FilterOpts) (*OnRampAllowlistEnabledSetIterator, error)

	WatchAllowlistEnabledSet(opts *bind.WatchOpts, sink chan<- *OnRampAllowlistEnabledSet) (event.Subscription, error)

	ParseAllowlistEnabledSet(log types.Log) (*OnRampAllowlistEnabledSet, error)

	FilterAllowlistSet(opts *bind.FilterOpts) (*OnRampAllowlistSetIterator, error)

	WatchAllowlistSet(opts *bind.WatchOpts, sink chan<- *OnRampAllowlistSet) (event.Subscription, error)

	ParseAllowlistSet(log types.Log) (*OnRampAllowlistSet, error)

	FilterCCIPSendRequested(opts *bind.FilterOpts) (*OnRampCCIPSendRequestedIterator, error)

	WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *OnRampCCIPSendRequested) (event.Subscription, error)

	ParseCCIPSendRequested(log types.Log) (*OnRampCCIPSendRequested, error)

	FilterFeeCharged(opts *bind.FilterOpts) (*OnRampFeeChargedIterator, error)

	WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *OnRampFeeCharged) (event.Subscription, error)

	ParseFeeCharged(log types.Log) (*OnRampFeeCharged, error)

	FilterFeedAdded(opts *bind.FilterOpts) (*OnRampFeedAddedIterator, error)

	WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *OnRampFeedAdded) (event.Subscription, error)

	ParseFeedAdded(log types.Log) (*OnRampFeedAdded, error)

	FilterFeedRemoved(opts *bind.FilterOpts) (*OnRampFeedRemovedIterator, error)

	WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *OnRampFeedRemoved) (event.Subscription, error)

	ParseFeedRemoved(log types.Log) (*OnRampFeedRemoved, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*OnRampFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*OnRampFeesWithdrawn, error)

	FilterNewTokenBucketConstructed(opts *bind.FilterOpts) (*OnRampNewTokenBucketConstructedIterator, error)

	WatchNewTokenBucketConstructed(opts *bind.WatchOpts, sink chan<- *OnRampNewTokenBucketConstructed) (event.Subscription, error)

	ParseNewTokenBucketConstructed(log types.Log) (*OnRampNewTokenBucketConstructed, error)

	FilterOnRampConfigSet(opts *bind.FilterOpts) (*OnRampOnRampConfigSetIterator, error)

	WatchOnRampConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampOnRampConfigSet) (event.Subscription, error)

	ParseOnRampConfigSet(log types.Log) (*OnRampOnRampConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OnRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*OnRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *OnRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*OnRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*OnRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *OnRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*OnRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*OnRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *OnRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*OnRampPoolRemoved, error)

	FilterRouterSet(opts *bind.FilterOpts) (*OnRampRouterSetIterator, error)

	WatchRouterSet(opts *bind.WatchOpts, sink chan<- *OnRampRouterSet) (event.Subscription, error)

	ParseRouterSet(log types.Log) (*OnRampRouterSet, error)

	FilterUnpaused(opts *bind.FilterOpts) (*OnRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OnRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*OnRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
