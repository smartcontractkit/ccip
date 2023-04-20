// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package forwarder

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

type IForwarderForwardRequest struct {
	From           common.Address
	Target         common.Address
	Nonce          *big.Int
	Data           []byte
	ValidUntilTime *big.Int
}

var ForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ForwardFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"}],\"name\":\"InvalidTypeName\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unavailbleNonce\",\"type\":\"uint256\"}],\"name\":\"NonceUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"RequestExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnregisteredDomainSeparator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnregisteredTypeHash\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"domainValue\",\"type\":\"bytes\"}],\"name\":\"DomainRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnValue\",\"type\":\"bytes\"}],\"name\":\"ForwardSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"typeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeStr\",\"type\":\"string\"}],\"name\":\"RequestTypeRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_TYPE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENERIC_PARAMS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntilTime\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"}],\"name\":\"_getEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntilTime\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"registerDomainSeparator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"typeSuffix\",\"type\":\"string\"}],\"name\":\"registerRequestType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"s_domains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"s_typeHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntilTime\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b503380600081620000695760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156200009c576200009c81620000f1565b50505060006040518060800160405280604b815260200162001c68604b9139604051602001620000cd919062000234565b60408051601f198184030181529190529050620000ea816200019c565b50620002ad565b336001600160a01b038216036200014b5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000060565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b805160208083019190912060008181526002909252604091829020805460ff19166001179055905181907f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb20290620001f590859062000278565b60405180910390a25050565b60005b838110156200021e57818101518382015260200162000204565b838111156200022e576000848401525b50505050565b6e08cdee4eec2e4c8a4cae2eacae6e85608b1b8152600082516200026081600f85016020870162000201565b602960f81b600f939091019283015250601001919050565b60208152600082518060208401526200029981604085016020870162000201565b601f01601f19169190910160400192915050565b6119ab80620002bd6000396000f3fe6080604052600436106100e15760003560e01c80638da5cb5b1161007f578063b5aa376711610059578063b5aa376714610271578063c3f28abd14610291578063d9210be5146102a6578063f2fde38b146102c657600080fd5b80638da5cb5b146101fc5780639a0e8929146102315780639c7b45921461025157600080fd5b806371714e17116100bb57806371714e171461017457806379ba5097146101945780637c954790146101ab5780638191a616146101cc57600080fd5b806301ffc9a7146100ed578063066a310c146101225780630b9623f11461014457600080fd5b366100e857005b600080fd5b3480156100f957600080fd5b5061010d610108366004611316565b6102e6565b60405190151581526020015b60405180910390f35b34801561012e57600080fd5b5061013761037f565b60405161011991906113d9565b34801561015057600080fd5b5061010d61015f3660046113ec565b60036020526000908152604090205460ff1681565b34801561018057600080fd5b5061013761018f36600461145f565b61039b565b3480156101a057600080fd5b506101a9610443565b005b6101be6101b93660046114cf565b610545565b604051610119929190611577565b3480156101d857600080fd5b5061010d6101e73660046113ec565b60026020526000908152604090205460ff1681565b34801561020857600080fd5b5060005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610119565b34801561023d57600080fd5b506101a961024c3660046114cf565b6107d9565b34801561025d57600080fd5b506101a961026c36600461159a565b6107fa565b34801561027d57600080fd5b5061013761028c36600461159a565b610899565b34801561029d57600080fd5b50610137610917565b3480156102b257600080fd5b506101a96102c136600461159a565b610933565b3480156102d257600080fd5b506101a96102e13660046115ed565b610a8d565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167fa3c180ce00000000000000000000000000000000000000000000000000000000148061037957507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6040518060800160405280604b8152602001611902604b913981565b6060836103ab60208701876115ed565b73ffffffffffffffffffffffffffffffffffffffff166103d160408801602089016115ed565b73ffffffffffffffffffffffffffffffffffffffff1660408801356103f960608a018a611623565b604051610407929190611688565b60405190819003812061042a959493929160808c0135908a908a90602001611698565b6040516020818303038152906040529050949350505050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146104c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6000606061055889898989898989610aa1565b61056189610c88565b608089013515801590610578575042896080013511155b156105bb576040517f6bdd243600000000000000000000000000000000000000000000000000000000815242600482015260808a013560248201526044016104c0565b60006105ca60608b018b611623565b6105d760208d018d6115ed565b6040516020016105e9939291906116d5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152909150610629908b0160208c016115ed565b73ffffffffffffffffffffffffffffffffffffffff168160405161064d919061170e565b6000604051808303816000865af19150503d806000811461068a576040519150601f19603f3d011682016040523d82523d6000602084013e61068f565b606091505b5090935091508261073357815160000361072b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f466f727761726465642063616c6c20726576657274656420776974686f75742060448201527f726561736f6e000000000000000000000000000000000000000000000000000060648201526084016104c0565b815182602001fd5b8861074460408c0160208d016115ed565b73ffffffffffffffffffffffffffffffffffffffff1661076760208d018d6115ed565b73ffffffffffffffffffffffffffffffffffffffff167f94e72b9bdade69580d865c1293804239e91513c10e675ca3345fc848f9daba2e8d604001358e80606001906107b39190611623565b886040516107c49493929190611773565b60405180910390a45097509795505050505050565b6107e287610d8a565b6107f187878787878787610aa1565b50505050505050565b610802610e12565b600061081085858585610899565b80516020808301919091206000818152600390925260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590519192509081907f4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8906108899085906113d9565b60405180910390a2505050505050565b606060405180608001604052806052815260200161194d605291398051906020012085856040516108cb929190611688565b604051809103902084846040516108e3929190611688565b60408051918290038220602083019490945281019190915260608101919091524660808201523060a082015260c00161042a565b60405180608001604052806052815260200161194d6052913981565b60005b83811015610a38576000858583818110610952576109526117aa565b909101357fff00000000000000000000000000000000000000000000000000000000000000169150507f28000000000000000000000000000000000000000000000000000000000000008114806109ea57507f29000000000000000000000000000000000000000000000000000000000000007fff000000000000000000000000000000000000000000000000000000000000008216145b15610a255785856040517f91ea68d00000000000000000000000000000000000000000000000000000000081526004016104c09291906117d9565b5080610a30816117ed565b915050610936565b50600084846040518060800160405280604b8152602001611902604b91398585604051602001610a6c95949392919061184c565b6040516020818303038152906040529050610a8681610e95565b5050505050565b610a95610e12565b610a9e81610f16565b50565b60008681526003602052604090205460ff16610ae9576040517faa54b09400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008581526002602052604090205460ff16610b31576040517fbb6d969a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600086610b408988888861039b565b8051602091820120604051610b879392017f190100000000000000000000000000000000000000000000000000000000000081526002810192909252602282015260420190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012090503215801590610c475750610bd660208901896115ed565b73ffffffffffffffffffffffffffffffffffffffff16610c2e84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250869392505061100b9050565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610c7e576040517f73a8ee1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050505050565b60046000610c9960208401846115ed565b73ffffffffffffffffffffffffffffffffffffffff1681526020808201929092526040908101600090812084830135825290925290205460ff1615610d1057604080517f8d7cf7220000000000000000000000000000000000000000000000000000000081529082013560048201526024016104c0565b600160046000610d2360208501856115ed565b73ffffffffffffffffffffffffffffffffffffffff168152602080820192909252604090810160009081209482013581529390915290912080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b60046000610d9b60208401846115ed565b73ffffffffffffffffffffffffffffffffffffffff1681526020808201929092526040908101600090812084830135825290925290205460ff1615610a9e57604080517f8d7cf7220000000000000000000000000000000000000000000000000000000081529082013560048201526024016104c0565b60005473ffffffffffffffffffffffffffffffffffffffff163314610e93576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016104c0565b565b80516020808301919091206000818152600290925260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055905181907f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb20290610f0a9085906113d9565b60405180910390a25050565b3373ffffffffffffffffffffffffffffffffffffffff821603610f95576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016104c0565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600080600061101a858561102f565b9150915061102781611074565b509392505050565b60008082516041036110655760208301516040840151606085015160001a61105987828585611227565b9450945050505061106d565b506000905060025b9250929050565b6000816004811115611088576110886118d2565b036110905750565b60018160048111156110a4576110a46118d2565b0361110b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016104c0565b600281600481111561111f5761111f6118d2565b03611186576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016104c0565b600381600481111561119a5761119a6118d2565b03610a9e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016104c0565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561125e575060009050600361130d565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156112b2573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166113065760006001925092505061130d565b9150600090505b94509492505050565b60006020828403121561132857600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461135857600080fd5b9392505050565b60005b8381101561137a578181015183820152602001611362565b83811115611389576000848401525b50505050565b600081518084526113a781602086016020860161135f565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611358602083018461138f565b6000602082840312156113fe57600080fd5b5035919050565b600060a0828403121561141757600080fd5b50919050565b60008083601f84011261142f57600080fd5b50813567ffffffffffffffff81111561144757600080fd5b60208301915083602082850101111561106d57600080fd5b6000806000806060858703121561147557600080fd5b843567ffffffffffffffff8082111561148d57600080fd5b61149988838901611405565b95506020870135945060408701359150808211156114b657600080fd5b506114c38782880161141d565b95989497509550505050565b600080600080600080600060a0888a0312156114ea57600080fd5b873567ffffffffffffffff8082111561150257600080fd5b61150e8b838c01611405565b985060208a0135975060408a0135965060608a013591508082111561153257600080fd5b61153e8b838c0161141d565b909650945060808a013591508082111561155757600080fd5b506115648a828b0161141d565b989b979a50959850939692959293505050565b8215158152604060208201526000611592604083018461138f565b949350505050565b600080600080604085870312156115b057600080fd5b843567ffffffffffffffff808211156115c857600080fd5b6115d48883890161141d565b909650945060208701359150808211156114b657600080fd5b6000602082840312156115ff57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461135857600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261165857600080fd5b83018035915067ffffffffffffffff82111561167357600080fd5b60200191503681900382131561106d57600080fd5b8183823760009101908152919050565b8881528760208201528660408201528560608201528460808201528360a0820152818360c08301376000910160c001908152979650505050505050565b8284823760609190911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000169101908152601401919050565b6000825161172081846020870161135f565b9190910192915050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b84815260606020820152600061178d60608301858761172a565b828103604084015261179f818561138f565b979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60208152600061159260208301848661172a565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611845577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b8486823760008582017f28000000000000000000000000000000000000000000000000000000000000008152855161188b816001840160208a0161135f565b7f2c00000000000000000000000000000000000000000000000000000000000000600192909101918201528385600283013760009301600201928352509095945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfe616464726573732066726f6d2c61646472657373207461726765742c75696e74323536206e6f6e63652c627974657320646174612c75696e743235362076616c6964556e74696c54696d65454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e747261637429a164736f6c634300080f000a616464726573732066726f6d2c61646472657373207461726765742c75696e74323536206e6f6e63652c627974657320646174612c75696e743235362076616c6964556e74696c54696d65",
}

var ForwarderABI = ForwarderMetaData.ABI

var ForwarderBin = ForwarderMetaData.Bin

func DeployForwarder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Forwarder, error) {
	parsed, err := ForwarderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ForwarderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Forwarder{ForwarderCaller: ForwarderCaller{contract: contract}, ForwarderTransactor: ForwarderTransactor{contract: contract}, ForwarderFilterer: ForwarderFilterer{contract: contract}}, nil
}

type Forwarder struct {
	address common.Address
	abi     abi.ABI
	ForwarderCaller
	ForwarderTransactor
	ForwarderFilterer
}

type ForwarderCaller struct {
	contract *bind.BoundContract
}

type ForwarderTransactor struct {
	contract *bind.BoundContract
}

type ForwarderFilterer struct {
	contract *bind.BoundContract
}

type ForwarderSession struct {
	Contract     *Forwarder
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ForwarderCallerSession struct {
	Contract *ForwarderCaller
	CallOpts bind.CallOpts
}

type ForwarderTransactorSession struct {
	Contract     *ForwarderTransactor
	TransactOpts bind.TransactOpts
}

type ForwarderRaw struct {
	Contract *Forwarder
}

type ForwarderCallerRaw struct {
	Contract *ForwarderCaller
}

type ForwarderTransactorRaw struct {
	Contract *ForwarderTransactor
}

func NewForwarder(address common.Address, backend bind.ContractBackend) (*Forwarder, error) {
	abi, err := abi.JSON(strings.NewReader(ForwarderABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Forwarder{address: address, abi: abi, ForwarderCaller: ForwarderCaller{contract: contract}, ForwarderTransactor: ForwarderTransactor{contract: contract}, ForwarderFilterer: ForwarderFilterer{contract: contract}}, nil
}

func NewForwarderCaller(address common.Address, caller bind.ContractCaller) (*ForwarderCaller, error) {
	contract, err := bindForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderCaller{contract: contract}, nil
}

func NewForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*ForwarderTransactor, error) {
	contract, err := bindForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderTransactor{contract: contract}, nil
}

func NewForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*ForwarderFilterer, error) {
	contract, err := bindForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForwarderFilterer{contract: contract}, nil
}

func bindForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_Forwarder *ForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forwarder.Contract.ForwarderCaller.contract.Call(opts, result, method, params...)
}

func (_Forwarder *ForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.Contract.ForwarderTransactor.contract.Transfer(opts)
}

func (_Forwarder *ForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forwarder.Contract.ForwarderTransactor.contract.Transact(opts, method, params...)
}

func (_Forwarder *ForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forwarder.Contract.contract.Call(opts, result, method, params...)
}

func (_Forwarder *ForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.Contract.contract.Transfer(opts)
}

func (_Forwarder *ForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forwarder.Contract.contract.Transact(opts, method, params...)
}

func (_Forwarder *ForwarderCaller) EIP712DOMAINTYPE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "EIP712_DOMAIN_TYPE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Forwarder *ForwarderSession) EIP712DOMAINTYPE() (string, error) {
	return _Forwarder.Contract.EIP712DOMAINTYPE(&_Forwarder.CallOpts)
}

func (_Forwarder *ForwarderCallerSession) EIP712DOMAINTYPE() (string, error) {
	return _Forwarder.Contract.EIP712DOMAINTYPE(&_Forwarder.CallOpts)
}

func (_Forwarder *ForwarderCaller) GENERICPARAMS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "GENERIC_PARAMS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Forwarder *ForwarderSession) GENERICPARAMS() (string, error) {
	return _Forwarder.Contract.GENERICPARAMS(&_Forwarder.CallOpts)
}

func (_Forwarder *ForwarderCallerSession) GENERICPARAMS() (string, error) {
	return _Forwarder.Contract.GENERICPARAMS(&_Forwarder.CallOpts)
}

func (_Forwarder *ForwarderCaller) GetEncoded(opts *bind.CallOpts, req IForwarderForwardRequest, requestTypeHash [32]byte, suffixData []byte) ([]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "_getEncoded", req, requestTypeHash, suffixData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_Forwarder *ForwarderSession) GetEncoded(req IForwarderForwardRequest, requestTypeHash [32]byte, suffixData []byte) ([]byte, error) {
	return _Forwarder.Contract.GetEncoded(&_Forwarder.CallOpts, req, requestTypeHash, suffixData)
}

func (_Forwarder *ForwarderCallerSession) GetEncoded(req IForwarderForwardRequest, requestTypeHash [32]byte, suffixData []byte) ([]byte, error) {
	return _Forwarder.Contract.GetEncoded(&_Forwarder.CallOpts, req, requestTypeHash, suffixData)
}

func (_Forwarder *ForwarderCaller) GetDomainSeparator(opts *bind.CallOpts, name string, version string) ([]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getDomainSeparator", name, version)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_Forwarder *ForwarderSession) GetDomainSeparator(name string, version string) ([]byte, error) {
	return _Forwarder.Contract.GetDomainSeparator(&_Forwarder.CallOpts, name, version)
}

func (_Forwarder *ForwarderCallerSession) GetDomainSeparator(name string, version string) ([]byte, error) {
	return _Forwarder.Contract.GetDomainSeparator(&_Forwarder.CallOpts, name, version)
}

func (_Forwarder *ForwarderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Forwarder *ForwarderSession) Owner() (common.Address, error) {
	return _Forwarder.Contract.Owner(&_Forwarder.CallOpts)
}

func (_Forwarder *ForwarderCallerSession) Owner() (common.Address, error) {
	return _Forwarder.Contract.Owner(&_Forwarder.CallOpts)
}

func (_Forwarder *ForwarderCaller) SDomains(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "s_domains", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Forwarder *ForwarderSession) SDomains(arg0 [32]byte) (bool, error) {
	return _Forwarder.Contract.SDomains(&_Forwarder.CallOpts, arg0)
}

func (_Forwarder *ForwarderCallerSession) SDomains(arg0 [32]byte) (bool, error) {
	return _Forwarder.Contract.SDomains(&_Forwarder.CallOpts, arg0)
}

func (_Forwarder *ForwarderCaller) STypeHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "s_typeHashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Forwarder *ForwarderSession) STypeHashes(arg0 [32]byte) (bool, error) {
	return _Forwarder.Contract.STypeHashes(&_Forwarder.CallOpts, arg0)
}

func (_Forwarder *ForwarderCallerSession) STypeHashes(arg0 [32]byte) (bool, error) {
	return _Forwarder.Contract.STypeHashes(&_Forwarder.CallOpts, arg0)
}

func (_Forwarder *ForwarderCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Forwarder *ForwarderSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Forwarder.Contract.SupportsInterface(&_Forwarder.CallOpts, interfaceId)
}

func (_Forwarder *ForwarderCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Forwarder.Contract.SupportsInterface(&_Forwarder.CallOpts, interfaceId)
}

func (_Forwarder *ForwarderCaller) Verify(opts *bind.CallOpts, req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) error {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "verify", req, domainSeparator, requestTypeHash, suffixData, sig)

	if err != nil {
		return err
	}

	return err

}

func (_Forwarder *ForwarderSession) Verify(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) error {
	return _Forwarder.Contract.Verify(&_Forwarder.CallOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

func (_Forwarder *ForwarderCallerSession) Verify(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) error {
	return _Forwarder.Contract.Verify(&_Forwarder.CallOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

func (_Forwarder *ForwarderTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "acceptOwnership")
}

func (_Forwarder *ForwarderSession) AcceptOwnership() (*types.Transaction, error) {
	return _Forwarder.Contract.AcceptOwnership(&_Forwarder.TransactOpts)
}

func (_Forwarder *ForwarderTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Forwarder.Contract.AcceptOwnership(&_Forwarder.TransactOpts)
}

func (_Forwarder *ForwarderTransactor) Execute(opts *bind.TransactOpts, req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "execute", req, domainSeparator, requestTypeHash, suffixData, sig)
}

func (_Forwarder *ForwarderSession) Execute(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.Execute(&_Forwarder.TransactOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

func (_Forwarder *ForwarderTransactorSession) Execute(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.Execute(&_Forwarder.TransactOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

func (_Forwarder *ForwarderTransactor) RegisterDomainSeparator(opts *bind.TransactOpts, name string, version string) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "registerDomainSeparator", name, version)
}

func (_Forwarder *ForwarderSession) RegisterDomainSeparator(name string, version string) (*types.Transaction, error) {
	return _Forwarder.Contract.RegisterDomainSeparator(&_Forwarder.TransactOpts, name, version)
}

func (_Forwarder *ForwarderTransactorSession) RegisterDomainSeparator(name string, version string) (*types.Transaction, error) {
	return _Forwarder.Contract.RegisterDomainSeparator(&_Forwarder.TransactOpts, name, version)
}

func (_Forwarder *ForwarderTransactor) RegisterRequestType(opts *bind.TransactOpts, typeName string, typeSuffix string) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "registerRequestType", typeName, typeSuffix)
}

func (_Forwarder *ForwarderSession) RegisterRequestType(typeName string, typeSuffix string) (*types.Transaction, error) {
	return _Forwarder.Contract.RegisterRequestType(&_Forwarder.TransactOpts, typeName, typeSuffix)
}

func (_Forwarder *ForwarderTransactorSession) RegisterRequestType(typeName string, typeSuffix string) (*types.Transaction, error) {
	return _Forwarder.Contract.RegisterRequestType(&_Forwarder.TransactOpts, typeName, typeSuffix)
}

func (_Forwarder *ForwarderTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "transferOwnership", to)
}

func (_Forwarder *ForwarderSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Forwarder.Contract.TransferOwnership(&_Forwarder.TransactOpts, to)
}

func (_Forwarder *ForwarderTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Forwarder.Contract.TransferOwnership(&_Forwarder.TransactOpts, to)
}

func (_Forwarder *ForwarderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.contract.RawTransact(opts, nil)
}

func (_Forwarder *ForwarderSession) Receive() (*types.Transaction, error) {
	return _Forwarder.Contract.Receive(&_Forwarder.TransactOpts)
}

func (_Forwarder *ForwarderTransactorSession) Receive() (*types.Transaction, error) {
	return _Forwarder.Contract.Receive(&_Forwarder.TransactOpts)
}

type ForwarderDomainRegisteredIterator struct {
	Event *ForwarderDomainRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ForwarderDomainRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderDomainRegistered)
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
		it.Event = new(ForwarderDomainRegistered)
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

func (it *ForwarderDomainRegisteredIterator) Error() error {
	return it.fail
}

func (it *ForwarderDomainRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ForwarderDomainRegistered struct {
	DomainSeparator [32]byte
	DomainValue     []byte
	Raw             types.Log
}

func (_Forwarder *ForwarderFilterer) FilterDomainRegistered(opts *bind.FilterOpts, domainSeparator [][32]byte) (*ForwarderDomainRegisteredIterator, error) {

	var domainSeparatorRule []interface{}
	for _, domainSeparatorItem := range domainSeparator {
		domainSeparatorRule = append(domainSeparatorRule, domainSeparatorItem)
	}

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "DomainRegistered", domainSeparatorRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderDomainRegisteredIterator{contract: _Forwarder.contract, event: "DomainRegistered", logs: logs, sub: sub}, nil
}

func (_Forwarder *ForwarderFilterer) WatchDomainRegistered(opts *bind.WatchOpts, sink chan<- *ForwarderDomainRegistered, domainSeparator [][32]byte) (event.Subscription, error) {

	var domainSeparatorRule []interface{}
	for _, domainSeparatorItem := range domainSeparator {
		domainSeparatorRule = append(domainSeparatorRule, domainSeparatorItem)
	}

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "DomainRegistered", domainSeparatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ForwarderDomainRegistered)
				if err := _Forwarder.contract.UnpackLog(event, "DomainRegistered", log); err != nil {
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

func (_Forwarder *ForwarderFilterer) ParseDomainRegistered(log types.Log) (*ForwarderDomainRegistered, error) {
	event := new(ForwarderDomainRegistered)
	if err := _Forwarder.contract.UnpackLog(event, "DomainRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ForwarderForwardSucceededIterator struct {
	Event *ForwarderForwardSucceeded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ForwarderForwardSucceededIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderForwardSucceeded)
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
		it.Event = new(ForwarderForwardSucceeded)
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

func (it *ForwarderForwardSucceededIterator) Error() error {
	return it.fail
}

func (it *ForwarderForwardSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ForwarderForwardSucceeded struct {
	From            common.Address
	Target          common.Address
	DomainSeparator [32]byte
	Nonce           *big.Int
	Data            []byte
	ReturnValue     []byte
	Raw             types.Log
}

func (_Forwarder *ForwarderFilterer) FilterForwardSucceeded(opts *bind.FilterOpts, from []common.Address, target []common.Address, domainSeparator [][32]byte) (*ForwarderForwardSucceededIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var domainSeparatorRule []interface{}
	for _, domainSeparatorItem := range domainSeparator {
		domainSeparatorRule = append(domainSeparatorRule, domainSeparatorItem)
	}

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "ForwardSucceeded", fromRule, targetRule, domainSeparatorRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderForwardSucceededIterator{contract: _Forwarder.contract, event: "ForwardSucceeded", logs: logs, sub: sub}, nil
}

func (_Forwarder *ForwarderFilterer) WatchForwardSucceeded(opts *bind.WatchOpts, sink chan<- *ForwarderForwardSucceeded, from []common.Address, target []common.Address, domainSeparator [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var domainSeparatorRule []interface{}
	for _, domainSeparatorItem := range domainSeparator {
		domainSeparatorRule = append(domainSeparatorRule, domainSeparatorItem)
	}

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "ForwardSucceeded", fromRule, targetRule, domainSeparatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ForwarderForwardSucceeded)
				if err := _Forwarder.contract.UnpackLog(event, "ForwardSucceeded", log); err != nil {
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

func (_Forwarder *ForwarderFilterer) ParseForwardSucceeded(log types.Log) (*ForwarderForwardSucceeded, error) {
	event := new(ForwarderForwardSucceeded)
	if err := _Forwarder.contract.UnpackLog(event, "ForwardSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ForwarderOwnershipTransferRequestedIterator struct {
	Event *ForwarderOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ForwarderOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderOwnershipTransferRequested)
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
		it.Event = new(ForwarderOwnershipTransferRequested)
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

func (it *ForwarderOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ForwarderOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ForwarderOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Forwarder *ForwarderFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ForwarderOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderOwnershipTransferRequestedIterator{contract: _Forwarder.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_Forwarder *ForwarderFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ForwarderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ForwarderOwnershipTransferRequested)
				if err := _Forwarder.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_Forwarder *ForwarderFilterer) ParseOwnershipTransferRequested(log types.Log) (*ForwarderOwnershipTransferRequested, error) {
	event := new(ForwarderOwnershipTransferRequested)
	if err := _Forwarder.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ForwarderOwnershipTransferredIterator struct {
	Event *ForwarderOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ForwarderOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderOwnershipTransferred)
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
		it.Event = new(ForwarderOwnershipTransferred)
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

func (it *ForwarderOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ForwarderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ForwarderOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Forwarder *ForwarderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ForwarderOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderOwnershipTransferredIterator{contract: _Forwarder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Forwarder *ForwarderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ForwarderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ForwarderOwnershipTransferred)
				if err := _Forwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_Forwarder *ForwarderFilterer) ParseOwnershipTransferred(log types.Log) (*ForwarderOwnershipTransferred, error) {
	event := new(ForwarderOwnershipTransferred)
	if err := _Forwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ForwarderRequestTypeRegisteredIterator struct {
	Event *ForwarderRequestTypeRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ForwarderRequestTypeRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderRequestTypeRegistered)
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
		it.Event = new(ForwarderRequestTypeRegistered)
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

func (it *ForwarderRequestTypeRegisteredIterator) Error() error {
	return it.fail
}

func (it *ForwarderRequestTypeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ForwarderRequestTypeRegistered struct {
	TypeHash [32]byte
	TypeStr  string
	Raw      types.Log
}

func (_Forwarder *ForwarderFilterer) FilterRequestTypeRegistered(opts *bind.FilterOpts, typeHash [][32]byte) (*ForwarderRequestTypeRegisteredIterator, error) {

	var typeHashRule []interface{}
	for _, typeHashItem := range typeHash {
		typeHashRule = append(typeHashRule, typeHashItem)
	}

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "RequestTypeRegistered", typeHashRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderRequestTypeRegisteredIterator{contract: _Forwarder.contract, event: "RequestTypeRegistered", logs: logs, sub: sub}, nil
}

func (_Forwarder *ForwarderFilterer) WatchRequestTypeRegistered(opts *bind.WatchOpts, sink chan<- *ForwarderRequestTypeRegistered, typeHash [][32]byte) (event.Subscription, error) {

	var typeHashRule []interface{}
	for _, typeHashItem := range typeHash {
		typeHashRule = append(typeHashRule, typeHashItem)
	}

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "RequestTypeRegistered", typeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ForwarderRequestTypeRegistered)
				if err := _Forwarder.contract.UnpackLog(event, "RequestTypeRegistered", log); err != nil {
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

func (_Forwarder *ForwarderFilterer) ParseRequestTypeRegistered(log types.Log) (*ForwarderRequestTypeRegistered, error) {
	event := new(ForwarderRequestTypeRegistered)
	if err := _Forwarder.contract.UnpackLog(event, "RequestTypeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Forwarder *Forwarder) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Forwarder.abi.Events["DomainRegistered"].ID:
		return _Forwarder.ParseDomainRegistered(log)
	case _Forwarder.abi.Events["ForwardSucceeded"].ID:
		return _Forwarder.ParseForwardSucceeded(log)
	case _Forwarder.abi.Events["OwnershipTransferRequested"].ID:
		return _Forwarder.ParseOwnershipTransferRequested(log)
	case _Forwarder.abi.Events["OwnershipTransferred"].ID:
		return _Forwarder.ParseOwnershipTransferred(log)
	case _Forwarder.abi.Events["RequestTypeRegistered"].ID:
		return _Forwarder.ParseRequestTypeRegistered(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ForwarderDomainRegistered) Topic() common.Hash {
	return common.HexToHash("0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8")
}

func (ForwarderForwardSucceeded) Topic() common.Hash {
	return common.HexToHash("0x94e72b9bdade69580d865c1293804239e91513c10e675ca3345fc848f9daba2e")
}

func (ForwarderOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ForwarderOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (ForwarderRequestTypeRegistered) Topic() common.Hash {
	return common.HexToHash("0x64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb202")
}

func (_Forwarder *Forwarder) Address() common.Address {
	return _Forwarder.address
}

type ForwarderInterface interface {
	EIP712DOMAINTYPE(opts *bind.CallOpts) (string, error)

	GENERICPARAMS(opts *bind.CallOpts) (string, error)

	GetEncoded(opts *bind.CallOpts, req IForwarderForwardRequest, requestTypeHash [32]byte, suffixData []byte) ([]byte, error)

	GetDomainSeparator(opts *bind.CallOpts, name string, version string) ([]byte, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SDomains(opts *bind.CallOpts, arg0 [32]byte) (bool, error)

	STypeHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	Verify(opts *bind.CallOpts, req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) error

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (*types.Transaction, error)

	RegisterDomainSeparator(opts *bind.TransactOpts, name string, version string) (*types.Transaction, error)

	RegisterRequestType(opts *bind.TransactOpts, typeName string, typeSuffix string) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterDomainRegistered(opts *bind.FilterOpts, domainSeparator [][32]byte) (*ForwarderDomainRegisteredIterator, error)

	WatchDomainRegistered(opts *bind.WatchOpts, sink chan<- *ForwarderDomainRegistered, domainSeparator [][32]byte) (event.Subscription, error)

	ParseDomainRegistered(log types.Log) (*ForwarderDomainRegistered, error)

	FilterForwardSucceeded(opts *bind.FilterOpts, from []common.Address, target []common.Address, domainSeparator [][32]byte) (*ForwarderForwardSucceededIterator, error)

	WatchForwardSucceeded(opts *bind.WatchOpts, sink chan<- *ForwarderForwardSucceeded, from []common.Address, target []common.Address, domainSeparator [][32]byte) (event.Subscription, error)

	ParseForwardSucceeded(log types.Log) (*ForwarderForwardSucceeded, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ForwarderOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ForwarderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ForwarderOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ForwarderOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ForwarderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ForwarderOwnershipTransferred, error)

	FilterRequestTypeRegistered(opts *bind.FilterOpts, typeHash [][32]byte) (*ForwarderRequestTypeRegisteredIterator, error)

	WatchRequestTypeRegistered(opts *bind.WatchOpts, sink chan<- *ForwarderRequestTypeRegistered, typeHash [][32]byte) (event.Subscription, error)

	ParseRequestTypeRegistered(log types.Log) (*ForwarderRequestTypeRegistered, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
