// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_lbtc_token_pool

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

var MockLBTCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AllowListNotEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotARampOnRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotTransferToSelf\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"DisabledNonZeroRateLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"expected\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"actual\",\"type\":\"uint8\"}],\"name\":\"InvalidDecimalArgs\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"InvalidRateLimitRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolData\",\"type\":\"bytes\"}],\"name\":\"InvalidRemoteChainDecimals\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidRemotePoolForChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidSourcePoolAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeProposedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"NonExistentChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"remoteDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"localDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"remoteAmount\",\"type\":\"uint256\"}],\"name\":\"OverflowDetected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateLimitMustBeDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remoteToken\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rateLimitAdmin\",\"type\":\"address\"}],\"name\":\"RateLimitAdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"RemotePoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"RemotePoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRouter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"addRemotePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"remoteTokenAddress\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"internalType\":\"structTokenPool.ChainUpdate[]\",\"name\":\"chainsToAdd\",\"type\":\"tuple[]\"}],\"name\":\"applyChainUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentInboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentOutboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRateLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRemotePools\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRemoteToken\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRmnProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedChains\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"isRemotePool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"isSupportedChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isSupportedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"internalType\":\"structPool.LockOrBurnInV1\",\"name\":\"lockOrBurnIn\",\"type\":\"tuple\"}],\"name\":\"lockOrBurn\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destPoolData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.LockOrBurnOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"originalSender\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainTokenData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.ReleaseOrMintInV1\",\"name\":\"releaseOrMintIn\",\"type\":\"tuple\"}],\"name\":\"releaseOrMint\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"destinationAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"removeRemotePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundConfig\",\"type\":\"tuple\"}],\"name\":\"setChainRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rateLimitAdmin\",\"type\":\"address\"}],\"name\":\"setRateLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620037f1380380620037f1833981016040819052620000359162000590565b836008848484336000816200005d57604051639b15e16f60e01b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b038481169190911790915581161562000090576200009081620001eb565b50506001600160a01b0385161580620000b057506001600160a01b038116155b80620000c357506001600160a01b038216155b15620000e2576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b03808616608081905290831660c0526040805163313ce56760e01b8152905163313ce567916004808201926020929091908290030181865afa92505050801562000152575060408051601f3d908101601f191682019092526200014f91810190620006a0565b60015b1562000192578060ff168560ff161462000190576040516332ad3e0760e11b815260ff80871660048301528216602482015260440160405180910390fd5b505b60ff841660a052600480546001600160a01b0319166001600160a01b038316179055825115801560e052620001dc57604080516000815260208101909152620001dc908462000265565b5050505050505050506200071a565b336001600160a01b038216036200021557604051636d6c4ee560e11b815260040160405180910390fd5b600080546001600160a01b0319166001600160a01b03838116918217835560015460405192939116917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60e05162000286576040516335f4a7b360e01b815260040160405180910390fd5b60005b825181101562000311576000838281518110620002aa57620002aa620006cc565b60209081029190910101519050620002c4600282620003c2565b1562000307576040516001600160a01b03821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b5060010162000289565b5060005b8151811015620003bd576000828281518110620003365762000336620006cc565b6020026020010151905060006001600160a01b0316816001600160a01b031603620003625750620003b4565b6200036f600282620003e2565b15620003b2576040516001600160a01b03821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b60010162000315565b505050565b6000620003d9836001600160a01b038416620003f9565b90505b92915050565b6000620003d9836001600160a01b038416620004fd565b60008181526001830160205260408120548015620004f257600062000420600183620006e2565b85549091506000906200043690600190620006e2565b9050808214620004a25760008660000182815481106200045a576200045a620006cc565b9060005260206000200154905080876000018481548110620004805762000480620006cc565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080620004b657620004b662000704565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050620003dc565b6000915050620003dc565b60008181526001830160205260408120546200054657508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155620003dc565b506000620003dc565b6001600160a01b03811681146200056557600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b80516200058b816200054f565b919050565b60008060008060808587031215620005a757600080fd5b8451620005b4816200054f565b602086810151919550906001600160401b0380821115620005d457600080fd5b818801915088601f830112620005e957600080fd5b815181811115620005fe57620005fe62000568565b8060051b604051601f19603f8301168101818110858211171562000626576200062662000568565b60405291825284820192508381018501918b8311156200064557600080fd5b938501935b828510156200066e576200065e856200057e565b845293850193928501926200064a565b80985050505050505062000685604086016200057e565b915062000695606086016200057e565b905092959194509250565b600060208284031215620006b357600080fd5b815160ff81168114620006c557600080fd5b9392505050565b634e487b7160e01b600052603260045260246000fd5b81810381811115620003dc57634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60805160a05160c05160e051613081620007706000396000818161054f015261194b01526000610529015260006102e00152600081816102470152818161029c015281816106a40152610bd801526130816000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80639a4575b911610104578063c0d78655116100a2578063dc0bd97111610071578063dc0bd97114610527578063e0351e131461054d578063e8a1da1714610573578063f2fde38b1461058657600080fd5b8063c0d78655146104d9578063c4bffe2b146104ec578063c75eea9c14610501578063cf7401f31461051457600080fd5b8063acfecf91116100de578063acfecf9114610426578063af58d59f14610439578063b0f479a1146104a8578063b7946580146104c657600080fd5b80639a4575b9146103d1578063a42a7b8b146103f1578063a7cd63b71461041157600080fd5b806354c8a4f31161017157806379ba50971161014b57806379ba5097146103855780637d54534e1461038d5780638926f54f146103a05780638da5cb5b146103b357600080fd5b806354c8a4f31461033f57806362ddd3c4146103545780636d3d1a581461036757600080fd5b8063240028e8116101ad578063240028e81461028c57806324f65ee7146102d9578063390775371461030a5780634c5ef0ed1461032c57600080fd5b806301ffc9a7146101d4578063181f5a77146101fc57806321df0da714610245575b600080fd5b6101e76101e23660046123e9565b610599565b60405190151581526020015b60405180910390f35b6102386040518060400160405280601781526020017f4d6f636b4c425443546f6b656e506f6f6c20312e352e3100000000000000000081525081565b6040516101f39190612499565b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101f3565b6101e761029a3660046124ac565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff90811691161490565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101f3565b61031d6103183660046124e2565b61067e565b604051905181526020016101f3565b6101e761033a36600461253b565b6107f3565b61035261034d36600461260a565b61083d565b005b61035261036236600461253b565b6108b8565b60095473ffffffffffffffffffffffffffffffffffffffff16610267565b610352610955565b61035261039b3660046124ac565b610a23565b6101e76103ae366004612676565b610aa4565b60015473ffffffffffffffffffffffffffffffffffffffff16610267565b6103e46103df366004612691565b610abb565b6040516101f391906126cc565b6104046103ff366004612676565b610cb3565b6040516101f39190612723565b610419610e1e565b6040516101f391906127a5565b61035261043436600461253b565b610e2f565b61044c610447366004612676565b610f47565b6040516101f3919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b60045473ffffffffffffffffffffffffffffffffffffffff16610267565b6102386104d4366004612676565b61101c565b6103526104e73660046124ac565b6110cc565b6104f46111a7565b6040516101f391906127ff565b61044c61050f366004612676565b61125f565b61035261052236600461297e565b611331565b7f0000000000000000000000000000000000000000000000000000000000000000610267565b7f00000000000000000000000000000000000000000000000000000000000000006101e7565b61035261058136600461260a565b6113b5565b6103526105943660046124ac565b6118c7565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167faff2afbf00000000000000000000000000000000000000000000000000000000148061062c57507fffffffff0000000000000000000000000000000000000000000000000000000082167f0e64dd2900000000000000000000000000000000000000000000000000000000145b8061067857507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b60408051602081019091526000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166340c10f196106d960608501604086016124ac565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff909116600482015260608501356024820152604401600060405180830381600087803b15801561074957600080fd5b505af115801561075d573d6000803e3d6000fd5b506107729250505060608301604084016124ac565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f084606001356040516107d491815260200190565b60405180910390a3506040805160208101909152606090910135815290565b600061083583836040516108089291906129c3565b604080519182900390912067ffffffffffffffff87166000908152600760205291909120600501906118db565b949350505050565b6108456118f6565b6108b28484808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152505060408051602080880282810182019093528782529093508792508691829185019084908082843760009201919091525061194992505050565b50505050565b6108c06118f6565b6108c983610aa4565b610910576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024015b60405180910390fd5b6109508383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611aff92505050565b505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146109a6576040517f02b543c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560008054909116815560405173ffffffffffffffffffffffffffffffffffffffff909216929183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610a2b6118f6565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d091749060200160405180910390a150565b6000610678600567ffffffffffffffff84166118db565b6040805180820190915260608082526020820152606080604051602001610b05907f1234abcd00000000000000000000000000000000000000000000000000000000815260040190565b6040516020818303038152906040529150600282604051610b2691906129d3565b602060405180830381855afa158015610b43573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190610b6691906129ef565b604051602001610b7891815260200190565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152908290527f42966c680000000000000000000000000000000000000000000000000000000082526060860135600483015291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906342966c6890602401600060405180830381600087803b158015610c3157600080fd5b505af1158015610c45573d6000803e3d6000fd5b5050604051606087013581523392507f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7915060200160405180910390a26040518060400160405280610ca38660200160208101906104d49190612676565b8152602001919091529392505050565b67ffffffffffffffff8116600090815260076020526040812060609190610cdc90600501611bf9565b90506000815167ffffffffffffffff811115610cfa57610cfa612841565b604051908082528060200260200182016040528015610d2d57816020015b6060815260200190600190039081610d185790505b50905060005b8251811015610e165760086000848381518110610d5257610d52612a08565b602002602001015181526020019081526020016000208054610d7390612a37565b80601f0160208091040260200160405190810160405280929190818152602001828054610d9f90612a37565b8015610dec5780601f10610dc157610100808354040283529160200191610dec565b820191906000526020600020905b815481529060010190602001808311610dcf57829003601f168201915b5050505050828281518110610e0357610e03612a08565b6020908102919091010152600101610d33565b509392505050565b6060610e2a6002611bf9565b905090565b610e376118f6565b610e4083610aa4565b610e82576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602401610907565b610ec28282604051610e959291906129c3565b604080519182900390912067ffffffffffffffff8616600090815260076020529190912060050190611c06565b610efe578282826040517f74f23c7c00000000000000000000000000000000000000000000000000000000815260040161090793929190612ad3565b8267ffffffffffffffff167f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d768383604051610f3a929190612af7565b60405180910390a2505050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845260028201546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff16151594820194909452600390910154808416606083015291909104909116608082015261067890611c12565b67ffffffffffffffff8116600090815260076020526040902060040180546060919061104790612a37565b80601f016020809104026020016040519081016040528092919081815260200182805461107390612a37565b80156110c05780601f10611095576101008083540402835291602001916110c0565b820191906000526020600020905b8154815290600101906020018083116110a357829003601f168201915b50505050509050919050565b6110d46118f6565b73ffffffffffffffffffffffffffffffffffffffff8116611121576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684910160405180910390a15050565b606060006111b56005611bf9565b90506000815167ffffffffffffffff8111156111d3576111d3612841565b6040519080825280602002602001820160405280156111fc578160200160208202803683370190505b50905060005b82518110156112585782818151811061121d5761121d612a08565b602002602001015182828151811061123757611237612a08565b67ffffffffffffffff90921660209283029190910190910152600101611202565b5092915050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845281546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff16151594820194909452600190910154808416606083015291909104909116608082015261067890611c12565b60095473ffffffffffffffffffffffffffffffffffffffff163314801590611371575060015473ffffffffffffffffffffffffffffffffffffffff163314155b156113aa576040517f8e4a23d6000000000000000000000000000000000000000000000000000000008152336004820152602401610907565b610950838383611cc4565b6113bd6118f6565b60005b838110156115aa5760008585838181106113dc576113dc612a08565b90506020020160208101906113f19190612676565b9050611408600567ffffffffffffffff8316611c06565b61144a576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610907565b67ffffffffffffffff8116600090815260076020526040812061146f90600501611bf9565b905060005b81518110156114db576114d282828151811061149257611492612a08565b6020026020010151600760008667ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600501611c0690919063ffffffff16565b50600101611474565b5067ffffffffffffffff8216600090815260076020526040812080547fffffffffffffffffffffff00000000000000000000000000000000000000000090811682556001820183905560028201805490911690556003810182905590611544600483018261237c565b600582016000818161155682826123b6565b505060405167ffffffffffffffff871681527f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d85991694506020019250611598915050565b60405180910390a150506001016113c0565b5060005b818110156118c05760008383838181106115ca576115ca612a08565b90506020028101906115dc9190612b0b565b6115e590612bcd565b90506115f681606001516000611dae565b61160581608001516000611dae565b806040015151600003611644576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805161165c9060059067ffffffffffffffff16611eeb565b6116a15780516040517f1d5ad3c500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610907565b805167ffffffffffffffff16600090815260076020908152604091829020825160a08082018552606080870180518601516fffffffffffffffffffffffffffffffff90811680865263ffffffff42168689018190528351511515878b0181905284518a0151841686890181905294518b0151841660809889018190528954740100000000000000000000000000000000000000009283027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff7001000000000000000000000000000000008087027fffffffffffffffffffffffff000000000000000000000000000000000000000094851690981788178216929092178d5592810290971760018c01558c519889018d52898e0180518d01518716808b528a8e019590955280515115158a8f018190528151909d01518716988a01899052518d0151909516979098018790526002890180549a9091029990931617179094169590951790925590920290911760038201559082015160048201906118249082612d44565b5060005b8260200151518110156118685761186083600001518460200151838151811061185357611853612a08565b6020026020010151611aff565b600101611828565b507f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c282600001518360400151846060015185608001516040516118ae9493929190612e5e565b60405180910390a150506001016115ae565b5050505050565b6118cf6118f6565b6118d881611ef7565b50565b600081815260018301602052604081205415155b9392505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611947576040517f2b5c74de00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b7f00000000000000000000000000000000000000000000000000000000000000006119a0576040517f35f4a7b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8251811015611a365760008382815181106119c0576119c0612a08565b602002602001015190506119de816002611fbb90919063ffffffff16565b15611a2d5760405173ffffffffffffffffffffffffffffffffffffffff821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b506001016119a3565b5060005b8151811015610950576000828281518110611a5757611a57612a08565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611a9b5750611af7565b611aa6600282611fdd565b15611af55760405173ffffffffffffffffffffffffffffffffffffffff821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b600101611a3a565b8051600003611b3a576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160208083019190912067ffffffffffffffff8416600090815260079092526040909120611b6c9060050182611eeb565b611ba65782826040517f393b8ad2000000000000000000000000000000000000000000000000000000008152600401610907929190612ef7565b6000818152600860205260409020611bbe8382612d44565b508267ffffffffffffffff167f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea83604051610f3a9190612499565b606060006118ef83611fff565b60006118ef838361205a565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152611ca082606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff1642611c849190612f49565b85608001516fffffffffffffffffffffffffffffffff1661214d565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b611ccd83610aa4565b611d0f576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602401610907565b611d1a826000611dae565b67ffffffffffffffff83166000908152600760205260409020611d3d9083612175565b611d48816000611dae565b67ffffffffffffffff83166000908152600760205260409020611d6e9060020182612175565b7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b838383604051611da193929190612f5c565b60405180910390a1505050565b815115611e795781602001516fffffffffffffffffffffffffffffffff1682604001516fffffffffffffffffffffffffffffffff16101580611e04575060408201516fffffffffffffffffffffffffffffffff16155b15611e3d57816040517f8020d1240000000000000000000000000000000000000000000000000000000081526004016109079190612fdf565b8015611e75576040517f433fc33d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b60408201516fffffffffffffffffffffffffffffffff16151580611eb2575060208201516fffffffffffffffffffffffffffffffff1615155b15611e7557816040517fd68af9cc0000000000000000000000000000000000000000000000000000000081526004016109079190612fdf565b60006118ef8383612317565b3373ffffffffffffffffffffffffffffffffffffffff821603611f46576040517fdad89dca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217835560015460405192939116917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006118ef8373ffffffffffffffffffffffffffffffffffffffff841661205a565b60006118ef8373ffffffffffffffffffffffffffffffffffffffff8416612317565b6060816000018054806020026020016040519081016040528092919081815260200182805480156110c057602002820191906000526020600020905b81548152602001906001019080831161203b5750505050509050919050565b6000818152600183016020526040812054801561214357600061207e600183612f49565b855490915060009061209290600190612f49565b90508082146120f75760008660000182815481106120b2576120b2612a08565b90600052602060002001549050808760000184815481106120d5576120d5612a08565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806121085761210861301b565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610678565b6000915050610678565b600061216c8561215d848661304a565b6121679087613061565b612366565b95945050505050565b815460009061219e90700100000000000000000000000000000000900463ffffffff1642612f49565b9050801561224057600183015483546121e6916fffffffffffffffffffffffffffffffff8082169281169185917001000000000000000000000000000000009091041661214d565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b60208201518354612266916fffffffffffffffffffffffffffffffff9081169116612366565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1990611da1908490612fdf565b600081815260018301602052604081205461235e57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610678565b506000610678565b600081831061237557816118ef565b5090919050565b50805461238890612a37565b6000825580601f10612398575050565b601f0160209004906000526020600020908101906118d891906123d0565b50805460008255906000526020600020908101906118d891905b5b808211156123e557600081556001016123d1565b5090565b6000602082840312156123fb57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146118ef57600080fd5b60005b8381101561244657818101518382015260200161242e565b50506000910152565b6000815180845261246781602086016020860161242b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006118ef602083018461244f565b6000602082840312156124be57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146118ef57600080fd5b6000602082840312156124f457600080fd5b813567ffffffffffffffff81111561250b57600080fd5b820161010081850312156118ef57600080fd5b803567ffffffffffffffff8116811461253657600080fd5b919050565b60008060006040848603121561255057600080fd5b6125598461251e565b9250602084013567ffffffffffffffff8082111561257657600080fd5b818601915086601f83011261258a57600080fd5b81358181111561259957600080fd5b8760208285010111156125ab57600080fd5b6020830194508093505050509250925092565b60008083601f8401126125d057600080fd5b50813567ffffffffffffffff8111156125e857600080fd5b6020830191508360208260051b850101111561260357600080fd5b9250929050565b6000806000806040858703121561262057600080fd5b843567ffffffffffffffff8082111561263857600080fd5b612644888389016125be565b9096509450602087013591508082111561265d57600080fd5b5061266a878288016125be565b95989497509550505050565b60006020828403121561268857600080fd5b6118ef8261251e565b6000602082840312156126a357600080fd5b813567ffffffffffffffff8111156126ba57600080fd5b820160a081850312156118ef57600080fd5b6020815260008251604060208401526126e8606084018261244f565b905060208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084830301604085015261216c828261244f565b600060208083016020845280855180835260408601915060408160051b87010192506020870160005b82811015612798577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc088860301845261278685835161244f565b9450928501929085019060010161274c565b5092979650505050505050565b6020808252825182820181905260009190848201906040850190845b818110156127f357835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016127c1565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b818110156127f357835167ffffffffffffffff168352928401929184019160010161281b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff8111828210171561289357612893612841565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156128e0576128e0612841565b604052919050565b80356fffffffffffffffffffffffffffffffff8116811461253657600080fd5b60006060828403121561291a57600080fd5b6040516060810181811067ffffffffffffffff8211171561293d5761293d612841565b6040529050808235801515811461295357600080fd5b8152612961602084016128e8565b6020820152612972604084016128e8565b60408201525092915050565b600080600060e0848603121561299357600080fd5b61299c8461251e565b92506129ab8560208601612908565b91506129ba8560808601612908565b90509250925092565b8183823760009101908152919050565b600082516129e581846020870161242b565b9190910192915050565b600060208284031215612a0157600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600181811c90821680612a4b57607f821691505b602082108103612a84577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b67ffffffffffffffff8416815260406020820152600061216c604083018486612a8a565b602081526000610835602083018486612a8a565b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee18336030181126129e557600080fd5b600082601f830112612b5057600080fd5b813567ffffffffffffffff811115612b6a57612b6a612841565b612b9b60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601612899565b818152846020838601011115612bb057600080fd5b816020850160208301376000918101602001919091529392505050565b60006101208236031215612be057600080fd5b612be8612870565b612bf18361251e565b815260208084013567ffffffffffffffff80821115612c0f57600080fd5b9085019036601f830112612c2257600080fd5b813581811115612c3457612c34612841565b8060051b612c43858201612899565b9182528381018501918581019036841115612c5d57600080fd5b86860192505b83831015612c9957823585811115612c7b5760008081fd5b612c893689838a0101612b3f565b8352509186019190860190612c63565b8087890152505050506040860135925080831115612cb657600080fd5b5050612cc436828601612b3f565b604083015250612cd73660608501612908565b6060820152612ce93660c08501612908565b608082015292915050565b601f821115610950576000816000526020600020601f850160051c81016020861015612d1d5750805b601f850160051c820191505b81811015612d3c57828155600101612d29565b505050505050565b815167ffffffffffffffff811115612d5e57612d5e612841565b612d7281612d6c8454612a37565b84612cf4565b602080601f831160018114612dc55760008415612d8f5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555612d3c565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015612e1257888601518255948401946001909101908401612df3565b5085821015612e4e57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b600061010067ffffffffffffffff87168352806020840152612e828184018761244f565b8551151560408581019190915260208701516fffffffffffffffffffffffffffffffff9081166060870152908701511660808501529150612ec09050565b8251151560a083015260208301516fffffffffffffffffffffffffffffffff90811660c084015260408401511660e083015261216c565b67ffffffffffffffff83168152604060208201526000610835604083018461244f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561067857610678612f1a565b67ffffffffffffffff8416815260e08101612fa860208301858051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b82511515608083015260208301516fffffffffffffffffffffffffffffffff90811660a084015260408401511660c0830152610835565b6060810161067882848051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b808202811582820484141761067857610678612f1a565b8082018082111561067857610678612f1a56fea164736f6c6343000818000a",
}

var MockLBTCTokenPoolABI = MockLBTCTokenPoolMetaData.ABI

var MockLBTCTokenPoolBin = MockLBTCTokenPoolMetaData.Bin

func DeployMockLBTCTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, allowlist []common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *MockLBTCTokenPool, error) {
	parsed, err := MockLBTCTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockLBTCTokenPoolBin), backend, token, allowlist, rmnProxy, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockLBTCTokenPool{address: address, abi: *parsed, MockLBTCTokenPoolCaller: MockLBTCTokenPoolCaller{contract: contract}, MockLBTCTokenPoolTransactor: MockLBTCTokenPoolTransactor{contract: contract}, MockLBTCTokenPoolFilterer: MockLBTCTokenPoolFilterer{contract: contract}}, nil
}

type MockLBTCTokenPool struct {
	address common.Address
	abi     abi.ABI
	MockLBTCTokenPoolCaller
	MockLBTCTokenPoolTransactor
	MockLBTCTokenPoolFilterer
}

type MockLBTCTokenPoolCaller struct {
	contract *bind.BoundContract
}

type MockLBTCTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type MockLBTCTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type MockLBTCTokenPoolSession struct {
	Contract     *MockLBTCTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockLBTCTokenPoolCallerSession struct {
	Contract *MockLBTCTokenPoolCaller
	CallOpts bind.CallOpts
}

type MockLBTCTokenPoolTransactorSession struct {
	Contract     *MockLBTCTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type MockLBTCTokenPoolRaw struct {
	Contract *MockLBTCTokenPool
}

type MockLBTCTokenPoolCallerRaw struct {
	Contract *MockLBTCTokenPoolCaller
}

type MockLBTCTokenPoolTransactorRaw struct {
	Contract *MockLBTCTokenPoolTransactor
}

func NewMockLBTCTokenPool(address common.Address, backend bind.ContractBackend) (*MockLBTCTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(MockLBTCTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockLBTCTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPool{address: address, abi: abi, MockLBTCTokenPoolCaller: MockLBTCTokenPoolCaller{contract: contract}, MockLBTCTokenPoolTransactor: MockLBTCTokenPoolTransactor{contract: contract}, MockLBTCTokenPoolFilterer: MockLBTCTokenPoolFilterer{contract: contract}}, nil
}

func NewMockLBTCTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*MockLBTCTokenPoolCaller, error) {
	contract, err := bindMockLBTCTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolCaller{contract: contract}, nil
}

func NewMockLBTCTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*MockLBTCTokenPoolTransactor, error) {
	contract, err := bindMockLBTCTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolTransactor{contract: contract}, nil
}

func NewMockLBTCTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*MockLBTCTokenPoolFilterer, error) {
	contract, err := bindMockLBTCTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolFilterer{contract: contract}, nil
}

func bindMockLBTCTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockLBTCTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockLBTCTokenPool.Contract.MockLBTCTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.MockLBTCTokenPoolTransactor.contract.Transfer(opts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.MockLBTCTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockLBTCTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.contract.Transfer(opts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _MockLBTCTokenPool.Contract.GetAllowList(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _MockLBTCTokenPool.Contract.GetAllowList(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _MockLBTCTokenPool.Contract.GetAllowListEnabled(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _MockLBTCTokenPool.Contract.GetAllowListEnabled(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _MockLBTCTokenPool.Contract.GetCurrentInboundRateLimiterState(&_MockLBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _MockLBTCTokenPool.Contract.GetCurrentInboundRateLimiterState(&_MockLBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _MockLBTCTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_MockLBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _MockLBTCTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_MockLBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _MockLBTCTokenPool.Contract.GetRateLimitAdmin(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _MockLBTCTokenPool.Contract.GetRateLimitAdmin(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _MockLBTCTokenPool.Contract.GetRemotePools(&_MockLBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _MockLBTCTokenPool.Contract.GetRemotePools(&_MockLBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _MockLBTCTokenPool.Contract.GetRemoteToken(&_MockLBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _MockLBTCTokenPool.Contract.GetRemoteToken(&_MockLBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _MockLBTCTokenPool.Contract.GetRmnProxy(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _MockLBTCTokenPool.Contract.GetRmnProxy(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) GetRouter() (common.Address, error) {
	return _MockLBTCTokenPool.Contract.GetRouter(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) GetRouter() (common.Address, error) {
	return _MockLBTCTokenPool.Contract.GetRouter(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _MockLBTCTokenPool.Contract.GetSupportedChains(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _MockLBTCTokenPool.Contract.GetSupportedChains(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) GetToken() (common.Address, error) {
	return _MockLBTCTokenPool.Contract.GetToken(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _MockLBTCTokenPool.Contract.GetToken(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _MockLBTCTokenPool.Contract.GetTokenDecimals(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _MockLBTCTokenPool.Contract.GetTokenDecimals(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _MockLBTCTokenPool.Contract.IsRemotePool(&_MockLBTCTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _MockLBTCTokenPool.Contract.IsRemotePool(&_MockLBTCTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _MockLBTCTokenPool.Contract.IsSupportedChain(&_MockLBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _MockLBTCTokenPool.Contract.IsSupportedChain(&_MockLBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _MockLBTCTokenPool.Contract.IsSupportedToken(&_MockLBTCTokenPool.CallOpts, token)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _MockLBTCTokenPool.Contract.IsSupportedToken(&_MockLBTCTokenPool.CallOpts, token)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) Owner() (common.Address, error) {
	return _MockLBTCTokenPool.Contract.Owner(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) Owner() (common.Address, error) {
	return _MockLBTCTokenPool.Contract.Owner(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockLBTCTokenPool.Contract.SupportsInterface(&_MockLBTCTokenPool.CallOpts, interfaceId)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockLBTCTokenPool.Contract.SupportsInterface(&_MockLBTCTokenPool.CallOpts, interfaceId)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) TypeAndVersion() (string, error) {
	return _MockLBTCTokenPool.Contract.TypeAndVersion(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _MockLBTCTokenPool.Contract.TypeAndVersion(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockLBTCTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.AcceptOwnership(&_MockLBTCTokenPool.TransactOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.AcceptOwnership(&_MockLBTCTokenPool.TransactOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockLBTCTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.AddRemotePool(&_MockLBTCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.AddRemotePool(&_MockLBTCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.ApplyAllowListUpdates(&_MockLBTCTokenPool.TransactOpts, removes, adds)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.ApplyAllowListUpdates(&_MockLBTCTokenPool.TransactOpts, removes, adds)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _MockLBTCTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.ApplyChainUpdates(&_MockLBTCTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.ApplyChainUpdates(&_MockLBTCTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _MockLBTCTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.LockOrBurn(&_MockLBTCTokenPool.TransactOpts, lockOrBurnIn)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.LockOrBurn(&_MockLBTCTokenPool.TransactOpts, lockOrBurnIn)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _MockLBTCTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.ReleaseOrMint(&_MockLBTCTokenPool.TransactOpts, releaseOrMintIn)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.ReleaseOrMint(&_MockLBTCTokenPool.TransactOpts, releaseOrMintIn)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockLBTCTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.RemoveRemotePool(&_MockLBTCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.RemoveRemotePool(&_MockLBTCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _MockLBTCTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.SetChainRateLimiterConfig(&_MockLBTCTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.SetChainRateLimiterConfig(&_MockLBTCTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.SetRateLimitAdmin(&_MockLBTCTokenPool.TransactOpts, rateLimitAdmin)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.SetRateLimitAdmin(&_MockLBTCTokenPool.TransactOpts, rateLimitAdmin)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.contract.Transact(opts, "setRouter", newRouter)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.SetRouter(&_MockLBTCTokenPool.TransactOpts, newRouter)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.SetRouter(&_MockLBTCTokenPool.TransactOpts, newRouter)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.TransferOwnership(&_MockLBTCTokenPool.TransactOpts, to)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.TransferOwnership(&_MockLBTCTokenPool.TransactOpts, to)
}

type MockLBTCTokenPoolAllowListAddIterator struct {
	Event *MockLBTCTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolAllowListAdd)
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
		it.Event = new(MockLBTCTokenPoolAllowListAdd)
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

func (it *MockLBTCTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*MockLBTCTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolAllowListAddIterator{contract: _MockLBTCTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolAllowListAdd)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*MockLBTCTokenPoolAllowListAdd, error) {
	event := new(MockLBTCTokenPoolAllowListAdd)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolAllowListRemoveIterator struct {
	Event *MockLBTCTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolAllowListRemove)
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
		it.Event = new(MockLBTCTokenPoolAllowListRemove)
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

func (it *MockLBTCTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*MockLBTCTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolAllowListRemoveIterator{contract: _MockLBTCTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolAllowListRemove)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*MockLBTCTokenPoolAllowListRemove, error) {
	event := new(MockLBTCTokenPoolAllowListRemove)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolBurnedIterator struct {
	Event *MockLBTCTokenPoolBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolBurned)
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
		it.Event = new(MockLBTCTokenPoolBurned)
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

func (it *MockLBTCTokenPoolBurnedIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*MockLBTCTokenPoolBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolBurnedIterator{contract: _MockLBTCTokenPool.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolBurned)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseBurned(log types.Log) (*MockLBTCTokenPoolBurned, error) {
	event := new(MockLBTCTokenPoolBurned)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolChainAddedIterator struct {
	Event *MockLBTCTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolChainAdded)
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
		it.Event = new(MockLBTCTokenPoolChainAdded)
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

func (it *MockLBTCTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*MockLBTCTokenPoolChainAddedIterator, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolChainAddedIterator{contract: _MockLBTCTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolChainAdded)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseChainAdded(log types.Log) (*MockLBTCTokenPoolChainAdded, error) {
	event := new(MockLBTCTokenPoolChainAdded)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolChainConfiguredIterator struct {
	Event *MockLBTCTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolChainConfigured)
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
		it.Event = new(MockLBTCTokenPoolChainConfigured)
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

func (it *MockLBTCTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*MockLBTCTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolChainConfiguredIterator{contract: _MockLBTCTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolChainConfigured)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseChainConfigured(log types.Log) (*MockLBTCTokenPoolChainConfigured, error) {
	event := new(MockLBTCTokenPoolChainConfigured)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolChainRemovedIterator struct {
	Event *MockLBTCTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolChainRemoved)
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
		it.Event = new(MockLBTCTokenPoolChainRemoved)
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

func (it *MockLBTCTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*MockLBTCTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolChainRemovedIterator{contract: _MockLBTCTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolChainRemoved)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseChainRemoved(log types.Log) (*MockLBTCTokenPoolChainRemoved, error) {
	event := new(MockLBTCTokenPoolChainRemoved)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolConfigChangedIterator struct {
	Event *MockLBTCTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolConfigChanged)
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
		it.Event = new(MockLBTCTokenPoolConfigChanged)
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

func (it *MockLBTCTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*MockLBTCTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolConfigChangedIterator{contract: _MockLBTCTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolConfigChanged)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseConfigChanged(log types.Log) (*MockLBTCTokenPoolConfigChanged, error) {
	event := new(MockLBTCTokenPoolConfigChanged)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolLockedIterator struct {
	Event *MockLBTCTokenPoolLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolLocked)
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
		it.Event = new(MockLBTCTokenPoolLocked)
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

func (it *MockLBTCTokenPoolLockedIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*MockLBTCTokenPoolLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolLockedIterator{contract: _MockLBTCTokenPool.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolLocked)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseLocked(log types.Log) (*MockLBTCTokenPoolLocked, error) {
	event := new(MockLBTCTokenPoolLocked)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolMintedIterator struct {
	Event *MockLBTCTokenPoolMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolMinted)
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
		it.Event = new(MockLBTCTokenPoolMinted)
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

func (it *MockLBTCTokenPoolMintedIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*MockLBTCTokenPoolMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolMintedIterator{contract: _MockLBTCTokenPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolMinted)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseMinted(log types.Log) (*MockLBTCTokenPoolMinted, error) {
	event := new(MockLBTCTokenPoolMinted)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolOwnershipTransferRequestedIterator struct {
	Event *MockLBTCTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolOwnershipTransferRequested)
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
		it.Event = new(MockLBTCTokenPoolOwnershipTransferRequested)
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

func (it *MockLBTCTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockLBTCTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolOwnershipTransferRequestedIterator{contract: _MockLBTCTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolOwnershipTransferRequested)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*MockLBTCTokenPoolOwnershipTransferRequested, error) {
	event := new(MockLBTCTokenPoolOwnershipTransferRequested)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolOwnershipTransferredIterator struct {
	Event *MockLBTCTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolOwnershipTransferred)
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
		it.Event = new(MockLBTCTokenPoolOwnershipTransferred)
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

func (it *MockLBTCTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockLBTCTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolOwnershipTransferredIterator{contract: _MockLBTCTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolOwnershipTransferred)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*MockLBTCTokenPoolOwnershipTransferred, error) {
	event := new(MockLBTCTokenPoolOwnershipTransferred)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolRateLimitAdminSetIterator struct {
	Event *MockLBTCTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolRateLimitAdminSet)
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
		it.Event = new(MockLBTCTokenPoolRateLimitAdminSet)
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

func (it *MockLBTCTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*MockLBTCTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolRateLimitAdminSetIterator{contract: _MockLBTCTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolRateLimitAdminSet)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*MockLBTCTokenPoolRateLimitAdminSet, error) {
	event := new(MockLBTCTokenPoolRateLimitAdminSet)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolReleasedIterator struct {
	Event *MockLBTCTokenPoolReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolReleased)
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
		it.Event = new(MockLBTCTokenPoolReleased)
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

func (it *MockLBTCTokenPoolReleasedIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*MockLBTCTokenPoolReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolReleasedIterator{contract: _MockLBTCTokenPool.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolReleased)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseReleased(log types.Log) (*MockLBTCTokenPoolReleased, error) {
	event := new(MockLBTCTokenPoolReleased)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolRemotePoolAddedIterator struct {
	Event *MockLBTCTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolRemotePoolAdded)
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
		it.Event = new(MockLBTCTokenPoolRemotePoolAdded)
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

func (it *MockLBTCTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockLBTCTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolRemotePoolAddedIterator{contract: _MockLBTCTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolRemotePoolAdded)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*MockLBTCTokenPoolRemotePoolAdded, error) {
	event := new(MockLBTCTokenPoolRemotePoolAdded)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolRemotePoolRemovedIterator struct {
	Event *MockLBTCTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolRemotePoolRemoved)
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
		it.Event = new(MockLBTCTokenPoolRemotePoolRemoved)
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

func (it *MockLBTCTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockLBTCTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolRemotePoolRemovedIterator{contract: _MockLBTCTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolRemotePoolRemoved)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*MockLBTCTokenPoolRemotePoolRemoved, error) {
	event := new(MockLBTCTokenPoolRemotePoolRemoved)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockLBTCTokenPoolRouterUpdatedIterator struct {
	Event *MockLBTCTokenPoolRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolRouterUpdated)
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
		it.Event = new(MockLBTCTokenPoolRouterUpdated)
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

func (it *MockLBTCTokenPoolRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*MockLBTCTokenPoolRouterUpdatedIterator, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolRouterUpdatedIterator{contract: _MockLBTCTokenPool.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolRouterUpdated)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseRouterUpdated(log types.Log) (*MockLBTCTokenPoolRouterUpdated, error) {
	event := new(MockLBTCTokenPoolRouterUpdated)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MockLBTCTokenPool.abi.Events["AllowListAdd"].ID:
		return _MockLBTCTokenPool.ParseAllowListAdd(log)
	case _MockLBTCTokenPool.abi.Events["AllowListRemove"].ID:
		return _MockLBTCTokenPool.ParseAllowListRemove(log)
	case _MockLBTCTokenPool.abi.Events["Burned"].ID:
		return _MockLBTCTokenPool.ParseBurned(log)
	case _MockLBTCTokenPool.abi.Events["ChainAdded"].ID:
		return _MockLBTCTokenPool.ParseChainAdded(log)
	case _MockLBTCTokenPool.abi.Events["ChainConfigured"].ID:
		return _MockLBTCTokenPool.ParseChainConfigured(log)
	case _MockLBTCTokenPool.abi.Events["ChainRemoved"].ID:
		return _MockLBTCTokenPool.ParseChainRemoved(log)
	case _MockLBTCTokenPool.abi.Events["ConfigChanged"].ID:
		return _MockLBTCTokenPool.ParseConfigChanged(log)
	case _MockLBTCTokenPool.abi.Events["Locked"].ID:
		return _MockLBTCTokenPool.ParseLocked(log)
	case _MockLBTCTokenPool.abi.Events["Minted"].ID:
		return _MockLBTCTokenPool.ParseMinted(log)
	case _MockLBTCTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _MockLBTCTokenPool.ParseOwnershipTransferRequested(log)
	case _MockLBTCTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _MockLBTCTokenPool.ParseOwnershipTransferred(log)
	case _MockLBTCTokenPool.abi.Events["RateLimitAdminSet"].ID:
		return _MockLBTCTokenPool.ParseRateLimitAdminSet(log)
	case _MockLBTCTokenPool.abi.Events["Released"].ID:
		return _MockLBTCTokenPool.ParseReleased(log)
	case _MockLBTCTokenPool.abi.Events["RemotePoolAdded"].ID:
		return _MockLBTCTokenPool.ParseRemotePoolAdded(log)
	case _MockLBTCTokenPool.abi.Events["RemotePoolRemoved"].ID:
		return _MockLBTCTokenPool.ParseRemotePoolRemoved(log)
	case _MockLBTCTokenPool.abi.Events["RouterUpdated"].ID:
		return _MockLBTCTokenPool.ParseRouterUpdated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MockLBTCTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (MockLBTCTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (MockLBTCTokenPoolBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
}

func (MockLBTCTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (MockLBTCTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (MockLBTCTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (MockLBTCTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (MockLBTCTokenPoolLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (MockLBTCTokenPoolMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (MockLBTCTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (MockLBTCTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (MockLBTCTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (MockLBTCTokenPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (MockLBTCTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (MockLBTCTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (MockLBTCTokenPoolRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (_MockLBTCTokenPool *MockLBTCTokenPool) Address() common.Address {
	return _MockLBTCTokenPool.address
}

type MockLBTCTokenPoolInterface interface {
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

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

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

	FilterAllowListAdd(opts *bind.FilterOpts) (*MockLBTCTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*MockLBTCTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*MockLBTCTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*MockLBTCTokenPoolAllowListRemove, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*MockLBTCTokenPoolBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*MockLBTCTokenPoolBurned, error)

	FilterChainAdded(opts *bind.FilterOpts) (*MockLBTCTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*MockLBTCTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*MockLBTCTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*MockLBTCTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*MockLBTCTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*MockLBTCTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*MockLBTCTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*MockLBTCTokenPoolConfigChanged, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*MockLBTCTokenPoolLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*MockLBTCTokenPoolLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*MockLBTCTokenPoolMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*MockLBTCTokenPoolMinted, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockLBTCTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*MockLBTCTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockLBTCTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*MockLBTCTokenPoolOwnershipTransferred, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*MockLBTCTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*MockLBTCTokenPoolRateLimitAdminSet, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*MockLBTCTokenPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*MockLBTCTokenPoolReleased, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockLBTCTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*MockLBTCTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockLBTCTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*MockLBTCTokenPoolRemotePoolRemoved, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*MockLBTCTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*MockLBTCTokenPoolRouterUpdated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
