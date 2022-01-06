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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"encodedConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmited\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_offRamp\",\"outputs\":[{\"internalType\":\"contractOffRampInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620028ab380380620028ab83398101604081905262000034916200018d565b600133806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c081620000e1565b505050151560f81b60805260601b6001600160601b03191660a052620001bf565b6001600160a01b0381163314156200013c5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001a057600080fd5b81516001600160a01b0381168114620001b857600080fd5b9392505050565b60805160f81c60a05160601c6126ba620001f16000396000818160f501526114340152600061048b01526126ba6000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c806381ff704811610076578063b1dc65a41161005b578063b1dc65a4146101a9578063e3d0e712146101bc578063f2fde38b146101cf57600080fd5b806381ff70481461015b5780638da5cb5b1461018b57600080fd5b8063181f5a77146100a8578063583a0132146100f057806379ba50971461013c5780638141183414610146575b600080fd5b604080518082018252601581527f4d6573736167654578656375746f7220312e302e300000000000000000000000602082015290516100e791906121a0565b60405180910390f35b6101177f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e7565b6101446101e2565b005b61014e6102e4565b6040516100e79190612000565b6004546002546040805163ffffffff808516825264010000000090940490931660208401528201526060016100e7565b60005473ffffffffffffffffffffffffffffffffffffffff16610117565b6101446101b7366004611c72565b610353565b6101446101ca366004611ba5565b6109f7565b6101446101dd366004611b81565b6113dc565b60015473ffffffffffffffffffffffffffffffffffffffff163314610268576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600780548060200260200160405190810160405280929190818152602001828054801561034957602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161031e575b5050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916103a39184918491908e908e90819084018382808284376000920191909152506113f092505050565b6040805183815263ffffffff600884901c1660208201527fd78f2a94a6a9ba96eb1197c7833ce19ec0fef80881049b0bd8ced9ee533739e3910160405180910390a16040805160608101825260025480825260035460ff80821660208501526101009091041692820192909252908314610479576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d617463680000000000000000000000604482015260640161025f565b6104878b8b8b8b8b8b6114e8565b60007f0000000000000000000000000000000000000000000000000000000000000000156104e4576002826020015183604001516104c59190612446565b6104cf919061246b565b6104da906001612446565b60ff1690506104fa565b60208201516104f4906001612446565b60ff1690505b888114610563576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e617475726573000000000000604482015260640161025f565b8887146105cc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604482015260640161025f565b3360009081526005602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561060f5761060f6125cf565b6002811115610620576106206125cf565b905250905060028160200151600281111561063d5761063d6125cf565b14801561068457506007816000015160ff168154811061065f5761065f61262d565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b6106ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604482015260640161025f565b505050505060008888604051610701929190611ff0565b604051908190038120610718918c90602001611fd4565b6040516020818303038152906040528051906020012090506107386117c3565b604080518082019091526000808252602082015260005b888110156109d557600060018588846020811061076e5761076e61262d565b61077b91901a601b612446565b8d8d8681811061078d5761078d61262d565b905060200201358c8c878181106107a6576107a661262d565b90506020020135604051600081526020016040526040516107e3949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610805573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff811660009081526005602090815290849020838501909452835460ff80821685529296509294508401916101009004166002811115610885576108856125cf565b6002811115610896576108966125cf565b90525092506001836020015160028111156108b3576108b36125cf565b1461091a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e0000604482015260640161025f565b8251849060ff16601f81106109315761093161262d565b60200201511561099d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e6174757265000000000000000000000000604482015260640161025f565b600184846000015160ff16601f81106109b8576109b861262d565b9115156020909202015250806109cd81612538565b91505061074f565b5050505063ffffffff81106109ec576109ec612571565b505050505050505050565b855185518560ff16601f831115610a6a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e65727300000000000000000000000000000000604482015260640161025f565b60008111610ad4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f7369746976650000000000000000000000000000604482015260640161025f565b818314610b62576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e00000000000000000000000000000000000000000000000000000000606482015260840161025f565b610b6d8160036124b4565b8311610bd5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604482015260640161025f565b610bdd61159f565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b60065415610dd057600654600090610c35906001906124f1565b9050600060068281548110610c4c57610c4c61262d565b60009182526020822001546007805473ffffffffffffffffffffffffffffffffffffffff90921693509084908110610c8657610c8661262d565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff85811684526005909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600680549192509080610d0657610d066125fe565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690550190556007805480610d6f57610d6f6125fe565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550610c1b915050565b60005b8151518110156112375760006005600084600001518481518110610df957610df961262d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115610e4357610e436125cf565b14610eaa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e65722061646472657373000000000000000000604482015260640161025f565b6040805180820190915260ff82168152600160208201528251805160059160009185908110610edb57610edb61262d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115610f7c57610f7c6125cf565b021790555060009150610f8c9050565b6005600084602001518481518110610fa657610fa661262d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115610ff057610ff06125cf565b14611057576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d6974746572206164647265737300000000604482015260640161025f565b6040805180820190915260ff82168152602081016002815250600560008460200151848151811061108a5761108a61262d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561112b5761112b6125cf565b0217905550508251805160069250839081106111495761114961262d565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90931692909217909155820151805160079190839081106111c5576111c561262d565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790558061122f81612538565b915050610dd3565b506040810151600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff4381168202928317855590830481169360019390926000926112c992869290821691161761241e565b92506101000a81548163ffffffff021916908363ffffffff1602179055506113284630600460009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151611622565b6002819055825180516003805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560045460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05986113c7988b98919763ffffffff909216969095919491939192612258565b60405180910390a15050505050505050505050565b6113e461159f565b6113ed816116cd565b50565b6000818060200190518101906114069190611d57565b905060005b81518110156114e15760008282815181106114285761142861262d565b602002602001015190507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dc53c76e8260000151836020015184604001516040518463ffffffff1660e01b815260040161149b93929190612013565b600060405180830381600087803b1580156114b557600080fd5b505af11580156114c9573d6000803e3d6000fd5b505050505080806114d990612538565b91505061140b565b5050505050565b60006114f58260206124b4565b6115008560206124b4565b61150c88610144612406565b6115169190612406565b6115209190612406565b61152b906000612406565b9050368114611596576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d617463680000000000000000604482015260640161025f565b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314611620576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161025f565b565b6000808a8a8a8a8a8a8a8a8a604051602001611646999897969594939291906121b3565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff811633141561174d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161025f565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604051806103e00160405280601f906020820280368337509192915050565b80516117ed8161268b565b919050565b600082601f83011261180357600080fd5b813560206118186118138361239c565b61234d565b80838252828201915082860187848660051b890101111561183857600080fd5b60005b8581101561186057813561184e8161268b565b8452928401929084019060010161183b565b5090979650505050505050565b60008083601f84011261187f57600080fd5b50813567ffffffffffffffff81111561189757600080fd5b6020830191508360208260051b85010111156118b257600080fd5b9250929050565b600082601f8301126118ca57600080fd5b815160206118da6118138361239c565b80838252828201915082860187848660051b89010111156118fa57600080fd5b60005b858110156118605781516119108161268b565b845292840192908401906001016118fd565b600082601f83011261193357600080fd5b815160206119436118138361239c565b80838252828201915082860187848660051b890101111561196357600080fd5b60005b8581101561186057815184529284019290840190600101611966565b600082601f83011261199357600080fd5b81356119a1611813826123c0565b8181528460208386010111156119b657600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126119e457600080fd5b81516119f2611813826123c0565b818152846020838601011115611a0757600080fd5b611a18826020830160208701612508565b949350505050565b600060a08284031215611a3257600080fd5b611a3a6122de565b90508151815260208201516020820152604082015160408201526060820151611a628161268b565b6060820152608082015167ffffffffffffffff80821115611a8257600080fd5b9083019060c08286031215611a9657600080fd5b611a9e612307565b611aa7836117e2565b8152602083015182811115611abb57600080fd5b611ac7878286016119d3565b602083015250604083015182811115611adf57600080fd5b611aeb878286016118b9565b604083015250606083015182811115611b0357600080fd5b611b0f87828601611922565b606083015250611b21608084016117e2565b608082015260a083015182811115611b3857600080fd5b611b44878286016119d3565b60a083015250608084015250909392505050565b803567ffffffffffffffff811681146117ed57600080fd5b803560ff811681146117ed57600080fd5b600060208284031215611b9357600080fd5b8135611b9e8161268b565b9392505050565b60008060008060008060c08789031215611bbe57600080fd5b863567ffffffffffffffff80821115611bd657600080fd5b611be28a838b016117f2565b97506020890135915080821115611bf857600080fd5b611c048a838b016117f2565b9650611c1260408a01611b70565b95506060890135915080821115611c2857600080fd5b611c348a838b01611982565b9450611c4260808a01611b58565b935060a0890135915080821115611c5857600080fd5b50611c6589828a01611982565b9150509295509295509295565b60008060008060008060008060e0898b031215611c8e57600080fd5b606089018a811115611c9f57600080fd5b8998503567ffffffffffffffff80821115611cb957600080fd5b818b0191508b601f830112611ccd57600080fd5b813581811115611cdc57600080fd5b8c6020828501011115611cee57600080fd5b6020830199508098505060808b0135915080821115611d0c57600080fd5b611d188c838d0161186d565b909750955060a08b0135915080821115611d3157600080fd5b50611d3e8b828c0161186d565b999c989b50969995989497949560c00135949350505050565b600060208284031215611d6957600080fd5b815167ffffffffffffffff811115611d8057600080fd5b83601f8285010112611d9157600080fd5b80830151611da16118138261239c565b808282526020820191506020848701018760208560051b878a0101011115611dc857600080fd5b60005b8481101561186057815167ffffffffffffffff811115611dea57600080fd5b888701016060818b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0011215611e2057600080fd5b611e2861232a565b602082015167ffffffffffffffff811115611e4257600080fd5b8b603f8285010112611e5357600080fd5b60208184010151611e666118138261239c565b808282526020820191506040848701018f60408560051b878a0101011115611e8d57600080fd5b600094505b83851015611eb157805183526001949094019360209283019201611e92565b508452505050604082015167ffffffffffffffff811115611ed157600080fd5b611ee08c602083860101611a20565b602083810191909152606093909301516040830152508552938401939190910190600101611dcb565b600081518084526020808501945080840160005b83811015611f4f57815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611f1d565b509495945050505050565b600081518084526020808501945080840160005b83811015611f4f57815187529582019590820190600101611f6e565b60008151808452611fa2816020860160208601612508565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b8281526060826020830137600060809190910190815292915050565b8183823760009101908152919050565b602081526000611b9e6020830184611f09565b606080825284519082018190526000906020906080840190828801845b8281101561204c57815184529284019290840190600101612030565b5050508381038285015285518152818601518282015260408601516040820152606086015173ffffffffffffffffffffffffffffffffffffffff80821660608401526080880151915060a060808401528082511660a08401528382015160c0808501526120bd610160850182611f8a565b60408401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60868303810160e0880152815180845291880193506000929091908801905b8084101561212357845186168252938801936001939093019290880190612101565b506060860151975081878203016101008801526121408189611f5a565b9750506080850151935061216d61012087018573ffffffffffffffffffffffffffffffffffffffff169052565b60a0850151945080868803016101408701525050505061218d8382611f8a565b9350505050826040830152949350505050565b602081526000611b9e6020830184611f8a565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526121fa8285018b611f09565b9150838203608085015261220e828a611f09565b915060ff881660a085015283820360c085015261222b8288611f8a565b90861660e085015283810361010085015290506122488185611f8a565b9c9b505050505050505050505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526122888184018a611f09565b9050828103608084015261229c8189611f09565b905060ff871660a084015282810360c08401526122b98187611f8a565b905067ffffffffffffffff851660e08401528281036101008401526122488185611f8a565b60405160a0810167ffffffffffffffff811182821017156123015761230161265c565b60405290565b60405160c0810167ffffffffffffffff811182821017156123015761230161265c565b6040516060810167ffffffffffffffff811182821017156123015761230161265c565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156123945761239461265c565b604052919050565b600067ffffffffffffffff8211156123b6576123b661265c565b5060051b60200190565b600067ffffffffffffffff8211156123da576123da61265c565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60008219821115612419576124196125a0565b500190565b600063ffffffff80831681851680830382111561243d5761243d6125a0565b01949350505050565b600060ff821660ff84168060ff03821115612463576124636125a0565b019392505050565b600060ff8316806124a5577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8060ff84160491505092915050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156124ec576124ec6125a0565b500290565b600082821015612503576125036125a0565b500390565b60005b8381101561252357818101518382015260200161250b565b83811115612532576000848401525b50505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561256a5761256a6125a0565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff811681146113ed57600080fdfea164736f6c6343000806000a",
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
	EncodedConfigVersion      uint64
	Encoded                   []byte
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

type MessageExecutorTransmitedIterator struct {
	Event *MessageExecutorTransmited

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MessageExecutorTransmitedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageExecutorTransmited)
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
		it.Event = new(MessageExecutorTransmited)
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

func (it *MessageExecutorTransmitedIterator) Error() error {
	return it.fail
}

func (it *MessageExecutorTransmitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MessageExecutorTransmited struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_MessageExecutor *MessageExecutorFilterer) FilterTransmited(opts *bind.FilterOpts) (*MessageExecutorTransmitedIterator, error) {

	logs, sub, err := _MessageExecutor.contract.FilterLogs(opts, "Transmited")
	if err != nil {
		return nil, err
	}
	return &MessageExecutorTransmitedIterator{contract: _MessageExecutor.contract, event: "Transmited", logs: logs, sub: sub}, nil
}

func (_MessageExecutor *MessageExecutorFilterer) WatchTransmited(opts *bind.WatchOpts, sink chan<- *MessageExecutorTransmited) (event.Subscription, error) {

	logs, sub, err := _MessageExecutor.contract.WatchLogs(opts, "Transmited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MessageExecutorTransmited)
				if err := _MessageExecutor.contract.UnpackLog(event, "Transmited", log); err != nil {
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

func (_MessageExecutor *MessageExecutorFilterer) ParseTransmited(log types.Log) (*MessageExecutorTransmited, error) {
	event := new(MessageExecutorTransmited)
	if err := _MessageExecutor.contract.UnpackLog(event, "Transmited", log); err != nil {
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

func (_MessageExecutor *MessageExecutor) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MessageExecutor.abi.Events["ConfigSet"].ID:
		return _MessageExecutor.ParseConfigSet(log)
	case _MessageExecutor.abi.Events["OwnershipTransferRequested"].ID:
		return _MessageExecutor.ParseOwnershipTransferRequested(log)
	case _MessageExecutor.abi.Events["OwnershipTransferred"].ID:
		return _MessageExecutor.ParseOwnershipTransferred(log)
	case _MessageExecutor.abi.Events["Transmited"].ID:
		return _MessageExecutor.ParseTransmited(log)

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

func (MessageExecutorTransmited) Topic() common.Hash {
	return common.HexToHash("0xd78f2a94a6a9ba96eb1197c7833ce19ec0fef80881049b0bd8ced9ee533739e3")
}

func (_MessageExecutor *MessageExecutor) Address() common.Address {
	return _MessageExecutor.address
}

type MessageExecutorInterface interface {
	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

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

	FilterTransmited(opts *bind.FilterOpts) (*MessageExecutorTransmitedIterator, error)

	WatchTransmited(opts *bind.WatchOpts, sink chan<- *MessageExecutorTransmited) (event.Subscription, error)

	ParseTransmited(log types.Log) (*MessageExecutorTransmited, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
