// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ccipSender

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

type CCIPBaseApprovedSenderUpdate struct {
	DestChainSelector uint64
	Sender            []byte
}

type CCIPBaseChainUpdate struct {
	ChainSelector  uint64
	Allowed        bool
	Recipient      []byte
	ExtraArgsBytes []byte
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

var CCIPSenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"ApprovedSenderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"ApprovedSenderRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRouter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"CCIPRouterModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraArgsBytes\",\"type\":\"bytes\"}],\"name\":\"ChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"removeChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawnByOwner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgsBytes\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPBase.ChainUpdate[]\",\"name\":\"chains\",\"type\":\"tuple[]\"}],\"name\":\"applyChainUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"senderAddr\",\"type\":\"bytes\"}],\"name\":\"isApprovedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"s_chainConfigs\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgsBytes\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPBase.ApprovedSenderUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPBase.ApprovedSenderUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"}],\"name\":\"updateApprovedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"updateRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200255b3803806200255b8339810160408190526200003491620001bc565b8033806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf8162000111565b5050506001600160a01b038116620000ea576040516342bcdf7f60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905550620001ee565b336001600160a01b038216036200016b5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001cf57600080fd5b81516001600160a01b0381168114620001e757600080fd5b9392505050565b61235d80620001fe6000396000f3fe6080604052600436106100c05760003560e01c80638da5cb5b11610074578063c851cc321161004e578063c851cc321461021d578063effde2401461023d578063f2fde38b1461025e57600080fd5b80638da5cb5b146101865780639fe74e26146101d2578063b0f479a1146101f257600080fd5b80635e35359e116100a55780635e35359e1461012f57806379ba5097146101515780638462a2b91461016657600080fd5b80630e958d6b146100cc57806335f170ef1461010157600080fd5b366100c757005b600080fd5b3480156100d857600080fd5b506100ec6100e736600461190d565b61027e565b60405190151581526020015b60405180910390f35b34801561010d57600080fd5b5061012161011c366004611990565b6102c9565b6040516100f8929190611a20565b34801561013b57600080fd5b5061014f61014a366004611a72565b6103f5565b005b34801561015d57600080fd5b5061014f6104ca565b34801561017257600080fd5b5061014f610181366004611afa565b6105cc565b34801561019257600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f8565b3480156101de57600080fd5b5061014f6101ed366004611b66565b61090d565b3480156101fe57600080fd5b5060025473ffffffffffffffffffffffffffffffffffffffff166101ad565b34801561022957600080fd5b5061014f610238366004611ba8565b610ac6565b61025061024b366004611cf8565b610b92565b6040519081526020016100f8565b34801561026a57600080fd5b5061014f610279366004611ba8565b61106b565b67ffffffffffffffff831660009081526003602052604080822090516002909101906102ad9085908590611e12565b9081526040519081900360200190205460ff1690509392505050565b6003602052600090815260409020805481906102e490611e22565b80601f016020809104026020016040519081016040528092919081815260200182805461031090611e22565b801561035d5780601f106103325761010080835404028352916020019161035d565b820191906000526020600020905b81548152906001019060200180831161034057829003601f168201915b50505050509080600101805461037290611e22565b80601f016020809104026020016040519081016040528092919081815260200182805461039e90611e22565b80156103eb5780601f106103c0576101008083540402835291602001916103eb565b820191906000526020600020905b8154815290600101906020018083116103ce57829003601f168201915b5050505050905082565b6103fd61107f565b73ffffffffffffffffffffffffffffffffffffffff831661043d5761043873ffffffffffffffffffffffffffffffffffffffff831682611102565b61045e565b61045e73ffffffffffffffffffffffffffffffffffffffff8416838361125c565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f6832d9be2410a86571981e1e60fd4c1f9ea2a1034d6102a2b7d6c5e480adf02e836040516104bd91815260200190565b60405180910390a3505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610550576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6105d461107f565b60005b8181101561076757600360008484848181106105f5576105f5611e75565b90506020028101906106079190611ea4565b610615906020810190611990565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020183838381811061064c5761064c611e75565b905060200281019061065e9190611ea4565b61066c906020810190611ee2565b60405161067a929190611e12565b90815260405190819003602001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558282828181106106c1576106c1611e75565b90506020028101906106d39190611ea4565b6106e1906020810190611ee2565b6040516106ef929190611e12565b604051809103902083838381811061070957610709611e75565b905060200281019061071b9190611ea4565b610729906020810190611990565b67ffffffffffffffff167f021290bab0d93f4d9a243bd430e45dd4bc8238451e9abbba6fab4463677dfce960405160405180910390a36001016105d7565b5060005b838110156109065760016003600087878581811061078b5761078b611e75565b905060200281019061079d9190611ea4565b6107ab906020810190611990565b67ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206002018686848181106107e2576107e2611e75565b90506020028101906107f49190611ea4565b610802906020810190611ee2565b604051610810929190611e12565b90815260405190819003602001902080549115157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0090921691909117905584848281811061086057610860611e75565b90506020028101906108729190611ea4565b610880906020810190611ee2565b60405161088e929190611e12565b60405180910390208585838181106108a8576108a8611e75565b90506020028101906108ba9190611ea4565b6108c8906020810190611990565b67ffffffffffffffff167f72d9f73bb7cb11065e15df29d61e803a0eba356d509a7025a6f51ebdea07f9e760405160405180910390a360010161076b565b5050505050565b61091561107f565b60005b81811015610ac157600083838381811061093457610934611e75565b90506020028101906109469190611f47565b61094f90611f89565b905080602001516109ba57805167ffffffffffffffff16600090815260036020526040812061097d916118a2565b805160405167ffffffffffffffff909116907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d85991690600090a2610ab8565b8060400151516000036109f9576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080820151825167ffffffffffffffff16600090815260036020529190912090610a249082612086565b506060810151815167ffffffffffffffff16600090815260036020526040902060010190610a529082612086565b508060400151604051610a6591906121a0565b6040518091039020816000015167ffffffffffffffff167f1ced5bcae649ed29cebfa0010298ad6794bf3822e8cb754a6eee5353a9a872128360600151604051610aaf91906121b2565b60405180910390a35b50600101610918565b505050565b610ace61107f565b73ffffffffffffffffffffffffffffffffffffffff8116610b1b576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f3672b589036f39ac008505b790fcb05d484d70b65680ec64c089a3c173fdc4c890600090a35050565b67ffffffffffffffff841660009081526003602052604081208054869190610bb990611e22565b9050600003610c00576040517fd79f2ea400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610547565b6040805160a08101825267ffffffffffffffff8816600090815260036020529182208054829190610c3090611e22565b80601f0160208091040260200160405190810160405280929190818152602001828054610c5c90611e22565b8015610ca95780601f10610c7e57610100808354040283529160200191610ca9565b820191906000526020600020905b815481529060010190602001808311610c8c57829003601f168201915b505050505081526020018681526020018781526020018573ffffffffffffffffffffffffffffffffffffffff168152602001600360008a67ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206001018054610d1090611e22565b80601f0160208091040260200160405190810160405280929190818152602001828054610d3c90611e22565b8015610d895780601f10610d5e57610100808354040283529160200191610d89565b820191906000526020600020905b815481529060010190602001808311610d6c57829003601f168201915b5050505050815250905060005b8651811015610e8f57610e063330898481518110610db657610db6611e75565b6020026020010151602001518a8581518110610dd457610dd4611e75565b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff16611330909392919063ffffffff16565b6002548751610e879173ffffffffffffffffffffffffffffffffffffffff1690899084908110610e3857610e38611e75565b602002602001015160200151898481518110610e5657610e56611e75565b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff166113949092919063ffffffff16565b600101610d96565b506002546040517f20487ded00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff16906320487ded90610ee9908b9086906004016121c5565b602060405180830381865afa158015610f06573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f2a91906122da565b905073ffffffffffffffffffffffffffffffffffffffff851615610f9157610f6a73ffffffffffffffffffffffffffffffffffffffff8616333084611330565b600254610f919073ffffffffffffffffffffffffffffffffffffffff878116911683611394565b60025473ffffffffffffffffffffffffffffffffffffffff908116906396f4e9f990871615610fc1576000610fc3565b825b8a856040518463ffffffff1660e01b8152600401610fe29291906121c5565b60206040518083038185885af1158015611000573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061102591906122da565b93507f54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a8460405161105891815260200190565b60405180910390a1505050949350505050565b61107361107f565b61107c81611492565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314611100576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610547565b565b8047101561116c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610547565b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d80600081146111c6576040519150601f19603f3d011682016040523d82523d6000602084013e6111cb565b606091505b5050905080610ac1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610547565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610ac19084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611587565b60405173ffffffffffffffffffffffffffffffffffffffff8085166024830152831660448201526064810182905261138e9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016112ae565b50505050565b6040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa15801561140b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061142f91906122da565b61143991906122f3565b60405173ffffffffffffffffffffffffffffffffffffffff851660248201526044810182905290915061138e9085907f095ea7b300000000000000000000000000000000000000000000000000000000906064016112ae565b3373ffffffffffffffffffffffffffffffffffffffff821603611511576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610547565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006115e9826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166116939092919063ffffffff16565b805190915015610ac157808060200190518101906116079190612333565b610ac1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610547565b60606116a284846000856116aa565b949350505050565b60608247101561173c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610547565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161176591906121a0565b60006040518083038185875af1925050503d80600081146117a2576040519150601f19603f3d011682016040523d82523d6000602084013e6117a7565b606091505b50915091506117b8878383876117c3565b979650505050505050565b606083156118595782516000036118525773ffffffffffffffffffffffffffffffffffffffff85163b611852576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610547565b50816116a2565b6116a2838381511561186e5781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054791906121b2565b5080546118ae90611e22565b6000825580601f106118be575050565b601f01602090049060005260206000209081019061107c91905b808211156118ec57600081556001016118d8565b5090565b803567ffffffffffffffff8116811461190857600080fd5b919050565b60008060006040848603121561192257600080fd5b61192b846118f0565b9250602084013567ffffffffffffffff8082111561194857600080fd5b818601915086601f83011261195c57600080fd5b81358181111561196b57600080fd5b87602082850101111561197d57600080fd5b6020830194508093505050509250925092565b6000602082840312156119a257600080fd5b6119ab826118f0565b9392505050565b60005b838110156119cd5781810151838201526020016119b5565b50506000910152565b600081518084526119ee8160208601602086016119b2565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b604081526000611a3360408301856119d6565b8281036020840152611a4581856119d6565b95945050505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461190857600080fd5b600080600060608486031215611a8757600080fd5b611a9084611a4e565b9250611a9e60208501611a4e565b9150604084013590509250925092565b60008083601f840112611ac057600080fd5b50813567ffffffffffffffff811115611ad857600080fd5b6020830191508360208260051b8501011115611af357600080fd5b9250929050565b60008060008060408587031215611b1057600080fd5b843567ffffffffffffffff80821115611b2857600080fd5b611b3488838901611aae565b90965094506020870135915080821115611b4d57600080fd5b50611b5a87828801611aae565b95989497509550505050565b60008060208385031215611b7957600080fd5b823567ffffffffffffffff811115611b9057600080fd5b611b9c85828601611aae565b90969095509350505050565b600060208284031215611bba57600080fd5b6119ab82611a4e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715611c1557611c15611bc3565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611c6257611c62611bc3565b604052919050565b600082601f830112611c7b57600080fd5b813567ffffffffffffffff811115611c9557611c95611bc3565b611cc660207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611c1b565b818152846020838601011115611cdb57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060808587031215611d0e57600080fd5b611d17856118f0565b935060208086013567ffffffffffffffff80821115611d3557600080fd5b818801915088601f830112611d4957600080fd5b813581811115611d5b57611d5b611bc3565b611d69848260051b01611c1b565b81815260069190911b8301840190848101908b831115611d8857600080fd5b938501935b82851015611dd2576040858d031215611da65760008081fd5b611dae611bf2565b611db786611a4e565b81528587013587820152825260409094019390850190611d8d565b975050506040880135925080831115611dea57600080fd5b5050611df887828801611c6a565b925050611e0760608601611a4e565b905092959194509250565b8183823760009101908152919050565b600181811c90821680611e3657607f821691505b602082108103611e6f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc1833603018112611ed857600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611f1757600080fd5b83018035915067ffffffffffffffff821115611f3257600080fd5b602001915036819003821315611af357600080fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112611ed857600080fd5b801515811461107c57600080fd5b600060808236031215611f9b57600080fd5b6040516080810167ffffffffffffffff8282108183111715611fbf57611fbf611bc3565b81604052611fcc856118f0565b835260208501359150611fde82611f7b565b8160208401526040850135915080821115611ff857600080fd5b61200436838701611c6a565b6040840152606085013591508082111561201d57600080fd5b5061202a36828601611c6a565b60608301525092915050565b601f821115610ac1576000816000526020600020601f850160051c8101602086101561205f5750805b601f850160051c820191505b8181101561207e5782815560010161206b565b505050505050565b815167ffffffffffffffff8111156120a0576120a0611bc3565b6120b4816120ae8454611e22565b84612036565b602080601f83116001811461210757600084156120d15750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b17855561207e565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561215457888601518255948401946001909101908401612135565b508582101561219057878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b60008251611ed88184602087016119b2565b6020815260006119ab60208301846119d6565b6000604067ffffffffffffffff851683526020604081850152845160a060408601526121f460e08601826119d6565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08087840301606088015261222f83836119d6565b6040890151888203830160808a01528051808352908601945060009350908501905b80841015612290578451805173ffffffffffffffffffffffffffffffffffffffff16835286015186830152938501936001939093019290860190612251565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a015295506122cc81876119d6565b9a9950505050505050505050565b6000602082840312156122ec57600080fd5b5051919050565b8082018082111561232d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b60006020828403121561234557600080fd5b81516119ab81611f7b56fea164736f6c6343000818000a",
}

var CCIPSenderABI = CCIPSenderMetaData.ABI

var CCIPSenderBin = CCIPSenderMetaData.Bin

func DeployCCIPSender(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address) (common.Address, *types.Transaction, *CCIPSender, error) {
	parsed, err := CCIPSenderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCIPSenderBin), backend, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCIPSender{address: address, abi: *parsed, CCIPSenderCaller: CCIPSenderCaller{contract: contract}, CCIPSenderTransactor: CCIPSenderTransactor{contract: contract}, CCIPSenderFilterer: CCIPSenderFilterer{contract: contract}}, nil
}

type CCIPSender struct {
	address common.Address
	abi     abi.ABI
	CCIPSenderCaller
	CCIPSenderTransactor
	CCIPSenderFilterer
}

type CCIPSenderCaller struct {
	contract *bind.BoundContract
}

type CCIPSenderTransactor struct {
	contract *bind.BoundContract
}

type CCIPSenderFilterer struct {
	contract *bind.BoundContract
}

type CCIPSenderSession struct {
	Contract     *CCIPSender
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CCIPSenderCallerSession struct {
	Contract *CCIPSenderCaller
	CallOpts bind.CallOpts
}

type CCIPSenderTransactorSession struct {
	Contract     *CCIPSenderTransactor
	TransactOpts bind.TransactOpts
}

type CCIPSenderRaw struct {
	Contract *CCIPSender
}

type CCIPSenderCallerRaw struct {
	Contract *CCIPSenderCaller
}

type CCIPSenderTransactorRaw struct {
	Contract *CCIPSenderTransactor
}

func NewCCIPSender(address common.Address, backend bind.ContractBackend) (*CCIPSender, error) {
	abi, err := abi.JSON(strings.NewReader(CCIPSenderABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCCIPSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCIPSender{address: address, abi: abi, CCIPSenderCaller: CCIPSenderCaller{contract: contract}, CCIPSenderTransactor: CCIPSenderTransactor{contract: contract}, CCIPSenderFilterer: CCIPSenderFilterer{contract: contract}}, nil
}

func NewCCIPSenderCaller(address common.Address, caller bind.ContractCaller) (*CCIPSenderCaller, error) {
	contract, err := bindCCIPSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCIPSenderCaller{contract: contract}, nil
}

func NewCCIPSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*CCIPSenderTransactor, error) {
	contract, err := bindCCIPSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCIPSenderTransactor{contract: contract}, nil
}

func NewCCIPSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*CCIPSenderFilterer, error) {
	contract, err := bindCCIPSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCIPSenderFilterer{contract: contract}, nil
}

func bindCCIPSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCIPSenderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CCIPSender *CCIPSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCIPSender.Contract.CCIPSenderCaller.contract.Call(opts, result, method, params...)
}

func (_CCIPSender *CCIPSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPSender.Contract.CCIPSenderTransactor.contract.Transfer(opts)
}

func (_CCIPSender *CCIPSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCIPSender.Contract.CCIPSenderTransactor.contract.Transact(opts, method, params...)
}

func (_CCIPSender *CCIPSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCIPSender.Contract.contract.Call(opts, result, method, params...)
}

func (_CCIPSender *CCIPSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPSender.Contract.contract.Transfer(opts)
}

func (_CCIPSender *CCIPSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCIPSender.Contract.contract.Transact(opts, method, params...)
}

func (_CCIPSender *CCIPSenderCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCIPSender.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCIPSender *CCIPSenderSession) GetRouter() (common.Address, error) {
	return _CCIPSender.Contract.GetRouter(&_CCIPSender.CallOpts)
}

func (_CCIPSender *CCIPSenderCallerSession) GetRouter() (common.Address, error) {
	return _CCIPSender.Contract.GetRouter(&_CCIPSender.CallOpts)
}

func (_CCIPSender *CCIPSenderCaller) IsApprovedSender(opts *bind.CallOpts, sourceChainSelector uint64, senderAddr []byte) (bool, error) {
	var out []interface{}
	err := _CCIPSender.contract.Call(opts, &out, "isApprovedSender", sourceChainSelector, senderAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CCIPSender *CCIPSenderSession) IsApprovedSender(sourceChainSelector uint64, senderAddr []byte) (bool, error) {
	return _CCIPSender.Contract.IsApprovedSender(&_CCIPSender.CallOpts, sourceChainSelector, senderAddr)
}

func (_CCIPSender *CCIPSenderCallerSession) IsApprovedSender(sourceChainSelector uint64, senderAddr []byte) (bool, error) {
	return _CCIPSender.Contract.IsApprovedSender(&_CCIPSender.CallOpts, sourceChainSelector, senderAddr)
}

func (_CCIPSender *CCIPSenderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCIPSender.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCIPSender *CCIPSenderSession) Owner() (common.Address, error) {
	return _CCIPSender.Contract.Owner(&_CCIPSender.CallOpts)
}

func (_CCIPSender *CCIPSenderCallerSession) Owner() (common.Address, error) {
	return _CCIPSender.Contract.Owner(&_CCIPSender.CallOpts)
}

func (_CCIPSender *CCIPSenderCaller) SChainConfigs(opts *bind.CallOpts, destChainSelector uint64) (SChainConfigs,

	error) {
	var out []interface{}
	err := _CCIPSender.contract.Call(opts, &out, "s_chainConfigs", destChainSelector)

	outstruct := new(SChainConfigs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recipient = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ExtraArgsBytes = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

func (_CCIPSender *CCIPSenderSession) SChainConfigs(destChainSelector uint64) (SChainConfigs,

	error) {
	return _CCIPSender.Contract.SChainConfigs(&_CCIPSender.CallOpts, destChainSelector)
}

func (_CCIPSender *CCIPSenderCallerSession) SChainConfigs(destChainSelector uint64) (SChainConfigs,

	error) {
	return _CCIPSender.Contract.SChainConfigs(&_CCIPSender.CallOpts, destChainSelector)
}

func (_CCIPSender *CCIPSenderTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPSender.contract.Transact(opts, "acceptOwnership")
}

func (_CCIPSender *CCIPSenderSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCIPSender.Contract.AcceptOwnership(&_CCIPSender.TransactOpts)
}

func (_CCIPSender *CCIPSenderTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCIPSender.Contract.AcceptOwnership(&_CCIPSender.TransactOpts)
}

func (_CCIPSender *CCIPSenderTransactor) ApplyChainUpdates(opts *bind.TransactOpts, chains []CCIPBaseChainUpdate) (*types.Transaction, error) {
	return _CCIPSender.contract.Transact(opts, "applyChainUpdates", chains)
}

func (_CCIPSender *CCIPSenderSession) ApplyChainUpdates(chains []CCIPBaseChainUpdate) (*types.Transaction, error) {
	return _CCIPSender.Contract.ApplyChainUpdates(&_CCIPSender.TransactOpts, chains)
}

func (_CCIPSender *CCIPSenderTransactorSession) ApplyChainUpdates(chains []CCIPBaseChainUpdate) (*types.Transaction, error) {
	return _CCIPSender.Contract.ApplyChainUpdates(&_CCIPSender.TransactOpts, chains)
}

func (_CCIPSender *CCIPSenderTransactor) CcipSend(opts *bind.TransactOpts, destChainSelector uint64, tokenAmounts []ClientEVMTokenAmount, data []byte, feeToken common.Address) (*types.Transaction, error) {
	return _CCIPSender.contract.Transact(opts, "ccipSend", destChainSelector, tokenAmounts, data, feeToken)
}

func (_CCIPSender *CCIPSenderSession) CcipSend(destChainSelector uint64, tokenAmounts []ClientEVMTokenAmount, data []byte, feeToken common.Address) (*types.Transaction, error) {
	return _CCIPSender.Contract.CcipSend(&_CCIPSender.TransactOpts, destChainSelector, tokenAmounts, data, feeToken)
}

func (_CCIPSender *CCIPSenderTransactorSession) CcipSend(destChainSelector uint64, tokenAmounts []ClientEVMTokenAmount, data []byte, feeToken common.Address) (*types.Transaction, error) {
	return _CCIPSender.Contract.CcipSend(&_CCIPSender.TransactOpts, destChainSelector, tokenAmounts, data, feeToken)
}

func (_CCIPSender *CCIPSenderTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCIPSender.contract.Transact(opts, "transferOwnership", to)
}

func (_CCIPSender *CCIPSenderSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCIPSender.Contract.TransferOwnership(&_CCIPSender.TransactOpts, to)
}

func (_CCIPSender *CCIPSenderTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCIPSender.Contract.TransferOwnership(&_CCIPSender.TransactOpts, to)
}

func (_CCIPSender *CCIPSenderTransactor) UpdateApprovedSenders(opts *bind.TransactOpts, adds []CCIPBaseApprovedSenderUpdate, removes []CCIPBaseApprovedSenderUpdate) (*types.Transaction, error) {
	return _CCIPSender.contract.Transact(opts, "updateApprovedSenders", adds, removes)
}

func (_CCIPSender *CCIPSenderSession) UpdateApprovedSenders(adds []CCIPBaseApprovedSenderUpdate, removes []CCIPBaseApprovedSenderUpdate) (*types.Transaction, error) {
	return _CCIPSender.Contract.UpdateApprovedSenders(&_CCIPSender.TransactOpts, adds, removes)
}

func (_CCIPSender *CCIPSenderTransactorSession) UpdateApprovedSenders(adds []CCIPBaseApprovedSenderUpdate, removes []CCIPBaseApprovedSenderUpdate) (*types.Transaction, error) {
	return _CCIPSender.Contract.UpdateApprovedSenders(&_CCIPSender.TransactOpts, adds, removes)
}

func (_CCIPSender *CCIPSenderTransactor) UpdateRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _CCIPSender.contract.Transact(opts, "updateRouter", newRouter)
}

func (_CCIPSender *CCIPSenderSession) UpdateRouter(newRouter common.Address) (*types.Transaction, error) {
	return _CCIPSender.Contract.UpdateRouter(&_CCIPSender.TransactOpts, newRouter)
}

func (_CCIPSender *CCIPSenderTransactorSession) UpdateRouter(newRouter common.Address) (*types.Transaction, error) {
	return _CCIPSender.Contract.UpdateRouter(&_CCIPSender.TransactOpts, newRouter)
}

func (_CCIPSender *CCIPSenderTransactor) WithdrawTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCIPSender.contract.Transact(opts, "withdrawTokens", token, to, amount)
}

func (_CCIPSender *CCIPSenderSession) WithdrawTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCIPSender.Contract.WithdrawTokens(&_CCIPSender.TransactOpts, token, to, amount)
}

func (_CCIPSender *CCIPSenderTransactorSession) WithdrawTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCIPSender.Contract.WithdrawTokens(&_CCIPSender.TransactOpts, token, to, amount)
}

func (_CCIPSender *CCIPSenderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPSender.contract.RawTransact(opts, nil)
}

func (_CCIPSender *CCIPSenderSession) Receive() (*types.Transaction, error) {
	return _CCIPSender.Contract.Receive(&_CCIPSender.TransactOpts)
}

func (_CCIPSender *CCIPSenderTransactorSession) Receive() (*types.Transaction, error) {
	return _CCIPSender.Contract.Receive(&_CCIPSender.TransactOpts)
}

type CCIPSenderApprovedSenderAddedIterator struct {
	Event *CCIPSenderApprovedSenderAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPSenderApprovedSenderAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPSenderApprovedSenderAdded)
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
		it.Event = new(CCIPSenderApprovedSenderAdded)
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

func (it *CCIPSenderApprovedSenderAddedIterator) Error() error {
	return it.fail
}

func (it *CCIPSenderApprovedSenderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPSenderApprovedSenderAdded struct {
	DestChainSelector uint64
	Recipient         common.Hash
	Raw               types.Log
}

func (_CCIPSender *CCIPSenderFilterer) FilterApprovedSenderAdded(opts *bind.FilterOpts, destChainSelector []uint64, recipient [][]byte) (*CCIPSenderApprovedSenderAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CCIPSender.contract.FilterLogs(opts, "ApprovedSenderAdded", destChainSelectorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CCIPSenderApprovedSenderAddedIterator{contract: _CCIPSender.contract, event: "ApprovedSenderAdded", logs: logs, sub: sub}, nil
}

func (_CCIPSender *CCIPSenderFilterer) WatchApprovedSenderAdded(opts *bind.WatchOpts, sink chan<- *CCIPSenderApprovedSenderAdded, destChainSelector []uint64, recipient [][]byte) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CCIPSender.contract.WatchLogs(opts, "ApprovedSenderAdded", destChainSelectorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPSenderApprovedSenderAdded)
				if err := _CCIPSender.contract.UnpackLog(event, "ApprovedSenderAdded", log); err != nil {
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

func (_CCIPSender *CCIPSenderFilterer) ParseApprovedSenderAdded(log types.Log) (*CCIPSenderApprovedSenderAdded, error) {
	event := new(CCIPSenderApprovedSenderAdded)
	if err := _CCIPSender.contract.UnpackLog(event, "ApprovedSenderAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPSenderApprovedSenderRemovedIterator struct {
	Event *CCIPSenderApprovedSenderRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPSenderApprovedSenderRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPSenderApprovedSenderRemoved)
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
		it.Event = new(CCIPSenderApprovedSenderRemoved)
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

func (it *CCIPSenderApprovedSenderRemovedIterator) Error() error {
	return it.fail
}

func (it *CCIPSenderApprovedSenderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPSenderApprovedSenderRemoved struct {
	DestChainSelector uint64
	Recipient         common.Hash
	Raw               types.Log
}

func (_CCIPSender *CCIPSenderFilterer) FilterApprovedSenderRemoved(opts *bind.FilterOpts, destChainSelector []uint64, recipient [][]byte) (*CCIPSenderApprovedSenderRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CCIPSender.contract.FilterLogs(opts, "ApprovedSenderRemoved", destChainSelectorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CCIPSenderApprovedSenderRemovedIterator{contract: _CCIPSender.contract, event: "ApprovedSenderRemoved", logs: logs, sub: sub}, nil
}

func (_CCIPSender *CCIPSenderFilterer) WatchApprovedSenderRemoved(opts *bind.WatchOpts, sink chan<- *CCIPSenderApprovedSenderRemoved, destChainSelector []uint64, recipient [][]byte) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CCIPSender.contract.WatchLogs(opts, "ApprovedSenderRemoved", destChainSelectorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPSenderApprovedSenderRemoved)
				if err := _CCIPSender.contract.UnpackLog(event, "ApprovedSenderRemoved", log); err != nil {
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

func (_CCIPSender *CCIPSenderFilterer) ParseApprovedSenderRemoved(log types.Log) (*CCIPSenderApprovedSenderRemoved, error) {
	event := new(CCIPSenderApprovedSenderRemoved)
	if err := _CCIPSender.contract.UnpackLog(event, "ApprovedSenderRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPSenderCCIPRouterModifiedIterator struct {
	Event *CCIPSenderCCIPRouterModified

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPSenderCCIPRouterModifiedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPSenderCCIPRouterModified)
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
		it.Event = new(CCIPSenderCCIPRouterModified)
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

func (it *CCIPSenderCCIPRouterModifiedIterator) Error() error {
	return it.fail
}

func (it *CCIPSenderCCIPRouterModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPSenderCCIPRouterModified struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_CCIPSender *CCIPSenderFilterer) FilterCCIPRouterModified(opts *bind.FilterOpts, oldRouter []common.Address, newRouter []common.Address) (*CCIPSenderCCIPRouterModifiedIterator, error) {

	var oldRouterRule []interface{}
	for _, oldRouterItem := range oldRouter {
		oldRouterRule = append(oldRouterRule, oldRouterItem)
	}
	var newRouterRule []interface{}
	for _, newRouterItem := range newRouter {
		newRouterRule = append(newRouterRule, newRouterItem)
	}

	logs, sub, err := _CCIPSender.contract.FilterLogs(opts, "CCIPRouterModified", oldRouterRule, newRouterRule)
	if err != nil {
		return nil, err
	}
	return &CCIPSenderCCIPRouterModifiedIterator{contract: _CCIPSender.contract, event: "CCIPRouterModified", logs: logs, sub: sub}, nil
}

func (_CCIPSender *CCIPSenderFilterer) WatchCCIPRouterModified(opts *bind.WatchOpts, sink chan<- *CCIPSenderCCIPRouterModified, oldRouter []common.Address, newRouter []common.Address) (event.Subscription, error) {

	var oldRouterRule []interface{}
	for _, oldRouterItem := range oldRouter {
		oldRouterRule = append(oldRouterRule, oldRouterItem)
	}
	var newRouterRule []interface{}
	for _, newRouterItem := range newRouter {
		newRouterRule = append(newRouterRule, newRouterItem)
	}

	logs, sub, err := _CCIPSender.contract.WatchLogs(opts, "CCIPRouterModified", oldRouterRule, newRouterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPSenderCCIPRouterModified)
				if err := _CCIPSender.contract.UnpackLog(event, "CCIPRouterModified", log); err != nil {
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

func (_CCIPSender *CCIPSenderFilterer) ParseCCIPRouterModified(log types.Log) (*CCIPSenderCCIPRouterModified, error) {
	event := new(CCIPSenderCCIPRouterModified)
	if err := _CCIPSender.contract.UnpackLog(event, "CCIPRouterModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPSenderChainAddedIterator struct {
	Event *CCIPSenderChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPSenderChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPSenderChainAdded)
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
		it.Event = new(CCIPSenderChainAdded)
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

func (it *CCIPSenderChainAddedIterator) Error() error {
	return it.fail
}

func (it *CCIPSenderChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPSenderChainAdded struct {
	RemoteChainSelector uint64
	Recipient           common.Hash
	ExtraArgsBytes      []byte
	Raw                 types.Log
}

func (_CCIPSender *CCIPSenderFilterer) FilterChainAdded(opts *bind.FilterOpts, remoteChainSelector []uint64, recipient [][]byte) (*CCIPSenderChainAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CCIPSender.contract.FilterLogs(opts, "ChainAdded", remoteChainSelectorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CCIPSenderChainAddedIterator{contract: _CCIPSender.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_CCIPSender *CCIPSenderFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *CCIPSenderChainAdded, remoteChainSelector []uint64, recipient [][]byte) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CCIPSender.contract.WatchLogs(opts, "ChainAdded", remoteChainSelectorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPSenderChainAdded)
				if err := _CCIPSender.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_CCIPSender *CCIPSenderFilterer) ParseChainAdded(log types.Log) (*CCIPSenderChainAdded, error) {
	event := new(CCIPSenderChainAdded)
	if err := _CCIPSender.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPSenderChainRemovedIterator struct {
	Event *CCIPSenderChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPSenderChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPSenderChainRemoved)
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
		it.Event = new(CCIPSenderChainRemoved)
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

func (it *CCIPSenderChainRemovedIterator) Error() error {
	return it.fail
}

func (it *CCIPSenderChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPSenderChainRemoved struct {
	RemoveChainSelector uint64
	Raw                 types.Log
}

func (_CCIPSender *CCIPSenderFilterer) FilterChainRemoved(opts *bind.FilterOpts, removeChainSelector []uint64) (*CCIPSenderChainRemovedIterator, error) {

	var removeChainSelectorRule []interface{}
	for _, removeChainSelectorItem := range removeChainSelector {
		removeChainSelectorRule = append(removeChainSelectorRule, removeChainSelectorItem)
	}

	logs, sub, err := _CCIPSender.contract.FilterLogs(opts, "ChainRemoved", removeChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCIPSenderChainRemovedIterator{contract: _CCIPSender.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_CCIPSender *CCIPSenderFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *CCIPSenderChainRemoved, removeChainSelector []uint64) (event.Subscription, error) {

	var removeChainSelectorRule []interface{}
	for _, removeChainSelectorItem := range removeChainSelector {
		removeChainSelectorRule = append(removeChainSelectorRule, removeChainSelectorItem)
	}

	logs, sub, err := _CCIPSender.contract.WatchLogs(opts, "ChainRemoved", removeChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPSenderChainRemoved)
				if err := _CCIPSender.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_CCIPSender *CCIPSenderFilterer) ParseChainRemoved(log types.Log) (*CCIPSenderChainRemoved, error) {
	event := new(CCIPSenderChainRemoved)
	if err := _CCIPSender.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPSenderMessageSentIterator struct {
	Event *CCIPSenderMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPSenderMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPSenderMessageSent)
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
		it.Event = new(CCIPSenderMessageSent)
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

func (it *CCIPSenderMessageSentIterator) Error() error {
	return it.fail
}

func (it *CCIPSenderMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPSenderMessageSent struct {
	MessageId [32]byte
	Raw       types.Log
}

func (_CCIPSender *CCIPSenderFilterer) FilterMessageSent(opts *bind.FilterOpts) (*CCIPSenderMessageSentIterator, error) {

	logs, sub, err := _CCIPSender.contract.FilterLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return &CCIPSenderMessageSentIterator{contract: _CCIPSender.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

func (_CCIPSender *CCIPSenderFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *CCIPSenderMessageSent) (event.Subscription, error) {

	logs, sub, err := _CCIPSender.contract.WatchLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPSenderMessageSent)
				if err := _CCIPSender.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

func (_CCIPSender *CCIPSenderFilterer) ParseMessageSent(log types.Log) (*CCIPSenderMessageSent, error) {
	event := new(CCIPSenderMessageSent)
	if err := _CCIPSender.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPSenderOwnershipTransferRequestedIterator struct {
	Event *CCIPSenderOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPSenderOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPSenderOwnershipTransferRequested)
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
		it.Event = new(CCIPSenderOwnershipTransferRequested)
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

func (it *CCIPSenderOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CCIPSenderOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPSenderOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCIPSender *CCIPSenderFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCIPSenderOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCIPSender.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCIPSenderOwnershipTransferRequestedIterator{contract: _CCIPSender.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CCIPSender *CCIPSenderFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCIPSenderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCIPSender.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPSenderOwnershipTransferRequested)
				if err := _CCIPSender.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CCIPSender *CCIPSenderFilterer) ParseOwnershipTransferRequested(log types.Log) (*CCIPSenderOwnershipTransferRequested, error) {
	event := new(CCIPSenderOwnershipTransferRequested)
	if err := _CCIPSender.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPSenderOwnershipTransferredIterator struct {
	Event *CCIPSenderOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPSenderOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPSenderOwnershipTransferred)
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
		it.Event = new(CCIPSenderOwnershipTransferred)
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

func (it *CCIPSenderOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CCIPSenderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPSenderOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCIPSender *CCIPSenderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCIPSenderOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCIPSender.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCIPSenderOwnershipTransferredIterator{contract: _CCIPSender.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CCIPSender *CCIPSenderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCIPSenderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCIPSender.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPSenderOwnershipTransferred)
				if err := _CCIPSender.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CCIPSender *CCIPSenderFilterer) ParseOwnershipTransferred(log types.Log) (*CCIPSenderOwnershipTransferred, error) {
	event := new(CCIPSenderOwnershipTransferred)
	if err := _CCIPSender.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPSenderTokensWithdrawnByOwnerIterator struct {
	Event *CCIPSenderTokensWithdrawnByOwner

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPSenderTokensWithdrawnByOwnerIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPSenderTokensWithdrawnByOwner)
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
		it.Event = new(CCIPSenderTokensWithdrawnByOwner)
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

func (it *CCIPSenderTokensWithdrawnByOwnerIterator) Error() error {
	return it.fail
}

func (it *CCIPSenderTokensWithdrawnByOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPSenderTokensWithdrawnByOwner struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_CCIPSender *CCIPSenderFilterer) FilterTokensWithdrawnByOwner(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*CCIPSenderTokensWithdrawnByOwnerIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCIPSender.contract.FilterLogs(opts, "TokensWithdrawnByOwner", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCIPSenderTokensWithdrawnByOwnerIterator{contract: _CCIPSender.contract, event: "TokensWithdrawnByOwner", logs: logs, sub: sub}, nil
}

func (_CCIPSender *CCIPSenderFilterer) WatchTokensWithdrawnByOwner(opts *bind.WatchOpts, sink chan<- *CCIPSenderTokensWithdrawnByOwner, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCIPSender.contract.WatchLogs(opts, "TokensWithdrawnByOwner", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPSenderTokensWithdrawnByOwner)
				if err := _CCIPSender.contract.UnpackLog(event, "TokensWithdrawnByOwner", log); err != nil {
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

func (_CCIPSender *CCIPSenderFilterer) ParseTokensWithdrawnByOwner(log types.Log) (*CCIPSenderTokensWithdrawnByOwner, error) {
	event := new(CCIPSenderTokensWithdrawnByOwner)
	if err := _CCIPSender.contract.UnpackLog(event, "TokensWithdrawnByOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SChainConfigs struct {
	Recipient      []byte
	ExtraArgsBytes []byte
}

func (_CCIPSender *CCIPSender) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CCIPSender.abi.Events["ApprovedSenderAdded"].ID:
		return _CCIPSender.ParseApprovedSenderAdded(log)
	case _CCIPSender.abi.Events["ApprovedSenderRemoved"].ID:
		return _CCIPSender.ParseApprovedSenderRemoved(log)
	case _CCIPSender.abi.Events["CCIPRouterModified"].ID:
		return _CCIPSender.ParseCCIPRouterModified(log)
	case _CCIPSender.abi.Events["ChainAdded"].ID:
		return _CCIPSender.ParseChainAdded(log)
	case _CCIPSender.abi.Events["ChainRemoved"].ID:
		return _CCIPSender.ParseChainRemoved(log)
	case _CCIPSender.abi.Events["MessageSent"].ID:
		return _CCIPSender.ParseMessageSent(log)
	case _CCIPSender.abi.Events["OwnershipTransferRequested"].ID:
		return _CCIPSender.ParseOwnershipTransferRequested(log)
	case _CCIPSender.abi.Events["OwnershipTransferred"].ID:
		return _CCIPSender.ParseOwnershipTransferred(log)
	case _CCIPSender.abi.Events["TokensWithdrawnByOwner"].ID:
		return _CCIPSender.ParseTokensWithdrawnByOwner(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CCIPSenderApprovedSenderAdded) Topic() common.Hash {
	return common.HexToHash("0x72d9f73bb7cb11065e15df29d61e803a0eba356d509a7025a6f51ebdea07f9e7")
}

func (CCIPSenderApprovedSenderRemoved) Topic() common.Hash {
	return common.HexToHash("0x021290bab0d93f4d9a243bd430e45dd4bc8238451e9abbba6fab4463677dfce9")
}

func (CCIPSenderCCIPRouterModified) Topic() common.Hash {
	return common.HexToHash("0x3672b589036f39ac008505b790fcb05d484d70b65680ec64c089a3c173fdc4c8")
}

func (CCIPSenderChainAdded) Topic() common.Hash {
	return common.HexToHash("0x1ced5bcae649ed29cebfa0010298ad6794bf3822e8cb754a6eee5353a9a87212")
}

func (CCIPSenderChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (CCIPSenderMessageSent) Topic() common.Hash {
	return common.HexToHash("0x54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a")
}

func (CCIPSenderOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCIPSenderOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CCIPSenderTokensWithdrawnByOwner) Topic() common.Hash {
	return common.HexToHash("0x6832d9be2410a86571981e1e60fd4c1f9ea2a1034d6102a2b7d6c5e480adf02e")
}

func (_CCIPSender *CCIPSender) Address() common.Address {
	return _CCIPSender.address
}

type CCIPSenderInterface interface {
	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IsApprovedSender(opts *bind.CallOpts, sourceChainSelector uint64, senderAddr []byte) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SChainConfigs(opts *bind.CallOpts, destChainSelector uint64) (SChainConfigs,

		error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, chains []CCIPBaseChainUpdate) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destChainSelector uint64, tokenAmounts []ClientEVMTokenAmount, data []byte, feeToken common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateApprovedSenders(opts *bind.TransactOpts, adds []CCIPBaseApprovedSenderUpdate, removes []CCIPBaseApprovedSenderUpdate) (*types.Transaction, error)

	UpdateRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	WithdrawTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterApprovedSenderAdded(opts *bind.FilterOpts, destChainSelector []uint64, recipient [][]byte) (*CCIPSenderApprovedSenderAddedIterator, error)

	WatchApprovedSenderAdded(opts *bind.WatchOpts, sink chan<- *CCIPSenderApprovedSenderAdded, destChainSelector []uint64, recipient [][]byte) (event.Subscription, error)

	ParseApprovedSenderAdded(log types.Log) (*CCIPSenderApprovedSenderAdded, error)

	FilterApprovedSenderRemoved(opts *bind.FilterOpts, destChainSelector []uint64, recipient [][]byte) (*CCIPSenderApprovedSenderRemovedIterator, error)

	WatchApprovedSenderRemoved(opts *bind.WatchOpts, sink chan<- *CCIPSenderApprovedSenderRemoved, destChainSelector []uint64, recipient [][]byte) (event.Subscription, error)

	ParseApprovedSenderRemoved(log types.Log) (*CCIPSenderApprovedSenderRemoved, error)

	FilterCCIPRouterModified(opts *bind.FilterOpts, oldRouter []common.Address, newRouter []common.Address) (*CCIPSenderCCIPRouterModifiedIterator, error)

	WatchCCIPRouterModified(opts *bind.WatchOpts, sink chan<- *CCIPSenderCCIPRouterModified, oldRouter []common.Address, newRouter []common.Address) (event.Subscription, error)

	ParseCCIPRouterModified(log types.Log) (*CCIPSenderCCIPRouterModified, error)

	FilterChainAdded(opts *bind.FilterOpts, remoteChainSelector []uint64, recipient [][]byte) (*CCIPSenderChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *CCIPSenderChainAdded, remoteChainSelector []uint64, recipient [][]byte) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*CCIPSenderChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts, removeChainSelector []uint64) (*CCIPSenderChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *CCIPSenderChainRemoved, removeChainSelector []uint64) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*CCIPSenderChainRemoved, error)

	FilterMessageSent(opts *bind.FilterOpts) (*CCIPSenderMessageSentIterator, error)

	WatchMessageSent(opts *bind.WatchOpts, sink chan<- *CCIPSenderMessageSent) (event.Subscription, error)

	ParseMessageSent(log types.Log) (*CCIPSenderMessageSent, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCIPSenderOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCIPSenderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCIPSenderOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCIPSenderOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCIPSenderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCIPSenderOwnershipTransferred, error)

	FilterTokensWithdrawnByOwner(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*CCIPSenderTokensWithdrawnByOwnerIterator, error)

	WatchTokensWithdrawnByOwner(opts *bind.WatchOpts, sink chan<- *CCIPSenderTokensWithdrawnByOwner, token []common.Address, to []common.Address) (event.Subscription, error)

	ParseTokensWithdrawnByOwner(log types.Log) (*CCIPSenderTokensWithdrawnByOwner, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
