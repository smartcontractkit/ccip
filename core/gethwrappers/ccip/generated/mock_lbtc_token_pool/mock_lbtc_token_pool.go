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
	Allowed                   bool
	RemotePoolAddress         []byte
	RemoteTokenAddress        []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

var MockLBTCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"destPoolData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AllowListNotEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotARampOnRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"DisabledNonZeroRateLimit\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"InvalidRateLimitRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidSourcePoolAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"NonExistentChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateLimitMustBeDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remoteToken\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"previousPoolAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"RemotePoolSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRouter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"remoteTokenAddress\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"internalType\":\"structTokenPool.ChainUpdate[]\",\"name\":\"chains\",\"type\":\"tuple[]\"}],\"name\":\"applyChainUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentInboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentOutboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRateLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRemotePool\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRemoteToken\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRmnProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedChains\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_destPoolData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"isSupportedChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isSupportedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"internalType\":\"structPool.LockOrBurnInV1\",\"name\":\"lockOrBurnIn\",\"type\":\"tuple\"}],\"name\":\"lockOrBurn\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destPoolData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.LockOrBurnOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"originalSender\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainTokenData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.ReleaseOrMintInV1\",\"name\":\"releaseOrMintIn\",\"type\":\"tuple\"}],\"name\":\"releaseOrMint\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"destinationAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundConfig\",\"type\":\"tuple\"}],\"name\":\"setChainRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rateLimitAdmin\",\"type\":\"address\"}],\"name\":\"setRateLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"setRemotePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200363f3803806200363f83398101604081905262000034916200061f565b8484848433806000816200008f5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c257620000c28162000189565b5050506001600160a01b0384161580620000e357506001600160a01b038116155b80620000f657506001600160a01b038216155b1562000115576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b0384811660805282811660a052600480546001600160a01b031916918316919091179055825115801560c052620001685760408051600081526020810190915262000168908462000234565b5050505080600990816200017d9190620007d1565b505050505050620008eb565b336001600160a01b03821603620001e35760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000086565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60c05162000255576040516335f4a7b360e01b815260040160405180910390fd5b60005b8251811015620002e05760008382815181106200027957620002796200089d565b602090810291909101015190506200029360028262000391565b15620002d6576040516001600160a01b03821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b5060010162000258565b5060005b81518110156200038c5760008282815181106200030557620003056200089d565b6020026020010151905060006001600160a01b0316816001600160a01b03160362000331575062000383565b6200033e600282620003b1565b1562000381576040516001600160a01b03821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b600101620002e4565b505050565b6000620003a8836001600160a01b038416620003c8565b90505b92915050565b6000620003a8836001600160a01b038416620004cc565b60008181526001830160205260408120548015620004c1576000620003ef600183620008b3565b85549091506000906200040590600190620008b3565b9050808214620004715760008660000182815481106200042957620004296200089d565b90600052602060002001549050808760000184815481106200044f576200044f6200089d565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080620004855762000485620008d5565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050620003ab565b6000915050620003ab565b60008181526001830160205260408120546200051557508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155620003ab565b506000620003ab565b6001600160a01b03811681146200053457600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171562000578576200057862000537565b604052919050565b80516200058d816200051e565b919050565b600082601f830112620005a457600080fd5b81516001600160401b03811115620005c057620005c062000537565b6020620005d6601f8301601f191682016200054d565b8281528582848701011115620005eb57600080fd5b60005b838110156200060b578581018301518282018401528201620005ee565b506000928101909101919091529392505050565b600080600080600060a086880312156200063857600080fd5b855162000645816200051e565b602087810151919650906001600160401b03808211156200066557600080fd5b818901915089601f8301126200067a57600080fd5b8151818111156200068f576200068f62000537565b8060051b620006a08582016200054d565b918252838101850191858101908d841115620006bb57600080fd5b948601945b83861015620006e95785519250620006d8836200051e565b8282529486019490860190620006c0565b9950620006fd9250505060408a0162000580565b95506200070d60608a0162000580565b945060808901519250808311156200072457600080fd5b5050620007348882890162000592565b9150509295509295909350565b600181811c908216806200075657607f821691505b6020821081036200077757634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200038c576000816000526020600020601f850160051c81016020861015620007a85750805b601f850160051c820191505b81811015620007c957828155600101620007b4565b505050505050565b81516001600160401b03811115620007ed57620007ed62000537565b6200080581620007fe845462000741565b846200077d565b602080601f8311600181146200083d5760008415620008245750858301515b600019600386901b1c1916600185901b178555620007c9565b600085815260208120601f198616915b828110156200086e578886015182559484019460019091019084016200084d565b50858210156200088d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b81810381811115620003ab57634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60805160a05160c051612d0862000937600039600081816104f001526116da015260006104ca01526000818161024401528181610299015281816107700152610c570152612d086000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c80638da5cb5b116100f9578063c4bffe2b11610097578063db6327dc11610071578063db6327dc146104b5578063dc0bd971146104c8578063e0351e13146104ee578063f2fde38b1461051457600080fd5b8063c4bffe2b1461047a578063c75eea9c1461048f578063cf7401f3146104a257600080fd5b8063af58d59f116100d3578063af58d59f146103c7578063b0f479a114610436578063b794658014610454578063c0d786551461046757600080fd5b80638da5cb5b146103745780639a4575b914610392578063a7cd63b7146103b257600080fd5b8063390775371161016657806378a010b21161014057806378a010b21461033357806379ba5097146103465780637d54534e1461034e5780638926f54f1461036157600080fd5b806339077537146102de57806354c8a4f3146103005780636d3d1a581461031557600080fd5b806321df0da71161019757806321df0da714610242578063240028e81461028957806332a7a822146102d657600080fd5b806301ffc9a7146101be5780630a2fd493146101e6578063181f5a7714610206575b600080fd5b6101d16101cc3660046120ad565b610527565b60405190151581526020015b60405180910390f35b6101f96101f436600461210c565b61060c565b6040516101dd919061218b565b6101f96040518060400160405280601781526020017f4d6f636b4c425443546f6b656e506f6f6c20312e352e3100000000000000000081525081565b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101dd565b6101d161029736600461219e565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff90811691161490565b6101f96106bc565b6102f16102ec3660046121d4565b61074a565b604051905181526020016101dd565b61031361030e36600461225c565b6108bf565b005b60085473ffffffffffffffffffffffffffffffffffffffff16610264565b6103136103413660046122c8565b61093a565b610313610aae565b61031361035c36600461219e565b610bab565b6101d161036f36600461210c565b610bfa565b60005473ffffffffffffffffffffffffffffffffffffffff16610264565b6103a56103a036600461234b565b610c11565b6040516101dd9190612386565b6103ba610dbc565b6040516101dd91906123e6565b6103da6103d536600461210c565b610dcd565b6040516101dd919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b60045473ffffffffffffffffffffffffffffffffffffffff16610264565b6101f961046236600461210c565b610ea2565b61031361047536600461219e565b610ecd565b610482610fa8565b6040516101dd9190612440565b6103da61049d36600461210c565b611060565b6103136104b0366004612579565b611132565b6103136104c33660046125be565b6111bb565b7f0000000000000000000000000000000000000000000000000000000000000000610264565b7f00000000000000000000000000000000000000000000000000000000000000006101d1565b61031361052236600461219e565b611641565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167faff2afbf0000000000000000000000000000000000000000000000000000000014806105ba57507fffffffff0000000000000000000000000000000000000000000000000000000082167f0e64dd2900000000000000000000000000000000000000000000000000000000145b8061060657507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b67ffffffffffffffff8116600090815260076020526040902060040180546060919061063790612600565b80601f016020809104026020016040519081016040528092919081815260200182805461066390612600565b80156106b05780601f10610685576101008083540402835291602001916106b0565b820191906000526020600020905b81548152906001019060200180831161069357829003601f168201915b50505050509050919050565b600980546106c990612600565b80601f01602080910402602001604051908101604052809291908181526020018280546106f590612600565b80156107425780601f1061071757610100808354040283529160200191610742565b820191906000526020600020905b81548152906001019060200180831161072557829003601f168201915b505050505081565b60408051602081019091526000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166340c10f196107a5606085016040860161219e565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff909116600482015260608501356024820152604401600060405180830381600087803b15801561081557600080fd5b505af1158015610829573d6000803e3d6000fd5b5061083e92505050606083016040840161219e565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f084606001356040516108a091815260200190565b60405180910390a3506040805160208101909152606090910135815290565b6108c7611655565b610934848480806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250506040805160208088028281018201909352878252909350879250869182918501908490808284376000920191909152506116d892505050565b50505050565b610942611655565b61094b83610bfa565b610992576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024015b60405180910390fd5b67ffffffffffffffff8316600090815260076020526040812060040180546109b990612600565b80601f01602080910402602001604051908101604052809291908181526020018280546109e590612600565b8015610a325780601f10610a0757610100808354040283529160200191610a32565b820191906000526020600020905b815481529060010190602001808311610a1557829003601f168201915b5050505067ffffffffffffffff8616600090815260076020526040902091925050600401610a618385836126a3565b508367ffffffffffffffff167fdb4d6220746a38cbc5335f7e108f7de80f482f4d23350253dfd0917df75a14bf828585604051610aa0939291906127be565b60405180910390a250505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610b2f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610989565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610bb3611655565b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6000610606600567ffffffffffffffff841661188e565b60408051808201909152606080825260208201526040517f42966c68000000000000000000000000000000000000000000000000000000008152606083013560048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906342966c6890602401600060405180830381600087803b158015610cb057600080fd5b505af1158015610cc4573d6000803e3d6000fd5b5050604051606085013581523392507f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7915060200160405180910390a26040518060400160405280610d22846020016020810190610462919061210c565b815260200160098054610d3490612600565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6090612600565b8015610dad5780601f10610d8257610100808354040283529160200191610dad565b820191906000526020600020905b815481529060010190602001808311610d9057829003601f168201915b50505050508152509050919050565b6060610dc860026118a9565b905090565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845260028201546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff161515948201949094526003909101548084166060830152919091049091166080820152610606906118b6565b67ffffffffffffffff8116600090815260076020526040902060050180546060919061063790612600565b610ed5611655565b73ffffffffffffffffffffffffffffffffffffffff8116610f22576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684910160405180910390a15050565b60606000610fb660056118a9565b90506000815167ffffffffffffffff811115610fd457610fd4612482565b604051908082528060200260200182016040528015610ffd578160200160208202803683370190505b50905060005b82518110156110595782818151811061101e5761101e612822565b602002602001015182828151811061103857611038612822565b67ffffffffffffffff90921660209283029190910190910152600101611003565b5092915050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845281546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff161515948201949094526001909101548084166060830152919091049091166080820152610606906118b6565b60085473ffffffffffffffffffffffffffffffffffffffff163314801590611172575060005473ffffffffffffffffffffffffffffffffffffffff163314155b156111ab576040517f8e4a23d6000000000000000000000000000000000000000000000000000000008152336004820152602401610989565b6111b6838383611968565b505050565b6111c3611655565b60005b818110156111b65760008383838181106111e2576111e2612822565b90506020028101906111f49190612851565b6111fd9061293a565b90506112128160800151826020015115611a52565b6112258160a00151826020015115611a52565b8060200151156115215780516112479060059067ffffffffffffffff16611b8f565b61128c5780516040517f1d5ad3c500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610989565b60408101515115806112a15750606081015151155b156112d8576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805161012081018252608083810180516020908101516fffffffffffffffffffffffffffffffff9081168486019081524263ffffffff90811660a0808901829052865151151560c08a01528651860151851660e08a015295518901518416610100890152918752875180860189529489018051850151841686528585019290925281515115158589015281518401518316606080870191909152915188015183168587015283870194855288880151878901908152828a015183890152895167ffffffffffffffff1660009081526007865289902088518051825482890151838e01519289167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316177001000000000000000000000000000000009188168202177fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff90811674010000000000000000000000000000000000000000941515850217865584890151948d0151948a16948a168202949094176001860155995180516002860180549b8301519f830151918b169b9093169a909a179d9096168a029c909c179091169615150295909517909855908101519401519381169316909102919091176003820155915190919060048201906114b990826129ee565b50606082015160058201906114ce90826129ee565b505081516060830151608084015160a08501516040517f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c295506115149493929190612b08565b60405180910390a1611638565b80516115399060059067ffffffffffffffff16611b9b565b61157e5780516040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610989565b805167ffffffffffffffff16600090815260076020526040812080547fffffffffffffffffffffff000000000000000000000000000000000000000000908116825560018201839055600282018054909116905560038101829055906115e7600483018261205f565b6115f560058301600061205f565b5050805160405167ffffffffffffffff90911681527f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599169060200160405180910390a15b506001016111c6565b611649611655565b61165281611ba7565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146116d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610989565b565b7f000000000000000000000000000000000000000000000000000000000000000061172f576040517f35f4a7b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b82518110156117c557600083828151811061174f5761174f612822565b6020026020010151905061176d816002611c9c90919063ffffffff16565b156117bc5760405173ffffffffffffffffffffffffffffffffffffffff821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b50600101611732565b5060005b81518110156111b65760008282815181106117e6576117e6612822565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361182a5750611886565b611835600282611cbe565b156118845760405173ffffffffffffffffffffffffffffffffffffffff821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b6001016117c9565b600081815260018301602052604081205415155b9392505050565b606060006118a283611ce0565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261194482606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff16426119289190612bd0565b85608001516fffffffffffffffffffffffffffffffff16611d3b565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b61197183610bfa565b6119b3576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602401610989565b6119be826000611a52565b67ffffffffffffffff831660009081526007602052604090206119e19083611d65565b6119ec816000611a52565b67ffffffffffffffff83166000908152600760205260409020611a129060020182611d65565b7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b838383604051611a4593929190612be3565b60405180910390a1505050565b815115611b1d5781602001516fffffffffffffffffffffffffffffffff1682604001516fffffffffffffffffffffffffffffffff16101580611aa8575060408201516fffffffffffffffffffffffffffffffff16155b15611ae157816040517f8020d1240000000000000000000000000000000000000000000000000000000081526004016109899190612c66565b8015611b19576040517f433fc33d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b60408201516fffffffffffffffffffffffffffffffff16151580611b56575060208201516fffffffffffffffffffffffffffffffff1615155b15611b1957816040517fd68af9cc0000000000000000000000000000000000000000000000000000000081526004016109899190612c66565b60006118a28383611f07565b60006118a28383611f56565b3373ffffffffffffffffffffffffffffffffffffffff821603611c26576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610989565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006118a28373ffffffffffffffffffffffffffffffffffffffff8416611f56565b60006118a28373ffffffffffffffffffffffffffffffffffffffff8416611f07565b6060816000018054806020026020016040519081016040528092919081815260200182805480156106b057602002820191906000526020600020905b815481526020019060010190808311611d1c5750505050509050919050565b6000611d5a85611d4b8486612ca2565b611d559087612cb9565b612049565b90505b949350505050565b8154600090611d8e90700100000000000000000000000000000000900463ffffffff1642612bd0565b90508015611e305760018301548354611dd6916fffffffffffffffffffffffffffffffff80821692811691859170010000000000000000000000000000000090910416611d3b565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b60208201518354611e56916fffffffffffffffffffffffffffffffff9081169116612049565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1990611a45908490612c66565b6000818152600183016020526040812054611f4e57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610606565b506000610606565b6000818152600183016020526040812054801561203f576000611f7a600183612bd0565b8554909150600090611f8e90600190612bd0565b9050808214611ff3576000866000018281548110611fae57611fae612822565b9060005260206000200154905080876000018481548110611fd157611fd1612822565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061200457612004612ccc565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610606565b6000915050610606565b600081831061205857816118a2565b5090919050565b50805461206b90612600565b6000825580601f1061207b575050565b601f01602090049060005260206000209081019061165291905b808211156120a95760008155600101612095565b5090565b6000602082840312156120bf57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146118a257600080fd5b803567ffffffffffffffff8116811461210757600080fd5b919050565b60006020828403121561211e57600080fd5b6118a2826120ef565b6000815180845260005b8181101561214d57602081850181015186830182015201612131565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006118a26020830184612127565b6000602082840312156121b057600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146118a257600080fd5b6000602082840312156121e657600080fd5b813567ffffffffffffffff8111156121fd57600080fd5b820161010081850312156118a257600080fd5b60008083601f84011261222257600080fd5b50813567ffffffffffffffff81111561223a57600080fd5b6020830191508360208260051b850101111561225557600080fd5b9250929050565b6000806000806040858703121561227257600080fd5b843567ffffffffffffffff8082111561228a57600080fd5b61229688838901612210565b909650945060208701359150808211156122af57600080fd5b506122bc87828801612210565b95989497509550505050565b6000806000604084860312156122dd57600080fd5b6122e6846120ef565b9250602084013567ffffffffffffffff8082111561230357600080fd5b818601915086601f83011261231757600080fd5b81358181111561232657600080fd5b87602082850101111561233857600080fd5b6020830194508093505050509250925092565b60006020828403121561235d57600080fd5b813567ffffffffffffffff81111561237457600080fd5b820160a081850312156118a257600080fd5b6020815260008251604060208401526123a26060840182612127565b905060208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160408501526123dd8282612127565b95945050505050565b6020808252825182820181905260009190848201906040850190845b8181101561243457835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101612402565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561243457835167ffffffffffffffff168352928401929184019160010161245c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160c0810167ffffffffffffffff811182821017156124d4576124d4612482565b60405290565b8035801515811461210757600080fd5b80356fffffffffffffffffffffffffffffffff8116811461210757600080fd5b60006060828403121561251c57600080fd5b6040516060810181811067ffffffffffffffff8211171561253f5761253f612482565b60405290508061254e836124da565b815261255c602084016124ea565b602082015261256d604084016124ea565b60408201525092915050565b600080600060e0848603121561258e57600080fd5b612597846120ef565b92506125a6856020860161250a565b91506125b5856080860161250a565b90509250925092565b600080602083850312156125d157600080fd5b823567ffffffffffffffff8111156125e857600080fd5b6125f485828601612210565b90969095509350505050565b600181811c9082168061261457607f821691505b60208210810361264d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f8211156111b6576000816000526020600020601f850160051c8101602086101561267c5750805b601f850160051c820191505b8181101561269b57828155600101612688565b505050505050565b67ffffffffffffffff8311156126bb576126bb612482565b6126cf836126c98354612600565b83612653565b6000601f84116001811461272157600085156126eb5750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b1783556127b7565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b828110156127705786850135825560209485019460019092019101612750565b50868210156127ab577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b6040815260006127d16040830186612127565b82810360208401528381528385602083013760006020858301015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f860116820101915050949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec183360301811261288557600080fd5b9190910192915050565b600082601f8301126128a057600080fd5b813567ffffffffffffffff808211156128bb576128bb612482565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561290157612901612482565b8160405283815286602085880101111561291a57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000610140823603121561294d57600080fd5b6129556124b1565b61295e836120ef565b815261296c602084016124da565b6020820152604083013567ffffffffffffffff8082111561298c57600080fd5b6129983683870161288f565b604084015260608501359150808211156129b157600080fd5b506129be3682860161288f565b6060830152506129d1366080850161250a565b60808201526129e33660e0850161250a565b60a082015292915050565b815167ffffffffffffffff811115612a0857612a08612482565b612a1c81612a168454612600565b84612653565b602080601f831160018114612a6f5760008415612a395750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b17855561269b565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015612abc57888601518255948401946001909101908401612a9d565b5085821015612af857878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b600061010067ffffffffffffffff87168352806020840152612b2c81840187612127565b8551151560408581019190915260208701516fffffffffffffffffffffffffffffffff9081166060870152908701511660808501529150612b6a9050565b8251151560a083015260208301516fffffffffffffffffffffffffffffffff90811660c084015260408401511660e08301526123dd565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561060657610606612ba1565b67ffffffffffffffff8416815260e08101612c2f60208301858051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b82511515608083015260208301516fffffffffffffffffffffffffffffffff90811660a084015260408401511660c0830152611d5d565b6060810161060682848051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b808202811582820484141761060657610606612ba1565b8082018082111561060657610606612ba1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000818000a",
}

var MockLBTCTokenPoolABI = MockLBTCTokenPoolMetaData.ABI

var MockLBTCTokenPoolBin = MockLBTCTokenPoolMetaData.Bin

func DeployMockLBTCTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, allowlist []common.Address, rmnProxy common.Address, router common.Address, destPoolData []byte) (common.Address, *types.Transaction, *MockLBTCTokenPool, error) {
	parsed, err := MockLBTCTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockLBTCTokenPoolBin), backend, token, allowlist, rmnProxy, router, destPoolData)
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) GetRemotePool(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "getRemotePool", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) GetRemotePool(remoteChainSelector uint64) ([]byte, error) {
	return _MockLBTCTokenPool.Contract.GetRemotePool(&_MockLBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) GetRemotePool(remoteChainSelector uint64) ([]byte, error) {
	return _MockLBTCTokenPool.Contract.GetRemotePool(&_MockLBTCTokenPool.CallOpts, remoteChainSelector)
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolCaller) IDestPoolData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _MockLBTCTokenPool.contract.Call(opts, &out, "i_destPoolData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) IDestPoolData() ([]byte, error) {
	return _MockLBTCTokenPool.Contract.IDestPoolData(&_MockLBTCTokenPool.CallOpts)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolCallerSession) IDestPoolData() ([]byte, error) {
	return _MockLBTCTokenPool.Contract.IDestPoolData(&_MockLBTCTokenPool.CallOpts)
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.ApplyAllowListUpdates(&_MockLBTCTokenPool.TransactOpts, removes, adds)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.ApplyAllowListUpdates(&_MockLBTCTokenPool.TransactOpts, removes, adds)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _MockLBTCTokenPool.contract.Transact(opts, "applyChainUpdates", chains)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) ApplyChainUpdates(chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.ApplyChainUpdates(&_MockLBTCTokenPool.TransactOpts, chains)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorSession) ApplyChainUpdates(chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.ApplyChainUpdates(&_MockLBTCTokenPool.TransactOpts, chains)
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactor) SetRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockLBTCTokenPool.contract.Transact(opts, "setRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolSession) SetRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.SetRemotePool(&_MockLBTCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolTransactorSession) SetRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockLBTCTokenPool.Contract.SetRemotePool(&_MockLBTCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
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

type MockLBTCTokenPoolRemotePoolSetIterator struct {
	Event *MockLBTCTokenPoolRemotePoolSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockLBTCTokenPoolRemotePoolSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockLBTCTokenPoolRemotePoolSet)
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
		it.Event = new(MockLBTCTokenPoolRemotePoolSet)
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

func (it *MockLBTCTokenPoolRemotePoolSetIterator) Error() error {
	return it.fail
}

func (it *MockLBTCTokenPoolRemotePoolSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockLBTCTokenPoolRemotePoolSet struct {
	RemoteChainSelector uint64
	PreviousPoolAddress []byte
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) FilterRemotePoolSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockLBTCTokenPoolRemotePoolSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.FilterLogs(opts, "RemotePoolSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockLBTCTokenPoolRemotePoolSetIterator{contract: _MockLBTCTokenPool.contract, event: "RemotePoolSet", logs: logs, sub: sub}, nil
}

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) WatchRemotePoolSet(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolRemotePoolSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockLBTCTokenPool.contract.WatchLogs(opts, "RemotePoolSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockLBTCTokenPoolRemotePoolSet)
				if err := _MockLBTCTokenPool.contract.UnpackLog(event, "RemotePoolSet", log); err != nil {
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

func (_MockLBTCTokenPool *MockLBTCTokenPoolFilterer) ParseRemotePoolSet(log types.Log) (*MockLBTCTokenPoolRemotePoolSet, error) {
	event := new(MockLBTCTokenPoolRemotePoolSet)
	if err := _MockLBTCTokenPool.contract.UnpackLog(event, "RemotePoolSet", log); err != nil {
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
	case _MockLBTCTokenPool.abi.Events["Released"].ID:
		return _MockLBTCTokenPool.ParseReleased(log)
	case _MockLBTCTokenPool.abi.Events["RemotePoolSet"].ID:
		return _MockLBTCTokenPool.ParseRemotePoolSet(log)
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

func (MockLBTCTokenPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (MockLBTCTokenPoolRemotePoolSet) Topic() common.Hash {
	return common.HexToHash("0xdb4d6220746a38cbc5335f7e108f7de80f482f4d23350253dfd0917df75a14bf")
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

	GetRemotePool(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	IDestPoolData(opts *bind.CallOpts) ([]byte, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, chains []TokenPoolChainUpdate) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

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

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*MockLBTCTokenPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*MockLBTCTokenPoolReleased, error)

	FilterRemotePoolSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockLBTCTokenPoolRemotePoolSetIterator, error)

	WatchRemotePoolSet(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolRemotePoolSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolSet(log types.Log) (*MockLBTCTokenPoolRemotePoolSet, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*MockLBTCTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *MockLBTCTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*MockLBTCTokenPoolRouterUpdated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
