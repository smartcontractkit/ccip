// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ccipReceiver

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

type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type ICCIPClientBaseapprovedSenderUpdate struct {
	DestChainSelector uint64
	Sender            []byte
}

var CCIPReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrorCase\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageNotFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"MessageFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"disableChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraArgsBytes\",\"type\":\"bytes\"}],\"name\":\"enableChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"getMessageContents\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"getMessageStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"senderAddr\",\"type\":\"bytes\"}],\"name\":\"isApprovedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"processMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"retryFailedMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"s_chains\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgsBytes\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"simRevert\",\"type\":\"bool\"}],\"name\":\"setSimRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"internalType\":\"structICCIPClientBase.approvedSenderUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"internalType\":\"structICCIPClientBase.approvedSenderUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"}],\"name\":\"updateApprovedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNativeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620029e6380380620029e68339810160408190526200003491620001a8565b8033806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf81620000fd565b5050506001600160a01b038116620000ea576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b031660805250620001da565b336001600160a01b03821603620001575760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001bb57600080fd5b81516001600160a01b0381168114620001d357600080fd5b9392505050565b6080516127e9620001fd6000396000818161035e0152610b4801526127e96000f3fe60806040526004361061011b5760003560e01c806379ba50971161009c578063b0f479a11161006e578063d8469e4011610056578063d8469e40146103a2578063e4ca8754146103c2578063f2fde38b146103e257005b8063b0f479a11461034f578063cf6730f81461038257005b806379ba5097146102ae5780638462a2b9146102c357806385572ffb146102e35780638da5cb5b1461030357005b806352f813c3116100ed5780635dc5ebdb116100d55780635dc5ebdb146102335780635e35359e146102615780636939cd971461028157005b806352f813c3146101f3578063536c6bfa1461021357005b80630e958d6b14610124578063181f5a771461015957806341eade46146101a55780635075a9d4146101c557005b3661012257005b005b34801561013057600080fd5b5061014461013f366004611bc1565b610402565b60405190151581526020015b60405180910390f35b34801561016557600080fd5b50604080518082018252601681527f43434950526563656976657220312e302e302d64657600000000000000000000602082015290516101509190611c84565b3480156101b157600080fd5b506101226101c0366004611c97565b61044c565b3480156101d157600080fd5b506101e56101e0366004611cb4565b61048b565b604051908152602001610150565b3480156101ff57600080fd5b5061012261020e366004611cdb565b61049e565b34801561021f57600080fd5b5061012261022e366004611d1a565b6104d7565b34801561023f57600080fd5b5061025361024e366004611c97565b6104ed565b604051610150929190611d46565b34801561026d57600080fd5b5061012261027c366004611d74565b610619565b34801561028d57600080fd5b506102a161029c366004611cb4565b610642565b6040516101509190611db5565b3480156102ba57600080fd5b5061012261084d565b3480156102cf57600080fd5b506101226102de366004611ee1565b61094f565b3480156102ef57600080fd5b506101226102fe366004611f4d565b610b30565b34801561030f57600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610150565b34801561035b57600080fd5b507f000000000000000000000000000000000000000000000000000000000000000061032a565b34801561038e57600080fd5b5061012261039d366004611f4d565b610d57565b3480156103ae57600080fd5b506101226103bd366004611f88565b610e92565b3480156103ce57600080fd5b506101226103dd366004611cb4565b610ef4565b3480156103ee57600080fd5b506101226103fd36600461200b565b611172565b67ffffffffffffffff8316600090815260026020819052604080832090519101906104309085908590612028565b9081526040519081900360200190205460ff1690509392505050565b610454611186565b67ffffffffffffffff81166000908152600260205260408120906104788282611b14565b610486600183016000611b14565b505050565b6000610498600483611209565b92915050565b6104a6611186565b600780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b6104df611186565b6104e9828261121c565b5050565b60026020526000908152604090208054819061050890612038565b80601f016020809104026020016040519081016040528092919081815260200182805461053490612038565b80156105815780601f1061055657610100808354040283529160200191610581565b820191906000526020600020905b81548152906001019060200180831161056457829003601f168201915b50505050509080600101805461059690612038565b80601f01602080910402602001604051908101604052809291908181526020018280546105c290612038565b801561060f5780601f106105e45761010080835404028352916020019161060f565b820191906000526020600020905b8154815290600101906020018083116105f257829003601f168201915b5050505050905082565b610621611186565b61048673ffffffffffffffffffffffffffffffffffffffff84168383611376565b6040805160a08082018352600080835260208084018290526060848601819052808501819052608085015285825260038152908490208451928301855280548352600181015467ffffffffffffffff16918301919091526002810180549394929391928401916106b190612038565b80601f01602080910402602001604051908101604052809291908181526020018280546106dd90612038565b801561072a5780601f106106ff5761010080835404028352916020019161072a565b820191906000526020600020905b81548152906001019060200180831161070d57829003601f168201915b5050505050815260200160038201805461074390612038565b80601f016020809104026020016040519081016040528092919081815260200182805461076f90612038565b80156107bc5780601f10610791576101008083540402835291602001916107bc565b820191906000526020600020905b81548152906001019060200180831161079f57829003601f168201915b5050505050815260200160048201805480602002602001604051908101604052809291908181526020016000905b8282101561083f5760008481526020908190206040805180820190915260028502909101805473ffffffffffffffffffffffffffffffffffffffff1682526001908101548284015290835290920191016107ea565b505050915250909392505050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146108d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610957611186565b60005b81811015610a3a57600260008484848181106109785761097861208b565b905060200281019061098a91906120ba565b610998906020810190611c97565b67ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206002018383838181106109cf576109cf61208b565b90506020028101906109e191906120ba565b6109ef9060208101906120f8565b6040516109fd929190612028565b90815260405190819003602001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905560010161095a565b5060005b83811015610b2957600160026000878785818110610a5e57610a5e61208b565b9050602002810190610a7091906120ba565b610a7e906020810190611c97565b67ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600201868684818110610ab557610ab561208b565b9050602002810190610ac791906120ba565b610ad59060208101906120f8565b604051610ae3929190612028565b90815260405190819003602001902080549115157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909216919091179055600101610a3e565b5050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610ba1576040517fd7f733340000000000000000000000000000000000000000000000000000000081523360048201526024016108ca565b610bb16040820160208301611c97565b67ffffffffffffffff811660009081526002602052604090208054610bd590612038565b9050600003610c1c576040517fd79f2ea400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016108ca565b6040517fcf6730f8000000000000000000000000000000000000000000000000000000008152309063cf6730f890610c5890859060040161226a565b600060405180830381600087803b158015610c7257600080fd5b505af1925050508015610c83575060015b610d27573d808015610cb1576040519150601f19603f3d011682016040523d82523d6000602084013e610cb6565b606091505b50610cc8833560015b60049190611403565b50823560009081526003602052604090208390610ce5828261266b565b50506040518335907f55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f90610d1a908490611c84565b60405180910390a2505050565b6040518235907fdf6958669026659bac75ba986685e11a7d271284989f565f2802522663e9a70f90600090a25050565b333014610d90576040517f14d4a4e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610da06040820160208301611c97565b610dad60408301836120f8565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525067ffffffffffffffff8616815260026020819052604091829020915191019350610e0a925084915061276b565b9081526040519081900360200190205460ff16610e5557806040517f5075bb380000000000000000000000000000000000000000000000000000000081526004016108ca9190611c84565b60075460ff1615610486576040517f79f79e0b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e9a611186565b67ffffffffffffffff85166000908152600260205260409020610ebe8486836123ef565b508015610b295767ffffffffffffffff85166000908152600260205260409020600101610eec8284836123ef565b505050505050565b610efc611186565b6001610f09600483611209565b14610f43576040517fb6e78260000000000000000000000000000000000000000000000000000000008152600481018290526024016108ca565b610f4e816000610cbf565b506000818152600360209081526040808320815160a08101835281548152600182015467ffffffffffffffff16938101939093526002810180549192840191610f9690612038565b80601f0160208091040260200160405190810160405280929190818152602001828054610fc290612038565b801561100f5780601f10610fe45761010080835404028352916020019161100f565b820191906000526020600020905b815481529060010190602001808311610ff257829003601f168201915b5050505050815260200160038201805461102890612038565b80601f016020809104026020016040519081016040528092919081815260200182805461105490612038565b80156110a15780601f10611076576101008083540402835291602001916110a1565b820191906000526020600020905b81548152906001019060200180831161108457829003601f168201915b5050505050815260200160048201805480602002602001604051908101604052809291908181526020016000905b828210156111245760008481526020908190206040805180820190915260028502909101805473ffffffffffffffffffffffffffffffffffffffff1682526001908101548284015290835290920191016110cf565b5050505081525050905061113781611418565b6111426004836114be565b5060405182907fef3bf8c64bc480286c4f3503b870ceb23e648d2d902e31fb7bb46680da6de8ad90600090a25050565b61117a611186565b611183816114ca565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314611207576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016108ca565b565b600061121583836115bf565b9392505050565b80471015611286576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e636500000060448201526064016108ca565b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d80600081146112e0576040519150601f19603f3d011682016040523d82523d6000602084013e6112e5565b606091505b5050905080610486576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d6179206861766520726576657274656400000000000060648201526084016108ca565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610486908490611649565b6000611410848484611755565b949350505050565b60005b8160800151518110156104e9576000826080015182815181106114405761144061208b565b60200260200101516020015190506000836080015183815181106114665761146661208b565b60200260200101516000015190506114b461149660005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff83169084611376565b505060010161141b565b60006112158383611772565b3373ffffffffffffffffffffffffffffffffffffffff821603611549576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016108ca565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000818152600283016020526040812054801515806115e357506115e3848461178f565b611215576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b6579000060448201526064016108ca565b60006116ab826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661179b9092919063ffffffff16565b80519091501561048657808060200190518101906116c9919061277d565b610486576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016108ca565b6000828152600284016020526040812082905561141084846117aa565b6000818152600283016020526040812081905561121583836117b6565b600061121583836117c2565b606061141084846000856117da565b600061121583836118f3565b60006112158383611942565b60008181526001830160205260408120541515611215565b60608247101561186c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016108ca565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611895919061276b565b60006040518083038185875af1925050503d80600081146118d2576040519150601f19603f3d011682016040523d82523d6000602084013e6118d7565b606091505b50915091506118e887838387611a35565b979650505050505050565b600081815260018301602052604081205461193a57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610498565b506000610498565b60008181526001830160205260408120548015611a2b57600061196660018361279a565b855490915060009061197a9060019061279a565b90508181146119df57600086600001828154811061199a5761199a61208b565b90600052602060002001549050808760000184815481106119bd576119bd61208b565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806119f0576119f06127ad565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610498565b6000915050610498565b60608315611acb578251600003611ac45773ffffffffffffffffffffffffffffffffffffffff85163b611ac4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016108ca565b5081611410565b6114108383815115611ae05781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ca9190611c84565b508054611b2090612038565b6000825580601f10611b30575050565b601f01602090049060005260206000209081019061118391905b80821115611b5e5760008155600101611b4a565b5090565b67ffffffffffffffff8116811461118357600080fd5b60008083601f840112611b8a57600080fd5b50813567ffffffffffffffff811115611ba257600080fd5b602083019150836020828501011115611bba57600080fd5b9250929050565b600080600060408486031215611bd657600080fd5b8335611be181611b62565b9250602084013567ffffffffffffffff811115611bfd57600080fd5b611c0986828701611b78565b9497909650939450505050565b60005b83811015611c31578181015183820152602001611c19565b50506000910152565b60008151808452611c52816020860160208601611c16565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006112156020830184611c3a565b600060208284031215611ca957600080fd5b813561121581611b62565b600060208284031215611cc657600080fd5b5035919050565b801515811461118357600080fd5b600060208284031215611ced57600080fd5b813561121581611ccd565b73ffffffffffffffffffffffffffffffffffffffff8116811461118357600080fd5b60008060408385031215611d2d57600080fd5b8235611d3881611cf8565b946020939093013593505050565b604081526000611d596040830185611c3a565b8281036020840152611d6b8185611c3a565b95945050505050565b600080600060608486031215611d8957600080fd5b8335611d9481611cf8565b92506020840135611da481611cf8565b929592945050506040919091013590565b6000602080835283518184015280840151604067ffffffffffffffff821660408601526040860151915060a06060860152611df360c0860183611c3a565b915060608601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe080878503016080880152611e2f8483611c3a565b608089015188820390920160a089015281518082529186019450600092508501905b80831015611e90578451805173ffffffffffffffffffffffffffffffffffffffff16835286015186830152938501936001929092019190830190611e51565b50979650505050505050565b60008083601f840112611eae57600080fd5b50813567ffffffffffffffff811115611ec657600080fd5b6020830191508360208260051b8501011115611bba57600080fd5b60008060008060408587031215611ef757600080fd5b843567ffffffffffffffff80821115611f0f57600080fd5b611f1b88838901611e9c565b90965094506020870135915080821115611f3457600080fd5b50611f4187828801611e9c565b95989497509550505050565b600060208284031215611f5f57600080fd5b813567ffffffffffffffff811115611f7657600080fd5b820160a0818503121561121557600080fd5b600080600080600060608688031215611fa057600080fd5b8535611fab81611b62565b9450602086013567ffffffffffffffff80821115611fc857600080fd5b611fd489838a01611b78565b90965094506040880135915080821115611fed57600080fd5b50611ffa88828901611b78565b969995985093965092949392505050565b60006020828403121561201d57600080fd5b813561121581611cf8565b8183823760009101908152919050565b600181811c9082168061204c57607f821691505b602082108103612085577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc18336030181126120ee57600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261212d57600080fd5b83018035915067ffffffffffffffff82111561214857600080fd5b602001915036819003821315611bba57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261219257600080fd5b830160208101925035905067ffffffffffffffff8111156121b257600080fd5b803603821315611bba57600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8183526000602080850194508260005b8581101561225f57813561222d81611cf8565b73ffffffffffffffffffffffffffffffffffffffff16875281830135838801526040968701969091019060010161221a565b509495945050505050565b60208152813560208201526000602083013561228581611b62565b67ffffffffffffffff80821660408501526122a3604086018661215d565b925060a060608601526122ba60c0860184836121c1565b9250506122ca606086018661215d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808786030160808801526123008583856121c1565b9450608088013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe188360301831261233957600080fd5b6020928801928301923591508382111561235257600080fd5b8160061b360383131561236457600080fd5b8685030160a08701526118e884828461220a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f821115610486576000816000526020600020601f850160051c810160208610156123d05750805b601f850160051c820191505b81811015610eec578281556001016123dc565b67ffffffffffffffff83111561240757612407612378565b61241b836124158354612038565b836123a7565b6000601f84116001811461246d57600085156124375750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355610b29565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b828110156124bc578685013582556020948501946001909201910161249c565b50868210156124f7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b813561254381611cf8565b73ffffffffffffffffffffffffffffffffffffffff81167fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825550602082013560018201555050565b680100000000000000008311156125a9576125a9612378565b8054838255808410156126365760017f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80831683146125ea576125ea612509565b80861686146125fb576125fb612509565b5060008360005260206000208360011b81018760011b820191505b80821015612631578282558284830155600282019150612616565b505050505b5060008181526020812083915b85811015610eec576126558383612538565b6040929092019160029190910190600101612643565b8135815560018101602083013561268181611b62565b67ffffffffffffffff8082167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008454161783556126c160408601866120f8565b935091506126d38383600287016123ef565b6126e060608601866120f8565b935091506126f28383600387016123ef565b608085013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe185360301831261272957600080fd5b91840191823591508082111561273e57600080fd5b506020820191508060061b360382131561275757600080fd5b612765818360048601612590565b50505050565b600082516120ee818460208701611c16565b60006020828403121561278f57600080fd5b815161121581611ccd565b8181038181111561049857610498612509565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000818000a",
}

var CCIPReceiverABI = CCIPReceiverMetaData.ABI

var CCIPReceiverBin = CCIPReceiverMetaData.Bin

func DeployCCIPReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address) (common.Address, *types.Transaction, *CCIPReceiver, error) {
	parsed, err := CCIPReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCIPReceiverBin), backend, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCIPReceiver{address: address, abi: *parsed, CCIPReceiverCaller: CCIPReceiverCaller{contract: contract}, CCIPReceiverTransactor: CCIPReceiverTransactor{contract: contract}, CCIPReceiverFilterer: CCIPReceiverFilterer{contract: contract}}, nil
}

type CCIPReceiver struct {
	address common.Address
	abi     abi.ABI
	CCIPReceiverCaller
	CCIPReceiverTransactor
	CCIPReceiverFilterer
}

type CCIPReceiverCaller struct {
	contract *bind.BoundContract
}

type CCIPReceiverTransactor struct {
	contract *bind.BoundContract
}

type CCIPReceiverFilterer struct {
	contract *bind.BoundContract
}

type CCIPReceiverSession struct {
	Contract     *CCIPReceiver
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CCIPReceiverCallerSession struct {
	Contract *CCIPReceiverCaller
	CallOpts bind.CallOpts
}

type CCIPReceiverTransactorSession struct {
	Contract     *CCIPReceiverTransactor
	TransactOpts bind.TransactOpts
}

type CCIPReceiverRaw struct {
	Contract *CCIPReceiver
}

type CCIPReceiverCallerRaw struct {
	Contract *CCIPReceiverCaller
}

type CCIPReceiverTransactorRaw struct {
	Contract *CCIPReceiverTransactor
}

func NewCCIPReceiver(address common.Address, backend bind.ContractBackend) (*CCIPReceiver, error) {
	abi, err := abi.JSON(strings.NewReader(CCIPReceiverABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCCIPReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCIPReceiver{address: address, abi: abi, CCIPReceiverCaller: CCIPReceiverCaller{contract: contract}, CCIPReceiverTransactor: CCIPReceiverTransactor{contract: contract}, CCIPReceiverFilterer: CCIPReceiverFilterer{contract: contract}}, nil
}

func NewCCIPReceiverCaller(address common.Address, caller bind.ContractCaller) (*CCIPReceiverCaller, error) {
	contract, err := bindCCIPReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCIPReceiverCaller{contract: contract}, nil
}

func NewCCIPReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*CCIPReceiverTransactor, error) {
	contract, err := bindCCIPReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCIPReceiverTransactor{contract: contract}, nil
}

func NewCCIPReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*CCIPReceiverFilterer, error) {
	contract, err := bindCCIPReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCIPReceiverFilterer{contract: contract}, nil
}

func bindCCIPReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCIPReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CCIPReceiver *CCIPReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCIPReceiver.Contract.CCIPReceiverCaller.contract.Call(opts, result, method, params...)
}

func (_CCIPReceiver *CCIPReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.CCIPReceiverTransactor.contract.Transfer(opts)
}

func (_CCIPReceiver *CCIPReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.CCIPReceiverTransactor.contract.Transact(opts, method, params...)
}

func (_CCIPReceiver *CCIPReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCIPReceiver.Contract.contract.Call(opts, result, method, params...)
}

func (_CCIPReceiver *CCIPReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.contract.Transfer(opts)
}

func (_CCIPReceiver *CCIPReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.contract.Transact(opts, method, params...)
}

func (_CCIPReceiver *CCIPReceiverCaller) GetMessageContents(opts *bind.CallOpts, messageId [32]byte) (ClientAny2EVMMessage, error) {
	var out []interface{}
	err := _CCIPReceiver.contract.Call(opts, &out, "getMessageContents", messageId)

	if err != nil {
		return *new(ClientAny2EVMMessage), err
	}

	out0 := *abi.ConvertType(out[0], new(ClientAny2EVMMessage)).(*ClientAny2EVMMessage)

	return out0, err

}

func (_CCIPReceiver *CCIPReceiverSession) GetMessageContents(messageId [32]byte) (ClientAny2EVMMessage, error) {
	return _CCIPReceiver.Contract.GetMessageContents(&_CCIPReceiver.CallOpts, messageId)
}

func (_CCIPReceiver *CCIPReceiverCallerSession) GetMessageContents(messageId [32]byte) (ClientAny2EVMMessage, error) {
	return _CCIPReceiver.Contract.GetMessageContents(&_CCIPReceiver.CallOpts, messageId)
}

func (_CCIPReceiver *CCIPReceiverCaller) GetMessageStatus(opts *bind.CallOpts, messageId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CCIPReceiver.contract.Call(opts, &out, "getMessageStatus", messageId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CCIPReceiver *CCIPReceiverSession) GetMessageStatus(messageId [32]byte) (*big.Int, error) {
	return _CCIPReceiver.Contract.GetMessageStatus(&_CCIPReceiver.CallOpts, messageId)
}

func (_CCIPReceiver *CCIPReceiverCallerSession) GetMessageStatus(messageId [32]byte) (*big.Int, error) {
	return _CCIPReceiver.Contract.GetMessageStatus(&_CCIPReceiver.CallOpts, messageId)
}

func (_CCIPReceiver *CCIPReceiverCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCIPReceiver.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCIPReceiver *CCIPReceiverSession) GetRouter() (common.Address, error) {
	return _CCIPReceiver.Contract.GetRouter(&_CCIPReceiver.CallOpts)
}

func (_CCIPReceiver *CCIPReceiverCallerSession) GetRouter() (common.Address, error) {
	return _CCIPReceiver.Contract.GetRouter(&_CCIPReceiver.CallOpts)
}

func (_CCIPReceiver *CCIPReceiverCaller) IsApprovedSender(opts *bind.CallOpts, sourceChainSelector uint64, senderAddr []byte) (bool, error) {
	var out []interface{}
	err := _CCIPReceiver.contract.Call(opts, &out, "isApprovedSender", sourceChainSelector, senderAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CCIPReceiver *CCIPReceiverSession) IsApprovedSender(sourceChainSelector uint64, senderAddr []byte) (bool, error) {
	return _CCIPReceiver.Contract.IsApprovedSender(&_CCIPReceiver.CallOpts, sourceChainSelector, senderAddr)
}

func (_CCIPReceiver *CCIPReceiverCallerSession) IsApprovedSender(sourceChainSelector uint64, senderAddr []byte) (bool, error) {
	return _CCIPReceiver.Contract.IsApprovedSender(&_CCIPReceiver.CallOpts, sourceChainSelector, senderAddr)
}

func (_CCIPReceiver *CCIPReceiverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCIPReceiver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCIPReceiver *CCIPReceiverSession) Owner() (common.Address, error) {
	return _CCIPReceiver.Contract.Owner(&_CCIPReceiver.CallOpts)
}

func (_CCIPReceiver *CCIPReceiverCallerSession) Owner() (common.Address, error) {
	return _CCIPReceiver.Contract.Owner(&_CCIPReceiver.CallOpts)
}

func (_CCIPReceiver *CCIPReceiverCaller) SChains(opts *bind.CallOpts, arg0 uint64) (SChains,

	error) {
	var out []interface{}
	err := _CCIPReceiver.contract.Call(opts, &out, "s_chains", arg0)

	outstruct := new(SChains)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recipient = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ExtraArgsBytes = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

func (_CCIPReceiver *CCIPReceiverSession) SChains(arg0 uint64) (SChains,

	error) {
	return _CCIPReceiver.Contract.SChains(&_CCIPReceiver.CallOpts, arg0)
}

func (_CCIPReceiver *CCIPReceiverCallerSession) SChains(arg0 uint64) (SChains,

	error) {
	return _CCIPReceiver.Contract.SChains(&_CCIPReceiver.CallOpts, arg0)
}

func (_CCIPReceiver *CCIPReceiverCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CCIPReceiver.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CCIPReceiver *CCIPReceiverSession) TypeAndVersion() (string, error) {
	return _CCIPReceiver.Contract.TypeAndVersion(&_CCIPReceiver.CallOpts)
}

func (_CCIPReceiver *CCIPReceiverCallerSession) TypeAndVersion() (string, error) {
	return _CCIPReceiver.Contract.TypeAndVersion(&_CCIPReceiver.CallOpts)
}

func (_CCIPReceiver *CCIPReceiverTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "acceptOwnership")
}

func (_CCIPReceiver *CCIPReceiverSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCIPReceiver.Contract.AcceptOwnership(&_CCIPReceiver.TransactOpts)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCIPReceiver.Contract.AcceptOwnership(&_CCIPReceiver.TransactOpts)
}

func (_CCIPReceiver *CCIPReceiverTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "ccipReceive", message)
}

func (_CCIPReceiver *CCIPReceiverSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.CcipReceive(&_CCIPReceiver.TransactOpts, message)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.CcipReceive(&_CCIPReceiver.TransactOpts, message)
}

func (_CCIPReceiver *CCIPReceiverTransactor) DisableChain(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "disableChain", chainSelector)
}

func (_CCIPReceiver *CCIPReceiverSession) DisableChain(chainSelector uint64) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.DisableChain(&_CCIPReceiver.TransactOpts, chainSelector)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) DisableChain(chainSelector uint64) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.DisableChain(&_CCIPReceiver.TransactOpts, chainSelector)
}

func (_CCIPReceiver *CCIPReceiverTransactor) EnableChain(opts *bind.TransactOpts, chainSelector uint64, recipient []byte, _extraArgsBytes []byte) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "enableChain", chainSelector, recipient, _extraArgsBytes)
}

func (_CCIPReceiver *CCIPReceiverSession) EnableChain(chainSelector uint64, recipient []byte, _extraArgsBytes []byte) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.EnableChain(&_CCIPReceiver.TransactOpts, chainSelector, recipient, _extraArgsBytes)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) EnableChain(chainSelector uint64, recipient []byte, _extraArgsBytes []byte) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.EnableChain(&_CCIPReceiver.TransactOpts, chainSelector, recipient, _extraArgsBytes)
}

func (_CCIPReceiver *CCIPReceiverTransactor) ProcessMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "processMessage", message)
}

func (_CCIPReceiver *CCIPReceiverSession) ProcessMessage(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.ProcessMessage(&_CCIPReceiver.TransactOpts, message)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) ProcessMessage(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.ProcessMessage(&_CCIPReceiver.TransactOpts, message)
}

func (_CCIPReceiver *CCIPReceiverTransactor) RetryFailedMessage(opts *bind.TransactOpts, messageId [32]byte) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "retryFailedMessage", messageId)
}

func (_CCIPReceiver *CCIPReceiverSession) RetryFailedMessage(messageId [32]byte) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.RetryFailedMessage(&_CCIPReceiver.TransactOpts, messageId)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) RetryFailedMessage(messageId [32]byte) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.RetryFailedMessage(&_CCIPReceiver.TransactOpts, messageId)
}

func (_CCIPReceiver *CCIPReceiverTransactor) SetSimRevert(opts *bind.TransactOpts, simRevert bool) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "setSimRevert", simRevert)
}

func (_CCIPReceiver *CCIPReceiverSession) SetSimRevert(simRevert bool) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.SetSimRevert(&_CCIPReceiver.TransactOpts, simRevert)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) SetSimRevert(simRevert bool) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.SetSimRevert(&_CCIPReceiver.TransactOpts, simRevert)
}

func (_CCIPReceiver *CCIPReceiverTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "transferOwnership", to)
}

func (_CCIPReceiver *CCIPReceiverSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.TransferOwnership(&_CCIPReceiver.TransactOpts, to)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.TransferOwnership(&_CCIPReceiver.TransactOpts, to)
}

func (_CCIPReceiver *CCIPReceiverTransactor) UpdateApprovedSenders(opts *bind.TransactOpts, adds []ICCIPClientBaseapprovedSenderUpdate, removes []ICCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "updateApprovedSenders", adds, removes)
}

func (_CCIPReceiver *CCIPReceiverSession) UpdateApprovedSenders(adds []ICCIPClientBaseapprovedSenderUpdate, removes []ICCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.UpdateApprovedSenders(&_CCIPReceiver.TransactOpts, adds, removes)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) UpdateApprovedSenders(adds []ICCIPClientBaseapprovedSenderUpdate, removes []ICCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.UpdateApprovedSenders(&_CCIPReceiver.TransactOpts, adds, removes)
}

func (_CCIPReceiver *CCIPReceiverTransactor) WithdrawNativeToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "withdrawNativeToken", to, amount)
}

func (_CCIPReceiver *CCIPReceiverSession) WithdrawNativeToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.WithdrawNativeToken(&_CCIPReceiver.TransactOpts, to, amount)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) WithdrawNativeToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.WithdrawNativeToken(&_CCIPReceiver.TransactOpts, to, amount)
}

func (_CCIPReceiver *CCIPReceiverTransactor) WithdrawTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "withdrawTokens", token, to, amount)
}

func (_CCIPReceiver *CCIPReceiverSession) WithdrawTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.WithdrawTokens(&_CCIPReceiver.TransactOpts, token, to, amount)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) WithdrawTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.WithdrawTokens(&_CCIPReceiver.TransactOpts, token, to, amount)
}

func (_CCIPReceiver *CCIPReceiverTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CCIPReceiver.contract.RawTransact(opts, calldata)
}

func (_CCIPReceiver *CCIPReceiverSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.Fallback(&_CCIPReceiver.TransactOpts, calldata)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.Fallback(&_CCIPReceiver.TransactOpts, calldata)
}

func (_CCIPReceiver *CCIPReceiverTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPReceiver.contract.RawTransact(opts, nil)
}

func (_CCIPReceiver *CCIPReceiverSession) Receive() (*types.Transaction, error) {
	return _CCIPReceiver.Contract.Receive(&_CCIPReceiver.TransactOpts)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) Receive() (*types.Transaction, error) {
	return _CCIPReceiver.Contract.Receive(&_CCIPReceiver.TransactOpts)
}

type CCIPReceiverMessageFailedIterator struct {
	Event *CCIPReceiverMessageFailed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPReceiverMessageFailedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPReceiverMessageFailed)
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
		it.Event = new(CCIPReceiverMessageFailed)
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

func (it *CCIPReceiverMessageFailedIterator) Error() error {
	return it.fail
}

func (it *CCIPReceiverMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPReceiverMessageFailed struct {
	MessageId [32]byte
	Reason    []byte
	Raw       types.Log
}

func (_CCIPReceiver *CCIPReceiverFilterer) FilterMessageFailed(opts *bind.FilterOpts, messageId [][32]byte) (*CCIPReceiverMessageFailedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _CCIPReceiver.contract.FilterLogs(opts, "MessageFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &CCIPReceiverMessageFailedIterator{contract: _CCIPReceiver.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

func (_CCIPReceiver *CCIPReceiverFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *CCIPReceiverMessageFailed, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _CCIPReceiver.contract.WatchLogs(opts, "MessageFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPReceiverMessageFailed)
				if err := _CCIPReceiver.contract.UnpackLog(event, "MessageFailed", log); err != nil {
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

func (_CCIPReceiver *CCIPReceiverFilterer) ParseMessageFailed(log types.Log) (*CCIPReceiverMessageFailed, error) {
	event := new(CCIPReceiverMessageFailed)
	if err := _CCIPReceiver.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPReceiverMessageRecoveredIterator struct {
	Event *CCIPReceiverMessageRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPReceiverMessageRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPReceiverMessageRecovered)
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
		it.Event = new(CCIPReceiverMessageRecovered)
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

func (it *CCIPReceiverMessageRecoveredIterator) Error() error {
	return it.fail
}

func (it *CCIPReceiverMessageRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPReceiverMessageRecovered struct {
	MessageId [32]byte
	Raw       types.Log
}

func (_CCIPReceiver *CCIPReceiverFilterer) FilterMessageRecovered(opts *bind.FilterOpts, messageId [][32]byte) (*CCIPReceiverMessageRecoveredIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _CCIPReceiver.contract.FilterLogs(opts, "MessageRecovered", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &CCIPReceiverMessageRecoveredIterator{contract: _CCIPReceiver.contract, event: "MessageRecovered", logs: logs, sub: sub}, nil
}

func (_CCIPReceiver *CCIPReceiverFilterer) WatchMessageRecovered(opts *bind.WatchOpts, sink chan<- *CCIPReceiverMessageRecovered, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _CCIPReceiver.contract.WatchLogs(opts, "MessageRecovered", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPReceiverMessageRecovered)
				if err := _CCIPReceiver.contract.UnpackLog(event, "MessageRecovered", log); err != nil {
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

func (_CCIPReceiver *CCIPReceiverFilterer) ParseMessageRecovered(log types.Log) (*CCIPReceiverMessageRecovered, error) {
	event := new(CCIPReceiverMessageRecovered)
	if err := _CCIPReceiver.contract.UnpackLog(event, "MessageRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPReceiverMessageSucceededIterator struct {
	Event *CCIPReceiverMessageSucceeded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPReceiverMessageSucceededIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPReceiverMessageSucceeded)
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
		it.Event = new(CCIPReceiverMessageSucceeded)
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

func (it *CCIPReceiverMessageSucceededIterator) Error() error {
	return it.fail
}

func (it *CCIPReceiverMessageSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPReceiverMessageSucceeded struct {
	MessageId [32]byte
	Raw       types.Log
}

func (_CCIPReceiver *CCIPReceiverFilterer) FilterMessageSucceeded(opts *bind.FilterOpts, messageId [][32]byte) (*CCIPReceiverMessageSucceededIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _CCIPReceiver.contract.FilterLogs(opts, "MessageSucceeded", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &CCIPReceiverMessageSucceededIterator{contract: _CCIPReceiver.contract, event: "MessageSucceeded", logs: logs, sub: sub}, nil
}

func (_CCIPReceiver *CCIPReceiverFilterer) WatchMessageSucceeded(opts *bind.WatchOpts, sink chan<- *CCIPReceiverMessageSucceeded, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _CCIPReceiver.contract.WatchLogs(opts, "MessageSucceeded", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPReceiverMessageSucceeded)
				if err := _CCIPReceiver.contract.UnpackLog(event, "MessageSucceeded", log); err != nil {
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

func (_CCIPReceiver *CCIPReceiverFilterer) ParseMessageSucceeded(log types.Log) (*CCIPReceiverMessageSucceeded, error) {
	event := new(CCIPReceiverMessageSucceeded)
	if err := _CCIPReceiver.contract.UnpackLog(event, "MessageSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPReceiverOwnershipTransferRequestedIterator struct {
	Event *CCIPReceiverOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPReceiverOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPReceiverOwnershipTransferRequested)
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
		it.Event = new(CCIPReceiverOwnershipTransferRequested)
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

func (it *CCIPReceiverOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CCIPReceiverOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPReceiverOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCIPReceiver *CCIPReceiverFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCIPReceiverOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCIPReceiver.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCIPReceiverOwnershipTransferRequestedIterator{contract: _CCIPReceiver.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CCIPReceiver *CCIPReceiverFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCIPReceiverOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCIPReceiver.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPReceiverOwnershipTransferRequested)
				if err := _CCIPReceiver.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CCIPReceiver *CCIPReceiverFilterer) ParseOwnershipTransferRequested(log types.Log) (*CCIPReceiverOwnershipTransferRequested, error) {
	event := new(CCIPReceiverOwnershipTransferRequested)
	if err := _CCIPReceiver.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPReceiverOwnershipTransferredIterator struct {
	Event *CCIPReceiverOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPReceiverOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPReceiverOwnershipTransferred)
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
		it.Event = new(CCIPReceiverOwnershipTransferred)
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

func (it *CCIPReceiverOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CCIPReceiverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPReceiverOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCIPReceiver *CCIPReceiverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCIPReceiverOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCIPReceiver.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCIPReceiverOwnershipTransferredIterator{contract: _CCIPReceiver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CCIPReceiver *CCIPReceiverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCIPReceiverOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCIPReceiver.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPReceiverOwnershipTransferred)
				if err := _CCIPReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CCIPReceiver *CCIPReceiverFilterer) ParseOwnershipTransferred(log types.Log) (*CCIPReceiverOwnershipTransferred, error) {
	event := new(CCIPReceiverOwnershipTransferred)
	if err := _CCIPReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SChains struct {
	Recipient      []byte
	ExtraArgsBytes []byte
}

func (_CCIPReceiver *CCIPReceiver) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CCIPReceiver.abi.Events["MessageFailed"].ID:
		return _CCIPReceiver.ParseMessageFailed(log)
	case _CCIPReceiver.abi.Events["MessageRecovered"].ID:
		return _CCIPReceiver.ParseMessageRecovered(log)
	case _CCIPReceiver.abi.Events["MessageSucceeded"].ID:
		return _CCIPReceiver.ParseMessageSucceeded(log)
	case _CCIPReceiver.abi.Events["OwnershipTransferRequested"].ID:
		return _CCIPReceiver.ParseOwnershipTransferRequested(log)
	case _CCIPReceiver.abi.Events["OwnershipTransferred"].ID:
		return _CCIPReceiver.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CCIPReceiverMessageFailed) Topic() common.Hash {
	return common.HexToHash("0x55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f")
}

func (CCIPReceiverMessageRecovered) Topic() common.Hash {
	return common.HexToHash("0xef3bf8c64bc480286c4f3503b870ceb23e648d2d902e31fb7bb46680da6de8ad")
}

func (CCIPReceiverMessageSucceeded) Topic() common.Hash {
	return common.HexToHash("0xdf6958669026659bac75ba986685e11a7d271284989f565f2802522663e9a70f")
}

func (CCIPReceiverOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCIPReceiverOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_CCIPReceiver *CCIPReceiver) Address() common.Address {
	return _CCIPReceiver.address
}

type CCIPReceiverInterface interface {
	GetMessageContents(opts *bind.CallOpts, messageId [32]byte) (ClientAny2EVMMessage, error)

	GetMessageStatus(opts *bind.CallOpts, messageId [32]byte) (*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IsApprovedSender(opts *bind.CallOpts, sourceChainSelector uint64, senderAddr []byte) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SChains(opts *bind.CallOpts, arg0 uint64) (SChains,

		error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	DisableChain(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error)

	EnableChain(opts *bind.TransactOpts, chainSelector uint64, recipient []byte, _extraArgsBytes []byte) (*types.Transaction, error)

	ProcessMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	RetryFailedMessage(opts *bind.TransactOpts, messageId [32]byte) (*types.Transaction, error)

	SetSimRevert(opts *bind.TransactOpts, simRevert bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateApprovedSenders(opts *bind.TransactOpts, adds []ICCIPClientBaseapprovedSenderUpdate, removes []ICCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error)

	WithdrawNativeToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)

	WithdrawTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterMessageFailed(opts *bind.FilterOpts, messageId [][32]byte) (*CCIPReceiverMessageFailedIterator, error)

	WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *CCIPReceiverMessageFailed, messageId [][32]byte) (event.Subscription, error)

	ParseMessageFailed(log types.Log) (*CCIPReceiverMessageFailed, error)

	FilterMessageRecovered(opts *bind.FilterOpts, messageId [][32]byte) (*CCIPReceiverMessageRecoveredIterator, error)

	WatchMessageRecovered(opts *bind.WatchOpts, sink chan<- *CCIPReceiverMessageRecovered, messageId [][32]byte) (event.Subscription, error)

	ParseMessageRecovered(log types.Log) (*CCIPReceiverMessageRecovered, error)

	FilterMessageSucceeded(opts *bind.FilterOpts, messageId [][32]byte) (*CCIPReceiverMessageSucceededIterator, error)

	WatchMessageSucceeded(opts *bind.WatchOpts, sink chan<- *CCIPReceiverMessageSucceeded, messageId [][32]byte) (event.Subscription, error)

	ParseMessageSucceeded(log types.Log) (*CCIPReceiverMessageSucceeded, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCIPReceiverOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCIPReceiverOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCIPReceiverOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCIPReceiverOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCIPReceiverOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCIPReceiverOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
