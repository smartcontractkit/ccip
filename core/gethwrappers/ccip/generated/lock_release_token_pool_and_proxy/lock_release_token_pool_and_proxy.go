// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lock_release_token_pool_and_proxy

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

type PoolLockOrBurnInV1 struct {
	Receiver            []byte
	RemoteChainSelector uint64
	OriginalSender      common.Address
	Amount              *big.Int
	LocalToken          common.Address
}

type PoolLockOrBurnOutV1 struct {
	DestPoolAddress []byte
	DestPoolData    []byte
}

type PoolReleaseOrMintInV1 struct {
	OriginalSender      []byte
	RemoteChainSelector uint64
	Receiver            common.Address
	Amount              *big.Int
	SourcePoolAddress   []byte
	SourcePoolData      []byte
	OffchainTokenData   []byte
}

type PoolReleaseOrMintOutV1 struct {
	LocalToken        common.Address
	DestinationAmount *big.Int
}

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

type TokenPoolChainUpdate struct {
	RemoteChainSelector       uint64
	Allowed                   bool
	RemotePoolAddress         []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

var LockReleaseTokenPoolAndProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"acceptLiquidity\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowListNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotARampOnRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"DisabledNonZeroRateLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"InvalidRatelimitRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidSourcePoolAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidityNotAccepted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"NonExistentChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateLimitMustBeDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPoolPriorTo1_5\",\"name\":\"oldPool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPoolPriorTo1_5\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"LegacyPoolChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"previousPoolAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"RemotePoolSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRouter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensConsumed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"internalType\":\"structTokenPool.ChainUpdate[]\",\"name\":\"chains\",\"type\":\"tuple[]\"}],\"name\":\"applyChainUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canAcceptLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentInboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentOutboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLockReleaseInterfaceId\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRateLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRebalancer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRemotePool\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRmnProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedChains\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"isSupportedChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isSupportedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"internalType\":\"structPool.LockOrBurnInV1\",\"name\":\"lockOrBurnIn\",\"type\":\"tuple\"}],\"name\":\"lockOrBurn\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"destPoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destPoolData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.LockOrBurnOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"provideLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"originalSender\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainTokenData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.ReleaseOrMintInV1\",\"name\":\"releaseOrMintIn\",\"type\":\"tuple\"}],\"name\":\"releaseOrMint\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundConfig\",\"type\":\"tuple\"}],\"name\":\"setChainRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPoolPriorTo1_5\",\"name\":\"prevPool\",\"type\":\"address\"}],\"name\":\"setPreviousPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rateLimitAdmin\",\"type\":\"address\"}],\"name\":\"setRateLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rebalancer\",\"type\":\"address\"}],\"name\":\"setRebalancer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"setRemotePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162004cbc38038062004cbc83398101604081905262000035916200056d565b84848483838383833380600081620000945760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c757620000c78162000186565b5050506001600160a01b0384161580620000e857506001600160a01b038116155b80620000fb57506001600160a01b038216155b156200011a576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b0384811660805282811660a052600480546001600160a01b031916918316919091179055825115801560c0526200016d576040805160008152602081019091526200016d908462000231565b5050505094151560e05250620006de9650505050505050565b336001600160a01b03821603620001e05760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016200008b565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60c05162000252576040516335f4a7b360e01b815260040160405180910390fd5b60005b8251811015620002dd57600083828151811062000276576200027662000690565b60209081029190910101519050620002906002826200038e565b15620002d3576040516001600160a01b03821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b5060010162000255565b5060005b81518110156200038957600082828151811062000302576200030262000690565b6020026020010151905060006001600160a01b0316816001600160a01b0316036200032e575062000380565b6200033b600282620003ae565b156200037e576040516001600160a01b03821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b600101620002e1565b505050565b6000620003a5836001600160a01b038416620003c5565b90505b92915050565b6000620003a5836001600160a01b038416620004c9565b60008181526001830160205260408120548015620004be576000620003ec600183620006a6565b85549091506000906200040290600190620006a6565b90508181146200046e57600086600001828154811062000426576200042662000690565b90600052602060002001549050808760000184815481106200044c576200044c62000690565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080620004825762000482620006c8565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050620003a8565b6000915050620003a8565b60008181526001830160205260408120546200051257508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155620003a8565b506000620003a8565b6001600160a01b03811681146200053157600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b805162000557816200051b565b919050565b805180151581146200055757600080fd5b600080600080600060a086880312156200058657600080fd5b855162000593816200051b565b602087810151919650906001600160401b0380821115620005b357600080fd5b818901915089601f830112620005c857600080fd5b815181811115620005dd57620005dd62000534565b8060051b604051601f19603f8301168101818110858211171562000605576200060562000534565b60405291825284820192508381018501918c8311156200062457600080fd5b938501935b828510156200064d576200063d856200054a565b8452938501939285019262000629565b80995050505050505062000664604087016200054a565b925062000674606087016200055c565b915062000684608087016200054a565b90509295509295909350565b634e487b7160e01b600052603260045260246000fd5b81810381811115620003a857634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60805160a05160c05160e05161453b620007816000396000818161055401526114690152600081816105ee01528181611e2301526129e10152600081816105c801528181611bd201526120d4015260008181610292015281816102e70152818161082b015281816108fd015281816109cc01528181610aa90152818161152b01528181611ff1015281816121da015281816129770152612bd1015261453b6000f3fe608060405234801561001057600080fd5b50600436106102265760003560e01c80638da5cb5b1161012a578063c0d78655116100bd578063dc0bd9711161008c578063eb521a4c11610071578063eb521a4c14610612578063f2fde38b14610625578063f6e2145e1461063857600080fd5b8063dc0bd971146105c6578063e0351e13146105ec57600080fd5b8063c0d7865514610578578063c4bffe2b1461058b578063c75eea9c146105a0578063cf7401f3146105b357600080fd5b8063a8d87a3b116100f9578063a8d87a3b146104b2578063af58d59f146104c5578063b0f479a114610534578063bb98546b1461055257600080fd5b80638da5cb5b1461044c5780639766b9321461046a5780639a4575b91461047d578063a7cd63b71461049d57600080fd5b806354c8a4f3116101bd57806379ba50971161018c57806383826b2b1161017157806383826b2b146103f85780638926f54f1461040b5780638bfca18c1461041e57600080fd5b806379ba5097146103dd5780637d54534e146103e557600080fd5b806354c8a4f3146103865780636cfd1553146103995780636d3d1a58146103ac57806378a010b2146103ca57600080fd5b806321df0da7116101f957806321df0da714610290578063240028e8146102d75780634059f55b14610324578063432a6ba31461036857600080fd5b806301ffc9a71461022b5780630a2fd493146102535780630a861f2a14610273578063181f5a7714610288575b600080fd5b61023e61023936600461357b565b61064b565b60405190151581526020015b60405180910390f35b6102666102613660046135da565b6106f3565b60405161024a9190613663565b610286610281366004613676565b6107a3565b005b610266610954565b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161024a565b61023e6102e53660046136bc565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff90811691161490565b6103376103323660046136d9565b610970565b60408051825173ffffffffffffffffffffffffffffffffffffffff168152602092830151928101929092520161024a565b60095473ffffffffffffffffffffffffffffffffffffffff166102b2565b610286610394366004613760565b610ad9565b6102866103a73660046136bc565b610b54565b600a5473ffffffffffffffffffffffffffffffffffffffff166102b2565b6102866103d83660046137cc565b610ba3565b610286610d12565b6102866103f33660046136bc565b610e0f565b61023e61040636600461384f565b610e5e565b61023e6104193660046135da565b610f2b565b6040517f98a4717700000000000000000000000000000000000000000000000000000000815260200161024a565b60005473ffffffffffffffffffffffffffffffffffffffff166102b2565b6102866104783660046136bc565b610f42565b61049061048b366004613886565b610fd1565b60405161024a91906138c1565b6104a561109a565b60405161024a9190613921565b6102b26104c03660046135da565b503090565b6104d86104d33660046135da565b6110ab565b60405161024a919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b60045473ffffffffffffffffffffffffffffffffffffffff166102b2565b7f000000000000000000000000000000000000000000000000000000000000000061023e565b6102866105863660046136bc565b611180565b610593611254565b60405161024a919061397b565b6104d86105ae3660046135da565b61130c565b6102866105c1366004613b26565b6113de565b7f00000000000000000000000000000000000000000000000000000000000000006102b2565b7f000000000000000000000000000000000000000000000000000000000000000061023e565b610286610620366004613676565b611467565b6102866106333660046136bc565b611583565b610286610646366004613b6b565b611597565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f98a471770000000000000000000000000000000000000000000000000000000014806106de57507fffffffff0000000000000000000000000000000000000000000000000000000082167fe1d4056600000000000000000000000000000000000000000000000000000000145b806106ed57506106ed826119dc565b92915050565b67ffffffffffffffff8116600090815260076020526040902060040180546060919061071e90613bad565b80601f016020809104026020016040519081016040528092919081815260200182805461074a90613bad565b80156107975780601f1061076c57610100808354040283529160200191610797565b820191906000526020600020905b81548152906001019060200180831161077a57829003601f168201915b50505050509050919050565b60095473ffffffffffffffffffffffffffffffffffffffff1633146107fb576040517f8e4a23d60000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015281907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610887573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ab9190613c00565b10156108e3576040517fbb55fd2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61092473ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163383611ac0565b604051819033907fc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf984017171990600090a350565b6040518060600160405280602681526020016145096026913981565b604080518082019091526000808252602082015261099561099083613cb5565b611b94565b60085473ffffffffffffffffffffffffffffffffffffffff166109fc576109f773ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016336060850135611ac0565b610a0d565b610a0d610a0883613cb5565b611d05565b610a1d60608301604084016136bc565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f528460600135604051610a7f91815260200190565b60405180910390a3506040805180820190915273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152606090910135602082015290565b610ae1611d9e565b610b4e84848080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808802828101820190935287825290935087925086918291850190849080828437600092019190915250611e2192505050565b50505050565b610b5c611d9e565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b610bab611d9e565b610bb483610f2b565b610bf6576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024016107f2565b67ffffffffffffffff831660009081526007602052604081206004018054610c1d90613bad565b80601f0160208091040260200160405190810160405280929190818152602001828054610c4990613bad565b8015610c965780601f10610c6b57610100808354040283529160200191610c96565b820191906000526020600020905b815481529060010190602001808311610c7957829003601f168201915b5050505067ffffffffffffffff8616600090815260076020526040902091925050600401610cc5838583613de8565b508367ffffffffffffffff167fdb4d6220746a38cbc5335f7e108f7de80f482f4d23350253dfd0917df75a14bf828585604051610d0493929190613f02565b60405180910390a250505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610d93576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016107f2565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610e17611d9e565b600a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600073ffffffffffffffffffffffffffffffffffffffff8216301480610f245750600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86169281019290925273ffffffffffffffffffffffffffffffffffffffff848116602484015216906383826b2b90604401602060405180830381865afa158015610f00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f249190613f66565b9392505050565b60006106ed600567ffffffffffffffff8416611fd7565b610f4a611d9e565b6008805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f81accd0a7023865eaa51b3399dd0eafc488bf3ba238402911e1659cfe860f22891015b60405180910390a15050565b6040805180820190915260608082526020820152610ff6610ff183613f83565b611fef565b60085473ffffffffffffffffffffffffffffffffffffffff16156110255761102561102083613f83565b6121b7565b6040516060830135815233907f9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd600089060200160405180910390a2604051806040016040528061107f84602001602081019061026191906135da565b81526040805160208181019092526000815291015292915050565b60606110a660026122d5565b905090565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845260028201546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff1615159482019490945260039091015480841660608301529190910490911660808201526106ed906122e2565b611188611d9e565b73ffffffffffffffffffffffffffffffffffffffff81166111d5576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849101610fc5565b6060600061126260056122d5565b90506000815167ffffffffffffffff811115611280576112806139bd565b6040519080825280602002602001820160405280156112a9578160200160208202803683370190505b50905060005b8251811015611305578281815181106112ca576112ca61400a565b60200260200101518282815181106112e4576112e461400a565b67ffffffffffffffff909216602092830291909101909101526001016112af565b5092915050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845281546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff1615159482019490945260019091015480841660608301529190910490911660808201526106ed906122e2565b600a5473ffffffffffffffffffffffffffffffffffffffff16331480159061141e575060005473ffffffffffffffffffffffffffffffffffffffff163314155b15611457576040517f8e4a23d60000000000000000000000000000000000000000000000000000000081523360048201526024016107f2565b611462838383612394565b505050565b7f00000000000000000000000000000000000000000000000000000000000000006114be576040517fe93f8fa400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60095473ffffffffffffffffffffffffffffffffffffffff163314611511576040517f8e4a23d60000000000000000000000000000000000000000000000000000000081523360048201526024016107f2565b61155373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633308461247e565b604051819033907fc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb31208890600090a350565b61158b611d9e565b611594816124dc565b50565b61159f611d9e565b60005b818110156114625760008383838181106115be576115be61400a565b90506020028101906115d09190614039565b6115d990614077565b90506115ee81606001518260200151156125d1565b61160181608001518260200151156125d1565b8060200151156118ca5780516116239060059067ffffffffffffffff1661270a565b6116685780516040517f1d5ad3c500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016107f2565b8060400151516000036116a7576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805161010081018252606083810180516020908101516fffffffffffffffffffffffffffffffff9081168486019081524263ffffffff9081166080808901829052865151151560a0808b01919091528751870151861660c08b015296518a0151851660e08a015292885288519586018952828a01805186015185168752868601919091528051511515868a015280518501518416868801525188015183168583015283870194855288880151878901908152895167ffffffffffffffff1660009081526007865289902088518051825482890151838e01519289167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316177001000000000000000000000000000000009188168202177fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff908116740100000000000000000000000000000000000000009415158502178655848d015194890151948a16948a168202949094176001860155995180516002860180549b8301519f830151918b169b9093169a909a179d9096168a029c909c1790911696151502959095179098559485015194015193811693169091029190911760038201559151909190600482019061187d90826140fb565b50508151606083015160808401516040517f0f135cbb9afa12a8bf3bbd071c117bcca4ddeca6160ef7f33d012a81b9c0c47194506118bd93929190614215565b60405180910390a16119d3565b80516118e29060059067ffffffffffffffff16612716565b6119275780516040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016107f2565b805167ffffffffffffffff16600090815260076020526040812080547fffffffffffffffffffffff00000000000000000000000000000000000000000090811682556001820183905560028201805490911690556003810182905590611990600483018261352d565b5050805160405167ffffffffffffffff90911681527f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599169060200160405180910390a15b506001016115a2565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167faff2afbf000000000000000000000000000000000000000000000000000000001480611a6f57507fffffffff0000000000000000000000000000000000000000000000000000000082167f773a5d4500000000000000000000000000000000000000000000000000000000145b806106ed57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a7000000000000000000000000000000000000000000000000000000001492915050565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526114629084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152612722565b60208101516040517f58babe3300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906358babe3390602401602060405180830381865afa158015611c2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c529190613f66565b15611c89576040517f53ad11d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c96816020015161282e565b611ca381602001516106f3565b8051906020012081608001518051906020012014611cf35780608001516040517f24eb47e50000000000000000000000000000000000000000000000000000000081526004016107f29190613663565b61159481602001518260600151612954565b6008548151606083015160208401516040517f8627fad600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90941693638627fad693611d699390923392600401614298565b600060405180830381600087803b158015611d8357600080fd5b505af1158015611d97573d6000803e3d6000fd5b5050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314611e1f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016107f2565b565b7f0000000000000000000000000000000000000000000000000000000000000000611e78576040517f35f4a7b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8251811015611f0e576000838281518110611e9857611e9861400a565b60200260200101519050611eb681600261299b90919063ffffffff16565b15611f055760405173ffffffffffffffffffffffffffffffffffffffff821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b50600101611e7b565b5060005b8151811015611462576000828281518110611f2f57611f2f61400a565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611f735750611fcf565b611f7e6002826129bd565b15611fcd5760405173ffffffffffffffffffffffffffffffffffffffff821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b600101611f12565b60008181526001830160205260408120541515610f24565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff16146120965760808101516040517f961c9a4f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016107f2565b60208101516040517f58babe3300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906358babe3390602401602060405180830381865afa158015612130573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121549190613f66565b1561218b576040517f53ad11d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61219881604001516129df565b6121a58160200151612a63565b61159481602001518260600151612bb1565b60085460608201516122049173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811692911690611ac0565b60085460408083015183516060850151602086015193517f9687544500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9095169463968754459461226c949392916004016142f9565b6000604051808303816000875af115801561228b573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526122d19190810190614359565b5050565b60606000610f2483612bf5565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261237082606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff164261235491906143f6565b85608001516fffffffffffffffffffffffffffffffff16612c50565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b61239d83610f2b565b6123df576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024016107f2565b6123ea8260006125d1565b67ffffffffffffffff8316600090815260076020526040902061240d9083612c7a565b6124188160006125d1565b67ffffffffffffffff8316600090815260076020526040902061243e9060020182612c7a565b7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b83838360405161247193929190614215565b60405180910390a1505050565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052610b4e9085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611b12565b3373ffffffffffffffffffffffffffffffffffffffff82160361255b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016107f2565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8151156126985781602001516fffffffffffffffffffffffffffffffff1682604001516fffffffffffffffffffffffffffffffff16101580612627575060408201516fffffffffffffffffffffffffffffffff16155b1561266057816040517f70505e560000000000000000000000000000000000000000000000000000000081526004016107f29190614409565b80156122d1576040517f433fc33d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408201516fffffffffffffffffffffffffffffffff161515806126d1575060208201516fffffffffffffffffffffffffffffffff1615155b156122d157816040517fd68af9cc0000000000000000000000000000000000000000000000000000000081526004016107f29190614409565b6000610f248383612e1c565b6000610f248383612e6b565b6000612784826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16612f5e9092919063ffffffff16565b80519091501561146257808060200190518101906127a29190613f66565b611462576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016107f2565b61283781610f2b565b612879576040517fa9902c7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016107f2565b600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84169281019290925233602483015273ffffffffffffffffffffffffffffffffffffffff16906383826b2b90604401602060405180830381865afa1580156128f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061291c9190613f66565b611594576040517f728fe07b0000000000000000000000000000000000000000000000000000000081523360048201526024016107f2565b67ffffffffffffffff821660009081526007602052604090206122d190600201827f0000000000000000000000000000000000000000000000000000000000000000612f6d565b6000610f248373ffffffffffffffffffffffffffffffffffffffff8416612e6b565b6000610f248373ffffffffffffffffffffffffffffffffffffffff8416612e1c565b7f00000000000000000000000000000000000000000000000000000000000000008015612a145750612a126002826132f0565b155b15611594576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016107f2565b612a6c81610f2b565b612aae576040517fa9902c7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016107f2565b600480546040517fa8d87a3b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84169281019290925273ffffffffffffffffffffffffffffffffffffffff169063a8d87a3b90602401602060405180830381865afa158015612b27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b4b9190614445565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611594576040517f728fe07b0000000000000000000000000000000000000000000000000000000081523360048201526024016107f2565b67ffffffffffffffff821660009081526007602052604090206122d190827f0000000000000000000000000000000000000000000000000000000000000000612f6d565b60608160000180548060200260200160405190810160405280929190818152602001828054801561079757602002820191906000526020600020905b815481526020019060010190808311612c315750505050509050919050565b6000612c6f85612c608486614462565b612c6a9087614479565b61331f565b90505b949350505050565b8154600090612ca390700100000000000000000000000000000000900463ffffffff16426143f6565b90508015612d455760018301548354612ceb916fffffffffffffffffffffffffffffffff80821692811691859170010000000000000000000000000000000090910416612c50565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b60208201518354612d6b916fffffffffffffffffffffffffffffffff908116911661331f565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1990612471908490614409565b6000818152600183016020526040812054612e63575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556106ed565b5060006106ed565b60008181526001830160205260408120548015612f54576000612e8f6001836143f6565b8554909150600090612ea3906001906143f6565b9050818114612f08576000866000018281548110612ec357612ec361400a565b9060005260206000200154905080876000018481548110612ee657612ee661400a565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080612f1957612f1961448c565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506106ed565b60009150506106ed565b6060612c728484600085613335565b825474010000000000000000000000000000000000000000900460ff161580612f94575081155b15612f9e57505050565b825460018401546fffffffffffffffffffffffffffffffff80831692911690600090612fe490700100000000000000000000000000000000900463ffffffff16426143f6565b905080156130a45781831115613026576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018601546130609083908590849070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16612c50565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff167001000000000000000000000000000000004263ffffffff160217875592505b8482101561315b5773ffffffffffffffffffffffffffffffffffffffff8416613103576040517ff94ebcd100000000000000000000000000000000000000000000000000000000815260048101839052602481018690526044016107f2565b6040517f1a76572a000000000000000000000000000000000000000000000000000000008152600481018390526024810186905273ffffffffffffffffffffffffffffffffffffffff851660448201526064016107f2565b8483101561326e5760018681015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1690600090829061319f90826143f6565b6131a9878a6143f6565b6131b39190614479565b6131bd91906144bb565b905073ffffffffffffffffffffffffffffffffffffffff8616613216576040517f15279c0800000000000000000000000000000000000000000000000000000000815260048101829052602481018690526044016107f2565b6040517fd0c8d23a000000000000000000000000000000000000000000000000000000008152600481018290526024810186905273ffffffffffffffffffffffffffffffffffffffff871660448201526064016107f2565b61327885846143f6565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515610f24565b600081831061332e5781610f24565b5090919050565b6060824710156133c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016107f2565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516133f091906144f6565b60006040518083038185875af1925050503d806000811461342d576040519150601f19603f3d011682016040523d82523d6000602084013e613432565b606091505b50915091506134438783838761344e565b979650505050505050565b606083156134e45782516000036134dd5773ffffffffffffffffffffffffffffffffffffffff85163b6134dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016107f2565b5081612c72565b612c7283838151156134f95781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f29190613663565b50805461353990613bad565b6000825580601f10613549575050565b601f01602090049060005260206000209081019061159491905b808211156135775760008155600101613563565b5090565b60006020828403121561358d57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610f2457600080fd5b803567ffffffffffffffff811681146135d557600080fd5b919050565b6000602082840312156135ec57600080fd5b610f24826135bd565b60005b838110156136105781810151838201526020016135f8565b50506000910152565b600081518084526136318160208601602086016135f5565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610f246020830184613619565b60006020828403121561368857600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461159457600080fd5b80356135d58161368f565b6000602082840312156136ce57600080fd5b8135610f248161368f565b6000602082840312156136eb57600080fd5b813567ffffffffffffffff81111561370257600080fd5b820160e08185031215610f2457600080fd5b60008083601f84011261372657600080fd5b50813567ffffffffffffffff81111561373e57600080fd5b6020830191508360208260051b850101111561375957600080fd5b9250929050565b6000806000806040858703121561377657600080fd5b843567ffffffffffffffff8082111561378e57600080fd5b61379a88838901613714565b909650945060208701359150808211156137b357600080fd5b506137c087828801613714565b95989497509550505050565b6000806000604084860312156137e157600080fd5b6137ea846135bd565b9250602084013567ffffffffffffffff8082111561380757600080fd5b818601915086601f83011261381b57600080fd5b81358181111561382a57600080fd5b87602082850101111561383c57600080fd5b6020830194508093505050509250925092565b6000806040838503121561386257600080fd5b61386b836135bd565b9150602083013561387b8161368f565b809150509250929050565b60006020828403121561389857600080fd5b813567ffffffffffffffff8111156138af57600080fd5b820160a08185031215610f2457600080fd5b6020815260008251604060208401526138dd6060840182613619565b905060208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160408501526139188282613619565b95945050505050565b6020808252825182820181905260009190848201906040850190845b8181101561396f57835173ffffffffffffffffffffffffffffffffffffffff168352928401929184019160010161393d565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561396f57835167ffffffffffffffff1683529284019291840191600101613997565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715613a0f57613a0f6139bd565b60405290565b60405160a0810167ffffffffffffffff81118282101715613a0f57613a0f6139bd565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613a7f57613a7f6139bd565b604052919050565b801515811461159457600080fd5b80356fffffffffffffffffffffffffffffffff811681146135d557600080fd5b600060608284031215613ac757600080fd5b6040516060810181811067ffffffffffffffff82111715613aea57613aea6139bd565b6040529050808235613afb81613a87565b8152613b0960208401613a95565b6020820152613b1a60408401613a95565b60408201525092915050565b600080600060e08486031215613b3b57600080fd5b613b44846135bd565b9250613b538560208601613ab5565b9150613b628560808601613ab5565b90509250925092565b60008060208385031215613b7e57600080fd5b823567ffffffffffffffff811115613b9557600080fd5b613ba185828601613714565b90969095509350505050565b600181811c90821680613bc157607f821691505b602082108103613bfa577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b600060208284031215613c1257600080fd5b5051919050565b600067ffffffffffffffff821115613c3357613c336139bd565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112613c7057600080fd5b8135613c83613c7e82613c19565b613a38565b818152846020838601011115613c9857600080fd5b816020850160208301376000918101602001919091529392505050565b600060e08236031215613cc757600080fd5b613ccf6139ec565b823567ffffffffffffffff80821115613ce757600080fd5b613cf336838701613c5f565b8352613d01602086016135bd565b6020840152613d12604086016136b1565b6040840152606085013560608401526080850135915080821115613d3557600080fd5b613d4136838701613c5f565b608084015260a0850135915080821115613d5a57600080fd5b613d6636838701613c5f565b60a084015260c0850135915080821115613d7f57600080fd5b50613d8c36828601613c5f565b60c08301525092915050565b601f821115611462576000816000526020600020601f850160051c81016020861015613dc15750805b601f850160051c820191505b81811015613de057828155600101613dcd565b505050505050565b67ffffffffffffffff831115613e0057613e006139bd565b613e1483613e0e8354613bad565b83613d98565b6000601f841160018114613e665760008515613e305750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355611d97565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b82811015613eb55786850135825560209485019460019092019101613e95565b5086821015613ef0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b604081526000613f156040830186613619565b82810360208401528381528385602083013760006020858301015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f860116820101915050949350505050565b600060208284031215613f7857600080fd5b8151610f2481613a87565b600060a08236031215613f9557600080fd5b613f9d613a15565b823567ffffffffffffffff811115613fb457600080fd5b613fc036828601613c5f565b825250613fcf602084016135bd565b60208201526040830135613fe28161368f565b6040820152606083810135908201526080830135613fff8161368f565b608082015292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee183360301811261406d57600080fd5b9190910192915050565b6000610120823603121561408a57600080fd5b614092613a15565b61409b836135bd565b815260208301356140ab81613a87565b6020820152604083013567ffffffffffffffff8111156140ca57600080fd5b6140d636828601613c5f565b6040830152506140e93660608501613ab5565b6060820152613fff3660c08501613ab5565b815167ffffffffffffffff811115614115576141156139bd565b614129816141238454613bad565b84613d98565b602080601f83116001811461417c57600084156141465750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555613de0565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156141c9578886015182559484019460019091019084016141aa565b508582101561420557878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b67ffffffffffffffff8416815260e0810161426160208301858051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b82511515608083015260208301516fffffffffffffffffffffffffffffffff90811660a084015260408401511660c0830152612c72565b60a0815260006142ab60a0830187613619565b73ffffffffffffffffffffffffffffffffffffffff8616602084015284604084015267ffffffffffffffff841660608401528281036080840152600081526020810191505095945050505050565b73ffffffffffffffffffffffffffffffffffffffff8516815260a06020820152600061432860a0830186613619565b60408301949094525067ffffffffffffffff9190911660608201528082036080909101526000815260200192915050565b60006020828403121561436b57600080fd5b815167ffffffffffffffff81111561438257600080fd5b8201601f8101841361439357600080fd5b80516143a1613c7e82613c19565b8181528560208385010111156143b657600080fd5b6139188260208301602086016135f5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156106ed576106ed6143c7565b606081016106ed82848051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b60006020828403121561445757600080fd5b8151610f248161368f565b80820281158282048414176106ed576106ed6143c7565b808201808211156106ed576106ed6143c7565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000826144f1577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000825161406d8184602087016135f556fe4c6f636b52656c65617365546f6b656e506f6f6c416e6450726f787920312e352e302d646576a164736f6c6343000818000a",
}

var LockReleaseTokenPoolAndProxyABI = LockReleaseTokenPoolAndProxyMetaData.ABI

var LockReleaseTokenPoolAndProxyBin = LockReleaseTokenPoolAndProxyMetaData.Bin

func DeployLockReleaseTokenPoolAndProxy(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, allowlist []common.Address, rmnProxy common.Address, acceptLiquidity bool, router common.Address) (common.Address, *types.Transaction, *LockReleaseTokenPoolAndProxy, error) {
	parsed, err := LockReleaseTokenPoolAndProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LockReleaseTokenPoolAndProxyBin), backend, token, allowlist, rmnProxy, acceptLiquidity, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LockReleaseTokenPoolAndProxy{address: address, abi: *parsed, LockReleaseTokenPoolAndProxyCaller: LockReleaseTokenPoolAndProxyCaller{contract: contract}, LockReleaseTokenPoolAndProxyTransactor: LockReleaseTokenPoolAndProxyTransactor{contract: contract}, LockReleaseTokenPoolAndProxyFilterer: LockReleaseTokenPoolAndProxyFilterer{contract: contract}}, nil
}

type LockReleaseTokenPoolAndProxy struct {
	address common.Address
	abi     abi.ABI
	LockReleaseTokenPoolAndProxyCaller
	LockReleaseTokenPoolAndProxyTransactor
	LockReleaseTokenPoolAndProxyFilterer
}

type LockReleaseTokenPoolAndProxyCaller struct {
	contract *bind.BoundContract
}

type LockReleaseTokenPoolAndProxyTransactor struct {
	contract *bind.BoundContract
}

type LockReleaseTokenPoolAndProxyFilterer struct {
	contract *bind.BoundContract
}

type LockReleaseTokenPoolAndProxySession struct {
	Contract     *LockReleaseTokenPoolAndProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type LockReleaseTokenPoolAndProxyCallerSession struct {
	Contract *LockReleaseTokenPoolAndProxyCaller
	CallOpts bind.CallOpts
}

type LockReleaseTokenPoolAndProxyTransactorSession struct {
	Contract     *LockReleaseTokenPoolAndProxyTransactor
	TransactOpts bind.TransactOpts
}

type LockReleaseTokenPoolAndProxyRaw struct {
	Contract *LockReleaseTokenPoolAndProxy
}

type LockReleaseTokenPoolAndProxyCallerRaw struct {
	Contract *LockReleaseTokenPoolAndProxyCaller
}

type LockReleaseTokenPoolAndProxyTransactorRaw struct {
	Contract *LockReleaseTokenPoolAndProxyTransactor
}

func NewLockReleaseTokenPoolAndProxy(address common.Address, backend bind.ContractBackend) (*LockReleaseTokenPoolAndProxy, error) {
	abi, err := abi.JSON(strings.NewReader(LockReleaseTokenPoolAndProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindLockReleaseTokenPoolAndProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxy{address: address, abi: abi, LockReleaseTokenPoolAndProxyCaller: LockReleaseTokenPoolAndProxyCaller{contract: contract}, LockReleaseTokenPoolAndProxyTransactor: LockReleaseTokenPoolAndProxyTransactor{contract: contract}, LockReleaseTokenPoolAndProxyFilterer: LockReleaseTokenPoolAndProxyFilterer{contract: contract}}, nil
}

func NewLockReleaseTokenPoolAndProxyCaller(address common.Address, caller bind.ContractCaller) (*LockReleaseTokenPoolAndProxyCaller, error) {
	contract, err := bindLockReleaseTokenPoolAndProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyCaller{contract: contract}, nil
}

func NewLockReleaseTokenPoolAndProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*LockReleaseTokenPoolAndProxyTransactor, error) {
	contract, err := bindLockReleaseTokenPoolAndProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyTransactor{contract: contract}, nil
}

func NewLockReleaseTokenPoolAndProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*LockReleaseTokenPoolAndProxyFilterer, error) {
	contract, err := bindLockReleaseTokenPoolAndProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyFilterer{contract: contract}, nil
}

func bindLockReleaseTokenPoolAndProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LockReleaseTokenPoolAndProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LockReleaseTokenPoolAndProxy.Contract.LockReleaseTokenPoolAndProxyCaller.contract.Call(opts, result, method, params...)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.LockReleaseTokenPoolAndProxyTransactor.contract.Transfer(opts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.LockReleaseTokenPoolAndProxyTransactor.contract.Transact(opts, method, params...)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LockReleaseTokenPoolAndProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.contract.Transfer(opts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.contract.Transact(opts, method, params...)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) CanAcceptLiquidity(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "canAcceptLiquidity")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) CanAcceptLiquidity() (bool, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.CanAcceptLiquidity(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) CanAcceptLiquidity() (bool, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.CanAcceptLiquidity(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) GetAllowList() ([]common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetAllowList(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) GetAllowList() ([]common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetAllowList(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) GetAllowListEnabled() (bool, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetAllowListEnabled(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) GetAllowListEnabled() (bool, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetAllowListEnabled(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetCurrentInboundRateLimiterState(&_LockReleaseTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetCurrentInboundRateLimiterState(&_LockReleaseTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetCurrentOutboundRateLimiterState(&_LockReleaseTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetCurrentOutboundRateLimiterState(&_LockReleaseTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) GetLockReleaseInterfaceId(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "getLockReleaseInterfaceId")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) GetLockReleaseInterfaceId() ([4]byte, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetLockReleaseInterfaceId(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) GetLockReleaseInterfaceId() ([4]byte, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetLockReleaseInterfaceId(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) GetOnRamp(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "getOnRamp", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) GetOnRamp(arg0 uint64) (common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetOnRamp(&_LockReleaseTokenPoolAndProxy.CallOpts, arg0)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) GetOnRamp(arg0 uint64) (common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetOnRamp(&_LockReleaseTokenPoolAndProxy.CallOpts, arg0)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) GetRateLimitAdmin() (common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetRateLimitAdmin(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetRateLimitAdmin(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) GetRebalancer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "getRebalancer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) GetRebalancer() (common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetRebalancer(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) GetRebalancer() (common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetRebalancer(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) GetRemotePool(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "getRemotePool", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) GetRemotePool(remoteChainSelector uint64) ([]byte, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetRemotePool(&_LockReleaseTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) GetRemotePool(remoteChainSelector uint64) ([]byte, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetRemotePool(&_LockReleaseTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) GetRmnProxy() (common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetRmnProxy(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) GetRmnProxy() (common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetRmnProxy(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) GetRouter() (common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetRouter(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) GetRouter() (common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetRouter(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) GetSupportedChains() ([]uint64, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetSupportedChains(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) GetSupportedChains() ([]uint64, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetSupportedChains(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) GetToken() (common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetToken(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) GetToken() (common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.GetToken(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) IsOffRamp(opts *bind.CallOpts, sourceChainSelector uint64, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "isOffRamp", sourceChainSelector, offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) IsOffRamp(sourceChainSelector uint64, offRamp common.Address) (bool, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.IsOffRamp(&_LockReleaseTokenPoolAndProxy.CallOpts, sourceChainSelector, offRamp)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) IsOffRamp(sourceChainSelector uint64, offRamp common.Address) (bool, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.IsOffRamp(&_LockReleaseTokenPoolAndProxy.CallOpts, sourceChainSelector, offRamp)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.IsSupportedChain(&_LockReleaseTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.IsSupportedChain(&_LockReleaseTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) IsSupportedToken(token common.Address) (bool, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.IsSupportedToken(&_LockReleaseTokenPoolAndProxy.CallOpts, token)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.IsSupportedToken(&_LockReleaseTokenPoolAndProxy.CallOpts, token)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) Owner() (common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.Owner(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) Owner() (common.Address, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.Owner(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.SupportsInterface(&_LockReleaseTokenPoolAndProxy.CallOpts, interfaceId)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.SupportsInterface(&_LockReleaseTokenPoolAndProxy.CallOpts, interfaceId)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LockReleaseTokenPoolAndProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) TypeAndVersion() (string, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.TypeAndVersion(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyCallerSession) TypeAndVersion() (string, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.TypeAndVersion(&_LockReleaseTokenPoolAndProxy.CallOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.contract.Transact(opts, "acceptOwnership")
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.AcceptOwnership(&_LockReleaseTokenPoolAndProxy.TransactOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.AcceptOwnership(&_LockReleaseTokenPoolAndProxy.TransactOpts)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.ApplyAllowListUpdates(&_LockReleaseTokenPoolAndProxy.TransactOpts, removes, adds)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.ApplyAllowListUpdates(&_LockReleaseTokenPoolAndProxy.TransactOpts, removes, adds)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactor) ApplyChainUpdates(opts *bind.TransactOpts, chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.contract.Transact(opts, "applyChainUpdates", chains)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) ApplyChainUpdates(chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.ApplyChainUpdates(&_LockReleaseTokenPoolAndProxy.TransactOpts, chains)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorSession) ApplyChainUpdates(chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.ApplyChainUpdates(&_LockReleaseTokenPoolAndProxy.TransactOpts, chains)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.LockOrBurn(&_LockReleaseTokenPoolAndProxy.TransactOpts, lockOrBurnIn)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.LockOrBurn(&_LockReleaseTokenPoolAndProxy.TransactOpts, lockOrBurnIn)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactor) ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.contract.Transact(opts, "provideLiquidity", amount)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.ProvideLiquidity(&_LockReleaseTokenPoolAndProxy.TransactOpts, amount)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.ProvideLiquidity(&_LockReleaseTokenPoolAndProxy.TransactOpts, amount)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.ReleaseOrMint(&_LockReleaseTokenPoolAndProxy.TransactOpts, releaseOrMintIn)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.ReleaseOrMint(&_LockReleaseTokenPoolAndProxy.TransactOpts, releaseOrMintIn)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.SetChainRateLimiterConfig(&_LockReleaseTokenPoolAndProxy.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.SetChainRateLimiterConfig(&_LockReleaseTokenPoolAndProxy.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactor) SetPreviousPool(opts *bind.TransactOpts, prevPool common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.contract.Transact(opts, "setPreviousPool", prevPool)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) SetPreviousPool(prevPool common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.SetPreviousPool(&_LockReleaseTokenPoolAndProxy.TransactOpts, prevPool)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorSession) SetPreviousPool(prevPool common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.SetPreviousPool(&_LockReleaseTokenPoolAndProxy.TransactOpts, prevPool)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.SetRateLimitAdmin(&_LockReleaseTokenPoolAndProxy.TransactOpts, rateLimitAdmin)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.SetRateLimitAdmin(&_LockReleaseTokenPoolAndProxy.TransactOpts, rateLimitAdmin)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactor) SetRebalancer(opts *bind.TransactOpts, rebalancer common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.contract.Transact(opts, "setRebalancer", rebalancer)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) SetRebalancer(rebalancer common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.SetRebalancer(&_LockReleaseTokenPoolAndProxy.TransactOpts, rebalancer)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorSession) SetRebalancer(rebalancer common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.SetRebalancer(&_LockReleaseTokenPoolAndProxy.TransactOpts, rebalancer)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactor) SetRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.contract.Transact(opts, "setRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) SetRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.SetRemotePool(&_LockReleaseTokenPoolAndProxy.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorSession) SetRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.SetRemotePool(&_LockReleaseTokenPoolAndProxy.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.contract.Transact(opts, "setRouter", newRouter)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.SetRouter(&_LockReleaseTokenPoolAndProxy.TransactOpts, newRouter)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.SetRouter(&_LockReleaseTokenPoolAndProxy.TransactOpts, newRouter)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.TransferOwnership(&_LockReleaseTokenPoolAndProxy.TransactOpts, to)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.TransferOwnership(&_LockReleaseTokenPoolAndProxy.TransactOpts, to)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactor) WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.contract.Transact(opts, "withdrawLiquidity", amount)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxySession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.WithdrawLiquidity(&_LockReleaseTokenPoolAndProxy.TransactOpts, amount)
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyTransactorSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPoolAndProxy.Contract.WithdrawLiquidity(&_LockReleaseTokenPoolAndProxy.TransactOpts, amount)
}

type LockReleaseTokenPoolAndProxyAllowListAddIterator struct {
	Event *LockReleaseTokenPoolAndProxyAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyAllowListAdd)
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
		it.Event = new(LockReleaseTokenPoolAndProxyAllowListAdd)
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

func (it *LockReleaseTokenPoolAndProxyAllowListAddIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyAllowListAddIterator, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyAllowListAddIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyAllowListAdd)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseAllowListAdd(log types.Log) (*LockReleaseTokenPoolAndProxyAllowListAdd, error) {
	event := new(LockReleaseTokenPoolAndProxyAllowListAdd)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyAllowListRemoveIterator struct {
	Event *LockReleaseTokenPoolAndProxyAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyAllowListRemove)
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
		it.Event = new(LockReleaseTokenPoolAndProxyAllowListRemove)
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

func (it *LockReleaseTokenPoolAndProxyAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyAllowListRemoveIterator, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyAllowListRemoveIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyAllowListRemove)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseAllowListRemove(log types.Log) (*LockReleaseTokenPoolAndProxyAllowListRemove, error) {
	event := new(LockReleaseTokenPoolAndProxyAllowListRemove)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyBurnedIterator struct {
	Event *LockReleaseTokenPoolAndProxyBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyBurned)
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
		it.Event = new(LockReleaseTokenPoolAndProxyBurned)
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

func (it *LockReleaseTokenPoolAndProxyBurnedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*LockReleaseTokenPoolAndProxyBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyBurnedIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyBurned)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "Burned", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseBurned(log types.Log) (*LockReleaseTokenPoolAndProxyBurned, error) {
	event := new(LockReleaseTokenPoolAndProxyBurned)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyChainAddedIterator struct {
	Event *LockReleaseTokenPoolAndProxyChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyChainAdded)
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
		it.Event = new(LockReleaseTokenPoolAndProxyChainAdded)
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

func (it *LockReleaseTokenPoolAndProxyChainAddedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyChainAdded struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterChainAdded(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyChainAddedIterator, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyChainAddedIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyChainAdded) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyChainAdded)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseChainAdded(log types.Log) (*LockReleaseTokenPoolAndProxyChainAdded, error) {
	event := new(LockReleaseTokenPoolAndProxyChainAdded)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyChainConfiguredIterator struct {
	Event *LockReleaseTokenPoolAndProxyChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyChainConfigured)
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
		it.Event = new(LockReleaseTokenPoolAndProxyChainConfigured)
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

func (it *LockReleaseTokenPoolAndProxyChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyChainConfiguredIterator, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyChainConfiguredIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyChainConfigured) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyChainConfigured)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseChainConfigured(log types.Log) (*LockReleaseTokenPoolAndProxyChainConfigured, error) {
	event := new(LockReleaseTokenPoolAndProxyChainConfigured)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyChainRemovedIterator struct {
	Event *LockReleaseTokenPoolAndProxyChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyChainRemoved)
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
		it.Event = new(LockReleaseTokenPoolAndProxyChainRemoved)
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

func (it *LockReleaseTokenPoolAndProxyChainRemovedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyChainRemovedIterator, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyChainRemovedIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyChainRemoved) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyChainRemoved)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseChainRemoved(log types.Log) (*LockReleaseTokenPoolAndProxyChainRemoved, error) {
	event := new(LockReleaseTokenPoolAndProxyChainRemoved)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyConfigChangedIterator struct {
	Event *LockReleaseTokenPoolAndProxyConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyConfigChanged)
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
		it.Event = new(LockReleaseTokenPoolAndProxyConfigChanged)
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

func (it *LockReleaseTokenPoolAndProxyConfigChangedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyConfigChangedIterator, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyConfigChangedIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyConfigChanged) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyConfigChanged)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseConfigChanged(log types.Log) (*LockReleaseTokenPoolAndProxyConfigChanged, error) {
	event := new(LockReleaseTokenPoolAndProxyConfigChanged)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyLegacyPoolChangedIterator struct {
	Event *LockReleaseTokenPoolAndProxyLegacyPoolChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyLegacyPoolChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyLegacyPoolChanged)
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
		it.Event = new(LockReleaseTokenPoolAndProxyLegacyPoolChanged)
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

func (it *LockReleaseTokenPoolAndProxyLegacyPoolChangedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyLegacyPoolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyLegacyPoolChanged struct {
	OldPool common.Address
	NewPool common.Address
	Raw     types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterLegacyPoolChanged(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyLegacyPoolChangedIterator, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "LegacyPoolChanged")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyLegacyPoolChangedIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "LegacyPoolChanged", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchLegacyPoolChanged(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyLegacyPoolChanged) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "LegacyPoolChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyLegacyPoolChanged)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "LegacyPoolChanged", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseLegacyPoolChanged(log types.Log) (*LockReleaseTokenPoolAndProxyLegacyPoolChanged, error) {
	event := new(LockReleaseTokenPoolAndProxyLegacyPoolChanged)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "LegacyPoolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyLiquidityAddedIterator struct {
	Event *LockReleaseTokenPoolAndProxyLiquidityAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyLiquidityAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyLiquidityAdded)
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
		it.Event = new(LockReleaseTokenPoolAndProxyLiquidityAdded)
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

func (it *LockReleaseTokenPoolAndProxyLiquidityAddedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyLiquidityAdded struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolAndProxyLiquidityAddedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "LiquidityAdded", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyLiquidityAddedIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyLiquidityAdded, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "LiquidityAdded", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyLiquidityAdded)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseLiquidityAdded(log types.Log) (*LockReleaseTokenPoolAndProxyLiquidityAdded, error) {
	event := new(LockReleaseTokenPoolAndProxyLiquidityAdded)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyLiquidityRemovedIterator struct {
	Event *LockReleaseTokenPoolAndProxyLiquidityRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyLiquidityRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyLiquidityRemoved)
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
		it.Event = new(LockReleaseTokenPoolAndProxyLiquidityRemoved)
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

func (it *LockReleaseTokenPoolAndProxyLiquidityRemovedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyLiquidityRemoved struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolAndProxyLiquidityRemovedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "LiquidityRemoved", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyLiquidityRemovedIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyLiquidityRemoved, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "LiquidityRemoved", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyLiquidityRemoved)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseLiquidityRemoved(log types.Log) (*LockReleaseTokenPoolAndProxyLiquidityRemoved, error) {
	event := new(LockReleaseTokenPoolAndProxyLiquidityRemoved)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyLockedIterator struct {
	Event *LockReleaseTokenPoolAndProxyLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyLocked)
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
		it.Event = new(LockReleaseTokenPoolAndProxyLocked)
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

func (it *LockReleaseTokenPoolAndProxyLockedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*LockReleaseTokenPoolAndProxyLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyLockedIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyLocked)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "Locked", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseLocked(log types.Log) (*LockReleaseTokenPoolAndProxyLocked, error) {
	event := new(LockReleaseTokenPoolAndProxyLocked)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyMintedIterator struct {
	Event *LockReleaseTokenPoolAndProxyMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyMinted)
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
		it.Event = new(LockReleaseTokenPoolAndProxyMinted)
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

func (it *LockReleaseTokenPoolAndProxyMintedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LockReleaseTokenPoolAndProxyMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyMintedIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyMinted)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "Minted", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseMinted(log types.Log) (*LockReleaseTokenPoolAndProxyMinted, error) {
	event := new(LockReleaseTokenPoolAndProxyMinted)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyOwnershipTransferRequestedIterator struct {
	Event *LockReleaseTokenPoolAndProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyOwnershipTransferRequested)
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
		it.Event = new(LockReleaseTokenPoolAndProxyOwnershipTransferRequested)
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

func (it *LockReleaseTokenPoolAndProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockReleaseTokenPoolAndProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyOwnershipTransferRequestedIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyOwnershipTransferRequested)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*LockReleaseTokenPoolAndProxyOwnershipTransferRequested, error) {
	event := new(LockReleaseTokenPoolAndProxyOwnershipTransferRequested)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyOwnershipTransferredIterator struct {
	Event *LockReleaseTokenPoolAndProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyOwnershipTransferred)
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
		it.Event = new(LockReleaseTokenPoolAndProxyOwnershipTransferred)
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

func (it *LockReleaseTokenPoolAndProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockReleaseTokenPoolAndProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyOwnershipTransferredIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyOwnershipTransferred)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseOwnershipTransferred(log types.Log) (*LockReleaseTokenPoolAndProxyOwnershipTransferred, error) {
	event := new(LockReleaseTokenPoolAndProxyOwnershipTransferred)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyReleasedIterator struct {
	Event *LockReleaseTokenPoolAndProxyReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyReleased)
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
		it.Event = new(LockReleaseTokenPoolAndProxyReleased)
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

func (it *LockReleaseTokenPoolAndProxyReleasedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LockReleaseTokenPoolAndProxyReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyReleasedIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyReleased)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "Released", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseReleased(log types.Log) (*LockReleaseTokenPoolAndProxyReleased, error) {
	event := new(LockReleaseTokenPoolAndProxyReleased)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyRemotePoolSetIterator struct {
	Event *LockReleaseTokenPoolAndProxyRemotePoolSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyRemotePoolSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyRemotePoolSet)
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
		it.Event = new(LockReleaseTokenPoolAndProxyRemotePoolSet)
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

func (it *LockReleaseTokenPoolAndProxyRemotePoolSetIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyRemotePoolSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyRemotePoolSet struct {
	RemoteChainSelector uint64
	PreviousPoolAddress []byte
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterRemotePoolSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolAndProxyRemotePoolSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "RemotePoolSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyRemotePoolSetIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "RemotePoolSet", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchRemotePoolSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyRemotePoolSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "RemotePoolSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyRemotePoolSet)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "RemotePoolSet", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseRemotePoolSet(log types.Log) (*LockReleaseTokenPoolAndProxyRemotePoolSet, error) {
	event := new(LockReleaseTokenPoolAndProxyRemotePoolSet)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "RemotePoolSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyRouterUpdatedIterator struct {
	Event *LockReleaseTokenPoolAndProxyRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyRouterUpdated)
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
		it.Event = new(LockReleaseTokenPoolAndProxyRouterUpdated)
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

func (it *LockReleaseTokenPoolAndProxyRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyRouterUpdatedIterator, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyRouterUpdatedIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyRouterUpdated)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseRouterUpdated(log types.Log) (*LockReleaseTokenPoolAndProxyRouterUpdated, error) {
	event := new(LockReleaseTokenPoolAndProxyRouterUpdated)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAndProxyTokensConsumedIterator struct {
	Event *LockReleaseTokenPoolAndProxyTokensConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAndProxyTokensConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAndProxyTokensConsumed)
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
		it.Event = new(LockReleaseTokenPoolAndProxyTokensConsumed)
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

func (it *LockReleaseTokenPoolAndProxyTokensConsumedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAndProxyTokensConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAndProxyTokensConsumed struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) FilterTokensConsumed(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyTokensConsumedIterator, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.FilterLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAndProxyTokensConsumedIterator{contract: _LockReleaseTokenPoolAndProxy.contract, event: "TokensConsumed", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyTokensConsumed) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPoolAndProxy.contract.WatchLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAndProxyTokensConsumed)
				if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
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

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxyFilterer) ParseTokensConsumed(log types.Log) (*LockReleaseTokenPoolAndProxyTokensConsumed, error) {
	event := new(LockReleaseTokenPoolAndProxyTokensConsumed)
	if err := _LockReleaseTokenPoolAndProxy.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxy) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _LockReleaseTokenPoolAndProxy.abi.Events["AllowListAdd"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseAllowListAdd(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["AllowListRemove"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseAllowListRemove(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["Burned"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseBurned(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["ChainAdded"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseChainAdded(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["ChainConfigured"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseChainConfigured(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["ChainRemoved"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseChainRemoved(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["ConfigChanged"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseConfigChanged(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["LegacyPoolChanged"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseLegacyPoolChanged(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["LiquidityAdded"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseLiquidityAdded(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["LiquidityRemoved"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseLiquidityRemoved(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["Locked"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseLocked(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["Minted"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseMinted(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["OwnershipTransferRequested"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseOwnershipTransferRequested(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["OwnershipTransferred"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseOwnershipTransferred(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["Released"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseReleased(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["RemotePoolSet"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseRemotePoolSet(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["RouterUpdated"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseRouterUpdated(log)
	case _LockReleaseTokenPoolAndProxy.abi.Events["TokensConsumed"].ID:
		return _LockReleaseTokenPoolAndProxy.ParseTokensConsumed(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (LockReleaseTokenPoolAndProxyAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (LockReleaseTokenPoolAndProxyAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (LockReleaseTokenPoolAndProxyBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
}

func (LockReleaseTokenPoolAndProxyChainAdded) Topic() common.Hash {
	return common.HexToHash("0x0f135cbb9afa12a8bf3bbd071c117bcca4ddeca6160ef7f33d012a81b9c0c471")
}

func (LockReleaseTokenPoolAndProxyChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (LockReleaseTokenPoolAndProxyChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (LockReleaseTokenPoolAndProxyConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (LockReleaseTokenPoolAndProxyLegacyPoolChanged) Topic() common.Hash {
	return common.HexToHash("0x81accd0a7023865eaa51b3399dd0eafc488bf3ba238402911e1659cfe860f228")
}

func (LockReleaseTokenPoolAndProxyLiquidityAdded) Topic() common.Hash {
	return common.HexToHash("0xc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088")
}

func (LockReleaseTokenPoolAndProxyLiquidityRemoved) Topic() common.Hash {
	return common.HexToHash("0xc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719")
}

func (LockReleaseTokenPoolAndProxyLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (LockReleaseTokenPoolAndProxyMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (LockReleaseTokenPoolAndProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (LockReleaseTokenPoolAndProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (LockReleaseTokenPoolAndProxyReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (LockReleaseTokenPoolAndProxyRemotePoolSet) Topic() common.Hash {
	return common.HexToHash("0xdb4d6220746a38cbc5335f7e108f7de80f482f4d23350253dfd0917df75a14bf")
}

func (LockReleaseTokenPoolAndProxyRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (LockReleaseTokenPoolAndProxyTokensConsumed) Topic() common.Hash {
	return common.HexToHash("0x1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a")
}

func (_LockReleaseTokenPoolAndProxy *LockReleaseTokenPoolAndProxy) Address() common.Address {
	return _LockReleaseTokenPoolAndProxy.address
}

type LockReleaseTokenPoolAndProxyInterface interface {
	CanAcceptLiquidity(opts *bind.CallOpts) (bool, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetLockReleaseInterfaceId(opts *bind.CallOpts) ([4]byte, error)

	GetOnRamp(opts *bind.CallOpts, arg0 uint64) (common.Address, error)

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRebalancer(opts *bind.CallOpts) (common.Address, error)

	GetRemotePool(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	IsOffRamp(opts *bind.CallOpts, sourceChainSelector uint64, offRamp common.Address) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, chains []TokenPoolChainUpdate) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetPreviousPool(opts *bind.TransactOpts, prevPool common.Address) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRebalancer(opts *bind.TransactOpts, rebalancer common.Address) (*types.Transaction, error)

	SetRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*LockReleaseTokenPoolAndProxyAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*LockReleaseTokenPoolAndProxyAllowListRemove, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*LockReleaseTokenPoolAndProxyBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*LockReleaseTokenPoolAndProxyBurned, error)

	FilterChainAdded(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*LockReleaseTokenPoolAndProxyChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*LockReleaseTokenPoolAndProxyChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*LockReleaseTokenPoolAndProxyChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*LockReleaseTokenPoolAndProxyConfigChanged, error)

	FilterLegacyPoolChanged(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyLegacyPoolChangedIterator, error)

	WatchLegacyPoolChanged(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyLegacyPoolChanged) (event.Subscription, error)

	ParseLegacyPoolChanged(log types.Log) (*LockReleaseTokenPoolAndProxyLegacyPoolChanged, error)

	FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolAndProxyLiquidityAddedIterator, error)

	WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyLiquidityAdded, provider []common.Address, amount []*big.Int) (event.Subscription, error)

	ParseLiquidityAdded(log types.Log) (*LockReleaseTokenPoolAndProxyLiquidityAdded, error)

	FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolAndProxyLiquidityRemovedIterator, error)

	WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyLiquidityRemoved, provider []common.Address, amount []*big.Int) (event.Subscription, error)

	ParseLiquidityRemoved(log types.Log) (*LockReleaseTokenPoolAndProxyLiquidityRemoved, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*LockReleaseTokenPoolAndProxyLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*LockReleaseTokenPoolAndProxyLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LockReleaseTokenPoolAndProxyMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*LockReleaseTokenPoolAndProxyMinted, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockReleaseTokenPoolAndProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*LockReleaseTokenPoolAndProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockReleaseTokenPoolAndProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*LockReleaseTokenPoolAndProxyOwnershipTransferred, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LockReleaseTokenPoolAndProxyReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*LockReleaseTokenPoolAndProxyReleased, error)

	FilterRemotePoolSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolAndProxyRemotePoolSetIterator, error)

	WatchRemotePoolSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyRemotePoolSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolSet(log types.Log) (*LockReleaseTokenPoolAndProxyRemotePoolSet, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*LockReleaseTokenPoolAndProxyRouterUpdated, error)

	FilterTokensConsumed(opts *bind.FilterOpts) (*LockReleaseTokenPoolAndProxyTokensConsumedIterator, error)

	WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAndProxyTokensConsumed) (event.Subscription, error)

	ParseTokensConsumed(log types.Log) (*LockReleaseTokenPoolAndProxyTokensConsumed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
