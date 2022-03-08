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
	Bin: "0x60c06040523480156200001157600080fd5b50604051620028fe380380620028fe833981016040819052620000349162000186565b600133806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c081620000da565b50505015156080526001600160a01b031660a052620001b8565b6001600160a01b038116331415620001355760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000602082840312156200019957600080fd5b81516001600160a01b0381168114620001b157600080fd5b9392505050565b60805160a051612719620001e5600039600081816101100152611474015260006104cb01526127196000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c80638da5cb5b11610076578063b1dc65a41161005b578063b1dc65a4146101e4578063e3d0e712146101f7578063f2fde38b1461020a57600080fd5b80638da5cb5b146101a6578063afcb95d7146101c457600080fd5b806379ba5097116100a757806379ba509714610157578063814118341461016157806381ff70481461017657600080fd5b8063181f5a77146100c3578063583a01321461010b575b600080fd5b604080518082018252601581527f4d6573736167654578656375746f7220312e302e3000000000000000000000006020820152905161010291906118b1565b60405180910390f35b6101327f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610102565b61015f61021d565b005b61016961031f565b604051610102919061191c565b6004546002546040805163ffffffff80851682526401000000009094049093166020840152820152606001610102565b60005473ffffffffffffffffffffffffffffffffffffffff16610132565b604080516001815260006020820181905291810191909152606001610102565b61015f6101f236600461197b565b61038e565b61015f610205366004611ccc565b610a37565b61015f610218366004611d99565b61141c565b60015473ffffffffffffffffffffffffffffffffffffffff1633146102a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600780548060200260200160405190810160405280929190818152602001828054801561038457602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610359575b5050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916103e491849163ffffffff851691908e908e908190840183828082843760009201919091525061143092505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260025480825260035460ff808216602085015261010090910416928201929092529083146104b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d617463680000000000000000000000604482015260640161029a565b6104c78b8b8b8b8b8b61153d565b60007f000000000000000000000000000000000000000000000000000000000000000015610524576002826020015183604001516105059190611e14565b61050f9190611e39565b61051a906001611e14565b60ff16905061053a565b6020820151610534906001611e14565b60ff1690505b8881146105a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e617475726573000000000000604482015260640161029a565b88871461060c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604482015260640161029a565b3360009081526005602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561064f5761064f611e82565b600281111561066057610660611e82565b905250905060028160200151600281111561067d5761067d611e82565b1480156106c457506007816000015160ff168154811061069f5761069f611db6565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b61072a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604482015260640161029a565b505050505060008888604051610741929190611eb1565b604051908190038120610758918c90602001611ec1565b604051602081830303815290604052805190602001209050610778611818565b604080518082019091526000808252602082015260005b88811015610a155760006001858884602081106107ae576107ae611db6565b6107bb91901a601b611e14565b8d8d868181106107cd576107cd611db6565b905060200201358c8c878181106107e6576107e6611db6565b9050602002013560405160008152602001604052604051610823949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610845573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff811660009081526005602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156108c5576108c5611e82565b60028111156108d6576108d6611e82565b90525092506001836020015160028111156108f3576108f3611e82565b1461095a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e0000604482015260640161029a565b8251849060ff16601f811061097157610971611db6565b6020020151156109dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e6174757265000000000000000000000000604482015260640161029a565b600184846000015160ff16601f81106109f8576109f8611db6565b911515602090920201525080610a0d81611edd565b91505061078f565b5050505063ffffffff8110610a2c57610a2c611f16565b505050505050505050565b855185518560ff16601f831115610aaa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e65727300000000000000000000000000000000604482015260640161029a565b60008111610b14576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f7369746976650000000000000000000000000000604482015260640161029a565b818314610ba2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e00000000000000000000000000000000000000000000000000000000606482015260840161029a565b610bad816003611f45565b8311610c15576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604482015260640161029a565b610c1d6115f4565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b60065415610e1057600654600090610c7590600190611f82565b9050600060068281548110610c8c57610c8c611db6565b60009182526020822001546007805473ffffffffffffffffffffffffffffffffffffffff90921693509084908110610cc657610cc6611db6565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff85811684526005909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600680549192509080610d4657610d46611f99565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690550190556007805480610daf57610daf611f99565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550610c5b915050565b60005b8151518110156112775760006005600084600001518481518110610e3957610e39611db6565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115610e8357610e83611e82565b14610eea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e65722061646472657373000000000000000000604482015260640161029a565b6040805180820190915260ff82168152600160208201528251805160059160009185908110610f1b57610f1b611db6565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115610fbc57610fbc611e82565b021790555060009150610fcc9050565b6005600084602001518481518110610fe657610fe6611db6565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561103057611030611e82565b14611097576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d6974746572206164647265737300000000604482015260640161029a565b6040805180820190915260ff8216815260208101600281525060056000846020015184815181106110ca576110ca611db6565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561116b5761116b611e82565b02179055505082518051600692508390811061118957611189611db6565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600791908390811061120557611205611db6565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790558061126f81611edd565b915050610e13565b506040810151600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092611309928692908216911617611fc8565b92506101000a81548163ffffffff021916908363ffffffff1602179055506113684630600460009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151611677565b6002819055825180516003805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560045460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598611407988b98919763ffffffff909216969095919491939192611ff0565b60405180910390a15050505050505050505050565b6114246115f4565b61142d81611722565b50565b60008180602001905181019061144691906122d5565b905060005b815181101561153657600082828151811061146857611468611db6565b602002602001015190507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633b8d08ef8260400151604051806040016040528085600001518152602001856020015181525060006040518463ffffffff1660e01b81526004016114f0939291906124d9565b600060405180830381600087803b15801561150a57600080fd5b505af115801561151e573d6000803e3d6000fd5b5050505050808061152e90611edd565b91505061144b565b5050505050565b600061154a826020611f45565b611555856020611f45565b6115618861014461265f565b61156b919061265f565b611575919061265f565b61158090600061265f565b90503681146115eb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d617463680000000000000000604482015260640161029a565b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314611675576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161029a565b565b6000808a8a8a8a8a8a8a8a8a60405160200161169b99989796959493929190612677565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff81163314156117a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161029a565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604051806103e00160405280601f906020820280368337509192915050565b60005b8381101561185257818101518382015260200161183a565b83811115611861576000848401525b50505050565b6000815180845261187f816020860160208601611837565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006118c46020830184611867565b9392505050565b600081518084526020808501945080840160005b8381101561191157815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016118df565b509495945050505050565b6020815260006118c460208301846118cb565b60008083601f84011261194157600080fd5b50813567ffffffffffffffff81111561195957600080fd5b6020830191508360208260051b850101111561197457600080fd5b9250929050565b60008060008060008060008060e0898b03121561199757600080fd5b606089018a8111156119a857600080fd5b8998503567ffffffffffffffff808211156119c257600080fd5b818b0191508b601f8301126119d657600080fd5b8135818111156119e557600080fd5b8c60208285010111156119f757600080fd5b6020830199508098505060808b0135915080821115611a1557600080fd5b611a218c838d0161192f565b909750955060a08b0135915080821115611a3a57600080fd5b50611a478b828c0161192f565b999c989b50969995989497949560c00135949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff81118282101715611ab257611ab2611a60565b60405290565b60405160e0810167ffffffffffffffff81118282101715611ab257611ab2611a60565b6040516060810167ffffffffffffffff81118282101715611ab257611ab2611a60565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611b4557611b45611a60565b604052919050565b600067ffffffffffffffff821115611b6757611b67611a60565b5060051b60200190565b73ffffffffffffffffffffffffffffffffffffffff8116811461142d57600080fd5b600082601f830112611ba457600080fd5b81356020611bb9611bb483611b4d565b611afe565b82815260059290921b84018101918181019086841115611bd857600080fd5b8286015b84811015611bfc578035611bef81611b71565b8352918301918301611bdc565b509695505050505050565b803560ff81168114611c1857600080fd5b919050565b600067ffffffffffffffff821115611c3757611c37611a60565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112611c7457600080fd5b8135611c82611bb482611c1d565b818152846020838601011115611c9757600080fd5b816020850160208301376000918101602001919091529392505050565b803567ffffffffffffffff81168114611c1857600080fd5b60008060008060008060c08789031215611ce557600080fd5b863567ffffffffffffffff80821115611cfd57600080fd5b611d098a838b01611b93565b97506020890135915080821115611d1f57600080fd5b611d2b8a838b01611b93565b9650611d3960408a01611c07565b95506060890135915080821115611d4f57600080fd5b611d5b8a838b01611c63565b9450611d6960808a01611cb4565b935060a0890135915080821115611d7f57600080fd5b50611d8c89828a01611c63565b9150509295509295509295565b600060208284031215611dab57600080fd5b81356118c481611b71565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600060ff821660ff84168060ff03821115611e3157611e31611de5565b019392505050565b600060ff831680611e73577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8060ff84160491505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611f0f57611f0f611de5565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611f7d57611f7d611de5565b500290565b600082821015611f9457611f94611de5565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600063ffffffff808316818516808303821115611fe757611fe7611de5565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526120208184018a6118cb565b9050828103608084015261203481896118cb565b905060ff871660a084015282810360c08401526120518187611867565b905067ffffffffffffffff851660e08401528281036101008401526120768185611867565b9c9b505050505050505050505050565b8051611c1881611b71565b600082601f8301126120a257600080fd5b815160206120b2611bb483611b4d565b82815260059290921b840181019181810190868411156120d157600080fd5b8286015b84811015611bfc5780516120e881611b71565b83529183019183016120d5565b600082601f83011261210657600080fd5b81516020612116611bb483611b4d565b82815260059290921b8401810191818101908684111561213557600080fd5b8286015b84811015611bfc5780518352918301918301612139565b600082601f83011261216157600080fd5b815161216f611bb482611c1d565b81815284602083860101111561218457600080fd5b612195826020830160208701611837565b949350505050565b6000608082840312156121af57600080fd5b6121b7611a8f565b9050815181526020820151602082015260408201516121d581611b71565b6040820152606082015167ffffffffffffffff808211156121f557600080fd5b9083019060e0828603121561220957600080fd5b612211611ab8565b82518281111561222057600080fd5b61222c87828601612091565b82525060208301518281111561224157600080fd5b61224d878286016120f5565b6020830152506040830151604082015261226960608401612086565b606082015261227a60808401612086565b608082015260a08301518281111561229157600080fd5b61229d87828601612150565b60a08301525060c0830151828111156122b557600080fd5b6122c187828601612150565b60c083015250606084015250909392505050565b600060208083850312156122e857600080fd5b825167ffffffffffffffff8082111561230057600080fd5b818501915085601f83011261231457600080fd5b8151612322611bb482611b4d565b81815260059190911b8301840190848101908883111561234157600080fd5b8585015b8381101561244c5780518581111561235c57600080fd5b86016060818c037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001121561239057600080fd5b612398611adb565b88820151878111156123a957600080fd5b8201603f81018d136123ba57600080fd5b898101516123ca611bb482611b4d565b81815260059190911b8201604001908b8101908f8311156123ea57600080fd5b6040840193505b8284101561240a5783518252928c0192908c01906123f1565b845250505060408201518982015260608201518781111561242a57600080fd5b6124388d8b8386010161219d565b604083015250845250918601918601612345565b5098975050505050505050565b600081518084526020808501945080840160005b838110156119115781518752958201959082019060010161246d565b805160408084528151908401819052600091602091908201906060860190845b818110156124c5578351835292840192918401916001016124a9565b505093820151949091019390935250919050565b606081528351606082015260006020808601516080840152604086015173ffffffffffffffffffffffffffffffffffffffff80821660a086015260608801519150608060c08601526101c08501825160e0808801528181518084526101e0890191508683019350600092505b8083101561256757835185168252928601926001929092019190860190612545565b508585015193507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff20925082888203016101008901526125a68185612459565b93505050604083015161012087015260608301516125dd61014088018273ffffffffffffffffffffffffffffffffffffffff169052565b50608083015173ffffffffffffffffffffffffffffffffffffffff1661016087015260a083015186830382016101808801526126198382611867565b92505060c0830151925080868303016101a0870152506126398183611867565b9150508381038285015261264d8187612489565b92505050612195604083018415159052565b6000821982111561267257612672611de5565b500190565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526126be8285018b6118cb565b915083820360808501526126d2828a6118cb565b915060ff881660a085015283820360c08501526126ef8288611867565b90861660e08501528381036101008501529050612076818561186756fea164736f6c634300080c000a",
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
