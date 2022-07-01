// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package afn_contract

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

type AFNInterfaceHeartbeat struct {
	Round            *big.Int
	Timestamp        *big.Int
	CommitteeVersion *big.Int
}

var AFNContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"weightThresholdForHeartbeat\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightThresholdForBadSignal\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyVoted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"IncorrectRound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"InvalidVoter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustRecoverFromBadSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RecoveryNotNecessary\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AFNBadSignal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"parties\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goodQuorum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"badQuorum\",\"type\":\"uint256\"}],\"name\":\"AFNConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"committeeVersion\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structAFNInterface.Heartbeat\",\"name\":\"heartbeat\",\"type\":\"tuple\"}],\"name\":\"AFNHeartbeat\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"BadVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"GoodVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RecoveredFromBadSignal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBadVotersAndVotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitteeVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGoodVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"getLastGoodVoteByParticipant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastHeartbeat\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"committeeVersion\",\"type\":\"uint256\"}],\"internalType\":\"structAFNInterface.Heartbeat\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParticipants\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"getWeightByParticipant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWeightThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"good\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bad\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasBadSignal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"hasVotedBad\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"weightThresholdForHeartbeat\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightThresholdForBadSignal\",\"type\":\"uint256\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteBad\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"voteGood\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001fa438038062001fa4833981016040819052620000349162000721565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000e2565b505050620000d8848484846001806200018d60201b60201c565b5050505062000904565b336001600160a01b038216036200013c5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b845186511415806200019e57508551155b80620001a8575083155b80620001b2575082155b80620001bc575081155b80620001c6575080155b15620001e5576040516306b7c75960e31b815260040160405180910390fd5b600060038054806020026020016040519081016040528092919081815260200182805480156200023f57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831162000220575b5050505050905060005b8151811015620002ae576000600260008484815181106200026e576200026e6200080f565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000208190555080620002a6906200083b565b905062000249565b50600485905560058490556007829055620002c86200049e565b620002d2620004bc565b8651600090620002ea9060039060208b0190620005a3565b5060005b8851811015620004275760006001600160a01b03168982815181106200031857620003186200080f565b60200260200101516001600160a01b03160362000348576040516306b7c75960e31b815260040160405180910390fd5b8781815181106200035d576200035d6200080f565b6020026020010151600003620003865760405163585b926360e01b815260040160405180910390fd5b8781815181106200039b576200039b6200080f565b6020026020010151600260008b8481518110620003bc57620003bc6200080f565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002081905550878181518110620003fd57620003fd6200080f565b60200260200101518262000412919062000857565b91506200041f816200083b565b9050620002ee565b50858110806200043657508481105b1562000455576040516306b7c75960e31b815260040160405180910390fd5b7f69af5b8b5b348d6b619cb6b338b5cfd865aa9e8cedd36a4a69257a9a07ebedaa888888886040516200048c949392919062000872565b60405180910390a15050505050505050565b60068054906000620004b0836200083b565b90915550506000600c55565b6000600e8054806020026020016040519081016040528092919081815260200182805480156200051657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311620004f7575b5050505050905060005b81518110156200058a576000600d60008484815181106200054557620005456200080f565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191691151591909117905562000582816200083b565b905062000520565b506000600f819055620005a090600e906200060d565b50565b828054828255906000526020600020908101928215620005fb579160200282015b82811115620005fb57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620005c4565b506200060992915062000628565b5090565b5080546000825590600052602060002090810190620005a091905b5b8082111562000609576000815560010162000629565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156200068057620006806200063f565b604052919050565b60006001600160401b03821115620006a457620006a46200063f565b5060051b60200190565b600082601f830112620006c057600080fd5b81516020620006d9620006d38362000688565b62000655565b82815260059290921b84018101918181019086841115620006f957600080fd5b8286015b84811015620007165780518352918301918301620006fd565b509695505050505050565b600080600080608085870312156200073857600080fd5b84516001600160401b03808211156200075057600080fd5b818701915087601f8301126200076557600080fd5b8151602062000778620006d38362000688565b82815260059290921b8401810191818101908b8411156200079857600080fd5b948201945b83861015620007cf5785516001600160a01b0381168114620007bf5760008081fd5b825294820194908201906200079d565b918a0151919850909350505080821115620007e957600080fd5b50620007f887828801620006ae565b604087015160609097015195989097509350505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820162000850576200085062000825565b5060010190565b600082198211156200086d576200086d62000825565b500190565b6080808252855190820181905260009060209060a0840190828901845b82811015620008b65781516001600160a01b0316845292840192908401906001016200088f565b5050508381038285015286518082528783019183019060005b81811015620008ed57835183529284019291840191600101620008cf565b505060408501969096525050506060015292915050565b61169080620009146000396000f3fe608060405234801561001057600080fd5b50600436106101515760003560e01c8063a60e6511116100cd578063cf72b39b11610081578063f20081d111610066578063f20081d114610343578063f2fde38b1461034b578063f438c9c01461035e57600080fd5b8063cf72b39b14610325578063d08041b11461033057600080fd5b8063c1b35c09116100b2578063c1b35c09146102d1578063c3453fa514610307578063ce7460241461031d57600080fd5b8063a60e6511146102b6578063acea368b146102be57600080fd5b80635aa68ac01161012457806379ba50971161010957806379ba50971461027e5780638da5cb5b146102865780639f8743f7146102ae57600080fd5b80635aa68ac01461022557806379adb16e1461023a57600080fd5b8063181f5a77146101565780632cb145d4146101a85780632ea95371146101b2578063343157b4146101fb575b600080fd5b6101926040518060400160405280600981526020017f41464e20302e302e31000000000000000000000000000000000000000000000081525081565b60405161019f919061123a565b60405180910390f35b6101b0610379565b005b6101eb6101c03660046112d6565b73ffffffffffffffffffffffffffffffffffffffff166000908152600d602052604090205460ff1690565b604051901515815260200161019f565b6102036105b1565b604080518251815260208084015190820152918101519082015260600161019f565b61022d6105fa565b60405161019f9190611349565b6102706102483660046112d6565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205490565b60405190815260200161019f565b6101b0610669565b60005460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019f565b600654610270565b600754610270565b6101b06102cc366004611469565b610766565b6102706102df3660046112d6565b73ffffffffffffffffffffffffffffffffffffffff166000908152600b602052604090205490565b61030f61079d565b60405161019f929190611539565b6101b0610817565b60105460ff166101eb565b6101b061033e36600461155b565b6108b6565b600c54610270565b6101b06103593660046112d6565b610b28565b6004546005546040805192835260208301919091520161019f565b60105460ff16156103b6576040517fc28cc95000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336000818152600260205260408120549081900361041d576040517f669f262e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600d602052604090205460ff161561047d576040517f7c9a1cf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000818152600d6020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155600e8054918201815582527fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd0180547fffffffffffffffffffffffff000000000000000000000000000000000000000016909217909155600f80548392906105379084906115a3565b9091555050600554600f54106105ad57601080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f73907f5e30313a1ab6e1815608b22b40911f1a7decec69d5df18a2298002bacb906105a49042815260200190565b60405180910390a15b5050565b6105d560405180606001604052806000815260200160008152602001600081525090565b506040805160608101825260085481526009546020820152600a549181019190915290565b6060600380548060200260200160405190810160405280929190818152602001828054801561065f57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610634575b5050505050905090565b60015473ffffffffffffffffffffffffffffffffffffffff1633146106ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610414565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61076e610b3c565b61079784848484600654600161078491906115a3565b6007546107929060016115a3565b610bbf565b50505050565b60606000600e600f548180548060200260200160405190810160405280929190818152602001828054801561080857602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116107dd575b50505050509150915091509091565b61081f610b3c565b60105460ff1661085b576040517fe147761200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610863610f5f565b601080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f3e48434bea67b1e259c2380d289dcb6372257ab2c37bc86f0e1acf83a7b07ac090600090a1565b6006548181146108fc576040517f43a010e10000000000000000000000000000000000000000000000000000000081526004810182905260248101839052604401610414565b60105460ff1615610939576040517fc28cc95000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336000818152600260205260408120549003610999576040517f669f262e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610414565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600b60205260409020548290036109f8576040517f7c9a1cf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166000908152600b602090815260408083208590556002909152812054600c805491929091610a3e9084906115a3565b90915550506040805173ffffffffffffffffffffffffffffffffffffffff83168152602081018490527f5489e43df72470c733e49d6f7bc612d52f64600fb2801593290ec32fcf144791910160405180910390a1600454600c5410610b23576040805160608101825283815267ffffffffffffffff4216602082018190526007549282018390526008859055600955600a91909155610adb611070565b604080518251815260208084015190820152828201518183015290517f90b45dcfd48782731999668957597f8b47e29aaa1d53ef2ad07612429777bed39181900360600190a1505b505050565b610b30610b3c565b610b398161108c565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610bbd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610414565b565b84518651141580610bcf57508551155b80610bd8575083155b80610be1575082155b80610bea575081155b80610bf3575080155b15610c2a576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006003805480602002602001604051908101604052809291908181526020018280548015610c8f57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610c64575b5050505050905060005b8151811015610d1157600060026000848481518110610cba57610cba6115bb565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080610d0a906115ea565b9050610c99565b50600485905560058490556007829055610d29611070565b610d31610f5f565b8651600090610d479060039060208b0190611181565b5060005b8851811015610ed357600073ffffffffffffffffffffffffffffffffffffffff16898281518110610d7e57610d7e6115bb565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1603610dd3576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b878181518110610de557610de56115bb565b6020026020010151600003610e26576040517f585b926300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b878181518110610e3857610e386115bb565b6020026020010151600260008b8481518110610e5657610e566115bb565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550878181518110610eae57610eae6115bb565b602002602001015182610ec191906115a3565b9150610ecc816115ea565b9050610d4b565b5085811080610ee157508481105b15610f18576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f69af5b8b5b348d6b619cb6b338b5cfd865aa9e8cedd36a4a69257a9a07ebedaa88888888604051610f4d9493929190611622565b60405180910390a15050505050505050565b6000600e805480602002602001604051908101604052809291908181526020018280548015610fc457602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610f99575b5050505050905060005b815181101561105c576000600d6000848481518110610fef57610fef6115bb565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055611055816115ea565b9050610fce565b506000600f819055610b3990600e9061120b565b60068054906000611080836115ea565b90915550506000600c55565b3373ffffffffffffffffffffffffffffffffffffffff82160361110b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610414565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156111fb579160200282015b828111156111fb57825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9091161782556020909201916001909101906111a1565b50611207929150611225565b5090565b5080546000825590600052602060002090810190610b3991905b5b808211156112075760008155600101611226565b600060208083528351808285015260005b818110156112675785810183015185820160400152820161124b565b81811115611279576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146112d157600080fd5b919050565b6000602082840312156112e857600080fd5b6112f1826112ad565b9392505050565b600081518084526020808501945080840160005b8381101561133e57815173ffffffffffffffffffffffffffffffffffffffff168752958201959082019060010161130c565b509495945050505050565b6020815260006112f160208301846112f8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156113d2576113d261135c565b604052919050565b600067ffffffffffffffff8211156113f4576113f461135c565b5060051b60200190565b600082601f83011261140f57600080fd5b8135602061142461141f836113da565b61138b565b82815260059290921b8401810191818101908684111561144357600080fd5b8286015b8481101561145e5780358352918301918301611447565b509695505050505050565b6000806000806080858703121561147f57600080fd5b843567ffffffffffffffff8082111561149757600080fd5b818701915087601f8301126114ab57600080fd5b813560206114bb61141f836113da565b82815260059290921b8401810191818101908b8411156114da57600080fd5b948201945b838610156114ff576114f0866112ad565b825294820194908201906114df565b9850508801359250508082111561151557600080fd5b50611522878288016113fe565b949794965050505060408301359260600135919050565b60408152600061154c60408301856112f8565b90508260208301529392505050565b60006020828403121561156d57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156115b6576115b6611574565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361161b5761161b611574565b5060010190565b60808152600061163560808301876112f8565b82810360208481019190915286518083528782019282019060005b8181101561166c57845183529383019391830191600101611650565b50506040850196909652505050606001529291505056fea164736f6c634300080f000a",
}

var AFNContractABI = AFNContractMetaData.ABI

var AFNContractBin = AFNContractMetaData.Bin

func DeployAFNContract(auth *bind.TransactOpts, backend bind.ContractBackend, participants []common.Address, weights []*big.Int, weightThresholdForHeartbeat *big.Int, weightThresholdForBadSignal *big.Int) (common.Address, *types.Transaction, *AFNContract, error) {
	parsed, err := AFNContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AFNContractBin), backend, participants, weights, weightThresholdForHeartbeat, weightThresholdForBadSignal)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AFNContract{AFNContractCaller: AFNContractCaller{contract: contract}, AFNContractTransactor: AFNContractTransactor{contract: contract}, AFNContractFilterer: AFNContractFilterer{contract: contract}}, nil
}

type AFNContract struct {
	address common.Address
	abi     abi.ABI
	AFNContractCaller
	AFNContractTransactor
	AFNContractFilterer
}

type AFNContractCaller struct {
	contract *bind.BoundContract
}

type AFNContractTransactor struct {
	contract *bind.BoundContract
}

type AFNContractFilterer struct {
	contract *bind.BoundContract
}

type AFNContractSession struct {
	Contract     *AFNContract
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type AFNContractCallerSession struct {
	Contract *AFNContractCaller
	CallOpts bind.CallOpts
}

type AFNContractTransactorSession struct {
	Contract     *AFNContractTransactor
	TransactOpts bind.TransactOpts
}

type AFNContractRaw struct {
	Contract *AFNContract
}

type AFNContractCallerRaw struct {
	Contract *AFNContractCaller
}

type AFNContractTransactorRaw struct {
	Contract *AFNContractTransactor
}

func NewAFNContract(address common.Address, backend bind.ContractBackend) (*AFNContract, error) {
	abi, err := abi.JSON(strings.NewReader(AFNContractABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindAFNContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AFNContract{address: address, abi: abi, AFNContractCaller: AFNContractCaller{contract: contract}, AFNContractTransactor: AFNContractTransactor{contract: contract}, AFNContractFilterer: AFNContractFilterer{contract: contract}}, nil
}

func NewAFNContractCaller(address common.Address, caller bind.ContractCaller) (*AFNContractCaller, error) {
	contract, err := bindAFNContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AFNContractCaller{contract: contract}, nil
}

func NewAFNContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AFNContractTransactor, error) {
	contract, err := bindAFNContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AFNContractTransactor{contract: contract}, nil
}

func NewAFNContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AFNContractFilterer, error) {
	contract, err := bindAFNContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AFNContractFilterer{contract: contract}, nil
}

func bindAFNContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AFNContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_AFNContract *AFNContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AFNContract.Contract.AFNContractCaller.contract.Call(opts, result, method, params...)
}

func (_AFNContract *AFNContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AFNContract.Contract.AFNContractTransactor.contract.Transfer(opts)
}

func (_AFNContract *AFNContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AFNContract.Contract.AFNContractTransactor.contract.Transact(opts, method, params...)
}

func (_AFNContract *AFNContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AFNContract.Contract.contract.Call(opts, result, method, params...)
}

func (_AFNContract *AFNContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AFNContract.Contract.contract.Transfer(opts)
}

func (_AFNContract *AFNContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AFNContract.Contract.contract.Transact(opts, method, params...)
}

func (_AFNContract *AFNContractCaller) GetBadVotersAndVotes(opts *bind.CallOpts) (GetBadVotersAndVotes,

	error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getBadVotersAndVotes")

	outstruct := new(GetBadVotersAndVotes)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Voters = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Votes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_AFNContract *AFNContractSession) GetBadVotersAndVotes() (GetBadVotersAndVotes,

	error) {
	return _AFNContract.Contract.GetBadVotersAndVotes(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCallerSession) GetBadVotersAndVotes() (GetBadVotersAndVotes,

	error) {
	return _AFNContract.Contract.GetBadVotersAndVotes(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCaller) GetCommitteeVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getCommitteeVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AFNContract *AFNContractSession) GetCommitteeVersion() (*big.Int, error) {
	return _AFNContract.Contract.GetCommitteeVersion(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCallerSession) GetCommitteeVersion() (*big.Int, error) {
	return _AFNContract.Contract.GetCommitteeVersion(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCaller) GetGoodVotes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getGoodVotes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AFNContract *AFNContractSession) GetGoodVotes() (*big.Int, error) {
	return _AFNContract.Contract.GetGoodVotes(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCallerSession) GetGoodVotes() (*big.Int, error) {
	return _AFNContract.Contract.GetGoodVotes(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCaller) GetLastGoodVoteByParticipant(opts *bind.CallOpts, participant common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getLastGoodVoteByParticipant", participant)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AFNContract *AFNContractSession) GetLastGoodVoteByParticipant(participant common.Address) (*big.Int, error) {
	return _AFNContract.Contract.GetLastGoodVoteByParticipant(&_AFNContract.CallOpts, participant)
}

func (_AFNContract *AFNContractCallerSession) GetLastGoodVoteByParticipant(participant common.Address) (*big.Int, error) {
	return _AFNContract.Contract.GetLastGoodVoteByParticipant(&_AFNContract.CallOpts, participant)
}

func (_AFNContract *AFNContractCaller) GetLastHeartbeat(opts *bind.CallOpts) (AFNInterfaceHeartbeat, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getLastHeartbeat")

	if err != nil {
		return *new(AFNInterfaceHeartbeat), err
	}

	out0 := *abi.ConvertType(out[0], new(AFNInterfaceHeartbeat)).(*AFNInterfaceHeartbeat)

	return out0, err

}

func (_AFNContract *AFNContractSession) GetLastHeartbeat() (AFNInterfaceHeartbeat, error) {
	return _AFNContract.Contract.GetLastHeartbeat(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCallerSession) GetLastHeartbeat() (AFNInterfaceHeartbeat, error) {
	return _AFNContract.Contract.GetLastHeartbeat(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCaller) GetParticipants(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getParticipants")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_AFNContract *AFNContractSession) GetParticipants() ([]common.Address, error) {
	return _AFNContract.Contract.GetParticipants(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCallerSession) GetParticipants() ([]common.Address, error) {
	return _AFNContract.Contract.GetParticipants(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCaller) GetRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AFNContract *AFNContractSession) GetRound() (*big.Int, error) {
	return _AFNContract.Contract.GetRound(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCallerSession) GetRound() (*big.Int, error) {
	return _AFNContract.Contract.GetRound(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCaller) GetWeightByParticipant(opts *bind.CallOpts, participant common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getWeightByParticipant", participant)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AFNContract *AFNContractSession) GetWeightByParticipant(participant common.Address) (*big.Int, error) {
	return _AFNContract.Contract.GetWeightByParticipant(&_AFNContract.CallOpts, participant)
}

func (_AFNContract *AFNContractCallerSession) GetWeightByParticipant(participant common.Address) (*big.Int, error) {
	return _AFNContract.Contract.GetWeightByParticipant(&_AFNContract.CallOpts, participant)
}

func (_AFNContract *AFNContractCaller) GetWeightThresholds(opts *bind.CallOpts) (GetWeightThresholds,

	error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getWeightThresholds")

	outstruct := new(GetWeightThresholds)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Good = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Bad = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_AFNContract *AFNContractSession) GetWeightThresholds() (GetWeightThresholds,

	error) {
	return _AFNContract.Contract.GetWeightThresholds(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCallerSession) GetWeightThresholds() (GetWeightThresholds,

	error) {
	return _AFNContract.Contract.GetWeightThresholds(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCaller) HasBadSignal(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "hasBadSignal")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AFNContract *AFNContractSession) HasBadSignal() (bool, error) {
	return _AFNContract.Contract.HasBadSignal(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCallerSession) HasBadSignal() (bool, error) {
	return _AFNContract.Contract.HasBadSignal(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCaller) HasVotedBad(opts *bind.CallOpts, participant common.Address) (bool, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "hasVotedBad", participant)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AFNContract *AFNContractSession) HasVotedBad(participant common.Address) (bool, error) {
	return _AFNContract.Contract.HasVotedBad(&_AFNContract.CallOpts, participant)
}

func (_AFNContract *AFNContractCallerSession) HasVotedBad(participant common.Address) (bool, error) {
	return _AFNContract.Contract.HasVotedBad(&_AFNContract.CallOpts, participant)
}

func (_AFNContract *AFNContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_AFNContract *AFNContractSession) Owner() (common.Address, error) {
	return _AFNContract.Contract.Owner(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCallerSession) Owner() (common.Address, error) {
	return _AFNContract.Contract.Owner(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_AFNContract *AFNContractSession) TypeAndVersion() (string, error) {
	return _AFNContract.Contract.TypeAndVersion(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCallerSession) TypeAndVersion() (string, error) {
	return _AFNContract.Contract.TypeAndVersion(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AFNContract.contract.Transact(opts, "acceptOwnership")
}

func (_AFNContract *AFNContractSession) AcceptOwnership() (*types.Transaction, error) {
	return _AFNContract.Contract.AcceptOwnership(&_AFNContract.TransactOpts)
}

func (_AFNContract *AFNContractTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AFNContract.Contract.AcceptOwnership(&_AFNContract.TransactOpts)
}

func (_AFNContract *AFNContractTransactor) Recover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AFNContract.contract.Transact(opts, "recover")
}

func (_AFNContract *AFNContractSession) Recover() (*types.Transaction, error) {
	return _AFNContract.Contract.Recover(&_AFNContract.TransactOpts)
}

func (_AFNContract *AFNContractTransactorSession) Recover() (*types.Transaction, error) {
	return _AFNContract.Contract.Recover(&_AFNContract.TransactOpts)
}

func (_AFNContract *AFNContractTransactor) SetConfig(opts *bind.TransactOpts, participants []common.Address, weights []*big.Int, weightThresholdForHeartbeat *big.Int, weightThresholdForBadSignal *big.Int) (*types.Transaction, error) {
	return _AFNContract.contract.Transact(opts, "setConfig", participants, weights, weightThresholdForHeartbeat, weightThresholdForBadSignal)
}

func (_AFNContract *AFNContractSession) SetConfig(participants []common.Address, weights []*big.Int, weightThresholdForHeartbeat *big.Int, weightThresholdForBadSignal *big.Int) (*types.Transaction, error) {
	return _AFNContract.Contract.SetConfig(&_AFNContract.TransactOpts, participants, weights, weightThresholdForHeartbeat, weightThresholdForBadSignal)
}

func (_AFNContract *AFNContractTransactorSession) SetConfig(participants []common.Address, weights []*big.Int, weightThresholdForHeartbeat *big.Int, weightThresholdForBadSignal *big.Int) (*types.Transaction, error) {
	return _AFNContract.Contract.SetConfig(&_AFNContract.TransactOpts, participants, weights, weightThresholdForHeartbeat, weightThresholdForBadSignal)
}

func (_AFNContract *AFNContractTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AFNContract.contract.Transact(opts, "transferOwnership", to)
}

func (_AFNContract *AFNContractSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AFNContract.Contract.TransferOwnership(&_AFNContract.TransactOpts, to)
}

func (_AFNContract *AFNContractTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AFNContract.Contract.TransferOwnership(&_AFNContract.TransactOpts, to)
}

func (_AFNContract *AFNContractTransactor) VoteBad(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AFNContract.contract.Transact(opts, "voteBad")
}

func (_AFNContract *AFNContractSession) VoteBad() (*types.Transaction, error) {
	return _AFNContract.Contract.VoteBad(&_AFNContract.TransactOpts)
}

func (_AFNContract *AFNContractTransactorSession) VoteBad() (*types.Transaction, error) {
	return _AFNContract.Contract.VoteBad(&_AFNContract.TransactOpts)
}

func (_AFNContract *AFNContractTransactor) VoteGood(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error) {
	return _AFNContract.contract.Transact(opts, "voteGood", round)
}

func (_AFNContract *AFNContractSession) VoteGood(round *big.Int) (*types.Transaction, error) {
	return _AFNContract.Contract.VoteGood(&_AFNContract.TransactOpts, round)
}

func (_AFNContract *AFNContractTransactorSession) VoteGood(round *big.Int) (*types.Transaction, error) {
	return _AFNContract.Contract.VoteGood(&_AFNContract.TransactOpts, round)
}

type AFNContractAFNBadSignalIterator struct {
	Event *AFNContractAFNBadSignal

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AFNContractAFNBadSignalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AFNContractAFNBadSignal)
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
		it.Event = new(AFNContractAFNBadSignal)
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

func (it *AFNContractAFNBadSignalIterator) Error() error {
	return it.fail
}

func (it *AFNContractAFNBadSignalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AFNContractAFNBadSignal struct {
	Timestamp *big.Int
	Raw       types.Log
}

func (_AFNContract *AFNContractFilterer) FilterAFNBadSignal(opts *bind.FilterOpts) (*AFNContractAFNBadSignalIterator, error) {

	logs, sub, err := _AFNContract.contract.FilterLogs(opts, "AFNBadSignal")
	if err != nil {
		return nil, err
	}
	return &AFNContractAFNBadSignalIterator{contract: _AFNContract.contract, event: "AFNBadSignal", logs: logs, sub: sub}, nil
}

func (_AFNContract *AFNContractFilterer) WatchAFNBadSignal(opts *bind.WatchOpts, sink chan<- *AFNContractAFNBadSignal) (event.Subscription, error) {

	logs, sub, err := _AFNContract.contract.WatchLogs(opts, "AFNBadSignal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AFNContractAFNBadSignal)
				if err := _AFNContract.contract.UnpackLog(event, "AFNBadSignal", log); err != nil {
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

func (_AFNContract *AFNContractFilterer) ParseAFNBadSignal(log types.Log) (*AFNContractAFNBadSignal, error) {
	event := new(AFNContractAFNBadSignal)
	if err := _AFNContract.contract.UnpackLog(event, "AFNBadSignal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AFNContractAFNConfigSetIterator struct {
	Event *AFNContractAFNConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AFNContractAFNConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AFNContractAFNConfigSet)
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
		it.Event = new(AFNContractAFNConfigSet)
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

func (it *AFNContractAFNConfigSetIterator) Error() error {
	return it.fail
}

func (it *AFNContractAFNConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AFNContractAFNConfigSet struct {
	Parties    []common.Address
	Weights    []*big.Int
	GoodQuorum *big.Int
	BadQuorum  *big.Int
	Raw        types.Log
}

func (_AFNContract *AFNContractFilterer) FilterAFNConfigSet(opts *bind.FilterOpts) (*AFNContractAFNConfigSetIterator, error) {

	logs, sub, err := _AFNContract.contract.FilterLogs(opts, "AFNConfigSet")
	if err != nil {
		return nil, err
	}
	return &AFNContractAFNConfigSetIterator{contract: _AFNContract.contract, event: "AFNConfigSet", logs: logs, sub: sub}, nil
}

func (_AFNContract *AFNContractFilterer) WatchAFNConfigSet(opts *bind.WatchOpts, sink chan<- *AFNContractAFNConfigSet) (event.Subscription, error) {

	logs, sub, err := _AFNContract.contract.WatchLogs(opts, "AFNConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AFNContractAFNConfigSet)
				if err := _AFNContract.contract.UnpackLog(event, "AFNConfigSet", log); err != nil {
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

func (_AFNContract *AFNContractFilterer) ParseAFNConfigSet(log types.Log) (*AFNContractAFNConfigSet, error) {
	event := new(AFNContractAFNConfigSet)
	if err := _AFNContract.contract.UnpackLog(event, "AFNConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AFNContractAFNHeartbeatIterator struct {
	Event *AFNContractAFNHeartbeat

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AFNContractAFNHeartbeatIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AFNContractAFNHeartbeat)
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
		it.Event = new(AFNContractAFNHeartbeat)
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

func (it *AFNContractAFNHeartbeatIterator) Error() error {
	return it.fail
}

func (it *AFNContractAFNHeartbeatIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AFNContractAFNHeartbeat struct {
	Heartbeat AFNInterfaceHeartbeat
	Raw       types.Log
}

func (_AFNContract *AFNContractFilterer) FilterAFNHeartbeat(opts *bind.FilterOpts) (*AFNContractAFNHeartbeatIterator, error) {

	logs, sub, err := _AFNContract.contract.FilterLogs(opts, "AFNHeartbeat")
	if err != nil {
		return nil, err
	}
	return &AFNContractAFNHeartbeatIterator{contract: _AFNContract.contract, event: "AFNHeartbeat", logs: logs, sub: sub}, nil
}

func (_AFNContract *AFNContractFilterer) WatchAFNHeartbeat(opts *bind.WatchOpts, sink chan<- *AFNContractAFNHeartbeat) (event.Subscription, error) {

	logs, sub, err := _AFNContract.contract.WatchLogs(opts, "AFNHeartbeat")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AFNContractAFNHeartbeat)
				if err := _AFNContract.contract.UnpackLog(event, "AFNHeartbeat", log); err != nil {
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

func (_AFNContract *AFNContractFilterer) ParseAFNHeartbeat(log types.Log) (*AFNContractAFNHeartbeat, error) {
	event := new(AFNContractAFNHeartbeat)
	if err := _AFNContract.contract.UnpackLog(event, "AFNHeartbeat", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AFNContractBadVoteIterator struct {
	Event *AFNContractBadVote

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AFNContractBadVoteIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AFNContractBadVote)
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
		it.Event = new(AFNContractBadVote)
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

func (it *AFNContractBadVoteIterator) Error() error {
	return it.fail
}

func (it *AFNContractBadVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AFNContractBadVote struct {
	Voter common.Address
	Round *big.Int
	Raw   types.Log
}

func (_AFNContract *AFNContractFilterer) FilterBadVote(opts *bind.FilterOpts) (*AFNContractBadVoteIterator, error) {

	logs, sub, err := _AFNContract.contract.FilterLogs(opts, "BadVote")
	if err != nil {
		return nil, err
	}
	return &AFNContractBadVoteIterator{contract: _AFNContract.contract, event: "BadVote", logs: logs, sub: sub}, nil
}

func (_AFNContract *AFNContractFilterer) WatchBadVote(opts *bind.WatchOpts, sink chan<- *AFNContractBadVote) (event.Subscription, error) {

	logs, sub, err := _AFNContract.contract.WatchLogs(opts, "BadVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AFNContractBadVote)
				if err := _AFNContract.contract.UnpackLog(event, "BadVote", log); err != nil {
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

func (_AFNContract *AFNContractFilterer) ParseBadVote(log types.Log) (*AFNContractBadVote, error) {
	event := new(AFNContractBadVote)
	if err := _AFNContract.contract.UnpackLog(event, "BadVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AFNContractGoodVoteIterator struct {
	Event *AFNContractGoodVote

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AFNContractGoodVoteIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AFNContractGoodVote)
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
		it.Event = new(AFNContractGoodVote)
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

func (it *AFNContractGoodVoteIterator) Error() error {
	return it.fail
}

func (it *AFNContractGoodVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AFNContractGoodVote struct {
	Voter common.Address
	Round *big.Int
	Raw   types.Log
}

func (_AFNContract *AFNContractFilterer) FilterGoodVote(opts *bind.FilterOpts) (*AFNContractGoodVoteIterator, error) {

	logs, sub, err := _AFNContract.contract.FilterLogs(opts, "GoodVote")
	if err != nil {
		return nil, err
	}
	return &AFNContractGoodVoteIterator{contract: _AFNContract.contract, event: "GoodVote", logs: logs, sub: sub}, nil
}

func (_AFNContract *AFNContractFilterer) WatchGoodVote(opts *bind.WatchOpts, sink chan<- *AFNContractGoodVote) (event.Subscription, error) {

	logs, sub, err := _AFNContract.contract.WatchLogs(opts, "GoodVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AFNContractGoodVote)
				if err := _AFNContract.contract.UnpackLog(event, "GoodVote", log); err != nil {
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

func (_AFNContract *AFNContractFilterer) ParseGoodVote(log types.Log) (*AFNContractGoodVote, error) {
	event := new(AFNContractGoodVote)
	if err := _AFNContract.contract.UnpackLog(event, "GoodVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AFNContractOwnershipTransferRequestedIterator struct {
	Event *AFNContractOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AFNContractOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AFNContractOwnershipTransferRequested)
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
		it.Event = new(AFNContractOwnershipTransferRequested)
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

func (it *AFNContractOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *AFNContractOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AFNContractOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_AFNContract *AFNContractFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AFNContractOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AFNContract.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AFNContractOwnershipTransferRequestedIterator{contract: _AFNContract.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_AFNContract *AFNContractFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AFNContractOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AFNContract.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AFNContractOwnershipTransferRequested)
				if err := _AFNContract.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_AFNContract *AFNContractFilterer) ParseOwnershipTransferRequested(log types.Log) (*AFNContractOwnershipTransferRequested, error) {
	event := new(AFNContractOwnershipTransferRequested)
	if err := _AFNContract.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AFNContractOwnershipTransferredIterator struct {
	Event *AFNContractOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AFNContractOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AFNContractOwnershipTransferred)
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
		it.Event = new(AFNContractOwnershipTransferred)
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

func (it *AFNContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *AFNContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AFNContractOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_AFNContract *AFNContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AFNContractOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AFNContract.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AFNContractOwnershipTransferredIterator{contract: _AFNContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_AFNContract *AFNContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AFNContractOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AFNContract.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AFNContractOwnershipTransferred)
				if err := _AFNContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_AFNContract *AFNContractFilterer) ParseOwnershipTransferred(log types.Log) (*AFNContractOwnershipTransferred, error) {
	event := new(AFNContractOwnershipTransferred)
	if err := _AFNContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AFNContractRecoveredFromBadSignalIterator struct {
	Event *AFNContractRecoveredFromBadSignal

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AFNContractRecoveredFromBadSignalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AFNContractRecoveredFromBadSignal)
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
		it.Event = new(AFNContractRecoveredFromBadSignal)
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

func (it *AFNContractRecoveredFromBadSignalIterator) Error() error {
	return it.fail
}

func (it *AFNContractRecoveredFromBadSignalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AFNContractRecoveredFromBadSignal struct {
	Raw types.Log
}

func (_AFNContract *AFNContractFilterer) FilterRecoveredFromBadSignal(opts *bind.FilterOpts) (*AFNContractRecoveredFromBadSignalIterator, error) {

	logs, sub, err := _AFNContract.contract.FilterLogs(opts, "RecoveredFromBadSignal")
	if err != nil {
		return nil, err
	}
	return &AFNContractRecoveredFromBadSignalIterator{contract: _AFNContract.contract, event: "RecoveredFromBadSignal", logs: logs, sub: sub}, nil
}

func (_AFNContract *AFNContractFilterer) WatchRecoveredFromBadSignal(opts *bind.WatchOpts, sink chan<- *AFNContractRecoveredFromBadSignal) (event.Subscription, error) {

	logs, sub, err := _AFNContract.contract.WatchLogs(opts, "RecoveredFromBadSignal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AFNContractRecoveredFromBadSignal)
				if err := _AFNContract.contract.UnpackLog(event, "RecoveredFromBadSignal", log); err != nil {
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

func (_AFNContract *AFNContractFilterer) ParseRecoveredFromBadSignal(log types.Log) (*AFNContractRecoveredFromBadSignal, error) {
	event := new(AFNContractRecoveredFromBadSignal)
	if err := _AFNContract.contract.UnpackLog(event, "RecoveredFromBadSignal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetBadVotersAndVotes struct {
	Voters []common.Address
	Votes  *big.Int
}
type GetWeightThresholds struct {
	Good *big.Int
	Bad  *big.Int
}

func (_AFNContract *AFNContract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _AFNContract.abi.Events["AFNBadSignal"].ID:
		return _AFNContract.ParseAFNBadSignal(log)
	case _AFNContract.abi.Events["AFNConfigSet"].ID:
		return _AFNContract.ParseAFNConfigSet(log)
	case _AFNContract.abi.Events["AFNHeartbeat"].ID:
		return _AFNContract.ParseAFNHeartbeat(log)
	case _AFNContract.abi.Events["BadVote"].ID:
		return _AFNContract.ParseBadVote(log)
	case _AFNContract.abi.Events["GoodVote"].ID:
		return _AFNContract.ParseGoodVote(log)
	case _AFNContract.abi.Events["OwnershipTransferRequested"].ID:
		return _AFNContract.ParseOwnershipTransferRequested(log)
	case _AFNContract.abi.Events["OwnershipTransferred"].ID:
		return _AFNContract.ParseOwnershipTransferred(log)
	case _AFNContract.abi.Events["RecoveredFromBadSignal"].ID:
		return _AFNContract.ParseRecoveredFromBadSignal(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (AFNContractAFNBadSignal) Topic() common.Hash {
	return common.HexToHash("0x73907f5e30313a1ab6e1815608b22b40911f1a7decec69d5df18a2298002bacb")
}

func (AFNContractAFNConfigSet) Topic() common.Hash {
	return common.HexToHash("0x69af5b8b5b348d6b619cb6b338b5cfd865aa9e8cedd36a4a69257a9a07ebedaa")
}

func (AFNContractAFNHeartbeat) Topic() common.Hash {
	return common.HexToHash("0x90b45dcfd48782731999668957597f8b47e29aaa1d53ef2ad07612429777bed3")
}

func (AFNContractBadVote) Topic() common.Hash {
	return common.HexToHash("0x0b21c4350e3db4e6e412b398113b3769fa4fcf4582c88579705b3d42002a41fd")
}

func (AFNContractGoodVote) Topic() common.Hash {
	return common.HexToHash("0x5489e43df72470c733e49d6f7bc612d52f64600fb2801593290ec32fcf144791")
}

func (AFNContractOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (AFNContractOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (AFNContractRecoveredFromBadSignal) Topic() common.Hash {
	return common.HexToHash("0x3e48434bea67b1e259c2380d289dcb6372257ab2c37bc86f0e1acf83a7b07ac0")
}

func (_AFNContract *AFNContract) Address() common.Address {
	return _AFNContract.address
}

type AFNContractInterface interface {
	GetBadVotersAndVotes(opts *bind.CallOpts) (GetBadVotersAndVotes,

		error)

	GetCommitteeVersion(opts *bind.CallOpts) (*big.Int, error)

	GetGoodVotes(opts *bind.CallOpts) (*big.Int, error)

	GetLastGoodVoteByParticipant(opts *bind.CallOpts, participant common.Address) (*big.Int, error)

	GetLastHeartbeat(opts *bind.CallOpts) (AFNInterfaceHeartbeat, error)

	GetParticipants(opts *bind.CallOpts) ([]common.Address, error)

	GetRound(opts *bind.CallOpts) (*big.Int, error)

	GetWeightByParticipant(opts *bind.CallOpts, participant common.Address) (*big.Int, error)

	GetWeightThresholds(opts *bind.CallOpts) (GetWeightThresholds,

		error)

	HasBadSignal(opts *bind.CallOpts) (bool, error)

	HasVotedBad(opts *bind.CallOpts, participant common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Recover(opts *bind.TransactOpts) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, participants []common.Address, weights []*big.Int, weightThresholdForHeartbeat *big.Int, weightThresholdForBadSignal *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	VoteBad(opts *bind.TransactOpts) (*types.Transaction, error)

	VoteGood(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error)

	FilterAFNBadSignal(opts *bind.FilterOpts) (*AFNContractAFNBadSignalIterator, error)

	WatchAFNBadSignal(opts *bind.WatchOpts, sink chan<- *AFNContractAFNBadSignal) (event.Subscription, error)

	ParseAFNBadSignal(log types.Log) (*AFNContractAFNBadSignal, error)

	FilterAFNConfigSet(opts *bind.FilterOpts) (*AFNContractAFNConfigSetIterator, error)

	WatchAFNConfigSet(opts *bind.WatchOpts, sink chan<- *AFNContractAFNConfigSet) (event.Subscription, error)

	ParseAFNConfigSet(log types.Log) (*AFNContractAFNConfigSet, error)

	FilterAFNHeartbeat(opts *bind.FilterOpts) (*AFNContractAFNHeartbeatIterator, error)

	WatchAFNHeartbeat(opts *bind.WatchOpts, sink chan<- *AFNContractAFNHeartbeat) (event.Subscription, error)

	ParseAFNHeartbeat(log types.Log) (*AFNContractAFNHeartbeat, error)

	FilterBadVote(opts *bind.FilterOpts) (*AFNContractBadVoteIterator, error)

	WatchBadVote(opts *bind.WatchOpts, sink chan<- *AFNContractBadVote) (event.Subscription, error)

	ParseBadVote(log types.Log) (*AFNContractBadVote, error)

	FilterGoodVote(opts *bind.FilterOpts) (*AFNContractGoodVoteIterator, error)

	WatchGoodVote(opts *bind.WatchOpts, sink chan<- *AFNContractGoodVote) (event.Subscription, error)

	ParseGoodVote(log types.Log) (*AFNContractGoodVote, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AFNContractOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AFNContractOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*AFNContractOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AFNContractOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AFNContractOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*AFNContractOwnershipTransferred, error)

	FilterRecoveredFromBadSignal(opts *bind.FilterOpts) (*AFNContractRecoveredFromBadSignalIterator, error)

	WatchRecoveredFromBadSignal(opts *bind.WatchOpts, sink chan<- *AFNContractRecoveredFromBadSignal) (event.Subscription, error)

	ParseRecoveredFromBadSignal(log types.Log) (*AFNContractRecoveredFromBadSignal, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
