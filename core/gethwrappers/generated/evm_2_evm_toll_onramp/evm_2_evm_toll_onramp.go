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

type AggregateRateLimiterInterfaceRateLimiterConfig struct {
	Rate     *big.Int
	Capacity *big.Int
}

type AggregateRateLimiterInterfaceTokenBucket struct {
	Rate        *big.Int
	Capacity    *big.Int
	Tokens      *big.Int
	LastUpdated *big.Int
}

type Any2EVMTollOnRampInterfaceFeeConfig struct {
	Fees      []*big.Int
	FeeTokens []common.Address
}

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

type CCIPEVM2EVMTollMessage struct {
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"},{\"internalType\":\"contractAny2EVMTollOnRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTokenAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFeeConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"AllowListSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCCIP.EVM2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OnRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVM2AnyTollMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"getRequiredFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_destinationChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"setAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setAllowlistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"}],\"internalType\":\"structAny2EVMTollOnRampInterface.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"setFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162003c8b38038062003c8b8339810160408190526200003491620006f0565b6000805460ff191681558a908a908a908a908a908a908a908a908a908a908390839088908b908b908a903390819081620000b55760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ef57620000ef81620003ec565b5050506001600160a01b0381166200011a57604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200015c5760405162d8548360e71b815260040160405180910390fd5b8151620001719060049060208501906200049d565b5060005b825181101562000253576000828281518110620001965762000196620007f8565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001e057620001e0620007f8565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600590925220805460ff191660011790556200024b816200080e565b905062000175565b5050815115905062000283576006805460ff191660011790558051620002819060079060208401906200049d565b505b60005b8151811015620002f057600160086000848481518110620002ab57620002ab620007f8565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055620002e8816200080e565b905062000286565b5050600980546001600160a01b039283166001600160a01b03199182161790915560408051608080820183528551808352602096870180518885018190529051848601819052426060909501859052600c92909255600d55600e55600f919091559d909d5260a09b909b528551601180549388015197909d01516001600160401b03908116600160801b02600160801b600160c01b031998821668010000000000000000026001600160801b031990951691909216179290921795909516179099555050601280549790911696909516959095179093555050601080546001600160401b031916905550620008369a5050505050505050505050565b336001600160a01b03821603620004465760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ac565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620004f5579160200282015b82811115620004f557825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620004be565b506200050392915062000507565b5090565b5b8082111562000503576000815560010162000508565b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146200054a57600080fd5b50565b600082601f8301126200055f57600080fd5b815160206001600160401b03808311156200057e576200057e6200051e565b8260051b604051601f19603f83011681018181108482111715620005a657620005a66200051e565b604052938452858101830193838101925087851115620005c557600080fd5b83870191505b84821015620005f1578151620005e18162000534565b83529183019190830190620005cb565b979650505050505050565b8051620006098162000534565b919050565b80516001600160401b03811681146200060957600080fd5b6000606082840312156200063957600080fd5b604051606081016001600160401b03811182821017156200065e576200065e6200051e565b6040529050806200066f836200060e565b81526200067f602084016200060e565b602082015262000692604084016200060e565b60408201525092915050565b600060408284031215620006b157600080fd5b604080519081016001600160401b0381118282101715620006d657620006d66200051e565b604052825181526020928301519281019290925250919050565b6000806000806000806000806000806101a08b8d0312156200071157600080fd5b8a5160208c015160408d0151919b5099506001600160401b03808211156200073857600080fd5b620007468e838f016200054d565b995060608d01519150808211156200075d57600080fd5b6200076b8e838f016200054d565b985060808d01519150808211156200078257600080fd5b50620007918d828e016200054d565b965050620007a260a08c01620005fc565b9450620007b38c60c08d0162000626565b9350620007c58c6101208d016200069e565b9250620007d66101608c01620005fc565b9150620007e76101808c01620005fc565b90509295989b9194979a5092959850565b634e487b7160e01b600052603260045260246000fd5b6000600182016200082f57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a0516134286200086360003960006105630152600081816103b3015261086401526134286000f3fe608060405234801561001057600080fd5b50600436106102ad5760003560e01c8063744b92e21161017b578063bbe4f6db116100d8578063d7644ba21161008c578063f2fde38b11610071578063f2fde38b146106c2578063f3fae9fa146106d5578063f78faa32146106e857600080fd5b8063d7644ba21461069c578063eb511dd4146106af57600080fd5b8063c3f909d4116100bd578063c3f909d4146105bc578063c5eff3d01461065e578063d0d5de611461067357600080fd5b8063bbe4f6db146102b2578063c0d78655146105a957600080fd5b806390c2339b1161012f578063a721719511610114578063a72171951461055e578063b0f479a114610585578063b4069b311461059657600080fd5b806390c2339b14610510578063918725431461054b57600080fd5b80638456cb59116101605780638456cb59146104ea57806389c06568146104f25780638da5cb5b146104fa57600080fd5b8063744b92e2146104cf57806379ba5097146104e257600080fd5b80634120fccd1161022957806359e96b5b116101dd5780635c975abb116101c25780635c975abb1461049c578063671dc337146104a7578063681fba16146104ba57600080fd5b806359e96b5b1461045d5780635b16ebb71461047057600080fd5b80634741062e1161020e5780634741062e14610419578063552b818b14610439578063599f64311461044c57600080fd5b80634120fccd146103fe5780634352fa9f1461040657600080fd5b8063181f5a771161028057806328094b591161026557806328094b59146103ae57806339aa9264146103e35780633f4ba83a146103f657600080fd5b8063181f5a77146103545780632222dd421461039d57600080fd5b806304c2a34a146102b257806305afe24a146102fb578063108ee5fc14610327578063147809b31461033c575b600080fd5b6102de6102c03660046128de565b6001600160a01b039081166000908152600360205260409020541690565b6040516001600160a01b0390911681526020015b60405180910390f35b61030e610309366004612ae7565b6106f3565b60405167ffffffffffffffff90911681526020016102f2565b61033a6103353660046128de565b610984565b005b610344610a3a565b60405190151581526020016102f2565b6103906040518060400160405280601781526020017f45564d3245564d546f6c6c4f6e52616d7020312e302e3000000000000000000081525081565b6040516102f29190612c41565b6002546001600160a01b03166102de565b6103d57f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016102f2565b61033a6103f13660046128de565b610ac7565b61033a610afe565b61030e610b10565b61033a610414366004612c54565b610b30565b61042c610427366004612cb8565b610d85565b6040516102f29190612d28565b61033a610447366004612d3b565b610e4d565b6009546001600160a01b03166102de565b61033a61046b366004612db0565b610fd7565b61034461047e3660046128de565b6001600160a01b031660009081526005602052604090205460ff1690565b60005460ff16610344565b61033a6104b5366004612df1565b61103b565b6104c261108d565b6040516102f29190612e42565b61033a6104dd366004612e55565b611152565b61033a61148a565b61033a61156d565b6104c261157d565b60005461010090046001600160a01b03166102de565b6105186115df565b6040516102f291908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b61033a610559366004612e8e565b611680565b6103d57f000000000000000000000000000000000000000000000000000000000000000081565b6012546001600160a01b03166102de565b6102de6105a43660046128de565b6117ac565b61033a6105b73660046128de565b61189a565b61062b6040805160608101825260008082526020820181905291810191909152506040805160608101825260115467ffffffffffffffff808216835268010000000000000000820481166020840152700100000000000000000000000000000000909104169181019190915290565b60408051825167ffffffffffffffff908116825260208085015182169083015292820151909216908201526060016102f2565b6106666118fd565b6040516102f29190612edd565b6103d56106813660046128de565b6001600160a01b031660009081526014602052604090205490565b61033a6106aa366004612f38565b61195d565b61033a6106bd366004612e55565b6119a6565b61033a6106d03660046128de565b611b89565b61033a6106e3366004612f55565b611b9d565b60065460ff16610344565b6000805460ff161561074c5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561079f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c39190612ffc565b156107f9576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6012546001600160a01b0316331461083d576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108568360200151518460400151856060015185611d4a565b6040805161014081019091527f000000000000000000000000000000000000000000000000000000000000000081526010805460009291602083019184906108a79067ffffffffffffffff16613048565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff168152602001846001600160a01b0316815260200185600001516001600160a01b0316815260200185602001518152602001856040015181526020018560600151815260200185608001516001600160a01b031681526020018560a0015181526020018560c0015181525090507fab2ca9da6d303be28d1a5e854e3e170be286e07696245e77f8ea11f55367d48181604051610971919061306f565b60405180910390a1602001519392505050565b61098c61204d565b6001600160a01b0381166109cc576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28910160405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa158015610a9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac19190612ffc565b15905090565b610acf61204d565b6009805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b610b0661204d565b610b0e6120ac565b565b601054600090610b2b9067ffffffffffffffff166001613156565b905090565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015610b6d57506009546001600160a01b03163314155b15610ba4576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610be0576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b5460005b81811015610c3a57600a6000600b8381548110610c0557610c05613182565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610c33816131b1565b9050610be6565b5060005b82811015610d6a576000858281518110610c5a57610c5a613182565b6020026020010151905060006001600160a01b0316816001600160a01b031603610cb0576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610cc257610cc2613182565b6020026020010151600a6000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610d2757610d27613182565b6020026020010151604051610d519291906001600160a01b03929092168252602082015260400190565b60405180910390a150610d63816131b1565b9050610c3e565b508351610d7e90600b9060208701906127d6565b5050505050565b80516060908067ffffffffffffffff811115610da357610da36128fb565b604051908082528060200260200182016040528015610dcc578160200160208202803683370190505b50915060005b81811015610e4657600a6000858381518110610df057610df0613182565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610e2b57610e2b613182565b6020908102919091010152610e3f816131b1565b9050610dd2565b5050919050565b610e5561204d565b60006007805480602002602001604051908101604052809291908181526020018280548015610ead57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e8f575b5050505050905060005b8151811015610f1a57600060086000848481518110610ed857610ed8613182565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055610f13816131b1565b9050610eb7565b50610f2760078484612844565b5060005b82811015610f9857600160086000868685818110610f4b57610f4b613182565b9050602002016020810190610f6091906128de565b6001600160a01b031681526020810191909152604001600020805460ff1916911515919091179055610f91816131b1565b9050610f2b565b507ff8adc5fee247b62a85f63294cb46e4be61da815e4756bc57672a83b24faf0dda8383604051610fca9291906131cb565b60405180910390a1505050565b610fdf61204d565b610ff36001600160a01b0384168383612148565b604080516001600160a01b038086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa890606001610fca565b61104361204d565b8060116110508282613224565b9050507fcc6ce9e57c1de2adf58a81e94b96b43d77ea6973e3f08e6ea4fe83d62ae60e9e816040516110829190613312565b60405180910390a150565b60045460609067ffffffffffffffff8111156110ab576110ab6128fb565b6040519080825280602002602001820160405280156110d4578160200160208202803683370190505b50905060005b60045481101561114e57611114600482815481106110fa576110fa613182565b6000918252602090912001546001600160a01b03166117ac565b82828151811061112657611126613182565b6001600160a01b0390921660209283029190910190910152611147816131b1565b90506110da565b5090565b61115a61204d565b6004546000819003611198576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290611226576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614611275576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004611284600185613365565b8154811061129457611294613182565b9060005260206000200160009054906101000a90046001600160a01b03169050600482602001516bffffffffffffffffffffffff16815481106112d9576112d9613182565b6000918252602090912001546001600160a01b031660046112fb600186613365565b8154811061130b5761130b613182565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600483602001516bffffffffffffffffffffffff168154811061135f5761135f613182565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff90921674010000000000000000000000000000000000000000029190921617905560048054806113e9576113e961337c565b600082815260208082206000199084018101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b038781168084526003835260408085208590559188168085526005845293829020805460ff191690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b031633146114e45760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610743565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61157561204d565b610b0e6121cd565b606060048054806020026020016040519081016040528092919081815260200182805480156115d557602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116115b7575b5050505050905090565b61160a6040518060800160405280600081526020016000815260200160008152602001600081525090565b60408051608081018252600c548152600d546020820152600e5491810191909152600f546060820181905242906000906116449083613365565b602084015184519192506116709161165c90846133ab565b856040015161166b91906133ca565b612255565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b0316141580156116bd57506009546001600160a01b03163314155b156116f4576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff11611748576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611752600c61226b565b6020810151600d8190558151600c55600e5461176e9190612255565b600e55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d9101611082565b6001600160a01b0380821660009081526003602052604081205490911680611800576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa15801561186f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061189391906133e2565b9392505050565b6118a261204d565b6012805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527fc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d1590602001611082565b606060078054806020026020016040519081016040528092919081815260200182805480156115d5576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116115b7575050505050905090565b61196561204d565b6006805460ff19168215159081179091556040519081527fccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df03290602001611082565b6119ae61204d565b6001600160a01b03821615806119cb57506001600160a01b038116155b15611a02576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015611a91576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01805473ffffffffffffffffffffffffffffffffffffffff1916841790558482526005815290859020805460ff19169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c9101610fca565b611b9161204d565b611b9a81612318565b50565b611ba561204d565b60208101515181515114611be5576040517f5601467a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b601354811015611c40576014600060138381548110611c0957611c09613182565b60009182526020808320909101546001600160a01b0316835282019290925260400181205580611c38816131b1565b915050611be8565b5060005b816020015151811015611d2f5760006001600160a01b031682602001518281518110611c7257611c72613182565b60200260200101516001600160a01b031603611cba576040517f5601467a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151805182908110611cce57611cce613182565b60200260200101516014600084602001518481518110611cf057611cf0613182565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020819055508080611d27906131b1565b915050611c44565b506020808201518051611d469260139201906127d6565b5050565b6012546001600160a01b0316611d8c576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038116611dcc576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60115468010000000000000000900467ffffffffffffffff16841115611e40576011546040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090910467ffffffffffffffff16600482015260248101859052604401610743565b8251601154700100000000000000000000000000000000900467ffffffffffffffff16811180611e71575082518114155b15611ea8576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60065460ff168015611ed357506001600160a01b03821660009081526008602052604090205460ff16155b15611f15576040517fd0d259760000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610743565b611f1f84846123d4565b60005b81811015612045576000858281518110611f3e57611f3e613182565b602002602001015190506000611f6c826001600160a01b039081166000908152600360205260409020541690565b90506001600160a01b038116611fb9576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610743565b806001600160a01b031663503c2858878581518110611fda57611fda613182565b60200260200101516040518263ffffffff1660e01b815260040161200091815260200190565b600060405180830381600087803b15801561201a57600080fd5b505af115801561202e573d6000803e3d6000fd5b5050505050508061203e906131b1565b9050611f22565b505050505050565b60005461010090046001600160a01b03163314610b0e5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610743565b60005460ff166120fe5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610743565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526121c8908490612562565b505050565b60005460ff16156122205760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610743565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861212b3390565b60008183106122645781611893565b5090919050565b60018101546002820154429114806122865750808260030154145b1561228f575050565b8160010154826002015411156122d1576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260030154826122e39190613365565b6001840154845491925061230a916122fb90846133ab565b856002015461166b91906133ca565b600284015550600390910155565b336001600160a01b038216036123705760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610743565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000805b83518110156124cb576000600a60008684815181106123f9576123f9613182565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020549050806000036124885784828151811061243e5761243e613182565b60200260200101516040517f9a655f7b00000000000000000000000000000000000000000000000000000000815260040161074391906001600160a01b0391909116815260200190565b83828151811061249a5761249a613182565b6020026020010151816124ad91906133ab565b6124b790846133ca565b925050806124c4906131b1565b90506123d8565b5080156121c8576124dc600c61226b565b600e54811115612518576040517f3bfa6f3800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600c600201600082825461252d9190613365565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba37890602001610fca565b60006125b7826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166126479092919063ffffffff16565b8051909150156121c857808060200190518101906125d59190612ffc565b6121c85760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610743565b6060612656848460008561265e565b949350505050565b6060824710156126d65760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610743565b843b6127245760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610743565b600080866001600160a01b0316858760405161274091906133ff565b60006040518083038185875af1925050503d806000811461277d576040519150601f19603f3d011682016040523d82523d6000602084013e612782565b606091505b509150915061279282828661279d565b979650505050505050565b606083156127ac575081611893565b8251156127bc5782518084602001fd5b8160405162461bcd60e51b81526004016107439190612c41565b828054828255906000526020600020908101928215612838579160200282015b82811115612838578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039091161782556020909201916001909101906127f6565b5061114e9291506128a4565b828054828255906000526020600020908101928215612838579160200282015b8281111561283857815473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03843516178255602090920191600190910190612864565b5b8082111561114e57600081556001016128a5565b6001600160a01b0381168114611b9a57600080fd5b80356128d9816128b9565b919050565b6000602082840312156128f057600080fd5b8135611893816128b9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff8111828210171561294d5761294d6128fb565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561297c5761297c6128fb565b604052919050565b600082601f83011261299557600080fd5b813567ffffffffffffffff8111156129af576129af6128fb565b6129c26020601f19601f84011601612953565b8181528460208386010111156129d757600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115612a0e57612a0e6128fb565b5060051b60200190565b600082601f830112612a2957600080fd5b81356020612a3e612a39836129f4565b612953565b82815260059290921b84018101918181019086841115612a5d57600080fd5b8286015b84811015612a81578035612a74816128b9565b8352918301918301612a61565b509695505050505050565b600082601f830112612a9d57600080fd5b81356020612aad612a39836129f4565b82815260059290921b84018101918181019086841115612acc57600080fd5b8286015b84811015612a815780358352918301918301612ad0565b60008060408385031215612afa57600080fd5b823567ffffffffffffffff80821115612b1257600080fd5b9084019060e08287031215612b2657600080fd5b612b2e61292a565b612b37836128ce565b8152602083013582811115612b4b57600080fd5b612b5788828601612984565b602083015250604083013582811115612b6f57600080fd5b612b7b88828601612a18565b604083015250606083013582811115612b9357600080fd5b612b9f88828601612a8c565b606083015250612bb1608084016128ce565b608082015260a083013560a082015260c083013560c0820152809450505050612bdc602084016128ce565b90509250929050565b60005b83811015612c00578181015183820152602001612be8565b83811115612c0f576000848401525b50505050565b60008151808452612c2d816020860160208601612be5565b601f01601f19169290920160200192915050565b6020815260006118936020830184612c15565b60008060408385031215612c6757600080fd5b823567ffffffffffffffff80821115612c7f57600080fd5b612c8b86838701612a18565b93506020850135915080821115612ca157600080fd5b50612cae85828601612a8c565b9150509250929050565b600060208284031215612cca57600080fd5b813567ffffffffffffffff811115612ce157600080fd5b61265684828501612a18565b600081518084526020808501945080840160005b83811015612d1d57815187529582019590820190600101612d01565b509495945050505050565b6020815260006118936020830184612ced565b60008060208385031215612d4e57600080fd5b823567ffffffffffffffff80821115612d6657600080fd5b818501915085601f830112612d7a57600080fd5b813581811115612d8957600080fd5b8660208260051b8501011115612d9e57600080fd5b60209290920196919550909350505050565b600080600060608486031215612dc557600080fd5b8335612dd0816128b9565b92506020840135612de0816128b9565b929592945050506040919091013590565b600060608284031215612e0357600080fd5b50919050565b600081518084526020808501945080840160005b83811015612d1d5781516001600160a01b031687529582019590820190600101612e1d565b6020815260006118936020830184612e09565b60008060408385031215612e6857600080fd5b8235612e73816128b9565b91506020830135612e83816128b9565b809150509250929050565b600060408284031215612ea057600080fd5b6040516040810181811067ffffffffffffffff82111715612ec357612ec36128fb565b604052823581526020928301359281019290925250919050565b6020808252825182820181905260009190848201906040850190845b81811015612f1e5783516001600160a01b031683529284019291840191600101612ef9565b50909695505050505050565b8015158114611b9a57600080fd5b600060208284031215612f4a57600080fd5b813561189381612f2a565b600060208284031215612f6757600080fd5b813567ffffffffffffffff80821115612f7f57600080fd5b9083019060408286031215612f9357600080fd5b604051604081018181108382111715612fae57612fae6128fb565b604052823582811115612fc057600080fd5b612fcc87828601612a8c565b825250602083013582811115612fe157600080fd5b612fed87828601612a18565b60208301525095945050505050565b60006020828403121561300e57600080fd5b815161189381612f2a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681810361306557613065613019565b6001019392505050565b602081528151602082015260006020830151613097604084018267ffffffffffffffff169052565b5060408301516001600160a01b03811660608401525060608301516001600160a01b03811660808401525060808301516101408060a08501526130de610160850183612c15565b915060a0850151601f19808685030160c08701526130fc8483612e09565b935060c08701519150808685030160e08701525061311a8382612ced565b92505060e0850151610100613139818701836001600160a01b03169052565b860151610120868101919091529095015193019290925250919050565b600067ffffffffffffffff80831681851680830382111561317957613179613019565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060001982036131c4576131c4613019565b5060010190565b60208082528181018390526000908460408401835b86811015612a815782356131f3816128b9565b6001600160a01b0316825291830191908301906001016131e0565b67ffffffffffffffff81168114611b9a57600080fd5b813561322f8161320e565b67ffffffffffffffff811690508154817fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000821617835560208401356132738161320e565b6fffffffffffffffff00000000000000008160401b16905080837fffffffffffffffffffffffffffffffff0000000000000000000000000000000084161717845560408501356132c28161320e565b77ffffffffffffffff000000000000000000000000000000008160801b16847fffffffffffffffff0000000000000000000000000000000000000000000000008516178317178555505050505050565b6060810182356133218161320e565b67ffffffffffffffff908116835260208401359061333e8261320e565b90811660208401526040840135906133558261320e565b8082166040850152505092915050565b60008282101561337757613377613019565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008160001904831182151516156133c5576133c5613019565b500290565b600082198211156133dd576133dd613019565b500190565b6000602082840312156133f457600080fd5b8151611893816128b9565b60008251613411818460208701612be5565b919091019291505056fea164736f6c634300080f000a",
}

var EVM2EVMTollOnRampABI = EVM2EVMTollOnRampMetaData.ABI

var EVM2EVMTollOnRampBin = EVM2EVMTollOnRampMetaData.Bin

func DeployEVM2EVMTollOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, chainId *big.Int, destinationChainId *big.Int, tokens []common.Address, pools []common.Address, allowlist []common.Address, afn common.Address, config BaseOnRampInterfaceOnRampConfig, rateLimiterConfig AggregateRateLimiterInterfaceRateLimiterConfig, tokenLimitsAdmin common.Address, router common.Address) (common.Address, *types.Transaction, *EVM2EVMTollOnRamp, error) {
	parsed, err := EVM2EVMTollOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMTollOnRampBin), backend, chainId, destinationChainId, tokens, pools, allowlist, afn, config, rateLimiterConfig, tokenLimitsAdmin, router)
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(AggregateRateLimiterInterfaceTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(AggregateRateLimiterInterfaceTokenBucket)).(*AggregateRateLimiterInterfaceTokenBucket)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMTollOnRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMTollOnRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOnRamp.CallOpts)
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.GetPricesForTokens(&_EVM2EVMTollOnRamp.CallOpts, tokens)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.GetPricesForTokens(&_EVM2EVMTollOnRamp.CallOpts, tokens)
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMTollOnRamp.CallOpts)
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) IChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "i_chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) IChainId() (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.IChainId(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) IChainId() (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.IChainId(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) IDestinationChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "i_destinationChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) IDestinationChainId() (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.IDestinationChainId(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) IDestinationChainId() (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.IDestinationChainId(&_EVM2EVMTollOnRamp.CallOpts)
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetFeeConfig(opts *bind.TransactOpts, feeConfig Any2EVMTollOnRampInterfaceFeeConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setFeeConfig", feeConfig)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetFeeConfig(feeConfig Any2EVMTollOnRampInterfaceFeeConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetFeeConfig(&_EVM2EVMTollOnRamp.TransactOpts, feeConfig)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetFeeConfig(feeConfig Any2EVMTollOnRampInterfaceFeeConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetFeeConfig(&_EVM2EVMTollOnRamp.TransactOpts, feeConfig)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetPrices(&_EVM2EVMTollOnRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetPrices(&_EVM2EVMTollOnRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetRateLimiterConfig(&_EVM2EVMTollOnRamp.TransactOpts, config)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetRateLimiterConfig(&_EVM2EVMTollOnRamp.TransactOpts, config)
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMTollOnRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMTollOnRamp.TransactOpts, newAdmin)
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
	Message CCIPEVM2EVMTollMessage
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

type EVM2EVMTollOnRampConfigChangedIterator struct {
	Event *EVM2EVMTollOnRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampConfigChanged)
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
		it.Event = new(EVM2EVMTollOnRampConfigChanged)
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

func (it *EVM2EVMTollOnRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMTollOnRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampConfigChangedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampConfigChanged)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMTollOnRampConfigChanged, error) {
	event := new(EVM2EVMTollOnRampConfigChanged)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

type EVM2EVMTollOnRampTokenPriceChangedIterator struct {
	Event *EVM2EVMTollOnRampTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampTokenPriceChanged)
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
		it.Event = new(EVM2EVMTollOnRampTokenPriceChanged)
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

func (it *EVM2EVMTollOnRampTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMTollOnRampTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampTokenPriceChangedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampTokenPriceChanged)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMTollOnRampTokenPriceChanged, error) {
	event := new(EVM2EVMTollOnRampTokenPriceChanged)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMTollOnRampTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampTokensRemovedFromBucket)
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
		it.Event = new(EVM2EVMTollOnRampTokensRemovedFromBucket)
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

func (it *EVM2EVMTollOnRampTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMTollOnRampTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampTokensRemovedFromBucketIterator{contract: _EVM2EVMTollOnRamp.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampTokensRemovedFromBucket)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMTollOnRampTokensRemovedFromBucket, error) {
	event := new(EVM2EVMTollOnRampTokensRemovedFromBucket)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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
	case _EVM2EVMTollOnRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMTollOnRamp.ParseConfigChanged(log)
	case _EVM2EVMTollOnRamp.abi.Events["FeeCharged"].ID:
		return _EVM2EVMTollOnRamp.ParseFeeCharged(log)
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
	case _EVM2EVMTollOnRamp.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMTollOnRamp.ParseTokenPriceChanged(log)
	case _EVM2EVMTollOnRamp.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMTollOnRamp.ParseTokensRemovedFromBucket(log)
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

func (EVM2EVMTollOnRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMTollOnRampFeeCharged) Topic() common.Hash {
	return common.HexToHash("0x945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d")
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

func (EVM2EVMTollOnRampTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMTollOnRampTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMTollOnRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRamp) Address() common.Address {
	return _EVM2EVMTollOnRamp.address
}

type EVM2EVMTollOnRampInterface interface {
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetAllowlist(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowlistEnabled(opts *bind.CallOpts) (bool, error)

	GetConfig(opts *bind.CallOpts) (BaseOnRampInterfaceOnRampConfig, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)

	GetRequiredFee(opts *bind.CallOpts, feeToken common.Address) (*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetTokenPool(opts *bind.CallOpts, token common.Address) (common.Address, error)

	IChainId(opts *bind.CallOpts) (*big.Int, error)

	IDestinationChainId(opts *bind.CallOpts) (*big.Int, error)

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

	IsPool(opts *bind.CallOpts, addr common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, message CCIPEVM2AnyTollMessage, originalSender common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error)

	SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error)

	SetFeeConfig(opts *bind.TransactOpts, feeConfig Any2EVMTollOnRampInterfaceFeeConfig) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

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

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMTollOnRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMTollOnRampConfigChanged, error)

	FilterFeeCharged(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeeChargedIterator, error)

	WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeeCharged) (event.Subscription, error)

	ParseFeeCharged(log types.Log) (*EVM2EVMTollOnRampFeeCharged, error)

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

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMTollOnRampTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMTollOnRampTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMTollOnRampTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMTollOnRampTokensRemovedFromBucket, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOnRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMTollOnRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
