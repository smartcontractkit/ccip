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
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"parties\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"goodQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"badQuorum\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AFNBadSignal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"committeeVersion\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structAFNInterface.Heartbeat\",\"name\":\"heartbeat\",\"type\":\"tuple\"}],\"name\":\"AFNHeartbeat\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"BadVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"parties\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goodQuorum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"badQuorum\",\"type\":\"uint256\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"GoodVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RecoveredFromBadSignal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBadVotersAndVotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitteeVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"getGoodVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"party\",\"type\":\"address\"}],\"name\":\"getLastGoodVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastHeartbeat\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"committeeVersion\",\"type\":\"uint256\"}],\"internalType\":\"structAFNInterface.Heartbeat\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParties\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuorums\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"good\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bad\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"party\",\"type\":\"address\"}],\"name\":\"getWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasBadSignal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"party\",\"type\":\"address\"}],\"name\":\"hasVotedBad\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"parties\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"goodQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"badQuorum\",\"type\":\"uint256\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteBad\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"voteGood\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001c1238038062001c128339810160408190526200003491620004e4565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000e2565b505050620000d8848484846001806200018e60201b60201c565b5050505062000719565b6001600160a01b0381163314156200013d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b845186511415806200019f57508551155b80620001a9575083155b80620001b3575082155b80620001bd575081155b80620001c7575080155b15620001e6576040516306b7c75960e31b815260040160405180910390fd5b600060038054806020026020016040519081016040528092919081815260200182805480156200024057602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831162000221575b5050505050905060005b8151811015620002b1576000600260008484815181106200026f576200026f620006ed565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020819055508080620002a890620006c3565b9150506200024a565b5060048590556005849055600683905560078290558651620002db9060039060208a0190620003e9565b5060005b8751811015620003a057868181518110620002fe57620002fe620006ed565b602002602001015160001415620003285760405163585b926360e01b815260040160405180910390fd5b8681815181106200033d576200033d620006ed565b6020026020010151600260008a84815181106200035e576200035e620006ed565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000208190555080806200039790620006c3565b915050620002df565b507f973528bd06888da95feea86250f8fc2401650b59cbb4cb47dda24c2b79bd9d0487878787604051620003d89493929190620005d8565b60405180910390a150505050505050565b82805482825590600052602060002090810192821562000441579160200282015b828111156200044157825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200040a565b506200044f92915062000453565b5090565b5b808211156200044f576000815560010162000454565b600082601f8301126200047c57600080fd5b81516020620004956200048f836200069d565b6200066a565b80838252828201915082860187848660051b8901011115620004b657600080fd5b60005b85811015620004d757815184529284019290840190600101620004b9565b5090979650505050505050565b60008060008060808587031215620004fb57600080fd5b84516001600160401b03808211156200051357600080fd5b818701915087601f8301126200052857600080fd5b815160206200053b6200048f836200069d565b8083825282820191508286018c848660051b89010111156200055c57600080fd5b600096505b84871015620005975780516001600160a01b03811681146200058257600080fd5b83526001969096019591830191830162000561565b50918a0151919850909350505080821115620005b257600080fd5b50620005c1878288016200046a565b604087015160609097015195989097509350505050565b6080808252855190820181905260009060209060a0840190828901845b828110156200061c5781516001600160a01b031684529284019290840190600101620005f5565b5050508381038285015286518082528783019183019060005b81811015620006535783518352928401929184019160010162000635565b505060408501969096525050506060015292915050565b604051601f8201601f191681016001600160401b038111828210171562000695576200069562000703565b604052919050565b60006001600160401b03821115620006b957620006b962000703565b5060051b60200190565b6000600019821415620006e657634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6114e980620007296000396000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c8063a8311aa8116100b2578063ce74602411610081578063d08041b111610066578063d08041b1146102cd578063f2fde38b146102e0578063fb5d8c74146102f357600080fd5b8063ce746024146102ba578063cf72b39b146102c257600080fd5b8063a8311aa814610246578063ac6c52511461025b578063acea368b14610291578063c3453fa5146102a457600080fd5b806379ba50971161010957806392325ab2116100ee57806392325ab2146102085780639f8743f714610236578063a60e65111461023e57600080fd5b806379ba5097146101d85780638da5cb5b146101e057600080fd5b80632cb145d41461013b5780632ea9537114610145578063343157b41461019357806377642b73146101bd575b600080fd5b610143610329565b005b61017e610153366004611192565b73ffffffffffffffffffffffffffffffffffffffff166000908152600d602052604090205460ff1690565b60405190151581526020015b60405180910390f35b61019b61055e565b604080518251815260208084015190820152918101519082015260600161018a565b6004546005546040805192835260208301919091520161018a565b6101436105a7565b60005460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161018a565b61022861021636600461128b565b6000908152600c602052604090205490565b60405190815260200161018a565b600654610228565b600754610228565b61024e6106a4565b60405161018a91906112f5565b610228610269366004611192565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205490565b61014361029f3660046111b4565b610713565b6102ac61074a565b60405161018a929190611369565b6101436107c4565b60105460ff1661017e565b6101436102db36600461128b565b61096f565b6101436102ee366004611192565b610c00565b610228610301366004611192565b73ffffffffffffffffffffffffffffffffffffffff166000908152600b602052604090205490565b60105460ff1615610366576040517fc28cc95000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b33600081815260026020526040902054806103ca576040517f669f262e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600d602052604090205460ff161561042a576040517f7c9a1cf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000818152600d6020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155600e8054918201815582527fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd0180547fffffffffffffffffffffffff000000000000000000000000000000000000000016909217909155600f80548392906104e49084906113fe565b9091555050600554600f541061055a57601080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f73907f5e30313a1ab6e1815608b22b40911f1a7decec69d5df18a2298002bacb906105519042815260200190565b60405180910390a15b5050565b61058260405180606001604052806000815260200160008152602001600081525090565b506040805160608101825260085481526009546020820152600a549181019190915290565b60015473ffffffffffffffffffffffffffffffffffffffff163314610628576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016103c1565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600380548060200260200160405190810160405280929190818152602001828054801561070957602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116106de575b5050505050905090565b61071b610c14565b61074484848484600654600161073191906113fe565b60075461073f9060016113fe565b610c97565b50505050565b60606000600e600f54818054806020026020016040519081016040528092919081815260200182805480156107b557602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161078a575b50505050509150915091509091565b6107cc610c14565b60105460ff16610808576040517fe147761200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600e80548060200260200160405190810160405280929190818152602001828054801561086d57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610842575b5050505050905060005b8151811015610907576000600d60008484815181106108985761089861147e565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055806108ff81611416565b915050610877565b506000600f81905561091b90600e9061103a565b601080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f3e48434bea67b1e259c2380d289dcb6372257ab2c37bc86f0e1acf83a7b07ac090600090a150565b6006548181146109b5576040517f43a010e100000000000000000000000000000000000000000000000000000000815260048101829052602481018390526044016103c1565b60105460ff16156109f2576040517fc28cc95000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b33600081815260026020526040902054610a50576040517f669f262e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016103c1565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600b6020526040902054821415610aaf576040517f7c9a1cf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166000908152600b602090815260408083208590556002825280832054858452600c9092528220805491929091610afd9084906113fe565b90915550506040805173ffffffffffffffffffffffffffffffffffffffff83168152602081018490527f5489e43df72470c733e49d6f7bc612d52f64600fb2801593290ec32fcf144791910160405180910390a16004546000838152600c602052604090205410610bfb576040805160608101825283815267ffffffffffffffff4216602082018190526007549282018390526008859055600955600a9190915560068054906000610bae83611416565b9091555050604080518251815260208084015190820152828201518183015290517f90b45dcfd48782731999668957597f8b47e29aaa1d53ef2ad07612429777bed39181900360600190a1505b505050565b610c08610c14565b610c1181610f44565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610c95576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016103c1565b565b84518651141580610ca757508551155b80610cb0575083155b80610cb9575082155b80610cc2575081155b80610ccb575080155b15610d02576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006003805480602002602001604051908101604052809291908181526020018280548015610d6757602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610d3c575b5050505050905060005b8151811015610deb57600060026000848481518110610d9257610d9261147e565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508080610de390611416565b915050610d71565b5060048590556005849055600683905560078290558651610e139060039060208a0190611058565b5060005b8751811015610efd57868181518110610e3257610e3261147e565b602002602001015160001415610e74576040517f585b926300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b868181518110610e8657610e8661147e565b6020026020010151600260008a8481518110610ea457610ea461147e565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508080610ef590611416565b915050610e17565b507f973528bd06888da95feea86250f8fc2401650b59cbb4cb47dda24c2b79bd9d0487878787604051610f339493929190611308565b60405180910390a150505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116331415610fc4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016103c1565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b5080546000825590600052602060002090810190610c1191906110e2565b8280548282559060005260206000209081019282156110d2579160200282015b828111156110d257825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190611078565b506110de9291506110e2565b5090565b5b808211156110de57600081556001016110e3565b803573ffffffffffffffffffffffffffffffffffffffff8116811461111b57600080fd5b919050565b600082601f83011261113157600080fd5b81356020611146611141836113da565b61138b565b80838252828201915082860187848660051b890101111561116657600080fd5b60005b8581101561118557813584529284019290840190600101611169565b5090979650505050505050565b6000602082840312156111a457600080fd5b6111ad826110f7565b9392505050565b600080600080608085870312156111ca57600080fd5b843567ffffffffffffffff808211156111e257600080fd5b818701915087601f8301126111f657600080fd5b81356020611206611141836113da565b8083825282820191508286018c848660051b890101111561122657600080fd5b600096505b848710156112505761123c816110f7565b83526001969096019591830191830161122b565b509850508801359250508082111561126757600080fd5b5061127487828801611120565b949794965050505060408301359260600135919050565b60006020828403121561129d57600080fd5b5035919050565b600081518084526020808501945080840160005b838110156112ea57815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016112b8565b509495945050505050565b6020815260006111ad60208301846112a4565b60808152600061131b60808301876112a4565b82810360208481019190915286518083528782019282019060005b8181101561135257845183529383019391830191600101611336565b505060408501969096525050506060015292915050565b60408152600061137c60408301856112a4565b90508260208301529392505050565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156113d2576113d26114ad565b604052919050565b600067ffffffffffffffff8211156113f4576113f46114ad565b5060051b60200190565b600082198211156114115761141161144f565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156114485761144861144f565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c6343000806000a",
}

var AFNContractABI = AFNContractMetaData.ABI

var AFNContractBin = AFNContractMetaData.Bin

func DeployAFNContract(auth *bind.TransactOpts, backend bind.ContractBackend, parties []common.Address, weights []*big.Int, goodQuorum *big.Int, badQuorum *big.Int) (common.Address, *types.Transaction, *AFNContract, error) {
	parsed, err := AFNContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AFNContractBin), backend, parties, weights, goodQuorum, badQuorum)
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

func (_AFNContract *AFNContractCaller) GetGoodVotes(opts *bind.CallOpts, round *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getGoodVotes", round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AFNContract *AFNContractSession) GetGoodVotes(round *big.Int) (*big.Int, error) {
	return _AFNContract.Contract.GetGoodVotes(&_AFNContract.CallOpts, round)
}

func (_AFNContract *AFNContractCallerSession) GetGoodVotes(round *big.Int) (*big.Int, error) {
	return _AFNContract.Contract.GetGoodVotes(&_AFNContract.CallOpts, round)
}

func (_AFNContract *AFNContractCaller) GetLastGoodVote(opts *bind.CallOpts, party common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getLastGoodVote", party)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AFNContract *AFNContractSession) GetLastGoodVote(party common.Address) (*big.Int, error) {
	return _AFNContract.Contract.GetLastGoodVote(&_AFNContract.CallOpts, party)
}

func (_AFNContract *AFNContractCallerSession) GetLastGoodVote(party common.Address) (*big.Int, error) {
	return _AFNContract.Contract.GetLastGoodVote(&_AFNContract.CallOpts, party)
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

func (_AFNContract *AFNContractCaller) GetParties(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getParties")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_AFNContract *AFNContractSession) GetParties() ([]common.Address, error) {
	return _AFNContract.Contract.GetParties(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCallerSession) GetParties() ([]common.Address, error) {
	return _AFNContract.Contract.GetParties(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCaller) GetQuorums(opts *bind.CallOpts) (GetQuorums,

	error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getQuorums")

	outstruct := new(GetQuorums)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Good = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Bad = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_AFNContract *AFNContractSession) GetQuorums() (GetQuorums,

	error) {
	return _AFNContract.Contract.GetQuorums(&_AFNContract.CallOpts)
}

func (_AFNContract *AFNContractCallerSession) GetQuorums() (GetQuorums,

	error) {
	return _AFNContract.Contract.GetQuorums(&_AFNContract.CallOpts)
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

func (_AFNContract *AFNContractCaller) GetWeight(opts *bind.CallOpts, party common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "getWeight", party)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AFNContract *AFNContractSession) GetWeight(party common.Address) (*big.Int, error) {
	return _AFNContract.Contract.GetWeight(&_AFNContract.CallOpts, party)
}

func (_AFNContract *AFNContractCallerSession) GetWeight(party common.Address) (*big.Int, error) {
	return _AFNContract.Contract.GetWeight(&_AFNContract.CallOpts, party)
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

func (_AFNContract *AFNContractCaller) HasVotedBad(opts *bind.CallOpts, party common.Address) (bool, error) {
	var out []interface{}
	err := _AFNContract.contract.Call(opts, &out, "hasVotedBad", party)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AFNContract *AFNContractSession) HasVotedBad(party common.Address) (bool, error) {
	return _AFNContract.Contract.HasVotedBad(&_AFNContract.CallOpts, party)
}

func (_AFNContract *AFNContractCallerSession) HasVotedBad(party common.Address) (bool, error) {
	return _AFNContract.Contract.HasVotedBad(&_AFNContract.CallOpts, party)
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

func (_AFNContract *AFNContractTransactor) SetConfig(opts *bind.TransactOpts, parties []common.Address, weights []*big.Int, goodQuorum *big.Int, badQuorum *big.Int) (*types.Transaction, error) {
	return _AFNContract.contract.Transact(opts, "setConfig", parties, weights, goodQuorum, badQuorum)
}

func (_AFNContract *AFNContractSession) SetConfig(parties []common.Address, weights []*big.Int, goodQuorum *big.Int, badQuorum *big.Int) (*types.Transaction, error) {
	return _AFNContract.Contract.SetConfig(&_AFNContract.TransactOpts, parties, weights, goodQuorum, badQuorum)
}

func (_AFNContract *AFNContractTransactorSession) SetConfig(parties []common.Address, weights []*big.Int, goodQuorum *big.Int, badQuorum *big.Int) (*types.Transaction, error) {
	return _AFNContract.Contract.SetConfig(&_AFNContract.TransactOpts, parties, weights, goodQuorum, badQuorum)
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

type AFNContractConfigSetIterator struct {
	Event *AFNContractConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AFNContractConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AFNContractConfigSet)
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
		it.Event = new(AFNContractConfigSet)
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

func (it *AFNContractConfigSetIterator) Error() error {
	return it.fail
}

func (it *AFNContractConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AFNContractConfigSet struct {
	Parties    []common.Address
	Weights    []*big.Int
	GoodQuorum *big.Int
	BadQuorum  *big.Int
	Raw        types.Log
}

func (_AFNContract *AFNContractFilterer) FilterConfigSet(opts *bind.FilterOpts) (*AFNContractConfigSetIterator, error) {

	logs, sub, err := _AFNContract.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &AFNContractConfigSetIterator{contract: _AFNContract.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_AFNContract *AFNContractFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *AFNContractConfigSet) (event.Subscription, error) {

	logs, sub, err := _AFNContract.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AFNContractConfigSet)
				if err := _AFNContract.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_AFNContract *AFNContractFilterer) ParseConfigSet(log types.Log) (*AFNContractConfigSet, error) {
	event := new(AFNContractConfigSet)
	if err := _AFNContract.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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
type GetQuorums struct {
	Good *big.Int
	Bad  *big.Int
}

func (_AFNContract *AFNContract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _AFNContract.abi.Events["AFNBadSignal"].ID:
		return _AFNContract.ParseAFNBadSignal(log)
	case _AFNContract.abi.Events["AFNHeartbeat"].ID:
		return _AFNContract.ParseAFNHeartbeat(log)
	case _AFNContract.abi.Events["BadVote"].ID:
		return _AFNContract.ParseBadVote(log)
	case _AFNContract.abi.Events["ConfigSet"].ID:
		return _AFNContract.ParseConfigSet(log)
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

func (AFNContractAFNHeartbeat) Topic() common.Hash {
	return common.HexToHash("0x90b45dcfd48782731999668957597f8b47e29aaa1d53ef2ad07612429777bed3")
}

func (AFNContractBadVote) Topic() common.Hash {
	return common.HexToHash("0x0b21c4350e3db4e6e412b398113b3769fa4fcf4582c88579705b3d42002a41fd")
}

func (AFNContractConfigSet) Topic() common.Hash {
	return common.HexToHash("0x973528bd06888da95feea86250f8fc2401650b59cbb4cb47dda24c2b79bd9d04")
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

	GetGoodVotes(opts *bind.CallOpts, round *big.Int) (*big.Int, error)

	GetLastGoodVote(opts *bind.CallOpts, party common.Address) (*big.Int, error)

	GetLastHeartbeat(opts *bind.CallOpts) (AFNInterfaceHeartbeat, error)

	GetParties(opts *bind.CallOpts) ([]common.Address, error)

	GetQuorums(opts *bind.CallOpts) (GetQuorums,

		error)

	GetRound(opts *bind.CallOpts) (*big.Int, error)

	GetWeight(opts *bind.CallOpts, party common.Address) (*big.Int, error)

	HasBadSignal(opts *bind.CallOpts) (bool, error)

	HasVotedBad(opts *bind.CallOpts, party common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Recover(opts *bind.TransactOpts) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, parties []common.Address, weights []*big.Int, goodQuorum *big.Int, badQuorum *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	VoteBad(opts *bind.TransactOpts) (*types.Transaction, error)

	VoteGood(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error)

	FilterAFNBadSignal(opts *bind.FilterOpts) (*AFNContractAFNBadSignalIterator, error)

	WatchAFNBadSignal(opts *bind.WatchOpts, sink chan<- *AFNContractAFNBadSignal) (event.Subscription, error)

	ParseAFNBadSignal(log types.Log) (*AFNContractAFNBadSignal, error)

	FilterAFNHeartbeat(opts *bind.FilterOpts) (*AFNContractAFNHeartbeatIterator, error)

	WatchAFNHeartbeat(opts *bind.WatchOpts, sink chan<- *AFNContractAFNHeartbeat) (event.Subscription, error)

	ParseAFNHeartbeat(log types.Log) (*AFNContractAFNHeartbeat, error)

	FilterBadVote(opts *bind.FilterOpts) (*AFNContractBadVoteIterator, error)

	WatchBadVote(opts *bind.WatchOpts, sink chan<- *AFNContractBadVote) (event.Subscription, error)

	ParseBadVote(log types.Log) (*AFNContractBadVote, error)

	FilterConfigSet(opts *bind.FilterOpts) (*AFNContractConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *AFNContractConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*AFNContractConfigSet, error)

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
