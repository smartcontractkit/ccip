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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientNativeFeeTokenAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBeingApproved\",\"type\":\"bool\"}],\"name\":\"ApprovedSenderModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRouter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"CCIPRouterModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraArgsBytes\",\"type\":\"bytes\"}],\"name\":\"ChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"removeChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawnByOwner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgsBytes\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPBase.ChainUpdate[]\",\"name\":\"chains\",\"type\":\"tuple[]\"}],\"name\":\"applyChainUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"senderAddr\",\"type\":\"bytes\"}],\"name\":\"isApprovedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"modifyRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"s_chainConfigs\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgsBytes\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPBase.ApprovedSenderUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPBase.ApprovedSenderUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"}],\"name\":\"updateApprovedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNativeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002759380380620027598339810160408190526200003491620001bc565b8033806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf8162000111565b5050506001600160a01b038116620000ea576040516342bcdf7f60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905550620001ee565b336001600160a01b038216036200016b5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001cf57600080fd5b81516001600160a01b0381168114620001e757600080fd5b9392505050565b61255b80620001fe6000396000f3fe6080604052600436106100cb5760003560e01c80638462a2b911610074578063b0f479a11161004e578063b0f479a11461023d578063effde24014610268578063f2fde38b1461028957600080fd5b80638462a2b9146101b15780638da5cb5b146101d15780639fe74e261461021d57600080fd5b8063536c6bfa116100a5578063536c6bfa1461015c5780635e35359e1461017c57806379ba50971461019c57600080fd5b80630e958d6b146100d757806335f170ef1461010c5780634e12498b1461013a57600080fd5b366100d257005b600080fd5b3480156100e357600080fd5b506100f76100f2366004611c11565b6102a9565b60405190151581526020015b60405180910390f35b34801561011857600080fd5b5061012c610127366004611c64565b6102f4565b604051610103929190611cf4565b34801561014657600080fd5b5061015a610155366004611d4f565b610420565b005b34801561016857600080fd5b5061015a610177366004611d6c565b6104ec565b34801561018857600080fd5b5061015a610197366004611d98565b610550565b3480156101a857600080fd5b5061015a6105e5565b3480156101bd57600080fd5b5061015a6101cc366004611e1e565b6106e7565b3480156101dd57600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610103565b34801561022957600080fd5b5061015a610238366004611e8a565b610a48565b34801561024957600080fd5b5060025473ffffffffffffffffffffffffffffffffffffffff166101f8565b61027b610276366004611ecc565b610dc8565b604051908152602001610103565b34801561029557600080fd5b5061015a6102a4366004611d4f565b611326565b67ffffffffffffffff831660009081526003602052604080822090516002909101906102d89085908590611f8d565b9081526040519081900360200190205460ff1690509392505050565b60036020526000908152604090208054819061030f90611f9d565b80601f016020809104026020016040519081016040528092919081815260200182805461033b90611f9d565b80156103885780601f1061035d57610100808354040283529160200191610388565b820191906000526020600020905b81548152906001019060200180831161036b57829003601f168201915b50505050509080600101805461039d90611f9d565b80601f01602080910402602001604051908101604052809291908181526020018280546103c990611f9d565b80156104165780601f106103eb57610100808354040283529160200191610416565b820191906000526020600020905b8154815290600101906020018083116103f957829003601f168201915b5050505050905082565b61042861133a565b73ffffffffffffffffffffffffffffffffffffffff8116610475576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f3672b589036f39ac008505b790fcb05d484d70b65680ec64c089a3c173fdc4c890600090a35050565b6104f461133a565b6104fe82826113bd565b60405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907f6832d9be2410a86571981e1e60fd4c1f9ea2a1034d6102a2b7d6c5e480adf02e9060200160405180910390a35050565b61055861133a565b61057973ffffffffffffffffffffffffffffffffffffffff84168383611517565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f6832d9be2410a86571981e1e60fd4c1f9ea2a1034d6102a2b7d6c5e480adf02e836040516105d891815260200190565b60405180910390a3505050565b60015473ffffffffffffffffffffffffffffffffffffffff16331461066b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6106ef61133a565b60005b81811015610892576003600084848481811061071057610710611ff0565b9050602002810190610722919061201f565b610730906020810190611c64565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020183838381811061076757610767611ff0565b9050602002810190610779919061201f565b61078790602081019061205d565b604051610795929190611f8d565b90815260405190819003602001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558282828181106107dc576107dc611ff0565b90506020028101906107ee919061201f565b6107fc90602081019061205d565b60405161080a929190611f8d565b604051809103902083838381811061082457610824611ff0565b9050602002810190610836919061201f565b610844906020810190611c64565b67ffffffffffffffff167f8b7bd63124894be95396727f5b0df205ee011641e77fdab67e084f997ceb90816000604051610882911515815260200190565b60405180910390a36001016106f2565b5060005b83811015610a41576001600360008787858181106108b6576108b6611ff0565b90506020028101906108c8919061201f565b6108d6906020810190611c64565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020186868481811061090d5761090d611ff0565b905060200281019061091f919061201f565b61092d90602081019061205d565b60405161093b929190611f8d565b90815260405190819003602001902080549115157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0090921691909117905584848281811061098b5761098b611ff0565b905060200281019061099d919061201f565b6109ab90602081019061205d565b6040516109b9929190611f8d565b60405180910390208585838181106109d3576109d3611ff0565b90506020028101906109e5919061201f565b6109f3906020810190611c64565b67ffffffffffffffff167f8b7bd63124894be95396727f5b0df205ee011641e77fdab67e084f997ceb90816001604051610a31911515815260200190565b60405180910390a3600101610896565b5050505050565b610a5061133a565b60005b81811015610dc357828282818110610a6d57610a6d611ff0565b9050602002810190610a7f91906120c2565b610a90906040810190602001612104565b610b635760036000848484818110610aaa57610aaa611ff0565b9050602002810190610abc91906120c2565b610aca906020810190611c64565b67ffffffffffffffff16815260208101919091526040016000908120610aef91611b5d565b7f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916838383818110610b2257610b22611ff0565b9050602002810190610b3491906120c2565b610b42906020810190611c64565b60405167ffffffffffffffff909116815260200160405180910390a1610dbb565b828282818110610b7557610b75611ff0565b9050602002810190610b8791906120c2565b610b9590604081019061205d565b9050600003610bd0576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060036000858585818110610be857610be8611ff0565b9050602002810190610bfa91906120c2565b610c08906020810190611c64565b67ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000209050838383818110610c3e57610c3e611ff0565b9050602002810190610c5091906120c2565b610c5e90604081019061205d565b8291610c6b9190836121a0565b50838383818110610c7e57610c7e611ff0565b9050602002810190610c9091906120c2565b610c9e90606081019061205d565b159050610ce957838383818110610cb757610cb7611ff0565b9050602002810190610cc991906120c2565b610cd790606081019061205d565b6001830191610ce79190836121a0565b505b7f1ced5bcae649ed29cebfa0010298ad6794bf3822e8cb754a6eee5353a9a87212848484818110610d1c57610d1c611ff0565b9050602002810190610d2e91906120c2565b610d3c906020810190611c64565b858585818110610d4e57610d4e611ff0565b9050602002810190610d6091906120c2565b610d6e90604081019061205d565b878787818110610d8057610d80611ff0565b9050602002810190610d9291906120c2565b610da090606081019061205d565b604051610db1959493929190612303565b60405180910390a1505b600101610a53565b505050565b67ffffffffffffffff8616600090815260036020526040812080548891908190610df190611f9d565b9050600003610e38576040517fd79f2ea400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff83166004820152602401610662565b6040805160a08101825267ffffffffffffffff8b16600090815260036020529182208054829190610e6890611f9d565b80601f0160208091040260200160405190810160405280929190818152602001828054610e9490611f9d565b8015610ee15780601f10610eb657610100808354040283529160200191610ee1565b820191906000526020600020905b815481529060010190602001808311610ec457829003601f168201915b5050505050815260200188888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525093855250506040805160208e8102820181019092528d815293810193928e92508d9182919085015b82821015610f7157610f6260408302860136819003810190612346565b81526020019060010190610f45565b505050505081526020018673ffffffffffffffffffffffffffffffffffffffff168152602001600360008d67ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206001018054610fcc90611f9d565b80601f0160208091040260200160405190810160405280929190818152602001828054610ff890611f9d565b80156110455780601f1061101a57610100808354040283529160200191611045565b820191906000526020600020905b81548152906001019060200180831161102857829003601f168201915b5050505050815250905060005b88811015611147576110c133308c8c8581811061107157611071611ff0565b905060400201602001358d8d8681811061108d5761108d611ff0565b6110a39260206040909202019081019150611d4f565b73ffffffffffffffffffffffffffffffffffffffff169291906115eb565b60025461113f9073ffffffffffffffffffffffffffffffffffffffff168b8b848181106110f0576110f0611ff0565b905060400201602001358c8c8581811061110c5761110c611ff0565b6111229260206040909202019081019150611d4f565b73ffffffffffffffffffffffffffffffffffffffff16919061164f565b600101611052565b506002546040517f20487ded00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff16906320487ded906111a1908e90869060040161239e565b602060405180830381865afa1580156111be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111e291906124b3565b905073ffffffffffffffffffffffffffffffffffffffff8616156112495761122273ffffffffffffffffffffffffffffffffffffffff87163330846115eb565b6002546112499073ffffffffffffffffffffffffffffffffffffffff88811691168361164f565b60025473ffffffffffffffffffffffffffffffffffffffff908116906396f4e9f99088161561127957600061127b565b825b8d856040518463ffffffff1660e01b815260040161129a92919061239e565b60206040518083038185885af11580156112b8573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906112dd91906124b3565b94507f54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a8560405161131091815260200190565b60405180910390a1505050509695505050505050565b61132e61133a565b6113378161174d565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146113bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610662565b565b80471015611427576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610662565b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114611481576040519150601f19603f3d011682016040523d82523d6000602084013e611486565b606091505b5050905080610dc3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610662565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610dc39084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611842565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526116499085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611569565b50505050565b6040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa1580156116c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116ea91906124b3565b6116f491906124cc565b60405173ffffffffffffffffffffffffffffffffffffffff85166024820152604481018290529091506116499085907f095ea7b30000000000000000000000000000000000000000000000000000000090606401611569565b3373ffffffffffffffffffffffffffffffffffffffff8216036117cc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610662565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006118a4826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661194e9092919063ffffffff16565b805190915015610dc357808060200190518101906118c2919061250c565b610dc3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610662565b606061195d8484600085611965565b949350505050565b6060824710156119f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610662565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611a209190612529565b60006040518083038185875af1925050503d8060008114611a5d576040519150601f19603f3d011682016040523d82523d6000602084013e611a62565b606091505b5091509150611a7387838387611a7e565b979650505050505050565b60608315611b14578251600003611b0d5773ffffffffffffffffffffffffffffffffffffffff85163b611b0d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610662565b508161195d565b61195d8383815115611b295781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610662919061253b565b508054611b6990611f9d565b6000825580601f10611b79575050565b601f01602090049060005260206000209081019061133791905b80821115611ba75760008155600101611b93565b5090565b803567ffffffffffffffff81168114611bc357600080fd5b919050565b60008083601f840112611bda57600080fd5b50813567ffffffffffffffff811115611bf257600080fd5b602083019150836020828501011115611c0a57600080fd5b9250929050565b600080600060408486031215611c2657600080fd5b611c2f84611bab565b9250602084013567ffffffffffffffff811115611c4b57600080fd5b611c5786828701611bc8565b9497909650939450505050565b600060208284031215611c7657600080fd5b611c7f82611bab565b9392505050565b60005b83811015611ca1578181015183820152602001611c89565b50506000910152565b60008151808452611cc2816020860160208601611c86565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b604081526000611d076040830185611caa565b8281036020840152611d198185611caa565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461133757600080fd5b8035611bc381611d22565b600060208284031215611d6157600080fd5b8135611c7f81611d22565b60008060408385031215611d7f57600080fd5b8235611d8a81611d22565b946020939093013593505050565b600080600060608486031215611dad57600080fd5b8335611db881611d22565b92506020840135611dc881611d22565b929592945050506040919091013590565b60008083601f840112611deb57600080fd5b50813567ffffffffffffffff811115611e0357600080fd5b6020830191508360208260051b8501011115611c0a57600080fd5b60008060008060408587031215611e3457600080fd5b843567ffffffffffffffff80821115611e4c57600080fd5b611e5888838901611dd9565b90965094506020870135915080821115611e7157600080fd5b50611e7e87828801611dd9565b95989497509550505050565b60008060208385031215611e9d57600080fd5b823567ffffffffffffffff811115611eb457600080fd5b611ec085828601611dd9565b90969095509350505050565b60008060008060008060808789031215611ee557600080fd5b611eee87611bab565b9550602087013567ffffffffffffffff80821115611f0b57600080fd5b818901915089601f830112611f1f57600080fd5b813581811115611f2e57600080fd5b8a60208260061b8501011115611f4357600080fd5b602083019750809650506040890135915080821115611f6157600080fd5b50611f6e89828a01611bc8565b9094509250611f81905060608801611d44565b90509295509295509295565b8183823760009101908152919050565b600181811c90821680611fb157607f821691505b602082108103611fea577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc183360301811261205357600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261209257600080fd5b83018035915067ffffffffffffffff8211156120ad57600080fd5b602001915036819003821315611c0a57600080fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8183360301811261205357600080fd5b801515811461133757600080fd5b60006020828403121561211657600080fd5b8135611c7f816120f6565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f821115610dc3576000816000526020600020601f850160051c810160208610156121795750805b601f850160051c820191505b8181101561219857828155600101612185565b505050505050565b67ffffffffffffffff8311156121b8576121b8612121565b6121cc836121c68354611f9d565b83612150565b6000601f84116001811461221e57600085156121e85750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355610a41565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b8281101561226d578685013582556020948501946001909201910161224d565b50868210156122a8577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b67ffffffffffffffff861681526060602082015260006123276060830186886122ba565b828103604084015261233a8185876122ba565b98975050505050505050565b60006040828403121561235857600080fd5b6040516040810181811067ffffffffffffffff8211171561237b5761237b612121565b604052823561238981611d22565b81526020928301359281019290925250919050565b6000604067ffffffffffffffff851683526020604081850152845160a060408601526123cd60e0860182611caa565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808784030160608801526124088383611caa565b6040890151888203830160808a01528051808352908601945060009350908501905b80841015612469578451805173ffffffffffffffffffffffffffffffffffffffff1683528601518683015293850193600193909301929086019061242a565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a015295506124a58187611caa565b9a9950505050505050505050565b6000602082840312156124c557600080fd5b5051919050565b80820180821115612506577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b60006020828403121561251e57600080fd5b8151611c7f816120f6565b60008251612053818460208701611c86565b602081526000611c7f6020830184611caa56fea164736f6c6343000818000a",
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

func (_CCIPSender *CCIPSenderTransactor) ModifyRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _CCIPSender.contract.Transact(opts, "modifyRouter", newRouter)
}

func (_CCIPSender *CCIPSenderSession) ModifyRouter(newRouter common.Address) (*types.Transaction, error) {
	return _CCIPSender.Contract.ModifyRouter(&_CCIPSender.TransactOpts, newRouter)
}

func (_CCIPSender *CCIPSenderTransactorSession) ModifyRouter(newRouter common.Address) (*types.Transaction, error) {
	return _CCIPSender.Contract.ModifyRouter(&_CCIPSender.TransactOpts, newRouter)
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

func (_CCIPSender *CCIPSenderTransactor) WithdrawNativeToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCIPSender.contract.Transact(opts, "withdrawNativeToken", to, amount)
}

func (_CCIPSender *CCIPSenderSession) WithdrawNativeToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCIPSender.Contract.WithdrawNativeToken(&_CCIPSender.TransactOpts, to, amount)
}

func (_CCIPSender *CCIPSenderTransactorSession) WithdrawNativeToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCIPSender.Contract.WithdrawNativeToken(&_CCIPSender.TransactOpts, to, amount)
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

type CCIPSenderApprovedSenderModifiedIterator struct {
	Event *CCIPSenderApprovedSenderModified

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPSenderApprovedSenderModifiedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPSenderApprovedSenderModified)
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
		it.Event = new(CCIPSenderApprovedSenderModified)
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

func (it *CCIPSenderApprovedSenderModifiedIterator) Error() error {
	return it.fail
}

func (it *CCIPSenderApprovedSenderModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPSenderApprovedSenderModified struct {
	DestChainSelector uint64
	Recipient         common.Hash
	IsBeingApproved   bool
	Raw               types.Log
}

func (_CCIPSender *CCIPSenderFilterer) FilterApprovedSenderModified(opts *bind.FilterOpts, destChainSelector []uint64, recipient [][]byte) (*CCIPSenderApprovedSenderModifiedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CCIPSender.contract.FilterLogs(opts, "ApprovedSenderModified", destChainSelectorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CCIPSenderApprovedSenderModifiedIterator{contract: _CCIPSender.contract, event: "ApprovedSenderModified", logs: logs, sub: sub}, nil
}

func (_CCIPSender *CCIPSenderFilterer) WatchApprovedSenderModified(opts *bind.WatchOpts, sink chan<- *CCIPSenderApprovedSenderModified, destChainSelector []uint64, recipient [][]byte) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CCIPSender.contract.WatchLogs(opts, "ApprovedSenderModified", destChainSelectorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPSenderApprovedSenderModified)
				if err := _CCIPSender.contract.UnpackLog(event, "ApprovedSenderModified", log); err != nil {
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

func (_CCIPSender *CCIPSenderFilterer) ParseApprovedSenderModified(log types.Log) (*CCIPSenderApprovedSenderModified, error) {
	event := new(CCIPSenderApprovedSenderModified)
	if err := _CCIPSender.contract.UnpackLog(event, "ApprovedSenderModified", log); err != nil {
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
	Recipient           []byte
	ExtraArgsBytes      []byte
	Raw                 types.Log
}

func (_CCIPSender *CCIPSenderFilterer) FilterChainAdded(opts *bind.FilterOpts) (*CCIPSenderChainAddedIterator, error) {

	logs, sub, err := _CCIPSender.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &CCIPSenderChainAddedIterator{contract: _CCIPSender.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_CCIPSender *CCIPSenderFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *CCIPSenderChainAdded) (event.Subscription, error) {

	logs, sub, err := _CCIPSender.contract.WatchLogs(opts, "ChainAdded")
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

func (_CCIPSender *CCIPSenderFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*CCIPSenderChainRemovedIterator, error) {

	logs, sub, err := _CCIPSender.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &CCIPSenderChainRemovedIterator{contract: _CCIPSender.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_CCIPSender *CCIPSenderFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *CCIPSenderChainRemoved) (event.Subscription, error) {

	logs, sub, err := _CCIPSender.contract.WatchLogs(opts, "ChainRemoved")
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

type CCIPSenderMessageReceivedIterator struct {
	Event *CCIPSenderMessageReceived

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPSenderMessageReceivedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPSenderMessageReceived)
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
		it.Event = new(CCIPSenderMessageReceived)
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

func (it *CCIPSenderMessageReceivedIterator) Error() error {
	return it.fail
}

func (it *CCIPSenderMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPSenderMessageReceived struct {
	MessageId [32]byte
	Raw       types.Log
}

func (_CCIPSender *CCIPSenderFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*CCIPSenderMessageReceivedIterator, error) {

	logs, sub, err := _CCIPSender.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &CCIPSenderMessageReceivedIterator{contract: _CCIPSender.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

func (_CCIPSender *CCIPSenderFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *CCIPSenderMessageReceived) (event.Subscription, error) {

	logs, sub, err := _CCIPSender.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPSenderMessageReceived)
				if err := _CCIPSender.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

func (_CCIPSender *CCIPSenderFilterer) ParseMessageReceived(log types.Log) (*CCIPSenderMessageReceived, error) {
	event := new(CCIPSenderMessageReceived)
	if err := _CCIPSender.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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
	case _CCIPSender.abi.Events["ApprovedSenderModified"].ID:
		return _CCIPSender.ParseApprovedSenderModified(log)
	case _CCIPSender.abi.Events["CCIPRouterModified"].ID:
		return _CCIPSender.ParseCCIPRouterModified(log)
	case _CCIPSender.abi.Events["ChainAdded"].ID:
		return _CCIPSender.ParseChainAdded(log)
	case _CCIPSender.abi.Events["ChainRemoved"].ID:
		return _CCIPSender.ParseChainRemoved(log)
	case _CCIPSender.abi.Events["MessageReceived"].ID:
		return _CCIPSender.ParseMessageReceived(log)
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

func (CCIPSenderApprovedSenderModified) Topic() common.Hash {
	return common.HexToHash("0x8b7bd63124894be95396727f5b0df205ee011641e77fdab67e084f997ceb9081")
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

func (CCIPSenderMessageReceived) Topic() common.Hash {
	return common.HexToHash("0xe29dc34207c78fc0f6048a32f159139c33339c6d6df8b07dcd33f6d699ff2327")
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

	ModifyRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateApprovedSenders(opts *bind.TransactOpts, adds []CCIPBaseApprovedSenderUpdate, removes []CCIPBaseApprovedSenderUpdate) (*types.Transaction, error)

	WithdrawNativeToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)

	WithdrawTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterApprovedSenderModified(opts *bind.FilterOpts, destChainSelector []uint64, recipient [][]byte) (*CCIPSenderApprovedSenderModifiedIterator, error)

	WatchApprovedSenderModified(opts *bind.WatchOpts, sink chan<- *CCIPSenderApprovedSenderModified, destChainSelector []uint64, recipient [][]byte) (event.Subscription, error)

	ParseApprovedSenderModified(log types.Log) (*CCIPSenderApprovedSenderModified, error)

	FilterCCIPRouterModified(opts *bind.FilterOpts, oldRouter []common.Address, newRouter []common.Address) (*CCIPSenderCCIPRouterModifiedIterator, error)

	WatchCCIPRouterModified(opts *bind.WatchOpts, sink chan<- *CCIPSenderCCIPRouterModified, oldRouter []common.Address, newRouter []common.Address) (event.Subscription, error)

	ParseCCIPRouterModified(log types.Log) (*CCIPSenderCCIPRouterModified, error)

	FilterChainAdded(opts *bind.FilterOpts) (*CCIPSenderChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *CCIPSenderChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*CCIPSenderChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*CCIPSenderChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *CCIPSenderChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*CCIPSenderChainRemoved, error)

	FilterMessageReceived(opts *bind.FilterOpts) (*CCIPSenderMessageReceivedIterator, error)

	WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *CCIPSenderMessageReceived) (event.Subscription, error)

	ParseMessageReceived(log types.Log) (*CCIPSenderMessageReceived, error)

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
