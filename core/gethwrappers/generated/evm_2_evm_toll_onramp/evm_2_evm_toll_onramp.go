// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_toll_onramp

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
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated"
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

type BaseOnRampInterfaceOnRampConfig struct {
	RelayingFeeJuels uint64
	MaxDataSize      uint64
	MaxTokensLength  uint64
}

type CCIPEVM2AnyTollMessage struct {
	Receiver       common.Address
	Data           []byte
	Tokens         []common.Address
	Amounts        []*big.Int
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	GasLimit       *big.Int
}

type CCIPEVM2EVMTollEvent struct {
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

var EVM2EVMTollOnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"contractAggregatorV2V3Interface[]\",\"name\":\"feeds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"internalType\":\"contractAny2EVMTollOnRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTokenAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceFeedConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenConfigMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenFeedMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"AllowListSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCCIP.EVM2EVMTollEvent\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OnRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DESTINATION_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"addFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVM2AnyTollMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"getRequiredFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"removeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"setAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setAllowlistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162003f4538038062003f458339810160408190526200003491620007e6565b6000805460ff1916815589908990899089908990899089908990899084908890879082908a9088903390819081620000b35760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ed57620000ed816200049f565b5050506001600160a01b0381166200011857604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200015a5760405162d8548360e71b815260040160405180910390fd5b81516200016f90600490602085019062000550565b5060005b825181101562000251576000828281518110620001945762000194620008ed565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001de57620001de620008ed565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600590925220805460ff19166001179055620002498162000903565b905062000173565b5050508051825114620002775760405163ee9d106b60e01b815260040160405180910390fd5b81516200028c90600790602085019062000550565b5060005b825181101562000357576000828281518110620002b157620002b1620008ed565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060066000868581518110620002fb57620002fb620008ed565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b0292909116919091179055506200034f8162000903565b905062000290565b5050815115905062000387576008805460ff1916600117905580516200038590600990602084019062000550565b505b60005b8151811015620003f4576001600a6000848481518110620003af57620003af620008ed565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055620003ec8162000903565b90506200038a565b505060809890985260a09690965250508351600c805460208701516040909701516001600160401b03908116600160801b02600160801b600160c01b031998821668010000000000000000026001600160801b031990931691909416171795909516179093555050600d80546001600160a01b039093166001600160a01b0319909316929092179091555050600b80546001600160401b0319169055506200092b9650505050505050565b336001600160a01b03821603620004f95760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000aa565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620005a8579160200282015b82811115620005a857825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000571565b50620005b6929150620005ba565b5090565b5b80821115620005b65760008155600101620005bb565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620006125762000612620005d1565b604052919050565b60006001600160401b03821115620006365762000636620005d1565b5060051b60200190565b6001600160a01b03811681146200065657600080fd5b50565b600082601f8301126200066b57600080fd5b81516020620006846200067e836200061a565b620005e7565b82815260059290921b84018101918181019086841115620006a457600080fd5b8286015b84811015620006cc578051620006be8162000640565b8352918301918301620006a8565b509695505050505050565b600082601f830112620006e957600080fd5b81516020620006fc6200067e836200061a565b82815260059290921b840181019181810190868411156200071c57600080fd5b8286015b84811015620006cc578051620007368162000640565b835291830191830162000720565b8051620007518162000640565b919050565b80516001600160401b03811681146200075157600080fd5b6000606082840312156200078157600080fd5b604051606081016001600160401b0381118282101715620007a657620007a6620005d1565b604052905080620007b78362000756565b8152620007c76020840162000756565b6020820152620007da6040840162000756565b60408201525092915050565b60008060008060008060008060006101608a8c0312156200080657600080fd5b895160208b015160408c0151919a5098506001600160401b03808211156200082d57600080fd5b6200083b8d838e0162000659565b985060608c01519150808211156200085257600080fd5b620008608d838e0162000659565b975060808c01519150808211156200087757600080fd5b620008858d838e01620006d7565b965060a08c01519150808211156200089c57600080fd5b50620008ab8c828d01620006d7565b945050620008bc60c08b0162000744565b9250620008cd8b60e08c016200076e565b9150620008de6101408b0162000744565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b6000600182016200092457634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a0516135ed6200095860003960006103d50152600081816104ea015261083b01526135ed6000f3fe608060405234801561001057600080fd5b506004361061025c5760003560e01c8063744b92e211610145578063bbe4f6db116100bd578063d0d5de611161008c578063eb511dd411610071578063eb511dd414610658578063f2fde38b1461066b578063f78faa321461067e57600080fd5b8063d0d5de6114610632578063d7644ba21461064557600080fd5b8063bbe4f6db14610261578063c0d7865514610568578063c3f909d41461057b578063c5eff3d01461061d57600080fd5b806385e1f4d0116101145780638da5cb5b116100f95780638da5cb5b14610514578063b0f479a114610537578063b4069b311461055557600080fd5b806385e1f4d0146104e557806389c065681461050c57600080fd5b8063744b92e2146104ba57806379ba5097146104cd57806381be8fa4146104d55780638456cb59146104dd57600080fd5b80633f4ba83a116101d857806359e96b5b116101a75780635c975abb1161018c5780635c975abb14610487578063671dc33714610492578063681fba16146104a557600080fd5b806359e96b5b1461043b5780635b16ebb71461044e57600080fd5b80633f4ba83a146104055780634120fccd1461040d578063552b818b146104155780635853c6271461042857600080fd5b806316b8e7311161022f5780632222dd42116102145780632222dd421461039f5780632b898c25146103bd5780632ea02369146103d057600080fd5b806316b8e7311461031d578063181f5a771461035657600080fd5b806304c2a34a1461026157806305afe24a146102c4578063108ee5fc146102f0578063147809b314610305575b600080fd5b61029a61026f366004612b42565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600360205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6102d76102d2366004612d87565b610689565b60405167ffffffffffffffff90911681526020016102bb565b6103036102fe366004612b42565b610982565b005b61030d610a5d565b60405190151581526020016102bb565b61029a61032b366004612b42565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600660205260409020541690565b6103926040518060400160405280601781526020017f45564d3245564d546f6c6c4f6e52616d7020312e302e3000000000000000000081525081565b6040516102bb9190612eff565b60025473ffffffffffffffffffffffffffffffffffffffff1661029a565b6103036103cb366004612f12565b610af7565b6103f77f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016102bb565b610303610ec7565b6102d7610ed9565b610303610423366004612f4b565b610ef9565b610303610436366004612f12565b6110e6565b610303610449366004612fc0565b6112f5565b61030d61045c366004612b42565b73ffffffffffffffffffffffffffffffffffffffff1660009081526005602052604090205460ff1690565b60005460ff1661030d565b6103036104a0366004613001565b611373565b6104ad6113c5565b6040516102bb919061306a565b6103036104c8366004612f12565b6114a4565b610303611899565b6104ad6119bb565b610303611a2a565b6103f77f000000000000000000000000000000000000000000000000000000000000000081565b6104ad611a3a565b600054610100900473ffffffffffffffffffffffffffffffffffffffff1661029a565b600d5473ffffffffffffffffffffffffffffffffffffffff1661029a565b61029a610563366004612b42565b611aa7565b610303610576366004612b42565b611baf565b6105ea60408051606081018252600080825260208201819052918101919091525060408051606081018252600c5467ffffffffffffffff808216835268010000000000000000820481166020840152700100000000000000000000000000000000909104169181019190915290565b60408051825167ffffffffffffffff908116825260208085015182169083015292820151909216908201526060016102bb565b610625611c2a565b6040516102bb919061307d565b6103f7610640366004612b42565b611c97565b6103036106533660046130e5565b611db6565b610303610666366004612f12565b611e1d565b610303610679366004612b42565b61205d565b60085460ff1661030d565b6000805460ff16156106fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610769573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061078d9190613102565b156107c3576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5473ffffffffffffffffffffffffffffffffffffffff163314610814576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61082d8360200151518460400151856060015185612071565b6040805161014081019091527f00000000000000000000000000000000000000000000000000000000000000008152600b8054600092916020830191849061087e9067ffffffffffffffff1661314e565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff168152602001856000015173ffffffffffffffffffffffffffffffffffffffff168152602001856020015181526020018560400151815260200185606001518152602001856080015173ffffffffffffffffffffffffffffffffffffffff1681526020018560a0015181526020018560c0015181525090507fab2ca9da6d303be28d1a5e854e3e170be286e07696245e77f8ea11f55367d4818160405161096f91906131a5565b60405180910390a1602001519392505050565b61098a6123d2565b73ffffffffffffffffffffffffffffffffffffffff81166109d7576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28910160405180910390a15050565b600254604080517f46f8e6d7000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff16916346f8e6d79160048083019260209291908290030181865afa158015610acd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610af19190613102565b15905090565b610aff6123d2565b6007546000819003610b3d576040517f2e70248b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260066020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610bd8576040517f3917193900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610c41576040517f9403a50500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006007610c506001856132d1565b81548110610c6057610c606132e8565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600782602001516bffffffffffffffffffffffff1681548110610cb257610cb26132e8565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166007610ce16001866132d1565b81548110610cf157610cf16132e8565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600783602001516bffffffffffffffffffffffff1681548110610d5f57610d5f6132e8565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526006909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556007805480610e0157610e01613317565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600683526040808520949094558351908152908716918101919091527fa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d91015b60405180910390a15050505050565b610ecf6123d2565b610ed7612458565b565b600b54600090610ef49067ffffffffffffffff166001613346565b905090565b610f016123d2565b60006009805480602002602001604051908101604052809291908181526020018280548015610f6657602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610f3b575b5050505050905060005b8151811015610ffe576000600a6000848481518110610f9157610f916132e8565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055610ff781613372565b9050610f70565b5061100b60098484612a7c565b5060005b828110156110a7576001600a600086868581811061102f5761102f6132e8565b90506020020160208101906110449190612b42565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790556110a081613372565b905061100f565b507ff8adc5fee247b62a85f63294cb46e4be61da815e4756bc57672a83b24faf0dda83836040516110d99291906133aa565b60405180910390a1505050565b6110ee6123d2565b73ffffffffffffffffffffffffffffffffffffffff82161580611125575073ffffffffffffffffffffffffffffffffffffffff8116155b1561115c576040517fee9d106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260066020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156111f8576040517f965ffb7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600780546bffffffffffffffffffffffff908116602080870191825288861660008181526006835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68890920180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821790558351908152908101919091527f037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb60130591016110d9565b6112fd6123d2565b61131e73ffffffffffffffffffffffffffffffffffffffff84168383612539565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8906060016110d9565b61137b6123d2565b80600c6113888282613410565b9050507fcc6ce9e57c1de2adf58a81e94b96b43d77ea6973e3f08e6ea4fe83d62ae60e9e816040516113ba91906134fe565b60405180910390a150565b60045460609067ffffffffffffffff8111156113e3576113e3612b5f565b60405190808252806020026020018201604052801561140c578160200160208202803683370190505b50905060005b6004548110156114a05761145960048281548110611432576114326132e8565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16611aa7565b82828151811061146b5761146b6132e8565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015261149981613372565b9050611412565b5090565b6114ac6123d2565b60045460008190036114ea576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290611585576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff16146115ee576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060046115fd6001856132d1565b8154811061160d5761160d6132e8565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600482602001516bffffffffffffffffffffffff168154811061165f5761165f6132e8565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16600461168e6001866132d1565b8154811061169e5761169e6132e8565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600483602001516bffffffffffffffffffffffff168154811061170c5761170c6132e8565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff90921674010000000000000000000000000000000000000000029190921617905560048054806117ae576117ae613317565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600383526040808520859055918816808552600584529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101610eb8565b60015473ffffffffffffffffffffffffffffffffffffffff16331461191a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016106f3565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60606007805480602002602001604051908101604052809291908181526020018280548015611a2057602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116119f5575b5050505050905090565b611a326123d2565b610ed76125cb565b60606004805480602002602001604051908101604052809291908181526020018280548015611a205760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116119f5575050505050905090565b73ffffffffffffffffffffffffffffffffffffffff80821660009081526003602052604081205490911680611b08576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611b84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ba89190613551565b9392505050565b611bb76123d2565b600d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15906020016113ba565b60606009805480602002602001604051908101604052809291908181526020018280548015611a205760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116119f5575050505050905090565b600080611cc98373ffffffffffffffffffffffffffffffffffffffff9081166000908152600660205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff8116611d30576040517feef7849700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024016106f3565b8073ffffffffffffffffffffffffffffffffffffffff166350d25bcd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d9f919061356e565b600c54611ba8919067ffffffffffffffff16613587565b611dbe6123d2565b600880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168215159081179091556040519081527fccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df032906020016113ba565b611e256123d2565b73ffffffffffffffffffffffffffffffffffffffff82161580611e5c575073ffffffffffffffffffffffffffffffffffffffff8116155b15611e93576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015611f2f576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600581529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c91016110d9565b6120656123d2565b61206e8161268b565b50565b600d5473ffffffffffffffffffffffffffffffffffffffff166120c0576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811661210d576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c5468010000000000000000900467ffffffffffffffff1684111561218157600c546040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090910467ffffffffffffffff166004820152602481018590526044016106f3565b8251600c54700100000000000000000000000000000000900467ffffffffffffffff168111806121b2575082518114155b156121e9576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085460ff168015612221575073ffffffffffffffffffffffffffffffffffffffff82166000908152600a602052604090205460ff16155b15612270576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016106f3565b60005b818110156123ca57600085828151811061228f5761228f6132e8565b6020026020010151905060006122ca8273ffffffffffffffffffffffffffffffffffffffff9081166000908152600360205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff8116612331576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016106f3565b8073ffffffffffffffffffffffffffffffffffffffff1663503c285887858151811061235f5761235f6132e8565b60200260200101516040518263ffffffff1660e01b815260040161238591815260200190565b600060405180830381600087803b15801561239f57600080fd5b505af11580156123b3573d6000803e3d6000fd5b505050505050806123c390613372565b9050612273565b505050505050565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610ed7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016106f3565b60005460ff166124c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016106f3565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526125c6908490612786565b505050565b60005460ff1615612638576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016106f3565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861250f3390565b3373ffffffffffffffffffffffffffffffffffffffff82160361270a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016106f3565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006127e8826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166128929092919063ffffffff16565b8051909150156125c657808060200190518101906128069190613102565b6125c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016106f3565b60606128a184846000856128a9565b949350505050565b60608247101561293b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016106f3565b843b6129a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016106f3565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516129cc91906135c4565b60006040518083038185875af1925050503d8060008114612a09576040519150601f19603f3d011682016040523d82523d6000602084013e612a0e565b606091505b5091509150612a1e828286612a29565b979650505050505050565b60608315612a38575081611ba8565b825115612a485782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f39190612eff565b828054828255906000526020600020908101928215612af4579160200282015b82811115612af45781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff843516178255602090920191600190910190612a9c565b506114a09291505b808211156114a05760008155600101612afc565b73ffffffffffffffffffffffffffffffffffffffff8116811461206e57600080fd5b8035612b3d81612b10565b919050565b600060208284031215612b5457600080fd5b8135611ba881612b10565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715612bb157612bb1612b5f565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612bfe57612bfe612b5f565b604052919050565b600082601f830112612c1757600080fd5b813567ffffffffffffffff811115612c3157612c31612b5f565b612c6260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601612bb7565b818152846020838601011115612c7757600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115612cae57612cae612b5f565b5060051b60200190565b600082601f830112612cc957600080fd5b81356020612cde612cd983612c94565b612bb7565b82815260059290921b84018101918181019086841115612cfd57600080fd5b8286015b84811015612d21578035612d1481612b10565b8352918301918301612d01565b509695505050505050565b600082601f830112612d3d57600080fd5b81356020612d4d612cd983612c94565b82815260059290921b84018101918181019086841115612d6c57600080fd5b8286015b84811015612d215780358352918301918301612d70565b60008060408385031215612d9a57600080fd5b823567ffffffffffffffff80821115612db257600080fd5b9084019060e08287031215612dc657600080fd5b612dce612b8e565b612dd783612b32565b8152602083013582811115612deb57600080fd5b612df788828601612c06565b602083015250604083013582811115612e0f57600080fd5b612e1b88828601612cb8565b604083015250606083013582811115612e3357600080fd5b612e3f88828601612d2c565b606083015250612e5160808401612b32565b608082015260a083013560a082015260c083013560c0820152809450505050612e7c60208401612b32565b90509250929050565b60005b83811015612ea0578181015183820152602001612e88565b83811115612eaf576000848401525b50505050565b60008151808452612ecd816020860160208601612e85565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611ba86020830184612eb5565b60008060408385031215612f2557600080fd5b8235612f3081612b10565b91506020830135612f4081612b10565b809150509250929050565b60008060208385031215612f5e57600080fd5b823567ffffffffffffffff80821115612f7657600080fd5b818501915085601f830112612f8a57600080fd5b813581811115612f9957600080fd5b8660208260051b8501011115612fae57600080fd5b60209290920196919550909350505050565b600080600060608486031215612fd557600080fd5b8335612fe081612b10565b92506020840135612ff081612b10565b929592945050506040919091013590565b60006060828403121561301357600080fd5b50919050565b600081518084526020808501945080840160005b8381101561305f57815173ffffffffffffffffffffffffffffffffffffffff168752958201959082019060010161302d565b509495945050505050565b602081526000611ba86020830184613019565b6020808252825182820181905260009190848201906040850190845b818110156130cb57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101613099565b50909695505050505050565b801515811461206e57600080fd5b6000602082840312156130f757600080fd5b8135611ba8816130d7565b60006020828403121561311457600080fd5b8151611ba8816130d7565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681810361316b5761316b61311f565b6001019392505050565b600081518084526020808501945080840160005b8381101561305f57815187529582019590820190600101613189565b6020815281516020820152600060208301516131cd604084018267ffffffffffffffff169052565b50604083015173ffffffffffffffffffffffffffffffffffffffff8116606084015250606083015173ffffffffffffffffffffffffffffffffffffffff811660808401525060808301516101408060a085015261322e610160850183612eb5565b915060a08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808685030160c087015261326a8483613019565b935060c08701519150808685030160e0870152506132888382613175565b92505060e08501516101006132b48187018373ffffffffffffffffffffffffffffffffffffffff169052565b860151610120868101919091529095015193019290925250919050565b6000828210156132e3576132e361311f565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600067ffffffffffffffff8083168185168083038211156133695761336961311f565b01949350505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036133a3576133a361311f565b5060010190565b60208082528181018390526000908460408401835b86811015612d215782356133d281612b10565b73ffffffffffffffffffffffffffffffffffffffff16825291830191908301906001016133bf565b67ffffffffffffffff8116811461206e57600080fd5b813561341b816133fa565b67ffffffffffffffff811690508154817fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008216178355602084013561345f816133fa565b6fffffffffffffffff00000000000000008160401b16905080837fffffffffffffffffffffffffffffffff0000000000000000000000000000000084161717845560408501356134ae816133fa565b77ffffffffffffffff000000000000000000000000000000008160801b16847fffffffffffffffff0000000000000000000000000000000000000000000000008516178317178555505050505050565b60608101823561350d816133fa565b67ffffffffffffffff908116835260208401359061352a826133fa565b9081166020840152604084013590613541826133fa565b8082166040850152505092915050565b60006020828403121561356357600080fd5b8151611ba881612b10565b60006020828403121561358057600080fd5b5051919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156135bf576135bf61311f565b500290565b600082516135d6818460208701612e85565b919091019291505056fea164736f6c634300080f000a",
}

var EVM2EVMTollOnRampABI = EVM2EVMTollOnRampMetaData.ABI

var EVM2EVMTollOnRampBin = EVM2EVMTollOnRampMetaData.Bin

func DeployEVM2EVMTollOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, chainId *big.Int, destinationChainId *big.Int, tokens []common.Address, pools []common.Address, feeds []common.Address, allowlist []common.Address, afn common.Address, config BaseOnRampInterfaceOnRampConfig, router common.Address) (common.Address, *types.Transaction, *EVM2EVMTollOnRamp, error) {
	parsed, err := EVM2EVMTollOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMTollOnRampBin), backend, chainId, destinationChainId, tokens, pools, feeds, allowlist, afn, config, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMTollOnRamp{EVM2EVMTollOnRampCaller: EVM2EVMTollOnRampCaller{contract: contract}, EVM2EVMTollOnRampTransactor: EVM2EVMTollOnRampTransactor{contract: contract}, EVM2EVMTollOnRampFilterer: EVM2EVMTollOnRampFilterer{contract: contract}}, nil
}

type EVM2EVMTollOnRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMTollOnRampCaller
	EVM2EVMTollOnRampTransactor
	EVM2EVMTollOnRampFilterer
}

type EVM2EVMTollOnRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOnRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOnRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOnRampSession struct {
	Contract     *EVM2EVMTollOnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMTollOnRampCallerSession struct {
	Contract *EVM2EVMTollOnRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMTollOnRampTransactorSession struct {
	Contract     *EVM2EVMTollOnRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMTollOnRampRaw struct {
	Contract *EVM2EVMTollOnRamp
}

type EVM2EVMTollOnRampCallerRaw struct {
	Contract *EVM2EVMTollOnRampCaller
}

type EVM2EVMTollOnRampTransactorRaw struct {
	Contract *EVM2EVMTollOnRampTransactor
}

func NewEVM2EVMTollOnRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMTollOnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMTollOnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMTollOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRamp{address: address, abi: abi, EVM2EVMTollOnRampCaller: EVM2EVMTollOnRampCaller{contract: contract}, EVM2EVMTollOnRampTransactor: EVM2EVMTollOnRampTransactor{contract: contract}, EVM2EVMTollOnRampFilterer: EVM2EVMTollOnRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMTollOnRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMTollOnRampCaller, error) {
	contract, err := bindEVM2EVMTollOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampCaller{contract: contract}, nil
}

func NewEVM2EVMTollOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMTollOnRampTransactor, error) {
	contract, err := bindEVM2EVMTollOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampTransactor{contract: contract}, nil
}

func NewEVM2EVMTollOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMTollOnRampFilterer, error) {
	contract, err := bindEVM2EVMTollOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampFilterer{contract: contract}, nil
}

func bindEVM2EVMTollOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMTollOnRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMTollOnRamp.Contract.EVM2EVMTollOnRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.EVM2EVMTollOnRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.EVM2EVMTollOnRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMTollOnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) CHAINID() (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.CHAINID(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) CHAINID() (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.CHAINID(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "DESTINATION_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.DESTINATIONCHAINID(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.DESTINATIONCHAINID(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetAFN() (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetAFN(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetAFN(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetAllowlist(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getAllowlist")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetAllowlist() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetAllowlist(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetAllowlist() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetAllowlist(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetAllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getAllowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetAllowlistEnabled() (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.GetAllowlistEnabled(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetAllowlistEnabled() (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.GetAllowlistEnabled(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetConfig(opts *bind.CallOpts) (BaseOnRampInterfaceOnRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOnRampInterfaceOnRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOnRampInterfaceOnRampConfig)).(*BaseOnRampInterfaceOnRampConfig)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetConfig() (BaseOnRampInterfaceOnRampConfig, error) {
	return _EVM2EVMTollOnRamp.Contract.GetConfig(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetConfig() (BaseOnRampInterfaceOnRampConfig, error) {
	return _EVM2EVMTollOnRamp.Contract.GetConfig(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetDestinationToken(&_EVM2EVMTollOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetDestinationToken(&_EVM2EVMTollOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetDestinationTokens(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetDestinationTokens(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getExpectedNextSequenceNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetExpectedNextSequenceNumber() (uint64, error) {
	return _EVM2EVMTollOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetExpectedNextSequenceNumber() (uint64, error) {
	return _EVM2EVMTollOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getFeed", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetFeed(token common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetFeed(&_EVM2EVMTollOnRamp.CallOpts, token)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetFeed(token common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetFeed(&_EVM2EVMTollOnRamp.CallOpts, token)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getFeedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetFeedTokens() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetFeedTokens(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetFeedTokens() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetFeedTokens(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetPool(&_EVM2EVMTollOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetPool(&_EVM2EVMTollOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetPoolTokens(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetPoolTokens(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetRequiredFee(opts *bind.CallOpts, feeToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getRequiredFee", feeToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetRequiredFee(feeToken common.Address) (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.GetRequiredFee(&_EVM2EVMTollOnRamp.CallOpts, feeToken)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetRequiredFee(feeToken common.Address) (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.GetRequiredFee(&_EVM2EVMTollOnRamp.CallOpts, feeToken)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetRouter() (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetRouter(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetRouter(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetTokenPool(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getTokenPool", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetTokenPool(token common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetTokenPool(&_EVM2EVMTollOnRamp.CallOpts, token)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetTokenPool(token common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetTokenPool(&_EVM2EVMTollOnRamp.CallOpts, token)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.IsAFNHealthy(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.IsAFNHealthy(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) IsPool(addr common.Address) (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.IsPool(&_EVM2EVMTollOnRamp.CallOpts, addr)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) IsPool(addr common.Address) (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.IsPool(&_EVM2EVMTollOnRamp.CallOpts, addr)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) Owner() (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.Owner(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.Owner(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) Paused() (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.Paused(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.Paused(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMTollOnRamp.Contract.TypeAndVersion(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMTollOnRamp.Contract.TypeAndVersion(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.AcceptOwnership(&_EVM2EVMTollOnRamp.TransactOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.AcceptOwnership(&_EVM2EVMTollOnRamp.TransactOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "addFeed", token, feed)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.AddFeed(&_EVM2EVMTollOnRamp.TransactOpts, token, feed)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.AddFeed(&_EVM2EVMTollOnRamp.TransactOpts, token, feed)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.AddPool(&_EVM2EVMTollOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.AddPool(&_EVM2EVMTollOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) ForwardFromRouter(opts *bind.TransactOpts, message CCIPEVM2AnyTollMessage, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "forwardFromRouter", message, originalSender)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) ForwardFromRouter(message CCIPEVM2AnyTollMessage, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.ForwardFromRouter(&_EVM2EVMTollOnRamp.TransactOpts, message, originalSender)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) ForwardFromRouter(message CCIPEVM2AnyTollMessage, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.ForwardFromRouter(&_EVM2EVMTollOnRamp.TransactOpts, message, originalSender)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.Pause(&_EVM2EVMTollOnRamp.TransactOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.Pause(&_EVM2EVMTollOnRamp.TransactOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "removeFeed", token, feed)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.RemoveFeed(&_EVM2EVMTollOnRamp.TransactOpts, token, feed)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.RemoveFeed(&_EVM2EVMTollOnRamp.TransactOpts, token, feed)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.RemovePool(&_EVM2EVMTollOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.RemovePool(&_EVM2EVMTollOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetAFN(&_EVM2EVMTollOnRamp.TransactOpts, afn)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetAFN(&_EVM2EVMTollOnRamp.TransactOpts, afn)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setAllowlist", allowlist)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetAllowlist(&_EVM2EVMTollOnRamp.TransactOpts, allowlist)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetAllowlist(&_EVM2EVMTollOnRamp.TransactOpts, allowlist)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setAllowlistEnabled", enabled)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetAllowlistEnabled(&_EVM2EVMTollOnRamp.TransactOpts, enabled)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetAllowlistEnabled(&_EVM2EVMTollOnRamp.TransactOpts, enabled)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetConfig(opts *bind.TransactOpts, config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setConfig", config)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetConfig(config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetConfig(&_EVM2EVMTollOnRamp.TransactOpts, config)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetConfig(config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetConfig(&_EVM2EVMTollOnRamp.TransactOpts, config)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetRouter(&_EVM2EVMTollOnRamp.TransactOpts, router)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetRouter(&_EVM2EVMTollOnRamp.TransactOpts, router)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.TransferOwnership(&_EVM2EVMTollOnRamp.TransactOpts, to)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.TransferOwnership(&_EVM2EVMTollOnRamp.TransactOpts, to)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.Unpause(&_EVM2EVMTollOnRamp.TransactOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.Unpause(&_EVM2EVMTollOnRamp.TransactOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "withdrawAccumulatedFees", feeToken, recipient, amount)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.WithdrawAccumulatedFees(&_EVM2EVMTollOnRamp.TransactOpts, feeToken, recipient, amount)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.WithdrawAccumulatedFees(&_EVM2EVMTollOnRamp.TransactOpts, feeToken, recipient, amount)
}

type EVM2EVMTollOnRampAFNSetIterator struct {
	Event *EVM2EVMTollOnRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampAFNSet)
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
		it.Event = new(EVM2EVMTollOnRampAFNSet)
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

func (it *EVM2EVMTollOnRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampAFNSetIterator{contract: _EVM2EVMTollOnRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampAFNSet)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseAFNSet(log types.Log) (*EVM2EVMTollOnRampAFNSet, error) {
	event := new(EVM2EVMTollOnRampAFNSet)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampAllowListEnabledSetIterator struct {
	Event *EVM2EVMTollOnRampAllowListEnabledSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampAllowListEnabledSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampAllowListEnabledSet)
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
		it.Event = new(EVM2EVMTollOnRampAllowListEnabledSet)
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

func (it *EVM2EVMTollOnRampAllowListEnabledSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampAllowListEnabledSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampAllowListEnabledSet struct {
	Enabled bool
	Raw     types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterAllowListEnabledSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampAllowListEnabledSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "AllowListEnabledSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampAllowListEnabledSetIterator{contract: _EVM2EVMTollOnRamp.contract, event: "AllowListEnabledSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchAllowListEnabledSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampAllowListEnabledSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "AllowListEnabledSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampAllowListEnabledSet)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "AllowListEnabledSet", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseAllowListEnabledSet(log types.Log) (*EVM2EVMTollOnRampAllowListEnabledSet, error) {
	event := new(EVM2EVMTollOnRampAllowListEnabledSet)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "AllowListEnabledSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampAllowListSetIterator struct {
	Event *EVM2EVMTollOnRampAllowListSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampAllowListSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampAllowListSet)
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
		it.Event = new(EVM2EVMTollOnRampAllowListSet)
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

func (it *EVM2EVMTollOnRampAllowListSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampAllowListSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampAllowListSet struct {
	Allowlist []common.Address
	Raw       types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterAllowListSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampAllowListSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "AllowListSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampAllowListSetIterator{contract: _EVM2EVMTollOnRamp.contract, event: "AllowListSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchAllowListSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampAllowListSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "AllowListSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampAllowListSet)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "AllowListSet", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseAllowListSet(log types.Log) (*EVM2EVMTollOnRampAllowListSet, error) {
	event := new(EVM2EVMTollOnRampAllowListSet)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "AllowListSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampCCIPSendRequestedIterator struct {
	Event *EVM2EVMTollOnRampCCIPSendRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampCCIPSendRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampCCIPSendRequested)
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
		it.Event = new(EVM2EVMTollOnRampCCIPSendRequested)
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

func (it *EVM2EVMTollOnRampCCIPSendRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampCCIPSendRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampCCIPSendRequested struct {
	Message CCIPEVM2EVMTollEvent
	Raw     types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMTollOnRampCCIPSendRequestedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampCCIPSendRequestedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "CCIPSendRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampCCIPSendRequested) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampCCIPSendRequested)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseCCIPSendRequested(log types.Log) (*EVM2EVMTollOnRampCCIPSendRequested, error) {
	event := new(EVM2EVMTollOnRampCCIPSendRequested)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampFeeChargedIterator struct {
	Event *EVM2EVMTollOnRampFeeCharged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampFeeChargedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampFeeCharged)
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
		it.Event = new(EVM2EVMTollOnRampFeeCharged)
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

func (it *EVM2EVMTollOnRampFeeChargedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampFeeCharged struct {
	From common.Address
	To   common.Address
	Fee  *big.Int
	Raw  types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterFeeCharged(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeeChargedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampFeeChargedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeeCharged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampFeeCharged)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeeCharged", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseFeeCharged(log types.Log) (*EVM2EVMTollOnRampFeeCharged, error) {
	event := new(EVM2EVMTollOnRampFeeCharged)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampFeedAddedIterator struct {
	Event *EVM2EVMTollOnRampFeedAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampFeedAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampFeedAdded)
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
		it.Event = new(EVM2EVMTollOnRampFeedAdded)
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

func (it *EVM2EVMTollOnRampFeedAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampFeedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampFeedAdded struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterFeedAdded(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeedAddedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampFeedAddedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "FeedAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeedAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampFeedAdded)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeedAdded", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseFeedAdded(log types.Log) (*EVM2EVMTollOnRampFeedAdded, error) {
	event := new(EVM2EVMTollOnRampFeedAdded)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampFeedRemovedIterator struct {
	Event *EVM2EVMTollOnRampFeedRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampFeedRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampFeedRemoved)
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
		it.Event = new(EVM2EVMTollOnRampFeedRemoved)
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

func (it *EVM2EVMTollOnRampFeedRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampFeedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampFeedRemoved struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterFeedRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeedRemovedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampFeedRemovedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "FeedRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeedRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampFeedRemoved)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseFeedRemoved(log types.Log) (*EVM2EVMTollOnRampFeedRemoved, error) {
	event := new(EVM2EVMTollOnRampFeedRemoved)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampFeesWithdrawnIterator struct {
	Event *EVM2EVMTollOnRampFeesWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampFeesWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampFeesWithdrawn)
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
		it.Event = new(EVM2EVMTollOnRampFeesWithdrawn)
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

func (it *EVM2EVMTollOnRampFeesWithdrawnIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampFeesWithdrawn struct {
	FeeToken  common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeesWithdrawnIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampFeesWithdrawnIterator{contract: _EVM2EVMTollOnRamp.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampFeesWithdrawn)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseFeesWithdrawn(log types.Log) (*EVM2EVMTollOnRampFeesWithdrawn, error) {
	event := new(EVM2EVMTollOnRampFeesWithdrawn)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampOnRampConfigSetIterator struct {
	Event *EVM2EVMTollOnRampOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampOnRampConfigSet)
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
		it.Event = new(EVM2EVMTollOnRampOnRampConfigSet)
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

func (it *EVM2EVMTollOnRampOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampOnRampConfigSet struct {
	Config BaseOnRampInterfaceOnRampConfig
	Raw    types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterOnRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampOnRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "OnRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampOnRampConfigSetIterator{contract: _EVM2EVMTollOnRamp.contract, event: "OnRampConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchOnRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "OnRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampOnRampConfigSet)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "OnRampConfigSet", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseOnRampConfigSet(log types.Log) (*EVM2EVMTollOnRampOnRampConfigSet, error) {
	event := new(EVM2EVMTollOnRampOnRampConfigSet)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "OnRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMTollOnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMTollOnRampOwnershipTransferRequested)
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

func (it *EVM2EVMTollOnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampOwnershipTransferRequestedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampOwnershipTransferRequested)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMTollOnRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMTollOnRampOwnershipTransferRequested)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampOwnershipTransferredIterator struct {
	Event *EVM2EVMTollOnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMTollOnRampOwnershipTransferred)
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

func (it *EVM2EVMTollOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampOwnershipTransferredIterator{contract: _EVM2EVMTollOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampOwnershipTransferred)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMTollOnRampOwnershipTransferred, error) {
	event := new(EVM2EVMTollOnRampOwnershipTransferred)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampPausedIterator struct {
	Event *EVM2EVMTollOnRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampPaused)
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
		it.Event = new(EVM2EVMTollOnRampPaused)
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

func (it *EVM2EVMTollOnRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMTollOnRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampPausedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampPaused)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParsePaused(log types.Log) (*EVM2EVMTollOnRampPaused, error) {
	event := new(EVM2EVMTollOnRampPaused)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampPoolAddedIterator struct {
	Event *EVM2EVMTollOnRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampPoolAdded)
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
		it.Event = new(EVM2EVMTollOnRampPoolAdded)
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

func (it *EVM2EVMTollOnRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMTollOnRampPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampPoolAddedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampPoolAdded)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMTollOnRampPoolAdded, error) {
	event := new(EVM2EVMTollOnRampPoolAdded)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampPoolRemovedIterator struct {
	Event *EVM2EVMTollOnRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampPoolRemoved)
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
		it.Event = new(EVM2EVMTollOnRampPoolRemoved)
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

func (it *EVM2EVMTollOnRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOnRampPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampPoolRemovedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampPoolRemoved)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMTollOnRampPoolRemoved, error) {
	event := new(EVM2EVMTollOnRampPoolRemoved)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampRouterSetIterator struct {
	Event *EVM2EVMTollOnRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampRouterSet)
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
		it.Event = new(EVM2EVMTollOnRampRouterSet)
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

func (it *EVM2EVMTollOnRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterRouterSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampRouterSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampRouterSetIterator{contract: _EVM2EVMTollOnRamp.contract, event: "RouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampRouterSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampRouterSet)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "RouterSet", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseRouterSet(log types.Log) (*EVM2EVMTollOnRampRouterSet, error) {
	event := new(EVM2EVMTollOnRampRouterSet)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "RouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampUnpausedIterator struct {
	Event *EVM2EVMTollOnRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampUnpaused)
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
		it.Event = new(EVM2EVMTollOnRampUnpaused)
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

func (it *EVM2EVMTollOnRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOnRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampUnpausedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampUnpaused)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMTollOnRampUnpaused, error) {
	event := new(EVM2EVMTollOnRampUnpaused)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMTollOnRamp.abi.Events["AFNSet"].ID:
		return _EVM2EVMTollOnRamp.ParseAFNSet(log)
	case _EVM2EVMTollOnRamp.abi.Events["AllowListEnabledSet"].ID:
		return _EVM2EVMTollOnRamp.ParseAllowListEnabledSet(log)
	case _EVM2EVMTollOnRamp.abi.Events["AllowListSet"].ID:
		return _EVM2EVMTollOnRamp.ParseAllowListSet(log)
	case _EVM2EVMTollOnRamp.abi.Events["CCIPSendRequested"].ID:
		return _EVM2EVMTollOnRamp.ParseCCIPSendRequested(log)
	case _EVM2EVMTollOnRamp.abi.Events["FeeCharged"].ID:
		return _EVM2EVMTollOnRamp.ParseFeeCharged(log)
	case _EVM2EVMTollOnRamp.abi.Events["FeedAdded"].ID:
		return _EVM2EVMTollOnRamp.ParseFeedAdded(log)
	case _EVM2EVMTollOnRamp.abi.Events["FeedRemoved"].ID:
		return _EVM2EVMTollOnRamp.ParseFeedRemoved(log)
	case _EVM2EVMTollOnRamp.abi.Events["FeesWithdrawn"].ID:
		return _EVM2EVMTollOnRamp.ParseFeesWithdrawn(log)
	case _EVM2EVMTollOnRamp.abi.Events["OnRampConfigSet"].ID:
		return _EVM2EVMTollOnRamp.ParseOnRampConfigSet(log)
	case _EVM2EVMTollOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMTollOnRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMTollOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMTollOnRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMTollOnRamp.abi.Events["Paused"].ID:
		return _EVM2EVMTollOnRamp.ParsePaused(log)
	case _EVM2EVMTollOnRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMTollOnRamp.ParsePoolAdded(log)
	case _EVM2EVMTollOnRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMTollOnRamp.ParsePoolRemoved(log)
	case _EVM2EVMTollOnRamp.abi.Events["RouterSet"].ID:
		return _EVM2EVMTollOnRamp.ParseRouterSet(log)
	case _EVM2EVMTollOnRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMTollOnRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMTollOnRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMTollOnRampAllowListEnabledSet) Topic() common.Hash {
	return common.HexToHash("0xccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df032")
}

func (EVM2EVMTollOnRampAllowListSet) Topic() common.Hash {
	return common.HexToHash("0xf8adc5fee247b62a85f63294cb46e4be61da815e4756bc57672a83b24faf0dda")
}

func (EVM2EVMTollOnRampCCIPSendRequested) Topic() common.Hash {
	return common.HexToHash("0xab2ca9da6d303be28d1a5e854e3e170be286e07696245e77f8ea11f55367d481")
}

func (EVM2EVMTollOnRampFeeCharged) Topic() common.Hash {
	return common.HexToHash("0x945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d")
}

func (EVM2EVMTollOnRampFeedAdded) Topic() common.Hash {
	return common.HexToHash("0x037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb601305")
}

func (EVM2EVMTollOnRampFeedRemoved) Topic() common.Hash {
	return common.HexToHash("0xa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d")
}

func (EVM2EVMTollOnRampFeesWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8")
}

func (EVM2EVMTollOnRampOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xcc6ce9e57c1de2adf58a81e94b96b43d77ea6973e3f08e6ea4fe83d62ae60e9e")
}

func (EVM2EVMTollOnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMTollOnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMTollOnRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMTollOnRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMTollOnRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMTollOnRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0xc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15")
}

func (EVM2EVMTollOnRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRamp) Address() common.Address {
	return _EVM2EVMTollOnRamp.address
}

type EVM2EVMTollOnRampInterface interface {
	CHAINID(opts *bind.CallOpts) (*big.Int, error)

	DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetAllowlist(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowlistEnabled(opts *bind.CallOpts) (bool, error)

	GetConfig(opts *bind.CallOpts) (BaseOnRampInterfaceOnRampConfig, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error)

	GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetRequiredFee(opts *bind.CallOpts, feeToken common.Address) (*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetTokenPool(opts *bind.CallOpts, token common.Address) (common.Address, error)

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

	IsPool(opts *bind.CallOpts, addr common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, message CCIPEVM2AnyTollMessage, originalSender common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error)

	SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMTollOnRampAFNSet, error)

	FilterAllowListEnabledSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampAllowListEnabledSetIterator, error)

	WatchAllowListEnabledSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampAllowListEnabledSet) (event.Subscription, error)

	ParseAllowListEnabledSet(log types.Log) (*EVM2EVMTollOnRampAllowListEnabledSet, error)

	FilterAllowListSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampAllowListSetIterator, error)

	WatchAllowListSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampAllowListSet) (event.Subscription, error)

	ParseAllowListSet(log types.Log) (*EVM2EVMTollOnRampAllowListSet, error)

	FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMTollOnRampCCIPSendRequestedIterator, error)

	WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampCCIPSendRequested) (event.Subscription, error)

	ParseCCIPSendRequested(log types.Log) (*EVM2EVMTollOnRampCCIPSendRequested, error)

	FilterFeeCharged(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeeChargedIterator, error)

	WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeeCharged) (event.Subscription, error)

	ParseFeeCharged(log types.Log) (*EVM2EVMTollOnRampFeeCharged, error)

	FilterFeedAdded(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeedAddedIterator, error)

	WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeedAdded) (event.Subscription, error)

	ParseFeedAdded(log types.Log) (*EVM2EVMTollOnRampFeedAdded, error)

	FilterFeedRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeedRemovedIterator, error)

	WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeedRemoved) (event.Subscription, error)

	ParseFeedRemoved(log types.Log) (*EVM2EVMTollOnRampFeedRemoved, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*EVM2EVMTollOnRampFeesWithdrawn, error)

	FilterOnRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampOnRampConfigSetIterator, error)

	WatchOnRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampOnRampConfigSet) (event.Subscription, error)

	ParseOnRampConfigSet(log types.Log) (*EVM2EVMTollOnRampOnRampConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMTollOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMTollOnRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMTollOnRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMTollOnRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMTollOnRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMTollOnRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOnRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMTollOnRampPoolRemoved, error)

	FilterRouterSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampRouterSetIterator, error)

	WatchRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampRouterSet) (event.Subscription, error)

	ParseRouterSet(log types.Log) (*EVM2EVMTollOnRampRouterSet, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOnRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMTollOnRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
