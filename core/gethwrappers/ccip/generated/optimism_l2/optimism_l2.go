// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimism_l2

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

var OptimismL2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_otherBridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"DepositFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHBridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHBridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractCrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_BRIDGE\",\"outputs\":[{\"internalType\":\"contractStandardBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeETHTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1TokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractCrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162002cbc38038062002cbc8339810160408190526200003491620001a0565b7342000000000000000000000000000000000000076080526001600160a01b03811660a052620000636200006a565b50620001d2565b600054610100900460ff16158080156200008b5750600054600160ff909116105b80620000bb5750620000a8306200019160201b620004941760201c565b158015620000bb575060005460ff166001145b620001235760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b6000805460ff19166001179055801562000147576000805461ff0019166101001790555b80156200018e576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6001600160a01b03163b151590565b600060208284031215620001b357600080fd5b81516001600160a01b0381168114620001cb57600080fd5b9392505050565b60805160a051612a6962000253600039600081816102690152818161039d015281816105cd01528181610a910152818161147301526117b40152600081816102c20152818161044c015281816105a30152818161060401528181610a6701528181610ac801528181610d550152818161143601526117780152612a696000f3fe60806040526004361061012d5760003560e01c8063662a633a116100a55780638f601f6611610074578063a3a7954811610059578063a3a795481461046e578063c89701a21461025a578063e11013dd1461048157600080fd5b80638f601f66146103f4578063927ede2d1461043a57600080fd5b8063662a633a146103785780637f46ddb21461038b5780638129fc1c146103bf57806387087623146103d457600080fd5b806336c717c1116100fc578063540abf73116100e1578063540abf73146102e657806354fd4d50146103065780635c975abb1461035c57600080fd5b806336c717c11461025a5780633cb747bf146102b357600080fd5b80630166a07a1461020157806309fc8843146102215780631635f5fd1461023457806332b7006d1461024757600080fd5b366101fc57333b156101c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b6101fa73deaddeaddeaddeaddeaddeaddeaddeaddead000033333462030d40604051806020016040528060008152506104b0565b005b600080fd5b34801561020d57600080fd5b506101fa61021c366004612479565b61058b565b6101fa61022f36600461252a565b610978565b6101fa61024236600461257d565b610a4f565b6101fa6102553660046125f0565b610f1c565b34801561026657600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156102bf57600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610289565b3480156102f257600080fd5b506101fa610301366004612644565b610ff6565b34801561031257600080fd5b5061034f6040518060400160405280600581526020017f312e372e3000000000000000000000000000000000000000000000000000000081525081565b6040516102aa9190612731565b34801561036857600080fd5b50604051600081526020016102aa565b6101fa610386366004612479565b61103b565b34801561039757600080fd5b506102897f000000000000000000000000000000000000000000000000000000000000000081565b3480156103cb57600080fd5b506101fa6110ae565b3480156103e057600080fd5b506101fa6103ef366004612744565b611238565b34801561040057600080fd5b5061042c61040f3660046127c7565b600260209081526000928352604080842090915290825290205481565b6040519081526020016102aa565b34801561044657600080fd5b506102897f000000000000000000000000000000000000000000000000000000000000000081565b6101fa61047c366004612744565b61130c565b6101fa61048f366004612800565b611350565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b7fffffffffffffffffffffffff215221522152215221522152215221522153000073ffffffffffffffffffffffffffffffffffffffff8716016104ff576104fa8585858585611399565b610583565b60008673ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa15801561054c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105709190612863565b90506105818782888888888861157d565b505b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156106a957507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561066d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106919190612863565b73ffffffffffffffffffffffffffffffffffffffff16145b61075b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101bd565b610764876118c4565b156108b2576107738787611926565b610825576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101bd565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b15801561089557600080fd5b505af11580156108a9573d6000803e3d6000fd5b50505050610934565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a16835292905220546108f09084906128af565b73ffffffffffffffffffffffffffffffffffffffff8089166000818152600260209081526040808320948c1683529390529190912091909155610934908585611a46565b610581878787878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611b1a92505050565b333b15610a07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101bd565b610a4a3333348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061139992505050565b505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148015610b6d57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b31573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b559190612863565b73ffffffffffffffffffffffffffffffffffffffff16145b610c1f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101bd565b823414610cae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e7420726571756972656400000000000060648201526084016101bd565b3073ffffffffffffffffffffffffffffffffffffffff851603610d53576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c66000000000000000000000000000000000000000000000000000000000060648201526084016101bd565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610e2e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201527f657373656e67657200000000000000000000000000000000000000000000000060648201526084016101bd565b610e7085858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611ba892505050565b6000610e8d855a8660405180602001604052806000815250611c49565b905080610583576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c6564000000000000000000000000000000000000000000000000000000000060648201526084016101bd565b333b15610fab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101bd565b610fef853333878787878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104b092505050565b5050505050565b61058187873388888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061157d92505050565b73ffffffffffffffffffffffffffffffffffffffff8716158015611088575073ffffffffffffffffffffffffffffffffffffffff861673deaddeaddeaddeaddeaddeaddeaddeaddead0000145b1561109f5761109a8585858585610a4f565b610581565b6105818688878787878761058b565b600054610100900460ff16158080156110ce5750600054600160ff909116105b806110e85750303b1580156110e8575060005460ff166001145b611174576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016101bd565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156111d257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b801561123557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b333b156112c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101bd565b61058386863333888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061157d92505050565b610583863387878787878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104b092505050565b6113933385348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061139992505050565b50505050565b823414611428576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5374616e646172644272696467653a206272696467696e6720455448206d757360448201527f7420696e636c7564652073756666696369656e74204554482076616c7565000060648201526084016101bd565b61143485858584611c63565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633dbb202b847f0000000000000000000000000000000000000000000000000000000000000000631635f5fd60e01b898989886040516024016114b194939291906128c6565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b90921682526115449291889060040161290f565b6000604051808303818588803b15801561155d57600080fd5b505af1158015611571573d6000803e3d6000fd5b50505050505050505050565b611586876118c4565b156116d4576115958787611926565b611647576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101bd565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201859052881690639dc29fac90604401600060405180830381600087803b1580156116b757600080fd5b505af11580156116cb573d6000803e3d6000fd5b50505050611768565b6116f673ffffffffffffffffffffffffffffffffffffffff8816863086611d04565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a1683529290522054611734908490612954565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600260209081526040808320938b16835292905220555b611776878787878786611d62565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633dbb202b7f0000000000000000000000000000000000000000000000000000000000000000630166a07a60e01b898b8a8a8a896040516024016117f69695949392919061296c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b90921682526118899291879060040161290f565b600060405180830381600087803b1580156118a357600080fd5b505af11580156118b7573d6000803e3d6000fd5b5050505050505050505050565b60006118f0827f1d1d8b6300000000000000000000000000000000000000000000000000000000611df0565b806119205750611920827fec4fc8e300000000000000000000000000000000000000000000000000000000611df0565b92915050565b6000611952837f1d1d8b6300000000000000000000000000000000000000000000000000000000611df0565b156119fb578273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119c69190612863565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050611920565b8273ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119a2573d6000803e3d6000fd5b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610a4a9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611e13565b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89868686604051611b92939291906129c7565b60405180910390a4610583868686868686611f1f565b8373ffffffffffffffffffffffffffffffffffffffff1673deaddeaddeaddeaddeaddeaddeaddeaddead000073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89868686604051611c35939291906129c7565b60405180910390a461139384848484611fa7565b600080600080845160208601878a8af19695505050505050565b8373ffffffffffffffffffffffffffffffffffffffff1673deaddeaddeaddeaddeaddeaddeaddeaddead000073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e868686604051611cf0939291906129c7565b60405180910390a461139384848484612014565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526113939085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611a98565b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e868686604051611dda939291906129c7565b60405180910390a4610583868686868686612073565b6000611dfb836120eb565b8015611e0c5750611e0c838361214f565b9392505050565b6000611e75826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661221e9092919063ffffffff16565b805190915015610a4a5780806020019051810190611e939190612a05565b610a4a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101bd565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd868686604051611f97939291906129c7565b60405180910390a4505050505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d8484604051612006929190612a27565b60405180910390a350505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af58484604051612006929190612a27565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf868686604051611f97939291906129c7565b6000612117827f01ffc9a70000000000000000000000000000000000000000000000000000000061214f565b80156119205750612148827fffffffff0000000000000000000000000000000000000000000000000000000061214f565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015612207575060208210155b80156122135750600081115b979650505050505050565b606061222d8484600085612235565b949350505050565b6060824710156122c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101bd565b73ffffffffffffffffffffffffffffffffffffffff85163b612345576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101bd565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161236e9190612a40565b60006040518083038185875af1925050503d80600081146123ab576040519150601f19603f3d011682016040523d82523d6000602084013e6123b0565b606091505b5091509150612213828286606083156123ca575081611e0c565b8251156123da5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101bd9190612731565b73ffffffffffffffffffffffffffffffffffffffff8116811461123557600080fd5b60008083601f84011261244257600080fd5b50813567ffffffffffffffff81111561245a57600080fd5b60208301915083602082850101111561247257600080fd5b9250929050565b600080600080600080600060c0888a03121561249457600080fd5b873561249f8161240e565b965060208801356124af8161240e565b955060408801356124bf8161240e565b945060608801356124cf8161240e565b93506080880135925060a088013567ffffffffffffffff8111156124f257600080fd5b6124fe8a828b01612430565b989b979a50959850939692959293505050565b803563ffffffff8116811461252557600080fd5b919050565b60008060006040848603121561253f57600080fd5b61254884612511565b9250602084013567ffffffffffffffff81111561256457600080fd5b61257086828701612430565b9497909650939450505050565b60008060008060006080868803121561259557600080fd5b85356125a08161240e565b945060208601356125b08161240e565b935060408601359250606086013567ffffffffffffffff8111156125d357600080fd5b6125df88828901612430565b969995985093965092949392505050565b60008060008060006080868803121561260857600080fd5b85356126138161240e565b94506020860135935061262860408701612511565b9250606086013567ffffffffffffffff8111156125d357600080fd5b600080600080600080600060c0888a03121561265f57600080fd5b873561266a8161240e565b9650602088013561267a8161240e565b9550604088013561268a8161240e565b94506060880135935061269f60808901612511565b925060a088013567ffffffffffffffff8111156124f257600080fd5b60005b838110156126d65781810151838201526020016126be565b838111156113935750506000910152565b600081518084526126ff8160208601602086016126bb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611e0c60208301846126e7565b60008060008060008060a0878903121561275d57600080fd5b86356127688161240e565b955060208701356127788161240e565b94506040870135935061278d60608801612511565b9250608087013567ffffffffffffffff8111156127a957600080fd5b6127b589828a01612430565b979a9699509497509295939492505050565b600080604083850312156127da57600080fd5b82356127e58161240e565b915060208301356127f58161240e565b809150509250929050565b6000806000806060858703121561281657600080fd5b84356128218161240e565b935061282f60208601612511565b9250604085013567ffffffffffffffff81111561284b57600080fd5b61285787828801612430565b95989497509550505050565b60006020828403121561287557600080fd5b8151611e0c8161240e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156128c1576128c1612880565b500390565b600073ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152508360408301526080606083015261290560808301846126e7565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815260606020820152600061293e60608301856126e7565b905063ffffffff83166040830152949350505050565b6000821982111561296757612967612880565b500190565b600073ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a08301526129bb60c08301846126e7565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681528260208201526060604082015260006129fc60608301846126e7565b95945050505050565b600060208284031215612a1757600080fd5b81518015158114611e0c57600080fd5b82815260406020820152600061222d60408301846126e7565b60008251612a528184602087016126bb565b919091019291505056fea164736f6c634300080f000a",
}

var OptimismL2ABI = OptimismL2MetaData.ABI

var OptimismL2Bin = OptimismL2MetaData.Bin

func DeployOptimismL2(auth *bind.TransactOpts, backend bind.ContractBackend, _otherBridge common.Address) (common.Address, *types.Transaction, *OptimismL2, error) {
	parsed, err := OptimismL2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptimismL2Bin), backend, _otherBridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptimismL2{address: address, abi: *parsed, OptimismL2Caller: OptimismL2Caller{contract: contract}, OptimismL2Transactor: OptimismL2Transactor{contract: contract}, OptimismL2Filterer: OptimismL2Filterer{contract: contract}}, nil
}

type OptimismL2 struct {
	address common.Address
	abi     abi.ABI
	OptimismL2Caller
	OptimismL2Transactor
	OptimismL2Filterer
}

type OptimismL2Caller struct {
	contract *bind.BoundContract
}

type OptimismL2Transactor struct {
	contract *bind.BoundContract
}

type OptimismL2Filterer struct {
	contract *bind.BoundContract
}

type OptimismL2Session struct {
	Contract     *OptimismL2
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OptimismL2CallerSession struct {
	Contract *OptimismL2Caller
	CallOpts bind.CallOpts
}

type OptimismL2TransactorSession struct {
	Contract     *OptimismL2Transactor
	TransactOpts bind.TransactOpts
}

type OptimismL2Raw struct {
	Contract *OptimismL2
}

type OptimismL2CallerRaw struct {
	Contract *OptimismL2Caller
}

type OptimismL2TransactorRaw struct {
	Contract *OptimismL2Transactor
}

func NewOptimismL2(address common.Address, backend bind.ContractBackend) (*OptimismL2, error) {
	abi, err := abi.JSON(strings.NewReader(OptimismL2ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOptimismL2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimismL2{address: address, abi: abi, OptimismL2Caller: OptimismL2Caller{contract: contract}, OptimismL2Transactor: OptimismL2Transactor{contract: contract}, OptimismL2Filterer: OptimismL2Filterer{contract: contract}}, nil
}

func NewOptimismL2Caller(address common.Address, caller bind.ContractCaller) (*OptimismL2Caller, error) {
	contract, err := bindOptimismL2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismL2Caller{contract: contract}, nil
}

func NewOptimismL2Transactor(address common.Address, transactor bind.ContractTransactor) (*OptimismL2Transactor, error) {
	contract, err := bindOptimismL2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismL2Transactor{contract: contract}, nil
}

func NewOptimismL2Filterer(address common.Address, filterer bind.ContractFilterer) (*OptimismL2Filterer, error) {
	contract, err := bindOptimismL2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimismL2Filterer{contract: contract}, nil
}

func bindOptimismL2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimismL2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OptimismL2 *OptimismL2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismL2.Contract.OptimismL2Caller.contract.Call(opts, result, method, params...)
}

func (_OptimismL2 *OptimismL2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismL2.Contract.OptimismL2Transactor.contract.Transfer(opts)
}

func (_OptimismL2 *OptimismL2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismL2.Contract.OptimismL2Transactor.contract.Transact(opts, method, params...)
}

func (_OptimismL2 *OptimismL2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismL2.Contract.contract.Call(opts, result, method, params...)
}

func (_OptimismL2 *OptimismL2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismL2.Contract.contract.Transfer(opts)
}

func (_OptimismL2 *OptimismL2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismL2.Contract.contract.Transact(opts, method, params...)
}

func (_OptimismL2 *OptimismL2Caller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismL2.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismL2 *OptimismL2Session) MESSENGER() (common.Address, error) {
	return _OptimismL2.Contract.MESSENGER(&_OptimismL2.CallOpts)
}

func (_OptimismL2 *OptimismL2CallerSession) MESSENGER() (common.Address, error) {
	return _OptimismL2.Contract.MESSENGER(&_OptimismL2.CallOpts)
}

func (_OptimismL2 *OptimismL2Caller) OTHERBRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismL2.contract.Call(opts, &out, "OTHER_BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismL2 *OptimismL2Session) OTHERBRIDGE() (common.Address, error) {
	return _OptimismL2.Contract.OTHERBRIDGE(&_OptimismL2.CallOpts)
}

func (_OptimismL2 *OptimismL2CallerSession) OTHERBRIDGE() (common.Address, error) {
	return _OptimismL2.Contract.OTHERBRIDGE(&_OptimismL2.CallOpts)
}

func (_OptimismL2 *OptimismL2Caller) Deposits(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OptimismL2.contract.Call(opts, &out, "deposits", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OptimismL2 *OptimismL2Session) Deposits(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _OptimismL2.Contract.Deposits(&_OptimismL2.CallOpts, arg0, arg1)
}

func (_OptimismL2 *OptimismL2CallerSession) Deposits(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _OptimismL2.Contract.Deposits(&_OptimismL2.CallOpts, arg0, arg1)
}

func (_OptimismL2 *OptimismL2Caller) L1TokenBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismL2.contract.Call(opts, &out, "l1TokenBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismL2 *OptimismL2Session) L1TokenBridge() (common.Address, error) {
	return _OptimismL2.Contract.L1TokenBridge(&_OptimismL2.CallOpts)
}

func (_OptimismL2 *OptimismL2CallerSession) L1TokenBridge() (common.Address, error) {
	return _OptimismL2.Contract.L1TokenBridge(&_OptimismL2.CallOpts)
}

func (_OptimismL2 *OptimismL2Caller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismL2.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismL2 *OptimismL2Session) Messenger() (common.Address, error) {
	return _OptimismL2.Contract.Messenger(&_OptimismL2.CallOpts)
}

func (_OptimismL2 *OptimismL2CallerSession) Messenger() (common.Address, error) {
	return _OptimismL2.Contract.Messenger(&_OptimismL2.CallOpts)
}

func (_OptimismL2 *OptimismL2Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OptimismL2.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OptimismL2 *OptimismL2Session) Version() (string, error) {
	return _OptimismL2.Contract.Version(&_OptimismL2.CallOpts)
}

func (_OptimismL2 *OptimismL2CallerSession) Version() (string, error) {
	return _OptimismL2.Contract.Version(&_OptimismL2.CallOpts)
}

func (_OptimismL2 *OptimismL2Transactor) BridgeERC20(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.contract.Transact(opts, "bridgeERC20", _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2Session) BridgeERC20(_localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.BridgeERC20(&_OptimismL2.TransactOpts, _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2TransactorSession) BridgeERC20(_localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.BridgeERC20(&_OptimismL2.TransactOpts, _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2Transactor) BridgeERC20To(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.contract.Transact(opts, "bridgeERC20To", _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2Session) BridgeERC20To(_localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.BridgeERC20To(&_OptimismL2.TransactOpts, _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2TransactorSession) BridgeERC20To(_localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.BridgeERC20To(&_OptimismL2.TransactOpts, _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2Transactor) BridgeETH(opts *bind.TransactOpts, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.contract.Transact(opts, "bridgeETH", _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2Session) BridgeETH(_minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.BridgeETH(&_OptimismL2.TransactOpts, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2TransactorSession) BridgeETH(_minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.BridgeETH(&_OptimismL2.TransactOpts, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2Transactor) BridgeETHTo(opts *bind.TransactOpts, _to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.contract.Transact(opts, "bridgeETHTo", _to, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2Session) BridgeETHTo(_to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.BridgeETHTo(&_OptimismL2.TransactOpts, _to, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2TransactorSession) BridgeETHTo(_to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.BridgeETHTo(&_OptimismL2.TransactOpts, _to, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2Transactor) FinalizeBridgeERC20(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.contract.Transact(opts, "finalizeBridgeERC20", _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

func (_OptimismL2 *OptimismL2Session) FinalizeBridgeERC20(_localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.FinalizeBridgeERC20(&_OptimismL2.TransactOpts, _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

func (_OptimismL2 *OptimismL2TransactorSession) FinalizeBridgeERC20(_localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.FinalizeBridgeERC20(&_OptimismL2.TransactOpts, _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

func (_OptimismL2 *OptimismL2Transactor) FinalizeBridgeETH(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.contract.Transact(opts, "finalizeBridgeETH", _from, _to, _amount, _extraData)
}

func (_OptimismL2 *OptimismL2Session) FinalizeBridgeETH(_from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.FinalizeBridgeETH(&_OptimismL2.TransactOpts, _from, _to, _amount, _extraData)
}

func (_OptimismL2 *OptimismL2TransactorSession) FinalizeBridgeETH(_from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.FinalizeBridgeETH(&_OptimismL2.TransactOpts, _from, _to, _amount, _extraData)
}

func (_OptimismL2 *OptimismL2Transactor) FinalizeDeposit(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.contract.Transact(opts, "finalizeDeposit", _l1Token, _l2Token, _from, _to, _amount, _extraData)
}

func (_OptimismL2 *OptimismL2Session) FinalizeDeposit(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.FinalizeDeposit(&_OptimismL2.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _extraData)
}

func (_OptimismL2 *OptimismL2TransactorSession) FinalizeDeposit(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.FinalizeDeposit(&_OptimismL2.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _extraData)
}

func (_OptimismL2 *OptimismL2Transactor) Withdraw(opts *bind.TransactOpts, _l2Token common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.contract.Transact(opts, "withdraw", _l2Token, _amount, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2Session) Withdraw(_l2Token common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.Withdraw(&_OptimismL2.TransactOpts, _l2Token, _amount, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2TransactorSession) Withdraw(_l2Token common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.Withdraw(&_OptimismL2.TransactOpts, _l2Token, _amount, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2Transactor) WithdrawTo(opts *bind.TransactOpts, _l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.contract.Transact(opts, "withdrawTo", _l2Token, _to, _amount, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2Session) WithdrawTo(_l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.WithdrawTo(&_OptimismL2.TransactOpts, _l2Token, _to, _amount, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2TransactorSession) WithdrawTo(_l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _OptimismL2.Contract.WithdrawTo(&_OptimismL2.TransactOpts, _l2Token, _to, _amount, _minGasLimit, _extraData)
}

func (_OptimismL2 *OptimismL2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismL2.contract.RawTransact(opts, nil)
}

func (_OptimismL2 *OptimismL2Session) Receive() (*types.Transaction, error) {
	return _OptimismL2.Contract.Receive(&_OptimismL2.TransactOpts)
}

func (_OptimismL2 *OptimismL2TransactorSession) Receive() (*types.Transaction, error) {
	return _OptimismL2.Contract.Receive(&_OptimismL2.TransactOpts)
}

type OptimismL2DepositFinalizedIterator struct {
	Event *OptimismL2DepositFinalized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismL2DepositFinalizedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismL2DepositFinalized)
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
		it.Event = new(OptimismL2DepositFinalized)
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

func (it *OptimismL2DepositFinalizedIterator) Error() error {
	return it.fail
}

func (it *OptimismL2DepositFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismL2DepositFinalized struct {
	L1Token   common.Address
	L2Token   common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log
}

func (_OptimismL2 *OptimismL2Filterer) FilterDepositFinalized(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*OptimismL2DepositFinalizedIterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OptimismL2.contract.FilterLogs(opts, "DepositFinalized", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &OptimismL2DepositFinalizedIterator{contract: _OptimismL2.contract, event: "DepositFinalized", logs: logs, sub: sub}, nil
}

func (_OptimismL2 *OptimismL2Filterer) WatchDepositFinalized(opts *bind.WatchOpts, sink chan<- *OptimismL2DepositFinalized, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OptimismL2.contract.WatchLogs(opts, "DepositFinalized", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismL2DepositFinalized)
				if err := _OptimismL2.contract.UnpackLog(event, "DepositFinalized", log); err != nil {
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

func (_OptimismL2 *OptimismL2Filterer) ParseDepositFinalized(log types.Log) (*OptimismL2DepositFinalized, error) {
	event := new(OptimismL2DepositFinalized)
	if err := _OptimismL2.contract.UnpackLog(event, "DepositFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismL2ERC20BridgeFinalizedIterator struct {
	Event *OptimismL2ERC20BridgeFinalized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismL2ERC20BridgeFinalizedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismL2ERC20BridgeFinalized)
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
		it.Event = new(OptimismL2ERC20BridgeFinalized)
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

func (it *OptimismL2ERC20BridgeFinalizedIterator) Error() error {
	return it.fail
}

func (it *OptimismL2ERC20BridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismL2ERC20BridgeFinalized struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         types.Log
}

func (_OptimismL2 *OptimismL2Filterer) FilterERC20BridgeFinalized(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*OptimismL2ERC20BridgeFinalizedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OptimismL2.contract.FilterLogs(opts, "ERC20BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &OptimismL2ERC20BridgeFinalizedIterator{contract: _OptimismL2.contract, event: "ERC20BridgeFinalized", logs: logs, sub: sub}, nil
}

func (_OptimismL2 *OptimismL2Filterer) WatchERC20BridgeFinalized(opts *bind.WatchOpts, sink chan<- *OptimismL2ERC20BridgeFinalized, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OptimismL2.contract.WatchLogs(opts, "ERC20BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismL2ERC20BridgeFinalized)
				if err := _OptimismL2.contract.UnpackLog(event, "ERC20BridgeFinalized", log); err != nil {
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

func (_OptimismL2 *OptimismL2Filterer) ParseERC20BridgeFinalized(log types.Log) (*OptimismL2ERC20BridgeFinalized, error) {
	event := new(OptimismL2ERC20BridgeFinalized)
	if err := _OptimismL2.contract.UnpackLog(event, "ERC20BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismL2ERC20BridgeInitiatedIterator struct {
	Event *OptimismL2ERC20BridgeInitiated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismL2ERC20BridgeInitiatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismL2ERC20BridgeInitiated)
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
		it.Event = new(OptimismL2ERC20BridgeInitiated)
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

func (it *OptimismL2ERC20BridgeInitiatedIterator) Error() error {
	return it.fail
}

func (it *OptimismL2ERC20BridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismL2ERC20BridgeInitiated struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         types.Log
}

func (_OptimismL2 *OptimismL2Filterer) FilterERC20BridgeInitiated(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*OptimismL2ERC20BridgeInitiatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OptimismL2.contract.FilterLogs(opts, "ERC20BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &OptimismL2ERC20BridgeInitiatedIterator{contract: _OptimismL2.contract, event: "ERC20BridgeInitiated", logs: logs, sub: sub}, nil
}

func (_OptimismL2 *OptimismL2Filterer) WatchERC20BridgeInitiated(opts *bind.WatchOpts, sink chan<- *OptimismL2ERC20BridgeInitiated, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OptimismL2.contract.WatchLogs(opts, "ERC20BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismL2ERC20BridgeInitiated)
				if err := _OptimismL2.contract.UnpackLog(event, "ERC20BridgeInitiated", log); err != nil {
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

func (_OptimismL2 *OptimismL2Filterer) ParseERC20BridgeInitiated(log types.Log) (*OptimismL2ERC20BridgeInitiated, error) {
	event := new(OptimismL2ERC20BridgeInitiated)
	if err := _OptimismL2.contract.UnpackLog(event, "ERC20BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismL2ETHBridgeFinalizedIterator struct {
	Event *OptimismL2ETHBridgeFinalized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismL2ETHBridgeFinalizedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismL2ETHBridgeFinalized)
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
		it.Event = new(OptimismL2ETHBridgeFinalized)
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

func (it *OptimismL2ETHBridgeFinalizedIterator) Error() error {
	return it.fail
}

func (it *OptimismL2ETHBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismL2ETHBridgeFinalized struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log
}

func (_OptimismL2 *OptimismL2Filterer) FilterETHBridgeFinalized(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismL2ETHBridgeFinalizedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismL2.contract.FilterLogs(opts, "ETHBridgeFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismL2ETHBridgeFinalizedIterator{contract: _OptimismL2.contract, event: "ETHBridgeFinalized", logs: logs, sub: sub}, nil
}

func (_OptimismL2 *OptimismL2Filterer) WatchETHBridgeFinalized(opts *bind.WatchOpts, sink chan<- *OptimismL2ETHBridgeFinalized, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismL2.contract.WatchLogs(opts, "ETHBridgeFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismL2ETHBridgeFinalized)
				if err := _OptimismL2.contract.UnpackLog(event, "ETHBridgeFinalized", log); err != nil {
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

func (_OptimismL2 *OptimismL2Filterer) ParseETHBridgeFinalized(log types.Log) (*OptimismL2ETHBridgeFinalized, error) {
	event := new(OptimismL2ETHBridgeFinalized)
	if err := _OptimismL2.contract.UnpackLog(event, "ETHBridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismL2ETHBridgeInitiatedIterator struct {
	Event *OptimismL2ETHBridgeInitiated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismL2ETHBridgeInitiatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismL2ETHBridgeInitiated)
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
		it.Event = new(OptimismL2ETHBridgeInitiated)
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

func (it *OptimismL2ETHBridgeInitiatedIterator) Error() error {
	return it.fail
}

func (it *OptimismL2ETHBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismL2ETHBridgeInitiated struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log
}

func (_OptimismL2 *OptimismL2Filterer) FilterETHBridgeInitiated(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismL2ETHBridgeInitiatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismL2.contract.FilterLogs(opts, "ETHBridgeInitiated", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismL2ETHBridgeInitiatedIterator{contract: _OptimismL2.contract, event: "ETHBridgeInitiated", logs: logs, sub: sub}, nil
}

func (_OptimismL2 *OptimismL2Filterer) WatchETHBridgeInitiated(opts *bind.WatchOpts, sink chan<- *OptimismL2ETHBridgeInitiated, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismL2.contract.WatchLogs(opts, "ETHBridgeInitiated", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismL2ETHBridgeInitiated)
				if err := _OptimismL2.contract.UnpackLog(event, "ETHBridgeInitiated", log); err != nil {
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

func (_OptimismL2 *OptimismL2Filterer) ParseETHBridgeInitiated(log types.Log) (*OptimismL2ETHBridgeInitiated, error) {
	event := new(OptimismL2ETHBridgeInitiated)
	if err := _OptimismL2.contract.UnpackLog(event, "ETHBridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismL2WithdrawalInitiatedIterator struct {
	Event *OptimismL2WithdrawalInitiated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismL2WithdrawalInitiatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismL2WithdrawalInitiated)
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
		it.Event = new(OptimismL2WithdrawalInitiated)
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

func (it *OptimismL2WithdrawalInitiatedIterator) Error() error {
	return it.fail
}

func (it *OptimismL2WithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismL2WithdrawalInitiated struct {
	L1Token   common.Address
	L2Token   common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log
}

func (_OptimismL2 *OptimismL2Filterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*OptimismL2WithdrawalInitiatedIterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OptimismL2.contract.FilterLogs(opts, "WithdrawalInitiated", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &OptimismL2WithdrawalInitiatedIterator{contract: _OptimismL2.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

func (_OptimismL2 *OptimismL2Filterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *OptimismL2WithdrawalInitiated, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OptimismL2.contract.WatchLogs(opts, "WithdrawalInitiated", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismL2WithdrawalInitiated)
				if err := _OptimismL2.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
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

func (_OptimismL2 *OptimismL2Filterer) ParseWithdrawalInitiated(log types.Log) (*OptimismL2WithdrawalInitiated, error) {
	event := new(OptimismL2WithdrawalInitiated)
	if err := _OptimismL2.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_OptimismL2 *OptimismL2) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OptimismL2.abi.Events["DepositFinalized"].ID:
		return _OptimismL2.ParseDepositFinalized(log)
	case _OptimismL2.abi.Events["ERC20BridgeFinalized"].ID:
		return _OptimismL2.ParseERC20BridgeFinalized(log)
	case _OptimismL2.abi.Events["ERC20BridgeInitiated"].ID:
		return _OptimismL2.ParseERC20BridgeInitiated(log)
	case _OptimismL2.abi.Events["ETHBridgeFinalized"].ID:
		return _OptimismL2.ParseETHBridgeFinalized(log)
	case _OptimismL2.abi.Events["ETHBridgeInitiated"].ID:
		return _OptimismL2.ParseETHBridgeInitiated(log)
	case _OptimismL2.abi.Events["WithdrawalInitiated"].ID:
		return _OptimismL2.ParseWithdrawalInitiated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (OptimismL2DepositFinalized) Topic() common.Hash {
	return common.HexToHash("0xb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89")
}

func (OptimismL2ERC20BridgeFinalized) Topic() common.Hash {
	return common.HexToHash("0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd")
}

func (OptimismL2ERC20BridgeInitiated) Topic() common.Hash {
	return common.HexToHash("0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf")
}

func (OptimismL2ETHBridgeFinalized) Topic() common.Hash {
	return common.HexToHash("0x31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d")
}

func (OptimismL2ETHBridgeInitiated) Topic() common.Hash {
	return common.HexToHash("0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5")
}

func (OptimismL2WithdrawalInitiated) Topic() common.Hash {
	return common.HexToHash("0x73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e")
}

func (_OptimismL2 *OptimismL2) Address() common.Address {
	return _OptimismL2.address
}

type OptimismL2Interface interface {
	MESSENGER(opts *bind.CallOpts) (common.Address, error)

	OTHERBRIDGE(opts *bind.CallOpts) (common.Address, error)

	Deposits(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error)

	L1TokenBridge(opts *bind.CallOpts) (common.Address, error)

	Messenger(opts *bind.CallOpts) (common.Address, error)

	Version(opts *bind.CallOpts) (string, error)

	BridgeERC20(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error)

	BridgeERC20To(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error)

	BridgeETH(opts *bind.TransactOpts, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error)

	BridgeETHTo(opts *bind.TransactOpts, _to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error)

	FinalizeBridgeERC20(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error)

	FinalizeBridgeETH(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error)

	FinalizeDeposit(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, _l2Token common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error)

	WithdrawTo(opts *bind.TransactOpts, _l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterDepositFinalized(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*OptimismL2DepositFinalizedIterator, error)

	WatchDepositFinalized(opts *bind.WatchOpts, sink chan<- *OptimismL2DepositFinalized, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error)

	ParseDepositFinalized(log types.Log) (*OptimismL2DepositFinalized, error)

	FilterERC20BridgeFinalized(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*OptimismL2ERC20BridgeFinalizedIterator, error)

	WatchERC20BridgeFinalized(opts *bind.WatchOpts, sink chan<- *OptimismL2ERC20BridgeFinalized, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error)

	ParseERC20BridgeFinalized(log types.Log) (*OptimismL2ERC20BridgeFinalized, error)

	FilterERC20BridgeInitiated(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*OptimismL2ERC20BridgeInitiatedIterator, error)

	WatchERC20BridgeInitiated(opts *bind.WatchOpts, sink chan<- *OptimismL2ERC20BridgeInitiated, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error)

	ParseERC20BridgeInitiated(log types.Log) (*OptimismL2ERC20BridgeInitiated, error)

	FilterETHBridgeFinalized(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismL2ETHBridgeFinalizedIterator, error)

	WatchETHBridgeFinalized(opts *bind.WatchOpts, sink chan<- *OptimismL2ETHBridgeFinalized, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseETHBridgeFinalized(log types.Log) (*OptimismL2ETHBridgeFinalized, error)

	FilterETHBridgeInitiated(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismL2ETHBridgeInitiatedIterator, error)

	WatchETHBridgeInitiated(opts *bind.WatchOpts, sink chan<- *OptimismL2ETHBridgeInitiated, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseETHBridgeInitiated(log types.Log) (*OptimismL2ETHBridgeInitiated, error)

	FilterWithdrawalInitiated(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*OptimismL2WithdrawalInitiatedIterator, error)

	WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *OptimismL2WithdrawalInitiated, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error)

	ParseWithdrawalInitiated(log types.Log) (*OptimismL2WithdrawalInitiated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
