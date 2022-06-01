// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package onramp_router

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
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated"
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

type CCIPEVMToAnyTollMessage struct {
	Receiver       common.Address
	Data           []byte
	Tokens         []common.Address
	Amounts        []*big.Int
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	GasLimit       *big.Int
}

var OnRampRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMToAnyTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"contractOnRampInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"setOnRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b611713806101576000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80638da5cb5b11610076578063e7c62c8c1161005b578063e7c62c8c146101ca578063f1927cae146101f6578063f2fde38b1461020957600080fd5b80638da5cb5b14610155578063d8a98f8c1461019457600080fd5b8063181f5a77146100a85780635221c1f0146100f057806359e96b5b1461013857806379ba50971461014d575b600080fd5b604080518082018252601281527f4f6e52616d70526f7574657220302e302e310000000000000000000000000000602082015290516100e79190611044565b60405180910390f35b6101286100fe366004611057565b60009081526002602052604090205473ffffffffffffffffffffffffffffffffffffffff16151590565b60405190151581526020016100e7565b61014b6101463660046110a2565b61021c565b005b61014b6102a2565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e7565b61016f6101a2366004611057565b60009081526002602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b6101dd6101d836600461130b565b6103a4565b60405167ffffffffffffffff90911681526020016100e7565b61014b610204366004611402565b61051f565b61014b610217366004611432565b610624565b610224610638565b61024573ffffffffffffffffffffffffffffffffffffffff841683836106bb565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa89060600160405180910390a1505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610328576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b600082815260026020526040812054339073ffffffffffffffffffffffffffffffffffffffff1680610405576040517f45abe4ae0000000000000000000000000000000000000000000000000000000081526004810186905260240161031f565b83606001515184604001515114610448576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610467828660400151876060015188608001518960a00151610794565b9050808560a00181815161047b919061147e565b9052506040517f05afe24a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316906305afe24a906104d290889087906004016114d0565b6020604051808303816000875af11580156104f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105159190611601565b9695505050505050565b610527610638565b60008281526002602052604090205473ffffffffffffffffffffffffffffffffffffffff8083169116036105a6576040517fe31de3b20000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff8216602482015260440161031f565b60008281526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091559051909184917f4b680ef9fa79bb5f36e7559d7b33fd57a8336f78cc120c8cd93333b5ade624cb9190a35050565b61062c610638565b61063581610b7d565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146106b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161031f565b565b60405173ffffffffffffffffffffffffffffffffffffffff831660248201526044810182905261078f9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610c72565b505050565b60003373ffffffffffffffffffffffffffffffffffffffff8416156109ed576040517fd0d5de6100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015288169063d0d5de61906024016020604051808303816000875af1158015610821573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610845919061162b565b915061086973ffffffffffffffffffffffffffffffffffffffff8516823085610d7e565b610873828461147e565b6040805173ffffffffffffffffffffffffffffffffffffffff841681523060208201529081018490529093507f945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d9060600160405180910390a16040517f04c2a34a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152600091908916906304c2a34a906024016020604051808303816000875af115801561093e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109629190611644565b905073ffffffffffffffffffffffffffffffffffffffff81166109c9576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8616600482015260240161031f565b6109eb73ffffffffffffffffffffffffffffffffffffffff8616838387610d7e565b505b60005b8651811015610b72576000878281518110610a0d57610a0d611661565b60209081029190910101516040517f04c2a34a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301529192506000918b16906304c2a34a906024016020604051808303816000875af1158015610a8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab09190611644565b905073ffffffffffffffffffffffffffffffffffffffff8116610b17576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015260240161031f565b610b5d84828a8681518110610b2e57610b2e611661565b60200260200101518573ffffffffffffffffffffffffffffffffffffffff16610d7e909392919063ffffffff16565b50508080610b6a90611690565b9150506109f0565b505095945050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603610bfc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161031f565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000610cd4826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16610de29092919063ffffffff16565b80519091501561078f5780806020019051810190610cf291906116c8565b61078f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161031f565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052610ddc9085907f23b872dd000000000000000000000000000000000000000000000000000000009060840161070d565b50505050565b6060610df18484600085610dfb565b90505b9392505050565b606082471015610e8d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161031f565b843b610ef5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161031f565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610f1e91906116ea565b60006040518083038185875af1925050503d8060008114610f5b576040519150601f19603f3d011682016040523d82523d6000602084013e610f60565b606091505b5091509150610f70828286610f7b565b979650505050505050565b60608315610f8a575081610df4565b825115610f9a5782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031f9190611044565b60005b83811015610fe9578181015183820152602001610fd1565b83811115610ddc5750506000910152565b60008151808452611012816020860160208601610fce565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610df46020830184610ffa565b60006020828403121561106957600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461063557600080fd5b803561109d81611070565b919050565b6000806000606084860312156110b757600080fd5b83356110c281611070565b925060208401356110d281611070565b929592945050506040919091013590565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715611135576111356110e3565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611182576111826110e3565b604052919050565b600082601f83011261119b57600080fd5b813567ffffffffffffffff8111156111b5576111b56110e3565b6111e660207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161113b565b8181528460208386010111156111fb57600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115611232576112326110e3565b5060051b60200190565b600082601f83011261124d57600080fd5b8135602061126261125d83611218565b61113b565b82815260059290921b8401810191818101908684111561128157600080fd5b8286015b848110156112a557803561129881611070565b8352918301918301611285565b509695505050505050565b600082601f8301126112c157600080fd5b813560206112d161125d83611218565b82815260059290921b840181019181810190868411156112f057600080fd5b8286015b848110156112a557803583529183019183016112f4565b6000806040838503121561131e57600080fd5b82359150602083013567ffffffffffffffff8082111561133d57600080fd5b9084019060e0828703121561135157600080fd5b611359611112565b61136283611092565b815260208301358281111561137657600080fd5b6113828882860161118a565b60208301525060408301358281111561139a57600080fd5b6113a68882860161123c565b6040830152506060830135828111156113be57600080fd5b6113ca888286016112b0565b6060830152506113dc60808401611092565b608082015260a083013560a082015260c083013560c08201528093505050509250929050565b6000806040838503121561141557600080fd5b82359150602083013561142781611070565b809150509250929050565b60006020828403121561144457600080fd5b8135610df481611070565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156114905761149061144f565b500390565b600081518084526020808501945080840160005b838110156114c5578151875295820195908201906001016114a9565b509495945050505050565b60408152600073ffffffffffffffffffffffffffffffffffffffff80855116604084015260208086015160e0606086015261150f610120860182610ffa565b60408801517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087830381016080890152815180845291850193506000929091908501905b8084101561157557845187168252938501936001939093019290850190611553565b5060608a01519550818882030160a08901526115918187611495565b95505050505060808601516115be60c086018273ffffffffffffffffffffffffffffffffffffffff169052565b5060a086015160e085015260c08601516101008501528192506115f88185018673ffffffffffffffffffffffffffffffffffffffff169052565b50509392505050565b60006020828403121561161357600080fd5b815167ffffffffffffffff81168114610df457600080fd5b60006020828403121561163d57600080fd5b5051919050565b60006020828403121561165657600080fd5b8151610df481611070565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036116c1576116c161144f565b5060010190565b6000602082840312156116da57600080fd5b81518015158114610df457600080fd5b600082516116fc818460208701610fce565b919091019291505056fea164736f6c634300080d000a",
}

var OnRampRouterABI = OnRampRouterMetaData.ABI

var OnRampRouterBin = OnRampRouterMetaData.Bin

func DeployOnRampRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OnRampRouter, error) {
	parsed, err := OnRampRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnRampRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnRampRouter{OnRampRouterCaller: OnRampRouterCaller{contract: contract}, OnRampRouterTransactor: OnRampRouterTransactor{contract: contract}, OnRampRouterFilterer: OnRampRouterFilterer{contract: contract}}, nil
}

type OnRampRouter struct {
	address common.Address
	abi     abi.ABI
	OnRampRouterCaller
	OnRampRouterTransactor
	OnRampRouterFilterer
}

type OnRampRouterCaller struct {
	contract *bind.BoundContract
}

type OnRampRouterTransactor struct {
	contract *bind.BoundContract
}

type OnRampRouterFilterer struct {
	contract *bind.BoundContract
}

type OnRampRouterSession struct {
	Contract     *OnRampRouter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OnRampRouterCallerSession struct {
	Contract *OnRampRouterCaller
	CallOpts bind.CallOpts
}

type OnRampRouterTransactorSession struct {
	Contract     *OnRampRouterTransactor
	TransactOpts bind.TransactOpts
}

type OnRampRouterRaw struct {
	Contract *OnRampRouter
}

type OnRampRouterCallerRaw struct {
	Contract *OnRampRouterCaller
}

type OnRampRouterTransactorRaw struct {
	Contract *OnRampRouterTransactor
}

func NewOnRampRouter(address common.Address, backend bind.ContractBackend) (*OnRampRouter, error) {
	abi, err := abi.JSON(strings.NewReader(OnRampRouterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOnRampRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnRampRouter{address: address, abi: abi, OnRampRouterCaller: OnRampRouterCaller{contract: contract}, OnRampRouterTransactor: OnRampRouterTransactor{contract: contract}, OnRampRouterFilterer: OnRampRouterFilterer{contract: contract}}, nil
}

func NewOnRampRouterCaller(address common.Address, caller bind.ContractCaller) (*OnRampRouterCaller, error) {
	contract, err := bindOnRampRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampRouterCaller{contract: contract}, nil
}

func NewOnRampRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*OnRampRouterTransactor, error) {
	contract, err := bindOnRampRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampRouterTransactor{contract: contract}, nil
}

func NewOnRampRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*OnRampRouterFilterer, error) {
	contract, err := bindOnRampRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnRampRouterFilterer{contract: contract}, nil
}

func bindOnRampRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OnRampRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_OnRampRouter *OnRampRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRampRouter.Contract.OnRampRouterCaller.contract.Call(opts, result, method, params...)
}

func (_OnRampRouter *OnRampRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRampRouter.Contract.OnRampRouterTransactor.contract.Transfer(opts)
}

func (_OnRampRouter *OnRampRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRampRouter.Contract.OnRampRouterTransactor.contract.Transact(opts, method, params...)
}

func (_OnRampRouter *OnRampRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRampRouter.Contract.contract.Call(opts, result, method, params...)
}

func (_OnRampRouter *OnRampRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRampRouter.Contract.contract.Transfer(opts)
}

func (_OnRampRouter *OnRampRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRampRouter.Contract.contract.Transact(opts, method, params...)
}

func (_OnRampRouter *OnRampRouterCaller) GetOnRamp(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OnRampRouter.contract.Call(opts, &out, "getOnRamp", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRampRouter *OnRampRouterSession) GetOnRamp(chainId *big.Int) (common.Address, error) {
	return _OnRampRouter.Contract.GetOnRamp(&_OnRampRouter.CallOpts, chainId)
}

func (_OnRampRouter *OnRampRouterCallerSession) GetOnRamp(chainId *big.Int) (common.Address, error) {
	return _OnRampRouter.Contract.GetOnRamp(&_OnRampRouter.CallOpts, chainId)
}

func (_OnRampRouter *OnRampRouterCaller) IsChainSupported(opts *bind.CallOpts, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _OnRampRouter.contract.Call(opts, &out, "isChainSupported", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OnRampRouter *OnRampRouterSession) IsChainSupported(chainId *big.Int) (bool, error) {
	return _OnRampRouter.Contract.IsChainSupported(&_OnRampRouter.CallOpts, chainId)
}

func (_OnRampRouter *OnRampRouterCallerSession) IsChainSupported(chainId *big.Int) (bool, error) {
	return _OnRampRouter.Contract.IsChainSupported(&_OnRampRouter.CallOpts, chainId)
}

func (_OnRampRouter *OnRampRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnRampRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRampRouter *OnRampRouterSession) Owner() (common.Address, error) {
	return _OnRampRouter.Contract.Owner(&_OnRampRouter.CallOpts)
}

func (_OnRampRouter *OnRampRouterCallerSession) Owner() (common.Address, error) {
	return _OnRampRouter.Contract.Owner(&_OnRampRouter.CallOpts)
}

func (_OnRampRouter *OnRampRouterCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OnRampRouter.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OnRampRouter *OnRampRouterSession) TypeAndVersion() (string, error) {
	return _OnRampRouter.Contract.TypeAndVersion(&_OnRampRouter.CallOpts)
}

func (_OnRampRouter *OnRampRouterCallerSession) TypeAndVersion() (string, error) {
	return _OnRampRouter.Contract.TypeAndVersion(&_OnRampRouter.CallOpts)
}

func (_OnRampRouter *OnRampRouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRampRouter.contract.Transact(opts, "acceptOwnership")
}

func (_OnRampRouter *OnRampRouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRampRouter.Contract.AcceptOwnership(&_OnRampRouter.TransactOpts)
}

func (_OnRampRouter *OnRampRouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRampRouter.Contract.AcceptOwnership(&_OnRampRouter.TransactOpts)
}

func (_OnRampRouter *OnRampRouterTransactor) CcipSend(opts *bind.TransactOpts, destinationChainId *big.Int, message CCIPEVMToAnyTollMessage) (*types.Transaction, error) {
	return _OnRampRouter.contract.Transact(opts, "ccipSend", destinationChainId, message)
}

func (_OnRampRouter *OnRampRouterSession) CcipSend(destinationChainId *big.Int, message CCIPEVMToAnyTollMessage) (*types.Transaction, error) {
	return _OnRampRouter.Contract.CcipSend(&_OnRampRouter.TransactOpts, destinationChainId, message)
}

func (_OnRampRouter *OnRampRouterTransactorSession) CcipSend(destinationChainId *big.Int, message CCIPEVMToAnyTollMessage) (*types.Transaction, error) {
	return _OnRampRouter.Contract.CcipSend(&_OnRampRouter.TransactOpts, destinationChainId, message)
}

func (_OnRampRouter *OnRampRouterTransactor) SetOnRamp(opts *bind.TransactOpts, chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
	return _OnRampRouter.contract.Transact(opts, "setOnRamp", chainId, onRamp)
}

func (_OnRampRouter *OnRampRouterSession) SetOnRamp(chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
	return _OnRampRouter.Contract.SetOnRamp(&_OnRampRouter.TransactOpts, chainId, onRamp)
}

func (_OnRampRouter *OnRampRouterTransactorSession) SetOnRamp(chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
	return _OnRampRouter.Contract.SetOnRamp(&_OnRampRouter.TransactOpts, chainId, onRamp)
}

func (_OnRampRouter *OnRampRouterTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OnRampRouter.contract.Transact(opts, "transferOwnership", to)
}

func (_OnRampRouter *OnRampRouterSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRampRouter.Contract.TransferOwnership(&_OnRampRouter.TransactOpts, to)
}

func (_OnRampRouter *OnRampRouterTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRampRouter.Contract.TransferOwnership(&_OnRampRouter.TransactOpts, to)
}

func (_OnRampRouter *OnRampRouterTransactor) WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OnRampRouter.contract.Transact(opts, "withdrawAccumulatedFees", feeToken, recipient, amount)
}

func (_OnRampRouter *OnRampRouterSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OnRampRouter.Contract.WithdrawAccumulatedFees(&_OnRampRouter.TransactOpts, feeToken, recipient, amount)
}

func (_OnRampRouter *OnRampRouterTransactorSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OnRampRouter.Contract.WithdrawAccumulatedFees(&_OnRampRouter.TransactOpts, feeToken, recipient, amount)
}

type OnRampRouterFeeChargedIterator struct {
	Event *OnRampRouterFeeCharged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampRouterFeeChargedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampRouterFeeCharged)
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
		it.Event = new(OnRampRouterFeeCharged)
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

func (it *OnRampRouterFeeChargedIterator) Error() error {
	return it.fail
}

func (it *OnRampRouterFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampRouterFeeCharged struct {
	From common.Address
	To   common.Address
	Fee  *big.Int
	Raw  types.Log
}

func (_OnRampRouter *OnRampRouterFilterer) FilterFeeCharged(opts *bind.FilterOpts) (*OnRampRouterFeeChargedIterator, error) {

	logs, sub, err := _OnRampRouter.contract.FilterLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return &OnRampRouterFeeChargedIterator{contract: _OnRampRouter.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

func (_OnRampRouter *OnRampRouterFilterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *OnRampRouterFeeCharged) (event.Subscription, error) {

	logs, sub, err := _OnRampRouter.contract.WatchLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampRouterFeeCharged)
				if err := _OnRampRouter.contract.UnpackLog(event, "FeeCharged", log); err != nil {
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

func (_OnRampRouter *OnRampRouterFilterer) ParseFeeCharged(log types.Log) (*OnRampRouterFeeCharged, error) {
	event := new(OnRampRouterFeeCharged)
	if err := _OnRampRouter.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampRouterFeesWithdrawnIterator struct {
	Event *OnRampRouterFeesWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampRouterFeesWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampRouterFeesWithdrawn)
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
		it.Event = new(OnRampRouterFeesWithdrawn)
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

func (it *OnRampRouterFeesWithdrawnIterator) Error() error {
	return it.fail
}

func (it *OnRampRouterFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampRouterFeesWithdrawn struct {
	FeeToken  common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_OnRampRouter *OnRampRouterFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*OnRampRouterFeesWithdrawnIterator, error) {

	logs, sub, err := _OnRampRouter.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &OnRampRouterFeesWithdrawnIterator{contract: _OnRampRouter.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

func (_OnRampRouter *OnRampRouterFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampRouterFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _OnRampRouter.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampRouterFeesWithdrawn)
				if err := _OnRampRouter.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

func (_OnRampRouter *OnRampRouterFilterer) ParseFeesWithdrawn(log types.Log) (*OnRampRouterFeesWithdrawn, error) {
	event := new(OnRampRouterFeesWithdrawn)
	if err := _OnRampRouter.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampRouterOnRampSetIterator struct {
	Event *OnRampRouterOnRampSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampRouterOnRampSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampRouterOnRampSet)
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
		it.Event = new(OnRampRouterOnRampSet)
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

func (it *OnRampRouterOnRampSetIterator) Error() error {
	return it.fail
}

func (it *OnRampRouterOnRampSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampRouterOnRampSet struct {
	ChainId *big.Int
	OnRamp  common.Address
	Raw     types.Log
}

func (_OnRampRouter *OnRampRouterFilterer) FilterOnRampSet(opts *bind.FilterOpts, chainId []*big.Int, onRamp []common.Address) (*OnRampRouterOnRampSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var onRampRule []interface{}
	for _, onRampItem := range onRamp {
		onRampRule = append(onRampRule, onRampItem)
	}

	logs, sub, err := _OnRampRouter.contract.FilterLogs(opts, "OnRampSet", chainIdRule, onRampRule)
	if err != nil {
		return nil, err
	}
	return &OnRampRouterOnRampSetIterator{contract: _OnRampRouter.contract, event: "OnRampSet", logs: logs, sub: sub}, nil
}

func (_OnRampRouter *OnRampRouterFilterer) WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *OnRampRouterOnRampSet, chainId []*big.Int, onRamp []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var onRampRule []interface{}
	for _, onRampItem := range onRamp {
		onRampRule = append(onRampRule, onRampItem)
	}

	logs, sub, err := _OnRampRouter.contract.WatchLogs(opts, "OnRampSet", chainIdRule, onRampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampRouterOnRampSet)
				if err := _OnRampRouter.contract.UnpackLog(event, "OnRampSet", log); err != nil {
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

func (_OnRampRouter *OnRampRouterFilterer) ParseOnRampSet(log types.Log) (*OnRampRouterOnRampSet, error) {
	event := new(OnRampRouterOnRampSet)
	if err := _OnRampRouter.contract.UnpackLog(event, "OnRampSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampRouterOwnershipTransferRequestedIterator struct {
	Event *OnRampRouterOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampRouterOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampRouterOwnershipTransferRequested)
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
		it.Event = new(OnRampRouterOwnershipTransferRequested)
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

func (it *OnRampRouterOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OnRampRouterOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampRouterOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRampRouter *OnRampRouterFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampRouterOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampRouter.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampRouterOwnershipTransferRequestedIterator{contract: _OnRampRouter.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OnRampRouter *OnRampRouterFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampRouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampRouter.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampRouterOwnershipTransferRequested)
				if err := _OnRampRouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OnRampRouter *OnRampRouterFilterer) ParseOwnershipTransferRequested(log types.Log) (*OnRampRouterOwnershipTransferRequested, error) {
	event := new(OnRampRouterOwnershipTransferRequested)
	if err := _OnRampRouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampRouterOwnershipTransferredIterator struct {
	Event *OnRampRouterOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampRouterOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampRouterOwnershipTransferred)
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
		it.Event = new(OnRampRouterOwnershipTransferred)
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

func (it *OnRampRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OnRampRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampRouterOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRampRouter *OnRampRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampRouterOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampRouter.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampRouterOwnershipTransferredIterator{contract: _OnRampRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OnRampRouter *OnRampRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampRouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampRouter.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampRouterOwnershipTransferred)
				if err := _OnRampRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OnRampRouter *OnRampRouterFilterer) ParseOwnershipTransferred(log types.Log) (*OnRampRouterOwnershipTransferred, error) {
	event := new(OnRampRouterOwnershipTransferred)
	if err := _OnRampRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_OnRampRouter *OnRampRouter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OnRampRouter.abi.Events["FeeCharged"].ID:
		return _OnRampRouter.ParseFeeCharged(log)
	case _OnRampRouter.abi.Events["FeesWithdrawn"].ID:
		return _OnRampRouter.ParseFeesWithdrawn(log)
	case _OnRampRouter.abi.Events["OnRampSet"].ID:
		return _OnRampRouter.ParseOnRampSet(log)
	case _OnRampRouter.abi.Events["OwnershipTransferRequested"].ID:
		return _OnRampRouter.ParseOwnershipTransferRequested(log)
	case _OnRampRouter.abi.Events["OwnershipTransferred"].ID:
		return _OnRampRouter.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (OnRampRouterFeeCharged) Topic() common.Hash {
	return common.HexToHash("0x945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d")
}

func (OnRampRouterFeesWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8")
}

func (OnRampRouterOnRampSet) Topic() common.Hash {
	return common.HexToHash("0x4b680ef9fa79bb5f36e7559d7b33fd57a8336f78cc120c8cd93333b5ade624cb")
}

func (OnRampRouterOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OnRampRouterOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_OnRampRouter *OnRampRouter) Address() common.Address {
	return _OnRampRouter.address
}

type OnRampRouterInterface interface {
	GetOnRamp(opts *bind.CallOpts, chainId *big.Int) (common.Address, error)

	IsChainSupported(opts *bind.CallOpts, chainId *big.Int) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainId *big.Int, message CCIPEVMToAnyTollMessage) (*types.Transaction, error)

	SetOnRamp(opts *bind.TransactOpts, chainId *big.Int, onRamp common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	FilterFeeCharged(opts *bind.FilterOpts) (*OnRampRouterFeeChargedIterator, error)

	WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *OnRampRouterFeeCharged) (event.Subscription, error)

	ParseFeeCharged(log types.Log) (*OnRampRouterFeeCharged, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*OnRampRouterFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampRouterFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*OnRampRouterFeesWithdrawn, error)

	FilterOnRampSet(opts *bind.FilterOpts, chainId []*big.Int, onRamp []common.Address) (*OnRampRouterOnRampSetIterator, error)

	WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *OnRampRouterOnRampSet, chainId []*big.Int, onRamp []common.Address) (event.Subscription, error)

	ParseOnRampSet(log types.Log) (*OnRampRouterOnRampSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampRouterOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampRouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OnRampRouterOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampRouterOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampRouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OnRampRouterOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
