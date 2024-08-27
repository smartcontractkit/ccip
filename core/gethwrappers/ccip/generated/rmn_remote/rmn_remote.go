// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rmn_remote

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

type MerkleRoot struct {
	SourceChainSelector uint64
	MinSeqNr            uint64
	MaxSeqNr            uint64
	MerkleRoot          [32]byte
	OnRampAddress       []byte
}

type RMNRemoteConfig struct {
	RmnHomeContractConfigDigest [32]byte
	Signers                     []RMNRemoteSigner
	MinSigners                  uint64
}

type RMNRemoteSigner struct {
	OnchainPublicKey common.Address
	NodeIndex        uint64
}

type RMNRemoteVersionedConfig struct {
	Version uint32
	Config  RMNRemoteConfig
}

type Signature struct {
	R [32]byte
	S [32]byte
}

var RMNRemoteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ConfigNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateOnchainPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignerOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinSignersTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfOrderSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedSigner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rmnHomeContractConfigDigest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"onchainPublicKey\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nodeIndex\",\"type\":\"uint64\"}],\"internalType\":\"structRMNRemote.Signer[]\",\"name\":\"signers\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"minSigners\",\"type\":\"uint64\"}],\"internalType\":\"structRMNRemote.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structRMNRemote.VersionedConfig\",\"name\":\"versionedConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersionedConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rmnHomeContractConfigDigest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"onchainPublicKey\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nodeIndex\",\"type\":\"uint64\"}],\"internalType\":\"structRMNRemote.Signer[]\",\"name\":\"signers\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"minSigners\",\"type\":\"uint64\"}],\"internalType\":\"structRMNRemote.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"internalType\":\"structRMNRemote.VersionedConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rmnHomeContractConfigDigest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"onchainPublicKey\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nodeIndex\",\"type\":\"uint64\"}],\"internalType\":\"structRMNRemote.Signer[]\",\"name\":\"signers\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"minSigners\",\"type\":\"uint64\"}],\"internalType\":\"structRMNRemote.Config\",\"name\":\"newConfig\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"onRampAddress\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleRoot[]\",\"name\":\"destLaneUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200178a3803806200178a83398101604081905262000034916200017e565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000d3565b5050506001600160401b0316608052620001b0565b336001600160a01b038216036200012d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000602082840312156200019157600080fd5b81516001600160401b0381168114620001a957600080fd5b9392505050565b6080516115c1620001c9600039600050506115c16000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806379ba50971161005b57806379ba50971461011b5780638da5cb5b14610123578063f2fde38b1461014b578063f8a8fd6d146100f257600080fd5b8063181f5a771461008d578063198f0f77146100df5780631add205f146100f457806330dfc30814610109575b600080fd5b6100c96040518060400160405280601381526020017f524d4e52656d6f746520312e362e302d6465760000000000000000000000000081525081565b6040516100d69190610bea565b60405180910390f35b6100f26100ed366004610c04565b61015e565b005b6100fc61058f565b6040516100d69190610c3f565b6100f2610117366004610ebd565b5050565b6100f26108aa565b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100d6565b6100f261015936600461109a565b6109ac565b6101666109c2565b60015b61017660208301836110b7565b90508110156102465761018c60208301836110b7565b8281811061019c5761019c611126565b90506040020160200160208101906101b49190611155565b67ffffffffffffffff166101cb60208401846110b7565b6101d66001856111a1565b8181106101e5576101e5611126565b90506040020160200160208101906101fd9190611155565b67ffffffffffffffff161061023e576040517f4485151700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600101610169565b5061025460208201826110b7565b90506102666060830160408401611155565b67ffffffffffffffff1611156102a8576040517ffba0d9e600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025b6001810154156103a45760018082018054600692600092916102cd91906111a1565b815481106102dd576102dd611126565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905560018101805480610347576103476111ba565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffff000000000000000000000000000000000000000000000000000000001690550190556102ab565b5060005b6103b560208301836110b7565b90508110156104ea57600660006103cf60208501856110b7565b848181106103df576103df611126565b6103f5926020604090920201908101915061109a565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040016000205460ff1615610456576040517f28cae27d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016006600061046960208601866110b7565b8581811061047957610479611126565b61048f926020604090920201908101915061109a565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790556001016103a8565b508060026104f882826112a2565b5050600580546000919082906105139063ffffffff166113dd565b91906101000a81548163ffffffff021916908363ffffffff160217905590507f6cc65868ae41a007e6c3ed18ce591c123dd4e5864b421888c68ce92dae98cea460405180604001604052808363ffffffff1681526020018461057490611400565b90526040516105839190610c3f565b60405180910390a15050565b610597610b38565b60408051808201825260055463ffffffff1681528151606081018352600280548252600380548551602082810282018101909752818152949580870195858201939092909160009084015b82821015610650576000848152602090819020604080518082019091529084015473ffffffffffffffffffffffffffffffffffffffff8116825274010000000000000000000000000000000000000000900467ffffffffffffffff16818301528252600190920191016105e2565b505050908252506002919091015467ffffffffffffffff166020909101529052919050565b845181101561085c57600085828151811061069257610692611126565b602002602001015190506000600186601b84600001518560200151604051600081526020016040526040516106e3949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610705573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661077d576040517f8baa579f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16106107e2576040517fbbe15e7f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811660009081526006602052604090205460ff16610841576040517faaaa914100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b92508261084d8561157c565b94505050806001019050610675565b5060045467ffffffffffffffff168210156108a3576040517f59fa4a9300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610930576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6109b46109c2565b6109bd81610a43565b50565b565b60005473ffffffffffffffffffffffffffffffffffffffff1633146109c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610927565b3373ffffffffffffffffffffffffffffffffffffffff821603610ac2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610927565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040518060400160405280600063ffffffff168152602001610b8160405180606001604052806000801916815260200160608152602001600067ffffffffffffffff1681525090565b905290565b6000815180845260005b81811015610bac57602081850181015186830182015201610b90565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610bfd6020830184610b86565b9392505050565b600060208284031215610c1657600080fd5b813567ffffffffffffffff811115610c2d57600080fd5b820160608185031215610bfd57600080fd5b6000602080835263ffffffff8451168184015280840151604080604086015260c0850182516060870152838301516060608088015281815180845260e0890191508683019350600092505b80831015610cd3578351805173ffffffffffffffffffffffffffffffffffffffff16835287015167ffffffffffffffff1687830152928601926001929092019190840190610c8a565b50604085015167ffffffffffffffff811660a08a0152955098975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610d4957610d49610cf7565b60405290565b60405160a0810167ffffffffffffffff81118282101715610d4957610d49610cf7565b6040516060810167ffffffffffffffff81118282101715610d4957610d49610cf7565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610ddc57610ddc610cf7565b604052919050565b600067ffffffffffffffff821115610dfe57610dfe610cf7565b5060051b60200190565b67ffffffffffffffff811681146109bd57600080fd5b8035610e2981610e08565b919050565b600082601f830112610e3f57600080fd5b81356020610e54610e4f83610de4565b610d95565b82815260069290921b84018101918181019086841115610e7357600080fd5b8286015b84811015610eb25760408189031215610e905760008081fd5b610e98610d26565b813581528482013585820152835291830191604001610e77565b509695505050505050565b60008060408385031215610ed057600080fd5b823567ffffffffffffffff80821115610ee857600080fd5b818501915085601f830112610efc57600080fd5b81356020610f0c610e4f83610de4565b82815260059290921b84018101918181019089841115610f2b57600080fd5b8286015b8481101561104a57803586811115610f4657600080fd5b87017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060a0828e0382011215610f7b57600080fd5b610f83610d4f565b86830135610f9081610e08565b81526040830135610fa081610e08565b818801526060830135610fb281610e08565b60408201526080830135606082015260a083013589811115610fd357600080fd5b8084019350508d603f840112610fe857600080fd5b8683013589811115610ffc57610ffc610cf7565b61100c8884601f84011601610d95565b92508083528e604082860101111561102357600080fd5b80604085018985013760009083018801526080810191909152845250918301918301610f2f565b509650508601359250508082111561106157600080fd5b5061106e85828601610e2e565b9150509250929050565b73ffffffffffffffffffffffffffffffffffffffff811681146109bd57600080fd5b6000602082840312156110ac57600080fd5b8135610bfd81611078565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126110ec57600080fd5b83018035915067ffffffffffffffff82111561110757600080fd5b6020019150600681901b360382131561111f57600080fd5b9250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561116757600080fd5b8135610bfd81610e08565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156111b4576111b4611172565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081356111b481610e08565b813561120181611078565b73ffffffffffffffffffffffffffffffffffffffff811690508154817fffffffffffffffffffffffff00000000000000000000000000000000000000008216178355602084013561125181610e08565b7bffffffffffffffff00000000000000000000000000000000000000008160a01b16837fffffffff000000000000000000000000000000000000000000000000000000008416171784555050505050565b81358155600180820160208401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18536030181126112e057600080fd5b8401803567ffffffffffffffff8111156112f957600080fd5b6020820191508060061b360382131561131157600080fd5b6801000000000000000081111561132a5761132a610cf7565b82548184558082101561135f576000848152602081208381019083015b8082101561135b5782825590870190611347565b5050505b50600092835260208320925b8181101561138f5761137d83856111f6565b9284019260409290920191840161136b565b50505050506101176113a3604084016111e9565b6002830167ffffffffffffffff82167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008254161781555050565b600063ffffffff8083168181036113f6576113f6611172565b6001019392505050565b60006060823603121561141257600080fd5b61141a610d72565b8235815260208084013567ffffffffffffffff81111561143957600080fd5b840136601f82011261144a57600080fd5b8035611458610e4f82610de4565b81815260069190911b8201830190838101903683111561147757600080fd5b928401925b828410156114cd57604084360312156114955760008081fd5b61149d610d26565b84356114a881611078565b8152848601356114b781610e08565b818701528252604093909301929084019061147c565b808587015250505050506114e360408401610e1e565b604082015292915050565b81811015611569578c89037ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee00184528751805187168a528a81015187168b8b01528b81015187168c8b015287810151888b0152850151858a01849052611556848b0182610b86565b99505096890196928901926001016114ee565b50969d9c50505050505050505050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036115ad576115ad611172565b506001019056fea164736f6c6343000818000a",
}

var RMNRemoteABI = RMNRemoteMetaData.ABI

var RMNRemoteBin = RMNRemoteMetaData.Bin

func DeployRMNRemote(auth *bind.TransactOpts, backend bind.ContractBackend, chainSelector uint64) (common.Address, *types.Transaction, *RMNRemote, error) {
	parsed, err := RMNRemoteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RMNRemoteBin), backend, chainSelector)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RMNRemote{address: address, abi: *parsed, RMNRemoteCaller: RMNRemoteCaller{contract: contract}, RMNRemoteTransactor: RMNRemoteTransactor{contract: contract}, RMNRemoteFilterer: RMNRemoteFilterer{contract: contract}}, nil
}

type RMNRemote struct {
	address common.Address
	abi     abi.ABI
	RMNRemoteCaller
	RMNRemoteTransactor
	RMNRemoteFilterer
}

type RMNRemoteCaller struct {
	contract *bind.BoundContract
}

type RMNRemoteTransactor struct {
	contract *bind.BoundContract
}

type RMNRemoteFilterer struct {
	contract *bind.BoundContract
}

type RMNRemoteSession struct {
	Contract     *RMNRemote
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type RMNRemoteCallerSession struct {
	Contract *RMNRemoteCaller
	CallOpts bind.CallOpts
}

type RMNRemoteTransactorSession struct {
	Contract     *RMNRemoteTransactor
	TransactOpts bind.TransactOpts
}

type RMNRemoteRaw struct {
	Contract *RMNRemote
}

type RMNRemoteCallerRaw struct {
	Contract *RMNRemoteCaller
}

type RMNRemoteTransactorRaw struct {
	Contract *RMNRemoteTransactor
}

func NewRMNRemote(address common.Address, backend bind.ContractBackend) (*RMNRemote, error) {
	abi, err := abi.JSON(strings.NewReader(RMNRemoteABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindRMNRemote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RMNRemote{address: address, abi: abi, RMNRemoteCaller: RMNRemoteCaller{contract: contract}, RMNRemoteTransactor: RMNRemoteTransactor{contract: contract}, RMNRemoteFilterer: RMNRemoteFilterer{contract: contract}}, nil
}

func NewRMNRemoteCaller(address common.Address, caller bind.ContractCaller) (*RMNRemoteCaller, error) {
	contract, err := bindRMNRemote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RMNRemoteCaller{contract: contract}, nil
}

func NewRMNRemoteTransactor(address common.Address, transactor bind.ContractTransactor) (*RMNRemoteTransactor, error) {
	contract, err := bindRMNRemote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RMNRemoteTransactor{contract: contract}, nil
}

func NewRMNRemoteFilterer(address common.Address, filterer bind.ContractFilterer) (*RMNRemoteFilterer, error) {
	contract, err := bindRMNRemote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RMNRemoteFilterer{contract: contract}, nil
}

func bindRMNRemote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RMNRemoteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_RMNRemote *RMNRemoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RMNRemote.Contract.RMNRemoteCaller.contract.Call(opts, result, method, params...)
}

func (_RMNRemote *RMNRemoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RMNRemote.Contract.RMNRemoteTransactor.contract.Transfer(opts)
}

func (_RMNRemote *RMNRemoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RMNRemote.Contract.RMNRemoteTransactor.contract.Transact(opts, method, params...)
}

func (_RMNRemote *RMNRemoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RMNRemote.Contract.contract.Call(opts, result, method, params...)
}

func (_RMNRemote *RMNRemoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RMNRemote.Contract.contract.Transfer(opts)
}

func (_RMNRemote *RMNRemoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RMNRemote.Contract.contract.Transact(opts, method, params...)
}

func (_RMNRemote *RMNRemoteCaller) GetVersionedConfig(opts *bind.CallOpts) (RMNRemoteVersionedConfig, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "getVersionedConfig")

	if err != nil {
		return *new(RMNRemoteVersionedConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(RMNRemoteVersionedConfig)).(*RMNRemoteVersionedConfig)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) GetVersionedConfig() (RMNRemoteVersionedConfig, error) {
	return _RMNRemote.Contract.GetVersionedConfig(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) GetVersionedConfig() (RMNRemoteVersionedConfig, error) {
	return _RMNRemote.Contract.GetVersionedConfig(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) Owner() (common.Address, error) {
	return _RMNRemote.Contract.Owner(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) Owner() (common.Address, error) {
	return _RMNRemote.Contract.Owner(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) TypeAndVersion() (string, error) {
	return _RMNRemote.Contract.TypeAndVersion(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) TypeAndVersion() (string, error) {
	return _RMNRemote.Contract.TypeAndVersion(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCaller) Verify(opts *bind.CallOpts, destLaneUpdates []MerkleRoot, signatures []Signature) error {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "verify", destLaneUpdates, signatures)

	if err != nil {
		return err
	}

	return err

}

func (_RMNRemote *RMNRemoteSession) Verify(destLaneUpdates []MerkleRoot, signatures []Signature) error {
	return _RMNRemote.Contract.Verify(&_RMNRemote.CallOpts, destLaneUpdates, signatures)
}

func (_RMNRemote *RMNRemoteCallerSession) Verify(destLaneUpdates []MerkleRoot, signatures []Signature) error {
	return _RMNRemote.Contract.Verify(&_RMNRemote.CallOpts, destLaneUpdates, signatures)
}

func (_RMNRemote *RMNRemoteTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "acceptOwnership")
}

func (_RMNRemote *RMNRemoteSession) AcceptOwnership() (*types.Transaction, error) {
	return _RMNRemote.Contract.AcceptOwnership(&_RMNRemote.TransactOpts)
}

func (_RMNRemote *RMNRemoteTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _RMNRemote.Contract.AcceptOwnership(&_RMNRemote.TransactOpts)
}

func (_RMNRemote *RMNRemoteTransactor) SetConfig(opts *bind.TransactOpts, newConfig RMNRemoteConfig) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "setConfig", newConfig)
}

func (_RMNRemote *RMNRemoteSession) SetConfig(newConfig RMNRemoteConfig) (*types.Transaction, error) {
	return _RMNRemote.Contract.SetConfig(&_RMNRemote.TransactOpts, newConfig)
}

func (_RMNRemote *RMNRemoteTransactorSession) SetConfig(newConfig RMNRemoteConfig) (*types.Transaction, error) {
	return _RMNRemote.Contract.SetConfig(&_RMNRemote.TransactOpts, newConfig)
}

func (_RMNRemote *RMNRemoteTransactor) Test(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "test")
}

func (_RMNRemote *RMNRemoteSession) Test() (*types.Transaction, error) {
	return _RMNRemote.Contract.Test(&_RMNRemote.TransactOpts)
}

func (_RMNRemote *RMNRemoteTransactorSession) Test() (*types.Transaction, error) {
	return _RMNRemote.Contract.Test(&_RMNRemote.TransactOpts)
}

func (_RMNRemote *RMNRemoteTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "transferOwnership", to)
}

func (_RMNRemote *RMNRemoteSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RMNRemote.Contract.TransferOwnership(&_RMNRemote.TransactOpts, to)
}

func (_RMNRemote *RMNRemoteTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RMNRemote.Contract.TransferOwnership(&_RMNRemote.TransactOpts, to)
}

type RMNRemoteConfigSetIterator struct {
	Event *RMNRemoteConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNRemoteConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNRemoteConfigSet)
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
		it.Event = new(RMNRemoteConfigSet)
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

func (it *RMNRemoteConfigSetIterator) Error() error {
	return it.fail
}

func (it *RMNRemoteConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNRemoteConfigSet struct {
	VersionedConfig RMNRemoteVersionedConfig
	Raw             types.Log
}

func (_RMNRemote *RMNRemoteFilterer) FilterConfigSet(opts *bind.FilterOpts) (*RMNRemoteConfigSetIterator, error) {

	logs, sub, err := _RMNRemote.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &RMNRemoteConfigSetIterator{contract: _RMNRemote.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_RMNRemote *RMNRemoteFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *RMNRemoteConfigSet) (event.Subscription, error) {

	logs, sub, err := _RMNRemote.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNRemoteConfigSet)
				if err := _RMNRemote.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_RMNRemote *RMNRemoteFilterer) ParseConfigSet(log types.Log) (*RMNRemoteConfigSet, error) {
	event := new(RMNRemoteConfigSet)
	if err := _RMNRemote.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNRemoteOwnershipTransferRequestedIterator struct {
	Event *RMNRemoteOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNRemoteOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNRemoteOwnershipTransferRequested)
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
		it.Event = new(RMNRemoteOwnershipTransferRequested)
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

func (it *RMNRemoteOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *RMNRemoteOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNRemoteOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RMNRemote *RMNRemoteFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNRemoteOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMNRemote.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RMNRemoteOwnershipTransferRequestedIterator{contract: _RMNRemote.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_RMNRemote *RMNRemoteFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RMNRemoteOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMNRemote.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNRemoteOwnershipTransferRequested)
				if err := _RMNRemote.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_RMNRemote *RMNRemoteFilterer) ParseOwnershipTransferRequested(log types.Log) (*RMNRemoteOwnershipTransferRequested, error) {
	event := new(RMNRemoteOwnershipTransferRequested)
	if err := _RMNRemote.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNRemoteOwnershipTransferredIterator struct {
	Event *RMNRemoteOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNRemoteOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNRemoteOwnershipTransferred)
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
		it.Event = new(RMNRemoteOwnershipTransferred)
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

func (it *RMNRemoteOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *RMNRemoteOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNRemoteOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RMNRemote *RMNRemoteFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNRemoteOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMNRemote.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RMNRemoteOwnershipTransferredIterator{contract: _RMNRemote.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_RMNRemote *RMNRemoteFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RMNRemoteOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMNRemote.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNRemoteOwnershipTransferred)
				if err := _RMNRemote.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_RMNRemote *RMNRemoteFilterer) ParseOwnershipTransferred(log types.Log) (*RMNRemoteOwnershipTransferred, error) {
	event := new(RMNRemoteOwnershipTransferred)
	if err := _RMNRemote.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_RMNRemote *RMNRemote) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _RMNRemote.abi.Events["ConfigSet"].ID:
		return _RMNRemote.ParseConfigSet(log)
	case _RMNRemote.abi.Events["OwnershipTransferRequested"].ID:
		return _RMNRemote.ParseOwnershipTransferRequested(log)
	case _RMNRemote.abi.Events["OwnershipTransferred"].ID:
		return _RMNRemote.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (RMNRemoteConfigSet) Topic() common.Hash {
	return common.HexToHash("0x6cc65868ae41a007e6c3ed18ce591c123dd4e5864b421888c68ce92dae98cea4")
}

func (RMNRemoteOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (RMNRemoteOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_RMNRemote *RMNRemote) Address() common.Address {
	return _RMNRemote.address
}

type RMNRemoteInterface interface {
	GetVersionedConfig(opts *bind.CallOpts) (RMNRemoteVersionedConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	Verify(opts *bind.CallOpts, destLaneUpdates []MerkleRoot, signatures []Signature) error

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, newConfig RMNRemoteConfig) (*types.Transaction, error)

	Test(opts *bind.TransactOpts) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*RMNRemoteConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *RMNRemoteConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*RMNRemoteConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNRemoteOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RMNRemoteOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*RMNRemoteOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNRemoteOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RMNRemoteOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*RMNRemoteOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
