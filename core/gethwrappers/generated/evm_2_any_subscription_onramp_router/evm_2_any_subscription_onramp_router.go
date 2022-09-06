// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_any_subscription_onramp_router

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

type Any2EVMSubscriptionOnRampRouterInterfaceRouterConfig struct {
	Fee      *big.Int
	FeeToken common.Address
	FeeAdmin common.Address
}

type CCIPEVM2AnySubscriptionMessage struct {
	Receiver common.Address
	Data     []byte
	Tokens   []common.Address
	Amounts  []*big.Int
	GasLimit *big.Int
}

var EVM2AnySubscriptionOnRampRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAdmin\",\"type\":\"address\"}],\"internalType\":\"structAny2EVMSubscriptionOnRampRouterInterface.RouterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FeeTokenAmountTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"FundingTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractAny2EVMSubscriptionOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByFeeAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"name\":\"FeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractAny2EVMSubscriptionOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SubscriptionUnfunded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVM2AnySubscriptionMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"contractAny2EVMSubscriptionOnRampInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"newFee\",\"type\":\"uint96\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractAny2EVMSubscriptionOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"setOnRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unfundSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200171f3803806200171f8339810160408190526200003491620001d7565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be8162000113565b5050815160208301516001600160601b039091166c010000000000000000000000006001600160a01b039283160217600455604090920151600580546001600160a01b03191691909316179091555062000269565b336001600160a01b038216036200016d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b0381168114620001d457600080fd5b50565b600060608284031215620001ea57600080fd5b604051606081016001600160401b03811182821017156200021b57634e487b7160e01b600052604160045260246000fd5b60405282516001600160601b03811681146200023657600080fd5b815260208301516200024881620001be565b602082015260408301516200025d81620001be565b60408201529392505050565b6114a680620002796000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063ae275dce1161008c578063d8a98f8c11610066578063d8a98f8c1461020b578063f1927cae14610234578063f2fde38b14610247578063f8b2cb4f1461025a57600080fd5b8063ae275dce146101c5578063c1060653146101d8578063ced72f87146101eb57600080fd5b806359e96b5b116100c857806359e96b5b1461017057806379ba5097146101855780638da5cb5b1461018d57806395e712db146101b257600080fd5b80630d58bf0c146100ef578063181f5a77146101205780635221c1f014610135575b600080fd5b6101026100fd366004610ff0565b610291565b60405167ffffffffffffffff90911681526020015b60405180910390f35b6101286103dc565b6040516101179190611124565b610160610143366004611137565b6000908152600260205260409020546001600160a01b0316151590565b6040519015158152602001610117565b61018361017e366004611150565b6103f8565b005b610183610464565b6000546001600160a01b03165b6040516001600160a01b039091168152602001610117565b6101836101c0366004611137565b61052d565b6101836101d3366004611191565b6105bf565b6101836101e6366004611137565b610674565b6004546040516bffffffffffffffffffffffff9091168152602001610117565b61019a610219366004611137565b6000908152600260205260409020546001600160a01b031690565b6101836102423660046111bf565b610706565b6101836102553660046111ef565b6107e4565b6102836102683660046111ef565b6001600160a01b031660009081526003602052604090205490565b604051908152602001610117565b6000828152600260205260408120546001600160a01b0316806102e8576040517f45abe4ae000000000000000000000000000000000000000000000000000000008152600481018590526024015b60405180910390fd5b6004546bffffffffffffffffffffffff16156103375760045433600090815260036020526040812080546bffffffffffffffffffffffff9093169290919061033190849061123b565b90915550505b61034a81846040015185606001516107f8565b6040517fae990dce0000000000000000000000000000000000000000000000000000000081526001600160a01b0382169063ae990dce90610391908690339060040161128d565b6020604051808303816000875af11580156103b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d49190611370565b949350505050565b6040518060600160405280602581526020016114756025913981565b61040061094d565b6104146001600160a01b03841683836109a9565b604080516001600160a01b038086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa89060600160405180910390a1505050565b6001546001600160a01b031633146104be5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016102df565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b336000818152600360205260408120805484929061054c90849061123b565b9091555050600454610578906c0100000000000000000000000090046001600160a01b031682846109a9565b806001600160a01b03167f437ce891210910c3800b0cb0fa2ee1dad361d5f396e4c457707a9f7ab918fd39836040516105b391815260200190565b60405180910390a25050565b6005546001600160a01b03163314610603576040517f112cedd700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff83169081179091556040519081527ff53f31763bcf350b90021051ebd7bbbc5e269027d22f73fd987c13db1426b3729060200160405180910390a150565b600454339061069e906c0100000000000000000000000090046001600160a01b0316823085610a57565b6001600160a01b038116600090815260036020526040812080548492906106c690849061139a565b90915550506040518281526001600160a01b038216907fc89bca949929d103fee7b5eae37fdafa6f82a94463c8e9ea2ec5c6b488705680906020016105b3565b61070e61094d565b6000828152600260205260409020546001600160a01b03808316911603610773576040517fe31de3b2000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b03821660248201526044016102df565b60008281526002602052604080822080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0385169081179091559051909184917f4b680ef9fa79bb5f36e7559d7b33fd57a8336f78cc120c8cd93333b5ade624cb9190a35050565b6107ec61094d565b6107f581610aa8565b50565b60005b8251811015610947576000838281518110610818576108186113b2565b60209081029190910101516040517f04c2a34a0000000000000000000000000000000000000000000000000000000081526001600160a01b0380831660048301529192506000918716906304c2a34a906024016020604051808303816000875af115801561088a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ae91906113e1565b90506001600160a01b0381166108fb576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016102df565b6109343382868681518110610912576109126113b2565b6020026020010151856001600160a01b0316610a57909392919063ffffffff16565b505080610940906113fe565b90506107fb565b50505050565b6000546001600160a01b031633146109a75760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016102df565b565b6040516001600160a01b038316602482015260448101829052610a529084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610b69565b505050565b6040516001600160a01b03808516602483015283166044820152606481018290526109479085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016109ee565b336001600160a01b03821603610b005760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016102df565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000610bbe826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610c4e9092919063ffffffff16565b805190915015610a525780806020019051810190610bdc9190611436565b610a525760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016102df565b6060610c5d8484600085610c67565b90505b9392505050565b606082471015610cdf5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016102df565b843b610d2d5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102df565b600080866001600160a01b03168587604051610d499190611458565b60006040518083038185875af1925050503d8060008114610d86576040519150601f19603f3d011682016040523d82523d6000602084013e610d8b565b606091505b5091509150610d9b828286610da6565b979650505050505050565b60608315610db5575081610c60565b825115610dc55782518084602001fd5b8160405162461bcd60e51b81526004016102df9190611124565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff81118282101715610e3157610e31610ddf565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610e6057610e60610ddf565b604052919050565b6001600160a01b03811681146107f557600080fd5b8035610e8881610e68565b919050565b600082601f830112610e9e57600080fd5b813567ffffffffffffffff811115610eb857610eb8610ddf565b610ecb6020601f19601f84011601610e37565b818152846020838601011115610ee057600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115610f1757610f17610ddf565b5060051b60200190565b600082601f830112610f3257600080fd5b81356020610f47610f4283610efd565b610e37565b82815260059290921b84018101918181019086841115610f6657600080fd5b8286015b84811015610f8a578035610f7d81610e68565b8352918301918301610f6a565b509695505050505050565b600082601f830112610fa657600080fd5b81356020610fb6610f4283610efd565b82815260059290921b84018101918181019086841115610fd557600080fd5b8286015b84811015610f8a5780358352918301918301610fd9565b6000806040838503121561100357600080fd5b82359150602083013567ffffffffffffffff8082111561102257600080fd5b9084019060a0828703121561103657600080fd5b61103e610e0e565b61104783610e7d565b815260208301358281111561105b57600080fd5b61106788828601610e8d565b60208301525060408301358281111561107f57600080fd5b61108b88828601610f21565b6040830152506060830135828111156110a357600080fd5b6110af88828601610f95565b606083015250608083013560808201528093505050509250929050565b60005b838110156110e75781810151838201526020016110cf565b838111156109475750506000910152565b600081518084526111108160208601602086016110cc565b601f01601f19169290920160200192915050565b602081526000610c6060208301846110f8565b60006020828403121561114957600080fd5b5035919050565b60008060006060848603121561116557600080fd5b833561117081610e68565b9250602084013561118081610e68565b929592945050506040919091013590565b6000602082840312156111a357600080fd5b81356bffffffffffffffffffffffff81168114610c6057600080fd5b600080604083850312156111d257600080fd5b8235915060208301356111e481610e68565b809150509250929050565b60006020828403121561120157600080fd5b8135610c6081610e68565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561124d5761124d61120c565b500390565b600081518084526020808501945080840160005b8381101561128257815187529582019590820190600101611266565b509495945050505050565b6040815260006001600160a01b0380855116604084015260208086015160a060608601526112be60e08601826110f8565b60408801517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087830381016080890152815180845291850193506000929091908501905b8084101561132457845187168252938501936001939093019290850190611302565b5060608a01519550818882030160a08901526113408187611252565b955050505050608086015160c0850152819250611367818501866001600160a01b03169052565b50509392505050565b60006020828403121561138257600080fd5b815167ffffffffffffffff81168114610c6057600080fd5b600082198211156113ad576113ad61120c565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156113f357600080fd5b8151610c6081610e68565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361142f5761142f61120c565b5060010190565b60006020828403121561144857600080fd5b81518015158114610c6057600080fd5b6000825161146a8184602087016110cc565b919091019291505056fe45564d32416e79537562736372697074696f6e4f6e52616d70526f7574657220312e302e30a164736f6c634300080f000a",
}

var EVM2AnySubscriptionOnRampRouterABI = EVM2AnySubscriptionOnRampRouterMetaData.ABI

var EVM2AnySubscriptionOnRampRouterBin = EVM2AnySubscriptionOnRampRouterMetaData.Bin

func DeployEVM2AnySubscriptionOnRampRouter(auth *bind.TransactOpts, backend bind.ContractBackend, config Any2EVMSubscriptionOnRampRouterInterfaceRouterConfig) (common.Address, *types.Transaction, *EVM2AnySubscriptionOnRampRouter, error) {
	parsed, err := EVM2AnySubscriptionOnRampRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2AnySubscriptionOnRampRouterBin), backend, config)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2AnySubscriptionOnRampRouter{EVM2AnySubscriptionOnRampRouterCaller: EVM2AnySubscriptionOnRampRouterCaller{contract: contract}, EVM2AnySubscriptionOnRampRouterTransactor: EVM2AnySubscriptionOnRampRouterTransactor{contract: contract}, EVM2AnySubscriptionOnRampRouterFilterer: EVM2AnySubscriptionOnRampRouterFilterer{contract: contract}}, nil
}

type EVM2AnySubscriptionOnRampRouter struct {
	address common.Address
	abi     abi.ABI
	EVM2AnySubscriptionOnRampRouterCaller
	EVM2AnySubscriptionOnRampRouterTransactor
	EVM2AnySubscriptionOnRampRouterFilterer
}

type EVM2AnySubscriptionOnRampRouterCaller struct {
	contract *bind.BoundContract
}

type EVM2AnySubscriptionOnRampRouterTransactor struct {
	contract *bind.BoundContract
}

type EVM2AnySubscriptionOnRampRouterFilterer struct {
	contract *bind.BoundContract
}

type EVM2AnySubscriptionOnRampRouterSession struct {
	Contract     *EVM2AnySubscriptionOnRampRouter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2AnySubscriptionOnRampRouterCallerSession struct {
	Contract *EVM2AnySubscriptionOnRampRouterCaller
	CallOpts bind.CallOpts
}

type EVM2AnySubscriptionOnRampRouterTransactorSession struct {
	Contract     *EVM2AnySubscriptionOnRampRouterTransactor
	TransactOpts bind.TransactOpts
}

type EVM2AnySubscriptionOnRampRouterRaw struct {
	Contract *EVM2AnySubscriptionOnRampRouter
}

type EVM2AnySubscriptionOnRampRouterCallerRaw struct {
	Contract *EVM2AnySubscriptionOnRampRouterCaller
}

type EVM2AnySubscriptionOnRampRouterTransactorRaw struct {
	Contract *EVM2AnySubscriptionOnRampRouterTransactor
}

func NewEVM2AnySubscriptionOnRampRouter(address common.Address, backend bind.ContractBackend) (*EVM2AnySubscriptionOnRampRouter, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2AnySubscriptionOnRampRouterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2AnySubscriptionOnRampRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2AnySubscriptionOnRampRouter{address: address, abi: abi, EVM2AnySubscriptionOnRampRouterCaller: EVM2AnySubscriptionOnRampRouterCaller{contract: contract}, EVM2AnySubscriptionOnRampRouterTransactor: EVM2AnySubscriptionOnRampRouterTransactor{contract: contract}, EVM2AnySubscriptionOnRampRouterFilterer: EVM2AnySubscriptionOnRampRouterFilterer{contract: contract}}, nil
}

func NewEVM2AnySubscriptionOnRampRouterCaller(address common.Address, caller bind.ContractCaller) (*EVM2AnySubscriptionOnRampRouterCaller, error) {
	contract, err := bindEVM2AnySubscriptionOnRampRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2AnySubscriptionOnRampRouterCaller{contract: contract}, nil
}

func NewEVM2AnySubscriptionOnRampRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2AnySubscriptionOnRampRouterTransactor, error) {
	contract, err := bindEVM2AnySubscriptionOnRampRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2AnySubscriptionOnRampRouterTransactor{contract: contract}, nil
}

func NewEVM2AnySubscriptionOnRampRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2AnySubscriptionOnRampRouterFilterer, error) {
	contract, err := bindEVM2AnySubscriptionOnRampRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2AnySubscriptionOnRampRouterFilterer{contract: contract}, nil
}

func bindEVM2AnySubscriptionOnRampRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2AnySubscriptionOnRampRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2AnySubscriptionOnRampRouter.Contract.EVM2AnySubscriptionOnRampRouterCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.EVM2AnySubscriptionOnRampRouterTransactor.contract.Transfer(opts)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.EVM2AnySubscriptionOnRampRouterTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2AnySubscriptionOnRampRouter.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.contract.Transfer(opts)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterCaller) GetBalance(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EVM2AnySubscriptionOnRampRouter.contract.Call(opts, &out, "getBalance", sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterSession) GetBalance(sender common.Address) (*big.Int, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.GetBalance(&_EVM2AnySubscriptionOnRampRouter.CallOpts, sender)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterCallerSession) GetBalance(sender common.Address) (*big.Int, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.GetBalance(&_EVM2AnySubscriptionOnRampRouter.CallOpts, sender)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterCaller) GetFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2AnySubscriptionOnRampRouter.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterSession) GetFee() (*big.Int, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.GetFee(&_EVM2AnySubscriptionOnRampRouter.CallOpts)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterCallerSession) GetFee() (*big.Int, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.GetFee(&_EVM2AnySubscriptionOnRampRouter.CallOpts)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterCaller) GetOnRamp(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EVM2AnySubscriptionOnRampRouter.contract.Call(opts, &out, "getOnRamp", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterSession) GetOnRamp(chainId *big.Int) (common.Address, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.GetOnRamp(&_EVM2AnySubscriptionOnRampRouter.CallOpts, chainId)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterCallerSession) GetOnRamp(chainId *big.Int) (common.Address, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.GetOnRamp(&_EVM2AnySubscriptionOnRampRouter.CallOpts, chainId)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterCaller) IsChainSupported(opts *bind.CallOpts, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _EVM2AnySubscriptionOnRampRouter.contract.Call(opts, &out, "isChainSupported", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterSession) IsChainSupported(chainId *big.Int) (bool, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.IsChainSupported(&_EVM2AnySubscriptionOnRampRouter.CallOpts, chainId)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterCallerSession) IsChainSupported(chainId *big.Int) (bool, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.IsChainSupported(&_EVM2AnySubscriptionOnRampRouter.CallOpts, chainId)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2AnySubscriptionOnRampRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterSession) Owner() (common.Address, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.Owner(&_EVM2AnySubscriptionOnRampRouter.CallOpts)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterCallerSession) Owner() (common.Address, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.Owner(&_EVM2AnySubscriptionOnRampRouter.CallOpts)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2AnySubscriptionOnRampRouter.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterSession) TypeAndVersion() (string, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.TypeAndVersion(&_EVM2AnySubscriptionOnRampRouter.CallOpts)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterCallerSession) TypeAndVersion() (string, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.TypeAndVersion(&_EVM2AnySubscriptionOnRampRouter.CallOpts)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.AcceptOwnership(&_EVM2AnySubscriptionOnRampRouter.TransactOpts)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.AcceptOwnership(&_EVM2AnySubscriptionOnRampRouter.TransactOpts)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactor) CcipSend(opts *bind.TransactOpts, destinationChainId *big.Int, message CCIPEVM2AnySubscriptionMessage) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.contract.Transact(opts, "ccipSend", destinationChainId, message)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterSession) CcipSend(destinationChainId *big.Int, message CCIPEVM2AnySubscriptionMessage) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.CcipSend(&_EVM2AnySubscriptionOnRampRouter.TransactOpts, destinationChainId, message)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactorSession) CcipSend(destinationChainId *big.Int, message CCIPEVM2AnySubscriptionMessage) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.CcipSend(&_EVM2AnySubscriptionOnRampRouter.TransactOpts, destinationChainId, message)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactor) FundSubscription(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.contract.Transact(opts, "fundSubscription", amount)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterSession) FundSubscription(amount *big.Int) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.FundSubscription(&_EVM2AnySubscriptionOnRampRouter.TransactOpts, amount)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactorSession) FundSubscription(amount *big.Int) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.FundSubscription(&_EVM2AnySubscriptionOnRampRouter.TransactOpts, amount)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactor) SetFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.contract.Transact(opts, "setFee", newFee)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.SetFee(&_EVM2AnySubscriptionOnRampRouter.TransactOpts, newFee)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactorSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.SetFee(&_EVM2AnySubscriptionOnRampRouter.TransactOpts, newFee)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactor) SetOnRamp(opts *bind.TransactOpts, chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.contract.Transact(opts, "setOnRamp", chainId, onRamp)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterSession) SetOnRamp(chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.SetOnRamp(&_EVM2AnySubscriptionOnRampRouter.TransactOpts, chainId, onRamp)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactorSession) SetOnRamp(chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.SetOnRamp(&_EVM2AnySubscriptionOnRampRouter.TransactOpts, chainId, onRamp)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.TransferOwnership(&_EVM2AnySubscriptionOnRampRouter.TransactOpts, to)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.TransferOwnership(&_EVM2AnySubscriptionOnRampRouter.TransactOpts, to)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactor) UnfundSubscription(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.contract.Transact(opts, "unfundSubscription", amount)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterSession) UnfundSubscription(amount *big.Int) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.UnfundSubscription(&_EVM2AnySubscriptionOnRampRouter.TransactOpts, amount)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactorSession) UnfundSubscription(amount *big.Int) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.UnfundSubscription(&_EVM2AnySubscriptionOnRampRouter.TransactOpts, amount)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactor) WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.contract.Transact(opts, "withdrawAccumulatedFees", feeToken, recipient, amount)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.WithdrawAccumulatedFees(&_EVM2AnySubscriptionOnRampRouter.TransactOpts, feeToken, recipient, amount)
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterTransactorSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2AnySubscriptionOnRampRouter.Contract.WithdrawAccumulatedFees(&_EVM2AnySubscriptionOnRampRouter.TransactOpts, feeToken, recipient, amount)
}

type EVM2AnySubscriptionOnRampRouterFeeChargedIterator struct {
	Event *EVM2AnySubscriptionOnRampRouterFeeCharged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2AnySubscriptionOnRampRouterFeeChargedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2AnySubscriptionOnRampRouterFeeCharged)
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
		it.Event = new(EVM2AnySubscriptionOnRampRouterFeeCharged)
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

func (it *EVM2AnySubscriptionOnRampRouterFeeChargedIterator) Error() error {
	return it.fail
}

func (it *EVM2AnySubscriptionOnRampRouterFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2AnySubscriptionOnRampRouterFeeCharged struct {
	From common.Address
	To   common.Address
	Fee  *big.Int
	Raw  types.Log
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) FilterFeeCharged(opts *bind.FilterOpts) (*EVM2AnySubscriptionOnRampRouterFeeChargedIterator, error) {

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.FilterLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return &EVM2AnySubscriptionOnRampRouterFeeChargedIterator{contract: _EVM2AnySubscriptionOnRampRouter.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterFeeCharged) (event.Subscription, error) {

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.WatchLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2AnySubscriptionOnRampRouterFeeCharged)
				if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "FeeCharged", log); err != nil {
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

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) ParseFeeCharged(log types.Log) (*EVM2AnySubscriptionOnRampRouterFeeCharged, error) {
	event := new(EVM2AnySubscriptionOnRampRouterFeeCharged)
	if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2AnySubscriptionOnRampRouterFeeSetIterator struct {
	Event *EVM2AnySubscriptionOnRampRouterFeeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2AnySubscriptionOnRampRouterFeeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2AnySubscriptionOnRampRouterFeeSet)
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
		it.Event = new(EVM2AnySubscriptionOnRampRouterFeeSet)
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

func (it *EVM2AnySubscriptionOnRampRouterFeeSetIterator) Error() error {
	return it.fail
}

func (it *EVM2AnySubscriptionOnRampRouterFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2AnySubscriptionOnRampRouterFeeSet struct {
	Arg0 *big.Int
	Raw  types.Log
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) FilterFeeSet(opts *bind.FilterOpts) (*EVM2AnySubscriptionOnRampRouterFeeSetIterator, error) {

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.FilterLogs(opts, "FeeSet")
	if err != nil {
		return nil, err
	}
	return &EVM2AnySubscriptionOnRampRouterFeeSetIterator{contract: _EVM2AnySubscriptionOnRampRouter.contract, event: "FeeSet", logs: logs, sub: sub}, nil
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) WatchFeeSet(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterFeeSet) (event.Subscription, error) {

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.WatchLogs(opts, "FeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2AnySubscriptionOnRampRouterFeeSet)
				if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "FeeSet", log); err != nil {
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

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) ParseFeeSet(log types.Log) (*EVM2AnySubscriptionOnRampRouterFeeSet, error) {
	event := new(EVM2AnySubscriptionOnRampRouterFeeSet)
	if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "FeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2AnySubscriptionOnRampRouterFeesWithdrawnIterator struct {
	Event *EVM2AnySubscriptionOnRampRouterFeesWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2AnySubscriptionOnRampRouterFeesWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2AnySubscriptionOnRampRouterFeesWithdrawn)
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
		it.Event = new(EVM2AnySubscriptionOnRampRouterFeesWithdrawn)
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

func (it *EVM2AnySubscriptionOnRampRouterFeesWithdrawnIterator) Error() error {
	return it.fail
}

func (it *EVM2AnySubscriptionOnRampRouterFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2AnySubscriptionOnRampRouterFeesWithdrawn struct {
	FeeToken  common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*EVM2AnySubscriptionOnRampRouterFeesWithdrawnIterator, error) {

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &EVM2AnySubscriptionOnRampRouterFeesWithdrawnIterator{contract: _EVM2AnySubscriptionOnRampRouter.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2AnySubscriptionOnRampRouterFeesWithdrawn)
				if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) ParseFeesWithdrawn(log types.Log) (*EVM2AnySubscriptionOnRampRouterFeesWithdrawn, error) {
	event := new(EVM2AnySubscriptionOnRampRouterFeesWithdrawn)
	if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2AnySubscriptionOnRampRouterOnRampSetIterator struct {
	Event *EVM2AnySubscriptionOnRampRouterOnRampSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2AnySubscriptionOnRampRouterOnRampSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2AnySubscriptionOnRampRouterOnRampSet)
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
		it.Event = new(EVM2AnySubscriptionOnRampRouterOnRampSet)
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

func (it *EVM2AnySubscriptionOnRampRouterOnRampSetIterator) Error() error {
	return it.fail
}

func (it *EVM2AnySubscriptionOnRampRouterOnRampSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2AnySubscriptionOnRampRouterOnRampSet struct {
	ChainId *big.Int
	OnRamp  common.Address
	Raw     types.Log
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) FilterOnRampSet(opts *bind.FilterOpts, chainId []*big.Int, onRamp []common.Address) (*EVM2AnySubscriptionOnRampRouterOnRampSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var onRampRule []interface{}
	for _, onRampItem := range onRamp {
		onRampRule = append(onRampRule, onRampItem)
	}

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.FilterLogs(opts, "OnRampSet", chainIdRule, onRampRule)
	if err != nil {
		return nil, err
	}
	return &EVM2AnySubscriptionOnRampRouterOnRampSetIterator{contract: _EVM2AnySubscriptionOnRampRouter.contract, event: "OnRampSet", logs: logs, sub: sub}, nil
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterOnRampSet, chainId []*big.Int, onRamp []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var onRampRule []interface{}
	for _, onRampItem := range onRamp {
		onRampRule = append(onRampRule, onRampItem)
	}

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.WatchLogs(opts, "OnRampSet", chainIdRule, onRampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2AnySubscriptionOnRampRouterOnRampSet)
				if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "OnRampSet", log); err != nil {
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

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) ParseOnRampSet(log types.Log) (*EVM2AnySubscriptionOnRampRouterOnRampSet, error) {
	event := new(EVM2AnySubscriptionOnRampRouterOnRampSet)
	if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "OnRampSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2AnySubscriptionOnRampRouterOwnershipTransferRequestedIterator struct {
	Event *EVM2AnySubscriptionOnRampRouterOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2AnySubscriptionOnRampRouterOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2AnySubscriptionOnRampRouterOwnershipTransferRequested)
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
		it.Event = new(EVM2AnySubscriptionOnRampRouterOwnershipTransferRequested)
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

func (it *EVM2AnySubscriptionOnRampRouterOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2AnySubscriptionOnRampRouterOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2AnySubscriptionOnRampRouterOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2AnySubscriptionOnRampRouterOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2AnySubscriptionOnRampRouterOwnershipTransferRequestedIterator{contract: _EVM2AnySubscriptionOnRampRouter.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2AnySubscriptionOnRampRouterOwnershipTransferRequested)
				if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2AnySubscriptionOnRampRouterOwnershipTransferRequested, error) {
	event := new(EVM2AnySubscriptionOnRampRouterOwnershipTransferRequested)
	if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2AnySubscriptionOnRampRouterOwnershipTransferredIterator struct {
	Event *EVM2AnySubscriptionOnRampRouterOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2AnySubscriptionOnRampRouterOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2AnySubscriptionOnRampRouterOwnershipTransferred)
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
		it.Event = new(EVM2AnySubscriptionOnRampRouterOwnershipTransferred)
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

func (it *EVM2AnySubscriptionOnRampRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2AnySubscriptionOnRampRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2AnySubscriptionOnRampRouterOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2AnySubscriptionOnRampRouterOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2AnySubscriptionOnRampRouterOwnershipTransferredIterator{contract: _EVM2AnySubscriptionOnRampRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2AnySubscriptionOnRampRouterOwnershipTransferred)
				if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2AnySubscriptionOnRampRouterOwnershipTransferred, error) {
	event := new(EVM2AnySubscriptionOnRampRouterOwnershipTransferred)
	if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2AnySubscriptionOnRampRouterSubscriptionFundedIterator struct {
	Event *EVM2AnySubscriptionOnRampRouterSubscriptionFunded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2AnySubscriptionOnRampRouterSubscriptionFundedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2AnySubscriptionOnRampRouterSubscriptionFunded)
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
		it.Event = new(EVM2AnySubscriptionOnRampRouterSubscriptionFunded)
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

func (it *EVM2AnySubscriptionOnRampRouterSubscriptionFundedIterator) Error() error {
	return it.fail
}

func (it *EVM2AnySubscriptionOnRampRouterSubscriptionFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2AnySubscriptionOnRampRouterSubscriptionFunded struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) FilterSubscriptionFunded(opts *bind.FilterOpts, sender []common.Address) (*EVM2AnySubscriptionOnRampRouterSubscriptionFundedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.FilterLogs(opts, "SubscriptionFunded", senderRule)
	if err != nil {
		return nil, err
	}
	return &EVM2AnySubscriptionOnRampRouterSubscriptionFundedIterator{contract: _EVM2AnySubscriptionOnRampRouter.contract, event: "SubscriptionFunded", logs: logs, sub: sub}, nil
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterSubscriptionFunded, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.WatchLogs(opts, "SubscriptionFunded", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2AnySubscriptionOnRampRouterSubscriptionFunded)
				if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
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

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) ParseSubscriptionFunded(log types.Log) (*EVM2AnySubscriptionOnRampRouterSubscriptionFunded, error) {
	event := new(EVM2AnySubscriptionOnRampRouterSubscriptionFunded)
	if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2AnySubscriptionOnRampRouterSubscriptionUnfundedIterator struct {
	Event *EVM2AnySubscriptionOnRampRouterSubscriptionUnfunded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2AnySubscriptionOnRampRouterSubscriptionUnfundedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2AnySubscriptionOnRampRouterSubscriptionUnfunded)
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
		it.Event = new(EVM2AnySubscriptionOnRampRouterSubscriptionUnfunded)
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

func (it *EVM2AnySubscriptionOnRampRouterSubscriptionUnfundedIterator) Error() error {
	return it.fail
}

func (it *EVM2AnySubscriptionOnRampRouterSubscriptionUnfundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2AnySubscriptionOnRampRouterSubscriptionUnfunded struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) FilterSubscriptionUnfunded(opts *bind.FilterOpts, sender []common.Address) (*EVM2AnySubscriptionOnRampRouterSubscriptionUnfundedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.FilterLogs(opts, "SubscriptionUnfunded", senderRule)
	if err != nil {
		return nil, err
	}
	return &EVM2AnySubscriptionOnRampRouterSubscriptionUnfundedIterator{contract: _EVM2AnySubscriptionOnRampRouter.contract, event: "SubscriptionUnfunded", logs: logs, sub: sub}, nil
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) WatchSubscriptionUnfunded(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterSubscriptionUnfunded, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2AnySubscriptionOnRampRouter.contract.WatchLogs(opts, "SubscriptionUnfunded", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2AnySubscriptionOnRampRouterSubscriptionUnfunded)
				if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "SubscriptionUnfunded", log); err != nil {
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

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouterFilterer) ParseSubscriptionUnfunded(log types.Log) (*EVM2AnySubscriptionOnRampRouterSubscriptionUnfunded, error) {
	event := new(EVM2AnySubscriptionOnRampRouterSubscriptionUnfunded)
	if err := _EVM2AnySubscriptionOnRampRouter.contract.UnpackLog(event, "SubscriptionUnfunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2AnySubscriptionOnRampRouter.abi.Events["FeeCharged"].ID:
		return _EVM2AnySubscriptionOnRampRouter.ParseFeeCharged(log)
	case _EVM2AnySubscriptionOnRampRouter.abi.Events["FeeSet"].ID:
		return _EVM2AnySubscriptionOnRampRouter.ParseFeeSet(log)
	case _EVM2AnySubscriptionOnRampRouter.abi.Events["FeesWithdrawn"].ID:
		return _EVM2AnySubscriptionOnRampRouter.ParseFeesWithdrawn(log)
	case _EVM2AnySubscriptionOnRampRouter.abi.Events["OnRampSet"].ID:
		return _EVM2AnySubscriptionOnRampRouter.ParseOnRampSet(log)
	case _EVM2AnySubscriptionOnRampRouter.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2AnySubscriptionOnRampRouter.ParseOwnershipTransferRequested(log)
	case _EVM2AnySubscriptionOnRampRouter.abi.Events["OwnershipTransferred"].ID:
		return _EVM2AnySubscriptionOnRampRouter.ParseOwnershipTransferred(log)
	case _EVM2AnySubscriptionOnRampRouter.abi.Events["SubscriptionFunded"].ID:
		return _EVM2AnySubscriptionOnRampRouter.ParseSubscriptionFunded(log)
	case _EVM2AnySubscriptionOnRampRouter.abi.Events["SubscriptionUnfunded"].ID:
		return _EVM2AnySubscriptionOnRampRouter.ParseSubscriptionUnfunded(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2AnySubscriptionOnRampRouterFeeCharged) Topic() common.Hash {
	return common.HexToHash("0x945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d")
}

func (EVM2AnySubscriptionOnRampRouterFeeSet) Topic() common.Hash {
	return common.HexToHash("0xf53f31763bcf350b90021051ebd7bbbc5e269027d22f73fd987c13db1426b372")
}

func (EVM2AnySubscriptionOnRampRouterFeesWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8")
}

func (EVM2AnySubscriptionOnRampRouterOnRampSet) Topic() common.Hash {
	return common.HexToHash("0x4b680ef9fa79bb5f36e7559d7b33fd57a8336f78cc120c8cd93333b5ade624cb")
}

func (EVM2AnySubscriptionOnRampRouterOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2AnySubscriptionOnRampRouterOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2AnySubscriptionOnRampRouterSubscriptionFunded) Topic() common.Hash {
	return common.HexToHash("0xc89bca949929d103fee7b5eae37fdafa6f82a94463c8e9ea2ec5c6b488705680")
}

func (EVM2AnySubscriptionOnRampRouterSubscriptionUnfunded) Topic() common.Hash {
	return common.HexToHash("0x437ce891210910c3800b0cb0fa2ee1dad361d5f396e4c457707a9f7ab918fd39")
}

func (_EVM2AnySubscriptionOnRampRouter *EVM2AnySubscriptionOnRampRouter) Address() common.Address {
	return _EVM2AnySubscriptionOnRampRouter.address
}

type EVM2AnySubscriptionOnRampRouterInterface interface {
	GetBalance(opts *bind.CallOpts, sender common.Address) (*big.Int, error)

	GetFee(opts *bind.CallOpts) (*big.Int, error)

	GetOnRamp(opts *bind.CallOpts, chainId *big.Int) (common.Address, error)

	IsChainSupported(opts *bind.CallOpts, chainId *big.Int) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainId *big.Int, message CCIPEVM2AnySubscriptionMessage) (*types.Transaction, error)

	FundSubscription(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	SetFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error)

	SetOnRamp(opts *bind.TransactOpts, chainId *big.Int, onRamp common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UnfundSubscription(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	FilterFeeCharged(opts *bind.FilterOpts) (*EVM2AnySubscriptionOnRampRouterFeeChargedIterator, error)

	WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterFeeCharged) (event.Subscription, error)

	ParseFeeCharged(log types.Log) (*EVM2AnySubscriptionOnRampRouterFeeCharged, error)

	FilterFeeSet(opts *bind.FilterOpts) (*EVM2AnySubscriptionOnRampRouterFeeSetIterator, error)

	WatchFeeSet(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterFeeSet) (event.Subscription, error)

	ParseFeeSet(log types.Log) (*EVM2AnySubscriptionOnRampRouterFeeSet, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*EVM2AnySubscriptionOnRampRouterFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*EVM2AnySubscriptionOnRampRouterFeesWithdrawn, error)

	FilterOnRampSet(opts *bind.FilterOpts, chainId []*big.Int, onRamp []common.Address) (*EVM2AnySubscriptionOnRampRouterOnRampSetIterator, error)

	WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterOnRampSet, chainId []*big.Int, onRamp []common.Address) (event.Subscription, error)

	ParseOnRampSet(log types.Log) (*EVM2AnySubscriptionOnRampRouterOnRampSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2AnySubscriptionOnRampRouterOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2AnySubscriptionOnRampRouterOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2AnySubscriptionOnRampRouterOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2AnySubscriptionOnRampRouterOwnershipTransferred, error)

	FilterSubscriptionFunded(opts *bind.FilterOpts, sender []common.Address) (*EVM2AnySubscriptionOnRampRouterSubscriptionFundedIterator, error)

	WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterSubscriptionFunded, sender []common.Address) (event.Subscription, error)

	ParseSubscriptionFunded(log types.Log) (*EVM2AnySubscriptionOnRampRouterSubscriptionFunded, error)

	FilterSubscriptionUnfunded(opts *bind.FilterOpts, sender []common.Address) (*EVM2AnySubscriptionOnRampRouterSubscriptionUnfundedIterator, error)

	WatchSubscriptionUnfunded(opts *bind.WatchOpts, sink chan<- *EVM2AnySubscriptionOnRampRouterSubscriptionUnfunded, sender []common.Address) (event.Subscription, error)

	ParseSubscriptionUnfunded(log types.Log) (*EVM2AnySubscriptionOnRampRouterSubscriptionUnfunded, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
