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

type CCIPClientBaseapprovedSenderUpdate struct {
	DestChainSelector uint64
	Sender            []byte
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

var CCIPSenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientNativeFeeTokenAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"disableChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraArgsBytes\",\"type\":\"bytes\"}],\"name\":\"enableChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"senderAddr\",\"type\":\"bytes\"}],\"name\":\"isApprovedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"s_chainConfigs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgsBytes\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPClientBase.approvedSenderUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPClientBase.approvedSenderUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"}],\"name\":\"updateApprovedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNativeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620021a3380380620021a38339810160408190526200003491620001a8565b8033806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf81620000fd565b5050506001600160a01b038116620000ea576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b031660805250620001da565b336001600160a01b03821603620001575760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001bb57600080fd5b81516001600160a01b0381168114620001d357600080fd5b9392505050565b608051611f91620002126000396000818161028b01528181610b9401528181610c5d01528181610d310152610d6d0152611f916000f3fe6080604052600436106100d65760003560e01c806379ba50971161007f578063b0f479a111610059578063b0f479a11461027c578063d8469e40146102af578063effde240146102cf578063f2fde38b146102f0576100dd565b806379ba5097146101fb5780638462a2b9146102105780638da5cb5b14610230576100dd565b806341eade46116100b057806341eade461461019b578063536c6bfa146101bb5780635e35359e146101db576100dd565b80630e958d6b146100eb578063181f5a771461012057806335f170ef1461016c576100dd565b366100dd57005b3480156100e957600080fd5b005b3480156100f757600080fd5b5061010b6101063660046116ed565b610310565b60405190151581526020015b60405180910390f35b34801561012c57600080fd5b50604080518082018252601481527f4343495053656e64657220312e302e302d6465760000000000000000000000006020820152905161011791906117ae565b34801561017857600080fd5b5061018c6101873660046117c8565b61035b565b604051610117939291906117e3565b3480156101a757600080fd5b506100e96101b63660046117c8565b610492565b3480156101c757600080fd5b506100e96101d636600461183c565b6104dd565b3480156101e757600080fd5b506100e96101f6366004611873565b6104f3565b34801561020757600080fd5b506100e9610521565b34801561021c57600080fd5b506100e961022b3660046118f9565b610623565b34801561023c57600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610117565b34801561028857600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610257565b3480156102bb57600080fd5b506100e96102ca366004611965565b610804565b6102e26102dd3660046119e6565b610885565b604051908152602001610117565b3480156102fc57600080fd5b506100e961030b366004611aa7565b610e50565b67ffffffffffffffff8316600090815260026020526040808220905160039091019061033f9085908590611ac4565b9081526040519081900360200190205460ff1690509392505050565b6002602052600090815260409020805460018201805460ff909216929161038190611ad4565b80601f01602080910402602001604051908101604052809291908181526020018280546103ad90611ad4565b80156103fa5780601f106103cf576101008083540402835291602001916103fa565b820191906000526020600020905b8154815290600101906020018083116103dd57829003601f168201915b50505050509080600201805461040f90611ad4565b80601f016020809104026020016040519081016040528092919081815260200182805461043b90611ad4565b80156104885780601f1061045d57610100808354040283529160200191610488565b820191906000526020600020905b81548152906001019060200180831161046b57829003601f168201915b5050505050905083565b61049a610e64565b67ffffffffffffffff16600090815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b6104e5610e64565b6104ef8282610ee7565b5050565b6104fb610e64565b61051c73ffffffffffffffffffffffffffffffffffffffff84168383611041565b505050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146105a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61062b610e64565b60005b8181101561070e576002600084848481811061064c5761064c611b27565b905060200281019061065e9190611b56565b61066c9060208101906117c8565b67ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206003018383838181106106a3576106a3611b27565b90506020028101906106b59190611b56565b6106c3906020810190611b94565b6040516106d1929190611ac4565b90815260405190819003602001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905560010161062e565b5060005b838110156107fd5760016002600087878581811061073257610732611b27565b90506020028101906107449190611b56565b6107529060208101906117c8565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060030186868481811061078957610789611b27565b905060200281019061079b9190611b56565b6107a9906020810190611b94565b6040516107b7929190611ac4565b90815260405190819003602001902080549115157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909216919091179055600101610712565b5050505050565b61080c610e64565b67ffffffffffffffff8516600090815260026020526040902060018101610834858783611c70565b50811561084c576002810161084a838583611c70565b505b805460ff161561087d5780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681555b505050505050565b67ffffffffffffffff86166000908152600260205260408120600181018054899291906108b190611ad4565b159050806108c05750805460ff165b15610903576040517fd79f2ea400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8316600482015260240161059e565b6040805160a08101825267ffffffffffffffff8b1660009081526002602052918220600101805482919061093690611ad4565b80601f016020809104026020016040519081016040528092919081815260200182805461096290611ad4565b80156109af5780601f10610984576101008083540402835291602001916109af565b820191906000526020600020905b81548152906001019060200180831161099257829003601f168201915b5050505050815260200188888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525093855250506040805160208e8102820181019092528d815293810193928e92508d9182919085015b82821015610a3f57610a3060408302860136819003810190611d8a565b81526020019060010190610a13565b505050505081526020018673ffffffffffffffffffffffffffffffffffffffff168152602001600260008d67ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206002018054610a9a90611ad4565b80601f0160208091040260200160405190810160405280929190818152602001828054610ac690611ad4565b8015610b135780601f10610ae857610100808354040283529160200191610b13565b820191906000526020600020905b815481529060010190602001808311610af657829003601f168201915b5050505050815250905060005b88811015610c1c57610b8f33308c8c85818110610b3f57610b3f611b27565b905060400201602001358d8d86818110610b5b57610b5b611b27565b610b719260206040909202019081019150611aa7565b73ffffffffffffffffffffffffffffffffffffffff16929190611115565b610c147f00000000000000000000000000000000000000000000000000000000000000008b8b84818110610bc557610bc5611b27565b905060400201602001358c8c85818110610be157610be1611b27565b610bf79260206040909202019081019150611aa7565b73ffffffffffffffffffffffffffffffffffffffff169190611179565b600101610b20565b506040517f20487ded00000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906320487ded90610c94908e908690600401611de2565b602060405180830381865afa158015610cb1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd59190611ef7565b905073ffffffffffffffffffffffffffffffffffffffff861615610d5657610d1573ffffffffffffffffffffffffffffffffffffffff8716333084611115565b610d5673ffffffffffffffffffffffffffffffffffffffff87167f000000000000000000000000000000000000000000000000000000000000000083611179565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116906396f4e9f990881615610da3576000610da5565b825b8d856040518463ffffffff1660e01b8152600401610dc4929190611de2565b60206040518083038185885af1158015610de2573d6000803e3d6000fd5b50505050506040513d601f19601f82011682018060405250810190610e079190611ef7565b94507f54791b38f3859327992a1ca0590ad3c0f08feba98d1a4f56ab0dca74d203392a85604051610e3a91815260200190565b60405180910390a1505050509695505050505050565b610e58610e64565b610e6181611277565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610ee5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161059e565b565b80471015610f51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604482015260640161059e565b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114610fab576040519150601f19603f3d011682016040523d82523d6000602084013e610fb0565b606091505b505090508061051c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d61792068617665207265766572746564000000000000606482015260840161059e565b60405173ffffffffffffffffffffffffffffffffffffffff831660248201526044810182905261051c9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261136c565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526111739085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611093565b50505050565b6040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa1580156111f0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112149190611ef7565b61121e9190611f10565b60405173ffffffffffffffffffffffffffffffffffffffff85166024820152604481018290529091506111739085907f095ea7b30000000000000000000000000000000000000000000000000000000090606401611093565b3373ffffffffffffffffffffffffffffffffffffffff8216036112f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161059e565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006113ce826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166114789092919063ffffffff16565b80519091501561051c57808060200190518101906113ec9190611f50565b61051c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161059e565b6060611487848460008561148f565b949350505050565b606082471015611521576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161059e565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161154a9190611f72565b60006040518083038185875af1925050503d8060008114611587576040519150601f19603f3d011682016040523d82523d6000602084013e61158c565b606091505b509150915061159d878383876115a8565b979650505050505050565b6060831561163e5782516000036116375773ffffffffffffffffffffffffffffffffffffffff85163b611637576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161059e565b5081611487565b61148783838151156116535781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059e91906117ae565b803567ffffffffffffffff8116811461169f57600080fd5b919050565b60008083601f8401126116b657600080fd5b50813567ffffffffffffffff8111156116ce57600080fd5b6020830191508360208285010111156116e657600080fd5b9250929050565b60008060006040848603121561170257600080fd5b61170b84611687565b9250602084013567ffffffffffffffff81111561172757600080fd5b611733868287016116a4565b9497909650939450505050565b60005b8381101561175b578181015183820152602001611743565b50506000910152565b6000815180845261177c816020860160208601611740565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006117c16020830184611764565b9392505050565b6000602082840312156117da57600080fd5b6117c182611687565b83151581526060602082015260006117fe6060830185611764565b82810360408401526118108185611764565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610e6157600080fd5b6000806040838503121561184f57600080fd5b823561185a8161181a565b946020939093013593505050565b803561169f8161181a565b60008060006060848603121561188857600080fd5b83356118938161181a565b925060208401356118a38161181a565b929592945050506040919091013590565b60008083601f8401126118c657600080fd5b50813567ffffffffffffffff8111156118de57600080fd5b6020830191508360208260051b85010111156116e657600080fd5b6000806000806040858703121561190f57600080fd5b843567ffffffffffffffff8082111561192757600080fd5b611933888389016118b4565b9096509450602087013591508082111561194c57600080fd5b50611959878288016118b4565b95989497509550505050565b60008060008060006060868803121561197d57600080fd5b61198686611687565b9450602086013567ffffffffffffffff808211156119a357600080fd5b6119af89838a016116a4565b909650945060408801359150808211156119c857600080fd5b506119d5888289016116a4565b969995985093965092949392505050565b600080600080600080608087890312156119ff57600080fd5b611a0887611687565b9550602087013567ffffffffffffffff80821115611a2557600080fd5b818901915089601f830112611a3957600080fd5b813581811115611a4857600080fd5b8a60208260061b8501011115611a5d57600080fd5b602083019750809650506040890135915080821115611a7b57600080fd5b50611a8889828a016116a4565b9094509250611a9b905060608801611868565b90509295509295509295565b600060208284031215611ab957600080fd5b81356117c18161181a565b8183823760009101908152919050565b600181811c90821680611ae857607f821691505b602082108103611b21577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc1833603018112611b8a57600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611bc957600080fd5b83018035915067ffffffffffffffff821115611be457600080fd5b6020019150368190038213156116e657600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82111561051c576000816000526020600020601f850160051c81016020861015611c515750805b601f850160051c820191505b8181101561087d57828155600101611c5d565b67ffffffffffffffff831115611c8857611c88611bf9565b611c9c83611c968354611ad4565b83611c28565b6000601f841160018114611cee5760008515611cb85750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b1783556107fd565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b82811015611d3d5786850135825560209485019460019092019101611d1d565b5086821015611d78577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b600060408284031215611d9c57600080fd5b6040516040810181811067ffffffffffffffff82111715611dbf57611dbf611bf9565b6040528235611dcd8161181a565b81526020928301359281019290925250919050565b6000604067ffffffffffffffff851683526020604081850152845160a06040860152611e1160e0860182611764565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080878403016060880152611e4c8383611764565b6040890151888203830160808a01528051808352908601945060009350908501905b80841015611ead578451805173ffffffffffffffffffffffffffffffffffffffff16835286015186830152938501936001939093019290860190611e6e565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a01529550611ee98187611764565b9a9950505050505050505050565b600060208284031215611f0957600080fd5b5051919050565b80820180821115611f4a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b600060208284031215611f6257600080fd5b815180151581146117c157600080fd5b60008251611b8a81846020870161174056fea164736f6c6343000818000a",
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

func (_CCIPSender *CCIPSenderCaller) SChainConfigs(opts *bind.CallOpts, arg0 uint64) (SChainConfigs,

	error) {
	var out []interface{}
	err := _CCIPSender.contract.Call(opts, &out, "s_chainConfigs", arg0)

	outstruct := new(SChainConfigs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsDisabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Recipient = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.ExtraArgsBytes = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

func (_CCIPSender *CCIPSenderSession) SChainConfigs(arg0 uint64) (SChainConfigs,

	error) {
	return _CCIPSender.Contract.SChainConfigs(&_CCIPSender.CallOpts, arg0)
}

func (_CCIPSender *CCIPSenderCallerSession) SChainConfigs(arg0 uint64) (SChainConfigs,

	error) {
	return _CCIPSender.Contract.SChainConfigs(&_CCIPSender.CallOpts, arg0)
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

func (_CCIPSender *CCIPSenderTransactor) UpdateApprovedSenders(opts *bind.TransactOpts, adds []CCIPClientBaseapprovedSenderUpdate, removes []CCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error) {
	return _CCIPSender.contract.Transact(opts, "updateApprovedSenders", adds, removes)
}

func (_CCIPSender *CCIPSenderSession) UpdateApprovedSenders(adds []CCIPClientBaseapprovedSenderUpdate, removes []CCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error) {
	return _CCIPSender.Contract.UpdateApprovedSenders(&_CCIPSender.TransactOpts, adds, removes)
}

func (_CCIPSender *CCIPSenderTransactorSession) UpdateApprovedSenders(adds []CCIPClientBaseapprovedSenderUpdate, removes []CCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error) {
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

type SChainConfigs struct {
	IsDisabled     bool
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

	SChainConfigs(opts *bind.CallOpts, arg0 uint64) (SChainConfigs,

		error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destChainSelector uint64, tokenAmounts []ClientEVMTokenAmount, data []byte, feeToken common.Address) (*types.Transaction, error)

	DisableChain(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error)

	EnableChain(opts *bind.TransactOpts, chainSelector uint64, recipient []byte, _extraArgsBytes []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateApprovedSenders(opts *bind.TransactOpts, adds []CCIPClientBaseapprovedSenderUpdate, removes []CCIPClientBaseapprovedSenderUpdate) (*types.Transaction, error)

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
