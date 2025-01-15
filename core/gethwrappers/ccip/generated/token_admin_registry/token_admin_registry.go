package token_admin_registry

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

type TokenAdminRegistryTokenConfig struct {
	Administrator        common.Address
	PendingAdministrator common.Address
	TokenPool            common.Address
}

var TokenAdminRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InvalidTokenPoolToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"OnlyAdministrator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"OnlyPendingAdministrator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OnlyRegistryModuleOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currentAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdministratorTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdministratorTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousPool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"PoolSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"RegistryModuleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"RegistryModuleRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"addRegistryModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"startIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxCount\",\"type\":\"uint64\"}],\"name\":\"getAllConfiguredTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendingAdministrator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenPool\",\"type\":\"address\"}],\"internalType\":\"structTokenAdminRegistry.TokenConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"}],\"name\":\"isAdministrator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"isRegistryModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"}],\"name\":\"proposeAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"removeRegistryModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"setPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b611449806101576000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80637d3f255211610097578063cb67e3b111610066578063cb67e3b1146102bc578063ddadfa8e14610374578063e677ae3714610387578063f2fde38b1461039a57600080fd5b80637d3f2552146101e05780638da5cb5b14610203578063bbe4f6db14610242578063c1af6e031461027f57600080fd5b80634e847fc7116100d35780634e847fc7146101925780635e63547a146101a557806372d64a81146101c557806379ba5097146101d857600080fd5b806310cbcf1814610105578063156194da1461011a578063181f5a771461012d5780633dc457721461017f575b600080fd5b61011861011336600461116c565b6103ad565b005b61011861012836600461116c565b61040a565b6101696040518060400160405280601881526020017f546f6b656e41646d696e526567697374727920312e352e30000000000000000081525081565b6040516101769190611187565b60405180910390f35b61011861018d36600461116c565b61050f565b6101186101a03660046111f4565b610573565b6101b86101b3366004611227565b6107d3565b604051610176919061129c565b6101b86101d336600461130e565b6108cc565b6101186109e2565b6101f36101ee36600461116c565b610adf565b6040519015158152602001610176565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610176565b61021d61025036600461116c565b73ffffffffffffffffffffffffffffffffffffffff908116600090815260026020819052604090912001541690565b6101f361028d3660046111f4565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260026020526040902054821691161490565b6103356102ca36600461116c565b60408051606080820183526000808352602080840182905292840181905273ffffffffffffffffffffffffffffffffffffffff948516815260028084529084902084519283018552805486168352600181015486169383019390935291909101549092169082015290565b60408051825173ffffffffffffffffffffffffffffffffffffffff90811682526020808501518216908301529282015190921690820152606001610176565b6101186103823660046111f4565b610aec565b6101186103953660046111f4565b610bf6565b6101186103a836600461116c565b610dbe565b6103b5610dcf565b6103c0600582610e52565b156104075760405173ffffffffffffffffffffffffffffffffffffffff8216907f93eaa26dcb9275e56bacb1d33fdbf402262da6f0f4baf2a6e2cd154b73f387f890600090a25b50565b73ffffffffffffffffffffffffffffffffffffffff808216600090815260026020526040902060018101549091163314610493576040517f3edffe7500000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff831660248201526044015b60405180910390fd5b8054337fffffffffffffffffffffffff00000000000000000000000000000000000000009182168117835560018301805490921690915560405173ffffffffffffffffffffffffffffffffffffffff8416907f399b55200f7f639a63d76efe3dcfa9156ce367058d6b673041b84a628885f5a790600090a35050565b610517610dcf565b610522600582610e7b565b156104075760405173ffffffffffffffffffffffffffffffffffffffff821681527f3cabf004338366bfeaeb610ad827cb58d16b588017c509501f2c97c83caae7b29060200160405180910390a150565b73ffffffffffffffffffffffffffffffffffffffff80831660009081526002602052604090205483911633146105f3576040517fed5d85b500000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff8216602482015260440161048a565b73ffffffffffffffffffffffffffffffffffffffff8216158015906106a557506040517f240028e800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015283169063240028e890602401602060405180830381865afa15801561067f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a39190611338565b155b156106f4576040517f962b60e600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260240161048a565b73ffffffffffffffffffffffffffffffffffffffff808416600090815260026020819052604090912090810180548584167fffffffffffffffffffffffff0000000000000000000000000000000000000000821681179092559192919091169081146107cc578373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f754449ec3aff3bd528bfce43ae9319c4a381b67fcd1d20097b3b24dacaecc35d60405160405180910390a45b5050505050565b606060008267ffffffffffffffff8111156107f0576107f061135a565b604051908082528060200260200182016040528015610819578160200160208202803683370190505b50905060005b838110156108c2576002600086868481811061083d5761083d611389565b9050602002016020810190610852919061116c565b73ffffffffffffffffffffffffffffffffffffffff9081168252602082019290925260400160002060020154835191169083908390811061089557610895611389565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015260010161081f565b5090505b92915050565b606060006108da6003610e9d565b9050808467ffffffffffffffff16106108f357506108c6565b67ffffffffffffffff80841690829061090e908716836113e7565b111561092b5761092867ffffffffffffffff8616836113fa565b90505b8067ffffffffffffffff8111156109445761094461135a565b60405190808252806020026020018201604052801561096d578160200160208202803683370190505b50925060005b818110156109d95761099a6109928267ffffffffffffffff89166113e7565b600390610ea7565b8482815181106109ac576109ac611389565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152600101610973565b50505092915050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610a63576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161048a565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60006108c6600583610eb3565b73ffffffffffffffffffffffffffffffffffffffff8083166000908152600260205260409020548391163314610b6c576040517fed5d85b500000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff8216602482015260440161048a565b73ffffffffffffffffffffffffffffffffffffffff8381166000818152600260205260408082206001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001695881695861790559051909392339290917fc54c3051ff16e63bb9203214432372aca006c589e3653619b577a3265675b7169190a450505050565b610bff33610adf565b158015610c24575060005473ffffffffffffffffffffffffffffffffffffffff163314155b15610c5d576040517f51ca1ec300000000000000000000000000000000000000000000000000000000815233600482015260240161048a565b73ffffffffffffffffffffffffffffffffffffffff8116610caa576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8083166000908152600260205260409020805490911615610d24576040517f45ed80e900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260240161048a565b6001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8416179055610d71600384610e7b565b5060405173ffffffffffffffffffffffffffffffffffffffff808416916000918616907fc54c3051ff16e63bb9203214432372aca006c589e3653619b577a3265675b716908390a4505050565b610dc6610dcf565b61040781610ee2565b60005473ffffffffffffffffffffffffffffffffffffffff163314610e50576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161048a565b565b6000610e748373ffffffffffffffffffffffffffffffffffffffff8416610fd7565b9392505050565b6000610e748373ffffffffffffffffffffffffffffffffffffffff84166110ca565b60006108c6825490565b6000610e748383611119565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515610e74565b3373ffffffffffffffffffffffffffffffffffffffff821603610f61576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161048a565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600081815260018301602052604081205480156110c0576000610ffb6001836113fa565b855490915060009061100f906001906113fa565b905080821461107457600086600001828154811061102f5761102f611389565b906000526020600020015490508087600001848154811061105257611052611389565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806110855761108561140d565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506108c6565b60009150506108c6565b6000818152600183016020526040812054611111575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556108c6565b5060006108c6565b600082600001828154811061113057611130611389565b9060005260206000200154905092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461116757600080fd5b919050565b60006020828403121561117e57600080fd5b610e7482611143565b60006020808352835180602085015260005b818110156111b557858101830151858201604001528201611199565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b6000806040838503121561120757600080fd5b61121083611143565b915061121e60208401611143565b90509250929050565b6000806020838503121561123a57600080fd5b823567ffffffffffffffff8082111561125257600080fd5b818501915085601f83011261126657600080fd5b81358181111561127557600080fd5b8660208260051b850101111561128a57600080fd5b60209290920196919550909350505050565b6020808252825182820181905260009190848201906040850190845b818110156112ea57835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016112b8565b50909695505050505050565b803567ffffffffffffffff8116811461116757600080fd5b6000806040838503121561132157600080fd5b61132a836112f6565b915061121e602084016112f6565b60006020828403121561134a57600080fd5b81518015158114610e7457600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156108c6576108c66113b8565b818103818111156108c6576108c66113b8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000818000a",
}

var TokenAdminRegistryABI = TokenAdminRegistryMetaData.ABI

var TokenAdminRegistryBin = TokenAdminRegistryMetaData.Bin

func DeployTokenAdminRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *generated.Transaction, *TokenAdminRegistry, error) {
	parsed, err := TokenAdminRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}
	if generated.IsZKSync(backend) {
		address, ethTx, contractBind, _ := generated.DeployContract(auth, parsed, common.FromHex(TokenAdminRegistryZKBin), backend)
		contractReturn := &TokenAdminRegistry{address: address, abi: *parsed, TokenAdminRegistryCaller: TokenAdminRegistryCaller{contract: contractBind}, TokenAdminRegistryTransactor: TokenAdminRegistryTransactor{contract: contractBind}, TokenAdminRegistryFilterer: TokenAdminRegistryFilterer{contract: contractBind}}
		return address, ethTx, contractReturn, err
	}
	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenAdminRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, &generated.Transaction{Transaction: tx, HashZks: tx.Hash()}, &TokenAdminRegistry{address: address, abi: *parsed, TokenAdminRegistryCaller: TokenAdminRegistryCaller{contract: contract}, TokenAdminRegistryTransactor: TokenAdminRegistryTransactor{contract: contract}, TokenAdminRegistryFilterer: TokenAdminRegistryFilterer{contract: contract}}, nil
}

type TokenAdminRegistry struct {
	address common.Address
	abi     abi.ABI
	TokenAdminRegistryCaller
	TokenAdminRegistryTransactor
	TokenAdminRegistryFilterer
}

type TokenAdminRegistryCaller struct {
	contract *bind.BoundContract
}

type TokenAdminRegistryTransactor struct {
	contract *bind.BoundContract
}

type TokenAdminRegistryFilterer struct {
	contract *bind.BoundContract
}

type TokenAdminRegistrySession struct {
	Contract     *TokenAdminRegistry
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TokenAdminRegistryCallerSession struct {
	Contract *TokenAdminRegistryCaller
	CallOpts bind.CallOpts
}

type TokenAdminRegistryTransactorSession struct {
	Contract     *TokenAdminRegistryTransactor
	TransactOpts bind.TransactOpts
}

type TokenAdminRegistryRaw struct {
	Contract *TokenAdminRegistry
}

type TokenAdminRegistryCallerRaw struct {
	Contract *TokenAdminRegistryCaller
}

type TokenAdminRegistryTransactorRaw struct {
	Contract *TokenAdminRegistryTransactor
}

func NewTokenAdminRegistry(address common.Address, backend bind.ContractBackend) (*TokenAdminRegistry, error) {
	abi, err := abi.JSON(strings.NewReader(TokenAdminRegistryABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindTokenAdminRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenAdminRegistry{address: address, abi: abi, TokenAdminRegistryCaller: TokenAdminRegistryCaller{contract: contract}, TokenAdminRegistryTransactor: TokenAdminRegistryTransactor{contract: contract}, TokenAdminRegistryFilterer: TokenAdminRegistryFilterer{contract: contract}}, nil
}

func NewTokenAdminRegistryCaller(address common.Address, caller bind.ContractCaller) (*TokenAdminRegistryCaller, error) {
	contract, err := bindTokenAdminRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenAdminRegistryCaller{contract: contract}, nil
}

func NewTokenAdminRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenAdminRegistryTransactor, error) {
	contract, err := bindTokenAdminRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenAdminRegistryTransactor{contract: contract}, nil
}

func NewTokenAdminRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenAdminRegistryFilterer, error) {
	contract, err := bindTokenAdminRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenAdminRegistryFilterer{contract: contract}, nil
}

func bindTokenAdminRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenAdminRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_TokenAdminRegistry *TokenAdminRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenAdminRegistry.Contract.TokenAdminRegistryCaller.contract.Call(opts, result, method, params...)
}

func (_TokenAdminRegistry *TokenAdminRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.TokenAdminRegistryTransactor.contract.Transfer(opts)
}

func (_TokenAdminRegistry *TokenAdminRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.TokenAdminRegistryTransactor.contract.Transact(opts, method, params...)
}

func (_TokenAdminRegistry *TokenAdminRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenAdminRegistry.Contract.contract.Call(opts, result, method, params...)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.contract.Transfer(opts)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.contract.Transact(opts, method, params...)
}

func (_TokenAdminRegistry *TokenAdminRegistryCaller) GetAllConfiguredTokens(opts *bind.CallOpts, startIndex uint64, maxCount uint64) ([]common.Address, error) {
	var out []interface{}
	err := _TokenAdminRegistry.contract.Call(opts, &out, "getAllConfiguredTokens", startIndex, maxCount)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenAdminRegistry *TokenAdminRegistrySession) GetAllConfiguredTokens(startIndex uint64, maxCount uint64) ([]common.Address, error) {
	return _TokenAdminRegistry.Contract.GetAllConfiguredTokens(&_TokenAdminRegistry.CallOpts, startIndex, maxCount)
}

func (_TokenAdminRegistry *TokenAdminRegistryCallerSession) GetAllConfiguredTokens(startIndex uint64, maxCount uint64) ([]common.Address, error) {
	return _TokenAdminRegistry.Contract.GetAllConfiguredTokens(&_TokenAdminRegistry.CallOpts, startIndex, maxCount)
}

func (_TokenAdminRegistry *TokenAdminRegistryCaller) GetPool(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _TokenAdminRegistry.contract.Call(opts, &out, "getPool", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TokenAdminRegistry *TokenAdminRegistrySession) GetPool(token common.Address) (common.Address, error) {
	return _TokenAdminRegistry.Contract.GetPool(&_TokenAdminRegistry.CallOpts, token)
}

func (_TokenAdminRegistry *TokenAdminRegistryCallerSession) GetPool(token common.Address) (common.Address, error) {
	return _TokenAdminRegistry.Contract.GetPool(&_TokenAdminRegistry.CallOpts, token)
}

func (_TokenAdminRegistry *TokenAdminRegistryCaller) GetPools(opts *bind.CallOpts, tokens []common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _TokenAdminRegistry.contract.Call(opts, &out, "getPools", tokens)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenAdminRegistry *TokenAdminRegistrySession) GetPools(tokens []common.Address) ([]common.Address, error) {
	return _TokenAdminRegistry.Contract.GetPools(&_TokenAdminRegistry.CallOpts, tokens)
}

func (_TokenAdminRegistry *TokenAdminRegistryCallerSession) GetPools(tokens []common.Address) ([]common.Address, error) {
	return _TokenAdminRegistry.Contract.GetPools(&_TokenAdminRegistry.CallOpts, tokens)
}

func (_TokenAdminRegistry *TokenAdminRegistryCaller) GetTokenConfig(opts *bind.CallOpts, token common.Address) (TokenAdminRegistryTokenConfig, error) {
	var out []interface{}
	err := _TokenAdminRegistry.contract.Call(opts, &out, "getTokenConfig", token)

	if err != nil {
		return *new(TokenAdminRegistryTokenConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenAdminRegistryTokenConfig)).(*TokenAdminRegistryTokenConfig)

	return out0, err

}

func (_TokenAdminRegistry *TokenAdminRegistrySession) GetTokenConfig(token common.Address) (TokenAdminRegistryTokenConfig, error) {
	return _TokenAdminRegistry.Contract.GetTokenConfig(&_TokenAdminRegistry.CallOpts, token)
}

func (_TokenAdminRegistry *TokenAdminRegistryCallerSession) GetTokenConfig(token common.Address) (TokenAdminRegistryTokenConfig, error) {
	return _TokenAdminRegistry.Contract.GetTokenConfig(&_TokenAdminRegistry.CallOpts, token)
}

func (_TokenAdminRegistry *TokenAdminRegistryCaller) IsAdministrator(opts *bind.CallOpts, localToken common.Address, administrator common.Address) (bool, error) {
	var out []interface{}
	err := _TokenAdminRegistry.contract.Call(opts, &out, "isAdministrator", localToken, administrator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_TokenAdminRegistry *TokenAdminRegistrySession) IsAdministrator(localToken common.Address, administrator common.Address) (bool, error) {
	return _TokenAdminRegistry.Contract.IsAdministrator(&_TokenAdminRegistry.CallOpts, localToken, administrator)
}

func (_TokenAdminRegistry *TokenAdminRegistryCallerSession) IsAdministrator(localToken common.Address, administrator common.Address) (bool, error) {
	return _TokenAdminRegistry.Contract.IsAdministrator(&_TokenAdminRegistry.CallOpts, localToken, administrator)
}

func (_TokenAdminRegistry *TokenAdminRegistryCaller) IsRegistryModule(opts *bind.CallOpts, module common.Address) (bool, error) {
	var out []interface{}
	err := _TokenAdminRegistry.contract.Call(opts, &out, "isRegistryModule", module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_TokenAdminRegistry *TokenAdminRegistrySession) IsRegistryModule(module common.Address) (bool, error) {
	return _TokenAdminRegistry.Contract.IsRegistryModule(&_TokenAdminRegistry.CallOpts, module)
}

func (_TokenAdminRegistry *TokenAdminRegistryCallerSession) IsRegistryModule(module common.Address) (bool, error) {
	return _TokenAdminRegistry.Contract.IsRegistryModule(&_TokenAdminRegistry.CallOpts, module)
}

func (_TokenAdminRegistry *TokenAdminRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenAdminRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TokenAdminRegistry *TokenAdminRegistrySession) Owner() (common.Address, error) {
	return _TokenAdminRegistry.Contract.Owner(&_TokenAdminRegistry.CallOpts)
}

func (_TokenAdminRegistry *TokenAdminRegistryCallerSession) Owner() (common.Address, error) {
	return _TokenAdminRegistry.Contract.Owner(&_TokenAdminRegistry.CallOpts)
}

func (_TokenAdminRegistry *TokenAdminRegistryCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenAdminRegistry.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_TokenAdminRegistry *TokenAdminRegistrySession) TypeAndVersion() (string, error) {
	return _TokenAdminRegistry.Contract.TypeAndVersion(&_TokenAdminRegistry.CallOpts)
}

func (_TokenAdminRegistry *TokenAdminRegistryCallerSession) TypeAndVersion() (string, error) {
	return _TokenAdminRegistry.Contract.TypeAndVersion(&_TokenAdminRegistry.CallOpts)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactor) AcceptAdminRole(opts *bind.TransactOpts, localToken common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.contract.Transact(opts, "acceptAdminRole", localToken)
}

func (_TokenAdminRegistry *TokenAdminRegistrySession) AcceptAdminRole(localToken common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.AcceptAdminRole(&_TokenAdminRegistry.TransactOpts, localToken)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactorSession) AcceptAdminRole(localToken common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.AcceptAdminRole(&_TokenAdminRegistry.TransactOpts, localToken)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenAdminRegistry.contract.Transact(opts, "acceptOwnership")
}

func (_TokenAdminRegistry *TokenAdminRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.AcceptOwnership(&_TokenAdminRegistry.TransactOpts)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.AcceptOwnership(&_TokenAdminRegistry.TransactOpts)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactor) AddRegistryModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.contract.Transact(opts, "addRegistryModule", module)
}

func (_TokenAdminRegistry *TokenAdminRegistrySession) AddRegistryModule(module common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.AddRegistryModule(&_TokenAdminRegistry.TransactOpts, module)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactorSession) AddRegistryModule(module common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.AddRegistryModule(&_TokenAdminRegistry.TransactOpts, module)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactor) ProposeAdministrator(opts *bind.TransactOpts, localToken common.Address, administrator common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.contract.Transact(opts, "proposeAdministrator", localToken, administrator)
}

func (_TokenAdminRegistry *TokenAdminRegistrySession) ProposeAdministrator(localToken common.Address, administrator common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.ProposeAdministrator(&_TokenAdminRegistry.TransactOpts, localToken, administrator)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactorSession) ProposeAdministrator(localToken common.Address, administrator common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.ProposeAdministrator(&_TokenAdminRegistry.TransactOpts, localToken, administrator)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactor) RemoveRegistryModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.contract.Transact(opts, "removeRegistryModule", module)
}

func (_TokenAdminRegistry *TokenAdminRegistrySession) RemoveRegistryModule(module common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.RemoveRegistryModule(&_TokenAdminRegistry.TransactOpts, module)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactorSession) RemoveRegistryModule(module common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.RemoveRegistryModule(&_TokenAdminRegistry.TransactOpts, module)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactor) SetPool(opts *bind.TransactOpts, localToken common.Address, pool common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.contract.Transact(opts, "setPool", localToken, pool)
}

func (_TokenAdminRegistry *TokenAdminRegistrySession) SetPool(localToken common.Address, pool common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.SetPool(&_TokenAdminRegistry.TransactOpts, localToken, pool)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactorSession) SetPool(localToken common.Address, pool common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.SetPool(&_TokenAdminRegistry.TransactOpts, localToken, pool)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactor) TransferAdminRole(opts *bind.TransactOpts, localToken common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.contract.Transact(opts, "transferAdminRole", localToken, newAdmin)
}

func (_TokenAdminRegistry *TokenAdminRegistrySession) TransferAdminRole(localToken common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.TransferAdminRole(&_TokenAdminRegistry.TransactOpts, localToken, newAdmin)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactorSession) TransferAdminRole(localToken common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.TransferAdminRole(&_TokenAdminRegistry.TransactOpts, localToken, newAdmin)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.contract.Transact(opts, "transferOwnership", to)
}

func (_TokenAdminRegistry *TokenAdminRegistrySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.TransferOwnership(&_TokenAdminRegistry.TransactOpts, to)
}

func (_TokenAdminRegistry *TokenAdminRegistryTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _TokenAdminRegistry.Contract.TransferOwnership(&_TokenAdminRegistry.TransactOpts, to)
}

type TokenAdminRegistryAdministratorTransferRequestedIterator struct {
	Event *TokenAdminRegistryAdministratorTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenAdminRegistryAdministratorTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAdminRegistryAdministratorTransferRequested)
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
		it.Event = new(TokenAdminRegistryAdministratorTransferRequested)
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

func (it *TokenAdminRegistryAdministratorTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *TokenAdminRegistryAdministratorTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenAdminRegistryAdministratorTransferRequested struct {
	Token        common.Address
	CurrentAdmin common.Address
	NewAdmin     common.Address
	Raw          types.Log
}

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) FilterAdministratorTransferRequested(opts *bind.FilterOpts, token []common.Address, currentAdmin []common.Address, newAdmin []common.Address) (*TokenAdminRegistryAdministratorTransferRequestedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var currentAdminRule []interface{}
	for _, currentAdminItem := range currentAdmin {
		currentAdminRule = append(currentAdminRule, currentAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _TokenAdminRegistry.contract.FilterLogs(opts, "AdministratorTransferRequested", tokenRule, currentAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &TokenAdminRegistryAdministratorTransferRequestedIterator{contract: _TokenAdminRegistry.contract, event: "AdministratorTransferRequested", logs: logs, sub: sub}, nil
}

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) WatchAdministratorTransferRequested(opts *bind.WatchOpts, sink chan<- *TokenAdminRegistryAdministratorTransferRequested, token []common.Address, currentAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var currentAdminRule []interface{}
	for _, currentAdminItem := range currentAdmin {
		currentAdminRule = append(currentAdminRule, currentAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _TokenAdminRegistry.contract.WatchLogs(opts, "AdministratorTransferRequested", tokenRule, currentAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenAdminRegistryAdministratorTransferRequested)
				if err := _TokenAdminRegistry.contract.UnpackLog(event, "AdministratorTransferRequested", log); err != nil {
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

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) ParseAdministratorTransferRequested(log types.Log) (*TokenAdminRegistryAdministratorTransferRequested, error) {
	event := new(TokenAdminRegistryAdministratorTransferRequested)
	if err := _TokenAdminRegistry.contract.UnpackLog(event, "AdministratorTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenAdminRegistryAdministratorTransferredIterator struct {
	Event *TokenAdminRegistryAdministratorTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenAdminRegistryAdministratorTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAdminRegistryAdministratorTransferred)
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
		it.Event = new(TokenAdminRegistryAdministratorTransferred)
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

func (it *TokenAdminRegistryAdministratorTransferredIterator) Error() error {
	return it.fail
}

func (it *TokenAdminRegistryAdministratorTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenAdminRegistryAdministratorTransferred struct {
	Token    common.Address
	NewAdmin common.Address
	Raw      types.Log
}

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) FilterAdministratorTransferred(opts *bind.FilterOpts, token []common.Address, newAdmin []common.Address) (*TokenAdminRegistryAdministratorTransferredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _TokenAdminRegistry.contract.FilterLogs(opts, "AdministratorTransferred", tokenRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &TokenAdminRegistryAdministratorTransferredIterator{contract: _TokenAdminRegistry.contract, event: "AdministratorTransferred", logs: logs, sub: sub}, nil
}

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) WatchAdministratorTransferred(opts *bind.WatchOpts, sink chan<- *TokenAdminRegistryAdministratorTransferred, token []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _TokenAdminRegistry.contract.WatchLogs(opts, "AdministratorTransferred", tokenRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenAdminRegistryAdministratorTransferred)
				if err := _TokenAdminRegistry.contract.UnpackLog(event, "AdministratorTransferred", log); err != nil {
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

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) ParseAdministratorTransferred(log types.Log) (*TokenAdminRegistryAdministratorTransferred, error) {
	event := new(TokenAdminRegistryAdministratorTransferred)
	if err := _TokenAdminRegistry.contract.UnpackLog(event, "AdministratorTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenAdminRegistryOwnershipTransferRequestedIterator struct {
	Event *TokenAdminRegistryOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenAdminRegistryOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAdminRegistryOwnershipTransferRequested)
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
		it.Event = new(TokenAdminRegistryOwnershipTransferRequested)
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

func (it *TokenAdminRegistryOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *TokenAdminRegistryOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenAdminRegistryOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenAdminRegistryOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenAdminRegistry.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenAdminRegistryOwnershipTransferRequestedIterator{contract: _TokenAdminRegistry.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *TokenAdminRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenAdminRegistry.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenAdminRegistryOwnershipTransferRequested)
				if err := _TokenAdminRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) ParseOwnershipTransferRequested(log types.Log) (*TokenAdminRegistryOwnershipTransferRequested, error) {
	event := new(TokenAdminRegistryOwnershipTransferRequested)
	if err := _TokenAdminRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenAdminRegistryOwnershipTransferredIterator struct {
	Event *TokenAdminRegistryOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenAdminRegistryOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAdminRegistryOwnershipTransferred)
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
		it.Event = new(TokenAdminRegistryOwnershipTransferred)
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

func (it *TokenAdminRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *TokenAdminRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenAdminRegistryOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenAdminRegistryOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenAdminRegistry.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenAdminRegistryOwnershipTransferredIterator{contract: _TokenAdminRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenAdminRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenAdminRegistry.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenAdminRegistryOwnershipTransferred)
				if err := _TokenAdminRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*TokenAdminRegistryOwnershipTransferred, error) {
	event := new(TokenAdminRegistryOwnershipTransferred)
	if err := _TokenAdminRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenAdminRegistryPoolSetIterator struct {
	Event *TokenAdminRegistryPoolSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenAdminRegistryPoolSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAdminRegistryPoolSet)
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
		it.Event = new(TokenAdminRegistryPoolSet)
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

func (it *TokenAdminRegistryPoolSetIterator) Error() error {
	return it.fail
}

func (it *TokenAdminRegistryPoolSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenAdminRegistryPoolSet struct {
	Token        common.Address
	PreviousPool common.Address
	NewPool      common.Address
	Raw          types.Log
}

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) FilterPoolSet(opts *bind.FilterOpts, token []common.Address, previousPool []common.Address, newPool []common.Address) (*TokenAdminRegistryPoolSetIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var previousPoolRule []interface{}
	for _, previousPoolItem := range previousPool {
		previousPoolRule = append(previousPoolRule, previousPoolItem)
	}
	var newPoolRule []interface{}
	for _, newPoolItem := range newPool {
		newPoolRule = append(newPoolRule, newPoolItem)
	}

	logs, sub, err := _TokenAdminRegistry.contract.FilterLogs(opts, "PoolSet", tokenRule, previousPoolRule, newPoolRule)
	if err != nil {
		return nil, err
	}
	return &TokenAdminRegistryPoolSetIterator{contract: _TokenAdminRegistry.contract, event: "PoolSet", logs: logs, sub: sub}, nil
}

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) WatchPoolSet(opts *bind.WatchOpts, sink chan<- *TokenAdminRegistryPoolSet, token []common.Address, previousPool []common.Address, newPool []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var previousPoolRule []interface{}
	for _, previousPoolItem := range previousPool {
		previousPoolRule = append(previousPoolRule, previousPoolItem)
	}
	var newPoolRule []interface{}
	for _, newPoolItem := range newPool {
		newPoolRule = append(newPoolRule, newPoolItem)
	}

	logs, sub, err := _TokenAdminRegistry.contract.WatchLogs(opts, "PoolSet", tokenRule, previousPoolRule, newPoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenAdminRegistryPoolSet)
				if err := _TokenAdminRegistry.contract.UnpackLog(event, "PoolSet", log); err != nil {
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

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) ParsePoolSet(log types.Log) (*TokenAdminRegistryPoolSet, error) {
	event := new(TokenAdminRegistryPoolSet)
	if err := _TokenAdminRegistry.contract.UnpackLog(event, "PoolSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenAdminRegistryRegistryModuleAddedIterator struct {
	Event *TokenAdminRegistryRegistryModuleAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenAdminRegistryRegistryModuleAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAdminRegistryRegistryModuleAdded)
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
		it.Event = new(TokenAdminRegistryRegistryModuleAdded)
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

func (it *TokenAdminRegistryRegistryModuleAddedIterator) Error() error {
	return it.fail
}

func (it *TokenAdminRegistryRegistryModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenAdminRegistryRegistryModuleAdded struct {
	Module common.Address
	Raw    types.Log
}

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) FilterRegistryModuleAdded(opts *bind.FilterOpts) (*TokenAdminRegistryRegistryModuleAddedIterator, error) {

	logs, sub, err := _TokenAdminRegistry.contract.FilterLogs(opts, "RegistryModuleAdded")
	if err != nil {
		return nil, err
	}
	return &TokenAdminRegistryRegistryModuleAddedIterator{contract: _TokenAdminRegistry.contract, event: "RegistryModuleAdded", logs: logs, sub: sub}, nil
}

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) WatchRegistryModuleAdded(opts *bind.WatchOpts, sink chan<- *TokenAdminRegistryRegistryModuleAdded) (event.Subscription, error) {

	logs, sub, err := _TokenAdminRegistry.contract.WatchLogs(opts, "RegistryModuleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenAdminRegistryRegistryModuleAdded)
				if err := _TokenAdminRegistry.contract.UnpackLog(event, "RegistryModuleAdded", log); err != nil {
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

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) ParseRegistryModuleAdded(log types.Log) (*TokenAdminRegistryRegistryModuleAdded, error) {
	event := new(TokenAdminRegistryRegistryModuleAdded)
	if err := _TokenAdminRegistry.contract.UnpackLog(event, "RegistryModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenAdminRegistryRegistryModuleRemovedIterator struct {
	Event *TokenAdminRegistryRegistryModuleRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenAdminRegistryRegistryModuleRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAdminRegistryRegistryModuleRemoved)
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
		it.Event = new(TokenAdminRegistryRegistryModuleRemoved)
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

func (it *TokenAdminRegistryRegistryModuleRemovedIterator) Error() error {
	return it.fail
}

func (it *TokenAdminRegistryRegistryModuleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenAdminRegistryRegistryModuleRemoved struct {
	Module common.Address
	Raw    types.Log
}

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) FilterRegistryModuleRemoved(opts *bind.FilterOpts, module []common.Address) (*TokenAdminRegistryRegistryModuleRemovedIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _TokenAdminRegistry.contract.FilterLogs(opts, "RegistryModuleRemoved", moduleRule)
	if err != nil {
		return nil, err
	}
	return &TokenAdminRegistryRegistryModuleRemovedIterator{contract: _TokenAdminRegistry.contract, event: "RegistryModuleRemoved", logs: logs, sub: sub}, nil
}

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) WatchRegistryModuleRemoved(opts *bind.WatchOpts, sink chan<- *TokenAdminRegistryRegistryModuleRemoved, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _TokenAdminRegistry.contract.WatchLogs(opts, "RegistryModuleRemoved", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenAdminRegistryRegistryModuleRemoved)
				if err := _TokenAdminRegistry.contract.UnpackLog(event, "RegistryModuleRemoved", log); err != nil {
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

func (_TokenAdminRegistry *TokenAdminRegistryFilterer) ParseRegistryModuleRemoved(log types.Log) (*TokenAdminRegistryRegistryModuleRemoved, error) {
	event := new(TokenAdminRegistryRegistryModuleRemoved)
	if err := _TokenAdminRegistry.contract.UnpackLog(event, "RegistryModuleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_TokenAdminRegistry *TokenAdminRegistry) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _TokenAdminRegistry.abi.Events["AdministratorTransferRequested"].ID:
		return _TokenAdminRegistry.ParseAdministratorTransferRequested(log)
	case _TokenAdminRegistry.abi.Events["AdministratorTransferred"].ID:
		return _TokenAdminRegistry.ParseAdministratorTransferred(log)
	case _TokenAdminRegistry.abi.Events["OwnershipTransferRequested"].ID:
		return _TokenAdminRegistry.ParseOwnershipTransferRequested(log)
	case _TokenAdminRegistry.abi.Events["OwnershipTransferred"].ID:
		return _TokenAdminRegistry.ParseOwnershipTransferred(log)
	case _TokenAdminRegistry.abi.Events["PoolSet"].ID:
		return _TokenAdminRegistry.ParsePoolSet(log)
	case _TokenAdminRegistry.abi.Events["RegistryModuleAdded"].ID:
		return _TokenAdminRegistry.ParseRegistryModuleAdded(log)
	case _TokenAdminRegistry.abi.Events["RegistryModuleRemoved"].ID:
		return _TokenAdminRegistry.ParseRegistryModuleRemoved(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (TokenAdminRegistryAdministratorTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xc54c3051ff16e63bb9203214432372aca006c589e3653619b577a3265675b716")
}

func (TokenAdminRegistryAdministratorTransferred) Topic() common.Hash {
	return common.HexToHash("0x399b55200f7f639a63d76efe3dcfa9156ce367058d6b673041b84a628885f5a7")
}

func (TokenAdminRegistryOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (TokenAdminRegistryOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (TokenAdminRegistryPoolSet) Topic() common.Hash {
	return common.HexToHash("0x754449ec3aff3bd528bfce43ae9319c4a381b67fcd1d20097b3b24dacaecc35d")
}

func (TokenAdminRegistryRegistryModuleAdded) Topic() common.Hash {
	return common.HexToHash("0x3cabf004338366bfeaeb610ad827cb58d16b588017c509501f2c97c83caae7b2")
}

func (TokenAdminRegistryRegistryModuleRemoved) Topic() common.Hash {
	return common.HexToHash("0x93eaa26dcb9275e56bacb1d33fdbf402262da6f0f4baf2a6e2cd154b73f387f8")
}

func (_TokenAdminRegistry *TokenAdminRegistry) Address() common.Address {
	return _TokenAdminRegistry.address
}

type TokenAdminRegistryInterface interface {
	GetAllConfiguredTokens(opts *bind.CallOpts, startIndex uint64, maxCount uint64) ([]common.Address, error)

	GetPool(opts *bind.CallOpts, token common.Address) (common.Address, error)

	GetPools(opts *bind.CallOpts, tokens []common.Address) ([]common.Address, error)

	GetTokenConfig(opts *bind.CallOpts, token common.Address) (TokenAdminRegistryTokenConfig, error)

	IsAdministrator(opts *bind.CallOpts, localToken common.Address, administrator common.Address) (bool, error)

	IsRegistryModule(opts *bind.CallOpts, module common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptAdminRole(opts *bind.TransactOpts, localToken common.Address) (*types.Transaction, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRegistryModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error)

	ProposeAdministrator(opts *bind.TransactOpts, localToken common.Address, administrator common.Address) (*types.Transaction, error)

	RemoveRegistryModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error)

	SetPool(opts *bind.TransactOpts, localToken common.Address, pool common.Address) (*types.Transaction, error)

	TransferAdminRole(opts *bind.TransactOpts, localToken common.Address, newAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAdministratorTransferRequested(opts *bind.FilterOpts, token []common.Address, currentAdmin []common.Address, newAdmin []common.Address) (*TokenAdminRegistryAdministratorTransferRequestedIterator, error)

	WatchAdministratorTransferRequested(opts *bind.WatchOpts, sink chan<- *TokenAdminRegistryAdministratorTransferRequested, token []common.Address, currentAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error)

	ParseAdministratorTransferRequested(log types.Log) (*TokenAdminRegistryAdministratorTransferRequested, error)

	FilterAdministratorTransferred(opts *bind.FilterOpts, token []common.Address, newAdmin []common.Address) (*TokenAdminRegistryAdministratorTransferredIterator, error)

	WatchAdministratorTransferred(opts *bind.WatchOpts, sink chan<- *TokenAdminRegistryAdministratorTransferred, token []common.Address, newAdmin []common.Address) (event.Subscription, error)

	ParseAdministratorTransferred(log types.Log) (*TokenAdminRegistryAdministratorTransferred, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenAdminRegistryOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *TokenAdminRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*TokenAdminRegistryOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenAdminRegistryOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenAdminRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*TokenAdminRegistryOwnershipTransferred, error)

	FilterPoolSet(opts *bind.FilterOpts, token []common.Address, previousPool []common.Address, newPool []common.Address) (*TokenAdminRegistryPoolSetIterator, error)

	WatchPoolSet(opts *bind.WatchOpts, sink chan<- *TokenAdminRegistryPoolSet, token []common.Address, previousPool []common.Address, newPool []common.Address) (event.Subscription, error)

	ParsePoolSet(log types.Log) (*TokenAdminRegistryPoolSet, error)

	FilterRegistryModuleAdded(opts *bind.FilterOpts) (*TokenAdminRegistryRegistryModuleAddedIterator, error)

	WatchRegistryModuleAdded(opts *bind.WatchOpts, sink chan<- *TokenAdminRegistryRegistryModuleAdded) (event.Subscription, error)

	ParseRegistryModuleAdded(log types.Log) (*TokenAdminRegistryRegistryModuleAdded, error)

	FilterRegistryModuleRemoved(opts *bind.FilterOpts, module []common.Address) (*TokenAdminRegistryRegistryModuleRemovedIterator, error)

	WatchRegistryModuleRemoved(opts *bind.WatchOpts, sink chan<- *TokenAdminRegistryRegistryModuleRemoved, module []common.Address) (event.Subscription, error)

	ParseRegistryModuleRemoved(log types.Log) (*TokenAdminRegistryRegistryModuleRemoved, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}

var TokenAdminRegistryZKBin = ("0x00030000000000020006000000000002000200000001035500000000030100190000006003300270000001360030019d0000008004000039000000400040043f0000000100200190000000470000c13d0000013602300197000000040020008c000004540000413d000000000301043b000000e0033002700000013c0030009c000000570000213d000001480030009c0000006d0000213d0000014e0030009c000000eb0000213d000001510030009c000001970000613d000001520030009c000004540000c13d000000240020008c000004540000413d0000000002000416000000000002004b000004540000c13d0000000401100370000000000101043b000400000001001d000001530010009c000004540000213d0000000401000029000000000010043f0000000201000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000000201043b0000000101200039000000000301041a00000153043001970000000006000411000000000064004b000003020000c13d000000000402041a0000013704400197000000000464019f000000000042041b0000013702300197000000000021041b0000000001000414000001360010009c0000013601008041000000c00110021000000154011001c70000800d02000039000000030300003900000176040000410000000405000029000001740000013d0000000001000416000000000001004b000004540000c13d0000000001000411000000000001004b000000640000c13d0000013901000041000000800010043f0000002001000039000000840010043f0000001801000039000000a40010043f0000013a01000041000000c40010043f0000013b01000041000004d4000104300000013d0030009c000000a40000213d000001430030009c000001390000213d000001460030009c000001c20000613d000001470030009c000004540000c13d0000000001000416000000000001004b000004540000c13d000000000100041a0000026c0000013d000000000200041a0000013702200197000000000112019f000000000010041b0000002001000039000001000010044300000120000004430000013801000041000004d30001042e000001490030009c000001580000213d0000014c0030009c000001d60000613d0000014d0030009c000004540000c13d000000240020008c000004540000413d0000000003000416000000000003004b000004540000c13d0000000403100370000000000303043b0000015b0030009c000004540000213d0000002304300039000000000024004b000004540000813d0000000404300039000000000441034f000000000404043b000200000004001d0000015b0040009c000004540000213d000100240030003d000000020300002900000005033002100000000104300029000000000024004b000004540000213d0000003f043000390000016804400197000001690040009c000004560000213d0000008004400039000000400040043f0000000204000029000000800040043f0000001f0430018f000000000003004b0000009d0000613d000000000121034f000000a002300039000000a003000039000000001501043c0000000003530436000000000023004b000000990000c13d000000000004004b000000020000006b000003920000c13d000000400100043d000400000001001d00000080020000390000031e0000013d0000013e0030009c000001790000213d000001410030009c0000021c0000613d000001420030009c000004540000c13d000000440020008c000004540000413d0000000002000416000000000002004b000004540000c13d0000000402100370000000000202043b000400000002001d000001530020009c000004540000213d0000002401100370000000000101043b000300000001001d000001530010009c000004540000213d0000000401000029000000000010043f0000000201000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000000101043b000000000101041a00000153011001970000000004000411000000000041004b000003290000c13d0000000401000029000000000010043f0000000201000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000000101043b0000000101100039000000000201041a00000137022001970000000307000029000000000272019f000000000021041b0000000001000414000001360010009c0000013601008041000000c00110021000000154011001c70000800d0200003900000004030000390000015d0400004100000004050000290000000006000411000001740000013d0000014f0030009c0000024c0000613d000001500030009c000004540000c13d000000240020008c000004540000413d0000000002000416000000000002004b000004540000c13d0000000401100370000000000101043b000400000001001d000001530010009c000004540000213d000000000100041a00000153011001970000000002000411000000000012004b000002ee0000c13d0000000401000029000000000010043f0000000601000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000000101043b000000000101041a000000000001004b000001770000c13d0000000501000039000000000201041a0000016e0020009c000004560000813d0000000103200039000000000031041b0000016f0220009a0000000403000029000000000032041b000000000101041a000300000001001d000000000030043f0000000601000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000000101043b0000000302000029000000000021041b000000400100043d00000004020000290000000000210435000001360010009c000001360100804100000040011002100000000002000414000001360020009c0000013602008041000000c002200210000000000112019f00000170011001c70000800d0200003900000001030000390000017104000041000001740000013d000001440030009c0000025c0000613d000001450030009c000004540000c13d000000440020008c000004540000413d0000000002000416000000000002004b000004540000c13d0000000402100370000000000202043b000001530020009c000004540000213d0000002401100370000000000101043b000400000001001d000001530010009c000004540000213d000000000020043f0000000201000039000000200010043f000000000100001904d204b70000040f000000000101041a0000015301100197000000040010006c00000000010000390000000101006039000000800010043f0000016301000041000004d30001042e0000014a0030009c000002700000613d0000014b0030009c000004540000c13d0000000001000416000000000001004b000004540000c13d0000000101000039000000000201041a00000153032001970000000006000411000000000036004b000002f80000c13d000000000300041a0000013704300197000000000464019f000000000040041b0000013702200197000000000021041b00000000010004140000015305300197000001360010009c0000013601008041000000c00110021000000154011001c70000800d020000390000000303000039000001660400004104d204c80000040f0000000100200190000004540000613d0000000001000019000004d30001042e0000013f0030009c000002bf0000613d000001400030009c000004540000c13d000000240020008c000004540000413d0000000002000416000000000002004b000004540000c13d0000000401100370000000000601043b000001530060009c000004540000213d000000000100041a00000153011001970000000005000411000000000015004b000002ee0000c13d000000000056004b0000030a0000c13d0000013901000041000000800010043f0000002001000039000000840010043f0000001701000039000000a40010043f0000015601000041000000c40010043f0000013b01000041000004d400010430000000240020008c000004540000413d0000000002000416000000000002004b000004540000c13d0000000401100370000000000101043b000400000001001d000001530010009c000004540000213d000000000100041a00000153011001970000000002000411000000000012004b000002ee0000c13d0000000401000029000000000010043f0000000601000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000000101043b000000000101041a000300000001001d000000000001004b000001770000613d0000000504000039000000000204041a000000000002004b000003b60000c13d0000017201000041000000000010043f0000001101000039000000040010043f0000015901000041000004d400010430000000240020008c000004540000413d0000000002000416000000000002004b000004540000c13d0000000401100370000000000101043b000001530010009c000004540000213d04d204a30000040f000000000001004b0000000001000039000000010100c039000000400200043d0000000000120435000001360020009c0000013602008041000000400120021000000164011001c7000004d30001042e000000440020008c000004540000413d0000000002000416000000000002004b000004540000c13d0000000402100370000000000202043b000400000002001d000001530020009c000004540000213d0000002401100370000000000101043b000300000001001d000001530010009c000004540000213d0000000401000029000000000010043f0000000201000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000000101043b000000000101041a00000153021001970000000001000411000000000012004b000003370000c13d0000000303000029000000000003004b000003d50000c13d0000000401000029000000000010043f0000000201000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000000101043b0000000201100039000000000201041a000001370320019700000003033001af000000000031041b0000015306200197000000030060006c000001770000613d0000000001000414000001360010009c0000013601008041000000c00110021000000154011001c70000800d0200003900000004030000390000016d0400004100000004050000290000000307000029000001740000013d000000240020008c000004540000413d0000000002000416000000000002004b000004540000c13d0000000401100370000000000101043b000001530010009c000004540000213d000000e002000039000000400020043f000000800000043f000000a00000043f000000c00000043f000000000010043f0000000201000039000000200010043f000000000100001904d204b70000040f000400000001001d000000e00100003904d204880000040f0000000403000029000000000103041a0000015301100197000000e00010043f0000000102300039000000000202041a0000015302200197000001000020043f0000000202300039000000000202041a0000015302200197000001200020043f000000400200043d0000000001120436000001000300043d00000153033001970000000000310435000001200100043d000001530110019700000040032000390000000000130435000001360020009c0000013602008041000000400120021000000162011001c7000004d30001042e0000000001000416000000000001004b000004540000c13d000000c001000039000000400010043f0000001801000039000000800010043f0000017302000041000000a00020043f0000002003000039000000c00030043f000000e00010043f000001000020043f000001180000043f0000017401000041000004d30001042e000000240020008c000004540000413d0000000002000416000000000002004b000004540000c13d0000000401100370000000000101043b000001530010009c000004540000213d000000000010043f0000000201000039000000200010043f000000000100001904d204b70000040f0000000201100039000000000101041a0000015301100197000000800010043f0000016301000041000004d30001042e000000440020008c000004540000413d0000000002000416000000000002004b000004540000c13d0000000402100370000000000302043b0000015b0030009c000004540000213d0000002401100370000000000601043b0000015b0060009c000004540000213d00000000010004150000000304000039000000000504041a000000000235004b000003180000a13d0000015b0660019700000000083600190000000007000415000000060770008a0000000507700210000000000058004b0000028f0000a13d0000000007000415000000050770008a00000005077002100000015b0020009c0000000006020019000004560000213d00000005096002100000003f02900039000001670a200197000000400200043d00000000082a00190000000000a8004b000000000a000039000000010a0040390000015b0080009c000004560000213d0000000100a00190000004560000c13d000000400080043f00000000086204360000001f0a90018f000000000009004b000002a80000613d0000000009980019000000000b000031000000020bb00367000000000c08001900000000bd0b043c000000000cdc043600000000009c004b000002a40000c13d00000000000a004b0000000507700270000000000702001f000000000006004b000003190000613d00000000070000190000000009370019000000000059004b000004090000813d000000000040043f000000000a02043300000000007a004b000004090000a13d000000050a700210000000000a8a00190000015c0990009a000000000909041a000001530990019700000000009a04350000000107700039000000000067004b000002ae0000413d000003190000013d000000440020008c000004540000413d0000000002000416000000000002004b000004540000c13d0000000402100370000000000202043b000400000002001d000001530020009c000004540000213d0000002401100370000000000101043b000300000001001d000001530010009c000004540000213d0000000001000411000000000010043f0000000601000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000000101043b000000000101041a000000000001004b000002e40000c13d000000000100041a00000153011001970000000002000411000000000012004b000003e30000c13d000000030000006b000003450000c13d000000400100043d0000015e020000410000000000210435000001360010009c000001360100804100000040011002100000015f011001c7000004d4000104300000013901000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f0000017701000041000000c40010043f0000013b01000041000004d4000104300000013901000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f0000016501000041000000c40010043f0000013b01000041000004d400010430000000400100043d000000240210003900000004030000290000000000320435000001750200004100000000002104350000015302600197000003300000013d0000000101000039000000000201041a0000013702200197000000000262019f000000000021041b0000000001000414000001360010009c0000013601008041000000c00110021000000154011001c70000800d0200003900000003030000390000015504000041000001740000013d0000006002000039000000000300041500000000013100490000000001000002000000400100043d000400000001001d04d204930000040f00000004020000290000000001210049000001360010009c00000136010080410000006001100210000001360020009c00000136020080410000004002200210000000000121019f000004d30001042e000000400100043d00000024021000390000000403000029000000000032043500000160020000410000000000210435000001530240019700000004031000390000000000230435000001360010009c0000013601008041000000400110021000000161011001c7000004d400010430000000400200043d00000024032000390000000404000029000000000043043500000160030000410000000000320435000001530110019700000004032000390000000000130435000001360020009c0000013602008041000000400120021000000161011001c7000004d4000104300000000401000029000000000010043f0000000201000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000000101043b000000000201041a00000153002001980000045c0000c13d0000000101100039000000000201041a000001370220019700000003022001af000000000021041b0000000401000029000000000010043f0000000401000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000000101043b000000000101041a000000000001004b000003860000c13d0000000301000039000000000201041a0000015b0020009c000004560000213d0000000103200039000000000031041b0000015c0220009a0000000403000029000000000032041b000000000101041a000200000001001d000000000030043f0000000401000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000000101043b0000000202000029000000000021041b0000000001000414000001360010009c0000013601008041000000c00110021000000154011001c70000800d0200003900000004030000390000015d04000041000000040500002900000000060000190000000307000029000001740000013d0000000003000019000400000003001d0000000502300210000300000002001d00000001012000290000000201100367000000000101043b000001530010009c000004540000213d000000000010043f0000000201000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000800200043d0000000403000029000000000032004b000004090000a13d0000000302000029000000a002200039000000000101043b0000000201100039000000000101041a000001530110019700000000001204350000000103300039000000020030006c000003930000413d000000a00000013d0000000303000029000000010130008a000000000023004b000003e90000c13d000001780230009a000000000002041b000000000014041b0000000401000029000000000010043f0000000601000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000000101043b000000000001041b0000000001000414000001360010009c0000013601008041000000c00110021000000154011001c70000800d0200003900000002030000390000017904000041000000450000013d000000400b00043d0000016a0100004100000000001b04350000000401b00039000000040200002900000000002104350000000001000414000000040030008c0000040f0000c13d0000000103000031000000200030008c000000200400003900000000040340190000043c0000013d000000400100043d0000015802000041000000000021043500000004021000390000000003000411000004610000013d000000000012004b000004090000a13d000001780130009a000001780220009a000000000202041a000000000021041b000000000020043f0000000601000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004540000613d000000000101043b0000000302000029000000000021041b0000000501000039000000000101041a0000000003010019000000000001004b000004850000c13d0000017201000041000000000010043f0000003101000039000000040010043f0000015901000041000004d4000104300000017201000041000000000010043f0000003201000039000000040010043f0000015901000041000004d4000104300000013600b0009c000001360200004100000000020b40190000004002200210000001360010009c0000013601008041000000c001100210000000000121019f00000159011001c7000000000203001900020000000b001d04d204cd0000040f000000000301001900000060033002700000013603300197000000200030008c000000200400003900000000040340190000001f0640018f0000002007400190000000020b00002900000002057000290000042c0000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b000004280000c13d000000000006004b000004390000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000004670000613d0000001f01400039000000600210018f00000000040b00190000000001b20019000000000021004b000000000200003900000001020040390000015b0010009c000004560000213d0000000100200190000004560000c13d000000400010043f000000200030008c000004540000413d0000000002040433000000000002004b0000000003000039000000010300c039000000000032004b000004540000c13d000000000002004b000001fb0000c13d0000016c020000410000045e0000013d0000000001000019000004d4000104300000017201000041000000000010043f0000004101000039000000040010043f0000015901000041000004d400010430000000400100043d0000015a020000410000000000210435000000040210003900000004030000290000000000320435000001360010009c0000013601008041000000400110021000000159011001c7000004d4000104300000001f0530018f0000016b06300198000000400200043d0000000004620019000004720000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b0000046e0000c13d000000000005004b0000047f0000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000001360020009c00000136020080410000004002200210000000000112019f000004d400010430000000010130008a0000000504000039000003ba0000013d0000017a0010009c0000048d0000813d0000006001100039000000400010043f000000000001042d0000017201000041000000000010043f0000004101000039000000040010043f0000015901000041000004d40001043000000020030000390000000004310436000000000302043300000000003404350000004001100039000000000003004b000004a20000613d000000000400001900000020022000390000000005020433000001530550019700000000015104360000000104400039000000000034004b0000049b0000413d000000000001042d000000000010043f0000000601000039000000200010043f0000000001000414000001360010009c0000013601008041000000c00110021000000157011001c7000080100200003904d204cd0000040f0000000100200190000004b50000613d000000000101043b000000000101041a000000000001004b0000000001000039000000010100c039000000000001042d0000000001000019000004d4000104300000000002000414000001360020009c0000013602008041000000c002200210000001360010009c00000136010080410000004001100210000000000121019f00000157011001c7000080100200003904d204cd0000040f0000000100200190000004c60000613d000000000101043b000000000001042d0000000001000019000004d400010430000004cb002104210000000102000039000000000001042d0000000002000019000000000001042d000004d0002104230000000102000039000000000001042d0000000002000019000000000001042d000004d200000432000004d30001042e000004d40001043000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000020000000000000000000000000000004000000100000000000000000008c379a00000000000000000000000000000000000000000000000000000000043616e6e6f7420736574206f776e657220746f207a65726f00000000000000000000000000000000000000000000000000000064000000800000000000000000000000000000000000000000000000000000000000000000000000007d3f255100000000000000000000000000000000000000000000000000000000cb67e3b000000000000000000000000000000000000000000000000000000000e677ae3600000000000000000000000000000000000000000000000000000000e677ae3700000000000000000000000000000000000000000000000000000000f2fde38b00000000000000000000000000000000000000000000000000000000cb67e3b100000000000000000000000000000000000000000000000000000000ddadfa8e00000000000000000000000000000000000000000000000000000000bbe4f6da00000000000000000000000000000000000000000000000000000000bbe4f6db00000000000000000000000000000000000000000000000000000000c1af6e03000000000000000000000000000000000000000000000000000000007d3f2552000000000000000000000000000000000000000000000000000000008da5cb5b000000000000000000000000000000000000000000000000000000004e847fc60000000000000000000000000000000000000000000000000000000072d64a800000000000000000000000000000000000000000000000000000000072d64a810000000000000000000000000000000000000000000000000000000079ba5097000000000000000000000000000000000000000000000000000000004e847fc7000000000000000000000000000000000000000000000000000000005e63547a00000000000000000000000000000000000000000000000000000000181f5a7600000000000000000000000000000000000000000000000000000000181f5a77000000000000000000000000000000000000000000000000000000003dc457720000000000000000000000000000000000000000000000000000000010cbcf1800000000000000000000000000000000000000000000000000000000156194da000000000000000000000000ffffffffffffffffffffffffffffffffffffffff0200000000000000000000000000000000000000000000000000000000000000ed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127843616e6e6f74207472616e7366657220746f2073656c66000000000000000000020000000000000000000000000000000000004000000000000000000000000051ca1ec300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002400000000000000000000000045ed80e900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffff3da8a5f161a6c3ff06a60736d0ed24d7963cc6a5c4fafd2fa1dae9bb908e07a5c54c3051ff16e63bb9203214432372aca006c589e3653619b577a3265675b716d92e233d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000ed5d85b50000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000440000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000002000000080000000000000000000000000000000000000000000000000000000200000000000000000000000004d7573742062652070726f706f736564206f776e6572000000000000000000008be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e000000000000000000000000000000000000000000000003fffffffffffffffe07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0000000000000000000000000000000000000000000000000ffffffffffffff7f240028e80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffe0962b60e600000000000000000000000000000000000000000000000000000000754449ec3aff3bd528bfce43ae9319c4a381b67fcd1d20097b3b24dacaecc35d0000000000000000000000000000000000000000000000010000000000000000fc949c7b4a13586e39d89eead2f38644f9fb3efb5a0490b14f8fc0ceab44c25002000000000000000000000000000000000000200000000000000000000000003cabf004338366bfeaeb610ad827cb58d16b588017c509501f2c97c83caae7b24e487b7100000000000000000000000000000000000000000000000000000000546f6b656e41646d696e526567697374727920312e352e3000000000000000000000000000000000000000000000000000000060000000c000000000000000003edffe7500000000000000000000000000000000000000000000000000000000399b55200f7f639a63d76efe3dcfa9156ce367058d6b673041b84a628885f5a74f6e6c792063616c6c61626c65206279206f776e657200000000000000000000fc949c7b4a13586e39d89eead2f38644f9fb3efb5a0490b14f8fc0ceab44c25193eaa26dcb9275e56bacb1d33fdbf402262da6f0f4baf2a6e2cd154b73f387f8000000000000000000000000000000000000000000000000ffffffffffffffa0")
