// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_l1_bridge_adapter

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

type MockL1BridgeAdapterFinalizePayload struct {
	Nonce  *big.Int
	Amount *big.Int
}

type MockL1BridgeAdapterPayload struct {
	Action uint8
	Data   []byte
}

type MockL1BridgeAdapterProvePayload struct {
	Nonce *big.Int
}

var MockL1BridgeAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFinalizationAction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MsgShouldNotContainValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"NonceAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"NonceNotProven\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structMockL1BridgeAdapter.FinalizePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"encodeFinalizePayload\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumMockL1BridgeAdapter.FinalizationAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structMockL1BridgeAdapter.Payload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"encodePayload\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structMockL1BridgeAdapter.ProvePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"encodeProvePayload\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"localReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"bridgeSpecificPayload\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeFeeInNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"provideLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600160005534801561001557600080fd5b5060405161107e38038061107e83398101604081905261003491610045565b6001600160a01b0316608052610075565b60006020828403121561005757600080fd5b81516001600160a01b038116811461006e57600080fd5b9392505050565b608051610fd96100a5600039600081816101b701528181610289015281816104f3015261067a0152610fd96000f3fe60806040526004361061007b5760003560e01c8063a71d98b71161004e578063a71d98b714610111578063aee0c38814610131578063eb521a4c1461014c578063f19e1eb61461016c57600080fd5b80630a861f2a146100805780632e4b1fc9146100a2578063331e5ff0146100c357806338314bb2146100e1575b600080fd5b34801561008c57600080fd5b506100a061009b366004610a8c565b610187565b005b3480156100ae57600080fd5b50604051600081526020015b60405180910390f35b3480156100cf57600080fd5b506100a06100de366004610b6f565b50565b3480156100ed57600080fd5b506101016100fc366004610cd3565b6102e0565b60405190151581526020016100ba565b61012461011f366004610d34565b610560565b6040516100ba9190610e21565b34801561013d57600080fd5b506100a06100de366004610e3b565b34801561015857600080fd5b506100a0610167366004610a8c565b610660565b34801561017857600080fd5b506100a06100de366004610e6d565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015281907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610213573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102379190610e91565b101561026f576040517fbb55fd2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102b073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633836106d2565b604051819033907fc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf984017171990600090a350565b6000806102ef83850185610b6f565b905060008151600181111561030657610306610eaa565b036103c157600081602001518060200190518101906103259190610ed9565b805160009081526001602052604090205490915060ff161561037e5780516040517f91cab50400000000000000000000000000000000000000000000000000000000815260048101919091526024015b60405180910390fd5b516000908152600160208190526040822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909117905591506105589050565b6001815160018111156103d6576103d6610eaa565b0361052657600081602001518060200190518101906103f59190610efd565b805160009081526001602052604090205490915060ff166104485780516040517f974f61110000000000000000000000000000000000000000000000000000000081526004810191909152602401610375565b805160009081526002602052604090205460ff16156104995780516040517f91cab5040000000000000000000000000000000000000000000000000000000081526004810191909152602401610375565b8051600090815260026020908152604090912080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905581015161051b9073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169088906106d2565b600192505050610558565b6040517fee2ef09800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b949350505050565b6040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810184905260609073ffffffffffffffffffffffffffffffffffffffff8816906323b872dd906064016020604051808303816000875af11580156105dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106009190610f2f565b5060008054818061061083610f51565b9190505560405160200161062691815260200190565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815291905298975050505050505050565b6106a273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163330846107ab565b604051819033907fc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb31208890600090a350565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526107a69084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261080f565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526108099085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401610724565b50505050565b6000610871826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661091b9092919063ffffffff16565b8051909150156107a6578080602001905181019061088f9190610f2f565b6107a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610375565b60606105588484600085856000808673ffffffffffffffffffffffffffffffffffffffff16858760405161094f9190610fb0565b60006040518083038185875af1925050503d806000811461098c576040519150601f19603f3d011682016040523d82523d6000602084013e610991565b606091505b50915091506109a2878383876109ad565b979650505050505050565b60608315610a43578251600003610a3c5773ffffffffffffffffffffffffffffffffffffffff85163b610a3c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610375565b5081610558565b6105588383815115610a585781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103759190610e21565b600060208284031215610a9e57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610af757610af7610aa5565b60405290565b6040516020810167ffffffffffffffff81118282101715610af757610af7610aa5565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610b6757610b67610aa5565b604052919050565b60006020808385031215610b8257600080fd5b823567ffffffffffffffff80821115610b9a57600080fd5b9084019060408287031215610bae57600080fd5b610bb6610ad4565b823560028110610bc557600080fd5b81528284013582811115610bd857600080fd5b80840193505086601f840112610bed57600080fd5b823582811115610bff57610bff610aa5565b610c2f857fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610b20565b92508083528785828601011115610c4557600080fd5b8085850186850137600090830185015292830152509392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c8557600080fd5b919050565b60008083601f840112610c9c57600080fd5b50813567ffffffffffffffff811115610cb457600080fd5b602083019150836020828501011115610ccc57600080fd5b9250929050565b60008060008060608587031215610ce957600080fd5b610cf285610c61565b9350610d0060208601610c61565b9250604085013567ffffffffffffffff811115610d1c57600080fd5b610d2887828801610c8a565b95989497509550505050565b60008060008060008060a08789031215610d4d57600080fd5b610d5687610c61565b9550610d6460208801610c61565b9450610d7260408801610c61565b935060608701359250608087013567ffffffffffffffff811115610d9557600080fd5b610da189828a01610c8a565b979a9699509497509295939492505050565b60005b83811015610dce578181015183820152602001610db6565b50506000910152565b60008151808452610def816020860160208601610db3565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610e346020830184610dd7565b9392505050565b600060408284031215610e4d57600080fd5b610e55610ad4565b82358152602083013560208201528091505092915050565b600060208284031215610e7f57600080fd5b610e87610afd565b9135825250919050565b600060208284031215610ea357600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600060208284031215610eeb57600080fd5b610ef3610afd565b9151825250919050565b600060408284031215610f0f57600080fd5b610f17610ad4565b82518152602083015160208201528091505092915050565b600060208284031215610f4157600080fd5b81518015158114610e3457600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610fa9577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b60008251610fc2818460208701610db3565b919091019291505056fea164736f6c6343000813000a",
}

var MockL1BridgeAdapterABI = MockL1BridgeAdapterMetaData.ABI

var MockL1BridgeAdapterBin = MockL1BridgeAdapterMetaData.Bin

func DeployMockL1BridgeAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *MockL1BridgeAdapter, error) {
	parsed, err := MockL1BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockL1BridgeAdapterBin), backend, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockL1BridgeAdapter{address: address, abi: *parsed, MockL1BridgeAdapterCaller: MockL1BridgeAdapterCaller{contract: contract}, MockL1BridgeAdapterTransactor: MockL1BridgeAdapterTransactor{contract: contract}, MockL1BridgeAdapterFilterer: MockL1BridgeAdapterFilterer{contract: contract}}, nil
}

type MockL1BridgeAdapter struct {
	address common.Address
	abi     abi.ABI
	MockL1BridgeAdapterCaller
	MockL1BridgeAdapterTransactor
	MockL1BridgeAdapterFilterer
}

type MockL1BridgeAdapterCaller struct {
	contract *bind.BoundContract
}

type MockL1BridgeAdapterTransactor struct {
	contract *bind.BoundContract
}

type MockL1BridgeAdapterFilterer struct {
	contract *bind.BoundContract
}

type MockL1BridgeAdapterSession struct {
	Contract     *MockL1BridgeAdapter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockL1BridgeAdapterCallerSession struct {
	Contract *MockL1BridgeAdapterCaller
	CallOpts bind.CallOpts
}

type MockL1BridgeAdapterTransactorSession struct {
	Contract     *MockL1BridgeAdapterTransactor
	TransactOpts bind.TransactOpts
}

type MockL1BridgeAdapterRaw struct {
	Contract *MockL1BridgeAdapter
}

type MockL1BridgeAdapterCallerRaw struct {
	Contract *MockL1BridgeAdapterCaller
}

type MockL1BridgeAdapterTransactorRaw struct {
	Contract *MockL1BridgeAdapterTransactor
}

func NewMockL1BridgeAdapter(address common.Address, backend bind.ContractBackend) (*MockL1BridgeAdapter, error) {
	abi, err := abi.JSON(strings.NewReader(MockL1BridgeAdapterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockL1BridgeAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockL1BridgeAdapter{address: address, abi: abi, MockL1BridgeAdapterCaller: MockL1BridgeAdapterCaller{contract: contract}, MockL1BridgeAdapterTransactor: MockL1BridgeAdapterTransactor{contract: contract}, MockL1BridgeAdapterFilterer: MockL1BridgeAdapterFilterer{contract: contract}}, nil
}

func NewMockL1BridgeAdapterCaller(address common.Address, caller bind.ContractCaller) (*MockL1BridgeAdapterCaller, error) {
	contract, err := bindMockL1BridgeAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockL1BridgeAdapterCaller{contract: contract}, nil
}

func NewMockL1BridgeAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*MockL1BridgeAdapterTransactor, error) {
	contract, err := bindMockL1BridgeAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockL1BridgeAdapterTransactor{contract: contract}, nil
}

func NewMockL1BridgeAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*MockL1BridgeAdapterFilterer, error) {
	contract, err := bindMockL1BridgeAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockL1BridgeAdapterFilterer{contract: contract}, nil
}

func bindMockL1BridgeAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockL1BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL1BridgeAdapter.Contract.MockL1BridgeAdapterCaller.contract.Call(opts, result, method, params...)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.MockL1BridgeAdapterTransactor.contract.Transfer(opts)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.MockL1BridgeAdapterTransactor.contract.Transact(opts, method, params...)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL1BridgeAdapter.Contract.contract.Call(opts, result, method, params...)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.contract.Transfer(opts)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.contract.Transact(opts, method, params...)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterCaller) EncodeFinalizePayload(opts *bind.CallOpts, payload MockL1BridgeAdapterFinalizePayload) error {
	var out []interface{}
	err := _MockL1BridgeAdapter.contract.Call(opts, &out, "encodeFinalizePayload", payload)

	if err != nil {
		return err
	}

	return err

}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) EncodeFinalizePayload(payload MockL1BridgeAdapterFinalizePayload) error {
	return _MockL1BridgeAdapter.Contract.EncodeFinalizePayload(&_MockL1BridgeAdapter.CallOpts, payload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterCallerSession) EncodeFinalizePayload(payload MockL1BridgeAdapterFinalizePayload) error {
	return _MockL1BridgeAdapter.Contract.EncodeFinalizePayload(&_MockL1BridgeAdapter.CallOpts, payload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterCaller) EncodePayload(opts *bind.CallOpts, payload MockL1BridgeAdapterPayload) error {
	var out []interface{}
	err := _MockL1BridgeAdapter.contract.Call(opts, &out, "encodePayload", payload)

	if err != nil {
		return err
	}

	return err

}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) EncodePayload(payload MockL1BridgeAdapterPayload) error {
	return _MockL1BridgeAdapter.Contract.EncodePayload(&_MockL1BridgeAdapter.CallOpts, payload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterCallerSession) EncodePayload(payload MockL1BridgeAdapterPayload) error {
	return _MockL1BridgeAdapter.Contract.EncodePayload(&_MockL1BridgeAdapter.CallOpts, payload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterCaller) EncodeProvePayload(opts *bind.CallOpts, payload MockL1BridgeAdapterProvePayload) error {
	var out []interface{}
	err := _MockL1BridgeAdapter.contract.Call(opts, &out, "encodeProvePayload", payload)

	if err != nil {
		return err
	}

	return err

}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) EncodeProvePayload(payload MockL1BridgeAdapterProvePayload) error {
	return _MockL1BridgeAdapter.Contract.EncodeProvePayload(&_MockL1BridgeAdapter.CallOpts, payload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterCallerSession) EncodeProvePayload(payload MockL1BridgeAdapterProvePayload) error {
	return _MockL1BridgeAdapter.Contract.EncodeProvePayload(&_MockL1BridgeAdapter.CallOpts, payload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterCaller) GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockL1BridgeAdapter.contract.Call(opts, &out, "getBridgeFeeInNative")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _MockL1BridgeAdapter.Contract.GetBridgeFeeInNative(&_MockL1BridgeAdapter.CallOpts)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterCallerSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _MockL1BridgeAdapter.Contract.GetBridgeFeeInNative(&_MockL1BridgeAdapter.CallOpts)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, localReceiver common.Address, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.contract.Transact(opts, "finalizeWithdrawERC20", arg0, localReceiver, bridgeSpecificPayload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) FinalizeWithdrawERC20(arg0 common.Address, localReceiver common.Address, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.FinalizeWithdrawERC20(&_MockL1BridgeAdapter.TransactOpts, arg0, localReceiver, bridgeSpecificPayload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorSession) FinalizeWithdrawERC20(arg0 common.Address, localReceiver common.Address, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.FinalizeWithdrawERC20(&_MockL1BridgeAdapter.TransactOpts, arg0, localReceiver, bridgeSpecificPayload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactor) ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.contract.Transact(opts, "provideLiquidity", amount)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.ProvideLiquidity(&_MockL1BridgeAdapter.TransactOpts, amount)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.ProvideLiquidity(&_MockL1BridgeAdapter.TransactOpts, amount)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactor) SendERC20(opts *bind.TransactOpts, localToken common.Address, arg1 common.Address, arg2 common.Address, amount *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.contract.Transact(opts, "sendERC20", localToken, arg1, arg2, amount, arg4)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) SendERC20(localToken common.Address, arg1 common.Address, arg2 common.Address, amount *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.SendERC20(&_MockL1BridgeAdapter.TransactOpts, localToken, arg1, arg2, amount, arg4)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorSession) SendERC20(localToken common.Address, arg1 common.Address, arg2 common.Address, amount *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.SendERC20(&_MockL1BridgeAdapter.TransactOpts, localToken, arg1, arg2, amount, arg4)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactor) WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.contract.Transact(opts, "withdrawLiquidity", amount)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.WithdrawLiquidity(&_MockL1BridgeAdapter.TransactOpts, amount)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.WithdrawLiquidity(&_MockL1BridgeAdapter.TransactOpts, amount)
}

type MockL1BridgeAdapterLiquidityAddedIterator struct {
	Event *MockL1BridgeAdapterLiquidityAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockL1BridgeAdapterLiquidityAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL1BridgeAdapterLiquidityAdded)
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
		it.Event = new(MockL1BridgeAdapterLiquidityAdded)
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

func (it *MockL1BridgeAdapterLiquidityAddedIterator) Error() error {
	return it.fail
}

func (it *MockL1BridgeAdapterLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockL1BridgeAdapterLiquidityAdded struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*MockL1BridgeAdapterLiquidityAddedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _MockL1BridgeAdapter.contract.FilterLogs(opts, "LiquidityAdded", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &MockL1BridgeAdapterLiquidityAddedIterator{contract: _MockL1BridgeAdapter.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *MockL1BridgeAdapterLiquidityAdded, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _MockL1BridgeAdapter.contract.WatchLogs(opts, "LiquidityAdded", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockL1BridgeAdapterLiquidityAdded)
				if err := _MockL1BridgeAdapter.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

func (_MockL1BridgeAdapter *MockL1BridgeAdapterFilterer) ParseLiquidityAdded(log types.Log) (*MockL1BridgeAdapterLiquidityAdded, error) {
	event := new(MockL1BridgeAdapterLiquidityAdded)
	if err := _MockL1BridgeAdapter.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockL1BridgeAdapterLiquidityRemovedIterator struct {
	Event *MockL1BridgeAdapterLiquidityRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockL1BridgeAdapterLiquidityRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL1BridgeAdapterLiquidityRemoved)
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
		it.Event = new(MockL1BridgeAdapterLiquidityRemoved)
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

func (it *MockL1BridgeAdapterLiquidityRemovedIterator) Error() error {
	return it.fail
}

func (it *MockL1BridgeAdapterLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockL1BridgeAdapterLiquidityRemoved struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*MockL1BridgeAdapterLiquidityRemovedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _MockL1BridgeAdapter.contract.FilterLogs(opts, "LiquidityRemoved", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &MockL1BridgeAdapterLiquidityRemovedIterator{contract: _MockL1BridgeAdapter.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *MockL1BridgeAdapterLiquidityRemoved, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _MockL1BridgeAdapter.contract.WatchLogs(opts, "LiquidityRemoved", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockL1BridgeAdapterLiquidityRemoved)
				if err := _MockL1BridgeAdapter.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

func (_MockL1BridgeAdapter *MockL1BridgeAdapterFilterer) ParseLiquidityRemoved(log types.Log) (*MockL1BridgeAdapterLiquidityRemoved, error) {
	event := new(MockL1BridgeAdapterLiquidityRemoved)
	if err := _MockL1BridgeAdapter.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MockL1BridgeAdapter.abi.Events["LiquidityAdded"].ID:
		return _MockL1BridgeAdapter.ParseLiquidityAdded(log)
	case _MockL1BridgeAdapter.abi.Events["LiquidityRemoved"].ID:
		return _MockL1BridgeAdapter.ParseLiquidityRemoved(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MockL1BridgeAdapterLiquidityAdded) Topic() common.Hash {
	return common.HexToHash("0xc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088")
}

func (MockL1BridgeAdapterLiquidityRemoved) Topic() common.Hash {
	return common.HexToHash("0xc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719")
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapter) Address() common.Address {
	return _MockL1BridgeAdapter.address
}

type MockL1BridgeAdapterInterface interface {
	EncodeFinalizePayload(opts *bind.CallOpts, payload MockL1BridgeAdapterFinalizePayload) error

	EncodePayload(opts *bind.CallOpts, payload MockL1BridgeAdapterPayload) error

	EncodeProvePayload(opts *bind.CallOpts, payload MockL1BridgeAdapterProvePayload) error

	GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error)

	FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, localReceiver common.Address, bridgeSpecificPayload []byte) (*types.Transaction, error)

	ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, localToken common.Address, arg1 common.Address, arg2 common.Address, amount *big.Int, arg4 []byte) (*types.Transaction, error)

	WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*MockL1BridgeAdapterLiquidityAddedIterator, error)

	WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *MockL1BridgeAdapterLiquidityAdded, provider []common.Address, amount []*big.Int) (event.Subscription, error)

	ParseLiquidityAdded(log types.Log) (*MockL1BridgeAdapterLiquidityAdded, error)

	FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*MockL1BridgeAdapterLiquidityRemovedIterator, error)

	WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *MockL1BridgeAdapterLiquidityRemoved, provider []common.Address, amount []*big.Int) (event.Subscription, error)

	ParseLiquidityRemoved(log types.Log) (*MockL1BridgeAdapterLiquidityRemoved, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
