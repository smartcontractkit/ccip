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

var ArbitrumL1BridgeAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIL1GatewayRouter\",\"name\":\"l1GatewayRouter\",\"type\":\"address\"},{\"internalType\":\"contractIOutbox\",\"name\":\"l1Outbox\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MsgShouldNotContainValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NoGatewayForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GAS_PRICE_BID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SUBMISSION_COST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structArbitrumL1BridgeAdapter.ArbitrumFinalizationPayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"exposeForEncoding\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arbitrumFinalizationPayload\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeFeeInNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"}],\"name\":\"getL2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60c0604052600080546001600160401b031916905534801561002057600080fd5b506040516200142038038062001420833981016040819052610041916100ab565b6001600160a01b038216158061005e57506001600160a01b038116155b1561007c57604051635e9c404d60e11b815260040160405180910390fd5b6001600160a01b039182166080521660a0526100e5565b6001600160a01b03811681146100a857600080fd5b50565b600080604083850312156100be57600080fd5b82516100c981610093565b60208401519092506100da81610093565b809150509250929050565b60805160a05161130762000119600039600061023b0152600081816103460152818161049101526105ed01526113076000f3fe60806040526004361061007b5760003560e01c80635f2a9f411161004e5780635f2a9f411461010f57806379a35b4b14610126578063c985069c14610139578063ea6c2f801461017e57600080fd5b80632e4b1fc91461008057806332eb7905146100a857806338314bb2146100c05780633de17fe4146100e2575b600080fd5b34801561008c57600080fd5b50610095610199565b6040519081526020015b60405180910390f35b3480156100b457600080fd5b506100956311e1a30081565b3480156100cc57600080fd5b506100e06100db366004610b6b565b6101c2565b005b3480156100ee57600080fd5b506101026100fd366004610dc1565b6102b1565b60405161009f9190610f19565b34801561011b57600080fd5b50610095620186a081565b610102610134366004610f33565b6102da565b34801561014557600080fd5b50610159610154366004610f84565b6105a5565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161009f565b34801561018a57600080fd5b506100956602d79883d2000081565b60006101ac6311e1a300620186a0610fd0565b6101bd906602d79883d20000610fe7565b905090565b60006101d082840184610dc1565b805160208201516040808401516060850151608086015160a087015160c088015160e08901516101008a015196517f08635a95000000000000000000000000000000000000000000000000000000008152999a5073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016996308635a95996102789990989097969594939291600401611035565b600060405180830381600087803b15801561029257600080fd5b505af11580156102a6573d6000803e3d6000fd5b505050505050505050565b6060816040516020016102c491906110b2565b6040516020818303038152906040529050919050565b60606102fe73ffffffffffffffffffffffffffffffffffffffff8616333085610660565b6040517fbda009fe00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063bda009fe90602401602060405180830381865afa15801561038f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b39190611196565b905073ffffffffffffffffffffffffffffffffffffffff811661041f576040517f6c1460f400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff871660048201526024015b60405180910390fd5b61044073ffffffffffffffffffffffffffffffffffffffff87168285610742565b600061044a610199565b90508034101561048f576040517fe2c5a8f700000000000000000000000000000000000000000000000000000000815260048101829052346024820152604401610416565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16634fb1a07b3489888989620186a06311e1a3006602d79883d20000604051806020016040528060008152506040516020016105029291906111b3565b6040516020818303038152906040526040518963ffffffff1660e01b815260040161053397969594939291906111cc565b60006040518083038185885af1158015610551573d6000803e3d6000fd5b50505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610598919081019061122c565b925050505b949350505050565b6040517fa7e28d4800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063a7e28d4890602401602060405180830381865afa158015610636573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061065a9190611196565b92915050565b60405173ffffffffffffffffffffffffffffffffffffffff8085166024830152831660448201526064810182905261073c9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526108c9565b50505050565b8015806107e257506040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff838116602483015284169063dd62ed3e90604401602060405180830381865afa1580156107bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e091906112a3565b155b61086e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608401610416565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526108c49084907f095ea7b300000000000000000000000000000000000000000000000000000000906064016106ba565b505050565b600061092b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166109d59092919063ffffffff16565b8051909150156108c4578080602001905181019061094991906112bc565b6108c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610416565b606061059d8484600085856000808673ffffffffffffffffffffffffffffffffffffffff168587604051610a0991906112de565b60006040518083038185875af1925050503d8060008114610a46576040519150601f19603f3d011682016040523d82523d6000602084013e610a4b565b606091505b50915091506105988783838760608315610aed578251600003610ae65773ffffffffffffffffffffffffffffffffffffffff85163b610ae6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610416565b508161059d565b61059d8383815115610b025781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104169190610f19565b73ffffffffffffffffffffffffffffffffffffffff81168114610b5857600080fd5b50565b8035610b6681610b36565b919050565b60008060008060608587031215610b8157600080fd5b8435610b8c81610b36565b93506020850135610b9c81610b36565b9250604085013567ffffffffffffffff80821115610bb957600080fd5b818701915087601f830112610bcd57600080fd5b813581811115610bdc57600080fd5b886020828501011115610bee57600080fd5b95989497505060200194505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff81118282101715610c5057610c50610bfd565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610c9d57610c9d610bfd565b604052919050565b600082601f830112610cb657600080fd5b8135602067ffffffffffffffff821115610cd257610cd2610bfd565b8160051b610ce1828201610c56565b9283528481018201928281019087851115610cfb57600080fd5b83870192505b84831015610d1a57823582529183019190830190610d01565b979650505050505050565b600067ffffffffffffffff821115610d3f57610d3f610bfd565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112610d7c57600080fd5b8135610d8f610d8a82610d25565b610c56565b818152846020838601011115610da457600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215610dd357600080fd5b813567ffffffffffffffff80821115610deb57600080fd5b908301906101208286031215610e0057600080fd5b610e08610c2c565b823582811115610e1757600080fd5b610e2387828601610ca5565b82525060208301356020820152610e3c60408401610b5b565b6040820152610e4d60608401610b5b565b60608201526080830135608082015260a083013560a082015260c083013560c082015260e083013560e08201526101008084013583811115610e8e57600080fd5b610e9a88828701610d6b565b918301919091525095945050505050565b60005b83811015610ec6578181015183820152602001610eae565b50506000910152565b60008151808452610ee7816020860160208601610eab565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610f2c6020830184610ecf565b9392505050565b60008060008060808587031215610f4957600080fd5b8435610f5481610b36565b93506020850135610f6481610b36565b92506040850135610f7481610b36565b9396929550929360600135925050565b600060208284031215610f9657600080fd5b8135610f2c81610b36565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808202811582820484141761065a5761065a610fa1565b8082018082111561065a5761065a610fa1565b600081518084526020808501945080840160005b8381101561102a5781518752958201959082019060010161100e565b509495945050505050565b60006101208083526110498184018d610ffa565b90508a602084015273ffffffffffffffffffffffffffffffffffffffff808b166040850152808a166060850152508760808401528660a08401528560c08401528460e08401528281036101008401526110a28185610ecf565b9c9b505050505050505050505050565b60208152600082516101208060208501526110d1610140850183610ffa565b9150602085015160408501526040850151611104606086018273ffffffffffffffffffffffffffffffffffffffff169052565b50606085015173ffffffffffffffffffffffffffffffffffffffff8116608086015250608085015160a085015260a085015160c085015260c085015160e085015260e08501516101008181870152808701519150507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018286015261118c8382610ecf565b9695505050505050565b6000602082840312156111a857600080fd5b8151610f2c81610b36565b82815260406020820152600061059d6040830184610ecf565b600073ffffffffffffffffffffffffffffffffffffffff808a16835280891660208401528088166040840152508560608301528460808301528360a083015260e060c083015261121f60e0830184610ecf565b9998505050505050505050565b60006020828403121561123e57600080fd5b815167ffffffffffffffff81111561125557600080fd5b8201601f8101841361126657600080fd5b8051611274610d8a82610d25565b81815285602083850101111561128957600080fd5b61129a826020830160208601610eab565b95945050505050565b6000602082840312156112b557600080fd5b5051919050565b6000602082840312156112ce57600080fd5b81518015158114610f2c57600080fd5b600082516112f0818460208701610eab565b919091019291505056fea164736f6c6343000813000a",
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

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) GASPRICEBID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "GAS_PRICE_BID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) GASPRICEBID() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GASPRICEBID(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) GASPRICEBID() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GASPRICEBID(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) MAXGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "MAX_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) MAXGAS() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.MAXGAS(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) MAXGAS() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.MAXGAS(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) MAXSUBMISSIONCOST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "MAX_SUBMISSION_COST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) MAXSUBMISSIONCOST() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.MAXSUBMISSIONCOST(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) MAXSUBMISSIONCOST() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.MAXSUBMISSIONCOST(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) ExposeForEncoding(opts *bind.CallOpts, payload ArbitrumL1BridgeAdapterArbitrumFinalizationPayload) ([]byte, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "exposeForEncoding", payload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) ExposeForEncoding(payload ArbitrumL1BridgeAdapterArbitrumFinalizationPayload) ([]byte, error) {
	return _ArbitrumL1BridgeAdapter.Contract.ExposeForEncoding(&_ArbitrumL1BridgeAdapter.CallOpts, payload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) ExposeForEncoding(payload ArbitrumL1BridgeAdapterArbitrumFinalizationPayload) ([]byte, error) {
	return _ArbitrumL1BridgeAdapter.Contract.ExposeForEncoding(&_ArbitrumL1BridgeAdapter.CallOpts, payload)
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

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactor) SendERC20(opts *bind.TransactOpts, localToken common.Address, arg1 common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.contract.Transact(opts, "sendERC20", localToken, arg1, recipient, amount)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) SendERC20(localToken common.Address, arg1 common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.SendERC20(&_ArbitrumL1BridgeAdapter.TransactOpts, localToken, arg1, recipient, amount)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactorSession) SendERC20(localToken common.Address, arg1 common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.SendERC20(&_ArbitrumL1BridgeAdapter.TransactOpts, localToken, arg1, recipient, amount)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapter) Address() common.Address {
	return _ArbitrumL1BridgeAdapter.address
}

type ArbitrumL1BridgeAdapterInterface interface {
	GASPRICEBID(opts *bind.CallOpts) (*big.Int, error)

	MAXGAS(opts *bind.CallOpts) (*big.Int, error)

	MAXSUBMISSIONCOST(opts *bind.CallOpts) (*big.Int, error)

	ExposeForEncoding(opts *bind.CallOpts, payload ArbitrumL1BridgeAdapterArbitrumFinalizationPayload) ([]byte, error)

	GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error)

	GetL2Token(opts *bind.CallOpts, l1Token common.Address) (common.Address, error)

	FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, localToken common.Address, arg1 common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	Address() common.Address
}
