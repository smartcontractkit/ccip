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
	Bin: "0x60806040523480156200001157600080fd5b506040516200226e3803806200226e83398101604081905262000034916200082d565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000e1565b505050620000d78484848460016200018c60201b60201c565b5050505062000a10565b336001600160a01b038216036200013b5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b835185511415806200019d57508451155b80620001a7575082155b80620001b1575081155b80620001bb575080155b15620001da576040516306b7c75960e31b815260040160405180910390fd5b600060038054806020026020016040519081016040528092919081815260200182805480156200023457602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831162000215575b5050505050905060005b8151811015620002a157600260008383815181106200026157620002616200091b565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206000905580620002999062000947565b90506200023e565b50600484905560058390556006829055620002bb62000491565b620002c6816200056d565b8551600090620002de9060039060208a0190620006a7565b5060005b87518110156200041b5760006001600160a01b03168882815181106200030c576200030c6200091b565b60200260200101516001600160a01b0316036200033c576040516306b7c75960e31b815260040160405180910390fd5b8681815181106200035157620003516200091b565b60200260200101516000036200037a5760405163585b926360e01b815260040160405180910390fd5b8681815181106200038f576200038f6200091b565b6020026020010151600260008a8481518110620003b057620003b06200091b565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002081905550868181518110620003f157620003f16200091b565b60200260200101518262000406919062000963565b9150620004138162000947565b9050620002e2565b50848110806200042a57508381105b1562000449576040516306b7c75960e31b815260040160405180910390fd5b7f69af5b8b5b348d6b619cb6b338b5cfd865aa9e8cedd36a4a69257a9a07ebedaa878787876040516200048094939291906200097e565b60405180910390a150505050505050565b60006007805480602002602001604051908101604052809291908181526020018280548015620004eb57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311620004cc575b5050505050905060005b81518110156200055657600860008383815181106200051857620005186200091b565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191690556200054e8162000947565b9050620004f5565b50620005656007600062000711565b506000600955565b6000600c805480602002602001604051908101604052809291908181526020018280548015620005bd57602002820191906000526020600020905b815481526020019060010190808311620005a8575b5050505050905060005b815181101562000694576000828281518110620005e857620005e86200091b565b60200260200101519050600b60008281526020019081526020016000206000905560005b84518110156200067e57600d600083815260200190815260200160002060008683815181106200064057620006406200091b565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff19169055620006768162000947565b90506200060c565b5050806200068c9062000947565b9050620005c7565b50620006a3600c600062000711565b5050565b828054828255906000526020600020908101928215620006ff579160200282015b82811115620006ff57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620006c8565b506200070d92915062000734565b5090565b508054600082559060005260206000209081019062000731919062000734565b50565b5b808211156200070d576000815560010162000735565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156200078c576200078c6200074b565b604052919050565b60006001600160401b03821115620007b057620007b06200074b565b5060051b60200190565b600082601f830112620007cc57600080fd5b81516020620007e5620007df8362000794565b62000761565b82815260059290921b840181019181810190868411156200080557600080fd5b8286015b8481101562000822578051835291830191830162000809565b509695505050505050565b600080600080608085870312156200084457600080fd5b84516001600160401b03808211156200085c57600080fd5b818701915087601f8301126200087157600080fd5b8151602062000884620007df8362000794565b82815260059290921b8401810191818101908b841115620008a457600080fd5b948201945b83861015620008db5785516001600160a01b0381168114620008cb5760008081fd5b82529482019490820190620008a9565b918a0151919850909350505080821115620008f557600080fd5b506200090487828801620007ba565b604087015160609097015195989097509350505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016200095c576200095c62000931565b5060010190565b6000821982111562000979576200097962000931565b500190565b6080808252855190820181905260009060209060a0840190828901845b82811015620009c25781516001600160a01b0316845292840192908401906001016200099b565b5050508381038285015286518082528783019183019060005b81811015620009f957835183529284019291840191600101620009db565b505060408501969096525050506060015292915050565b61184e8062000a206000396000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c806379adb16e116100b25780639dc6edc711610081578063f2fde38b11610066578063f2fde38b14610327578063f438c9c01461033a578063ff888fb11461035557600080fd5b80639dc6edc714610309578063c3453fa51461031157600080fd5b806379adb16e1461027557806379ba5097146102b95780638da5cb5b146102c15780638e1d4e61146102e957600080fd5b80633cd4f66911610109578063508ede92116100ee578063508ede9214610208578063518565651461024d5780635aa68ac01461026057600080fd5b80633cd4f669146101f357806346f8e6d7146101fb57600080fd5b8063181f5a771461013b5780632cb145d41461018d5780632ea9537114610197578063365f15ec146101e0575b600080fd5b6101776040518060400160405280600981526020017f41464e20312e302e30000000000000000000000000000000000000000000000081525081565b6040516101849190611359565b60405180910390f35b610195610378565b005b6101d06101a53660046113f5565b73ffffffffffffffffffffffffffffffffffffffff1660009081526008602052604090205460ff1690565b6040519015158152602001610184565b6101956101ee366004611417565b6105c9565b61019561089f565b60055460095410156101d0565b6101d061021636600461148c565b6000908152600d6020908152604080832073ffffffffffffffffffffffffffffffffffffffff949094168352929052205460ff1690565b61019561025b3660046115c3565b610918565b610268610941565b60405161018491906116e4565b6102ab6102833660046113f5565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205490565b604051908152602001610184565b6101956109b0565b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610184565b6102ab6102f73660046116f7565b6000908152600b602052604090205490565b6006546102ab565b610319610aad565b604051610184929190611710565b6101956103353660046113f5565b610b27565b60045460055460408051928352602083019190915201610184565b6101d06103633660046116f7565b6000908152600a602052604090205460ff1690565b600554600954106103b5576040517fc28cc95000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336000818152600260205260408120549081900361041c576040517f669f262e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821660009081526008602052604090205460ff161561047c576040517f7c9a1cf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216600081815260086020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915560078054918201815582527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690921790915560098054839290610536908490611761565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316907fa5889da6c2d25ef72eaae82bb0b8acf51eeebdd6bd12f1a24360de7d9b9cfa289060200160405180910390a2600554600954106105c5576040514281527f73907f5e30313a1ab6e1815608b22b40911f1a7decec69d5df18a2298002bacb9060200160405180910390a15b5050565b60055460095410610606576040517fc28cc95000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360008181526002602052604081205490819003610668576040517f669f262e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610413565b60005b8381101561089857600085858381811061068757610687611779565b9050602002013590506106a9816000908152600a602052604090205460ff1690565b156106b45750610888565b6000818152600d6020908152604080832073ffffffffffffffffffffffffffffffffffffffff8816845290915290205460ff16156106f25750610888565b6000818152600d6020908152604080832073ffffffffffffffffffffffffffffffffffffffff88168452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055838352600b9091528120549081900361079357600c80546001810182556000919091527fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c7018290555b600061079f8583611761565b6000848152600b60205260409081902082905551909150839073ffffffffffffffffffffffffffffffffffffffff8816907f262f79a5a063a0af3e27989b0b0f0ae1e2c19257d27efe01a7f0cab7b3b470a4906107ff9089815260200190565b60405180910390a36004548110610884576000838152600a60205260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183907f719fab74b843fdceffa591cc0a3445a9dddc9e1e304471baed67e8408a1405c79061087b9084815260200190565b60405180910390a25b5050505b610891816117a8565b905061066b565b5050505050565b6108a7610b3b565b60055460095410156108e5576040517fe147761200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108ed610bbe565b6040517f3e48434bea67b1e259c2380d289dcb6372257ab2c37bc86f0e1acf83a7b07ac090600090a1565b610920610b3b565b61093b8484848460065460016109369190611761565b610cc7565b50505050565b606060038054806020026020016040519081016040528092919081815260200182805480156109a657602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161097b575b5050505050905090565b60015473ffffffffffffffffffffffffffffffffffffffff163314610a31576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610413565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60606000600760095481805480602002602001604051908101604052809291908181526020018280548015610b1857602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610aed575b50505050509150915091509091565b610b2f610b3b565b610b388161105c565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610bbc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610413565b565b60006007805480602002602001604051908101604052809291908181526020018280548015610c2357602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610bf8575b5050505050905060005b8151811015610cb25760086000838381518110610c4c57610c4c611779565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055610cab816117a8565b9050610c2d565b50610cbf6007600061129c565b506000600955565b83518551141580610cd757508451155b80610ce0575082155b80610ce9575081155b80610cf2575080155b15610d29576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006003805480602002602001604051908101604052809291908181526020018280548015610d8e57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610d63575b5050505050905060005b8151811015610e0e5760026000838381518110610db757610db7611779565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905580610e07906117a8565b9050610d98565b50600484905560058390556006829055610e26610bbe565b610e2f81611151565b8551600090610e459060039060208a01906112ba565b5060005b8751811015610fd157600073ffffffffffffffffffffffffffffffffffffffff16888281518110610e7c57610e7c611779565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1603610ed1576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b868181518110610ee357610ee3611779565b6020026020010151600003610f24576040517f585b926300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b868181518110610f3657610f36611779565b6020026020010151600260008a8481518110610f5457610f54611779565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550868181518110610fac57610fac611779565b602002602001015182610fbf9190611761565b9150610fca816117a8565b9050610e49565b5084811080610fdf57508381105b15611016576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f69af5b8b5b348d6b619cb6b338b5cfd865aa9e8cedd36a4a69257a9a07ebedaa8787878760405161104b94939291906117e0565b60405180910390a150505050505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036110db576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610413565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000600c80548060200260200160405190810160405280929190818152602001828054801561119f57602002820191906000526020600020905b81548152602001906001019080831161118b575b5050505050905060005b81518110156112935760008282815181106111c6576111c6611779565b60200260200101519050600b60008281526020019081526020016000206000905560005b845181101561128057600d6000838152602001908152602001600020600086838151811061121a5761121a611779565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055611279816117a8565b90506111ea565b50508061128c906117a8565b90506111a9565b506105c5600c60005b5080546000825590600052602060002090810190610b389190611344565b828054828255906000526020600020908101928215611334579160200282015b8281111561133457825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9091161782556020909201916001909101906112da565b50611340929150611344565b5090565b5b808211156113405760008155600101611345565b600060208083528351808285015260005b818110156113865785810183015185820160400152820161136a565b81811115611398576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146113f057600080fd5b919050565b60006020828403121561140757600080fd5b611410826113cc565b9392505050565b6000806020838503121561142a57600080fd5b823567ffffffffffffffff8082111561144257600080fd5b818501915085601f83011261145657600080fd5b81358181111561146557600080fd5b8660208260051b850101111561147a57600080fd5b60209290920196919550909350505050565b6000806040838503121561149f57600080fd5b6114a8836113cc565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561152c5761152c6114b6565b604052919050565b600067ffffffffffffffff82111561154e5761154e6114b6565b5060051b60200190565b600082601f83011261156957600080fd5b8135602061157e61157983611534565b6114e5565b82815260059290921b8401810191818101908684111561159d57600080fd5b8286015b848110156115b857803583529183019183016115a1565b509695505050505050565b600080600080608085870312156115d957600080fd5b843567ffffffffffffffff808211156115f157600080fd5b818701915087601f83011261160557600080fd5b8135602061161561157983611534565b82815260059290921b8401810191818101908b84111561163457600080fd5b948201945b838610156116595761164a866113cc565b82529482019490820190611639565b9850508801359250508082111561166f57600080fd5b5061167c87828801611558565b949794965050505060408301359260600135919050565b600081518084526020808501945080840160005b838110156116d957815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016116a7565b509495945050505050565b6020815260006114106020830184611693565b60006020828403121561170957600080fd5b5035919050565b6040815260006117236040830185611693565b90508260208301529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561177457611774611732565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036117d9576117d9611732565b5060010190565b6080815260006117f36080830187611693565b82810360208481019190915286518083528782019282019060005b8181101561182a5784518352938301939183019160010161180e565b50506040850196909652505050606001529291505056fea164736f6c634300080f000a",
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
