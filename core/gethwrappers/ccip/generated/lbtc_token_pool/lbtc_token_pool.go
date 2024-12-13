// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lbtc_token_pool

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
	RemotePoolAddresses       [][]byte
	RemoteTokenAddress        []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

var LombardTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lbtc_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ccipRouter_\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowlist_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"rmnProxy_\",\"type\":\"address\"},{\"internalType\":\"contractCLAdapter\",\"name\":\"adapter_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"attestationEnable_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowListNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotARampOnRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotTransferToSelf\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"DisabledNonZeroRateLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"expected\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"actual\",\"type\":\"uint8\"}],\"name\":\"InvalidDecimalArgs\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"InvalidRateLimitRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolData\",\"type\":\"bytes\"}],\"name\":\"InvalidRemoteChainDecimals\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidRemotePoolForChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidSourcePoolAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeProposedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"NonExistentChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"remoteDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"localDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"remoteAmount\",\"type\":\"uint256\"}],\"name\":\"OverflowDetected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateLimitMustBeDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remoteToken\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rateLimitAdmin\",\"type\":\"address\"}],\"name\":\"RateLimitAdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"RemotePoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"RemotePoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRouter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensConsumed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adapter\",\"outputs\":[{\"internalType\":\"contractCLAdapter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"addRemotePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"remoteTokenAddress\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"internalType\":\"structTokenPool.ChainUpdate[]\",\"name\":\"chainsToAdd\",\"type\":\"tuple[]\"}],\"name\":\"applyChainUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentInboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentOutboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRateLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRemotePools\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRemoteToken\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRmnProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedChains\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAttestationEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"isRemotePool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"isSupportedChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isSupportedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"internalType\":\"structPool.LockOrBurnInV1\",\"name\":\"lockOrBurnIn\",\"type\":\"tuple\"}],\"name\":\"lockOrBurn\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destPoolData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.LockOrBurnOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"originalSender\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainTokenData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.ReleaseOrMintInV1\",\"name\":\"releaseOrMintIn\",\"type\":\"tuple\"}],\"name\":\"releaseOrMint\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"destinationAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"removeRemotePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundConfig\",\"type\":\"tuple\"}],\"name\":\"setChainRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rateLimitAdmin\",\"type\":\"address\"}],\"name\":\"setRateLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620046dd380380620046dd8339810160408190526200003591620005d1565b856008858588336000816200005d57604051639b15e16f60e01b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0384811691909117909155811615620000905762000090816200021b565b50506001600160a01b0385161580620000b057506001600160a01b038116155b80620000c357506001600160a01b038216155b15620000e2576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b03808616608081905290831660c0526040805163313ce56760e01b8152905163313ce567916004808201926020929091908290030181865afa92505050801562000152575060408051601f3d908101601f191682019092526200014f9181019062000708565b60015b1562000192578060ff168560ff161462000190576040516332ad3e0760e11b815260ff80871660048301528216602482015260440160405180910390fd5b505b60ff841660a052600480546001600160a01b0319166001600160a01b038316179055825115801560e052620001dc57604080516000815260208101909152620001dc908462000295565b5050600a8054941515600160a01b026001600160a81b03199095166001600160a01b039096169590951793909317909355506200078295505050505050565b336001600160a01b038216036200024557604051636d6c4ee560e11b815260040160405180910390fd5b600080546001600160a01b0319166001600160a01b03838116918217835560015460405192939116917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60e051620002b6576040516335f4a7b360e01b815260040160405180910390fd5b60005b825181101562000341576000838281518110620002da57620002da62000734565b60209081029190910101519050620002f4600282620003f2565b1562000337576040516001600160a01b03821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b50600101620002b9565b5060005b8151811015620003ed57600082828151811062000366576200036662000734565b6020026020010151905060006001600160a01b0316816001600160a01b031603620003925750620003e4565b6200039f60028262000412565b15620003e2576040516001600160a01b03821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b60010162000345565b505050565b600062000409836001600160a01b03841662000429565b90505b92915050565b600062000409836001600160a01b0384166200052d565b6000818152600183016020526040812054801562000522576000620004506001836200074a565b855490915060009062000466906001906200074a565b9050808214620004d25760008660000182815481106200048a576200048a62000734565b9060005260206000200154905080876000018481548110620004b057620004b062000734565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080620004e657620004e66200076c565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506200040c565b60009150506200040c565b600081815260018301602052604081205462000576575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556200040c565b5060006200040c565b6001600160a01b03811681146200059557600080fd5b50565b8051620005a5816200057f565b919050565b634e487b7160e01b600052604160045260246000fd5b80518015158114620005a557600080fd5b60008060008060008060c08789031215620005eb57600080fd5b8651620005f8816200057f565b809650506020808801516200060d816200057f565b60408901519096506001600160401b03808211156200062b57600080fd5b818a0191508a601f8301126200064057600080fd5b815181811115620006555762000655620005aa565b8060051b604051601f19603f830116810181811085821117156200067d576200067d620005aa565b60405291825284820192508381018501918d8311156200069c57600080fd5b938501935b82851015620006c557620006b58562000598565b84529385019392850192620006a1565b809950505050505050620006dc6060880162000598565b9250620006ec6080880162000598565b9150620006fc60a08801620005c0565b90509295509295509295565b6000602082840312156200071b57600080fd5b815160ff811681146200072d57600080fd5b9392505050565b634e487b7160e01b600052603260045260246000fd5b818103818111156200040c57634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60805160a05160c05160e051613ee9620007f46000396000818161056701528181611d1501526126c401526000818161054101528181611b0a0152612001015260006102c601526000818161024e0152818161028201528181610bec0152818161265a01526128af0152613ee96000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c80639a4575b911610104578063c0d78655116100a2578063dc0bd97111610071578063dc0bd9711461053f578063e0351e1314610565578063e8a1da171461058b578063f2fde38b1461059e57600080fd5b8063c0d78655146104f1578063c4bffe2b14610504578063c75eea9c14610519578063cf7401f31461052c57600080fd5b8063acfecf91116100de578063acfecf9114610431578063af58d59f14610444578063b0f479a1146104b3578063b7946580146104d157600080fd5b80639a4575b9146103dc578063a42a7b8b146103fc578063a7cd63b71461041c57600080fd5b806354c8a4f31161017c57806379ba50971161014b57806379ba5097146103905780637d54534e146103985780638926f54f146103ab5780638da5cb5b146103be57600080fd5b806354c8a4f31461032557806355b961561461033a57806362ddd3c41461035f5780636d3d1a581461037257600080fd5b8063240028e8116101b8578063240028e81461027257806324f65ee7146102bf57806339077537146102f05780634c5ef0ed1461031257600080fd5b806301ffc9a7146101df57806303eadcfc1461020757806321df0da71461024c575b600080fd5b6101f26101ed36600461306f565b6105b1565b60405190151581526020015b60405180910390f35b600a546102279073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101fe565b7f0000000000000000000000000000000000000000000000000000000000000000610227565b6101f26102803660046130d3565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff90811691161490565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101fe565b6103036102fe3660046130f0565b610696565b604051905181526020016101fe565b6101f2610320366004613142565b6108b5565b610338610333366004613213565b6108ff565b005b600a546101f29074010000000000000000000000000000000000000000900460ff1681565b61033861036d366004613142565b61097a565b60095473ffffffffffffffffffffffffffffffffffffffff16610227565b610338610a17565b6103386103a63660046130d3565b610ae5565b6101f26103b936600461327f565b610b66565b60015473ffffffffffffffffffffffffffffffffffffffff16610227565b6103ef6103ea36600461329c565b610b7d565b6040516101fe9190613345565b61040f61040a36600461327f565b610e59565b6040516101fe919061339c565b610424610fc4565b6040516101fe919061341e565b61033861043f366004613142565b610fd5565b61045761045236600461327f565b6110ed565b6040516101fe919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b60045473ffffffffffffffffffffffffffffffffffffffff16610227565b6104e46104df36600461327f565b6111c2565b6040516101fe9190613478565b6103386104ff3660046130d3565b611272565b61050c61134d565b6040516101fe919061348b565b61045761052736600461327f565b611405565b61033861053a3660046135ef565b6114d7565b7f0000000000000000000000000000000000000000000000000000000000000000610227565b7f00000000000000000000000000000000000000000000000000000000000000006101f2565b610338610599366004613213565b61155b565b6103386105ac3660046130d3565b611a6d565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167faff2afbf00000000000000000000000000000000000000000000000000000000148061064457507fffffffff0000000000000000000000000000000000000000000000000000000082167f0e64dd2900000000000000000000000000000000000000000000000000000000145b8061069057507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b6040805160208101909152600081526106ae82611a81565b600a5460009074010000000000000000000000000000000000000000900460ff161561077757600a5473ffffffffffffffffffffffffffffffffffffffff16635391a405610702604086016020870161327f565b61070f60e0870187613636565b6040518463ffffffff1660e01b815260040161072d939291906136e4565b6020604051808303816000875af115801561074c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107709190613708565b9050610816565b600a5473ffffffffffffffffffffffffffffffffffffffff16630c373d746107a5604086016020870161327f565b6107b260c0870187613636565b6040518463ffffffff1660e01b81526004016107d0939291906136e4565b6020604051808303816000875af11580156107ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108139190613708565b90505b61082660608401604085016130d3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f08367ffffffffffffffff1660405161088e91815260200190565b60405180910390a3604080516020810190915267ffffffffffffffff909116815292915050565b60006108f783836040516108ca929190613725565b604080519182900390912067ffffffffffffffff8716600090815260076020529190912060050190611ca5565b949350505050565b610907611cc0565b61097484848080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808802828101820190935287825290935087925086918291850190849080828437600092019190915250611d1392505050565b50505050565b610982611cc0565b61098b83610b66565b6109d2576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024015b60405180910390fd5b610a128383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611ec992505050565b505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610a68576040517f02b543c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560008054909116815560405173ffffffffffffffffffffffffffffffffffffffff909216929183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610aed611cc0565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d091749060200160405180910390a150565b6000610690600567ffffffffffffffff8416611ca5565b6040805180820190915260608082526020820152610b9a82611fc3565b600a546040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152606084013560248201527f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb906044016020604051808303816000875af1158015610c37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c5b9190613735565b50600a54600090819073ffffffffffffffffffffffffffffffffffffffff1663550e7ab2610c8f604087016020880161327f565b610c998780613636565b88606001356040518563ffffffff1660e01b8152600401610cbd9493929190613752565b6000604051808303816000875af1158015610cdc573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610d2291908101906137cd565b9092509050610d3760608501604086016130d3565b73ffffffffffffffffffffffffffffffffffffffff167f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df783604051610d7e91815260200190565b60405180910390a2600a5460609074010000000000000000000000000000000000000000900460ff1615610e2457600282604051610dbc9190613855565b602060405180830381855afa158015610dd9573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190610dfc9190613871565b604051602001610e0e91815260200190565b6040516020818303038152906040529050610e27565b50805b6040518060400160405280610e488760200160208101906104df919061327f565b815260200191909152949350505050565b67ffffffffffffffff8116600090815260076020526040812060609190610e829060050161214f565b90506000815167ffffffffffffffff811115610ea057610ea06134cd565b604051908082528060200260200182016040528015610ed357816020015b6060815260200190600190039081610ebe5790505b50905060005b8251811015610fbc5760086000848381518110610ef857610ef861388a565b602002602001015181526020019081526020016000208054610f19906138b9565b80601f0160208091040260200160405190810160405280929190818152602001828054610f45906138b9565b8015610f925780601f10610f6757610100808354040283529160200191610f92565b820191906000526020600020905b815481529060010190602001808311610f7557829003601f168201915b5050505050828281518110610fa957610fa961388a565b6020908102919091010152600101610ed9565b509392505050565b6060610fd0600261214f565b905090565b610fdd611cc0565b610fe683610b66565b611028576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024016109c9565b611068828260405161103b929190613725565b604080519182900390912067ffffffffffffffff861660009081526007602052919091206005019061215c565b6110a4578282826040517f74f23c7c0000000000000000000000000000000000000000000000000000000081526004016109c9939291906136e4565b8267ffffffffffffffff167f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7683836040516110e092919061390c565b60405180910390a2505050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845260028201546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff16151594820194909452600390910154808416606083015291909104909116608082015261069090612168565b67ffffffffffffffff811660009081526007602052604090206004018054606091906111ed906138b9565b80601f0160208091040260200160405190810160405280929190818152602001828054611219906138b9565b80156112665780601f1061123b57610100808354040283529160200191611266565b820191906000526020600020905b81548152906001019060200180831161124957829003601f168201915b50505050509050919050565b61127a611cc0565b73ffffffffffffffffffffffffffffffffffffffff81166112c7576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684910160405180910390a15050565b6060600061135b600561214f565b90506000815167ffffffffffffffff811115611379576113796134cd565b6040519080825280602002602001820160405280156113a2578160200160208202803683370190505b50905060005b82518110156113fe578281815181106113c3576113c361388a565b60200260200101518282815181106113dd576113dd61388a565b67ffffffffffffffff909216602092830291909101909101526001016113a8565b5092915050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845281546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff16151594820194909452600190910154808416606083015291909104909116608082015261069090612168565b60095473ffffffffffffffffffffffffffffffffffffffff163314801590611517575060015473ffffffffffffffffffffffffffffffffffffffff163314155b15611550576040517f8e4a23d60000000000000000000000000000000000000000000000000000000081523360048201526024016109c9565b610a1283838361221a565b611563611cc0565b60005b838110156117505760008585838181106115825761158261388a565b9050602002016020810190611597919061327f565b90506115ae600567ffffffffffffffff831661215c565b6115f0576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016109c9565b67ffffffffffffffff811660009081526007602052604081206116159060050161214f565b905060005b8151811015611681576116788282815181106116385761163861388a565b6020026020010151600760008667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060050161215c90919063ffffffff16565b5060010161161a565b5067ffffffffffffffff8216600090815260076020526040812080547fffffffffffffffffffffff000000000000000000000000000000000000000000908116825560018201839055600282018054909116905560038101829055906116ea6004830182613002565b60058201600081816116fc828261303c565b505060405167ffffffffffffffff871681527f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599169450602001925061173e915050565b60405180910390a15050600101611566565b5060005b81811015611a665760008383838181106117705761177061388a565b90506020028101906117829190613920565b61178b90613a44565b905061179c81606001516000612304565b6117ab81608001516000612304565b8060400151516000036117ea576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516118029060059067ffffffffffffffff16612441565b6118475780516040517f1d5ad3c500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016109c9565b805167ffffffffffffffff16600090815260076020908152604091829020825160a08082018552606080870180518601516fffffffffffffffffffffffffffffffff90811680865263ffffffff42168689018190528351511515878b0181905284518a0151841686890181905294518b0151841660809889018190528954740100000000000000000000000000000000000000009283027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff7001000000000000000000000000000000008087027fffffffffffffffffffffffff000000000000000000000000000000000000000094851690981788178216929092178d5592810290971760018c01558c519889018d52898e0180518d01518716808b528a8e019590955280515115158a8f018190528151909d01518716988a01899052518d0151909516979098018790526002890180549a9091029990931617179094169590951790925590920290911760038201559082015160048201906119ca9082613b54565b5060005b826020015151811015611a0e57611a068360000151846020015183815181106119f9576119f961388a565b6020026020010151611ec9565b6001016119ce565b507f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c28260000151836040015184606001518560800151604051611a549493929190613c6e565b60405180910390a15050600101611754565b5050505050565b611a75611cc0565b611a7e8161244d565b50565b611a9461028060a08301608084016130d3565b611af357611aa860a08201608083016130d3565b6040517f961c9a4f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016109c9565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016632cbc26bb611b3f604084016020850161327f565b60405160e083901b7fffffffff0000000000000000000000000000000000000000000000000000000016815260809190911b77ffffffffffffffff00000000000000000000000000000000166004820152602401602060405180830381865afa158015611bb0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bd49190613735565b15611c0b576040517f53ad11d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c23611c1e604083016020840161327f565b612511565b611c43611c36604083016020840161327f565b61032060a0840184613636565b611c8857611c5460a0820182613636565b6040517f24eb47e50000000000000000000000000000000000000000000000000000000081526004016109c992919061390c565b611a7e611c9b604083016020840161327f565b8260600135612637565b600081815260018301602052604081205415155b9392505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611d11576040517f2b5c74de00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b7f0000000000000000000000000000000000000000000000000000000000000000611d6a576040517f35f4a7b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8251811015611e00576000838281518110611d8a57611d8a61388a565b60200260200101519050611da881600261267e90919063ffffffff16565b15611df75760405173ffffffffffffffffffffffffffffffffffffffff821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b50600101611d6d565b5060005b8151811015610a12576000828281518110611e2157611e2161388a565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611e655750611ec1565b611e706002826126a0565b15611ebf5760405173ffffffffffffffffffffffffffffffffffffffff821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b600101611e04565b8051600003611f04576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160208083019190912067ffffffffffffffff8416600090815260079092526040909120611f369060050182612441565b611f705782826040517f393b8ad20000000000000000000000000000000000000000000000000000000081526004016109c9929190613d07565b6000818152600860205260409020611f888382613b54565b508267ffffffffffffffff167f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea836040516110e09190613478565b611fd661028060a08301608084016130d3565b611fea57611aa860a08201608083016130d3565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016632cbc26bb612036604084016020850161327f565b60405160e083901b7fffffffff0000000000000000000000000000000000000000000000000000000016815260809190911b77ffffffffffffffff00000000000000000000000000000000166004820152602401602060405180830381865afa1580156120a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120cb9190613735565b15612102576040517f53ad11d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61211a61211560608301604084016130d3565b6126c2565b61213261212d604083016020840161327f565b612741565b611a7e612145604083016020840161327f565b826060013561288f565b60606000611cb9836128d3565b6000611cb9838361292e565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526121f682606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff16426121da9190613d59565b85608001516fffffffffffffffffffffffffffffffff16612a21565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b61222383610b66565b612265576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024016109c9565b612270826000612304565b67ffffffffffffffff831660009081526007602052604090206122939083612a49565b61229e816000612304565b67ffffffffffffffff831660009081526007602052604090206122c49060020182612a49565b7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b8383836040516122f793929190613d6c565b60405180910390a1505050565b8151156123cf5781602001516fffffffffffffffffffffffffffffffff1682604001516fffffffffffffffffffffffffffffffff1610158061235a575060408201516fffffffffffffffffffffffffffffffff16155b1561239357816040517f8020d1240000000000000000000000000000000000000000000000000000000081526004016109c99190613def565b80156123cb576040517f433fc33d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b60408201516fffffffffffffffffffffffffffffffff16151580612408575060208201516fffffffffffffffffffffffffffffffff1615155b156123cb57816040517fd68af9cc0000000000000000000000000000000000000000000000000000000081526004016109c99190613def565b6000611cb98383612beb565b3373ffffffffffffffffffffffffffffffffffffffff82160361249c576040517fdad89dca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217835560015460405192939116917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b61251a81610b66565b61255c576040517fa9902c7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016109c9565b600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84169281019290925233602483015273ffffffffffffffffffffffffffffffffffffffff16906383826b2b90604401602060405180830381865afa1580156125db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125ff9190613735565b611a7e576040517f728fe07b0000000000000000000000000000000000000000000000000000000081523360048201526024016109c9565b67ffffffffffffffff821660009081526007602052604090206123cb90600201827f0000000000000000000000000000000000000000000000000000000000000000612c3a565b6000611cb98373ffffffffffffffffffffffffffffffffffffffff841661292e565b6000611cb98373ffffffffffffffffffffffffffffffffffffffff8416612beb565b7f000000000000000000000000000000000000000000000000000000000000000015611a7e576126f3600282612fbd565b611a7e576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016109c9565b61274a81610b66565b61278c576040517fa9902c7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016109c9565b600480546040517fa8d87a3b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84169281019290925273ffffffffffffffffffffffffffffffffffffffff169063a8d87a3b90602401602060405180830381865afa158015612805573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128299190613e2b565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611a7e576040517f728fe07b0000000000000000000000000000000000000000000000000000000081523360048201526024016109c9565b67ffffffffffffffff821660009081526007602052604090206123cb90827f0000000000000000000000000000000000000000000000000000000000000000612c3a565b60608160000180548060200260200160405190810160405280929190818152602001828054801561126657602002820191906000526020600020905b81548152602001906001019080831161290f5750505050509050919050565b60008181526001830160205260408120548015612a17576000612952600183613d59565b855490915060009061296690600190613d59565b90508082146129cb5760008660000182815481106129865761298661388a565b90600052602060002001549050808760000184815481106129a9576129a961388a565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806129dc576129dc613e48565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610690565b6000915050610690565b6000612a4085612a318486613e77565b612a3b9087613e8e565b612fec565b95945050505050565b8154600090612a7290700100000000000000000000000000000000900463ffffffff1642613d59565b90508015612b145760018301548354612aba916fffffffffffffffffffffffffffffffff80821692811691859170010000000000000000000000000000000090910416612a21565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b60208201518354612b3a916fffffffffffffffffffffffffffffffff9081169116612fec565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19906122f7908490613def565b6000818152600183016020526040812054612c3257508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610690565b506000610690565b825474010000000000000000000000000000000000000000900460ff161580612c61575081155b15612c6b57505050565b825460018401546fffffffffffffffffffffffffffffffff80831692911690600090612cb190700100000000000000000000000000000000900463ffffffff1642613d59565b90508015612d715781831115612cf3576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001860154612d2d9083908590849070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16612a21565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff167001000000000000000000000000000000004263ffffffff160217875592505b84821015612e285773ffffffffffffffffffffffffffffffffffffffff8416612dd0576040517ff94ebcd100000000000000000000000000000000000000000000000000000000815260048101839052602481018690526044016109c9565b6040517f1a76572a000000000000000000000000000000000000000000000000000000008152600481018390526024810186905273ffffffffffffffffffffffffffffffffffffffff851660448201526064016109c9565b84831015612f3b5760018681015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16906000908290612e6c9082613d59565b612e76878a613d59565b612e809190613e8e565b612e8a9190613ea1565b905073ffffffffffffffffffffffffffffffffffffffff8616612ee3576040517f15279c0800000000000000000000000000000000000000000000000000000000815260048101829052602481018690526044016109c9565b6040517fd0c8d23a000000000000000000000000000000000000000000000000000000008152600481018290526024810186905273ffffffffffffffffffffffffffffffffffffffff871660448201526064016109c9565b612f458584613d59565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515611cb9565b6000818310612ffb5781611cb9565b5090919050565b50805461300e906138b9565b6000825580601f1061301e575050565b601f016020900490600052602060002090810190611a7e9190613056565b5080546000825590600052602060002090810190611a7e91905b5b8082111561306b5760008155600101613057565b5090565b60006020828403121561308157600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611cb957600080fd5b73ffffffffffffffffffffffffffffffffffffffff81168114611a7e57600080fd5b6000602082840312156130e557600080fd5b8135611cb9816130b1565b60006020828403121561310257600080fd5b813567ffffffffffffffff81111561311957600080fd5b82016101008185031215611cb957600080fd5b67ffffffffffffffff81168114611a7e57600080fd5b60008060006040848603121561315757600080fd5b83356131628161312c565b9250602084013567ffffffffffffffff8082111561317f57600080fd5b818601915086601f83011261319357600080fd5b8135818111156131a257600080fd5b8760208285010111156131b457600080fd5b6020830194508093505050509250925092565b60008083601f8401126131d957600080fd5b50813567ffffffffffffffff8111156131f157600080fd5b6020830191508360208260051b850101111561320c57600080fd5b9250929050565b6000806000806040858703121561322957600080fd5b843567ffffffffffffffff8082111561324157600080fd5b61324d888389016131c7565b9096509450602087013591508082111561326657600080fd5b50613273878288016131c7565b95989497509550505050565b60006020828403121561329157600080fd5b8135611cb98161312c565b6000602082840312156132ae57600080fd5b813567ffffffffffffffff8111156132c557600080fd5b820160a08185031215611cb957600080fd5b60005b838110156132f25781810151838201526020016132da565b50506000910152565b600081518084526133138160208601602086016132d7565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600082516040602084015261336160608401826132fb565b905060208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848303016040850152612a4082826132fb565b600060208083016020845280855180835260408601915060408160051b87010192506020870160005b82811015613411577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08886030184526133ff8583516132fb565b945092850192908501906001016133c5565b5092979650505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561346c57835173ffffffffffffffffffffffffffffffffffffffff168352928401929184019160010161343a565b50909695505050505050565b602081526000611cb960208301846132fb565b6020808252825182820181905260009190848201906040850190845b8181101561346c57835167ffffffffffffffff16835292840192918401916001016134a7565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613543576135436134cd565b604052919050565b8015158114611a7e57600080fd5b80356fffffffffffffffffffffffffffffffff8116811461357957600080fd5b919050565b60006060828403121561359057600080fd5b6040516060810181811067ffffffffffffffff821117156135b3576135b36134cd565b60405290508082356135c48161354b565b81526135d260208401613559565b60208201526135e360408401613559565b60408201525092915050565b600080600060e0848603121561360457600080fd5b833561360f8161312c565b925061361e856020860161357e565b915061362d856080860161357e565b90509250925092565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261366b57600080fd5b83018035915067ffffffffffffffff82111561368657600080fd5b60200191503681900382131561320c57600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b67ffffffffffffffff84168152604060208201526000612a4060408301848661369b565b60006020828403121561371a57600080fd5b8151611cb98161312c565b8183823760009101908152919050565b60006020828403121561374757600080fd5b8151611cb98161354b565b67ffffffffffffffff8516815260606020820152600061377660608301858761369b565b905082604083015295945050505050565b600067ffffffffffffffff8211156137a1576137a16134cd565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600080604083850312156137e057600080fd5b82519150602083015167ffffffffffffffff8111156137fe57600080fd5b8301601f8101851361380f57600080fd5b805161382261381d82613787565b6134fc565b81815286602083850101111561383757600080fd5b6138488260208301602086016132d7565b8093505050509250929050565b600082516138678184602087016132d7565b9190910192915050565b60006020828403121561388357600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600181811c908216806138cd57607f821691505b602082108103613906577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b6020815260006108f760208301848661369b565b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee183360301811261386757600080fd5b600082601f83011261396557600080fd5b813561397361381d82613787565b81815284602083860101111561398857600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126139b657600080fd5b8135602067ffffffffffffffff808311156139d3576139d36134cd565b8260051b6139e28382016134fc565b93845285810183019383810190888611156139fc57600080fd5b84880192505b85831015613a3857823584811115613a1a5760008081fd5b613a288a87838c0101613954565b8352509184019190840190613a02565b98975050505050505050565b60006101208236031215613a5757600080fd5b60405160a0810167ffffffffffffffff8282108183111715613a7b57613a7b6134cd565b8160405284359150613a8c8261312c565b90825260208401359080821115613aa257600080fd5b613aae368387016139a5565b60208401526040850135915080821115613ac757600080fd5b50613ad436828601613954565b604083015250613ae7366060850161357e565b6060820152613af93660c0850161357e565b608082015292915050565b601f821115610a12576000816000526020600020601f850160051c81016020861015613b2d5750805b601f850160051c820191505b81811015613b4c57828155600101613b39565b505050505050565b815167ffffffffffffffff811115613b6e57613b6e6134cd565b613b8281613b7c84546138b9565b84613b04565b602080601f831160018114613bd55760008415613b9f5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555613b4c565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015613c2257888601518255948401946001909101908401613c03565b5085821015613c5e57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b600061010067ffffffffffffffff87168352806020840152613c92818401876132fb565b8551151560408581019190915260208701516fffffffffffffffffffffffffffffffff9081166060870152908701511660808501529150613cd09050565b8251151560a083015260208301516fffffffffffffffffffffffffffffffff90811660c084015260408401511660e0830152612a40565b67ffffffffffffffff831681526040602082015260006108f760408301846132fb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561069057610690613d2a565b67ffffffffffffffff8416815260e08101613db860208301858051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b82511515608083015260208301516fffffffffffffffffffffffffffffffff90811660a084015260408401511660c08301526108f7565b6060810161069082848051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b600060208284031215613e3d57600080fd5b8151611cb9816130b1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b808202811582820484141761069057610690613d2a565b8082018082111561069057610690613d2a565b600082613ed7577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea164736f6c6343000818000a",
}

var LombardTokenPoolABI = LombardTokenPoolMetaData.ABI

var LombardTokenPoolBin = LombardTokenPoolMetaData.Bin

func DeployLombardTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, lbtc_ common.Address, ccipRouter_ common.Address, allowlist_ []common.Address, rmnProxy_ common.Address, adapter_ common.Address, attestationEnable_ bool) (common.Address, *types.Transaction, *LombardTokenPool, error) {
	parsed, err := LombardTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LombardTokenPoolBin), backend, lbtc_, ccipRouter_, allowlist_, rmnProxy_, adapter_, attestationEnable_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LombardTokenPool{address: address, abi: *parsed, LombardTokenPoolCaller: LombardTokenPoolCaller{contract: contract}, LombardTokenPoolTransactor: LombardTokenPoolTransactor{contract: contract}, LombardTokenPoolFilterer: LombardTokenPoolFilterer{contract: contract}}, nil
}

type LombardTokenPool struct {
	address common.Address
	abi     abi.ABI
	LombardTokenPoolCaller
	LombardTokenPoolTransactor
	LombardTokenPoolFilterer
}

type LombardTokenPoolCaller struct {
	contract *bind.BoundContract
}

type LombardTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type LombardTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type LombardTokenPoolSession struct {
	Contract     *LombardTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type LombardTokenPoolCallerSession struct {
	Contract *LombardTokenPoolCaller
	CallOpts bind.CallOpts
}

type LombardTokenPoolTransactorSession struct {
	Contract     *LombardTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type LombardTokenPoolRaw struct {
	Contract *LombardTokenPool
}

type LombardTokenPoolCallerRaw struct {
	Contract *LombardTokenPoolCaller
}

type LombardTokenPoolTransactorRaw struct {
	Contract *LombardTokenPoolTransactor
}

func NewLombardTokenPool(address common.Address, backend bind.ContractBackend) (*LombardTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(LombardTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindLombardTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPool{address: address, abi: abi, LombardTokenPoolCaller: LombardTokenPoolCaller{contract: contract}, LombardTokenPoolTransactor: LombardTokenPoolTransactor{contract: contract}, LombardTokenPoolFilterer: LombardTokenPoolFilterer{contract: contract}}, nil
}

func NewLombardTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*LombardTokenPoolCaller, error) {
	contract, err := bindLombardTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolCaller{contract: contract}, nil
}

func NewLombardTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LombardTokenPoolTransactor, error) {
	contract, err := bindLombardTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolTransactor{contract: contract}, nil
}

func NewLombardTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LombardTokenPoolFilterer, error) {
	contract, err := bindLombardTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolFilterer{contract: contract}, nil
}

func bindLombardTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LombardTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_LombardTokenPool *LombardTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LombardTokenPool.Contract.LombardTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_LombardTokenPool *LombardTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.LombardTokenPoolTransactor.contract.Transfer(opts)
}

func (_LombardTokenPool *LombardTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.LombardTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_LombardTokenPool *LombardTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LombardTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_LombardTokenPool *LombardTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.contract.Transfer(opts)
}

func (_LombardTokenPool *LombardTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_LombardTokenPool *LombardTokenPoolCaller) Adapter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "adapter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) Adapter() (common.Address, error) {
	return _LombardTokenPool.Contract.Adapter(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) Adapter() (common.Address, error) {
	return _LombardTokenPool.Contract.Adapter(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _LombardTokenPool.Contract.GetAllowList(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _LombardTokenPool.Contract.GetAllowList(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _LombardTokenPool.Contract.GetAllowListEnabled(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _LombardTokenPool.Contract.GetAllowListEnabled(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LombardTokenPool.Contract.GetCurrentInboundRateLimiterState(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LombardTokenPool.Contract.GetCurrentInboundRateLimiterState(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LombardTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LombardTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _LombardTokenPool.Contract.GetRateLimitAdmin(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _LombardTokenPool.Contract.GetRateLimitAdmin(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _LombardTokenPool.Contract.GetRemotePools(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _LombardTokenPool.Contract.GetRemotePools(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _LombardTokenPool.Contract.GetRemoteToken(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _LombardTokenPool.Contract.GetRemoteToken(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _LombardTokenPool.Contract.GetRmnProxy(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _LombardTokenPool.Contract.GetRmnProxy(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetRouter() (common.Address, error) {
	return _LombardTokenPool.Contract.GetRouter(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetRouter() (common.Address, error) {
	return _LombardTokenPool.Contract.GetRouter(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _LombardTokenPool.Contract.GetSupportedChains(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _LombardTokenPool.Contract.GetSupportedChains(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetToken() (common.Address, error) {
	return _LombardTokenPool.Contract.GetToken(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _LombardTokenPool.Contract.GetToken(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _LombardTokenPool.Contract.GetTokenDecimals(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _LombardTokenPool.Contract.GetTokenDecimals(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) IsAttestationEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "isAttestationEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) IsAttestationEnabled() (bool, error) {
	return _LombardTokenPool.Contract.IsAttestationEnabled(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) IsAttestationEnabled() (bool, error) {
	return _LombardTokenPool.Contract.IsAttestationEnabled(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _LombardTokenPool.Contract.IsRemotePool(&_LombardTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _LombardTokenPool.Contract.IsRemotePool(&_LombardTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _LombardTokenPool.Contract.IsSupportedChain(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _LombardTokenPool.Contract.IsSupportedChain(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LombardTokenPool.Contract.IsSupportedToken(&_LombardTokenPool.CallOpts, token)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LombardTokenPool.Contract.IsSupportedToken(&_LombardTokenPool.CallOpts, token)
}

func (_LombardTokenPool *LombardTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) Owner() (common.Address, error) {
	return _LombardTokenPool.Contract.Owner(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) Owner() (common.Address, error) {
	return _LombardTokenPool.Contract.Owner(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LombardTokenPool.Contract.SupportsInterface(&_LombardTokenPool.CallOpts, interfaceId)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LombardTokenPool.Contract.SupportsInterface(&_LombardTokenPool.CallOpts, interfaceId)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_LombardTokenPool *LombardTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _LombardTokenPool.Contract.AcceptOwnership(&_LombardTokenPool.TransactOpts)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LombardTokenPool.Contract.AcceptOwnership(&_LombardTokenPool.TransactOpts)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.AddRemotePool(&_LombardTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.AddRemotePool(&_LombardTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_LombardTokenPool *LombardTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ApplyAllowListUpdates(&_LombardTokenPool.TransactOpts, removes, adds)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ApplyAllowListUpdates(&_LombardTokenPool.TransactOpts, removes, adds)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_LombardTokenPool *LombardTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ApplyChainUpdates(&_LombardTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ApplyChainUpdates(&_LombardTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_LombardTokenPool *LombardTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.LockOrBurn(&_LombardTokenPool.TransactOpts, lockOrBurnIn)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.LockOrBurn(&_LombardTokenPool.TransactOpts, lockOrBurnIn)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_LombardTokenPool *LombardTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ReleaseOrMint(&_LombardTokenPool.TransactOpts, releaseOrMintIn)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ReleaseOrMint(&_LombardTokenPool.TransactOpts, releaseOrMintIn)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.RemoveRemotePool(&_LombardTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.RemoveRemotePool(&_LombardTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_LombardTokenPool *LombardTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetChainRateLimiterConfig(&_LombardTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetChainRateLimiterConfig(&_LombardTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_LombardTokenPool *LombardTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetRateLimitAdmin(&_LombardTokenPool.TransactOpts, rateLimitAdmin)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetRateLimitAdmin(&_LombardTokenPool.TransactOpts, rateLimitAdmin)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "setRouter", newRouter)
}

func (_LombardTokenPool *LombardTokenPoolSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetRouter(&_LombardTokenPool.TransactOpts, newRouter)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetRouter(&_LombardTokenPool.TransactOpts, newRouter)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_LombardTokenPool *LombardTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.TransferOwnership(&_LombardTokenPool.TransactOpts, to)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.TransferOwnership(&_LombardTokenPool.TransactOpts, to)
}

type LombardTokenPoolAllowListAddIterator struct {
	Event *LombardTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolAllowListAdd)
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
		it.Event = new(LombardTokenPoolAllowListAdd)
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

func (it *LombardTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*LombardTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolAllowListAddIterator{contract: _LombardTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolAllowListAdd)
				if err := _LombardTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*LombardTokenPoolAllowListAdd, error) {
	event := new(LombardTokenPoolAllowListAdd)
	if err := _LombardTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolAllowListRemoveIterator struct {
	Event *LombardTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolAllowListRemove)
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
		it.Event = new(LombardTokenPoolAllowListRemove)
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

func (it *LombardTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*LombardTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolAllowListRemoveIterator{contract: _LombardTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolAllowListRemove)
				if err := _LombardTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*LombardTokenPoolAllowListRemove, error) {
	event := new(LombardTokenPoolAllowListRemove)
	if err := _LombardTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolBurnedIterator struct {
	Event *LombardTokenPoolBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolBurned)
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
		it.Event = new(LombardTokenPoolBurned)
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

func (it *LombardTokenPoolBurnedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*LombardTokenPoolBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolBurnedIterator{contract: _LombardTokenPool.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolBurned)
				if err := _LombardTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseBurned(log types.Log) (*LombardTokenPoolBurned, error) {
	event := new(LombardTokenPoolBurned)
	if err := _LombardTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolChainAddedIterator struct {
	Event *LombardTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolChainAdded)
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
		it.Event = new(LombardTokenPoolChainAdded)
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

func (it *LombardTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*LombardTokenPoolChainAddedIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolChainAddedIterator{contract: _LombardTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolChainAdded)
				if err := _LombardTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseChainAdded(log types.Log) (*LombardTokenPoolChainAdded, error) {
	event := new(LombardTokenPoolChainAdded)
	if err := _LombardTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolChainConfiguredIterator struct {
	Event *LombardTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolChainConfigured)
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
		it.Event = new(LombardTokenPoolChainConfigured)
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

func (it *LombardTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*LombardTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolChainConfiguredIterator{contract: _LombardTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolChainConfigured)
				if err := _LombardTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseChainConfigured(log types.Log) (*LombardTokenPoolChainConfigured, error) {
	event := new(LombardTokenPoolChainConfigured)
	if err := _LombardTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolChainRemovedIterator struct {
	Event *LombardTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolChainRemoved)
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
		it.Event = new(LombardTokenPoolChainRemoved)
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

func (it *LombardTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*LombardTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolChainRemovedIterator{contract: _LombardTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolChainRemoved)
				if err := _LombardTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseChainRemoved(log types.Log) (*LombardTokenPoolChainRemoved, error) {
	event := new(LombardTokenPoolChainRemoved)
	if err := _LombardTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolConfigChangedIterator struct {
	Event *LombardTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolConfigChanged)
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
		it.Event = new(LombardTokenPoolConfigChanged)
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

func (it *LombardTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*LombardTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolConfigChangedIterator{contract: _LombardTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolConfigChanged)
				if err := _LombardTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseConfigChanged(log types.Log) (*LombardTokenPoolConfigChanged, error) {
	event := new(LombardTokenPoolConfigChanged)
	if err := _LombardTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolLockedIterator struct {
	Event *LombardTokenPoolLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolLocked)
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
		it.Event = new(LombardTokenPoolLocked)
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

func (it *LombardTokenPoolLockedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*LombardTokenPoolLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolLockedIterator{contract: _LombardTokenPool.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolLocked)
				if err := _LombardTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseLocked(log types.Log) (*LombardTokenPoolLocked, error) {
	event := new(LombardTokenPoolLocked)
	if err := _LombardTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolMintedIterator struct {
	Event *LombardTokenPoolMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolMinted)
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
		it.Event = new(LombardTokenPoolMinted)
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

func (it *LombardTokenPoolMintedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LombardTokenPoolMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolMintedIterator{contract: _LombardTokenPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolMinted)
				if err := _LombardTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseMinted(log types.Log) (*LombardTokenPoolMinted, error) {
	event := new(LombardTokenPoolMinted)
	if err := _LombardTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolOwnershipTransferRequestedIterator struct {
	Event *LombardTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolOwnershipTransferRequested)
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
		it.Event = new(LombardTokenPoolOwnershipTransferRequested)
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

func (it *LombardTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolOwnershipTransferRequestedIterator{contract: _LombardTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolOwnershipTransferRequested)
				if err := _LombardTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*LombardTokenPoolOwnershipTransferRequested, error) {
	event := new(LombardTokenPoolOwnershipTransferRequested)
	if err := _LombardTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolOwnershipTransferredIterator struct {
	Event *LombardTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolOwnershipTransferred)
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
		it.Event = new(LombardTokenPoolOwnershipTransferred)
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

func (it *LombardTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolOwnershipTransferredIterator{contract: _LombardTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolOwnershipTransferred)
				if err := _LombardTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*LombardTokenPoolOwnershipTransferred, error) {
	event := new(LombardTokenPoolOwnershipTransferred)
	if err := _LombardTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolRateLimitAdminSetIterator struct {
	Event *LombardTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolRateLimitAdminSet)
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
		it.Event = new(LombardTokenPoolRateLimitAdminSet)
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

func (it *LombardTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*LombardTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolRateLimitAdminSetIterator{contract: _LombardTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolRateLimitAdminSet)
				if err := _LombardTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*LombardTokenPoolRateLimitAdminSet, error) {
	event := new(LombardTokenPoolRateLimitAdminSet)
	if err := _LombardTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolReleasedIterator struct {
	Event *LombardTokenPoolReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolReleased)
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
		it.Event = new(LombardTokenPoolReleased)
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

func (it *LombardTokenPoolReleasedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LombardTokenPoolReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolReleasedIterator{contract: _LombardTokenPool.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolReleased)
				if err := _LombardTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseReleased(log types.Log) (*LombardTokenPoolReleased, error) {
	event := new(LombardTokenPoolReleased)
	if err := _LombardTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolRemotePoolAddedIterator struct {
	Event *LombardTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolRemotePoolAdded)
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
		it.Event = new(LombardTokenPoolRemotePoolAdded)
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

func (it *LombardTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolRemotePoolAddedIterator{contract: _LombardTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolRemotePoolAdded)
				if err := _LombardTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*LombardTokenPoolRemotePoolAdded, error) {
	event := new(LombardTokenPoolRemotePoolAdded)
	if err := _LombardTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolRemotePoolRemovedIterator struct {
	Event *LombardTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolRemotePoolRemoved)
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
		it.Event = new(LombardTokenPoolRemotePoolRemoved)
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

func (it *LombardTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolRemotePoolRemovedIterator{contract: _LombardTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolRemotePoolRemoved)
				if err := _LombardTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*LombardTokenPoolRemotePoolRemoved, error) {
	event := new(LombardTokenPoolRemotePoolRemoved)
	if err := _LombardTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolRouterUpdatedIterator struct {
	Event *LombardTokenPoolRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolRouterUpdated)
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
		it.Event = new(LombardTokenPoolRouterUpdated)
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

func (it *LombardTokenPoolRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*LombardTokenPoolRouterUpdatedIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolRouterUpdatedIterator{contract: _LombardTokenPool.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolRouterUpdated)
				if err := _LombardTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseRouterUpdated(log types.Log) (*LombardTokenPoolRouterUpdated, error) {
	event := new(LombardTokenPoolRouterUpdated)
	if err := _LombardTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolTokensConsumedIterator struct {
	Event *LombardTokenPoolTokensConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolTokensConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolTokensConsumed)
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
		it.Event = new(LombardTokenPoolTokensConsumed)
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

func (it *LombardTokenPoolTokensConsumedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolTokensConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolTokensConsumed struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterTokensConsumed(opts *bind.FilterOpts) (*LombardTokenPoolTokensConsumedIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolTokensConsumedIterator{contract: _LombardTokenPool.contract, event: "TokensConsumed", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolTokensConsumed) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolTokensConsumed)
				if err := _LombardTokenPool.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
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

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseTokensConsumed(log types.Log) (*LombardTokenPoolTokensConsumed, error) {
	event := new(LombardTokenPoolTokensConsumed)
	if err := _LombardTokenPool.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_LombardTokenPool *LombardTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _LombardTokenPool.abi.Events["AllowListAdd"].ID:
		return _LombardTokenPool.ParseAllowListAdd(log)
	case _LombardTokenPool.abi.Events["AllowListRemove"].ID:
		return _LombardTokenPool.ParseAllowListRemove(log)
	case _LombardTokenPool.abi.Events["Burned"].ID:
		return _LombardTokenPool.ParseBurned(log)
	case _LombardTokenPool.abi.Events["ChainAdded"].ID:
		return _LombardTokenPool.ParseChainAdded(log)
	case _LombardTokenPool.abi.Events["ChainConfigured"].ID:
		return _LombardTokenPool.ParseChainConfigured(log)
	case _LombardTokenPool.abi.Events["ChainRemoved"].ID:
		return _LombardTokenPool.ParseChainRemoved(log)
	case _LombardTokenPool.abi.Events["ConfigChanged"].ID:
		return _LombardTokenPool.ParseConfigChanged(log)
	case _LombardTokenPool.abi.Events["Locked"].ID:
		return _LombardTokenPool.ParseLocked(log)
	case _LombardTokenPool.abi.Events["Minted"].ID:
		return _LombardTokenPool.ParseMinted(log)
	case _LombardTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _LombardTokenPool.ParseOwnershipTransferRequested(log)
	case _LombardTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _LombardTokenPool.ParseOwnershipTransferred(log)
	case _LombardTokenPool.abi.Events["RateLimitAdminSet"].ID:
		return _LombardTokenPool.ParseRateLimitAdminSet(log)
	case _LombardTokenPool.abi.Events["Released"].ID:
		return _LombardTokenPool.ParseReleased(log)
	case _LombardTokenPool.abi.Events["RemotePoolAdded"].ID:
		return _LombardTokenPool.ParseRemotePoolAdded(log)
	case _LombardTokenPool.abi.Events["RemotePoolRemoved"].ID:
		return _LombardTokenPool.ParseRemotePoolRemoved(log)
	case _LombardTokenPool.abi.Events["RouterUpdated"].ID:
		return _LombardTokenPool.ParseRouterUpdated(log)
	case _LombardTokenPool.abi.Events["TokensConsumed"].ID:
		return _LombardTokenPool.ParseTokensConsumed(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (LombardTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (LombardTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (LombardTokenPoolBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
}

func (LombardTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (LombardTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (LombardTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (LombardTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (LombardTokenPoolLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (LombardTokenPoolMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (LombardTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (LombardTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (LombardTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (LombardTokenPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (LombardTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (LombardTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (LombardTokenPoolRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (LombardTokenPoolTokensConsumed) Topic() common.Hash {
	return common.HexToHash("0x1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a")
}

func (_LombardTokenPool *LombardTokenPool) Address() common.Address {
	return _LombardTokenPool.address
}

type LombardTokenPoolInterface interface {
	Adapter(opts *bind.CallOpts) (common.Address, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	IsAttestationEnabled(opts *bind.CallOpts) (bool, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*LombardTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*LombardTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*LombardTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*LombardTokenPoolAllowListRemove, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*LombardTokenPoolBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*LombardTokenPoolBurned, error)

	FilterChainAdded(opts *bind.FilterOpts) (*LombardTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*LombardTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*LombardTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*LombardTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*LombardTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*LombardTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*LombardTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*LombardTokenPoolConfigChanged, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*LombardTokenPoolLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*LombardTokenPoolLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LombardTokenPoolMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*LombardTokenPoolMinted, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*LombardTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*LombardTokenPoolOwnershipTransferred, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*LombardTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*LombardTokenPoolRateLimitAdminSet, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LombardTokenPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*LombardTokenPoolReleased, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*LombardTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*LombardTokenPoolRemotePoolRemoved, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*LombardTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*LombardTokenPoolRouterUpdated, error)

	FilterTokensConsumed(opts *bind.FilterOpts) (*LombardTokenPoolTokensConsumedIterator, error)

	WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolTokensConsumed) (event.Subscription, error)

	ParseTokensConsumed(log types.Log) (*LombardTokenPoolTokensConsumed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
