// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package custom_token_pool

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
	Ramp    common.Address
	Allowed bool
}

var CustomTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AllowListNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"OffRampAllowanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"OnRampAllowanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SynthBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SynthMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"ramp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"internalType\":\"structTokenPool.RampUpdate[]\",\"name\":\"onRamps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ramp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"internalType\":\"structTokenPool.RampUpdate[]\",\"name\":\"offRamps\",\"type\":\"tuple[]\"}],\"name\":\"applyRampUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"isOnRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"lockOrBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"releaseOrMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200232e3803806200232e833981016040819052620000349162000596565b604080516000808252602082019092528391508233806000816200009f5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000d257620000d281620001d0565b50506001805460ff60a01b19169055506001600160a01b0383166200010a57604051634655efd160e11b815260040160405180910390fd5b6040805160a08082018352602084810180516001600160801b039081168086524263ffffffff16938601849052875115158688018190529251821660608701819052968801519091166080958601819052600880546001600160a01b031916909217600160801b9485021760ff60a01b1916600160a01b909302929092179055029092176009556001600160a01b03851690528251158015909152620001c557604080516000815260208101909152620001c590836200027b565b5050505050620006c8565b336001600160a01b038216036200022a5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000096565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60a0516200029c576040516335f4a7b360e01b815260040160405180910390fd5b60005b825181101562000331576000838281518110620002c057620002c062000654565b60209081029190910101519050620002da600682620003ec565b156200031d576040516001600160a01b03821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b50620003298162000680565b90506200029f565b5060005b8151811015620003e757600082828151811062000356576200035662000654565b6020026020010151905060006001600160a01b0316816001600160a01b031603620003825750620003d4565b6200038f6006826200040c565b15620003d2576040516001600160a01b03821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b620003df8162000680565b905062000335565b505050565b600062000403836001600160a01b03841662000423565b90505b92915050565b600062000403836001600160a01b03841662000527565b600081815260018301602052604081205480156200051c5760006200044a6001836200069c565b855490915060009062000460906001906200069c565b9050818114620004cc57600086600001828154811062000484576200048462000654565b9060005260206000200154905080876000018481548110620004aa57620004aa62000654565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080620004e057620004e0620006b2565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000406565b600091505062000406565b6000818152600183016020526040812054620005705750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000406565b50600062000406565b80516001600160801b03811681146200059157600080fd5b919050565b6000808284036080811215620005ab57600080fd5b83516001600160a01b0381168114620005c357600080fd5b92506060601f1982011215620005d857600080fd5b50604051606081016001600160401b03811182821017156200060a57634e487b7160e01b600052604160045260246000fd5b604052602084015180151581146200062157600080fd5b8152620006316040850162000579565b6020820152620006446060850162000579565b6040820152809150509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016200069557620006956200066a565b5060010190565b818103818111156200040657620004066200066a565b634e487b7160e01b600052603160045260246000fd5b60805160a051611c39620006f5600039600081816103050152610d96015260006101780152611c396000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c80638456cb59116100b2578063a7cd63b711610081578063c92b283211610066578063c92b2832146102f0578063e0351e1314610303578063f2fde38b1461032957600080fd5b8063a7cd63b7146102c8578063af519112146102dd57600080fd5b80638456cb591461027c5780638627fad6146102845780638da5cb5b1461029757806396875445146102b557600080fd5b8063546719cd116101095780635c975abb116100ee5780635c975abb1461023e5780636f32b8721461026157806379ba50971461027457600080fd5b8063546719cd146101c757806354c8a4f31461022b57600080fd5b806301ffc9a71461013b5780631d7a74a01461016357806321df0da7146101765780633f4ba83a146101bd575b600080fd5b61014e610149366004611524565b61033c565b60405190151581526020015b60405180910390f35b61014e61017136600461158f565b6103d5565b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161015a565b6101c56103e2565b005b6101cf6103f4565b60405161015a919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b6101c56102393660046115f6565b6104a9565b60015474010000000000000000000000000000000000000000900460ff1661014e565b61014e61026f36600461158f565b610524565b6101c5610531565b6101c5610633565b6101c56102923660046117af565b610643565b60005473ffffffffffffffffffffffffffffffffffffffff16610198565b6101c56102c3366004611880565b610741565b6102d0610841565b60405161015a919061191e565b6101c56102eb366004611a3a565b6108ff565b6101c56102fe366004611abe565b610b0f565b7f000000000000000000000000000000000000000000000000000000000000000061014e565b6101c561033736600461158f565b610b25565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f317fa3340000000000000000000000000000000000000000000000000000000014806103cf57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b60006103cf600483610b36565b6103ea610b68565b6103f2610be9565b565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526008546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff1660208501527401000000000000000000000000000000000000000090920460ff1615159383019390935260095480841660608401520490911660808201526104a490610ce2565b905090565b6104b1610b68565b61051e84848080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808802828101820190935287825290935087925086918291850190849080828437600092019190915250610d9492505050565b50505050565b60006103cf600283610b36565b60015473ffffffffffffffffffffffffffffffffffffffff1633146105b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61063b610b68565b6103f2610f5a565b60015474010000000000000000000000000000000000000000900460ff16156106c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016105ae565b6106d1336103d5565b610707576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518381527fbb0b72e5f44e331506684da008a30e10d50658c29d8159f6c6ab40bf1e52e6009060200160405180910390a15050505050565b60015474010000000000000000000000000000000000000000900460ff16156107c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016105ae565b6107cf33610524565b610805576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518481527f02992093bca69a36949677658a77d359b510dc6232c68f9f118f7c0127a1b1479060200160405180910390a150505050505050565b6060600061084f6006611046565b67ffffffffffffffff81111561086757610867611662565b604051908082528060200260200182016040528015610890578160200160208202803683370190505b50905060005b6108a06006611046565b8110156108f9576108b2600682611050565b8282815181106108c4576108c4611b2a565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101526108f281611b88565b9050610896565b50919050565b610907610b68565b60005b8251811015610a0857600083828151811061092757610927611b2a565b60200260200101519050806020015161094d5780516109489060029061105c565b61095b565b805161095b9060029061107e565b156109f7577fbceff8f229c6dfcbf8bdcfb18726b84b0fd249b4803deb3948ff34d90401366284838151811061099357610993611b2a565b6020026020010151600001518584815181106109b1576109b1611b2a565b6020026020010151602001516040516109ee92919073ffffffffffffffffffffffffffffffffffffffff9290921682521515602082015260400190565b60405180910390a15b50610a0181611b88565b905061090a565b5060005b8151811015610b0a576000828281518110610a2957610a29611b2a565b602002602001015190508060200151610a4f578051610a4a9060049061105c565b610a5d565b8051610a5d9060049061107e565b15610af9577fd8c3333ded377884ced3869cd0bcb9be54ea664076df1f5d39c4689120313648838381518110610a9557610a95611b2a565b602002602001015160000151848481518110610ab357610ab3611b2a565b602002602001015160200151604051610af092919073ffffffffffffffffffffffffffffffffffffffff9290921682521515602082015260400190565b60405180910390a15b50610b0381611b88565b9050610a0c565b505050565b610b17610b68565b610b226008826110a0565b50565b610b2d610b68565b610b2281611285565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018301602052604081205415155b9392505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016105ae565b60015474010000000000000000000000000000000000000000900460ff16610c6d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016105ae565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152610d7082606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff1642610d549190611bc0565b85608001516fffffffffffffffffffffffffffffffff1661137a565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b7f0000000000000000000000000000000000000000000000000000000000000000610deb576040517f35f4a7b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8251811015610e89576000838281518110610e0b57610e0b611b2a565b60200260200101519050610e2981600661105c90919063ffffffff16565b15610e785760405173ffffffffffffffffffffffffffffffffffffffff821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b50610e8281611b88565b9050610dee565b5060005b8151811015610b0a576000828281518110610eaa57610eaa611b2a565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610eee5750610f4a565b610ef960068261107e565b15610f485760405173ffffffffffffffffffffffffffffffffffffffff821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b610f5381611b88565b9050610e8d565b60015474010000000000000000000000000000000000000000900460ff1615610fdf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016105ae565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610cb83390565b60006103cf825490565b6000610b6183836113a2565b6000610b618373ffffffffffffffffffffffffffffffffffffffff84166113cc565b6000610b618373ffffffffffffffffffffffffffffffffffffffff84166114bf565b81546000906110c990700100000000000000000000000000000000900463ffffffff1642611bc0565b9050801561116b5760018301548354611111916fffffffffffffffffffffffffffffffff8082169281169185917001000000000000000000000000000000009091041661137a565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b60208201518354611191916fffffffffffffffffffffffffffffffff908116911661150e565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19906112789084908151151581526020808301516fffffffffffffffffffffffffffffffff90811691830191909152604092830151169181019190915260600190565b60405180910390a1505050565b3373ffffffffffffffffffffffffffffffffffffffff821603611304576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016105ae565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006113998561138a8486611bd3565b6113949087611bea565b61150e565b95945050505050565b60008260000182815481106113b9576113b9611b2a565b9060005260206000200154905092915050565b600081815260018301602052604081205480156114b55760006113f0600183611bc0565b855490915060009061140490600190611bc0565b905081811461146957600086600001828154811061142457611424611b2a565b906000526020600020015490508087600001848154811061144757611447611b2a565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061147a5761147a611bfd565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506103cf565b60009150506103cf565b6000818152600183016020526040812054611506575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556103cf565b5060006103cf565b600081831061151d5781610b61565b5090919050565b60006020828403121561153657600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610b6157600080fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461158a57600080fd5b919050565b6000602082840312156115a157600080fd5b610b6182611566565b60008083601f8401126115bc57600080fd5b50813567ffffffffffffffff8111156115d457600080fd5b6020830191508360208260051b85010111156115ef57600080fd5b9250929050565b6000806000806040858703121561160c57600080fd5b843567ffffffffffffffff8082111561162457600080fd5b611630888389016115aa565b9096509450602087013591508082111561164957600080fd5b50611656878288016115aa565b95989497509550505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156116b4576116b4611662565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561170157611701611662565b604052919050565b600082601f83011261171a57600080fd5b813567ffffffffffffffff81111561173457611734611662565b61176560207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016116ba565b81815284602083860101111561177a57600080fd5b816020850160208301376000918101602001919091529392505050565b803567ffffffffffffffff8116811461158a57600080fd5b600080600080600060a086880312156117c757600080fd5b853567ffffffffffffffff808211156117df57600080fd5b6117eb89838a01611709565b96506117f960208901611566565b95506040880135945061180e60608901611797565b9350608088013591508082111561182457600080fd5b5061183188828901611709565b9150509295509295909350565b60008083601f84011261185057600080fd5b50813567ffffffffffffffff81111561186857600080fd5b6020830191508360208285010111156115ef57600080fd5b600080600080600080600060a0888a03121561189b57600080fd5b6118a488611566565b9650602088013567ffffffffffffffff808211156118c157600080fd5b6118cd8b838c0161183e565b909850965060408a013595508691506118e860608b01611797565b945060808a01359150808211156118fe57600080fd5b5061190b8a828b0161183e565b989b979a50959850939692959293505050565b6020808252825182820181905260009190848201906040850190845b8181101561196c57835173ffffffffffffffffffffffffffffffffffffffff168352928401929184019160010161193a565b50909695505050505050565b8035801515811461158a57600080fd5b600082601f83011261199957600080fd5b8135602067ffffffffffffffff8211156119b5576119b5611662565b6119c3818360051b016116ba565b82815260069290921b840181019181810190868411156119e257600080fd5b8286015b84811015611a2f57604081890312156119ff5760008081fd5b611a07611691565b611a1082611566565b8152611a1d858301611978565b818601528352918301916040016119e6565b509695505050505050565b60008060408385031215611a4d57600080fd5b823567ffffffffffffffff80821115611a6557600080fd5b611a7186838701611988565b93506020850135915080821115611a8757600080fd5b50611a9485828601611988565b9150509250929050565b80356fffffffffffffffffffffffffffffffff8116811461158a57600080fd5b600060608284031215611ad057600080fd5b6040516060810181811067ffffffffffffffff82111715611af357611af3611662565b604052611aff83611978565b8152611b0d60208401611a9e565b6020820152611b1e60408401611a9e565b60408201529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611bb957611bb9611b59565b5060010190565b818103818111156103cf576103cf611b59565b80820281158282048414176103cf576103cf611b59565b808201808211156103cf576103cf611b59565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000813000a",
}

var CustomTokenPoolABI = CustomTokenPoolMetaData.ABI

var CustomTokenPoolBin = CustomTokenPoolMetaData.Bin

func DeployCustomTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, rateLimiterConfig RateLimiterConfig) (common.Address, *types.Transaction, *CustomTokenPool, error) {
	parsed, err := CustomTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CustomTokenPoolBin), backend, token, rateLimiterConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CustomTokenPool{CustomTokenPoolCaller: CustomTokenPoolCaller{contract: contract}, CustomTokenPoolTransactor: CustomTokenPoolTransactor{contract: contract}, CustomTokenPoolFilterer: CustomTokenPoolFilterer{contract: contract}}, nil
}

type CustomTokenPool struct {
	address common.Address
	abi     abi.ABI
	CustomTokenPoolCaller
	CustomTokenPoolTransactor
	CustomTokenPoolFilterer
}

type CustomTokenPoolCaller struct {
	contract *bind.BoundContract
}

type CustomTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type CustomTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type CustomTokenPoolSession struct {
	Contract     *CustomTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CustomTokenPoolCallerSession struct {
	Contract *CustomTokenPoolCaller
	CallOpts bind.CallOpts
}

type CustomTokenPoolTransactorSession struct {
	Contract     *CustomTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type CustomTokenPoolRaw struct {
	Contract *CustomTokenPool
}

type CustomTokenPoolCallerRaw struct {
	Contract *CustomTokenPoolCaller
}

type CustomTokenPoolTransactorRaw struct {
	Contract *CustomTokenPoolTransactor
}

func NewCustomTokenPool(address common.Address, backend bind.ContractBackend) (*CustomTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(CustomTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCustomTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPool{address: address, abi: abi, CustomTokenPoolCaller: CustomTokenPoolCaller{contract: contract}, CustomTokenPoolTransactor: CustomTokenPoolTransactor{contract: contract}, CustomTokenPoolFilterer: CustomTokenPoolFilterer{contract: contract}}, nil
}

func NewCustomTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*CustomTokenPoolCaller, error) {
	contract, err := bindCustomTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolCaller{contract: contract}, nil
}

func NewCustomTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*CustomTokenPoolTransactor, error) {
	contract, err := bindCustomTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolTransactor{contract: contract}, nil
}

func NewCustomTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*CustomTokenPoolFilterer, error) {
	contract, err := bindCustomTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolFilterer{contract: contract}, nil
}

func bindCustomTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CustomTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CustomTokenPool *CustomTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustomTokenPool.Contract.CustomTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_CustomTokenPool *CustomTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.CustomTokenPoolTransactor.contract.Transfer(opts)
}

func (_CustomTokenPool *CustomTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.CustomTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_CustomTokenPool *CustomTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustomTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_CustomTokenPool *CustomTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.contract.Transfer(opts)
}

func (_CustomTokenPool *CustomTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_CustomTokenPool *CustomTokenPoolCaller) CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "currentRateLimiterState")

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _CustomTokenPool.Contract.CurrentRateLimiterState(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _CustomTokenPool.Contract.CurrentRateLimiterState(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _CustomTokenPool.Contract.GetAllowList(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _CustomTokenPool.Contract.GetAllowList(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _CustomTokenPool.Contract.GetAllowListEnabled(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _CustomTokenPool.Contract.GetAllowListEnabled(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) GetToken() (common.Address, error) {
	return _CustomTokenPool.Contract.GetToken(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _CustomTokenPool.Contract.GetToken(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCaller) IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "isOffRamp", offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _CustomTokenPool.Contract.IsOffRamp(&_CustomTokenPool.CallOpts, offRamp)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _CustomTokenPool.Contract.IsOffRamp(&_CustomTokenPool.CallOpts, offRamp)
}

func (_CustomTokenPool *CustomTokenPoolCaller) IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "isOnRamp", onRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _CustomTokenPool.Contract.IsOnRamp(&_CustomTokenPool.CallOpts, onRamp)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _CustomTokenPool.Contract.IsOnRamp(&_CustomTokenPool.CallOpts, onRamp)
}

func (_CustomTokenPool *CustomTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) Owner() (common.Address, error) {
	return _CustomTokenPool.Contract.Owner(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) Owner() (common.Address, error) {
	return _CustomTokenPool.Contract.Owner(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) Paused() (bool, error) {
	return _CustomTokenPool.Contract.Paused(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) Paused() (bool, error) {
	return _CustomTokenPool.Contract.Paused(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CustomTokenPool.Contract.SupportsInterface(&_CustomTokenPool.CallOpts, interfaceId)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CustomTokenPool.Contract.SupportsInterface(&_CustomTokenPool.CallOpts, interfaceId)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_CustomTokenPool *CustomTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.AcceptOwnership(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.AcceptOwnership(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_CustomTokenPool *CustomTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.ApplyAllowListUpdates(&_CustomTokenPool.TransactOpts, removes, adds)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.ApplyAllowListUpdates(&_CustomTokenPool.TransactOpts, removes, adds)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) ApplyRampUpdates(opts *bind.TransactOpts, onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "applyRampUpdates", onRamps, offRamps)
}

func (_CustomTokenPool *CustomTokenPoolSession) ApplyRampUpdates(onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.ApplyRampUpdates(&_CustomTokenPool.TransactOpts, onRamps, offRamps)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) ApplyRampUpdates(onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.ApplyRampUpdates(&_CustomTokenPool.TransactOpts, onRamps, offRamps)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, arg0 common.Address, arg1 []byte, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "lockOrBurn", arg0, arg1, amount, arg3, arg4)
}

func (_CustomTokenPool *CustomTokenPoolSession) LockOrBurn(arg0 common.Address, arg1 []byte, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.LockOrBurn(&_CustomTokenPool.TransactOpts, arg0, arg1, amount, arg3, arg4)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) LockOrBurn(arg0 common.Address, arg1 []byte, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.LockOrBurn(&_CustomTokenPool.TransactOpts, arg0, arg1, amount, arg3, arg4)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "pause")
}

func (_CustomTokenPool *CustomTokenPoolSession) Pause() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.Pause(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) Pause() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.Pause(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, arg0 []byte, arg1 common.Address, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "releaseOrMint", arg0, arg1, amount, arg3, arg4)
}

func (_CustomTokenPool *CustomTokenPoolSession) ReleaseOrMint(arg0 []byte, arg1 common.Address, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.ReleaseOrMint(&_CustomTokenPool.TransactOpts, arg0, arg1, amount, arg3, arg4)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) ReleaseOrMint(arg0 []byte, arg1 common.Address, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.ReleaseOrMint(&_CustomTokenPool.TransactOpts, arg0, arg1, amount, arg3, arg4)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_CustomTokenPool *CustomTokenPoolSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.SetRateLimiterConfig(&_CustomTokenPool.TransactOpts, config)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.SetRateLimiterConfig(&_CustomTokenPool.TransactOpts, config)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_CustomTokenPool *CustomTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.TransferOwnership(&_CustomTokenPool.TransactOpts, to)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.TransferOwnership(&_CustomTokenPool.TransactOpts, to)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "unpause")
}

func (_CustomTokenPool *CustomTokenPoolSession) Unpause() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.Unpause(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.Unpause(&_CustomTokenPool.TransactOpts)
}

type CustomTokenPoolAllowListAddIterator struct {
	Event *CustomTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolAllowListAdd)
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
		it.Event = new(CustomTokenPoolAllowListAdd)
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

func (it *CustomTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*CustomTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolAllowListAddIterator{contract: _CustomTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolAllowListAdd)
				if err := _CustomTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*CustomTokenPoolAllowListAdd, error) {
	event := new(CustomTokenPoolAllowListAdd)
	if err := _CustomTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolAllowListRemoveIterator struct {
	Event *CustomTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolAllowListRemove)
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
		it.Event = new(CustomTokenPoolAllowListRemove)
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

func (it *CustomTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*CustomTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolAllowListRemoveIterator{contract: _CustomTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolAllowListRemove)
				if err := _CustomTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*CustomTokenPoolAllowListRemove, error) {
	event := new(CustomTokenPoolAllowListRemove)
	if err := _CustomTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolBurnedIterator struct {
	Event *CustomTokenPoolBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolBurned)
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
		it.Event = new(CustomTokenPoolBurned)
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

func (it *CustomTokenPoolBurnedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*CustomTokenPoolBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolBurnedIterator{contract: _CustomTokenPool.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolBurned)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseBurned(log types.Log) (*CustomTokenPoolBurned, error) {
	event := new(CustomTokenPoolBurned)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolLockedIterator struct {
	Event *CustomTokenPoolLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolLocked)
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
		it.Event = new(CustomTokenPoolLocked)
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

func (it *CustomTokenPoolLockedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*CustomTokenPoolLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolLockedIterator{contract: _CustomTokenPool.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolLocked)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseLocked(log types.Log) (*CustomTokenPoolLocked, error) {
	event := new(CustomTokenPoolLocked)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolMintedIterator struct {
	Event *CustomTokenPoolMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolMinted)
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
		it.Event = new(CustomTokenPoolMinted)
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

func (it *CustomTokenPoolMintedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*CustomTokenPoolMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolMintedIterator{contract: _CustomTokenPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolMinted)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseMinted(log types.Log) (*CustomTokenPoolMinted, error) {
	event := new(CustomTokenPoolMinted)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolOffRampAllowanceSetIterator struct {
	Event *CustomTokenPoolOffRampAllowanceSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolOffRampAllowanceSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolOffRampAllowanceSet)
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
		it.Event = new(CustomTokenPoolOffRampAllowanceSet)
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

func (it *CustomTokenPoolOffRampAllowanceSetIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolOffRampAllowanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolOffRampAllowanceSet struct {
	OnRamp  common.Address
	Allowed bool
	Raw     types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterOffRampAllowanceSet(opts *bind.FilterOpts) (*CustomTokenPoolOffRampAllowanceSetIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "OffRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolOffRampAllowanceSetIterator{contract: _CustomTokenPool.contract, event: "OffRampAllowanceSet", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchOffRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOffRampAllowanceSet) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "OffRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolOffRampAllowanceSet)
				if err := _CustomTokenPool.contract.UnpackLog(event, "OffRampAllowanceSet", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseOffRampAllowanceSet(log types.Log) (*CustomTokenPoolOffRampAllowanceSet, error) {
	event := new(CustomTokenPoolOffRampAllowanceSet)
	if err := _CustomTokenPool.contract.UnpackLog(event, "OffRampAllowanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolOnRampAllowanceSetIterator struct {
	Event *CustomTokenPoolOnRampAllowanceSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolOnRampAllowanceSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolOnRampAllowanceSet)
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
		it.Event = new(CustomTokenPoolOnRampAllowanceSet)
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

func (it *CustomTokenPoolOnRampAllowanceSetIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolOnRampAllowanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolOnRampAllowanceSet struct {
	OnRamp  common.Address
	Allowed bool
	Raw     types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterOnRampAllowanceSet(opts *bind.FilterOpts) (*CustomTokenPoolOnRampAllowanceSetIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "OnRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolOnRampAllowanceSetIterator{contract: _CustomTokenPool.contract, event: "OnRampAllowanceSet", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchOnRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOnRampAllowanceSet) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "OnRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolOnRampAllowanceSet)
				if err := _CustomTokenPool.contract.UnpackLog(event, "OnRampAllowanceSet", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseOnRampAllowanceSet(log types.Log) (*CustomTokenPoolOnRampAllowanceSet, error) {
	event := new(CustomTokenPoolOnRampAllowanceSet)
	if err := _CustomTokenPool.contract.UnpackLog(event, "OnRampAllowanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolOwnershipTransferRequestedIterator struct {
	Event *CustomTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolOwnershipTransferRequested)
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
		it.Event = new(CustomTokenPoolOwnershipTransferRequested)
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

func (it *CustomTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CustomTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolOwnershipTransferRequestedIterator{contract: _CustomTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolOwnershipTransferRequested)
				if err := _CustomTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*CustomTokenPoolOwnershipTransferRequested, error) {
	event := new(CustomTokenPoolOwnershipTransferRequested)
	if err := _CustomTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolOwnershipTransferredIterator struct {
	Event *CustomTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolOwnershipTransferred)
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
		it.Event = new(CustomTokenPoolOwnershipTransferred)
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

func (it *CustomTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CustomTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolOwnershipTransferredIterator{contract: _CustomTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolOwnershipTransferred)
				if err := _CustomTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*CustomTokenPoolOwnershipTransferred, error) {
	event := new(CustomTokenPoolOwnershipTransferred)
	if err := _CustomTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolPausedIterator struct {
	Event *CustomTokenPoolPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolPaused)
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
		it.Event = new(CustomTokenPoolPaused)
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

func (it *CustomTokenPoolPausedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*CustomTokenPoolPausedIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolPausedIterator{contract: _CustomTokenPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolPaused) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolPaused)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParsePaused(log types.Log) (*CustomTokenPoolPaused, error) {
	event := new(CustomTokenPoolPaused)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolReleasedIterator struct {
	Event *CustomTokenPoolReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolReleased)
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
		it.Event = new(CustomTokenPoolReleased)
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

func (it *CustomTokenPoolReleasedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*CustomTokenPoolReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolReleasedIterator{contract: _CustomTokenPool.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolReleased)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseReleased(log types.Log) (*CustomTokenPoolReleased, error) {
	event := new(CustomTokenPoolReleased)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolSynthBurnedIterator struct {
	Event *CustomTokenPoolSynthBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolSynthBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolSynthBurned)
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
		it.Event = new(CustomTokenPoolSynthBurned)
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

func (it *CustomTokenPoolSynthBurnedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolSynthBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolSynthBurned struct {
	Amount *big.Int
	Raw    types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterSynthBurned(opts *bind.FilterOpts) (*CustomTokenPoolSynthBurnedIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "SynthBurned")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolSynthBurnedIterator{contract: _CustomTokenPool.contract, event: "SynthBurned", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchSynthBurned(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolSynthBurned) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "SynthBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolSynthBurned)
				if err := _CustomTokenPool.contract.UnpackLog(event, "SynthBurned", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseSynthBurned(log types.Log) (*CustomTokenPoolSynthBurned, error) {
	event := new(CustomTokenPoolSynthBurned)
	if err := _CustomTokenPool.contract.UnpackLog(event, "SynthBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolSynthMintedIterator struct {
	Event *CustomTokenPoolSynthMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolSynthMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolSynthMinted)
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
		it.Event = new(CustomTokenPoolSynthMinted)
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

func (it *CustomTokenPoolSynthMintedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolSynthMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolSynthMinted struct {
	Amount *big.Int
	Raw    types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterSynthMinted(opts *bind.FilterOpts) (*CustomTokenPoolSynthMintedIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "SynthMinted")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolSynthMintedIterator{contract: _CustomTokenPool.contract, event: "SynthMinted", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchSynthMinted(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolSynthMinted) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "SynthMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolSynthMinted)
				if err := _CustomTokenPool.contract.UnpackLog(event, "SynthMinted", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseSynthMinted(log types.Log) (*CustomTokenPoolSynthMinted, error) {
	event := new(CustomTokenPoolSynthMinted)
	if err := _CustomTokenPool.contract.UnpackLog(event, "SynthMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolUnpausedIterator struct {
	Event *CustomTokenPoolUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolUnpaused)
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
		it.Event = new(CustomTokenPoolUnpaused)
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

func (it *CustomTokenPoolUnpausedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CustomTokenPoolUnpausedIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolUnpausedIterator{contract: _CustomTokenPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolUnpaused)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseUnpaused(log types.Log) (*CustomTokenPoolUnpaused, error) {
	event := new(CustomTokenPoolUnpaused)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_CustomTokenPool *CustomTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CustomTokenPool.abi.Events["AllowListAdd"].ID:
		return _CustomTokenPool.ParseAllowListAdd(log)
	case _CustomTokenPool.abi.Events["AllowListRemove"].ID:
		return _CustomTokenPool.ParseAllowListRemove(log)
	case _CustomTokenPool.abi.Events["Burned"].ID:
		return _CustomTokenPool.ParseBurned(log)
	case _CustomTokenPool.abi.Events["Locked"].ID:
		return _CustomTokenPool.ParseLocked(log)
	case _CustomTokenPool.abi.Events["Minted"].ID:
		return _CustomTokenPool.ParseMinted(log)
	case _CustomTokenPool.abi.Events["OffRampAllowanceSet"].ID:
		return _CustomTokenPool.ParseOffRampAllowanceSet(log)
	case _CustomTokenPool.abi.Events["OnRampAllowanceSet"].ID:
		return _CustomTokenPool.ParseOnRampAllowanceSet(log)
	case _CustomTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _CustomTokenPool.ParseOwnershipTransferRequested(log)
	case _CustomTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _CustomTokenPool.ParseOwnershipTransferred(log)
	case _CustomTokenPool.abi.Events["Paused"].ID:
		return _CustomTokenPool.ParsePaused(log)
	case _CustomTokenPool.abi.Events["Released"].ID:
		return _CustomTokenPool.ParseReleased(log)
	case _CustomTokenPool.abi.Events["SynthBurned"].ID:
		return _CustomTokenPool.ParseSynthBurned(log)
	case _CustomTokenPool.abi.Events["SynthMinted"].ID:
		return _CustomTokenPool.ParseSynthMinted(log)
	case _CustomTokenPool.abi.Events["Unpaused"].ID:
		return _CustomTokenPool.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CustomTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (CustomTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (CustomTokenPoolBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
}

func (CustomTokenPoolLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (CustomTokenPoolMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (CustomTokenPoolOffRampAllowanceSet) Topic() common.Hash {
	return common.HexToHash("0xd8c3333ded377884ced3869cd0bcb9be54ea664076df1f5d39c4689120313648")
}

func (CustomTokenPoolOnRampAllowanceSet) Topic() common.Hash {
	return common.HexToHash("0xbceff8f229c6dfcbf8bdcfb18726b84b0fd249b4803deb3948ff34d904013662")
}

func (CustomTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CustomTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CustomTokenPoolPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (CustomTokenPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (CustomTokenPoolSynthBurned) Topic() common.Hash {
	return common.HexToHash("0x02992093bca69a36949677658a77d359b510dc6232c68f9f118f7c0127a1b147")
}

func (CustomTokenPoolSynthMinted) Topic() common.Hash {
	return common.HexToHash("0xbb0b72e5f44e331506684da008a30e10d50658c29d8159f6c6ab40bf1e52e600")
}

func (CustomTokenPoolUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_CustomTokenPool *CustomTokenPool) Address() common.Address {
	return _CustomTokenPool.address
}

type CustomTokenPoolInterface interface {
	CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyRampUpdates(opts *bind.TransactOpts, onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, arg0 common.Address, arg1 []byte, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, arg0 []byte, arg1 common.Address, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*CustomTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*CustomTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*CustomTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*CustomTokenPoolAllowListRemove, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*CustomTokenPoolBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*CustomTokenPoolBurned, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*CustomTokenPoolLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*CustomTokenPoolLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*CustomTokenPoolMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*CustomTokenPoolMinted, error)

	FilterOffRampAllowanceSet(opts *bind.FilterOpts) (*CustomTokenPoolOffRampAllowanceSetIterator, error)

	WatchOffRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOffRampAllowanceSet) (event.Subscription, error)

	ParseOffRampAllowanceSet(log types.Log) (*CustomTokenPoolOffRampAllowanceSet, error)

	FilterOnRampAllowanceSet(opts *bind.FilterOpts) (*CustomTokenPoolOnRampAllowanceSetIterator, error)

	WatchOnRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOnRampAllowanceSet) (event.Subscription, error)

	ParseOnRampAllowanceSet(log types.Log) (*CustomTokenPoolOnRampAllowanceSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CustomTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CustomTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CustomTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CustomTokenPoolOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*CustomTokenPoolPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*CustomTokenPoolPaused, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*CustomTokenPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*CustomTokenPoolReleased, error)

	FilterSynthBurned(opts *bind.FilterOpts) (*CustomTokenPoolSynthBurnedIterator, error)

	WatchSynthBurned(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolSynthBurned) (event.Subscription, error)

	ParseSynthBurned(log types.Log) (*CustomTokenPoolSynthBurned, error)

	FilterSynthMinted(opts *bind.FilterOpts) (*CustomTokenPoolSynthMintedIterator, error)

	WatchSynthMinted(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolSynthMinted) (event.Subscription, error)

	ParseSynthMinted(log types.Log) (*CustomTokenPoolSynthMinted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*CustomTokenPoolUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*CustomTokenPoolUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
