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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"encodedConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmited\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_offRamp\",\"outputs\":[{\"internalType\":\"contractOffRampInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200291d3803806200291d83398101604081905262000034916200018f565b80600133806000816200008e5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c157620000c181620000e3565b505050151560f81b60805260601b6001600160601b03191660a05250620001c1565b6001600160a01b0381163314156200013e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000085565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001a257600080fd5b81516001600160a01b0381168114620001ba57600080fd5b9392505050565b60805160f81c60a05160601c612729620001f460003960008181610110015261146e015260006104b901526127296000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c80638da5cb5b11610076578063b57671661161005b578063b5767166146101d7578063e3d0e712146101ea578063f2fde38b146101fd57600080fd5b80638da5cb5b146101a6578063b1dc65a4146101c457600080fd5b806379ba5097116100a757806379ba509714610157578063814118341461016157806381ff70481461017657600080fd5b8063181f5a77146100c3578063583a01321461010b575b600080fd5b604080518082018252601581527f4d6573736167654578656375746f7220312e302e30000000000000000000000060208201529051610102919061220f565b60405180910390f35b6101327f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610102565b61015f610210565b005b610169610312565b604051610102919061206f565b6004546002546040805163ffffffff80851682526401000000009094049093166020840152820152606001610102565b60005473ffffffffffffffffffffffffffffffffffffffff16610132565b61015f6101d2366004611cac565b610381565b61015f6101e5366004611f43565b610a25565b61015f6101f8366004611bdf565b610a34565b61015f61020b366004611bbb565b611419565b60015473ffffffffffffffffffffffffffffffffffffffff163314610296576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600780548060200260200160405190810160405280929190818152602001828054801561037757602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161034c575b5050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916103d19184918491908e908e908190840183828082843760009201919091525061142a92505050565b6040805183815263ffffffff600884901c1660208201527fd78f2a94a6a9ba96eb1197c7833ce19ec0fef80881049b0bd8ced9ee533739e3910160405180910390a16040805160608101825260025480825260035460ff808216602085015261010090910416928201929092529083146104a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d617463680000000000000000000000604482015260640161028d565b6104b58b8b8b8b8b8b611522565b60007f000000000000000000000000000000000000000000000000000000000000000015610512576002826020015183604001516104f391906124b5565b6104fd91906124da565b6105089060016124b5565b60ff169050610528565b60208201516105229060016124b5565b60ff1690505b888114610591576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e617475726573000000000000604482015260640161028d565b8887146105fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604482015260640161028d565b3360009081526005602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561063d5761063d61263e565b600281111561064e5761064e61263e565b905250905060028160200151600281111561066b5761066b61263e565b1480156106b257506007816000015160ff168154811061068d5761068d61269c565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b610718576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604482015260640161028d565b50505050506000888860405161072f92919061205f565b604051908190038120610746918c90602001612043565b6040516020818303038152906040528051906020012090506107666117fd565b604080518082019091526000808252602082015260005b88811015610a0357600060018588846020811061079c5761079c61269c565b6107a991901a601b6124b5565b8d8d868181106107bb576107bb61269c565b905060200201358c8c878181106107d4576107d461269c565b9050602002013560405160008152602001604052604051610811949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610833573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff811660009081526005602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156108b3576108b361263e565b60028111156108c4576108c461263e565b90525092506001836020015160028111156108e1576108e161263e565b14610948576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e0000604482015260640161028d565b8251849060ff16601f811061095f5761095f61269c565b6020020151156109cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e6174757265000000000000000000000000604482015260640161028d565b600184846000015160ff16601f81106109e6576109e661269c565b9115156020909202015250806109fb816125a7565b91505061077d565b5050505063ffffffff8110610a1a57610a1a6125e0565b505050505050505050565b610a316000808361142a565b50565b855185518560ff16601f831115610aa7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e65727300000000000000000000000000000000604482015260640161028d565b60008111610b11576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f7369746976650000000000000000000000000000604482015260640161028d565b818314610b9f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e00000000000000000000000000000000000000000000000000000000606482015260840161028d565b610baa816003612523565b8311610c12576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604482015260640161028d565b610c1a6115d9565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b60065415610e0d57600654600090610c7290600190612560565b9050600060068281548110610c8957610c8961269c565b60009182526020822001546007805473ffffffffffffffffffffffffffffffffffffffff90921693509084908110610cc357610cc361269c565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff85811684526005909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600680549192509080610d4357610d4361266d565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690550190556007805480610dac57610dac61266d565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550610c58915050565b60005b8151518110156112745760006005600084600001518481518110610e3657610e3661269c565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115610e8057610e8061263e565b14610ee7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e65722061646472657373000000000000000000604482015260640161028d565b6040805180820190915260ff82168152600160208201528251805160059160009185908110610f1857610f1861269c565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115610fb957610fb961263e565b021790555060009150610fc99050565b6005600084602001518481518110610fe357610fe361269c565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561102d5761102d61263e565b14611094576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d6974746572206164647265737300000000604482015260640161028d565b6040805180820190915260ff8216815260208101600281525060056000846020015184815181106110c7576110c761269c565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156111685761116861263e565b0217905550508251805160069250839081106111865761118661269c565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90931692909217909155820151805160079190839081106112025761120261269c565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790558061126c816125a7565b915050610e10565b506040810151600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff43811682029283178555908304811693600193909260009261130692869290821691161761248d565b92506101000a81548163ffffffff021916908363ffffffff1602179055506113654630600460009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a0015161165c565b6002819055825180516003805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560045460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598611404988b98919763ffffffff9092169690959194919391926122c7565b60405180910390a15050505050505050505050565b6114216115d9565b610a3181611707565b6000818060200190518101906114409190611d91565b905060005b815181101561151b5760008282815181106114625761146261269c565b602002602001015190507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dc53c76e8260000151836020015184604001516040518463ffffffff1660e01b81526004016114d593929190612082565b600060405180830381600087803b1580156114ef57600080fd5b505af1158015611503573d6000803e3d6000fd5b50505050508080611513906125a7565b915050611445565b5050505050565b600061152f826020612523565b61153a856020612523565b61154688610144612475565b6115509190612475565b61155a9190612475565b611565906000612475565b90503681146115d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d617463680000000000000000604482015260640161028d565b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461165a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161028d565b565b6000808a8a8a8a8a8a8a8a8a60405160200161168099989796959493929190612222565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116331415611787576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161028d565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604051806103e00160405280601f906020820280368337509192915050565b8051611827816126fa565b919050565b600082601f83011261183d57600080fd5b8135602061185261184d8361240b565b6123bc565b80838252828201915082860187848660051b890101111561187257600080fd5b60005b8581101561189a578135611888816126fa565b84529284019290840190600101611875565b5090979650505050505050565b60008083601f8401126118b957600080fd5b50813567ffffffffffffffff8111156118d157600080fd5b6020830191508360208260051b85010111156118ec57600080fd5b9250929050565b600082601f83011261190457600080fd5b8151602061191461184d8361240b565b80838252828201915082860187848660051b890101111561193457600080fd5b60005b8581101561189a57815161194a816126fa565b84529284019290840190600101611937565b600082601f83011261196d57600080fd5b8151602061197d61184d8361240b565b80838252828201915082860187848660051b890101111561199d57600080fd5b60005b8581101561189a578151845292840192908401906001016119a0565b600082601f8301126119cd57600080fd5b81356119db61184d8261242f565b8181528460208386010111156119f057600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112611a1e57600080fd5b8151611a2c61184d8261242f565b818152846020838601011115611a4157600080fd5b611a52826020830160208701612577565b949350505050565b600060a08284031215611a6c57600080fd5b611a7461234d565b90508151815260208201516020820152604082015160408201526060820151611a9c816126fa565b6060820152608082015167ffffffffffffffff80821115611abc57600080fd5b9083019060c08286031215611ad057600080fd5b611ad8612376565b611ae18361181c565b8152602083015182811115611af557600080fd5b611b0187828601611a0d565b602083015250604083015182811115611b1957600080fd5b611b25878286016118f3565b604083015250606083015182811115611b3d57600080fd5b611b498782860161195c565b606083015250611b5b6080840161181c565b608082015260a083015182811115611b7257600080fd5b611b7e87828601611a0d565b60a083015250608084015250909392505050565b803567ffffffffffffffff8116811461182757600080fd5b803560ff8116811461182757600080fd5b600060208284031215611bcd57600080fd5b8135611bd8816126fa565b9392505050565b60008060008060008060c08789031215611bf857600080fd5b863567ffffffffffffffff80821115611c1057600080fd5b611c1c8a838b0161182c565b97506020890135915080821115611c3257600080fd5b611c3e8a838b0161182c565b9650611c4c60408a01611baa565b95506060890135915080821115611c6257600080fd5b611c6e8a838b016119bc565b9450611c7c60808a01611b92565b935060a0890135915080821115611c9257600080fd5b50611c9f89828a016119bc565b9150509295509295509295565b60008060008060008060008060e0898b031215611cc857600080fd5b606089018a811115611cd957600080fd5b8998503567ffffffffffffffff80821115611cf357600080fd5b818b0191508b601f830112611d0757600080fd5b813581811115611d1657600080fd5b8c6020828501011115611d2857600080fd5b6020830199508098505060808b0135915080821115611d4657600080fd5b611d528c838d016118a7565b909750955060a08b0135915080821115611d6b57600080fd5b50611d788b828c016118a7565b999c989b50969995989497949560c00135949350505050565b600060208284031215611da357600080fd5b815167ffffffffffffffff811115611dba57600080fd5b83601f8285010112611dcb57600080fd5b80830151611ddb61184d8261240b565b808282526020820191506020848701018760208560051b878a0101011115611e0257600080fd5b60005b8481101561189a57815167ffffffffffffffff811115611e2457600080fd5b888701016060818b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0011215611e5a57600080fd5b611e62612399565b602082015167ffffffffffffffff811115611e7c57600080fd5b8b603f8285010112611e8d57600080fd5b60208184010151611ea061184d8261240b565b808282526020820191506040848701018f60408560051b878a0101011115611ec757600080fd5b600094505b83851015611eeb57805183526001949094019360209283019201611ecc565b508452505050604082015167ffffffffffffffff811115611f0b57600080fd5b611f1a8c602083860101611a5a565b602083810191909152606093909301516040830152508552938401939190910190600101611e05565b600060208284031215611f5557600080fd5b813567ffffffffffffffff811115611f6c57600080fd5b611a52848285016119bc565b600081518084526020808501945080840160005b83811015611fbe57815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611f8c565b509495945050505050565b600081518084526020808501945080840160005b83811015611fbe57815187529582019590820190600101611fdd565b60008151808452612011816020860160208601612577565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b8281526060826020830137600060809190910190815292915050565b8183823760009101908152919050565b602081526000611bd86020830184611f78565b606080825284519082018190526000906020906080840190828801845b828110156120bb5781518452928401929084019060010161209f565b5050508381038285015285518152818601518282015260408601516040820152606086015173ffffffffffffffffffffffffffffffffffffffff80821660608401526080880151915060a060808401528082511660a08401528382015160c08085015261212c610160850182611ff9565b60408401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60868303810160e0880152815180845291880193506000929091908801905b8084101561219257845186168252938801936001939093019290880190612170565b506060860151975081878203016101008801526121af8189611fc9565b975050608085015193506121dc61012087018573ffffffffffffffffffffffffffffffffffffffff169052565b60a085015194508086880301610140870152505050506121fc8382611ff9565b9350505050826040830152949350505050565b602081526000611bd86020830184611ff9565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526122698285018b611f78565b9150838203608085015261227d828a611f78565b915060ff881660a085015283820360c085015261229a8288611ff9565b90861660e085015283810361010085015290506122b78185611ff9565b9c9b505050505050505050505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526122f78184018a611f78565b9050828103608084015261230b8189611f78565b905060ff871660a084015282810360c08401526123288187611ff9565b905067ffffffffffffffff851660e08401528281036101008401526122b78185611ff9565b60405160a0810167ffffffffffffffff81118282101715612370576123706126cb565b60405290565b60405160c0810167ffffffffffffffff81118282101715612370576123706126cb565b6040516060810167ffffffffffffffff81118282101715612370576123706126cb565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612403576124036126cb565b604052919050565b600067ffffffffffffffff821115612425576124256126cb565b5060051b60200190565b600067ffffffffffffffff821115612449576124496126cb565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082198211156124885761248861260f565b500190565b600063ffffffff8083168185168083038211156124ac576124ac61260f565b01949350505050565b600060ff821660ff84168060ff038211156124d2576124d261260f565b019392505050565b600060ff831680612514577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8060ff84160491505092915050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561255b5761255b61260f565b500290565b6000828210156125725761257261260f565b500390565b60005b8381101561259257818101518382015260200161257a565b838111156125a1576000848401525b50505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156125d9576125d961260f565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff81168114610a3157600080fdfea164736f6c6343000806000a",
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
	EncodedConfigVersion      uint64
	Encoded                   []byte
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

type MessageExecutorHelperTransmitedIterator struct {
	Event *MessageExecutorHelperTransmited

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MessageExecutorHelperTransmitedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageExecutorHelperTransmited)
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
		it.Event = new(MessageExecutorHelperTransmited)
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

func (it *MessageExecutorHelperTransmitedIterator) Error() error {
	return it.fail
}

func (it *MessageExecutorHelperTransmitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MessageExecutorHelperTransmited struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) FilterTransmited(opts *bind.FilterOpts) (*MessageExecutorHelperTransmitedIterator, error) {

	logs, sub, err := _MessageExecutorHelper.contract.FilterLogs(opts, "Transmited")
	if err != nil {
		return nil, err
	}
	return &MessageExecutorHelperTransmitedIterator{contract: _MessageExecutorHelper.contract, event: "Transmited", logs: logs, sub: sub}, nil
}

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) WatchTransmited(opts *bind.WatchOpts, sink chan<- *MessageExecutorHelperTransmited) (event.Subscription, error) {

	logs, sub, err := _MessageExecutorHelper.contract.WatchLogs(opts, "Transmited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MessageExecutorHelperTransmited)
				if err := _MessageExecutorHelper.contract.UnpackLog(event, "Transmited", log); err != nil {
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

func (_MessageExecutorHelper *MessageExecutorHelperFilterer) ParseTransmited(log types.Log) (*MessageExecutorHelperTransmited, error) {
	event := new(MessageExecutorHelperTransmited)
	if err := _MessageExecutorHelper.contract.UnpackLog(event, "Transmited", log); err != nil {
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

func (_MessageExecutorHelper *MessageExecutorHelper) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MessageExecutorHelper.abi.Events["ConfigSet"].ID:
		return _MessageExecutorHelper.ParseConfigSet(log)
	case _MessageExecutorHelper.abi.Events["OwnershipTransferRequested"].ID:
		return _MessageExecutorHelper.ParseOwnershipTransferRequested(log)
	case _MessageExecutorHelper.abi.Events["OwnershipTransferred"].ID:
		return _MessageExecutorHelper.ParseOwnershipTransferred(log)
	case _MessageExecutorHelper.abi.Events["Transmited"].ID:
		return _MessageExecutorHelper.ParseTransmited(log)

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

func (MessageExecutorHelperTransmited) Topic() common.Hash {
	return common.HexToHash("0xd78f2a94a6a9ba96eb1197c7833ce19ec0fef80881049b0bd8ced9ee533739e3")
}

func (_MessageExecutorHelper *MessageExecutorHelper) Address() common.Address {
	return _MessageExecutorHelper.address
}

type MessageExecutorHelperInterface interface {
	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

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

	FilterTransmited(opts *bind.FilterOpts) (*MessageExecutorHelperTransmitedIterator, error)

	WatchTransmited(opts *bind.WatchOpts, sink chan<- *MessageExecutorHelperTransmited) (event.Subscription, error)

	ParseTransmited(log types.Log) (*MessageExecutorHelperTransmited, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
