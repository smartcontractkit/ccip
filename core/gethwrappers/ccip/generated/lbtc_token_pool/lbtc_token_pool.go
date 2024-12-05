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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lbtc_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ccipRouter_\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowlist_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"rmnProxy_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"attestationEnable_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowListNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotARampOnRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotTransferToSelf\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"DisabledNonZeroRateLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"expected\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"actual\",\"type\":\"uint8\"}],\"name\":\"InvalidDecimalArgs\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"InvalidRateLimitRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolData\",\"type\":\"bytes\"}],\"name\":\"InvalidRemoteChainDecimals\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidRemotePoolForChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidSourcePoolAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeProposedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"NonExistentChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"remoteDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"localDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"remoteAmount\",\"type\":\"uint256\"}],\"name\":\"OverflowDetected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateLimitMustBeDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remoteToken\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rateLimitAdmin\",\"type\":\"address\"}],\"name\":\"RateLimitAdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"RemotePoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"RemotePoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRouter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensConsumed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adapter\",\"outputs\":[{\"internalType\":\"contractCLAdapter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"addRemotePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"remoteTokenAddress\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"internalType\":\"structTokenPool.ChainUpdate[]\",\"name\":\"chainsToAdd\",\"type\":\"tuple[]\"}],\"name\":\"applyChainUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentInboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentOutboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRateLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRemotePools\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRemoteToken\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRmnProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedChains\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAttestationEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"isRemotePool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"isSupportedChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isSupportedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"internalType\":\"structPool.LockOrBurnInV1\",\"name\":\"lockOrBurnIn\",\"type\":\"tuple\"}],\"name\":\"lockOrBurn\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destPoolData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.LockOrBurnOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"originalSender\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainTokenData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.ReleaseOrMintInV1\",\"name\":\"releaseOrMintIn\",\"type\":\"tuple\"}],\"name\":\"releaseOrMint\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"destinationAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"removeRemotePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCLAdapter\",\"name\":\"adapter_\",\"type\":\"address\"}],\"name\":\"setAdapter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundConfig\",\"type\":\"tuple\"}],\"name\":\"setChainRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rateLimitAdmin\",\"type\":\"address\"}],\"name\":\"setRateLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620046b3380380620046b38339810160408190526200003591620005bf565b846008848487336000816200005d57604051639b15e16f60e01b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b03848116919091179091558116156200009057620000908162000209565b50506001600160a01b0385161580620000b057506001600160a01b038116155b80620000c357506001600160a01b038216155b15620000e2576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b03808616608081905290831660c0526040805163313ce56760e01b8152905163313ce567916004808201926020929091908290030181865afa92505050801562000152575060408051601f3d908101601f191682019092526200014f91810190620006e5565b60015b1562000192578060ff168560ff161462000190576040516332ad3e0760e11b815260ff80871660048301528216602482015260440160405180910390fd5b505b60ff841660a052600480546001600160a01b0319166001600160a01b038316179055825115801560e052620001dc57604080516000815260208101909152620001dc908462000283565b5050600a8054941515600160a01b0260ff60a01b1990951694909417909355506200075f95505050505050565b336001600160a01b038216036200023357604051636d6c4ee560e11b815260040160405180910390fd5b600080546001600160a01b0319166001600160a01b03838116918217835560015460405192939116917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60e051620002a4576040516335f4a7b360e01b815260040160405180910390fd5b60005b82518110156200032f576000838281518110620002c857620002c862000711565b60209081029190910101519050620002e2600282620003e0565b1562000325576040516001600160a01b03821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b50600101620002a7565b5060005b8151811015620003db57600082828151811062000354576200035462000711565b6020026020010151905060006001600160a01b0316816001600160a01b031603620003805750620003d2565b6200038d60028262000400565b15620003d0576040516001600160a01b03821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b60010162000333565b505050565b6000620003f7836001600160a01b03841662000417565b90505b92915050565b6000620003f7836001600160a01b0384166200051b565b60008181526001830160205260408120548015620005105760006200043e60018362000727565b8554909150600090620004549060019062000727565b9050808214620004c057600086600001828154811062000478576200047862000711565b90600052602060002001549050808760000184815481106200049e576200049e62000711565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080620004d457620004d462000749565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050620003fa565b6000915050620003fa565b60008181526001830160205260408120546200056457508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155620003fa565b506000620003fa565b6001600160a01b03811681146200058357600080fd5b50565b805162000593816200056d565b919050565b634e487b7160e01b600052604160045260246000fd5b805180151581146200059357600080fd5b600080600080600060a08688031215620005d857600080fd5b8551620005e5816200056d565b80955050602080870151620005fa816200056d565b60408801519095506001600160401b03808211156200061857600080fd5b818901915089601f8301126200062d57600080fd5b81518181111562000642576200064262000598565b8060051b604051601f19603f830116810181811085821117156200066a576200066a62000598565b60405291825284820192508381018501918c8311156200068957600080fd5b938501935b82851015620006b257620006a28562000586565b845293850193928501926200068e565b809850505050505050620006c96060870162000586565b9150620006d960808701620005ae565b90509295509295909350565b600060208284031215620006f857600080fd5b815160ff811681146200070a57600080fd5b9392505050565b634e487b7160e01b600052603260045260246000fd5b81810381811115620003fa57634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60805160a05160c05160e051613ee2620007d1600039600081816105c701528181611d3e01526126ed0152600081816105a101528181611b33015261202a015260006102d10152600081816102590152818161028d01528181610c150152818161268301526128d80152613ee26000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c80639a4575b91161010f578063c0d78655116100a2578063dc0bd97111610071578063dc0bd9711461059f578063e0351e13146105c5578063e8a1da17146105eb578063f2fde38b146105fe57600080fd5b8063c0d7865514610551578063c4bffe2b14610564578063c75eea9c14610579578063cf7401f31461058c57600080fd5b8063acfecf91116100de578063acfecf9114610491578063af58d59f146104a4578063b0f479a114610513578063b79465801461053157600080fd5b80639a4575b9146103e7578063a42a7b8b14610407578063a7cd63b714610427578063ab1da79c1461043c57600080fd5b806354c8a4f31161018757806379ba50971161015657806379ba50971461039b5780637d54534e146103a35780638926f54f146103b65780638da5cb5b146103c957600080fd5b806354c8a4f31461033057806355b961561461034557806362ddd3c41461036a5780636d3d1a581461037d57600080fd5b8063240028e8116101c3578063240028e81461027d57806324f65ee7146102ca57806339077537146102fb5780634c5ef0ed1461031d57600080fd5b806301ffc9a7146101ea57806303eadcfc1461021257806321df0da714610257575b600080fd5b6101fd6101f8366004613098565b610611565b60405190151581526020015b60405180910390f35b600a546102329073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610209565b7f0000000000000000000000000000000000000000000000000000000000000000610232565b6101fd61028b3660046130fc565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff90811691161490565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610209565b61030e610309366004613119565b6106f6565b60405190518152602001610209565b6101fd61032b366004613172565b6108de565b61034361033e366004613241565b610928565b005b600a546101fd9074010000000000000000000000000000000000000000900460ff1681565b610343610378366004613172565b6109a3565b60095473ffffffffffffffffffffffffffffffffffffffff16610232565b610343610a40565b6103436103b13660046130fc565b610b0e565b6101fd6103c43660046132ad565b610b8f565b60015473ffffffffffffffffffffffffffffffffffffffff16610232565b6103fa6103f53660046132c8565b610ba6565b6040516102099190613371565b61041a6104153660046132ad565b610e82565b60405161020991906133c8565b61042f610fed565b604051610209919061344a565b61034361044a3660046130fc565b600a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b61034361049f366004613172565b610ffe565b6104b76104b23660046132ad565b611116565b604051610209919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b60045473ffffffffffffffffffffffffffffffffffffffff16610232565b61054461053f3660046132ad565b6111eb565b60405161020991906134a4565b61034361055f3660046130fc565b61129b565b61056c611376565b60405161020991906134b7565b6104b76105873660046132ad565b61142e565b61034361059a36600461363f565b611500565b7f0000000000000000000000000000000000000000000000000000000000000000610232565b7f00000000000000000000000000000000000000000000000000000000000000006101fd565b6103436105f9366004613241565b611584565b61034361060c3660046130fc565b611a96565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167faff2afbf0000000000000000000000000000000000000000000000000000000014806106a457507fffffffff0000000000000000000000000000000000000000000000000000000082167f0e64dd2900000000000000000000000000000000000000000000000000000000145b806106f057507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b60408051602081019091526000815261070e82611aaa565b600a5474010000000000000000000000000000000000000000900460ff16156107c157600a5473ffffffffffffffffffffffffffffffffffffffff16635391a40561075f60408501602086016132ad565b61076c60e0860186613684565b6040518463ffffffff1660e01b815260040161078a93929190613732565b600060405180830381600087803b1580156107a457600080fd5b505af11580156107b8573d6000803e3d6000fd5b5050505061084d565b600a5473ffffffffffffffffffffffffffffffffffffffff16630c373d746107ef60408501602086016132ad565b6107fc60c0860186613684565b6040518463ffffffff1660e01b815260040161081a93929190613732565b600060405180830381600087803b15801561083457600080fd5b505af1158015610848573d6000803e3d6000fd5b505050505b61085d60608301604084016130fc565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f084606001356040516108bf91815260200190565b60405180910390a3506040805160208101909152606090910135815290565b600061092083836040516108f3929190613756565b604080519182900390912067ffffffffffffffff8716600090815260076020529190912060050190611cce565b949350505050565b610930611ce9565b61099d84848080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808802828101820190935287825290935087925086918291850190849080828437600092019190915250611d3c92505050565b50505050565b6109ab611ce9565b6109b483610b8f565b6109fb576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024015b60405180910390fd5b610a3b8383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611ef292505050565b505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610a91576040517f02b543c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560008054909116815560405173ffffffffffffffffffffffffffffffffffffffff909216929183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610b16611ce9565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d091749060200160405180910390a150565b60006106f0600567ffffffffffffffff8416611cce565b6040805180820190915260608082526020820152610bc382611fec565b600a546040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152606084013560248201527f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb906044016020604051808303816000875af1158015610c60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c849190613766565b50600a54600090819073ffffffffffffffffffffffffffffffffffffffff1663550e7ab2610cb860408701602088016132ad565b610cc28780613684565b88606001356040518563ffffffff1660e01b8152600401610ce69493929190613783565b6000604051808303816000875af1158015610d05573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610d4b91908101906137fe565b9092509050610d6060608501604086016130fc565b73ffffffffffffffffffffffffffffffffffffffff167f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df783604051610da791815260200190565b60405180910390a2600a5460609074010000000000000000000000000000000000000000900460ff1615610e4d57600282604051610de59190613886565b602060405180830381855afa158015610e02573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190610e2591906138a2565b604051602001610e3791815260200190565b6040516020818303038152906040529050610e50565b50805b6040518060400160405280610e7187602001602081019061053f91906132ad565b815260200191909152949350505050565b67ffffffffffffffff8116600090815260076020526040812060609190610eab90600501612178565b90506000815167ffffffffffffffff811115610ec957610ec96134f9565b604051908082528060200260200182016040528015610efc57816020015b6060815260200190600190039081610ee75790505b50905060005b8251811015610fe55760086000848381518110610f2157610f216138bb565b602002602001015181526020019081526020016000208054610f42906138ea565b80601f0160208091040260200160405190810160405280929190818152602001828054610f6e906138ea565b8015610fbb5780601f10610f9057610100808354040283529160200191610fbb565b820191906000526020600020905b815481529060010190602001808311610f9e57829003601f168201915b5050505050828281518110610fd257610fd26138bb565b6020908102919091010152600101610f02565b509392505050565b6060610ff96002612178565b905090565b611006611ce9565b61100f83610b8f565b611051576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024016109f2565b6110918282604051611064929190613756565b604080519182900390912067ffffffffffffffff8616600090815260076020529190912060050190612185565b6110cd578282826040517f74f23c7c0000000000000000000000000000000000000000000000000000000081526004016109f293929190613732565b8267ffffffffffffffff167f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76838360405161110992919061393d565b60405180910390a2505050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845260028201546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff1615159482019490945260039091015480841660608301529190910490911660808201526106f090612191565b67ffffffffffffffff81166000908152600760205260409020600401805460609190611216906138ea565b80601f0160208091040260200160405190810160405280929190818152602001828054611242906138ea565b801561128f5780601f106112645761010080835404028352916020019161128f565b820191906000526020600020905b81548152906001019060200180831161127257829003601f168201915b50505050509050919050565b6112a3611ce9565b73ffffffffffffffffffffffffffffffffffffffff81166112f0576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684910160405180910390a15050565b606060006113846005612178565b90506000815167ffffffffffffffff8111156113a2576113a26134f9565b6040519080825280602002602001820160405280156113cb578160200160208202803683370190505b50905060005b8251811015611427578281815181106113ec576113ec6138bb565b6020026020010151828281518110611406576114066138bb565b67ffffffffffffffff909216602092830291909101909101526001016113d1565b5092915050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845281546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff1615159482019490945260019091015480841660608301529190910490911660808201526106f090612191565b60095473ffffffffffffffffffffffffffffffffffffffff163314801590611540575060015473ffffffffffffffffffffffffffffffffffffffff163314155b15611579576040517f8e4a23d60000000000000000000000000000000000000000000000000000000081523360048201526024016109f2565b610a3b838383612243565b61158c611ce9565b60005b838110156117795760008585838181106115ab576115ab6138bb565b90506020020160208101906115c091906132ad565b90506115d7600567ffffffffffffffff8316612185565b611619576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016109f2565b67ffffffffffffffff8116600090815260076020526040812061163e90600501612178565b905060005b81518110156116aa576116a1828281518110611661576116616138bb565b6020026020010151600760008667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060050161218590919063ffffffff16565b50600101611643565b5067ffffffffffffffff8216600090815260076020526040812080547fffffffffffffffffffffff00000000000000000000000000000000000000000090811682556001820183905560028201805490911690556003810182905590611713600483018261302b565b60058201600081816117258282613065565b505060405167ffffffffffffffff871681527f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d85991694506020019250611767915050565b60405180910390a1505060010161158f565b5060005b81811015611a8f576000838383818110611799576117996138bb565b90506020028101906117ab9190613951565b6117b4906139d6565b90506117c58160600151600061232d565b6117d48160800151600061232d565b806040015151600003611813576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805161182b9060059067ffffffffffffffff1661246a565b6118705780516040517f1d5ad3c500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016109f2565b805167ffffffffffffffff16600090815260076020908152604091829020825160a08082018552606080870180518601516fffffffffffffffffffffffffffffffff90811680865263ffffffff42168689018190528351511515878b0181905284518a0151841686890181905294518b0151841660809889018190528954740100000000000000000000000000000000000000009283027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff7001000000000000000000000000000000008087027fffffffffffffffffffffffff000000000000000000000000000000000000000094851690981788178216929092178d5592810290971760018c01558c519889018d52898e0180518d01518716808b528a8e019590955280515115158a8f018190528151909d01518716988a01899052518d0151909516979098018790526002890180549a9091029990931617179094169590951790925590920290911760038201559082015160048201906119f39082613b4d565b5060005b826020015151811015611a3757611a2f836000015184602001518381518110611a2257611a226138bb565b6020026020010151611ef2565b6001016119f7565b507f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c28260000151836040015184606001518560800151604051611a7d9493929190613c67565b60405180910390a1505060010161177d565b5050505050565b611a9e611ce9565b611aa781612476565b50565b611abd61028b60a08301608084016130fc565b611b1c57611ad160a08201608083016130fc565b6040517f961c9a4f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016109f2565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016632cbc26bb611b6860408401602085016132ad565b60405160e083901b7fffffffff0000000000000000000000000000000000000000000000000000000016815260809190911b77ffffffffffffffff00000000000000000000000000000000166004820152602401602060405180830381865afa158015611bd9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bfd9190613766565b15611c34576040517f53ad11d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c4c611c4760408301602084016132ad565b61253a565b611c6c611c5f60408301602084016132ad565b61032b60a0840184613684565b611cb157611c7d60a0820182613684565b6040517f24eb47e50000000000000000000000000000000000000000000000000000000081526004016109f292919061393d565b611aa7611cc460408301602084016132ad565b8260600135612660565b600081815260018301602052604081205415155b9392505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611d3a576040517f2b5c74de00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b7f0000000000000000000000000000000000000000000000000000000000000000611d93576040517f35f4a7b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8251811015611e29576000838281518110611db357611db36138bb565b60200260200101519050611dd18160026126a790919063ffffffff16565b15611e205760405173ffffffffffffffffffffffffffffffffffffffff821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b50600101611d96565b5060005b8151811015610a3b576000828281518110611e4a57611e4a6138bb565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611e8e5750611eea565b611e996002826126c9565b15611ee85760405173ffffffffffffffffffffffffffffffffffffffff821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b600101611e2d565b8051600003611f2d576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160208083019190912067ffffffffffffffff8416600090815260079092526040909120611f5f906005018261246a565b611f995782826040517f393b8ad20000000000000000000000000000000000000000000000000000000081526004016109f2929190613d00565b6000818152600860205260409020611fb18382613b4d565b508267ffffffffffffffff167f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea8360405161110991906134a4565b611fff61028b60a08301608084016130fc565b61201357611ad160a08201608083016130fc565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016632cbc26bb61205f60408401602085016132ad565b60405160e083901b7fffffffff0000000000000000000000000000000000000000000000000000000016815260809190911b77ffffffffffffffff00000000000000000000000000000000166004820152602401602060405180830381865afa1580156120d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120f49190613766565b1561212b576040517f53ad11d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61214361213e60608301604084016130fc565b6126eb565b61215b61215660408301602084016132ad565b61276a565b611aa761216e60408301602084016132ad565b82606001356128b8565b60606000611ce2836128fc565b6000611ce28383612957565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261221f82606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff16426122039190613d52565b85608001516fffffffffffffffffffffffffffffffff16612a4a565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b61224c83610b8f565b61228e576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024016109f2565b61229982600061232d565b67ffffffffffffffff831660009081526007602052604090206122bc9083612a72565b6122c781600061232d565b67ffffffffffffffff831660009081526007602052604090206122ed9060020182612a72565b7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b83838360405161232093929190613d65565b60405180910390a1505050565b8151156123f85781602001516fffffffffffffffffffffffffffffffff1682604001516fffffffffffffffffffffffffffffffff16101580612383575060408201516fffffffffffffffffffffffffffffffff16155b156123bc57816040517f8020d1240000000000000000000000000000000000000000000000000000000081526004016109f29190613de8565b80156123f4576040517f433fc33d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b60408201516fffffffffffffffffffffffffffffffff16151580612431575060208201516fffffffffffffffffffffffffffffffff1615155b156123f457816040517fd68af9cc0000000000000000000000000000000000000000000000000000000081526004016109f29190613de8565b6000611ce28383612c14565b3373ffffffffffffffffffffffffffffffffffffffff8216036124c5576040517fdad89dca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217835560015460405192939116917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b61254381610b8f565b612585576040517fa9902c7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016109f2565b600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84169281019290925233602483015273ffffffffffffffffffffffffffffffffffffffff16906383826b2b90604401602060405180830381865afa158015612604573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126289190613766565b611aa7576040517f728fe07b0000000000000000000000000000000000000000000000000000000081523360048201526024016109f2565b67ffffffffffffffff821660009081526007602052604090206123f490600201827f0000000000000000000000000000000000000000000000000000000000000000612c63565b6000611ce28373ffffffffffffffffffffffffffffffffffffffff8416612957565b6000611ce28373ffffffffffffffffffffffffffffffffffffffff8416612c14565b7f000000000000000000000000000000000000000000000000000000000000000015611aa75761271c600282612fe6565b611aa7576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016109f2565b61277381610b8f565b6127b5576040517fa9902c7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016109f2565b600480546040517fa8d87a3b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84169281019290925273ffffffffffffffffffffffffffffffffffffffff169063a8d87a3b90602401602060405180830381865afa15801561282e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128529190613e24565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611aa7576040517f728fe07b0000000000000000000000000000000000000000000000000000000081523360048201526024016109f2565b67ffffffffffffffff821660009081526007602052604090206123f490827f0000000000000000000000000000000000000000000000000000000000000000612c63565b60608160000180548060200260200160405190810160405280929190818152602001828054801561128f57602002820191906000526020600020905b8154815260200190600101908083116129385750505050509050919050565b60008181526001830160205260408120548015612a4057600061297b600183613d52565b855490915060009061298f90600190613d52565b90508082146129f45760008660000182815481106129af576129af6138bb565b90600052602060002001549050808760000184815481106129d2576129d26138bb565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080612a0557612a05613e41565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506106f0565b60009150506106f0565b6000612a6985612a5a8486613e70565b612a649087613e87565b613015565b95945050505050565b8154600090612a9b90700100000000000000000000000000000000900463ffffffff1642613d52565b90508015612b3d5760018301548354612ae3916fffffffffffffffffffffffffffffffff80821692811691859170010000000000000000000000000000000090910416612a4a565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b60208201518354612b63916fffffffffffffffffffffffffffffffff9081169116613015565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1990612320908490613de8565b6000818152600183016020526040812054612c5b575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556106f0565b5060006106f0565b825474010000000000000000000000000000000000000000900460ff161580612c8a575081155b15612c9457505050565b825460018401546fffffffffffffffffffffffffffffffff80831692911690600090612cda90700100000000000000000000000000000000900463ffffffff1642613d52565b90508015612d9a5781831115612d1c576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001860154612d569083908590849070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16612a4a565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff167001000000000000000000000000000000004263ffffffff160217875592505b84821015612e515773ffffffffffffffffffffffffffffffffffffffff8416612df9576040517ff94ebcd100000000000000000000000000000000000000000000000000000000815260048101839052602481018690526044016109f2565b6040517f1a76572a000000000000000000000000000000000000000000000000000000008152600481018390526024810186905273ffffffffffffffffffffffffffffffffffffffff851660448201526064016109f2565b84831015612f645760018681015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16906000908290612e959082613d52565b612e9f878a613d52565b612ea99190613e87565b612eb39190613e9a565b905073ffffffffffffffffffffffffffffffffffffffff8616612f0c576040517f15279c0800000000000000000000000000000000000000000000000000000000815260048101829052602481018690526044016109f2565b6040517fd0c8d23a000000000000000000000000000000000000000000000000000000008152600481018290526024810186905273ffffffffffffffffffffffffffffffffffffffff871660448201526064016109f2565b612f6e8584613d52565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515611ce2565b60008183106130245781611ce2565b5090919050565b508054613037906138ea565b6000825580601f10613047575050565b601f016020900490600052602060002090810190611aa7919061307f565b5080546000825590600052602060002090810190611aa791905b5b808211156130945760008155600101613080565b5090565b6000602082840312156130aa57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611ce257600080fd5b73ffffffffffffffffffffffffffffffffffffffff81168114611aa757600080fd5b60006020828403121561310e57600080fd5b8135611ce2816130da565b60006020828403121561312b57600080fd5b813567ffffffffffffffff81111561314257600080fd5b82016101008185031215611ce257600080fd5b803567ffffffffffffffff8116811461316d57600080fd5b919050565b60008060006040848603121561318757600080fd5b61319084613155565b9250602084013567ffffffffffffffff808211156131ad57600080fd5b818601915086601f8301126131c157600080fd5b8135818111156131d057600080fd5b8760208285010111156131e257600080fd5b6020830194508093505050509250925092565b60008083601f84011261320757600080fd5b50813567ffffffffffffffff81111561321f57600080fd5b6020830191508360208260051b850101111561323a57600080fd5b9250929050565b6000806000806040858703121561325757600080fd5b843567ffffffffffffffff8082111561326f57600080fd5b61327b888389016131f5565b9096509450602087013591508082111561329457600080fd5b506132a1878288016131f5565b95989497509550505050565b6000602082840312156132bf57600080fd5b611ce282613155565b6000602082840312156132da57600080fd5b813567ffffffffffffffff8111156132f157600080fd5b820160a08185031215611ce257600080fd5b60005b8381101561331e578181015183820152602001613306565b50506000910152565b6000815180845261333f816020860160208601613303565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600082516040602084015261338d6060840182613327565b905060208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848303016040850152612a698282613327565b600060208083016020845280855180835260408601915060408160051b87010192506020870160005b8281101561343d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc088860301845261342b858351613327565b945092850192908501906001016133f1565b5092979650505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561349857835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101613466565b50909695505050505050565b602081526000611ce26020830184613327565b6020808252825182820181905260009190848201906040850190845b8181101561349857835167ffffffffffffffff16835292840192918401916001016134d3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff8111828210171561354b5761354b6134f9565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613598576135986134f9565b604052919050565b8015158114611aa757600080fd5b80356fffffffffffffffffffffffffffffffff8116811461316d57600080fd5b6000606082840312156135e057600080fd5b6040516060810181811067ffffffffffffffff82111715613603576136036134f9565b6040529050808235613614816135a0565b8152613622602084016135ae565b6020820152613633604084016135ae565b60408201525092915050565b600080600060e0848603121561365457600080fd5b61365d84613155565b925061366c85602086016135ce565b915061367b85608086016135ce565b90509250925092565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126136b957600080fd5b83018035915067ffffffffffffffff8211156136d457600080fd5b60200191503681900382131561323a57600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b67ffffffffffffffff84168152604060208201526000612a696040830184866136e9565b8183823760009101908152919050565b60006020828403121561377857600080fd5b8151611ce2816135a0565b67ffffffffffffffff851681526060602082015260006137a76060830185876136e9565b905082604083015295945050505050565b600067ffffffffffffffff8211156137d2576137d26134f9565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6000806040838503121561381157600080fd5b82519150602083015167ffffffffffffffff81111561382f57600080fd5b8301601f8101851361384057600080fd5b805161385361384e826137b8565b613551565b81815286602083850101111561386857600080fd5b613879826020830160208601613303565b8093505050509250929050565b60008251613898818460208701613303565b9190910192915050565b6000602082840312156138b457600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600181811c908216806138fe57607f821691505b602082108103613937577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b6020815260006109206020830184866136e9565b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee183360301811261389857600080fd5b600082601f83011261399657600080fd5b81356139a461384e826137b8565b8181528460208386010111156139b957600080fd5b816020850160208301376000918101602001919091529392505050565b600061012082360312156139e957600080fd5b6139f1613528565b6139fa83613155565b815260208084013567ffffffffffffffff80821115613a1857600080fd5b9085019036601f830112613a2b57600080fd5b813581811115613a3d57613a3d6134f9565b8060051b613a4c858201613551565b9182528381018501918581019036841115613a6657600080fd5b86860192505b83831015613aa257823585811115613a845760008081fd5b613a923689838a0101613985565b8352509186019190860190613a6c565b8087890152505050506040860135925080831115613abf57600080fd5b5050613acd36828601613985565b604083015250613ae036606085016135ce565b6060820152613af23660c085016135ce565b608082015292915050565b601f821115610a3b576000816000526020600020601f850160051c81016020861015613b265750805b601f850160051c820191505b81811015613b4557828155600101613b32565b505050505050565b815167ffffffffffffffff811115613b6757613b676134f9565b613b7b81613b7584546138ea565b84613afd565b602080601f831160018114613bce5760008415613b985750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555613b45565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015613c1b57888601518255948401946001909101908401613bfc565b5085821015613c5757878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b600061010067ffffffffffffffff87168352806020840152613c8b81840187613327565b8551151560408581019190915260208701516fffffffffffffffffffffffffffffffff9081166060870152908701511660808501529150613cc99050565b8251151560a083015260208301516fffffffffffffffffffffffffffffffff90811660c084015260408401511660e0830152612a69565b67ffffffffffffffff831681526040602082015260006109206040830184613327565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156106f0576106f0613d23565b67ffffffffffffffff8416815260e08101613db160208301858051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b82511515608083015260208301516fffffffffffffffffffffffffffffffff90811660a084015260408401511660c0830152610920565b606081016106f082848051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b600060208284031215613e3657600080fd5b8151611ce2816130da565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b80820281158282048414176106f0576106f0613d23565b808201808211156106f0576106f0613d23565b600082613ed0577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea164736f6c6343000818000a",
}

var LombardTokenPoolABI = LombardTokenPoolMetaData.ABI

var LombardTokenPoolBin = LombardTokenPoolMetaData.Bin

func DeployLombardTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, lbtc_ common.Address, ccipRouter_ common.Address, allowlist_ []common.Address, rmnProxy_ common.Address, attestationEnable_ bool) (common.Address, *types.Transaction, *LombardTokenPool, error) {
	parsed, err := LombardTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LombardTokenPoolBin), backend, lbtc_, ccipRouter_, allowlist_, rmnProxy_, attestationEnable_)
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

func (_LombardTokenPool *LombardTokenPoolTransactor) SetAdapter(opts *bind.TransactOpts, adapter_ common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "setAdapter", adapter_)
}

func (_LombardTokenPool *LombardTokenPoolSession) SetAdapter(adapter_ common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetAdapter(&_LombardTokenPool.TransactOpts, adapter_)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) SetAdapter(adapter_ common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetAdapter(&_LombardTokenPool.TransactOpts, adapter_)
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

	SetAdapter(opts *bind.TransactOpts, adapter_ common.Address) (*types.Transaction, error)

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
