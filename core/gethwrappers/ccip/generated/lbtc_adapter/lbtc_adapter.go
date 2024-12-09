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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"contractLombardTokenPool\",\"name\":\"tokenPool_\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"executionGasLimit_\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Adapter_ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CLAttemptToOverrideChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CLAttemptToOverrideChainSelector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CLRefundFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CLUnauthorizedTokenPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CLZeroChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CLZeroChanSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBridge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"NotEnoughToPayFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroPayload\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBridge\",\"name\":\"oldBridge\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIBridge\",\"name\":\"newBridge\",\"type\":\"address\"}],\"name\":\"BridgeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"CLChainSelectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CLTokenPoolDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"prevVal\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"newVal\",\"type\":\"uint128\"}],\"name\":\"ExecutionGasLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"}],\"name\":\"changeBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_toChain\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutionGasLimit\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_toChain\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getRemoteChainSelector\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"onChainData\",\"type\":\"bytes\"}],\"name\":\"initWithdrawalNoSignatures\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"initiateDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastBurnedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"lastPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offChainData\",\"type\":\"bytes\"}],\"name\":\"initiateWithdrawal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lbtc\",\"outputs\":[{\"internalType\":\"contractILBTC\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newVal\",\"type\":\"uint128\"}],\"name\":\"setExecutionGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chain\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"setRemoteChainSelector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPool\",\"outputs\":[{\"internalType\":\"contractLombardTokenPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200246c3803806200246c83398101604081905262000034916200018d565b600080546001600160a01b0319166001600160a01b038516179055620000616200005b3390565b620000c6565b6200006c8162000118565b600580546001600160a01b0319166001600160a01b0384169081179091556040519081527f45163ba2f75e282ba1000a5e166237acb454fa6b2b4cc778671bf977a01b13b49060200160405180910390a1505050620001ec565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6004546040516001600160801b038084169216907f45a4a024e3d155ab5b1634193775d00d4cbf4ef295e76d0bf1aa6dfeb114dd8490600090a3600480546001600160801b0319166001600160801b0392909216919091179055565b6001600160a01b03811681146200018a57600080fd5b50565b600080600060608486031215620001a357600080fd5b8351620001b08162000174565b6020850151909350620001c38162000174565b60408501519092506001600160801b0381168114620001e157600080fd5b809150509250925092565b61227080620001fc6000396000f3fe60806040526004361061010e5760003560e01c8063715886c4116100a5578063a1a6d50811610074578063e240550d11610059578063e240550d14610377578063e78cea9214610397578063f2fde38b146103c457600080fd5b8063a1a6d5081461031c578063b1ab7e221461035757600080fd5b8063715886c41461025d578063775710d4146102705780638da5cb5b146102a757806392b0680a146102d257600080fd5b80635391a405116100e15780635391a405146101e5578063550e7ab21461020557806368b3c91014610233578063715018a61461024857600080fd5b806308774410146101135780630c373d7414610135578063104e992914610173578063453bb95c146101c5575b600080fd5b34801561011f57600080fd5b5061013361012e3660046118be565b6103e4565b005b34801561014157600080fd5b50610155610150366004611941565b61046a565b60405167ffffffffffffffff90911681526020015b60405180910390f35b34801561017f57600080fd5b506005546101a09073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161016a565b3480156101d157600080fd5b506101336101e0366004611996565b61057e565b3480156101f157600080fd5b50610155610200366004611941565b610731565b34801561021157600080fd5b506102256102203660046119c6565b6108a0565b60405161016a929190611a90565b34801561023f57600080fd5b506101a0610c47565b34801561025457600080fd5b50610133610cde565b61013361026b366004611bc3565b610cf2565b34801561027c57600080fd5b5061015561028b366004611c38565b60026020526000908152604090205467ffffffffffffffff1681565b3480156102b357600080fd5b5060015473ffffffffffffffffffffffffffffffffffffffff166101a0565b3480156102de57600080fd5b506004546102fb906fffffffffffffffffffffffffffffffff1681565b6040516fffffffffffffffffffffffffffffffff909116815260200161016a565b34801561032857600080fd5b50610349610337366004611c51565b60036020526000908152604090205481565b60405190815260200161016a565b34801561036357600080fd5b50610349610372366004611c6e565b6110d8565b34801561038357600080fd5b50610133610392366004611cd2565b61122d565b3480156103a357600080fd5b506000546101a09073ffffffffffffffffffffffffffffffffffffffff1681565b3480156103d057600080fd5b506101336103df3660046118be565b611241565b6103ec6112f5565b6103f5816112fd565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917fd565484d693f5157abcceb853139678038bc740991b0a4dc3baa2426325bb3c09190a35050565b600061047461134a565b6104db600360008667ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000205484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061139d92505050565b6000546040517f0968f26400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690630968f264906105339086908690600401611d04565b6020604051808303816000875af1158015610552573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105769190611d51565b949350505050565b610586611423565b816105bd576040517fe531cb7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff16600003610601576040517fe531cb7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281526002602052604090205467ffffffffffffffff1615610651576040517fb72f782a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8116600090815260036020526040902054156106a2576040517fe95ca6dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082815260026020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff861690811790915580845260038352928190208590558051858152918201929092527f61805347ac73c642ed5bfbc228e1702a4d51591b36a35387250de0c0d75b0072910160405180910390a15050565b600061073b61134a565b60008061074a84860186611d6e565b67ffffffffffffffff88166000908152600360205260409020549193509150610773908361139d565b6000546040517f6b93aa5100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690636b93aa51906107cb9085908590600401611dd2565b600060405180830381600087803b1580156107e557600080fd5b505af11580156107f9573d6000803e3d6000fd5b50506000546040517f0968f26400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169250630968f2649150610853908590600401611e00565b6020604051808303816000875af1158015610872573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108969190611d51565b9695505050505050565b600060606108ac61134a565b6000600780546108bb90611e13565b905011156109a2576006549150600780546108d590611e13565b80601f016020809104026020016040519081016040528092919081815260200182805461090190611e13565b801561094e5780601f106109235761010080835404028352916020019161094e565b820191906000526020600020905b81548152906001019060200180831161093157829003601f168201915b509394506000935061095f92505050565b6040519080825280601f01601f191660200182016040528015610989576020820181803683370190505b506007906109979082611eb3565b506000600655610b43565b6109aa610c47565b6000546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526024810186905291169063095ea7b3906044016020604051808303816000875af1158015610a22573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a469190611fcd565b506000805467ffffffffffffffff88168252600360205260409091205473ffffffffffffffffffffffffffffffffffffffff909116906383031c4190610a8c8789611fef565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526004810192909252602482015267ffffffffffffffff861660448201526064016000604051808303816000875af1158015610af7573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610b3d919081019061202c565b90925090505b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166368b3c9106040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bd291906120af565b73ffffffffffffffffffffffffffffffffffffffff166342966c68836040518263ffffffff1660e01b8152600401610c0c91815260200190565b600060405180830381600087803b158015610c2657600080fd5b505af1158015610c3a573d6000803e3d6000fd5b5050505094509492505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166368b3c9106040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cb5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd991906120af565b905090565b610ce6611423565b610cf060006114a4565b565b73ffffffffffffffffffffffffffffffffffffffff861630146110d05760068290556007610d208282611eb3565b50600085815260026020908152604080832054815192830187905267ffffffffffffffff169291610d629101604051602081830303815290604052858561151b565b90506000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b0f479a16040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dd3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df791906120af565b905060008173ffffffffffffffffffffffffffffffffffffffff166320487ded85856040518363ffffffff1660e01b8152600401610e369291906120cc565b602060405180830381865afa158015610e53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e7791906121e1565b905080341015610ebb576040517faf09dec9000000000000000000000000000000000000000000000000000000008152600481018290526024015b60405180910390fd5b80341115610f8d576000610ecf82346121fa565b905060008b73ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114610f2b576040519150601f19603f3d011682016040523d82523d6000602084013e610f30565b606091505b5050905080610f8a576040517f9bb3185e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d16600482015260248101839052604401610eb2565b50505b610f95610c47565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015260248201899052919091169063095ea7b3906044016020604051808303816000875af115801561100c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110309190611fcd565b506040517f96f4e9f900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316906396f4e9f990839061108790889088906004016120cc565b60206040518083038185885af11580156110a5573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906110ca91906121e1565b50505050505b505050505050565b600554604080517fb0f479a1000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163b0f479a19160048083019260209291908290030181865afa158015611148573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061116c91906120af565b60008781526002602090815260409182902054825191820188905273ffffffffffffffffffffffffffffffffffffffff93909316926320487ded9267ffffffffffffffff909116916111cf9101604051602081830303815290604052878761151b565b6040518363ffffffff1660e01b81526004016111ec9291906120cc565b602060405180830381865afa158015611209573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089691906121e1565b611235611423565b61123e81611816565b50565b611249611423565b73ffffffffffffffffffffffffffffffffffffffff81166112ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610eb2565b61123e816114a4565b610cf0611423565b73ffffffffffffffffffffffffffffffffffffffff811661123e576040517fcf891a8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60055473ffffffffffffffffffffffffffffffffffffffff163314610cf0576040517f36b0b92e000000000000000000000000000000000000000000000000000000008152336004820152602401610eb2565b6000546040517f8783f52500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690638783f525906113f59085908590600401611a90565b600060405180830381600087803b15801561140f57600080fd5b505af11580156110d0573d6000803e3d6000fd5b60015473ffffffffffffffffffffffffffffffffffffffff163314610cf0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610eb2565b6001805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6115636040518060a00160405280606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b604080516001808252818301909252600091816020015b604080518082019091526000808252602082015281526020019060019003908161157a575050604080518082018083526000547f68b3c91000000000000000000000000000000000000000000000000000000000909152915192935091829173ffffffffffffffffffffffffffffffffffffffff16906368b3c910906044808501916020918187030181865afa158015611618573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061163c91906120af565b73ffffffffffffffffffffffffffffffffffffffff168152602001858152508160008151811061166e5761166e612234565b60200260200101819052506060600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166355b961566040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061170c9190611fcd565b61174e57835160000361174b576040517ff3cc293900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50825b6040805160a081018252878152602080820184905281830185905260006060830152825180840184526004546fffffffffffffffffffffffffffffffff1680825260019183019182528451602481019190915290511515604480830191909152845180830390910181526064909101909352820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f181dcf10000000000000000000000000000000000000000000000000000000001790529060808201529695505050505050565b6004546040516fffffffffffffffffffffffffffffffff8084169216907f45a4a024e3d155ab5b1634193775d00d4cbf4ef295e76d0bf1aa6dfeb114dd8490600090a3600480547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff92909216919091179055565b73ffffffffffffffffffffffffffffffffffffffff8116811461123e57600080fd5b6000602082840312156118d057600080fd5b81356118db8161189c565b9392505050565b67ffffffffffffffff8116811461123e57600080fd5b60008083601f84011261190a57600080fd5b50813567ffffffffffffffff81111561192257600080fd5b60208301915083602082850101111561193a57600080fd5b9250929050565b60008060006040848603121561195657600080fd5b8335611961816118e2565b9250602084013567ffffffffffffffff81111561197d57600080fd5b611989868287016118f8565b9497909650939450505050565b600080604083850312156119a957600080fd5b8235915060208301356119bb816118e2565b809150509250929050565b600080600080606085870312156119dc57600080fd5b84356119e7816118e2565b9350602085013567ffffffffffffffff811115611a0357600080fd5b611a0f878288016118f8565b9598909750949560400135949350505050565b60005b83811015611a3d578181015183820152602001611a25565b50506000910152565b60008151808452611a5e816020860160208601611a22565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b8281526040602082015260006105766040830184611a46565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611b1f57611b1f611aa9565b604052919050565b600067ffffffffffffffff821115611b4157611b41611aa9565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112611b7e57600080fd5b8135611b91611b8c82611b27565b611ad8565b818152846020838601011115611ba657600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c08789031215611bdc57600080fd5b8635611be78161189c565b95506020870135945060408701359350606087013592506080870135915060a087013567ffffffffffffffff811115611c1f57600080fd5b611c2b89828a01611b6d565b9150509295509295509295565b600060208284031215611c4a57600080fd5b5035919050565b600060208284031215611c6357600080fd5b81356118db816118e2565b600080600080600060a08688031215611c8657600080fd5b85359450602086013593506040860135925060608601359150608086013567ffffffffffffffff811115611cb957600080fd5b611cc588828901611b6d565b9150509295509295909350565b600060208284031215611ce457600080fd5b81356fffffffffffffffffffffffffffffffff811681146118db57600080fd5b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b600060208284031215611d6357600080fd5b81516118db816118e2565b60008060408385031215611d8157600080fd5b823567ffffffffffffffff80821115611d9957600080fd5b611da586838701611b6d565b93506020850135915080821115611dbb57600080fd5b50611dc885828601611b6d565b9150509250929050565b604081526000611de56040830185611a46565b8281036020840152611df78185611a46565b95945050505050565b6020815260006118db6020830184611a46565b600181811c90821680611e2757607f821691505b602082108103611e60577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f821115611eae576000816000526020600020601f850160051c81016020861015611e8f5750805b601f850160051c820191505b818110156110d057828155600101611e9b565b505050565b815167ffffffffffffffff811115611ecd57611ecd611aa9565b611ee181611edb8454611e13565b84611e66565b602080601f831160018114611f345760008415611efe5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556110d0565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611f8157888601518255948401946001909101908401611f62565b5085821015611fbd57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b600060208284031215611fdf57600080fd5b815180151581146118db57600080fd5b80356020831015612026577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602084900360031b1b165b92915050565b6000806040838503121561203f57600080fd5b82519150602083015167ffffffffffffffff81111561205d57600080fd5b8301601f8101851361206e57600080fd5b805161207c611b8c82611b27565b81815286602083850101111561209157600080fd5b6120a2826020830160208601611a22565b8093505050509250929050565b6000602082840312156120c157600080fd5b81516118db8161189c565b6000604067ffffffffffffffff851683526020604081850152845160a060408601526120fb60e0860182611a46565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808784030160608801526121368383611a46565b6040890151888203830160808a01528051808352908601945060009350908501905b80841015612197578451805173ffffffffffffffffffffffffffffffffffffffff16835286015186830152938501936001939093019290860190612158565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a015295506121d38187611a46565b9a9950505050505050505050565b6000602082840312156121f357600080fd5b5051919050565b81810381811115612026577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c6343000818000a",
}

var CLAdapterABI = CLAdapterMetaData.ABI

var CLAdapterBin = CLAdapterMetaData.Bin

func DeployCLAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, bridge_ common.Address, tokenPool_ common.Address, executionGasLimit_ *big.Int) (common.Address, *types.Transaction, *CLAdapter, error) {
	parsed, err := CLAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CLAdapterBin), backend, bridge_, tokenPool_, executionGasLimit_)
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
