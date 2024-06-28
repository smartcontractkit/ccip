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

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type ICCIPClientBaseapprovedSenderUpdate struct {
	DestChainSelector uint64
	Sender            []byte
}

var CCIPSenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientFeeTokenAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientNativeFeeTokenAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"disableChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraArgsBytes\",\"type\":\"bytes\"}],\"name\":\"enableChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"senderAddr\",\"type\":\"bytes\"}],\"name\":\"isApprovedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"s_chains\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgsBytes\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"internalType\":\"structICCIPClientBase.approvedSenderUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"internalType\":\"structICCIPClientBase.approvedSenderUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"}],\"name\":\"updateApprovedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNativeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162002467380380620024678339810160408190526200003491620001a8565b8033806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf81620000fd565b5050506001600160a01b038116620000ea576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b031660805250620001da565b336001600160a01b03821603620001575760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001bb57600080fd5b81516001600160a01b0381168114620001d357600080fd5b9392505050565b608051612247620002206000396000818161027201528181610ab901528181610c2601528181610d4401528181610e0801528181610ec00152610f5b01526122476000f3fe6080604052600436106100ca5760003560e01c806379ba509711610079578063b0f479a111610056578063b0f479a114610263578063d8469e4014610296578063effde240146102b6578063f2fde38b146102d757005b806379ba5097146101e25780638462a2b9146101f75780638da5cb5b1461021757005b8063536c6bfa116100a7578063536c6bfa146101745780635dc5ebdb146101945780635e35359e146101c257005b80630e958d6b146100d3578063181f5a771461010857806341eade461461015457005b366100d157005b005b3480156100df57600080fd5b506100f36100ee3660046119ac565b6102f7565b60405190151581526020015b60405180910390f35b34801561011457600080fd5b50604080518082018252601481527f4343495053656e64657220312e302e302d646576000000000000000000000000602082015290516100ff9190611a6d565b34801561016057600080fd5b506100d161016f366004611a87565b610341565b34801561018057600080fd5b506100d161018f366004611ac4565b610380565b3480156101a057600080fd5b506101b46101af366004611a87565b610396565b6040516100ff929190611af0565b3480156101ce57600080fd5b506100d16101dd366004611b29565b6104c2565b3480156101ee57600080fd5b506100d16104eb565b34801561020357600080fd5b506100d1610212366004611baf565b6105ed565b34801561022357600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ff565b34801561026f57600080fd5b507f000000000000000000000000000000000000000000000000000000000000000061023e565b3480156102a257600080fd5b506100d16102b1366004611c1b565b6107ce565b6102c96102c4366004611c9c565b610830565b6040519081526020016100ff565b3480156102e357600080fd5b506100d16102f2366004611d5d565b61103d565b67ffffffffffffffff8316600090815260026020819052604080832090519101906103259085908590611d7a565b9081526040519081900360200190205460ff1690509392505050565b610349611051565b67ffffffffffffffff811660009081526002602052604081209061036d82826118f8565b61037b6001830160006118f8565b505050565b610388611051565b61039282826110d4565b5050565b6002602052600090815260409020805481906103b190611d8a565b80601f01602080910402602001604051908101604052809291908181526020018280546103dd90611d8a565b801561042a5780601f106103ff5761010080835404028352916020019161042a565b820191906000526020600020905b81548152906001019060200180831161040d57829003601f168201915b50505050509080600101805461043f90611d8a565b80601f016020809104026020016040519081016040528092919081815260200182805461046b90611d8a565b80156104b85780601f1061048d576101008083540402835291602001916104b8565b820191906000526020600020905b81548152906001019060200180831161049b57829003601f168201915b5050505050905082565b6104ca611051565b61037b73ffffffffffffffffffffffffffffffffffffffff8416838361122e565b60015473ffffffffffffffffffffffffffffffffffffffff163314610571576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6105f5611051565b60005b818110156106d8576002600084848481811061061657610616611ddd565b90506020028101906106289190611e0c565b610636906020810190611a87565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020183838381811061066d5761066d611ddd565b905060200281019061067f9190611e0c565b61068d906020810190611e4a565b60405161069b929190611d7a565b90815260405190819003602001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556001016105f8565b5060005b838110156107c7576001600260008787858181106106fc576106fc611ddd565b905060200281019061070e9190611e0c565b61071c906020810190611a87565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020186868481811061075357610753611ddd565b90506020028101906107659190611e0c565b610773906020810190611e4a565b604051610781929190611d7a565b90815260405190819003602001902080549115157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009092169190911790556001016106dc565b5050505050565b6107d6611051565b67ffffffffffffffff851660009081526002602052604090206107fa848683611f26565b5080156107c75767ffffffffffffffff85166000908152600260205260409020600101610828828483611f26565b505050505050565b67ffffffffffffffff86166000908152600260205260408120805488919061085790611d8a565b905060000361089e576040517fd79f2ea400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610568565b6040805160a08101825267ffffffffffffffff8a166000908152600260205291822080548291906108ce90611d8a565b80601f01602080910402602001604051908101604052809291908181526020018280546108fa90611d8a565b80156109475780601f1061091c57610100808354040283529160200191610947565b820191906000526020600020905b81548152906001019060200180831161092a57829003601f168201915b5050505050815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525093855250506040805160208d8102820181019092528c815293810193928d92508c9182919085015b828210156109d7576109c860408302860136819003810190612040565b815260200190600101906109ab565b505050505081526020018573ffffffffffffffffffffffffffffffffffffffff168152602001600260008c67ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206001018054610a3290611d8a565b80601f0160208091040260200160405190810160405280929190818152602001828054610a5e90611d8a565b8015610aab5780601f10610a8057610100808354040283529160200191610aab565b820191906000526020600020905b815481529060010190602001808311610a8e57829003601f168201915b5050505050815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166320487ded8b846040518363ffffffff1660e01b8152600401610b12929190612098565b602060405180830381865afa158015610b2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5391906121ad565b905060005b88811015610da457610bc733308c8c85818110610b7757610b77611ddd565b905060400201602001358d8d86818110610b9357610b93611ddd565b610ba99260206040909202019081019150611d5d565b73ffffffffffffffffffffffffffffffffffffffff16929190611302565b8573ffffffffffffffffffffffffffffffffffffffff168a8a83818110610bf057610bf0611ddd565b610c069260206040909202019081019150611d5d565b73ffffffffffffffffffffffffffffffffffffffff1614610cab57610ca67f00000000000000000000000000000000000000000000000000000000000000008b8b84818110610c5757610c57611ddd565b905060400201602001358c8c85818110610c7357610c73611ddd565b610c899260206040909202019081019150611d5d565b73ffffffffffffffffffffffffffffffffffffffff169190611366565b610d9c565b8573ffffffffffffffffffffffffffffffffffffffff168a8a83818110610cd457610cd4611ddd565b610cea9260206040909202019081019150611d5d565b73ffffffffffffffffffffffffffffffffffffffff16148015610d22575073ffffffffffffffffffffffffffffffffffffffff861615155b15610d9c57610d3f3330848d8d86818110610b9357610b93611ddd565b610d9c7f0000000000000000000000000000000000000000000000000000000000000000838c8c85818110610d7657610d76611ddd565b90506040020160200135610d8a91906121c6565b8c8c85818110610c7357610c73611ddd565b600101610b58565b5073ffffffffffffffffffffffffffffffffffffffff851615801590610e7d57506040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116602483015286169063dd62ed3e90604401602060405180830381865afa158015610e57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e7b91906121ad565b155b15610eea57610ea473ffffffffffffffffffffffffffffffffffffffff8616333084611302565b610ee573ffffffffffffffffffffffffffffffffffffffff86167f000000000000000000000000000000000000000000000000000000000000000083611366565b610f44565b73ffffffffffffffffffffffffffffffffffffffff8516158015610f0d57508034105b15610f44576040517f07da6ee600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116906396f4e9f990871615610f91576000610f93565b825b8c856040518463ffffffff1660e01b8152600401610fb2929190612098565b60206040518083038185885af1158015610fd0573d6000803e3d6000fd5b50505050506040513d601f19601f82011682018060405250810190610ff591906121ad565b93507f54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a8460405161102891815260200190565b60405180910390a15050509695505050505050565b611045611051565b61104e816114e8565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146110d2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610568565b565b8047101561113e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610568565b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114611198576040519150601f19603f3d011682016040523d82523d6000602084013e61119d565b606091505b505090508061037b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610568565b60405173ffffffffffffffffffffffffffffffffffffffff831660248201526044810182905261037b9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526115dd565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526113609085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611280565b50505050565b80158061140657506040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff838116602483015284169063dd62ed3e90604401602060405180830381865afa1580156113e0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061140491906121ad565b155b611492576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608401610568565b60405173ffffffffffffffffffffffffffffffffffffffff831660248201526044810182905261037b9084907f095ea7b30000000000000000000000000000000000000000000000000000000090606401611280565b3373ffffffffffffffffffffffffffffffffffffffff821603611567576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610568565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600061163f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166116e99092919063ffffffff16565b80519091501561037b578080602001905181019061165d9190612206565b61037b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610568565b60606116f88484600085611700565b949350505050565b606082471015611792576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610568565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516117bb9190612228565b60006040518083038185875af1925050503d80600081146117f8576040519150601f19603f3d011682016040523d82523d6000602084013e6117fd565b606091505b509150915061180e87838387611819565b979650505050505050565b606083156118af5782516000036118a85773ffffffffffffffffffffffffffffffffffffffff85163b6118a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610568565b50816116f8565b6116f883838151156118c45781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105689190611a6d565b50805461190490611d8a565b6000825580601f10611914575050565b601f01602090049060005260206000209081019061104e91905b80821115611942576000815560010161192e565b5090565b803567ffffffffffffffff8116811461195e57600080fd5b919050565b60008083601f84011261197557600080fd5b50813567ffffffffffffffff81111561198d57600080fd5b6020830191508360208285010111156119a557600080fd5b9250929050565b6000806000604084860312156119c157600080fd5b6119ca84611946565b9250602084013567ffffffffffffffff8111156119e657600080fd5b6119f286828701611963565b9497909650939450505050565b60005b83811015611a1a578181015183820152602001611a02565b50506000910152565b60008151808452611a3b8160208601602086016119ff565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611a806020830184611a23565b9392505050565b600060208284031215611a9957600080fd5b611a8082611946565b73ffffffffffffffffffffffffffffffffffffffff8116811461104e57600080fd5b60008060408385031215611ad757600080fd5b8235611ae281611aa2565b946020939093013593505050565b604081526000611b036040830185611a23565b8281036020840152611b158185611a23565b95945050505050565b803561195e81611aa2565b600080600060608486031215611b3e57600080fd5b8335611b4981611aa2565b92506020840135611b5981611aa2565b929592945050506040919091013590565b60008083601f840112611b7c57600080fd5b50813567ffffffffffffffff811115611b9457600080fd5b6020830191508360208260051b85010111156119a557600080fd5b60008060008060408587031215611bc557600080fd5b843567ffffffffffffffff80821115611bdd57600080fd5b611be988838901611b6a565b90965094506020870135915080821115611c0257600080fd5b50611c0f87828801611b6a565b95989497509550505050565b600080600080600060608688031215611c3357600080fd5b611c3c86611946565b9450602086013567ffffffffffffffff80821115611c5957600080fd5b611c6589838a01611963565b90965094506040880135915080821115611c7e57600080fd5b50611c8b88828901611963565b969995985093965092949392505050565b60008060008060008060808789031215611cb557600080fd5b611cbe87611946565b9550602087013567ffffffffffffffff80821115611cdb57600080fd5b818901915089601f830112611cef57600080fd5b813581811115611cfe57600080fd5b8a60208260061b8501011115611d1357600080fd5b602083019750809650506040890135915080821115611d3157600080fd5b50611d3e89828a01611963565b9094509250611d51905060608801611b1e565b90509295509295509295565b600060208284031215611d6f57600080fd5b8135611a8081611aa2565b8183823760009101908152919050565b600181811c90821680611d9e57607f821691505b602082108103611dd7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc1833603018112611e4057600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611e7f57600080fd5b83018035915067ffffffffffffffff821115611e9a57600080fd5b6020019150368190038213156119a557600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82111561037b576000816000526020600020601f850160051c81016020861015611f075750805b601f850160051c820191505b8181101561082857828155600101611f13565b67ffffffffffffffff831115611f3e57611f3e611eaf565b611f5283611f4c8354611d8a565b83611ede565b6000601f841160018114611fa45760008515611f6e5750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b1783556107c7565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b82811015611ff35786850135825560209485019460019092019101611fd3565b508682101561202e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b60006040828403121561205257600080fd5b6040516040810181811067ffffffffffffffff8211171561207557612075611eaf565b604052823561208381611aa2565b81526020928301359281019290925250919050565b6000604067ffffffffffffffff851683526020604081850152845160a060408601526120c760e0860182611a23565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808784030160608801526121028383611a23565b6040890151888203830160808a01528051808352908601945060009350908501905b80841015612163578451805173ffffffffffffffffffffffffffffffffffffffff16835286015186830152938501936001939093019290860190612124565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a0152955061219f8187611a23565b9a9950505050505050505050565b6000602082840312156121bf57600080fd5b5051919050565b80820180821115612200577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b60006020828403121561221857600080fd5b81518015158114611a8057600080fd5b60008251611e408184602087016119ff56fea164736f6c6343000818000a",
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

func (_CCIPSender *CCIPSenderCaller) SChains(opts *bind.CallOpts, arg0 uint64) (SChains,

	error) {
	var out []interface{}
	err := _CCIPSender.contract.Call(opts, &out, "s_chains", arg0)

	outstruct := new(SChains)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recipient = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ExtraArgsBytes = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

func (_CCIPSender *CCIPSenderSession) SChains(arg0 uint64) (SChains,

	error) {
	return _CCIPSender.Contract.SChains(&_CCIPSender.CallOpts, arg0)
}

func (_CCIPSender *CCIPSenderCallerSession) SChains(arg0 uint64) (SChains,

	error) {
	return _CCIPSender.Contract.SChains(&_CCIPSender.CallOpts, arg0)
}

func (_CCIPSender *CCIPSenderCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CCIPSender.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CCIPSender *CCIPSenderSession) TypeAndVersion() (string, error) {
	return _CCIPSender.Contract.TypeAndVersion(&_CCIPSender.CallOpts)
}

func (_CCIPSender *CCIPSenderCallerSession) TypeAndVersion() (string, error) {
	return _CCIPSender.Contract.TypeAndVersion(&_CCIPSender.CallOpts)
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

func (_CCIPSender *CCIPSenderTransactor) CcipSend(opts *bind.TransactOpts, destChainSelector uint64, tokenAmounts []ClientEVMTokenAmount, data []byte, feeToken common.Address) (*types.Transaction, error) {
	return _CCIPSender.contract.Transact(opts, "ccipSend", destChainSelector, tokenAmounts, data, feeToken)
}

func (_CCIPSender *CCIPSenderSession) CcipSend(destChainSelector uint64, tokenAmounts []ClientEVMTokenAmount, data []byte, feeToken common.Address) (*types.Transaction, error) {
	return _CCIPSender.Contract.CcipSend(&_CCIPSender.TransactOpts, destChainSelector, tokenAmounts, data, feeToken)
}

func (_CCIPSender *CCIPSenderTransactorSession) CcipSend(destChainSelector uint64, tokenAmounts []ClientEVMTokenAmount, data []byte, feeToken common.Address) (*types.Transaction, error) {
	return _CCIPSender.Contract.CcipSend(&_CCIPSender.TransactOpts, destChainSelector, tokenAmounts, data, feeToken)
}

func (_CCIPSender *CCIPSenderTransactor) DisableChain(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error) {
	return _CCIPSender.contract.Transact(opts, "disableChain", chainSelector)
}

func (_CCIPSender *CCIPSenderSession) DisableChain(chainSelector uint64) (*types.Transaction, error) {
	return _CCIPSender.Contract.DisableChain(&_CCIPSender.TransactOpts, chainSelector)
}

func (_CCIPSender *CCIPSenderTransactorSession) DisableChain(chainSelector uint64) (*types.Transaction, error) {
	return _CCIPSender.Contract.DisableChain(&_CCIPSender.TransactOpts, chainSelector)
}

func (_CCIPSender *CCIPSenderTransactor) EnableChain(opts *bind.TransactOpts, chainSelector uint64, recipient []byte, _extraArgsBytes []byte) (*types.Transaction, error) {
	return _CCIPSender.contract.Transact(opts, "enableChain", chainSelector, recipient, _extraArgsBytes)
}

func (_CCIPSender *CCIPSenderSession) EnableChain(chainSelector uint64, recipient []byte, _extraArgsBytes []byte) (*types.Transaction, error) {
	return _CCIPSender.Contract.EnableChain(&_CCIPSender.TransactOpts, chainSelector, recipient, _extraArgsBytes)
}

func (_CCIPSender *CCIPSenderTransactorSession) EnableChain(chainSelector uint64, recipient []byte, _extraArgsBytes []byte) (*types.Transaction, error) {
	return _CCIPSender.Contract.EnableChain(&_CCIPSender.TransactOpts, chainSelector, recipient, _extraArgsBytes)
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

func (_CCIPSender *CCIPSenderTransactor) UpdateApprovedSenders(opts *bind.TransactOpts, adds []ICCIPClientBaseapprovedSenderUpdate, removes []ICCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error) {
	return _CCIPSender.contract.Transact(opts, "updateApprovedSenders", adds, removes)
}

func (_CCIPSender *CCIPSenderSession) UpdateApprovedSenders(adds []ICCIPClientBaseapprovedSenderUpdate, removes []ICCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error) {
	return _CCIPSender.Contract.UpdateApprovedSenders(&_CCIPSender.TransactOpts, adds, removes)
}

func (_CCIPSender *CCIPSenderTransactorSession) UpdateApprovedSenders(adds []ICCIPClientBaseapprovedSenderUpdate, removes []ICCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error) {
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

func (_CCIPSender *CCIPSenderTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CCIPSender.contract.RawTransact(opts, calldata)
}

func (_CCIPSender *CCIPSenderSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CCIPSender.Contract.Fallback(&_CCIPSender.TransactOpts, calldata)
}

func (_CCIPSender *CCIPSenderTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CCIPSender.Contract.Fallback(&_CCIPSender.TransactOpts, calldata)
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

type SChains struct {
	Recipient      []byte
	ExtraArgsBytes []byte
}

func (_CCIPSender *CCIPSender) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CCIPSender.abi.Events["MessageReceived"].ID:
		return _CCIPSender.ParseMessageReceived(log)
	case _CCIPSender.abi.Events["MessageSent"].ID:
		return _CCIPSender.ParseMessageSent(log)
	case _CCIPSender.abi.Events["OwnershipTransferRequested"].ID:
		return _CCIPSender.ParseOwnershipTransferRequested(log)
	case _CCIPSender.abi.Events["OwnershipTransferred"].ID:
		return _CCIPSender.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
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

func (_CCIPSender *CCIPSender) Address() common.Address {
	return _CCIPSender.address
}

type CCIPSenderInterface interface {
	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IsApprovedSender(opts *bind.CallOpts, sourceChainSelector uint64, senderAddr []byte) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SChains(opts *bind.CallOpts, arg0 uint64) (SChains,

		error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destChainSelector uint64, tokenAmounts []ClientEVMTokenAmount, data []byte, feeToken common.Address) (*types.Transaction, error)

	DisableChain(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error)

	EnableChain(opts *bind.TransactOpts, chainSelector uint64, recipient []byte, _extraArgsBytes []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateApprovedSenders(opts *bind.TransactOpts, adds []ICCIPClientBaseapprovedSenderUpdate, removes []ICCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error)

	WithdrawNativeToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)

	WithdrawTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

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

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
