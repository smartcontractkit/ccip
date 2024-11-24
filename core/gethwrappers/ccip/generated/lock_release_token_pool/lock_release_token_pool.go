// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lock_release_token_pool

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
	zkSyncAccounts "github.com/zksync-sdk/zksync2-go/accounts"
	zkSyncClient "github.com/zksync-sdk/zksync2-go/clients"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
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
	DestTokenAddress []byte
	DestPoolData     []byte
}

type PoolReleaseOrMintInV1 struct {
	OriginalSender      []byte
	RemoteChainSelector uint64
	Receiver            common.Address
	Amount              *big.Int
	LocalToken          common.Address
	SourcePoolAddress   []byte
	SourcePoolData      []byte
	OffchainTokenData   []byte
}

type PoolReleaseOrMintOutV1 struct {
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
	RemoteTokenAddress        []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

var LockReleaseTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"acceptLiquidity\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowListNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotARampOnRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"DisabledNonZeroRateLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"InvalidRateLimitRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidSourcePoolAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidityNotAccepted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"NonExistentChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateLimitMustBeDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remoteToken\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"previousPoolAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"RemotePoolSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRouter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensConsumed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"remoteTokenAddress\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"internalType\":\"structTokenPool.ChainUpdate[]\",\"name\":\"chains\",\"type\":\"tuple[]\"}],\"name\":\"applyChainUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canAcceptLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentInboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentOutboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRateLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRebalancer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRemotePool\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRemoteToken\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRmnProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedChains\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"isSupportedChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isSupportedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"internalType\":\"structPool.LockOrBurnInV1\",\"name\":\"lockOrBurnIn\",\"type\":\"tuple\"}],\"name\":\"lockOrBurn\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destPoolData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.LockOrBurnOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"provideLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"originalSender\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainTokenData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.ReleaseOrMintInV1\",\"name\":\"releaseOrMintIn\",\"type\":\"tuple\"}],\"name\":\"releaseOrMint\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"destinationAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundConfig\",\"type\":\"tuple\"}],\"name\":\"setChainRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rateLimitAdmin\",\"type\":\"address\"}],\"name\":\"setRateLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rebalancer\",\"type\":\"address\"}],\"name\":\"setRebalancer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"setRemotePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620049ae380380620049ae833981016040819052620000359162000565565b848484833380600081620000905760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c357620000c3816200017e565b5050506001600160a01b0384161580620000e457506001600160a01b038116155b80620000f757506001600160a01b038216155b1562000116576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b0384811660805282811660a052600480546001600160a01b031916918316919091179055825115801560c052620001695760408051600081526020810190915262000169908462000229565b5050505090151560e05250620006d692505050565b336001600160a01b03821603620001d85760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000087565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60c0516200024a576040516335f4a7b360e01b815260040160405180910390fd5b60005b8251811015620002d55760008382815181106200026e576200026e62000688565b602090810291909101015190506200028860028262000386565b15620002cb576040516001600160a01b03821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b506001016200024d565b5060005b815181101562000381576000828281518110620002fa57620002fa62000688565b6020026020010151905060006001600160a01b0316816001600160a01b03160362000326575062000378565b62000333600282620003a6565b1562000376576040516001600160a01b03821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b600101620002d9565b505050565b60006200039d836001600160a01b038416620003bd565b90505b92915050565b60006200039d836001600160a01b038416620004c1565b60008181526001830160205260408120548015620004b6576000620003e46001836200069e565b8554909150600090620003fa906001906200069e565b9050808214620004665760008660000182815481106200041e576200041e62000688565b906000526020600020015490508087600001848154811062000444576200044462000688565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806200047a576200047a620006c0565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050620003a0565b6000915050620003a0565b60008181526001830160205260408120546200050a57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155620003a0565b506000620003a0565b6001600160a01b03811681146200052957600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b80516200054f8162000513565b919050565b805180151581146200054f57600080fd5b600080600080600060a086880312156200057e57600080fd5b85516200058b8162000513565b602087810151919650906001600160401b0380821115620005ab57600080fd5b818901915089601f830112620005c057600080fd5b815181811115620005d557620005d56200052c565b8060051b604051601f19603f83011681018181108582111715620005fd57620005fd6200052c565b60405291825284820192508381018501918c8311156200061c57600080fd5b938501935b828510156200064557620006358562000542565b8452938501939285019262000621565b8099505050505050506200065c6040870162000542565b92506200066c6060870162000554565b91506200067c6080870162000542565b90509295509295909350565b634e487b7160e01b600052603260045260246000fd5b81810381811115620003a057634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60805160a05160c05160e05161423c62000772600039600081816104ef015261174201526000818161059c01528181611cde015261278301526000818161057601528181611b0f0152611f94015260008181610290015281816102e50152818161077a0152818161084c015281816108ed0152818161180401528181611a2f01528181611eb401528181612719015261296e015261423c6000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c80638da5cb5b1161010f578063c4bffe2b116100a2578063dc0bd97111610071578063dc0bd97114610574578063e0351e131461059a578063eb521a4c146105c0578063f2fde38b146105d357600080fd5b8063c4bffe2b14610526578063c75eea9c1461053b578063cf7401f31461054e578063db6327dc1461056157600080fd5b8063b0f479a1116100de578063b0f479a1146104bc578063b7946580146104da578063bb98546b146104ed578063c0d786551461051357600080fd5b80638da5cb5b146103fa5780639a4575b914610418578063a7cd63b714610438578063af58d59f1461044d57600080fd5b806354c8a4f31161018757806378a010b21161015657806378a010b2146103b957806379ba5097146103cc5780637d54534e146103d45780638926f54f146103e757600080fd5b806354c8a4f31461036257806366320087146103755780636cfd1553146103885780636d3d1a581461039b57600080fd5b806321df0da7116101c357806321df0da71461028e578063240028e8146102d55780633907753714610322578063432a6ba31461034457600080fd5b806301ffc9a7146101f55780630a2fd4931461021d5780630a861f2a1461023d578063181f5a7714610252575b600080fd5b610208610203366004613318565b6105e6565b60405190151581526020015b60405180910390f35b61023061022b366004613377565b610642565b6040516102149190613400565b61025061024b366004613413565b6106f2565b005b6102306040518060400160405280601a81526020017f4c6f636b52656c65617365546f6b656e506f6f6c20312e352e3000000000000081525081565b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610214565b6102086102e3366004613459565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff90811691161490565b610335610330366004613476565b6108a3565b60405190518152602001610214565b60095473ffffffffffffffffffffffffffffffffffffffff166102b0565b6102506103703660046134fe565b6109a9565b61025061038336600461356a565b610a24565b610250610396366004613459565b610b00565b60085473ffffffffffffffffffffffffffffffffffffffff166102b0565b6102506103c7366004613596565b610b4f565b610250610cbe565b6102506103e2366004613459565b610dbb565b6102086103f5366004613377565b610e0a565b60005473ffffffffffffffffffffffffffffffffffffffff166102b0565b61042b610426366004613619565b610e21565b6040516102149190613654565b610440610ebb565b60405161021491906136b4565b61046061045b366004613377565b610ecc565b604051610214919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b60045473ffffffffffffffffffffffffffffffffffffffff166102b0565b6102306104e8366004613377565b610fa1565b7f0000000000000000000000000000000000000000000000000000000000000000610208565b610250610521366004613459565b610fcc565b61052e6110a7565b604051610214919061370e565b610460610549366004613377565b61115f565b61025061055c366004613876565b611231565b61025061056f3660046138bb565b6112ba565b7f00000000000000000000000000000000000000000000000000000000000000006102b0565b7f0000000000000000000000000000000000000000000000000000000000000000610208565b6102506105ce366004613413565b611740565b6102506105e1366004613459565b61185c565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167fe1d4056600000000000000000000000000000000000000000000000000000000148061063c575061063c82611870565b92915050565b67ffffffffffffffff8116600090815260076020526040902060040180546060919061066d906138fd565b80601f0160208091040260200160405190810160405280929190818152602001828054610699906138fd565b80156106e65780601f106106bb576101008083540402835291602001916106e6565b820191906000526020600020905b8154815290600101906020018083116106c957829003601f168201915b50505050509050919050565b60095473ffffffffffffffffffffffffffffffffffffffff16331461074a576040517f8e4a23d60000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015281907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa1580156107d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107fa9190613950565b1015610832576040517fbb55fd2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61087373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163383611954565b604051819033907fc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf984017171990600090a350565b6040805160208101909152600081526108c36108be83613a14565b611a28565b6109186108d66060840160408501613459565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906060850135611954565b6109286060830160408401613459565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52846060013560405161098a91815260200190565b60405180910390a3506040805160208101909152606090910135815290565b6109b1611c59565b610a1e84848080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808802828101820190935287825290935087925086918291850190849080828437600092019190915250611cdc92505050565b50505050565b610a2c611c59565b6040517f0a861f2a0000000000000000000000000000000000000000000000000000000081526004810182905273ffffffffffffffffffffffffffffffffffffffff831690630a861f2a90602401600060405180830381600087803b158015610a9457600080fd5b505af1158015610aa8573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff167f6fa7abcf1345d1d478e5ea0da6b5f26a90eadb0546ef15ed3833944fbfd1db6282604051610af491815260200190565b60405180910390a25050565b610b08611c59565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b610b57611c59565b610b6083610e0a565b610ba2576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602401610741565b67ffffffffffffffff831660009081526007602052604081206004018054610bc9906138fd565b80601f0160208091040260200160405190810160405280929190818152602001828054610bf5906138fd565b8015610c425780601f10610c1757610100808354040283529160200191610c42565b820191906000526020600020905b815481529060010190602001808311610c2557829003601f168201915b5050505067ffffffffffffffff8616600090815260076020526040902091925050600401610c71838583613b59565b508367ffffffffffffffff167fdb4d6220746a38cbc5335f7e108f7de80f482f4d23350253dfd0917df75a14bf828585604051610cb093929190613c74565b60405180910390a250505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610d3f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610741565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610dc3611c59565b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600061063c600567ffffffffffffffff8416611e92565b6040805180820190915260608082526020820152610e46610e4183613cd8565b611ead565b6040516060830135815233907f9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd600089060200160405180910390a26040518060400160405280610ea08460200160208101906104e89190613377565b81526040805160208181019092526000815291015292915050565b6060610ec76002612077565b905090565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845260028201546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff16151594820194909452600390910154808416606083015291909104909116608082015261063c90612084565b67ffffffffffffffff8116600090815260076020526040902060050180546060919061066d906138fd565b610fd4611c59565b73ffffffffffffffffffffffffffffffffffffffff8116611021576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684910160405180910390a15050565b606060006110b56005612077565b90506000815167ffffffffffffffff8111156110d3576110d3613750565b6040519080825280602002602001820160405280156110fc578160200160208202803683370190505b50905060005b82518110156111585782818151811061111d5761111d613d7a565b602002602001015182828151811061113757611137613d7a565b67ffffffffffffffff90921660209283029190910190910152600101611102565b5092915050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845281546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff16151594820194909452600190910154808416606083015291909104909116608082015261063c90612084565b60085473ffffffffffffffffffffffffffffffffffffffff163314801590611271575060005473ffffffffffffffffffffffffffffffffffffffff163314155b156112aa576040517f8e4a23d6000000000000000000000000000000000000000000000000000000008152336004820152602401610741565b6112b5838383612136565b505050565b6112c2611c59565b60005b818110156112b55760008383838181106112e1576112e1613d7a565b90506020028101906112f39190613da9565b6112fc90613de7565b90506113118160800151826020015115612220565b6113248160a00151826020015115612220565b8060200151156116205780516113469060059067ffffffffffffffff16612359565b61138b5780516040517f1d5ad3c500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610741565b60408101515115806113a05750606081015151155b156113d7576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805161012081018252608083810180516020908101516fffffffffffffffffffffffffffffffff9081168486019081524263ffffffff90811660a0808901829052865151151560c08a01528651860151851660e08a015295518901518416610100890152918752875180860189529489018051850151841686528585019290925281515115158589015281518401518316606080870191909152915188015183168587015283870194855288880151878901908152828a015183890152895167ffffffffffffffff1660009081526007865289902088518051825482890151838e01519289167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316177001000000000000000000000000000000009188168202177fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff90811674010000000000000000000000000000000000000000941515850217865584890151948d0151948a16948a168202949094176001860155995180516002860180549b8301519f830151918b169b9093169a909a179d9096168a029c909c179091169615150295909517909855908101519401519381169316909102919091176003820155915190919060048201906115b89082613e9b565b50606082015160058201906115cd9082613e9b565b505081516060830151608084015160a08501516040517f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c295506116139493929190613fb5565b60405180910390a1611737565b80516116389060059067ffffffffffffffff16612365565b61167d5780516040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610741565b805167ffffffffffffffff16600090815260076020526040812080547fffffffffffffffffffffff000000000000000000000000000000000000000000908116825560018201839055600282018054909116905560038101829055906116e660048301826132ca565b6116f46005830160006132ca565b5050805160405167ffffffffffffffff90911681527f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599169060200160405180910390a15b506001016112c5565b7f0000000000000000000000000000000000000000000000000000000000000000611797576040517fe93f8fa400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60095473ffffffffffffffffffffffffffffffffffffffff1633146117ea576040517f8e4a23d6000000000000000000000000000000000000000000000000000000008152336004820152602401610741565b61182c73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084612371565b604051819033907fc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb31208890600090a350565b611864611c59565b61186d816123cf565b50565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167faff2afbf00000000000000000000000000000000000000000000000000000000148061190357507fffffffff0000000000000000000000000000000000000000000000000000000082167f0e64dd2900000000000000000000000000000000000000000000000000000000145b8061063c57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a7000000000000000000000000000000000000000000000000000000001492915050565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526112b59084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526124c4565b60808101517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff908116911614611abd5760808101516040517f961c9a4f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610741565b60208101516040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815260809190911b77ffffffffffffffff000000000000000000000000000000001660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632cbc26bb90602401602060405180830381865afa158015611b6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b8f919061404e565b15611bc6576040517f53ad11d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611bd381602001516125d0565b6000611be28260200151610642565b9050805160001480611c06575080805190602001208260a001518051906020012014155b15611c43578160a001516040517f24eb47e50000000000000000000000000000000000000000000000000000000081526004016107419190613400565b611c55826020015183606001516126f6565b5050565b60005473ffffffffffffffffffffffffffffffffffffffff163314611cda576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610741565b565b7f0000000000000000000000000000000000000000000000000000000000000000611d33576040517f35f4a7b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8251811015611dc9576000838281518110611d5357611d53613d7a565b60200260200101519050611d7181600261273d90919063ffffffff16565b15611dc05760405173ffffffffffffffffffffffffffffffffffffffff821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b50600101611d36565b5060005b81518110156112b5576000828281518110611dea57611dea613d7a565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611e2e5750611e8a565b611e3960028261275f565b15611e885760405173ffffffffffffffffffffffffffffffffffffffff821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b600101611dcd565b600081815260018301602052604081205415155b9392505050565b60808101517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff908116911614611f425760808101516040517f961c9a4f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610741565b60208101516040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815260809190911b77ffffffffffffffff000000000000000000000000000000001660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632cbc26bb90602401602060405180830381865afa158015611ff0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612014919061404e565b1561204b576040517f53ad11d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6120588160400151612781565b6120658160200151612800565b61186d8160200151826060015161294e565b60606000611ea683612992565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261211282606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff16426120f6919061409a565b85608001516fffffffffffffffffffffffffffffffff166129ed565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b61213f83610e0a565b612181576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602401610741565b61218c826000612220565b67ffffffffffffffff831660009081526007602052604090206121af9083612a17565b6121ba816000612220565b67ffffffffffffffff831660009081526007602052604090206121e09060020182612a17565b7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b838383604051612213939291906140ad565b60405180910390a1505050565b8151156122e75781602001516fffffffffffffffffffffffffffffffff1682604001516fffffffffffffffffffffffffffffffff16101580612276575060408201516fffffffffffffffffffffffffffffffff16155b156122af57816040517f8020d1240000000000000000000000000000000000000000000000000000000081526004016107419190614130565b8015611c55576040517f433fc33d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408201516fffffffffffffffffffffffffffffffff16151580612320575060208201516fffffffffffffffffffffffffffffffff1615155b15611c5557816040517fd68af9cc0000000000000000000000000000000000000000000000000000000081526004016107419190614130565b6000611ea68383612bb9565b6000611ea68383612c08565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052610a1e9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016119a6565b3373ffffffffffffffffffffffffffffffffffffffff82160361244e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610741565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000612526826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16612cfb9092919063ffffffff16565b8051909150156112b55780806020019051810190612544919061404e565b6112b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610741565b6125d981610e0a565b61261b576040517fa9902c7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610741565b600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84169281019290925233602483015273ffffffffffffffffffffffffffffffffffffffff16906383826b2b90604401602060405180830381865afa15801561269a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126be919061404e565b61186d576040517f728fe07b000000000000000000000000000000000000000000000000000000008152336004820152602401610741565b67ffffffffffffffff82166000908152600760205260409020611c5590600201827f0000000000000000000000000000000000000000000000000000000000000000612d0a565b6000611ea68373ffffffffffffffffffffffffffffffffffffffff8416612c08565b6000611ea68373ffffffffffffffffffffffffffffffffffffffff8416612bb9565b7f00000000000000000000000000000000000000000000000000000000000000001561186d576127b260028261308d565b61186d576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610741565b61280981610e0a565b61284b576040517fa9902c7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610741565b600480546040517fa8d87a3b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84169281019290925273ffffffffffffffffffffffffffffffffffffffff169063a8d87a3b90602401602060405180830381865afa1580156128c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128e8919061416c565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461186d576040517f728fe07b000000000000000000000000000000000000000000000000000000008152336004820152602401610741565b67ffffffffffffffff82166000908152600760205260409020611c5590827f0000000000000000000000000000000000000000000000000000000000000000612d0a565b6060816000018054806020026020016040519081016040528092919081815260200182805480156106e657602002820191906000526020600020905b8154815260200190600101908083116129ce5750505050509050919050565b6000612a0c856129fd8486614189565b612a0790876141a0565b6130bc565b90505b949350505050565b8154600090612a4090700100000000000000000000000000000000900463ffffffff164261409a565b90508015612ae25760018301548354612a88916fffffffffffffffffffffffffffffffff808216928116918591700100000000000000000000000000000000909104166129ed565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b60208201518354612b08916fffffffffffffffffffffffffffffffff90811691166130bc565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1990612213908490614130565b6000818152600183016020526040812054612c005750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561063c565b50600061063c565b60008181526001830160205260408120548015612cf1576000612c2c60018361409a565b8554909150600090612c409060019061409a565b9050808214612ca5576000866000018281548110612c6057612c60613d7a565b9060005260206000200154905080876000018481548110612c8357612c83613d7a565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080612cb657612cb66141b3565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061063c565b600091505061063c565b6060612a0f84846000856130d2565b825474010000000000000000000000000000000000000000900460ff161580612d31575081155b15612d3b57505050565b825460018401546fffffffffffffffffffffffffffffffff80831692911690600090612d8190700100000000000000000000000000000000900463ffffffff164261409a565b90508015612e415781831115612dc3576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001860154612dfd9083908590849070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff166129ed565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff167001000000000000000000000000000000004263ffffffff160217875592505b84821015612ef85773ffffffffffffffffffffffffffffffffffffffff8416612ea0576040517ff94ebcd10000000000000000000000000000000000000000000000000000000081526004810183905260248101869052604401610741565b6040517f1a76572a000000000000000000000000000000000000000000000000000000008152600481018390526024810186905273ffffffffffffffffffffffffffffffffffffffff85166044820152606401610741565b8483101561300b5760018681015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16906000908290612f3c908261409a565b612f46878a61409a565b612f5091906141a0565b612f5a91906141e2565b905073ffffffffffffffffffffffffffffffffffffffff8616612fb3576040517f15279c080000000000000000000000000000000000000000000000000000000081526004810182905260248101869052604401610741565b6040517fd0c8d23a000000000000000000000000000000000000000000000000000000008152600481018290526024810186905273ffffffffffffffffffffffffffffffffffffffff87166044820152606401610741565b613015858461409a565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515611ea6565b60008183106130cb5781611ea6565b5090919050565b606082471015613164576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610741565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161318d919061421d565b60006040518083038185875af1925050503d80600081146131ca576040519150601f19603f3d011682016040523d82523d6000602084013e6131cf565b606091505b50915091506131e0878383876131eb565b979650505050505050565b6060831561328157825160000361327a5773ffffffffffffffffffffffffffffffffffffffff85163b61327a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610741565b5081612a0f565b612a0f83838151156132965781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107419190613400565b5080546132d6906138fd565b6000825580601f106132e6575050565b601f01602090049060005260206000209081019061186d91905b808211156133145760008155600101613300565b5090565b60006020828403121561332a57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611ea657600080fd5b803567ffffffffffffffff8116811461337257600080fd5b919050565b60006020828403121561338957600080fd5b611ea68261335a565b60005b838110156133ad578181015183820152602001613395565b50506000910152565b600081518084526133ce816020860160208601613392565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611ea660208301846133b6565b60006020828403121561342557600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461186d57600080fd5b80356133728161342c565b60006020828403121561346b57600080fd5b8135611ea68161342c565b60006020828403121561348857600080fd5b813567ffffffffffffffff81111561349f57600080fd5b82016101008185031215611ea657600080fd5b60008083601f8401126134c457600080fd5b50813567ffffffffffffffff8111156134dc57600080fd5b6020830191508360208260051b85010111156134f757600080fd5b9250929050565b6000806000806040858703121561351457600080fd5b843567ffffffffffffffff8082111561352c57600080fd5b613538888389016134b2565b9096509450602087013591508082111561355157600080fd5b5061355e878288016134b2565b95989497509550505050565b6000806040838503121561357d57600080fd5b82356135888161342c565b946020939093013593505050565b6000806000604084860312156135ab57600080fd5b6135b48461335a565b9250602084013567ffffffffffffffff808211156135d157600080fd5b818601915086601f8301126135e557600080fd5b8135818111156135f457600080fd5b87602082850101111561360657600080fd5b6020830194508093505050509250925092565b60006020828403121561362b57600080fd5b813567ffffffffffffffff81111561364257600080fd5b820160a08185031215611ea657600080fd5b60208152600082516040602084015261367060608401826133b6565b905060208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160408501526136ab82826133b6565b95945050505050565b6020808252825182820181905260009190848201906040850190845b8181101561370257835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016136d0565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561370257835167ffffffffffffffff168352928401929184019160010161372a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff811182821017156137a3576137a3613750565b60405290565b60405160c0810167ffffffffffffffff811182821017156137a3576137a3613750565b801515811461186d57600080fd5b8035613372816137cc565b80356fffffffffffffffffffffffffffffffff8116811461337257600080fd5b60006060828403121561381757600080fd5b6040516060810181811067ffffffffffffffff8211171561383a5761383a613750565b604052905080823561384b816137cc565b8152613859602084016137e5565b602082015261386a604084016137e5565b60408201525092915050565b600080600060e0848603121561388b57600080fd5b6138948461335a565b92506138a38560208601613805565b91506138b28560808601613805565b90509250925092565b600080602083850312156138ce57600080fd5b823567ffffffffffffffff8111156138e557600080fd5b6138f1858286016134b2565b90969095509350505050565b600181811c9082168061391157607f821691505b60208210810361394a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60006020828403121561396257600080fd5b5051919050565b600082601f83011261397a57600080fd5b813567ffffffffffffffff8082111561399557613995613750565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156139db576139db613750565b816040528381528660208588010111156139f457600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006101008236031215613a2757600080fd5b613a2f61377f565b823567ffffffffffffffff80821115613a4757600080fd5b613a5336838701613969565b8352613a616020860161335a565b6020840152613a726040860161344e565b604084015260608501356060840152613a8d6080860161344e565b608084015260a0850135915080821115613aa657600080fd5b613ab236838701613969565b60a084015260c0850135915080821115613acb57600080fd5b613ad736838701613969565b60c084015260e0850135915080821115613af057600080fd5b50613afd36828601613969565b60e08301525092915050565b601f8211156112b5576000816000526020600020601f850160051c81016020861015613b325750805b601f850160051c820191505b81811015613b5157828155600101613b3e565b505050505050565b67ffffffffffffffff831115613b7157613b71613750565b613b8583613b7f83546138fd565b83613b09565b6000601f841160018114613bd75760008515613ba15750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355613c6d565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b82811015613c265786850135825560209485019460019092019101613c06565b5086821015613c61577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b604081526000613c8760408301866133b6565b82810360208401528381528385602083013760006020858301015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f860116820101915050949350505050565b600060a08236031215613cea57600080fd5b60405160a0810167ffffffffffffffff8282108183111715613d0e57613d0e613750565b816040528435915080821115613d2357600080fd5b50613d3036828601613969565b825250613d3f6020840161335a565b60208201526040830135613d528161342c565b6040820152606083810135908201526080830135613d6f8161342c565b608082015292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec1833603018112613ddd57600080fd5b9190910192915050565b60006101408236031215613dfa57600080fd5b613e026137a9565b613e0b8361335a565b8152613e19602084016137da565b6020820152604083013567ffffffffffffffff80821115613e3957600080fd5b613e4536838701613969565b60408401526060850135915080821115613e5e57600080fd5b50613e6b36828601613969565b606083015250613e7e3660808501613805565b6080820152613e903660e08501613805565b60a082015292915050565b815167ffffffffffffffff811115613eb557613eb5613750565b613ec981613ec384546138fd565b84613b09565b602080601f831160018114613f1c5760008415613ee65750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555613b51565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015613f6957888601518255948401946001909101908401613f4a565b5085821015613fa557878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b600061010067ffffffffffffffff87168352806020840152613fd9818401876133b6565b8551151560408581019190915260208701516fffffffffffffffffffffffffffffffff90811660608701529087015116608085015291506140179050565b8251151560a083015260208301516fffffffffffffffffffffffffffffffff90811660c084015260408401511660e08301526136ab565b60006020828403121561406057600080fd5b8151611ea6816137cc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561063c5761063c61406b565b67ffffffffffffffff8416815260e081016140f960208301858051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b82511515608083015260208301516fffffffffffffffffffffffffffffffff90811660a084015260408401511660c0830152612a0f565b6060810161063c82848051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b60006020828403121561417e57600080fd5b8151611ea68161342c565b808202811582820484141761063c5761063c61406b565b8082018082111561063c5761063c61406b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600082614218577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60008251613ddd81846020870161339256fea164736f6c6343000818000a",
}

var LockReleaseTokenPoolABI = LockReleaseTokenPoolMetaData.ABI

var LockReleaseTokenPoolBin = LockReleaseTokenPoolMetaData.Bin

func DeployLockReleaseTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, allowlist []common.Address, rmnProxy common.Address, acceptLiquidity bool, router common.Address) (common.Address, *CustomTransaction, *LockReleaseTokenPool, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	switch chainId.Uint64() {

	case 324, 280, 300:
		return DeployZkSyncLockReleaseTokenPool(auth, backend, token, allowlist, rmnProxy, acceptLiquidity, router)
	}

	parsed, err := LockReleaseTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LockReleaseTokenPoolBin), backend, token, allowlist, rmnProxy, acceptLiquidity, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, &CustomTransaction{Transaction: tx, CustomHash: tx.Hash()}, &LockReleaseTokenPool{address: address, abi: *parsed, LockReleaseTokenPoolCaller: LockReleaseTokenPoolCaller{contract: contract}, LockReleaseTokenPoolTransactor: LockReleaseTokenPoolTransactor{contract: contract}, LockReleaseTokenPoolFilterer: LockReleaseTokenPoolFilterer{contract: contract}}, nil
}

type LockReleaseTokenPool struct {
	address common.Address
	abi     abi.ABI
	LockReleaseTokenPoolCaller
	LockReleaseTokenPoolTransactor
	LockReleaseTokenPoolFilterer
}

type LockReleaseTokenPoolCaller struct {
	contract *bind.BoundContract
}

type LockReleaseTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type LockReleaseTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type LockReleaseTokenPoolSession struct {
	Contract     *LockReleaseTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type LockReleaseTokenPoolCallerSession struct {
	Contract *LockReleaseTokenPoolCaller
	CallOpts bind.CallOpts
}

type LockReleaseTokenPoolTransactorSession struct {
	Contract     *LockReleaseTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type LockReleaseTokenPoolRaw struct {
	Contract *LockReleaseTokenPool
}

type LockReleaseTokenPoolCallerRaw struct {
	Contract *LockReleaseTokenPoolCaller
}

type LockReleaseTokenPoolTransactorRaw struct {
	Contract *LockReleaseTokenPoolTransactor
}

func NewLockReleaseTokenPool(address common.Address, backend bind.ContractBackend) (*LockReleaseTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(LockReleaseTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindLockReleaseTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPool{address: address, abi: abi, LockReleaseTokenPoolCaller: LockReleaseTokenPoolCaller{contract: contract}, LockReleaseTokenPoolTransactor: LockReleaseTokenPoolTransactor{contract: contract}, LockReleaseTokenPoolFilterer: LockReleaseTokenPoolFilterer{contract: contract}}, nil
}

func NewLockReleaseTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*LockReleaseTokenPoolCaller, error) {
	contract, err := bindLockReleaseTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolCaller{contract: contract}, nil
}

func NewLockReleaseTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LockReleaseTokenPoolTransactor, error) {
	contract, err := bindLockReleaseTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolTransactor{contract: contract}, nil
}

func NewLockReleaseTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LockReleaseTokenPoolFilterer, error) {
	contract, err := bindLockReleaseTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolFilterer{contract: contract}, nil
}

func bindLockReleaseTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LockReleaseTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LockReleaseTokenPool.Contract.LockReleaseTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockReleaseTokenPoolTransactor.contract.Transfer(opts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockReleaseTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LockReleaseTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.contract.Transfer(opts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) CanAcceptLiquidity(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "canAcceptLiquidity")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) CanAcceptLiquidity() (bool, error) {
	return _LockReleaseTokenPool.Contract.CanAcceptLiquidity(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) CanAcceptLiquidity() (bool, error) {
	return _LockReleaseTokenPool.Contract.CanAcceptLiquidity(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetAllowList(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetAllowList(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _LockReleaseTokenPool.Contract.GetAllowListEnabled(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _LockReleaseTokenPool.Contract.GetAllowListEnabled(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LockReleaseTokenPool.Contract.GetCurrentInboundRateLimiterState(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LockReleaseTokenPool.Contract.GetCurrentInboundRateLimiterState(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LockReleaseTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LockReleaseTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRateLimitAdmin(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRateLimitAdmin(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRebalancer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRebalancer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRebalancer() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRebalancer(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRebalancer() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRebalancer(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRemotePool(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRemotePool", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRemotePool(remoteChainSelector uint64) ([]byte, error) {
	return _LockReleaseTokenPool.Contract.GetRemotePool(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRemotePool(remoteChainSelector uint64) ([]byte, error) {
	return _LockReleaseTokenPool.Contract.GetRemotePool(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _LockReleaseTokenPool.Contract.GetRemoteToken(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _LockReleaseTokenPool.Contract.GetRemoteToken(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRmnProxy(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRmnProxy(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRouter() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRouter(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRouter() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRouter(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _LockReleaseTokenPool.Contract.GetSupportedChains(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _LockReleaseTokenPool.Contract.GetSupportedChains(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetToken() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetToken(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetToken(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsSupportedChain(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsSupportedChain(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsSupportedToken(&_LockReleaseTokenPool.CallOpts, token)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsSupportedToken(&_LockReleaseTokenPool.CallOpts, token)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) Owner() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.Owner(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) Owner() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.Owner(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LockReleaseTokenPool.Contract.SupportsInterface(&_LockReleaseTokenPool.CallOpts, interfaceId)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LockReleaseTokenPool.Contract.SupportsInterface(&_LockReleaseTokenPool.CallOpts, interfaceId)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) TypeAndVersion() (string, error) {
	return _LockReleaseTokenPool.Contract.TypeAndVersion(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _LockReleaseTokenPool.Contract.TypeAndVersion(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.AcceptOwnership(&_LockReleaseTokenPool.TransactOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.AcceptOwnership(&_LockReleaseTokenPool.TransactOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyAllowListUpdates(&_LockReleaseTokenPool.TransactOpts, removes, adds)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyAllowListUpdates(&_LockReleaseTokenPool.TransactOpts, removes, adds)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "applyChainUpdates", chains)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ApplyChainUpdates(chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyChainUpdates(&_LockReleaseTokenPool.TransactOpts, chains)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ApplyChainUpdates(chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyChainUpdates(&_LockReleaseTokenPool.TransactOpts, chains)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockOrBurn(&_LockReleaseTokenPool.TransactOpts, lockOrBurnIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockOrBurn(&_LockReleaseTokenPool.TransactOpts, lockOrBurnIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "provideLiquidity", amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ProvideLiquidity(&_LockReleaseTokenPool.TransactOpts, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ProvideLiquidity(&_LockReleaseTokenPool.TransactOpts, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ReleaseOrMint(&_LockReleaseTokenPool.TransactOpts, releaseOrMintIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ReleaseOrMint(&_LockReleaseTokenPool.TransactOpts, releaseOrMintIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetChainRateLimiterConfig(&_LockReleaseTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetChainRateLimiterConfig(&_LockReleaseTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetRateLimitAdmin(&_LockReleaseTokenPool.TransactOpts, rateLimitAdmin)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetRateLimitAdmin(&_LockReleaseTokenPool.TransactOpts, rateLimitAdmin)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetRebalancer(opts *bind.TransactOpts, rebalancer common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setRebalancer", rebalancer)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetRebalancer(rebalancer common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetRebalancer(&_LockReleaseTokenPool.TransactOpts, rebalancer)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetRebalancer(rebalancer common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetRebalancer(&_LockReleaseTokenPool.TransactOpts, rebalancer)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetRemotePool(&_LockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetRemotePool(&_LockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setRouter", newRouter)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetRouter(&_LockReleaseTokenPool.TransactOpts, newRouter)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetRouter(&_LockReleaseTokenPool.TransactOpts, newRouter)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) TransferLiquidity(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "transferLiquidity", from, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) TransferLiquidity(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.TransferLiquidity(&_LockReleaseTokenPool.TransactOpts, from, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) TransferLiquidity(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.TransferLiquidity(&_LockReleaseTokenPool.TransactOpts, from, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.TransferOwnership(&_LockReleaseTokenPool.TransactOpts, to)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.TransferOwnership(&_LockReleaseTokenPool.TransactOpts, to)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "withdrawLiquidity", amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.WithdrawLiquidity(&_LockReleaseTokenPool.TransactOpts, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.WithdrawLiquidity(&_LockReleaseTokenPool.TransactOpts, amount)
}

type LockReleaseTokenPoolAllowListAddIterator struct {
	Event *LockReleaseTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAllowListAdd)
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
		it.Event = new(LockReleaseTokenPoolAllowListAdd)
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

func (it *LockReleaseTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*LockReleaseTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAllowListAddIterator{contract: _LockReleaseTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAllowListAdd)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*LockReleaseTokenPoolAllowListAdd, error) {
	event := new(LockReleaseTokenPoolAllowListAdd)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAllowListRemoveIterator struct {
	Event *LockReleaseTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAllowListRemove)
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
		it.Event = new(LockReleaseTokenPoolAllowListRemove)
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

func (it *LockReleaseTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*LockReleaseTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAllowListRemoveIterator{contract: _LockReleaseTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAllowListRemove)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*LockReleaseTokenPoolAllowListRemove, error) {
	event := new(LockReleaseTokenPoolAllowListRemove)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolBurnedIterator struct {
	Event *LockReleaseTokenPoolBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolBurned)
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
		it.Event = new(LockReleaseTokenPoolBurned)
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

func (it *LockReleaseTokenPoolBurnedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*LockReleaseTokenPoolBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolBurnedIterator{contract: _LockReleaseTokenPool.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolBurned)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseBurned(log types.Log) (*LockReleaseTokenPoolBurned, error) {
	event := new(LockReleaseTokenPoolBurned)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolChainAddedIterator struct {
	Event *LockReleaseTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolChainAdded)
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
		it.Event = new(LockReleaseTokenPoolChainAdded)
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

func (it *LockReleaseTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainAddedIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolChainAddedIterator{contract: _LockReleaseTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolChainAdded)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseChainAdded(log types.Log) (*LockReleaseTokenPoolChainAdded, error) {
	event := new(LockReleaseTokenPoolChainAdded)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolChainConfiguredIterator struct {
	Event *LockReleaseTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolChainConfigured)
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
		it.Event = new(LockReleaseTokenPoolChainConfigured)
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

func (it *LockReleaseTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolChainConfiguredIterator{contract: _LockReleaseTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolChainConfigured)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseChainConfigured(log types.Log) (*LockReleaseTokenPoolChainConfigured, error) {
	event := new(LockReleaseTokenPoolChainConfigured)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolChainRemovedIterator struct {
	Event *LockReleaseTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolChainRemoved)
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
		it.Event = new(LockReleaseTokenPoolChainRemoved)
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

func (it *LockReleaseTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolChainRemovedIterator{contract: _LockReleaseTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolChainRemoved)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseChainRemoved(log types.Log) (*LockReleaseTokenPoolChainRemoved, error) {
	event := new(LockReleaseTokenPoolChainRemoved)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolConfigChangedIterator struct {
	Event *LockReleaseTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolConfigChanged)
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
		it.Event = new(LockReleaseTokenPoolConfigChanged)
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

func (it *LockReleaseTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*LockReleaseTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolConfigChangedIterator{contract: _LockReleaseTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolConfigChanged)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseConfigChanged(log types.Log) (*LockReleaseTokenPoolConfigChanged, error) {
	event := new(LockReleaseTokenPoolConfigChanged)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolLiquidityAddedIterator struct {
	Event *LockReleaseTokenPoolLiquidityAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolLiquidityAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolLiquidityAdded)
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
		it.Event = new(LockReleaseTokenPoolLiquidityAdded)
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

func (it *LockReleaseTokenPoolLiquidityAddedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolLiquidityAdded struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolLiquidityAddedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "LiquidityAdded", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolLiquidityAddedIterator{contract: _LockReleaseTokenPool.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityAdded, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "LiquidityAdded", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolLiquidityAdded)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseLiquidityAdded(log types.Log) (*LockReleaseTokenPoolLiquidityAdded, error) {
	event := new(LockReleaseTokenPoolLiquidityAdded)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolLiquidityRemovedIterator struct {
	Event *LockReleaseTokenPoolLiquidityRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolLiquidityRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolLiquidityRemoved)
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
		it.Event = new(LockReleaseTokenPoolLiquidityRemoved)
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

func (it *LockReleaseTokenPoolLiquidityRemovedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolLiquidityRemoved struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolLiquidityRemovedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "LiquidityRemoved", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolLiquidityRemovedIterator{contract: _LockReleaseTokenPool.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityRemoved, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "LiquidityRemoved", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolLiquidityRemoved)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseLiquidityRemoved(log types.Log) (*LockReleaseTokenPoolLiquidityRemoved, error) {
	event := new(LockReleaseTokenPoolLiquidityRemoved)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolLiquidityTransferredIterator struct {
	Event *LockReleaseTokenPoolLiquidityTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolLiquidityTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolLiquidityTransferred)
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
		it.Event = new(LockReleaseTokenPoolLiquidityTransferred)
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

func (it *LockReleaseTokenPoolLiquidityTransferredIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolLiquidityTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolLiquidityTransferred struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterLiquidityTransferred(opts *bind.FilterOpts, from []common.Address) (*LockReleaseTokenPoolLiquidityTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "LiquidityTransferred", fromRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolLiquidityTransferredIterator{contract: _LockReleaseTokenPool.contract, event: "LiquidityTransferred", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchLiquidityTransferred(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityTransferred, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "LiquidityTransferred", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolLiquidityTransferred)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityTransferred", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseLiquidityTransferred(log types.Log) (*LockReleaseTokenPoolLiquidityTransferred, error) {
	event := new(LockReleaseTokenPoolLiquidityTransferred)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolLockedIterator struct {
	Event *LockReleaseTokenPoolLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolLocked)
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
		it.Event = new(LockReleaseTokenPoolLocked)
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

func (it *LockReleaseTokenPoolLockedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*LockReleaseTokenPoolLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolLockedIterator{contract: _LockReleaseTokenPool.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolLocked)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseLocked(log types.Log) (*LockReleaseTokenPoolLocked, error) {
	event := new(LockReleaseTokenPoolLocked)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolMintedIterator struct {
	Event *LockReleaseTokenPoolMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolMinted)
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
		it.Event = new(LockReleaseTokenPoolMinted)
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

func (it *LockReleaseTokenPoolMintedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LockReleaseTokenPoolMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolMintedIterator{contract: _LockReleaseTokenPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolMinted)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseMinted(log types.Log) (*LockReleaseTokenPoolMinted, error) {
	event := new(LockReleaseTokenPoolMinted)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolOwnershipTransferRequestedIterator struct {
	Event *LockReleaseTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolOwnershipTransferRequested)
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
		it.Event = new(LockReleaseTokenPoolOwnershipTransferRequested)
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

func (it *LockReleaseTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockReleaseTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolOwnershipTransferRequestedIterator{contract: _LockReleaseTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolOwnershipTransferRequested)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*LockReleaseTokenPoolOwnershipTransferRequested, error) {
	event := new(LockReleaseTokenPoolOwnershipTransferRequested)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolOwnershipTransferredIterator struct {
	Event *LockReleaseTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolOwnershipTransferred)
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
		it.Event = new(LockReleaseTokenPoolOwnershipTransferred)
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

func (it *LockReleaseTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockReleaseTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolOwnershipTransferredIterator{contract: _LockReleaseTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolOwnershipTransferred)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*LockReleaseTokenPoolOwnershipTransferred, error) {
	event := new(LockReleaseTokenPoolOwnershipTransferred)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolReleasedIterator struct {
	Event *LockReleaseTokenPoolReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolReleased)
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
		it.Event = new(LockReleaseTokenPoolReleased)
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

func (it *LockReleaseTokenPoolReleasedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LockReleaseTokenPoolReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolReleasedIterator{contract: _LockReleaseTokenPool.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolReleased)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseReleased(log types.Log) (*LockReleaseTokenPoolReleased, error) {
	event := new(LockReleaseTokenPoolReleased)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolRemotePoolSetIterator struct {
	Event *LockReleaseTokenPoolRemotePoolSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolRemotePoolSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolRemotePoolSet)
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
		it.Event = new(LockReleaseTokenPoolRemotePoolSet)
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

func (it *LockReleaseTokenPoolRemotePoolSetIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolRemotePoolSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolRemotePoolSet struct {
	RemoteChainSelector uint64
	PreviousPoolAddress []byte
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterRemotePoolSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolRemotePoolSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "RemotePoolSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolRemotePoolSetIterator{contract: _LockReleaseTokenPool.contract, event: "RemotePoolSet", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchRemotePoolSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRemotePoolSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "RemotePoolSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolRemotePoolSet)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolSet", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseRemotePoolSet(log types.Log) (*LockReleaseTokenPoolRemotePoolSet, error) {
	event := new(LockReleaseTokenPoolRemotePoolSet)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolRouterUpdatedIterator struct {
	Event *LockReleaseTokenPoolRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolRouterUpdated)
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
		it.Event = new(LockReleaseTokenPoolRouterUpdated)
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

func (it *LockReleaseTokenPoolRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*LockReleaseTokenPoolRouterUpdatedIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolRouterUpdatedIterator{contract: _LockReleaseTokenPool.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolRouterUpdated)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseRouterUpdated(log types.Log) (*LockReleaseTokenPoolRouterUpdated, error) {
	event := new(LockReleaseTokenPoolRouterUpdated)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolTokensConsumedIterator struct {
	Event *LockReleaseTokenPoolTokensConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolTokensConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolTokensConsumed)
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
		it.Event = new(LockReleaseTokenPoolTokensConsumed)
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

func (it *LockReleaseTokenPoolTokensConsumedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolTokensConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolTokensConsumed struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterTokensConsumed(opts *bind.FilterOpts) (*LockReleaseTokenPoolTokensConsumedIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolTokensConsumedIterator{contract: _LockReleaseTokenPool.contract, event: "TokensConsumed", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolTokensConsumed) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolTokensConsumed)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseTokensConsumed(log types.Log) (*LockReleaseTokenPoolTokensConsumed, error) {
	event := new(LockReleaseTokenPoolTokensConsumed)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var LockReleaseTokenPoolZkBin string = ("0x000400000000000200180000000000020000006003100270000005990030019d0000059903300197000300000031035500020000000103550000000100200190000000260000c13d0000008004000039000000400040043f000000040030008c0000004a0000413d000000000201043b000000e002200270000005aa0020009c0000004c0000213d000005c20020009c000000ef0000213d000005ce0020009c0000018f0000213d000005d40020009c000001f70000213d000005d70020009c000007310000613d000005d80020009c0000004a0000c13d000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000401100370000000000101043b0000059d0010009c0000004a0000213d165f12170000040f000001c00000013d0000010004000039000000400040043f0000000002000416000000000002004b0000004a0000c13d0000001f023000390000059a022001970000010002200039000000400020043f0000001f0530018f0000059b063001980000010002600039000000380000613d000000000701034f000000007807043c0000000004840436000000000024004b000000340000c13d000000000005004b000000450000613d000000000161034f0000000304500210000000000502043300000000054501cf000000000545022f000000000101043b0000010004400089000000000141022f00000000014101cf000000000151019f0000000000120435000000a00030008c0000004a0000413d000001000100043d0000059c0010009c000001480000a13d00000000010000190000166100010430000005ab0020009c000001610000213d000005b70020009c000001b00000213d000005bd0020009c0000020f0000213d000005c00020009c000007420000613d000005c10020009c0000004a0000c13d000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000402100370000000000502043b0000059d0050009c0000004a0000213d0000000002530049000005e70020009c0000004a0000213d000000a40020008c0000004a0000413d0000006002000039000000800020043f000000a00020043f0000016002000039000000400020043f0000000404500039000000000641034f000000000606043b0000059d0060009c0000004a0000213d00000000065600190000002305600039000000000035004b0000004a0000813d0000000407600039000000000571034f000000000505043b0000059d0050009c0000015b0000213d0000001f095000390000062f099001970000003f099000390000062f09900197000006070090009c0000015b0000213d0000016009900039000000400090043f000001600050043f00000000065600190000002406600039000000000036004b0000004a0000213d0000002003700039000000000631034f0000062f075001980000001f0850018f0000018003700039000000900000613d0000018009000039000000000a06034f00000000ab0a043c0000000009b90436000000000039004b0000008c0000c13d000000000008004b0000009d0000613d000000000676034f0000000307800210000000000803043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f000000000063043500000180035000390000000000030435000000c00020043f000e00200040003d0000000e02100360000000000202043b0000059d0020009c0000004a0000213d000000e00020043f0000000e020000290000002002200039000000000321034f000000000303043b0000059c0030009c0000004a0000213d000001000030043f0000004003200039000000000331034f0000002002200039000000000121034f000000000101043b000c00000001001d000001200010043f000000000103043b000d00000001001d0000059c0010009c0000004a0000213d0000000d01000029000001400010043f000005dd0100004100000000001004430000000001000412000000040010044300000024000004430000000001000414000005990010009c0000059901008041000000c001100210000005de011001c70000800502000039165f165a0000040f000000010020019000000ebe0000613d000000400200043d000b00000002001d0000000402200039000000000101043b0000059c011001970000000d0010006b00000d860000c13d000000e00100043d00000609030000410000000b04000029000000000034043500000080011002100000060a011001970000000000120435000005dd01000041000000000010044300000000010004120000000400100443000000200100003900000024001004430000000001000414000005990010009c0000059901008041000000c001100210000005de011001c70000800502000039165f165a0000040f000000010020019000000ebe0000613d000000000201043b00000000010004140000059c02200197000000040020008c00000e4f0000c13d0000000103000031000000200030008c0000002004000039000000000403401900000e790000013d000005c30020009c000001c70000213d000005c90020009c0000024f0000213d000005cc0020009c000007470000613d000005cd0020009c0000004a0000c13d000000440030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000402100370000000000202043b000e00000002001d0000059c0020009c0000004a0000213d0000002401100370000000000301043b000000000100041a0000059c011001970000000002000411000000000012004b00000b090000c13d000d00000003001d000006150100004100000000001004430000000e0100002900000004001004430000000001000414000005990010009c0000059901008041000000c00110021000000616011001c70000800202000039165f165a0000040f000000010020019000000ebe0000613d000000000101043b000000000001004b0000004a0000613d000000400500043d0000061701000041000000000015043500000004015000390000000d04000029000000000041043500000000010004140000000e02000029000000040020008c000001360000613d000005990050009c000005990200004100000000020540190000004002200210000005990010009c0000059901008041000000c001100210000000000121019f000005e0011001c70000000e02000029000c00000005001d165f16550000040f0000000c050000290000000d040000290000006003100270000105990030019d0003000000010355000000010020019000000c340000613d0000059d0050009c0000015b0000213d000000400050043f0000000000450435000005990050009c000005990500804100000040015002100000000002000414000005990020009c0000059902008041000000c002200210000000000112019f000005a4011001c70000800d02000039000000020300003900000618040000410000000e0500002900000c2f0000013d000001200200043d0000059d0020009c0000004a0000213d0000001f04200039000000000034004b00000000050000190000059e050080410000059e04400197000000000004004b00000000060000190000059e060040410000059e0040009c000000000605c019000000000006004b0000004a0000c13d000001000420003900000000040404330000059d0040009c000006f10000a13d0000061a01000041000000000010043f0000004101000039000000040010043f000005e0010000410000166100010430000005ac0020009c000001e60000213d000005b20020009c000002580000213d000005b50020009c000008080000613d000005b60020009c0000004a0000c13d000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000401100370000000000101043b000e00000001001d0000059d0010009c0000004a0000213d165f12720000040f0000000e01000029000000000010043f0000000701000039000000200010043f00000040020000390000000001000019165f16220000040f000d00000001001d000000400100043d000e00000001001d165f11f30000040f0000000d05000029000000000405041a000005fe004001980000000002000039000000010200c0390000000e01000029000000400310003900000000002304350000008002400270000005990220019700000020031000390000000000230435000005ea0240019700000000002104350000000102500039000002390000013d000005cf0020009c0000069a0000213d000005d20020009c000008480000613d000005d30020009c0000004a0000c13d000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000401100370000000000101043b000e00000001001d0000059c0010009c0000004a0000213d0000000001000412001600000001001d001500000000003d000080050100003900000044030000390000000004000415000000160440008a0000000504400210000005dd02000041165f16370000040f0000059c011001970000000e0010006b00000000010000390000000101006039000000800010043f000005e601000041000016600001042e000005b80020009c000006a30000213d000005bb0020009c000008530000613d000005bc0020009c0000004a0000c13d000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000401100370000000000101043b0000059d0010009c0000004a0000213d165f12870000040f0000002002000039000000400300043d000e00000003001d0000000002230436165f11cb0000040f0000000e02000029000002460000013d000005c40020009c000006bb0000213d000005c70020009c000008590000613d000005c80020009c0000004a0000c13d0000000001000416000000000001004b0000004a0000c13d0000000101000039000000000201041a0000059c032001970000000006000411000000000036004b00000b2d0000c13d000000000300041a000005a004300197000000000464019f000000000040041b000005a002200197000000000021041b00000000010004140000059c05300197000005990010009c0000059901008041000000c001100210000005d9011001c70000800d020000390000000303000039000006130400004100000c2f0000013d000005ad0020009c000006d30000213d000005b00020009c000008960000613d000005b10020009c0000004a0000c13d0000000001000416000000000001004b0000004a0000c13d0000000001000412001000000001001d000f00400000003d000080050100003900000044030000390000000004000415000000100440008a00000acc0000013d000005d50020009c000008a70000613d000005d60020009c0000004a0000c13d0000000001000416000000000001004b0000004a0000c13d000000c001000039000000400010043f0000001a01000039000000800010043f0000062301000041000000a00010043f0000002001000039000000c00010043f0000008001000039000000e002000039165f11cb0000040f000000c00110008a000005990010009c0000059901008041000000600110021000000624011001c7000016600001042e000005be0020009c000008d40000613d000005bf0020009c0000004a0000c13d000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000401100370000000000101043b000e00000001001d0000059d0010009c0000004a0000213d165f12720000040f0000000e01000029000000000010043f0000000701000039000000200010043f00000040020000390000000001000019165f16220000040f000d00000001001d000000400100043d000e00000001001d165f11f30000040f0000000d050000290000000201500039000000000401041a000005fe004001980000000002000039000000010200c0390000000e01000029000000400310003900000000002304350000008002400270000005990220019700000020031000390000000000230435000005ea0240019700000000002104350000000302500039000000000402041a000000800210003900000080034002700000000000320435000005ea0340019700000060021000390000000000320435165f13e10000040f000000400100043d000d00000001001d0000000e02000029165f11fe0000040f0000000d020000290000000001210049000005990010009c00000599010080410000006001100210000005990020009c00000599020080410000004002200210000000000121019f000016600001042e000005ca0020009c000009070000613d000005cb0020009c0000004a0000c13d0000000001000416000000000001004b0000004a0000c13d0000000801000039000008570000013d000005b30020009c000009140000613d000005b40020009c0000004a0000c13d000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000402100370000000000202043b0000059d0020009c0000004a0000213d0000002304200039000000000034004b0000004a0000813d0000000404200039000000000141034f000000000101043b000300000001001d0000059d0010009c0000004a0000213d000800240020003d000000030100002900000005011002100000000801100029000000000031004b0000004a0000213d000000000100041a0000059c011001970000000002000411000000000012004b00000b090000c13d000000030000006b00000c320000613d000b00000000001d000002b90000013d000000000643001900000000000604350000000076050434000000000006004b0000000006000039000000010600c039000000400810003900000000006804350000000006070433000005ea066001970000006007100039000000000067043500000040055000390000000005050433000005ea05500197000000800610003900000000005604350000000065020434000000000005004b0000000005000039000000010500c039000000a00710003900000000005704350000000005060433000005ea05500197000000c006100039000000000056043500000040022000390000000002020433000005ea02200197000000e00510003900000000002504350000001f024000390000062f0220019700000000021200490000000002320019000005990020009c00000599020080410000006002200210000005990010009c00000599010080410000004001100210000000000112019f0000000002000414000005990020009c0000059902008041000000c002200210000000000112019f000005d9011001c70000800d020000390000000103000039000005f604000041165f16550000040f00000001002001900000004a0000613d0000000b020000290000000102200039000b00000002001d000000030020006c00000c320000813d0000000b01000029000000050110021000000008011000290000000202000367000000000112034f000000000101043b0000000004000031000000080340006a0000013f0330008a0000059e053001970000059e06100197000000000756013f000000000056004b00000000050000190000059e05004041000000000031004b00000000030000190000059e030080410000059e0070009c000000000503c019000000000005004b0000004a0000c13d00000008051000290000000001540049000005e70010009c0000004a0000213d000001400010008c0000004a0000413d000000400300043d000e00000003001d000005e80030009c0000015b0000213d0000000e03000029000000c003300039000000400030043f000000000352034f000000000303043b0000059d0030009c0000004a0000213d0000000e06000029000000000d3604360000002006500039000000000762034f000000000707043b000000000007004b0000000008000039000000010800c039000000000087004b0000004a0000c13d00000000007d04350000002006600039000000000762034f000000000707043b0000059d0070009c0000004a0000213d000000000b5700190000001f07b00039000000000047004b00000000080000190000059e080080410000059e097001970000059e07400197000000000a79013f000000000079004b00000000090000190000059e090040410000059e00a0009c000000000908c019000000000009004b0000004a0000c13d0000000008b2034f000000000808043b0000059d0080009c0000015b0000213d0000001f098000390000062f099001970000003f099000390000062f0a900197000000400900043d000000000aa9001900000000009a004b000000000c000039000000010c0040390000059d00a0009c0000015b0000213d0000000100c001900000015b0000c13d0000004000a0043f000000000a890436000000200bb00039000000000cb8001900000000004c004b0000004a0000213d000900000001001d00000000010d0019000000000cb2034f0000062f0d800198000000000bda0019000003220000613d000000000e0c034f000000000f0a001900000000e30e043c000000000f3f04360000000000bf004b0000031e0000c13d0000001f0e8001900000032f0000613d0000000003dc034f000000030ce00210000000000d0b0433000000000dcd01cf000000000dcd022f000000000303043b000001000cc000890000000003c3022f0000000003c301cf0000000003d3019f00000000003b043500000000038a001900000000000304350000000e030000290000004003300039000a00000003001d00000000009304350000002006600039000000000362034f000000000803043b0000059d0080009c000000000d01001900000009010000290000004a0000213d00000000095800190000001f03900039000000000043004b00000000050000190000059e050080410000059e03300197000000000873013f000000000073004b00000000030000190000059e030040410000059e0080009c000000000305c019000000000003004b0000004a0000c13d000000000392034f000000000503043b0000059d0050009c0000015b0000213d0000001f035000390000062f033001970000003f033000390000062f03300197000000400700043d0000000008370019000000000078004b000000000a000039000000010a0040390000059d0080009c0000015b0000213d0000000100a001900000015b0000c13d000000400080043f000000000857043600000020099000390000000003950019000000000043004b0000004a0000213d000000000992034f0000062f0a5001980000000004a800190000036b0000613d000000000b09034f000000000c08001900000000b30b043c000000000c3c043600000000004c004b000003670000c13d0000001f0b500190000003780000613d0000000003a9034f0000000309b00210000000000a040433000000000a9a01cf000000000a9a022f000000000303043b0000010009900089000000000393022f00000000039301cf0000000003a3019f0000000000340435000000000358001900000000000304350000000e030000290000006003300039000c00000003001d0000000000730435000000800410008a000005e70040009c0000004a0000213d000000600040008c0000004a0000413d000000400400043d000005e90040009c0000015b0000213d0000006003400039000000400030043f0000002005600039000000000352034f000000000603043b000000000006004b0000000003000039000000010300c039000000000036004b0000004a0000c13d00000000066404360000002005500039000000000352034f000000000703043b000005ea0070009c0000004a0000213d00000000007604350000002005500039000000000352034f000000000603043b000005ea0060009c0000004a0000213d000000400340003900000000006304350000000e030000290000008003300039000d00000003001d0000000000430435000000e00110008a000005e70010009c0000004a0000213d000000600010008c0000004a0000413d000000400100043d000005e90010009c0000015b0000213d0000006003100039000000400030043f0000002005500039000000000352034f000000000403043b000000000004004b0000000003000039000000010300c039000000000034004b0000004a0000c13d00000000044104360000002005500039000000000352034f000000000603043b000005ea0060009c0000004a0000213d00000000006404350000002003500039000000000232034f000000000202043b000005ea0020009c0000004a0000213d000000400510003900000000002504350000000e02000029000000a002200039000900000002001d000000000012043500000000020d04330000000d03000029000000000303043300000040063000390000000006060433000005ea076001970000000068030434000000000008004b000003d80000613d000000000007004b00000d8c0000613d0000000006060433000005ea06600197000000000067004b00000d8c0000813d000000000002004b000003dd0000c13d00000d350000013d000000000007004b00000d7e0000c13d0000000006060433000005ea0060019800000d7e0000c13d0000000003050433000005ea033001970000000005010433000000000005004b000003eb0000613d000000000003004b00000d930000613d0000000004040433000005ea04400197000000000043004b00000d930000813d000000000002004b000003f20000c13d00000d350000013d000000000003004b00000d820000c13d0000000003040433000005ea0030019800000d820000c13d000000000002004b000005cd0000613d0000000e0100002900000000010104330000059d01100197000700000001001d000000000010043f0000000601000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b000000000101041a000000000001004b00000d710000c13d0000000501000039000000000101041a0000059d0010009c0000015b0000213d00000001021000390000000503000039000000000023041b000005ef0110009a0000000702000029000000000021041b000000000103041a000600000001001d000000000020043f0000000601000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b0000000602000029000000000021041b0000000a0100002900000000010104330000000001010433000000000001004b00000c1a0000613d0000000c0100002900000000010104330000000001010433000000000001004b00000c1a0000613d0000000d01000029000000000201043300000040012000390000000001010433000600000001001d000700000002001d00000020012000390000000001010433000500000001001d000005f00100004100000000001004430000000001000414000005990010009c0000059901008041000000c001100210000005f1011001c70000800b02000039165f165a0000040f000000010020019000000ebe0000613d000000000201043b000000400100043d000005e20010009c0000015b0000213d0000000503000029000005ea043001970000000603000029000005ea05300197000005990320019700000007020000290000000002020433000000a006100039000000400060043f00000080061000390000000000560435000000000002004b0000000002000039000000010200c0390000004005100039000000000025043500000020021000390000000000320435000000600210003900000000004204350000000000410435000000400200043d000005e20020009c0000015b0000213d0000000904000029000000000404043300000020054000390000000005050433000000000604043300000040044000390000000004040433000000a007200039000000400070043f000005ea0440019700000080072000390000000000470435000000000006004b0000000004000039000000010400c0390000004006200039000000000046043500000020042000390000000000340435000005ea03500197000000600420003900000000003404350000000000320435000000400600043d000005f20060009c0000015b0000213d0000000a0300002900000000030304330000000c0400002900000000040404330000008005600039000000400050043f0000006005600039000500000005001d00000000004504350000004004600039000a00000004001d0000000000340435000600000006001d0000000001160436000700000001001d00000000002104350000000e0100002900000000010104330000059d01100197000000000010043f0000000701000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000060200002900000000020204330000000043020434000005ea03300197000000000601043b000000000106041a000005f301100197000000000131019f00000000030404330000008003300210000005f403300197000000000131019f00000040032000390000000003030433000000000003004b000005f5030000410000000003006019000000000131019f000000000016041b000000800120003900000060022000390000000002020433000005ea0220019700000000010104330000008001100210000000000121019f0000000102600039000000000012041b0000000201600039000000000201041a000005f302200197000000070300002900000000030304330000000054030434000005ea04400197000000000242019f00000000040504330000008004400210000005f404400197000000000242019f00000040043000390000000004040433000000000004004b000005f5040000410000000004006019000000000242019f000000000021041b000000800130003900000060023000390000000002020433000005ea0220019700000000010104330000008001100210000000000121019f0000000302600039000000000012041b0000000a01000029000000000301043300000000750304340000059d0050009c0000015b0000213d0000000404600039000000000104041a000000010010019000000001081002700000007f0880618f0000001f0080008c00000000020000390000000102002039000000000121013f000000010010019000000c5b0000c13d000000200080008c000600000006001d000a00000004001d000700000005001d000400000003001d000005020000413d000100000008001d000200000007001d000000000040043f0000000001000414000005990010009c0000059901008041000000c001100210000005a4011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d00000007050000290000001f025000390000000502200270000000200050008c0000000002004019000000000301043b00000001010000290000001f01100039000000050110027000000000011300190000000002230019000000000012004b00000006060000290000000a040000290000000207000029000005020000813d000000000002041b0000000102200039000000000012004b000004fe0000413d0000001f0050008c0000052e0000a13d000000000040043f0000000001000414000005990010009c0000059901008041000000c001100210000005a4011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d00000007070000290000062f02700198000000000101043b0000059d0000613d000000010320008a00000005033002700000000003310019000000010430003900000020030000390000000606000029000000040800002900000000058300190000000005050433000000000051041b00000020033000390000000101100039000000000041004b000005190000c13d000000000072004b0000052a0000813d0000000302700210000000f80220018f000006300220027f000006300220016700000000038300190000000003030433000000000223016f000000000021041b000000010170021000000001011001bf0000000a04000029000005390000013d000000000005004b000005320000613d0000000001070433000005330000013d00000000010000190000000302500210000006300220027f0000063002200167000000000121016f0000000102500210000000000121019f000000000014041b0000000501000029000000000301043300000000750304340000059d0050009c0000015b0000213d0000000504600039000000000104041a000000010010019000000001061002700000007f0660618f0000001f0060008c00000000020000390000000102002039000000000121013f000000010010019000000c5b0000c13d000000200060008c000a00000004001d000700000005001d000500000003001d0000056e0000413d000400000006001d000600000007001d000000000040043f0000000001000414000005990010009c0000059901008041000000c001100210000005a4011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d00000007050000290000001f025000390000000502200270000000200050008c0000000002004019000000000301043b00000004010000290000001f01100039000000050110027000000000011300190000000002230019000000000012004b0000000a0400002900000006070000290000056e0000813d000000000002041b0000000102200039000000000012004b0000056a0000413d0000001f0050008c000005990000a13d000000000040043f0000000001000414000005990010009c0000059901008041000000c001100210000005a4011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d00000007060000290000062f02600198000000000101043b000005c80000613d000000010320008a0000000503300270000000000331001900000001043000390000002003000039000000050700002900000000057300190000000005050433000000000051041b00000020033000390000000101100039000000000041004b000005840000c13d000000000062004b000005950000813d0000000302600210000000f80220018f000006300220027f000006300220016700000000037300190000000003030433000000000223016f000000000021041b000000010160021000000001011001bf0000000a04000029000005aa0000013d000000000005004b000005a30000613d0000000001070433000005a40000013d000000200300003900000006060000290000000408000029000000000072004b000005220000413d0000052a0000013d00000000010000190000000302500210000006300220027f0000063002200167000000000121016f0000000102500210000000000121019f000000000014041b000000090100002900000000020104330000000d0100002900000000050104330000000c0100002900000000030104330000000e010000290000000004010433000000400100043d0000002006100039000001000700003900000000007604350000059d0440019700000000004104350000010007100039000000006403043400000000004704350000012003100039000000000004004b0000027d0000613d000000000700001900000000083700190000000009760019000000000909043300000000009804350000002007700039000000000047004b000005c00000413d0000027d0000013d00000020030000390000000507000029000000000062004b0000058d0000413d000005950000013d0000000e0100002900000000010104330000059d01100197000d00000001001d000000000010043f0000000601000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b000000000301041a000000000003004b00000ec50000613d0000000501000039000000000201041a000000000002004b00000ebf0000613d000000010130008a000000000023004b000006030000613d000000000012004b00000e490000a13d000005f80130009a000005f80220009a000000000202041a000000000021041b000000000020043f0000000601000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039000c00000003001d165f165a0000040f00000001002001900000004a0000613d000000000101043b0000000c02000029000000000021041b0000000501000039000000000301041a000000000003004b00000eeb0000613d000000010130008a000005f80230009a000000000002041b0000000502000039000000000012041b0000000d01000029000000000010043f0000000601000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b000000000001041b0000000e0100002900000000010104330000059d01100197000000000010043f0000000701000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000301043b000000000003041b0000000101300039000000000001041b0000000201300039000000000001041b0000000301300039000000000001041b0000000404300039000000000104041a000000010010019000000001051002700000007f0550618f0000001f0050008c00000000020000390000000102002039000000000121013f000000010010019000000c5b0000c13d000000000005004b0000065a0000613d0000001f0050008c000006590000a13d000a00000005001d000d00000003001d000c00000004001d000000000040043f0000000001000414000005990010009c0000059901008041000000c001100210000005a4011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b0000000a020000290000001f02200039000000050220027000000000022100190000000103100039000000000023004b000006550000813d000000000003041b0000000103300039000000000023004b000006510000413d0000000c02000029000000000002041b00000000040100190000000d03000029000000000004041b0000000503300039000000000103041a000000010010019000000001041002700000007f0440618f0000001f0040008c00000000020000390000000102002039000000000121013f000000010010019000000c5b0000c13d000000000004004b000006850000613d0000001f0040008c000006840000a13d000c00000004001d000d00000003001d000000000030043f0000000001000414000005990010009c0000059901008041000000c001100210000005a4011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b0000000c020000290000001f02200039000000050220027000000000022100190000000103100039000000000023004b000006810000813d000000000003041b0000000103300039000000000023004b0000067d0000413d0000000d02000029000000000002041b0000000003010019000000000003041b0000000e0100002900000000010104330000059d01100197000000400200043d0000000000120435000005990020009c000005990200804100000040012002100000000002000414000005990020009c0000059902008041000000c002200210000000000112019f000005a4011001c70000800d020000390000000103000039000005f904000041165f16550000040f0000000100200190000002b40000c13d0000004a0000013d000005d00020009c0000096f0000613d000005d10020009c0000004a0000c13d0000000001000416000000000001004b0000004a0000c13d0000000901000039000008570000013d000005b90020009c00000ac20000613d000005ba0020009c0000004a0000c13d000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000401100370000000000101043b0000059c0010009c0000004a0000213d000000000200041a0000059c022001970000000003000411000000000023004b00000b090000c13d000000000001004b00000bb80000c13d000005f701000041000000800010043f00000605010000410000166100010430000005c50020009c00000ad50000613d000005c60020009c0000004a0000c13d000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000401100370000000000101043b0000059d0010009c0000004a0000213d165f14fb0000040f000000000001004b0000000001000039000000010100c039000000400200043d0000000000120435000005990020009c0000059902008041000000400120021000000611011001c7000016600001042e000005ae0020009c00000ae70000613d000005af0020009c0000004a0000c13d000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000401100370000000000601043b0000059c0060009c0000004a0000213d000000000100041a0000059c011001970000000005000411000000000015004b00000b090000c13d000000000056004b00000bc90000c13d000005a801000041000000800010043f0000002001000039000000840010043f0000001701000039000000a40010043f000005db01000041000000c40010043f000005dc01000041000016610001043000000005054002100000003f065000390000059f06600197000000400700043d0000000006670019000c00000007001d000000000076004b000000000700003900000001070040390000059d0060009c0000015b0000213d00000001007001900000015b0000c13d0000010007300039000000400060043f0000000c030000290000000003430436000b00000003001d00000120022000390000000003250019000000000073004b0000004a0000213d000000000004004b000007100000613d0000000b0400002900000000250204340000059c0050009c0000004a0000213d0000000004540436000000000032004b0000070a0000413d000001400200043d0000059c0020009c0000004a0000213d000001600400043d000000000004004b0000000003000039000000010300c039000900000004001d000000000034004b0000004a0000c13d000001800300043d0000059c0030009c0000004a0000213d0000000004000411000000000004004b00000bf50000c13d000000400100043d0000004402100039000005a7030000410000000000320435000000240210003900000018030000390000000000320435000005a8020000410000000000210435000000040210003900000020030000390000000000320435000005990010009c00000599010080410000004001100210000005a9011001c70000166100010430000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000401100370000000000201043b00000629002001980000004a0000c13d00000001010000390000062a0020009c00000b370000213d0000062d0020009c00000ad20000613d0000062e0020009c00000ad20000613d00000b3b0000013d0000000001000416000000000001004b0000004a0000c13d000000000100041a000008a30000013d000000440030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000402100370000000000202043b0000059d0020009c0000004a0000213d0000002305200039000000000035004b0000004a0000813d0000000405200039000000000551034f000000000905043b0000059d0090009c0000004a0000213d0000002407200039000000050b90021000000000087b0019000000000038004b0000004a0000213d0000002402100370000000000202043b0000059d0020009c0000004a0000213d0000002305200039000000000035004b0000004a0000813d0000000405200039000000000551034f000000000605043b0000059d0060009c0000004a0000213d0000002402200039000000050a60021000000000052a0019000000000035004b0000004a0000213d000000000300041a0000059c03300197000000000c00041100000000003c004b00000b090000c13d0000003f03b000390000059f03300197000005f20030009c0000015b0000213d0000008003300039000c00000003001d000000400030043f000000800090043f000000000009004b000007880000613d000000000371034f000000000303043b0000059c0030009c0000004a0000213d000000200440003900000000003404350000002007700039000000000087004b0000077d0000413d000000400300043d000c00000003001d0000003f03a000390000059f033001970000000c033000290000000c0030006c000000000400003900000001040040390000059d0030009c0000015b0000213d00000001004001900000015b0000c13d000000400030043f0000000c030000290000000003630436000b00000003001d000000000006004b000007a20000613d0000000c03000029000000000421034f000000000404043b0000059c0040009c0000004a0000213d000000200330003900000000004304350000002002200039000000000052004b000007990000413d000005dd01000041000000000010044300000000010004120000000400100443000000400100003900000024001004430000000001000414000005990010009c0000059901008041000000c001100210000005de011001c70000800502000039165f165a0000040f000000010020019000000ebe0000613d000000000101043b000000000001004b00000c170000613d000000800100043d000000000001004b00000de40000c13d0000000c010000290000000001010433000000000001004b00000c320000613d0000000003000019000007c20000013d00000001033000390000000c010000290000000001010433000000000013004b00000c320000813d00000005013002100000000b0110002900000000010104330000059c04100198000007bd0000613d000000000040043f0000000301000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039000d00000003001d000e00000004001d165f165a0000040f0000000e040000290000000d0300002900000001002001900000004a0000613d000000000101043b000000000101041a000000000001004b000007bd0000c13d0000000203000039000000000103041a0000059d0010009c0000015b0000213d0000000102100039000000000023041b000005a30110009a000000000041041b000000000103041a000a00000001001d000000000040043f0000000301000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f0000000e0300002900000001002001900000004a0000613d000000000101043b0000000a02000029000000000021041b000000400100043d0000000000310435000005990010009c000005990100804100000040011002100000000002000414000005990020009c0000059902008041000000c002200210000000000112019f000005a4011001c70000800d020000390000000103000039000005a504000041165f16550000040f0000000d030000290000000100200190000007bd0000c13d0000004a0000013d0000000002000416000000000002004b0000004a0000c13d0000000504000039000000000204041a000000800020043f000000000040043f000000000002004b00000b130000c13d000000a002000039000000400020043f000000000400001900000005064002100000003f05600039000006020550019700000000052500190000059d0050009c0000015b0000213d000000400050043f00000000044204360000001f0560018f000000000006004b000008260000613d000000000131034f00000000036400190000000006040019000000001701043c0000000006760436000000000036004b000008220000c13d000000000005004b000000800100043d000000000001004b000008380000613d00000000010000190000000003020433000000000013004b00000e490000a13d00000005031002100000000005340019000000a00330003900000000030304330000059d0330019700000000003504350000000101100039000000800300043d000000000031004b0000082b0000413d000000400100043d00000020030000390000000005310436000000000302043300000000003504350000004002100039000000000003004b000008fe0000613d000000000500001900000000460404340000059d0660019700000000026204360000000105500039000000000035004b000008410000413d000008fe0000013d0000000001000416000000000001004b0000004a0000c13d0000000001000412001800000001001d001700000000003d000080050100003900000044030000390000000004000415000000180440008a000008a00000013d0000000001000416000000000001004b0000004a0000c13d0000000401000039000000000101041a000008a30000013d000000440030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000402100370000000000202043b000e00000002001d0000059d0020009c0000004a0000213d0000002402100370000000000202043b0000059d0020009c0000004a0000213d0000002304200039000000000034004b0000004a0000813d0000000404200039000000000141034f000000000101043b000d00000001001d0000059d0010009c0000004a0000213d0000002402200039000c00000002001d0000000d01200029000000000031004b0000004a0000213d000000000100041a0000059c011001970000000002000411000000000012004b00000b090000c13d0000000e01000029000000000010043f0000000601000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b000000000101041a000000000001004b00000c410000c13d000000400100043d000005fa02000041000000000021043500000004021000390000000e030000290000000000320435000005990010009c00000599010080410000004001100210000005e0011001c700001661000104300000000001000416000000000001004b0000004a0000c13d0000000001000412001200000001001d001100200000003d000080050100003900000044030000390000000004000415000000120440008a0000000504400210000005dd02000041165f16370000040f0000059c01100197000000800010043f000005e601000041000016600001042e000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000401100370000000000201043b0000000901000039000000000101041a0000059c011001970000000003000411000000000031004b00000b280000c13d000e00000002001d000005dd0100004100000000001004430000000001000412000000040010044300000024000004430000000001000414000005990010009c0000059901008041000000c001100210000005de011001c70000800502000039165f165a0000040f000000010020019000000ebe0000613d000000000301043b000000400b00043d000006260100004100000000001b04350000000401b00039000000000200041000000000002104350000000001000414000d00000003001d0000059c02300197000000040020008c00000b700000c13d0000000103000031000000200030008c0000002004000039000000000403401900000b9c0000013d0000000001000416000000000001004b0000004a0000c13d0000000202000039000000000102041a000000800010043f000000000020043f0000002002000039000000000001004b0000000004020019000008ed0000613d000000a005000039000006060300004100000000040000190000000006050019000000000503041a000000000556043600000001033000390000000104400039000000000014004b000008e20000413d000000410160008a0000062f04100197000005f20040009c0000015b0000213d0000008001400039000000400010043f0000000000210435000000a002400039000000800300043d0000000000320435000000c002400039000000000003004b000008fe0000613d000000a004000039000000000500001900000000460404340000059c0660019700000000026204360000000105500039000000000035004b000008f80000413d0000000002120049000005990020009c00000599020080410000006002200210000005990010009c00000599010080410000004001100210000000000112019f000016600001042e000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000401100370000000000101043b000e00000001001d0000059c0010009c0000004a0000213d165f13ca0000040f000000090100003900000ae10000013d000000e40030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000402100370000000000202043b000e00000002001d0000059d0020009c0000004a0000213d000000e002000039000000400020043f0000002402100370000000000202043b000000000002004b0000000003000039000000010300c039000000000032004b0000004a0000c13d000000800020043f0000004402100370000000000202043b000005ea0020009c0000004a0000213d000000a00020043f0000006402100370000000000202043b000005ea0020009c0000004a0000213d000000c00020043f0000014002000039000000400020043f0000008402100370000000000202043b000000000002004b0000000003000039000000010300c039000000000032004b0000004a0000c13d000000e00020043f000000a402100370000000000202043b000005ea0020009c0000004a0000213d000001000020043f000000c401100370000000000101043b000005ea0010009c0000004a0000213d000001200010043f0000000801000039000000000101041a0000059c021001970000000001000411000000000021004b000009500000613d000000000200041a0000059c02200197000000000021004b00000cb10000c13d0000000e01000029000000000010043f0000000601000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b000000000101041a000000000001004b0000088b0000613d000000c00100043d000005ea01100197000000800200043d000000000002004b00000d380000c13d000000000001004b0000096b0000c13d000000a00100043d000005ea0010019800000d3e0000613d000000400200043d000e00000002001d000005ec0100004100000d9c0000013d000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000402100370000000000202043b000e00000002001d0000059d0020009c0000004a0000213d0000000e0230006a000005e70020009c0000004a0000213d000001040020008c0000004a0000413d000000800000043f000001a002000039000000400020043f0000000e040000290000000404400039000000000541034f000000000505043b0000059d0050009c0000004a0000213d0000000e065000290000002305600039000000000035004b0000004a0000813d0000000407600039000000000571034f000000000505043b0000059d0050009c0000015b0000213d0000001f085000390000062f088001970000003f088000390000062f088001970000061e0080009c0000015b0000213d000001a008800039000000400080043f000001a00050043f00000000065600190000002406600039000000000036004b0000004a0000213d0000002006700039000000000761034f0000062f085001980000001f0950018f000001c006800039000009a90000613d000001c00a000039000000000b07034f00000000bc0b043c000000000aca043600000000006a004b000009a50000c13d000000000009004b000009b60000613d000000000787034f0000000308900210000000000906043300000000098901cf000000000989022f000000000707043b0000010008800089000000000787022f00000000078701cf000000000797019f0000000000760435000001c0055000390000000000050435000000a00020043f0000002002400039000000000421034f000000000404043b0000059d0040009c0000004a0000213d000000c00040043f0000002002200039000000000421034f000000000404043b0000059c0040009c0000004a0000213d000000e00040043f0000002004200039000000000441034f0000004002200039000000000521034f000000000404043b000d00000004001d000001000040043f000000000405043b0000059c0040009c0000004a0000213d000001200040043f0000002002200039000000000421034f000000000404043b0000059d0040009c0000004a0000213d0000000e074000290000002304700039000000000034004b0000004a0000813d0000000408700039000000000481034f000000000404043b0000059d0040009c0000015b0000213d0000001f054000390000062f055001970000003f055000390000062f06500197000000400500043d0000000006650019000000000056004b000000000900003900000001090040390000059d0060009c0000015b0000213d00000001009001900000015b0000c13d000000400060043f000000000645043600000000074700190000002407700039000000000037004b0000004a0000213d0000002007800039000000000871034f0000062f094001980000001f0a40018f0000000007960019000009fd0000613d000000000b08034f000000000c06001900000000bd0b043c000000000cdc043600000000007c004b000009f90000c13d00000000000a004b00000a0a0000613d000000000898034f0000000309a00210000000000a070433000000000a9a01cf000000000a9a022f000000000808043b0000010009900089000000000898022f00000000089801cf0000000008a8019f000000000087043500000000044600190000000000040435000001400050043f0000002002200039000000000421034f000000000404043b0000059d0040009c0000004a0000213d0000000e074000290000002304700039000000000034004b0000004a0000813d0000000408700039000000000481034f000000000404043b0000059d0040009c0000015b0000213d0000001f054000390000062f055001970000003f055000390000062f06500197000000400500043d0000000006650019000000000056004b000000000900003900000001090040390000059d0060009c0000015b0000213d00000001009001900000015b0000c13d000000400060043f000000000645043600000000074700190000002407700039000000000037004b0000004a0000213d0000002007800039000000000871034f0000062f094001980000001f0a40018f000000000796001900000a3a0000613d000000000b08034f000000000c06001900000000bd0b043c000000000cdc043600000000007c004b00000a360000c13d00000000000a004b00000a470000613d000000000898034f0000000309a00210000000000a070433000000000a9a01cf000000000a9a022f000000000808043b0000010009900089000000000898022f00000000089801cf0000000008a8019f000000000087043500000000044600190000000000040435000001600050043f0000002002200039000000000221034f000000000202043b0000059d0020009c0000004a0000213d0000000e062000290000002302600039000000000032004b0000004a0000813d0000000407600039000000000271034f000000000202043b0000059d0020009c0000015b0000213d0000001f042000390000062f044001970000003f044000390000062f05400197000000400400043d0000000005540019000000000045004b000000000800003900000001080040390000059d0050009c0000015b0000213d00000001008001900000015b0000c13d000000400050043f000000000524043600000000062600190000002406600039000000000036004b0000004a0000213d0000002003700039000000000331034f0000062f062001980000001f0720018f000000000165001900000a770000613d000000000803034f0000000009050019000000008a08043c0000000009a90436000000000019004b00000a730000c13d000000000007004b00000a840000613d000000000363034f0000000306700210000000000701043300000000076701cf000000000767022f000000000303043b0000010006600089000000000363022f00000000036301cf000000000373019f000000000031043500000000012500190000000000010435000001800040043f000001200100043d000c00000001001d000005dd0100004100000000001004430000000001000412000000040010044300000024000004430000000001000414000005990010009c0000059901008041000000c001100210000005de011001c70000800502000039165f165a0000040f000000010020019000000ebe0000613d0000000c020000290000059c03200197000000400200043d000c00000002001d0000000402200039000000000101043b000900000001001d0000059c01100197000a00000003001d000000000013004b000010290000c13d000000c00100043d00000609030000410000000c04000029000000000034043500000080011002100000060a011001970000000000120435000005dd01000041000000000010044300000000010004120000000400100443000000200100003900000024001004430000000001000414000005990010009c0000059901008041000000c001100210000005de011001c70000800502000039165f165a0000040f000000010020019000000ebe0000613d000000000201043b00000000010004140000059c02200197000000040020008c0000102c0000c13d0000000103000031000000200030008c00000020040000390000000004034019000010560000013d0000000001000416000000000001004b0000004a0000c13d0000000001000412001400000001001d001300600000003d000080050100003900000044030000390000000004000415000000140440008a0000000504400210000005dd02000041165f16370000040f000000000001004b0000000001000039000000010100c039000000800010043f000005e601000041000016600001042e000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000401100370000000000101043b000e00000001001d0000059c0010009c0000004a0000213d165f13ca0000040f0000000801000039000000000201041a000005a0022001970000000e022001af000000000021041b0000000001000019000016600001042e000000240030008c0000004a0000413d0000000002000416000000000002004b0000004a0000c13d0000000401100370000000000101043b000e00000001001d000005dd01000041000000000010044300000000010004120000000400100443000000600100003900000024001004430000000001000414000005990010009c0000059901008041000000c001100210000005de011001c70000800502000039165f165a0000040f000000010020019000000ebe0000613d000000400300043d000000000101043b000000000001004b00000b3e0000c13d000005e4010000410000000000130435000005990030009c00000599030080410000004001300210000005e5011001c70000166100010430000005a801000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f0000061901000041000000c40010043f000005dc010000410000166100010430000000a007000039000005ff0400004100000000060000190000000005070019000000000704041a000000000775043600000001044000390000000106600039000000000026004b00000b160000413d000006000250009a000006010020009c0000015b0000413d000000410250008a0000062f022001970000008002200039000000800400043d000000400020043f0000059d0040009c0000015b0000213d000008140000013d000005df01000041000000800010043f000000840030043f00000625010000410000166100010430000005a801000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f0000061201000041000000c40010043f000005dc0100004100001661000104300000062b0020009c00000ad20000613d0000062c0020009c00000ad20000613d000000800000043f000005e601000041000016600001042e0000000901000039000000000101041a0000059c011001970000000004000411000000000041004b00000baf0000c13d0000002001300039000005e102000041000000000021043500000064013000390000000e0200002900000000002104350000004401300039000000000200041000000000002104350000002401300039000000000041043500000064010000390000000000130435000005e20030009c0000015b0000213d000000a001300039000000400010043f000005dd0100004100000000001004430000000001000412000000040010044300000024000004430000000001000414000005990010009c0000059901008041000000c001100210000005de011001c70000800502000039000d00000003001d165f165a0000040f000000010020019000000ebe0000613d000000000101043b0000000d02000029165f142f0000040f0000000001000414000005990010009c0000059901008041000000c001100210000005d9011001c70000800d020000390000000303000039000005e30400004100000c2d0000013d0000059900b0009c000005990300004100000000030b40190000004003300210000005990010009c0000059901008041000000c001100210000000000131019f000005e0011001c7000c0000000b001d165f165a0000040f0000000c0b00002900000060031002700000059903300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b001900000b8b0000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b00000b870000c13d000000000006004b00000b980000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000000bd70000613d0000001f01400039000000600210018f0000000001b20019000000000021004b000000000200003900000001020040390000059d0010009c0000015b0000213d00000001002001900000015b0000c13d000000400010043f000000200030008c0000004a0000413d00000000020b04330000000e03000029000000000032004b00000c220000813d000006280200004100000c1c0000013d000005df01000041000000000013043500000004013000390000000000410435000005990030009c00000599030080410000004001300210000005e0011001c700001661000104300000000402000039000000000302041a000005a004300197000000000414019f000000000042041b0000059c02300197000000800020043f000000a00010043f0000000001000414000005990010009c0000059901008041000000c00110021000000603011001c70000800d020000390000000103000039000006040400004100000c2f0000013d0000000101000039000000000201041a000005a002200197000000000262019f000000000021041b0000000001000414000005990010009c0000059901008041000000c001100210000005d9011001c70000800d020000390000000303000039000005da0400004100000c2f0000013d0000001f0530018f0000059b06300198000000400200043d000000000462001900000be20000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00000bde0000c13d000000000005004b00000bef0000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000005990020009c00000599020080410000004002200210000000000112019f0000166100010430000000000500041a000005a005500197000000000445019f000000000040041b000000000002004b00000c1a0000613d0000059c0010019800000c1a0000613d000000000003004b00000c1a0000613d000000800010043f000000a00020043f0000000405000039000000000105041a000005a001100197000000000131019f000000000015041b0000000c010000290000000001010433000000000001004b0000000001000039000000010100c039000000c00010043f000000000100001900000d5e0000613d000000400100043d000005a10010009c0000015b0000213d0000002002100039000000400020043f0000000000010435000000c00100043d000000000001004b00000c610000c13d000000400100043d0000061d0200004100000c1c0000013d000000400100043d000005f7020000410000000000210435000005990010009c00000599010080410000004001100210000005e5011001c700001661000104300000000d010000290000000002000411165f12f00000040f0000000001000414000005990010009c0000059901008041000000c001100210000005d9011001c70000800d020000390000000303000039000006270400004100000000050004110000000e06000029165f16550000040f00000001002001900000004a0000613d0000000001000019000016600001042e00000599033001970000001f0530018f0000059b06300198000000400200043d000000000462001900000be20000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00000c3c0000c13d00000be20000013d0000000e01000029000000000010043f0000000701000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b0000000401100039000000000201041a000000010320019000000001042002700000007f0440618f000b00000004001d0000001f0040008c00000000040000390000000104002039000000000442013f000000010040019000000cb60000613d0000061a01000041000000000010043f0000002201000039000000040010043f000005e00100004100001661000104300000000c020000290000000002020433000000000002004b00000d5e0000613d000000000200001900000c6d0000013d0000000e0200002900000001022000390000000c010000290000000001010433000000000012004b00000d5c0000813d000e00000002001d00000005012002100000000b0110002900000000010104330000059c0310019800000c670000613d000000000030043f0000000301000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039000d00000003001d165f165a0000040f0000000d0400002900000001002001900000004a0000613d000000000101043b000000000101041a000000000001004b00000c670000c13d0000000203000039000000000103041a0000059d0010009c0000015b0000213d0000000102100039000000000023041b000005a30110009a000000000041041b000000000103041a000a00000001001d000000000040043f0000000301000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f0000000d0300002900000001002001900000004a0000613d000000000101043b0000000a02000029000000000021041b000000400100043d0000000000310435000005990010009c000005990100804100000040011002100000000002000414000005990020009c0000059902008041000000c002200210000000000112019f000005a4011001c70000800d020000390000000103000039000005a504000041165f16550000040f000000010020019000000c670000c13d0000004a0000013d000005df02000041000001400020043f000001440010043f000005fb010000410000166100010430000000400400043d000900000004001d0000000b050000290000000004540436000a00000004001d000000000003004b00000cd60000613d000000000010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a4011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d0000000b06000029000000000006004b00000000020000190000000a0500002900000cdc0000613d000000000101043b00000000020000190000000003250019000000000401041a000000000043043500000001011000390000002002200039000000000062004b00000cce0000413d00000cdc0000013d00000631012001970000000a0200002900000000001204350000000b0000006b000000200200003900000000020060390000003f012000390000062f021001970000000901200029000000000021004b000000000200003900000001020040390000059d0010009c0000015b0000213d00000001002001900000015b0000c13d000000400010043f0000000e01000029000000000010043f0000000701000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b0000000401100039000b00000001001d000000000101041a000000010010019000000001021002700000007f0220618f000800000002001d0000001f0020008c00000000020000390000000102002039000000000121013f000000010010019000000c5b0000c13d0000000801000029000000200010008c00000d210000413d0000000b01000029000000000010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a4011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d0000000d030000290000001f023000390000000502200270000000200030008c0000000002004019000000000301043b00000008010000290000001f01100039000000050110027000000000011300190000000002230019000000000012004b00000d210000813d000000000002041b0000000102200039000000000012004b00000d1d0000413d0000000d010000290000001f0010008c00000eca0000a13d0000000b01000029000000000010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a4011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000200300008a0000000d02300180000000000101043b00000ef10000c13d000000000300001900000efc0000013d000000400100043d000005ed0200004100000c1c0000013d000000000001004b00000d990000613d000000a00200043d000005ea02200197000000000021004b00000d990000813d0000000e01000029000000000010043f0000000701000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b0000008002000039165f150f0000040f000001200100043d000005ea01100197000000e00200043d000000000002004b00000da20000c13d000000000001004b00000d580000c13d000001000100043d000005ea0010019800000da80000613d000000400200043d000e00000002001d000005ec0100004100000edb0000013d000000c00100043d00000004050000390000000904000029000000e00040043f000000800200043d00000140000004430000016000200443000000a00200043d00000020030000390000018000300443000001a0002004430000004002000039000001c000200443000001e00010044300000060010000390000020000100443000002200040044300000100003004430000012000500443000005a601000041000016600001042e0000000e010000290000000001010433000000400200043d000005ee0300004100000000003204350000059d0110019700000004032000390000000000130435000005990020009c00000599020080410000004001200210000005e0011001c70000166100010430000000400200043d000e00000002001d000005ec0100004100000d8f0000013d000000400300043d000e00000003001d000005ec0200004100000d960000013d00000608010000410000000b0300002900000000001304350000000d01000029000000000012043500000bb30000013d000000400200043d000e00000002001d000005eb0100004100000000001204350000000402200039000000000103001900000d9f0000013d000000400300043d000e00000003001d000005eb020000410000000000230435000000040230003900000d9f0000013d000000400200043d000e00000002001d000005eb01000041000000000012043500000004022000390000008001000039165f12e20000040f0000000e010000290000072c0000013d000000000001004b00000ed80000613d000001000200043d000005ea02200197000000000021004b00000ed80000813d0000000e01000029000000000010043f0000000701000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b0000000201100039000000e002000039165f150f0000040f000000400100043d0000000e020000290000000002210436000000800300043d000000000003004b0000000003000039000000010300c0390000000000320435000000a00200043d000005ea0220019700000040031000390000000000230435000000c00200043d000005ea0220019700000060031000390000000000230435000000e00200043d000000000002004b0000000002000039000000010200c03900000080031000390000000000230435000001000200043d000005ea02200197000000a0031000390000000000230435000001200200043d000005ea02200197000000c0031000390000000000230435000005990010009c000005990100804100000040011002100000000002000414000005990020009c0000059902008041000000c002200210000000000112019f000005fc011001c70000800d020000390000000103000039000005fd0400004100000bc80000013d000000000200001900000deb0000013d0000000d020000290000000102200039000000800100043d000000000012004b000007b70000813d000d00000002001d0000000501200210000000a00110003900000000010104330000059c01100197000e00000001001d000000000010043f0000000301000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b000000000301041a000000000003004b00000de60000613d0000000201000039000000000201041a000000000002004b00000ebf0000613d000000010130008a000000000023004b00000e230000613d000000000012004b00000e490000a13d0000061b0130009a0000061b0220009a000000000202041a000000000021041b000000000020043f0000000301000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039000a00000003001d165f165a0000040f0000000a0300002900000001002001900000004a0000613d000000000101043b000000000031041b0000000201000039000000000301041a000000000003004b00000eeb0000613d000000010130008a0000061b0230009a000000000002041b0000000202000039000000000012041b0000000e01000029000000000010043f0000000301000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b000000000001041b000000400100043d0000000e020000290000000000210435000005990010009c000005990100804100000040011002100000000002000414000005990020009c0000059902008041000000c002200210000000000112019f000005a4011001c70000800d0200003900000001030000390000061c04000041165f16550000040f000000010020019000000de60000c13d0000004a0000013d0000061a01000041000000000010043f0000003201000039000000040010043f000005e00100004100001661000104300000000b03000029000005990030009c00000599030080410000004003300210000005990010009c0000059901008041000000c001100210000000000131019f000005e0011001c7165f165a0000040f00000060031002700000059903300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000000b0570002900000e680000613d000000000801034f0000000b09000029000000008a08043c0000000009a90436000000000059004b00000e640000c13d000000000006004b00000e750000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000000edf0000613d0000001f01400039000000600210018f0000000b01200029000000000021004b000000000200003900000001020040390000059d0010009c0000015b0000213d00000001002001900000015b0000c13d000000400010043f000000200030008c0000004a0000413d0000000b020000290000000002020433000000000002004b0000000003000039000000010300c039000000000032004b0000004a0000c13d000000000002004b000010870000c13d000001000100043d000b00000001001d000005dd01000041000000000010044300000000010004120000000400100443000000400100003900000024001004430000000001000414000005990010009c0000059901008041000000c001100210000005de011001c70000800502000039165f165a0000040f000000010020019000000ebe0000613d000000000101043b000000000001004b00000f5a0000c13d000000e00100043d0000059d01100197000a00000001001d000000000010043f0000000601000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000400200043d000b00000002001d0000000402200039000000000101043b000000000101041a000000000001004b00000f730000c13d00000610010000410000000b0300002900000000001304350000000a0100002900000d8a0000013d000000000001042f0000061a01000041000000000010043f0000001101000039000000040010043f000005e00100004100001661000104300000000e010000290000000001010433000000400200043d000005fa0300004100000d750000013d0000000d0000006b000000000100001900000ed00000613d0000000c010000290000000201100367000000000101043b0000000d040000290000000302400210000006300220027f0000063002200167000000000121016f0000000102400210000000000121019f00000f0b0000013d000000400200043d000e00000002001d000005eb0100004100000000001204350000000402200039000000e00100003900000d9f0000013d0000001f0530018f0000059b06300198000000400200043d000000000462001900000be20000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00000ee60000c13d00000be20000013d0000061a01000041000000000010043f0000003101000039000000040010043f000005e0010000410000166100010430000000020400036700000000030000190000000c060000290000000005630019000000000554034f000000000505043b000000000051041b00000001011000390000002003300039000000000023004b00000ef40000413d0000000d0020006c00000f080000813d0000000d020000290000000302200210000000f80220018f000006300220027f00000630022001670000000c033000290000000203300367000000000303043b000000000223016f000000000021041b0000000d01000029000000010110021000000001011001bf0000000b02000029000000000012041b0000004002000039000000400100043d000000000221043600000009030000290000000003030433000000400410003900000000003404350000006004100039000000000003004b0000000a0800002900000f200000613d000000000500001900000000064500190000000007580019000000000707043300000000007604350000002005500039000000000035004b00000f190000413d000000000534001900000000000504350000001f033000390000062f033001970000000003340019000000000413004900000000004204350000000d0500002900000000025304360000062f045001980000001f0550018f00000000034200190000000c06000029000000020660036700000f350000613d000000000706034f0000000008020019000000007907043c0000000008980436000000000038004b00000f310000c13d000000000005004b00000f420000613d000000000446034f0000000305500210000000000603043300000000065601cf000000000656022f000000000404043b0000010005500089000000000454022f00000000045401cf000000000464019f00000000004304350000000d040000290000001f034000390000062f033001970000000004420019000000000004043500000000031300490000000002230019000005990020009c00000599020080410000006002200210000005990010009c00000599010080410000004001100210000000000112019f0000000002000414000005990020009c0000059902008041000000c002200210000000000112019f000005d9011001c70000800d0200003900000002030000390000061404000041000001460000013d0000000b010000290000059c01100197000b00000001001d000000000010043f0000000301000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b000000000101041a000000000001004b00000ea30000c13d000000400100043d0000060c02000041000000000021043500000004021000390000000b03000029000008900000013d0000000401000039000000000301041a0000060d010000410000000b0400002900000000001404350000000a01000029000000000012043500000000010004140000059c02300197000000040020008c00000f820000c13d0000000104000031000000200040008c000000200400803900000fac0000013d0000000b03000029000005990030009c00000599030080410000004003300210000005990010009c0000059901008041000000c001100210000000000131019f000005e0011001c7165f165a0000040f00000060031002700000059903300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000000b0570002900000f9b0000613d000000000801034f0000000b09000029000000008a08043c0000000009a90436000000000059004b00000f970000c13d000000000006004b00000fa80000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f00030000000103550000000100200190000010190000613d0000001f01400039000000600210018f0000000b01200029000000000021004b000000000200003900000001020040390000059d0010009c0000015b0000213d00000001002001900000015b0000c13d000000400010043f000000200040008c0000004a0000413d0000000b020000290000000002020433000005f50020009c0000004a0000813d0000000003000411000000000023004b000010250000c13d000000e00100043d0000059d01100197000001200200043d000b00000002001d000000000010043f0000000701000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b0000000b020000290000000d03000029165f157c0000040f000000400100043d0000000c020000290000000000210435000005990010009c000005990100804100000040011002100000000002000414000005990020009c0000059902008041000000c002200210000000000121019f000005a4011001c70000800d0200003900000002030000390000060f040000410000000005000411165f16550000040f00000001002001900000004a0000613d0000000e010000290000000201100367000000000101043b0000059d0010009c0000004a0000213d165f12870000040f000e00000001001d000000400100043d000d00000001001d165f11e80000040f0000000d010000290000000e020000290000000001210436000e00000001001d000000400100043d000c00000001001d165f11dd0000040f0000000c0200002900000000000204350000000e010000290000000000210435000000400400043d000c00000004001d000000200100003900000000021404360000000d010000290000000001010433000000400300003900000000003204350000006002400039165f11cb0000040f00000000020100190000000c040000290000000001410049000000200310008a0000000e01000029000000000101043300000040044000390000000000340435165f11cb0000040f0000000c020000290000000001210049000005990020009c00000599020080410000004002200210000005990010009c00000599010080410000006001100210000000000121019f000016600001042e0000001f0530018f0000059b06300198000000400200043d000000000462001900000be20000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000010200000c13d00000be20000013d0000060e0200004100000000002104350000000402100039000008900000013d00000608010000410000000c0300002900000ebb0000013d0000000c03000029000005990030009c00000599030080410000004003300210000005990010009c0000059901008041000000c001100210000000000131019f000005e0011001c7165f165a0000040f00000060031002700000059903300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000000c05700029000010450000613d000000000801034f0000000c09000029000000008a08043c0000000009a90436000000000059004b000010410000c13d000000000006004b000010520000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f00030000000103550000000100200190000010890000613d0000001f01400039000000600210018f0000000c01200029000000000021004b000000000200003900000001020040390000059d0010009c0000015b0000213d00000001002001900000015b0000c13d000000400010043f000000200030008c0000004a0000413d0000000c020000290000000002020433000000000002004b0000000003000039000000010300c039000000000032004b0000004a0000c13d000000000002004b000010870000c13d000000c00100043d0000059d01100197000b00000001001d000000000010043f0000000601000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000400200043d000c00000002001d0000000402200039000000000101043b000000000101041a000000000001004b000010950000c13d00000610010000410000000c0300002900000000001304350000000b0100002900000d8a0000013d0000060b0200004100000c1c0000013d0000001f0530018f0000059b06300198000000400200043d000000000462001900000be20000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000010900000c13d00000be20000013d0000000401000039000000000301041a0000061f010000410000000c0400002900000000001404350000000b01000029000000000012043500000024014000390000000002000411000000000021043500000000010004140000059c02300197000000040020008c000010a70000c13d0000000104000031000000200040008c0000002004008039000010d10000013d0000000c03000029000005990030009c00000599030080410000004003300210000005990010009c0000059901008041000000c001100210000000000131019f00000620011001c7165f165a0000040f00000060031002700000059903300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000000c05700029000010c00000613d000000000801034f0000000c09000029000000008a08043c0000000009a90436000000000059004b000010bc0000c13d000000000006004b000010cd0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f00030000000103550000000100200190000010ec0000613d0000001f01400039000000600210018f0000000c01200029000000000021004b000000000200003900000001020040390000059d0010009c0000015b0000213d00000001002001900000015b0000c13d000000400010043f000000200040008c0000004a0000413d0000000c020000290000000002020433000000000002004b0000000003000039000000010300c039000000000032004b0000004a0000c13d000000000002004b000010f80000c13d0000060e02000041000000000021043500000004021000390000000003000411000008900000013d0000001f0530018f0000059b06300198000000400200043d000000000462001900000be20000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000010f30000c13d00000be20000013d000000c00100043d0000059d01100197000000000010043f0000000701000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b0000000401100039000000000201041a000000010320019000000001042002700000007f0440618f000c00000004001d0000001f0040008c00000000040000390000000104002039000000000043004b00000c5b0000c13d000000400400043d000800000004001d0000000c050000290000000004540436000b00000004001d000000000003004b000011300000613d000000000010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a4011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d0000000c0000006b0000000002000019000011360000613d000000000101043b00000000020000190000000b03200029000000000401041a0000000000430435000000010110003900000020022000390000000c0020006c000011280000413d000011360000013d00000631012001970000000b0200002900000000001204350000000c0000006b000000200200003900000000020060390000003f012000390000062f011001970000000802100029000000000012004b000000000100003900000001010040390000059d0020009c0000015b0000213d00000001001001900000015b0000c13d0000000003020019000000400020043f00000008010000290000000001010433000c00000001001d000000000001004b000011b70000613d000001400100043d0000000021010434000005990010009c00000599010080410000006001100210000005990020009c00000599020080410000004002200210000000000112019f0000000002000414000005990020009c0000059902008041000000c002200210000000000121019f000005d9011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d0000000c02000029000005990020009c000005990200804100000060022002100000000b03000029000005990030009c00000599030080410000004003300210000000000232019f000000000101043b000c00000001001d0000000001000414000005990010009c0000059901008041000000c001100210000000000121019f000005d9011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d000000000101043b0000000c0010006b000011b60000c13d000000c00100043d0000059d01100197000001000200043d000c00000002001d000000000010043f0000000701000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000004a0000613d0000000e02000029000e00440020003d000000000101043b00000002011000390000000c020000290000000a03000029165f157c0000040f0000000e010000290000000201100367000000000201043b0000059c0020009c0000004a0000213d00000009010000290000000d03000029165f12f00000040f0000000e010000290000000201100367000000000601043b0000059c0060009c0000004a0000213d000000400100043d0000000d020000290000000000210435000005990010009c000005990100804100000040011002100000000002000414000005990020009c0000059902008041000000c002200210000000000112019f000005a4011001c70000800d02000039000000030300003900000621040000410000000005000411165f16550000040f00000001002001900000004a0000613d000000400100043d000e00000001001d165f11dd0000040f0000000d020000290000000e010000290000000000210435000000400100043d0000000000210435000005990010009c0000059901008041000000400110021000000611011001c7000016600001042e000000400300043d000001400100043d00000622020000410000000004030019000e00000003001d00000000002304350000000402300039000000200300003900000000003204350000002402400039165f11cb0000040f0000000e020000290000000001210049000005990010009c0000059901008041000005990020009c000005990200804100000060011002100000004002200210000000000121019f000016610001043000000000430104340000000001320436000000000003004b000011d70000613d000000000200001900000000052100190000000006240019000000000606043300000000006504350000002002200039000000000032004b000011d00000413d000000000231001900000000000204350000001f023000390000062f022001970000000001210019000000000001042d000006320010009c000011e20000813d0000002001100039000000400010043f000000000001042d0000061a01000041000000000010043f0000004101000039000000040010043f000005e0010000410000166100010430000006330010009c000011ed0000813d0000004001100039000000400010043f000000000001042d0000061a01000041000000000010043f0000004101000039000000040010043f000005e0010000410000166100010430000006340010009c000011f80000813d000000a001100039000000400010043f000000000001042d0000061a01000041000000000010043f0000004101000039000000040010043f000005e00100004100001661000104300000000043020434000005ea03300197000000000331043600000000040404330000059904400197000000000043043500000040032000390000000003030433000000000003004b0000000003000039000000010300c0390000004004100039000000000034043500000060032000390000000003030433000005ea033001970000006004100039000000000034043500000080022000390000000002020433000005ea0220019700000080031000390000000000230435000000a001100039000000000001042d00030000000000020000059d01100197000000000010043f0000000701000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f0000000100200190000012640000613d000000000101043b0000000405100039000000000205041a000000010320019000000001062002700000007f0660618f0000001f0060008c00000000040000390000000104002039000000000043004b000012660000c13d000000400100043d0000000004610436000000000003004b000012500000613d000100000004001d000200000006001d000300000001001d000000000050043f0000000001000414000005990010009c0000059901008041000000c001100210000005a4011001c70000801002000039165f165a0000040f0000000100200190000012640000613d0000000206000029000000000006004b000012560000613d000000000201043b0000000005000019000000030100002900000001070000290000000003570019000000000402041a000000000043043500000001022000390000002005500039000000000065004b000012480000413d000012580000013d00000631022001970000000000240435000000000006004b00000020050000390000000005006039000012580000013d000000000500001900000003010000290000003f035000390000062f023001970000000003120019000000000023004b000000000200003900000001020040390000059d0030009c0000126c0000213d00000001002001900000126c0000c13d000000400030043f000000000001042d000000000100001900001661000104300000061a01000041000000000010043f0000002201000039000000040010043f000005e00100004100001661000104300000061a01000041000000000010043f0000004101000039000000040010043f000005e0010000410000166100010430000000400100043d000006340010009c000012810000813d000000a002100039000000400020043f000000800210003900000000000204350000006002100039000000000002043500000040021000390000000000020435000000200210003900000000000204350000000000010435000000000001042d0000061a01000041000000000010043f0000004101000039000000040010043f000005e001000041000016610001043000030000000000020000059d01100197000000000010043f0000000701000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f0000000100200190000012d40000613d000000000101043b0000000505100039000000000205041a000000010320019000000001062002700000007f0660618f0000001f0060008c00000000040000390000000104002039000000000043004b000012d60000c13d000000400100043d0000000004610436000000000003004b000012c00000613d000100000004001d000200000006001d000300000001001d000000000050043f0000000001000414000005990010009c0000059901008041000000c001100210000005a4011001c70000801002000039165f165a0000040f0000000100200190000012d40000613d0000000206000029000000000006004b000012c60000613d000000000201043b0000000005000019000000030100002900000001070000290000000003570019000000000402041a000000000043043500000001022000390000002005500039000000000065004b000012b80000413d000012c80000013d00000631022001970000000000240435000000000006004b00000020050000390000000005006039000012c80000013d000000000500001900000003010000290000003f035000390000062f023001970000000003120019000000000023004b000000000200003900000001020040390000059d0030009c000012dc0000213d0000000100200190000012dc0000c13d000000400030043f000000000001042d000000000100001900001661000104300000061a01000041000000000010043f0000002201000039000000040010043f000005e00100004100001661000104300000061a01000041000000000010043f0000004101000039000000040010043f000005e00100004100001661000104300000000043010434000000000003004b0000000003000039000000010300c03900000000033204360000000004040433000005ea044001970000000000430435000000400220003900000040011000390000000001010433000005ea011001970000000000120435000000000001042d0002000000000002000000400400043d000000440540003900000000003504350000002003400039000006350500004100000000005304350000059c022001970000002405400039000000000025043500000044020000390000000000240435000006360040009c0000137e0000813d0000008009400039000000400090043f000005e80040009c0000137e0000213d0000059c0a100197000000c001400039000000400010043f00000020010000390000000000190435000000a00140003900000637020000410000000000210435000000000204043300000000010004140000000400a0008c0000133c0000c13d00000001020000390000000101000031000000000001004b000013540000613d0000059d0010009c0000137e0000213d0000001f041000390000062f044001970000003f044000390000062f04400197000000400c00043d00000000044c00190000000000c4004b000000000500003900000001050040390000059d0040009c0000137e0000213d00000001005001900000137e0000c13d000000400040043f000000000b1c04360000062f031001980000001f0410018f00000000013b001900000003050003670000132e0000613d000000000605034f00000000070b0019000000006806043c0000000007870436000000000017004b0000132a0000c13d000000000004004b000013560000613d000000000335034f0000000304400210000000000501043300000000054501cf000000000545022f000000000303043b0000010004400089000000000343022f00000000034301cf000000000353019f0000000000310435000013560000013d000005990030009c00000599030080410000004003300210000005990020009c00000599020080410000006002200210000000000232019f000005990010009c0000059901008041000000c001100210000000000112019f00000000020a0019000200000009001d00010000000a001d165f16550000040f000000010a0000290000000209000029000000010220018f00030000000103550000006001100270000105990010019d0000059901100197000000000001004b000013120000c13d000000600c000039000000800b00003900000000030c0433000000000002004b000013860000613d000000000003004b000013710000c13d00020000000c001d00010000000b001d000006150100004100000000001004430000000400a004430000000001000414000005990010009c0000059901008041000000c00110021000000616011001c70000800202000039165f165a0000040f0000000100200190000013b80000613d000000000101043b000000000001004b0000000201000029000013b90000613d0000000003010433000000000003004b000000010b0000290000137d0000613d000005e70030009c000013840000213d0000001f0030008c000013840000a13d00000000010b0433000000000001004b0000000002000039000000010200c039000000000021004b000013840000c13d000000000001004b0000139c0000613d000000000001042d0000061a01000041000000000010043f0000004101000039000000040010043f000005e001000041000016610001043000000000010000190000166100010430000000000003004b000013b00000c13d0000000001090019000000400400043d000200000004001d000005a80200004100000000002404350000000403400039000000200200003900000000002304350000002402400039165f11cb0000040f00000002020000290000000001210049000005990010009c0000059901008041000005990020009c000005990200804100000060011002100000004002200210000000000121019f0000166100010430000000400100043d00000064021000390000063803000041000000000032043500000044021000390000063903000041000000000032043500000024021000390000002a030000390000000000320435000005a8020000410000000000210435000000040210003900000020030000390000000000320435000005990010009c000005990100804100000040011002100000063a011001c700001661000104300000059900b0009c000005990b0080410000004002b00210000005990030009c00000599030080410000006001300210000000000121019f0000166100010430000000000001042f000000400100043d00000044021000390000063b03000041000000000032043500000024021000390000001d030000390000000000320435000005a8020000410000000000210435000000040210003900000020030000390000000000320435000005990010009c00000599010080410000004001100210000005a9011001c70000166100010430000000000100041a0000059c011001970000000002000411000000000012004b000013d00000c13d000000000001042d000000400100043d000000440210003900000619030000410000000000320435000000240210003900000016030000390000000000320435000005a8020000410000000000210435000000040210003900000020030000390000000000320435000005990010009c00000599010080410000004001100210000005a9011001c700001661000104300005000000000002000000400300043d000006340030009c000014280000813d000000a002300039000000400020043f00000080023000390000000000020435000000600230003900000000000204350000004002300039000000000002043500000020023000390000000000020435000000000003043500000060021000390000000002020433000100000002001d000500000001001d0000000012010434000300000002001d000200000001001d0000000001010433000400000001001d000005f00100004100000000001004430000000001000414000005990010009c0000059901008041000000c001100210000005f1011001c70000800b02000039165f165a0000040f00000001002001900000142e0000613d00000004020000290000059904200197000000000601043b000000000346004b0000000501000029000014220000413d00000080021000390000000002020433000005ea0520019700000000023500a9000000000046004b000014130000613d00000000033200d9000000000053004b000014220000c13d0000000303000029000005ea03300197000000000032001a000014220000413d00000000023200190000000103000029000005ea03300197000005ea04200197000000000023004b00000000030480190000000000310435000005990260019700000002030000290000000000230435000000000001042d0000061a01000041000000000010043f0000001101000039000000040010043f000005e00100004100001661000104300000061a01000041000000000010043f0000004101000039000000040010043f000005e0010000410000166100010430000000000001042f0002000000000002000000400900043d000006330090009c000014af0000813d0000059c0a1001970000004001900039000000400010043f00000020019000390000063703000041000000000031043500000020010000390000000000190435000000002302043400000000010004140000000400a0008c0000146d0000c13d00000001020000390000000101000031000000000001004b000014850000613d0000059d0010009c000014af0000213d0000001f041000390000062f044001970000003f044000390000062f04400197000000400c00043d00000000044c00190000000000c4004b000000000500003900000001050040390000059d0040009c000014af0000213d0000000100500190000014af0000c13d000000400040043f000000000b1c04360000062f031001980000001f0410018f00000000013b001900000003050003670000145f0000613d000000000605034f00000000070b0019000000006806043c0000000007870436000000000017004b0000145b0000c13d000000000004004b000014870000613d000000000335034f0000000304400210000000000501043300000000054501cf000000000545022f000000000303043b0000010004400089000000000343022f00000000034301cf000000000353019f0000000000310435000014870000013d000005990030009c00000599030080410000006003300210000005990020009c00000599020080410000004002200210000000000223019f000005990010009c0000059901008041000000c001100210000000000112019f00000000020a0019000200000009001d00010000000a001d165f16550000040f000000010a0000290000000209000029000000010220018f00030000000103550000006001100270000105990010019d0000059901100197000000000001004b000014430000c13d000000600c000039000000800b00003900000000030c0433000000000002004b000014b70000613d000000000003004b000014a20000c13d00020000000c001d00010000000b001d000006150100004100000000001004430000000400a004430000000001000414000005990010009c0000059901008041000000c00110021000000616011001c70000800202000039165f165a0000040f0000000100200190000014e90000613d000000000101043b000000000001004b0000000201000029000014ea0000613d0000000003010433000000000003004b000000010b000029000014ae0000613d000005e70030009c000014b50000213d0000001f0030008c000014b50000a13d00000000010b0433000000000001004b0000000002000039000000010200c039000000000021004b000014b50000c13d000000000001004b000014cd0000613d000000000001042d0000061a01000041000000000010043f0000004101000039000000040010043f000005e001000041000016610001043000000000010000190000166100010430000000000003004b000014e10000c13d0000000001090019000000400400043d000200000004001d000005a80200004100000000002404350000000403400039000000200200003900000000002304350000002402400039165f11cb0000040f00000002020000290000000001210049000005990010009c0000059901008041000005990020009c000005990200804100000060011002100000004002200210000000000121019f0000166100010430000000400100043d00000064021000390000063803000041000000000032043500000044021000390000063903000041000000000032043500000024021000390000002a030000390000000000320435000005a8020000410000000000210435000000040210003900000020030000390000000000320435000005990010009c000005990100804100000040011002100000063a011001c700001661000104300000059900b0009c000005990b0080410000004002b00210000005990030009c00000599030080410000006001300210000000000121019f0000166100010430000000000001042f000000400100043d00000044021000390000063b03000041000000000032043500000024021000390000001d030000390000000000320435000005a8020000410000000000210435000000040210003900000020030000390000000000320435000005990010009c00000599010080410000004001100210000005a9011001c70000166100010430000000000010043f0000000601000039000000200010043f0000000001000414000005990010009c0000059901008041000000c001100210000005a2011001c70000801002000039165f165a0000040f00000001002001900000150d0000613d000000000101043b000000000101041a000000000001004b0000000001000039000000010100c039000000000001042d000000000100001900001661000104300003000000000002000100000002001d000300000001001d000000000101041a000200000001001d000005f00100004100000000001004430000000001000414000005990010009c0000059901008041000000c001100210000005f1011001c70000800b02000039165f165a0000040f0000000100200190000015710000613d000000020800002900000080028002700000059903200197000000000201043b000000000532004b00000003070000290000156b0000413d00000001017000390000152a0000c13d000000000207041a0000153c0000013d000000000301041a000000800630027000000000045600a900000000055400d9000000000065004b0000156b0000c13d000005ea05800197000000000054001a0000156b0000413d000005ea033001970000000004540019000000000043004b0000000003048019000005a0048001970000008002200210000005f402200197000000000242019f000000000232019f000000010600002900000020036000390000000004030433000005ea04400197000005ea05200197000000000054004b00000000050440190000063c02200197000000000225019f0000000005060433000000000005004b0000000005000019000005f50500c041000000000252019f000000000027041b000000400260003900000000050204330000008005500210000000000445019f000000000041041b0000000001000039000000010100c039000000400400043d00000000011404360000000003030433000005ea0330019700000000003104350000000001020433000005ea0110019700000040024000390000000000120435000005990040009c000005990400804100000040014002100000000002000414000005990020009c0000059902008041000000c002200210000000000112019f0000063d011001c70000800d0200003900000001030000390000063e04000041165f16550000040f0000000100200190000015720000613d000000000001042d0000061a01000041000000000010043f0000001101000039000000040010043f000005e0010000410000166100010430000000000001042f000000000100001900001661000104300000059c04400197000000400510003900000000004504350000002004100039000000000034043500000000002104350000006001100039000000000001042d0006000000000002000000000401041a000005fe00400198000015d00000613d000000000002004b000015d00000613d000600000004001d000500000002001d000200000003001d000300000001001d0000000101100039000100000001001d000000000101041a000400000001001d000005f00100004100000000001004430000000001000414000005990010009c0000059901008041000000c001100210000005f1011001c70000800b02000039165f165a0000040f0000000100200190000015d10000613d000000060300002900000080023002700000059902200197000000000101043b000000000421004b000015ec0000413d000005ea033001970000000405000029000005ea02500197000015a20000c13d00000005040000290000000305000029000015b60000013d000000000023004b000015f40000213d000000800650027000000000056400a900000000044500d9000000000064004b000015ec0000c13d000000000035001a000015ec0000413d00000000033500190000008001100210000005f4011001970000000305000029000000000405041a0000063f04400197000000000114019f000000000015041b000000000032004b00000000030240190000000504000029000000000042004b000015d20000413d000000000143004b000015e30000413d000005ea01100197000000000205041a0000064102200197000000000112019f000000000015041b000000400100043d0000000000410435000005990010009c000005990100804100000040011002100000000002000414000005990020009c0000059902008041000000c002200210000000000112019f000005a4011001c70000800d0200003900000001030000390000064204000041165f16550000040f0000000100200190000015f20000613d000000000001042d000000000001042f000000400100043d0000000004010019000000040110003900000002030000290000059c00300198000015fc0000c13d000006460300004100000000003404350000000000210435000000240140003900000005020000290000000000210435000005990040009c0000059904008041000000400140021000000620011001c700001661000104300000000101000029000000000101041a0000008001100272000015ec0000613d00000005043000690000000002140019000000010220008a000000000042004b000016010000813d0000061a01000041000000000010043f0000001101000039000000040010043f000005e001000041000016610001043000000000010000190000166100010430000000400100043d00000640020000410000000000210435000005990010009c00000599010080410000004001100210000005e5011001c700001661000104300000064503000041000600000004001d00000000003404350000000503000029000016150000013d00000000021200d9000000400100043d0000000005010019000000040110003900000002040000290000059c00400198000016120000c13d00000644040000410000000000450435000000000021043500000024015000390000000000310435000005990050009c0000059905008041000000400150021000000620011001c700001661000104300000064304000041000600000005001d00000000004504350000000204000029165f15740000040f00000006020000290000000001210049000005990010009c00000599010080410000006001100210000005990020009c00000599020080410000004002200210000000000121019f0000166100010430000000000001042f000005990010009c00000599010080410000004001100210000005990020009c00000599020080410000006002200210000000000112019f0000000002000414000005990020009c0000059902008041000000c002200210000000000112019f000005d9011001c70000801002000039165f165a0000040f0000000100200190000016350000613d000000000101043b000000000001042d0000000001000019000016610001043000000000050100190000000000200443000000050030008c000016450000413d000000040100003900000000020000190000000506200210000000000664001900000005066002700000000006060031000000000161043a0000000102200039000000000031004b0000163d0000413d000005990030009c000005990300804100000060013002100000000002000414000005990020009c0000059902008041000000c002200210000000000112019f00000647011001c70000000002050019165f165a0000040f0000000100200190000016540000613d000000000101043b000000000001042d000000000001042f00001658002104210000000102000039000000000001042d0000000002000019000000000001042d0000165d002104230000000102000039000000000001042d0000000002000019000000000001042d0000165f00000432000016600001042e00001661000104300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000000000000000000000000001ffffffe000000000000000000000000000000000000000000000000000000000ffffffe0000000000000000000000000ffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffffff80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0ffffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffdf0200000000000000000000000000000000000040000000000000000000000000bfa87805ed57dc1f0d489ce33be4c4577d74ccde357eeeee058a32c55c44a53202000000000000000000000000000000000000200000000000000000000000002640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8000000020000000000000000000000000000014000000100000000000000000043616e6e6f7420736574206f776e657220746f207a65726f000000000000000008c379a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000000000000000000000008da5cb5a00000000000000000000000000000000000000000000000000000000c4bffe2a00000000000000000000000000000000000000000000000000000000dc0bd97000000000000000000000000000000000000000000000000000000000eb521a4b00000000000000000000000000000000000000000000000000000000eb521a4c00000000000000000000000000000000000000000000000000000000f2fde38b00000000000000000000000000000000000000000000000000000000dc0bd97100000000000000000000000000000000000000000000000000000000e0351e1300000000000000000000000000000000000000000000000000000000cf7401f200000000000000000000000000000000000000000000000000000000cf7401f300000000000000000000000000000000000000000000000000000000db6327dc00000000000000000000000000000000000000000000000000000000c4bffe2b00000000000000000000000000000000000000000000000000000000c75eea9c00000000000000000000000000000000000000000000000000000000b0f479a000000000000000000000000000000000000000000000000000000000bb98546a00000000000000000000000000000000000000000000000000000000bb98546b00000000000000000000000000000000000000000000000000000000c0d7865500000000000000000000000000000000000000000000000000000000b0f479a100000000000000000000000000000000000000000000000000000000b794658000000000000000000000000000000000000000000000000000000000a7cd63b600000000000000000000000000000000000000000000000000000000a7cd63b700000000000000000000000000000000000000000000000000000000af58d59f000000000000000000000000000000000000000000000000000000008da5cb5b000000000000000000000000000000000000000000000000000000009a4575b90000000000000000000000000000000000000000000000000000000054c8a4f20000000000000000000000000000000000000000000000000000000078a010b1000000000000000000000000000000000000000000000000000000007d54534d000000000000000000000000000000000000000000000000000000007d54534e000000000000000000000000000000000000000000000000000000008926f54f0000000000000000000000000000000000000000000000000000000078a010b20000000000000000000000000000000000000000000000000000000079ba5097000000000000000000000000000000000000000000000000000000006cfd1552000000000000000000000000000000000000000000000000000000006cfd1553000000000000000000000000000000000000000000000000000000006d3d1a580000000000000000000000000000000000000000000000000000000054c8a4f300000000000000000000000000000000000000000000000000000000663200870000000000000000000000000000000000000000000000000000000021df0da60000000000000000000000000000000000000000000000000000000039077536000000000000000000000000000000000000000000000000000000003907753700000000000000000000000000000000000000000000000000000000432a6ba30000000000000000000000000000000000000000000000000000000021df0da700000000000000000000000000000000000000000000000000000000240028e8000000000000000000000000000000000000000000000000000000000a861f29000000000000000000000000000000000000000000000000000000000a861f2a00000000000000000000000000000000000000000000000000000000181f5a770000000000000000000000000000000000000000000000000000000001ffc9a7000000000000000000000000000000000000000000000000000000000a2fd4930200000000000000000000000000000000000000000000000000000000000000ed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127843616e6e6f74207472616e7366657220746f2073656c660000000000000000000000000000000000000000000000000000000064000000800000000000000000310ab089e4439a4c15d089f94afb7896ff553aecb10793d0ab882de59d99a32e02000002000000000000000000000000000000440000000000000000000000008e4a23d600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002400000000000000000000000023b872dd00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffff5fc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088e93f8fa400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000200000008000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffff3f000000000000000000000000000000000000000000000000ffffffffffffff9f00000000000000000000000000000000ffffffffffffffffffffffffffffffff8020d12400000000000000000000000000000000000000000000000000000000d68af9cc00000000000000000000000000000000000000000000000000000000433fc33d000000000000000000000000000000000000000000000000000000001d5ad3c500000000000000000000000000000000000000000000000000000000fc949c7b4a13586e39d89eead2f38644f9fb3efb5a0490b14f8fc0ceab44c250796b89b91644bc98cd93958e4c9038275d622183e25ac5af08cc6b5d955391320200000200000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffff7fffffffffffffffffffffff000000000000000000000000000000000000000000000000000000000000000000ffffffff0000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000008d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c28579befe00000000000000000000000000000000000000000000000000000000fc949c7b4a13586e39d89eead2f38644f9fb3efb5a0490b14f8fc0ceab44c2515204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599161e670e4b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002400000140000000000000000002000000000000000000000000000000000000e00000000000000000000000000350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b0000000000000000000000ff0000000000000000000000000000000000000000036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0000000000000000000000000000000000000000000000000ffffffffffffffc1ffffffffffffffffffffffffffffffffffffffffffffffff000000000000008000000000000000000000000000000000000000000000003fffffffffffffffe0020000000000000000000000000000000000004000000080000000000000000002dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16840000000000000000000000000000000000000004000000800000000000000000405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace000000000000000000000000000000000000000000000000fffffffffffffe9f961c9a4f000000000000000000000000000000000000000000000000000000002cbc26bb000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffff0000000000000000000000000000000053ad11d800000000000000000000000000000000000000000000000000000000d0d2597600000000000000000000000000000000000000000000000000000000a8d87a3b00000000000000000000000000000000000000000000000000000000728fe07b000000000000000000000000000000000000000000000000000000009f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008a9902c7e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000004d7573742062652070726f706f736564206f776e6572000000000000000000008be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0db4d6220746a38cbc5335f7e108f7de80f482f4d23350253dfd0917df75a14bf1806aa1896bbf26568e884a7374b41e002500962caba6a15023a8d90e8508b8302000002000000000000000000000000000000240000000000000000000000000a861f2a000000000000000000000000000000000000000000000000000000006fa7abcf1345d1d478e5ea0da6b5f26a90eadb0546ef15ed3833944fbfd1db624f6e6c792063616c6c61626c65206279206f776e6572000000000000000000004e487b7100000000000000000000000000000000000000000000000000000000bfa87805ed57dc1f0d489ce33be4c4577d74ccde357eeeee058a32c55c44a533800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756635f4a7b300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000fffffffffffffe5f83826b2b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000440000000000000000000000002d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f5224eb47e5000000000000000000000000000000000000000000000000000000004c6f636b52656c65617365546f6b656e506f6f6c20312e352e300000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000002400000080000000000000000070a0823100000000000000000000000000000000000000000000000000000000c2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719bb55fd270000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffaff2afbeffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1d4056600000000000000000000000000000000000000000000000000000000aff2afbf0000000000000000000000000000000000000000000000000000000001ffc9a7000000000000000000000000000000000000000000000000000000000e64dd2900000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000000000000000000000000000000000000000000000ffffffffffffffe0000000000000000000000000000000000000000000000000ffffffffffffffc0000000000000000000000000000000000000000000000000ffffffffffffff60a9059cbb00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffff805361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646f742073756363656564000000000000000000000000000000000000000000005361666545524332303a204552433230206f7065726174696f6e20646964206e0000000000000000000000000000000000000084000000000000000000000000416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000ffffffffffffffffffffff00ffffffff0000000000000000000000000000000002000000000000000000000000000000000000600000000000000000000000009ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19ffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9725942a00000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffff000000000000000000000000000000001871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690ad0c8d23a0000000000000000000000000000000000000000000000000000000015279c08000000000000000000000000000000000000000000000000000000001a76572a00000000000000000000000000000000000000000000000000000000f94ebcd100000000000000000000000000000000000000000000000000000000020000020000000000000000000000000000000000000000000000000000000017d7e45b984975a46dbdaf6d028174dba298b890a7b749a087b67443a2c5fa43")

func (_LockReleaseTokenPool *LockReleaseTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _LockReleaseTokenPool.abi.Events["AllowListAdd"].ID:
		return _LockReleaseTokenPool.ParseAllowListAdd(log)
	case _LockReleaseTokenPool.abi.Events["AllowListRemove"].ID:
		return _LockReleaseTokenPool.ParseAllowListRemove(log)
	case _LockReleaseTokenPool.abi.Events["Burned"].ID:
		return _LockReleaseTokenPool.ParseBurned(log)
	case _LockReleaseTokenPool.abi.Events["ChainAdded"].ID:
		return _LockReleaseTokenPool.ParseChainAdded(log)
	case _LockReleaseTokenPool.abi.Events["ChainConfigured"].ID:
		return _LockReleaseTokenPool.ParseChainConfigured(log)
	case _LockReleaseTokenPool.abi.Events["ChainRemoved"].ID:
		return _LockReleaseTokenPool.ParseChainRemoved(log)
	case _LockReleaseTokenPool.abi.Events["ConfigChanged"].ID:
		return _LockReleaseTokenPool.ParseConfigChanged(log)
	case _LockReleaseTokenPool.abi.Events["LiquidityAdded"].ID:
		return _LockReleaseTokenPool.ParseLiquidityAdded(log)
	case _LockReleaseTokenPool.abi.Events["LiquidityRemoved"].ID:
		return _LockReleaseTokenPool.ParseLiquidityRemoved(log)
	case _LockReleaseTokenPool.abi.Events["LiquidityTransferred"].ID:
		return _LockReleaseTokenPool.ParseLiquidityTransferred(log)
	case _LockReleaseTokenPool.abi.Events["Locked"].ID:
		return _LockReleaseTokenPool.ParseLocked(log)
	case _LockReleaseTokenPool.abi.Events["Minted"].ID:
		return _LockReleaseTokenPool.ParseMinted(log)
	case _LockReleaseTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _LockReleaseTokenPool.ParseOwnershipTransferRequested(log)
	case _LockReleaseTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _LockReleaseTokenPool.ParseOwnershipTransferred(log)
	case _LockReleaseTokenPool.abi.Events["Released"].ID:
		return _LockReleaseTokenPool.ParseReleased(log)
	case _LockReleaseTokenPool.abi.Events["RemotePoolSet"].ID:
		return _LockReleaseTokenPool.ParseRemotePoolSet(log)
	case _LockReleaseTokenPool.abi.Events["RouterUpdated"].ID:
		return _LockReleaseTokenPool.ParseRouterUpdated(log)
	case _LockReleaseTokenPool.abi.Events["TokensConsumed"].ID:
		return _LockReleaseTokenPool.ParseTokensConsumed(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (LockReleaseTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (LockReleaseTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (LockReleaseTokenPoolBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
}

func (LockReleaseTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (LockReleaseTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (LockReleaseTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (LockReleaseTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (LockReleaseTokenPoolLiquidityAdded) Topic() common.Hash {
	return common.HexToHash("0xc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088")
}

func (LockReleaseTokenPoolLiquidityRemoved) Topic() common.Hash {
	return common.HexToHash("0xc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719")
}

func (LockReleaseTokenPoolLiquidityTransferred) Topic() common.Hash {
	return common.HexToHash("0x6fa7abcf1345d1d478e5ea0da6b5f26a90eadb0546ef15ed3833944fbfd1db62")
}

func (LockReleaseTokenPoolLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (LockReleaseTokenPoolMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (LockReleaseTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (LockReleaseTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (LockReleaseTokenPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (LockReleaseTokenPoolRemotePoolSet) Topic() common.Hash {
	return common.HexToHash("0xdb4d6220746a38cbc5335f7e108f7de80f482f4d23350253dfd0917df75a14bf")
}

func (LockReleaseTokenPoolRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (LockReleaseTokenPoolTokensConsumed) Topic() common.Hash {
	return common.HexToHash("0x1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a")
}

func (_LockReleaseTokenPool *LockReleaseTokenPool) Address() common.Address {
	return _LockReleaseTokenPool.address
}

type CustomTransaction struct {
	*types.Transaction
	CustomHash common.Hash
}

func (tx *CustomTransaction) Hash() common.Hash {
	return tx.CustomHash
}

func ConvertToTransaction(resp zktypes.TransactionResponse) *CustomTransaction {
	dtx := &types.DynamicFeeTx{
		ChainID:   resp.ChainID.ToInt(),
		Nonce:     uint64(resp.Nonce),
		GasTipCap: resp.MaxPriorityFeePerGas.ToInt(),
		GasFeeCap: resp.MaxFeePerGas.ToInt(),
		To:        &resp.To,
		Value:     resp.Value.ToInt(),
		Data:      resp.Data,
		Gas:       uint64(resp.Gas),
	}

	tx := types.NewTx(dtx)
	customTransaction := CustomTransaction{Transaction: tx, CustomHash: resp.Hash}
	return &customTransaction
}

func DeployZkSyncLockReleaseTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, params ...interface{}) (common.Address, *CustomTransaction, *LockReleaseTokenPool, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	fmt.Println("Deploying zksync contract")
	zksyncClient := zkSyncClient.NewClient(client.Client())
	fmt.Println("getting wallet")
	wallet := auth.Context.Value("wallet").(*zkSyncAccounts.Wallet)
	fmt.Println("got wallet")
	fmt.Println("getting bytes")
	decodedBytes := common.FromHex(LockReleaseTokenPoolZkBin)
	fmt.Println("deploying")
	LockReleaseTokenPoolAbi, err := LockReleaseTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	constructor, err := LockReleaseTokenPoolAbi.Pack("", params...)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	hash, err := wallet.DeployWithCreate(nil, zkSyncAccounts.CreateTransaction{
		Bytecode: decodedBytes,
		Calldata: constructor,
	})
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	fmt.Println("hash of tx", hash)
	receipt, err := zksyncClient.WaitMined(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	tx, _, err := zksyncClient.TransactionByHash(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	fmt.Println("tx hash", tx.Hash)

	ethTx := ConvertToTransaction(*tx)
	address := receipt.ContractAddress

	parsed, err := LockReleaseTokenPoolMetaData.GetAbi()
	contractBind := bind.NewBoundContract(address, *parsed, backend, backend, backend)

	contractReturn := &LockReleaseTokenPool{address: address, abi: *parsed, LockReleaseTokenPoolCaller: LockReleaseTokenPoolCaller{contract: contractBind}, LockReleaseTokenPoolTransactor: LockReleaseTokenPoolTransactor{contract: contractBind}, LockReleaseTokenPoolFilterer: LockReleaseTokenPoolFilterer{contract: contractBind}}

	return address, ethTx, contractReturn, err
}

type LockReleaseTokenPoolInterface interface {
	CanAcceptLiquidity(opts *bind.CallOpts) (bool, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRebalancer(opts *bind.CallOpts) (common.Address, error)

	GetRemotePool(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

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

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRebalancer(opts *bind.TransactOpts, rebalancer common.Address) (*types.Transaction, error)

	SetRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferLiquidity(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*LockReleaseTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*LockReleaseTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*LockReleaseTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*LockReleaseTokenPoolAllowListRemove, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*LockReleaseTokenPoolBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*LockReleaseTokenPoolBurned, error)

	FilterChainAdded(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*LockReleaseTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*LockReleaseTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*LockReleaseTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*LockReleaseTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*LockReleaseTokenPoolConfigChanged, error)

	FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolLiquidityAddedIterator, error)

	WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityAdded, provider []common.Address, amount []*big.Int) (event.Subscription, error)

	ParseLiquidityAdded(log types.Log) (*LockReleaseTokenPoolLiquidityAdded, error)

	FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolLiquidityRemovedIterator, error)

	WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityRemoved, provider []common.Address, amount []*big.Int) (event.Subscription, error)

	ParseLiquidityRemoved(log types.Log) (*LockReleaseTokenPoolLiquidityRemoved, error)

	FilterLiquidityTransferred(opts *bind.FilterOpts, from []common.Address) (*LockReleaseTokenPoolLiquidityTransferredIterator, error)

	WatchLiquidityTransferred(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityTransferred, from []common.Address) (event.Subscription, error)

	ParseLiquidityTransferred(log types.Log) (*LockReleaseTokenPoolLiquidityTransferred, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*LockReleaseTokenPoolLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*LockReleaseTokenPoolLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LockReleaseTokenPoolMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*LockReleaseTokenPoolMinted, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockReleaseTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*LockReleaseTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockReleaseTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*LockReleaseTokenPoolOwnershipTransferred, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LockReleaseTokenPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*LockReleaseTokenPoolReleased, error)

	FilterRemotePoolSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolRemotePoolSetIterator, error)

	WatchRemotePoolSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRemotePoolSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolSet(log types.Log) (*LockReleaseTokenPoolRemotePoolSet, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*LockReleaseTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*LockReleaseTokenPoolRouterUpdated, error)

	FilterTokensConsumed(opts *bind.FilterOpts) (*LockReleaseTokenPoolTokensConsumedIterator, error)

	WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolTokensConsumed) (event.Subscription, error)

	ParseTokensConsumed(log types.Log) (*LockReleaseTokenPoolTokensConsumed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
