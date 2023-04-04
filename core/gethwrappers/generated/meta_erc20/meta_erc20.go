// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package meta_erc20

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
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated"
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
)

var MetaERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_ccip_router\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_test_only_force_cross_chain_transfer\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredBalance\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MustBeTrustedForwarder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawFailure\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTrustedForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"}],\"name\":\"metaTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIForwarder\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"setForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001809380380620018098339810160408190526200003491620001ba565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600180546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be816200010e565b50505060038390553360009081526004602052604090209290925560078054921515600160a01b026001600160a81b03199093166001600160a01b03929092169190911791909117905562000212565b336001600160a01b03821603620001685760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600280546001600160a01b0319166001600160a01b03838116918217909255600154604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b600080600060608486031215620001d057600080fd5b835160208501519093506001600160a01b0381168114620001f057600080fd5b604085015190925080151581146200020757600080fd5b809150509250925092565b6115e780620002226000396000f3fe6080604052600436106101635760003560e01c806379ba5097116100c0578063a9059cbb11610074578063ce1b815f11610059578063ce1b815f14610453578063dd62ed3e1461047e578063f2fde38b146104b657600080fd5b8063a9059cbb14610413578063b9998a241461043357600080fd5b80638da5cb5b116100a55780638da5cb5b1461035e57806395d89b41146103aa5780639dc29fac146103f357600080fd5b806379ba5097146103295780637b1837de1461033e57600080fd5b8063313ce5671161011757806350431ce4116100fc57806350431ce4146102ab578063572b6c05146102c057806370a08231146102fc57600080fd5b8063313ce5671461026457806340c10f191461028b57600080fd5b8063178293441161014857806317829344146101fe57806318160ddd1461022057806323b872dd1461024457600080fd5b806306fdde031461016f578063095ea7b3146101ce57600080fd5b3661016a57005b600080fd5b34801561017b57600080fd5b506101b86040518060400160405280600981526020017f42616e6b546f6b656e000000000000000000000000000000000000000000000081525081565b6040516101c591906112d4565b60405180910390f35b3480156101da57600080fd5b506101ee6101e9366004611310565b6104d6565b60405190151581526020016101c5565b34801561020a57600080fd5b5061021e61021936600461133c565b6104f3565b005b34801561022c57600080fd5b5061023660035481565b6040519081526020016101c5565b34801561025057600080fd5b506101ee61025f36600461138b565b610851565b34801561027057600080fd5b50610279601281565b60405160ff90911681526020016101c5565b34801561029757600080fd5b5061021e6102a6366004611310565b6109c0565b3480156102b757600080fd5b5061021e610ae0565b3480156102cc57600080fd5b506101ee6102db3660046113cc565b60005473ffffffffffffffffffffffffffffffffffffffff91821691161490565b34801561030857600080fd5b506102366103173660046113cc565b60046020526000908152604090205481565b34801561033557600080fd5b5061021e610ba5565b34801561034a57600080fd5b5061021e610359366004611310565b610ca6565b34801561036a57600080fd5b5060015473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c5565b3480156103b657600080fd5b506101b86040518060400160405280600981526020017f42414e4b544f4b454e000000000000000000000000000000000000000000000081525081565b3480156103ff57600080fd5b5061021e61040e366004611310565b610cc3565b34801561041f57600080fd5b506101ee61042e366004611310565b610eb1565b34801561043f57600080fd5b5061021e61044e3660046113cc565b610ebe565b34801561045f57600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff16610385565b34801561048a57600080fd5b506102366104993660046113e9565b600560209081526000928352604080842090915290825290205481565b3480156104c257600080fd5b5061021e6104d13660046113cc565b610f0c565b60006104ea6104e3610f1d565b8484610f7c565b50600192915050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461054b576040517fa2f64cc50000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61055481610fe2565b61056f5761056a610563610f1d565b8484611021565b505050565b604080516001808252818301909252600091816020015b604080518082019091526000808252602082015281526020019060019003908161058657905050905060405180604001604052803073ffffffffffffffffffffffffffffffffffffffff16815260200184815250816000815181106105ed576105ed611422565b60209081029190910101526040805160a0810190915273ffffffffffffffffffffffffffffffffffffffff851660c08201526000908060e081016040516020818303038152906040528152602001604051806020016040528060008152508152602001838152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001610709604051806040016040528062030d408152602001600015158152506040805182516024820152602092830151151560448083019190915282518083039091018152606490910190915290810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f97a657c90000000000000000000000000000000000000000000000000000000017905290565b90526007546040517f20487ded00000000000000000000000000000000000000000000000000000000815291925060009173ffffffffffffffffffffffffffffffffffffffff909116906320487ded906107699087908690600401611451565b602060405180830381865afa158015610786573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107aa9190611563565b6007546040517f96f4e9f900000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff16906396f4e9f99083906108059088908790600401611451565b60206040518083038185885af1158015610823573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906108489190611563565b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526005602052604081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90826108a1610f1d565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146109ab5773ffffffffffffffffffffffffffffffffffffffff841660009081526005602052604081208391610910610f1d565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461095591906115ab565b73ffffffffffffffffffffffffffffffffffffffff8516600090815260056020526040812090610983610f1d565b73ffffffffffffffffffffffffffffffffffffffff1681526020810191909152604001600020555b6109b6848484611021565b5060019392505050565b73ffffffffffffffffffffffffffffffffffffffff8216610a3d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610542565b8060036000828254610a4f91906115c2565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526004602052604081208054839290610a899084906115c2565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b610ae86110f0565b476000610b0a60015473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114610b61576040519150601f19603f3d011682016040523d82523d6000602084013e610b66565b606091505b5050905080610ba1576040517f1a0263ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b60025473ffffffffffffffffffffffffffffffffffffffff163314610c26576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610542565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560028054909116905560405173ffffffffffffffffffffffffffffffffffffffff909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b610cb8610cb1610f1d565b3083611021565b610ba1308383610f7c565b73ffffffffffffffffffffffffffffffffffffffff8216610d66576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610542565b73ffffffffffffffffffffffffffffffffffffffff821660009081526004602052604090205481811015610e1c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610542565b73ffffffffffffffffffffffffffffffffffffffff83166000908152600460205260408120838303905560038054849290610e589084906115ab565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020015b60405180910390a3505050565b60006104ea610563610f1d565b610ec66110f0565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff831617905550565b50565b610f146110f0565b610f0981611173565b600060143610801590610f47575060005473ffffffffffffffffffffffffffffffffffffffff1633145b15610f7757507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec36013560601c90565b503390565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526005602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259101610ea4565b60075460009074010000000000000000000000000000000000000000900460ff161561101057506001919050565b5067ffffffffffffffff1646141590565b73ffffffffffffffffffffffffffffffffffffffff83166000908152600460205260409020546110529082906115ab565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260046020526040808220939093559084168152205461108f9082906115c2565b73ffffffffffffffffffffffffffffffffffffffff80841660008181526004602052604090819020939093559151908516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90610ea49085815260200190565b60015473ffffffffffffffffffffffffffffffffffffffff163314611171576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610542565b565b3373ffffffffffffffffffffffffffffffffffffffff8216036111f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610542565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600154604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b6000815180845260005b8181101561128f57602081850181015186830182015201611273565b818111156112a1576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006112e76020830184611269565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610f0957600080fd5b6000806040838503121561132357600080fd5b823561132e816112ee565b946020939093013593505050565b60008060006060848603121561135157600080fd5b833561135c816112ee565b925060208401359150604084013567ffffffffffffffff8116811461138057600080fd5b809150509250925092565b6000806000606084860312156113a057600080fd5b83356113ab816112ee565b925060208401356113bb816112ee565b929592945050506040919091013590565b6000602082840312156113de57600080fd5b81356112e7816112ee565b600080604083850312156113fc57600080fd5b8235611407816112ee565b91506020830135611417816112ee565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000604067ffffffffffffffff8516835260208181850152845160a08386015261147e60e0860182611269565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808784030160608801526114b98383611269565b88860151888203830160808a01528051808352908601945060009350908501905b80841015611519578451805173ffffffffffffffffffffffffffffffffffffffff168352860151868301529385019360019390930192908601906114da565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a015295506115558187611269565b9a9950505050505050505050565b60006020828403121561157557600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156115bd576115bd61157c565b500390565b600082198211156115d5576115d561157c565b50019056fea164736f6c634300080f000a",
}

var MetaERC20ABI = MetaERC20MetaData.ABI

var MetaERC20Bin = MetaERC20MetaData.Bin

func DeployMetaERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _totalSupply *big.Int, _ccip_router common.Address, _test_only_force_cross_chain_transfer bool) (common.Address, *types.Transaction, *MetaERC20, error) {
	parsed, err := MetaERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MetaERC20Bin), backend, _totalSupply, _ccip_router, _test_only_force_cross_chain_transfer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MetaERC20{MetaERC20Caller: MetaERC20Caller{contract: contract}, MetaERC20Transactor: MetaERC20Transactor{contract: contract}, MetaERC20Filterer: MetaERC20Filterer{contract: contract}}, nil
}

type MetaERC20 struct {
	address common.Address
	abi     abi.ABI
	MetaERC20Caller
	MetaERC20Transactor
	MetaERC20Filterer
}

type MetaERC20Caller struct {
	contract *bind.BoundContract
}

type MetaERC20Transactor struct {
	contract *bind.BoundContract
}

type MetaERC20Filterer struct {
	contract *bind.BoundContract
}

type MetaERC20Session struct {
	Contract     *MetaERC20
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MetaERC20CallerSession struct {
	Contract *MetaERC20Caller
	CallOpts bind.CallOpts
}

type MetaERC20TransactorSession struct {
	Contract     *MetaERC20Transactor
	TransactOpts bind.TransactOpts
}

type MetaERC20Raw struct {
	Contract *MetaERC20
}

type MetaERC20CallerRaw struct {
	Contract *MetaERC20Caller
}

type MetaERC20TransactorRaw struct {
	Contract *MetaERC20Transactor
}

func NewMetaERC20(address common.Address, backend bind.ContractBackend) (*MetaERC20, error) {
	abi, err := abi.JSON(strings.NewReader(MetaERC20ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMetaERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaERC20{address: address, abi: abi, MetaERC20Caller: MetaERC20Caller{contract: contract}, MetaERC20Transactor: MetaERC20Transactor{contract: contract}, MetaERC20Filterer: MetaERC20Filterer{contract: contract}}, nil
}

func NewMetaERC20Caller(address common.Address, caller bind.ContractCaller) (*MetaERC20Caller, error) {
	contract, err := bindMetaERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaERC20Caller{contract: contract}, nil
}

func NewMetaERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*MetaERC20Transactor, error) {
	contract, err := bindMetaERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaERC20Transactor{contract: contract}, nil
}

func NewMetaERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*MetaERC20Filterer, error) {
	contract, err := bindMetaERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaERC20Filterer{contract: contract}, nil
}

func bindMetaERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_MetaERC20 *MetaERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaERC20.Contract.MetaERC20Caller.contract.Call(opts, result, method, params...)
}

func (_MetaERC20 *MetaERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaERC20.Contract.MetaERC20Transactor.contract.Transfer(opts)
}

func (_MetaERC20 *MetaERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaERC20.Contract.MetaERC20Transactor.contract.Transact(opts, method, params...)
}

func (_MetaERC20 *MetaERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaERC20.Contract.contract.Call(opts, result, method, params...)
}

func (_MetaERC20 *MetaERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaERC20.Contract.contract.Transfer(opts)
}

func (_MetaERC20 *MetaERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaERC20.Contract.contract.Transact(opts, method, params...)
}

func (_MetaERC20 *MetaERC20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MetaERC20.Contract.Allowance(&_MetaERC20.CallOpts, arg0, arg1)
}

func (_MetaERC20 *MetaERC20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MetaERC20.Contract.Allowance(&_MetaERC20.CallOpts, arg0, arg1)
}

func (_MetaERC20 *MetaERC20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MetaERC20.Contract.BalanceOf(&_MetaERC20.CallOpts, arg0)
}

func (_MetaERC20 *MetaERC20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MetaERC20.Contract.BalanceOf(&_MetaERC20.CallOpts, arg0)
}

func (_MetaERC20 *MetaERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) Decimals() (uint8, error) {
	return _MetaERC20.Contract.Decimals(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) Decimals() (uint8, error) {
	return _MetaERC20.Contract.Decimals(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20Caller) GetTrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "getTrustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) GetTrustedForwarder() (common.Address, error) {
	return _MetaERC20.Contract.GetTrustedForwarder(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) GetTrustedForwarder() (common.Address, error) {
	return _MetaERC20.Contract.GetTrustedForwarder(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20Caller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _MetaERC20.Contract.IsTrustedForwarder(&_MetaERC20.CallOpts, forwarder)
}

func (_MetaERC20 *MetaERC20CallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _MetaERC20.Contract.IsTrustedForwarder(&_MetaERC20.CallOpts, forwarder)
}

func (_MetaERC20 *MetaERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) Name() (string, error) {
	return _MetaERC20.Contract.Name(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) Name() (string, error) {
	return _MetaERC20.Contract.Name(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) Owner() (common.Address, error) {
	return _MetaERC20.Contract.Owner(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) Owner() (common.Address, error) {
	return _MetaERC20.Contract.Owner(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) Symbol() (string, error) {
	return _MetaERC20.Contract.Symbol(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) Symbol() (string, error) {
	return _MetaERC20.Contract.Symbol(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) TotalSupply() (*big.Int, error) {
	return _MetaERC20.Contract.TotalSupply(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _MetaERC20.Contract.TotalSupply(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "acceptOwnership")
}

func (_MetaERC20 *MetaERC20Session) AcceptOwnership() (*types.Transaction, error) {
	return _MetaERC20.Contract.AcceptOwnership(&_MetaERC20.TransactOpts)
}

func (_MetaERC20 *MetaERC20TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MetaERC20.Contract.AcceptOwnership(&_MetaERC20.TransactOpts)
}

func (_MetaERC20 *MetaERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "approve", spender, amount)
}

func (_MetaERC20 *MetaERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Approve(&_MetaERC20.TransactOpts, spender, amount)
}

func (_MetaERC20 *MetaERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Approve(&_MetaERC20.TransactOpts, spender, amount)
}

func (_MetaERC20 *MetaERC20Transactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "burn", account, amount)
}

func (_MetaERC20 *MetaERC20Session) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Burn(&_MetaERC20.TransactOpts, account, amount)
}

func (_MetaERC20 *MetaERC20TransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Burn(&_MetaERC20.TransactOpts, account, amount)
}

func (_MetaERC20 *MetaERC20Transactor) Fund(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "fund", spender, amount)
}

func (_MetaERC20 *MetaERC20Session) Fund(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Fund(&_MetaERC20.TransactOpts, spender, amount)
}

func (_MetaERC20 *MetaERC20TransactorSession) Fund(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Fund(&_MetaERC20.TransactOpts, spender, amount)
}

func (_MetaERC20 *MetaERC20Transactor) MetaTransfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, destinationChainId uint64) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "metaTransfer", receiver, amount, destinationChainId)
}

func (_MetaERC20 *MetaERC20Session) MetaTransfer(receiver common.Address, amount *big.Int, destinationChainId uint64) (*types.Transaction, error) {
	return _MetaERC20.Contract.MetaTransfer(&_MetaERC20.TransactOpts, receiver, amount, destinationChainId)
}

func (_MetaERC20 *MetaERC20TransactorSession) MetaTransfer(receiver common.Address, amount *big.Int, destinationChainId uint64) (*types.Transaction, error) {
	return _MetaERC20.Contract.MetaTransfer(&_MetaERC20.TransactOpts, receiver, amount, destinationChainId)
}

func (_MetaERC20 *MetaERC20Transactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "mint", account, amount)
}

func (_MetaERC20 *MetaERC20Session) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Mint(&_MetaERC20.TransactOpts, account, amount)
}

func (_MetaERC20 *MetaERC20TransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Mint(&_MetaERC20.TransactOpts, account, amount)
}

func (_MetaERC20 *MetaERC20Transactor) SetForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "setForwarder", forwarder)
}

func (_MetaERC20 *MetaERC20Session) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _MetaERC20.Contract.SetForwarder(&_MetaERC20.TransactOpts, forwarder)
}

func (_MetaERC20 *MetaERC20TransactorSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _MetaERC20.Contract.SetForwarder(&_MetaERC20.TransactOpts, forwarder)
}

func (_MetaERC20 *MetaERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "transfer", to, amount)
}

func (_MetaERC20 *MetaERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Transfer(&_MetaERC20.TransactOpts, to, amount)
}

func (_MetaERC20 *MetaERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Transfer(&_MetaERC20.TransactOpts, to, amount)
}

func (_MetaERC20 *MetaERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

func (_MetaERC20 *MetaERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.TransferFrom(&_MetaERC20.TransactOpts, from, to, amount)
}

func (_MetaERC20 *MetaERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.TransferFrom(&_MetaERC20.TransactOpts, from, to, amount)
}

func (_MetaERC20 *MetaERC20Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "transferOwnership", to)
}

func (_MetaERC20 *MetaERC20Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MetaERC20.Contract.TransferOwnership(&_MetaERC20.TransactOpts, to)
}

func (_MetaERC20 *MetaERC20TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MetaERC20.Contract.TransferOwnership(&_MetaERC20.TransactOpts, to)
}

func (_MetaERC20 *MetaERC20Transactor) WithdrawNative(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "withdrawNative")
}

func (_MetaERC20 *MetaERC20Session) WithdrawNative() (*types.Transaction, error) {
	return _MetaERC20.Contract.WithdrawNative(&_MetaERC20.TransactOpts)
}

func (_MetaERC20 *MetaERC20TransactorSession) WithdrawNative() (*types.Transaction, error) {
	return _MetaERC20.Contract.WithdrawNative(&_MetaERC20.TransactOpts)
}

func (_MetaERC20 *MetaERC20Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaERC20.contract.RawTransact(opts, nil)
}

func (_MetaERC20 *MetaERC20Session) Receive() (*types.Transaction, error) {
	return _MetaERC20.Contract.Receive(&_MetaERC20.TransactOpts)
}

func (_MetaERC20 *MetaERC20TransactorSession) Receive() (*types.Transaction, error) {
	return _MetaERC20.Contract.Receive(&_MetaERC20.TransactOpts)
}

type MetaERC20ApprovalIterator struct {
	Event *MetaERC20Approval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MetaERC20ApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaERC20Approval)
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
		it.Event = new(MetaERC20Approval)
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

func (it *MetaERC20ApprovalIterator) Error() error {
	return it.fail
}

func (it *MetaERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MetaERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_MetaERC20 *MetaERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MetaERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MetaERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MetaERC20ApprovalIterator{contract: _MetaERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_MetaERC20 *MetaERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MetaERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MetaERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MetaERC20Approval)
				if err := _MetaERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_MetaERC20 *MetaERC20Filterer) ParseApproval(log types.Log) (*MetaERC20Approval, error) {
	event := new(MetaERC20Approval)
	if err := _MetaERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MetaERC20OwnershipTransferRequestedIterator struct {
	Event *MetaERC20OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MetaERC20OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaERC20OwnershipTransferRequested)
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
		it.Event = new(MetaERC20OwnershipTransferRequested)
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

func (it *MetaERC20OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *MetaERC20OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MetaERC20OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MetaERC20 *MetaERC20Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MetaERC20OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetaERC20.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MetaERC20OwnershipTransferRequestedIterator{contract: _MetaERC20.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_MetaERC20 *MetaERC20Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MetaERC20OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetaERC20.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MetaERC20OwnershipTransferRequested)
				if err := _MetaERC20.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_MetaERC20 *MetaERC20Filterer) ParseOwnershipTransferRequested(log types.Log) (*MetaERC20OwnershipTransferRequested, error) {
	event := new(MetaERC20OwnershipTransferRequested)
	if err := _MetaERC20.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MetaERC20OwnershipTransferredIterator struct {
	Event *MetaERC20OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MetaERC20OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaERC20OwnershipTransferred)
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
		it.Event = new(MetaERC20OwnershipTransferred)
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

func (it *MetaERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *MetaERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MetaERC20OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MetaERC20 *MetaERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MetaERC20OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetaERC20.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MetaERC20OwnershipTransferredIterator{contract: _MetaERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_MetaERC20 *MetaERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MetaERC20OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetaERC20.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MetaERC20OwnershipTransferred)
				if err := _MetaERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_MetaERC20 *MetaERC20Filterer) ParseOwnershipTransferred(log types.Log) (*MetaERC20OwnershipTransferred, error) {
	event := new(MetaERC20OwnershipTransferred)
	if err := _MetaERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MetaERC20TransferIterator struct {
	Event *MetaERC20Transfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MetaERC20TransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaERC20Transfer)
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
		it.Event = new(MetaERC20Transfer)
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

func (it *MetaERC20TransferIterator) Error() error {
	return it.fail
}

func (it *MetaERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MetaERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_MetaERC20 *MetaERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MetaERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetaERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MetaERC20TransferIterator{contract: _MetaERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_MetaERC20 *MetaERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MetaERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetaERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MetaERC20Transfer)
				if err := _MetaERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_MetaERC20 *MetaERC20Filterer) ParseTransfer(log types.Log) (*MetaERC20Transfer, error) {
	event := new(MetaERC20Transfer)
	if err := _MetaERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_MetaERC20 *MetaERC20) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MetaERC20.abi.Events["Approval"].ID:
		return _MetaERC20.ParseApproval(log)
	case _MetaERC20.abi.Events["OwnershipTransferRequested"].ID:
		return _MetaERC20.ParseOwnershipTransferRequested(log)
	case _MetaERC20.abi.Events["OwnershipTransferred"].ID:
		return _MetaERC20.ParseOwnershipTransferred(log)
	case _MetaERC20.abi.Events["Transfer"].ID:
		return _MetaERC20.ParseTransfer(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MetaERC20Approval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (MetaERC20OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (MetaERC20OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (MetaERC20Transfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (_MetaERC20 *MetaERC20) Address() common.Address {
	return _MetaERC20.address
}

type MetaERC20Interface interface {
	Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	GetTrustedForwarder(opts *bind.CallOpts) (common.Address, error)

	IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error)

	Name(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)

	Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	Fund(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)

	MetaTransfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, destinationChainId uint64) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	SetForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawNative(opts *bind.TransactOpts) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MetaERC20ApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *MetaERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*MetaERC20Approval, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MetaERC20OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MetaERC20OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*MetaERC20OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MetaERC20OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MetaERC20OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*MetaERC20OwnershipTransferred, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MetaERC20TransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *MetaERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*MetaERC20Transfer, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
