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

type CCIPClientBaseapprovedSenderUpdate struct {
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrorCase\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageNotFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"MessageFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"disableChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraArgsBytes\",\"type\":\"bytes\"}],\"name\":\"enableChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"getMessageContents\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"getMessageStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"senderAddr\",\"type\":\"bytes\"}],\"name\":\"isApprovedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"processMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"forwardingAddress\",\"type\":\"address\"}],\"name\":\"retryFailedMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"s_chainConfigs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgsBytes\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"simRevert\",\"type\":\"bool\"}],\"name\":\"setSimRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPClientBase.approvedSenderUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPClientBase.approvedSenderUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"}],\"name\":\"updateApprovedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNativeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162002a8738038062002a878339810160408190526200003491620001a8565b8033806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf81620000fd565b5050506001600160a01b038116620000ea576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b031660805250620001da565b336001600160a01b03821603620001575760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001bb57600080fd5b81516001600160a01b0381168114620001d357600080fd5b9392505050565b60805161288a620001fd6000396000818161039d0152610e3a015261288a6000f3fe60806040526004361061012d5760003560e01c80636939cd97116100a55780638da5cb5b11610074578063cf6730f811610059578063cf6730f8146103c1578063d8469e40146103e1578063f2fde38b1461040157610134565b80638da5cb5b14610342578063b0f479a11461038e57610134565b80636939cd97146102c057806379ba5097146102ed5780638462a2b91461030257806385572ffb1461032257610134565b806341eade46116100fc57806352f813c3116100e157806352f813c314610260578063536c6bfa146102805780635e35359e146102a057610134565b806341eade46146102125780635075a9d41461023257610134565b80630e958d6b14610142578063181f5a771461017757806335f170ef146101c3578063369f7f66146101f257610134565b3661013457005b34801561014057600080fd5b005b34801561014e57600080fd5b5061016261015d366004611c12565b610421565b60405190151581526020015b60405180910390f35b34801561018357600080fd5b50604080518082018252601681527f43434950526563656976657220312e302e302d646576000000000000000000006020820152905161016e9190611cd5565b3480156101cf57600080fd5b506101e36101de366004611ce8565b61046c565b60405161016e93929190611d05565b3480156101fe57600080fd5b5061014061020d366004611d5e565b6105a3565b34801561021e57600080fd5b5061014061022d366004611ce8565b61085e565b34801561023e57600080fd5b5061025261024d366004611d8e565b6108a9565b60405190815260200161016e565b34801561026c57600080fd5b5061014061027b366004611db5565b6108bc565b34801561028c57600080fd5b5061014061029b366004611dd2565b6108f5565b3480156102ac57600080fd5b506101406102bb366004611dfe565b61090b565b3480156102cc57600080fd5b506102e06102db366004611d8e565b610939565b60405161016e9190611e3f565b3480156102f957600080fd5b50610140610b44565b34801561030e57600080fd5b5061014061031d366004611f6b565b610c41565b34801561032e57600080fd5b5061014061033d366004611fd7565b610e22565b34801561034e57600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161016e565b34801561039a57600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610369565b3480156103cd57600080fd5b506101406103dc366004611fd7565b611055565b3480156103ed57600080fd5b506101406103fc366004612012565b611191565b34801561040d57600080fd5b5061014061041c366004612095565b611212565b67ffffffffffffffff8316600090815260026020526040808220905160039091019061045090859085906120b2565b9081526040519081900360200190205460ff1690509392505050565b6002602052600090815260409020805460018201805460ff9092169291610492906120c2565b80601f01602080910402602001604051908101604052809291908181526020018280546104be906120c2565b801561050b5780601f106104e05761010080835404028352916020019161050b565b820191906000526020600020905b8154815290600101906020018083116104ee57829003601f168201915b505050505090806002018054610520906120c2565b80601f016020809104026020016040519081016040528092919081815260200182805461054c906120c2565b80156105995780601f1061056e57610100808354040283529160200191610599565b820191906000526020600020905b81548152906001019060200180831161057c57829003601f168201915b5050505050905083565b6105ab611226565b60016105b86004846112a9565b146105f7576040517fb6e78260000000000000000000000000000000000000000000000000000000008152600481018390526024015b60405180910390fd5b6106078260005b600491906112bc565b506000828152600360209081526040808320815160a08101835281548152600182015467ffffffffffffffff1693810193909352600281018054919284019161064f906120c2565b80601f016020809104026020016040519081016040528092919081815260200182805461067b906120c2565b80156106c85780601f1061069d576101008083540402835291602001916106c8565b820191906000526020600020905b8154815290600101906020018083116106ab57829003601f168201915b505050505081526020016003820180546106e1906120c2565b80601f016020809104026020016040519081016040528092919081815260200182805461070d906120c2565b801561075a5780601f1061072f5761010080835404028352916020019161075a565b820191906000526020600020905b81548152906001019060200180831161073d57829003601f168201915b5050505050815260200160048201805480602002602001604051908101604052809291908181526020016000905b828210156107dd5760008481526020908190206040805180820190915260028502909101805473ffffffffffffffffffffffffffffffffffffffff168252600190810154828401529083529092019101610788565b505050915250506040805173ffffffffffffffffffffffffffffffffffffffff85166020820152919250610822918391016040516020818303038152906040526112d1565b61082d600484611376565b5060405183907fef3bf8c64bc480286c4f3503b870ceb23e648d2d902e31fb7bb46680da6de8ad90600090a2505050565b610866611226565b67ffffffffffffffff16600090815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b60006108b66004836112a9565b92915050565b6108c4611226565b600780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b6108fd611226565b6109078282611382565b5050565b610913611226565b61093473ffffffffffffffffffffffffffffffffffffffff841683836114dc565b505050565b6040805160a08082018352600080835260208084018290526060848601819052808501819052608085015285825260038152908490208451928301855280548352600181015467ffffffffffffffff16918301919091526002810180549394929391928401916109a8906120c2565b80601f01602080910402602001604051908101604052809291908181526020018280546109d4906120c2565b8015610a215780601f106109f657610100808354040283529160200191610a21565b820191906000526020600020905b815481529060010190602001808311610a0457829003601f168201915b50505050508152602001600382018054610a3a906120c2565b80601f0160208091040260200160405190810160405280929190818152602001828054610a66906120c2565b8015610ab35780601f10610a8857610100808354040283529160200191610ab3565b820191906000526020600020905b815481529060010190602001808311610a9657829003601f168201915b5050505050815260200160048201805480602002602001604051908101604052809291908181526020016000905b82821015610b365760008481526020908190206040805180820190915260028502909101805473ffffffffffffffffffffffffffffffffffffffff168252600190810154828401529083529092019101610ae1565b505050915250909392505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610bc5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016105ee565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610c49611226565b60005b81811015610d2c5760026000848484818110610c6a57610c6a612115565b9050602002810190610c7c9190612144565b610c8a906020810190611ce8565b67ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600301838383818110610cc157610cc1612115565b9050602002810190610cd39190612144565b610ce1906020810190612182565b604051610cef9291906120b2565b90815260405190819003602001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055600101610c4c565b5060005b83811015610e1b57600160026000878785818110610d5057610d50612115565b9050602002810190610d629190612144565b610d70906020810190611ce8565b67ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600301868684818110610da757610da7612115565b9050602002810190610db99190612144565b610dc7906020810190612182565b604051610dd59291906120b2565b90815260405190819003602001902080549115157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909216919091179055600101610d30565b5050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610e93576040517fd7f733340000000000000000000000000000000000000000000000000000000081523360048201526024016105ee565b610ea36040820160208301611ce8565b67ffffffffffffffff81166000908152600260205260409020600181018054610ecb906120c2565b15905080610eda5750805460ff165b15610f1d576040517fd79f2ea400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff831660048201526024016105ee565b6040517fcf6730f8000000000000000000000000000000000000000000000000000000008152309063cf6730f890610f599086906004016122f4565b600060405180830381600087803b158015610f7357600080fd5b505af1925050508015610f84575060015b611024573d808015610fb2576040519150601f19603f3d011682016040523d82523d6000602084013e610fb7565b606091505b50610fc4843560016105fe565b50833560009081526003602052604090208490610fe182826126f5565b50506040518435907f55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f90611016908490611cd5565b60405180910390a250505050565b6040518335907fdf6958669026659bac75ba986685e11a7d271284989f565f2802522663e9a70f90600090a2505050565b33301461108e576040517f14d4a4e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61109e6040820160208301611ce8565b6110ab6040830183612182565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525067ffffffffffffffff86168152600260205260409081902090516003909101935061110992508491506127ef565b9081526040519081900360200190205460ff1661115457806040517f5075bb380000000000000000000000000000000000000000000000000000000081526004016105ee9190611cd5565b60075460ff1615610934576040517f79f79e0b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611199611226565b67ffffffffffffffff85166000908152600260205260409020600181016111c1858783612479565b5081156111d957600281016111d7838583612479565b505b805460ff161561120a5780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681555b505050505050565b61121a611226565b61122381611569565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146112a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016105ee565b565b60006112b5838361165e565b9392505050565b60006112c98484846116e8565b949350505050565b6000818060200190518101906112e79190612801565b905060005b8360800151518110156113705760008460800151828151811061131157611311612115565b602002602001015160200151905060008560800151838151811061133757611337612115565b602090810291909101015151905061136673ffffffffffffffffffffffffffffffffffffffff821685846114dc565b50506001016112ec565b50505050565b60006112b58383611705565b804710156113ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e636500000060448201526064016105ee565b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114611446576040519150601f19603f3d011682016040523d82523d6000602084013e61144b565b606091505b5050905080610934576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d6179206861766520726576657274656400000000000060648201526084016105ee565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610934908490611722565b3373ffffffffffffffffffffffffffffffffffffffff8216036115e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016105ee565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000818152600283016020526040812054801515806116825750611682848461182e565b6112b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b6579000060448201526064016105ee565b600082815260028401602052604081208290556112c9848461183a565b600081815260028301602052604081208190556112b58383611846565b6000611784826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166118529092919063ffffffff16565b80519091501561093457808060200190518101906117a2919061281e565b610934576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016105ee565b60006112b58383611861565b60006112b58383611879565b60006112b583836118c8565b60606112c984846000856119bb565b600081815260018301602052604081205415156112b5565b60008181526001830160205260408120546118c0575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556108b6565b5060006108b6565b600081815260018301602052604081205480156119b15760006118ec60018361283b565b85549091506000906119009060019061283b565b905081811461196557600086600001828154811061192057611920612115565b906000526020600020015490508087600001848154811061194357611943612115565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806119765761197661284e565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506108b6565b60009150506108b6565b606082471015611a4d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016105ee565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611a7691906127ef565b60006040518083038185875af1925050503d8060008114611ab3576040519150601f19603f3d011682016040523d82523d6000602084013e611ab8565b606091505b5091509150611ac987838387611ad4565b979650505050505050565b60608315611b6a578251600003611b635773ffffffffffffffffffffffffffffffffffffffff85163b611b63576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016105ee565b50816112c9565b6112c98383815115611b7f5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ee9190611cd5565b67ffffffffffffffff8116811461122357600080fd5b60008083601f840112611bdb57600080fd5b50813567ffffffffffffffff811115611bf357600080fd5b602083019150836020828501011115611c0b57600080fd5b9250929050565b600080600060408486031215611c2757600080fd5b8335611c3281611bb3565b9250602084013567ffffffffffffffff811115611c4e57600080fd5b611c5a86828701611bc9565b9497909650939450505050565b60005b83811015611c82578181015183820152602001611c6a565b50506000910152565b60008151808452611ca3816020860160208601611c67565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006112b56020830184611c8b565b600060208284031215611cfa57600080fd5b81356112b581611bb3565b8315158152606060208201526000611d206060830185611c8b565b8281036040840152611d328185611c8b565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461122357600080fd5b60008060408385031215611d7157600080fd5b823591506020830135611d8381611d3c565b809150509250929050565b600060208284031215611da057600080fd5b5035919050565b801515811461122357600080fd5b600060208284031215611dc757600080fd5b81356112b581611da7565b60008060408385031215611de557600080fd5b8235611df081611d3c565b946020939093013593505050565b600080600060608486031215611e1357600080fd5b8335611e1e81611d3c565b92506020840135611e2e81611d3c565b929592945050506040919091013590565b6000602080835283518184015280840151604067ffffffffffffffff821660408601526040860151915060a06060860152611e7d60c0860183611c8b565b915060608601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe080878503016080880152611eb98483611c8b565b608089015188820390920160a089015281518082529186019450600092508501905b80831015611f1a578451805173ffffffffffffffffffffffffffffffffffffffff16835286015186830152938501936001929092019190830190611edb565b50979650505050505050565b60008083601f840112611f3857600080fd5b50813567ffffffffffffffff811115611f5057600080fd5b6020830191508360208260051b8501011115611c0b57600080fd5b60008060008060408587031215611f8157600080fd5b843567ffffffffffffffff80821115611f9957600080fd5b611fa588838901611f26565b90965094506020870135915080821115611fbe57600080fd5b50611fcb87828801611f26565b95989497509550505050565b600060208284031215611fe957600080fd5b813567ffffffffffffffff81111561200057600080fd5b820160a081850312156112b557600080fd5b60008060008060006060868803121561202a57600080fd5b853561203581611bb3565b9450602086013567ffffffffffffffff8082111561205257600080fd5b61205e89838a01611bc9565b9096509450604088013591508082111561207757600080fd5b5061208488828901611bc9565b969995985093965092949392505050565b6000602082840312156120a757600080fd5b81356112b581611d3c565b8183823760009101908152919050565b600181811c908216806120d657607f821691505b60208210810361210f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc183360301811261217857600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126121b757600080fd5b83018035915067ffffffffffffffff8211156121d257600080fd5b602001915036819003821315611c0b57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261221c57600080fd5b830160208101925035905067ffffffffffffffff81111561223c57600080fd5b803603821315611c0b57600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8183526000602080850194508260005b858110156122e95781356122b781611d3c565b73ffffffffffffffffffffffffffffffffffffffff1687528183013583880152604096870196909101906001016122a4565b509495945050505050565b60208152813560208201526000602083013561230f81611bb3565b67ffffffffffffffff808216604085015261232d60408601866121e7565b925060a0606086015261234460c08601848361224b565b92505061235460608601866121e7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08087860301608088015261238a85838561224b565b9450608088013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18836030183126123c357600080fd5b602092880192830192359150838211156123dc57600080fd5b8160061b36038313156123ee57600080fd5b8685030160a0870152611ac9848284612294565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f821115610934576000816000526020600020601f850160051c8101602086101561245a5750805b601f850160051c820191505b8181101561120a57828155600101612466565b67ffffffffffffffff83111561249157612491612402565b6124a58361249f83546120c2565b83612431565b6000601f8411600181146124f757600085156124c15750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355610e1b565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b828110156125465786850135825560209485019460019092019101612526565b5086821015612581577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81356125cd81611d3c565b73ffffffffffffffffffffffffffffffffffffffff81167fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825550602082013560018201555050565b6801000000000000000083111561263357612633612402565b8054838255808410156126c05760017f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808316831461267457612674612593565b808616861461268557612685612593565b5060008360005260206000208360011b81018760011b820191505b808210156126bb5782825582848301556002820191506126a0565b505050505b5060008181526020812083915b8581101561120a576126df83836125c2565b60409290920191600291909101906001016126cd565b8135815560018101602083013561270b81611bb3565b67ffffffffffffffff8082167fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000084541617835561274b6040860186612182565b9350915061275d838360028701612479565b61276a6060860186612182565b9350915061277c838360038701612479565b608085013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18536030183126127b357600080fd5b9184019182359150808211156127c857600080fd5b506020820191508060061b36038213156127e157600080fd5b61137081836004860161261a565b60008251612178818460208701611c67565b60006020828403121561281357600080fd5b81516112b581611d3c565b60006020828403121561283057600080fd5b81516112b581611da7565b818103818111156108b6576108b6612593565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000818000a",
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

func (_CCIPReceiver *CCIPReceiverCaller) SChainConfigs(opts *bind.CallOpts, arg0 uint64) (SChainConfigs,

	error) {
	var out []interface{}
	err := _CCIPReceiver.contract.Call(opts, &out, "s_chainConfigs", arg0)

	outstruct := new(SChainConfigs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsDisabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Recipient = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.ExtraArgsBytes = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

func (_CCIPReceiver *CCIPReceiverSession) SChainConfigs(arg0 uint64) (SChainConfigs,

	error) {
	return _CCIPReceiver.Contract.SChainConfigs(&_CCIPReceiver.CallOpts, arg0)
}

func (_CCIPReceiver *CCIPReceiverCallerSession) SChainConfigs(arg0 uint64) (SChainConfigs,

	error) {
	return _CCIPReceiver.Contract.SChainConfigs(&_CCIPReceiver.CallOpts, arg0)
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

func (_CCIPReceiver *CCIPReceiverTransactor) RetryFailedMessage(opts *bind.TransactOpts, messageId [32]byte, forwardingAddress common.Address) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "retryFailedMessage", messageId, forwardingAddress)
}

func (_CCIPReceiver *CCIPReceiverSession) RetryFailedMessage(messageId [32]byte, forwardingAddress common.Address) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.RetryFailedMessage(&_CCIPReceiver.TransactOpts, messageId, forwardingAddress)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) RetryFailedMessage(messageId [32]byte, forwardingAddress common.Address) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.RetryFailedMessage(&_CCIPReceiver.TransactOpts, messageId, forwardingAddress)
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

func (_CCIPReceiver *CCIPReceiverTransactor) UpdateApprovedSenders(opts *bind.TransactOpts, adds []CCIPClientBaseapprovedSenderUpdate, removes []CCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error) {
	return _CCIPReceiver.contract.Transact(opts, "updateApprovedSenders", adds, removes)
}

func (_CCIPReceiver *CCIPReceiverSession) UpdateApprovedSenders(adds []CCIPClientBaseapprovedSenderUpdate, removes []CCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error) {
	return _CCIPReceiver.Contract.UpdateApprovedSenders(&_CCIPReceiver.TransactOpts, adds, removes)
}

func (_CCIPReceiver *CCIPReceiverTransactorSession) UpdateApprovedSenders(adds []CCIPClientBaseapprovedSenderUpdate, removes []CCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error) {
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

type SChainConfigs struct {
	IsDisabled     bool
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

	SChainConfigs(opts *bind.CallOpts, arg0 uint64) (SChainConfigs,

		error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	DisableChain(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error)

	EnableChain(opts *bind.TransactOpts, chainSelector uint64, recipient []byte, _extraArgsBytes []byte) (*types.Transaction, error)

	ProcessMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	RetryFailedMessage(opts *bind.TransactOpts, messageId [32]byte, forwardingAddress common.Address) (*types.Transaction, error)

	SetSimRevert(opts *bind.TransactOpts, simRevert bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateApprovedSenders(opts *bind.TransactOpts, adds []CCIPClientBaseapprovedSenderUpdate, removes []CCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error)

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
