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

type IRMNV2Signature struct {
	R [32]byte
	S [32]byte
}

type InternalMerkleRoot struct {
	SourceChainSelector uint64
	OnRampAddress       []byte
	MinSeqNr            uint64
	MaxSeqNr            uint64
	MerkleRoot          [32]byte
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

var RMNRemoteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"subject\",\"type\":\"bytes16\"}],\"name\":\"AlreadyCursed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConfigNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateOnchainPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignerOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinSignersTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"subject\",\"type\":\"bytes16\"}],\"name\":\"NotCursed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfOrderSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedSigner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rmnHomeContractConfigDigest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"onchainPublicKey\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nodeIndex\",\"type\":\"uint64\"}],\"internalType\":\"structRMNRemote.Signer[]\",\"name\":\"signers\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"minSigners\",\"type\":\"uint64\"}],\"internalType\":\"structRMNRemote.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structRMNRemote.VersionedConfig\",\"name\":\"versionedConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes16[]\",\"name\":\"subjects\",\"type\":\"bytes16[]\"}],\"name\":\"Cursed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes16[]\",\"name\":\"subjects\",\"type\":\"bytes16[]\"}],\"name\":\"Uncursed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"subject\",\"type\":\"bytes16\"}],\"name\":\"curse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16[]\",\"name\":\"subjects\",\"type\":\"bytes16[]\"}],\"name\":\"curse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainSelector\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCursedSubjects\",\"outputs\":[{\"internalType\":\"bytes16[]\",\"name\":\"subjects\",\"type\":\"bytes16[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersionedConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rmnHomeContractConfigDigest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"onchainPublicKey\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nodeIndex\",\"type\":\"uint64\"}],\"internalType\":\"structRMNRemote.Signer[]\",\"name\":\"signers\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"minSigners\",\"type\":\"uint64\"}],\"internalType\":\"structRMNRemote.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"internalType\":\"structRMNRemote.VersionedConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"subject\",\"type\":\"bytes16\"}],\"name\":\"isCursed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCursed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rmnHomeContractConfigDigest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"onchainPublicKey\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nodeIndex\",\"type\":\"uint64\"}],\"internalType\":\"structRMNRemote.Signer[]\",\"name\":\"signers\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"minSigners\",\"type\":\"uint64\"}],\"internalType\":\"structRMNRemote.Config\",\"name\":\"newConfig\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"subject\",\"type\":\"bytes16\"}],\"name\":\"uncurse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16[]\",\"name\":\"subjects\",\"type\":\"bytes16[]\"}],\"name\":\"uncurse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offrampAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"onRampAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.MerkleRoot[]\",\"name\":\"destLaneUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIRMNV2.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620022df380380620022df83398101604081905262000034916200017e565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000d3565b5050506001600160401b0316608052620001b0565b336001600160a01b038216036200012d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000602082840312156200019157600080fd5b81516001600160401b0381168114620001a957600080fd5b9392505050565b60805161210c620001d3600039600081816101f00152610a5d015261210c6000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806370a9089e116100975780639a19b329116100665780639a19b32914610242578063d881e09214610255578063f2fde38b1461026a578063f8bb876e1461027d57600080fd5b806370a9089e146101c757806379ba5097146101da5780637e72ffa6146101e25780638da5cb5b1461021a57600080fd5b80632cbc26bb116100d35780632cbc26bb14610176578063397796f71461019957806362eed415146101a15780636d2d3993146101b457600080fd5b8063181f5a77146100fa578063198f0f771461014c5780631add205f14610161575b600080fd5b6101366040518060400160405280601381526020017f524d4e52656d6f746520312e362e302d6465760000000000000000000000000081525081565b6040516101439190611535565b60405180910390f35b61015f61015a36600461154f565b610290565b005b6101696106c1565b604051610143919061158a565b610189610184366004611677565b6107a7565b6040519015158152602001610143565b610189610849565b61015f6101af366004611677565b610901565b61015f6101c2366004611677565b610975565b61015f6101d5366004611875565b6109e5565b61015f610d37565b60405167ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610143565b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610143565b61015f610250366004611a46565b610e39565b61025d6110cf565b6040516101439190611ade565b61015f610278366004611b44565b61115d565b61015f61028b366004611a46565b611171565b61029861130b565b60015b6102a86020830183611b61565b9050811015610378576102be6020830183611b61565b828181106102ce576102ce611bd0565b90506040020160200160208101906102e69190611bff565b67ffffffffffffffff166102fd6020840184611b61565b610308600185611c4b565b81811061031757610317611bd0565b905060400201602001602081019061032f9190611bff565b67ffffffffffffffff1610610370576040517f4485151700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60010161029b565b506103866020820182611b61565b90506103986060830160408401611bff565b67ffffffffffffffff1611156103da576040517ffba0d9e600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025b6001810154156104d65760018082018054600892600092916103ff9190611c4b565b8154811061040f5761040f611bd0565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556001810180548061047957610479611c5e565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffff000000000000000000000000000000000000000000000000000000001690550190556103dd565b5060005b6104e76020830183611b61565b905081101561061c57600860006105016020850185611b61565b8481811061051157610511611bd0565b6105279260206040909202019081019150611b44565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040016000205460ff1615610588576040517f28cae27d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016008600061059b6020860186611b61565b858181106105ab576105ab611bd0565b6105c19260206040909202019081019150611b44565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790556001016104da565b5080600261062a8282611d46565b5050600580546000919082906106459063ffffffff16611e81565b91906101000a81548163ffffffff021916908363ffffffff160217905590507f6cc65868ae41a007e6c3ed18ce591c123dd4e5864b421888c68ce92dae98cea460405180604001604052808363ffffffff168152602001846106a690611ea4565b90526040516106b5919061158a565b60405180910390a15050565b6106c9611483565b60408051808201825260055463ffffffff1681528151606081018352600280548252600380548551602082810282018101909752818152949580870195858201939092909160009084015b82821015610782576000848152602090819020604080518082019091529084015473ffffffffffffffffffffffffffffffffffffffff8116825274010000000000000000000000000000000000000000900467ffffffffffffffff1681830152825260019092019101610714565b505050908252506002919091015467ffffffffffffffff166020909101529052919050565b60065460009081036107bb57506000919050565b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000821660009081526007602052604090205415158061084357507f010000000000000000000000000000010000000000000000000000000000000060005260076020527f70b766b11586b6b505ed3893938b0cc6c6c98bd6f65e969ac311168d34e4f9e25415155b92915050565b600654600090810361085b5750600090565b7f010000000000000000000000000000000000000000000000000000000000000060005260076020527f7dde556524061d0ce70b736a6e842a48e4927608bf87fd31432ced12a03ffeb8541515806108fc57507f010000000000000000000000000000010000000000000000000000000000000060005260076020527f70b766b11586b6b505ed3893938b0cc6c6c98bd6f65e969ac311168d34e4f9e25415155b905090565b60408051600180825281830190925260009160208083019080368337019050509050818160008151811061093757610937611bd0565b7fffffffffffffffffffffffffffffffff000000000000000000000000000000009092166020928302919091019091015261097181611171565b5050565b6040805160018082528183019092526000916020808301908036833701905050905081816000815181106109ab576109ab611bd0565b7fffffffffffffffffffffffffffffffff000000000000000000000000000000009092166020928302919091019091015261097181610e39565b60055463ffffffff16600003610a27576040517face124bc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f9651943783dbf81935a60e98f218a9d9b5b28823fb2228bbd91320d632facf536040518060c001604052804681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020013073ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001600260000154815260200185815250604051602001610ae3929190611f92565b60405160208183030381529060405280519060200120905060008060005b8451811015610ce8576000858281518110610b1e57610b1e611bd0565b602002602001015190506000600186601b8460000151856020015160405160008152602001604052604051610b6f949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610b91573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116610c09576040517f8baa579f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610c6e576040517fbbe15e7f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811660009081526008602052604090205460ff16610ccd576040517faaaa914100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b925082610cd9856120c7565b94505050806001019050610b01565b5060045467ffffffffffffffff16821015610d2f576040517f59fa4a9300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610dbd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610e4161130b565b60005b8151811015611094576000828281518110610e6157610e61611bd0565b6020908102919091018101517fffffffffffffffffffffffffffffffff00000000000000000000000000000000811660009081526007909252604082205490925090819003610f00576040517f73281fa10000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff0000000000000000000000000000000083166004820152602401610db4565b6000610f0d600183611c4b565b60068054919250600091610f2390600190611c4b565b81548110610f3357610f33611bd0565b90600052602060002090600291828204019190066010029054906101000a900460801b90508060068381548110610f6c57610f6c611bd0565b90600052602060002090600291828204019190066010026101000a8154816fffffffffffffffffffffffffffffffff021916908360801c02179055508260076000836fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001908152602001600020819055506006805480610ff757610ff7611c5e565b6000828152602080822060027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9490940193840401805460018086166010026101000a6fffffffffffffffffffffffffffffffff0219909116909155929093557fffffffffffffffffffffffffffffffff000000000000000000000000000000009690961686526007909152604085209490945550505001610e44565b507f0676e709c9cc74fa0519fd78f7c33be0f1b2b0bae0507c724aef7229379c6ba1816040516110c49190611ade565b60405180910390a150565b6060600680548060200260200160405190810160405280929190818152602001828054801561115357602002820191906000526020600020906000905b82829054906101000a900460801b6fffffffffffffffffffffffffffffffff191681526020019060100190602082600f0104928301926001038202915080841161110c5790505b5050505050905090565b61116561130b565b61116e8161138e565b50565b61117961130b565b60005b81518110156112db57600082828151811061119957611199611bd0565b6020026020010151905060076000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020019081526020016000205460001461123a576040517f19d5c79b0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff0000000000000000000000000000000082166004820152602401610db4565b60068054600180820183557ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f600283040180546fffffffffffffffffffffffffffffffff9383166010026101000a9384021916608086901c939093029290921790915590547fffffffffffffffffffffffffffffffff000000000000000000000000000000009092166000908152600760205260409020919091550161117c565b507f1716e663a90a76d3b6c7e5f680673d1b051454c19c627e184c8daf28f3104f74816040516110c49190611ade565b60005473ffffffffffffffffffffffffffffffffffffffff16331461138c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610db4565b565b3373ffffffffffffffffffffffffffffffffffffffff82160361140d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610db4565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040518060400160405280600063ffffffff1681526020016114cc60405180606001604052806000801916815260200160608152602001600067ffffffffffffffff1681525090565b905290565b6000815180845260005b818110156114f7576020818501810151868301820152016114db565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061154860208301846114d1565b9392505050565b60006020828403121561156157600080fd5b813567ffffffffffffffff81111561157857600080fd5b82016060818503121561154857600080fd5b6000602080835263ffffffff8451168184015280840151604080604086015260c0850182516060870152838301516060608088015281815180845260e0890191508683019350600092505b8083101561161e578351805173ffffffffffffffffffffffffffffffffffffffff16835287015167ffffffffffffffff16878301529286019260019290920191908401906115d5565b50604085015167ffffffffffffffff811660a08a0152955098975050505050505050565b80357fffffffffffffffffffffffffffffffff000000000000000000000000000000008116811461167257600080fd5b919050565b60006020828403121561168957600080fd5b61154882611642565b73ffffffffffffffffffffffffffffffffffffffff8116811461116e57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715611706576117066116b4565b60405290565b60405160a0810167ffffffffffffffff81118282101715611706576117066116b4565b6040516060810167ffffffffffffffff81118282101715611706576117066116b4565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611799576117996116b4565b604052919050565b600067ffffffffffffffff8211156117bb576117bb6116b4565b5060051b60200190565b67ffffffffffffffff8116811461116e57600080fd5b8035611672816117c5565b600082601f8301126117f757600080fd5b8135602061180c611807836117a1565b611752565b82815260069290921b8401810191818101908684111561182b57600080fd5b8286015b8481101561186a57604081890312156118485760008081fd5b6118506116e3565b81358152848201358582015283529183019160400161182f565b509695505050505050565b60008060006060848603121561188a57600080fd5b6118948435611692565b83359250602084013567ffffffffffffffff808211156118b357600080fd5b818601915086601f8301126118c757600080fd5b6118d461180783356117a1565b82358082526020808301929160051b850101898111156118f357600080fd5b602085015b81811015611a1557848135111561190e57600080fd5b803586017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060a081838f0301121561194557600080fd5b61194d61170c565b602083013561195b816117c5565b815260408301358881111561196f57600080fd5b8301603f81018f1361198057600080fd5b602081013589811115611995576119956116b4565b6119a6602085601f84011601611752565b93508084528f60408284010111156119bd57600080fd5b8060408301602086013760006020828601015250508160208201526119e4606084016117db565b60408201526119f5608084016117db565b606082015260a092909201356080830152508452602093840193016118f8565b5090955050506040860135915080821115611a2f57600080fd5b50611a3c868287016117e6565b9150509250925092565b60006020808385031215611a5957600080fd5b823567ffffffffffffffff811115611a7057600080fd5b8301601f81018513611a8157600080fd5b8035611a8f611807826117a1565b81815260059190911b82018301908381019087831115611aae57600080fd5b928401925b82841015611ad357611ac484611642565b82529284019290840190611ab3565b979650505050505050565b6020808252825182820181905260009190848201906040850190845b81811015611b385783517fffffffffffffffffffffffffffffffff000000000000000000000000000000001683529284019291840191600101611afa565b50909695505050505050565b600060208284031215611b5657600080fd5b813561154881611692565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611b9657600080fd5b83018035915067ffffffffffffffff821115611bb157600080fd5b6020019150600681901b3603821315611bc957600080fd5b9250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215611c1157600080fd5b8135611548816117c5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561084357610843611c1c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008135610843816117c5565b8135611ca581611692565b73ffffffffffffffffffffffffffffffffffffffff811690508154817fffffffffffffffffffffffff000000000000000000000000000000000000000082161783556020840135611cf5816117c5565b7bffffffffffffffff00000000000000000000000000000000000000008160a01b16837fffffffff000000000000000000000000000000000000000000000000000000008416171784555050505050565b81358155600180820160208401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1853603018112611d8457600080fd5b8401803567ffffffffffffffff811115611d9d57600080fd5b6020820191508060061b3603821315611db557600080fd5b68010000000000000000811115611dce57611dce6116b4565b825481845580821015611e03576000848152602081208381019083015b80821015611dff5782825590870190611deb565b5050505b50600092835260208320925b81811015611e3357611e218385611c9a565b92840192604092909201918401611e0f565b5050505050610971611e4760408401611c8d565b6002830167ffffffffffffffff82167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008254161781555050565b600063ffffffff808316818103611e9a57611e9a611c1c565b6001019392505050565b600060608236031215611eb657600080fd5b611ebe61172f565b8235815260208084013567ffffffffffffffff811115611edd57600080fd5b840136601f820112611eee57600080fd5b8035611efc611807826117a1565b81815260069190911b82018301908381019036831115611f1b57600080fd5b928401925b82841015611f715760408436031215611f395760008081fd5b611f416116e3565b8435611f4c81611692565b815284860135611f5b816117c5565b8187015282526040939093019290840190611f20565b80858701525050505050611f87604084016117db565b604082015292915050565b60006040848352602060408185015261010084018551604086015281860151606067ffffffffffffffff808316606089015260408901519250608073ffffffffffffffffffffffffffffffffffffffff80851660808b015260608b0151945060a081861660a08c015260808c015160c08c015260a08c0151955060c060e08c015286915085518088526101209750878c019250878160051b8d01019750888701965060005b818110156120b4577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee08d8a030184528751868151168a528a810151848c8c0152612083858c01826114d1565b828e015189168c8f01528983015189168a8d0152918701519a87019a909a5298509689019692890192600101612037565b50969d9c50505050505050505050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036120f8576120f8611c1c565b506001019056fea164736f6c6343000818000a",
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

func (_RMNRemote *RMNRemoteCaller) GetChainSelector(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "getChainSelector")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) GetChainSelector() (uint64, error) {
	return _RMNRemote.Contract.GetChainSelector(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) GetChainSelector() (uint64, error) {
	return _RMNRemote.Contract.GetChainSelector(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCaller) GetCursedSubjects(opts *bind.CallOpts) ([][16]byte, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "getCursedSubjects")

	if err != nil {
		return *new([][16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][16]byte)).(*[][16]byte)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) GetCursedSubjects() ([][16]byte, error) {
	return _RMNRemote.Contract.GetCursedSubjects(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) GetCursedSubjects() ([][16]byte, error) {
	return _RMNRemote.Contract.GetCursedSubjects(&_RMNRemote.CallOpts)
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

func (_RMNRemote *RMNRemoteCaller) IsCursed(opts *bind.CallOpts, subject [16]byte) (bool, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "isCursed", subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) IsCursed(subject [16]byte) (bool, error) {
	return _RMNRemote.Contract.IsCursed(&_RMNRemote.CallOpts, subject)
}

func (_RMNRemote *RMNRemoteCallerSession) IsCursed(subject [16]byte) (bool, error) {
	return _RMNRemote.Contract.IsCursed(&_RMNRemote.CallOpts, subject)
}

func (_RMNRemote *RMNRemoteCaller) IsCursed0(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "isCursed0")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_RMNRemote *RMNRemoteSession) IsCursed0() (bool, error) {
	return _RMNRemote.Contract.IsCursed0(&_RMNRemote.CallOpts)
}

func (_RMNRemote *RMNRemoteCallerSession) IsCursed0() (bool, error) {
	return _RMNRemote.Contract.IsCursed0(&_RMNRemote.CallOpts)
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

func (_RMNRemote *RMNRemoteCaller) Verify(opts *bind.CallOpts, offrampAddress common.Address, destLaneUpdates []InternalMerkleRoot, signatures []IRMNV2Signature) error {
	var out []interface{}
	err := _RMNRemote.contract.Call(opts, &out, "verify", offrampAddress, destLaneUpdates, signatures)

	if err != nil {
		return err
	}

	return err

}

func (_RMNRemote *RMNRemoteSession) Verify(offrampAddress common.Address, destLaneUpdates []InternalMerkleRoot, signatures []IRMNV2Signature) error {
	return _RMNRemote.Contract.Verify(&_RMNRemote.CallOpts, offrampAddress, destLaneUpdates, signatures)
}

func (_RMNRemote *RMNRemoteCallerSession) Verify(offrampAddress common.Address, destLaneUpdates []InternalMerkleRoot, signatures []IRMNV2Signature) error {
	return _RMNRemote.Contract.Verify(&_RMNRemote.CallOpts, offrampAddress, destLaneUpdates, signatures)
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

func (_RMNRemote *RMNRemoteTransactor) Curse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "curse", subject)
}

func (_RMNRemote *RMNRemoteSession) Curse(subject [16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Curse(&_RMNRemote.TransactOpts, subject)
}

func (_RMNRemote *RMNRemoteTransactorSession) Curse(subject [16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Curse(&_RMNRemote.TransactOpts, subject)
}

func (_RMNRemote *RMNRemoteTransactor) Curse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "curse0", subjects)
}

func (_RMNRemote *RMNRemoteSession) Curse0(subjects [][16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Curse0(&_RMNRemote.TransactOpts, subjects)
}

func (_RMNRemote *RMNRemoteTransactorSession) Curse0(subjects [][16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Curse0(&_RMNRemote.TransactOpts, subjects)
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

func (_RMNRemote *RMNRemoteTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "transferOwnership", to)
}

func (_RMNRemote *RMNRemoteSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RMNRemote.Contract.TransferOwnership(&_RMNRemote.TransactOpts, to)
}

func (_RMNRemote *RMNRemoteTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RMNRemote.Contract.TransferOwnership(&_RMNRemote.TransactOpts, to)
}

func (_RMNRemote *RMNRemoteTransactor) Uncurse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "uncurse", subject)
}

func (_RMNRemote *RMNRemoteSession) Uncurse(subject [16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Uncurse(&_RMNRemote.TransactOpts, subject)
}

func (_RMNRemote *RMNRemoteTransactorSession) Uncurse(subject [16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Uncurse(&_RMNRemote.TransactOpts, subject)
}

func (_RMNRemote *RMNRemoteTransactor) Uncurse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error) {
	return _RMNRemote.contract.Transact(opts, "uncurse0", subjects)
}

func (_RMNRemote *RMNRemoteSession) Uncurse0(subjects [][16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Uncurse0(&_RMNRemote.TransactOpts, subjects)
}

func (_RMNRemote *RMNRemoteTransactorSession) Uncurse0(subjects [][16]byte) (*types.Transaction, error) {
	return _RMNRemote.Contract.Uncurse0(&_RMNRemote.TransactOpts, subjects)
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

type RMNRemoteCursedIterator struct {
	Event *RMNRemoteCursed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNRemoteCursedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNRemoteCursed)
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
		it.Event = new(RMNRemoteCursed)
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

func (it *RMNRemoteCursedIterator) Error() error {
	return it.fail
}

func (it *RMNRemoteCursedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNRemoteCursed struct {
	Subjects [][16]byte
	Raw      types.Log
}

func (_RMNRemote *RMNRemoteFilterer) FilterCursed(opts *bind.FilterOpts) (*RMNRemoteCursedIterator, error) {

	logs, sub, err := _RMNRemote.contract.FilterLogs(opts, "Cursed")
	if err != nil {
		return nil, err
	}
	return &RMNRemoteCursedIterator{contract: _RMNRemote.contract, event: "Cursed", logs: logs, sub: sub}, nil
}

func (_RMNRemote *RMNRemoteFilterer) WatchCursed(opts *bind.WatchOpts, sink chan<- *RMNRemoteCursed) (event.Subscription, error) {

	logs, sub, err := _RMNRemote.contract.WatchLogs(opts, "Cursed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNRemoteCursed)
				if err := _RMNRemote.contract.UnpackLog(event, "Cursed", log); err != nil {
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

func (_RMNRemote *RMNRemoteFilterer) ParseCursed(log types.Log) (*RMNRemoteCursed, error) {
	event := new(RMNRemoteCursed)
	if err := _RMNRemote.contract.UnpackLog(event, "Cursed", log); err != nil {
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

type RMNRemoteUncursedIterator struct {
	Event *RMNRemoteUncursed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNRemoteUncursedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNRemoteUncursed)
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
		it.Event = new(RMNRemoteUncursed)
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

func (it *RMNRemoteUncursedIterator) Error() error {
	return it.fail
}

func (it *RMNRemoteUncursedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNRemoteUncursed struct {
	Subjects [][16]byte
	Raw      types.Log
}

func (_RMNRemote *RMNRemoteFilterer) FilterUncursed(opts *bind.FilterOpts) (*RMNRemoteUncursedIterator, error) {

	logs, sub, err := _RMNRemote.contract.FilterLogs(opts, "Uncursed")
	if err != nil {
		return nil, err
	}
	return &RMNRemoteUncursedIterator{contract: _RMNRemote.contract, event: "Uncursed", logs: logs, sub: sub}, nil
}

func (_RMNRemote *RMNRemoteFilterer) WatchUncursed(opts *bind.WatchOpts, sink chan<- *RMNRemoteUncursed) (event.Subscription, error) {

	logs, sub, err := _RMNRemote.contract.WatchLogs(opts, "Uncursed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNRemoteUncursed)
				if err := _RMNRemote.contract.UnpackLog(event, "Uncursed", log); err != nil {
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

func (_RMNRemote *RMNRemoteFilterer) ParseUncursed(log types.Log) (*RMNRemoteUncursed, error) {
	event := new(RMNRemoteUncursed)
	if err := _RMNRemote.contract.UnpackLog(event, "Uncursed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_RMNRemote *RMNRemote) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _RMNRemote.abi.Events["ConfigSet"].ID:
		return _RMNRemote.ParseConfigSet(log)
	case _RMNRemote.abi.Events["Cursed"].ID:
		return _RMNRemote.ParseCursed(log)
	case _RMNRemote.abi.Events["OwnershipTransferRequested"].ID:
		return _RMNRemote.ParseOwnershipTransferRequested(log)
	case _RMNRemote.abi.Events["OwnershipTransferred"].ID:
		return _RMNRemote.ParseOwnershipTransferred(log)
	case _RMNRemote.abi.Events["Uncursed"].ID:
		return _RMNRemote.ParseUncursed(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (RMNRemoteConfigSet) Topic() common.Hash {
	return common.HexToHash("0x6cc65868ae41a007e6c3ed18ce591c123dd4e5864b421888c68ce92dae98cea4")
}

func (RMNRemoteCursed) Topic() common.Hash {
	return common.HexToHash("0x1716e663a90a76d3b6c7e5f680673d1b051454c19c627e184c8daf28f3104f74")
}

func (RMNRemoteOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (RMNRemoteOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (RMNRemoteUncursed) Topic() common.Hash {
	return common.HexToHash("0x0676e709c9cc74fa0519fd78f7c33be0f1b2b0bae0507c724aef7229379c6ba1")
}

func (_RMNRemote *RMNRemote) Address() common.Address {
	return _RMNRemote.address
}

type RMNRemoteInterface interface {
	GetChainSelector(opts *bind.CallOpts) (uint64, error)

	GetCursedSubjects(opts *bind.CallOpts) ([][16]byte, error)

	GetVersionedConfig(opts *bind.CallOpts) (RMNRemoteVersionedConfig, error)

	IsCursed(opts *bind.CallOpts, subject [16]byte) (bool, error)

	IsCursed0(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	Verify(opts *bind.CallOpts, offrampAddress common.Address, destLaneUpdates []InternalMerkleRoot, signatures []IRMNV2Signature) error

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Curse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error)

	Curse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, newConfig RMNRemoteConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Uncurse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error)

	Uncurse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*RMNRemoteConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *RMNRemoteConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*RMNRemoteConfigSet, error)

	FilterCursed(opts *bind.FilterOpts) (*RMNRemoteCursedIterator, error)

	WatchCursed(opts *bind.WatchOpts, sink chan<- *RMNRemoteCursed) (event.Subscription, error)

	ParseCursed(log types.Log) (*RMNRemoteCursed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNRemoteOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RMNRemoteOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*RMNRemoteOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNRemoteOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RMNRemoteOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*RMNRemoteOwnershipTransferred, error)

	FilterUncursed(opts *bind.FilterOpts) (*RMNRemoteUncursedIterator, error)

	WatchUncursed(opts *bind.WatchOpts, sink chan<- *RMNRemoteUncursed) (event.Subscription, error)

	ParseUncursed(log types.Log) (*RMNRemoteUncursed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
