// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package single_token_sender

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

var EOASingleTokenSenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractSingleTokenOnRamp\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DESTINATION_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ON_RAMP\",\"outputs\":[{\"internalType\":\"contractSingleTokenOnRamp\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rampDetails\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destinationChainToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"sendTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b50604051610fef380380610fef83398101604081905261002f9161004d565b6001600160601b0319606092831b8116608052911b1660a05261009f565b6000806040838503121561006057600080fd5b825161006b81610087565b602084015190925061007c81610087565b809150509250929050565b6001600160a01b038116811461009c57600080fd5b50565b60805160601c60a05160601c610ef56100fa60003960008181607101526103440152600081816101010152818161021701528181610474015281816105f30152818161068e0152818161072e01526107ce0152610ef56000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80632c3b6d23116100505780632c3b6d23146100fc5780634217e2871461012357806379040f451461014457600080fd5b80630ab7dea91461006c578063181f5a77146100bd575b600080fd5b6100937f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b604080518082018252601a81527f454f4153696e676c65546f6b656e53656e64657220312e302e30000000000000602082015290516100b49190610d44565b6100937f000000000000000000000000000000000000000000000000000000000000000081565b610136610131366004610c09565b610183565b6040519081526020016100b4565b61014c610687565b6040805173ffffffffffffffffffffffffffffffffffffffff948516815260208101939093529216918101919091526060016100b4565b600073ffffffffffffffffffffffffffffffffffffffff84166101ef576040517ffdc6604f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024015b60405180910390fd5b60408051600180825281830190925260609160009190602080830190803683370190505090507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166382bfefc86040518163ffffffff1660e01b815260040160206040518083038186803b15801561027b57600080fd5b505afa15801561028f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102b39190610c6d565b816000815181106102c6576102c6610e94565b73ffffffffffffffffffffffffffffffffffffffff9290921660209283029190910190910152604080516001808252818301909252600091816020016020820280368337019050509050858160008151811061032457610324610e94565b602002602001018181525050600033905060006040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168152602001838b6040516020016103b392919073ffffffffffffffffffffffffffffffffffffffff92831681529116602082015260400190565b60405160208183030381529060405281526020018581526020018481526020018873ffffffffffffffffffffffffffffffffffffffff16815260200186815250905061043c82308a8760008151811061040e5761040e610e94565b602002602001015173ffffffffffffffffffffffffffffffffffffffff16610871909392919063ffffffff16565b8360008151811061044f5761044f610e94565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16637535d2466040518163ffffffff1660e01b815260040160206040518083038186803b1580156104d857600080fd5b505afa1580156104ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105109190610c6d565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602481018b9052604401602060405180830381600087803b15801561057d57600080fd5b505af1158015610591573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b59190610c4b565b506040517f0649d29200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690630649d29290610628908490600401610d57565b602060405180830381600087803b15801561064257600080fd5b505af1158015610656573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061067a9190610c8a565b9998505050505050505050565b60008060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166382bfefc86040518163ffffffff1660e01b815260040160206040518083038186803b1580156106f257600080fd5b505afa158015610706573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072a9190610c6d565b92507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632ea023696040518163ffffffff1660e01b815260040160206040518083038186803b15801561079257600080fd5b505afa1580156107a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ca9190610c8a565b91507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e3a604f86040518163ffffffff1660e01b815260040160206040518083038186803b15801561083257600080fd5b505afa158015610846573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061086a9190610c6d565b9050909192565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261090690859061090c565b50505050565b600061096e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16610a1d9092919063ffffffff16565b805190915015610a18578080602001905181019061098c9190610c4b565b610a18576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101e6565b505050565b6060610a2c8484600085610a36565b90505b9392505050565b606082471015610ac8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101e6565b843b610b30576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101e6565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610b599190610d28565b60006040518083038185875af1925050503d8060008114610b96576040519150601f19603f3d011682016040523d82523d6000602084013e610b9b565b606091505b5091509150610bab828286610bb6565b979650505050505050565b60608315610bc5575081610a2f565b825115610bd55782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e69190610d44565b600080600060608486031215610c1e57600080fd5b8335610c2981610ec3565b9250602084013591506040840135610c4081610ec3565b809150509250925092565b600060208284031215610c5d57600080fd5b81518015158114610a2f57600080fd5b600060208284031215610c7f57600080fd5b8151610a2f81610ec3565b600060208284031215610c9c57600080fd5b5051919050565b600081518084526020808501945080840160005b83811015610cd357815187529582019590820190600101610cb7565b509495945050505050565b60008151808452610cf6816020860160208601610e68565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60008251610d3a818460208701610e68565b9190910192915050565b602081526000610a2f6020830184610cde565b6000602080835273ffffffffffffffffffffffffffffffffffffffff80855116828501528185015160c06040860152610d9360e0860182610cde565b60408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe087830381016060890152815180845291860193506000929091908601905b80841015610df957845186168252938601936001939093019290860190610dd7565b506060890151955081888203016080890152610e158187610ca3565b95505060808801519350610e4160a088018573ffffffffffffffffffffffffffffffffffffffff169052565b60a08801519350808786030160c0880152505050610e5f8282610cde565b95945050505050565b60005b83811015610e83578181015183820152602001610e6b565b838111156109065750506000910152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff81168114610ee557600080fd5b5056fea164736f6c6343000806000a",
}

var EOASingleTokenSenderABI = EOASingleTokenSenderMetaData.ABI

var EOASingleTokenSenderBin = EOASingleTokenSenderMetaData.Bin

func DeployEOASingleTokenSender(auth *bind.TransactOpts, backend bind.ContractBackend, onRamp common.Address, destinationContract common.Address) (common.Address, *types.Transaction, *EOASingleTokenSender, error) {
	parsed, err := EOASingleTokenSenderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EOASingleTokenSenderBin), backend, onRamp, destinationContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EOASingleTokenSender{EOASingleTokenSenderCaller: EOASingleTokenSenderCaller{contract: contract}, EOASingleTokenSenderTransactor: EOASingleTokenSenderTransactor{contract: contract}, EOASingleTokenSenderFilterer: EOASingleTokenSenderFilterer{contract: contract}}, nil
}

type EOASingleTokenSender struct {
	address common.Address
	abi     abi.ABI
	EOASingleTokenSenderCaller
	EOASingleTokenSenderTransactor
	EOASingleTokenSenderFilterer
}

type EOASingleTokenSenderCaller struct {
	contract *bind.BoundContract
}

type EOASingleTokenSenderTransactor struct {
	contract *bind.BoundContract
}

type EOASingleTokenSenderFilterer struct {
	contract *bind.BoundContract
}

type EOASingleTokenSenderSession struct {
	Contract     *EOASingleTokenSender
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EOASingleTokenSenderCallerSession struct {
	Contract *EOASingleTokenSenderCaller
	CallOpts bind.CallOpts
}

type EOASingleTokenSenderTransactorSession struct {
	Contract     *EOASingleTokenSenderTransactor
	TransactOpts bind.TransactOpts
}

type EOASingleTokenSenderRaw struct {
	Contract *EOASingleTokenSender
}

type EOASingleTokenSenderCallerRaw struct {
	Contract *EOASingleTokenSenderCaller
}

type EOASingleTokenSenderTransactorRaw struct {
	Contract *EOASingleTokenSenderTransactor
}

func NewEOASingleTokenSender(address common.Address, backend bind.ContractBackend) (*EOASingleTokenSender, error) {
	abi, err := abi.JSON(strings.NewReader(EOASingleTokenSenderABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEOASingleTokenSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EOASingleTokenSender{address: address, abi: abi, EOASingleTokenSenderCaller: EOASingleTokenSenderCaller{contract: contract}, EOASingleTokenSenderTransactor: EOASingleTokenSenderTransactor{contract: contract}, EOASingleTokenSenderFilterer: EOASingleTokenSenderFilterer{contract: contract}}, nil
}

func NewEOASingleTokenSenderCaller(address common.Address, caller bind.ContractCaller) (*EOASingleTokenSenderCaller, error) {
	contract, err := bindEOASingleTokenSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EOASingleTokenSenderCaller{contract: contract}, nil
}

func NewEOASingleTokenSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*EOASingleTokenSenderTransactor, error) {
	contract, err := bindEOASingleTokenSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EOASingleTokenSenderTransactor{contract: contract}, nil
}

func NewEOASingleTokenSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*EOASingleTokenSenderFilterer, error) {
	contract, err := bindEOASingleTokenSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EOASingleTokenSenderFilterer{contract: contract}, nil
}

func bindEOASingleTokenSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EOASingleTokenSenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EOASingleTokenSender *EOASingleTokenSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EOASingleTokenSender.Contract.EOASingleTokenSenderCaller.contract.Call(opts, result, method, params...)
}

func (_EOASingleTokenSender *EOASingleTokenSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EOASingleTokenSender.Contract.EOASingleTokenSenderTransactor.contract.Transfer(opts)
}

func (_EOASingleTokenSender *EOASingleTokenSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EOASingleTokenSender.Contract.EOASingleTokenSenderTransactor.contract.Transact(opts, method, params...)
}

func (_EOASingleTokenSender *EOASingleTokenSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EOASingleTokenSender.Contract.contract.Call(opts, result, method, params...)
}

func (_EOASingleTokenSender *EOASingleTokenSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EOASingleTokenSender.Contract.contract.Transfer(opts)
}

func (_EOASingleTokenSender *EOASingleTokenSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EOASingleTokenSender.Contract.contract.Transact(opts, method, params...)
}

func (_EOASingleTokenSender *EOASingleTokenSenderCaller) DESTINATIONCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EOASingleTokenSender.contract.Call(opts, &out, "DESTINATION_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EOASingleTokenSender *EOASingleTokenSenderSession) DESTINATIONCONTRACT() (common.Address, error) {
	return _EOASingleTokenSender.Contract.DESTINATIONCONTRACT(&_EOASingleTokenSender.CallOpts)
}

func (_EOASingleTokenSender *EOASingleTokenSenderCallerSession) DESTINATIONCONTRACT() (common.Address, error) {
	return _EOASingleTokenSender.Contract.DESTINATIONCONTRACT(&_EOASingleTokenSender.CallOpts)
}

func (_EOASingleTokenSender *EOASingleTokenSenderCaller) ONRAMP(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EOASingleTokenSender.contract.Call(opts, &out, "ON_RAMP")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EOASingleTokenSender *EOASingleTokenSenderSession) ONRAMP() (common.Address, error) {
	return _EOASingleTokenSender.Contract.ONRAMP(&_EOASingleTokenSender.CallOpts)
}

func (_EOASingleTokenSender *EOASingleTokenSenderCallerSession) ONRAMP() (common.Address, error) {
	return _EOASingleTokenSender.Contract.ONRAMP(&_EOASingleTokenSender.CallOpts)
}

func (_EOASingleTokenSender *EOASingleTokenSenderCaller) RampDetails(opts *bind.CallOpts) (RampDetails,

	error) {
	var out []interface{}
	err := _EOASingleTokenSender.contract.Call(opts, &out, "rampDetails")

	outstruct := new(RampDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.DestinationChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DestinationChainToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_EOASingleTokenSender *EOASingleTokenSenderSession) RampDetails() (RampDetails,

	error) {
	return _EOASingleTokenSender.Contract.RampDetails(&_EOASingleTokenSender.CallOpts)
}

func (_EOASingleTokenSender *EOASingleTokenSenderCallerSession) RampDetails() (RampDetails,

	error) {
	return _EOASingleTokenSender.Contract.RampDetails(&_EOASingleTokenSender.CallOpts)
}

func (_EOASingleTokenSender *EOASingleTokenSenderCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EOASingleTokenSender.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EOASingleTokenSender *EOASingleTokenSenderSession) TypeAndVersion() (string, error) {
	return _EOASingleTokenSender.Contract.TypeAndVersion(&_EOASingleTokenSender.CallOpts)
}

func (_EOASingleTokenSender *EOASingleTokenSenderCallerSession) TypeAndVersion() (string, error) {
	return _EOASingleTokenSender.Contract.TypeAndVersion(&_EOASingleTokenSender.CallOpts)
}

func (_EOASingleTokenSender *EOASingleTokenSenderTransactor) SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, amount *big.Int, executor common.Address) (*types.Transaction, error) {
	return _EOASingleTokenSender.contract.Transact(opts, "sendTokens", destinationAddress, amount, executor)
}

func (_EOASingleTokenSender *EOASingleTokenSenderSession) SendTokens(destinationAddress common.Address, amount *big.Int, executor common.Address) (*types.Transaction, error) {
	return _EOASingleTokenSender.Contract.SendTokens(&_EOASingleTokenSender.TransactOpts, destinationAddress, amount, executor)
}

func (_EOASingleTokenSender *EOASingleTokenSenderTransactorSession) SendTokens(destinationAddress common.Address, amount *big.Int, executor common.Address) (*types.Transaction, error) {
	return _EOASingleTokenSender.Contract.SendTokens(&_EOASingleTokenSender.TransactOpts, destinationAddress, amount, executor)
}

type RampDetails struct {
	Token                 common.Address
	DestinationChainId    *big.Int
	DestinationChainToken common.Address
}

func (_EOASingleTokenSender *EOASingleTokenSender) Address() common.Address {
	return _EOASingleTokenSender.address
}

type EOASingleTokenSenderInterface interface {
	DESTINATIONCONTRACT(opts *bind.CallOpts) (common.Address, error)

	ONRAMP(opts *bind.CallOpts) (common.Address, error)

	RampDetails(opts *bind.CallOpts) (RampDetails,

		error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, amount *big.Int, executor common.Address) (*types.Transaction, error)

	Address() common.Address
}
