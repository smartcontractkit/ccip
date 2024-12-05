// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lbtc_adapter

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

var CLAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"contractLombardTokenPool\",\"name\":\"tokenPool_\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"executionGasLimit_\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"ccipRouter_\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowlist_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"rmnProxy_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"attestationEnable_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Adapter_ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CLAttemptToOverrideChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CLAttemptToOverrideChainSelector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CLRefundFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CLUnauthorizedTokenPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CLZeroChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CLZeroChanSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBridge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"NotEnoughToPayFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroPayload\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBridge\",\"name\":\"oldBridge\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIBridge\",\"name\":\"newBridge\",\"type\":\"address\"}],\"name\":\"BridgeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"CLChainSelectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CLTokenPoolDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"prevVal\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"newVal\",\"type\":\"uint128\"}],\"name\":\"ExecutionGasLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"}],\"name\":\"changeBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_toChain\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutionGasLimit\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_toChain\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getRemoteChainSelector\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"onChainData\",\"type\":\"bytes\"}],\"name\":\"initWithdrawalNoSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"initiateDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastBurnedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"lastPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offChainData\",\"type\":\"bytes\"}],\"name\":\"initiateWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lbtc\",\"outputs\":[{\"internalType\":\"contractILBTC\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newVal\",\"type\":\"uint128\"}],\"name\":\"setExecutionGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chain\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"setRemoteChainSelector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPool\",\"outputs\":[{\"internalType\":\"contractLombardTokenPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200256c3803806200256c8339810160408190526200003491620001ca565b600080546001600160a01b0319166001600160a01b038916179055620000616200005b3390565b620000ca565b6200006c856200011c565b600580546001600160a01b0319166001600160a01b0388169081179091556040519081527f45163ba2f75e282ba1000a5e166237acb454fa6b2b4cc778671bf977a01b13b49060200160405180910390a15050505050505062000326565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6004546040516001600160801b038084169216907f45a4a024e3d155ab5b1634193775d00d4cbf4ef295e76d0bf1aa6dfeb114dd8490600090a3600480546001600160801b0319166001600160801b0392909216919091179055565b6001600160a01b03811681146200018e57600080fd5b50565b80516200019e8162000178565b919050565b634e487b7160e01b600052604160045260246000fd5b805180151581146200019e57600080fd5b600080600080600080600060e0888a031215620001e657600080fd5b8751620001f38162000178565b80975050602080890151620002088162000178565b60408a01519097506001600160801b03811681146200022657600080fd5b60608a0151909650620002398162000178565b60808a01519095506001600160401b03808211156200025757600080fd5b818b0191508b601f8301126200026c57600080fd5b815181811115620002815762000281620001a3565b8060051b604051601f19603f83011681018181108582111715620002a957620002a9620001a3565b60405291825284820192508381018501918e831115620002c857600080fd5b938501935b82851015620002f157620002e18562000191565b84529385019392850192620002cd565b8098505050505050506200030860a0890162000191565b91506200031860c08901620001b9565b905092959891949750929550565b61223680620003366000396000f3fe60806040526004361061010e5760003560e01c8063715886c4116100a5578063a1a6d50811610074578063e240550d11610059578063e240550d14610377578063e78cea9214610397578063f2fde38b146103c457600080fd5b8063a1a6d5081461031c578063b1ab7e221461035757600080fd5b8063715886c414610244578063775710d4146102575780638da5cb5b146102a757806392b0680a146102d257600080fd5b80635391a405116100e15780635391a405146101cc578063550e7ab2146101ec57806368b3c9101461021a578063715018a61461022f57600080fd5b806308774410146101135780630c373d7414610135578063104e992914610155578063453bb95c146101ac575b600080fd5b34801561011f57600080fd5b5061013361012e36600461189c565b6103e4565b005b34801561014157600080fd5b50610133610150366004611926565b61046a565b34801561016157600080fd5b506005546101829073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101b857600080fd5b506101336101c7366004611979565b610568565b3480156101d857600080fd5b506101336101e7366004611926565b61071b565b3480156101f857600080fd5b5061020c6102073660046119a5565b610874565b6040516101a3929190611a6d565b34801561022657600080fd5b50610182610c1b565b34801561023b57600080fd5b50610133610cb2565b610133610252366004611ba8565b610cc6565b34801561026357600080fd5b5061028e610272366004611c1d565b60026020526000908152604090205467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016101a3565b3480156102b357600080fd5b5060015473ffffffffffffffffffffffffffffffffffffffff16610182565b3480156102de57600080fd5b506004546102fb906fffffffffffffffffffffffffffffffff1681565b6040516fffffffffffffffffffffffffffffffff90911681526020016101a3565b34801561032857600080fd5b50610349610337366004611c36565b60036020526000908152604090205481565b6040519081526020016101a3565b34801561036357600080fd5b50610349610372366004611c51565b6110ac565b34801561038357600080fd5b50610133610392366004611cb5565b61120b565b3480156103a357600080fd5b506000546101829073ffffffffffffffffffffffffffffffffffffffff1681565b3480156103d057600080fd5b506101336103df36600461189c565b61121f565b6103ec6112d3565b6103f5816112db565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917fd565484d693f5157abcceb853139678038bc740991b0a4dc3baa2426325bb3c09190a35050565b610472611328565b6104d9600360008567ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000205483838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061137b92505050565b6000546040517f0968f26400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690630968f264906105319085908590600401611ce7565b600060405180830381600087803b15801561054b57600080fd5b505af115801561055f573d6000803e3d6000fd5b50505050505050565b610570611401565b816105a7576040517fe531cb7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff166000036105eb576040517fe531cb7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281526002602052604090205467ffffffffffffffff161561063b576040517fb72f782a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff81166000908152600360205260409020541561068c576040517fe95ca6dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082815260026020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff861690811790915580845260038352928190208590558051858152918201929092527f61805347ac73c642ed5bfbc228e1702a4d51591b36a35387250de0c0d75b0072910160405180910390a15050565b610723611328565b60008061073283850185611d34565b67ffffffffffffffff8716600090815260036020526040902054919350915061075b908361137b565b6000546040517f6b93aa5100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690636b93aa51906107b39085908590600401611d98565b600060405180830381600087803b1580156107cd57600080fd5b505af11580156107e1573d6000803e3d6000fd5b50506000546040517f0968f26400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169250630968f264915061083b908590600401611dc6565b600060405180830381600087803b15801561085557600080fd5b505af1158015610869573d6000803e3d6000fd5b505050505050505050565b60006060610880611328565b60006007805461088f90611dd9565b90501115610976576006549150600780546108a990611dd9565b80601f01602080910402602001604051908101604052809291908181526020018280546108d590611dd9565b80156109225780601f106108f757610100808354040283529160200191610922565b820191906000526020600020905b81548152906001019060200180831161090557829003601f168201915b509394506000935061093392505050565b6040519080825280601f01601f19166020018201604052801561095d576020820181803683370190505b5060079061096b9082611e79565b506000600655610b17565b61097e610c1b565b6000546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526024810186905291169063095ea7b3906044016020604051808303816000875af11580156109f6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1a9190611f93565b506000805467ffffffffffffffff88168252600360205260409091205473ffffffffffffffffffffffffffffffffffffffff909116906383031c4190610a608789611fb5565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526004810192909252602482015267ffffffffffffffff861660448201526064016000604051808303816000875af1158015610acb573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610b119190810190611ff2565b90925090505b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166368b3c9106040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba69190612075565b73ffffffffffffffffffffffffffffffffffffffff166342966c68836040518263ffffffff1660e01b8152600401610be091815260200190565b600060405180830381600087803b158015610bfa57600080fd5b505af1158015610c0e573d6000803e3d6000fd5b5050505094509492505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166368b3c9106040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cad9190612075565b905090565b610cba611401565b610cc46000611482565b565b73ffffffffffffffffffffffffffffffffffffffff861630146110a45760068290556007610cf48282611e79565b50600085815260026020908152604080832054815192830187905267ffffffffffffffff169291610d36910160405160208183030381529060405285856114f9565b90506000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b0f479a16040518163ffffffff1660e01b8152600401602060405180830381865afa158015610da7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dcb9190612075565b905060008173ffffffffffffffffffffffffffffffffffffffff166320487ded85856040518363ffffffff1660e01b8152600401610e0a929190612092565b602060405180830381865afa158015610e27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e4b91906121a7565b905080341015610e8f576040517faf09dec9000000000000000000000000000000000000000000000000000000008152600481018290526024015b60405180910390fd5b80341115610f61576000610ea382346121c0565b905060008b73ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114610eff576040519150601f19603f3d011682016040523d82523d6000602084013e610f04565b606091505b5050905080610f5e576040517f9bb3185e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d16600482015260248101839052604401610e86565b50505b610f69610c1b565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015260248201899052919091169063095ea7b3906044016020604051808303816000875af1158015610fe0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110049190611f93565b506040517f96f4e9f900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316906396f4e9f990839061105b9088908890600401612092565b60206040518083038185885af1158015611079573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061109e91906121a7565b50505050505b505050505050565b600554604080517fb0f479a1000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163b0f479a19160048083019260209291908290030181865afa15801561111c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111409190612075565b60008781526002602090815260409182902054825191820188905273ffffffffffffffffffffffffffffffffffffffff93909316926320487ded9267ffffffffffffffff909116916111a3910160405160208183030381529060405287876114f9565b6040518363ffffffff1660e01b81526004016111c0929190612092565b602060405180830381865afa1580156111dd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061120191906121a7565b9695505050505050565b611213611401565b61121c816117f4565b50565b611227611401565b73ffffffffffffffffffffffffffffffffffffffff81166112ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610e86565b61121c81611482565b610cc4611401565b73ffffffffffffffffffffffffffffffffffffffff811661121c576040517fcf891a8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60055473ffffffffffffffffffffffffffffffffffffffff163314610cc4576040517f36b0b92e000000000000000000000000000000000000000000000000000000008152336004820152602401610e86565b6000546040517f8783f52500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690638783f525906113d39085908590600401611a6d565b600060405180830381600087803b1580156113ed57600080fd5b505af11580156110a4573d6000803e3d6000fd5b60015473ffffffffffffffffffffffffffffffffffffffff163314610cc4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610e86565b6001805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6115416040518060a00160405280606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b604080516001808252818301909252600091816020015b6040805180820190915260008082526020820152815260200190600190039081611558575050604080518082018083526000547f68b3c91000000000000000000000000000000000000000000000000000000000909152915192935091829173ffffffffffffffffffffffffffffffffffffffff16906368b3c910906044808501916020918187030181865afa1580156115f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061161a9190612075565b73ffffffffffffffffffffffffffffffffffffffff168152602001858152508160008151811061164c5761164c6121fa565b60200260200101819052506060600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166355b961566040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116ea9190611f93565b61172c578351600003611729576040517ff3cc293900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50825b6040805160a081018252878152602080820184905281830185905260006060830152825180840184526004546fffffffffffffffffffffffffffffffff1680825260019183019182528451602481019190915290511515604480830191909152845180830390910181526064909101909352820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f181dcf10000000000000000000000000000000000000000000000000000000001790529060808201529695505050505050565b6004546040516fffffffffffffffffffffffffffffffff8084169216907f45a4a024e3d155ab5b1634193775d00d4cbf4ef295e76d0bf1aa6dfeb114dd8490600090a3600480547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff92909216919091179055565b73ffffffffffffffffffffffffffffffffffffffff8116811461121c57600080fd5b6000602082840312156118ae57600080fd5b81356118b98161187a565b9392505050565b803567ffffffffffffffff811681146118d857600080fd5b919050565b60008083601f8401126118ef57600080fd5b50813567ffffffffffffffff81111561190757600080fd5b60208301915083602082850101111561191f57600080fd5b9250929050565b60008060006040848603121561193b57600080fd5b611944846118c0565b9250602084013567ffffffffffffffff81111561196057600080fd5b61196c868287016118dd565b9497909650939450505050565b6000806040838503121561198c57600080fd5b8235915061199c602084016118c0565b90509250929050565b600080600080606085870312156119bb57600080fd5b6119c4856118c0565b9350602085013567ffffffffffffffff8111156119e057600080fd5b6119ec878288016118dd565b9598909750949560400135949350505050565b60005b83811015611a1a578181015183820152602001611a02565b50506000910152565b60008151808452611a3b8160208601602086016119ff565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b828152604060208201526000611a866040830184611a23565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611b0457611b04611a8e565b604052919050565b600067ffffffffffffffff821115611b2657611b26611a8e565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112611b6357600080fd5b8135611b76611b7182611b0c565b611abd565b818152846020838601011115611b8b57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c08789031215611bc157600080fd5b8635611bcc8161187a565b95506020870135945060408701359350606087013592506080870135915060a087013567ffffffffffffffff811115611c0457600080fd5b611c1089828a01611b52565b9150509295509295509295565b600060208284031215611c2f57600080fd5b5035919050565b600060208284031215611c4857600080fd5b6118b9826118c0565b600080600080600060a08688031215611c6957600080fd5b85359450602086013593506040860135925060608601359150608086013567ffffffffffffffff811115611c9c57600080fd5b611ca888828901611b52565b9150509295509295909350565b600060208284031215611cc757600080fd5b81356fffffffffffffffffffffffffffffffff811681146118b957600080fd5b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b60008060408385031215611d4757600080fd5b823567ffffffffffffffff80821115611d5f57600080fd5b611d6b86838701611b52565b93506020850135915080821115611d8157600080fd5b50611d8e85828601611b52565b9150509250929050565b604081526000611dab6040830185611a23565b8281036020840152611dbd8185611a23565b95945050505050565b6020815260006118b96020830184611a23565b600181811c90821680611ded57607f821691505b602082108103611e26577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f821115611e74576000816000526020600020601f850160051c81016020861015611e555750805b601f850160051c820191505b818110156110a457828155600101611e61565b505050565b815167ffffffffffffffff811115611e9357611e93611a8e565b611ea781611ea18454611dd9565b84611e2c565b602080601f831160018114611efa5760008415611ec45750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556110a4565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611f4757888601518255948401946001909101908401611f28565b5085821015611f8357878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b600060208284031215611fa557600080fd5b815180151581146118b957600080fd5b80356020831015611fec577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602084900360031b1b165b92915050565b6000806040838503121561200557600080fd5b82519150602083015167ffffffffffffffff81111561202357600080fd5b8301601f8101851361203457600080fd5b8051612042611b7182611b0c565b81815286602083850101111561205757600080fd5b6120688260208301602086016119ff565b8093505050509250929050565b60006020828403121561208757600080fd5b81516118b98161187a565b6000604067ffffffffffffffff851683526020604081850152845160a060408601526120c160e0860182611a23565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808784030160608801526120fc8383611a23565b6040890151888203830160808a01528051808352908601945060009350908501905b8084101561215d578451805173ffffffffffffffffffffffffffffffffffffffff1683528601518683015293850193600193909301929086019061211e565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a015295506121998187611a23565b9a9950505050505050505050565b6000602082840312156121b957600080fd5b5051919050565b81810381811115611fec577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c6343000818000a",
}

var CLAdapterABI = CLAdapterMetaData.ABI

var CLAdapterBin = CLAdapterMetaData.Bin

func DeployCLAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, bridge_ common.Address, tokenPool_ common.Address, executionGasLimit_ *big.Int, ccipRouter_ common.Address, allowlist_ []common.Address, rmnProxy_ common.Address, attestationEnable_ bool) (common.Address, *types.Transaction, *CLAdapter, error) {
	parsed, err := CLAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CLAdapterBin), backend, bridge_, tokenPool_, executionGasLimit_, ccipRouter_, allowlist_, rmnProxy_, attestationEnable_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CLAdapter{address: address, abi: *parsed, CLAdapterCaller: CLAdapterCaller{contract: contract}, CLAdapterTransactor: CLAdapterTransactor{contract: contract}, CLAdapterFilterer: CLAdapterFilterer{contract: contract}}, nil
}

type CLAdapter struct {
	address common.Address
	abi     abi.ABI
	CLAdapterCaller
	CLAdapterTransactor
	CLAdapterFilterer
}

type CLAdapterCaller struct {
	contract *bind.BoundContract
}

type CLAdapterTransactor struct {
	contract *bind.BoundContract
}

type CLAdapterFilterer struct {
	contract *bind.BoundContract
}

type CLAdapterSession struct {
	Contract     *CLAdapter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CLAdapterCallerSession struct {
	Contract *CLAdapterCaller
	CallOpts bind.CallOpts
}

type CLAdapterTransactorSession struct {
	Contract     *CLAdapterTransactor
	TransactOpts bind.TransactOpts
}

type CLAdapterRaw struct {
	Contract *CLAdapter
}

type CLAdapterCallerRaw struct {
	Contract *CLAdapterCaller
}

type CLAdapterTransactorRaw struct {
	Contract *CLAdapterTransactor
}

func NewCLAdapter(address common.Address, backend bind.ContractBackend) (*CLAdapter, error) {
	abi, err := abi.JSON(strings.NewReader(CLAdapterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCLAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CLAdapter{address: address, abi: abi, CLAdapterCaller: CLAdapterCaller{contract: contract}, CLAdapterTransactor: CLAdapterTransactor{contract: contract}, CLAdapterFilterer: CLAdapterFilterer{contract: contract}}, nil
}

func NewCLAdapterCaller(address common.Address, caller bind.ContractCaller) (*CLAdapterCaller, error) {
	contract, err := bindCLAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CLAdapterCaller{contract: contract}, nil
}

func NewCLAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*CLAdapterTransactor, error) {
	contract, err := bindCLAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CLAdapterTransactor{contract: contract}, nil
}

func NewCLAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*CLAdapterFilterer, error) {
	contract, err := bindCLAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CLAdapterFilterer{contract: contract}, nil
}

func bindCLAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CLAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CLAdapter *CLAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CLAdapter.Contract.CLAdapterCaller.contract.Call(opts, result, method, params...)
}

func (_CLAdapter *CLAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CLAdapter.Contract.CLAdapterTransactor.contract.Transfer(opts)
}

func (_CLAdapter *CLAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CLAdapter.Contract.CLAdapterTransactor.contract.Transact(opts, method, params...)
}

func (_CLAdapter *CLAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CLAdapter.Contract.contract.Call(opts, result, method, params...)
}

func (_CLAdapter *CLAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CLAdapter.Contract.contract.Transfer(opts)
}

func (_CLAdapter *CLAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CLAdapter.Contract.contract.Transact(opts, method, params...)
}

func (_CLAdapter *CLAdapterCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CLAdapter.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CLAdapter *CLAdapterSession) Bridge() (common.Address, error) {
	return _CLAdapter.Contract.Bridge(&_CLAdapter.CallOpts)
}

func (_CLAdapter *CLAdapterCallerSession) Bridge() (common.Address, error) {
	return _CLAdapter.Contract.Bridge(&_CLAdapter.CallOpts)
}

func (_CLAdapter *CLAdapterCaller) GetChain(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _CLAdapter.contract.Call(opts, &out, "getChain", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_CLAdapter *CLAdapterSession) GetChain(arg0 uint64) ([32]byte, error) {
	return _CLAdapter.Contract.GetChain(&_CLAdapter.CallOpts, arg0)
}

func (_CLAdapter *CLAdapterCallerSession) GetChain(arg0 uint64) ([32]byte, error) {
	return _CLAdapter.Contract.GetChain(&_CLAdapter.CallOpts, arg0)
}

func (_CLAdapter *CLAdapterCaller) GetExecutionGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CLAdapter.contract.Call(opts, &out, "getExecutionGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CLAdapter *CLAdapterSession) GetExecutionGasLimit() (*big.Int, error) {
	return _CLAdapter.Contract.GetExecutionGasLimit(&_CLAdapter.CallOpts)
}

func (_CLAdapter *CLAdapterCallerSession) GetExecutionGasLimit() (*big.Int, error) {
	return _CLAdapter.Contract.GetExecutionGasLimit(&_CLAdapter.CallOpts)
}

func (_CLAdapter *CLAdapterCaller) GetFee(opts *bind.CallOpts, _toChain [32]byte, arg1 [32]byte, _toAddress [32]byte, _amount *big.Int, _payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _CLAdapter.contract.Call(opts, &out, "getFee", _toChain, arg1, _toAddress, _amount, _payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CLAdapter *CLAdapterSession) GetFee(_toChain [32]byte, arg1 [32]byte, _toAddress [32]byte, _amount *big.Int, _payload []byte) (*big.Int, error) {
	return _CLAdapter.Contract.GetFee(&_CLAdapter.CallOpts, _toChain, arg1, _toAddress, _amount, _payload)
}

func (_CLAdapter *CLAdapterCallerSession) GetFee(_toChain [32]byte, arg1 [32]byte, _toAddress [32]byte, _amount *big.Int, _payload []byte) (*big.Int, error) {
	return _CLAdapter.Contract.GetFee(&_CLAdapter.CallOpts, _toChain, arg1, _toAddress, _amount, _payload)
}

func (_CLAdapter *CLAdapterCaller) GetRemoteChainSelector(opts *bind.CallOpts, arg0 [32]byte) (uint64, error) {
	var out []interface{}
	err := _CLAdapter.contract.Call(opts, &out, "getRemoteChainSelector", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_CLAdapter *CLAdapterSession) GetRemoteChainSelector(arg0 [32]byte) (uint64, error) {
	return _CLAdapter.Contract.GetRemoteChainSelector(&_CLAdapter.CallOpts, arg0)
}

func (_CLAdapter *CLAdapterCallerSession) GetRemoteChainSelector(arg0 [32]byte) (uint64, error) {
	return _CLAdapter.Contract.GetRemoteChainSelector(&_CLAdapter.CallOpts, arg0)
}

func (_CLAdapter *CLAdapterCaller) Lbtc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CLAdapter.contract.Call(opts, &out, "lbtc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CLAdapter *CLAdapterSession) Lbtc() (common.Address, error) {
	return _CLAdapter.Contract.Lbtc(&_CLAdapter.CallOpts)
}

func (_CLAdapter *CLAdapterCallerSession) Lbtc() (common.Address, error) {
	return _CLAdapter.Contract.Lbtc(&_CLAdapter.CallOpts)
}

func (_CLAdapter *CLAdapterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CLAdapter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CLAdapter *CLAdapterSession) Owner() (common.Address, error) {
	return _CLAdapter.Contract.Owner(&_CLAdapter.CallOpts)
}

func (_CLAdapter *CLAdapterCallerSession) Owner() (common.Address, error) {
	return _CLAdapter.Contract.Owner(&_CLAdapter.CallOpts)
}

func (_CLAdapter *CLAdapterCaller) TokenPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CLAdapter.contract.Call(opts, &out, "tokenPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CLAdapter *CLAdapterSession) TokenPool() (common.Address, error) {
	return _CLAdapter.Contract.TokenPool(&_CLAdapter.CallOpts)
}

func (_CLAdapter *CLAdapterCallerSession) TokenPool() (common.Address, error) {
	return _CLAdapter.Contract.TokenPool(&_CLAdapter.CallOpts)
}

func (_CLAdapter *CLAdapterTransactor) ChangeBridge(opts *bind.TransactOpts, bridge_ common.Address) (*types.Transaction, error) {
	return _CLAdapter.contract.Transact(opts, "changeBridge", bridge_)
}

func (_CLAdapter *CLAdapterSession) ChangeBridge(bridge_ common.Address) (*types.Transaction, error) {
	return _CLAdapter.Contract.ChangeBridge(&_CLAdapter.TransactOpts, bridge_)
}

func (_CLAdapter *CLAdapterTransactorSession) ChangeBridge(bridge_ common.Address) (*types.Transaction, error) {
	return _CLAdapter.Contract.ChangeBridge(&_CLAdapter.TransactOpts, bridge_)
}

func (_CLAdapter *CLAdapterTransactor) Deposit(opts *bind.TransactOpts, fromAddress common.Address, _toChain [32]byte, arg2 [32]byte, _toAddress [32]byte, _amount *big.Int, _payload []byte) (*types.Transaction, error) {
	return _CLAdapter.contract.Transact(opts, "deposit", fromAddress, _toChain, arg2, _toAddress, _amount, _payload)
}

func (_CLAdapter *CLAdapterSession) Deposit(fromAddress common.Address, _toChain [32]byte, arg2 [32]byte, _toAddress [32]byte, _amount *big.Int, _payload []byte) (*types.Transaction, error) {
	return _CLAdapter.Contract.Deposit(&_CLAdapter.TransactOpts, fromAddress, _toChain, arg2, _toAddress, _amount, _payload)
}

func (_CLAdapter *CLAdapterTransactorSession) Deposit(fromAddress common.Address, _toChain [32]byte, arg2 [32]byte, _toAddress [32]byte, _amount *big.Int, _payload []byte) (*types.Transaction, error) {
	return _CLAdapter.Contract.Deposit(&_CLAdapter.TransactOpts, fromAddress, _toChain, arg2, _toAddress, _amount, _payload)
}

func (_CLAdapter *CLAdapterTransactor) InitWithdrawalNoSignatures(opts *bind.TransactOpts, remoteSelector uint64, onChainData []byte) (*types.Transaction, error) {
	return _CLAdapter.contract.Transact(opts, "initWithdrawalNoSignatures", remoteSelector, onChainData)
}

func (_CLAdapter *CLAdapterSession) InitWithdrawalNoSignatures(remoteSelector uint64, onChainData []byte) (*types.Transaction, error) {
	return _CLAdapter.Contract.InitWithdrawalNoSignatures(&_CLAdapter.TransactOpts, remoteSelector, onChainData)
}

func (_CLAdapter *CLAdapterTransactorSession) InitWithdrawalNoSignatures(remoteSelector uint64, onChainData []byte) (*types.Transaction, error) {
	return _CLAdapter.Contract.InitWithdrawalNoSignatures(&_CLAdapter.TransactOpts, remoteSelector, onChainData)
}

func (_CLAdapter *CLAdapterTransactor) InitiateDeposit(opts *bind.TransactOpts, remoteChainSelector uint64, receiver []byte, amount *big.Int) (*types.Transaction, error) {
	return _CLAdapter.contract.Transact(opts, "initiateDeposit", remoteChainSelector, receiver, amount)
}

func (_CLAdapter *CLAdapterSession) InitiateDeposit(remoteChainSelector uint64, receiver []byte, amount *big.Int) (*types.Transaction, error) {
	return _CLAdapter.Contract.InitiateDeposit(&_CLAdapter.TransactOpts, remoteChainSelector, receiver, amount)
}

func (_CLAdapter *CLAdapterTransactorSession) InitiateDeposit(remoteChainSelector uint64, receiver []byte, amount *big.Int) (*types.Transaction, error) {
	return _CLAdapter.Contract.InitiateDeposit(&_CLAdapter.TransactOpts, remoteChainSelector, receiver, amount)
}

func (_CLAdapter *CLAdapterTransactor) InitiateWithdrawal(opts *bind.TransactOpts, remoteSelector uint64, offChainData []byte) (*types.Transaction, error) {
	return _CLAdapter.contract.Transact(opts, "initiateWithdrawal", remoteSelector, offChainData)
}

func (_CLAdapter *CLAdapterSession) InitiateWithdrawal(remoteSelector uint64, offChainData []byte) (*types.Transaction, error) {
	return _CLAdapter.Contract.InitiateWithdrawal(&_CLAdapter.TransactOpts, remoteSelector, offChainData)
}

func (_CLAdapter *CLAdapterTransactorSession) InitiateWithdrawal(remoteSelector uint64, offChainData []byte) (*types.Transaction, error) {
	return _CLAdapter.Contract.InitiateWithdrawal(&_CLAdapter.TransactOpts, remoteSelector, offChainData)
}

func (_CLAdapter *CLAdapterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CLAdapter.contract.Transact(opts, "renounceOwnership")
}

func (_CLAdapter *CLAdapterSession) RenounceOwnership() (*types.Transaction, error) {
	return _CLAdapter.Contract.RenounceOwnership(&_CLAdapter.TransactOpts)
}

func (_CLAdapter *CLAdapterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CLAdapter.Contract.RenounceOwnership(&_CLAdapter.TransactOpts)
}

func (_CLAdapter *CLAdapterTransactor) SetExecutionGasLimit(opts *bind.TransactOpts, newVal *big.Int) (*types.Transaction, error) {
	return _CLAdapter.contract.Transact(opts, "setExecutionGasLimit", newVal)
}

func (_CLAdapter *CLAdapterSession) SetExecutionGasLimit(newVal *big.Int) (*types.Transaction, error) {
	return _CLAdapter.Contract.SetExecutionGasLimit(&_CLAdapter.TransactOpts, newVal)
}

func (_CLAdapter *CLAdapterTransactorSession) SetExecutionGasLimit(newVal *big.Int) (*types.Transaction, error) {
	return _CLAdapter.Contract.SetExecutionGasLimit(&_CLAdapter.TransactOpts, newVal)
}

func (_CLAdapter *CLAdapterTransactor) SetRemoteChainSelector(opts *bind.TransactOpts, chain [32]byte, chainSelector uint64) (*types.Transaction, error) {
	return _CLAdapter.contract.Transact(opts, "setRemoteChainSelector", chain, chainSelector)
}

func (_CLAdapter *CLAdapterSession) SetRemoteChainSelector(chain [32]byte, chainSelector uint64) (*types.Transaction, error) {
	return _CLAdapter.Contract.SetRemoteChainSelector(&_CLAdapter.TransactOpts, chain, chainSelector)
}

func (_CLAdapter *CLAdapterTransactorSession) SetRemoteChainSelector(chain [32]byte, chainSelector uint64) (*types.Transaction, error) {
	return _CLAdapter.Contract.SetRemoteChainSelector(&_CLAdapter.TransactOpts, chain, chainSelector)
}

func (_CLAdapter *CLAdapterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CLAdapter.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_CLAdapter *CLAdapterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CLAdapter.Contract.TransferOwnership(&_CLAdapter.TransactOpts, newOwner)
}

func (_CLAdapter *CLAdapterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CLAdapter.Contract.TransferOwnership(&_CLAdapter.TransactOpts, newOwner)
}

type CLAdapterBridgeChangedIterator struct {
	Event *CLAdapterBridgeChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CLAdapterBridgeChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CLAdapterBridgeChanged)
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
		it.Event = new(CLAdapterBridgeChanged)
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

func (it *CLAdapterBridgeChangedIterator) Error() error {
	return it.fail
}

func (it *CLAdapterBridgeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CLAdapterBridgeChanged struct {
	OldBridge common.Address
	NewBridge common.Address
	Raw       types.Log
}

func (_CLAdapter *CLAdapterFilterer) FilterBridgeChanged(opts *bind.FilterOpts, oldBridge []common.Address, newBridge []common.Address) (*CLAdapterBridgeChangedIterator, error) {

	var oldBridgeRule []interface{}
	for _, oldBridgeItem := range oldBridge {
		oldBridgeRule = append(oldBridgeRule, oldBridgeItem)
	}
	var newBridgeRule []interface{}
	for _, newBridgeItem := range newBridge {
		newBridgeRule = append(newBridgeRule, newBridgeItem)
	}

	logs, sub, err := _CLAdapter.contract.FilterLogs(opts, "BridgeChanged", oldBridgeRule, newBridgeRule)
	if err != nil {
		return nil, err
	}
	return &CLAdapterBridgeChangedIterator{contract: _CLAdapter.contract, event: "BridgeChanged", logs: logs, sub: sub}, nil
}

func (_CLAdapter *CLAdapterFilterer) WatchBridgeChanged(opts *bind.WatchOpts, sink chan<- *CLAdapterBridgeChanged, oldBridge []common.Address, newBridge []common.Address) (event.Subscription, error) {

	var oldBridgeRule []interface{}
	for _, oldBridgeItem := range oldBridge {
		oldBridgeRule = append(oldBridgeRule, oldBridgeItem)
	}
	var newBridgeRule []interface{}
	for _, newBridgeItem := range newBridge {
		newBridgeRule = append(newBridgeRule, newBridgeItem)
	}

	logs, sub, err := _CLAdapter.contract.WatchLogs(opts, "BridgeChanged", oldBridgeRule, newBridgeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CLAdapterBridgeChanged)
				if err := _CLAdapter.contract.UnpackLog(event, "BridgeChanged", log); err != nil {
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

func (_CLAdapter *CLAdapterFilterer) ParseBridgeChanged(log types.Log) (*CLAdapterBridgeChanged, error) {
	event := new(CLAdapterBridgeChanged)
	if err := _CLAdapter.contract.UnpackLog(event, "BridgeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CLAdapterCLChainSelectorSetIterator struct {
	Event *CLAdapterCLChainSelectorSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CLAdapterCLChainSelectorSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CLAdapterCLChainSelectorSet)
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
		it.Event = new(CLAdapterCLChainSelectorSet)
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

func (it *CLAdapterCLChainSelectorSetIterator) Error() error {
	return it.fail
}

func (it *CLAdapterCLChainSelectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CLAdapterCLChainSelectorSet struct {
	Arg0 [32]byte
	Arg1 uint64
	Raw  types.Log
}

func (_CLAdapter *CLAdapterFilterer) FilterCLChainSelectorSet(opts *bind.FilterOpts) (*CLAdapterCLChainSelectorSetIterator, error) {

	logs, sub, err := _CLAdapter.contract.FilterLogs(opts, "CLChainSelectorSet")
	if err != nil {
		return nil, err
	}
	return &CLAdapterCLChainSelectorSetIterator{contract: _CLAdapter.contract, event: "CLChainSelectorSet", logs: logs, sub: sub}, nil
}

func (_CLAdapter *CLAdapterFilterer) WatchCLChainSelectorSet(opts *bind.WatchOpts, sink chan<- *CLAdapterCLChainSelectorSet) (event.Subscription, error) {

	logs, sub, err := _CLAdapter.contract.WatchLogs(opts, "CLChainSelectorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CLAdapterCLChainSelectorSet)
				if err := _CLAdapter.contract.UnpackLog(event, "CLChainSelectorSet", log); err != nil {
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

func (_CLAdapter *CLAdapterFilterer) ParseCLChainSelectorSet(log types.Log) (*CLAdapterCLChainSelectorSet, error) {
	event := new(CLAdapterCLChainSelectorSet)
	if err := _CLAdapter.contract.UnpackLog(event, "CLChainSelectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CLAdapterCLTokenPoolDeployedIterator struct {
	Event *CLAdapterCLTokenPoolDeployed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CLAdapterCLTokenPoolDeployedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CLAdapterCLTokenPoolDeployed)
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
		it.Event = new(CLAdapterCLTokenPoolDeployed)
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

func (it *CLAdapterCLTokenPoolDeployedIterator) Error() error {
	return it.fail
}

func (it *CLAdapterCLTokenPoolDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CLAdapterCLTokenPoolDeployed struct {
	Arg0 common.Address
	Raw  types.Log
}

func (_CLAdapter *CLAdapterFilterer) FilterCLTokenPoolDeployed(opts *bind.FilterOpts) (*CLAdapterCLTokenPoolDeployedIterator, error) {

	logs, sub, err := _CLAdapter.contract.FilterLogs(opts, "CLTokenPoolDeployed")
	if err != nil {
		return nil, err
	}
	return &CLAdapterCLTokenPoolDeployedIterator{contract: _CLAdapter.contract, event: "CLTokenPoolDeployed", logs: logs, sub: sub}, nil
}

func (_CLAdapter *CLAdapterFilterer) WatchCLTokenPoolDeployed(opts *bind.WatchOpts, sink chan<- *CLAdapterCLTokenPoolDeployed) (event.Subscription, error) {

	logs, sub, err := _CLAdapter.contract.WatchLogs(opts, "CLTokenPoolDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CLAdapterCLTokenPoolDeployed)
				if err := _CLAdapter.contract.UnpackLog(event, "CLTokenPoolDeployed", log); err != nil {
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

func (_CLAdapter *CLAdapterFilterer) ParseCLTokenPoolDeployed(log types.Log) (*CLAdapterCLTokenPoolDeployed, error) {
	event := new(CLAdapterCLTokenPoolDeployed)
	if err := _CLAdapter.contract.UnpackLog(event, "CLTokenPoolDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CLAdapterExecutionGasLimitSetIterator struct {
	Event *CLAdapterExecutionGasLimitSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CLAdapterExecutionGasLimitSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CLAdapterExecutionGasLimitSet)
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
		it.Event = new(CLAdapterExecutionGasLimitSet)
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

func (it *CLAdapterExecutionGasLimitSetIterator) Error() error {
	return it.fail
}

func (it *CLAdapterExecutionGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CLAdapterExecutionGasLimitSet struct {
	PrevVal *big.Int
	NewVal  *big.Int
	Raw     types.Log
}

func (_CLAdapter *CLAdapterFilterer) FilterExecutionGasLimitSet(opts *bind.FilterOpts, prevVal []*big.Int, newVal []*big.Int) (*CLAdapterExecutionGasLimitSetIterator, error) {

	var prevValRule []interface{}
	for _, prevValItem := range prevVal {
		prevValRule = append(prevValRule, prevValItem)
	}
	var newValRule []interface{}
	for _, newValItem := range newVal {
		newValRule = append(newValRule, newValItem)
	}

	logs, sub, err := _CLAdapter.contract.FilterLogs(opts, "ExecutionGasLimitSet", prevValRule, newValRule)
	if err != nil {
		return nil, err
	}
	return &CLAdapterExecutionGasLimitSetIterator{contract: _CLAdapter.contract, event: "ExecutionGasLimitSet", logs: logs, sub: sub}, nil
}

func (_CLAdapter *CLAdapterFilterer) WatchExecutionGasLimitSet(opts *bind.WatchOpts, sink chan<- *CLAdapterExecutionGasLimitSet, prevVal []*big.Int, newVal []*big.Int) (event.Subscription, error) {

	var prevValRule []interface{}
	for _, prevValItem := range prevVal {
		prevValRule = append(prevValRule, prevValItem)
	}
	var newValRule []interface{}
	for _, newValItem := range newVal {
		newValRule = append(newValRule, newValItem)
	}

	logs, sub, err := _CLAdapter.contract.WatchLogs(opts, "ExecutionGasLimitSet", prevValRule, newValRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CLAdapterExecutionGasLimitSet)
				if err := _CLAdapter.contract.UnpackLog(event, "ExecutionGasLimitSet", log); err != nil {
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

func (_CLAdapter *CLAdapterFilterer) ParseExecutionGasLimitSet(log types.Log) (*CLAdapterExecutionGasLimitSet, error) {
	event := new(CLAdapterExecutionGasLimitSet)
	if err := _CLAdapter.contract.UnpackLog(event, "ExecutionGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CLAdapterOwnershipTransferredIterator struct {
	Event *CLAdapterOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CLAdapterOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CLAdapterOwnershipTransferred)
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
		it.Event = new(CLAdapterOwnershipTransferred)
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

func (it *CLAdapterOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CLAdapterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CLAdapterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log
}

func (_CLAdapter *CLAdapterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CLAdapterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CLAdapter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CLAdapterOwnershipTransferredIterator{contract: _CLAdapter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CLAdapter *CLAdapterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CLAdapterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CLAdapter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CLAdapterOwnershipTransferred)
				if err := _CLAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CLAdapter *CLAdapterFilterer) ParseOwnershipTransferred(log types.Log) (*CLAdapterOwnershipTransferred, error) {
	event := new(CLAdapterOwnershipTransferred)
	if err := _CLAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_CLAdapter *CLAdapter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CLAdapter.abi.Events["BridgeChanged"].ID:
		return _CLAdapter.ParseBridgeChanged(log)
	case _CLAdapter.abi.Events["CLChainSelectorSet"].ID:
		return _CLAdapter.ParseCLChainSelectorSet(log)
	case _CLAdapter.abi.Events["CLTokenPoolDeployed"].ID:
		return _CLAdapter.ParseCLTokenPoolDeployed(log)
	case _CLAdapter.abi.Events["ExecutionGasLimitSet"].ID:
		return _CLAdapter.ParseExecutionGasLimitSet(log)
	case _CLAdapter.abi.Events["OwnershipTransferred"].ID:
		return _CLAdapter.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CLAdapterBridgeChanged) Topic() common.Hash {
	return common.HexToHash("0xd565484d693f5157abcceb853139678038bc740991b0a4dc3baa2426325bb3c0")
}

func (CLAdapterCLChainSelectorSet) Topic() common.Hash {
	return common.HexToHash("0x61805347ac73c642ed5bfbc228e1702a4d51591b36a35387250de0c0d75b0072")
}

func (CLAdapterCLTokenPoolDeployed) Topic() common.Hash {
	return common.HexToHash("0x45163ba2f75e282ba1000a5e166237acb454fa6b2b4cc778671bf977a01b13b4")
}

func (CLAdapterExecutionGasLimitSet) Topic() common.Hash {
	return common.HexToHash("0x45a4a024e3d155ab5b1634193775d00d4cbf4ef295e76d0bf1aa6dfeb114dd84")
}

func (CLAdapterOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_CLAdapter *CLAdapter) Address() common.Address {
	return _CLAdapter.address
}

type CLAdapterInterface interface {
	Bridge(opts *bind.CallOpts) (common.Address, error)

	GetChain(opts *bind.CallOpts, arg0 uint64) ([32]byte, error)

	GetExecutionGasLimit(opts *bind.CallOpts) (*big.Int, error)

	GetFee(opts *bind.CallOpts, _toChain [32]byte, arg1 [32]byte, _toAddress [32]byte, _amount *big.Int, _payload []byte) (*big.Int, error)

	GetRemoteChainSelector(opts *bind.CallOpts, arg0 [32]byte) (uint64, error)

	Lbtc(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TokenPool(opts *bind.CallOpts) (common.Address, error)

	ChangeBridge(opts *bind.TransactOpts, bridge_ common.Address) (*types.Transaction, error)

	Deposit(opts *bind.TransactOpts, fromAddress common.Address, _toChain [32]byte, arg2 [32]byte, _toAddress [32]byte, _amount *big.Int, _payload []byte) (*types.Transaction, error)

	InitWithdrawalNoSignatures(opts *bind.TransactOpts, remoteSelector uint64, onChainData []byte) (*types.Transaction, error)

	InitiateDeposit(opts *bind.TransactOpts, remoteChainSelector uint64, receiver []byte, amount *big.Int) (*types.Transaction, error)

	InitiateWithdrawal(opts *bind.TransactOpts, remoteSelector uint64, offChainData []byte) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetExecutionGasLimit(opts *bind.TransactOpts, newVal *big.Int) (*types.Transaction, error)

	SetRemoteChainSelector(opts *bind.TransactOpts, chain [32]byte, chainSelector uint64) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	FilterBridgeChanged(opts *bind.FilterOpts, oldBridge []common.Address, newBridge []common.Address) (*CLAdapterBridgeChangedIterator, error)

	WatchBridgeChanged(opts *bind.WatchOpts, sink chan<- *CLAdapterBridgeChanged, oldBridge []common.Address, newBridge []common.Address) (event.Subscription, error)

	ParseBridgeChanged(log types.Log) (*CLAdapterBridgeChanged, error)

	FilterCLChainSelectorSet(opts *bind.FilterOpts) (*CLAdapterCLChainSelectorSetIterator, error)

	WatchCLChainSelectorSet(opts *bind.WatchOpts, sink chan<- *CLAdapterCLChainSelectorSet) (event.Subscription, error)

	ParseCLChainSelectorSet(log types.Log) (*CLAdapterCLChainSelectorSet, error)

	FilterCLTokenPoolDeployed(opts *bind.FilterOpts) (*CLAdapterCLTokenPoolDeployedIterator, error)

	WatchCLTokenPoolDeployed(opts *bind.WatchOpts, sink chan<- *CLAdapterCLTokenPoolDeployed) (event.Subscription, error)

	ParseCLTokenPoolDeployed(log types.Log) (*CLAdapterCLTokenPoolDeployed, error)

	FilterExecutionGasLimitSet(opts *bind.FilterOpts, prevVal []*big.Int, newVal []*big.Int) (*CLAdapterExecutionGasLimitSetIterator, error)

	WatchExecutionGasLimitSet(opts *bind.WatchOpts, sink chan<- *CLAdapterExecutionGasLimitSet, prevVal []*big.Int, newVal []*big.Int) (event.Subscription, error)

	ParseExecutionGasLimitSet(log types.Log) (*CLAdapterExecutionGasLimitSet, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CLAdapterOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CLAdapterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CLAdapterOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
