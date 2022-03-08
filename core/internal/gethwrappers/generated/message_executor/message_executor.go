// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package message_executor

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

var MessageExecutorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_offRamp\",\"outputs\":[{\"internalType\":\"contractOffRampInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620028ec380380620028ec83398101604081905262000034916200018d565b600133806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c081620000e1565b505050151560f81b60805260601b6001600160601b03191660a052620001bf565b6001600160a01b0381163314156200013c5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001a057600080fd5b81516001600160a01b0381168114620001b857600080fd5b9392505050565b60805160f81c60a05160601c6126fa620001f2600039600081816101100152611474015260006104cb01526126fa6000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c80638da5cb5b11610076578063b1dc65a41161005b578063b1dc65a4146101e4578063e3d0e712146101f7578063f2fde38b1461020a57600080fd5b80638da5cb5b146101a6578063afcb95d7146101c457600080fd5b806379ba5097116100a757806379ba509714610157578063814118341461016157806381ff70481461017657600080fd5b8063181f5a77146100c3578063583a01321461010b575b600080fd5b604080518082018252601581527f4d6573736167654578656375746f7220312e302e3000000000000000000000006020820152905161010291906121e0565b60405180910390f35b6101327f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610102565b61015f61021d565b005b61016961031f565b6040516101029190612040565b6004546002546040805163ffffffff80851682526401000000009094049093166020840152820152606001610102565b60005473ffffffffffffffffffffffffffffffffffffffff16610132565b604080516001815260006020820181905291810191909152606001610102565b61015f6101f2366004611cb2565b61038e565b61015f610205366004611be5565b610a37565b61015f610218366004611bc1565b61141c565b60015473ffffffffffffffffffffffffffffffffffffffff1633146102a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600780548060200260200160405190810160405280929190818152602001828054801561038457602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610359575b5050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916103e491849163ffffffff851691908e908e908190840183828082843760009201919091525061143092505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260025480825260035460ff808216602085015261010090910416928201929092529083146104b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d617463680000000000000000000000604482015260640161029a565b6104c78b8b8b8b8b8b611528565b60007f000000000000000000000000000000000000000000000000000000000000000015610524576002826020015183604001516105059190612486565b61050f91906124ab565b61051a906001612486565b60ff16905061053a565b6020820151610534906001612486565b60ff1690505b8881146105a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e617475726573000000000000604482015260640161029a565b88871461060c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604482015260640161029a565b3360009081526005602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561064f5761064f61260f565b60028111156106605761066061260f565b905250905060028160200151600281111561067d5761067d61260f565b1480156106c457506007816000015160ff168154811061069f5761069f61266d565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b61072a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604482015260640161029a565b505050505060008888604051610741929190612030565b604051908190038120610758918c90602001612014565b604051602081830303815290604052805190602001209050610778611803565b604080518082019091526000808252602082015260005b88811015610a155760006001858884602081106107ae576107ae61266d565b6107bb91901a601b612486565b8d8d868181106107cd576107cd61266d565b905060200201358c8c878181106107e6576107e661266d565b9050602002013560405160008152602001604052604051610823949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610845573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff811660009081526005602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156108c5576108c561260f565b60028111156108d6576108d661260f565b90525092506001836020015160028111156108f3576108f361260f565b1461095a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e0000604482015260640161029a565b8251849060ff16601f81106109715761097161266d565b6020020151156109dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e6174757265000000000000000000000000604482015260640161029a565b600184846000015160ff16601f81106109f8576109f861266d565b911515602090920201525080610a0d81612578565b91505061078f565b5050505063ffffffff8110610a2c57610a2c6125b1565b505050505050505050565b855185518560ff16601f831115610aaa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e65727300000000000000000000000000000000604482015260640161029a565b60008111610b14576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f7369746976650000000000000000000000000000604482015260640161029a565b818314610ba2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e00000000000000000000000000000000000000000000000000000000606482015260840161029a565b610bad8160036124f4565b8311610c15576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604482015260640161029a565b610c1d6115df565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b60065415610e1057600654600090610c7590600190612531565b9050600060068281548110610c8c57610c8c61266d565b60009182526020822001546007805473ffffffffffffffffffffffffffffffffffffffff90921693509084908110610cc657610cc661266d565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff85811684526005909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600680549192509080610d4657610d4661263e565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690550190556007805480610daf57610daf61263e565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550610c5b915050565b60005b8151518110156112775760006005600084600001518481518110610e3957610e3961266d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115610e8357610e8361260f565b14610eea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e65722061646472657373000000000000000000604482015260640161029a565b6040805180820190915260ff82168152600160208201528251805160059160009185908110610f1b57610f1b61266d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115610fbc57610fbc61260f565b021790555060009150610fcc9050565b6005600084602001518481518110610fe657610fe661266d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff1660028111156110305761103061260f565b14611097576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d6974746572206164647265737300000000604482015260640161029a565b6040805180820190915260ff8216815260208101600281525060056000846020015184815181106110ca576110ca61266d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561116b5761116b61260f565b0217905550508251805160069250839081106111895761118961266d565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90931692909217909155820151805160079190839081106112055761120561266d565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790558061126f81612578565b915050610e13565b506040810151600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff43811682029283178555908304811693600193909260009261130992869290821691161761245e565b92506101000a81548163ffffffff021916908363ffffffff1602179055506113684630600460009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151611662565b6002819055825180516003805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560045460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598611407988b98919763ffffffff909216969095919491939192612298565b60405180910390a15050505050505050505050565b6114246115df565b61142d8161170d565b50565b6000818060200190518101906114469190611d97565b905060005b81518110156115215760008282815181106114685761146861266d565b602002602001015190507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dc53c76e8260000151836020015184604001516040518463ffffffff1660e01b81526004016114db93929190612053565b600060405180830381600087803b1580156114f557600080fd5b505af1158015611509573d6000803e3d6000fd5b5050505050808061151990612578565b91505061144b565b5050505050565b60006115358260206124f4565b6115408560206124f4565b61154c88610144612446565b6115569190612446565b6115609190612446565b61156b906000612446565b90503681146115d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d617463680000000000000000604482015260640161029a565b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314611660576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161029a565b565b6000808a8a8a8a8a8a8a8a8a604051602001611686999897969594939291906121f3565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff811633141561178d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161029a565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604051806103e00160405280601f906020820280368337509192915050565b805161182d816126cb565b919050565b600082601f83011261184357600080fd5b81356020611858611853836123dc565b61238d565b80838252828201915082860187848660051b890101111561187857600080fd5b60005b858110156118a057813561188e816126cb565b8452928401929084019060010161187b565b5090979650505050505050565b60008083601f8401126118bf57600080fd5b50813567ffffffffffffffff8111156118d757600080fd5b6020830191508360208260051b85010111156118f257600080fd5b9250929050565b600082601f83011261190a57600080fd5b8151602061191a611853836123dc565b80838252828201915082860187848660051b890101111561193a57600080fd5b60005b858110156118a0578151611950816126cb565b8452928401929084019060010161193d565b600082601f83011261197357600080fd5b81516020611983611853836123dc565b80838252828201915082860187848660051b89010111156119a357600080fd5b60005b858110156118a0578151845292840192908401906001016119a6565b600082601f8301126119d357600080fd5b81356119e161185382612400565b8181528460208386010111156119f657600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112611a2457600080fd5b8151611a3261185382612400565b818152846020838601011115611a4757600080fd5b611a58826020830160208701612548565b949350505050565b600060a08284031215611a7257600080fd5b611a7a61231e565b90508151815260208201516020820152604082015160408201526060820151611aa2816126cb565b6060820152608082015167ffffffffffffffff80821115611ac257600080fd5b9083019060c08286031215611ad657600080fd5b611ade612347565b611ae783611822565b8152602083015182811115611afb57600080fd5b611b0787828601611a13565b602083015250604083015182811115611b1f57600080fd5b611b2b878286016118f9565b604083015250606083015182811115611b4357600080fd5b611b4f87828601611962565b606083015250611b6160808401611822565b608082015260a083015182811115611b7857600080fd5b611b8487828601611a13565b60a083015250608084015250909392505050565b803567ffffffffffffffff8116811461182d57600080fd5b803560ff8116811461182d57600080fd5b600060208284031215611bd357600080fd5b8135611bde816126cb565b9392505050565b60008060008060008060c08789031215611bfe57600080fd5b863567ffffffffffffffff80821115611c1657600080fd5b611c228a838b01611832565b97506020890135915080821115611c3857600080fd5b611c448a838b01611832565b9650611c5260408a01611bb0565b95506060890135915080821115611c6857600080fd5b611c748a838b016119c2565b9450611c8260808a01611b98565b935060a0890135915080821115611c9857600080fd5b50611ca589828a016119c2565b9150509295509295509295565b60008060008060008060008060e0898b031215611cce57600080fd5b606089018a811115611cdf57600080fd5b8998503567ffffffffffffffff80821115611cf957600080fd5b818b0191508b601f830112611d0d57600080fd5b813581811115611d1c57600080fd5b8c6020828501011115611d2e57600080fd5b6020830199508098505060808b0135915080821115611d4c57600080fd5b611d588c838d016118ad565b909750955060a08b0135915080821115611d7157600080fd5b50611d7e8b828c016118ad565b999c989b50969995989497949560c00135949350505050565b600060208284031215611da957600080fd5b815167ffffffffffffffff811115611dc057600080fd5b83601f8285010112611dd157600080fd5b80830151611de1611853826123dc565b808282526020820191506020848701018760208560051b878a0101011115611e0857600080fd5b60005b848110156118a057815167ffffffffffffffff811115611e2a57600080fd5b888701016060818b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0011215611e6057600080fd5b611e6861236a565b602082015167ffffffffffffffff811115611e8257600080fd5b8b603f8285010112611e9357600080fd5b60208184010151611ea6611853826123dc565b808282526020820191506040848701018f60408560051b878a0101011115611ecd57600080fd5b600094505b83851015611ef157805183526001949094019360209283019201611ed2565b508452505050604082015167ffffffffffffffff811115611f1157600080fd5b611f208c602083860101611a60565b602083810191909152606093909301516040830152508552938401939190910190600101611e0b565b600081518084526020808501945080840160005b83811015611f8f57815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611f5d565b509495945050505050565b600081518084526020808501945080840160005b83811015611f8f57815187529582019590820190600101611fae565b60008151808452611fe2816020860160208601612548565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b8281526060826020830137600060809190910190815292915050565b8183823760009101908152919050565b602081526000611bde6020830184611f49565b606080825284519082018190526000906020906080840190828801845b8281101561208c57815184529284019290840190600101612070565b5050508381038285015285518152818601518282015260408601516040820152606086015173ffffffffffffffffffffffffffffffffffffffff80821660608401526080880151915060a060808401528082511660a08401528382015160c0808501526120fd610160850182611fca565b60408401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60868303810160e0880152815180845291880193506000929091908801905b8084101561216357845186168252938801936001939093019290880190612141565b506060860151975081878203016101008801526121808189611f9a565b975050608085015193506121ad61012087018573ffffffffffffffffffffffffffffffffffffffff169052565b60a085015194508086880301610140870152505050506121cd8382611fca565b9350505050826040830152949350505050565b602081526000611bde6020830184611fca565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b16604085015281606085015261223a8285018b611f49565b9150838203608085015261224e828a611f49565b915060ff881660a085015283820360c085015261226b8288611fca565b90861660e085015283810361010085015290506122888185611fca565b9c9b505050505050505050505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526122c88184018a611f49565b905082810360808401526122dc8189611f49565b905060ff871660a084015282810360c08401526122f98187611fca565b905067ffffffffffffffff851660e08401528281036101008401526122888185611fca565b60405160a0810167ffffffffffffffff811182821017156123415761234161269c565b60405290565b60405160c0810167ffffffffffffffff811182821017156123415761234161269c565b6040516060810167ffffffffffffffff811182821017156123415761234161269c565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156123d4576123d461269c565b604052919050565b600067ffffffffffffffff8211156123f6576123f661269c565b5060051b60200190565b600067ffffffffffffffff82111561241a5761241a61269c565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60008219821115612459576124596125e0565b500190565b600063ffffffff80831681851680830382111561247d5761247d6125e0565b01949350505050565b600060ff821660ff84168060ff038211156124a3576124a36125e0565b019392505050565b600060ff8316806124e5577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8060ff84160491505092915050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561252c5761252c6125e0565b500290565b600082821015612543576125436125e0565b500390565b60005b8381101561256357818101518382015260200161254b565b83811115612572576000848401525b50505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156125aa576125aa6125e0565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461142d57600080fdfea164736f6c6343000806000a",
}

var MessageExecutorABI = MessageExecutorMetaData.ABI

var MessageExecutorBin = MessageExecutorMetaData.Bin

func DeployMessageExecutor(auth *bind.TransactOpts, backend bind.ContractBackend, offRamp common.Address) (common.Address, *types.Transaction, *MessageExecutor, error) {
	parsed, err := MessageExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageExecutorBin), backend, offRamp)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageExecutor{MessageExecutorCaller: MessageExecutorCaller{contract: contract}, MessageExecutorTransactor: MessageExecutorTransactor{contract: contract}, MessageExecutorFilterer: MessageExecutorFilterer{contract: contract}}, nil
}

type MessageExecutor struct {
	address common.Address
	abi     abi.ABI
	MessageExecutorCaller
	MessageExecutorTransactor
	MessageExecutorFilterer
}

type MessageExecutorCaller struct {
	contract *bind.BoundContract
}

type MessageExecutorTransactor struct {
	contract *bind.BoundContract
}

type MessageExecutorFilterer struct {
	contract *bind.BoundContract
}

type MessageExecutorSession struct {
	Contract     *MessageExecutor
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MessageExecutorCallerSession struct {
	Contract *MessageExecutorCaller
	CallOpts bind.CallOpts
}

type MessageExecutorTransactorSession struct {
	Contract     *MessageExecutorTransactor
	TransactOpts bind.TransactOpts
}

type MessageExecutorRaw struct {
	Contract *MessageExecutor
}

type MessageExecutorCallerRaw struct {
	Contract *MessageExecutorCaller
}

type MessageExecutorTransactorRaw struct {
	Contract *MessageExecutorTransactor
}

func NewMessageExecutor(address common.Address, backend bind.ContractBackend) (*MessageExecutor, error) {
	abi, err := abi.JSON(strings.NewReader(MessageExecutorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMessageExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageExecutor{address: address, abi: abi, MessageExecutorCaller: MessageExecutorCaller{contract: contract}, MessageExecutorTransactor: MessageExecutorTransactor{contract: contract}, MessageExecutorFilterer: MessageExecutorFilterer{contract: contract}}, nil
}

func NewMessageExecutorCaller(address common.Address, caller bind.ContractCaller) (*MessageExecutorCaller, error) {
	contract, err := bindMessageExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageExecutorCaller{contract: contract}, nil
}

func NewMessageExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageExecutorTransactor, error) {
	contract, err := bindMessageExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageExecutorTransactor{contract: contract}, nil
}

func NewMessageExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageExecutorFilterer, error) {
	contract, err := bindMessageExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageExecutorFilterer{contract: contract}, nil
}

func bindMessageExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageExecutorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_MessageExecutor *MessageExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageExecutor.Contract.MessageExecutorCaller.contract.Call(opts, result, method, params...)
}

func (_MessageExecutor *MessageExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageExecutor.Contract.MessageExecutorTransactor.contract.Transfer(opts)
}

func (_MessageExecutor *MessageExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageExecutor.Contract.MessageExecutorTransactor.contract.Transact(opts, method, params...)
}

func (_MessageExecutor *MessageExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageExecutor.Contract.contract.Call(opts, result, method, params...)
}

func (_MessageExecutor *MessageExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageExecutor.Contract.contract.Transfer(opts)
}

func (_MessageExecutor *MessageExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageExecutor.Contract.contract.Transact(opts, method, params...)
}

func (_MessageExecutor *MessageExecutorCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _MessageExecutor.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_MessageExecutor *MessageExecutorSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _MessageExecutor.Contract.LatestConfigDetails(&_MessageExecutor.CallOpts)
}

func (_MessageExecutor *MessageExecutorCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _MessageExecutor.Contract.LatestConfigDetails(&_MessageExecutor.CallOpts)
}

func (_MessageExecutor *MessageExecutorCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _MessageExecutor.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_MessageExecutor *MessageExecutorSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _MessageExecutor.Contract.LatestConfigDigestAndEpoch(&_MessageExecutor.CallOpts)
}

func (_MessageExecutor *MessageExecutorCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _MessageExecutor.Contract.LatestConfigDigestAndEpoch(&_MessageExecutor.CallOpts)
}

func (_MessageExecutor *MessageExecutorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageExecutor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MessageExecutor *MessageExecutorSession) Owner() (common.Address, error) {
	return _MessageExecutor.Contract.Owner(&_MessageExecutor.CallOpts)
}

func (_MessageExecutor *MessageExecutorCallerSession) Owner() (common.Address, error) {
	return _MessageExecutor.Contract.Owner(&_MessageExecutor.CallOpts)
}

func (_MessageExecutor *MessageExecutorCaller) SOffRamp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageExecutor.contract.Call(opts, &out, "s_offRamp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MessageExecutor *MessageExecutorSession) SOffRamp() (common.Address, error) {
	return _MessageExecutor.Contract.SOffRamp(&_MessageExecutor.CallOpts)
}

func (_MessageExecutor *MessageExecutorCallerSession) SOffRamp() (common.Address, error) {
	return _MessageExecutor.Contract.SOffRamp(&_MessageExecutor.CallOpts)
}

func (_MessageExecutor *MessageExecutorCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MessageExecutor.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_MessageExecutor *MessageExecutorSession) Transmitters() ([]common.Address, error) {
	return _MessageExecutor.Contract.Transmitters(&_MessageExecutor.CallOpts)
}

func (_MessageExecutor *MessageExecutorCallerSession) Transmitters() ([]common.Address, error) {
	return _MessageExecutor.Contract.Transmitters(&_MessageExecutor.CallOpts)
}

func (_MessageExecutor *MessageExecutorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MessageExecutor.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MessageExecutor *MessageExecutorSession) TypeAndVersion() (string, error) {
	return _MessageExecutor.Contract.TypeAndVersion(&_MessageExecutor.CallOpts)
}

func (_MessageExecutor *MessageExecutorCallerSession) TypeAndVersion() (string, error) {
	return _MessageExecutor.Contract.TypeAndVersion(&_MessageExecutor.CallOpts)
}

func (_MessageExecutor *MessageExecutorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageExecutor.contract.Transact(opts, "acceptOwnership")
}

func (_MessageExecutor *MessageExecutorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MessageExecutor.Contract.AcceptOwnership(&_MessageExecutor.TransactOpts)
}

func (_MessageExecutor *MessageExecutorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MessageExecutor.Contract.AcceptOwnership(&_MessageExecutor.TransactOpts)
}

func (_MessageExecutor *MessageExecutorTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _MessageExecutor.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_MessageExecutor *MessageExecutorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _MessageExecutor.Contract.SetConfig(&_MessageExecutor.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_MessageExecutor *MessageExecutorTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _MessageExecutor.Contract.SetConfig(&_MessageExecutor.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_MessageExecutor *MessageExecutorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MessageExecutor.contract.Transact(opts, "transferOwnership", to)
}

func (_MessageExecutor *MessageExecutorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MessageExecutor.Contract.TransferOwnership(&_MessageExecutor.TransactOpts, to)
}

func (_MessageExecutor *MessageExecutorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MessageExecutor.Contract.TransferOwnership(&_MessageExecutor.TransactOpts, to)
}

func (_MessageExecutor *MessageExecutorTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _MessageExecutor.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_MessageExecutor *MessageExecutorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _MessageExecutor.Contract.Transmit(&_MessageExecutor.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_MessageExecutor *MessageExecutorTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _MessageExecutor.Contract.Transmit(&_MessageExecutor.TransactOpts, reportContext, report, rs, ss, rawVs)
}

type MessageExecutorConfigSetIterator struct {
	Event *MessageExecutorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MessageExecutorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageExecutorConfigSet)
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
		it.Event = new(MessageExecutorConfigSet)
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

func (it *MessageExecutorConfigSetIterator) Error() error {
	return it.fail
}

func (it *MessageExecutorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MessageExecutorConfigSet struct {
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

func (_MessageExecutor *MessageExecutorFilterer) FilterConfigSet(opts *bind.FilterOpts) (*MessageExecutorConfigSetIterator, error) {

	logs, sub, err := _MessageExecutor.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &MessageExecutorConfigSetIterator{contract: _MessageExecutor.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_MessageExecutor *MessageExecutorFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *MessageExecutorConfigSet) (event.Subscription, error) {

	logs, sub, err := _MessageExecutor.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MessageExecutorConfigSet)
				if err := _MessageExecutor.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_MessageExecutor *MessageExecutorFilterer) ParseConfigSet(log types.Log) (*MessageExecutorConfigSet, error) {
	event := new(MessageExecutorConfigSet)
	if err := _MessageExecutor.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MessageExecutorOwnershipTransferRequestedIterator struct {
	Event *MessageExecutorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MessageExecutorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageExecutorOwnershipTransferRequested)
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
		it.Event = new(MessageExecutorOwnershipTransferRequested)
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

func (it *MessageExecutorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *MessageExecutorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MessageExecutorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MessageExecutor *MessageExecutorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MessageExecutorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MessageExecutor.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MessageExecutorOwnershipTransferRequestedIterator{contract: _MessageExecutor.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_MessageExecutor *MessageExecutorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MessageExecutorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MessageExecutor.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MessageExecutorOwnershipTransferRequested)
				if err := _MessageExecutor.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_MessageExecutor *MessageExecutorFilterer) ParseOwnershipTransferRequested(log types.Log) (*MessageExecutorOwnershipTransferRequested, error) {
	event := new(MessageExecutorOwnershipTransferRequested)
	if err := _MessageExecutor.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MessageExecutorOwnershipTransferredIterator struct {
	Event *MessageExecutorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MessageExecutorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageExecutorOwnershipTransferred)
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
		it.Event = new(MessageExecutorOwnershipTransferred)
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

func (it *MessageExecutorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *MessageExecutorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MessageExecutorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MessageExecutor *MessageExecutorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MessageExecutorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MessageExecutor.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MessageExecutorOwnershipTransferredIterator{contract: _MessageExecutor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_MessageExecutor *MessageExecutorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageExecutorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MessageExecutor.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MessageExecutorOwnershipTransferred)
				if err := _MessageExecutor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_MessageExecutor *MessageExecutorFilterer) ParseOwnershipTransferred(log types.Log) (*MessageExecutorOwnershipTransferred, error) {
	event := new(MessageExecutorOwnershipTransferred)
	if err := _MessageExecutor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MessageExecutorTransmittedIterator struct {
	Event *MessageExecutorTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MessageExecutorTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageExecutorTransmitted)
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
		it.Event = new(MessageExecutorTransmitted)
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

func (it *MessageExecutorTransmittedIterator) Error() error {
	return it.fail
}

func (it *MessageExecutorTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MessageExecutorTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_MessageExecutor *MessageExecutorFilterer) FilterTransmitted(opts *bind.FilterOpts) (*MessageExecutorTransmittedIterator, error) {

	logs, sub, err := _MessageExecutor.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &MessageExecutorTransmittedIterator{contract: _MessageExecutor.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_MessageExecutor *MessageExecutorFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *MessageExecutorTransmitted) (event.Subscription, error) {

	logs, sub, err := _MessageExecutor.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MessageExecutorTransmitted)
				if err := _MessageExecutor.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_MessageExecutor *MessageExecutorFilterer) ParseTransmitted(log types.Log) (*MessageExecutorTransmitted, error) {
	event := new(MessageExecutorTransmitted)
	if err := _MessageExecutor.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_MessageExecutor *MessageExecutor) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MessageExecutor.abi.Events["ConfigSet"].ID:
		return _MessageExecutor.ParseConfigSet(log)
	case _MessageExecutor.abi.Events["OwnershipTransferRequested"].ID:
		return _MessageExecutor.ParseOwnershipTransferRequested(log)
	case _MessageExecutor.abi.Events["OwnershipTransferred"].ID:
		return _MessageExecutor.ParseOwnershipTransferred(log)
	case _MessageExecutor.abi.Events["Transmitted"].ID:
		return _MessageExecutor.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MessageExecutorConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (MessageExecutorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (MessageExecutorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (MessageExecutorTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (_MessageExecutor *MessageExecutor) Address() common.Address {
	return _MessageExecutor.address
}

type MessageExecutorInterface interface {
	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SOffRamp(opts *bind.CallOpts) (common.Address, error)

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*MessageExecutorConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *MessageExecutorConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*MessageExecutorConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MessageExecutorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MessageExecutorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*MessageExecutorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MessageExecutorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageExecutorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*MessageExecutorOwnershipTransferred, error)

	FilterTransmitted(opts *bind.FilterOpts) (*MessageExecutorTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *MessageExecutorTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*MessageExecutorTransmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
