// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ccip_reader_tester

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

type CCIPReaderTesterSourceChainConfig struct {
	IsEnabled bool
	MinSeqNr  uint64
	OnRamp    []byte
}

var CCIPReaderTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"getSourceChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPReaderTester.SourceChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPReaderTester.SourceChainConfig\",\"name\":\"sourceChainConfig\",\"type\":\"tuple\"}],\"name\":\"setSourceChainConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610667806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063e9d68a8e1461003b578063f831af8114610064575b600080fd5b61004e610049366004610222565b610079565b60405161005b9190610244565b60405180910390f35b610077610072366004610381565b61016a565b005b60408051606080820183526000808352602080840182905283850183905267ffffffffffffffff86811683528282529185902085519384018652805460ff8116151585526101009004909216908301526001810180549394929391928401916100e19061049c565b80601f016020809104026020016040519081016040528092919081815260200182805461010d9061049c565b801561015a5780601f1061012f5761010080835404028352916020019161015a565b820191906000526020600020905b81548152906001019060200180831161013d57829003601f168201915b5050505050815250509050919050565b67ffffffffffffffff8281166000908152602081815260409182902084518154928601517fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000009093169015157fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff1617610100929094169190910292909217825582015182919060018201906101fe9082610540565b5050505050565b803567ffffffffffffffff8116811461021d57600080fd5b919050565b60006020828403121561023457600080fd5b61023d82610205565b9392505050565b6000602080835283511515602084015267ffffffffffffffff60208501511660408401526040840151606080850152805180608086015260005b8181101561029a5782810184015186820160a00152830161027e565b50600060a0828701015260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116860101935050505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561032c5761032c6102da565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610379576103796102da565b604052919050565b6000806040838503121561039457600080fd5b61039d83610205565b915060208084013567ffffffffffffffff808211156103bb57600080fd5b90850190606082880312156103cf57600080fd5b6103d7610309565b823580151581146103e757600080fd5b81526103f4838501610205565b8482015260408301358281111561040a57600080fd5b80840193505087601f84011261041f57600080fd5b823582811115610431576104316102da565b610461857fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610332565b9250808352888582860101111561047757600080fd5b8085850186850137600085828501015250816040820152809450505050509250929050565b600181811c908216806104b057607f821691505b6020821081036104e9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f82111561053b576000816000526020600020601f850160051c810160208610156105185750805b601f850160051c820191505b8181101561053757828155600101610524565b5050505b505050565b815167ffffffffffffffff81111561055a5761055a6102da565b61056e81610568845461049c565b846104ef565b602080601f8311600181146105c1576000841561058b5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555610537565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561060e578886015182559484019460019091019084016105ef565b508582101561064a57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c6343000818000a",
}

var CCIPReaderTesterABI = CCIPReaderTesterMetaData.ABI

var CCIPReaderTesterBin = CCIPReaderTesterMetaData.Bin

func DeployCCIPReaderTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CCIPReaderTester, error) {
	parsed, err := CCIPReaderTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCIPReaderTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCIPReaderTester{address: address, abi: *parsed, CCIPReaderTesterCaller: CCIPReaderTesterCaller{contract: contract}, CCIPReaderTesterTransactor: CCIPReaderTesterTransactor{contract: contract}, CCIPReaderTesterFilterer: CCIPReaderTesterFilterer{contract: contract}}, nil
}

type CCIPReaderTester struct {
	address common.Address
	abi     abi.ABI
	CCIPReaderTesterCaller
	CCIPReaderTesterTransactor
	CCIPReaderTesterFilterer
}

type CCIPReaderTesterCaller struct {
	contract *bind.BoundContract
}

type CCIPReaderTesterTransactor struct {
	contract *bind.BoundContract
}

type CCIPReaderTesterFilterer struct {
	contract *bind.BoundContract
}

type CCIPReaderTesterSession struct {
	Contract     *CCIPReaderTester
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CCIPReaderTesterCallerSession struct {
	Contract *CCIPReaderTesterCaller
	CallOpts bind.CallOpts
}

type CCIPReaderTesterTransactorSession struct {
	Contract     *CCIPReaderTesterTransactor
	TransactOpts bind.TransactOpts
}

type CCIPReaderTesterRaw struct {
	Contract *CCIPReaderTester
}

type CCIPReaderTesterCallerRaw struct {
	Contract *CCIPReaderTesterCaller
}

type CCIPReaderTesterTransactorRaw struct {
	Contract *CCIPReaderTesterTransactor
}

func NewCCIPReaderTester(address common.Address, backend bind.ContractBackend) (*CCIPReaderTester, error) {
	abi, err := abi.JSON(strings.NewReader(CCIPReaderTesterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCCIPReaderTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCIPReaderTester{address: address, abi: abi, CCIPReaderTesterCaller: CCIPReaderTesterCaller{contract: contract}, CCIPReaderTesterTransactor: CCIPReaderTesterTransactor{contract: contract}, CCIPReaderTesterFilterer: CCIPReaderTesterFilterer{contract: contract}}, nil
}

func NewCCIPReaderTesterCaller(address common.Address, caller bind.ContractCaller) (*CCIPReaderTesterCaller, error) {
	contract, err := bindCCIPReaderTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCIPReaderTesterCaller{contract: contract}, nil
}

func NewCCIPReaderTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*CCIPReaderTesterTransactor, error) {
	contract, err := bindCCIPReaderTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCIPReaderTesterTransactor{contract: contract}, nil
}

func NewCCIPReaderTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*CCIPReaderTesterFilterer, error) {
	contract, err := bindCCIPReaderTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCIPReaderTesterFilterer{contract: contract}, nil
}

func bindCCIPReaderTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCIPReaderTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CCIPReaderTester *CCIPReaderTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCIPReaderTester.Contract.CCIPReaderTesterCaller.contract.Call(opts, result, method, params...)
}

func (_CCIPReaderTester *CCIPReaderTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPReaderTester.Contract.CCIPReaderTesterTransactor.contract.Transfer(opts)
}

func (_CCIPReaderTester *CCIPReaderTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCIPReaderTester.Contract.CCIPReaderTesterTransactor.contract.Transact(opts, method, params...)
}

func (_CCIPReaderTester *CCIPReaderTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCIPReaderTester.Contract.contract.Call(opts, result, method, params...)
}

func (_CCIPReaderTester *CCIPReaderTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPReaderTester.Contract.contract.Transfer(opts)
}

func (_CCIPReaderTester *CCIPReaderTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCIPReaderTester.Contract.contract.Transact(opts, method, params...)
}

func (_CCIPReaderTester *CCIPReaderTesterCaller) GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (CCIPReaderTesterSourceChainConfig, error) {
	var out []interface{}
	err := _CCIPReaderTester.contract.Call(opts, &out, "getSourceChainConfig", sourceChainSelector)

	if err != nil {
		return *new(CCIPReaderTesterSourceChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCIPReaderTesterSourceChainConfig)).(*CCIPReaderTesterSourceChainConfig)

	return out0, err

}

func (_CCIPReaderTester *CCIPReaderTesterSession) GetSourceChainConfig(sourceChainSelector uint64) (CCIPReaderTesterSourceChainConfig, error) {
	return _CCIPReaderTester.Contract.GetSourceChainConfig(&_CCIPReaderTester.CallOpts, sourceChainSelector)
}

func (_CCIPReaderTester *CCIPReaderTesterCallerSession) GetSourceChainConfig(sourceChainSelector uint64) (CCIPReaderTesterSourceChainConfig, error) {
	return _CCIPReaderTester.Contract.GetSourceChainConfig(&_CCIPReaderTester.CallOpts, sourceChainSelector)
}

func (_CCIPReaderTester *CCIPReaderTesterTransactor) SetSourceChainConfig(opts *bind.TransactOpts, sourceChainSelector uint64, sourceChainConfig CCIPReaderTesterSourceChainConfig) (*types.Transaction, error) {
	return _CCIPReaderTester.contract.Transact(opts, "setSourceChainConfig", sourceChainSelector, sourceChainConfig)
}

func (_CCIPReaderTester *CCIPReaderTesterSession) SetSourceChainConfig(sourceChainSelector uint64, sourceChainConfig CCIPReaderTesterSourceChainConfig) (*types.Transaction, error) {
	return _CCIPReaderTester.Contract.SetSourceChainConfig(&_CCIPReaderTester.TransactOpts, sourceChainSelector, sourceChainConfig)
}

func (_CCIPReaderTester *CCIPReaderTesterTransactorSession) SetSourceChainConfig(sourceChainSelector uint64, sourceChainConfig CCIPReaderTesterSourceChainConfig) (*types.Transaction, error) {
	return _CCIPReaderTester.Contract.SetSourceChainConfig(&_CCIPReaderTester.TransactOpts, sourceChainSelector, sourceChainConfig)
}

func (_CCIPReaderTester *CCIPReaderTester) Address() common.Address {
	return _CCIPReaderTester.address
}

type CCIPReaderTesterInterface interface {
	GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (CCIPReaderTesterSourceChainConfig, error)

	SetSourceChainConfig(opts *bind.TransactOpts, sourceChainSelector uint64, sourceChainConfig CCIPReaderTesterSourceChainConfig) (*types.Transaction, error)

	Address() common.Address
}
