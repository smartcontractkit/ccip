// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdc_token_pool

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
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
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
	_ = abi.ConvertType
)

type RateLimiterConfig struct {
	IsEnabled bool
	Capacity  *big.Int
	Rate      *big.Int
}

type RateLimiterTokenBucket struct {
	Tokens      *big.Int
	LastUpdated uint32
	IsEnabled   bool
	Capacity    *big.Int
	Rate        *big.Int
}

type TokenPoolRampUpdate struct {
	Ramp              common.Address
	Allowed           bool
	RateLimiterConfig RateLimiterConfig
}

type USDCTokenPoolDomain struct {
	AllowedCaller    [32]byte
	DomainIdentifier uint32
	Enabled          bool
}

type USDCTokenPoolDomainUpdate struct {
	AllowedCaller     [32]byte
	DomainIdentifier  uint32
	DestChainSelector uint64
	Enabled           bool
}

var USDCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractITokenMessenger\",\"name\":\"tokenMessenger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowListNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadARMSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"expected\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"got\",\"type\":\"uint32\"}],\"name\":\"InvalidDestinationDomain\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"allowedCaller\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"domainIdentifier\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structUSDCTokenPool.DomainUpdate\",\"name\":\"domain\",\"type\":\"tuple\"}],\"name\":\"InvalidDomain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"}],\"name\":\"InvalidMessageVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expected\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"}],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"expected\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"got\",\"type\":\"uint32\"}],\"name\":\"InvalidSourceDomain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"}],\"name\":\"InvalidTokenMessengerVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ramp\",\"type\":\"address\"}],\"name\":\"NonExistentRamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ramp\",\"type\":\"address\"}],\"name\":\"RampAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"domain\",\"type\":\"uint64\"}],\"name\":\"UnknownDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnlockingUSDCFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenMessenger\",\"type\":\"address\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"allowedCaller\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"domainIdentifier\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"name\":\"DomainsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"OffRampAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"OnRampAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"OnRampConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SUPPORTED_USDC_VERSION\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"ramp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"internalType\":\"structTokenPool.RampUpdate[]\",\"name\":\"onRamps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ramp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"internalType\":\"structTokenPool.RampUpdate[]\",\"name\":\"offRamps\",\"type\":\"tuple[]\"}],\"name\":\"applyRampUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"currentOffRampRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"currentOnRampRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getArmProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"getDomain\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"allowedCaller\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"domainIdentifier\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structUSDCTokenPool.Domain\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRamps\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOnRamps\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUSDCInterfaceId\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_localDomainIdentifier\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_messageTransmitter\",\"outputs\":[{\"internalType\":\"contractIMessageTransmitter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_tokenMessenger\",\"outputs\":[{\"internalType\":\"contractITokenMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"isOnRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"destinationReceiver\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"lockOrBurn\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"releaseOrMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"allowedCaller\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"domainIdentifier\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"name\":\"domains\",\"type\":\"tuple[]\"}],\"name\":\"setDomains\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setOffRampRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setOnRampRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b50604051620041b4380380620041b4833981016040819052620000359162000809565b82828233806000816200008f5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c257620000c2816200041f565b5050506001600160a01b038316620000ed576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b03808416608052811660a052815115801560c0526200012857604080516000815260208101909152620001289083620004ca565b5050506001600160a01b03841662000153576040516306b7c75960e31b815260040160405180910390fd5b6000846001600160a01b0316632c1219216040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000194573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001ba91906200091c565b90506000816001600160a01b03166354fd4d506040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001fd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000223919062000943565b905063ffffffff811615620002ba57816001600160a01b03166354fd4d506040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000271573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000297919062000943565b6040516334697c6b60e11b815263ffffffff909116600482015260240162000086565b6000866001600160a01b0316639cdbb1816040518163ffffffff1660e01b81526004016020604051808303816000875af1158015620002fd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000323919062000943565b905063ffffffff81161562000354576040516316ba39c560e31b815263ffffffff8216600482015260240162000086565b6001600160a01b0380881660e05283166101008190526040805163234d8e3d60e21b81529051638d3638f4916004808201926020929091908290030181865afa158015620003a6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620003cc919062000943565b63ffffffff16610120526040516001600160a01b03881681527f2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea95449060200160405180910390a150505050505050620009df565b336001600160a01b03821603620004795760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000086565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60c051620004eb576040516335f4a7b360e01b815260040160405180910390fd5b60005b8251811015620005805760008382815181106200050f576200050f6200096b565b60209081029190910101519050620005296002826200063b565b156200056c576040516001600160a01b03821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b50620005788162000997565b9050620004ee565b5060005b815181101562000636576000828281518110620005a557620005a56200096b565b6020026020010151905060006001600160a01b0316816001600160a01b031603620005d1575062000623565b620005de6002826200065b565b1562000621576040516001600160a01b03821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b6200062e8162000997565b905062000584565b505050565b600062000652836001600160a01b03841662000672565b90505b92915050565b600062000652836001600160a01b03841662000776565b600081815260018301602052604081205480156200076b57600062000699600183620009b3565b8554909150600090620006af90600190620009b3565b90508181146200071b576000866000018281548110620006d357620006d36200096b565b9060005260206000200154905080876000018481548110620006f957620006f96200096b565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806200072f576200072f620009c9565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000655565b600091505062000655565b6000818152600183016020526040812054620007bf5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000655565b50600062000655565b6001600160a01b0381168114620007de57600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b80516200080481620007c8565b919050565b600080600080608085870312156200082057600080fd5b84516200082d81620007c8565b809450506020808601516200084281620007c8565b60408701519094506001600160401b03808211156200086057600080fd5b818801915088601f8301126200087557600080fd5b8151818111156200088a576200088a620007e1565b8060051b604051601f19603f83011681018181108582111715620008b257620008b2620007e1565b60405291825284820192508381018501918b831115620008d157600080fd5b938501935b82851015620008fa57620008ea85620007f7565b84529385019392850192620008d6565b8097505050505050506200091160608601620007f7565b905092959194509250565b6000602082840312156200092f57600080fd5b81516200093c81620007c8565b9392505050565b6000602082840312156200095657600080fd5b815163ffffffff811681146200093c57600080fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201620009ac57620009ac62000981565b5060010190565b8181038181111562000655576200065562000981565b634e487b7160e01b600052603160045260246000fd5b60805160a05160c05160e051610100516101205161373162000a836000396000818161030d0152818161101a015281816118ba01526119180152600081816105950152610bce0152600081816102e601528181610e9c0152610f5301526000818161055901528181610d27015261139e015260006102aa01526000818161026301528181610e7a01528181610f20015281816117de01526119d601526137316000f3fe608060405234801561001057600080fd5b50600436106101c35760003560e01c80638627fad6116100f9578063b3a3fb4111610097578063dfadfa3511610071578063dfadfa35146104b9578063e0351e1314610557578063f2fde38b1461057d578063fbf84dd71461059057600080fd5b8063b3a3fb4114610480578063c49907b514610493578063d612b945146104a657600080fd5b806396875445116100d357806396875445146104555780639fdf13ff14610468578063a40e69c714610470578063a7cd63b71461047857600080fd5b80638627fad61461040f57806387381314146104225780638da5cb5b1461043757600080fd5b80636155cda0116101665780636f32b872116101405780636f32b872146103725780637448b3c7146103855780637787e7ab1461039857806379ba50971461040757600080fd5b80636155cda0146102e15780636b716b0d146103085780636d1081391461034457600080fd5b80631d7a74a0116101a25780631d7a74a01461024e57806321df0da7146102615780635246492f146102a857806354c8a4f3146102ce57600080fd5b806241d3c1146101c857806301ffc9a7146101dd578063181f5a7714610205575b600080fd5b6101db6101d6366004612a7a565b6105b7565b005b6101f06101eb366004612aef565b61075e565b60405190151581526020015b60405180910390f35b6102416040518060400160405280601381526020017f55534443546f6b656e506f6f6c20312e322e300000000000000000000000000081525081565b6040516101fc9190612b9f565b6101f061025c366004612bdb565b6107ba565b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101fc565b7f0000000000000000000000000000000000000000000000000000000000000000610283565b6101db6102dc366004612c42565b6107c7565b6102837f000000000000000000000000000000000000000000000000000000000000000081565b61032f7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101fc565b6040517fd6aca1be0000000000000000000000000000000000000000000000000000000081526020016101fc565b6101f0610380366004612bdb565b610842565b6101db610393366004612df7565b61084f565b6103ab6103a6366004612bdb565b61090e565b6040516101fc919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b6101db6109ec565b6101db61041d366004612edd565b610ae9565b61042a610cd2565b6040516101fc9190612f70565b60005473ffffffffffffffffffffffffffffffffffffffff16610283565b61024161046336600461300c565b610ce3565b61032f600081565b61042a611072565b61042a61107e565b6103ab61048e366004612bdb565b61108a565b6101db6104a13660046130f0565b611168565b6101db6104b4366004612df7565b61117c565b61052d6104c7366004613150565b60408051606080820183526000808352602080840182905292840181905267ffffffffffffffff949094168452600a82529282902082519384018352805484526001015463ffffffff811691840191909152640100000000900460ff1615159082015290565b604080518251815260208084015163ffffffff1690820152918101511515908201526060016101fc565b7f00000000000000000000000000000000000000000000000000000000000000006101f0565b6101db61058b366004612bdb565b61123b565b6102837f000000000000000000000000000000000000000000000000000000000000000081565b6105bf61124f565b60005b818110156107205760008383838181106105de576105de61316d565b9050608002018036038101906105f491906131ae565b805190915015806106115750604081015167ffffffffffffffff16155b1561068057604080517fa087bd2900000000000000000000000000000000000000000000000000000000815282516004820152602083015163ffffffff1660248201529082015167ffffffffffffffff1660448201526060820151151560648201526084015b60405180910390fd5b60408051606080820183528351825260208085015163ffffffff9081168285019081529286015115158486019081529585015167ffffffffffffffff166000908152600a90925293902091518255516001909101805493511515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000909416919092161791909117905561071981613259565b90506105c2565b507f1889010d2535a0ab1643678d1da87fbbe8b87b2f585b47ddb72ec622aef9ee568282604051610752929190613291565b60405180910390a15050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167fd6aca1be0000000000000000000000000000000000000000000000000000000014806107b457506107b4826112d2565b92915050565b60006107b460078361136a565b6107cf61124f565b61083c8484808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152505060408051602080880282810182019093528782529093508792508691829185019084908082843760009201919091525061139c92505050565b50505050565b60006107b460048361136a565b61085761124f565b61086082610842565b6108ae576040517f498f12f600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610677565b73ffffffffffffffffffffffffffffffffffffffff821660009081526006602052604090206108dd9082611567565b7f578db78e348076074dbff64a94073a83e9a65aa6766b8c75fdc89282b0e30ed6828260405161075292919061331a565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915273ffffffffffffffffffffffffffffffffffffffff8216600090815260066020908152604091829020825160a08101845281546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff1615159482019490945260019091015480841660608301529190910490911660808201526107b490611716565b60015473ffffffffffffffffffffffffffffffffffffffff163314610a6d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610677565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610af2336107ba565b610b28576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b31836117c8565b60008082806020019051810190610b4891906133b7565b91509150600082806020019051810190610b62919061341b565b9050600082806020019051810190610b7a919061345c565b9050610b8a816000015183611802565b805160208201516040517f57ecfd2800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016926357ecfd2892610c01926004016134ed565b6020604051808303816000875af1158015610c20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c44919061351b565b610c7a576040517fbf969f2200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405187815273ffffffffffffffffffffffffffffffffffffffff89169033907f9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f09060200160405180910390a3505050505050505050565b6060610cde60046119b3565b905090565b6060610cee33610842565b610d24576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b877f00000000000000000000000000000000000000000000000000000000000000008015610d5a5750610d5860028261136a565b155b15610da9576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610677565b67ffffffffffffffff85166000908152600a602090815260409182902082516060810184528154815260019091015463ffffffff81169282019290925264010000000090910460ff16151591810182905290610e3d576040517fd201c48a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff87166004820152602401610677565b610e46876119c0565b6000610e556020828b8d613538565b610e5e91613562565b9050610ec173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167f00000000000000000000000000000000000000000000000000000000000000008a6119fa565b602082015182516040517ff856ddb6000000000000000000000000000000000000000000000000000000008152600481018b905263ffffffff90921660248301526044820183905273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116606484015260848301919091526000917f00000000000000000000000000000000000000000000000000000000000000009091169063f856ddb69060a4016020604051808303816000875af1158015610f9e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fc2919061359e565b6040518a815290915033907f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df79060200160405180910390a260408051808201825267ffffffffffffffff9290921680835263ffffffff7f00000000000000000000000000000000000000000000000000000000000000008116602094850190815283519485019290925290511682820152805180830382018152606090920190529b9a5050505050505050505050565b6060610cde60076119b3565b6060610cde60026119b3565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915273ffffffffffffffffffffffffffffffffffffffff8216600090815260096020908152604091829020825160a08101845281546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff1615159482019490945260019091015480841660608301529190910490911660808201526107b490611716565b61117061124f565b61083c84848484611bb3565b61118461124f565b61118d826107ba565b6111db576040517f498f12f600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610677565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260096020526040902061120a9082611567565b7fb3ba339cfbb8ef80d7a29ce5493051cb90e64fcfa85d7124efc1adfa4c68399f828260405161075292919061331a565b61124361124f565b61124c81612163565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146112d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610677565b565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f317fa3340000000000000000000000000000000000000000000000000000000014806107b457507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a7000000000000000000000000000000000000000000000000000000001492915050565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018301602052604081205415155b9392505050565b7f00000000000000000000000000000000000000000000000000000000000000006113f3576040517f35f4a7b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b82518110156114915760008382815181106114135761141361316d565b6020026020010151905061143181600261225890919063ffffffff16565b156114805760405173ffffffffffffffffffffffffffffffffffffffff821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b5061148a81613259565b90506113f6565b5060005b81518110156115625760008282815181106114b2576114b261316d565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036114f65750611552565b61150160028261227a565b156115505760405173ffffffffffffffffffffffffffffffffffffffff821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b61155b81613259565b9050611495565b505050565b815460009061159090700100000000000000000000000000000000900463ffffffff16426135bb565b9050801561163257600183015483546115d8916fffffffffffffffffffffffffffffffff8082169281169185917001000000000000000000000000000000009091041661229c565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b60208201518354611658916fffffffffffffffffffffffffffffffff90811691166122c6565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19906117099084906135ce565b60405180910390a1505050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526117a482606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff164261178891906135bb565b85608001516fffffffffffffffffffffffffffffffff1661229c565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b33600090815260096020526040902061124c90827f00000000000000000000000000000000000000000000000000000000000000006122dc565b600482015163ffffffff81161561184d576040517f68d2f8d600000000000000000000000000000000000000000000000000000000815263ffffffff82166004820152602401610677565b6008830151600c8401516014850151602085015163ffffffff8085169116146118b85760208501516040517fe366a11700000000000000000000000000000000000000000000000000000000815263ffffffff91821660048201529084166024820152604401610677565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168263ffffffff161461194d576040517f77e4802600000000000000000000000000000000000000000000000000000000815263ffffffff7f00000000000000000000000000000000000000000000000000000000000000008116600483015283166024820152604401610677565b845167ffffffffffffffff8281169116146119ab5784516040517ff917ffea00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff91821660048201529082166024820152604401610677565b505050505050565b606060006113958361265f565b33600090815260066020526040902061124c90827f00000000000000000000000000000000000000000000000000000000000000006122dc565b801580611a9a57506040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff838116602483015284169063dd62ed3e90604401602060405180830381865afa158015611a74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a98919061360a565b155b611b26576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608401610677565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b3000000000000000000000000000000000000000000000000000000001790526115629084906126bb565b611bbb61124f565b60005b83811015611ed8576000858583818110611bda57611bda61316d565b905060a00201803603810190611bf09190613623565b9050806020015115611dc8578051611c0a9060049061227a565b15611d7b576040805160a08101825282820180516020908101516fffffffffffffffffffffffffffffffff908116845263ffffffff4281168386019081528451511515868801908152855185015184166060880190815286518901518516608089019081528a5173ffffffffffffffffffffffffffffffffffffffff1660009081526006909752958990209751885493519251151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff939095167001000000000000000000000000000000009081027fffffffffffffffffffffffff0000000000000000000000000000000000000000909516918716919091179390931791909116929092178655905192518216029116176001909201919091558251905191517f0b594bb0555ff7b252e0c789ccc9d8903fec294172064308727d570505cee1ac92611d6e929161331a565b60405180910390a1611ec7565b80516040517fd3eb6bc500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610677565b8051611dd690600490612258565b15611e7a57805173ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604080822080547fffffffffffffffffffffff00000000000000000000000000000000000000000016815560010191909155815190517f7fd064821314ad863a0714a3f1229375ace6b6427ed5544b7b2ba1c47b1b529491611d6e9173ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b80516040517f498f12f600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610677565b50611ed181613259565b9050611bbe565b5060005b8181101561215c576000838383818110611ef857611ef861316d565b905060a00201803603810190611f0e9190613623565b9050806020015115612099578051611f289060079061227a565b15611d7b576040805160a08101825282820180516020908101516fffffffffffffffffffffffffffffffff908116845263ffffffff4281168386019081528451511515868801908152855185015184166060880190815286518901518516608089019081528a5173ffffffffffffffffffffffffffffffffffffffff1660009081526009909752958990209751885493519251151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff939095167001000000000000000000000000000000009081027fffffffffffffffffffffffff0000000000000000000000000000000000000000909516918716919091179390931791909116929092178655905192518216029116176001909201919091558251905191517f395b7374909d2b54e5796f53c898ebf41d767c86c78ea86519acf2b805852d889261208c929161331a565b60405180910390a161214b565b80516120a790600790612258565b15611e7a57805173ffffffffffffffffffffffffffffffffffffffff1660009081526009602052604080822080547fffffffffffffffffffffff00000000000000000000000000000000000000000016815560010191909155815190517fcf91daec21e3510e2f2aea4b09d08c235d5c6844980be709f282ef591dbf420c9161208c9173ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b5061215581613259565b9050611edc565b5050505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036121e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610677565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006113958373ffffffffffffffffffffffffffffffffffffffff84166127c7565b60006113958373ffffffffffffffffffffffffffffffffffffffff84166128ba565b60006122bb856122ac8486613674565b6122b6908761368b565b6122c6565b90505b949350505050565b60008183106122d55781611395565b5090919050565b825474010000000000000000000000000000000000000000900460ff161580612303575081155b1561230d57505050565b825460018401546fffffffffffffffffffffffffffffffff8083169291169060009061235390700100000000000000000000000000000000900463ffffffff16426135bb565b905080156124135781831115612395576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018601546123cf9083908590849070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1661229c565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff167001000000000000000000000000000000004263ffffffff160217875592505b848210156124ca5773ffffffffffffffffffffffffffffffffffffffff8416612472576040517ff94ebcd10000000000000000000000000000000000000000000000000000000081526004810183905260248101869052604401610677565b6040517f1a76572a000000000000000000000000000000000000000000000000000000008152600481018390526024810186905273ffffffffffffffffffffffffffffffffffffffff85166044820152606401610677565b848310156125dd5760018681015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1690600090829061250e90826135bb565b612518878a6135bb565b612522919061368b565b61252c919061369e565b905073ffffffffffffffffffffffffffffffffffffffff8616612585576040517f15279c080000000000000000000000000000000000000000000000000000000081526004810182905260248101869052604401610677565b6040517fd0c8d23a000000000000000000000000000000000000000000000000000000008152600481018290526024810186905273ffffffffffffffffffffffffffffffffffffffff87166044820152606401610677565b6125e785846135bb565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b6060816000018054806020026020016040519081016040528092919081815260200182805480156126af57602002820191906000526020600020905b81548152602001906001019080831161269b575b50505050509050919050565b600061271d826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166129099092919063ffffffff16565b805190915015611562578080602001905181019061273b919061351b565b611562576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610677565b600081815260018301602052604081205480156128b05760006127eb6001836135bb565b85549091506000906127ff906001906135bb565b905081811461286457600086600001828154811061281f5761281f61316d565b90600052602060002001549050808760000184815481106128425761284261316d565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080612875576128756136d9565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506107b4565b60009150506107b4565b6000818152600183016020526040812054612901575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556107b4565b5060006107b4565b60606122be8484600085856000808673ffffffffffffffffffffffffffffffffffffffff16858760405161293d9190613708565b60006040518083038185875af1925050503d806000811461297a576040519150601f19603f3d011682016040523d82523d6000602084013e61297f565b606091505b50915091506129908783838761299b565b979650505050505050565b60608315612a31578251600003612a2a5773ffffffffffffffffffffffffffffffffffffffff85163b612a2a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610677565b50816122be565b6122be8383815115612a465781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106779190612b9f565b60008060208385031215612a8d57600080fd5b823567ffffffffffffffff80821115612aa557600080fd5b818501915085601f830112612ab957600080fd5b813581811115612ac857600080fd5b8660208260071b8501011115612add57600080fd5b60209290920196919550909350505050565b600060208284031215612b0157600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461139557600080fd5b60005b83811015612b4c578181015183820152602001612b34565b50506000910152565b60008151808452612b6d816020860160208601612b31565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006113956020830184612b55565b803573ffffffffffffffffffffffffffffffffffffffff81168114612bd657600080fd5b919050565b600060208284031215612bed57600080fd5b61139582612bb2565b60008083601f840112612c0857600080fd5b50813567ffffffffffffffff811115612c2057600080fd5b6020830191508360208260051b8501011115612c3b57600080fd5b9250929050565b60008060008060408587031215612c5857600080fd5b843567ffffffffffffffff80821115612c7057600080fd5b612c7c88838901612bf6565b90965094506020870135915080821115612c9557600080fd5b50612ca287828801612bf6565b95989497509550505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715612d0057612d00612cae565b60405290565b6040805190810167ffffffffffffffff81118282101715612d0057612d00612cae565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612d7057612d70612cae565b604052919050565b801515811461124c57600080fd5b80356fffffffffffffffffffffffffffffffff81168114612bd657600080fd5b600060608284031215612db857600080fd5b612dc0612cdd565b90508135612dcd81612d78565b8152612ddb60208301612d86565b6020820152612dec60408301612d86565b604082015292915050565b60008060808385031215612e0a57600080fd5b612e1383612bb2565b9150612e228460208501612da6565b90509250929050565b600067ffffffffffffffff821115612e4557612e45612cae565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112612e8257600080fd5b8135612e95612e9082612e2b565b612d29565b818152846020838601011115612eaa57600080fd5b816020850160208301376000918101602001919091529392505050565b67ffffffffffffffff8116811461124c57600080fd5b600080600080600060a08688031215612ef557600080fd5b853567ffffffffffffffff80821115612f0d57600080fd5b612f1989838a01612e71565b9650612f2760208901612bb2565b95506040880135945060608801359150612f4082612ec7565b90925060808701359080821115612f5657600080fd5b50612f6388828901612e71565b9150509295509295909350565b6020808252825182820181905260009190848201906040850190845b81811015612fbe57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101612f8c565b50909695505050505050565b60008083601f840112612fdc57600080fd5b50813567ffffffffffffffff811115612ff457600080fd5b602083019150836020828501011115612c3b57600080fd5b600080600080600080600060a0888a03121561302757600080fd5b61303088612bb2565b9650602088013567ffffffffffffffff8082111561304d57600080fd5b6130598b838c01612fca565b909850965060408a0135955060608a0135915061307582612ec7565b9093506080890135908082111561308b57600080fd5b506130988a828b01612fca565b989b979a50959850939692959293505050565b60008083601f8401126130bd57600080fd5b50813567ffffffffffffffff8111156130d557600080fd5b60208301915083602060a083028501011115612c3b57600080fd5b6000806000806040858703121561310657600080fd5b843567ffffffffffffffff8082111561311e57600080fd5b61312a888389016130ab565b9096509450602087013591508082111561314357600080fd5b50612ca2878288016130ab565b60006020828403121561316257600080fd5b813561139581612ec7565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b63ffffffff8116811461124c57600080fd5b6000608082840312156131c057600080fd5b6040516080810181811067ffffffffffffffff821117156131e3576131e3612cae565b6040528235815260208301356131f88161319c565b6020820152604083013561320b81612ec7565b6040820152606083013561321e81612d78565b60608201529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361328a5761328a61322a565b5060010190565b6020808252818101839052600090604080840186845b8781101561330d5781358352848201356132c08161319c565b63ffffffff1683860152818401356132d781612ec7565b67ffffffffffffffff16838501526060828101356132f481612d78565b15159084015260809283019291909101906001016132a7565b5090979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff831681526080810161139560208301848051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b600082601f83011261338357600080fd5b8151613391612e9082612e2b565b8181528460208386010111156133a657600080fd5b6122be826020830160208701612b31565b600080604083850312156133ca57600080fd5b825167ffffffffffffffff808211156133e257600080fd5b6133ee86838701613372565b9350602085015191508082111561340457600080fd5b5061341185828601613372565b9150509250929050565b60006040828403121561342d57600080fd5b613435612d06565b825161344081612ec7565b815260208301516134508161319c565b60208201529392505050565b60006020828403121561346e57600080fd5b815167ffffffffffffffff8082111561348657600080fd5b908301906040828603121561349a57600080fd5b6134a2612d06565b8251828111156134b157600080fd5b6134bd87828601613372565b8252506020830151828111156134d257600080fd5b6134de87828601613372565b60208301525095945050505050565b6040815260006135006040830185612b55565b82810360208401526135128185612b55565b95945050505050565b60006020828403121561352d57600080fd5b815161139581612d78565b6000808585111561354857600080fd5b8386111561355557600080fd5b5050820193919092039150565b803560208310156107b4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602084900360031b1b1692915050565b6000602082840312156135b057600080fd5b815161139581612ec7565b818103818111156107b4576107b461322a565b606081016107b482848051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b60006020828403121561361c57600080fd5b5051919050565b600060a0828403121561363557600080fd5b61363d612cdd565b61364683612bb2565b8152602083013561365681612d78565b60208201526136688460408501612da6565b60408201529392505050565b80820281158282048414176107b4576107b461322a565b808201808211156107b4576107b461322a565b6000826136d4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000825161371a818460208701612b31565b919091019291505056fea164736f6c6343000813000a",
}

var USDCTokenPoolABI = USDCTokenPoolMetaData.ABI

var USDCTokenPoolBin = USDCTokenPoolMetaData.Bin

func DeployUSDCTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger common.Address, token common.Address, allowlist []common.Address, armProxy common.Address) (common.Address, *types.Transaction, *USDCTokenPool, error) {
	parsed, err := USDCTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(USDCTokenPoolBin), backend, tokenMessenger, token, allowlist, armProxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &USDCTokenPool{USDCTokenPoolCaller: USDCTokenPoolCaller{contract: contract}, USDCTokenPoolTransactor: USDCTokenPoolTransactor{contract: contract}, USDCTokenPoolFilterer: USDCTokenPoolFilterer{contract: contract}}, nil
}

type USDCTokenPool struct {
	address common.Address
	abi     abi.ABI
	USDCTokenPoolCaller
	USDCTokenPoolTransactor
	USDCTokenPoolFilterer
}

type USDCTokenPoolCaller struct {
	contract *bind.BoundContract
}

type USDCTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type USDCTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type USDCTokenPoolSession struct {
	Contract     *USDCTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type USDCTokenPoolCallerSession struct {
	Contract *USDCTokenPoolCaller
	CallOpts bind.CallOpts
}

type USDCTokenPoolTransactorSession struct {
	Contract     *USDCTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type USDCTokenPoolRaw struct {
	Contract *USDCTokenPool
}

type USDCTokenPoolCallerRaw struct {
	Contract *USDCTokenPoolCaller
}

type USDCTokenPoolTransactorRaw struct {
	Contract *USDCTokenPoolTransactor
}

func NewUSDCTokenPool(address common.Address, backend bind.ContractBackend) (*USDCTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(USDCTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindUSDCTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPool{address: address, abi: abi, USDCTokenPoolCaller: USDCTokenPoolCaller{contract: contract}, USDCTokenPoolTransactor: USDCTokenPoolTransactor{contract: contract}, USDCTokenPoolFilterer: USDCTokenPoolFilterer{contract: contract}}, nil
}

func NewUSDCTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*USDCTokenPoolCaller, error) {
	contract, err := bindUSDCTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCaller{contract: contract}, nil
}

func NewUSDCTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*USDCTokenPoolTransactor, error) {
	contract, err := bindUSDCTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolTransactor{contract: contract}, nil
}

func NewUSDCTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*USDCTokenPoolFilterer, error) {
	contract, err := bindUSDCTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolFilterer{contract: contract}, nil
}

func bindUSDCTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := USDCTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_USDCTokenPool *USDCTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCTokenPool.Contract.USDCTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_USDCTokenPool *USDCTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.USDCTokenPoolTransactor.contract.Transfer(opts)
}

func (_USDCTokenPool *USDCTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.USDCTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_USDCTokenPool *USDCTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_USDCTokenPool *USDCTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.contract.Transfer(opts)
}

func (_USDCTokenPool *USDCTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_USDCTokenPool *USDCTokenPoolCaller) SUPPORTEDUSDCVERSION(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "SUPPORTED_USDC_VERSION")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) SUPPORTEDUSDCVERSION() (uint32, error) {
	return _USDCTokenPool.Contract.SUPPORTEDUSDCVERSION(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) SUPPORTEDUSDCVERSION() (uint32, error) {
	return _USDCTokenPool.Contract.SUPPORTEDUSDCVERSION(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) CurrentOffRampRateLimiterState(opts *bind.CallOpts, offRamp common.Address) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "currentOffRampRateLimiterState", offRamp)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) CurrentOffRampRateLimiterState(offRamp common.Address) (RateLimiterTokenBucket, error) {
	return _USDCTokenPool.Contract.CurrentOffRampRateLimiterState(&_USDCTokenPool.CallOpts, offRamp)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) CurrentOffRampRateLimiterState(offRamp common.Address) (RateLimiterTokenBucket, error) {
	return _USDCTokenPool.Contract.CurrentOffRampRateLimiterState(&_USDCTokenPool.CallOpts, offRamp)
}

func (_USDCTokenPool *USDCTokenPoolCaller) CurrentOnRampRateLimiterState(opts *bind.CallOpts, onRamp common.Address) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "currentOnRampRateLimiterState", onRamp)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) CurrentOnRampRateLimiterState(onRamp common.Address) (RateLimiterTokenBucket, error) {
	return _USDCTokenPool.Contract.CurrentOnRampRateLimiterState(&_USDCTokenPool.CallOpts, onRamp)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) CurrentOnRampRateLimiterState(onRamp common.Address) (RateLimiterTokenBucket, error) {
	return _USDCTokenPool.Contract.CurrentOnRampRateLimiterState(&_USDCTokenPool.CallOpts, onRamp)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _USDCTokenPool.Contract.GetAllowList(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _USDCTokenPool.Contract.GetAllowList(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _USDCTokenPool.Contract.GetAllowListEnabled(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _USDCTokenPool.Contract.GetAllowListEnabled(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetArmProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getArmProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetArmProxy() (common.Address, error) {
	return _USDCTokenPool.Contract.GetArmProxy(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetArmProxy() (common.Address, error) {
	return _USDCTokenPool.Contract.GetArmProxy(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetDomain(opts *bind.CallOpts, chainSelector uint64) (USDCTokenPoolDomain, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getDomain", chainSelector)

	if err != nil {
		return *new(USDCTokenPoolDomain), err
	}

	out0 := *abi.ConvertType(out[0], new(USDCTokenPoolDomain)).(*USDCTokenPoolDomain)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetDomain(chainSelector uint64) (USDCTokenPoolDomain, error) {
	return _USDCTokenPool.Contract.GetDomain(&_USDCTokenPool.CallOpts, chainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetDomain(chainSelector uint64) (USDCTokenPoolDomain, error) {
	return _USDCTokenPool.Contract.GetDomain(&_USDCTokenPool.CallOpts, chainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetOffRamps(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getOffRamps")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetOffRamps() ([]common.Address, error) {
	return _USDCTokenPool.Contract.GetOffRamps(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetOffRamps() ([]common.Address, error) {
	return _USDCTokenPool.Contract.GetOffRamps(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetOnRamps(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getOnRamps")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetOnRamps() ([]common.Address, error) {
	return _USDCTokenPool.Contract.GetOnRamps(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetOnRamps() ([]common.Address, error) {
	return _USDCTokenPool.Contract.GetOnRamps(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetToken() (common.Address, error) {
	return _USDCTokenPool.Contract.GetToken(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _USDCTokenPool.Contract.GetToken(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetUSDCInterfaceId(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getUSDCInterfaceId")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetUSDCInterfaceId() ([4]byte, error) {
	return _USDCTokenPool.Contract.GetUSDCInterfaceId(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetUSDCInterfaceId() ([4]byte, error) {
	return _USDCTokenPool.Contract.GetUSDCInterfaceId(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) ILocalDomainIdentifier(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "i_localDomainIdentifier")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) ILocalDomainIdentifier() (uint32, error) {
	return _USDCTokenPool.Contract.ILocalDomainIdentifier(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) ILocalDomainIdentifier() (uint32, error) {
	return _USDCTokenPool.Contract.ILocalDomainIdentifier(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) IMessageTransmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "i_messageTransmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) IMessageTransmitter() (common.Address, error) {
	return _USDCTokenPool.Contract.IMessageTransmitter(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) IMessageTransmitter() (common.Address, error) {
	return _USDCTokenPool.Contract.IMessageTransmitter(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) ITokenMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "i_tokenMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) ITokenMessenger() (common.Address, error) {
	return _USDCTokenPool.Contract.ITokenMessenger(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) ITokenMessenger() (common.Address, error) {
	return _USDCTokenPool.Contract.ITokenMessenger(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "isOffRamp", offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _USDCTokenPool.Contract.IsOffRamp(&_USDCTokenPool.CallOpts, offRamp)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _USDCTokenPool.Contract.IsOffRamp(&_USDCTokenPool.CallOpts, offRamp)
}

func (_USDCTokenPool *USDCTokenPoolCaller) IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "isOnRamp", onRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _USDCTokenPool.Contract.IsOnRamp(&_USDCTokenPool.CallOpts, onRamp)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _USDCTokenPool.Contract.IsOnRamp(&_USDCTokenPool.CallOpts, onRamp)
}

func (_USDCTokenPool *USDCTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) Owner() (common.Address, error) {
	return _USDCTokenPool.Contract.Owner(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) Owner() (common.Address, error) {
	return _USDCTokenPool.Contract.Owner(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _USDCTokenPool.Contract.SupportsInterface(&_USDCTokenPool.CallOpts, interfaceId)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _USDCTokenPool.Contract.SupportsInterface(&_USDCTokenPool.CallOpts, interfaceId)
}

func (_USDCTokenPool *USDCTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) TypeAndVersion() (string, error) {
	return _USDCTokenPool.Contract.TypeAndVersion(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _USDCTokenPool.Contract.TypeAndVersion(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_USDCTokenPool *USDCTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _USDCTokenPool.Contract.AcceptOwnership(&_USDCTokenPool.TransactOpts)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _USDCTokenPool.Contract.AcceptOwnership(&_USDCTokenPool.TransactOpts)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_USDCTokenPool *USDCTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyAllowListUpdates(&_USDCTokenPool.TransactOpts, removes, adds)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyAllowListUpdates(&_USDCTokenPool.TransactOpts, removes, adds)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) ApplyRampUpdates(opts *bind.TransactOpts, onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "applyRampUpdates", onRamps, offRamps)
}

func (_USDCTokenPool *USDCTokenPoolSession) ApplyRampUpdates(onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyRampUpdates(&_USDCTokenPool.TransactOpts, onRamps, offRamps)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) ApplyRampUpdates(onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyRampUpdates(&_USDCTokenPool.TransactOpts, onRamps, offRamps)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, originalSender common.Address, destinationReceiver []byte, amount *big.Int, destChainSelector uint64, arg4 []byte) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "lockOrBurn", originalSender, destinationReceiver, amount, destChainSelector, arg4)
}

func (_USDCTokenPool *USDCTokenPoolSession) LockOrBurn(originalSender common.Address, destinationReceiver []byte, amount *big.Int, destChainSelector uint64, arg4 []byte) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.LockOrBurn(&_USDCTokenPool.TransactOpts, originalSender, destinationReceiver, amount, destChainSelector, arg4)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) LockOrBurn(originalSender common.Address, destinationReceiver []byte, amount *big.Int, destChainSelector uint64, arg4 []byte) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.LockOrBurn(&_USDCTokenPool.TransactOpts, originalSender, destinationReceiver, amount, destChainSelector, arg4)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, arg0 []byte, receiver common.Address, amount *big.Int, arg3 uint64, extraData []byte) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "releaseOrMint", arg0, receiver, amount, arg3, extraData)
}

func (_USDCTokenPool *USDCTokenPoolSession) ReleaseOrMint(arg0 []byte, receiver common.Address, amount *big.Int, arg3 uint64, extraData []byte) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ReleaseOrMint(&_USDCTokenPool.TransactOpts, arg0, receiver, amount, arg3, extraData)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) ReleaseOrMint(arg0 []byte, receiver common.Address, amount *big.Int, arg3 uint64, extraData []byte) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ReleaseOrMint(&_USDCTokenPool.TransactOpts, arg0, receiver, amount, arg3, extraData)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) SetDomains(opts *bind.TransactOpts, domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "setDomains", domains)
}

func (_USDCTokenPool *USDCTokenPoolSession) SetDomains(domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetDomains(&_USDCTokenPool.TransactOpts, domains)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) SetDomains(domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetDomains(&_USDCTokenPool.TransactOpts, domains)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) SetOffRampRateLimiterConfig(opts *bind.TransactOpts, offRamp common.Address, config RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "setOffRampRateLimiterConfig", offRamp, config)
}

func (_USDCTokenPool *USDCTokenPoolSession) SetOffRampRateLimiterConfig(offRamp common.Address, config RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetOffRampRateLimiterConfig(&_USDCTokenPool.TransactOpts, offRamp, config)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) SetOffRampRateLimiterConfig(offRamp common.Address, config RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetOffRampRateLimiterConfig(&_USDCTokenPool.TransactOpts, offRamp, config)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) SetOnRampRateLimiterConfig(opts *bind.TransactOpts, onRamp common.Address, config RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "setOnRampRateLimiterConfig", onRamp, config)
}

func (_USDCTokenPool *USDCTokenPoolSession) SetOnRampRateLimiterConfig(onRamp common.Address, config RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetOnRampRateLimiterConfig(&_USDCTokenPool.TransactOpts, onRamp, config)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) SetOnRampRateLimiterConfig(onRamp common.Address, config RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetOnRampRateLimiterConfig(&_USDCTokenPool.TransactOpts, onRamp, config)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_USDCTokenPool *USDCTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.TransferOwnership(&_USDCTokenPool.TransactOpts, to)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.TransferOwnership(&_USDCTokenPool.TransactOpts, to)
}

type USDCTokenPoolAllowListAddIterator struct {
	Event *USDCTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolAllowListAdd)
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
		it.Event = new(USDCTokenPoolAllowListAdd)
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

func (it *USDCTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*USDCTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolAllowListAddIterator{contract: _USDCTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolAllowListAdd)
				if err := _USDCTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*USDCTokenPoolAllowListAdd, error) {
	event := new(USDCTokenPoolAllowListAdd)
	if err := _USDCTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolAllowListRemoveIterator struct {
	Event *USDCTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolAllowListRemove)
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
		it.Event = new(USDCTokenPoolAllowListRemove)
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

func (it *USDCTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*USDCTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolAllowListRemoveIterator{contract: _USDCTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolAllowListRemove)
				if err := _USDCTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*USDCTokenPoolAllowListRemove, error) {
	event := new(USDCTokenPoolAllowListRemove)
	if err := _USDCTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolBurnedIterator struct {
	Event *USDCTokenPoolBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolBurned)
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
		it.Event = new(USDCTokenPoolBurned)
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

func (it *USDCTokenPoolBurnedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*USDCTokenPoolBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolBurnedIterator{contract: _USDCTokenPool.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolBurned)
				if err := _USDCTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseBurned(log types.Log) (*USDCTokenPoolBurned, error) {
	event := new(USDCTokenPoolBurned)
	if err := _USDCTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolConfigSetIterator struct {
	Event *USDCTokenPoolConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolConfigSet)
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
		it.Event = new(USDCTokenPoolConfigSet)
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

func (it *USDCTokenPoolConfigSetIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolConfigSet struct {
	TokenMessenger common.Address
	Raw            types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterConfigSet(opts *bind.FilterOpts) (*USDCTokenPoolConfigSetIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolConfigSetIterator{contract: _USDCTokenPool.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolConfigSet) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolConfigSet)
				if err := _USDCTokenPool.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseConfigSet(log types.Log) (*USDCTokenPoolConfigSet, error) {
	event := new(USDCTokenPoolConfigSet)
	if err := _USDCTokenPool.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolDomainsSetIterator struct {
	Event *USDCTokenPoolDomainsSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolDomainsSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolDomainsSet)
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
		it.Event = new(USDCTokenPoolDomainsSet)
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

func (it *USDCTokenPoolDomainsSetIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolDomainsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolDomainsSet struct {
	Arg0 []USDCTokenPoolDomainUpdate
	Raw  types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterDomainsSet(opts *bind.FilterOpts) (*USDCTokenPoolDomainsSetIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "DomainsSet")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolDomainsSetIterator{contract: _USDCTokenPool.contract, event: "DomainsSet", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolDomainsSet) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "DomainsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolDomainsSet)
				if err := _USDCTokenPool.contract.UnpackLog(event, "DomainsSet", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseDomainsSet(log types.Log) (*USDCTokenPoolDomainsSet, error) {
	event := new(USDCTokenPoolDomainsSet)
	if err := _USDCTokenPool.contract.UnpackLog(event, "DomainsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolLockedIterator struct {
	Event *USDCTokenPoolLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolLocked)
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
		it.Event = new(USDCTokenPoolLocked)
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

func (it *USDCTokenPoolLockedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*USDCTokenPoolLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolLockedIterator{contract: _USDCTokenPool.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolLocked)
				if err := _USDCTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseLocked(log types.Log) (*USDCTokenPoolLocked, error) {
	event := new(USDCTokenPoolLocked)
	if err := _USDCTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolMintedIterator struct {
	Event *USDCTokenPoolMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolMinted)
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
		it.Event = new(USDCTokenPoolMinted)
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

func (it *USDCTokenPoolMintedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*USDCTokenPoolMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolMintedIterator{contract: _USDCTokenPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolMinted)
				if err := _USDCTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseMinted(log types.Log) (*USDCTokenPoolMinted, error) {
	event := new(USDCTokenPoolMinted)
	if err := _USDCTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolOffRampAddedIterator struct {
	Event *USDCTokenPoolOffRampAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolOffRampAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolOffRampAdded)
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
		it.Event = new(USDCTokenPoolOffRampAdded)
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

func (it *USDCTokenPoolOffRampAddedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolOffRampAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolOffRampAdded struct {
	OffRamp           common.Address
	RateLimiterConfig RateLimiterConfig
	Raw               types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterOffRampAdded(opts *bind.FilterOpts) (*USDCTokenPoolOffRampAddedIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "OffRampAdded")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolOffRampAddedIterator{contract: _USDCTokenPool.contract, event: "OffRampAdded", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOffRampAdded) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "OffRampAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolOffRampAdded)
				if err := _USDCTokenPool.contract.UnpackLog(event, "OffRampAdded", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseOffRampAdded(log types.Log) (*USDCTokenPoolOffRampAdded, error) {
	event := new(USDCTokenPoolOffRampAdded)
	if err := _USDCTokenPool.contract.UnpackLog(event, "OffRampAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolOffRampConfiguredIterator struct {
	Event *USDCTokenPoolOffRampConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolOffRampConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolOffRampConfigured)
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
		it.Event = new(USDCTokenPoolOffRampConfigured)
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

func (it *USDCTokenPoolOffRampConfiguredIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolOffRampConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolOffRampConfigured struct {
	OffRamp           common.Address
	RateLimiterConfig RateLimiterConfig
	Raw               types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterOffRampConfigured(opts *bind.FilterOpts) (*USDCTokenPoolOffRampConfiguredIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "OffRampConfigured")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolOffRampConfiguredIterator{contract: _USDCTokenPool.contract, event: "OffRampConfigured", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchOffRampConfigured(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOffRampConfigured) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "OffRampConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolOffRampConfigured)
				if err := _USDCTokenPool.contract.UnpackLog(event, "OffRampConfigured", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseOffRampConfigured(log types.Log) (*USDCTokenPoolOffRampConfigured, error) {
	event := new(USDCTokenPoolOffRampConfigured)
	if err := _USDCTokenPool.contract.UnpackLog(event, "OffRampConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolOffRampRemovedIterator struct {
	Event *USDCTokenPoolOffRampRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolOffRampRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolOffRampRemoved)
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
		it.Event = new(USDCTokenPoolOffRampRemoved)
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

func (it *USDCTokenPoolOffRampRemovedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolOffRampRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolOffRampRemoved struct {
	OffRamp common.Address
	Raw     types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterOffRampRemoved(opts *bind.FilterOpts) (*USDCTokenPoolOffRampRemovedIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "OffRampRemoved")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolOffRampRemovedIterator{contract: _USDCTokenPool.contract, event: "OffRampRemoved", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOffRampRemoved) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "OffRampRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolOffRampRemoved)
				if err := _USDCTokenPool.contract.UnpackLog(event, "OffRampRemoved", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseOffRampRemoved(log types.Log) (*USDCTokenPoolOffRampRemoved, error) {
	event := new(USDCTokenPoolOffRampRemoved)
	if err := _USDCTokenPool.contract.UnpackLog(event, "OffRampRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolOnRampAddedIterator struct {
	Event *USDCTokenPoolOnRampAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolOnRampAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolOnRampAdded)
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
		it.Event = new(USDCTokenPoolOnRampAdded)
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

func (it *USDCTokenPoolOnRampAddedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolOnRampAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolOnRampAdded struct {
	OnRamp            common.Address
	RateLimiterConfig RateLimiterConfig
	Raw               types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterOnRampAdded(opts *bind.FilterOpts) (*USDCTokenPoolOnRampAddedIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "OnRampAdded")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolOnRampAddedIterator{contract: _USDCTokenPool.contract, event: "OnRampAdded", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchOnRampAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOnRampAdded) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "OnRampAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolOnRampAdded)
				if err := _USDCTokenPool.contract.UnpackLog(event, "OnRampAdded", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseOnRampAdded(log types.Log) (*USDCTokenPoolOnRampAdded, error) {
	event := new(USDCTokenPoolOnRampAdded)
	if err := _USDCTokenPool.contract.UnpackLog(event, "OnRampAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolOnRampConfiguredIterator struct {
	Event *USDCTokenPoolOnRampConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolOnRampConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolOnRampConfigured)
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
		it.Event = new(USDCTokenPoolOnRampConfigured)
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

func (it *USDCTokenPoolOnRampConfiguredIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolOnRampConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolOnRampConfigured struct {
	OnRamp            common.Address
	RateLimiterConfig RateLimiterConfig
	Raw               types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterOnRampConfigured(opts *bind.FilterOpts) (*USDCTokenPoolOnRampConfiguredIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "OnRampConfigured")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolOnRampConfiguredIterator{contract: _USDCTokenPool.contract, event: "OnRampConfigured", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchOnRampConfigured(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOnRampConfigured) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "OnRampConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolOnRampConfigured)
				if err := _USDCTokenPool.contract.UnpackLog(event, "OnRampConfigured", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseOnRampConfigured(log types.Log) (*USDCTokenPoolOnRampConfigured, error) {
	event := new(USDCTokenPoolOnRampConfigured)
	if err := _USDCTokenPool.contract.UnpackLog(event, "OnRampConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolOnRampRemovedIterator struct {
	Event *USDCTokenPoolOnRampRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolOnRampRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolOnRampRemoved)
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
		it.Event = new(USDCTokenPoolOnRampRemoved)
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

func (it *USDCTokenPoolOnRampRemovedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolOnRampRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolOnRampRemoved struct {
	OnRamp common.Address
	Raw    types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterOnRampRemoved(opts *bind.FilterOpts) (*USDCTokenPoolOnRampRemovedIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "OnRampRemoved")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolOnRampRemovedIterator{contract: _USDCTokenPool.contract, event: "OnRampRemoved", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchOnRampRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOnRampRemoved) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "OnRampRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolOnRampRemoved)
				if err := _USDCTokenPool.contract.UnpackLog(event, "OnRampRemoved", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseOnRampRemoved(log types.Log) (*USDCTokenPoolOnRampRemoved, error) {
	event := new(USDCTokenPoolOnRampRemoved)
	if err := _USDCTokenPool.contract.UnpackLog(event, "OnRampRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolOwnershipTransferRequestedIterator struct {
	Event *USDCTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolOwnershipTransferRequested)
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
		it.Event = new(USDCTokenPoolOwnershipTransferRequested)
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

func (it *USDCTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolOwnershipTransferRequestedIterator{contract: _USDCTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolOwnershipTransferRequested)
				if err := _USDCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*USDCTokenPoolOwnershipTransferRequested, error) {
	event := new(USDCTokenPoolOwnershipTransferRequested)
	if err := _USDCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolOwnershipTransferredIterator struct {
	Event *USDCTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolOwnershipTransferred)
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
		it.Event = new(USDCTokenPoolOwnershipTransferred)
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

func (it *USDCTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolOwnershipTransferredIterator{contract: _USDCTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolOwnershipTransferred)
				if err := _USDCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*USDCTokenPoolOwnershipTransferred, error) {
	event := new(USDCTokenPoolOwnershipTransferred)
	if err := _USDCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolReleasedIterator struct {
	Event *USDCTokenPoolReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolReleased)
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
		it.Event = new(USDCTokenPoolReleased)
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

func (it *USDCTokenPoolReleasedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*USDCTokenPoolReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolReleasedIterator{contract: _USDCTokenPool.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolReleased)
				if err := _USDCTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseReleased(log types.Log) (*USDCTokenPoolReleased, error) {
	event := new(USDCTokenPoolReleased)
	if err := _USDCTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_USDCTokenPool *USDCTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _USDCTokenPool.abi.Events["AllowListAdd"].ID:
		return _USDCTokenPool.ParseAllowListAdd(log)
	case _USDCTokenPool.abi.Events["AllowListRemove"].ID:
		return _USDCTokenPool.ParseAllowListRemove(log)
	case _USDCTokenPool.abi.Events["Burned"].ID:
		return _USDCTokenPool.ParseBurned(log)
	case _USDCTokenPool.abi.Events["ConfigSet"].ID:
		return _USDCTokenPool.ParseConfigSet(log)
	case _USDCTokenPool.abi.Events["DomainsSet"].ID:
		return _USDCTokenPool.ParseDomainsSet(log)
	case _USDCTokenPool.abi.Events["Locked"].ID:
		return _USDCTokenPool.ParseLocked(log)
	case _USDCTokenPool.abi.Events["Minted"].ID:
		return _USDCTokenPool.ParseMinted(log)
	case _USDCTokenPool.abi.Events["OffRampAdded"].ID:
		return _USDCTokenPool.ParseOffRampAdded(log)
	case _USDCTokenPool.abi.Events["OffRampConfigured"].ID:
		return _USDCTokenPool.ParseOffRampConfigured(log)
	case _USDCTokenPool.abi.Events["OffRampRemoved"].ID:
		return _USDCTokenPool.ParseOffRampRemoved(log)
	case _USDCTokenPool.abi.Events["OnRampAdded"].ID:
		return _USDCTokenPool.ParseOnRampAdded(log)
	case _USDCTokenPool.abi.Events["OnRampConfigured"].ID:
		return _USDCTokenPool.ParseOnRampConfigured(log)
	case _USDCTokenPool.abi.Events["OnRampRemoved"].ID:
		return _USDCTokenPool.ParseOnRampRemoved(log)
	case _USDCTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _USDCTokenPool.ParseOwnershipTransferRequested(log)
	case _USDCTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _USDCTokenPool.ParseOwnershipTransferred(log)
	case _USDCTokenPool.abi.Events["Released"].ID:
		return _USDCTokenPool.ParseReleased(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (USDCTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (USDCTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (USDCTokenPoolBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
}

func (USDCTokenPoolConfigSet) Topic() common.Hash {
	return common.HexToHash("0x2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea9544")
}

func (USDCTokenPoolDomainsSet) Topic() common.Hash {
	return common.HexToHash("0x1889010d2535a0ab1643678d1da87fbbe8b87b2f585b47ddb72ec622aef9ee56")
}

func (USDCTokenPoolLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (USDCTokenPoolMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (USDCTokenPoolOffRampAdded) Topic() common.Hash {
	return common.HexToHash("0x395b7374909d2b54e5796f53c898ebf41d767c86c78ea86519acf2b805852d88")
}

func (USDCTokenPoolOffRampConfigured) Topic() common.Hash {
	return common.HexToHash("0xb3ba339cfbb8ef80d7a29ce5493051cb90e64fcfa85d7124efc1adfa4c68399f")
}

func (USDCTokenPoolOffRampRemoved) Topic() common.Hash {
	return common.HexToHash("0xcf91daec21e3510e2f2aea4b09d08c235d5c6844980be709f282ef591dbf420c")
}

func (USDCTokenPoolOnRampAdded) Topic() common.Hash {
	return common.HexToHash("0x0b594bb0555ff7b252e0c789ccc9d8903fec294172064308727d570505cee1ac")
}

func (USDCTokenPoolOnRampConfigured) Topic() common.Hash {
	return common.HexToHash("0x578db78e348076074dbff64a94073a83e9a65aa6766b8c75fdc89282b0e30ed6")
}

func (USDCTokenPoolOnRampRemoved) Topic() common.Hash {
	return common.HexToHash("0x7fd064821314ad863a0714a3f1229375ace6b6427ed5544b7b2ba1c47b1b5294")
}

func (USDCTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (USDCTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (USDCTokenPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (_USDCTokenPool *USDCTokenPool) Address() common.Address {
	return _USDCTokenPool.address
}

type USDCTokenPoolInterface interface {
	SUPPORTEDUSDCVERSION(opts *bind.CallOpts) (uint32, error)

	CurrentOffRampRateLimiterState(opts *bind.CallOpts, offRamp common.Address) (RateLimiterTokenBucket, error)

	CurrentOnRampRateLimiterState(opts *bind.CallOpts, onRamp common.Address) (RateLimiterTokenBucket, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetArmProxy(opts *bind.CallOpts) (common.Address, error)

	GetDomain(opts *bind.CallOpts, chainSelector uint64) (USDCTokenPoolDomain, error)

	GetOffRamps(opts *bind.CallOpts) ([]common.Address, error)

	GetOnRamps(opts *bind.CallOpts) ([]common.Address, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetUSDCInterfaceId(opts *bind.CallOpts) ([4]byte, error)

	ILocalDomainIdentifier(opts *bind.CallOpts) (uint32, error)

	IMessageTransmitter(opts *bind.CallOpts) (common.Address, error)

	ITokenMessenger(opts *bind.CallOpts) (common.Address, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyRampUpdates(opts *bind.TransactOpts, onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, originalSender common.Address, destinationReceiver []byte, amount *big.Int, destChainSelector uint64, arg4 []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, arg0 []byte, receiver common.Address, amount *big.Int, arg3 uint64, extraData []byte) (*types.Transaction, error)

	SetDomains(opts *bind.TransactOpts, domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error)

	SetOffRampRateLimiterConfig(opts *bind.TransactOpts, offRamp common.Address, config RateLimiterConfig) (*types.Transaction, error)

	SetOnRampRateLimiterConfig(opts *bind.TransactOpts, onRamp common.Address, config RateLimiterConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*USDCTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*USDCTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*USDCTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*USDCTokenPoolAllowListRemove, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*USDCTokenPoolBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*USDCTokenPoolBurned, error)

	FilterConfigSet(opts *bind.FilterOpts) (*USDCTokenPoolConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*USDCTokenPoolConfigSet, error)

	FilterDomainsSet(opts *bind.FilterOpts) (*USDCTokenPoolDomainsSetIterator, error)

	WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolDomainsSet) (event.Subscription, error)

	ParseDomainsSet(log types.Log) (*USDCTokenPoolDomainsSet, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*USDCTokenPoolLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*USDCTokenPoolLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*USDCTokenPoolMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*USDCTokenPoolMinted, error)

	FilterOffRampAdded(opts *bind.FilterOpts) (*USDCTokenPoolOffRampAddedIterator, error)

	WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOffRampAdded) (event.Subscription, error)

	ParseOffRampAdded(log types.Log) (*USDCTokenPoolOffRampAdded, error)

	FilterOffRampConfigured(opts *bind.FilterOpts) (*USDCTokenPoolOffRampConfiguredIterator, error)

	WatchOffRampConfigured(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOffRampConfigured) (event.Subscription, error)

	ParseOffRampConfigured(log types.Log) (*USDCTokenPoolOffRampConfigured, error)

	FilterOffRampRemoved(opts *bind.FilterOpts) (*USDCTokenPoolOffRampRemovedIterator, error)

	WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOffRampRemoved) (event.Subscription, error)

	ParseOffRampRemoved(log types.Log) (*USDCTokenPoolOffRampRemoved, error)

	FilterOnRampAdded(opts *bind.FilterOpts) (*USDCTokenPoolOnRampAddedIterator, error)

	WatchOnRampAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOnRampAdded) (event.Subscription, error)

	ParseOnRampAdded(log types.Log) (*USDCTokenPoolOnRampAdded, error)

	FilterOnRampConfigured(opts *bind.FilterOpts) (*USDCTokenPoolOnRampConfiguredIterator, error)

	WatchOnRampConfigured(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOnRampConfigured) (event.Subscription, error)

	ParseOnRampConfigured(log types.Log) (*USDCTokenPoolOnRampConfigured, error)

	FilterOnRampRemoved(opts *bind.FilterOpts) (*USDCTokenPoolOnRampRemovedIterator, error)

	WatchOnRampRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOnRampRemoved) (event.Subscription, error)

	ParseOnRampRemoved(log types.Log) (*USDCTokenPoolOnRampRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*USDCTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*USDCTokenPoolOwnershipTransferred, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*USDCTokenPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*USDCTokenPoolReleased, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
