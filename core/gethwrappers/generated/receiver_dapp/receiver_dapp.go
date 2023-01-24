// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package receiver_dapp

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

type CommonAny2EVMMessage struct {
	SourceChainId        uint64
	Sender               []byte
	Data                 []byte
	DestTokensAndAmounts []CommonEVMTokenAndAmount
}

type CommonEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

var ReceiverDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"FeeTokenSet\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCommon.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161095d38038061095d83398101604081905261002f916100c6565b8060006001600160a01b038216610060576040516335fdcccd60e21b81526000600482015260240160405180910390fd5b6001600160a01b0382166080526100768161007e565b5050506100f6565b600080546001600160a01b0319166001600160a01b038316908117825560405190917f722ff84c1234b2482061def5c82c6b5080c117b3cbb69d686844a051e4b8e7f391a250565b6000602082840312156100d857600080fd5b81516001600160a01b03811681146100ef57600080fd5b9392505050565b6080516108466101176000396000818160f4015261020801526108466000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80633015b91c116100505780633015b91c146100dd578063b0f479a1146100f2578063ca709a251461013957600080fd5b806301ffc9a71461006c578063181f5a7714610094575b600080fd5b61007f61007a3660046103df565b610157565b60405190151581526020015b60405180910390f35b6100d06040518060400160405280601281526020017f52656365697665724461707020322e302e30000000000000000000000000000081525081565b60405161008b9190610428565b6100f06100eb36600461047d565b6101f0565b005b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161008b565b60005473ffffffffffffffffffffffffffffffffffffffff16610114565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f3015b91c0000000000000000000000000000000000000000000000000000000014806101ea57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610265576040517fd7f7333400000000000000000000000000000000000000000000000000000000815233600482015260240160405180910390fd5b610276610271826106a3565b610279565b50565b6102768160400151826060015160008280602001905181019061029c919061074f565b91505060005b82518110156103d95760008382815181106102bf576102bf610789565b6020026020010151602001519050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561030957508015155b156103c85783828151811061032057610320610789565b6020908102919091010151516040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018490529091169063a9059cbb906044016020604051808303816000875af11580156103a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103c691906107b8565b505b506103d2816107da565b90506102a2565b50505050565b6000602082840312156103f157600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461042157600080fd5b9392505050565b600060208083528351808285015260005b8181101561045557858101830151858201604001528201610439565b81811115610467576000604083870101525b50601f01601f1916929092016040019392505050565b60006020828403121561048f57600080fd5b813567ffffffffffffffff8111156104a657600080fd5b82016080818503121561042157600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561050a5761050a6104b8565b60405290565b6040516080810167ffffffffffffffff8111828210171561050a5761050a6104b8565b604051601f8201601f1916810167ffffffffffffffff8111828210171561055c5761055c6104b8565b604052919050565b600082601f83011261057557600080fd5b813567ffffffffffffffff81111561058f5761058f6104b8565b6105a26020601f19601f84011601610533565b8181528460208386010111156105b757600080fd5b816020850160208301376000918101602001919091529392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461027657600080fd5b600082601f83011261060757600080fd5b8135602067ffffffffffffffff821115610623576106236104b8565b610631818360051b01610533565b82815260069290921b8401810191818101908684111561065057600080fd5b8286015b84811015610698576040818903121561066d5760008081fd5b6106756104e7565b8135610680816105d4565b81528185013585820152835291830191604001610654565b509695505050505050565b6000608082360312156106b557600080fd5b6106bd610510565b823567ffffffffffffffff80821682146106d657600080fd5b908252602084013590808211156106ec57600080fd5b6106f836838701610564565b6020840152604085013591508082111561071157600080fd5b61071d36838701610564565b6040840152606085013591508082111561073657600080fd5b50610743368286016105f6565b60608301525092915050565b6000806040838503121561076257600080fd5b825161076d816105d4565b602084015190925061077e816105d4565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156107ca57600080fd5b8151801515811461042157600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610832577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c634300080f000a",
}

var ReceiverDappABI = ReceiverDappMetaData.ABI

var ReceiverDappBin = ReceiverDappMetaData.Bin

func DeployReceiverDapp(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address) (common.Address, *types.Transaction, *ReceiverDapp, error) {
	parsed, err := ReceiverDappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReceiverDappBin), backend, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReceiverDapp{ReceiverDappCaller: ReceiverDappCaller{contract: contract}, ReceiverDappTransactor: ReceiverDappTransactor{contract: contract}, ReceiverDappFilterer: ReceiverDappFilterer{contract: contract}}, nil
}

type ReceiverDapp struct {
	address common.Address
	abi     abi.ABI
	ReceiverDappCaller
	ReceiverDappTransactor
	ReceiverDappFilterer
}

type ReceiverDappCaller struct {
	contract *bind.BoundContract
}

type ReceiverDappTransactor struct {
	contract *bind.BoundContract
}

type ReceiverDappFilterer struct {
	contract *bind.BoundContract
}

type ReceiverDappSession struct {
	Contract     *ReceiverDapp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ReceiverDappCallerSession struct {
	Contract *ReceiverDappCaller
	CallOpts bind.CallOpts
}

type ReceiverDappTransactorSession struct {
	Contract     *ReceiverDappTransactor
	TransactOpts bind.TransactOpts
}

type ReceiverDappRaw struct {
	Contract *ReceiverDapp
}

type ReceiverDappCallerRaw struct {
	Contract *ReceiverDappCaller
}

type ReceiverDappTransactorRaw struct {
	Contract *ReceiverDappTransactor
}

func NewReceiverDapp(address common.Address, backend bind.ContractBackend) (*ReceiverDapp, error) {
	abi, err := abi.JSON(strings.NewReader(ReceiverDappABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindReceiverDapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiverDapp{address: address, abi: abi, ReceiverDappCaller: ReceiverDappCaller{contract: contract}, ReceiverDappTransactor: ReceiverDappTransactor{contract: contract}, ReceiverDappFilterer: ReceiverDappFilterer{contract: contract}}, nil
}

func NewReceiverDappCaller(address common.Address, caller bind.ContractCaller) (*ReceiverDappCaller, error) {
	contract, err := bindReceiverDapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverDappCaller{contract: contract}, nil
}

func NewReceiverDappTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiverDappTransactor, error) {
	contract, err := bindReceiverDapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverDappTransactor{contract: contract}, nil
}

func NewReceiverDappFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiverDappFilterer, error) {
	contract, err := bindReceiverDapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiverDappFilterer{contract: contract}, nil
}

func bindReceiverDapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReceiverDappABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_ReceiverDapp *ReceiverDappRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiverDapp.Contract.ReceiverDappCaller.contract.Call(opts, result, method, params...)
}

func (_ReceiverDapp *ReceiverDappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.ReceiverDappTransactor.contract.Transfer(opts)
}

func (_ReceiverDapp *ReceiverDappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.ReceiverDappTransactor.contract.Transact(opts, method, params...)
}

func (_ReceiverDapp *ReceiverDappCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiverDapp.Contract.contract.Call(opts, result, method, params...)
}

func (_ReceiverDapp *ReceiverDappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.contract.Transfer(opts)
}

func (_ReceiverDapp *ReceiverDappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.contract.Transact(opts, method, params...)
}

func (_ReceiverDapp *ReceiverDappCaller) GetFeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReceiverDapp.contract.Call(opts, &out, "getFeeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ReceiverDapp *ReceiverDappSession) GetFeeToken() (common.Address, error) {
	return _ReceiverDapp.Contract.GetFeeToken(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappCallerSession) GetFeeToken() (common.Address, error) {
	return _ReceiverDapp.Contract.GetFeeToken(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReceiverDapp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ReceiverDapp *ReceiverDappSession) GetRouter() (common.Address, error) {
	return _ReceiverDapp.Contract.GetRouter(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappCallerSession) GetRouter() (common.Address, error) {
	return _ReceiverDapp.Contract.GetRouter(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ReceiverDapp.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ReceiverDapp *ReceiverDappSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ReceiverDapp.Contract.SupportsInterface(&_ReceiverDapp.CallOpts, interfaceId)
}

func (_ReceiverDapp *ReceiverDappCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ReceiverDapp.Contract.SupportsInterface(&_ReceiverDapp.CallOpts, interfaceId)
}

func (_ReceiverDapp *ReceiverDappCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ReceiverDapp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ReceiverDapp *ReceiverDappSession) TypeAndVersion() (string, error) {
	return _ReceiverDapp.Contract.TypeAndVersion(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappCallerSession) TypeAndVersion() (string, error) {
	return _ReceiverDapp.Contract.TypeAndVersion(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappTransactor) CcipReceive(opts *bind.TransactOpts, message CommonAny2EVMMessage) (*types.Transaction, error) {
	return _ReceiverDapp.contract.Transact(opts, "ccipReceive", message)
}

func (_ReceiverDapp *ReceiverDappSession) CcipReceive(message CommonAny2EVMMessage) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.CcipReceive(&_ReceiverDapp.TransactOpts, message)
}

func (_ReceiverDapp *ReceiverDappTransactorSession) CcipReceive(message CommonAny2EVMMessage) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.CcipReceive(&_ReceiverDapp.TransactOpts, message)
}

type ReceiverDappFeeTokenSetIterator struct {
	Event *ReceiverDappFeeTokenSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ReceiverDappFeeTokenSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiverDappFeeTokenSet)
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
		it.Event = new(ReceiverDappFeeTokenSet)
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

func (it *ReceiverDappFeeTokenSetIterator) Error() error {
	return it.fail
}

func (it *ReceiverDappFeeTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ReceiverDappFeeTokenSet struct {
	FeeToken common.Address
	Raw      types.Log
}

func (_ReceiverDapp *ReceiverDappFilterer) FilterFeeTokenSet(opts *bind.FilterOpts, feeToken []common.Address) (*ReceiverDappFeeTokenSetIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _ReceiverDapp.contract.FilterLogs(opts, "FeeTokenSet", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &ReceiverDappFeeTokenSetIterator{contract: _ReceiverDapp.contract, event: "FeeTokenSet", logs: logs, sub: sub}, nil
}

func (_ReceiverDapp *ReceiverDappFilterer) WatchFeeTokenSet(opts *bind.WatchOpts, sink chan<- *ReceiverDappFeeTokenSet, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _ReceiverDapp.contract.WatchLogs(opts, "FeeTokenSet", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ReceiverDappFeeTokenSet)
				if err := _ReceiverDapp.contract.UnpackLog(event, "FeeTokenSet", log); err != nil {
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

func (_ReceiverDapp *ReceiverDappFilterer) ParseFeeTokenSet(log types.Log) (*ReceiverDappFeeTokenSet, error) {
	event := new(ReceiverDappFeeTokenSet)
	if err := _ReceiverDapp.contract.UnpackLog(event, "FeeTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_ReceiverDapp *ReceiverDapp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _ReceiverDapp.abi.Events["FeeTokenSet"].ID:
		return _ReceiverDapp.ParseFeeTokenSet(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ReceiverDappFeeTokenSet) Topic() common.Hash {
	return common.HexToHash("0x722ff84c1234b2482061def5c82c6b5080c117b3cbb69d686844a051e4b8e7f3")
}

func (_ReceiverDapp *ReceiverDapp) Address() common.Address {
	return _ReceiverDapp.address
}

type ReceiverDappInterface interface {
	GetFeeToken(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	CcipReceive(opts *bind.TransactOpts, message CommonAny2EVMMessage) (*types.Transaction, error)

	FilterFeeTokenSet(opts *bind.FilterOpts, feeToken []common.Address) (*ReceiverDappFeeTokenSetIterator, error)

	WatchFeeTokenSet(opts *bind.WatchOpts, sink chan<- *ReceiverDappFeeTokenSet, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenSet(log types.Log) (*ReceiverDappFeeTokenSet, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
