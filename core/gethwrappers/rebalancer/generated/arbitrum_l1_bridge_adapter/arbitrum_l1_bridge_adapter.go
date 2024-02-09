// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum_l1_bridge_adapter

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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIL1GatewayRouter\",\"name\":\"l1GatewayRouter\",\"type\":\"address\"},{\"internalType\":\"contractIOutbox\",\"name\":\"l1Outbox\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MsgShouldNotContainValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NoGatewayForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unimplemented\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"outboundTransferResult\",\"type\":\"bytes\"}],\"name\":\"ArbitrumL1ToL2ERC20Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structArbitrumL1BridgeAdapter.ArbitrumFinalizationPayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"ArbitrumL2ToL1ERC20Finalized\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structArbitrumL1BridgeAdapter.ArbitrumFinalizationPayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"exposeArbitrumFinalizationPayload\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"}],\"internalType\":\"structArbitrumL1BridgeAdapter.SendERC20Params\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exposeSendERC20Params\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arbitrumFinalizationPayload\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeFeeInNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"}],\"name\":\"getL2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"bridgeSpecificPayload\",\"type\":\"bytes\"}],\"name\":\"sendERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60c0604052600080546001600160401b03191690553480156200002157600080fd5b50604051620016a6380380620016a68339810160408190526200004491620000b2565b6001600160a01b03821615806200006257506001600160a01b038116155b156200008157604051635e9c404d60e11b815260040160405180910390fd5b6001600160a01b039182166080521660a052620000f1565b6001600160a01b0381168114620000af57600080fd5b50565b60008060408385031215620000c657600080fd5b8251620000d38162000099565b6020840151909250620000e68162000099565b809150509250929050565b60805160a0516115816200012560003960006101ff015260008181610354015281816104cc01526106c301526115816000f3fe6080604052600436106100655760003560e01c8063b5399c9e11610043578063b5399c9e146100d4578063c7665dd2146100f2578063c985069c1461010d57600080fd5b80632e4b1fc91461006a57806338314bb214610092578063a71d98b7146100b4575b600080fd5b34801561007657600080fd5b5061007f610152565b6040519081526020015b60405180910390f35b34801561009e57600080fd5b506100b26100ad366004610d35565b610186565b005b6100c76100c2366004610d9a565b6102e8565b6040516100899190610e8d565b3480156100e057600080fd5b506100b26100ef366004610f4f565b50565b3480156100fe57600080fd5b506100b26100ef3660046110bc565b34801561011957600080fd5b5061012d6101283660046111a6565b61067b565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610089565b60006040517f6e12839900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610194828401846110bc565b805160208201516040808401516060850151608086015160a087015160c088015160e08901516101008a015196517f08635a95000000000000000000000000000000000000000000000000000000008152999a5073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016996308635a959961023c99909890979695949392916004016111fe565b600060405180830381600087803b15801561025657600080fd5b505af115801561026a573d6000803e3d6000fd5b50505050806060015173ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff167f72ecb2b30e3013e4dcbd23949793c0ad87766498cc85c69a637847c1a790b8f68360e00151846040516102d992919061127b565b60405180910390a35050505050565b606061030c73ffffffffffffffffffffffffffffffffffffffff8816333087610736565b6040517fbda009fe00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff88811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063bda009fe90602401602060405180830381865afa15801561039d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103c19190611367565b905073ffffffffffffffffffffffffffffffffffffffff811661042d576040517f6c1460f400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff891660048201526024015b60405180910390fd5b61044e73ffffffffffffffffffffffffffffffffffffffff89168287610818565b600061045c84860186610f4f565b9050600081602001518260400151836000015161047991906113b3565b61048391906113ca565b9050803410156104c8576040517f03da4d2300000000000000000000000000000000000000000000000000000000815234600482015260248101829052604401610424565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16634fb1a07b348d8c8d8d89600001518a604001518b602001516040518060200160405280600081525060405160200161053b9291906113dd565b6040516020818303038152906040526040518963ffffffff1660e01b815260040161056c97969594939291906113f6565b60006040518083038185885af115801561058a573d6000803e3d6000fd5b50505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526105d19190810190611456565b6000805491925073ffffffffffffffffffffffffffffffffffffffff808c1692908e169182917f81d643ecc07b47f5fa42aefae481ea6dda3186401486151c9e46b03ee36fd7c69190819061062f9067ffffffffffffffff166114cd565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790558c86604051610665939291906114f4565b60405180910390a49a9950505050505050505050565b6040517fa7e28d4800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063a7e28d4890602401602060405180830381865afa15801561070c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107309190611367565b92915050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526108129085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261099f565b50505050565b8015806108b857506040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff838116602483015284169063dd62ed3e90604401602060405180830381865afa158015610892573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b6919061151d565b155b610944576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608401610424565b60405173ffffffffffffffffffffffffffffffffffffffff831660248201526044810182905261099a9084907f095ea7b30000000000000000000000000000000000000000000000000000000090606401610790565b505050565b6000610a01826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16610aab9092919063ffffffff16565b80519091501561099a5780806020019051810190610a1f9190611536565b61099a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610424565b6060610aba8484600085610ac2565b949350505050565b606082471015610b54576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610424565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610b7d9190611558565b60006040518083038185875af1925050503d8060008114610bba576040519150601f19603f3d011682016040523d82523d6000602084013e610bbf565b606091505b5091509150610bd087838387610bdb565b979650505050505050565b60608315610c71578251600003610c6a5773ffffffffffffffffffffffffffffffffffffffff85163b610c6a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610424565b5081610aba565b610aba8383815115610c865781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104249190610e8d565b73ffffffffffffffffffffffffffffffffffffffff811681146100ef57600080fd5b8035610ce781610cba565b919050565b60008083601f840112610cfe57600080fd5b50813567ffffffffffffffff811115610d1657600080fd5b602083019150836020828501011115610d2e57600080fd5b9250929050565b60008060008060608587031215610d4b57600080fd5b8435610d5681610cba565b93506020850135610d6681610cba565b9250604085013567ffffffffffffffff811115610d8257600080fd5b610d8e87828801610cec565b95989497509550505050565b60008060008060008060a08789031215610db357600080fd5b8635610dbe81610cba565b95506020870135610dce81610cba565b94506040870135610dde81610cba565b935060608701359250608087013567ffffffffffffffff811115610e0157600080fd5b610e0d89828a01610cec565b979a9699509497509295939492505050565b60005b83811015610e3a578181015183820152602001610e22565b50506000910152565b60008151808452610e5b816020860160208601610e1f565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610ea06020830184610e43565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff81118282101715610efa57610efa610ea7565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610f4757610f47610ea7565b604052919050565b600060608284031215610f6157600080fd5b6040516060810181811067ffffffffffffffff82111715610f8457610f84610ea7565b80604052508235815260208301356020820152604083013560408201528091505092915050565b600082601f830112610fbc57600080fd5b8135602067ffffffffffffffff821115610fd857610fd8610ea7565b8160051b610fe7828201610f00565b928352848101820192828101908785111561100157600080fd5b83870192505b84831015610bd057823582529183019190830190611007565b600067ffffffffffffffff82111561103a5761103a610ea7565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261107757600080fd5b813561108a61108582611020565b610f00565b81815284602083860101111561109f57600080fd5b816020850160208301376000918101602001919091529392505050565b6000602082840312156110ce57600080fd5b813567ffffffffffffffff808211156110e657600080fd5b9083019061012082860312156110fb57600080fd5b611103610ed6565b82358281111561111257600080fd5b61111e87828601610fab565b8252506020830135602082015261113760408401610cdc565b604082015261114860608401610cdc565b60608201526080830135608082015260a083013560a082015260c083013560c082015260e083013560e0820152610100808401358381111561118957600080fd5b61119588828701611066565b918301919091525095945050505050565b6000602082840312156111b857600080fd5b8135610ea081610cba565b600081518084526020808501945080840160005b838110156111f3578151875295820195908201906001016111d7565b509495945050505050565b60006101208083526112128184018d6111c3565b90508a602084015273ffffffffffffffffffffffffffffffffffffffff808b166040850152808a166060850152508760808401528660a08401528560c08401528460e084015282810361010084015261126b8185610e43565b9c9b505050505050505050505050565b82815260406020820152600082516101208060408501526112a06101608501836111c3565b91506020850151606085015260408501516112d3608086018273ffffffffffffffffffffffffffffffffffffffff169052565b50606085015173ffffffffffffffffffffffffffffffffffffffff811660a086015250608085015160c085015260a085015160e085015260c0850151610100818187015260e08701518387015280870151925050507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08483030161014085015261135d8282610e43565b9695505050505050565b60006020828403121561137957600080fd5b8151610ea081610cba565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808202811582820484141761073057610730611384565b8082018082111561073057610730611384565b828152604060208201526000610aba6040830184610e43565b600073ffffffffffffffffffffffffffffffffffffffff808a16835280891660208401528088166040840152508560608301528460808301528360a083015260e060c083015261144960e0830184610e43565b9998505050505050505050565b60006020828403121561146857600080fd5b815167ffffffffffffffff81111561147f57600080fd5b8201601f8101841361149057600080fd5b805161149e61108582611020565b8181528560208385010111156114b357600080fd5b6114c4826020830160208601610e1f565b95945050505050565b600067ffffffffffffffff8083168181036114ea576114ea611384565b6001019392505050565b67ffffffffffffffff841681528260208201526060604082015260006114c46060830184610e43565b60006020828403121561152f57600080fd5b5051919050565b60006020828403121561154857600080fd5b81518015158114610ea057600080fd5b6000825161156a818460208701610e1f565b919091019291505056fea164736f6c6343000813000a",
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

type ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20SentIterator struct {
	Event *ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20SentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent)
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
		it.Event = new(ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent)
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

func (it *ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20SentIterator) Error() error {
	return it.fail
}

func (it *ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20SentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent struct {
	LocalToken             common.Address
	RemoteToken            common.Address
	Recipient              common.Address
	Nonce                  *big.Int
	Amount                 *big.Int
	OutboundTransferResult []byte
	Raw                    types.Log
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterFilterer) FilterArbitrumL1ToL2ERC20Sent(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, recipient []common.Address) (*ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20SentIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ArbitrumL1BridgeAdapter.contract.FilterLogs(opts, "ArbitrumL1ToL2ERC20Sent", localTokenRule, remoteTokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20SentIterator{contract: _ArbitrumL1BridgeAdapter.contract, event: "ArbitrumL1ToL2ERC20Sent", logs: logs, sub: sub}, nil
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterFilterer) WatchArbitrumL1ToL2ERC20Sent(opts *bind.WatchOpts, sink chan<- *ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent, localToken []common.Address, remoteToken []common.Address, recipient []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ArbitrumL1BridgeAdapter.contract.WatchLogs(opts, "ArbitrumL1ToL2ERC20Sent", localTokenRule, remoteTokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent)
				if err := _ArbitrumL1BridgeAdapter.contract.UnpackLog(event, "ArbitrumL1ToL2ERC20Sent", log); err != nil {
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

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterFilterer) ParseArbitrumL1ToL2ERC20Sent(log types.Log) (*ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent, error) {
	event := new(ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent)
	if err := _ArbitrumL1BridgeAdapter.contract.UnpackLog(event, "ArbitrumL1ToL2ERC20Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20FinalizedIterator struct {
	Event *ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20FinalizedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized)
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
		it.Event = new(ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized)
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

func (it *ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20FinalizedIterator) Error() error {
	return it.fail
}

func (it *ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20FinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized struct {
	RemoteSender  common.Address
	LocalReceiver common.Address
	Amount        *big.Int
	Payload       ArbitrumL1BridgeAdapterArbitrumFinalizationPayload
	Raw           types.Log
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterFilterer) FilterArbitrumL2ToL1ERC20Finalized(opts *bind.FilterOpts, remoteSender []common.Address, localReceiver []common.Address) (*ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20FinalizedIterator, error) {

	var remoteSenderRule []interface{}
	for _, remoteSenderItem := range remoteSender {
		remoteSenderRule = append(remoteSenderRule, remoteSenderItem)
	}
	var localReceiverRule []interface{}
	for _, localReceiverItem := range localReceiver {
		localReceiverRule = append(localReceiverRule, localReceiverItem)
	}

	logs, sub, err := _ArbitrumL1BridgeAdapter.contract.FilterLogs(opts, "ArbitrumL2ToL1ERC20Finalized", remoteSenderRule, localReceiverRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20FinalizedIterator{contract: _ArbitrumL1BridgeAdapter.contract, event: "ArbitrumL2ToL1ERC20Finalized", logs: logs, sub: sub}, nil
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterFilterer) WatchArbitrumL2ToL1ERC20Finalized(opts *bind.WatchOpts, sink chan<- *ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized, remoteSender []common.Address, localReceiver []common.Address) (event.Subscription, error) {

	var remoteSenderRule []interface{}
	for _, remoteSenderItem := range remoteSender {
		remoteSenderRule = append(remoteSenderRule, remoteSenderItem)
	}
	var localReceiverRule []interface{}
	for _, localReceiverItem := range localReceiver {
		localReceiverRule = append(localReceiverRule, localReceiverItem)
	}

	logs, sub, err := _ArbitrumL1BridgeAdapter.contract.WatchLogs(opts, "ArbitrumL2ToL1ERC20Finalized", remoteSenderRule, localReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized)
				if err := _ArbitrumL1BridgeAdapter.contract.UnpackLog(event, "ArbitrumL2ToL1ERC20Finalized", log); err != nil {
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

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterFilterer) ParseArbitrumL2ToL1ERC20Finalized(log types.Log) (*ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized, error) {
	event := new(ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized)
	if err := _ArbitrumL1BridgeAdapter.contract.UnpackLog(event, "ArbitrumL2ToL1ERC20Finalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _ArbitrumL1BridgeAdapter.abi.Events["ArbitrumL1ToL2ERC20Sent"].ID:
		return _ArbitrumL1BridgeAdapter.ParseArbitrumL1ToL2ERC20Sent(log)
	case _ArbitrumL1BridgeAdapter.abi.Events["ArbitrumL2ToL1ERC20Finalized"].ID:
		return _ArbitrumL1BridgeAdapter.ParseArbitrumL2ToL1ERC20Finalized(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent) Topic() common.Hash {
	return common.HexToHash("0x81d643ecc07b47f5fa42aefae481ea6dda3186401486151c9e46b03ee36fd7c6")
}

func (ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized) Topic() common.Hash {
	return common.HexToHash("0x72ecb2b30e3013e4dcbd23949793c0ad87766498cc85c69a637847c1a790b8f6")
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapter) Address() common.Address {
	return _ArbitrumL1BridgeAdapter.address
}

type ArbitrumL1BridgeAdapterInterface interface {
	ExposeArbitrumFinalizationPayload(opts *bind.CallOpts, payload ArbitrumL1BridgeAdapterArbitrumFinalizationPayload) error

	ExposeSendERC20Params(opts *bind.CallOpts, params ArbitrumL1BridgeAdapterSendERC20Params) error

	GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error)

	GetL2Token(opts *bind.CallOpts, l1Token common.Address) (common.Address, error)

	FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, localToken common.Address, arg1 common.Address, recipient common.Address, amount *big.Int, bridgeSpecificPayload []byte) (*types.Transaction, error)

	FilterArbitrumL1ToL2ERC20Sent(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, recipient []common.Address) (*ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20SentIterator, error)

	WatchArbitrumL1ToL2ERC20Sent(opts *bind.WatchOpts, sink chan<- *ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent, localToken []common.Address, remoteToken []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseArbitrumL1ToL2ERC20Sent(log types.Log) (*ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent, error)

	FilterArbitrumL2ToL1ERC20Finalized(opts *bind.FilterOpts, remoteSender []common.Address, localReceiver []common.Address) (*ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20FinalizedIterator, error)

	WatchArbitrumL2ToL1ERC20Finalized(opts *bind.WatchOpts, sink chan<- *ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized, remoteSender []common.Address, localReceiver []common.Address) (event.Subscription, error)

	ParseArbitrumL2ToL1ERC20Finalized(log types.Log) (*ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
