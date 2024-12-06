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
	Bin: "0x6101006040523480156200001257600080fd5b506040516200362a3803806200362a833981016040819052620000359162000590565b836012848484336000816200005d57604051639b15e16f60e01b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b038481169190911790915581161562000090576200009081620001eb565b50506001600160a01b0385161580620000b057506001600160a01b038116155b80620000c357506001600160a01b038216155b15620000e2576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b03808616608081905290831660c0526040805163313ce56760e01b8152905163313ce567916004808201926020929091908290030181865afa92505050801562000152575060408051601f3d908101601f191682019092526200014f91810190620006a0565b60015b1562000192578060ff168560ff161462000190576040516332ad3e0760e11b815260ff80871660048301528216602482015260440160405180910390fd5b505b60ff841660a052600480546001600160a01b0319166001600160a01b038316179055825115801560e052620001dc57604080516000815260208101909152620001dc908462000265565b5050505050505050506200071a565b336001600160a01b038216036200021557604051636d6c4ee560e11b815260040160405180910390fd5b600080546001600160a01b0319166001600160a01b03838116918217835560015460405192939116917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60e05162000286576040516335f4a7b360e01b815260040160405180910390fd5b60005b825181101562000311576000838281518110620002aa57620002aa620006cc565b60209081029190910101519050620002c4600282620003c2565b1562000307576040516001600160a01b03821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b5060010162000289565b5060005b8151811015620003bd576000828281518110620003365762000336620006cc565b6020026020010151905060006001600160a01b0316816001600160a01b031603620003625750620003b4565b6200036f600282620003e2565b15620003b2576040516001600160a01b03821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b60010162000315565b505050565b6000620003d9836001600160a01b038416620003f9565b90505b92915050565b6000620003d9836001600160a01b038416620004fd565b60008181526001830160205260408120548015620004f257600062000420600183620006e2565b85549091506000906200043690600190620006e2565b9050808214620004a25760008660000182815481106200045a576200045a620006cc565b9060005260206000200154905080876000018481548110620004805762000480620006cc565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080620004b657620004b662000704565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050620003dc565b6000915050620003dc565b60008181526001830160205260408120546200054657508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155620003dc565b506000620003dc565b6001600160a01b03811681146200056557600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b80516200058b816200054f565b919050565b60008060008060808587031215620005a757600080fd5b8451620005b4816200054f565b602086810151919550906001600160401b0380821115620005d457600080fd5b818801915088601f830112620005e957600080fd5b815181811115620005fe57620005fe62000568565b8060051b604051601f19603f8301168101818110858211171562000626576200062662000568565b60405291825284820192508381018501918b8311156200064557600080fd5b938501935b828510156200066e576200065e856200057e565b845293850193928501926200064a565b80985050505050505062000685604086016200057e565b915062000695606086016200057e565b905092959194509250565b600060208284031215620006b357600080fd5b815160ff81168114620006c557600080fd5b9392505050565b634e487b7160e01b600052603260045260246000fd5b81810381811115620003dc57634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60805160a05160c05160e051612ec8620007626000396000818161054f015261179201526000610529015260006102e0015260008181610247015261029c0152612ec86000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80639a4575b911610104578063c0d78655116100a2578063dc0bd97111610071578063dc0bd97114610527578063e0351e131461054d578063e8a1da1714610573578063f2fde38b1461058657600080fd5b8063c0d78655146104d9578063c4bffe2b146104ec578063c75eea9c14610501578063cf7401f31461051457600080fd5b8063acfecf91116100de578063acfecf9114610426578063af58d59f14610439578063b0f479a1146104a8578063b7946580146104c657600080fd5b80639a4575b9146103d1578063a42a7b8b146103f1578063a7cd63b71461041157600080fd5b806354c8a4f31161017157806379ba50971161014b57806379ba5097146103855780637d54534e1461038d5780638926f54f146103a05780638da5cb5b146103b357600080fd5b806354c8a4f31461033f57806362ddd3c4146103545780636d3d1a581461036757600080fd5b8063240028e8116101ad578063240028e81461028c57806324f65ee7146102d9578063390775371461030a5780634c5ef0ed1461032c57600080fd5b806301ffc9a7146101d4578063181f5a77146101fc57806321df0da714610245575b600080fd5b6101e76101e2366004612230565b610599565b60405190151581526020015b60405180910390f35b6102386040518060400160405280601781526020017f4d6f636b4c425443546f6b656e506f6f6c20312e352e3100000000000000000081525081565b6040516101f391906122e0565b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101f3565b6101e761029a3660046122f3565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff90811691161490565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101f3565b61031d610318366004612329565b61067e565b604051905181526020016101f3565b6101e761033a366004612382565b61071e565b61035261034d366004612451565b610768565b005b610352610362366004612382565b6107e3565b60095473ffffffffffffffffffffffffffffffffffffffff16610267565b610352610880565b61035261039b3660046122f3565b61094e565b6101e76103ae3660046124bd565b6109cf565b60015473ffffffffffffffffffffffffffffffffffffffff16610267565b6103e46103df3660046124d8565b6109e6565b6040516101f39190612513565b6104046103ff3660046124bd565b610afa565b6040516101f3919061256a565b610419610c65565b6040516101f391906125ec565b610352610434366004612382565b610c76565b61044c6104473660046124bd565b610d8e565b6040516101f3919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b60045473ffffffffffffffffffffffffffffffffffffffff16610267565b6102386104d43660046124bd565b610e63565b6103526104e73660046122f3565b610f13565b6104f4610fee565b6040516101f39190612646565b61044c61050f3660046124bd565b6110a6565b6103526105223660046127c5565b611178565b7f0000000000000000000000000000000000000000000000000000000000000000610267565b7f00000000000000000000000000000000000000000000000000000000000000006101e7565b610352610581366004612451565b6111fc565b6103526105943660046122f3565b61170e565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167faff2afbf00000000000000000000000000000000000000000000000000000000148061062c57507fffffffff0000000000000000000000000000000000000000000000000000000082167f0e64dd2900000000000000000000000000000000000000000000000000000000145b8061067857507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b60408051602081019091526000815261069d60608301604084016122f3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f084606001356040516106ff91815260200190565b60405180910390a3506040805160208101909152606090910135815290565b6000610760838360405161073392919061280a565b604080519182900390912067ffffffffffffffff8716600090815260076020529190912060050190611722565b949350505050565b61077061173d565b6107dd8484808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152505060408051602080880282810182019093528782529093508792508691829185019084908082843760009201919091525061179092505050565b50505050565b6107eb61173d565b6107f4836109cf565b61083b576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024015b60405180910390fd5b61087b8383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061194692505050565b505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146108d1576040517f02b543c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560008054909116815560405173ffffffffffffffffffffffffffffffffffffffff909216929183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61095661173d565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d091749060200160405180910390a150565b6000610678600567ffffffffffffffff8416611722565b6040805180820190915260608082526020820152606080604051602001610a30907f1234abcd00000000000000000000000000000000000000000000000000000000815260040190565b6040516020818303038152906040529150600282604051610a51919061281a565b602060405180830381855afa158015610a6e573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190610a919190612836565b604051602001610aa391815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152828201825292508190610aea906104d4908801602089016124bd565b8152602001919091529392505050565b67ffffffffffffffff8116600090815260076020526040812060609190610b2390600501611a40565b90506000815167ffffffffffffffff811115610b4157610b41612688565b604051908082528060200260200182016040528015610b7457816020015b6060815260200190600190039081610b5f5790505b50905060005b8251811015610c5d5760086000848381518110610b9957610b9961284f565b602002602001015181526020019081526020016000208054610bba9061287e565b80601f0160208091040260200160405190810160405280929190818152602001828054610be69061287e565b8015610c335780601f10610c0857610100808354040283529160200191610c33565b820191906000526020600020905b815481529060010190602001808311610c1657829003601f168201915b5050505050828281518110610c4a57610c4a61284f565b6020908102919091010152600101610b7a565b509392505050565b6060610c716002611a40565b905090565b610c7e61173d565b610c87836109cf565b610cc9576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602401610832565b610d098282604051610cdc92919061280a565b604080519182900390912067ffffffffffffffff8616600090815260076020529190912060050190611a4d565b610d45578282826040517f74f23c7c0000000000000000000000000000000000000000000000000000000081526004016108329392919061291a565b8267ffffffffffffffff167f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d768383604051610d8192919061293e565b60405180910390a2505050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845260028201546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff16151594820194909452600390910154808416606083015291909104909116608082015261067890611a59565b67ffffffffffffffff81166000908152600760205260409020600401805460609190610e8e9061287e565b80601f0160208091040260200160405190810160405280929190818152602001828054610eba9061287e565b8015610f075780601f10610edc57610100808354040283529160200191610f07565b820191906000526020600020905b815481529060010190602001808311610eea57829003601f168201915b50505050509050919050565b610f1b61173d565b73ffffffffffffffffffffffffffffffffffffffff8116610f68576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684910160405180910390a15050565b60606000610ffc6005611a40565b90506000815167ffffffffffffffff81111561101a5761101a612688565b604051908082528060200260200182016040528015611043578160200160208202803683370190505b50905060005b825181101561109f578281815181106110645761106461284f565b602002602001015182828151811061107e5761107e61284f565b67ffffffffffffffff90921660209283029190910190910152600101611049565b5092915050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845281546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff16151594820194909452600190910154808416606083015291909104909116608082015261067890611a59565b60095473ffffffffffffffffffffffffffffffffffffffff1633148015906111b8575060015473ffffffffffffffffffffffffffffffffffffffff163314155b156111f1576040517f8e4a23d6000000000000000000000000000000000000000000000000000000008152336004820152602401610832565b61087b838383611b0b565b61120461173d565b60005b838110156113f15760008585838181106112235761122361284f565b905060200201602081019061123891906124bd565b905061124f600567ffffffffffffffff8316611a4d565b611291576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610832565b67ffffffffffffffff811660009081526007602052604081206112b690600501611a40565b905060005b8151811015611322576113198282815181106112d9576112d961284f565b6020026020010151600760008667ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600501611a4d90919063ffffffff16565b506001016112bb565b5067ffffffffffffffff8216600090815260076020526040812080547fffffffffffffffffffffff0000000000000000000000000000000000000000009081168255600182018390556002820180549091169055600381018290559061138b60048301826121c3565b600582016000818161139d82826121fd565b505060405167ffffffffffffffff871681527f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916945060200192506113df915050565b60405180910390a15050600101611207565b5060005b818110156117075760008383838181106114115761141161284f565b90506020028101906114239190612952565b61142c90612a14565b905061143d81606001516000611bf5565b61144c81608001516000611bf5565b80604001515160000361148b576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516114a39060059067ffffffffffffffff16611d32565b6114e85780516040517f1d5ad3c500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610832565b805167ffffffffffffffff16600090815260076020908152604091829020825160a08082018552606080870180518601516fffffffffffffffffffffffffffffffff90811680865263ffffffff42168689018190528351511515878b0181905284518a0151841686890181905294518b0151841660809889018190528954740100000000000000000000000000000000000000009283027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff7001000000000000000000000000000000008087027fffffffffffffffffffffffff000000000000000000000000000000000000000094851690981788178216929092178d5592810290971760018c01558c519889018d52898e0180518d01518716808b528a8e019590955280515115158a8f018190528151909d01518716988a01899052518d0151909516979098018790526002890180549a90910299909316171790941695909517909255909202909117600382015590820151600482019061166b9082612b8b565b5060005b8260200151518110156116af576116a783600001518460200151838151811061169a5761169a61284f565b6020026020010151611946565b60010161166f565b507f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c282600001518360400151846060015185608001516040516116f59493929190612ca5565b60405180910390a150506001016113f5565b5050505050565b61171661173d565b61171f81611d3e565b50565b600081815260018301602052604081205415155b9392505050565b60015473ffffffffffffffffffffffffffffffffffffffff16331461178e576040517f2b5c74de00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b7f00000000000000000000000000000000000000000000000000000000000000006117e7576040517f35f4a7b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b825181101561187d5760008382815181106118075761180761284f565b60200260200101519050611825816002611e0290919063ffffffff16565b156118745760405173ffffffffffffffffffffffffffffffffffffffff821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b506001016117ea565b5060005b815181101561087b57600082828151811061189e5761189e61284f565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036118e2575061193e565b6118ed600282611e24565b1561193c5760405173ffffffffffffffffffffffffffffffffffffffff821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b600101611881565b8051600003611981576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160208083019190912067ffffffffffffffff84166000908152600790925260409091206119b39060050182611d32565b6119ed5782826040517f393b8ad2000000000000000000000000000000000000000000000000000000008152600401610832929190612d3e565b6000818152600860205260409020611a058382612b8b565b508267ffffffffffffffff167f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea83604051610d8191906122e0565b6060600061173683611e46565b60006117368383611ea1565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152611ae782606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff1642611acb9190612d90565b85608001516fffffffffffffffffffffffffffffffff16611f94565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b611b14836109cf565b611b56576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602401610832565b611b61826000611bf5565b67ffffffffffffffff83166000908152600760205260409020611b849083611fbc565b611b8f816000611bf5565b67ffffffffffffffff83166000908152600760205260409020611bb59060020182611fbc565b7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b838383604051611be893929190612da3565b60405180910390a1505050565b815115611cc05781602001516fffffffffffffffffffffffffffffffff1682604001516fffffffffffffffffffffffffffffffff16101580611c4b575060408201516fffffffffffffffffffffffffffffffff16155b15611c8457816040517f8020d1240000000000000000000000000000000000000000000000000000000081526004016108329190612e26565b8015611cbc576040517f433fc33d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b60408201516fffffffffffffffffffffffffffffffff16151580611cf9575060208201516fffffffffffffffffffffffffffffffff1615155b15611cbc57816040517fd68af9cc0000000000000000000000000000000000000000000000000000000081526004016108329190612e26565b6000611736838361215e565b3373ffffffffffffffffffffffffffffffffffffffff821603611d8d576040517fdad89dca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217835560015460405192939116917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006117368373ffffffffffffffffffffffffffffffffffffffff8416611ea1565b60006117368373ffffffffffffffffffffffffffffffffffffffff841661215e565b606081600001805480602002602001604051908101604052809291908181526020018280548015610f0757602002820191906000526020600020905b815481526020019060010190808311611e825750505050509050919050565b60008181526001830160205260408120548015611f8a576000611ec5600183612d90565b8554909150600090611ed990600190612d90565b9050808214611f3e576000866000018281548110611ef957611ef961284f565b9060005260206000200154905080876000018481548110611f1c57611f1c61284f565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080611f4f57611f4f612e62565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610678565b6000915050610678565b6000611fb385611fa48486612e91565b611fae9087612ea8565b6121ad565b95945050505050565b8154600090611fe590700100000000000000000000000000000000900463ffffffff1642612d90565b90508015612087576001830154835461202d916fffffffffffffffffffffffffffffffff80821692811691859170010000000000000000000000000000000090910416611f94565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b602082015183546120ad916fffffffffffffffffffffffffffffffff90811691166121ad565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1990611be8908490612e26565b60008181526001830160205260408120546121a557508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610678565b506000610678565b60008183106121bc5781611736565b5090919050565b5080546121cf9061287e565b6000825580601f106121df575050565b601f01602090049060005260206000209081019061171f9190612217565b508054600082559060005260206000209081019061171f91905b5b8082111561222c5760008155600101612218565b5090565b60006020828403121561224257600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461173657600080fd5b60005b8381101561228d578181015183820152602001612275565b50506000910152565b600081518084526122ae816020860160208601612272565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006117366020830184612296565b60006020828403121561230557600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461173657600080fd5b60006020828403121561233b57600080fd5b813567ffffffffffffffff81111561235257600080fd5b8201610100818503121561173657600080fd5b803567ffffffffffffffff8116811461237d57600080fd5b919050565b60008060006040848603121561239757600080fd5b6123a084612365565b9250602084013567ffffffffffffffff808211156123bd57600080fd5b818601915086601f8301126123d157600080fd5b8135818111156123e057600080fd5b8760208285010111156123f257600080fd5b6020830194508093505050509250925092565b60008083601f84011261241757600080fd5b50813567ffffffffffffffff81111561242f57600080fd5b6020830191508360208260051b850101111561244a57600080fd5b9250929050565b6000806000806040858703121561246757600080fd5b843567ffffffffffffffff8082111561247f57600080fd5b61248b88838901612405565b909650945060208701359150808211156124a457600080fd5b506124b187828801612405565b95989497509550505050565b6000602082840312156124cf57600080fd5b61173682612365565b6000602082840312156124ea57600080fd5b813567ffffffffffffffff81111561250157600080fd5b820160a0818503121561173657600080fd5b60208152600082516040602084015261252f6060840182612296565b905060208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848303016040850152611fb38282612296565b600060208083016020845280855180835260408601915060408160051b87010192506020870160005b828110156125df577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08886030184526125cd858351612296565b94509285019290850190600101612593565b5092979650505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561263a57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101612608565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561263a57835167ffffffffffffffff1683529284019291840191600101612662565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff811182821017156126da576126da612688565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561272757612727612688565b604052919050565b80356fffffffffffffffffffffffffffffffff8116811461237d57600080fd5b60006060828403121561276157600080fd5b6040516060810181811067ffffffffffffffff8211171561278457612784612688565b6040529050808235801515811461279a57600080fd5b81526127a86020840161272f565b60208201526127b96040840161272f565b60408201525092915050565b600080600060e084860312156127da57600080fd5b6127e384612365565b92506127f2856020860161274f565b9150612801856080860161274f565b90509250925092565b8183823760009101908152919050565b6000825161282c818460208701612272565b9190910192915050565b60006020828403121561284857600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600181811c9082168061289257607f821691505b6020821081036128cb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b67ffffffffffffffff84168152604060208201526000611fb36040830184866128d1565b6020815260006107606020830184866128d1565b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee183360301811261282c57600080fd5b600082601f83011261299757600080fd5b813567ffffffffffffffff8111156129b1576129b1612688565b6129e260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016126e0565b8181528460208386010111156129f757600080fd5b816020850160208301376000918101602001919091529392505050565b60006101208236031215612a2757600080fd5b612a2f6126b7565b612a3883612365565b815260208084013567ffffffffffffffff80821115612a5657600080fd5b9085019036601f830112612a6957600080fd5b813581811115612a7b57612a7b612688565b8060051b612a8a8582016126e0565b9182528381018501918581019036841115612aa457600080fd5b86860192505b83831015612ae057823585811115612ac25760008081fd5b612ad03689838a0101612986565b8352509186019190860190612aaa565b8087890152505050506040860135925080831115612afd57600080fd5b5050612b0b36828601612986565b604083015250612b1e366060850161274f565b6060820152612b303660c0850161274f565b608082015292915050565b601f82111561087b576000816000526020600020601f850160051c81016020861015612b645750805b601f850160051c820191505b81811015612b8357828155600101612b70565b505050505050565b815167ffffffffffffffff811115612ba557612ba5612688565b612bb981612bb3845461287e565b84612b3b565b602080601f831160018114612c0c5760008415612bd65750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555612b83565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015612c5957888601518255948401946001909101908401612c3a565b5085821015612c9557878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b600061010067ffffffffffffffff87168352806020840152612cc981840187612296565b8551151560408581019190915260208701516fffffffffffffffffffffffffffffffff9081166060870152908701511660808501529150612d079050565b8251151560a083015260208301516fffffffffffffffffffffffffffffffff90811660c084015260408401511660e0830152611fb3565b67ffffffffffffffff831681526040602082015260006107606040830184612296565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561067857610678612d61565b67ffffffffffffffff8416815260e08101612def60208301858051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b82511515608083015260208301516fffffffffffffffffffffffffffffffff90811660a084015260408401511660c0830152610760565b6060810161067882848051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b808202811582820484141761067857610678612d61565b8082018082111561067857610678612d6156fea164736f6c6343000818000a",
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
