// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry_module_owner_custom

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
	zkSyncAccounts "github.com/zksync-sdk/zksync2-go/accounts"
	zkSyncClient "github.com/zksync-sdk/zksync2-go/clients"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
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

var RegistryModuleOwnerCustomMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"CanOnlySelfRegister\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"}],\"name\":\"AdministratorRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"registerAdminViaGetCCIPAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"registerAdminViaOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161047e38038061047e83398101604081905261002f91610067565b6001600160a01b03811661005657604051639fabe1c160e01b815260040160405180910390fd5b6001600160a01b0316608052610097565b60006020828403121561007957600080fd5b81516001600160a01b038116811461009057600080fd5b9392505050565b6080516103cc6100b2600039600061024a01526103cc6000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063181f5a771461004657806396ea2f7a14610098578063ff12c354146100ad575b600080fd5b6100826040518060400160405280601f81526020017f52656769737472794d6f64756c654f776e6572437573746f6d20312e352e300081525081565b60405161008f91906102ef565b60405180910390f35b6100ab6100a636600461037e565b6100c0565b005b6100ab6100bb36600461037e565b61013b565b610138818273ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561010f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061013391906103a2565b61018a565b50565b610138818273ffffffffffffffffffffffffffffffffffffffff16638fd6a6ac6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561010f573d6000803e3d6000fd5b73ffffffffffffffffffffffffffffffffffffffff811633146101fd576040517fc454d18200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528316602482015260440160405180910390fd5b6040517fe677ae3700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff838116600483015282811660248301527f0000000000000000000000000000000000000000000000000000000000000000169063e677ae3790604401600060405180830381600087803b15801561028e57600080fd5b505af11580156102a2573d6000803e3d6000fd5b505060405173ffffffffffffffffffffffffffffffffffffffff8085169350851691507f09590fb70af4b833346363965e043a9339e8c7d378b8a2b903c75c277faec4f990600090a35050565b60006020808352835180602085015260005b8181101561031d57858101830151858201604001528201610301565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b73ffffffffffffffffffffffffffffffffffffffff8116811461013857600080fd5b60006020828403121561039057600080fd5b813561039b8161035c565b9392505050565b6000602082840312156103b457600080fd5b815161039b8161035c56fea164736f6c6343000818000a",
}

var RegistryModuleOwnerCustomABI = RegistryModuleOwnerCustomMetaData.ABI

var RegistryModuleOwnerCustomBin = RegistryModuleOwnerCustomMetaData.Bin

func DeployRegistryModuleOwnerCustom(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAdminRegistry common.Address) (common.Address, *CustomTransaction, *RegistryModuleOwnerCustom, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	switch chainId.Uint64() {

	case 324, 280, 300:
		return DeployZkSyncRegistryModuleOwnerCustom(auth, backend, tokenAdminRegistry)
	}

	parsed, err := RegistryModuleOwnerCustomMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegistryModuleOwnerCustomBin), backend, tokenAdminRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, &CustomTransaction{Transaction: tx, CustomHash: tx.Hash()}, &RegistryModuleOwnerCustom{address: address, abi: *parsed, RegistryModuleOwnerCustomCaller: RegistryModuleOwnerCustomCaller{contract: contract}, RegistryModuleOwnerCustomTransactor: RegistryModuleOwnerCustomTransactor{contract: contract}, RegistryModuleOwnerCustomFilterer: RegistryModuleOwnerCustomFilterer{contract: contract}}, nil
}

type RegistryModuleOwnerCustom struct {
	address common.Address
	abi     abi.ABI
	RegistryModuleOwnerCustomCaller
	RegistryModuleOwnerCustomTransactor
	RegistryModuleOwnerCustomFilterer
}

type RegistryModuleOwnerCustomCaller struct {
	contract *bind.BoundContract
}

type RegistryModuleOwnerCustomTransactor struct {
	contract *bind.BoundContract
}

type RegistryModuleOwnerCustomFilterer struct {
	contract *bind.BoundContract
}

type RegistryModuleOwnerCustomSession struct {
	Contract     *RegistryModuleOwnerCustom
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type RegistryModuleOwnerCustomCallerSession struct {
	Contract *RegistryModuleOwnerCustomCaller
	CallOpts bind.CallOpts
}

type RegistryModuleOwnerCustomTransactorSession struct {
	Contract     *RegistryModuleOwnerCustomTransactor
	TransactOpts bind.TransactOpts
}

type RegistryModuleOwnerCustomRaw struct {
	Contract *RegistryModuleOwnerCustom
}

type RegistryModuleOwnerCustomCallerRaw struct {
	Contract *RegistryModuleOwnerCustomCaller
}

type RegistryModuleOwnerCustomTransactorRaw struct {
	Contract *RegistryModuleOwnerCustomTransactor
}

func NewRegistryModuleOwnerCustom(address common.Address, backend bind.ContractBackend) (*RegistryModuleOwnerCustom, error) {
	abi, err := abi.JSON(strings.NewReader(RegistryModuleOwnerCustomABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindRegistryModuleOwnerCustom(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegistryModuleOwnerCustom{address: address, abi: abi, RegistryModuleOwnerCustomCaller: RegistryModuleOwnerCustomCaller{contract: contract}, RegistryModuleOwnerCustomTransactor: RegistryModuleOwnerCustomTransactor{contract: contract}, RegistryModuleOwnerCustomFilterer: RegistryModuleOwnerCustomFilterer{contract: contract}}, nil
}

func NewRegistryModuleOwnerCustomCaller(address common.Address, caller bind.ContractCaller) (*RegistryModuleOwnerCustomCaller, error) {
	contract, err := bindRegistryModuleOwnerCustom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryModuleOwnerCustomCaller{contract: contract}, nil
}

func NewRegistryModuleOwnerCustomTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryModuleOwnerCustomTransactor, error) {
	contract, err := bindRegistryModuleOwnerCustom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryModuleOwnerCustomTransactor{contract: contract}, nil
}

func NewRegistryModuleOwnerCustomFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryModuleOwnerCustomFilterer, error) {
	contract, err := bindRegistryModuleOwnerCustom(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryModuleOwnerCustomFilterer{contract: contract}, nil
}

func bindRegistryModuleOwnerCustom(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryModuleOwnerCustomMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryModuleOwnerCustom.Contract.RegistryModuleOwnerCustomCaller.contract.Call(opts, result, method, params...)
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryModuleOwnerCustom.Contract.RegistryModuleOwnerCustomTransactor.contract.Transfer(opts)
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryModuleOwnerCustom.Contract.RegistryModuleOwnerCustomTransactor.contract.Transact(opts, method, params...)
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryModuleOwnerCustom.Contract.contract.Call(opts, result, method, params...)
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryModuleOwnerCustom.Contract.contract.Transfer(opts)
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryModuleOwnerCustom.Contract.contract.Transact(opts, method, params...)
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RegistryModuleOwnerCustom.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomSession) TypeAndVersion() (string, error) {
	return _RegistryModuleOwnerCustom.Contract.TypeAndVersion(&_RegistryModuleOwnerCustom.CallOpts)
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomCallerSession) TypeAndVersion() (string, error) {
	return _RegistryModuleOwnerCustom.Contract.TypeAndVersion(&_RegistryModuleOwnerCustom.CallOpts)
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomTransactor) RegisterAdminViaGetCCIPAdmin(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _RegistryModuleOwnerCustom.contract.Transact(opts, "registerAdminViaGetCCIPAdmin", token)
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomSession) RegisterAdminViaGetCCIPAdmin(token common.Address) (*types.Transaction, error) {
	return _RegistryModuleOwnerCustom.Contract.RegisterAdminViaGetCCIPAdmin(&_RegistryModuleOwnerCustom.TransactOpts, token)
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomTransactorSession) RegisterAdminViaGetCCIPAdmin(token common.Address) (*types.Transaction, error) {
	return _RegistryModuleOwnerCustom.Contract.RegisterAdminViaGetCCIPAdmin(&_RegistryModuleOwnerCustom.TransactOpts, token)
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomTransactor) RegisterAdminViaOwner(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _RegistryModuleOwnerCustom.contract.Transact(opts, "registerAdminViaOwner", token)
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomSession) RegisterAdminViaOwner(token common.Address) (*types.Transaction, error) {
	return _RegistryModuleOwnerCustom.Contract.RegisterAdminViaOwner(&_RegistryModuleOwnerCustom.TransactOpts, token)
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomTransactorSession) RegisterAdminViaOwner(token common.Address) (*types.Transaction, error) {
	return _RegistryModuleOwnerCustom.Contract.RegisterAdminViaOwner(&_RegistryModuleOwnerCustom.TransactOpts, token)
}

type RegistryModuleOwnerCustomAdministratorRegisteredIterator struct {
	Event *RegistryModuleOwnerCustomAdministratorRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RegistryModuleOwnerCustomAdministratorRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryModuleOwnerCustomAdministratorRegistered)
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
		it.Event = new(RegistryModuleOwnerCustomAdministratorRegistered)
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

func (it *RegistryModuleOwnerCustomAdministratorRegisteredIterator) Error() error {
	return it.fail
}

func (it *RegistryModuleOwnerCustomAdministratorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RegistryModuleOwnerCustomAdministratorRegistered struct {
	Token         common.Address
	Administrator common.Address
	Raw           types.Log
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomFilterer) FilterAdministratorRegistered(opts *bind.FilterOpts, token []common.Address, administrator []common.Address) (*RegistryModuleOwnerCustomAdministratorRegisteredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var administratorRule []interface{}
	for _, administratorItem := range administrator {
		administratorRule = append(administratorRule, administratorItem)
	}

	logs, sub, err := _RegistryModuleOwnerCustom.contract.FilterLogs(opts, "AdministratorRegistered", tokenRule, administratorRule)
	if err != nil {
		return nil, err
	}
	return &RegistryModuleOwnerCustomAdministratorRegisteredIterator{contract: _RegistryModuleOwnerCustom.contract, event: "AdministratorRegistered", logs: logs, sub: sub}, nil
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomFilterer) WatchAdministratorRegistered(opts *bind.WatchOpts, sink chan<- *RegistryModuleOwnerCustomAdministratorRegistered, token []common.Address, administrator []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var administratorRule []interface{}
	for _, administratorItem := range administrator {
		administratorRule = append(administratorRule, administratorItem)
	}

	logs, sub, err := _RegistryModuleOwnerCustom.contract.WatchLogs(opts, "AdministratorRegistered", tokenRule, administratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RegistryModuleOwnerCustomAdministratorRegistered)
				if err := _RegistryModuleOwnerCustom.contract.UnpackLog(event, "AdministratorRegistered", log); err != nil {
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

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustomFilterer) ParseAdministratorRegistered(log types.Log) (*RegistryModuleOwnerCustomAdministratorRegistered, error) {
	event := new(RegistryModuleOwnerCustomAdministratorRegistered)
	if err := _RegistryModuleOwnerCustom.contract.UnpackLog(event, "AdministratorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var RegistryModuleOwnerCustomZkBin string = ("0x00010000000000020001000000000002000000600310027000000067033001970000000100200190000000220000c13d0000008002000039000000400020043f000000040030008c000000db0000413d000000000201043b000000e0022002700000006e0020009c0000004f0000613d0000006f0020009c0000008f0000613d000000700020009c000000db0000c13d0000000001000416000000000001004b000000db0000c13d000000c001000039000000400010043f0000001f01000039000000800010043f0000007402000041000000a00020043f0000002003000039000000c00030043f000000e00010043f000001000020043f0000011f0000043f0000007501000041000001970001042e0000000002000416000000000002004b000000db0000c13d0000001f023000390000006802200197000000a002200039000000400020043f0000001f0430018f0000006905300198000000a002500039000000330000613d000000a006000039000000000701034f000000007807043c0000000006860436000000000026004b0000002f0000c13d000000000004004b000000400000613d000000000151034f0000000304400210000000000502043300000000054501cf000000000545022f000000000101043b0000010004400089000000000141022f00000000014101cf000000000151019f0000000000120435000000200030008c000000db0000413d000000a00100043d0000006a0010009c000000db0000213d0000006a01100198000000a30000c13d000000400100043d0000006c020000410000000000210435000000670010009c000000670100804100000040011002100000006d011001c70000019800010430000000240030008c000000db0000413d0000000002000416000000000002004b000000db0000c13d0000000401100370000000000101043b0000006a0010009c000000db0000213d0000006a021001970000007103000041000000800030043f0000000003000414000000040020008c0000009e0000613d000100000001001d000000670030009c0000006703008041000000c00130021000000072011001c7019601910000040f00000060031002700000006703300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000080057001bf000000730000613d0000008008000039000000000901034f000000009a09043c0000000008a80436000000000058004b0000006f0000c13d000000000006004b000000800000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000000000003001f0000000100200190000000d10000c13d0000001f0530018f0000006906300198000000400200043d0000000004620019000000eb0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b0000008a0000c13d000000eb0000013d000000240030008c000000db0000413d0000000002000416000000000002004b000000db0000c13d0000000401100370000000000101043b0000006a0010009c000000db0000213d0000006a021001970000007303000041000000800030043f0000000003000414000000040020008c000000ac0000c13d0000000003000031000000200030008c00000020040000390000000004034019000000d20000013d000000800010043f0000014000000443000001600010044300000020010000390000010000100443000000010100003900000120001004430000006b01000041000001970001042e000100000001001d000000670030009c0000006703008041000000c00130021000000072011001c7019601910000040f000000800a00003900000060031002700000006703300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000080057001bf000000c10000613d000000000801034f000000008908043c000000000a9a043600000000005a004b000000bd0000c13d000000000006004b000000ce0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000000000003001f0000000100200190000000e00000613d00000001010000290000001f02400039000000600220018f00000080022001bf000000400020043f000000200030008c000000db0000413d000000800200043d0000006a0020009c000000dd0000a13d00000000010000190000019800010430019600fe0000040f0000000001000019000001970001042e0000001f0530018f0000006906300198000000400200043d0000000004620019000000eb0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000000e70000c13d000000000005004b000000f80000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000000670020009c00000067020080410000004002200210000000000112019f00000198000104300003000000000002000200000001001d0003006a0020019b0000000001000411000000030010006b000001570000c13d000000780100004100000000001004430000000001000412000000040010044300000024000004430000000001000414000000670010009c0000006701008041000000c00110021000000079011001c70000800502000039019601910000040f0000000100200190000001540000613d000000000101043b0000007a0200004100000000002004430000006a01100197000100000001001d00000004001004430000000001000414000000670010009c0000006701008041000000c0011002100000007b011001c70000800202000039019601910000040f0000000100200190000001540000613d000000000101043b000000000001004b000001550000613d000000400400043d0000002401400039000000030200002900000000002104350000007c01000041000000000014043500000002010000290000006a051001970000000401400039000000000051043500000000010004140000000102000029000000040020008c000001440000613d000000670040009c000000670300004100000000030440190000004003300210000000670010009c0000006701008041000000c001100210000000000131019f00000077011001c7000200000005001d000100000004001d0196018c0000040f000000010400002900000002050000290000006003100270000000670030019d00000001002001900000016c0000613d0000007d0040009c000001660000813d000000400040043f0000000001000414000000670010009c0000006701008041000000c0011002100000007e011001c70000800d0200003900000003030000390000007f0400004100000003060000290196018c0000040f0000000100200190000001550000613d000000000001042d000000000001042f0000000001000019000001980001043000000002010000290000006a01100197000000400200043d0000002403200039000000000013043500000076010000410000000000120435000000040120003900000003030000290000000000310435000000670020009c0000006702008041000000400120021000000077011001c700000198000104300000008001000041000000000010043f0000004101000039000000040010043f0000008101000041000001980001043000000067033001970000001f0530018f0000006906300198000000400200043d0000000004620019000001780000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000001740000c13d000000000005004b000001850000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000000670020009c00000067020080410000004002200210000000000112019f0000019800010430000000000001042f0000018f002104210000000102000039000000000001042d0000000002000019000000000001042d00000194002104230000000102000039000000000001042d0000000002000019000000000001042d0000019600000432000001970001042e000001980001043000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000000000000000000000000001ffffffe000000000000000000000000000000000000000000000000000000000ffffffe0000000000000000000000000ffffffffffffffffffffffffffffffffffffffff00000002000000000000000000000000000000800000010000000000000000009fabe1c100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000000000000000000ff12c3540000000000000000000000000000000000000000000000000000000096ea2f7a00000000000000000000000000000000000000000000000000000000181f5a778fd6a6ac0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000008000000000000000008da5cb5b0000000000000000000000000000000000000000000000000000000052656769737472794d6f64756c654f776e6572437573746f6d20312e352e30000000000000000000000000000000000000000060000000c00000000000000000c454d182000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000044000000000000000000000000310ab089e4439a4c15d089f94afb7896ff553aecb10793d0ab882de59d99a32e02000002000000000000000000000000000000440000000000000000000000001806aa1896bbf26568e884a7374b41e002500962caba6a15023a8d90e8508b830200000200000000000000000000000000000024000000000000000000000000e677ae37000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000020000000000000000000000000000000000000000000000000000000000000009590fb70af4b833346363965e043a9339e8c7d378b8a2b903c75c277faec4f94e487b71000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024000000000000000000000000cc96f0fc3b19b78dddeab7f470039c0066d2b516d7c1854b459e02acd493c97e")

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustom) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _RegistryModuleOwnerCustom.abi.Events["AdministratorRegistered"].ID:
		return _RegistryModuleOwnerCustom.ParseAdministratorRegistered(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (RegistryModuleOwnerCustomAdministratorRegistered) Topic() common.Hash {
	return common.HexToHash("0x09590fb70af4b833346363965e043a9339e8c7d378b8a2b903c75c277faec4f9")
}

func (_RegistryModuleOwnerCustom *RegistryModuleOwnerCustom) Address() common.Address {
	return _RegistryModuleOwnerCustom.address
}

type CustomTransaction struct {
	*types.Transaction
	CustomHash common.Hash
}

func (tx *CustomTransaction) Hash() common.Hash {
	return tx.CustomHash
}

func ConvertToTransaction(resp zktypes.TransactionResponse) *CustomTransaction {
	dtx := &types.DynamicFeeTx{
		ChainID:   resp.ChainID.ToInt(),
		Nonce:     uint64(resp.Nonce),
		GasTipCap: resp.MaxPriorityFeePerGas.ToInt(),
		GasFeeCap: resp.MaxFeePerGas.ToInt(),
		To:        &resp.To,
		Value:     resp.Value.ToInt(),
		Data:      resp.Data,
		Gas:       uint64(resp.Gas),
	}

	tx := types.NewTx(dtx)
	customTransaction := CustomTransaction{Transaction: tx, CustomHash: resp.Hash}
	return &customTransaction
}

func DeployZkSyncRegistryModuleOwnerCustom(auth *bind.TransactOpts, backend bind.ContractBackend, params ...interface{}) (common.Address, *CustomTransaction, *RegistryModuleOwnerCustom, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	
	zksyncClient := zkSyncClient.NewClient(client.Client())
	
	wallet := auth.Context.Value("wallet").(*zkSyncAccounts.Wallet)
	
	
	decodedBytes := common.FromHex(RegistryModuleOwnerCustomZkBin)
	
	RegistryModuleOwnerCustomAbi, err := RegistryModuleOwnerCustomMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	constructor, err := RegistryModuleOwnerCustomAbi.Pack("", params...)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	hash, err := wallet.DeployWithCreate(nil, zkSyncAccounts.CreateTransaction{
		Bytecode: decodedBytes,
		Calldata: constructor,
	})
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	
	receipt, err := zksyncClient.WaitMined(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	tx, _, err := zksyncClient.TransactionByHash(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	

	ethTx := ConvertToTransaction(*tx)
	address := receipt.ContractAddress

	parsed, err := RegistryModuleOwnerCustomMetaData.GetAbi()
	contractBind := bind.NewBoundContract(address, *parsed, backend, backend, backend)

	contractReturn := &RegistryModuleOwnerCustom{address: address, abi: *parsed, RegistryModuleOwnerCustomCaller: RegistryModuleOwnerCustomCaller{contract: contractBind}, RegistryModuleOwnerCustomTransactor: RegistryModuleOwnerCustomTransactor{contract: contractBind}, RegistryModuleOwnerCustomFilterer: RegistryModuleOwnerCustomFilterer{contract: contractBind}}

	return address, ethTx, contractReturn, err
}

type RegistryModuleOwnerCustomInterface interface {
	TypeAndVersion(opts *bind.CallOpts) (string, error)

	RegisterAdminViaGetCCIPAdmin(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error)

	RegisterAdminViaOwner(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error)

	FilterAdministratorRegistered(opts *bind.FilterOpts, token []common.Address, administrator []common.Address) (*RegistryModuleOwnerCustomAdministratorRegisteredIterator, error)

	WatchAdministratorRegistered(opts *bind.WatchOpts, sink chan<- *RegistryModuleOwnerCustomAdministratorRegistered, token []common.Address, administrator []common.Address) (event.Subscription, error)

	ParseAdministratorRegistered(log types.Log) (*RegistryModuleOwnerCustomAdministratorRegistered, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
