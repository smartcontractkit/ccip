// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum_l1_bridge_adapter

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

type ArbitrumL1BridgeAdapterArbitrumFinalizationPayload struct {
	Proof       [][32]byte
	Index       *big.Int
	L2Sender    common.Address
	To          common.Address
	L2Block     *big.Int
	L1Block     *big.Int
	L2Timestamp *big.Int
	Value       *big.Int
	Data        []byte
}

type ArbitrumL1BridgeAdapterSendERC20Params struct {
	GasLimit          *big.Int
	MaxSubmissionCost *big.Int
	MaxFeePerGas      *big.Int
}

var ArbitrumL1BridgeAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIL1GatewayRouter\",\"name\":\"l1GatewayRouter\",\"type\":\"address\"},{\"internalType\":\"contractIOutbox\",\"name\":\"l1Outbox\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MsgShouldNotContainValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NoGatewayForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GAS_PRICE_BID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SUBMISSION_COST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structArbitrumL1BridgeAdapter.ArbitrumFinalizationPayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"exposeArbitrumFinalizationPayload\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"}],\"internalType\":\"structArbitrumL1BridgeAdapter.SendERC20Params\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exposeSendERC20Params\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arbitrumFinalizationPayload\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeFeeInNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"}],\"name\":\"getL2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"bridgeSpecificPayload\",\"type\":\"bytes\"}],\"name\":\"sendERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60c0604052600080546001600160401b031916905534801561002057600080fd5b50604051620014bb380380620014bb833981016040819052610041916100ab565b6001600160a01b038216158061005e57506001600160a01b038116155b1561007c57604051635e9c404d60e11b815260040160405180910390fd5b6001600160a01b039182166080521660a0526100e5565b6001600160a01b03811681146100a857600080fd5b50565b600080604083850312156100be57600080fd5b82516100c981610093565b60208401519092506100da81610093565b809150509250929050565b60805160a0516113a262000119600039600061026f015260008181610351015281816104c7015261062201526113a26000f3fe6080604052600436106100965760003560e01c8063a71d98b711610069578063c7665dd21161004e578063c7665dd214610152578063c985069c1461016d578063ea6c2f80146101b257600080fd5b8063a71d98b714610114578063b5399c9e1461013457600080fd5b80632e4b1fc91461009b57806332eb7905146100c357806338314bb2146100db5780635f2a9f41146100fd575b600080fd5b3480156100a757600080fd5b506100b06101cd565b6040519081526020015b60405180910390f35b3480156100cf57600080fd5b506100b06311e1a30081565b3480156100e757600080fd5b506100fb6100f6366004610c94565b6101f6565b005b34801561010957600080fd5b506100b0620186a081565b610127610122366004610cf9565b6102e5565b6040516100ba9190610dec565b34801561014057600080fd5b506100fb61014f366004610eae565b50565b34801561015e57600080fd5b506100fb61014f36600461101b565b34801561017957600080fd5b5061018d610188366004611105565b6105da565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ba565b3480156101be57600080fd5b506100b06602d79883d2000081565b60006101e06311e1a300620186a0611151565b6101f1906602d79883d20000611168565b905090565b60006102048284018461101b565b805160208201516040808401516060850151608086015160a087015160c088015160e08901516101008a015196517f08635a95000000000000000000000000000000000000000000000000000000008152999a5073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016996308635a95996102ac999098909796959493929160040161117b565b600060405180830381600087803b1580156102c657600080fd5b505af11580156102da573d6000803e3d6000fd5b505050505050505050565b606061030973ffffffffffffffffffffffffffffffffffffffff8816333087610695565b6040517fbda009fe00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff88811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063bda009fe90602401602060405180830381865afa15801561039a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103be9190611231565b905073ffffffffffffffffffffffffffffffffffffffff811661042a576040517f6c1460f400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff891660048201526024015b60405180910390fd5b61044b73ffffffffffffffffffffffffffffffffffffffff89168287610777565b600061045984860186610eae565b905060008160200151826040015183600001516104769190611151565b6104809190611168565b9050803410156104c5576040517f03da4d2300000000000000000000000000000000000000000000000000000000815234600482015260248101829052604401610421565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16634fb1a07b348c8b8c8c886000015189604001518a602001516040518060200160405280600081525060405160200161053692919061124e565b6040516020818303038152906040526040518963ffffffff1660e01b81526004016105679796959493929190611267565b60006040518083038185885af1158015610585573d6000803e3d6000fd5b50505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526105cc91908101906112c7565b9a9950505050505050505050565b6040517fa7e28d4800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063a7e28d4890602401602060405180830381865afa15801561066b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061068f9190611231565b92915050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526107719085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526108fe565b50505050565b80158061081757506040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff838116602483015284169063dd62ed3e90604401602060405180830381865afa1580156107f1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610815919061133e565b155b6108a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608401610421565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526108f99084907f095ea7b300000000000000000000000000000000000000000000000000000000906064016106ef565b505050565b6000610960826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16610a0a9092919063ffffffff16565b8051909150156108f9578080602001905181019061097e9190611357565b6108f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610421565b6060610a198484600085610a21565b949350505050565b606082471015610ab3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610421565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610adc9190611379565b60006040518083038185875af1925050503d8060008114610b19576040519150601f19603f3d011682016040523d82523d6000602084013e610b1e565b606091505b5091509150610b2f87838387610b3a565b979650505050505050565b60608315610bd0578251600003610bc95773ffffffffffffffffffffffffffffffffffffffff85163b610bc9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610421565b5081610a19565b610a198383815115610be55781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104219190610dec565b73ffffffffffffffffffffffffffffffffffffffff8116811461014f57600080fd5b8035610c4681610c19565b919050565b60008083601f840112610c5d57600080fd5b50813567ffffffffffffffff811115610c7557600080fd5b602083019150836020828501011115610c8d57600080fd5b9250929050565b60008060008060608587031215610caa57600080fd5b8435610cb581610c19565b93506020850135610cc581610c19565b9250604085013567ffffffffffffffff811115610ce157600080fd5b610ced87828801610c4b565b95989497509550505050565b60008060008060008060a08789031215610d1257600080fd5b8635610d1d81610c19565b95506020870135610d2d81610c19565b94506040870135610d3d81610c19565b935060608701359250608087013567ffffffffffffffff811115610d6057600080fd5b610d6c89828a01610c4b565b979a9699509497509295939492505050565b60005b83811015610d99578181015183820152602001610d81565b50506000910152565b60008151808452610dba816020860160208601610d7e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610dff6020830184610da2565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff81118282101715610e5957610e59610e06565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610ea657610ea6610e06565b604052919050565b600060608284031215610ec057600080fd5b6040516060810181811067ffffffffffffffff82111715610ee357610ee3610e06565b80604052508235815260208301356020820152604083013560408201528091505092915050565b600082601f830112610f1b57600080fd5b8135602067ffffffffffffffff821115610f3757610f37610e06565b8160051b610f46828201610e5f565b9283528481018201928281019087851115610f6057600080fd5b83870192505b84831015610b2f57823582529183019190830190610f66565b600067ffffffffffffffff821115610f9957610f99610e06565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112610fd657600080fd5b8135610fe9610fe482610f7f565b610e5f565b818152846020838601011115610ffe57600080fd5b816020850160208301376000918101602001919091529392505050565b60006020828403121561102d57600080fd5b813567ffffffffffffffff8082111561104557600080fd5b90830190610120828603121561105a57600080fd5b611062610e35565b82358281111561107157600080fd5b61107d87828601610f0a565b8252506020830135602082015261109660408401610c3b565b60408201526110a760608401610c3b565b60608201526080830135608082015260a083013560a082015260c083013560c082015260e083013560e082015261010080840135838111156110e857600080fd5b6110f488828701610fc5565b918301919091525095945050505050565b60006020828403121561111757600080fd5b8135610dff81610c19565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808202811582820484141761068f5761068f611122565b8082018082111561068f5761068f611122565b6101208082528a51908201819052600090610140830190602090818e01845b828110156111b65781518552938301939083019060010161119a565b50505083018b905273ffffffffffffffffffffffffffffffffffffffff8a16604084015273ffffffffffffffffffffffffffffffffffffffff891660608401528760808401528660a08401528560c08401528460e08401528281036101008401526112218185610da2565b9c9b505050505050505050505050565b60006020828403121561124357600080fd5b8151610dff81610c19565b828152604060208201526000610a196040830184610da2565b600073ffffffffffffffffffffffffffffffffffffffff808a16835280891660208401528088166040840152508560608301528460808301528360a083015260e060c08301526112ba60e0830184610da2565b9998505050505050505050565b6000602082840312156112d957600080fd5b815167ffffffffffffffff8111156112f057600080fd5b8201601f8101841361130157600080fd5b805161130f610fe482610f7f565b81815285602083850101111561132457600080fd5b611335826020830160208601610d7e565b95945050505050565b60006020828403121561135057600080fd5b5051919050565b60006020828403121561136957600080fd5b81518015158114610dff57600080fd5b6000825161138b818460208701610d7e565b919091019291505056fea164736f6c6343000813000a",
}

var ArbitrumL1BridgeAdapterABI = ArbitrumL1BridgeAdapterMetaData.ABI

var ArbitrumL1BridgeAdapterBin = ArbitrumL1BridgeAdapterMetaData.Bin

func DeployArbitrumL1BridgeAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, l1GatewayRouter common.Address, l1Outbox common.Address) (common.Address, *types.Transaction, *ArbitrumL1BridgeAdapter, error) {
	parsed, err := ArbitrumL1BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArbitrumL1BridgeAdapterBin), backend, l1GatewayRouter, l1Outbox)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbitrumL1BridgeAdapter{address: address, abi: *parsed, ArbitrumL1BridgeAdapterCaller: ArbitrumL1BridgeAdapterCaller{contract: contract}, ArbitrumL1BridgeAdapterTransactor: ArbitrumL1BridgeAdapterTransactor{contract: contract}, ArbitrumL1BridgeAdapterFilterer: ArbitrumL1BridgeAdapterFilterer{contract: contract}}, nil
}

type ArbitrumL1BridgeAdapter struct {
	address common.Address
	abi     abi.ABI
	ArbitrumL1BridgeAdapterCaller
	ArbitrumL1BridgeAdapterTransactor
	ArbitrumL1BridgeAdapterFilterer
}

type ArbitrumL1BridgeAdapterCaller struct {
	contract *bind.BoundContract
}

type ArbitrumL1BridgeAdapterTransactor struct {
	contract *bind.BoundContract
}

type ArbitrumL1BridgeAdapterFilterer struct {
	contract *bind.BoundContract
}

type ArbitrumL1BridgeAdapterSession struct {
	Contract     *ArbitrumL1BridgeAdapter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ArbitrumL1BridgeAdapterCallerSession struct {
	Contract *ArbitrumL1BridgeAdapterCaller
	CallOpts bind.CallOpts
}

type ArbitrumL1BridgeAdapterTransactorSession struct {
	Contract     *ArbitrumL1BridgeAdapterTransactor
	TransactOpts bind.TransactOpts
}

type ArbitrumL1BridgeAdapterRaw struct {
	Contract *ArbitrumL1BridgeAdapter
}

type ArbitrumL1BridgeAdapterCallerRaw struct {
	Contract *ArbitrumL1BridgeAdapterCaller
}

type ArbitrumL1BridgeAdapterTransactorRaw struct {
	Contract *ArbitrumL1BridgeAdapterTransactor
}

func NewArbitrumL1BridgeAdapter(address common.Address, backend bind.ContractBackend) (*ArbitrumL1BridgeAdapter, error) {
	abi, err := abi.JSON(strings.NewReader(ArbitrumL1BridgeAdapterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindArbitrumL1BridgeAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapter{address: address, abi: abi, ArbitrumL1BridgeAdapterCaller: ArbitrumL1BridgeAdapterCaller{contract: contract}, ArbitrumL1BridgeAdapterTransactor: ArbitrumL1BridgeAdapterTransactor{contract: contract}, ArbitrumL1BridgeAdapterFilterer: ArbitrumL1BridgeAdapterFilterer{contract: contract}}, nil
}

func NewArbitrumL1BridgeAdapterCaller(address common.Address, caller bind.ContractCaller) (*ArbitrumL1BridgeAdapterCaller, error) {
	contract, err := bindArbitrumL1BridgeAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapterCaller{contract: contract}, nil
}

func NewArbitrumL1BridgeAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrumL1BridgeAdapterTransactor, error) {
	contract, err := bindArbitrumL1BridgeAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapterTransactor{contract: contract}, nil
}

func NewArbitrumL1BridgeAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrumL1BridgeAdapterFilterer, error) {
	contract, err := bindArbitrumL1BridgeAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapterFilterer{contract: contract}, nil
}

func bindArbitrumL1BridgeAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbitrumL1BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumL1BridgeAdapter.Contract.ArbitrumL1BridgeAdapterCaller.contract.Call(opts, result, method, params...)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.ArbitrumL1BridgeAdapterTransactor.contract.Transfer(opts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.ArbitrumL1BridgeAdapterTransactor.contract.Transact(opts, method, params...)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumL1BridgeAdapter.Contract.contract.Call(opts, result, method, params...)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.contract.Transfer(opts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.contract.Transact(opts, method, params...)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) GASPRICEBID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "GAS_PRICE_BID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) GASPRICEBID() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GASPRICEBID(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) GASPRICEBID() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GASPRICEBID(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) MAXGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "MAX_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) MAXGAS() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.MAXGAS(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) MAXGAS() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.MAXGAS(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) MAXSUBMISSIONCOST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "MAX_SUBMISSION_COST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) MAXSUBMISSIONCOST() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.MAXSUBMISSIONCOST(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) MAXSUBMISSIONCOST() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.MAXSUBMISSIONCOST(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) ExposeArbitrumFinalizationPayload(opts *bind.CallOpts, payload ArbitrumL1BridgeAdapterArbitrumFinalizationPayload) error {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "exposeArbitrumFinalizationPayload", payload)

	if err != nil {
		return err
	}

	return err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) ExposeArbitrumFinalizationPayload(payload ArbitrumL1BridgeAdapterArbitrumFinalizationPayload) error {
	return _ArbitrumL1BridgeAdapter.Contract.ExposeArbitrumFinalizationPayload(&_ArbitrumL1BridgeAdapter.CallOpts, payload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) ExposeArbitrumFinalizationPayload(payload ArbitrumL1BridgeAdapterArbitrumFinalizationPayload) error {
	return _ArbitrumL1BridgeAdapter.Contract.ExposeArbitrumFinalizationPayload(&_ArbitrumL1BridgeAdapter.CallOpts, payload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) ExposeSendERC20Params(opts *bind.CallOpts, params ArbitrumL1BridgeAdapterSendERC20Params) error {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "exposeSendERC20Params", params)

	if err != nil {
		return err
	}

	return err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) ExposeSendERC20Params(params ArbitrumL1BridgeAdapterSendERC20Params) error {
	return _ArbitrumL1BridgeAdapter.Contract.ExposeSendERC20Params(&_ArbitrumL1BridgeAdapter.CallOpts, params)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) ExposeSendERC20Params(params ArbitrumL1BridgeAdapterSendERC20Params) error {
	return _ArbitrumL1BridgeAdapter.Contract.ExposeSendERC20Params(&_ArbitrumL1BridgeAdapter.CallOpts, params)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "getBridgeFeeInNative")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GetBridgeFeeInNative(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GetBridgeFeeInNative(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) GetL2Token(opts *bind.CallOpts, l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "getL2Token", l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) GetL2Token(l1Token common.Address) (common.Address, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GetL2Token(&_ArbitrumL1BridgeAdapter.CallOpts, l1Token)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) GetL2Token(l1Token common.Address) (common.Address, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GetL2Token(&_ArbitrumL1BridgeAdapter.CallOpts, l1Token)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.contract.Transact(opts, "finalizeWithdrawERC20", arg0, arg1, arbitrumFinalizationPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.FinalizeWithdrawERC20(&_ArbitrumL1BridgeAdapter.TransactOpts, arg0, arg1, arbitrumFinalizationPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactorSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.FinalizeWithdrawERC20(&_ArbitrumL1BridgeAdapter.TransactOpts, arg0, arg1, arbitrumFinalizationPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactor) SendERC20(opts *bind.TransactOpts, localToken common.Address, arg1 common.Address, recipient common.Address, amount *big.Int, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.contract.Transact(opts, "sendERC20", localToken, arg1, recipient, amount, bridgeSpecificPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) SendERC20(localToken common.Address, arg1 common.Address, recipient common.Address, amount *big.Int, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.SendERC20(&_ArbitrumL1BridgeAdapter.TransactOpts, localToken, arg1, recipient, amount, bridgeSpecificPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactorSession) SendERC20(localToken common.Address, arg1 common.Address, recipient common.Address, amount *big.Int, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.SendERC20(&_ArbitrumL1BridgeAdapter.TransactOpts, localToken, arg1, recipient, amount, bridgeSpecificPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapter) Address() common.Address {
	return _ArbitrumL1BridgeAdapter.address
}

type ArbitrumL1BridgeAdapterInterface interface {
	GASPRICEBID(opts *bind.CallOpts) (*big.Int, error)

	MAXGAS(opts *bind.CallOpts) (*big.Int, error)

	MAXSUBMISSIONCOST(opts *bind.CallOpts) (*big.Int, error)

	ExposeArbitrumFinalizationPayload(opts *bind.CallOpts, payload ArbitrumL1BridgeAdapterArbitrumFinalizationPayload) error

	ExposeSendERC20Params(opts *bind.CallOpts, params ArbitrumL1BridgeAdapterSendERC20Params) error

	GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error)

	GetL2Token(opts *bind.CallOpts, l1Token common.Address) (common.Address, error)

	FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, localToken common.Address, arg1 common.Address, recipient common.Address, amount *big.Int, bridgeSpecificPayload []byte) (*types.Transaction, error)

	Address() common.Address
}
