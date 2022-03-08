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
	Bin: "0x60c06040523480156200001157600080fd5b506040516200294d3803806200294d83398101604081905262000034916200018f565b80600133806000816200008e5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c157620000c181620000e3565b505050151560f81b60805260601b6001600160601b03191660a05250620001c1565b6001600160a01b0381163314156200013e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000085565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001a257600080fd5b81516001600160a01b0381168114620001ba57600080fd5b9392505050565b60805160f81c60a05160601c612759620001f46000396000818161011b015261149e015260006104e901526127596000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80638da5cb5b11610081578063b57671661161005b578063b576716614610202578063e3d0e71214610215578063f2fde38b1461022857600080fd5b80638da5cb5b146101b1578063afcb95d7146101cf578063b1dc65a4146101ef57600080fd5b806379ba5097116100b257806379ba509714610162578063814118341461016c57806381ff70481461018157600080fd5b8063181f5a77146100ce578063583a013214610116575b600080fd5b604080518082018252601581527f4d6573736167654578656375746f7220312e302e3000000000000000000000006020820152905161010d919061223f565b60405180910390f35b61013d7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010d565b61016a61023b565b005b61017461033d565b60405161010d919061209f565b6004546002546040805163ffffffff8085168252640100000000909404909316602084015282015260600161010d565b60005473ffffffffffffffffffffffffffffffffffffffff1661013d565b60408051600181526000602082018190529181019190915260600161010d565b61016a6101fd366004611cdc565b6103ac565b61016a610210366004611f73565b610a55565b61016a610223366004611c0f565b610a64565b61016a610236366004611beb565b611449565b60015473ffffffffffffffffffffffffffffffffffffffff1633146102c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b606060078054806020026020016040519081016040528092919081815260200182805480156103a257602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610377575b5050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161040291849163ffffffff851691908e908e908190840183828082843760009201919091525061145a92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260025480825260035460ff808216602085015261010090910416928201929092529083146104d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d61746368000000000000000000000060448201526064016102b8565b6104e58b8b8b8b8b8b611552565b60007f0000000000000000000000000000000000000000000000000000000000000000156105425760028260200151836040015161052391906124e5565b61052d919061250a565b6105389060016124e5565b60ff169050610558565b60208201516105529060016124e5565b60ff1690505b8881146105c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e61747572657300000000000060448201526064016102b8565b88871461062a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e000060448201526064016102b8565b3360009081526005602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561066d5761066d61266e565b600281111561067e5761067e61266e565b905250905060028160200151600281111561069b5761069b61266e565b1480156106e257506007816000015160ff16815481106106bd576106bd6126cc565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b610748576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d6974746572000000000000000060448201526064016102b8565b50505050506000888860405161075f92919061208f565b604051908190038120610776918c90602001612073565b60405160208183030381529060405280519060200120905061079661182d565b604080518082019091526000808252602082015260005b88811015610a335760006001858884602081106107cc576107cc6126cc565b6107d991901a601b6124e5565b8d8d868181106107eb576107eb6126cc565b905060200201358c8c87818110610804576108046126cc565b9050602002013560405160008152602001604052604051610841949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610863573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff811660009081526005602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156108e3576108e361266e565b60028111156108f4576108f461266e565b90525092506001836020015160028111156109115761091161266e565b14610978576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e000060448201526064016102b8565b8251849060ff16601f811061098f5761098f6126cc565b6020020151156109fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e617475726500000000000000000000000060448201526064016102b8565b600184846000015160ff16601f8110610a1657610a166126cc565b911515602090920201525080610a2b816125d7565b9150506107ad565b5050505063ffffffff8110610a4a57610a4a612610565b505050505050505050565b610a616000808361145a565b50565b855185518560ff16601f831115610ad7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064016102b8565b60008111610b41576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016102b8565b818314610bcf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e0000000000000000000000000000000000000000000000000000000060648201526084016102b8565b610bda816003612553565b8311610c42576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016102b8565b610c4a611609565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b60065415610e3d57600654600090610ca290600190612590565b9050600060068281548110610cb957610cb96126cc565b60009182526020822001546007805473ffffffffffffffffffffffffffffffffffffffff90921693509084908110610cf357610cf36126cc565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff85811684526005909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600680549192509080610d7357610d7361269d565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690550190556007805480610ddc57610ddc61269d565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550610c88915050565b60005b8151518110156112a45760006005600084600001518481518110610e6657610e666126cc565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115610eb057610eb061266e565b14610f17576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016102b8565b6040805180820190915260ff82168152600160208201528251805160059160009185908110610f4857610f486126cc565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115610fe957610fe961266e565b021790555060009150610ff99050565b6005600084602001518481518110611013576110136126cc565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561105d5761105d61266e565b146110c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016102b8565b6040805180820190915260ff8216815260208101600281525060056000846020015184815181106110f7576110f76126cc565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156111985761119861266e565b0217905550508251805160069250839081106111b6576111b66126cc565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9093169290921790915582015180516007919083908110611232576112326126cc565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790558061129c816125d7565b915050610e40565b506040810151600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff4381168202928317855590830481169360019390926000926113369286929082169116176124bd565b92506101000a81548163ffffffff021916908363ffffffff1602179055506113954630600460009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a0015161168c565b6002819055825180516003805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560045460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598611434988b98919763ffffffff9092169690959194919391926122f7565b60405180910390a15050505050505050505050565b611451611609565b610a6181611737565b6000818060200190518101906114709190611dc1565b905060005b815181101561154b576000828281518110611492576114926126cc565b602002602001015190507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dc53c76e8260000151836020015184604001516040518463ffffffff1660e01b8152600401611505939291906120b2565b600060405180830381600087803b15801561151f57600080fd5b505af1158015611533573d6000803e3d6000fd5b50505050508080611543906125d7565b915050611475565b5050505050565b600061155f826020612553565b61156a856020612553565b611576886101446124a5565b61158091906124a5565b61158a91906124a5565b6115959060006124a5565b9050368114611600576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d61746368000000000000000060448201526064016102b8565b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461168a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016102b8565b565b6000808a8a8a8a8a8a8a8a8a6040516020016116b099989796959493929190612252565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff81163314156117b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016102b8565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604051806103e00160405280601f906020820280368337509192915050565b80516118578161272a565b919050565b600082601f83011261186d57600080fd5b8135602061188261187d8361243b565b6123ec565b80838252828201915082860187848660051b89010111156118a257600080fd5b60005b858110156118ca5781356118b88161272a565b845292840192908401906001016118a5565b5090979650505050505050565b60008083601f8401126118e957600080fd5b50813567ffffffffffffffff81111561190157600080fd5b6020830191508360208260051b850101111561191c57600080fd5b9250929050565b600082601f83011261193457600080fd5b8151602061194461187d8361243b565b80838252828201915082860187848660051b890101111561196457600080fd5b60005b858110156118ca57815161197a8161272a565b84529284019290840190600101611967565b600082601f83011261199d57600080fd5b815160206119ad61187d8361243b565b80838252828201915082860187848660051b89010111156119cd57600080fd5b60005b858110156118ca578151845292840192908401906001016119d0565b600082601f8301126119fd57600080fd5b8135611a0b61187d8261245f565b818152846020838601011115611a2057600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112611a4e57600080fd5b8151611a5c61187d8261245f565b818152846020838601011115611a7157600080fd5b611a828260208301602087016125a7565b949350505050565b600060a08284031215611a9c57600080fd5b611aa461237d565b90508151815260208201516020820152604082015160408201526060820151611acc8161272a565b6060820152608082015167ffffffffffffffff80821115611aec57600080fd5b9083019060c08286031215611b0057600080fd5b611b086123a6565b611b118361184c565b8152602083015182811115611b2557600080fd5b611b3187828601611a3d565b602083015250604083015182811115611b4957600080fd5b611b5587828601611923565b604083015250606083015182811115611b6d57600080fd5b611b798782860161198c565b606083015250611b8b6080840161184c565b608082015260a083015182811115611ba257600080fd5b611bae87828601611a3d565b60a083015250608084015250909392505050565b803567ffffffffffffffff8116811461185757600080fd5b803560ff8116811461185757600080fd5b600060208284031215611bfd57600080fd5b8135611c088161272a565b9392505050565b60008060008060008060c08789031215611c2857600080fd5b863567ffffffffffffffff80821115611c4057600080fd5b611c4c8a838b0161185c565b97506020890135915080821115611c6257600080fd5b611c6e8a838b0161185c565b9650611c7c60408a01611bda565b95506060890135915080821115611c9257600080fd5b611c9e8a838b016119ec565b9450611cac60808a01611bc2565b935060a0890135915080821115611cc257600080fd5b50611ccf89828a016119ec565b9150509295509295509295565b60008060008060008060008060e0898b031215611cf857600080fd5b606089018a811115611d0957600080fd5b8998503567ffffffffffffffff80821115611d2357600080fd5b818b0191508b601f830112611d3757600080fd5b813581811115611d4657600080fd5b8c6020828501011115611d5857600080fd5b6020830199508098505060808b0135915080821115611d7657600080fd5b611d828c838d016118d7565b909750955060a08b0135915080821115611d9b57600080fd5b50611da88b828c016118d7565b999c989b50969995989497949560c00135949350505050565b600060208284031215611dd357600080fd5b815167ffffffffffffffff811115611dea57600080fd5b83601f8285010112611dfb57600080fd5b80830151611e0b61187d8261243b565b808282526020820191506020848701018760208560051b878a0101011115611e3257600080fd5b60005b848110156118ca57815167ffffffffffffffff811115611e5457600080fd5b888701016060818b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0011215611e8a57600080fd5b611e926123c9565b602082015167ffffffffffffffff811115611eac57600080fd5b8b603f8285010112611ebd57600080fd5b60208184010151611ed061187d8261243b565b808282526020820191506040848701018f60408560051b878a0101011115611ef757600080fd5b600094505b83851015611f1b57805183526001949094019360209283019201611efc565b508452505050604082015167ffffffffffffffff811115611f3b57600080fd5b611f4a8c602083860101611a8a565b602083810191909152606093909301516040830152508552938401939190910190600101611e35565b600060208284031215611f8557600080fd5b813567ffffffffffffffff811115611f9c57600080fd5b611a82848285016119ec565b600081518084526020808501945080840160005b83811015611fee57815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611fbc565b509495945050505050565b600081518084526020808501945080840160005b83811015611fee5781518752958201959082019060010161200d565b600081518084526120418160208601602086016125a7565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b8281526060826020830137600060809190910190815292915050565b8183823760009101908152919050565b602081526000611c086020830184611fa8565b606080825284519082018190526000906020906080840190828801845b828110156120eb578151845292840192908401906001016120cf565b5050508381038285015285518152818601518282015260408601516040820152606086015173ffffffffffffffffffffffffffffffffffffffff80821660608401526080880151915060a060808401528082511660a08401528382015160c08085015261215c610160850182612029565b60408401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60868303810160e0880152815180845291880193506000929091908801905b808410156121c2578451861682529388019360019390930192908801906121a0565b506060860151975081878203016101008801526121df8189611ff9565b9750506080850151935061220c61012087018573ffffffffffffffffffffffffffffffffffffffff169052565b60a0850151945080868803016101408701525050505061222c8382612029565b9350505050826040830152949350505050565b602081526000611c086020830184612029565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526122998285018b611fa8565b915083820360808501526122ad828a611fa8565b915060ff881660a085015283820360c08501526122ca8288612029565b90861660e085015283810361010085015290506122e78185612029565b9c9b505050505050505050505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526123278184018a611fa8565b9050828103608084015261233b8189611fa8565b905060ff871660a084015282810360c08401526123588187612029565b905067ffffffffffffffff851660e08401528281036101008401526122e78185612029565b60405160a0810167ffffffffffffffff811182821017156123a0576123a06126fb565b60405290565b60405160c0810167ffffffffffffffff811182821017156123a0576123a06126fb565b6040516060810167ffffffffffffffff811182821017156123a0576123a06126fb565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612433576124336126fb565b604052919050565b600067ffffffffffffffff821115612455576124556126fb565b5060051b60200190565b600067ffffffffffffffff821115612479576124796126fb565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082198211156124b8576124b861263f565b500190565b600063ffffffff8083168185168083038211156124dc576124dc61263f565b01949350505050565b600060ff821660ff84168060ff038211156125025761250261263f565b019392505050565b600060ff831680612544577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8060ff84160491505092915050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561258b5761258b61263f565b500290565b6000828210156125a2576125a261263f565b500390565b60005b838110156125c25781810151838201526020016125aa565b838111156125d1576000848401525b50505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156126095761260961263f565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff81168114610a6157600080fdfea164736f6c6343000806000a",
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
