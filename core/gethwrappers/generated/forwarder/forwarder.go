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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ForwardFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"}],\"name\":\"InvalidTypeName\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"NonceMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"RequestExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnregisteredDomainSeparator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnregisteredTypeHash\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"domainValue\",\"type\":\"bytes\"}],\"name\":\"DomainRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnValue\",\"type\":\"bytes\"}],\"name\":\"ForwardSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"typeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeStr\",\"type\":\"string\"}],\"name\":\"RequestTypeRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_TYPE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENERIC_PARAMS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntilTime\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"}],\"name\":\"_getEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntilTime\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"registerDomainSeparator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"typeSuffix\",\"type\":\"string\"}],\"name\":\"registerRequestType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"s_domains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"s_typeHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntilTime\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b503380600081620000695760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156200009c576200009c81620000f1565b50505060006040518060800160405280604b815260200162001d00604b9139604051602001620000cd919062000234565b60408051601f198184030181529190529050620000ea816200019c565b50620002ad565b336001600160a01b038216036200014b5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000060565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b805160208083019190912060008181526002909252604091829020805460ff19166001179055905181907f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb20290620001f590859062000278565b60405180910390a25050565b60005b838110156200021e57818101518382015260200162000204565b838111156200022e576000848401525b50505050565b6e08cdee4eec2e4c8a4cae2eacae6e85608b1b8152600082516200026081600f85016020870162000201565b602960f81b600f939091019283015250601001919050565b60208152600082518060208401526200029981604085016020870162000201565b601f01601f19169190910160400192915050565b611a4380620002bd6000396000f3fe6080604052600436106100ec5760003560e01c80638191a6161161008a578063b5aa376711610059578063b5aa3767146102cd578063c3f28abd146102ed578063d9210be514610302578063f2fde38b1461032257600080fd5b80638191a616146102285780638da5cb5b146102585780639a0e89291461028d5780639c7b4592146102ad57600080fd5b80632d0335ab116100c65780632d0335ab1461017f57806371714e17146101d057806379ba5097146101f05780637c9547901461020757600080fd5b806301ffc9a7146100f8578063066a310c1461012d5780630b9623f11461014f57600080fd5b366100f357005b600080fd5b34801561010457600080fd5b5061011861011336600461139b565b610342565b60405190151581526020015b60405180910390f35b34801561013957600080fd5b506101426103db565b604051610124919061145e565b34801561015b57600080fd5b5061011861016a366004611471565b60036020526000908152604090205460ff1681565b34801561018b57600080fd5b506101c261019a36600461148a565b73ffffffffffffffffffffffffffffffffffffffff1660009081526004602052604090205490565b604051908152602001610124565b3480156101dc57600080fd5b506101426101eb36600461151a565b6103f7565b3480156101fc57600080fd5b5061020561049f565b005b61021a61021536600461158a565b6105a1565b604051610124929190611632565b34801561023457600080fd5b50610118610243366004611471565b60026020526000908152604090205460ff1681565b34801561026457600080fd5b5060005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610124565b34801561029957600080fd5b506102056102a836600461158a565b6107dd565b3480156102b957600080fd5b506102056102c8366004611655565b6107fe565b3480156102d957600080fd5b506101426102e8366004611655565b61089d565b3480156102f957600080fd5b5061014261091b565b34801561030e57600080fd5b5061020561031d366004611655565b610937565b34801561032e57600080fd5b5061020561033d36600461148a565b610ac9565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f8ec2b5650000000000000000000000000000000000000000000000000000000014806103d557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6040518060800160405280604b815260200161199a604b913981565b606083610407602087018761148a565b73ffffffffffffffffffffffffffffffffffffffff1661042d604088016020890161148a565b73ffffffffffffffffffffffffffffffffffffffff16604088013561045560608a018a6116a8565b60405161046392919061170d565b604051908190038120610486959493929160808c0135908a908a9060200161171d565b6040516020818303038152906040529050949350505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610525576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b600060606105b489898989898989610add565b6105bd89610cc4565b6080890135158015906105d4575042896080013511155b15610617576040517f6bdd243600000000000000000000000000000000000000000000000000000000815242600482015260808a0135602482015260440161051c565b600061062660608b018b6116a8565b61063360208d018d61148a565b6040516020016106459392919061175a565b604051602081830303815290604052905060008a602001602081019061066b919061148a565b90506106935a73ffffffffffffffffffffffffffffffffffffffff8316906000602086610da5565b90945092508361073657825160000361072e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f466f727761726465642063616c6c20726576657274656420776974686f75742060448201527f726561736f6e0000000000000000000000000000000000000000000000000000606482015260840161051c565b825183602001fd5b8961074760408d0160208e0161148a565b73ffffffffffffffffffffffffffffffffffffffff1661076a60208e018e61148a565b73ffffffffffffffffffffffffffffffffffffffff167f94e72b9bdade69580d865c1293804239e91513c10e675ca3345fc848f9daba2e8e604001358f80606001906107b691906116a8565b896040516107c794939291906117dc565b60405180910390a4505097509795505050505050565b6107e687610e30565b6107f587878787878787610add565b50505050505050565b610806610e97565b60006108148585858561089d565b80516020808301919091206000818152600390925260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590519192509081907f4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d89061088d90859061145e565b60405180910390a2505050505050565b60606040518060800160405280605281526020016119e5605291398051906020012085856040516108cf92919061170d565b604051809103902084846040516108e792919061170d565b60408051918290038220602083019490945281019190915260608101919091524660808201523060a082015260c001610486565b6040518060800160405280605281526020016119e56052913981565b600084848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509394505050505b8151811015610a7357600082828151811061098f5761098f611813565b01602001517fff000000000000000000000000000000000000000000000000000000000000001690507f2800000000000000000000000000000000000000000000000000000000000000811480610a2757507f29000000000000000000000000000000000000000000000000000000000000007fff000000000000000000000000000000000000000000000000000000000000008216145b15610a625786866040517f91ea68d000000000000000000000000000000000000000000000000000000000815260040161051c929190611842565b50610a6c81611856565b9050610972565b50600085856040518060800160405280604b815260200161199a604b91398686604051602001610aa79594939291906118b5565b6040516020818303038152906040529050610ac181610f1a565b505050505050565b610ad1610e97565b610ada81610f9b565b50565b60008681526003602052604090205460ff16610b25576040517faa54b09400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008581526002602052604090205460ff16610b6d576040517fbb6d969a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600086610b7c898888886103f7565b8051602091820120604051610bc39392017f190100000000000000000000000000000000000000000000000000000000000081526002810192909252602282015260420190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012090503215801590610c835750610c12602089018961148a565b73ffffffffffffffffffffffffffffffffffffffff16610c6a84848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525086939250506110909050565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610cba576040517f73a8ee1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050505050565b604081013560046000610cda602085018561148a565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040016000908120805491610d0e83611856565b9190505514610ada5760046000610d28602084018461148a565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481604001356040517f0f0f26af00000000000000000000000000000000000000000000000000000000815260040161051c929190918252602082015260400190565b6000606060008060008661ffff1667ffffffffffffffff811115610dcb57610dcb61193b565b6040519080825280601f01601f191660200182016040528015610df5576020820181803683370190505b5090506000808751602089018b8e8ef191503d925086831115610e16578692505b828152826000602083013e90999098509650505050505050565b604081013560046000610e46602085018561148a565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610ada5760046000610d28602084018461148a565b60005473ffffffffffffffffffffffffffffffffffffffff163314610f18576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161051c565b565b80516020808301919091206000818152600290925260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055905181907f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb20290610f8f90859061145e565b60405180910390a25050565b3373ffffffffffffffffffffffffffffffffffffffff82160361101a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161051c565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600080600061109f85856110b4565b915091506110ac816110f9565b509392505050565b60008082516041036110ea5760208301516040840151606085015160001a6110de878285856112ac565b945094505050506110f2565b506000905060025b9250929050565b600081600481111561110d5761110d61196a565b036111155750565b60018160048111156111295761112961196a565b03611190576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161051c565b60028160048111156111a4576111a461196a565b0361120b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161051c565b600381600481111561121f5761121f61196a565b03610ada576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161051c565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156112e35750600090506003611392565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611337573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661138b57600060019250925050611392565b9150600090505b94509492505050565b6000602082840312156113ad57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146113dd57600080fd5b9392505050565b60005b838110156113ff5781810151838201526020016113e7565b8381111561140e576000848401525b50505050565b6000815180845261142c8160208601602086016113e4565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006113dd6020830184611414565b60006020828403121561148357600080fd5b5035919050565b60006020828403121561149c57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146113dd57600080fd5b600060a082840312156114d257600080fd5b50919050565b60008083601f8401126114ea57600080fd5b50813567ffffffffffffffff81111561150257600080fd5b6020830191508360208285010111156110f257600080fd5b6000806000806060858703121561153057600080fd5b843567ffffffffffffffff8082111561154857600080fd5b611554888389016114c0565b955060208701359450604087013591508082111561157157600080fd5b5061157e878288016114d8565b95989497509550505050565b600080600080600080600060a0888a0312156115a557600080fd5b873567ffffffffffffffff808211156115bd57600080fd5b6115c98b838c016114c0565b985060208a0135975060408a0135965060608a01359150808211156115ed57600080fd5b6115f98b838c016114d8565b909650945060808a013591508082111561161257600080fd5b5061161f8a828b016114d8565b989b979a50959850939692959293505050565b821515815260406020820152600061164d6040830184611414565b949350505050565b6000806000806040858703121561166b57600080fd5b843567ffffffffffffffff8082111561168357600080fd5b61168f888389016114d8565b9096509450602087013591508082111561157157600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126116dd57600080fd5b83018035915067ffffffffffffffff8211156116f857600080fd5b6020019150368190038213156110f257600080fd5b8183823760009101908152919050565b8881528760208201528660408201528560608201528460808201528360a0820152818360c08301376000910160c001908152979650505050505050565b8284823760609190911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000169101908152601401919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8481526060602082015260006117f6606083018587611793565b82810360408401526118088185611414565b979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60208152600061164d602083018486611793565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036118ae577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b8486823760008582017f2800000000000000000000000000000000000000000000000000000000000000815285516118f4816001840160208a016113e4565b7f2c00000000000000000000000000000000000000000000000000000000000000600192909101918201528385600283013760009301600201928352509095945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfe616464726573732066726f6d2c61646472657373207461726765742c75696e74323536206e6f6e63652c627974657320646174612c75696e743235362076616c6964556e74696c54696d65454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e747261637429a164736f6c634300080f000a616464726573732066726f6d2c61646472657373207461726765742c75696e74323536206e6f6e63652c627974657320646174612c75696e743235362076616c6964556e74696c54696d65",
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

func (_Forwarder *ForwarderCaller) GetNonce(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getNonce", from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Forwarder *ForwarderSession) GetNonce(from common.Address) (*big.Int, error) {
	return _Forwarder.Contract.GetNonce(&_Forwarder.CallOpts, from)
}

func (_Forwarder *ForwarderCallerSession) GetNonce(from common.Address) (*big.Int, error) {
	return _Forwarder.Contract.GetNonce(&_Forwarder.CallOpts, from)
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

	GetNonce(opts *bind.CallOpts, from common.Address) (*big.Int, error)

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
