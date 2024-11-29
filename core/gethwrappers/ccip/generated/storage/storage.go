// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage

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

var StorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_oldNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"storedNumber\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060bc8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80632e64cec11460375780636057361d14604c575b600080fd5b60005460405190815260200160405180910390f35b605b60573660046097565b605d565b005b6000805482825560405190918391839133917f87f16aa184eca14ea45e132328a5effbb79b9f921657bd03d83608f26d76f3cf9190a45050565b60006020828403121560a857600080fd5b503591905056fea164736f6c6343000818000a",
}

var StorageABI = StorageMetaData.ABI

var StorageBin = StorageMetaData.Bin

func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *CustomTransaction, *Storage, error) {
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
		return DeployZkSyncStorage(auth, backend)
	}

	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, &CustomTransaction{Transaction: tx, CustomHash: tx.Hash()}, &Storage{address: address, abi: *parsed, StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

type Storage struct {
	address common.Address
	abi     abi.ABI
	StorageCaller
	StorageTransactor
	StorageFilterer
}

type StorageCaller struct {
	contract *bind.BoundContract
}

type StorageTransactor struct {
	contract *bind.BoundContract
}

type StorageFilterer struct {
	contract *bind.BoundContract
}

type StorageSession struct {
	Contract     *Storage
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type StorageCallerSession struct {
	Contract *StorageCaller
	CallOpts bind.CallOpts
}

type StorageTransactorSession struct {
	Contract     *StorageTransactor
	TransactOpts bind.TransactOpts
}

type StorageRaw struct {
	Contract *Storage
}

type StorageCallerRaw struct {
	Contract *StorageCaller
}

type StorageTransactorRaw struct {
	Contract *StorageTransactor
}

func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	abi, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{address: address, abi: abi, StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

func (_Storage *StorageCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Storage *StorageSession) Retrieve() (*big.Int, error) {
	return _Storage.Contract.Retrieve(&_Storage.CallOpts)
}

func (_Storage *StorageCallerSession) Retrieve() (*big.Int, error) {
	return _Storage.Contract.Retrieve(&_Storage.CallOpts)
}

func (_Storage *StorageTransactor) Store(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "store", num)
}

func (_Storage *StorageSession) Store(num *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Store(&_Storage.TransactOpts, num)
}

func (_Storage *StorageTransactorSession) Store(num *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Store(&_Storage.TransactOpts, num)
}

type StorageStoredNumberIterator struct {
	Event *StorageStoredNumber

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *StorageStoredNumberIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageStoredNumber)
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
		it.Event = new(StorageStoredNumber)
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

func (it *StorageStoredNumberIterator) Error() error {
	return it.fail
}

func (it *StorageStoredNumberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type StorageStoredNumber struct {
	From      common.Address
	OldNumber *big.Int
	Number    *big.Int
	Raw       types.Log
}

func (_Storage *StorageFilterer) FilterStoredNumber(opts *bind.FilterOpts, _from []common.Address, _oldNumber []*big.Int, _number []*big.Int) (*StorageStoredNumberIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _oldNumberRule []interface{}
	for _, _oldNumberItem := range _oldNumber {
		_oldNumberRule = append(_oldNumberRule, _oldNumberItem)
	}
	var _numberRule []interface{}
	for _, _numberItem := range _number {
		_numberRule = append(_numberRule, _numberItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "storedNumber", _fromRule, _oldNumberRule, _numberRule)
	if err != nil {
		return nil, err
	}
	return &StorageStoredNumberIterator{contract: _Storage.contract, event: "storedNumber", logs: logs, sub: sub}, nil
}

func (_Storage *StorageFilterer) WatchStoredNumber(opts *bind.WatchOpts, sink chan<- *StorageStoredNumber, _from []common.Address, _oldNumber []*big.Int, _number []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _oldNumberRule []interface{}
	for _, _oldNumberItem := range _oldNumber {
		_oldNumberRule = append(_oldNumberRule, _oldNumberItem)
	}
	var _numberRule []interface{}
	for _, _numberItem := range _number {
		_numberRule = append(_numberRule, _numberItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "storedNumber", _fromRule, _oldNumberRule, _numberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(StorageStoredNumber)
				if err := _Storage.contract.UnpackLog(event, "storedNumber", log); err != nil {
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

func (_Storage *StorageFilterer) ParseStoredNumber(log types.Log) (*StorageStoredNumber, error) {
	event := new(StorageStoredNumber)
	if err := _Storage.contract.UnpackLog(event, "storedNumber", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var StorageZkBin string = ("0x0000008003000039000000400030043f0000000100200190000000150000c13d00000060021002700000001002200197000000040020008c000000340000413d000000000301043b000000e003300270000000120030009c0000001d0000613d000000130030009c000000340000c13d0000000001000416000000000001004b000000340000c13d000000000100041a000000800010043f00000016010000410000003c0001042e0000000001000416000000000001004b000000340000c13d00000020010000390000010000100443000001200000044300000011010000410000003c0001042e000000240020008c000000340000413d0000000002000416000000000002004b000000340000c13d0000000401100370000000000701043b000000000600041a000000000070041b0000000001000414000000100010009c0000001001008041000000c00110021000000014011001c70000800d02000039000000040300003900000000050004110000001504000041003b00360000040f0000000100200190000000340000613d00000000010000190000003c0001042e00000000010000190000003d0001043000000039002104210000000102000039000000000001042d0000000002000019000000000001042d0000003b000004320000003c0001042e0000003d000104300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffff0000000200000000000000000000000000000040000001000000000000000000000000000000000000000000000000000000000000000000000000006057361d000000000000000000000000000000000000000000000000000000002e64cec1020000000000000000000000000000000000000000000000000000000000000087f16aa184eca14ea45e132328a5effbb79b9f921657bd03d83608f26d76f3cf00000000000000000000000000000000000000200000008000000000000000000000000000000000000000000000000000000000000000000000000000000000d05624d43b7436b857a8c30563e36489cd4edac61bd40b452ab9052253f0fef5")

func (_Storage *Storage) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Storage.abi.Events["StoredNumber"].ID:
		return _Storage.ParseStoredNumber(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (StorageStoredNumber) Topic() common.Hash {
	return common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
}

func (_Storage *Storage) Address() common.Address {
	return _Storage.address
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

func DeployZkSyncStorage(auth *bind.TransactOpts, backend bind.ContractBackend, params ...interface{}) (common.Address, *CustomTransaction, *Storage, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}

	zksyncClient := zkSyncClient.NewClient(client.Client())

	wallet := auth.Context.Value("wallet").(*zkSyncAccounts.Wallet)

	decodedBytes := common.FromHex(StorageZkBin)

	StorageAbi, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	constructor, err := StorageAbi.Pack("", params...)
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

	parsed, err := StorageMetaData.GetAbi()
	contractBind := bind.NewBoundContract(address, *parsed, backend, backend, backend)

	contractReturn := &Storage{address: address, abi: *parsed, StorageCaller: StorageCaller{contract: contractBind}, StorageTransactor: StorageTransactor{contract: contractBind}, StorageFilterer: StorageFilterer{contract: contractBind}}

	return address, ethTx, contractReturn, err
}

type StorageInterface interface {
	Retrieve(opts *bind.CallOpts) (*big.Int, error)

	Store(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error)

	FilterStoredNumber(opts *bind.FilterOpts, _from []common.Address, _oldNumber []*big.Int, _number []*big.Int) (*StorageStoredNumberIterator, error)

	WatchStoredNumber(opts *bind.WatchOpts, sink chan<- *StorageStoredNumber, _from []common.Address, _oldNumber []*big.Int, _number []*big.Int) (event.Subscription, error)

	ParseStoredNumber(log types.Log) (*StorageStoredNumber, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
