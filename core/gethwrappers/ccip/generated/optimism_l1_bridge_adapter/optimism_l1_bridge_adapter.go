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

type IL1CrossDomainMessengerL2MessageInclusionProof struct {
	StateRoot            [32]byte
	StateRootBatchHeader LibOVMCodecChainBatchHeader
	StateRootProof       LibOVMCodecChainInclusionProof
	StateTrieWitness     []byte
	StorageTrieWitness   []byte
}

type LibOVMCodecChainBatchHeader struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	ExtraData         []byte
}

type LibOVMCodecChainInclusionProof struct {
	Index    *big.Int
	Siblings [][32]byte
}

var OptimismL1BridgeAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIL1StandardBridge\",\"name\":\"l1Bridge\",\"type\":\"address\"},{\"internalType\":\"contractIWrappedNative\",\"name\":\"wrappedNative\",\"type\":\"address\"},{\"internalType\":\"contractIL1CrossDomainMessenger\",\"name\":\"l1CrossDomainMessenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositNativeToL2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20FromL2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawNativeFromL2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedNative\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_OVMCodec.ChainBatchHeader\",\"name\":\"stateRootBatchHeader\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLib_OVMCodec.ChainInclusionProof\",\"name\":\"stateRootProof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"stateTrieWitness\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"storageTrieWitness\",\"type\":\"bytes\"}],\"internalType\":\"structIL1CrossDomainMessenger.L2MessageInclusionProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"relayMessageFromL2ToL1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60e0604052600080546001600160401b03191690553480156200002157600080fd5b5060405162001609380380620016098339810160408190526200004491620000b7565b6001600160a01b03831615806200006257506001600160a01b038216155b156200008157604051635e9c404d60e11b815260040160405180910390fd5b6001600160a01b03928316608052821660a0521660c0526200010b565b6001600160a01b0381168114620000b457600080fd5b50565b600080600060608486031215620000cd57600080fd5b8351620000da816200009e565b6020850151909350620000ed816200009e565b604085015190925062000100816200009e565b809150509250925092565b60805160a05160c0516114a5620001646000396000818160e10152818161022001526102a0015260006107090152600081816101860152818161035d015281816103f50152818161054a015261068d01526114a56000f3fe6080604052600436106100655760003560e01c8063cb7a6e1611610043578063cb7a6e16146100b2578063e861e907146100d2578063f2bfa1e11461012957600080fd5b806318b3050c1461006a57806379a35b4b1461008c5780638b2e4a2c1461009f575b600080fd5b34801561007657600080fd5b5061008a610085366004610b49565b610149565b005b61008a61009a366004610bb8565b6101fc565b61008a6100ad366004610c03565b6104e8565b3480156100be57600080fd5b5061008a6100cd366004610c2d565b610633565b3480156100de57600080fd5b507f000000000000000000000000000000000000000000000000000000000000000060405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34801561013557600080fd5b5061008a610144366004610f12565b6106cc565b6040517f1532ec3400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690631532ec34906101c39088908890889088908890600401611090565b600060405180830381600087803b1580156101dd57600080fd5b505af11580156101f1573d6000803e3d6000fd5b505050505050505050565b61021e73ffffffffffffffffffffffffffffffffffffffff8516333084610746565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610320576040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632e1a7d4d90602401600060405180830381600087803b1580156102f957600080fd5b505af115801561030d573d6000803e3d6000fd5b5050505061031b82826104e8565b6104e2565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660048301526024820183905285169063095ea7b3906044016020604051808303816000875af11580156103b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d991906110d0565b506000805473ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163838b252091879187918791879167ffffffffffffffff168180610439836110f9565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060405160200161047f919067ffffffffffffffff91909116815260200190565b6040516020818303038152906040526040518763ffffffff1660e01b81526004016104af969594939291906111b5565b600060405180830381600087803b1580156104c957600080fd5b505af11580156104dd573d6000803e3d6000fd5b505050505b50505050565b80341461052f576040517f03da4d23000000000000000000000000000000000000000000000000000000008152346004820152602481018290526044015b60405180910390fd5b6000805473ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691639a2ac6d5913491869167ffffffffffffffff16818061058a836110f9565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040516020016105d0919067ffffffffffffffff91909116815260200190565b6040516020818303038152906040526040518563ffffffff1660e01b81526004016105fd93929190611214565b6000604051808303818588803b15801561061657600080fd5b505af115801561062a573d6000803e3d6000fd5b50505050505050565b600061064182840184611258565b8051602082015160408084015190517fa9f9e67500000000000000000000000000000000000000000000000000000000815293945073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169363a9f9e675936101c393909290918b918b918b908b906004016112c0565b6040517fd7fd19dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d7fd19dd906101c39088908890889088908890600401611374565b6040805173ffffffffffffffffffffffffffffffffffffffff8581166024830152848116604483015260648083018590528351808403909101815260849092018352602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564908401526104e2928792916000916108199185169084906108c8565b8051909150156108c3578080602001905181019061083791906110d0565b6108c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610526565b505050565b60606108d784846000856108df565b949350505050565b606082471015610971576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610526565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161099a9190611469565b60006040518083038185875af1925050503d80600081146109d7576040519150601f19603f3d011682016040523d82523d6000602084013e6109dc565b606091505b50915091506109ed878383876109f8565b979650505050505050565b60608315610a8e578251600003610a875773ffffffffffffffffffffffffffffffffffffffff85163b610a87576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610526565b50816108d7565b6108d78383815115610aa35781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105269190611485565b803573ffffffffffffffffffffffffffffffffffffffff81168114610afb57600080fd5b919050565b60008083601f840112610b1257600080fd5b50813567ffffffffffffffff811115610b2a57600080fd5b602083019150836020828501011115610b4257600080fd5b9250929050565b600080600080600060808688031215610b6157600080fd5b610b6a86610ad7565b9450610b7860208701610ad7565b935060408601359250606086013567ffffffffffffffff811115610b9b57600080fd5b610ba788828901610b00565b969995985093965092949392505050565b60008060008060808587031215610bce57600080fd5b610bd785610ad7565b9350610be560208601610ad7565b9250610bf360408601610ad7565b9396929550929360600135925050565b60008060408385031215610c1657600080fd5b610c1f83610ad7565b946020939093013593505050565b60008060008060608587031215610c4357600080fd5b610c4c85610ad7565b9350610c5a60208601610ad7565b9250604085013567ffffffffffffffff811115610c7657600080fd5b610c8287828801610b00565b95989497509550505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff81118282101715610ce057610ce0610c8e565b60405290565b6040805190810167ffffffffffffffff81118282101715610ce057610ce0610c8e565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610d5057610d50610c8e565b604052919050565b600082601f830112610d6957600080fd5b813567ffffffffffffffff811115610d8357610d83610c8e565b610db460207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610d09565b818152846020838601011115610dc957600080fd5b816020850160208301376000918101602001919091529392505050565b600060a08284031215610df857600080fd5b610e00610cbd565b905081358152602082013560208201526040820135604082015260608201356060820152608082013567ffffffffffffffff811115610e3e57600080fd5b610e4a84828501610d58565b60808301525092915050565b600060408284031215610e6857600080fd5b610e70610ce6565b90508135815260208083013567ffffffffffffffff80821115610e9257600080fd5b818501915085601f830112610ea657600080fd5b813581811115610eb857610eb8610c8e565b8060051b9150610ec9848301610d09565b8181529183018401918481019088841115610ee357600080fd5b938501935b83851015610f0157843582529385019390850190610ee8565b808688015250505050505092915050565b600080600080600060a08688031215610f2a57600080fd5b610f3386610ad7565b9450610f4160208701610ad7565b9350604086013567ffffffffffffffff80821115610f5e57600080fd5b610f6a89838a01610d58565b9450606088013593506080880135915080821115610f8757600080fd5b9087019060a0828a031215610f9b57600080fd5b610fa3610cbd565b82358152602083013582811115610fb957600080fd5b610fc58b828601610de6565b602083015250604083013582811115610fdd57600080fd5b610fe98b828601610e56565b60408301525060608301358281111561100157600080fd5b61100d8b828601610d58565b60608301525060808301358281111561102557600080fd5b6110318b828601610d58565b6080830152508093505050509295509295909350565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b600073ffffffffffffffffffffffffffffffffffffffff8088168352808716602084015250846040830152608060608301526109ed608083018486611047565b6000602082840312156110e257600080fd5b815180151581146110f257600080fd5b9392505050565b600067ffffffffffffffff80831681810361113d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6001019392505050565b60005b8381101561116257818101518382015260200161114a565b50506000910152565b60008151808452611183816020860160208601611147565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600073ffffffffffffffffffffffffffffffffffffffff8089168352808816602084015280871660408401525084606083015263ffffffff8416608083015260c060a083015261120860c083018461116b565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff8316602082015260606040820152600061124f606083018461116b565b95945050505050565b60006060828403121561126a57600080fd5b6040516060810181811067ffffffffffffffff8211171561128d5761128d610c8e565b60405261129983610ad7565b81526112a760208401610ad7565b6020820152604083013560408201528091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a083015261131060c083018486611047565b9998505050505050505050565b600060408301825184526020808401516040828701528281518085526060880191508383019450600092505b808310156113695784518252938301936001929092019190830190611349565b509695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525060a060408301526113ad60a083018661116b565b846060840152828103608084015283518152602084015160a06020830152805160a0830152602081015160c0830152604081015160e083015260608101516101008301526080810151905060a061012083015261140e61014083018261116b565b905060408501518282036040840152611427828261131d565b91505060608501518282036060840152611441828261116b565b9150506080850151828203608084015261145b828261116b565b9a9950505050505050505050565b6000825161147b818460208701611147565b9190910192915050565b6020815260006110f2602083018461116b56fea164736f6c6343000813000a",
}

var OptimismL1BridgeAdapterABI = OptimismL1BridgeAdapterMetaData.ABI

var OptimismL1BridgeAdapterBin = OptimismL1BridgeAdapterMetaData.Bin

func DeployOptimismL1BridgeAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, l1Bridge common.Address, wrappedNative common.Address, l1CrossDomainMessenger common.Address) (common.Address, *types.Transaction, *OptimismL1BridgeAdapter, error) {
	parsed, err := OptimismL1BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptimismL1BridgeAdapterBin), backend, l1Bridge, wrappedNative, l1CrossDomainMessenger)
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

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactor) DepositNativeToL2(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.contract.Transact(opts, "depositNativeToL2", recipient, amount)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterSession) DepositNativeToL2(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.DepositNativeToL2(&_OptimismL1BridgeAdapter.TransactOpts, recipient, amount)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactorSession) DepositNativeToL2(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.DepositNativeToL2(&_OptimismL1BridgeAdapter.TransactOpts, recipient, amount)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactor) FinalizeWithdrawERC20FromL2(opts *bind.TransactOpts, from common.Address, to common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.contract.Transact(opts, "finalizeWithdrawERC20FromL2", from, to, data)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterSession) FinalizeWithdrawERC20FromL2(from common.Address, to common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.FinalizeWithdrawERC20FromL2(&_OptimismL1BridgeAdapter.TransactOpts, from, to, data)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactorSession) FinalizeWithdrawERC20FromL2(from common.Address, to common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.FinalizeWithdrawERC20FromL2(&_OptimismL1BridgeAdapter.TransactOpts, from, to, data)
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

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactor) RelayMessageFromL2ToL1(opts *bind.TransactOpts, target common.Address, sender common.Address, message []byte, messageNonce *big.Int, proof IL1CrossDomainMessengerL2MessageInclusionProof) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.contract.Transact(opts, "relayMessageFromL2ToL1", target, sender, message, messageNonce, proof)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterSession) RelayMessageFromL2ToL1(target common.Address, sender common.Address, message []byte, messageNonce *big.Int, proof IL1CrossDomainMessengerL2MessageInclusionProof) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.RelayMessageFromL2ToL1(&_OptimismL1BridgeAdapter.TransactOpts, target, sender, message, messageNonce, proof)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactorSession) RelayMessageFromL2ToL1(target common.Address, sender common.Address, message []byte, messageNonce *big.Int, proof IL1CrossDomainMessengerL2MessageInclusionProof) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.RelayMessageFromL2ToL1(&_OptimismL1BridgeAdapter.TransactOpts, target, sender, message, messageNonce, proof)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactor) SendERC20(opts *bind.TransactOpts, l1Token common.Address, l2Token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.contract.Transact(opts, "sendERC20", l1Token, l2Token, recipient, amount)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterSession) SendERC20(l1Token common.Address, l2Token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.SendERC20(&_OptimismL1BridgeAdapter.TransactOpts, l1Token, l2Token, recipient, amount)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapterTransactorSession) SendERC20(l1Token common.Address, l2Token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OptimismL1BridgeAdapter.Contract.SendERC20(&_OptimismL1BridgeAdapter.TransactOpts, l1Token, l2Token, recipient, amount)
}

func (_OptimismL1BridgeAdapter *OptimismL1BridgeAdapter) Address() common.Address {
	return _OptimismL1BridgeAdapter.address
}

type OptimismL1BridgeAdapterInterface interface {
	GetWrappedNative(opts *bind.CallOpts) (common.Address, error)

	DepositNativeToL2(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	FinalizeWithdrawERC20FromL2(opts *bind.TransactOpts, from common.Address, to common.Address, data []byte) (*types.Transaction, error)

	FinalizeWithdrawNativeFromL2(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error)

	RelayMessageFromL2ToL1(opts *bind.TransactOpts, target common.Address, sender common.Address, message []byte, messageNonce *big.Int, proof IL1CrossDomainMessengerL2MessageInclusionProof) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, l1Token common.Address, l2Token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	Address() common.Address
}
