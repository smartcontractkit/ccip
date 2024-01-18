// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package no_op_ocr3

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

var NoOpOCR3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"ForkedChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"latestSequenceNumber\",\"type\":\"uint64\"}],\"name\":\"NonIncreasingSequenceNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR3Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b503380600081620000695760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156200009c576200009c81620000a9565b5050466080525062000154565b336001600160a01b03821603620001035760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000060565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60805161221b6200017060003960006110af015261221b6000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c806379ba50971161008c578063afcb95d711610066578063afcb95d7146101f9578063b11a19e81461015e578063b1dc65a414610232578063f2fde38b1461024557600080fd5b806379ba50971461019957806381ff7048146101a15780638da5cb5b146101d157600080fd5b80636900a3ae116100bd5780636900a3ae1461015e5780636a11ee901461017157806371aad10d1461018657600080fd5b8063181f5a77146100e457806356ca623e14610136578063666cab8d14610149575b600080fd5b6101206040518060400160405280600e81526020017f4e6f4f704f43523320312e302e3000000000000000000000000000000000000081525081565b60405161012d9190611a7d565b60405180910390f35b610120610144366004611ac0565b610258565b6101516102aa565b60405161012d9190611b2c565b61012061016c366004611b3f565b610319565b61018461017f366004611d14565b610331565b005b610120610194366004611de1565b610b49565b610184610e4f565b6004546002546040805163ffffffff8085168252640100000000909404909316602084015282015260600161012d565b60005460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161012d565b600254600454604080516001815260208101939093526801000000000000000090910467ffffffffffffffff169082015260600161012d565b610184610240366004611e6a565b610f4c565b610184610253366004611ac0565b61171a565b604051606082811b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166020830152906102a4906034015b604051602081830303815290604052610b49565b92915050565b6060600780548060200260200160405190810160405280929190818152602001828054801561030f57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116102e4575b5050505050905090565b60606102a48260405160200161029091815260200190565b855185518560ff16601f8311156103a9576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064015b60405180910390fd5b80600003610413576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016103a0565b8183146104a1576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e0000000000000000000000000000000000000000000000000000000060648201526084016103a0565b6104ac816003611f7e565b8311610514576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016103a0565b61051c61172e565b60065460005b8181101561061857600560006006838154811061054157610541611f95565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000169055600780546005929190849081106105b1576105b1611f95565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905561061181611fc4565b9050610522565b50895160005b818110156109f15760008c828151811061063a5761063a611f95565b602002602001015190506000600281111561065757610657611ffc565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260056020526040902054610100900460ff16600281111561069657610696611ffc565b146106fd576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016103a0565b73ffffffffffffffffffffffffffffffffffffffff811661074a576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260ff83168152602081016001905273ffffffffffffffffffffffffffffffffffffffff821660009081526005602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156107fa576107fa611ffc565b021790555090505060008c838151811061081657610816611f95565b602002602001015190506000600281111561083357610833611ffc565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260056020526040902054610100900460ff16600281111561087257610872611ffc565b146108d9576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016103a0565b73ffffffffffffffffffffffffffffffffffffffff8116610926576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260ff84168152602081016002905273ffffffffffffffffffffffffffffffffffffffff821660009081526005602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156109d6576109d6611ffc565b02179055509050505050806109ea90611fc4565b905061061e565b508a51610a059060069060208e0190611951565b508951610a199060079060208d0190611951565b506003805460ff838116610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000909216908c161717905560048054610a9f914691309190600090610a719063ffffffff1661202b565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168e8e8e8e8e8e6117b1565b600260000181905550600060048054906101000a900463ffffffff169050436004806101000a81548163ffffffff021916908363ffffffff1602179055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0581600260000154600460009054906101000a900463ffffffff168f8f8f8f8f8f604051610b339998979695949392919061204e565b60405180910390a1505050505050505050505050565b60408051808201909152601081527f30313233343536373839616263646566000000000000000000000000000000006020820152815160609190600090610b91906002611f7e565b610b9c9060026120e4565b67ffffffffffffffff811115610bb457610bb4611b58565b6040519080825280601f01601f191660200182016040528015610bde576020820181803683370190505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110610c1557610c15611f95565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610c7857610c78611f95565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060005b8451811015610e4757826004868381518110610cc857610cc8611f95565b016020015182517fff0000000000000000000000000000000000000000000000000000000000000090911690911c60f81c908110610d0857610d08611f95565b01602001517fff000000000000000000000000000000000000000000000000000000000000001682610d3b836002611f7e565b610d469060026120e4565b81518110610d5657610d56611f95565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535082858281518110610d9857610d98611f95565b602091010151815160f89190911c600f16908110610db857610db8611f95565b01602001517fff000000000000000000000000000000000000000000000000000000000000001682610deb836002611f7e565b610df69060036120e4565b81518110610e0657610e06611f95565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080610e3f81611fc4565b915050610caa565b509392505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610ed0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016103a0565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60045460208901359067ffffffffffffffff68010000000000000000909104811690821611610fd7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600660248201527f626173653a31000000000000000000000000000000000000000000000000000060448201526064016103a0565b600480547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8416021790556040805160608101825260025480825260035460ff808216602085015261010090910416928201929092528a359182146110ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600660248201527f626173653a32000000000000000000000000000000000000000000000000000060448201526064016103a0565b467f000000000000000000000000000000000000000000000000000000000000000014611135576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600660248201527f626173653a33000000000000000000000000000000000000000000000000000060448201526064016103a0565b6040805183815267ffffffffffffffff851660208201527fe893c2681d327421d89e1cb54fbe64645b4dcea668d6826130b62cf4c6eefea2910160405180910390a160208101516111879060016120f7565b60ff1687146111f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600660248201527f626173653a34000000000000000000000000000000000000000000000000000060448201526064016103a0565b86851461125b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600660248201527f626173653a35000000000000000000000000000000000000000000000000000060448201526064016103a0565b3360009081526005602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561129e5761129e611ffc565b60028111156112af576112af611ffc565b90525090506002816020015160028111156112cc576112cc611ffc565b14801561131357506007816000015160ff16815481106112ee576112ee611f95565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b611379576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600660248201527f626173653a36000000000000000000000000000000000000000000000000000060448201526064016103a0565b506000611387866020611f7e565b611392896020611f7e565b61139e8c6101446120e4565b6113a891906120e4565b6113b291906120e4565b905036811461141d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600660248201527f626173653a37000000000000000000000000000000000000000000000000000060448201526064016103a0565b5060008a8a604051611430929190612110565b604051908190038120611447918e90602001612120565b6040516020818303038152906040528051906020012090506114676119db565b8860005b818110156117095760006001858a846020811061148a5761148a611f95565b61149791901a601b6120f7565b8f8f868181106114a9576114a9611f95565b905060200201358e8e878181106114c2576114c2611f95565b90506020020135604051600081526020016040526040516114ff949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611521573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff8116600090815260056020908152848220848601909552845460ff80821686529397509195509293928401916101009091041660028111156115a4576115a4611ffc565b60028111156115b5576115b5611ffc565b90525090506001816020015160028111156115d2576115d2611ffc565b1461164e576115e082610258565b6040516020016115f09190612134565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526103a091600401611a7d565b8051859060ff16601f811061166557611665611f95565b6020020151156116d1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600660248201527f626173653a39000000000000000000000000000000000000000000000000000060448201526064016103a0565b600185826000015160ff16601f81106116ec576116ec611f95565b9115156020909202015250611702905081611fc4565b905061146b565b505050505050505050505050505050565b61172261172e565b61172b8161185c565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146117af576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016103a0565b565b6000808a8a8a8a8a8a8a8a8a6040516020016117d599989796959493929190612179565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036118db576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016103a0565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156119cb579160200282015b828111156119cb57825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190611971565b506119d79291506119fa565b5090565b604051806103e00160405280601f906020820280368337509192915050565b5b808211156119d757600081556001016119fb565b60005b83811015611a2a578181015183820152602001611a12565b50506000910152565b60008151808452611a4b816020860160208601611a0f565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611a906020830184611a33565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611abb57600080fd5b919050565b600060208284031215611ad257600080fd5b611a9082611a97565b600081518084526020808501945080840160005b83811015611b2157815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611aef565b509495945050505050565b602081526000611a906020830184611adb565b600060208284031215611b5157600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611bce57611bce611b58565b604052919050565b600082601f830112611be757600080fd5b8135602067ffffffffffffffff821115611c0357611c03611b58565b8160051b611c12828201611b87565b9283528481018201928281019087851115611c2c57600080fd5b83870192505b84831015611c5257611c4383611a97565b82529183019190830190611c32565b979650505050505050565b803560ff81168114611abb57600080fd5b600082601f830112611c7f57600080fd5b813567ffffffffffffffff811115611c9957611c99611b58565b611cca60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611b87565b818152846020838601011115611cdf57600080fd5b816020850160208301376000918101602001919091529392505050565b803567ffffffffffffffff81168114611abb57600080fd5b60008060008060008060c08789031215611d2d57600080fd5b863567ffffffffffffffff80821115611d4557600080fd5b611d518a838b01611bd6565b97506020890135915080821115611d6757600080fd5b611d738a838b01611bd6565b9650611d8160408a01611c5d565b95506060890135915080821115611d9757600080fd5b611da38a838b01611c6e565b9450611db160808a01611cfc565b935060a0890135915080821115611dc757600080fd5b50611dd489828a01611c6e565b9150509295509295509295565b600060208284031215611df357600080fd5b813567ffffffffffffffff811115611e0a57600080fd5b611e1684828501611c6e565b949350505050565b60008083601f840112611e3057600080fd5b50813567ffffffffffffffff811115611e4857600080fd5b6020830191508360208260051b8501011115611e6357600080fd5b9250929050565b60008060008060008060008060e0898b031215611e8657600080fd5b606089018a811115611e9757600080fd5b8998503567ffffffffffffffff80821115611eb157600080fd5b818b0191508b601f830112611ec557600080fd5b813581811115611ed457600080fd5b8c6020828501011115611ee657600080fd5b6020830199508098505060808b0135915080821115611f0457600080fd5b611f108c838d01611e1e565b909750955060a08b0135915080821115611f2957600080fd5b50611f368b828c01611e1e565b999c989b50969995989497949560c00135949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820281158282048414176102a4576102a4611f4f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611ff557611ff5611f4f565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600063ffffffff80831681810361204457612044611f4f565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b1660408501525080606084015261207e8184018a611adb565b905082810360808401526120928189611adb565b905060ff871660a084015282810360c08401526120af8187611a33565b905067ffffffffffffffff851660e08401528281036101008401526120d48185611a33565b9c9b505050505050505050505050565b808201808211156102a4576102a4611f4f565b60ff81811683821601908111156102a4576102a4611f4f565b8183823760009101908152919050565b828152606082602083013760800192915050565b7f626173653a383a2000000000000000000000000000000000000000000000000081526000825161216c816008850160208701611a0f565b9190910160080192915050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526121c08285018b611adb565b915083820360808501526121d4828a611adb565b915060ff881660a085015283820360c08501526121f18288611a33565b90861660e085015283810361010085015290506120d48185611a3356fea164736f6c6343000813000a",
}

var NoOpOCR3ABI = NoOpOCR3MetaData.ABI

var NoOpOCR3Bin = NoOpOCR3MetaData.Bin

func DeployNoOpOCR3(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NoOpOCR3, error) {
	parsed, err := NoOpOCR3MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NoOpOCR3Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NoOpOCR3{address: address, abi: *parsed, NoOpOCR3Caller: NoOpOCR3Caller{contract: contract}, NoOpOCR3Transactor: NoOpOCR3Transactor{contract: contract}, NoOpOCR3Filterer: NoOpOCR3Filterer{contract: contract}}, nil
}

type NoOpOCR3 struct {
	address common.Address
	abi     abi.ABI
	NoOpOCR3Caller
	NoOpOCR3Transactor
	NoOpOCR3Filterer
}

type NoOpOCR3Caller struct {
	contract *bind.BoundContract
}

type NoOpOCR3Transactor struct {
	contract *bind.BoundContract
}

type NoOpOCR3Filterer struct {
	contract *bind.BoundContract
}

type NoOpOCR3Session struct {
	Contract     *NoOpOCR3
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type NoOpOCR3CallerSession struct {
	Contract *NoOpOCR3Caller
	CallOpts bind.CallOpts
}

type NoOpOCR3TransactorSession struct {
	Contract     *NoOpOCR3Transactor
	TransactOpts bind.TransactOpts
}

type NoOpOCR3Raw struct {
	Contract *NoOpOCR3
}

type NoOpOCR3CallerRaw struct {
	Contract *NoOpOCR3Caller
}

type NoOpOCR3TransactorRaw struct {
	Contract *NoOpOCR3Transactor
}

func NewNoOpOCR3(address common.Address, backend bind.ContractBackend) (*NoOpOCR3, error) {
	abi, err := abi.JSON(strings.NewReader(NoOpOCR3ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindNoOpOCR3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NoOpOCR3{address: address, abi: abi, NoOpOCR3Caller: NoOpOCR3Caller{contract: contract}, NoOpOCR3Transactor: NoOpOCR3Transactor{contract: contract}, NoOpOCR3Filterer: NoOpOCR3Filterer{contract: contract}}, nil
}

func NewNoOpOCR3Caller(address common.Address, caller bind.ContractCaller) (*NoOpOCR3Caller, error) {
	contract, err := bindNoOpOCR3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NoOpOCR3Caller{contract: contract}, nil
}

func NewNoOpOCR3Transactor(address common.Address, transactor bind.ContractTransactor) (*NoOpOCR3Transactor, error) {
	contract, err := bindNoOpOCR3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NoOpOCR3Transactor{contract: contract}, nil
}

func NewNoOpOCR3Filterer(address common.Address, filterer bind.ContractFilterer) (*NoOpOCR3Filterer, error) {
	contract, err := bindNoOpOCR3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NoOpOCR3Filterer{contract: contract}, nil
}

func bindNoOpOCR3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NoOpOCR3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_NoOpOCR3 *NoOpOCR3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoOpOCR3.Contract.NoOpOCR3Caller.contract.Call(opts, result, method, params...)
}

func (_NoOpOCR3 *NoOpOCR3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOpOCR3.Contract.NoOpOCR3Transactor.contract.Transfer(opts)
}

func (_NoOpOCR3 *NoOpOCR3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOpOCR3.Contract.NoOpOCR3Transactor.contract.Transact(opts, method, params...)
}

func (_NoOpOCR3 *NoOpOCR3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoOpOCR3.Contract.contract.Call(opts, result, method, params...)
}

func (_NoOpOCR3 *NoOpOCR3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOpOCR3.Contract.contract.Transfer(opts)
}

func (_NoOpOCR3 *NoOpOCR3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOpOCR3.Contract.contract.Transact(opts, method, params...)
}

func (_NoOpOCR3 *NoOpOCR3Caller) GetTransmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _NoOpOCR3.contract.Call(opts, &out, "getTransmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_NoOpOCR3 *NoOpOCR3Session) GetTransmitters() ([]common.Address, error) {
	return _NoOpOCR3.Contract.GetTransmitters(&_NoOpOCR3.CallOpts)
}

func (_NoOpOCR3 *NoOpOCR3CallerSession) GetTransmitters() ([]common.Address, error) {
	return _NoOpOCR3.Contract.GetTransmitters(&_NoOpOCR3.CallOpts)
}

func (_NoOpOCR3 *NoOpOCR3Caller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _NoOpOCR3.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_NoOpOCR3 *NoOpOCR3Session) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _NoOpOCR3.Contract.LatestConfigDetails(&_NoOpOCR3.CallOpts)
}

func (_NoOpOCR3 *NoOpOCR3CallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _NoOpOCR3.Contract.LatestConfigDetails(&_NoOpOCR3.CallOpts)
}

func (_NoOpOCR3 *NoOpOCR3Caller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _NoOpOCR3.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.SequenceNumber = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

func (_NoOpOCR3 *NoOpOCR3Session) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _NoOpOCR3.Contract.LatestConfigDigestAndEpoch(&_NoOpOCR3.CallOpts)
}

func (_NoOpOCR3 *NoOpOCR3CallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _NoOpOCR3.Contract.LatestConfigDigestAndEpoch(&_NoOpOCR3.CallOpts)
}

func (_NoOpOCR3 *NoOpOCR3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NoOpOCR3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_NoOpOCR3 *NoOpOCR3Session) Owner() (common.Address, error) {
	return _NoOpOCR3.Contract.Owner(&_NoOpOCR3.CallOpts)
}

func (_NoOpOCR3 *NoOpOCR3CallerSession) Owner() (common.Address, error) {
	return _NoOpOCR3.Contract.Owner(&_NoOpOCR3.CallOpts)
}

func (_NoOpOCR3 *NoOpOCR3Caller) ToString(opts *bind.CallOpts, account common.Address) (string, error) {
	var out []interface{}
	err := _NoOpOCR3.contract.Call(opts, &out, "toString", account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_NoOpOCR3 *NoOpOCR3Session) ToString(account common.Address) (string, error) {
	return _NoOpOCR3.Contract.ToString(&_NoOpOCR3.CallOpts, account)
}

func (_NoOpOCR3 *NoOpOCR3CallerSession) ToString(account common.Address) (string, error) {
	return _NoOpOCR3.Contract.ToString(&_NoOpOCR3.CallOpts, account)
}

func (_NoOpOCR3 *NoOpOCR3Caller) ToString0(opts *bind.CallOpts, value *big.Int) (string, error) {
	var out []interface{}
	err := _NoOpOCR3.contract.Call(opts, &out, "toString0", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_NoOpOCR3 *NoOpOCR3Session) ToString0(value *big.Int) (string, error) {
	return _NoOpOCR3.Contract.ToString0(&_NoOpOCR3.CallOpts, value)
}

func (_NoOpOCR3 *NoOpOCR3CallerSession) ToString0(value *big.Int) (string, error) {
	return _NoOpOCR3.Contract.ToString0(&_NoOpOCR3.CallOpts, value)
}

func (_NoOpOCR3 *NoOpOCR3Caller) ToString1(opts *bind.CallOpts, data []byte) (string, error) {
	var out []interface{}
	err := _NoOpOCR3.contract.Call(opts, &out, "toString1", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_NoOpOCR3 *NoOpOCR3Session) ToString1(data []byte) (string, error) {
	return _NoOpOCR3.Contract.ToString1(&_NoOpOCR3.CallOpts, data)
}

func (_NoOpOCR3 *NoOpOCR3CallerSession) ToString1(data []byte) (string, error) {
	return _NoOpOCR3.Contract.ToString1(&_NoOpOCR3.CallOpts, data)
}

func (_NoOpOCR3 *NoOpOCR3Caller) ToString2(opts *bind.CallOpts, value [32]byte) (string, error) {
	var out []interface{}
	err := _NoOpOCR3.contract.Call(opts, &out, "toString2", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_NoOpOCR3 *NoOpOCR3Session) ToString2(value [32]byte) (string, error) {
	return _NoOpOCR3.Contract.ToString2(&_NoOpOCR3.CallOpts, value)
}

func (_NoOpOCR3 *NoOpOCR3CallerSession) ToString2(value [32]byte) (string, error) {
	return _NoOpOCR3.Contract.ToString2(&_NoOpOCR3.CallOpts, value)
}

func (_NoOpOCR3 *NoOpOCR3Caller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NoOpOCR3.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_NoOpOCR3 *NoOpOCR3Session) TypeAndVersion() (string, error) {
	return _NoOpOCR3.Contract.TypeAndVersion(&_NoOpOCR3.CallOpts)
}

func (_NoOpOCR3 *NoOpOCR3CallerSession) TypeAndVersion() (string, error) {
	return _NoOpOCR3.Contract.TypeAndVersion(&_NoOpOCR3.CallOpts)
}

func (_NoOpOCR3 *NoOpOCR3Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOpOCR3.contract.Transact(opts, "acceptOwnership")
}

func (_NoOpOCR3 *NoOpOCR3Session) AcceptOwnership() (*types.Transaction, error) {
	return _NoOpOCR3.Contract.AcceptOwnership(&_NoOpOCR3.TransactOpts)
}

func (_NoOpOCR3 *NoOpOCR3TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _NoOpOCR3.Contract.AcceptOwnership(&_NoOpOCR3.TransactOpts)
}

func (_NoOpOCR3 *NoOpOCR3Transactor) SetOCR3Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _NoOpOCR3.contract.Transact(opts, "setOCR3Config", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_NoOpOCR3 *NoOpOCR3Session) SetOCR3Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _NoOpOCR3.Contract.SetOCR3Config(&_NoOpOCR3.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_NoOpOCR3 *NoOpOCR3TransactorSession) SetOCR3Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _NoOpOCR3.Contract.SetOCR3Config(&_NoOpOCR3.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_NoOpOCR3 *NoOpOCR3Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _NoOpOCR3.contract.Transact(opts, "transferOwnership", to)
}

func (_NoOpOCR3 *NoOpOCR3Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _NoOpOCR3.Contract.TransferOwnership(&_NoOpOCR3.TransactOpts, to)
}

func (_NoOpOCR3 *NoOpOCR3TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _NoOpOCR3.Contract.TransferOwnership(&_NoOpOCR3.TransactOpts, to)
}

func (_NoOpOCR3 *NoOpOCR3Transactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _NoOpOCR3.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_NoOpOCR3 *NoOpOCR3Session) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _NoOpOCR3.Contract.Transmit(&_NoOpOCR3.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_NoOpOCR3 *NoOpOCR3TransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _NoOpOCR3.Contract.Transmit(&_NoOpOCR3.TransactOpts, reportContext, report, rs, ss, rawVs)
}

type NoOpOCR3ConfigSetIterator struct {
	Event *NoOpOCR3ConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NoOpOCR3ConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoOpOCR3ConfigSet)
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
		it.Event = new(NoOpOCR3ConfigSet)
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

func (it *NoOpOCR3ConfigSetIterator) Error() error {
	return it.fail
}

func (it *NoOpOCR3ConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NoOpOCR3ConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_NoOpOCR3 *NoOpOCR3Filterer) FilterConfigSet(opts *bind.FilterOpts) (*NoOpOCR3ConfigSetIterator, error) {

	logs, sub, err := _NoOpOCR3.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &NoOpOCR3ConfigSetIterator{contract: _NoOpOCR3.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_NoOpOCR3 *NoOpOCR3Filterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *NoOpOCR3ConfigSet) (event.Subscription, error) {

	logs, sub, err := _NoOpOCR3.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NoOpOCR3ConfigSet)
				if err := _NoOpOCR3.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_NoOpOCR3 *NoOpOCR3Filterer) ParseConfigSet(log types.Log) (*NoOpOCR3ConfigSet, error) {
	event := new(NoOpOCR3ConfigSet)
	if err := _NoOpOCR3.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NoOpOCR3OwnershipTransferRequestedIterator struct {
	Event *NoOpOCR3OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NoOpOCR3OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoOpOCR3OwnershipTransferRequested)
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
		it.Event = new(NoOpOCR3OwnershipTransferRequested)
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

func (it *NoOpOCR3OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *NoOpOCR3OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NoOpOCR3OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_NoOpOCR3 *NoOpOCR3Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NoOpOCR3OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NoOpOCR3.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NoOpOCR3OwnershipTransferRequestedIterator{contract: _NoOpOCR3.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_NoOpOCR3 *NoOpOCR3Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *NoOpOCR3OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NoOpOCR3.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NoOpOCR3OwnershipTransferRequested)
				if err := _NoOpOCR3.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_NoOpOCR3 *NoOpOCR3Filterer) ParseOwnershipTransferRequested(log types.Log) (*NoOpOCR3OwnershipTransferRequested, error) {
	event := new(NoOpOCR3OwnershipTransferRequested)
	if err := _NoOpOCR3.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NoOpOCR3OwnershipTransferredIterator struct {
	Event *NoOpOCR3OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NoOpOCR3OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoOpOCR3OwnershipTransferred)
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
		it.Event = new(NoOpOCR3OwnershipTransferred)
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

func (it *NoOpOCR3OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *NoOpOCR3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NoOpOCR3OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_NoOpOCR3 *NoOpOCR3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NoOpOCR3OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NoOpOCR3.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NoOpOCR3OwnershipTransferredIterator{contract: _NoOpOCR3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_NoOpOCR3 *NoOpOCR3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NoOpOCR3OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NoOpOCR3.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NoOpOCR3OwnershipTransferred)
				if err := _NoOpOCR3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_NoOpOCR3 *NoOpOCR3Filterer) ParseOwnershipTransferred(log types.Log) (*NoOpOCR3OwnershipTransferred, error) {
	event := new(NoOpOCR3OwnershipTransferred)
	if err := _NoOpOCR3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NoOpOCR3TransmittedIterator struct {
	Event *NoOpOCR3Transmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NoOpOCR3TransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoOpOCR3Transmitted)
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
		it.Event = new(NoOpOCR3Transmitted)
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

func (it *NoOpOCR3TransmittedIterator) Error() error {
	return it.fail
}

func (it *NoOpOCR3TransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NoOpOCR3Transmitted struct {
	ConfigDigest   [32]byte
	SequenceNumber uint64
	Raw            types.Log
}

func (_NoOpOCR3 *NoOpOCR3Filterer) FilterTransmitted(opts *bind.FilterOpts) (*NoOpOCR3TransmittedIterator, error) {

	logs, sub, err := _NoOpOCR3.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &NoOpOCR3TransmittedIterator{contract: _NoOpOCR3.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_NoOpOCR3 *NoOpOCR3Filterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *NoOpOCR3Transmitted) (event.Subscription, error) {

	logs, sub, err := _NoOpOCR3.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NoOpOCR3Transmitted)
				if err := _NoOpOCR3.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_NoOpOCR3 *NoOpOCR3Filterer) ParseTransmitted(log types.Log) (*NoOpOCR3Transmitted, error) {
	event := new(NoOpOCR3Transmitted)
	if err := _NoOpOCR3.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}
type LatestConfigDigestAndEpoch struct {
	ScanLogs       bool
	ConfigDigest   [32]byte
	SequenceNumber uint64
}

func (_NoOpOCR3 *NoOpOCR3) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _NoOpOCR3.abi.Events["ConfigSet"].ID:
		return _NoOpOCR3.ParseConfigSet(log)
	case _NoOpOCR3.abi.Events["OwnershipTransferRequested"].ID:
		return _NoOpOCR3.ParseOwnershipTransferRequested(log)
	case _NoOpOCR3.abi.Events["OwnershipTransferred"].ID:
		return _NoOpOCR3.ParseOwnershipTransferred(log)
	case _NoOpOCR3.abi.Events["Transmitted"].ID:
		return _NoOpOCR3.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (NoOpOCR3ConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (NoOpOCR3OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (NoOpOCR3OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (NoOpOCR3Transmitted) Topic() common.Hash {
	return common.HexToHash("0xe893c2681d327421d89e1cb54fbe64645b4dcea668d6826130b62cf4c6eefea2")
}

func (_NoOpOCR3 *NoOpOCR3) Address() common.Address {
	return _NoOpOCR3.address
}

type NoOpOCR3Interface interface {
	GetTransmitters(opts *bind.CallOpts) ([]common.Address, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	ToString(opts *bind.CallOpts, account common.Address) (string, error)

	ToString0(opts *bind.CallOpts, value *big.Int) (string, error)

	ToString1(opts *bind.CallOpts, data []byte) (string, error)

	ToString2(opts *bind.CallOpts, value [32]byte) (string, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetOCR3Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*NoOpOCR3ConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *NoOpOCR3ConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*NoOpOCR3ConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NoOpOCR3OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *NoOpOCR3OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*NoOpOCR3OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NoOpOCR3OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NoOpOCR3OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*NoOpOCR3OwnershipTransferred, error)

	FilterTransmitted(opts *bind.FilterOpts) (*NoOpOCR3TransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *NoOpOCR3Transmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*NoOpOCR3Transmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
