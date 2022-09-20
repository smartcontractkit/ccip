// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package toll_sender_dapp

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
)

var TollSenderDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractEVM2AnyTollOnRampRouterInterface\",\"name\":\"onRampRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidAddress\",\"type\":\"address\"}],\"name\":\"InvalidDestinationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"i_destinationChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_destinationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_onRampRouter\",\"outputs\":[{\"internalType\":\"contractEVM2AnyTollOnRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"sendTokens\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b50604051610c7d380380610c7d83398101604081905261002f91610064565b6001600160a01b0392831660805260a0919091521660c0526100a7565b6001600160a01b038116811461006157600080fd5b50565b60008060006060848603121561007957600080fd5b83516100848161004c565b60208501516040860151919450925061009c8161004c565b809150509250925092565b60805160a05160c051610b8c6100f16000396000818160c3015261037b015260008181610129015261034f0152600081816101020152818161025f01526103200152610b8c6000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c806389f9ad2a1161005057806389f9ad2a146100fd578063a721719514610124578063e83f967b1461015957600080fd5b8063181f5a771461006c5780635c1b583a146100be575b600080fd5b6100a86040518060400160405280601481526020017f546f6c6c53656e6465724461707020312e302e3000000000000000000000000081525081565b6040516100b591906107eb565b60405180910390f35b6100e57f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100b5565b6100e57f000000000000000000000000000000000000000000000000000000000000000081565b61014b7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100b5565b61016c6101673660046108ec565b610185565b60405167ffffffffffffffff90911681526020016100b5565b60006001600160a01b0384166101d7576040517ffdc6604f0000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024015b60405180910390fd5b60005b835181101561031d5761023533308584815181106101fa576101fa6109c2565b6020026020010151878581518110610214576102146109c2565b60200260200101516001600160a01b031661048d909392919063ffffffff16565b838181518110610247576102476109c2565b60200260200101516001600160a01b031663095ea7b37f0000000000000000000000000000000000000000000000000000000000000000858481518110610290576102906109c2565b60200260200101516040518363ffffffff1660e01b81526004016102c99291906001600160a01b03929092168252602082015260400190565b6020604051808303816000875af11580156102e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030c91906109d8565b50610316816109fa565b90506101da565b507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e7c62c8c7f00000000000000000000000000000000000000000000000000000000000000006040518060e001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200133896040516020016103d09291906001600160a01b0392831681529116602082015260400190565b604051602081830303815290604052815260200187815260200186815260200187600081518110610403576104036109c2565b60200260200101516001600160a01b031681526020016000815260200160008152506040518363ffffffff1660e01b8152600401610442929190610a5c565b6020604051808303816000875af1158015610461573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104859190610b39565b949350505050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261051590859061051b565b50505050565b6000610570826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166106059092919063ffffffff16565b805190915015610600578080602001905181019061058e91906109d8565b6106005760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101ce565b505050565b6060610485848460008561061b565b9392505050565b6060824710156106935760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101ce565b843b6106e15760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101ce565b600080866001600160a01b031685876040516106fd9190610b63565b60006040518083038185875af1925050503d806000811461073a576040519150601f19603f3d011682016040523d82523d6000602084013e61073f565b606091505b509150915061074f82828661075a565b979650505050505050565b60608315610769575081610614565b8251156107795782518084602001fd5b8160405162461bcd60e51b81526004016101ce91906107eb565b60005b838110156107ae578181015183820152602001610796565b838111156105155750506000910152565b600081518084526107d7816020860160208601610793565b601f01601f19169290920160200192915050565b60208152600061061460208301846107bf565b6001600160a01b038116811461081357600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561085557610855610816565b604052919050565b600067ffffffffffffffff82111561087757610877610816565b5060051b60200190565b600082601f83011261089257600080fd5b813560206108a76108a28361085d565b61082c565b82815260059290921b840181019181810190868411156108c657600080fd5b8286015b848110156108e157803583529183019183016108ca565b509695505050505050565b60008060006060848603121561090157600080fd5b833561090c816107fe565b925060208481013567ffffffffffffffff8082111561092a57600080fd5b818701915087601f83011261093e57600080fd5b813561094c6108a28261085d565b81815260059190911b8301840190848101908a83111561096b57600080fd5b938501935b82851015610992578435610983816107fe565b82529385019390850190610970565b9650505060408701359250808311156109aa57600080fd5b50506109b886828701610881565b9150509250925092565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156109ea57600080fd5b8151801515811461061457600080fd5b600060018201610a1a57634e487b7160e01b600052601160045260246000fd5b5060010190565b600081518084526020808501945080840160005b83811015610a5157815187529582019590820190600101610a35565b509495945050505050565b828152600060206040818401526001600160a01b038085511660408501528185015160e06060860152610a936101208601826107bf565b6040870151603f1987830381016080890152815180845291860193506000929091908601905b80841015610adb57845186168252938601936001939093019290860190610ab9565b5060608901519550818882030160a0890152610af78187610a21565b955050505050506080840151610b1860c08501826001600160a01b03169052565b5060a084015160e084015260c0840151610100840152809150509392505050565b600060208284031215610b4b57600080fd5b815167ffffffffffffffff8116811461061457600080fd5b60008251610b75818460208701610793565b919091019291505056fea164736f6c634300080f000a",
}

var TollSenderDappABI = TollSenderDappMetaData.ABI

var TollSenderDappBin = TollSenderDappMetaData.Bin

func DeployTollSenderDapp(auth *bind.TransactOpts, backend bind.ContractBackend, onRampRouter common.Address, destinationChainId *big.Int, destinationContract common.Address) (common.Address, *types.Transaction, *TollSenderDapp, error) {
	parsed, err := TollSenderDappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TollSenderDappBin), backend, onRampRouter, destinationChainId, destinationContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TollSenderDapp{TollSenderDappCaller: TollSenderDappCaller{contract: contract}, TollSenderDappTransactor: TollSenderDappTransactor{contract: contract}, TollSenderDappFilterer: TollSenderDappFilterer{contract: contract}}, nil
}

type TollSenderDapp struct {
	address common.Address
	abi     abi.ABI
	TollSenderDappCaller
	TollSenderDappTransactor
	TollSenderDappFilterer
}

type TollSenderDappCaller struct {
	contract *bind.BoundContract
}

type TollSenderDappTransactor struct {
	contract *bind.BoundContract
}

type TollSenderDappFilterer struct {
	contract *bind.BoundContract
}

type TollSenderDappSession struct {
	Contract     *TollSenderDapp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TollSenderDappCallerSession struct {
	Contract *TollSenderDappCaller
	CallOpts bind.CallOpts
}

type TollSenderDappTransactorSession struct {
	Contract     *TollSenderDappTransactor
	TransactOpts bind.TransactOpts
}

type TollSenderDappRaw struct {
	Contract *TollSenderDapp
}

type TollSenderDappCallerRaw struct {
	Contract *TollSenderDappCaller
}

type TollSenderDappTransactorRaw struct {
	Contract *TollSenderDappTransactor
}

func NewTollSenderDapp(address common.Address, backend bind.ContractBackend) (*TollSenderDapp, error) {
	abi, err := abi.JSON(strings.NewReader(TollSenderDappABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindTollSenderDapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TollSenderDapp{address: address, abi: abi, TollSenderDappCaller: TollSenderDappCaller{contract: contract}, TollSenderDappTransactor: TollSenderDappTransactor{contract: contract}, TollSenderDappFilterer: TollSenderDappFilterer{contract: contract}}, nil
}

func NewTollSenderDappCaller(address common.Address, caller bind.ContractCaller) (*TollSenderDappCaller, error) {
	contract, err := bindTollSenderDapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TollSenderDappCaller{contract: contract}, nil
}

func NewTollSenderDappTransactor(address common.Address, transactor bind.ContractTransactor) (*TollSenderDappTransactor, error) {
	contract, err := bindTollSenderDapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TollSenderDappTransactor{contract: contract}, nil
}

func NewTollSenderDappFilterer(address common.Address, filterer bind.ContractFilterer) (*TollSenderDappFilterer, error) {
	contract, err := bindTollSenderDapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TollSenderDappFilterer{contract: contract}, nil
}

func bindTollSenderDapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TollSenderDappABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_TollSenderDapp *TollSenderDappRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TollSenderDapp.Contract.TollSenderDappCaller.contract.Call(opts, result, method, params...)
}

func (_TollSenderDapp *TollSenderDappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.TollSenderDappTransactor.contract.Transfer(opts)
}

func (_TollSenderDapp *TollSenderDappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.TollSenderDappTransactor.contract.Transact(opts, method, params...)
}

func (_TollSenderDapp *TollSenderDappCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TollSenderDapp.Contract.contract.Call(opts, result, method, params...)
}

func (_TollSenderDapp *TollSenderDappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.contract.Transfer(opts)
}

func (_TollSenderDapp *TollSenderDappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.contract.Transact(opts, method, params...)
}

func (_TollSenderDapp *TollSenderDappCaller) IDestinationChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TollSenderDapp.contract.Call(opts, &out, "i_destinationChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TollSenderDapp *TollSenderDappSession) IDestinationChainId() (*big.Int, error) {
	return _TollSenderDapp.Contract.IDestinationChainId(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCallerSession) IDestinationChainId() (*big.Int, error) {
	return _TollSenderDapp.Contract.IDestinationChainId(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCaller) IDestinationContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TollSenderDapp.contract.Call(opts, &out, "i_destinationContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TollSenderDapp *TollSenderDappSession) IDestinationContract() (common.Address, error) {
	return _TollSenderDapp.Contract.IDestinationContract(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCallerSession) IDestinationContract() (common.Address, error) {
	return _TollSenderDapp.Contract.IDestinationContract(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCaller) IOnRampRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TollSenderDapp.contract.Call(opts, &out, "i_onRampRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TollSenderDapp *TollSenderDappSession) IOnRampRouter() (common.Address, error) {
	return _TollSenderDapp.Contract.IOnRampRouter(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCallerSession) IOnRampRouter() (common.Address, error) {
	return _TollSenderDapp.Contract.IOnRampRouter(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TollSenderDapp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_TollSenderDapp *TollSenderDappSession) TypeAndVersion() (string, error) {
	return _TollSenderDapp.Contract.TypeAndVersion(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCallerSession) TypeAndVersion() (string, error) {
	return _TollSenderDapp.Contract.TypeAndVersion(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappTransactor) SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TollSenderDapp.contract.Transact(opts, "sendTokens", destinationAddress, tokens, amounts)
}

func (_TollSenderDapp *TollSenderDappSession) SendTokens(destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.SendTokens(&_TollSenderDapp.TransactOpts, destinationAddress, tokens, amounts)
}

func (_TollSenderDapp *TollSenderDappTransactorSession) SendTokens(destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.SendTokens(&_TollSenderDapp.TransactOpts, destinationAddress, tokens, amounts)
}

func (_TollSenderDapp *TollSenderDapp) Address() common.Address {
	return _TollSenderDapp.address
}

type TollSenderDappInterface interface {
	IDestinationChainId(opts *bind.CallOpts) (*big.Int, error)

	IDestinationContract(opts *bind.CallOpts) (common.Address, error)

	IOnRampRouter(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error)

	Address() common.Address
}
