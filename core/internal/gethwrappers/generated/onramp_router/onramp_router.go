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

type CCIPMessagePayload struct {
	Tokens             []common.Address
	Amounts            []*big.Int
	DestinationChainId *big.Int
	Receiver           common.Address
	Executor           common.Address
	Data               []byte
}

var OnRampRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"contractOnRampInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"requestCrossChainSend\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"setOnRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6117ce806101576000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80637ce6855811610076578063d8a98f8c1161005b578063d8a98f8c146101c0578063f1927cae146101f6578063f2fde38b1461020957600080fd5b80637ce68558146101555780638da5cb5b1461018157600080fd5b8063181f5a77146100a85780635221c1f0146100f057806359e96b5b1461013857806379ba50971461014d575b600080fd5b604080518082018252601281527f4f6e52616d70526f7574657220302e302e310000000000000000000000000000602082015290516100e79190610efb565b60405180910390f35b6101286100fe366004610f4c565b60009081526002602052604090205473ffffffffffffffffffffffffffffffffffffffff16151590565b60405190151581526020016100e7565b61014b610146366004610f97565b61021c565b005b61014b6102a2565b610168610163366004610fd8565b6103a4565b60405167ffffffffffffffff90911681526020016100e7565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e7565b61019b6101ce366004610f4c565b60009081526002602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b61014b610204366004611013565b610510565b61014b610217366004611043565b610615565b610224610629565b61024573ffffffffffffffffffffffffffffffffffffffff841683836106ac565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa89060600160405180910390a1505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610328576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6040808201356000908152600260205290812054339073ffffffffffffffffffffffffffffffffffffffff168061040d57604080517f45abe4ae00000000000000000000000000000000000000000000000000000000815290850135600482015260240161031f565b61041a6020850185611060565b90506104268580611060565b90501461045f576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104718161046c866112f7565b610785565b6040517fd45cab5e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82169063d45cab5e906104c59087908690600401611518565b6020604051808303816000875af11580156104e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105089190611676565b949350505050565b610518610629565b60008281526002602052604090205473ffffffffffffffffffffffffffffffffffffffff808316911603610597576040517fe31de3b20000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff8216602482015260440161031f565b60008281526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091559051909184917f4b680ef9fa79bb5f36e7559d7b33fd57a8336f78cc120c8cd93333b5ade624cb9190a35050565b61061d610629565b61062681610a7e565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146106aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161031f565b565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526107809084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610b73565b505050565b805180513391600091829061079c5761079c6116a0565b60209081029190910101516040517fd0d5de6100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808316600483015291925060009186169063d0d5de61906024016020604051808303816000875af115801561081b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083f91906116cf565b905061086373ffffffffffffffffffffffffffffffffffffffff8316843084610c7f565b80846020015160008151811061087b5761087b6116a0565b6020026020010181815161088f9190611717565b9052506040805173ffffffffffffffffffffffffffffffffffffffff851681523060208201529081018290527f945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d9060600160405180910390a160005b845151811015610a765760008560000151828151811061090d5761090d6116a0565b60209081029190910101516040517f04c2a34a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301529192506000918916906304c2a34a906024016020604051808303816000875af115801561098c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b0919061172e565b905073ffffffffffffffffffffffffffffffffffffffff8116610a17576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015260240161031f565b610a61868289602001518681518110610a3257610a326116a0565b60200260200101518573ffffffffffffffffffffffffffffffffffffffff16610c7f909392919063ffffffff16565b50508080610a6e9061174b565b9150506108eb565b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603610afd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161031f565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000610bd5826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16610ce39092919063ffffffff16565b8051909150156107805780806020019051810190610bf39190611783565b610780576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161031f565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052610cdd9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016106fe565b50505050565b6060610cf28484600085610cfc565b90505b9392505050565b606082471015610d8e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161031f565b843b610df6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161031f565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610e1f91906117a5565b60006040518083038185875af1925050503d8060008114610e5c576040519150601f19603f3d011682016040523d82523d6000602084013e610e61565b606091505b5091509150610e71828286610e7c565b979650505050505050565b60608315610e8b575081610cf5565b825115610e9b5782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031f9190610efb565b60005b83811015610eea578181015183820152602001610ed2565b83811115610cdd5750506000910152565b6020815260008251806020840152610f1a816040850160208701610ecf565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b600060208284031215610f5e57600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461062657600080fd5b8035610f9281610f65565b919050565b600080600060608486031215610fac57600080fd5b8335610fb781610f65565b92506020840135610fc781610f65565b929592945050506040919091013590565b600060208284031215610fea57600080fd5b813567ffffffffffffffff81111561100157600080fd5b820160c08185031215610cf557600080fd5b6000806040838503121561102657600080fd5b82359150602083013561103881610f65565b809150509250929050565b60006020828403121561105557600080fd5b8135610cf581610f65565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261109557600080fd5b83018035915067ffffffffffffffff8211156110b057600080fd5b6020019150600581901b36038213156110c857600080fd5b9250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160c0810167ffffffffffffffff81118282101715611121576111216110cf565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561116e5761116e6110cf565b604052919050565b600067ffffffffffffffff821115611190576111906110cf565b5060051b60200190565b600082601f8301126111ab57600080fd5b813560206111c06111bb83611176565b611127565b82815260059290921b840181019181810190868411156111df57600080fd5b8286015b848110156112035780356111f681610f65565b83529183019183016111e3565b509695505050505050565b600082601f83011261121f57600080fd5b8135602061122f6111bb83611176565b82815260059290921b8401810191818101908684111561124e57600080fd5b8286015b848110156112035780358352918301918301611252565b600082601f83011261127a57600080fd5b813567ffffffffffffffff811115611294576112946110cf565b6112c560207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611127565b8181528460208386010111156112da57600080fd5b816020850160208301376000918101602001919091529392505050565b600060c0823603121561130957600080fd5b6113116110fe565b823567ffffffffffffffff8082111561132957600080fd5b6113353683870161119a565b8352602085013591508082111561134b57600080fd5b6113573683870161120e565b60208401526040850135604084015261137260608601610f87565b606084015261138360808601610f87565b608084015260a085013591508082111561139c57600080fd5b506113a936828601611269565b60a08301525092915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126113ea57600080fd5b830160208101925035905067ffffffffffffffff81111561140a57600080fd5b8060051b36038313156110c857600080fd5b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83111561144e57600080fd5b8260051b8083602087013760009401602001938452509192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126114a057600080fd5b830160208101925035905067ffffffffffffffff8111156114c057600080fd5b8036038313156110c857600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000610100820161152d85866113b5565b60c06040860152918290529060009061012085015b8183101561158057833561155581610f65565b73ffffffffffffffffffffffffffffffffffffffff1681526020938401936001939093019201611542565b61158d60208901896113b5565b945092507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc09150818682030160608701526115c981858561141c565b935050604087013560808601526115e260608801610f87565b73ffffffffffffffffffffffffffffffffffffffff811660a0870152915061160c60808801610f87565b73ffffffffffffffffffffffffffffffffffffffff811660c0870152915061163760a088018861146b565b9250818685030160e087015261164e8484836114cf565b945050505050610cf5602083018473ffffffffffffffffffffffffffffffffffffffff169052565b60006020828403121561168857600080fd5b815167ffffffffffffffff81168114610cf557600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156116e157600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611729576117296116e8565b500390565b60006020828403121561174057600080fd5b8151610cf581610f65565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361177c5761177c6116e8565b5060010190565b60006020828403121561179557600080fd5b81518015158114610cf557600080fd5b600082516117b7818460208701610ecf565b919091019291505056fea164736f6c634300080d000a",
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

func (_OnRampRouter *OnRampRouterTransactor) RequestCrossChainSend(opts *bind.TransactOpts, payload CCIPMessagePayload) (*types.Transaction, error) {
	return _OnRampRouter.contract.Transact(opts, "requestCrossChainSend", payload)
}

func (_OnRampRouter *OnRampRouterSession) RequestCrossChainSend(payload CCIPMessagePayload) (*types.Transaction, error) {
	return _OnRampRouter.Contract.RequestCrossChainSend(&_OnRampRouter.TransactOpts, payload)
}

func (_OnRampRouter *OnRampRouterTransactorSession) RequestCrossChainSend(payload CCIPMessagePayload) (*types.Transaction, error) {
	return _OnRampRouter.Contract.RequestCrossChainSend(&_OnRampRouter.TransactOpts, payload)
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

	RequestCrossChainSend(opts *bind.TransactOpts, payload CCIPMessagePayload) (*types.Transaction, error)

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
