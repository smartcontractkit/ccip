// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package legacy_burn_mint_token_pool

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

var LegacyBurnMintTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILegacyBurnMintERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AllowListNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ConsumingMoreThanMaxCapacity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"RateLimitReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"OffRampAllowanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"OnRampAllowanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"ramp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"internalType\":\"structTokenPool.RampUpdate[]\",\"name\":\"onRamps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ramp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"internalType\":\"structTokenPool.RampUpdate[]\",\"name\":\"offRamps\",\"type\":\"tuple[]\"}],\"name\":\"applyRampUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"isOnRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"lockOrBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"releaseOrMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620029ac380380620029ac8339810160408190526200003491620006e7565b82828233806000816200008e5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c157620000c18162000236565b50506001805460ff60a01b19169055506001600160a01b038316620000f957604051634655efd160e11b815260040160405180910390fd5b6040805160a08082018352602084810180516001600160801b039081168086524263ffffffff16938601849052875115158688018190529251821660608701819052968801519091166080958601819052600880546001600160a01b031916909217600160801b9485021760ff60a01b1916600160a01b909302929092179055029092176009556001600160a01b03851690528251158015909152620001b457604080516000815260208101909152620001b49083620002e1565b505060405163095ea7b360e01b815230600482015260001960248201526001600160a01b038516915063095ea7b3906044016020604051808303816000875af115801562000206573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200022c9190620007d1565b5050505062000863565b336001600160a01b03821603620002905760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000085565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60a05162000302576040516335f4a7b360e01b815260040160405180910390fd5b60005b825181101562000397576000838281518110620003265762000326620007ef565b602090810291909101015190506200034060068262000452565b1562000383576040516001600160a01b03821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b506200038f816200081b565b905062000305565b5060005b81518110156200044d576000828281518110620003bc57620003bc620007ef565b6020026020010151905060006001600160a01b0316816001600160a01b031603620003e857506200043a565b620003f560068262000472565b1562000438576040516001600160a01b03821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b62000445816200081b565b90506200039b565b505050565b600062000469836001600160a01b03841662000489565b90505b92915050565b600062000469836001600160a01b0384166200058d565b6000818152600183016020526040812054801562000582576000620004b060018362000837565b8554909150600090620004c69060019062000837565b905081811462000532576000866000018281548110620004ea57620004ea620007ef565b9060005260206000200154905080876000018481548110620005105762000510620007ef565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806200054657620005466200084d565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506200046c565b60009150506200046c565b6000818152600183016020526040812054620005d6575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556200046c565b5060006200046c565b6001600160a01b0381168114620005f557600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620006395762000639620005f8565b604052919050565b805180151581146200065257600080fd5b919050565b80516001600160801b03811681146200065257600080fd5b6000606082840312156200068257600080fd5b604051606081016001600160401b0381118282101715620006a757620006a7620005f8565b604052905080620006b88362000641565b8152620006c86020840162000657565b6020820152620006db6040840162000657565b60408201525092915050565b600080600060a08486031215620006fd57600080fd5b83516200070a81620005df565b602085810151919450906001600160401b03808211156200072a57600080fd5b818701915087601f8301126200073f57600080fd5b815181811115620007545762000754620005f8565b8060051b9150620007678483016200060e565b818152918301840191848101908a8411156200078257600080fd5b938501935b83851015620007b057845192506200079f83620005df565b828252938501939085019062000787565b809750505050505050620007c885604086016200066f565b90509250925092565b600060208284031215620007e457600080fd5b620004698262000641565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820162000830576200083062000805565b5060010190565b818103818111156200046c576200046c62000805565b634e487b7160e01b600052603160045260246000fd5b60805160a051612107620008a560003960008181610305015281816108d30152610f8f0152600081816101780152818161075c015261098a01526121076000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c80638456cb59116100b2578063a7cd63b711610081578063c92b283211610066578063c92b2832146102f0578063e0351e1314610303578063f2fde38b1461032957600080fd5b8063a7cd63b7146102c8578063af519112146102dd57600080fd5b80638456cb591461027c5780638627fad6146102845780638da5cb5b1461029757806396875445146102b557600080fd5b8063546719cd116101095780635c975abb116100ee5780635c975abb1461023e5780636f32b8721461026157806379ba50971461027457600080fd5b8063546719cd146101c757806354c8a4f31461022b57600080fd5b806301ffc9a71461013b5780631d7a74a01461016357806321df0da7146101765780633f4ba83a146101bd575b600080fd5b61014e6101493660046119b7565b61033c565b60405190151581526020015b60405180910390f35b61014e610171366004611a22565b6103d5565b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161015a565b6101c56103e2565b005b6101cf6103f4565b60405161015a919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b6101c5610239366004611a89565b6104a9565b60015474010000000000000000000000000000000000000000900460ff1661014e565b61014e61026f366004611a22565b610524565b6101c5610531565b6101c5610633565b6101c5610292366004611c42565b610643565b60005473ffffffffffffffffffffffffffffffffffffffff16610198565b6101c56102c3366004611d13565b61080c565b6102d0610a3a565b60405161015a9190611db1565b6101c56102eb366004611ecd565b610af8565b6101c56102fe366004611f51565b610d08565b7f000000000000000000000000000000000000000000000000000000000000000061014e565b6101c5610337366004611a22565b610d1e565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f317fa3340000000000000000000000000000000000000000000000000000000014806103cf57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b60006103cf600483610d2f565b6103ea610d61565b6103f2610de2565b565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526008546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff1660208501527401000000000000000000000000000000000000000090920460ff1615159383019390935260095480841660608401520490911660808201526104a490610edb565b905090565b6104b1610d61565b61051e84848080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808802828101820190935287825290935087925086918291850190849080828437600092019190915250610f8d92505050565b50505050565b60006103cf600283610d2f565b60015473ffffffffffffffffffffffffffffffffffffffff1633146105b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61063b610d61565b6103f2611153565b60015474010000000000000000000000000000000000000000900460ff16156106c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016105ae565b6106d1336103d5565b610707576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107108361123f565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590527f000000000000000000000000000000000000000000000000000000000000000016906340c10f1990604401600060405180830381600087803b1580156107a057600080fd5b505af11580156107b4573d6000803e3d6000fd5b505060405185815273ffffffffffffffffffffffffffffffffffffffff871692503391507f9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f09060200160405180910390a35050505050565b60015474010000000000000000000000000000000000000000900460ff1615610891576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016105ae565b61089a33610524565b6108d0576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b867f000000000000000000000000000000000000000000000000000000000000000080156109065750610904600682610d2f565b155b15610955576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016105ae565b6040517f9dc29fac000000000000000000000000000000000000000000000000000000008152306004820152602481018690527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639dc29fac90604401600060405180830381600087803b1580156109e357600080fd5b505af11580156109f7573d6000803e3d6000fd5b50506040518781523392507f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7915060200160405180910390a25050505050505050565b60606000610a48600661124a565b67ffffffffffffffff811115610a6057610a60611af5565b604051908082528060200260200182016040528015610a89578160200160208202803683370190505b50905060005b610a99600661124a565b811015610af257610aab600682611254565b828281518110610abd57610abd611fbd565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152610aeb8161201b565b9050610a8f565b50919050565b610b00610d61565b60005b8251811015610c01576000838281518110610b2057610b20611fbd565b602002602001015190508060200151610b46578051610b4190600290611260565b610b54565b8051610b5490600290611282565b15610bf0577fbceff8f229c6dfcbf8bdcfb18726b84b0fd249b4803deb3948ff34d904013662848381518110610b8c57610b8c611fbd565b602002602001015160000151858481518110610baa57610baa611fbd565b602002602001015160200151604051610be792919073ffffffffffffffffffffffffffffffffffffffff9290921682521515602082015260400190565b60405180910390a15b50610bfa8161201b565b9050610b03565b5060005b8151811015610d03576000828281518110610c2257610c22611fbd565b602002602001015190508060200151610c48578051610c4390600490611260565b610c56565b8051610c5690600490611282565b15610cf2577fd8c3333ded377884ced3869cd0bcb9be54ea664076df1f5d39c4689120313648838381518110610c8e57610c8e611fbd565b602002602001015160000151848481518110610cac57610cac611fbd565b602002602001015160200151604051610ce992919073ffffffffffffffffffffffffffffffffffffffff9290921682521515602082015260400190565b60405180910390a15b50610cfc8161201b565b9050610c05565b505050565b610d10610d61565b610d1b6008826112a4565b50565b610d26610d61565b610d1b81611489565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018301602052604081205415155b9392505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016105ae565b60015474010000000000000000000000000000000000000000900460ff16610e66576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016105ae565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152610f6982606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff1642610f4d9190612053565b85608001516fffffffffffffffffffffffffffffffff1661157e565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b7f0000000000000000000000000000000000000000000000000000000000000000610fe4576040517f35f4a7b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b825181101561108257600083828151811061100457611004611fbd565b6020026020010151905061102281600661126090919063ffffffff16565b156110715760405173ffffffffffffffffffffffffffffffffffffffff821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b5061107b8161201b565b9050610fe7565b5060005b8151811015610d035760008282815181106110a3576110a3611fbd565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036110e75750611143565b6110f2600682611282565b156111415760405173ffffffffffffffffffffffffffffffffffffffff821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b61114c8161201b565b9050611086565b60015474010000000000000000000000000000000000000000900460ff16156111d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016105ae565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610eb13390565b610d1b6008826115a6565b60006103cf825490565b6000610d5a8383611835565b6000610d5a8373ffffffffffffffffffffffffffffffffffffffff841661185f565b6000610d5a8373ffffffffffffffffffffffffffffffffffffffff8416611952565b81546000906112cd90700100000000000000000000000000000000900463ffffffff1642612053565b9050801561136f5760018301548354611315916fffffffffffffffffffffffffffffffff8082169281169185917001000000000000000000000000000000009091041661157e565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b60208201518354611395916fffffffffffffffffffffffffffffffff90811691166119a1565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c199061147c9084908151151581526020808301516fffffffffffffffffffffffffffffffff90811691830191909152604092830151169181019190915260600190565b60405180910390a1505050565b3373ffffffffffffffffffffffffffffffffffffffff821603611508576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016105ae565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600061159d8561158e8486612066565b611598908761207d565b6119a1565b95945050505050565b815474010000000000000000000000000000000000000000900460ff1615806115cd575080155b156115d6575050565b815460018301546fffffffffffffffffffffffffffffffff8083169291169060009061161c90700100000000000000000000000000000000900463ffffffff1642612053565b905080156116dc578183111561165e576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018501546116989083908590849070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1661157e565b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff167001000000000000000000000000000000004263ffffffff160217865592505b83821015611720576040517f48369c4300000000000000000000000000000000000000000000000000000000815260048101839052602481018590526044016105ae565b838310156117b45760018581015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff169081906117619082612053565b61176b8688612053565b611775919061207d565b61177f9190612090565b6040517fdc96cefa0000000000000000000000000000000000000000000000000000000081526004016105ae91815260200190565b6117be8484612053565b85547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161786556040518581529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a15050505050565b600082600001828154811061184c5761184c611fbd565b9060005260206000200154905092915050565b60008181526001830160205260408120548015611948576000611883600183612053565b855490915060009061189790600190612053565b90508181146118fc5760008660000182815481106118b7576118b7611fbd565b90600052602060002001549050808760000184815481106118da576118da611fbd565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061190d5761190d6120cb565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506103cf565b60009150506103cf565b6000818152600183016020526040812054611999575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556103cf565b5060006103cf565b60008183106119b05781610d5a565b5090919050565b6000602082840312156119c957600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610d5a57600080fd5b803573ffffffffffffffffffffffffffffffffffffffff81168114611a1d57600080fd5b919050565b600060208284031215611a3457600080fd5b610d5a826119f9565b60008083601f840112611a4f57600080fd5b50813567ffffffffffffffff811115611a6757600080fd5b6020830191508360208260051b8501011115611a8257600080fd5b9250929050565b60008060008060408587031215611a9f57600080fd5b843567ffffffffffffffff80821115611ab757600080fd5b611ac388838901611a3d565b90965094506020870135915080821115611adc57600080fd5b50611ae987828801611a3d565b95989497509550505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715611b4757611b47611af5565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611b9457611b94611af5565b604052919050565b600082601f830112611bad57600080fd5b813567ffffffffffffffff811115611bc757611bc7611af5565b611bf860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611b4d565b818152846020838601011115611c0d57600080fd5b816020850160208301376000918101602001919091529392505050565b803567ffffffffffffffff81168114611a1d57600080fd5b600080600080600060a08688031215611c5a57600080fd5b853567ffffffffffffffff80821115611c7257600080fd5b611c7e89838a01611b9c565b9650611c8c602089016119f9565b955060408801359450611ca160608901611c2a565b93506080880135915080821115611cb757600080fd5b50611cc488828901611b9c565b9150509295509295909350565b60008083601f840112611ce357600080fd5b50813567ffffffffffffffff811115611cfb57600080fd5b602083019150836020828501011115611a8257600080fd5b600080600080600080600060a0888a031215611d2e57600080fd5b611d37886119f9565b9650602088013567ffffffffffffffff80821115611d5457600080fd5b611d608b838c01611cd1565b909850965060408a01359550869150611d7b60608b01611c2a565b945060808a0135915080821115611d9157600080fd5b50611d9e8a828b01611cd1565b989b979a50959850939692959293505050565b6020808252825182820181905260009190848201906040850190845b81811015611dff57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101611dcd565b50909695505050505050565b80358015158114611a1d57600080fd5b600082601f830112611e2c57600080fd5b8135602067ffffffffffffffff821115611e4857611e48611af5565b611e56818360051b01611b4d565b82815260069290921b84018101918181019086841115611e7557600080fd5b8286015b84811015611ec25760408189031215611e925760008081fd5b611e9a611b24565b611ea3826119f9565b8152611eb0858301611e0b565b81860152835291830191604001611e79565b509695505050505050565b60008060408385031215611ee057600080fd5b823567ffffffffffffffff80821115611ef857600080fd5b611f0486838701611e1b565b93506020850135915080821115611f1a57600080fd5b50611f2785828601611e1b565b9150509250929050565b80356fffffffffffffffffffffffffffffffff81168114611a1d57600080fd5b600060608284031215611f6357600080fd5b6040516060810181811067ffffffffffffffff82111715611f8657611f86611af5565b604052611f9283611e0b565b8152611fa060208401611f31565b6020820152611fb160408401611f31565b60408201529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361204c5761204c611fec565b5060010190565b818103818111156103cf576103cf611fec565b80820281158282048414176103cf576103cf611fec565b808201808211156103cf576103cf611fec565b6000826120c6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000813000a",
}

var LegacyBurnMintTokenPoolABI = LegacyBurnMintTokenPoolMetaData.ABI

var LegacyBurnMintTokenPoolBin = LegacyBurnMintTokenPoolMetaData.Bin

func DeployLegacyBurnMintTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, allowlist []common.Address, rateLimiterConfig RateLimiterConfig) (common.Address, *types.Transaction, *LegacyBurnMintTokenPool, error) {
	parsed, err := LegacyBurnMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LegacyBurnMintTokenPoolBin), backend, token, allowlist, rateLimiterConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LegacyBurnMintTokenPool{LegacyBurnMintTokenPoolCaller: LegacyBurnMintTokenPoolCaller{contract: contract}, LegacyBurnMintTokenPoolTransactor: LegacyBurnMintTokenPoolTransactor{contract: contract}, LegacyBurnMintTokenPoolFilterer: LegacyBurnMintTokenPoolFilterer{contract: contract}}, nil
}

type LegacyBurnMintTokenPool struct {
	address common.Address
	abi     abi.ABI
	LegacyBurnMintTokenPoolCaller
	LegacyBurnMintTokenPoolTransactor
	LegacyBurnMintTokenPoolFilterer
}

type LegacyBurnMintTokenPoolCaller struct {
	contract *bind.BoundContract
}

type LegacyBurnMintTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type LegacyBurnMintTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type LegacyBurnMintTokenPoolSession struct {
	Contract     *LegacyBurnMintTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type LegacyBurnMintTokenPoolCallerSession struct {
	Contract *LegacyBurnMintTokenPoolCaller
	CallOpts bind.CallOpts
}

type LegacyBurnMintTokenPoolTransactorSession struct {
	Contract     *LegacyBurnMintTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type LegacyBurnMintTokenPoolRaw struct {
	Contract *LegacyBurnMintTokenPool
}

type LegacyBurnMintTokenPoolCallerRaw struct {
	Contract *LegacyBurnMintTokenPoolCaller
}

type LegacyBurnMintTokenPoolTransactorRaw struct {
	Contract *LegacyBurnMintTokenPoolTransactor
}

func NewLegacyBurnMintTokenPool(address common.Address, backend bind.ContractBackend) (*LegacyBurnMintTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(LegacyBurnMintTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindLegacyBurnMintTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPool{address: address, abi: abi, LegacyBurnMintTokenPoolCaller: LegacyBurnMintTokenPoolCaller{contract: contract}, LegacyBurnMintTokenPoolTransactor: LegacyBurnMintTokenPoolTransactor{contract: contract}, LegacyBurnMintTokenPoolFilterer: LegacyBurnMintTokenPoolFilterer{contract: contract}}, nil
}

func NewLegacyBurnMintTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*LegacyBurnMintTokenPoolCaller, error) {
	contract, err := bindLegacyBurnMintTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolCaller{contract: contract}, nil
}

func NewLegacyBurnMintTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LegacyBurnMintTokenPoolTransactor, error) {
	contract, err := bindLegacyBurnMintTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolTransactor{contract: contract}, nil
}

func NewLegacyBurnMintTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LegacyBurnMintTokenPoolFilterer, error) {
	contract, err := bindLegacyBurnMintTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolFilterer{contract: contract}, nil
}

func bindLegacyBurnMintTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LegacyBurnMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyBurnMintTokenPool.Contract.LegacyBurnMintTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.LegacyBurnMintTokenPoolTransactor.contract.Transfer(opts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.LegacyBurnMintTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyBurnMintTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.contract.Transfer(opts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCaller) CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _LegacyBurnMintTokenPool.contract.Call(opts, &out, "currentRateLimiterState")

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _LegacyBurnMintTokenPool.Contract.CurrentRateLimiterState(&_LegacyBurnMintTokenPool.CallOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCallerSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _LegacyBurnMintTokenPool.Contract.CurrentRateLimiterState(&_LegacyBurnMintTokenPool.CallOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LegacyBurnMintTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _LegacyBurnMintTokenPool.Contract.GetAllowList(&_LegacyBurnMintTokenPool.CallOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _LegacyBurnMintTokenPool.Contract.GetAllowList(&_LegacyBurnMintTokenPool.CallOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LegacyBurnMintTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _LegacyBurnMintTokenPool.Contract.GetAllowListEnabled(&_LegacyBurnMintTokenPool.CallOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _LegacyBurnMintTokenPool.Contract.GetAllowListEnabled(&_LegacyBurnMintTokenPool.CallOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyBurnMintTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) GetToken() (common.Address, error) {
	return _LegacyBurnMintTokenPool.Contract.GetToken(&_LegacyBurnMintTokenPool.CallOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _LegacyBurnMintTokenPool.Contract.GetToken(&_LegacyBurnMintTokenPool.CallOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCaller) IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _LegacyBurnMintTokenPool.contract.Call(opts, &out, "isOffRamp", offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _LegacyBurnMintTokenPool.Contract.IsOffRamp(&_LegacyBurnMintTokenPool.CallOpts, offRamp)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCallerSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _LegacyBurnMintTokenPool.Contract.IsOffRamp(&_LegacyBurnMintTokenPool.CallOpts, offRamp)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCaller) IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error) {
	var out []interface{}
	err := _LegacyBurnMintTokenPool.contract.Call(opts, &out, "isOnRamp", onRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _LegacyBurnMintTokenPool.Contract.IsOnRamp(&_LegacyBurnMintTokenPool.CallOpts, onRamp)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCallerSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _LegacyBurnMintTokenPool.Contract.IsOnRamp(&_LegacyBurnMintTokenPool.CallOpts, onRamp)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyBurnMintTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) Owner() (common.Address, error) {
	return _LegacyBurnMintTokenPool.Contract.Owner(&_LegacyBurnMintTokenPool.CallOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCallerSession) Owner() (common.Address, error) {
	return _LegacyBurnMintTokenPool.Contract.Owner(&_LegacyBurnMintTokenPool.CallOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LegacyBurnMintTokenPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) Paused() (bool, error) {
	return _LegacyBurnMintTokenPool.Contract.Paused(&_LegacyBurnMintTokenPool.CallOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCallerSession) Paused() (bool, error) {
	return _LegacyBurnMintTokenPool.Contract.Paused(&_LegacyBurnMintTokenPool.CallOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LegacyBurnMintTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LegacyBurnMintTokenPool.Contract.SupportsInterface(&_LegacyBurnMintTokenPool.CallOpts, interfaceId)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LegacyBurnMintTokenPool.Contract.SupportsInterface(&_LegacyBurnMintTokenPool.CallOpts, interfaceId)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.AcceptOwnership(&_LegacyBurnMintTokenPool.TransactOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.AcceptOwnership(&_LegacyBurnMintTokenPool.TransactOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.ApplyAllowListUpdates(&_LegacyBurnMintTokenPool.TransactOpts, removes, adds)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.ApplyAllowListUpdates(&_LegacyBurnMintTokenPool.TransactOpts, removes, adds)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactor) ApplyRampUpdates(opts *bind.TransactOpts, onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.contract.Transact(opts, "applyRampUpdates", onRamps, offRamps)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) ApplyRampUpdates(onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.ApplyRampUpdates(&_LegacyBurnMintTokenPool.TransactOpts, onRamps, offRamps)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactorSession) ApplyRampUpdates(onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.ApplyRampUpdates(&_LegacyBurnMintTokenPool.TransactOpts, onRamps, offRamps)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, originalSender common.Address, arg1 []byte, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.contract.Transact(opts, "lockOrBurn", originalSender, arg1, amount, arg3, arg4)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) LockOrBurn(originalSender common.Address, arg1 []byte, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.LockOrBurn(&_LegacyBurnMintTokenPool.TransactOpts, originalSender, arg1, amount, arg3, arg4)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactorSession) LockOrBurn(originalSender common.Address, arg1 []byte, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.LockOrBurn(&_LegacyBurnMintTokenPool.TransactOpts, originalSender, arg1, amount, arg3, arg4)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.contract.Transact(opts, "pause")
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) Pause() (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.Pause(&_LegacyBurnMintTokenPool.TransactOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactorSession) Pause() (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.Pause(&_LegacyBurnMintTokenPool.TransactOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, arg0 []byte, receiver common.Address, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.contract.Transact(opts, "releaseOrMint", arg0, receiver, amount, arg3, arg4)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) ReleaseOrMint(arg0 []byte, receiver common.Address, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.ReleaseOrMint(&_LegacyBurnMintTokenPool.TransactOpts, arg0, receiver, amount, arg3, arg4)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactorSession) ReleaseOrMint(arg0 []byte, receiver common.Address, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.ReleaseOrMint(&_LegacyBurnMintTokenPool.TransactOpts, arg0, receiver, amount, arg3, arg4)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.SetRateLimiterConfig(&_LegacyBurnMintTokenPool.TransactOpts, config)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactorSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.SetRateLimiterConfig(&_LegacyBurnMintTokenPool.TransactOpts, config)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.TransferOwnership(&_LegacyBurnMintTokenPool.TransactOpts, to)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.TransferOwnership(&_LegacyBurnMintTokenPool.TransactOpts, to)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.contract.Transact(opts, "unpause")
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolSession) Unpause() (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.Unpause(&_LegacyBurnMintTokenPool.TransactOpts)
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _LegacyBurnMintTokenPool.Contract.Unpause(&_LegacyBurnMintTokenPool.TransactOpts)
}

type LegacyBurnMintTokenPoolAllowListAddIterator struct {
	Event *LegacyBurnMintTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LegacyBurnMintTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyBurnMintTokenPoolAllowListAdd)
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
		it.Event = new(LegacyBurnMintTokenPoolAllowListAdd)
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

func (it *LegacyBurnMintTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *LegacyBurnMintTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LegacyBurnMintTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*LegacyBurnMintTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _LegacyBurnMintTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolAllowListAddIterator{contract: _LegacyBurnMintTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _LegacyBurnMintTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LegacyBurnMintTokenPoolAllowListAdd)
				if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*LegacyBurnMintTokenPoolAllowListAdd, error) {
	event := new(LegacyBurnMintTokenPoolAllowListAdd)
	if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LegacyBurnMintTokenPoolAllowListRemoveIterator struct {
	Event *LegacyBurnMintTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LegacyBurnMintTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyBurnMintTokenPoolAllowListRemove)
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
		it.Event = new(LegacyBurnMintTokenPoolAllowListRemove)
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

func (it *LegacyBurnMintTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *LegacyBurnMintTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LegacyBurnMintTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*LegacyBurnMintTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _LegacyBurnMintTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolAllowListRemoveIterator{contract: _LegacyBurnMintTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _LegacyBurnMintTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LegacyBurnMintTokenPoolAllowListRemove)
				if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*LegacyBurnMintTokenPoolAllowListRemove, error) {
	event := new(LegacyBurnMintTokenPoolAllowListRemove)
	if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LegacyBurnMintTokenPoolBurnedIterator struct {
	Event *LegacyBurnMintTokenPoolBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LegacyBurnMintTokenPoolBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyBurnMintTokenPoolBurned)
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
		it.Event = new(LegacyBurnMintTokenPoolBurned)
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

func (it *LegacyBurnMintTokenPoolBurnedIterator) Error() error {
	return it.fail
}

func (it *LegacyBurnMintTokenPoolBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LegacyBurnMintTokenPoolBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*LegacyBurnMintTokenPoolBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LegacyBurnMintTokenPool.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolBurnedIterator{contract: _LegacyBurnMintTokenPool.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LegacyBurnMintTokenPool.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LegacyBurnMintTokenPoolBurned)
				if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
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

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) ParseBurned(log types.Log) (*LegacyBurnMintTokenPoolBurned, error) {
	event := new(LegacyBurnMintTokenPoolBurned)
	if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LegacyBurnMintTokenPoolLockedIterator struct {
	Event *LegacyBurnMintTokenPoolLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LegacyBurnMintTokenPoolLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyBurnMintTokenPoolLocked)
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
		it.Event = new(LegacyBurnMintTokenPoolLocked)
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

func (it *LegacyBurnMintTokenPoolLockedIterator) Error() error {
	return it.fail
}

func (it *LegacyBurnMintTokenPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LegacyBurnMintTokenPoolLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*LegacyBurnMintTokenPoolLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LegacyBurnMintTokenPool.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolLockedIterator{contract: _LegacyBurnMintTokenPool.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LegacyBurnMintTokenPool.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LegacyBurnMintTokenPoolLocked)
				if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
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

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) ParseLocked(log types.Log) (*LegacyBurnMintTokenPoolLocked, error) {
	event := new(LegacyBurnMintTokenPoolLocked)
	if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LegacyBurnMintTokenPoolMintedIterator struct {
	Event *LegacyBurnMintTokenPoolMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LegacyBurnMintTokenPoolMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyBurnMintTokenPoolMinted)
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
		it.Event = new(LegacyBurnMintTokenPoolMinted)
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

func (it *LegacyBurnMintTokenPoolMintedIterator) Error() error {
	return it.fail
}

func (it *LegacyBurnMintTokenPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LegacyBurnMintTokenPoolMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LegacyBurnMintTokenPoolMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LegacyBurnMintTokenPool.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolMintedIterator{contract: _LegacyBurnMintTokenPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LegacyBurnMintTokenPool.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LegacyBurnMintTokenPoolMinted)
				if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
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

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) ParseMinted(log types.Log) (*LegacyBurnMintTokenPoolMinted, error) {
	event := new(LegacyBurnMintTokenPoolMinted)
	if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LegacyBurnMintTokenPoolOffRampAllowanceSetIterator struct {
	Event *LegacyBurnMintTokenPoolOffRampAllowanceSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LegacyBurnMintTokenPoolOffRampAllowanceSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyBurnMintTokenPoolOffRampAllowanceSet)
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
		it.Event = new(LegacyBurnMintTokenPoolOffRampAllowanceSet)
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

func (it *LegacyBurnMintTokenPoolOffRampAllowanceSetIterator) Error() error {
	return it.fail
}

func (it *LegacyBurnMintTokenPoolOffRampAllowanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LegacyBurnMintTokenPoolOffRampAllowanceSet struct {
	OnRamp  common.Address
	Allowed bool
	Raw     types.Log
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) FilterOffRampAllowanceSet(opts *bind.FilterOpts) (*LegacyBurnMintTokenPoolOffRampAllowanceSetIterator, error) {

	logs, sub, err := _LegacyBurnMintTokenPool.contract.FilterLogs(opts, "OffRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolOffRampAllowanceSetIterator{contract: _LegacyBurnMintTokenPool.contract, event: "OffRampAllowanceSet", logs: logs, sub: sub}, nil
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) WatchOffRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolOffRampAllowanceSet) (event.Subscription, error) {

	logs, sub, err := _LegacyBurnMintTokenPool.contract.WatchLogs(opts, "OffRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LegacyBurnMintTokenPoolOffRampAllowanceSet)
				if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "OffRampAllowanceSet", log); err != nil {
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

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) ParseOffRampAllowanceSet(log types.Log) (*LegacyBurnMintTokenPoolOffRampAllowanceSet, error) {
	event := new(LegacyBurnMintTokenPoolOffRampAllowanceSet)
	if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "OffRampAllowanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LegacyBurnMintTokenPoolOnRampAllowanceSetIterator struct {
	Event *LegacyBurnMintTokenPoolOnRampAllowanceSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LegacyBurnMintTokenPoolOnRampAllowanceSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyBurnMintTokenPoolOnRampAllowanceSet)
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
		it.Event = new(LegacyBurnMintTokenPoolOnRampAllowanceSet)
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

func (it *LegacyBurnMintTokenPoolOnRampAllowanceSetIterator) Error() error {
	return it.fail
}

func (it *LegacyBurnMintTokenPoolOnRampAllowanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LegacyBurnMintTokenPoolOnRampAllowanceSet struct {
	OnRamp  common.Address
	Allowed bool
	Raw     types.Log
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) FilterOnRampAllowanceSet(opts *bind.FilterOpts) (*LegacyBurnMintTokenPoolOnRampAllowanceSetIterator, error) {

	logs, sub, err := _LegacyBurnMintTokenPool.contract.FilterLogs(opts, "OnRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolOnRampAllowanceSetIterator{contract: _LegacyBurnMintTokenPool.contract, event: "OnRampAllowanceSet", logs: logs, sub: sub}, nil
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) WatchOnRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolOnRampAllowanceSet) (event.Subscription, error) {

	logs, sub, err := _LegacyBurnMintTokenPool.contract.WatchLogs(opts, "OnRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LegacyBurnMintTokenPoolOnRampAllowanceSet)
				if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "OnRampAllowanceSet", log); err != nil {
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

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) ParseOnRampAllowanceSet(log types.Log) (*LegacyBurnMintTokenPoolOnRampAllowanceSet, error) {
	event := new(LegacyBurnMintTokenPoolOnRampAllowanceSet)
	if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "OnRampAllowanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LegacyBurnMintTokenPoolOwnershipTransferRequestedIterator struct {
	Event *LegacyBurnMintTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LegacyBurnMintTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyBurnMintTokenPoolOwnershipTransferRequested)
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
		it.Event = new(LegacyBurnMintTokenPoolOwnershipTransferRequested)
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

func (it *LegacyBurnMintTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *LegacyBurnMintTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LegacyBurnMintTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LegacyBurnMintTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LegacyBurnMintTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolOwnershipTransferRequestedIterator{contract: _LegacyBurnMintTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LegacyBurnMintTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LegacyBurnMintTokenPoolOwnershipTransferRequested)
				if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*LegacyBurnMintTokenPoolOwnershipTransferRequested, error) {
	event := new(LegacyBurnMintTokenPoolOwnershipTransferRequested)
	if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LegacyBurnMintTokenPoolOwnershipTransferredIterator struct {
	Event *LegacyBurnMintTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LegacyBurnMintTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyBurnMintTokenPoolOwnershipTransferred)
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
		it.Event = new(LegacyBurnMintTokenPoolOwnershipTransferred)
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

func (it *LegacyBurnMintTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *LegacyBurnMintTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LegacyBurnMintTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LegacyBurnMintTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LegacyBurnMintTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolOwnershipTransferredIterator{contract: _LegacyBurnMintTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LegacyBurnMintTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LegacyBurnMintTokenPoolOwnershipTransferred)
				if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*LegacyBurnMintTokenPoolOwnershipTransferred, error) {
	event := new(LegacyBurnMintTokenPoolOwnershipTransferred)
	if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LegacyBurnMintTokenPoolPausedIterator struct {
	Event *LegacyBurnMintTokenPoolPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LegacyBurnMintTokenPoolPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyBurnMintTokenPoolPaused)
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
		it.Event = new(LegacyBurnMintTokenPoolPaused)
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

func (it *LegacyBurnMintTokenPoolPausedIterator) Error() error {
	return it.fail
}

func (it *LegacyBurnMintTokenPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LegacyBurnMintTokenPoolPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*LegacyBurnMintTokenPoolPausedIterator, error) {

	logs, sub, err := _LegacyBurnMintTokenPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolPausedIterator{contract: _LegacyBurnMintTokenPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolPaused) (event.Subscription, error) {

	logs, sub, err := _LegacyBurnMintTokenPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LegacyBurnMintTokenPoolPaused)
				if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) ParsePaused(log types.Log) (*LegacyBurnMintTokenPoolPaused, error) {
	event := new(LegacyBurnMintTokenPoolPaused)
	if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LegacyBurnMintTokenPoolReleasedIterator struct {
	Event *LegacyBurnMintTokenPoolReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LegacyBurnMintTokenPoolReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyBurnMintTokenPoolReleased)
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
		it.Event = new(LegacyBurnMintTokenPoolReleased)
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

func (it *LegacyBurnMintTokenPoolReleasedIterator) Error() error {
	return it.fail
}

func (it *LegacyBurnMintTokenPoolReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LegacyBurnMintTokenPoolReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LegacyBurnMintTokenPoolReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LegacyBurnMintTokenPool.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolReleasedIterator{contract: _LegacyBurnMintTokenPool.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LegacyBurnMintTokenPool.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LegacyBurnMintTokenPoolReleased)
				if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
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

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) ParseReleased(log types.Log) (*LegacyBurnMintTokenPoolReleased, error) {
	event := new(LegacyBurnMintTokenPoolReleased)
	if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LegacyBurnMintTokenPoolUnpausedIterator struct {
	Event *LegacyBurnMintTokenPoolUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LegacyBurnMintTokenPoolUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyBurnMintTokenPoolUnpaused)
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
		it.Event = new(LegacyBurnMintTokenPoolUnpaused)
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

func (it *LegacyBurnMintTokenPoolUnpausedIterator) Error() error {
	return it.fail
}

func (it *LegacyBurnMintTokenPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LegacyBurnMintTokenPoolUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LegacyBurnMintTokenPoolUnpausedIterator, error) {

	logs, sub, err := _LegacyBurnMintTokenPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LegacyBurnMintTokenPoolUnpausedIterator{contract: _LegacyBurnMintTokenPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _LegacyBurnMintTokenPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LegacyBurnMintTokenPoolUnpaused)
				if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPoolFilterer) ParseUnpaused(log types.Log) (*LegacyBurnMintTokenPoolUnpaused, error) {
	event := new(LegacyBurnMintTokenPoolUnpaused)
	if err := _LegacyBurnMintTokenPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _LegacyBurnMintTokenPool.abi.Events["AllowListAdd"].ID:
		return _LegacyBurnMintTokenPool.ParseAllowListAdd(log)
	case _LegacyBurnMintTokenPool.abi.Events["AllowListRemove"].ID:
		return _LegacyBurnMintTokenPool.ParseAllowListRemove(log)
	case _LegacyBurnMintTokenPool.abi.Events["Burned"].ID:
		return _LegacyBurnMintTokenPool.ParseBurned(log)
	case _LegacyBurnMintTokenPool.abi.Events["Locked"].ID:
		return _LegacyBurnMintTokenPool.ParseLocked(log)
	case _LegacyBurnMintTokenPool.abi.Events["Minted"].ID:
		return _LegacyBurnMintTokenPool.ParseMinted(log)
	case _LegacyBurnMintTokenPool.abi.Events["OffRampAllowanceSet"].ID:
		return _LegacyBurnMintTokenPool.ParseOffRampAllowanceSet(log)
	case _LegacyBurnMintTokenPool.abi.Events["OnRampAllowanceSet"].ID:
		return _LegacyBurnMintTokenPool.ParseOnRampAllowanceSet(log)
	case _LegacyBurnMintTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _LegacyBurnMintTokenPool.ParseOwnershipTransferRequested(log)
	case _LegacyBurnMintTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _LegacyBurnMintTokenPool.ParseOwnershipTransferred(log)
	case _LegacyBurnMintTokenPool.abi.Events["Paused"].ID:
		return _LegacyBurnMintTokenPool.ParsePaused(log)
	case _LegacyBurnMintTokenPool.abi.Events["Released"].ID:
		return _LegacyBurnMintTokenPool.ParseReleased(log)
	case _LegacyBurnMintTokenPool.abi.Events["Unpaused"].ID:
		return _LegacyBurnMintTokenPool.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (LegacyBurnMintTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (LegacyBurnMintTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (LegacyBurnMintTokenPoolBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
}

func (LegacyBurnMintTokenPoolLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (LegacyBurnMintTokenPoolMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (LegacyBurnMintTokenPoolOffRampAllowanceSet) Topic() common.Hash {
	return common.HexToHash("0xd8c3333ded377884ced3869cd0bcb9be54ea664076df1f5d39c4689120313648")
}

func (LegacyBurnMintTokenPoolOnRampAllowanceSet) Topic() common.Hash {
	return common.HexToHash("0xbceff8f229c6dfcbf8bdcfb18726b84b0fd249b4803deb3948ff34d904013662")
}

func (LegacyBurnMintTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (LegacyBurnMintTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (LegacyBurnMintTokenPoolPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (LegacyBurnMintTokenPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (LegacyBurnMintTokenPoolUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_LegacyBurnMintTokenPool *LegacyBurnMintTokenPool) Address() common.Address {
	return _LegacyBurnMintTokenPool.address
}

type LegacyBurnMintTokenPoolInterface interface {
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

	LockOrBurn(opts *bind.TransactOpts, originalSender common.Address, arg1 []byte, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, arg0 []byte, receiver common.Address, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*LegacyBurnMintTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*LegacyBurnMintTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*LegacyBurnMintTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*LegacyBurnMintTokenPoolAllowListRemove, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*LegacyBurnMintTokenPoolBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*LegacyBurnMintTokenPoolBurned, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*LegacyBurnMintTokenPoolLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*LegacyBurnMintTokenPoolLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LegacyBurnMintTokenPoolMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*LegacyBurnMintTokenPoolMinted, error)

	FilterOffRampAllowanceSet(opts *bind.FilterOpts) (*LegacyBurnMintTokenPoolOffRampAllowanceSetIterator, error)

	WatchOffRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolOffRampAllowanceSet) (event.Subscription, error)

	ParseOffRampAllowanceSet(log types.Log) (*LegacyBurnMintTokenPoolOffRampAllowanceSet, error)

	FilterOnRampAllowanceSet(opts *bind.FilterOpts) (*LegacyBurnMintTokenPoolOnRampAllowanceSetIterator, error)

	WatchOnRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolOnRampAllowanceSet) (event.Subscription, error)

	ParseOnRampAllowanceSet(log types.Log) (*LegacyBurnMintTokenPoolOnRampAllowanceSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LegacyBurnMintTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*LegacyBurnMintTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LegacyBurnMintTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*LegacyBurnMintTokenPoolOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*LegacyBurnMintTokenPoolPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*LegacyBurnMintTokenPoolPaused, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LegacyBurnMintTokenPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*LegacyBurnMintTokenPoolReleased, error)

	FilterUnpaused(opts *bind.FilterOpts) (*LegacyBurnMintTokenPoolUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LegacyBurnMintTokenPoolUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*LegacyBurnMintTokenPoolUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
