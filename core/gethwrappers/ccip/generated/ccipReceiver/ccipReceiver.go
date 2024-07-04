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

type CCIPClientBaseApprovedSenderUpdate struct {
	DestChainSelector uint64
	Sender            []byte
}

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

var CCIPReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageNotFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenReceiver\",\"type\":\"address\"}],\"name\":\"MessageAbandoned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"MessageFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"abandonFailedMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"disableChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraArgsBytes\",\"type\":\"bytes\"}],\"name\":\"enableChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"getMessageContents\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"getMessageStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"senderAddr\",\"type\":\"bytes\"}],\"name\":\"isApprovedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"processMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"retryFailedMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"s_chainConfigs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgsBytes\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPClientBase.ApprovedSenderUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPClientBase.ApprovedSenderUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"}],\"name\":\"updateApprovedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNativeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162002a8738038062002a878339810160408190526200003491620001a8565b8033806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf81620000fd565b5050506001600160a01b038116620000ea576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b031660805250620001da565b336001600160a01b03821603620001575760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001bb57600080fd5b81516001600160a01b0381168114620001d357600080fd5b9392505050565b60805161288a620001fd600039600081816103760152610e59015261288a6000f3fe60806040526004361061012d5760003560e01c806379ba5097116100a5578063b0f479a111610074578063d8469e4011610059578063d8469e40146103ba578063e4ca8754146103da578063f2fde38b146103fa57600080fd5b8063b0f479a114610367578063cf6730f81461039a57600080fd5b806379ba5097146102c65780638462a2b9146102db57806385572ffb146102fb5780638da5cb5b1461031b57600080fd5b80635075a9d4116100fc5780635e35359e116100e15780635e35359e146102595780636939cd97146102795780636d62d633146102a657600080fd5b80635075a9d41461020b578063536c6bfa1461023957600080fd5b80630e958d6b14610139578063181f5a771461016e57806335f170ef146101ba57806341eade46146101e957600080fd5b3661013457005b600080fd5b34801561014557600080fd5b50610159610154366004611c99565b61041a565b60405190151581526020015b60405180910390f35b34801561017a57600080fd5b50604080518082018252601681527f43434950526563656976657220312e302e302d64657600000000000000000000602082015290516101659190611d5c565b3480156101c657600080fd5b506101da6101d5366004611d6f565b610465565b60405161016593929190611d8c565b3480156101f557600080fd5b50610209610204366004611d6f565b61059c565b005b34801561021757600080fd5b5061022b610226366004611dc3565b6105e7565b604051908152602001610165565b34801561024557600080fd5b50610209610254366004611dfe565b6105fa565b34801561026557600080fd5b50610209610274366004611e2a565b610610565b34801561028557600080fd5b50610299610294366004611dc3565b61063e565b6040516101659190611e6b565b3480156102b257600080fd5b506102096102c1366004611f52565b610849565b3480156102d257600080fd5b50610209610b63565b3480156102e757600080fd5b506102096102f6366004611fc7565b610c60565b34801561030757600080fd5b50610209610316366004612033565b610e41565b34801561032757600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610165565b34801561037357600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610342565b3480156103a657600080fd5b506102096103b5366004612033565b611074565b3480156103c657600080fd5b506102096103d536600461206e565b611173565b3480156103e657600080fd5b506102096103f5366004611dc3565b6111f4565b34801561040657600080fd5b506102096104153660046120f1565b61145e565b67ffffffffffffffff83166000908152600260205260408082209051600390910190610449908590859061210e565b9081526040519081900360200190205460ff1690509392505050565b6002602052600090815260409020805460018201805460ff909216929161048b9061211e565b80601f01602080910402602001604051908101604052809291908181526020018280546104b79061211e565b80156105045780601f106104d957610100808354040283529160200191610504565b820191906000526020600020905b8154815290600101906020018083116104e757829003601f168201915b5050505050908060020180546105199061211e565b80601f01602080910402602001604051908101604052809291908181526020018280546105459061211e565b80156105925780601f1061056757610100808354040283529160200191610592565b820191906000526020600020905b81548152906001019060200180831161057557829003601f168201915b5050505050905083565b6105a4611472565b67ffffffffffffffff16600090815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b60006105f46004836114f5565b92915050565b610602611472565b61060c8282611508565b5050565b610618611472565b61063973ffffffffffffffffffffffffffffffffffffffff84168383611662565b505050565b6040805160a08082018352600080835260208084018290526060848601819052808501819052608085015285825260038152908490208451928301855280548352600181015467ffffffffffffffff16918301919091526002810180549394929391928401916106ad9061211e565b80601f01602080910402602001604051908101604052809291908181526020018280546106d99061211e565b80156107265780601f106106fb57610100808354040283529160200191610726565b820191906000526020600020905b81548152906001019060200180831161070957829003601f168201915b5050505050815260200160038201805461073f9061211e565b80601f016020809104026020016040519081016040528092919081815260200182805461076b9061211e565b80156107b85780601f1061078d576101008083540402835291602001916107b8565b820191906000526020600020905b81548152906001019060200180831161079b57829003601f168201915b5050505050815260200160048201805480602002602001604051908101604052809291908181526020016000905b8282101561083b5760008481526020908190206040805180820190915260028502909101805473ffffffffffffffffffffffffffffffffffffffff1682526001908101548284015290835290920191016107e6565b505050915250909392505050565b610851611472565b600161085e6004846114f5565b1461089d576040517fb6e78260000000000000000000000000000000000000000000000000000000008152600481018390526024015b60405180910390fd5b6108ad8260025b600491906116ef565b506000828152600360209081526040808320815160a08101835281548152600182015467ffffffffffffffff169381019390935260028101805491928401916108f59061211e565b80601f01602080910402602001604051908101604052809291908181526020018280546109219061211e565b801561096e5780601f106109435761010080835404028352916020019161096e565b820191906000526020600020905b81548152906001019060200180831161095157829003601f168201915b505050505081526020016003820180546109879061211e565b80601f01602080910402602001604051908101604052809291908181526020018280546109b39061211e565b8015610a005780601f106109d557610100808354040283529160200191610a00565b820191906000526020600020905b8154815290600101906020018083116109e357829003601f168201915b5050505050815260200160048201805480602002602001604051908101604052809291908181526020016000905b82821015610a835760008481526020908190206040805180820190915260028502909101805473ffffffffffffffffffffffffffffffffffffffff168252600190810154828401529083529092019101610a2e565b5050505081525050905060005b816080015151811015610b1257610b0a8383608001518381518110610ab757610ab7612171565b60200260200101516020015184608001518481518110610ad957610ad9612171565b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff166116629092919063ffffffff16565b600101610a90565b5060405173ffffffffffffffffffffffffffffffffffffffff8316815283907fd5038100bd3dc9631d3c3f4f61a3e53e9d466f40c47af9897292c7b35e32a9579060200160405180910390a2505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610be4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610894565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610c68611472565b60005b81811015610d4b5760026000848484818110610c8957610c89612171565b9050602002810190610c9b91906121a0565b610ca9906020810190611d6f565b67ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600301838383818110610ce057610ce0612171565b9050602002810190610cf291906121a0565b610d009060208101906121de565b604051610d0e92919061210e565b90815260405190819003602001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055600101610c6b565b5060005b83811015610e3a57600160026000878785818110610d6f57610d6f612171565b9050602002810190610d8191906121a0565b610d8f906020810190611d6f565b67ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600301868684818110610dc657610dc6612171565b9050602002810190610dd891906121a0565b610de69060208101906121de565b604051610df492919061210e565b90815260405190819003602001902080549115157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909216919091179055600101610d4f565b5050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610eb2576040517fd7f73334000000000000000000000000000000000000000000000000000000008152336004820152602401610894565b610ec26040820160208301611d6f565b67ffffffffffffffff81166000908152600260205260409020600181018054610eea9061211e565b15905080610ef95750805460ff165b15610f3c576040517fd79f2ea400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff83166004820152602401610894565b6040517fcf6730f8000000000000000000000000000000000000000000000000000000008152309063cf6730f890610f78908690600401612350565b600060405180830381600087803b158015610f9257600080fd5b505af1925050508015610fa3575060015b611043573d808015610fd1576040519150601f19603f3d011682016040523d82523d6000602084013e610fd6565b606091505b50610fe3843560016108a4565b508335600090815260036020526040902084906110008282612749565b50506040518435907f55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f90611035908490611d5c565b60405180910390a250505050565b6040518335907fdf6958669026659bac75ba986685e11a7d271284989f565f2802522663e9a70f90600090a2505050565b3330146110ad576040517f14d4a4e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110bd6040820160208301611d6f565b6110ca60408301836121de565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525067ffffffffffffffff8616815260026020526040908190209051600390910193506111289250849150612849565b9081526040519081900360200190205460ff1661063957806040517f5075bb380000000000000000000000000000000000000000000000000000000081526004016108949190611d5c565b61117b611472565b67ffffffffffffffff85166000908152600260205260409020600181016111a38587836124d5565b5081156111bb57600281016111b98385836124d5565b505b805460ff16156111ec5780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681555b505050505050565b60016112016004836114f5565b1461123b576040517fb6e7826000000000000000000000000000000000000000000000000000000000815260048101829052602401610894565b6112468160006108a4565b506000818152600360209081526040808320815160a08101835281548152600182015467ffffffffffffffff1693810193909352600281018054919284019161128e9061211e565b80601f01602080910402602001604051908101604052809291908181526020018280546112ba9061211e565b80156113075780601f106112dc57610100808354040283529160200191611307565b820191906000526020600020905b8154815290600101906020018083116112ea57829003601f168201915b505050505081526020016003820180546113209061211e565b80601f016020809104026020016040519081016040528092919081815260200182805461134c9061211e565b80156113995780601f1061136e57610100808354040283529160200191611399565b820191906000526020600020905b81548152906001019060200180831161137c57829003601f168201915b5050505050815260200160048201805480602002602001604051908101604052809291908181526020016000905b8282101561141c5760008481526020908190206040805180820190915260028502909101805473ffffffffffffffffffffffffffffffffffffffff1682526001908101548284015290835290920191016113c7565b5050505081525050905061142f81611704565b60405182907fef3bf8c64bc480286c4f3503b870ceb23e648d2d902e31fb7bb46680da6de8ad90600090a25050565b611466611472565b61146f8161170c565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146114f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610894565b565b60006115018383611801565b9392505050565b80471015611572576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610894565b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d80600081146115cc576040519150601f19603f3d011682016040523d82523d6000602084013e6115d1565b606091505b5050905080610639576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610894565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905261063990849061188b565b60006116fc848484611997565b949350505050565b61146f611472565b3373ffffffffffffffffffffffffffffffffffffffff82160361178b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610894565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600081815260028301602052604081205480151580611825575061182584846119b4565b611501576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b657900006044820152606401610894565b60006118ed826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166119c09092919063ffffffff16565b805190915015610639578080602001905181019061190b919061285b565b610639576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610894565b600082815260028401602052604081208290556116fc84846119cf565b600061150183836119db565b60606116fc84846000856119f3565b60006115018383611b0c565b60008181526001830160205260408120541515611501565b606082471015611a85576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610894565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611aae9190612849565b60006040518083038185875af1925050503d8060008114611aeb576040519150601f19603f3d011682016040523d82523d6000602084013e611af0565b606091505b5091509150611b0187838387611b5b565b979650505050505050565b6000818152600183016020526040812054611b53575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556105f4565b5060006105f4565b60608315611bf1578251600003611bea5773ffffffffffffffffffffffffffffffffffffffff85163b611bea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610894565b50816116fc565b6116fc8383815115611c065781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108949190611d5c565b67ffffffffffffffff8116811461146f57600080fd5b60008083601f840112611c6257600080fd5b50813567ffffffffffffffff811115611c7a57600080fd5b602083019150836020828501011115611c9257600080fd5b9250929050565b600080600060408486031215611cae57600080fd5b8335611cb981611c3a565b9250602084013567ffffffffffffffff811115611cd557600080fd5b611ce186828701611c50565b9497909650939450505050565b60005b83811015611d09578181015183820152602001611cf1565b50506000910152565b60008151808452611d2a816020860160208601611cee565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006115016020830184611d12565b600060208284031215611d8157600080fd5b813561150181611c3a565b8315158152606060208201526000611da76060830185611d12565b8281036040840152611db98185611d12565b9695505050505050565b600060208284031215611dd557600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461146f57600080fd5b60008060408385031215611e1157600080fd5b8235611e1c81611ddc565b946020939093013593505050565b600080600060608486031215611e3f57600080fd5b8335611e4a81611ddc565b92506020840135611e5a81611ddc565b929592945050506040919091013590565b6000602080835283518184015280840151604067ffffffffffffffff821660408601526040860151915060a06060860152611ea960c0860183611d12565b915060608601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe080878503016080880152611ee58483611d12565b608089015188820390920160a089015281518082529186019450600092508501905b80831015611f46578451805173ffffffffffffffffffffffffffffffffffffffff16835286015186830152938501936001929092019190830190611f07565b50979650505050505050565b60008060408385031215611f6557600080fd5b823591506020830135611f7781611ddc565b809150509250929050565b60008083601f840112611f9457600080fd5b50813567ffffffffffffffff811115611fac57600080fd5b6020830191508360208260051b8501011115611c9257600080fd5b60008060008060408587031215611fdd57600080fd5b843567ffffffffffffffff80821115611ff557600080fd5b61200188838901611f82565b9096509450602087013591508082111561201a57600080fd5b5061202787828801611f82565b95989497509550505050565b60006020828403121561204557600080fd5b813567ffffffffffffffff81111561205c57600080fd5b820160a0818503121561150157600080fd5b60008060008060006060868803121561208657600080fd5b853561209181611c3a565b9450602086013567ffffffffffffffff808211156120ae57600080fd5b6120ba89838a01611c50565b909650945060408801359150808211156120d357600080fd5b506120e088828901611c50565b969995985093965092949392505050565b60006020828403121561210357600080fd5b813561150181611ddc565b8183823760009101908152919050565b600181811c9082168061213257607f821691505b60208210810361216b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc18336030181126121d457600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261221357600080fd5b83018035915067ffffffffffffffff82111561222e57600080fd5b602001915036819003821315611c9257600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261227857600080fd5b830160208101925035905067ffffffffffffffff81111561229857600080fd5b803603821315611c9257600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8183526000602080850194508260005b8581101561234557813561231381611ddc565b73ffffffffffffffffffffffffffffffffffffffff168752818301358388015260409687019690910190600101612300565b509495945050505050565b60208152813560208201526000602083013561236b81611c3a565b67ffffffffffffffff80821660408501526123896040860186612243565b925060a060608601526123a060c0860184836122a7565b9250506123b06060860186612243565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808786030160808801526123e68583856122a7565b9450608088013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe188360301831261241f57600080fd5b6020928801928301923591508382111561243857600080fd5b8160061b360383131561244a57600080fd5b8685030160a0870152611b018482846122f0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f821115610639576000816000526020600020601f850160051c810160208610156124b65750805b601f850160051c820191505b818110156111ec578281556001016124c2565b67ffffffffffffffff8311156124ed576124ed61245e565b612501836124fb835461211e565b8361248d565b6000601f841160018114612553576000851561251d5750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355610e3a565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b828110156125a25786850135825560209485019460019092019101612582565b50868210156125dd577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b600181901b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8216821461264c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b919050565b813561265c81611ddc565b73ffffffffffffffffffffffffffffffffffffffff81167fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825550602082013560018201555050565b680100000000000000008311156126c2576126c261245e565b805483825580841015612714576126d8816125ef565b6126e1856125ef565b6000848152602081209283019291909101905b82821015612710578082558060018301556002820191506126f4565b5050505b5060008181526020812083915b858110156111ec576127338383612651565b6040929092019160029190910190600101612721565b8135815560018101602083013561275f81611c3a565b67ffffffffffffffff8082167fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000084541617835561279f60408601866121de565b935091506127b18383600287016124d5565b6127be60608601866121de565b935091506127d08383600387016124d5565b608085013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe185360301831261280757600080fd5b91840191823591508082111561281c57600080fd5b506020820191508060061b360382131561283557600080fd5b6128438183600486016126a9565b50505050565b600082516121d4818460208701611cee565b60006020828403121561286d57600080fd5b8151801515811461150157600080fdfea164736f6c6343000818000a",
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

func (_CCIPReceiver *CCIPReceiverCaller) SChainConfigs(opts *bind.CallOpts, destChainSelector uint64) (SChainConfigs,

	error) {
	var out []interface{}
	err := _CCIPReceiver.contract.Call(opts, &out, "s_chainConfigs", destChainSelector)

	outstruct := new(SChainConfigs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Disabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Recipient = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.ExtraArgsBytes = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

func (_CCIPReceiver *CCIPReceiverSession) SChainConfigs(destChainSelector uint64) (SChainConfigs,

	error) {
	return _CCIPReceiver.Contract.SChainConfigs(&_CCIPReceiver.CallOpts, destChainSelector)
}

func (_CCIPReceiver *CCIPReceiverCallerSession) SChainConfigs(destChainSelector uint64) (SChainConfigs,

	error) {
	return _CCIPReceiver.Contract.SChainConfigs(&_CCIPReceiver.CallOpts, destChainSelector)
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

func (_CCIPReceiver *CCIPReceiverTransactor) AbandonFailedMessage(opts *bind.TransactOpts, messageId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "abandonFailedMessage", messageId, receiver)
}

func (_CCIPReceiver *CCIPReceiverSession) AbandonFailedMessage(messageId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.AbandonFailedMessage(&_CCIPReceiver.TransactOpts, messageId, receiver)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) AbandonFailedMessage(messageId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.AbandonFailedMessage(&_CCIPReceiver.TransactOpts, messageId, receiver)
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

func (_CCIPReceiver *CCIPReceiverTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "transferOwnership", to)
}

func (_CCIPReceiver *CCIPReceiverSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.TransferOwnership(&_CCIPReceiver.TransactOpts, to)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.TransferOwnership(&_CCIPReceiver.TransactOpts, to)
}

func (_CCIPReceiver *CCIPReceiverTransactor) UpdateApprovedSenders(opts *bind.TransactOpts, adds []CCIPClientBaseApprovedSenderUpdate, removes []CCIPClientBaseApprovedSenderUpdate) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "updateApprovedSenders", adds, removes)
}

func (_CCIPReceiver *CCIPReceiverSession) UpdateApprovedSenders(adds []CCIPClientBaseApprovedSenderUpdate, removes []CCIPClientBaseApprovedSenderUpdate) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.UpdateApprovedSenders(&_CCIPReceiver.TransactOpts, adds, removes)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) UpdateApprovedSenders(adds []CCIPClientBaseApprovedSenderUpdate, removes []CCIPClientBaseApprovedSenderUpdate) (*types.Transaction, error) {
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

func (_CCIPReceiver *CCIPReceiverTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPReceiver.contract.RawTransact(opts, nil)
}

func (_CCIPReceiver *CCIPReceiverSession) Receive() (*types.Transaction, error) {
	return _CCIPReceiver.Contract.Receive(&_CCIPReceiver.TransactOpts)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) Receive() (*types.Transaction, error) {
	return _CCIPReceiver.Contract.Receive(&_CCIPReceiver.TransactOpts)
}

type CCIPReceiverMessageAbandonedIterator struct {
	Event *CCIPReceiverMessageAbandoned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPReceiverMessageAbandonedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPReceiverMessageAbandoned)
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
		it.Event = new(CCIPReceiverMessageAbandoned)
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

func (it *CCIPReceiverMessageAbandonedIterator) Error() error {
	return it.fail
}

func (it *CCIPReceiverMessageAbandonedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPReceiverMessageAbandoned struct {
	MessageId     [32]byte
	TokenReceiver common.Address
	Raw           types.Log
}

func (_CCIPReceiver *CCIPReceiverFilterer) FilterMessageAbandoned(opts *bind.FilterOpts, messageId [][32]byte) (*CCIPReceiverMessageAbandonedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _CCIPReceiver.contract.FilterLogs(opts, "MessageAbandoned", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &CCIPReceiverMessageAbandonedIterator{contract: _CCIPReceiver.contract, event: "MessageAbandoned", logs: logs, sub: sub}, nil
}

func (_CCIPReceiver *CCIPReceiverFilterer) WatchMessageAbandoned(opts *bind.WatchOpts, sink chan<- *CCIPReceiverMessageAbandoned, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _CCIPReceiver.contract.WatchLogs(opts, "MessageAbandoned", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPReceiverMessageAbandoned)
				if err := _CCIPReceiver.contract.UnpackLog(event, "MessageAbandoned", log); err != nil {
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

func (_CCIPReceiver *CCIPReceiverFilterer) ParseMessageAbandoned(log types.Log) (*CCIPReceiverMessageAbandoned, error) {
	event := new(CCIPReceiverMessageAbandoned)
	if err := _CCIPReceiver.contract.UnpackLog(event, "MessageAbandoned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

type SChainConfigs struct {
	Disabled       bool
	Recipient      []byte
	ExtraArgsBytes []byte
}

func (_CCIPReceiver *CCIPReceiver) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CCIPReceiver.abi.Events["MessageAbandoned"].ID:
		return _CCIPReceiver.ParseMessageAbandoned(log)
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

func (CCIPReceiverMessageAbandoned) Topic() common.Hash {
	return common.HexToHash("0xd5038100bd3dc9631d3c3f4f61a3e53e9d466f40c47af9897292c7b35e32a957")
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

	SChainConfigs(opts *bind.CallOpts, destChainSelector uint64) (SChainConfigs,

		error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AbandonFailedMessage(opts *bind.TransactOpts, messageId [32]byte, receiver common.Address) (*types.Transaction, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	DisableChain(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error)

	EnableChain(opts *bind.TransactOpts, chainSelector uint64, recipient []byte, _extraArgsBytes []byte) (*types.Transaction, error)

	ProcessMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	RetryFailedMessage(opts *bind.TransactOpts, messageId [32]byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateApprovedSenders(opts *bind.TransactOpts, adds []CCIPClientBaseApprovedSenderUpdate, removes []CCIPClientBaseApprovedSenderUpdate) (*types.Transaction, error)

	WithdrawNativeToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)

	WithdrawTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterMessageAbandoned(opts *bind.FilterOpts, messageId [][32]byte) (*CCIPReceiverMessageAbandonedIterator, error)

	WatchMessageAbandoned(opts *bind.WatchOpts, sink chan<- *CCIPReceiverMessageAbandoned, messageId [][32]byte) (event.Subscription, error)

	ParseMessageAbandoned(log types.Log) (*CCIPReceiverMessageAbandoned, error)

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
