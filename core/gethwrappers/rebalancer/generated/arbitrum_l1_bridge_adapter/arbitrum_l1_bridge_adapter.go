// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum_l1_bridge_adapter

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

type ArbitrumL1BridgeAdapterArbitrumFinalizationPayload struct {
	Proof       [][32]byte
	Index       *big.Int
	L2Sender    common.Address
	To          common.Address
	L2Block     *big.Int
	L1Block     *big.Int
	L2Timestamp *big.Int
	Value       *big.Int
	Data        []byte
}

type ArbitrumL1BridgeAdapterSendERC20Params struct {
	GasLimit          *big.Int
	MaxSubmissionCost *big.Int
	MaxFeePerGas      *big.Int
}

var ArbitrumL1BridgeAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIL1GatewayRouter\",\"name\":\"l1GatewayRouter\",\"type\":\"address\"},{\"internalType\":\"contractIOutbox\",\"name\":\"l1Outbox\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MsgShouldNotContainValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NoGatewayForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unimplemented\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structArbitrumL1BridgeAdapter.ArbitrumFinalizationPayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"exposeArbitrumFinalizationPayload\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"}],\"internalType\":\"structArbitrumL1BridgeAdapter.SendERC20Params\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exposeSendERC20Params\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arbitrumFinalizationPayload\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeFeeInNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"}],\"name\":\"getL2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"bridgeSpecificPayload\",\"type\":\"bytes\"}],\"name\":\"sendERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60c0604052600080546001600160401b031916905534801561002057600080fd5b506040516200144b3803806200144b833981016040819052610041916100ab565b6001600160a01b038216158061005e57506001600160a01b038116155b1561007c57604051635e9c404d60e11b815260040160405180910390fd5b6001600160a01b039182166080521660a0526100e5565b6001600160a01b03811681146100a857600080fd5b50565b600080604083850312156100be57600080fd5b82516100c981610093565b60208401519092506100da81610093565b809150509250929050565b60805160a0516113326200011960003960006101ff0152600081816102e10152818161045701526105b201526113326000f3fe6080604052600436106100655760003560e01c8063b5399c9e11610043578063b5399c9e146100d4578063c7665dd2146100f2578063c985069c1461010d57600080fd5b80632e4b1fc91461006a57806338314bb214610092578063a71d98b7146100b4575b600080fd5b34801561007657600080fd5b5061007f610152565b6040519081526020015b60405180910390f35b34801561009e57600080fd5b506100b26100ad366004610c24565b610186565b005b6100c76100c2366004610c89565b610275565b6040516100899190610d7c565b3480156100e057600080fd5b506100b26100ef366004610e3e565b50565b3480156100fe57600080fd5b506100b26100ef366004610fab565b34801561011957600080fd5b5061012d610128366004611095565b61056a565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610089565b60006040517f6e12839900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061019482840184610fab565b805160208201516040808401516060850151608086015160a087015160c088015160e08901516101008a015196517f08635a95000000000000000000000000000000000000000000000000000000008152999a5073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016996308635a959961023c99909890979695949392916004016110b2565b600060405180830381600087803b15801561025657600080fd5b505af115801561026a573d6000803e3d6000fd5b505050505050505050565b606061029973ffffffffffffffffffffffffffffffffffffffff8816333087610625565b6040517fbda009fe00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff88811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063bda009fe90602401602060405180830381865afa15801561032a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061034e9190611168565b905073ffffffffffffffffffffffffffffffffffffffff81166103ba576040517f6c1460f400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff891660048201526024015b60405180910390fd5b6103db73ffffffffffffffffffffffffffffffffffffffff89168287610707565b60006103e984860186610e3e565b9050600081602001518260400151836000015161040691906111b4565b61041091906111cb565b905080341015610455576040517f03da4d23000000000000000000000000000000000000000000000000000000008152346004820152602481018290526044016103b1565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16634fb1a07b348c8b8c8c886000015189604001518a60200151604051806020016040528060008152506040516020016104c69291906111de565b6040516020818303038152906040526040518963ffffffff1660e01b81526004016104f797969594939291906111f7565b60006040518083038185885af1158015610515573d6000803e3d6000fd5b50505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261055c9190810190611257565b9a9950505050505050505050565b6040517fa7e28d4800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063a7e28d4890602401602060405180830381865afa1580156105fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061f9190611168565b92915050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526107019085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261088e565b50505050565b8015806107a757506040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff838116602483015284169063dd62ed3e90604401602060405180830381865afa158015610781573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107a591906112ce565b155b610833576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e63650000000000000000000060648201526084016103b1565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526108899084907f095ea7b3000000000000000000000000000000000000000000000000000000009060640161067f565b505050565b60006108f0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661099a9092919063ffffffff16565b805190915015610889578080602001905181019061090e91906112e7565b610889576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016103b1565b60606109a984846000856109b1565b949350505050565b606082471015610a43576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016103b1565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610a6c9190611309565b60006040518083038185875af1925050503d8060008114610aa9576040519150601f19603f3d011682016040523d82523d6000602084013e610aae565b606091505b5091509150610abf87838387610aca565b979650505050505050565b60608315610b60578251600003610b595773ffffffffffffffffffffffffffffffffffffffff85163b610b59576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103b1565b50816109a9565b6109a98383815115610b755781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b19190610d7c565b73ffffffffffffffffffffffffffffffffffffffff811681146100ef57600080fd5b8035610bd681610ba9565b919050565b60008083601f840112610bed57600080fd5b50813567ffffffffffffffff811115610c0557600080fd5b602083019150836020828501011115610c1d57600080fd5b9250929050565b60008060008060608587031215610c3a57600080fd5b8435610c4581610ba9565b93506020850135610c5581610ba9565b9250604085013567ffffffffffffffff811115610c7157600080fd5b610c7d87828801610bdb565b95989497509550505050565b60008060008060008060a08789031215610ca257600080fd5b8635610cad81610ba9565b95506020870135610cbd81610ba9565b94506040870135610ccd81610ba9565b935060608701359250608087013567ffffffffffffffff811115610cf057600080fd5b610cfc89828a01610bdb565b979a9699509497509295939492505050565b60005b83811015610d29578181015183820152602001610d11565b50506000910152565b60008151808452610d4a816020860160208601610d0e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610d8f6020830184610d32565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff81118282101715610de957610de9610d96565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610e3657610e36610d96565b604052919050565b600060608284031215610e5057600080fd5b6040516060810181811067ffffffffffffffff82111715610e7357610e73610d96565b80604052508235815260208301356020820152604083013560408201528091505092915050565b600082601f830112610eab57600080fd5b8135602067ffffffffffffffff821115610ec757610ec7610d96565b8160051b610ed6828201610def565b9283528481018201928281019087851115610ef057600080fd5b83870192505b84831015610abf57823582529183019190830190610ef6565b600067ffffffffffffffff821115610f2957610f29610d96565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112610f6657600080fd5b8135610f79610f7482610f0f565b610def565b818152846020838601011115610f8e57600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215610fbd57600080fd5b813567ffffffffffffffff80821115610fd557600080fd5b908301906101208286031215610fea57600080fd5b610ff2610dc5565b82358281111561100157600080fd5b61100d87828601610e9a565b8252506020830135602082015261102660408401610bcb565b604082015261103760608401610bcb565b60608201526080830135608082015260a083013560a082015260c083013560c082015260e083013560e0820152610100808401358381111561107857600080fd5b61108488828701610f55565b918301919091525095945050505050565b6000602082840312156110a757600080fd5b8135610d8f81610ba9565b6101208082528a51908201819052600090610140830190602090818e01845b828110156110ed578151855293830193908301906001016110d1565b50505083018b905273ffffffffffffffffffffffffffffffffffffffff8a16604084015273ffffffffffffffffffffffffffffffffffffffff891660608401528760808401528660a08401528560c08401528460e08401528281036101008401526111588185610d32565b9c9b505050505050505050505050565b60006020828403121561117a57600080fd5b8151610d8f81610ba9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808202811582820484141761061f5761061f611185565b8082018082111561061f5761061f611185565b8281526040602082015260006109a96040830184610d32565b600073ffffffffffffffffffffffffffffffffffffffff808a16835280891660208401528088166040840152508560608301528460808301528360a083015260e060c083015261124a60e0830184610d32565b9998505050505050505050565b60006020828403121561126957600080fd5b815167ffffffffffffffff81111561128057600080fd5b8201601f8101841361129157600080fd5b805161129f610f7482610f0f565b8181528560208385010111156112b457600080fd5b6112c5826020830160208601610d0e565b95945050505050565b6000602082840312156112e057600080fd5b5051919050565b6000602082840312156112f957600080fd5b81518015158114610d8f57600080fd5b6000825161131b818460208701610d0e565b919091019291505056fea164736f6c6343000813000a",
}

var ArbitrumL1BridgeAdapterABI = ArbitrumL1BridgeAdapterMetaData.ABI

var ArbitrumL1BridgeAdapterBin = ArbitrumL1BridgeAdapterMetaData.Bin

func DeployArbitrumL1BridgeAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, l1GatewayRouter common.Address, l1Outbox common.Address) (common.Address, *types.Transaction, *ArbitrumL1BridgeAdapter, error) {
	parsed, err := ArbitrumL1BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArbitrumL1BridgeAdapterBin), backend, l1GatewayRouter, l1Outbox)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbitrumL1BridgeAdapter{address: address, abi: *parsed, ArbitrumL1BridgeAdapterCaller: ArbitrumL1BridgeAdapterCaller{contract: contract}, ArbitrumL1BridgeAdapterTransactor: ArbitrumL1BridgeAdapterTransactor{contract: contract}, ArbitrumL1BridgeAdapterFilterer: ArbitrumL1BridgeAdapterFilterer{contract: contract}}, nil
}

type ArbitrumL1BridgeAdapter struct {
	address common.Address
	abi     abi.ABI
	ArbitrumL1BridgeAdapterCaller
	ArbitrumL1BridgeAdapterTransactor
	ArbitrumL1BridgeAdapterFilterer
}

type ArbitrumL1BridgeAdapterCaller struct {
	contract *bind.BoundContract
}

type ArbitrumL1BridgeAdapterTransactor struct {
	contract *bind.BoundContract
}

type ArbitrumL1BridgeAdapterFilterer struct {
	contract *bind.BoundContract
}

type ArbitrumL1BridgeAdapterSession struct {
	Contract     *ArbitrumL1BridgeAdapter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ArbitrumL1BridgeAdapterCallerSession struct {
	Contract *ArbitrumL1BridgeAdapterCaller
	CallOpts bind.CallOpts
}

type ArbitrumL1BridgeAdapterTransactorSession struct {
	Contract     *ArbitrumL1BridgeAdapterTransactor
	TransactOpts bind.TransactOpts
}

type ArbitrumL1BridgeAdapterRaw struct {
	Contract *ArbitrumL1BridgeAdapter
}

type ArbitrumL1BridgeAdapterCallerRaw struct {
	Contract *ArbitrumL1BridgeAdapterCaller
}

type ArbitrumL1BridgeAdapterTransactorRaw struct {
	Contract *ArbitrumL1BridgeAdapterTransactor
}

func NewArbitrumL1BridgeAdapter(address common.Address, backend bind.ContractBackend) (*ArbitrumL1BridgeAdapter, error) {
	abi, err := abi.JSON(strings.NewReader(ArbitrumL1BridgeAdapterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindArbitrumL1BridgeAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapter{address: address, abi: abi, ArbitrumL1BridgeAdapterCaller: ArbitrumL1BridgeAdapterCaller{contract: contract}, ArbitrumL1BridgeAdapterTransactor: ArbitrumL1BridgeAdapterTransactor{contract: contract}, ArbitrumL1BridgeAdapterFilterer: ArbitrumL1BridgeAdapterFilterer{contract: contract}}, nil
}

func NewArbitrumL1BridgeAdapterCaller(address common.Address, caller bind.ContractCaller) (*ArbitrumL1BridgeAdapterCaller, error) {
	contract, err := bindArbitrumL1BridgeAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapterCaller{contract: contract}, nil
}

func NewArbitrumL1BridgeAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrumL1BridgeAdapterTransactor, error) {
	contract, err := bindArbitrumL1BridgeAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapterTransactor{contract: contract}, nil
}

func NewArbitrumL1BridgeAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrumL1BridgeAdapterFilterer, error) {
	contract, err := bindArbitrumL1BridgeAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapterFilterer{contract: contract}, nil
}

func bindArbitrumL1BridgeAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbitrumL1BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumL1BridgeAdapter.Contract.ArbitrumL1BridgeAdapterCaller.contract.Call(opts, result, method, params...)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.ArbitrumL1BridgeAdapterTransactor.contract.Transfer(opts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.ArbitrumL1BridgeAdapterTransactor.contract.Transact(opts, method, params...)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumL1BridgeAdapter.Contract.contract.Call(opts, result, method, params...)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.contract.Transfer(opts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.contract.Transact(opts, method, params...)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) ExposeArbitrumFinalizationPayload(opts *bind.CallOpts, payload ArbitrumL1BridgeAdapterArbitrumFinalizationPayload) error {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "exposeArbitrumFinalizationPayload", payload)

	if err != nil {
		return err
	}

	return err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) ExposeArbitrumFinalizationPayload(payload ArbitrumL1BridgeAdapterArbitrumFinalizationPayload) error {
	return _ArbitrumL1BridgeAdapter.Contract.ExposeArbitrumFinalizationPayload(&_ArbitrumL1BridgeAdapter.CallOpts, payload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) ExposeArbitrumFinalizationPayload(payload ArbitrumL1BridgeAdapterArbitrumFinalizationPayload) error {
	return _ArbitrumL1BridgeAdapter.Contract.ExposeArbitrumFinalizationPayload(&_ArbitrumL1BridgeAdapter.CallOpts, payload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) ExposeSendERC20Params(opts *bind.CallOpts, params ArbitrumL1BridgeAdapterSendERC20Params) error {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "exposeSendERC20Params", params)

	if err != nil {
		return err
	}

	return err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) ExposeSendERC20Params(params ArbitrumL1BridgeAdapterSendERC20Params) error {
	return _ArbitrumL1BridgeAdapter.Contract.ExposeSendERC20Params(&_ArbitrumL1BridgeAdapter.CallOpts, params)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) ExposeSendERC20Params(params ArbitrumL1BridgeAdapterSendERC20Params) error {
	return _ArbitrumL1BridgeAdapter.Contract.ExposeSendERC20Params(&_ArbitrumL1BridgeAdapter.CallOpts, params)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "getBridgeFeeInNative")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GetBridgeFeeInNative(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GetBridgeFeeInNative(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) GetL2Token(opts *bind.CallOpts, l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "getL2Token", l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) GetL2Token(l1Token common.Address) (common.Address, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GetL2Token(&_ArbitrumL1BridgeAdapter.CallOpts, l1Token)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) GetL2Token(l1Token common.Address) (common.Address, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GetL2Token(&_ArbitrumL1BridgeAdapter.CallOpts, l1Token)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.contract.Transact(opts, "finalizeWithdrawERC20", arg0, arg1, arbitrumFinalizationPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.FinalizeWithdrawERC20(&_ArbitrumL1BridgeAdapter.TransactOpts, arg0, arg1, arbitrumFinalizationPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactorSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.FinalizeWithdrawERC20(&_ArbitrumL1BridgeAdapter.TransactOpts, arg0, arg1, arbitrumFinalizationPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactor) SendERC20(opts *bind.TransactOpts, localToken common.Address, arg1 common.Address, recipient common.Address, amount *big.Int, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.contract.Transact(opts, "sendERC20", localToken, arg1, recipient, amount, bridgeSpecificPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) SendERC20(localToken common.Address, arg1 common.Address, recipient common.Address, amount *big.Int, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.SendERC20(&_ArbitrumL1BridgeAdapter.TransactOpts, localToken, arg1, recipient, amount, bridgeSpecificPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactorSession) SendERC20(localToken common.Address, arg1 common.Address, recipient common.Address, amount *big.Int, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.SendERC20(&_ArbitrumL1BridgeAdapter.TransactOpts, localToken, arg1, recipient, amount, bridgeSpecificPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapter) Address() common.Address {
	return _ArbitrumL1BridgeAdapter.address
}

type ArbitrumL1BridgeAdapterInterface interface {
	ExposeArbitrumFinalizationPayload(opts *bind.CallOpts, payload ArbitrumL1BridgeAdapterArbitrumFinalizationPayload) error

	ExposeSendERC20Params(opts *bind.CallOpts, params ArbitrumL1BridgeAdapterSendERC20Params) error

	GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error)

	GetL2Token(opts *bind.CallOpts, l1Token common.Address) (common.Address, error)

	FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, localToken common.Address, arg1 common.Address, recipient common.Address, amount *big.Int, bridgeSpecificPayload []byte) (*types.Transaction, error)

	Address() common.Address
}
