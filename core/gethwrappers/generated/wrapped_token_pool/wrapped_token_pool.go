// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrapped_token_pool

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

type IPoolRampUpdate struct {
	Ramp    common.Address
	Allowed bool
}

type RateLimiterConfig struct {
	IsEnabled bool
	Rate      *big.Int
	Capacity  *big.Int
}

type RateLimiterTokenBucket struct {
	Capacity    *big.Int
	Tokens      *big.Int
	Rate        *big.Int
	LastUpdated *big.Int
	IsEnabled   bool
}

var WrappedTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ConsumingMoreThanMaxCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ExceedsTokenLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"RateLimitReached\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"OffRampAllowanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"OnRampAllowanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"ramp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"internalType\":\"structIPool.RampUpdate[]\",\"name\":\"onRamps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ramp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"internalType\":\"structIPool.RampUpdate[]\",\"name\":\"offRamps\",\"type\":\"tuple[]\"}],\"name\":\"applyRampUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdated\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"isOnRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockOrBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"releaseOrMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162002a4038038062002a40833981016040819052620000349162000390565b838383828230863380600081620000925760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c557620000c581620001dc565b50506001805460ff60a01b19169055506001600160a01b038216620000fd57604051634655efd160e11b815260040160405180910390fd5b6040805160a0810182528282018051808352905160208084018290528501516001600160d01b03169383018490524264ffffffffff1660608401819052945115156080938401819052600692909255600755600880547fff0000000000000000000000000000000000000000000000000000000000000016909317600160d01b909402939093176001600160f81b0316600160f81b9093029290921790556001600160a01b03919091169052600c620001b7838262000508565b50600d620001c6828262000508565b50505060ff1660a05250620005d4945050505050565b336001600160a01b03821603620002365760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000089565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715620002c257620002c262000287565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620002f357620002f362000287565b604052919050565b600082601f8301126200030d57600080fd5b81516001600160401b0381111562000329576200032962000287565b60206200033f601f8301601f19168201620002c8565b82815285828487010111156200035457600080fd5b60005b838110156200037457858101830151828201840152820162000357565b83811115620003865760008385840101525b5095945050505050565b60008060008084860360c0811215620003a857600080fd5b85516001600160401b0380821115620003c057600080fd5b620003ce89838a01620002fb565b96506020880151915080821115620003e557600080fd5b50620003f488828901620002fb565b945050604086015160ff811681146200040c57600080fd5b92506060605f19820112156200042157600080fd5b506200042c6200029d565b606086015180151581146200044057600080fd5b815260808601516001600160d01b03811681146200045d57600080fd5b602082015260a095909501516040860152509194909350909190565b600181811c908216806200048e57607f821691505b602082108103620004af57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200050357600081815260208120601f850160051c81016020861015620004de5750805b601f850160051c820191505b81811015620004ff57828155600101620004ea565b5050505b505050565b81516001600160401b0381111562000524576200052462000287565b6200053c8162000535845462000479565b84620004b5565b602080601f8311600181146200057457600084156200055b5750858301515b600019600386901b1c1916600185901b178555620004ff565b600085815260208120601f198616915b82811015620005a55788860151825594840194600190910190840162000584565b5085821015620005c45787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a051612446620005fa60003960006102840152600061021001526124466000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c806370a08231116100ee578063a9059cbb11610097578063dd62ed3e11610071578063dd62ed3e14610413578063e2e59b3e14610459578063ea6192a21461046c578063f2fde38b1461047f57600080fd5b8063a9059cbb1461037e578063af51911214610391578063c38945a7146103a457600080fd5b80638da5cb5b116100c85780638da5cb5b1461034557806395d89b4114610363578063a457c2d71461036b57600080fd5b806370a08231146102ff57806379ba5097146103355780638456cb591461033d57600080fd5b80633091aee7116101505780633f4ba83a1161012a5780633f4ba83a146102c15780635c975abb146102c95780636f32b872146102ec57600080fd5b80633091aee714610268578063313ce5671461027d57806339509351146102ae57600080fd5b80631d7a74a0116101815780631d7a74a0146101fb57806321df0da71461020e57806323b872dd1461025557600080fd5b806306fdde03146101a8578063095ea7b3146101c657806318160ddd146101e9575b600080fd5b6101b0610492565b6040516101bd9190611eb2565b60405180910390f35b6101d96101d4366004611f4e565b610524565b60405190151581526020016101bd565b600b545b6040519081526020016101bd565b6101d9610209366004611f78565b61053b565b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101bd565b6101d9610263366004611f93565b610548565b61027b610276366004612086565b610633565b005b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101bd565b6101d96102bc366004611f4e565b610649565b61027b610692565b60015474010000000000000000000000000000000000000000900460ff166101d9565b6101d96102fa366004611f78565b6106a4565b6101ed61030d366004611f78565b73ffffffffffffffffffffffffffffffffffffffff1660009081526009602052604090205490565b61027b6106b1565b61027b6107ae565b60005473ffffffffffffffffffffffffffffffffffffffff16610230565b6101b06107be565b6101d9610379366004611f4e565b6107cd565b6101d961038c366004611f4e565b6108a5565b61027b61039f3660046121c0565b6108b2565b6103ac610ac2565b6040516101bd9190600060a082019050825182526020830151602083015279ffffffffffffffffffffffffffffffffffffffffffffffffffff604084015116604083015264ffffffffff606084015116606083015260808301511515608083015292915050565b6101ed610421366004612224565b73ffffffffffffffffffffffffffffffffffffffff9182166000908152600a6020908152604080832093909416825291909152205490565b61027b610467366004612257565b610b91565b61027b61047a366004611f4e565b610c98565b61027b61048d366004611f78565b610dc1565b6060600c80546104a19061227a565b80601f01602080910402602001604051908101604052809291908181526020018280546104cd9061227a565b801561051a5780601f106104ef5761010080835404028352916020019161051a565b820191906000526020600020905b8154815290600101906020018083116104fd57829003601f168201915b5050505050905090565b6000610531338484610dd2565b5060015b92915050565b6000610535600483610f85565b6000610555848484610fb7565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600a602090815260408083203384529091529020548281101561061b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206160448201527f6c6c6f77616e636500000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6106288533858403610dd2565b506001949350505050565b61063b61126b565b6106466006826112ec565b50565b336000818152600a6020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054909161053191859061068d9086906122fc565b610dd2565b61069a61126b565b6106a261141d565b565b6000610535600283610f85565b60015473ffffffffffffffffffffffffffffffffffffffff163314610732576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610612565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6107b661126b565b6106a2611516565b6060600d80546104a19061227a565b336000908152600a6020908152604080832073ffffffffffffffffffffffffffffffffffffffff861684529091528120548281101561088e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610612565b61089b3385858403610dd2565b5060019392505050565b6000610531338484610fb7565b6108ba61126b565b60005b82518110156109bb5760008382815181106108da576108da612314565b6020026020010151905080602001516109005780516108fb90600290611602565b61090e565b805161090e90600290611624565b156109aa577fbceff8f229c6dfcbf8bdcfb18726b84b0fd249b4803deb3948ff34d90401366284838151811061094657610946612314565b60200260200101516000015185848151811061096457610964612314565b6020026020010151602001516040516109a192919073ffffffffffffffffffffffffffffffffffffffff9290921682521515602082015260400190565b60405180910390a15b506109b481612343565b90506108bd565b5060005b8151811015610abd5760008282815181106109dc576109dc612314565b602002602001015190508060200151610a025780516109fd90600490611602565b610a10565b8051610a1090600490611624565b15610aac577fd8c3333ded377884ced3869cd0bcb9be54ea664076df1f5d39c4689120313648838381518110610a4857610a48612314565b602002602001015160000151848481518110610a6657610a66612314565b602002602001015160200151604051610aa392919073ffffffffffffffffffffffffffffffffffffffff9290921682521515602082015260400190565b60405180910390a15b50610ab681612343565b90506109bf565b505050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526006548152600754602082015260085479ffffffffffffffffffffffffffffffffffffffffffffffffffff8116928201929092527a010000000000000000000000000000000000000000000000000000820464ffffffffff1660608201527f010000000000000000000000000000000000000000000000000000000000000090910460ff1615156080820152610b8c90611646565b905090565b60015474010000000000000000000000000000000000000000900460ff1615610c16576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610612565b610c1f336106a4565b610c55576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c5f30836116e9565b60405182815233907f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df79060200160405180910390a25050565b60015474010000000000000000000000000000000000000000900460ff1615610d1d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610612565b610d263361053b565b610d5c576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d65816118d6565b610d6f82826118e1565b60405181815273ffffffffffffffffffffffffffffffffffffffff83169033907f9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0906020015b60405180910390a35050565b610dc961126b565b610646816119fa565b73ffffffffffffffffffffffffffffffffffffffff8316610e74576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610612565b73ffffffffffffffffffffffffffffffffffffffff8216610f17576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610612565b73ffffffffffffffffffffffffffffffffffffffff8381166000818152600a602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018301602052604081205415155b9392505050565b73ffffffffffffffffffffffffffffffffffffffff831661105a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610612565b73ffffffffffffffffffffffffffffffffffffffff82166110fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610612565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260096020526040902054818110156111b3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610612565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600960205260408082208585039055918516815290812080548492906111f79084906122fc565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161125d91815260200190565b60405180910390a350505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146106a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610612565b6112f582611aef565b8051600283018054604084015180865560208501517effffffffff00000000000000000000000000000000000000000000000000009092167f0100000000000000000000000000000000000000000000000000000000000000941515949094027fffffffffffff0000000000000000000000000000000000000000000000000000169390931779ffffffffffffffffffffffffffffffffffffffffffffffffffff90911617905560018301546113ab9190611c04565b60018301556040805182511515815260208084015179ffffffffffffffffffffffffffffffffffffffffffffffffffff169082015282820151918101919091527f44a2350342338075ac038f37b8d9e49e696e360492cb44cc6bc37fc117f19df8906060015b60405180910390a15050565b60015474010000000000000000000000000000000000000000900460ff166114a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610612565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60015474010000000000000000000000000000000000000000900460ff161561159b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610612565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586114ec3390565b6000610fb08373ffffffffffffffffffffffffffffffffffffffff8416611c1a565b6000610fb08373ffffffffffffffffffffffffffffffffffffffff8416611d0d565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526000826060015164ffffffffff164261168a919061237b565b835160408501519192506116d2916116be9079ffffffffffffffffffffffffffffffffffffffffffffffffffff1684612392565b85602001516116cd91906122fc565b611c04565b6020840152505064ffffffffff4216606082015290565b73ffffffffffffffffffffffffffffffffffffffff821661178c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610612565b73ffffffffffffffffffffffffffffffffffffffff821660009081526009602052604090205481811015611842576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610612565b73ffffffffffffffffffffffffffffffffffffffff831660009081526009602052604081208383039055600b805484929061187e90849061237b565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a3505050565b610646600682611d5c565b73ffffffffffffffffffffffffffffffffffffffff821661195e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610612565b80600b600082825461197091906122fc565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260096020526040812080548392906119aa9084906122fc565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610db5565b3373ffffffffffffffffffffffffffffffffffffffff821603611a79576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610612565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b805460018201541480611b2b5750600281015464ffffffffff7a0100000000000000000000000000000000000000000000000000009091041642145b15611b335750565b6002810154600090611b6a907a010000000000000000000000000000000000000000000000000000900464ffffffffff164261237b565b82546002840154919250611bad91611b9e9079ffffffffffffffffffffffffffffffffffffffffffffffffffff1684612392565b84600101546116cd91906122fc565b60018301555060020180547fff0000000000ffffffffffffffffffffffffffffffffffffffffffffffffffff167a0100000000000000000000000000000000000000000000000000004264ffffffffff1602179055565b6000818310611c135781610fb0565b5090919050565b60008181526001830160205260408120548015611d03576000611c3e60018361237b565b8554909150600090611c529060019061237b565b9050818114611cb7576000866000018281548110611c7257611c72612314565b9060005260206000200154905080876000018481548110611c9557611c95612314565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080611cc857611cc86123cf565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610535565b6000915050610535565b6000818152600183016020526040812054611d5457508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610535565b506000610535565b60028201547f0100000000000000000000000000000000000000000000000000000000000000900460ff161580611d91575080155b15611d9a575050565b611da382611aef565b8154811115611deb5781546040517f48369c43000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610612565b8082600101541015611e69576002820154600183015479ffffffffffffffffffffffffffffffffffffffffffffffffffff90911690611e2a908361237b565b611e3491906123fe565b6040517fdc96cefa00000000000000000000000000000000000000000000000000000000815260040161061291815260200190565b80826001016000828254611e7d919061237b565b90915550506040518181527f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a90602001611411565b600060208083528351808285015260005b81811015611edf57858101830151858201604001528201611ec3565b81811115611ef1576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611f4957600080fd5b919050565b60008060408385031215611f6157600080fd5b611f6a83611f25565b946020939093013593505050565b600060208284031215611f8a57600080fd5b610fb082611f25565b600080600060608486031215611fa857600080fd5b611fb184611f25565b9250611fbf60208501611f25565b9150604084013590509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561202157612021611fcf565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561206e5761206e611fcf565b604052919050565b80358015158114611f4957600080fd5b60006060828403121561209857600080fd5b6040516060810181811067ffffffffffffffff821117156120bb576120bb611fcf565b6040526120c783612076565b8152602083013579ffffffffffffffffffffffffffffffffffffffffffffffffffff811681146120f657600080fd5b60208201526040928301359281019290925250919050565b600082601f83011261211f57600080fd5b8135602067ffffffffffffffff82111561213b5761213b611fcf565b612149818360051b01612027565b82815260069290921b8401810191818101908684111561216857600080fd5b8286015b848110156121b557604081890312156121855760008081fd5b61218d611ffe565b61219682611f25565b81526121a3858301612076565b8186015283529183019160400161216c565b509695505050505050565b600080604083850312156121d357600080fd5b823567ffffffffffffffff808211156121eb57600080fd5b6121f78683870161210e565b9350602085013591508082111561220d57600080fd5b5061221a8582860161210e565b9150509250929050565b6000806040838503121561223757600080fd5b61224083611f25565b915061224e60208401611f25565b90509250929050565b6000806040838503121561226a57600080fd5b8235915061224e60208401611f25565b600181811c9082168061228e57607f821691505b6020821081036122c7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561230f5761230f6122cd565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612374576123746122cd565b5060010190565b60008282101561238d5761238d6122cd565b500390565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156123ca576123ca6122cd565b500290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600082612434577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea164736f6c634300080f000a",
}

var WrappedTokenPoolABI = WrappedTokenPoolMetaData.ABI

var WrappedTokenPoolBin = WrappedTokenPoolMetaData.Bin

func DeployWrappedTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals uint8, rateLimiterConfig RateLimiterConfig) (common.Address, *types.Transaction, *WrappedTokenPool, error) {
	parsed, err := WrappedTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WrappedTokenPoolBin), backend, name, symbol, decimals, rateLimiterConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WrappedTokenPool{WrappedTokenPoolCaller: WrappedTokenPoolCaller{contract: contract}, WrappedTokenPoolTransactor: WrappedTokenPoolTransactor{contract: contract}, WrappedTokenPoolFilterer: WrappedTokenPoolFilterer{contract: contract}}, nil
}

type WrappedTokenPool struct {
	address common.Address
	abi     abi.ABI
	WrappedTokenPoolCaller
	WrappedTokenPoolTransactor
	WrappedTokenPoolFilterer
}

type WrappedTokenPoolCaller struct {
	contract *bind.BoundContract
}

type WrappedTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type WrappedTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type WrappedTokenPoolSession struct {
	Contract     *WrappedTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type WrappedTokenPoolCallerSession struct {
	Contract *WrappedTokenPoolCaller
	CallOpts bind.CallOpts
}

type WrappedTokenPoolTransactorSession struct {
	Contract     *WrappedTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type WrappedTokenPoolRaw struct {
	Contract *WrappedTokenPool
}

type WrappedTokenPoolCallerRaw struct {
	Contract *WrappedTokenPoolCaller
}

type WrappedTokenPoolTransactorRaw struct {
	Contract *WrappedTokenPoolTransactor
}

func NewWrappedTokenPool(address common.Address, backend bind.ContractBackend) (*WrappedTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(WrappedTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindWrappedTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPool{address: address, abi: abi, WrappedTokenPoolCaller: WrappedTokenPoolCaller{contract: contract}, WrappedTokenPoolTransactor: WrappedTokenPoolTransactor{contract: contract}, WrappedTokenPoolFilterer: WrappedTokenPoolFilterer{contract: contract}}, nil
}

func NewWrappedTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*WrappedTokenPoolCaller, error) {
	contract, err := bindWrappedTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolCaller{contract: contract}, nil
}

func NewWrappedTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*WrappedTokenPoolTransactor, error) {
	contract, err := bindWrappedTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolTransactor{contract: contract}, nil
}

func NewWrappedTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*WrappedTokenPoolFilterer, error) {
	contract, err := bindWrappedTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolFilterer{contract: contract}, nil
}

func bindWrappedTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WrappedTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_WrappedTokenPool *WrappedTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedTokenPool.Contract.WrappedTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_WrappedTokenPool *WrappedTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.WrappedTokenPoolTransactor.contract.Transfer(opts)
}

func (_WrappedTokenPool *WrappedTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.WrappedTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_WrappedTokenPool *WrappedTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.contract.Transfer(opts)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_WrappedTokenPool *WrappedTokenPoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedTokenPool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WrappedTokenPool *WrappedTokenPoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedTokenPool.Contract.Allowance(&_WrappedTokenPool.CallOpts, owner, spender)
}

func (_WrappedTokenPool *WrappedTokenPoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedTokenPool.Contract.Allowance(&_WrappedTokenPool.CallOpts, owner, spender)
}

func (_WrappedTokenPool *WrappedTokenPoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedTokenPool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WrappedTokenPool *WrappedTokenPoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedTokenPool.Contract.BalanceOf(&_WrappedTokenPool.CallOpts, account)
}

func (_WrappedTokenPool *WrappedTokenPoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedTokenPool.Contract.BalanceOf(&_WrappedTokenPool.CallOpts, account)
}

func (_WrappedTokenPool *WrappedTokenPoolCaller) CurrentTokenBucketState(opts *bind.CallOpts) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _WrappedTokenPool.contract.Call(opts, &out, "currentTokenBucketState")

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_WrappedTokenPool *WrappedTokenPoolSession) CurrentTokenBucketState() (RateLimiterTokenBucket, error) {
	return _WrappedTokenPool.Contract.CurrentTokenBucketState(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCallerSession) CurrentTokenBucketState() (RateLimiterTokenBucket, error) {
	return _WrappedTokenPool.Contract.CurrentTokenBucketState(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WrappedTokenPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_WrappedTokenPool *WrappedTokenPoolSession) Decimals() (uint8, error) {
	return _WrappedTokenPool.Contract.Decimals(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCallerSession) Decimals() (uint8, error) {
	return _WrappedTokenPool.Contract.Decimals(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrappedTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_WrappedTokenPool *WrappedTokenPoolSession) GetToken() (common.Address, error) {
	return _WrappedTokenPool.Contract.GetToken(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _WrappedTokenPool.Contract.GetToken(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCaller) IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _WrappedTokenPool.contract.Call(opts, &out, "isOffRamp", offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_WrappedTokenPool *WrappedTokenPoolSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _WrappedTokenPool.Contract.IsOffRamp(&_WrappedTokenPool.CallOpts, offRamp)
}

func (_WrappedTokenPool *WrappedTokenPoolCallerSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _WrappedTokenPool.Contract.IsOffRamp(&_WrappedTokenPool.CallOpts, offRamp)
}

func (_WrappedTokenPool *WrappedTokenPoolCaller) IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error) {
	var out []interface{}
	err := _WrappedTokenPool.contract.Call(opts, &out, "isOnRamp", onRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_WrappedTokenPool *WrappedTokenPoolSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _WrappedTokenPool.Contract.IsOnRamp(&_WrappedTokenPool.CallOpts, onRamp)
}

func (_WrappedTokenPool *WrappedTokenPoolCallerSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _WrappedTokenPool.Contract.IsOnRamp(&_WrappedTokenPool.CallOpts, onRamp)
}

func (_WrappedTokenPool *WrappedTokenPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedTokenPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_WrappedTokenPool *WrappedTokenPoolSession) Name() (string, error) {
	return _WrappedTokenPool.Contract.Name(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCallerSession) Name() (string, error) {
	return _WrappedTokenPool.Contract.Name(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrappedTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_WrappedTokenPool *WrappedTokenPoolSession) Owner() (common.Address, error) {
	return _WrappedTokenPool.Contract.Owner(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCallerSession) Owner() (common.Address, error) {
	return _WrappedTokenPool.Contract.Owner(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WrappedTokenPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_WrappedTokenPool *WrappedTokenPoolSession) Paused() (bool, error) {
	return _WrappedTokenPool.Contract.Paused(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCallerSession) Paused() (bool, error) {
	return _WrappedTokenPool.Contract.Paused(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedTokenPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_WrappedTokenPool *WrappedTokenPoolSession) Symbol() (string, error) {
	return _WrappedTokenPool.Contract.Symbol(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCallerSession) Symbol() (string, error) {
	return _WrappedTokenPool.Contract.Symbol(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrappedTokenPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WrappedTokenPool *WrappedTokenPoolSession) TotalSupply() (*big.Int, error) {
	return _WrappedTokenPool.Contract.TotalSupply(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _WrappedTokenPool.Contract.TotalSupply(&_WrappedTokenPool.CallOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_WrappedTokenPool *WrappedTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.AcceptOwnership(&_WrappedTokenPool.TransactOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.AcceptOwnership(&_WrappedTokenPool.TransactOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactor) ApplyRampUpdates(opts *bind.TransactOpts, onRamps []IPoolRampUpdate, offRamps []IPoolRampUpdate) (*types.Transaction, error) {
	return _WrappedTokenPool.contract.Transact(opts, "applyRampUpdates", onRamps, offRamps)
}

func (_WrappedTokenPool *WrappedTokenPoolSession) ApplyRampUpdates(onRamps []IPoolRampUpdate, offRamps []IPoolRampUpdate) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.ApplyRampUpdates(&_WrappedTokenPool.TransactOpts, onRamps, offRamps)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorSession) ApplyRampUpdates(onRamps []IPoolRampUpdate, offRamps []IPoolRampUpdate) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.ApplyRampUpdates(&_WrappedTokenPool.TransactOpts, onRamps, offRamps)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.contract.Transact(opts, "approve", spender, amount)
}

func (_WrappedTokenPool *WrappedTokenPoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.Approve(&_WrappedTokenPool.TransactOpts, spender, amount)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.Approve(&_WrappedTokenPool.TransactOpts, spender, amount)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

func (_WrappedTokenPool *WrappedTokenPoolSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.DecreaseAllowance(&_WrappedTokenPool.TransactOpts, spender, subtractedValue)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.DecreaseAllowance(&_WrappedTokenPool.TransactOpts, spender, subtractedValue)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

func (_WrappedTokenPool *WrappedTokenPoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.IncreaseAllowance(&_WrappedTokenPool.TransactOpts, spender, addedValue)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.IncreaseAllowance(&_WrappedTokenPool.TransactOpts, spender, addedValue)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, amount *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _WrappedTokenPool.contract.Transact(opts, "lockOrBurn", amount, arg1)
}

func (_WrappedTokenPool *WrappedTokenPoolSession) LockOrBurn(amount *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.LockOrBurn(&_WrappedTokenPool.TransactOpts, amount, arg1)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorSession) LockOrBurn(amount *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.LockOrBurn(&_WrappedTokenPool.TransactOpts, amount, arg1)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedTokenPool.contract.Transact(opts, "pause")
}

func (_WrappedTokenPool *WrappedTokenPoolSession) Pause() (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.Pause(&_WrappedTokenPool.TransactOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorSession) Pause() (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.Pause(&_WrappedTokenPool.TransactOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.contract.Transact(opts, "releaseOrMint", recipient, amount)
}

func (_WrappedTokenPool *WrappedTokenPoolSession) ReleaseOrMint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.ReleaseOrMint(&_WrappedTokenPool.TransactOpts, recipient, amount)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorSession) ReleaseOrMint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.ReleaseOrMint(&_WrappedTokenPool.TransactOpts, recipient, amount)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error) {
	return _WrappedTokenPool.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_WrappedTokenPool *WrappedTokenPoolSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.SetRateLimiterConfig(&_WrappedTokenPool.TransactOpts, config)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.SetRateLimiterConfig(&_WrappedTokenPool.TransactOpts, config)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.contract.Transact(opts, "transfer", recipient, amount)
}

func (_WrappedTokenPool *WrappedTokenPoolSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.Transfer(&_WrappedTokenPool.TransactOpts, recipient, amount)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.Transfer(&_WrappedTokenPool.TransactOpts, recipient, amount)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

func (_WrappedTokenPool *WrappedTokenPoolSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.TransferFrom(&_WrappedTokenPool.TransactOpts, sender, recipient, amount)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.TransferFrom(&_WrappedTokenPool.TransactOpts, sender, recipient, amount)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _WrappedTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_WrappedTokenPool *WrappedTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.TransferOwnership(&_WrappedTokenPool.TransactOpts, to)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.TransferOwnership(&_WrappedTokenPool.TransactOpts, to)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedTokenPool.contract.Transact(opts, "unpause")
}

func (_WrappedTokenPool *WrappedTokenPoolSession) Unpause() (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.Unpause(&_WrappedTokenPool.TransactOpts)
}

func (_WrappedTokenPool *WrappedTokenPoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _WrappedTokenPool.Contract.Unpause(&_WrappedTokenPool.TransactOpts)
}

type WrappedTokenPoolApprovalIterator struct {
	Event *WrappedTokenPoolApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrappedTokenPoolApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenPoolApproval)
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
		it.Event = new(WrappedTokenPoolApproval)
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

func (it *WrappedTokenPoolApprovalIterator) Error() error {
	return it.fail
}

func (it *WrappedTokenPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrappedTokenPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WrappedTokenPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolApprovalIterator{contract: _WrappedTokenPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrappedTokenPoolApproval)
				if err := _WrappedTokenPool.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_WrappedTokenPool *WrappedTokenPoolFilterer) ParseApproval(log types.Log) (*WrappedTokenPoolApproval, error) {
	event := new(WrappedTokenPoolApproval)
	if err := _WrappedTokenPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrappedTokenPoolBurnedIterator struct {
	Event *WrappedTokenPoolBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrappedTokenPoolBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenPoolBurned)
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
		it.Event = new(WrappedTokenPoolBurned)
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

func (it *WrappedTokenPoolBurnedIterator) Error() error {
	return it.fail
}

func (it *WrappedTokenPoolBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrappedTokenPoolBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*WrappedTokenPoolBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolBurnedIterator{contract: _WrappedTokenPool.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrappedTokenPoolBurned)
				if err := _WrappedTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
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

func (_WrappedTokenPool *WrappedTokenPoolFilterer) ParseBurned(log types.Log) (*WrappedTokenPoolBurned, error) {
	event := new(WrappedTokenPoolBurned)
	if err := _WrappedTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrappedTokenPoolLockedIterator struct {
	Event *WrappedTokenPoolLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrappedTokenPoolLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenPoolLocked)
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
		it.Event = new(WrappedTokenPoolLocked)
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

func (it *WrappedTokenPoolLockedIterator) Error() error {
	return it.fail
}

func (it *WrappedTokenPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrappedTokenPoolLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*WrappedTokenPoolLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolLockedIterator{contract: _WrappedTokenPool.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrappedTokenPoolLocked)
				if err := _WrappedTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
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

func (_WrappedTokenPool *WrappedTokenPoolFilterer) ParseLocked(log types.Log) (*WrappedTokenPoolLocked, error) {
	event := new(WrappedTokenPoolLocked)
	if err := _WrappedTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrappedTokenPoolMintedIterator struct {
	Event *WrappedTokenPoolMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrappedTokenPoolMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenPoolMinted)
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
		it.Event = new(WrappedTokenPoolMinted)
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

func (it *WrappedTokenPoolMintedIterator) Error() error {
	return it.fail
}

func (it *WrappedTokenPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrappedTokenPoolMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*WrappedTokenPoolMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolMintedIterator{contract: _WrappedTokenPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrappedTokenPoolMinted)
				if err := _WrappedTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
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

func (_WrappedTokenPool *WrappedTokenPoolFilterer) ParseMinted(log types.Log) (*WrappedTokenPoolMinted, error) {
	event := new(WrappedTokenPoolMinted)
	if err := _WrappedTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrappedTokenPoolOffRampAllowanceSetIterator struct {
	Event *WrappedTokenPoolOffRampAllowanceSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrappedTokenPoolOffRampAllowanceSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenPoolOffRampAllowanceSet)
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
		it.Event = new(WrappedTokenPoolOffRampAllowanceSet)
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

func (it *WrappedTokenPoolOffRampAllowanceSetIterator) Error() error {
	return it.fail
}

func (it *WrappedTokenPoolOffRampAllowanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrappedTokenPoolOffRampAllowanceSet struct {
	OnRamp  common.Address
	Allowed bool
	Raw     types.Log
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) FilterOffRampAllowanceSet(opts *bind.FilterOpts) (*WrappedTokenPoolOffRampAllowanceSetIterator, error) {

	logs, sub, err := _WrappedTokenPool.contract.FilterLogs(opts, "OffRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolOffRampAllowanceSetIterator{contract: _WrappedTokenPool.contract, event: "OffRampAllowanceSet", logs: logs, sub: sub}, nil
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) WatchOffRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolOffRampAllowanceSet) (event.Subscription, error) {

	logs, sub, err := _WrappedTokenPool.contract.WatchLogs(opts, "OffRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrappedTokenPoolOffRampAllowanceSet)
				if err := _WrappedTokenPool.contract.UnpackLog(event, "OffRampAllowanceSet", log); err != nil {
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

func (_WrappedTokenPool *WrappedTokenPoolFilterer) ParseOffRampAllowanceSet(log types.Log) (*WrappedTokenPoolOffRampAllowanceSet, error) {
	event := new(WrappedTokenPoolOffRampAllowanceSet)
	if err := _WrappedTokenPool.contract.UnpackLog(event, "OffRampAllowanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrappedTokenPoolOnRampAllowanceSetIterator struct {
	Event *WrappedTokenPoolOnRampAllowanceSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrappedTokenPoolOnRampAllowanceSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenPoolOnRampAllowanceSet)
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
		it.Event = new(WrappedTokenPoolOnRampAllowanceSet)
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

func (it *WrappedTokenPoolOnRampAllowanceSetIterator) Error() error {
	return it.fail
}

func (it *WrappedTokenPoolOnRampAllowanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrappedTokenPoolOnRampAllowanceSet struct {
	OnRamp  common.Address
	Allowed bool
	Raw     types.Log
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) FilterOnRampAllowanceSet(opts *bind.FilterOpts) (*WrappedTokenPoolOnRampAllowanceSetIterator, error) {

	logs, sub, err := _WrappedTokenPool.contract.FilterLogs(opts, "OnRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolOnRampAllowanceSetIterator{contract: _WrappedTokenPool.contract, event: "OnRampAllowanceSet", logs: logs, sub: sub}, nil
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) WatchOnRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolOnRampAllowanceSet) (event.Subscription, error) {

	logs, sub, err := _WrappedTokenPool.contract.WatchLogs(opts, "OnRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrappedTokenPoolOnRampAllowanceSet)
				if err := _WrappedTokenPool.contract.UnpackLog(event, "OnRampAllowanceSet", log); err != nil {
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

func (_WrappedTokenPool *WrappedTokenPoolFilterer) ParseOnRampAllowanceSet(log types.Log) (*WrappedTokenPoolOnRampAllowanceSet, error) {
	event := new(WrappedTokenPoolOnRampAllowanceSet)
	if err := _WrappedTokenPool.contract.UnpackLog(event, "OnRampAllowanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrappedTokenPoolOwnershipTransferRequestedIterator struct {
	Event *WrappedTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrappedTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenPoolOwnershipTransferRequested)
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
		it.Event = new(WrappedTokenPoolOwnershipTransferRequested)
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

func (it *WrappedTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *WrappedTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrappedTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrappedTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolOwnershipTransferRequestedIterator{contract: _WrappedTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrappedTokenPoolOwnershipTransferRequested)
				if err := _WrappedTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_WrappedTokenPool *WrappedTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*WrappedTokenPoolOwnershipTransferRequested, error) {
	event := new(WrappedTokenPoolOwnershipTransferRequested)
	if err := _WrappedTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrappedTokenPoolOwnershipTransferredIterator struct {
	Event *WrappedTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrappedTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenPoolOwnershipTransferred)
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
		it.Event = new(WrappedTokenPoolOwnershipTransferred)
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

func (it *WrappedTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *WrappedTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrappedTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrappedTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolOwnershipTransferredIterator{contract: _WrappedTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrappedTokenPoolOwnershipTransferred)
				if err := _WrappedTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_WrappedTokenPool *WrappedTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*WrappedTokenPoolOwnershipTransferred, error) {
	event := new(WrappedTokenPoolOwnershipTransferred)
	if err := _WrappedTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrappedTokenPoolPausedIterator struct {
	Event *WrappedTokenPoolPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrappedTokenPoolPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenPoolPaused)
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
		it.Event = new(WrappedTokenPoolPaused)
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

func (it *WrappedTokenPoolPausedIterator) Error() error {
	return it.fail
}

func (it *WrappedTokenPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrappedTokenPoolPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*WrappedTokenPoolPausedIterator, error) {

	logs, sub, err := _WrappedTokenPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolPausedIterator{contract: _WrappedTokenPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolPaused) (event.Subscription, error) {

	logs, sub, err := _WrappedTokenPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrappedTokenPoolPaused)
				if err := _WrappedTokenPool.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_WrappedTokenPool *WrappedTokenPoolFilterer) ParsePaused(log types.Log) (*WrappedTokenPoolPaused, error) {
	event := new(WrappedTokenPoolPaused)
	if err := _WrappedTokenPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrappedTokenPoolReleasedIterator struct {
	Event *WrappedTokenPoolReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrappedTokenPoolReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenPoolReleased)
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
		it.Event = new(WrappedTokenPoolReleased)
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

func (it *WrappedTokenPoolReleasedIterator) Error() error {
	return it.fail
}

func (it *WrappedTokenPoolReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrappedTokenPoolReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*WrappedTokenPoolReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolReleasedIterator{contract: _WrappedTokenPool.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrappedTokenPoolReleased)
				if err := _WrappedTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
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

func (_WrappedTokenPool *WrappedTokenPoolFilterer) ParseReleased(log types.Log) (*WrappedTokenPoolReleased, error) {
	event := new(WrappedTokenPoolReleased)
	if err := _WrappedTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrappedTokenPoolTransferIterator struct {
	Event *WrappedTokenPoolTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrappedTokenPoolTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenPoolTransfer)
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
		it.Event = new(WrappedTokenPoolTransfer)
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

func (it *WrappedTokenPoolTransferIterator) Error() error {
	return it.fail
}

func (it *WrappedTokenPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrappedTokenPoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrappedTokenPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolTransferIterator{contract: _WrappedTokenPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedTokenPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrappedTokenPoolTransfer)
				if err := _WrappedTokenPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_WrappedTokenPool *WrappedTokenPoolFilterer) ParseTransfer(log types.Log) (*WrappedTokenPoolTransfer, error) {
	event := new(WrappedTokenPoolTransfer)
	if err := _WrappedTokenPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrappedTokenPoolUnpausedIterator struct {
	Event *WrappedTokenPoolUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrappedTokenPoolUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenPoolUnpaused)
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
		it.Event = new(WrappedTokenPoolUnpaused)
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

func (it *WrappedTokenPoolUnpausedIterator) Error() error {
	return it.fail
}

func (it *WrappedTokenPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrappedTokenPoolUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WrappedTokenPoolUnpausedIterator, error) {

	logs, sub, err := _WrappedTokenPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WrappedTokenPoolUnpausedIterator{contract: _WrappedTokenPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_WrappedTokenPool *WrappedTokenPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _WrappedTokenPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrappedTokenPoolUnpaused)
				if err := _WrappedTokenPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_WrappedTokenPool *WrappedTokenPoolFilterer) ParseUnpaused(log types.Log) (*WrappedTokenPoolUnpaused, error) {
	event := new(WrappedTokenPoolUnpaused)
	if err := _WrappedTokenPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_WrappedTokenPool *WrappedTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _WrappedTokenPool.abi.Events["Approval"].ID:
		return _WrappedTokenPool.ParseApproval(log)
	case _WrappedTokenPool.abi.Events["Burned"].ID:
		return _WrappedTokenPool.ParseBurned(log)
	case _WrappedTokenPool.abi.Events["Locked"].ID:
		return _WrappedTokenPool.ParseLocked(log)
	case _WrappedTokenPool.abi.Events["Minted"].ID:
		return _WrappedTokenPool.ParseMinted(log)
	case _WrappedTokenPool.abi.Events["OffRampAllowanceSet"].ID:
		return _WrappedTokenPool.ParseOffRampAllowanceSet(log)
	case _WrappedTokenPool.abi.Events["OnRampAllowanceSet"].ID:
		return _WrappedTokenPool.ParseOnRampAllowanceSet(log)
	case _WrappedTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _WrappedTokenPool.ParseOwnershipTransferRequested(log)
	case _WrappedTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _WrappedTokenPool.ParseOwnershipTransferred(log)
	case _WrappedTokenPool.abi.Events["Paused"].ID:
		return _WrappedTokenPool.ParsePaused(log)
	case _WrappedTokenPool.abi.Events["Released"].ID:
		return _WrappedTokenPool.ParseReleased(log)
	case _WrappedTokenPool.abi.Events["Transfer"].ID:
		return _WrappedTokenPool.ParseTransfer(log)
	case _WrappedTokenPool.abi.Events["Unpaused"].ID:
		return _WrappedTokenPool.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (WrappedTokenPoolApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (WrappedTokenPoolBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
}

func (WrappedTokenPoolLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (WrappedTokenPoolMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (WrappedTokenPoolOffRampAllowanceSet) Topic() common.Hash {
	return common.HexToHash("0xd8c3333ded377884ced3869cd0bcb9be54ea664076df1f5d39c4689120313648")
}

func (WrappedTokenPoolOnRampAllowanceSet) Topic() common.Hash {
	return common.HexToHash("0xbceff8f229c6dfcbf8bdcfb18726b84b0fd249b4803deb3948ff34d904013662")
}

func (WrappedTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (WrappedTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (WrappedTokenPoolPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (WrappedTokenPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (WrappedTokenPoolTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (WrappedTokenPoolUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_WrappedTokenPool *WrappedTokenPool) Address() common.Address {
	return _WrappedTokenPool.address
}

type WrappedTokenPoolInterface interface {
	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)

	CurrentTokenBucketState(opts *bind.CallOpts) (RateLimiterTokenBucket, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error)

	Name(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyRampUpdates(opts *bind.TransactOpts, onRamps []IPoolRampUpdate, offRamps []IPoolRampUpdate) (*types.Transaction, error)

	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)

	DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error)

	IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, amount *big.Int, arg1 common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WrappedTokenPoolApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*WrappedTokenPoolApproval, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*WrappedTokenPoolBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*WrappedTokenPoolBurned, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*WrappedTokenPoolLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*WrappedTokenPoolLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*WrappedTokenPoolMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*WrappedTokenPoolMinted, error)

	FilterOffRampAllowanceSet(opts *bind.FilterOpts) (*WrappedTokenPoolOffRampAllowanceSetIterator, error)

	WatchOffRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolOffRampAllowanceSet) (event.Subscription, error)

	ParseOffRampAllowanceSet(log types.Log) (*WrappedTokenPoolOffRampAllowanceSet, error)

	FilterOnRampAllowanceSet(opts *bind.FilterOpts) (*WrappedTokenPoolOnRampAllowanceSetIterator, error)

	WatchOnRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolOnRampAllowanceSet) (event.Subscription, error)

	ParseOnRampAllowanceSet(log types.Log) (*WrappedTokenPoolOnRampAllowanceSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrappedTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*WrappedTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrappedTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*WrappedTokenPoolOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*WrappedTokenPoolPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*WrappedTokenPoolPaused, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*WrappedTokenPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*WrappedTokenPoolReleased, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrappedTokenPoolTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*WrappedTokenPoolTransfer, error)

	FilterUnpaused(opts *bind.FilterOpts) (*WrappedTokenPoolUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WrappedTokenPoolUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*WrappedTokenPoolUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
