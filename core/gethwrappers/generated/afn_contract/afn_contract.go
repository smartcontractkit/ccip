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
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated"
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

var AFNContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"weightThresholdForBlessing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightThresholdForBadSignal\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyVoted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"InvalidVoter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustRecoverFromBadSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RecoveryNotNecessary\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AFNBadSignal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"parties\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goodQuorum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"badQuorum\",\"type\":\"uint256\"}],\"name\":\"AFNConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RecoveredFromBadSignal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"name\":\"RootBlessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"VoteBad\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"VoteToBless\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"badSignalReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBadVotersAndVotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfigVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParticipants\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getVotesToBlessRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"getWeightByParticipant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWeightThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blessing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"badSignal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"hasVotedBad\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"hasVotedToBlessRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootWithOrigin\",\"type\":\"bytes32\"}],\"name\":\"isBlessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverFromBadSignal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"weightThresholdForBlessing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightThresholdForBadSignal\",\"type\":\"uint256\"}],\"name\":\"setAFNConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteBad\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"rootsWithOrigin\",\"type\":\"bytes32[]\"}],\"name\":\"voteToBlessRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620020383803806200203883398101604081905262000034916200082d565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000e1565b505050620000d78484848460016200018c60201b60201c565b5050505062000a10565b336001600160a01b038216036200013b5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b835185511415806200019d57508451155b80620001a7575082155b80620001b1575081155b80620001bb575080155b15620001da576040516306b7c75960e31b815260040160405180910390fd5b600060038054806020026020016040519081016040528092919081815260200182805480156200023457602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831162000215575b5050505050905060005b8151811015620002a157600260008383815181106200026157620002616200091b565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206000905580620002999062000947565b90506200023e565b50600484905560058390556006829055620002bb62000491565b620002c6816200056d565b8551600090620002de9060039060208a0190620006a7565b5060005b87518110156200041b5760006001600160a01b03168882815181106200030c576200030c6200091b565b60200260200101516001600160a01b0316036200033c576040516306b7c75960e31b815260040160405180910390fd5b8681815181106200035157620003516200091b565b60200260200101516000036200037a5760405163585b926360e01b815260040160405180910390fd5b8681815181106200038f576200038f6200091b565b6020026020010151600260008a8481518110620003b057620003b06200091b565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002081905550868181518110620003f157620003f16200091b565b60200260200101518262000406919062000963565b9150620004138162000947565b9050620002e2565b50848110806200042a57508381105b1562000449576040516306b7c75960e31b815260040160405180910390fd5b7f69af5b8b5b348d6b619cb6b338b5cfd865aa9e8cedd36a4a69257a9a07ebedaa878787876040516200048094939291906200097e565b60405180910390a150505050505050565b60006007805480602002602001604051908101604052809291908181526020018280548015620004eb57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311620004cc575b5050505050905060005b81518110156200055657600860008383815181106200051857620005186200091b565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191690556200054e8162000947565b9050620004f5565b50620005656007600062000711565b506000600955565b6000600c805480602002602001604051908101604052809291908181526020018280548015620005bd57602002820191906000526020600020905b815481526020019060010190808311620005a8575b5050505050905060005b815181101562000694576000828281518110620005e857620005e86200091b565b60200260200101519050600b60008281526020019081526020016000206000905560005b84518110156200067e57600d600083815260200190815260200160002060008683815181106200064057620006406200091b565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff19169055620006768162000947565b90506200060c565b5050806200068c9062000947565b9050620005c7565b50620006a3600c600062000711565b5050565b828054828255906000526020600020908101928215620006ff579160200282015b82811115620006ff57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620006c8565b506200070d92915062000734565b5090565b508054600082559060005260206000209081019062000731919062000734565b50565b5b808211156200070d576000815560010162000735565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156200078c576200078c6200074b565b604052919050565b60006001600160401b03821115620007b057620007b06200074b565b5060051b60200190565b600082601f830112620007cc57600080fd5b81516020620007e5620007df8362000794565b62000761565b82815260059290921b840181019181810190868411156200080557600080fd5b8286015b8481101562000822578051835291830191830162000809565b509695505050505050565b600080600080608085870312156200084457600080fd5b84516001600160401b03808211156200085c57600080fd5b818701915087601f8301126200087157600080fd5b8151602062000884620007df8362000794565b82815260059290921b8401810191818101908b841115620008a457600080fd5b948201945b83861015620008db5785516001600160a01b0381168114620008cb5760008081fd5b82529482019490820190620008a9565b918a0151919850909350505080821115620008f557600080fd5b506200090487828801620007ba565b604087015160609097015195989097509350505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016200095c576200095c62000931565b5060010190565b6000821982111562000979576200097962000931565b500190565b6080808252855190820181905260009060209060a0840190828901845b82811015620009c25781516001600160a01b0316845292840192908401906001016200099b565b5050508381038285015286518082528783019183019060005b81811015620009f957835183529284019291840191600101620009db565b505060408501969096525050506060015292915050565b6116188062000a206000396000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c806379adb16e116100b25780639dc6edc711610081578063f2fde38b11610066578063f2fde38b146102f3578063f438c9c014610306578063ff888fb11461032157600080fd5b80639dc6edc7146102d5578063c3453fa5146102dd57600080fd5b806379adb16e1461025b57806379ba5097146102925780638da5cb5b1461029a5780638e1d4e61146102b557600080fd5b80633cd4f66911610109578063508ede92116100ee578063508ede92146101fb57806351856565146102335780635aa68ac01461024657600080fd5b80633cd4f669146101e657806346f8e6d7146101ee57600080fd5b8063181f5a771461013b5780632cb145d41461018d5780632ea9537114610197578063365f15ec146101d3575b600080fd5b6101776040518060400160405280600981526020017f41464e20312e302e30000000000000000000000000000000000000000000000081525081565b604051610184919061113d565b60405180910390f35b610195610344565b005b6101c36101a53660046111cc565b6001600160a01b031660009081526008602052604090205460ff1690565b6040519015158152602001610184565b6101956101e13660046111ee565b610543565b6101956107a9565b60055460095410156101c3565b6101c3610209366004611263565b6000908152600d602090815260408083206001600160a01b03949094168352929052205460ff1690565b61019561024136600461139a565b610822565b61024e61084b565b60405161018491906114ae565b6102846102693660046111cc565b6001600160a01b031660009081526002602052604090205490565b604051908152602001610184565b6101956108ad565b6000546040516001600160a01b039091168152602001610184565b6102846102c33660046114c1565b6000908152600b602052604090205490565b600654610284565b6102e5610990565b6040516101849291906114da565b6101956103013660046111cc565b6109fd565b60045460055460408051928352602083019190915201610184565b6101c361032f3660046114c1565b6000908152600a602052604090205460ff1690565b60055460095410610381576040517fc28cc95000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b33600081815260026020526040812054908190036103db576040517f669f262e0000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b6001600160a01b03821660009081526008602052604090205460ff161561042e576040517f7c9a1cf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382166000818152600860205260408120805460ff1916600190811790915560078054918201815582527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880180547fffffffffffffffffffffffff000000000000000000000000000000000000000016909217909155600980548392906104bd90849061152b565b90915550506040518181526001600160a01b038316907fa5889da6c2d25ef72eaae82bb0b8acf51eeebdd6bd12f1a24360de7d9b9cfa289060200160405180910390a26005546009541061053f576040514281527f73907f5e30313a1ab6e1815608b22b40911f1a7decec69d5df18a2298002bacb9060200160405180910390a15b5050565b60055460095410610580576040517fc28cc95000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b33600081815260026020526040812054908190036105d5576040517f669f262e0000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016103d2565b60005b838110156107a25760008585838181106105f4576105f4611543565b905060200201359050610616816000908152600a602052604090205460ff1690565b156106215750610792565b6000818152600d602090815260408083206001600160a01b038816845290915290205460ff16156106525750610792565b6000818152600d602090815260408083206001600160a01b03881684528252808320805460ff19166001179055838352600b909152812054908190036106c857600c80546001810182556000919091527fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c7018290555b60006106d4858361152b565b6000848152600b6020526040908190208290555190915083906001600160a01b038816907f262f79a5a063a0af3e27989b0b0f0ae1e2c19257d27efe01a7f0cab7b3b470a4906107279089815260200190565b60405180910390a3600454811061078e576000838152600a602052604090819020805460ff191660011790555183907f719fab74b843fdceffa591cc0a3445a9dddc9e1e304471baed67e8408a1405c7906107859084815260200190565b60405180910390a25b5050505b61079b81611572565b90506105d8565b5050505050565b6107b1610a11565b60055460095410156107ef576040517fe147761200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107f7610a87565b6040517f3e48434bea67b1e259c2380d289dcb6372257ab2c37bc86f0e1acf83a7b07ac090600090a1565b61082a610a11565b610845848484846006546001610840919061152b565b610b58565b50505050565b606060038054806020026020016040519081016040528092919081815260200182805480156108a357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610885575b5050505050905090565b6001546001600160a01b03163314610921576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016103d2565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b606060006007600954818054806020026020016040519081016040528092919081815260200182805480156109ee57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116109d0575b50505050509150915091509091565b610a05610a11565b610a0e81610e92565b50565b6000546001600160a01b03163314610a85576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016103d2565b565b60006007805480602002602001604051908101604052809291908181526020018280548015610adf57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610ac1575b5050505050905060005b8151811015610b435760086000838381518110610b0857610b08611543565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff19169055610b3c81611572565b9050610ae9565b50610b506007600061108d565b506000600955565b83518551141580610b6857508451155b80610b71575082155b80610b7a575081155b80610b83575080155b15610bba576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006003805480602002602001604051908101604052809291908181526020018280548015610c1257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610bf4575b5050505050905060005b8151811015610c785760026000838381518110610c3b57610c3b611543565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206000905580610c7190611572565b9050610c1c565b50600484905560058390556006829055610c90610a87565b610c9981610f6d565b8551600090610caf9060039060208a01906110ab565b5060005b8751811015610e075760006001600160a01b0316888281518110610cd957610cd9611543565b60200260200101516001600160a01b031603610d21576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b868181518110610d3357610d33611543565b6020026020010151600003610d74576040517f585b926300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b868181518110610d8657610d86611543565b6020026020010151600260008a8481518110610da457610da4611543565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002081905550868181518110610de257610de2611543565b602002602001015182610df5919061152b565b9150610e0081611572565b9050610cb3565b5084811080610e1557508381105b15610e4c576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f69af5b8b5b348d6b619cb6b338b5cfd865aa9e8cedd36a4a69257a9a07ebedaa87878787604051610e8194939291906115aa565b60405180910390a150505050505050565b336001600160a01b03821603610f04576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016103d2565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000600c805480602002602001604051908101604052809291908181526020018280548015610fbb57602002820191906000526020600020905b815481526020019060010190808311610fa7575b5050505050905060005b8151811015611084576000828281518110610fe257610fe2611543565b60200260200101519050600b60008281526020019081526020016000206000905560005b845181101561107157600d6000838152602001908152602001600020600086838151811061103657611036611543565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916905561106a81611572565b9050611006565b50508061107d90611572565b9050610fc5565b5061053f600c60005b5080546000825590600052602060002090810190610a0e9190611128565b828054828255906000526020600020908101928215611118579160200282015b8281111561111857825182547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039091161782556020909201916001909101906110cb565b50611124929150611128565b5090565b5b808211156111245760008155600101611129565b600060208083528351808285015260005b8181101561116a5785810183015185820160400152820161114e565b8181111561117c576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b80356001600160a01b03811681146111c757600080fd5b919050565b6000602082840312156111de57600080fd5b6111e7826111b0565b9392505050565b6000806020838503121561120157600080fd5b823567ffffffffffffffff8082111561121957600080fd5b818501915085601f83011261122d57600080fd5b81358181111561123c57600080fd5b8660208260051b850101111561125157600080fd5b60209290920196919550909350505050565b6000806040838503121561127657600080fd5b61127f836111b0565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156113035761130361128d565b604052919050565b600067ffffffffffffffff8211156113255761132561128d565b5060051b60200190565b600082601f83011261134057600080fd5b813560206113556113508361130b565b6112bc565b82815260059290921b8401810191818101908684111561137457600080fd5b8286015b8481101561138f5780358352918301918301611378565b509695505050505050565b600080600080608085870312156113b057600080fd5b843567ffffffffffffffff808211156113c857600080fd5b818701915087601f8301126113dc57600080fd5b813560206113ec6113508361130b565b82815260059290921b8401810191818101908b84111561140b57600080fd5b948201945b8386101561143057611421866111b0565b82529482019490820190611410565b9850508801359250508082111561144657600080fd5b506114538782880161132f565b949794965050505060408301359260600135919050565b600081518084526020808501945080840160005b838110156114a35781516001600160a01b03168752958201959082019060010161147e565b509495945050505050565b6020815260006111e7602083018461146a565b6000602082840312156114d357600080fd5b5035919050565b6040815260006114ed604083018561146a565b90508260208301529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561153e5761153e6114fc565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036115a3576115a36114fc565b5060010190565b6080815260006115bd608083018761146a565b82810360208481019190915286518083528782019282019060005b818110156115f4578451835293830193918301916001016115d8565b50506040850196909652505050606001529291505056fea164736f6c634300080f000a",
}

var AFNContractABI = AFNContractMetaData.ABI

var AFNContractBin = AFNContractMetaData.Bin

func DeployAFNContract(auth *bind.TransactOpts, backend bind.ContractBackend, participants []common.Address, weights []*big.Int, weightThresholdForBlessing *big.Int, weightThresholdForBadSignal *big.Int) (common.Address, *types.Transaction, *AFNContract, error) {
	parsed, err := AFNContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AFNContractBin), backend, participants, weights, weightThresholdForBlessing, weightThresholdForBadSignal)
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

func (_AFNContract *AFNContractCaller) BadSignalReceived(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "badSignalReceived")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AFNContract *AFNContractSession) BadSignalReceived() (bool, error) {
	return _AFNContract.Contract.BadSignalReceived(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCallerSession) BadSignalReceived() (bool, error) {
	return _AFNContract.Contract.BadSignalReceived(&_AFNContract.CallOpts)
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

func (_AFNContract *AFNContractCaller) GetConfigVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getConfigVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AFNContract *AFNContractSession) GetConfigVersion() (*big.Int, error) {
	return _AFNContract.Contract.GetConfigVersion(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCallerSession) GetConfigVersion() (*big.Int, error) {
	return _AFNContract.Contract.GetConfigVersion(&_AFNContract.CallOpts)
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

func (_AFNContract *AFNContractCaller) GetVotesToBlessRoot(opts *bind.CallOpts, root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getVotesToBlessRoot", root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AFNContract *AFNContractSession) GetVotesToBlessRoot(root [32]byte) (*big.Int, error) {
	return _AFNContract.Contract.GetVotesToBlessRoot(&_AFNContract.CallOpts, root)
}

func (_AFNContract *AFNContractCallerSession) GetVotesToBlessRoot(root [32]byte) (*big.Int, error) {
	return _AFNContract.Contract.GetVotesToBlessRoot(&_AFNContract.CallOpts, root)
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

	outstruct.Blessing = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BadSignal = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

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

func (_AFNContract *AFNContractCaller) HasVotedToBlessRoot(opts *bind.CallOpts, participant common.Address, root [32]byte) (bool, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "hasVotedToBlessRoot", participant, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AFNContract *AFNContractSession) HasVotedToBlessRoot(participant common.Address, root [32]byte) (bool, error) {
	return _AFNContract.Contract.HasVotedToBlessRoot(&_AFNContract.CallOpts, participant, root)
}

func (_AFNContract *AFNContractCallerSession) HasVotedToBlessRoot(participant common.Address, root [32]byte) (bool, error) {
	return _AFNContract.Contract.HasVotedToBlessRoot(&_AFNContract.CallOpts, participant, root)
}

func (_AFNContract *AFNContractCaller) IsBlessed(opts *bind.CallOpts, rootWithOrigin [32]byte) (bool, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "isBlessed", rootWithOrigin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AFNContract *AFNContractSession) IsBlessed(rootWithOrigin [32]byte) (bool, error) {
	return _AFNContract.Contract.IsBlessed(&_AFNContract.CallOpts, rootWithOrigin)
}

func (_AFNContract *AFNContractCallerSession) IsBlessed(rootWithOrigin [32]byte) (bool, error) {
	return _AFNContract.Contract.IsBlessed(&_AFNContract.CallOpts, rootWithOrigin)
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

func (_AFNContract *AFNContractTransactor) RecoverFromBadSignal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AFNContract.contract.Transact(opts, "recoverFromBadSignal")
}

func (_AFNContract *AFNContractSession) RecoverFromBadSignal() (*types.Transaction, error) {
	return _AFNContract.Contract.RecoverFromBadSignal(&_AFNContract.TransactOpts)
}

func (_AFNContract *AFNContractTransactorSession) RecoverFromBadSignal() (*types.Transaction, error) {
	return _AFNContract.Contract.RecoverFromBadSignal(&_AFNContract.TransactOpts)
}

func (_AFNContract *AFNContractTransactor) SetAFNConfig(opts *bind.TransactOpts, participants []common.Address, weights []*big.Int, weightThresholdForBlessing *big.Int, weightThresholdForBadSignal *big.Int) (*types.Transaction, error) {
	return _AFNContract.contract.Transact(opts, "setAFNConfig", participants, weights, weightThresholdForBlessing, weightThresholdForBadSignal)
}

func (_AFNContract *AFNContractSession) SetAFNConfig(participants []common.Address, weights []*big.Int, weightThresholdForBlessing *big.Int, weightThresholdForBadSignal *big.Int) (*types.Transaction, error) {
	return _AFNContract.Contract.SetAFNConfig(&_AFNContract.TransactOpts, participants, weights, weightThresholdForBlessing, weightThresholdForBadSignal)
}

func (_AFNContract *AFNContractTransactorSession) SetAFNConfig(participants []common.Address, weights []*big.Int, weightThresholdForBlessing *big.Int, weightThresholdForBadSignal *big.Int) (*types.Transaction, error) {
	return _AFNContract.Contract.SetAFNConfig(&_AFNContract.TransactOpts, participants, weights, weightThresholdForBlessing, weightThresholdForBadSignal)
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

func (_AFNContract *AFNContractTransactor) VoteToBlessRoots(opts *bind.TransactOpts, rootsWithOrigin [][32]byte) (*types.Transaction, error) {
	return _AFNContract.contract.Transact(opts, "voteToBlessRoots", rootsWithOrigin)
}

func (_AFNContract *AFNContractSession) VoteToBlessRoots(rootsWithOrigin [][32]byte) (*types.Transaction, error) {
	return _AFNContract.Contract.VoteToBlessRoots(&_AFNContract.TransactOpts, rootsWithOrigin)
}

func (_AFNContract *AFNContractTransactorSession) VoteToBlessRoots(rootsWithOrigin [][32]byte) (*types.Transaction, error) {
	return _AFNContract.Contract.VoteToBlessRoots(&_AFNContract.TransactOpts, rootsWithOrigin)
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

type AFNContractRootBlessedIterator struct {
	Event *AFNContractRootBlessed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AFNContractRootBlessedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AFNContractRootBlessed)
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
		it.Event = new(AFNContractRootBlessed)
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

func (it *AFNContractRootBlessedIterator) Error() error {
	return it.fail
}

func (it *AFNContractRootBlessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AFNContractRootBlessed struct {
	Root  [32]byte
	Votes *big.Int
	Raw   types.Log
}

func (_AFNContract *AFNContractFilterer) FilterRootBlessed(opts *bind.FilterOpts, root [][32]byte) (*AFNContractRootBlessedIterator, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _AFNContract.contract.FilterLogs(opts, "RootBlessed", rootRule)
	if err != nil {
		return nil, err
	}
	return &AFNContractRootBlessedIterator{contract: _AFNContract.contract, event: "RootBlessed", logs: logs, sub: sub}, nil
}

func (_AFNContract *AFNContractFilterer) WatchRootBlessed(opts *bind.WatchOpts, sink chan<- *AFNContractRootBlessed, root [][32]byte) (event.Subscription, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _AFNContract.contract.WatchLogs(opts, "RootBlessed", rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AFNContractRootBlessed)
				if err := _AFNContract.contract.UnpackLog(event, "RootBlessed", log); err != nil {
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

func (_AFNContract *AFNContractFilterer) ParseRootBlessed(log types.Log) (*AFNContractRootBlessed, error) {
	event := new(AFNContractRootBlessed)
	if err := _AFNContract.contract.UnpackLog(event, "RootBlessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AFNContractVoteBadIterator struct {
	Event *AFNContractVoteBad

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AFNContractVoteBadIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AFNContractVoteBad)
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
		it.Event = new(AFNContractVoteBad)
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

func (it *AFNContractVoteBadIterator) Error() error {
	return it.fail
}

func (it *AFNContractVoteBadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AFNContractVoteBad struct {
	Voter  common.Address
	Weight *big.Int
	Raw    types.Log
}

func (_AFNContract *AFNContractFilterer) FilterVoteBad(opts *bind.FilterOpts, voter []common.Address) (*AFNContractVoteBadIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _AFNContract.contract.FilterLogs(opts, "VoteBad", voterRule)
	if err != nil {
		return nil, err
	}
	return &AFNContractVoteBadIterator{contract: _AFNContract.contract, event: "VoteBad", logs: logs, sub: sub}, nil
}

func (_AFNContract *AFNContractFilterer) WatchVoteBad(opts *bind.WatchOpts, sink chan<- *AFNContractVoteBad, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _AFNContract.contract.WatchLogs(opts, "VoteBad", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AFNContractVoteBad)
				if err := _AFNContract.contract.UnpackLog(event, "VoteBad", log); err != nil {
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

func (_AFNContract *AFNContractFilterer) ParseVoteBad(log types.Log) (*AFNContractVoteBad, error) {
	event := new(AFNContractVoteBad)
	if err := _AFNContract.contract.UnpackLog(event, "VoteBad", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AFNContractVoteToBlessIterator struct {
	Event *AFNContractVoteToBless

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AFNContractVoteToBlessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AFNContractVoteToBless)
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
		it.Event = new(AFNContractVoteToBless)
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

func (it *AFNContractVoteToBlessIterator) Error() error {
	return it.fail
}

func (it *AFNContractVoteToBlessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AFNContractVoteToBless struct {
	Voter  common.Address
	Root   [32]byte
	Weight *big.Int
	Raw    types.Log
}

func (_AFNContract *AFNContractFilterer) FilterVoteToBless(opts *bind.FilterOpts, voter []common.Address, root [][32]byte) (*AFNContractVoteToBlessIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _AFNContract.contract.FilterLogs(opts, "VoteToBless", voterRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &AFNContractVoteToBlessIterator{contract: _AFNContract.contract, event: "VoteToBless", logs: logs, sub: sub}, nil
}

func (_AFNContract *AFNContractFilterer) WatchVoteToBless(opts *bind.WatchOpts, sink chan<- *AFNContractVoteToBless, voter []common.Address, root [][32]byte) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _AFNContract.contract.WatchLogs(opts, "VoteToBless", voterRule, rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AFNContractVoteToBless)
				if err := _AFNContract.contract.UnpackLog(event, "VoteToBless", log); err != nil {
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

func (_AFNContract *AFNContractFilterer) ParseVoteToBless(log types.Log) (*AFNContractVoteToBless, error) {
	event := new(AFNContractVoteToBless)
	if err := _AFNContract.contract.UnpackLog(event, "VoteToBless", log); err != nil {
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
	Blessing  *big.Int
	BadSignal *big.Int
}

func (_AFNContract *AFNContract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _AFNContract.abi.Events["AFNBadSignal"].ID:
		return _AFNContract.ParseAFNBadSignal(log)
	case _AFNContract.abi.Events["AFNConfigSet"].ID:
		return _AFNContract.ParseAFNConfigSet(log)
	case _AFNContract.abi.Events["OwnershipTransferRequested"].ID:
		return _AFNContract.ParseOwnershipTransferRequested(log)
	case _AFNContract.abi.Events["OwnershipTransferred"].ID:
		return _AFNContract.ParseOwnershipTransferred(log)
	case _AFNContract.abi.Events["RecoveredFromBadSignal"].ID:
		return _AFNContract.ParseRecoveredFromBadSignal(log)
	case _AFNContract.abi.Events["RootBlessed"].ID:
		return _AFNContract.ParseRootBlessed(log)
	case _AFNContract.abi.Events["VoteBad"].ID:
		return _AFNContract.ParseVoteBad(log)
	case _AFNContract.abi.Events["VoteToBless"].ID:
		return _AFNContract.ParseVoteToBless(log)

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

func (AFNContractOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (AFNContractOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (AFNContractRecoveredFromBadSignal) Topic() common.Hash {
	return common.HexToHash("0x3e48434bea67b1e259c2380d289dcb6372257ab2c37bc86f0e1acf83a7b07ac0")
}

func (AFNContractRootBlessed) Topic() common.Hash {
	return common.HexToHash("0x719fab74b843fdceffa591cc0a3445a9dddc9e1e304471baed67e8408a1405c7")
}

func (AFNContractVoteBad) Topic() common.Hash {
	return common.HexToHash("0xa5889da6c2d25ef72eaae82bb0b8acf51eeebdd6bd12f1a24360de7d9b9cfa28")
}

func (AFNContractVoteToBless) Topic() common.Hash {
	return common.HexToHash("0x262f79a5a063a0af3e27989b0b0f0ae1e2c19257d27efe01a7f0cab7b3b470a4")
}

func (_AFNContract *AFNContract) Address() common.Address {
	return _AFNContract.address
}

type AFNContractInterface interface {
	BadSignalReceived(opts *bind.CallOpts) (bool, error)

	GetBadVotersAndVotes(opts *bind.CallOpts) (GetBadVotersAndVotes,

		error)

	GetConfigVersion(opts *bind.CallOpts) (*big.Int, error)

	GetParticipants(opts *bind.CallOpts) ([]common.Address, error)

	GetVotesToBlessRoot(opts *bind.CallOpts, root [32]byte) (*big.Int, error)

	GetWeightByParticipant(opts *bind.CallOpts, participant common.Address) (*big.Int, error)

	GetWeightThresholds(opts *bind.CallOpts) (GetWeightThresholds,

		error)

	HasVotedBad(opts *bind.CallOpts, participant common.Address) (bool, error)

	HasVotedToBlessRoot(opts *bind.CallOpts, participant common.Address, root [32]byte) (bool, error)

	IsBlessed(opts *bind.CallOpts, rootWithOrigin [32]byte) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	RecoverFromBadSignal(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAFNConfig(opts *bind.TransactOpts, participants []common.Address, weights []*big.Int, weightThresholdForBlessing *big.Int, weightThresholdForBadSignal *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	VoteBad(opts *bind.TransactOpts) (*types.Transaction, error)

	VoteToBlessRoots(opts *bind.TransactOpts, rootsWithOrigin [][32]byte) (*types.Transaction, error)

	FilterAFNBadSignal(opts *bind.FilterOpts) (*AFNContractAFNBadSignalIterator, error)

	WatchAFNBadSignal(opts *bind.WatchOpts, sink chan<- *AFNContractAFNBadSignal) (event.Subscription, error)

	ParseAFNBadSignal(log types.Log) (*AFNContractAFNBadSignal, error)

	FilterAFNConfigSet(opts *bind.FilterOpts) (*AFNContractAFNConfigSetIterator, error)

	WatchAFNConfigSet(opts *bind.WatchOpts, sink chan<- *AFNContractAFNConfigSet) (event.Subscription, error)

	ParseAFNConfigSet(log types.Log) (*AFNContractAFNConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AFNContractOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AFNContractOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*AFNContractOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AFNContractOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AFNContractOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*AFNContractOwnershipTransferred, error)

	FilterRecoveredFromBadSignal(opts *bind.FilterOpts) (*AFNContractRecoveredFromBadSignalIterator, error)

	WatchRecoveredFromBadSignal(opts *bind.WatchOpts, sink chan<- *AFNContractRecoveredFromBadSignal) (event.Subscription, error)

	ParseRecoveredFromBadSignal(log types.Log) (*AFNContractRecoveredFromBadSignal, error)

	FilterRootBlessed(opts *bind.FilterOpts, root [][32]byte) (*AFNContractRootBlessedIterator, error)

	WatchRootBlessed(opts *bind.WatchOpts, sink chan<- *AFNContractRootBlessed, root [][32]byte) (event.Subscription, error)

	ParseRootBlessed(log types.Log) (*AFNContractRootBlessed, error)

	FilterVoteBad(opts *bind.FilterOpts, voter []common.Address) (*AFNContractVoteBadIterator, error)

	WatchVoteBad(opts *bind.WatchOpts, sink chan<- *AFNContractVoteBad, voter []common.Address) (event.Subscription, error)

	ParseVoteBad(log types.Log) (*AFNContractVoteBad, error)

	FilterVoteToBless(opts *bind.FilterOpts, voter []common.Address, root [][32]byte) (*AFNContractVoteToBlessIterator, error)

	WatchVoteToBless(opts *bind.WatchOpts, sink chan<- *AFNContractVoteToBless, voter []common.Address, root [][32]byte) (event.Subscription, error)

	ParseVoteToBless(log types.Log) (*AFNContractVoteToBless, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
