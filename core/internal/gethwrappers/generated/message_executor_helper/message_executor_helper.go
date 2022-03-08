// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package message_executor_helper

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
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated"
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
)

var MessageExecutorHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_offRamp\",\"outputs\":[{\"internalType\":\"contractOffRampInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200295f3803806200295f833981016040819052620000349162000188565b80600133806000816200008e5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c157620000c181620000dc565b50505015156080526001600160a01b031660a05250620001ba565b6001600160a01b038116331415620001375760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000085565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000602082840312156200019b57600080fd5b81516001600160a01b0381168114620001b357600080fd5b9392505050565b60805160a051612778620001e76000396000818161011b015261149e015260006104e901526127786000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80638da5cb5b11610081578063b57671661161005b578063b576716614610202578063e3d0e71214610215578063f2fde38b1461022857600080fd5b80638da5cb5b146101b1578063afcb95d7146101cf578063b1dc65a4146101ef57600080fd5b806379ba5097116100b257806379ba509714610162578063814118341461016c57806381ff70481461018157600080fd5b8063181f5a77146100ce578063583a013214610116575b600080fd5b604080518082018252601581527f4d6573736167654578656375746f7220312e302e3000000000000000000000006020820152905161010d91906118db565b60405180910390f35b61013d7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010d565b61016a61023b565b005b61017461033d565b60405161010d9190611946565b6004546002546040805163ffffffff8085168252640100000000909404909316602084015282015260600161010d565b60005473ffffffffffffffffffffffffffffffffffffffff1661013d565b60408051600181526000602082018190529181019190915260600161010d565b61016a6101fd3660046119a5565b6103ac565b61016a610210366004611c13565b610a55565b61016a610223366004611d33565b610a64565b61016a610236366004611e00565b611449565b60015473ffffffffffffffffffffffffffffffffffffffff1633146102c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b606060078054806020026020016040519081016040528092919081815260200182805480156103a257602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610377575b5050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161040291849163ffffffff851691908e908e908190840183828082843760009201919091525061145a92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260025480825260035460ff808216602085015261010090910416928201929092529083146104d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d61746368000000000000000000000060448201526064016102b8565b6104e58b8b8b8b8b8b611567565b60007f000000000000000000000000000000000000000000000000000000000000000015610542576002826020015183604001516105239190611e7b565b61052d9190611ea0565b610538906001611e7b565b60ff169050610558565b6020820151610552906001611e7b565b60ff1690505b8881146105c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e61747572657300000000000060448201526064016102b8565b88871461062a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e000060448201526064016102b8565b3360009081526005602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561066d5761066d611ee9565b600281111561067e5761067e611ee9565b905250905060028160200151600281111561069b5761069b611ee9565b1480156106e257506007816000015160ff16815481106106bd576106bd611e1d565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b610748576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d6974746572000000000000000060448201526064016102b8565b50505050506000888860405161075f929190611f18565b604051908190038120610776918c90602001611f28565b604051602081830303815290604052805190602001209050610796611842565b604080518082019091526000808252602082015260005b88811015610a335760006001858884602081106107cc576107cc611e1d565b6107d991901a601b611e7b565b8d8d868181106107eb576107eb611e1d565b905060200201358c8c8781811061080457610804611e1d565b9050602002013560405160008152602001604052604051610841949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610863573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff811660009081526005602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156108e3576108e3611ee9565b60028111156108f4576108f4611ee9565b905250925060018360200151600281111561091157610911611ee9565b14610978576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e000060448201526064016102b8565b8251849060ff16601f811061098f5761098f611e1d565b6020020151156109fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e617475726500000000000000000000000060448201526064016102b8565b600184846000015160ff16601f8110610a1657610a16611e1d565b911515602090920201525080610a2b81611f44565b9150506107ad565b5050505063ffffffff8110610a4a57610a4a611f7d565b505050505050505050565b610a616000808361145a565b50565b855185518560ff16601f831115610ad7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064016102b8565b60008111610b41576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016102b8565b818314610bcf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e0000000000000000000000000000000000000000000000000000000060648201526084016102b8565b610bda816003611fac565b8311610c42576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016102b8565b610c4a61161e565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b60065415610e3d57600654600090610ca290600190611fe9565b9050600060068281548110610cb957610cb9611e1d565b60009182526020822001546007805473ffffffffffffffffffffffffffffffffffffffff90921693509084908110610cf357610cf3611e1d565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff85811684526005909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600680549192509080610d7357610d73612000565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690550190556007805480610ddc57610ddc612000565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550610c88915050565b60005b8151518110156112a45760006005600084600001518481518110610e6657610e66611e1d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115610eb057610eb0611ee9565b14610f17576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016102b8565b6040805180820190915260ff82168152600160208201528251805160059160009185908110610f4857610f48611e1d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115610fe957610fe9611ee9565b021790555060009150610ff99050565b600560008460200151848151811061101357611013611e1d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561105d5761105d611ee9565b146110c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016102b8565b6040805180820190915260ff8216815260208101600281525060056000846020015184815181106110f7576110f7611e1d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561119857611198611ee9565b0217905550508251805160069250839081106111b6576111b6611e1d565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600791908390811061123257611232611e1d565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790558061129c81611f44565b915050610e40565b506040810151600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff43811682029283178555908304811693600193909260009261133692869290821691161761202f565b92506101000a81548163ffffffff021916908363ffffffff1602179055506113954630600460009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a001516116a1565b6002819055825180516003805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560045460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598611434988b98919763ffffffff909216969095919491939192612057565b60405180910390a15050505050505050505050565b61145161161e565b610a618161174c565b6000818060200190518101906114709190612334565b905060005b815181101561156057600082828151811061149257611492611e1d565b602002602001015190507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633b8d08ef8260400151604051806040016040528085600001518152602001856020015181525060006040518463ffffffff1660e01b815260040161151a93929190612538565b600060405180830381600087803b15801561153457600080fd5b505af1158015611548573d6000803e3d6000fd5b5050505050808061155890611f44565b915050611475565b5050505050565b6000611574826020611fac565b61157f856020611fac565b61158b886101446126be565b61159591906126be565b61159f91906126be565b6115aa9060006126be565b9050368114611615576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d61746368000000000000000060448201526064016102b8565b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461169f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016102b8565b565b6000808a8a8a8a8a8a8a8a8a6040516020016116c5999897969594939291906126d6565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff81163314156117cc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016102b8565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604051806103e00160405280601f906020820280368337509192915050565b60005b8381101561187c578181015183820152602001611864565b8381111561188b576000848401525b50505050565b600081518084526118a9816020860160208601611861565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006118ee6020830184611891565b9392505050565b600081518084526020808501945080840160005b8381101561193b57815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611909565b509495945050505050565b6020815260006118ee60208301846118f5565b60008083601f84011261196b57600080fd5b50813567ffffffffffffffff81111561198357600080fd5b6020830191508360208260051b850101111561199e57600080fd5b9250929050565b60008060008060008060008060e0898b0312156119c157600080fd5b606089018a8111156119d257600080fd5b8998503567ffffffffffffffff808211156119ec57600080fd5b818b0191508b601f830112611a0057600080fd5b813581811115611a0f57600080fd5b8c6020828501011115611a2157600080fd5b6020830199508098505060808b0135915080821115611a3f57600080fd5b611a4b8c838d01611959565b909750955060a08b0135915080821115611a6457600080fd5b50611a718b828c01611959565b999c989b50969995989497949560c00135949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff81118282101715611adc57611adc611a8a565b60405290565b60405160e0810167ffffffffffffffff81118282101715611adc57611adc611a8a565b6040516060810167ffffffffffffffff81118282101715611adc57611adc611a8a565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611b6f57611b6f611a8a565b604052919050565b600067ffffffffffffffff821115611b9157611b91611a8a565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112611bce57600080fd5b8135611be1611bdc82611b77565b611b28565b818152846020838601011115611bf657600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215611c2557600080fd5b813567ffffffffffffffff811115611c3c57600080fd5b611c4884828501611bbd565b949350505050565b600067ffffffffffffffff821115611c6a57611c6a611a8a565b5060051b60200190565b73ffffffffffffffffffffffffffffffffffffffff81168114610a6157600080fd5b600082601f830112611ca757600080fd5b81356020611cb7611bdc83611c50565b82815260059290921b84018101918181019086841115611cd657600080fd5b8286015b84811015611cfa578035611ced81611c74565b8352918301918301611cda565b509695505050505050565b803560ff81168114611d1657600080fd5b919050565b803567ffffffffffffffff81168114611d1657600080fd5b60008060008060008060c08789031215611d4c57600080fd5b863567ffffffffffffffff80821115611d6457600080fd5b611d708a838b01611c96565b97506020890135915080821115611d8657600080fd5b611d928a838b01611c96565b9650611da060408a01611d05565b95506060890135915080821115611db657600080fd5b611dc28a838b01611bbd565b9450611dd060808a01611d1b565b935060a0890135915080821115611de657600080fd5b50611df389828a01611bbd565b9150509295509295509295565b600060208284031215611e1257600080fd5b81356118ee81611c74565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600060ff821660ff84168060ff03821115611e9857611e98611e4c565b019392505050565b600060ff831680611eda577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8060ff84160491505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611f7657611f76611e4c565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611fe457611fe4611e4c565b500290565b600082821015611ffb57611ffb611e4c565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600063ffffffff80831681851680830382111561204e5761204e611e4c565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526120878184018a6118f5565b9050828103608084015261209b81896118f5565b905060ff871660a084015282810360c08401526120b88187611891565b905067ffffffffffffffff851660e08401528281036101008401526120dd8185611891565b9c9b505050505050505050505050565b8051611d1681611c74565b600082601f83011261210957600080fd5b81516020612119611bdc83611c50565b82815260059290921b8401810191818101908684111561213857600080fd5b8286015b84811015611cfa57805161214f81611c74565b835291830191830161213c565b600082601f83011261216d57600080fd5b8151602061217d611bdc83611c50565b82815260059290921b8401810191818101908684111561219c57600080fd5b8286015b84811015611cfa57805183529183019183016121a0565b600082601f8301126121c857600080fd5b81516121d6611bdc82611b77565b8181528460208386010111156121eb57600080fd5b611c48826020830160208701611861565b60006080828403121561220e57600080fd5b612216611ab9565b90508151815260208201516020820152604082015161223481611c74565b6040820152606082015167ffffffffffffffff8082111561225457600080fd5b9083019060e0828603121561226857600080fd5b612270611ae2565b82518281111561227f57600080fd5b61228b878286016120f8565b8252506020830151828111156122a057600080fd5b6122ac8782860161215c565b602083015250604083015160408201526122c8606084016120ed565b60608201526122d9608084016120ed565b608082015260a0830151828111156122f057600080fd5b6122fc878286016121b7565b60a08301525060c08301518281111561231457600080fd5b612320878286016121b7565b60c083015250606084015250909392505050565b6000602080838503121561234757600080fd5b825167ffffffffffffffff8082111561235f57600080fd5b818501915085601f83011261237357600080fd5b8151612381611bdc82611c50565b81815260059190911b830184019084810190888311156123a057600080fd5b8585015b838110156124ab578051858111156123bb57600080fd5b86016060818c037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00112156123ef57600080fd5b6123f7611b05565b888201518781111561240857600080fd5b8201603f81018d1361241957600080fd5b89810151612429611bdc82611c50565b81815260059190911b8201604001908b8101908f83111561244957600080fd5b6040840193505b828410156124695783518252928c0192908c0190612450565b845250505060408201518982015260608201518781111561248957600080fd5b6124978d8b838601016121fc565b6040830152508452509186019186016123a4565b5098975050505050505050565b600081518084526020808501945080840160005b8381101561193b578151875295820195908201906001016124cc565b805160408084528151908401819052600091602091908201906060860190845b8181101561252457835183529284019291840191600101612508565b505093820151949091019390935250919050565b606081528351606082015260006020808601516080840152604086015173ffffffffffffffffffffffffffffffffffffffff80821660a086015260608801519150608060c08601526101c08501825160e0808801528181518084526101e0890191508683019350600092505b808310156125c6578351851682529286019260019290920191908601906125a4565b508585015193507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff209250828882030161010089015261260581856124b8565b935050506040830151610120870152606083015161263c61014088018273ffffffffffffffffffffffffffffffffffffffff169052565b50608083015173ffffffffffffffffffffffffffffffffffffffff1661016087015260a083015186830382016101808801526126788382611891565b92505060c0830151925080868303016101a0870152506126988183611891565b915050838103828501526126ac81876124e8565b92505050611c48604083018415159052565b600082198211156126d1576126d1611e4c565b500190565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b16604085015281606085015261271d8285018b6118f5565b91508382036080850152612731828a6118f5565b915060ff881660a085015283820360c085015261274e8288611891565b90861660e085015283810361010085015290506120dd818561189156fea164736f6c634300080c000a",
}

var MessageExecutorHelperABI = MessageExecutorHelperMetaData.ABI

var MessageExecutorHelperBin = MessageExecutorHelperMetaData.Bin

func DeployMessageExecutorHelper(auth *bind.TransactOpts, backend bind.ContractBackend, offRamp common.Address) (common.Address, *types.Transaction, *MessageExecutorHelper, error) {
	parsed, err := MessageExecutorHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageExecutorHelperBin), backend, offRamp)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageExecutorHelper{MessageExecutorHelperCaller: MessageExecutorHelperCaller{contract: contract}, MessageExecutorHelperTransactor: MessageExecutorHelperTransactor{contract: contract}, MessageExecutorHelperFilterer: MessageExecutorHelperFilterer{contract: contract}}, nil
}

type MessageExecutorHelper struct {
	address common.Address
	abi     abi.ABI
	MessageExecutorHelperCaller
	MessageExecutorHelperTransactor
	MessageExecutorHelperFilterer
}

type MessageExecutorHelperCaller struct {
	contract *bind.BoundContract
}

type MessageExecutorHelperTransactor struct {
	contract *bind.BoundContract
}

type MessageExecutorHelperFilterer struct {
	contract *bind.BoundContract
}

type MessageExecutorHelperSession struct {
	Contract     *MessageExecutorHelper
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MessageExecutorHelperCallerSession struct {
	Contract *MessageExecutorHelperCaller
	CallOpts bind.CallOpts
}

type MessageExecutorHelperTransactorSession struct {
	Contract     *MessageExecutorHelperTransactor
	TransactOpts bind.TransactOpts
}

type MessageExecutorHelperRaw struct {
	Contract *MessageExecutorHelper
}

type MessageExecutorHelperCallerRaw struct {
	Contract *MessageExecutorHelperCaller
}

type MessageExecutorHelperTransactorRaw struct {
	Contract *MessageExecutorHelperTransactor
}

func NewMessageExecutorHelper(address common.Address, backend bind.ContractBackend) (*MessageExecutorHelper, error) {
	abi, err := abi.JSON(strings.NewReader(MessageExecutorHelperABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMessageExecutorHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageExecutorHelper{address: address, abi: abi, MessageExecutorHelperCaller: MessageExecutorHelperCaller{contract: contract}, MessageExecutorHelperTransactor: MessageExecutorHelperTransactor{contract: contract}, MessageExecutorHelperFilterer: MessageExecutorHelperFilterer{contract: contract}}, nil
}

func NewMessageExecutorHelperCaller(address common.Address, caller bind.ContractCaller) (*MessageExecutorHelperCaller, error) {
	contract, err := bindMessageExecutorHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageExecutorHelperCaller{contract: contract}, nil
}

func NewMessageExecutorHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageExecutorHelperTransactor, error) {
	contract, err := bindMessageExecutorHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageExecutorHelperTransactor{contract: contract}, nil
}

func NewMessageExecutorHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageExecutorHelperFilterer, error) {
	contract, err := bindMessageExecutorHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageExecutorHelperFilterer{contract: contract}, nil
}

func bindMessageExecutorHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageExecutorHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_MessageExecutorHelper *MessageExecutorHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageExecutorHelper.Contract.MessageExecutorHelperCaller.contract.Call(opts, result, method, params...)
}

func (_MessageExecutorHelper *MessageExecutorHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageExecutorHelper.Contract.MessageExecutorHelperTransactor.contract.Transfer(opts)
}

func (_MessageExecutorHelper *MessageExecutorHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageExecutorHelper.Contract.MessageExecutorHelperTransactor.contract.Transact(opts, method, params...)
}

func (_MessageExecutorHelper *MessageExecutorHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageExecutorHelper.Contract.contract.Call(opts, result, method, params...)
}

func (_MessageExecutorHelper *MessageExecutorHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageExecutorHelper.Contract.contract.Transfer(opts)
}

func (_MessageExecutorHelper *MessageExecutorHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageExecutorHelper.Contract.contract.Transact(opts, method, params...)
}

func (_MessageExecutorHelper *MessageExecutorHelperCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _MessageExecutorHelper.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_MessageExecutorHelper *MessageExecutorHelperSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _MessageExecutorHelper.Contract.LatestConfigDetails(&_MessageExecutorHelper.CallOpts)
}

func (_MessageExecutorHelper *MessageExecutorHelperCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _MessageExecutorHelper.Contract.LatestConfigDetails(&_MessageExecutorHelper.CallOpts)
}

func (_MessageExecutorHelper *MessageExecutorHelperCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _MessageExecutorHelper.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_MessageExecutorHelper *MessageExecutorHelperSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _MessageExecutorHelper.Contract.LatestConfigDigestAndEpoch(&_MessageExecutorHelper.CallOpts)
}

func (_MessageExecutorHelper *MessageExecutorHelperCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _MessageExecutorHelper.Contract.LatestConfigDigestAndEpoch(&_MessageExecutorHelper.CallOpts)
}

func (_MessageExecutorHelper *MessageExecutorHelperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageExecutorHelper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MessageExecutorHelper *MessageExecutorHelperSession) Owner() (common.Address, error) {
	return _MessageExecutorHelper.Contract.Owner(&_MessageExecutorHelper.CallOpts)
}

func (_MessageExecutorHelper *MessageExecutorHelperCallerSession) Owner() (common.Address, error) {
	return _MessageExecutorHelper.Contract.Owner(&_MessageExecutorHelper.CallOpts)
}

func (_MessageExecutorHelper *MessageExecutorHelperCaller) SOffRamp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageExecutorHelper.contract.Call(opts, &out, "s_offRamp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MessageExecutorHelper *MessageExecutorHelperSession) SOffRamp() (common.Address, error) {
	return _MessageExecutorHelper.Contract.SOffRamp(&_MessageExecutorHelper.CallOpts)
}

func (_MessageExecutorHelper *MessageExecutorHelperCallerSession) SOffRamp() (common.Address, error) {
	return _MessageExecutorHelper.Contract.SOffRamp(&_MessageExecutorHelper.CallOpts)
}

func (_MessageExecutorHelper *MessageExecutorHelperCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MessageExecutorHelper.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_MessageExecutorHelper *MessageExecutorHelperSession) Transmitters() ([]common.Address, error) {
	return _MessageExecutorHelper.Contract.Transmitters(&_MessageExecutorHelper.CallOpts)
}

func (_MessageExecutorHelper *MessageExecutorHelperCallerSession) Transmitters() ([]common.Address, error) {
	return _MessageExecutorHelper.Contract.Transmitters(&_MessageExecutorHelper.CallOpts)
}

func (_MessageExecutorHelper *MessageExecutorHelperCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MessageExecutorHelper.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MessageExecutorHelper *MessageExecutorHelperSession) TypeAndVersion() (string, error) {
	return _MessageExecutorHelper.Contract.TypeAndVersion(&_MessageExecutorHelper.CallOpts)
}

func (_MessageExecutorHelper *MessageExecutorHelperCallerSession) TypeAndVersion() (string, error) {
	return _MessageExecutorHelper.Contract.TypeAndVersion(&_MessageExecutorHelper.CallOpts)
}

func (_MessageExecutorHelper *MessageExecutorHelperTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageExecutorHelper.contract.Transact(opts, "acceptOwnership")
}

func (_MessageExecutorHelper *MessageExecutorHelperSession) AcceptOwnership() (*types.Transaction, error) {
	return _MessageExecutorHelper.Contract.AcceptOwnership(&_MessageExecutorHelper.TransactOpts)
}

func (_MessageExecutorHelper *MessageExecutorHelperTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MessageExecutorHelper.Contract.AcceptOwnership(&_MessageExecutorHelper.TransactOpts)
}

func (_MessageExecutorHelper *MessageExecutorHelperTransactor) Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error) {
	return _MessageExecutorHelper.contract.Transact(opts, "report", executableMessages)
}

func (_MessageExecutorHelper *MessageExecutorHelperSession) Report(executableMessages []byte) (*types.Transaction, error) {
	return _MessageExecutorHelper.Contract.Report(&_MessageExecutorHelper.TransactOpts, executableMessages)
}

func (_MessageExecutorHelper *MessageExecutorHelperTransactorSession) Report(executableMessages []byte) (*types.Transaction, error) {
	return _MessageExecutorHelper.Contract.Report(&_MessageExecutorHelper.TransactOpts, executableMessages)
}

func (_MessageExecutorHelper *MessageExecutorHelperTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _MessageExecutorHelper.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_MessageExecutorHelper *MessageExecutorHelperSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _MessageExecutorHelper.Contract.SetConfig(&_MessageExecutorHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_MessageExecutorHelper *MessageExecutorHelperTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _MessageExecutorHelper.Contract.SetConfig(&_MessageExecutorHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_MessageExecutorHelper *MessageExecutorHelperTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MessageExecutorHelper.contract.Transact(opts, "transferOwnership", to)
}

func (_MessageExecutorHelper *MessageExecutorHelperSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MessageExecutorHelper.Contract.TransferOwnership(&_MessageExecutorHelper.TransactOpts, to)
}

func (_MessageExecutorHelper *MessageExecutorHelperTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MessageExecutorHelper.Contract.TransferOwnership(&_MessageExecutorHelper.TransactOpts, to)
}

func (_MessageExecutorHelper *MessageExecutorHelperTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _MessageExecutorHelper.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_MessageExecutorHelper *MessageExecutorHelperSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _MessageExecutorHelper.Contract.Transmit(&_MessageExecutorHelper.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_MessageExecutorHelper *MessageExecutorHelperTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _MessageExecutorHelper.Contract.Transmit(&_MessageExecutorHelper.TransactOpts, reportContext, report, rs, ss, rawVs)
}

type MessageExecutorHelperConfigSetIterator struct {
	Event *MessageExecutorHelperConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MessageExecutorHelperConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageExecutorHelperConfigSet)
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
		it.Event = new(MessageExecutorHelperConfigSet)
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

func (it *MessageExecutorHelperConfigSetIterator) Error() error {
	return it.fail
}

func (it *MessageExecutorHelperConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MessageExecutorHelperConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) FilterConfigSet(opts *bind.FilterOpts) (*MessageExecutorHelperConfigSetIterator, error) {

	logs, sub, err := _MessageExecutorHelper.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &MessageExecutorHelperConfigSetIterator{contract: _MessageExecutorHelper.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *MessageExecutorHelperConfigSet) (event.Subscription, error) {

	logs, sub, err := _MessageExecutorHelper.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MessageExecutorHelperConfigSet)
				if err := _MessageExecutorHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) ParseConfigSet(log types.Log) (*MessageExecutorHelperConfigSet, error) {
	event := new(MessageExecutorHelperConfigSet)
	if err := _MessageExecutorHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MessageExecutorHelperOwnershipTransferRequestedIterator struct {
	Event *MessageExecutorHelperOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MessageExecutorHelperOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageExecutorHelperOwnershipTransferRequested)
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
		it.Event = new(MessageExecutorHelperOwnershipTransferRequested)
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

func (it *MessageExecutorHelperOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *MessageExecutorHelperOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MessageExecutorHelperOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MessageExecutorHelperOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MessageExecutorHelper.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MessageExecutorHelperOwnershipTransferRequestedIterator{contract: _MessageExecutorHelper.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MessageExecutorHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MessageExecutorHelper.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MessageExecutorHelperOwnershipTransferRequested)
				if err := _MessageExecutorHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) ParseOwnershipTransferRequested(log types.Log) (*MessageExecutorHelperOwnershipTransferRequested, error) {
	event := new(MessageExecutorHelperOwnershipTransferRequested)
	if err := _MessageExecutorHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MessageExecutorHelperOwnershipTransferredIterator struct {
	Event *MessageExecutorHelperOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MessageExecutorHelperOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageExecutorHelperOwnershipTransferred)
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
		it.Event = new(MessageExecutorHelperOwnershipTransferred)
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

func (it *MessageExecutorHelperOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *MessageExecutorHelperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MessageExecutorHelperOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MessageExecutorHelperOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MessageExecutorHelper.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MessageExecutorHelperOwnershipTransferredIterator{contract: _MessageExecutorHelper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageExecutorHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MessageExecutorHelper.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MessageExecutorHelperOwnershipTransferred)
				if err := _MessageExecutorHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) ParseOwnershipTransferred(log types.Log) (*MessageExecutorHelperOwnershipTransferred, error) {
	event := new(MessageExecutorHelperOwnershipTransferred)
	if err := _MessageExecutorHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MessageExecutorHelperTransmittedIterator struct {
	Event *MessageExecutorHelperTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MessageExecutorHelperTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageExecutorHelperTransmitted)
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
		it.Event = new(MessageExecutorHelperTransmitted)
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

func (it *MessageExecutorHelperTransmittedIterator) Error() error {
	return it.fail
}

func (it *MessageExecutorHelperTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MessageExecutorHelperTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) FilterTransmitted(opts *bind.FilterOpts) (*MessageExecutorHelperTransmittedIterator, error) {

	logs, sub, err := _MessageExecutorHelper.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &MessageExecutorHelperTransmittedIterator{contract: _MessageExecutorHelper.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *MessageExecutorHelperTransmitted) (event.Subscription, error) {

	logs, sub, err := _MessageExecutorHelper.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MessageExecutorHelperTransmitted)
				if err := _MessageExecutorHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) ParseTransmitted(log types.Log) (*MessageExecutorHelperTransmitted, error) {
	event := new(MessageExecutorHelperTransmitted)
	if err := _MessageExecutorHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}
type LatestConfigDigestAndEpoch struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}

func (_MessageExecutorHelper *MessageExecutorHelper) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MessageExecutorHelper.abi.Events["ConfigSet"].ID:
		return _MessageExecutorHelper.ParseConfigSet(log)
	case _MessageExecutorHelper.abi.Events["OwnershipTransferRequested"].ID:
		return _MessageExecutorHelper.ParseOwnershipTransferRequested(log)
	case _MessageExecutorHelper.abi.Events["OwnershipTransferred"].ID:
		return _MessageExecutorHelper.ParseOwnershipTransferred(log)
	case _MessageExecutorHelper.abi.Events["Transmitted"].ID:
		return _MessageExecutorHelper.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MessageExecutorHelperConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (MessageExecutorHelperOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (MessageExecutorHelperOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (MessageExecutorHelperTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (_MessageExecutorHelper *MessageExecutorHelper) Address() common.Address {
	return _MessageExecutorHelper.address
}

type MessageExecutorHelperInterface interface {
	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SOffRamp(opts *bind.CallOpts) (common.Address, error)

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*MessageExecutorHelperConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *MessageExecutorHelperConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*MessageExecutorHelperConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MessageExecutorHelperOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MessageExecutorHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*MessageExecutorHelperOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MessageExecutorHelperOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageExecutorHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*MessageExecutorHelperOwnershipTransferred, error)

	FilterTransmitted(opts *bind.FilterOpts) (*MessageExecutorHelperTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *MessageExecutorHelperTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*MessageExecutorHelperTransmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
