// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimism_l1_bridge_adapter

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

var OptimismL1BridgeAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIL1StandardBridge\",\"name\":\"l1Bridge\",\"type\":\"address\"},{\"internalType\":\"contractIWrappedNative\",\"name\":\"wrappedNative\",\"type\":\"address\"},{\"internalType\":\"contractIOptimismPortal\",\"name\":\"optimismPortal\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFinalizationAction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MsgShouldNotContainValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawNativeFromL2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeFeeInNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL1Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOptimismPortal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedNative\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e0604052600080546001600160401b03191690553480156200002157600080fd5b5060405162001935380380620019358339810160408190526200004491620000cb565b6001600160a01b03831615806200006257506001600160a01b038216155b806200007557506001600160a01b038116155b156200009457604051635e9c404d60e11b815260040160405180910390fd5b6001600160a01b0392831660805290821660a0521660c0526200011f565b6001600160a01b0381168114620000c857600080fd5b50565b600080600060608486031215620000e157600080fd5b8351620000ee81620000b2565b60208501519093506200010181620000b2565b60408501519092506200011481620000b2565b809150509250925092565b60805160a05160c0516117af620001866000396000818160f2015281816106b70152610791015260008181610199015281816104120152610492015260008181610166015281816101fa0152818161053e015281816105d3015261063501526117af6000f3fe6080604052600436106100745760003560e01c806354fd969f1161004e57806354fd969f146100e3578063a71d98b714610137578063c86d5bdd14610157578063e861e9071461018a57600080fd5b806318b3050c146100805780632e4b1fc9146100a257806338314bb2146100c357600080fd5b3661007b57005b600080fd5b34801561008c57600080fd5b506100a061009b366004610db7565b6101bd565b005b3480156100ae57600080fd5b50604051600081526020015b60405180910390f35b3480156100cf57600080fd5b506100a06100de366004610e2a565b610270565b3480156100ef57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ba565b61014a610145366004610e8f565b61033d565b6040516100ba9190610f82565b34801561016357600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610112565b34801561019657600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610112565b6040517f1532ec3400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690631532ec34906102379088908890889088908890600401610f9c565b600060405180830381600087803b15801561025157600080fd5b505af1158015610265573d6000803e3d6000fd5b505050505050505050565b600061027e8284018461112b565b9050600181516002811115610295576102956111f1565b036102c557600081602001518060200190518101906102b491906113b3565b90506102bf816106b5565b50610336565b6002815160028111156102da576102da6111f1565b0361030457600081602001518060200190518101906102f991906114b1565b90506102bf81610752565b6040517fee2ef09800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b606061036173ffffffffffffffffffffffffffffffffffffffff88163330876107c5565b34156103a0576040517f2543d86e0000000000000000000000000000000000000000000000000000000081523460048201526024015b60405180910390fd5b6000805467ffffffffffffffff1681806103b983611534565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040516020016103ff919067ffffffffffffffff91909116815260200190565b60405160208183030381529060405290507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff16036105b7576040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018690527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632e1a7d4d90602401600060405180830381600087803b1580156104eb57600080fd5b505af11580156104ff573d6000803e3d6000fd5b50506040517f9a2ac6d500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169250639a2ac6d59150879061057c908a906000908790600401611582565b6000604051808303818588803b15801561059557600080fd5b505af11580156105a9573d6000803e3d6000fd5b5050505050809150506106ab565b6105f873ffffffffffffffffffffffffffffffffffffffff89167f0000000000000000000000000000000000000000000000000000000000000000876108a7565b6040517f838b252000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063838b252090610675908b908b908b908b9060009089906004016115c6565b600060405180830381600087803b15801561068f57600080fd5b505af11580156106a3573d6000803e3d6000fd5b509293505050505b9695505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16634870496f82600001518360200151846040015185606001516040518563ffffffff1660e01b8152600401610724949392919061167c565b600060405180830381600087803b15801561073e57600080fd5b505af1158015610336573d6000803e3d6000fd5b80516040517f8c3152e900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691638c3152e9916107249190600401611738565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526108a19085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610a2e565b50505050565b80158061094757506040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff838116602483015284169063dd62ed3e90604401602060405180830381865afa158015610921573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610945919061174b565b155b6109d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608401610397565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610a299084907f095ea7b3000000000000000000000000000000000000000000000000000000009060640161081f565b505050565b6000610a90826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16610b3a9092919063ffffffff16565b805190915015610a295780806020019051810190610aae9190611764565b610a29576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610397565b6060610b498484600085610b51565b949350505050565b606082471015610be3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610397565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610c0c9190611786565b60006040518083038185875af1925050503d8060008114610c49576040519150601f19603f3d011682016040523d82523d6000602084013e610c4e565b606091505b5091509150610c5f87838387610c6a565b979650505050505050565b60608315610d00578251600003610cf95773ffffffffffffffffffffffffffffffffffffffff85163b610cf9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610397565b5081610b49565b610b498383815115610d155781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103979190610f82565b73ffffffffffffffffffffffffffffffffffffffff81168114610d6b57600080fd5b50565b60008083601f840112610d8057600080fd5b50813567ffffffffffffffff811115610d9857600080fd5b602083019150836020828501011115610db057600080fd5b9250929050565b600080600080600060808688031215610dcf57600080fd5b8535610dda81610d49565b94506020860135610dea81610d49565b935060408601359250606086013567ffffffffffffffff811115610e0d57600080fd5b610e1988828901610d6e565b969995985093965092949392505050565b60008060008060608587031215610e4057600080fd5b8435610e4b81610d49565b93506020850135610e5b81610d49565b9250604085013567ffffffffffffffff811115610e7757600080fd5b610e8387828801610d6e565b95989497509550505050565b60008060008060008060a08789031215610ea857600080fd5b8635610eb381610d49565b95506020870135610ec381610d49565b94506040870135610ed381610d49565b935060608701359250608087013567ffffffffffffffff811115610ef657600080fd5b610f0289828a01610d6e565b979a9699509497509295939492505050565b60005b83811015610f2f578181015183820152602001610f17565b50506000910152565b60008151808452610f50816020860160208601610f14565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610f956020830184610f38565b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525084604083015260806060830152826080830152828460a0840137600060a0848401015260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85011683010190509695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561106d5761106d61101b565b60405290565b6040516080810167ffffffffffffffff8111828210171561106d5761106d61101b565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156110dd576110dd61101b565b604052919050565b600067ffffffffffffffff8211156110ff576110ff61101b565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6000602080838503121561113e57600080fd5b823567ffffffffffffffff8082111561115657600080fd5b908401906040828703121561116a57600080fd5b61117261104a565b82356003811061118157600080fd5b8152828401358281111561119457600080fd5b80840193505086601f8401126111a957600080fd5b823591506111be6111b9836110e5565b611096565b82815287858486010111156111d257600080fd5b8285850186830137600092810185019290925292830152509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600082601f83011261123157600080fd5b815161123f6111b9826110e5565b81815284602083860101111561125457600080fd5b610b49826020830160208701610f14565b600060c0828403121561127757600080fd5b60405160c0810167ffffffffffffffff828210818311171561129b5761129b61101b565b8160405282935084518352602085015191506112b682610d49565b816020840152604085015191506112cc82610d49565b816040840152606085015160608401526080850151608084015260a08501519150808211156112fa57600080fd5b5061130785828601611220565b60a0830152505092915050565b600082601f83011261132557600080fd5b8151602067ffffffffffffffff808311156113425761134261101b565b8260051b611351838201611096565b938452858101830193838101908886111561136b57600080fd5b84880192505b858310156113a7578251848111156113895760008081fd5b6113978a87838c0101611220565b8352509184019190840190611371565b98975050505050505050565b6000602082840312156113c557600080fd5b815167ffffffffffffffff808211156113dd57600080fd5b9083019081850360e08112156113f257600080fd5b6113fa611073565b83518381111561140957600080fd5b61141588828701611265565b8252506020840151602082015260807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08301121561145257600080fd5b61145a611073565b6040858101518252606080870151602084015260808701518284015260a08701519083015282015260c084015191508282111561149657600080fd5b6114a287838601611314565b60608201529695505050505050565b6000602082840312156114c357600080fd5b815167ffffffffffffffff808211156114db57600080fd5b90830190602082860312156114ef57600080fd5b60405160208101818110838211171561150a5761150a61101b565b60405282518281111561151c57600080fd5b61152887828601611265565b82525095945050505050565b600067ffffffffffffffff808316818103611578577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6001019392505050565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff831660208201526060604082015260006115bd6060830184610f38565b95945050505050565b600073ffffffffffffffffffffffffffffffffffffffff8089168352808816602084015280871660408401525084606083015263ffffffff8416608083015260c060a08301526113a760c0830184610f38565b805182526000602082015173ffffffffffffffffffffffffffffffffffffffff80821660208601528060408501511660408601525050606082015160608401526080820151608084015260a082015160c060a0850152610b4960c0850182610f38565b60e08152600061168f60e0830187611619565b602086818501528551604085015280860151606085015260408601516080850152606086015160a085015283820360c08501528185518084528284019150828160051b85010183880160005b83811015611727577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0878403018552611715838351610f38565b948601949250908501906001016116db565b50909b9a5050505050505050505050565b602081526000610f956020830184611619565b60006020828403121561175d57600080fd5b5051919050565b60006020828403121561177657600080fd5b81518015158114610f9557600080fd5b60008251611798818460208701610f14565b919091019291505056fea164736f6c6343000813000a",
}

var OptimismL1BridgeAdapterABI = OptimismL1BridgeAdapterMetaData.ABI

var OptimismL1BridgeAdapterBin = OptimismL1BridgeAdapterMetaData.Bin

func DeployOptimismL1BridgeAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, l1Bridge common.Address, wrappedNative common.Address, optimismPortal common.Address) (common.Address, *types.Transaction, *OptimismL1BridgeAdapter, error) {
	parsed, err := OptimismL1BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptimismL1BridgeAdapterBin), backend, l1Bridge, wrappedNative, optimismPortal)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptimismL1BridgeAdapter{address: address, abi: *parsed, OptimismL1BridgeAdapterCaller: OptimismL1BridgeAdapterCaller{contract: contract}, OptimismL1BridgeAdapterTransactor: OptimismL1BridgeAdapterTransactor{contract: contract}, OptimismL1BridgeAdapterFilterer: OptimismL1BridgeAdapterFilterer{contract: contract}}, nil
}

type OptimismL1BridgeAdapter struct {
	address common.Address
	abi     abi.ABI
	OptimismL1BridgeAdapterCaller
	OptimismL1BridgeAdapterTransactor
	OptimismL1BridgeAdapterFilterer
}

type OptimismL1BridgeAdapterCaller struct {
	contract *bind.BoundContract
}

type OptimismL1BridgeAdapterTransactor struct {
	contract *bind.BoundContract
}

type OptimismL1BridgeAdapterFilterer struct {
	contract *bind.BoundContract
}

type OptimismL1BridgeAdapterSession struct {
	Contract     *OptimismL1BridgeAdapter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OptimismL1BridgeAdapterCallerSession struct {
	Contract *OptimismL1BridgeAdapterCaller
	CallOpts bind.CallOpts
}

type OptimismL1BridgeAdapterTransactorSession struct {
	Contract     *OptimismL1BridgeAdapterTransactor
	TransactOpts bind.TransactOpts
}

type OptimismL1BridgeAdapterRaw struct {
	Contract *OptimismL1BridgeAdapter
}

type OptimismL1BridgeAdapterCallerRaw struct {
	Contract *OptimismL1BridgeAdapterCaller
}

type OptimismL1BridgeAdapterTransactorRaw struct {
	Contract *OptimismL1BridgeAdapterTransactor
}

func NewOptimismL1BridgeAdapter(address common.Address, backend bind.ContractBackend) (*OptimismL1BridgeAdapter, error) {
	abi, err := abi.JSON(strings.NewReader(OptimismL1BridgeAdapterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOptimismL1BridgeAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimismL1BridgeAdapter{address: address, abi: abi, OptimismL1BridgeAdapterCaller: OptimismL1BridgeAdapterCaller{contract: contract}, OptimismL1BridgeAdapterTransactor: OptimismL1BridgeAdapterTransactor{contract: contract}, OptimismL1BridgeAdapterFilterer: OptimismL1BridgeAdapterFilterer{contract: contract}}, nil
}

func NewOptimismL1BridgeAdapterCaller(address common.Address, caller bind.ContractCaller) (*OptimismL1BridgeAdapterCaller, error) {
	contract, err := bindOptimismL1BridgeAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismL1BridgeAdapterCaller{contract: contract}, nil
}

func NewOptimismL1BridgeAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimismL1BridgeAdapterTransactor, error) {
	contract, err := bindOptimismL1BridgeAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismL1BridgeAdapterTransactor{contract: contract}, nil
}

func NewOptimismL1BridgeAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimismL1BridgeAdapterFilterer, error) {
	contract, err := bindOptimismL1BridgeAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimismL1BridgeAdapterFilterer{contract: contract}, nil
}

func bindOptimismL1BridgeAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimismL1BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismL1BridgeAdapter.Contract.OptimismL1BridgeAdapterCaller.contract.Call(opts, result, method, params...)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.OptimismL1BridgeAdapterTransactor.contract.Transfer(opts)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.OptimismL1BridgeAdapterTransactor.contract.Transact(opts, method, params...)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismL1BridgeAdapter.Contract.contract.Call(opts, result, method, params...)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.contract.Transfer(opts)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.contract.Transact(opts, method, params...)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterCaller) GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OptimismL1BridgeAdapter.contract.Call(opts, &out, "getBridgeFeeInNative")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _OptimismL1BridgeAdapter.Contract.GetBridgeFeeInNative(&_OptimismL1BridgeAdapter.CallOpts)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterCallerSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _OptimismL1BridgeAdapter.Contract.GetBridgeFeeInNative(&_OptimismL1BridgeAdapter.CallOpts)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterCaller) GetL1Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismL1BridgeAdapter.contract.Call(opts, &out, "getL1Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterSession) GetL1Bridge() (common.Address, error) {
	return _OptimismL1BridgeAdapter.Contract.GetL1Bridge(&_OptimismL1BridgeAdapter.CallOpts)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterCallerSession) GetL1Bridge() (common.Address, error) {
	return _OptimismL1BridgeAdapter.Contract.GetL1Bridge(&_OptimismL1BridgeAdapter.CallOpts)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterCaller) GetOptimismPortal(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismL1BridgeAdapter.contract.Call(opts, &out, "getOptimismPortal")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterSession) GetOptimismPortal() (common.Address, error) {
	return _OptimismL1BridgeAdapter.Contract.GetOptimismPortal(&_OptimismL1BridgeAdapter.CallOpts)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterCallerSession) GetOptimismPortal() (common.Address, error) {
	return _OptimismL1BridgeAdapter.Contract.GetOptimismPortal(&_OptimismL1BridgeAdapter.CallOpts)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterCaller) GetWrappedNative(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismL1BridgeAdapter.contract.Call(opts, &out, "getWrappedNative")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterSession) GetWrappedNative() (common.Address, error) {
	return _OptimismL1BridgeAdapter.Contract.GetWrappedNative(&_OptimismL1BridgeAdapter.CallOpts)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterCallerSession) GetWrappedNative() (common.Address, error) {
	return _OptimismL1BridgeAdapter.Contract.GetWrappedNative(&_OptimismL1BridgeAdapter.CallOpts)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.contract.Transact(opts, "finalizeWithdrawERC20", arg0, arg1, data)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.FinalizeWithdrawERC20(&_OptimismL1BridgeAdapter.TransactOpts, arg0, arg1, data)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactorSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.FinalizeWithdrawERC20(&_OptimismL1BridgeAdapter.TransactOpts, arg0, arg1, data)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactor) FinalizeWithdrawNativeFromL2(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.contract.Transact(opts, "finalizeWithdrawNativeFromL2", from, to, amount, data)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterSession) FinalizeWithdrawNativeFromL2(from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.FinalizeWithdrawNativeFromL2(&_OptimismL1BridgeAdapter.TransactOpts, from, to, amount, data)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactorSession) FinalizeWithdrawNativeFromL2(from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.FinalizeWithdrawNativeFromL2(&_OptimismL1BridgeAdapter.TransactOpts, from, to, amount, data)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactor) SendERC20(opts *bind.TransactOpts, localToken common.Address, remoteToken common.Address, recipient common.Address, amount *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.contract.Transact(opts, "sendERC20", localToken, remoteToken, recipient, amount, arg4)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterSession) SendERC20(localToken common.Address, remoteToken common.Address, recipient common.Address, amount *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.SendERC20(&_OptimismL1BridgeAdapter.TransactOpts, localToken, remoteToken, recipient, amount, arg4)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactorSession) SendERC20(localToken common.Address, remoteToken common.Address, recipient common.Address, amount *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.SendERC20(&_OptimismL1BridgeAdapter.TransactOpts, localToken, remoteToken, recipient, amount, arg4)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.contract.RawTransact(opts, nil)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterSession) Receive() (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.Receive(&_OptimismL1BridgeAdapter.TransactOpts)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactorSession) Receive() (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.Receive(&_OptimismL1BridgeAdapter.TransactOpts)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapter) Address() common.Address {
	return _OptimismL1BridgeAdapter.address
}

type OptimismL1BridgeAdapterInterface interface {
	GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error)

	GetL1Bridge(opts *bind.CallOpts) (common.Address, error)

	GetOptimismPortal(opts *bind.CallOpts) (common.Address, error)

	GetWrappedNative(opts *bind.CallOpts) (common.Address, error)

	FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, data []byte) (*types.Transaction, error)

	FinalizeWithdrawNativeFromL2(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, localToken common.Address, remoteToken common.Address, recipient common.Address, amount *big.Int, arg4 []byte) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	Address() common.Address
}
