// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fee_manager

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

type InternalFeeUpdate struct {
	SourceFeeToken              common.Address
	DestChainId                 uint64
	FeeTokenBaseUnitsPerUnitGas *big.Int
}

var FeeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceFeeToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"feeTokenBaseUnitsPerUnitGas\",\"type\":\"uint128\"}],\"internalType\":\"structInternal.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"feeUpdaters\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"stalenessThreshold\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawalAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByUpdaterOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"timePassed\",\"type\":\"uint256\"}],\"name\":\"StaleFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chain\",\"type\":\"uint64\"}],\"name\":\"TokenOrChainNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChain\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"feeTokenBaseUnitsPerUnitGas\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"GasFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"name\":\"getFeeTokenBaseUnitsPerUnitGas\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"feeTokenBaseUnitsPerUnitGas\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStalenessThreshold\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"isFeeUpdater\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"removeFeeUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"setFeeUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceFeeToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"feeTokenBaseUnitsPerUnitGas\",\"type\":\"uint128\"}],\"internalType\":\"structInternal.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"}],\"name\":\"updateFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200178a3803806200178a8339810160408190526200003491620004fb565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620001b5565b50505060005b8351811015620001545762000141848281518110620000e757620000e762000632565b60200260200101516000015185838151811062000108576200010862000632565b60200260200101516020015186848151811062000129576200012962000632565b6020026020010151604001516200026060201b60201c565b6200014c8162000648565b9050620000c4565b5060005b8251811015620001a2576200018f8382815181106200017b576200017b62000632565b60200260200101516200035660201b60201c565b6200019a8162000648565b905062000158565b5063ffffffff1660805250620006709050565b336001600160a01b038216036200020f5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b0383166200028857604051634655efd160e11b815260040160405180910390fd5b6040805180820182526001600160801b0380841682526001600160401b034281811660208086019182526001600160a01b038a166000818152600283528881208b87168083529352889020965187549351909516600160801b026001600160c01b0319909316949095169390931717909355925190917fb230bad3704091781ab962bc58267145fa2aa1c542698c04e10e9db6069fa2d391620003499186916001600160801b039290921682526001600160401b0316602082015260400190565b60405180910390a3505050565b6001600160a01b03811615620003b0576001600160a01b038116600081815260036020526040808220805460ff19166001179055517fa462a3423511fce5ad038a1130251b66e3b8c135fa2ca15777f45c72ea3037c59190a25b50565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715620003ee57620003ee620003b3565b60405290565b604051601f8201601f191681016001600160401b03811182821017156200041f576200041f620003b3565b604052919050565b60006001600160401b03821115620004435762000443620003b3565b5060051b60200190565b80516001600160a01b03811681146200046557600080fd5b919050565b600082601f8301126200047c57600080fd5b81516020620004956200048f8362000427565b620003f4565b82815260059290921b84018101918181019086841115620004b557600080fd5b8286015b84811015620004db57620004cd816200044d565b8352918301918301620004b9565b509695505050505050565b805163ffffffff811681146200046557600080fd5b600080600060608085870312156200051257600080fd5b84516001600160401b03808211156200052a57600080fd5b818701915087601f8301126200053f57600080fd5b81516020620005526200048f8362000427565b8281529185028401810191818101908b8411156200056f57600080fd5b948201945b83861015620005ee5786868d0312156200058e5760008081fd5b62000598620003c9565b620005a3876200044d565b8152838701518681168114620005b95760008081fd5b818501526040878101516001600160801b0381168114620005da5760008081fd5b908201528252948601949082019062000574565b918a01519198509094505050808311156200060857600080fd5b505062000618868287016200046a565b9250506200062960408501620004e6565b90509250925092565b634e487b7160e01b600052603260045260246000fd5b6000600182016200066957634e487b7160e01b600052601160045260246000fd5b5060010190565b6080516110f06200069a600039600081816101710152818161039c01526103f501526110f06000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c80639086658e11610076578063ae7fca181161005b578063ae7fca181461019b578063c60d9b98146101ae578063f2fde38b146101f757600080fd5b80639086658e14610154578063a6c94a731461016757600080fd5b8063604782e6116100a7578063604782e61461011157806379ba5097146101245780638da5cb5b1461012c57600080fd5b806301e33667146100c3578063268e5d48146100d8575b600080fd5b6100d66100d1366004610d2d565b61020a565b005b6100eb6100e6366004610d81565b610285565b6040516fffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100d661011f366004610db4565b610432565b6100d6610446565b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610108565b6100d6610162366004610e76565b610529565b60405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610108565b6100d66101a9366004610db4565b61061a565b6101e76101bc366004610db4565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205460ff1690565b6040519015158152602001610108565b6100d6610205366004610db4565b61062b565b61021261063c565b73ffffffffffffffffffffffffffffffffffffffff821661025f576040517f84c2102600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61028073ffffffffffffffffffffffffffffffffffffffff841683836106a5565b505050565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260026020908152604080832067ffffffffffffffff80861685529083528184208251808401909352546fffffffffffffffffffffffffffffffff81168352700100000000000000000000000000000000900416918101829052901580610318575080516fffffffffffffffffffffffffffffffff16155b1561037c576040517f102e3c2800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015267ffffffffffffffff841660248201526044015b60405180910390fd5b6000816020015167ffffffffffffffff16426103989190610fa6565b90507f000000000000000000000000000000000000000000000000000000000000000063ffffffff16811115610429576040517f55cf089a00000000000000000000000000000000000000000000000000000000815263ffffffff7f000000000000000000000000000000000000000000000000000000000000000016600482015260248101829052604401610373565b50519392505050565b61043a61063c565b61044381610732565b50565b60015473ffffffffffffffffffffffffffffffffffffffff1633146104ad5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610373565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60005473ffffffffffffffffffffffffffffffffffffffff16331480159061056157503360009081526003602052604090205460ff16155b15610598576040517f46f0815400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8151811015610616576106068282815181106105b9576105b9610fbd565b6020026020010151600001518383815181106105d7576105d7610fbd565b6020026020010151602001518484815181106105f5576105f5610fbd565b6020026020010151604001516107c5565b61060f81610fec565b905061059b565b5050565b61062261063c565b61044381610925565b61063361063c565b61044381610999565b60005473ffffffffffffffffffffffffffffffffffffffff1633146106a35760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610373565b565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610280908490610a74565b73ffffffffffffffffffffffffffffffffffffffff8116156104435773ffffffffffffffffffffffffffffffffffffffff811660008181526003602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055517fa462a3423511fce5ad038a1130251b66e3b8c135fa2ca15777f45c72ea3037c59190a250565b73ffffffffffffffffffffffffffffffffffffffff8316610812576040517f8cabdfa200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820182526fffffffffffffffffffffffffffffffff808416825267ffffffffffffffff42818116602080860191825273ffffffffffffffffffffffffffffffffffffffff8a166000818152600283528881208b87168083529352889020965187549351909516700100000000000000000000000000000000027fffffffffffffffff000000000000000000000000000000000000000000000000909316949095169390931717909355925190917fb230bad3704091781ab962bc58267145fa2aa1c542698c04e10e9db6069fa2d3916109189186916fffffffffffffffffffffffffffffffff92909216825267ffffffffffffffff16602082015260400190565b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff811660008181526003602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517f74a2c31badb27f0acfb9da3ef34c9e656ca1723881466e89a40f791f1c82ee719190a250565b3373ffffffffffffffffffffffffffffffffffffffff8216036109fe5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610373565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000610ad6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16610b669092919063ffffffff16565b8051909150156102805780806020019051810190610af49190611024565b6102805760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610373565b6060610b758484600085610b7f565b90505b9392505050565b606082471015610bf75760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610373565b843b610c455760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610373565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610c6e9190611076565b60006040518083038185875af1925050503d8060008114610cab576040519150601f19603f3d011682016040523d82523d6000602084013e610cb0565b606091505b5091509150610cc0828286610ccb565b979650505050505050565b60608315610cda575081610b78565b825115610cea5782518084602001fd5b8160405162461bcd60e51b81526004016103739190611092565b803573ffffffffffffffffffffffffffffffffffffffff81168114610d2857600080fd5b919050565b600080600060608486031215610d4257600080fd5b610d4b84610d04565b9250610d5960208501610d04565b9150604084013590509250925092565b803567ffffffffffffffff81168114610d2857600080fd5b60008060408385031215610d9457600080fd5b610d9d83610d04565b9150610dab60208401610d69565b90509250929050565b600060208284031215610dc657600080fd5b610b7882610d04565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715610e2157610e21610dcf565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610e6e57610e6e610dcf565b604052919050565b60006020808385031215610e8957600080fd5b823567ffffffffffffffff80821115610ea157600080fd5b818501915085601f830112610eb557600080fd5b813581811115610ec757610ec7610dcf565b610ed5848260051b01610e27565b81815284810192506060918202840185019188831115610ef457600080fd5b938501935b82851015610f6b5780858a031215610f115760008081fd5b610f19610dfe565b610f2286610d04565b8152610f2f878701610d69565b878201526040808701356fffffffffffffffffffffffffffffffff81168114610f585760008081fd5b9082015284529384019392850192610ef9565b50979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015610fb857610fb8610f77565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361101d5761101d610f77565b5060010190565b60006020828403121561103657600080fd5b81518015158114610b7857600080fd5b60005b83811015611061578181015183820152602001611049565b83811115611070576000848401525b50505050565b60008251611088818460208701611046565b9190910192915050565b60208152600082518060208401526110b1816040850160208701611046565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fea164736f6c634300080f000a",
}

var FeeManagerABI = FeeManagerMetaData.ABI

var FeeManagerBin = FeeManagerMetaData.Bin

func DeployFeeManager(auth *bind.TransactOpts, backend bind.ContractBackend, feeUpdates []InternalFeeUpdate, feeUpdaters []common.Address, stalenessThreshold uint32) (common.Address, *types.Transaction, *FeeManager, error) {
	parsed, err := FeeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeManagerBin), backend, feeUpdates, feeUpdaters, stalenessThreshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeManager{FeeManagerCaller: FeeManagerCaller{contract: contract}, FeeManagerTransactor: FeeManagerTransactor{contract: contract}, FeeManagerFilterer: FeeManagerFilterer{contract: contract}}, nil
}

type FeeManager struct {
	address common.Address
	abi     abi.ABI
	FeeManagerCaller
	FeeManagerTransactor
	FeeManagerFilterer
}

type FeeManagerCaller struct {
	contract *bind.BoundContract
}

type FeeManagerTransactor struct {
	contract *bind.BoundContract
}

type FeeManagerFilterer struct {
	contract *bind.BoundContract
}

type FeeManagerSession struct {
	Contract     *FeeManager
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type FeeManagerCallerSession struct {
	Contract *FeeManagerCaller
	CallOpts bind.CallOpts
}

type FeeManagerTransactorSession struct {
	Contract     *FeeManagerTransactor
	TransactOpts bind.TransactOpts
}

type FeeManagerRaw struct {
	Contract *FeeManager
}

type FeeManagerCallerRaw struct {
	Contract *FeeManagerCaller
}

type FeeManagerTransactorRaw struct {
	Contract *FeeManagerTransactor
}

func NewFeeManager(address common.Address, backend bind.ContractBackend) (*FeeManager, error) {
	abi, err := abi.JSON(strings.NewReader(FeeManagerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindFeeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeManager{address: address, abi: abi, FeeManagerCaller: FeeManagerCaller{contract: contract}, FeeManagerTransactor: FeeManagerTransactor{contract: contract}, FeeManagerFilterer: FeeManagerFilterer{contract: contract}}, nil
}

func NewFeeManagerCaller(address common.Address, caller bind.ContractCaller) (*FeeManagerCaller, error) {
	contract, err := bindFeeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeManagerCaller{contract: contract}, nil
}

func NewFeeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeManagerTransactor, error) {
	contract, err := bindFeeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeManagerTransactor{contract: contract}, nil
}

func NewFeeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeManagerFilterer, error) {
	contract, err := bindFeeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeManagerFilterer{contract: contract}, nil
}

func bindFeeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_FeeManager *FeeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeManager.Contract.FeeManagerCaller.contract.Call(opts, result, method, params...)
}

func (_FeeManager *FeeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeManager.Contract.FeeManagerTransactor.contract.Transfer(opts)
}

func (_FeeManager *FeeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeManager.Contract.FeeManagerTransactor.contract.Transact(opts, method, params...)
}

func (_FeeManager *FeeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeManager.Contract.contract.Call(opts, result, method, params...)
}

func (_FeeManager *FeeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeManager.Contract.contract.Transfer(opts)
}

func (_FeeManager *FeeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeManager.Contract.contract.Transact(opts, method, params...)
}

func (_FeeManager *FeeManagerCaller) GetFeeTokenBaseUnitsPerUnitGas(opts *bind.CallOpts, token common.Address, destChainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "getFeeTokenBaseUnitsPerUnitGas", token, destChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeManager *FeeManagerSession) GetFeeTokenBaseUnitsPerUnitGas(token common.Address, destChainId uint64) (*big.Int, error) {
	return _FeeManager.Contract.GetFeeTokenBaseUnitsPerUnitGas(&_FeeManager.CallOpts, token, destChainId)
}

func (_FeeManager *FeeManagerCallerSession) GetFeeTokenBaseUnitsPerUnitGas(token common.Address, destChainId uint64) (*big.Int, error) {
	return _FeeManager.Contract.GetFeeTokenBaseUnitsPerUnitGas(&_FeeManager.CallOpts, token, destChainId)
}

func (_FeeManager *FeeManagerCaller) GetStalenessThreshold(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "getStalenessThreshold")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_FeeManager *FeeManagerSession) GetStalenessThreshold() (uint32, error) {
	return _FeeManager.Contract.GetStalenessThreshold(&_FeeManager.CallOpts)
}

func (_FeeManager *FeeManagerCallerSession) GetStalenessThreshold() (uint32, error) {
	return _FeeManager.Contract.GetStalenessThreshold(&_FeeManager.CallOpts)
}

func (_FeeManager *FeeManagerCaller) IsFeeUpdater(opts *bind.CallOpts, feeUpdater common.Address) (bool, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "isFeeUpdater", feeUpdater)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_FeeManager *FeeManagerSession) IsFeeUpdater(feeUpdater common.Address) (bool, error) {
	return _FeeManager.Contract.IsFeeUpdater(&_FeeManager.CallOpts, feeUpdater)
}

func (_FeeManager *FeeManagerCallerSession) IsFeeUpdater(feeUpdater common.Address) (bool, error) {
	return _FeeManager.Contract.IsFeeUpdater(&_FeeManager.CallOpts, feeUpdater)
}

func (_FeeManager *FeeManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FeeManager *FeeManagerSession) Owner() (common.Address, error) {
	return _FeeManager.Contract.Owner(&_FeeManager.CallOpts)
}

func (_FeeManager *FeeManagerCallerSession) Owner() (common.Address, error) {
	return _FeeManager.Contract.Owner(&_FeeManager.CallOpts)
}

func (_FeeManager *FeeManagerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "acceptOwnership")
}

func (_FeeManager *FeeManagerSession) AcceptOwnership() (*types.Transaction, error) {
	return _FeeManager.Contract.AcceptOwnership(&_FeeManager.TransactOpts)
}

func (_FeeManager *FeeManagerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FeeManager.Contract.AcceptOwnership(&_FeeManager.TransactOpts)
}

func (_FeeManager *FeeManagerTransactor) RemoveFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "removeFeeUpdater", feeUpdater)
}

func (_FeeManager *FeeManagerSession) RemoveFeeUpdater(feeUpdater common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.RemoveFeeUpdater(&_FeeManager.TransactOpts, feeUpdater)
}

func (_FeeManager *FeeManagerTransactorSession) RemoveFeeUpdater(feeUpdater common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.RemoveFeeUpdater(&_FeeManager.TransactOpts, feeUpdater)
}

func (_FeeManager *FeeManagerTransactor) SetFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "setFeeUpdater", feeUpdater)
}

func (_FeeManager *FeeManagerSession) SetFeeUpdater(feeUpdater common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.SetFeeUpdater(&_FeeManager.TransactOpts, feeUpdater)
}

func (_FeeManager *FeeManagerTransactorSession) SetFeeUpdater(feeUpdater common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.SetFeeUpdater(&_FeeManager.TransactOpts, feeUpdater)
}

func (_FeeManager *FeeManagerTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "transferOwnership", to)
}

func (_FeeManager *FeeManagerSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.TransferOwnership(&_FeeManager.TransactOpts, to)
}

func (_FeeManager *FeeManagerTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.TransferOwnership(&_FeeManager.TransactOpts, to)
}

func (_FeeManager *FeeManagerTransactor) UpdateFees(opts *bind.TransactOpts, feeUpdates []InternalFeeUpdate) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "updateFees", feeUpdates)
}

func (_FeeManager *FeeManagerSession) UpdateFees(feeUpdates []InternalFeeUpdate) (*types.Transaction, error) {
	return _FeeManager.Contract.UpdateFees(&_FeeManager.TransactOpts, feeUpdates)
}

func (_FeeManager *FeeManagerTransactorSession) UpdateFees(feeUpdates []InternalFeeUpdate) (*types.Transaction, error) {
	return _FeeManager.Contract.UpdateFees(&_FeeManager.TransactOpts, feeUpdates)
}

func (_FeeManager *FeeManagerTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "withdrawToken", token, to, amount)
}

func (_FeeManager *FeeManagerSession) WithdrawToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FeeManager.Contract.WithdrawToken(&_FeeManager.TransactOpts, token, to, amount)
}

func (_FeeManager *FeeManagerTransactorSession) WithdrawToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FeeManager.Contract.WithdrawToken(&_FeeManager.TransactOpts, token, to, amount)
}

type FeeManagerFeeUpdaterRemovedIterator struct {
	Event *FeeManagerFeeUpdaterRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeManagerFeeUpdaterRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerFeeUpdaterRemoved)
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
		it.Event = new(FeeManagerFeeUpdaterRemoved)
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

func (it *FeeManagerFeeUpdaterRemovedIterator) Error() error {
	return it.fail
}

func (it *FeeManagerFeeUpdaterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeManagerFeeUpdaterRemoved struct {
	FeeUpdater common.Address
	Raw        types.Log
}

func (_FeeManager *FeeManagerFilterer) FilterFeeUpdaterRemoved(opts *bind.FilterOpts, feeUpdater []common.Address) (*FeeManagerFeeUpdaterRemovedIterator, error) {

	var feeUpdaterRule []interface{}
	for _, feeUpdaterItem := range feeUpdater {
		feeUpdaterRule = append(feeUpdaterRule, feeUpdaterItem)
	}

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "FeeUpdaterRemoved", feeUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &FeeManagerFeeUpdaterRemovedIterator{contract: _FeeManager.contract, event: "FeeUpdaterRemoved", logs: logs, sub: sub}, nil
}

func (_FeeManager *FeeManagerFilterer) WatchFeeUpdaterRemoved(opts *bind.WatchOpts, sink chan<- *FeeManagerFeeUpdaterRemoved, feeUpdater []common.Address) (event.Subscription, error) {

	var feeUpdaterRule []interface{}
	for _, feeUpdaterItem := range feeUpdater {
		feeUpdaterRule = append(feeUpdaterRule, feeUpdaterItem)
	}

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "FeeUpdaterRemoved", feeUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeManagerFeeUpdaterRemoved)
				if err := _FeeManager.contract.UnpackLog(event, "FeeUpdaterRemoved", log); err != nil {
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

func (_FeeManager *FeeManagerFilterer) ParseFeeUpdaterRemoved(log types.Log) (*FeeManagerFeeUpdaterRemoved, error) {
	event := new(FeeManagerFeeUpdaterRemoved)
	if err := _FeeManager.contract.UnpackLog(event, "FeeUpdaterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeManagerFeeUpdaterSetIterator struct {
	Event *FeeManagerFeeUpdaterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeManagerFeeUpdaterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerFeeUpdaterSet)
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
		it.Event = new(FeeManagerFeeUpdaterSet)
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

func (it *FeeManagerFeeUpdaterSetIterator) Error() error {
	return it.fail
}

func (it *FeeManagerFeeUpdaterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeManagerFeeUpdaterSet struct {
	FeeUpdater common.Address
	Raw        types.Log
}

func (_FeeManager *FeeManagerFilterer) FilterFeeUpdaterSet(opts *bind.FilterOpts, feeUpdater []common.Address) (*FeeManagerFeeUpdaterSetIterator, error) {

	var feeUpdaterRule []interface{}
	for _, feeUpdaterItem := range feeUpdater {
		feeUpdaterRule = append(feeUpdaterRule, feeUpdaterItem)
	}

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "FeeUpdaterSet", feeUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &FeeManagerFeeUpdaterSetIterator{contract: _FeeManager.contract, event: "FeeUpdaterSet", logs: logs, sub: sub}, nil
}

func (_FeeManager *FeeManagerFilterer) WatchFeeUpdaterSet(opts *bind.WatchOpts, sink chan<- *FeeManagerFeeUpdaterSet, feeUpdater []common.Address) (event.Subscription, error) {

	var feeUpdaterRule []interface{}
	for _, feeUpdaterItem := range feeUpdater {
		feeUpdaterRule = append(feeUpdaterRule, feeUpdaterItem)
	}

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "FeeUpdaterSet", feeUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeManagerFeeUpdaterSet)
				if err := _FeeManager.contract.UnpackLog(event, "FeeUpdaterSet", log); err != nil {
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

func (_FeeManager *FeeManagerFilterer) ParseFeeUpdaterSet(log types.Log) (*FeeManagerFeeUpdaterSet, error) {
	event := new(FeeManagerFeeUpdaterSet)
	if err := _FeeManager.contract.UnpackLog(event, "FeeUpdaterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeManagerGasFeeUpdatedIterator struct {
	Event *FeeManagerGasFeeUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeManagerGasFeeUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerGasFeeUpdated)
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
		it.Event = new(FeeManagerGasFeeUpdated)
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

func (it *FeeManagerGasFeeUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeManagerGasFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeManagerGasFeeUpdated struct {
	Token                       common.Address
	DestChain                   uint64
	FeeTokenBaseUnitsPerUnitGas *big.Int
	Timestamp                   uint64
	Raw                         types.Log
}

func (_FeeManager *FeeManagerFilterer) FilterGasFeeUpdated(opts *bind.FilterOpts, token []common.Address, destChain []uint64) (*FeeManagerGasFeeUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var destChainRule []interface{}
	for _, destChainItem := range destChain {
		destChainRule = append(destChainRule, destChainItem)
	}

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "GasFeeUpdated", tokenRule, destChainRule)
	if err != nil {
		return nil, err
	}
	return &FeeManagerGasFeeUpdatedIterator{contract: _FeeManager.contract, event: "GasFeeUpdated", logs: logs, sub: sub}, nil
}

func (_FeeManager *FeeManagerFilterer) WatchGasFeeUpdated(opts *bind.WatchOpts, sink chan<- *FeeManagerGasFeeUpdated, token []common.Address, destChain []uint64) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var destChainRule []interface{}
	for _, destChainItem := range destChain {
		destChainRule = append(destChainRule, destChainItem)
	}

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "GasFeeUpdated", tokenRule, destChainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeManagerGasFeeUpdated)
				if err := _FeeManager.contract.UnpackLog(event, "GasFeeUpdated", log); err != nil {
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

func (_FeeManager *FeeManagerFilterer) ParseGasFeeUpdated(log types.Log) (*FeeManagerGasFeeUpdated, error) {
	event := new(FeeManagerGasFeeUpdated)
	if err := _FeeManager.contract.UnpackLog(event, "GasFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeManagerOwnershipTransferRequestedIterator struct {
	Event *FeeManagerOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeManagerOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerOwnershipTransferRequested)
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
		it.Event = new(FeeManagerOwnershipTransferRequested)
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

func (it *FeeManagerOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *FeeManagerOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeManagerOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FeeManager *FeeManagerFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeManagerOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeeManagerOwnershipTransferRequestedIterator{contract: _FeeManager.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_FeeManager *FeeManagerFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FeeManagerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeManagerOwnershipTransferRequested)
				if err := _FeeManager.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_FeeManager *FeeManagerFilterer) ParseOwnershipTransferRequested(log types.Log) (*FeeManagerOwnershipTransferRequested, error) {
	event := new(FeeManagerOwnershipTransferRequested)
	if err := _FeeManager.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeManagerOwnershipTransferredIterator struct {
	Event *FeeManagerOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeManagerOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerOwnershipTransferred)
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
		it.Event = new(FeeManagerOwnershipTransferred)
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

func (it *FeeManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *FeeManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeManagerOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FeeManager *FeeManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeManagerOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeeManagerOwnershipTransferredIterator{contract: _FeeManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_FeeManager *FeeManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeManagerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeManagerOwnershipTransferred)
				if err := _FeeManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_FeeManager *FeeManagerFilterer) ParseOwnershipTransferred(log types.Log) (*FeeManagerOwnershipTransferred, error) {
	event := new(FeeManagerOwnershipTransferred)
	if err := _FeeManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_FeeManager *FeeManager) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _FeeManager.abi.Events["FeeUpdaterRemoved"].ID:
		return _FeeManager.ParseFeeUpdaterRemoved(log)
	case _FeeManager.abi.Events["FeeUpdaterSet"].ID:
		return _FeeManager.ParseFeeUpdaterSet(log)
	case _FeeManager.abi.Events["GasFeeUpdated"].ID:
		return _FeeManager.ParseGasFeeUpdated(log)
	case _FeeManager.abi.Events["OwnershipTransferRequested"].ID:
		return _FeeManager.ParseOwnershipTransferRequested(log)
	case _FeeManager.abi.Events["OwnershipTransferred"].ID:
		return _FeeManager.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (FeeManagerFeeUpdaterRemoved) Topic() common.Hash {
	return common.HexToHash("0x74a2c31badb27f0acfb9da3ef34c9e656ca1723881466e89a40f791f1c82ee71")
}

func (FeeManagerFeeUpdaterSet) Topic() common.Hash {
	return common.HexToHash("0xa462a3423511fce5ad038a1130251b66e3b8c135fa2ca15777f45c72ea3037c5")
}

func (FeeManagerGasFeeUpdated) Topic() common.Hash {
	return common.HexToHash("0xb230bad3704091781ab962bc58267145fa2aa1c542698c04e10e9db6069fa2d3")
}

func (FeeManagerOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (FeeManagerOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_FeeManager *FeeManager) Address() common.Address {
	return _FeeManager.address
}

type FeeManagerInterface interface {
	GetFeeTokenBaseUnitsPerUnitGas(opts *bind.CallOpts, token common.Address, destChainId uint64) (*big.Int, error)

	GetStalenessThreshold(opts *bind.CallOpts) (uint32, error)

	IsFeeUpdater(opts *bind.CallOpts, feeUpdater common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error)

	SetFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateFees(opts *bind.TransactOpts, feeUpdates []InternalFeeUpdate) (*types.Transaction, error)

	WithdrawToken(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	FilterFeeUpdaterRemoved(opts *bind.FilterOpts, feeUpdater []common.Address) (*FeeManagerFeeUpdaterRemovedIterator, error)

	WatchFeeUpdaterRemoved(opts *bind.WatchOpts, sink chan<- *FeeManagerFeeUpdaterRemoved, feeUpdater []common.Address) (event.Subscription, error)

	ParseFeeUpdaterRemoved(log types.Log) (*FeeManagerFeeUpdaterRemoved, error)

	FilterFeeUpdaterSet(opts *bind.FilterOpts, feeUpdater []common.Address) (*FeeManagerFeeUpdaterSetIterator, error)

	WatchFeeUpdaterSet(opts *bind.WatchOpts, sink chan<- *FeeManagerFeeUpdaterSet, feeUpdater []common.Address) (event.Subscription, error)

	ParseFeeUpdaterSet(log types.Log) (*FeeManagerFeeUpdaterSet, error)

	FilterGasFeeUpdated(opts *bind.FilterOpts, token []common.Address, destChain []uint64) (*FeeManagerGasFeeUpdatedIterator, error)

	WatchGasFeeUpdated(opts *bind.WatchOpts, sink chan<- *FeeManagerGasFeeUpdated, token []common.Address, destChain []uint64) (event.Subscription, error)

	ParseGasFeeUpdated(log types.Log) (*FeeManagerGasFeeUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeManagerOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FeeManagerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*FeeManagerOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeManagerOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeManagerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*FeeManagerOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
