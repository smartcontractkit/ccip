// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token_price_ocr

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
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
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
	_ = abi.ConvertType
)

type InternalTokenPriceUpdate struct {
	SourceToken common.Address
	UsdPerToken *big.Int
}

type TokenPriceOCRReport struct {
	PriceUpdates []InternalTokenPriceUpdate
}

var TokenPriceOCRMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPriceRegistry\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"ForkedChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"latestSequenceNumber\",\"type\":\"uint64\"}],\"name\":\"NonIncreasingSequenceNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint224\",\"name\":\"usdPerToken\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"name\":\"priceUpdates\",\"type\":\"tuple[]\"}],\"internalType\":\"structTokenPriceOCR.Report\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"exposeForEncoding\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getGasPriceUpdate\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"value\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPriceUpdate\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"value\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR3Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceRegistry\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"name\":\"setPriceRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620023e9380380620023e9833981016040819052620000349162000196565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000eb565b50504660805250600880546001600160a01b0319166001600160a01b0392909216919091179055620001c8565b336001600160a01b03821603620001455760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001a957600080fd5b81516001600160a01b0381168114620001c157600080fd5b9392505050565b6080516121fe620001eb60003960008181610eab0152610ef701526121fe6000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c806381ff70481161008c578063b1dc65a411610066578063b1dc65a41461021f578063ca137fba14610232578063f2fde38b146102e8578063fc77d708146102fb57600080fd5b806381ff70481461018e5780638da5cb5b146101be578063afcb95d7146101e657600080fd5b8063666cab8d116100bd578063666cab8d1461015e5780636a11ee901461017357806379ba50971461018657600080fd5b8063181f5a77146100e457806348dab75114610136578063508ee9de14610149575b600080fd5b6101206040518060400160405280601381526020017f546f6b656e50726963654f435220312e302e300000000000000000000000000081525081565b60405161012d9190611814565b60405180910390f35b61012061014436600461193e565b61036c565b61015c610157366004611a62565b610395565b005b6101666103e4565b60405161012d9190611ad0565b61015c610181366004611c0e565b610453565b61015c610c6b565b6004546002546040805163ffffffff8085168252640100000000909404909316602084015282015260600161012d565b60005460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161012d565b600254600454604080516001815260208101939093526801000000000000000090910467ffffffffffffffff169082015260600161012d565b61015c61022d366004611d27565b610d68565b6102af610240366004611a62565b73ffffffffffffffffffffffffffffffffffffffff166000908152600960205260409020547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116917c010000000000000000000000000000000000000000000000000000000090910463ffffffff1690565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff909316835263ffffffff90911660208301520161012d565b61015c6102f6366004611a62565b6113e2565b6102af610309366004611e0c565b67ffffffffffffffff166000908152600a60205260409020547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116917c010000000000000000000000000000000000000000000000000000000090910463ffffffff1690565b60608160405160200161037f9190611e96565b6040516020818303038152906040529050919050565b61039d6113f6565b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6060600780548060200260200160405190810160405280929190818152602001828054801561044957602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161041e575b5050505050905090565b855185518560ff16601f8311156104cb576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064015b60405180910390fd5b80600003610535576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016104c2565b8183146105c3576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e0000000000000000000000000000000000000000000000000000000060648201526084016104c2565b6105ce816003611ee8565b8311610636576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016104c2565b61063e6113f6565b60065460005b8181101561073a57600560006006838154811061066357610663611f05565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000169055600780546005929190849081106106d3576106d3611f05565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905561073381611f34565b9050610644565b50895160005b81811015610b135760008c828151811061075c5761075c611f05565b602002602001015190506000600281111561077957610779611f6c565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260056020526040902054610100900460ff1660028111156107b8576107b8611f6c565b1461081f576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016104c2565b73ffffffffffffffffffffffffffffffffffffffff811661086c576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260ff83168152602081016001905273ffffffffffffffffffffffffffffffffffffffff821660009081526005602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561091c5761091c611f6c565b021790555090505060008c838151811061093857610938611f05565b602002602001015190506000600281111561095557610955611f6c565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260056020526040902054610100900460ff16600281111561099457610994611f6c565b146109fb576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016104c2565b73ffffffffffffffffffffffffffffffffffffffff8116610a48576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260ff84168152602081016002905273ffffffffffffffffffffffffffffffffffffffff821660009081526005602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115610af857610af8611f6c565b0217905550905050505080610b0c90611f34565b9050610740565b508a51610b279060069060208e01906116f2565b508951610b3b9060079060208d01906116f2565b506003805460ff838116610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000909216908c161717905560048054610bc1914691309190600090610b939063ffffffff16611f9b565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168e8e8e8e8e8e611479565b600260000181905550600060048054906101000a900463ffffffff169050436004806101000a81548163ffffffff021916908363ffffffff1602179055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0581600260000154600460009054906101000a900463ffffffff168f8f8f8f8f8f604051610c5599989796959493929190611fbe565b60405180910390a1505050505050505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610cec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016104c2565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60045460208901359067ffffffffffffffff68010000000000000000909104811690821611610deb57600480546040517f6e376b6600000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808516938201939093526801000000000000000090910490911660248201526044016104c2565b610df6888883611524565b600480547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8416021790556040805160608101825260025480825260035460ff808216602085015261010090910416928201929092528a35918214610ea85780516040517f93df584c0000000000000000000000000000000000000000000000000000000081526004810191909152602481018390526044016104c2565b467f000000000000000000000000000000000000000000000000000000000000000014610f29576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201524660248201526044016104c2565b6040805183815267ffffffffffffffff851660208201527fe893c2681d327421d89e1cb54fbe64645b4dcea668d6826130b62cf4c6eefea2910160405180910390a16020810151610f7b906001612054565b60ff168714610fb6576040517f71253a2500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b868514610fef576040517fa75d88af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526005602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561103257611032611f6c565b600281111561104357611043611f6c565b905250905060028160200151600281111561106057611060611f6c565b1480156110a757506007816000015160ff168154811061108257611082611f05565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b6110dd576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060006110eb866020611ee8565b6110f6896020611ee8565b6111028c61014461206d565b61110c919061206d565b611116919061206d565b905036811461115a576040517f8e1192e1000000000000000000000000000000000000000000000000000000008152600481018290523660248201526044016104c2565b5060008a8a60405161116d929190612080565b604051908190038120611184918e90602001612090565b6040516020818303038152906040528051906020012090506111a461177c565b8860005b818110156113d15760006001858a84602081106111c7576111c7611f05565b6111d491901a601b612054565b8f8f868181106111e6576111e6611f05565b905060200201358e8e878181106111ff576111ff611f05565b905060200201356040516000815260200160405260405161123c949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa15801561125e573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff8116600090815260056020908152848220848601909552845460ff80821686529397509195509293928401916101009091041660028111156112e1576112e1611f6c565b60028111156112f2576112f2611f6c565b905250905060018160200151600281111561130f5761130f611f6c565b14611346576040517fca31867a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051859060ff16601f811061135d5761135d611f05565b602002015115611399576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600185826000015160ff16601f81106113b4576113b4611f05565b91151560209092020152506113ca905081611f34565b90506111a8565b505050505050505050505050505050565b6113ea6113f6565b6113f3816115fd565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314611477576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016104c2565b565b6000808a8a8a8a8a8a8a8a8a60405160200161149d999897969594939291906120a4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b60006115328385018561193e565b6008546040805180820182528351815281516000808252602082810190945294955073ffffffffffffffffffffffffffffffffffffffff90931693633937306f939192830191906115a5565b604080518082019091526000808252602082015281526020019060019003908161157e5790505b508152506040518263ffffffff1660e01b81526004016115c59190612139565b600060405180830381600087803b1580156115df57600080fd5b505af11580156115f3573d6000803e3d6000fd5b5050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff82160361167c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016104c2565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b82805482825590600052602060002090810192821561176c579160200282015b8281111561176c57825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190611712565b5061177892915061179b565b5090565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115611778576000815560010161179c565b6000815180845260005b818110156117d6576020818501810151868301820152016117ba565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061182760208301846117b0565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516020810167ffffffffffffffff811182821017156118805761188061182e565b60405290565b6040805190810167ffffffffffffffff811182821017156118805761188061182e565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156118f0576118f061182e565b604052919050565b600067ffffffffffffffff8211156119125761191261182e565b5060051b60200190565b73ffffffffffffffffffffffffffffffffffffffff811681146113f357600080fd5b6000602080838503121561195157600080fd5b823567ffffffffffffffff8082111561196957600080fd5b818501915082828703121561197d57600080fd5b61198561185d565b82358281111561199457600080fd5b80840193505086601f8401126119a957600080fd5b823591506119be6119b9836118f8565b6118a9565b82815260069290921b830184019184810190888411156119dd57600080fd5b938501935b83851015611a55576040858a0312156119fb5760008081fd5b611a03611886565b8535611a0e8161191c565b8152858701357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff81168114611a3f5760008081fd5b81880152825260409490940193908501906119e2565b8252509695505050505050565b600060208284031215611a7457600080fd5b81356118278161191c565b600081518084526020808501945080840160005b83811015611ac557815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611a93565b509495945050505050565b6020815260006118276020830184611a7f565b600082601f830112611af457600080fd5b81356020611b046119b9836118f8565b82815260059290921b84018101918181019086841115611b2357600080fd5b8286015b84811015611b47578035611b3a8161191c565b8352918301918301611b27565b509695505050505050565b803560ff81168114611b6357600080fd5b919050565b600082601f830112611b7957600080fd5b813567ffffffffffffffff811115611b9357611b9361182e565b611bc460207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016118a9565b818152846020838601011115611bd957600080fd5b816020850160208301376000918101602001919091529392505050565b803567ffffffffffffffff81168114611b6357600080fd5b60008060008060008060c08789031215611c2757600080fd5b863567ffffffffffffffff80821115611c3f57600080fd5b611c4b8a838b01611ae3565b97506020890135915080821115611c6157600080fd5b611c6d8a838b01611ae3565b9650611c7b60408a01611b52565b95506060890135915080821115611c9157600080fd5b611c9d8a838b01611b68565b9450611cab60808a01611bf6565b935060a0890135915080821115611cc157600080fd5b50611cce89828a01611b68565b9150509295509295509295565b60008083601f840112611ced57600080fd5b50813567ffffffffffffffff811115611d0557600080fd5b6020830191508360208260051b8501011115611d2057600080fd5b9250929050565b60008060008060008060008060e0898b031215611d4357600080fd5b606089018a811115611d5457600080fd5b8998503567ffffffffffffffff80821115611d6e57600080fd5b818b0191508b601f830112611d8257600080fd5b813581811115611d9157600080fd5b8c6020828501011115611da357600080fd5b6020830199508098505060808b0135915080821115611dc157600080fd5b611dcd8c838d01611cdb565b909750955060a08b0135915080821115611de657600080fd5b50611df38b828c01611cdb565b999c989b50969995989497949560c00135949350505050565b600060208284031215611e1e57600080fd5b61182782611bf6565b600081518084526020808501945080840160005b83811015611ac5578151805173ffffffffffffffffffffffffffffffffffffffff1688528301517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168388015260409096019590820190600101611e3b565b6020815260008251602080840152611eb16040840182611e27565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082028115828204841417611eff57611eff611eb9565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611f6557611f65611eb9565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600063ffffffff808316818103611fb457611fb4611eb9565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152611fee8184018a611a7f565b905082810360808401526120028189611a7f565b905060ff871660a084015282810360c084015261201f81876117b0565b905067ffffffffffffffff851660e084015282810361010084015261204481856117b0565b9c9b505050505050505050505050565b60ff8181168382160190811115611eff57611eff611eb9565b80820180821115611eff57611eff611eb9565b8183823760009101908152919050565b828152606082602083013760800192915050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526120eb8285018b611a7f565b915083820360808501526120ff828a611a7f565b915060ff881660a085015283820360c085015261211c82886117b0565b90861660e0850152838103610100850152905061204481856117b0565b600060208083528351604080838601526121566060860183611e27565b868401518682037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00183880152805180835290850193506000918501905b808310156121e5578451805167ffffffffffffffff1683528601517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1686830152938501936001929092019190830190612194565b5097965050505050505056fea164736f6c6343000813000a",
}

var TokenPriceOCRABI = TokenPriceOCRMetaData.ABI

var TokenPriceOCRBin = TokenPriceOCRMetaData.Bin

func DeployTokenPriceOCR(auth *bind.TransactOpts, backend bind.ContractBackend, priceRegistry common.Address) (common.Address, *types.Transaction, *TokenPriceOCR, error) {
	parsed, err := TokenPriceOCRMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenPriceOCRBin), backend, priceRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenPriceOCR{address: address, abi: *parsed, TokenPriceOCRCaller: TokenPriceOCRCaller{contract: contract}, TokenPriceOCRTransactor: TokenPriceOCRTransactor{contract: contract}, TokenPriceOCRFilterer: TokenPriceOCRFilterer{contract: contract}}, nil
}

type TokenPriceOCR struct {
	address common.Address
	abi     abi.ABI
	TokenPriceOCRCaller
	TokenPriceOCRTransactor
	TokenPriceOCRFilterer
}

type TokenPriceOCRCaller struct {
	contract *bind.BoundContract
}

type TokenPriceOCRTransactor struct {
	contract *bind.BoundContract
}

type TokenPriceOCRFilterer struct {
	contract *bind.BoundContract
}

type TokenPriceOCRSession struct {
	Contract     *TokenPriceOCR
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TokenPriceOCRCallerSession struct {
	Contract *TokenPriceOCRCaller
	CallOpts bind.CallOpts
}

type TokenPriceOCRTransactorSession struct {
	Contract     *TokenPriceOCRTransactor
	TransactOpts bind.TransactOpts
}

type TokenPriceOCRRaw struct {
	Contract *TokenPriceOCR
}

type TokenPriceOCRCallerRaw struct {
	Contract *TokenPriceOCRCaller
}

type TokenPriceOCRTransactorRaw struct {
	Contract *TokenPriceOCRTransactor
}

func NewTokenPriceOCR(address common.Address, backend bind.ContractBackend) (*TokenPriceOCR, error) {
	abi, err := abi.JSON(strings.NewReader(TokenPriceOCRABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindTokenPriceOCR(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenPriceOCR{address: address, abi: abi, TokenPriceOCRCaller: TokenPriceOCRCaller{contract: contract}, TokenPriceOCRTransactor: TokenPriceOCRTransactor{contract: contract}, TokenPriceOCRFilterer: TokenPriceOCRFilterer{contract: contract}}, nil
}

func NewTokenPriceOCRCaller(address common.Address, caller bind.ContractCaller) (*TokenPriceOCRCaller, error) {
	contract, err := bindTokenPriceOCR(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPriceOCRCaller{contract: contract}, nil
}

func NewTokenPriceOCRTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenPriceOCRTransactor, error) {
	contract, err := bindTokenPriceOCR(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPriceOCRTransactor{contract: contract}, nil
}

func NewTokenPriceOCRFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenPriceOCRFilterer, error) {
	contract, err := bindTokenPriceOCR(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenPriceOCRFilterer{contract: contract}, nil
}

func bindTokenPriceOCR(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenPriceOCRMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_TokenPriceOCR *TokenPriceOCRRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPriceOCR.Contract.TokenPriceOCRCaller.contract.Call(opts, result, method, params...)
}

func (_TokenPriceOCR *TokenPriceOCRRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.TokenPriceOCRTransactor.contract.Transfer(opts)
}

func (_TokenPriceOCR *TokenPriceOCRRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.TokenPriceOCRTransactor.contract.Transact(opts, method, params...)
}

func (_TokenPriceOCR *TokenPriceOCRCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPriceOCR.Contract.contract.Call(opts, result, method, params...)
}

func (_TokenPriceOCR *TokenPriceOCRTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.contract.Transfer(opts)
}

func (_TokenPriceOCR *TokenPriceOCRTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.contract.Transact(opts, method, params...)
}

func (_TokenPriceOCR *TokenPriceOCRCaller) ExposeForEncoding(opts *bind.CallOpts, report TokenPriceOCRReport) ([]byte, error) {
	var out []interface{}
	err := _TokenPriceOCR.contract.Call(opts, &out, "exposeForEncoding", report)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_TokenPriceOCR *TokenPriceOCRSession) ExposeForEncoding(report TokenPriceOCRReport) ([]byte, error) {
	return _TokenPriceOCR.Contract.ExposeForEncoding(&_TokenPriceOCR.CallOpts, report)
}

func (_TokenPriceOCR *TokenPriceOCRCallerSession) ExposeForEncoding(report TokenPriceOCRReport) ([]byte, error) {
	return _TokenPriceOCR.Contract.ExposeForEncoding(&_TokenPriceOCR.CallOpts, report)
}

func (_TokenPriceOCR *TokenPriceOCRCaller) GetGasPriceUpdate(opts *bind.CallOpts, destChainSelector uint64) (GetGasPriceUpdate,

	error) {
	var out []interface{}
	err := _TokenPriceOCR.contract.Call(opts, &out, "getGasPriceUpdate", destChainSelector)

	outstruct := new(GetGasPriceUpdate)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_TokenPriceOCR *TokenPriceOCRSession) GetGasPriceUpdate(destChainSelector uint64) (GetGasPriceUpdate,

	error) {
	return _TokenPriceOCR.Contract.GetGasPriceUpdate(&_TokenPriceOCR.CallOpts, destChainSelector)
}

func (_TokenPriceOCR *TokenPriceOCRCallerSession) GetGasPriceUpdate(destChainSelector uint64) (GetGasPriceUpdate,

	error) {
	return _TokenPriceOCR.Contract.GetGasPriceUpdate(&_TokenPriceOCR.CallOpts, destChainSelector)
}

func (_TokenPriceOCR *TokenPriceOCRCaller) GetTokenPriceUpdate(opts *bind.CallOpts, token common.Address) (GetTokenPriceUpdate,

	error) {
	var out []interface{}
	err := _TokenPriceOCR.contract.Call(opts, &out, "getTokenPriceUpdate", token)

	outstruct := new(GetTokenPriceUpdate)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_TokenPriceOCR *TokenPriceOCRSession) GetTokenPriceUpdate(token common.Address) (GetTokenPriceUpdate,

	error) {
	return _TokenPriceOCR.Contract.GetTokenPriceUpdate(&_TokenPriceOCR.CallOpts, token)
}

func (_TokenPriceOCR *TokenPriceOCRCallerSession) GetTokenPriceUpdate(token common.Address) (GetTokenPriceUpdate,

	error) {
	return _TokenPriceOCR.Contract.GetTokenPriceUpdate(&_TokenPriceOCR.CallOpts, token)
}

func (_TokenPriceOCR *TokenPriceOCRCaller) GetTransmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenPriceOCR.contract.Call(opts, &out, "getTransmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenPriceOCR *TokenPriceOCRSession) GetTransmitters() ([]common.Address, error) {
	return _TokenPriceOCR.Contract.GetTransmitters(&_TokenPriceOCR.CallOpts)
}

func (_TokenPriceOCR *TokenPriceOCRCallerSession) GetTransmitters() ([]common.Address, error) {
	return _TokenPriceOCR.Contract.GetTransmitters(&_TokenPriceOCR.CallOpts)
}

func (_TokenPriceOCR *TokenPriceOCRCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _TokenPriceOCR.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_TokenPriceOCR *TokenPriceOCRSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _TokenPriceOCR.Contract.LatestConfigDetails(&_TokenPriceOCR.CallOpts)
}

func (_TokenPriceOCR *TokenPriceOCRCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _TokenPriceOCR.Contract.LatestConfigDetails(&_TokenPriceOCR.CallOpts)
}

func (_TokenPriceOCR *TokenPriceOCRCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _TokenPriceOCR.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.SequenceNumber = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

func (_TokenPriceOCR *TokenPriceOCRSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _TokenPriceOCR.Contract.LatestConfigDigestAndEpoch(&_TokenPriceOCR.CallOpts)
}

func (_TokenPriceOCR *TokenPriceOCRCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _TokenPriceOCR.Contract.LatestConfigDigestAndEpoch(&_TokenPriceOCR.CallOpts)
}

func (_TokenPriceOCR *TokenPriceOCRCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenPriceOCR.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TokenPriceOCR *TokenPriceOCRSession) Owner() (common.Address, error) {
	return _TokenPriceOCR.Contract.Owner(&_TokenPriceOCR.CallOpts)
}

func (_TokenPriceOCR *TokenPriceOCRCallerSession) Owner() (common.Address, error) {
	return _TokenPriceOCR.Contract.Owner(&_TokenPriceOCR.CallOpts)
}

func (_TokenPriceOCR *TokenPriceOCRCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenPriceOCR.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_TokenPriceOCR *TokenPriceOCRSession) TypeAndVersion() (string, error) {
	return _TokenPriceOCR.Contract.TypeAndVersion(&_TokenPriceOCR.CallOpts)
}

func (_TokenPriceOCR *TokenPriceOCRCallerSession) TypeAndVersion() (string, error) {
	return _TokenPriceOCR.Contract.TypeAndVersion(&_TokenPriceOCR.CallOpts)
}

func (_TokenPriceOCR *TokenPriceOCRTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPriceOCR.contract.Transact(opts, "acceptOwnership")
}

func (_TokenPriceOCR *TokenPriceOCRSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.AcceptOwnership(&_TokenPriceOCR.TransactOpts)
}

func (_TokenPriceOCR *TokenPriceOCRTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.AcceptOwnership(&_TokenPriceOCR.TransactOpts)
}

func (_TokenPriceOCR *TokenPriceOCRTransactor) SetOCR3Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _TokenPriceOCR.contract.Transact(opts, "setOCR3Config", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_TokenPriceOCR *TokenPriceOCRSession) SetOCR3Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.SetOCR3Config(&_TokenPriceOCR.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_TokenPriceOCR *TokenPriceOCRTransactorSession) SetOCR3Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.SetOCR3Config(&_TokenPriceOCR.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_TokenPriceOCR *TokenPriceOCRTransactor) SetPriceRegistry(opts *bind.TransactOpts, priceRegistry common.Address) (*types.Transaction, error) {
	return _TokenPriceOCR.contract.Transact(opts, "setPriceRegistry", priceRegistry)
}

func (_TokenPriceOCR *TokenPriceOCRSession) SetPriceRegistry(priceRegistry common.Address) (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.SetPriceRegistry(&_TokenPriceOCR.TransactOpts, priceRegistry)
}

func (_TokenPriceOCR *TokenPriceOCRTransactorSession) SetPriceRegistry(priceRegistry common.Address) (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.SetPriceRegistry(&_TokenPriceOCR.TransactOpts, priceRegistry)
}

func (_TokenPriceOCR *TokenPriceOCRTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _TokenPriceOCR.contract.Transact(opts, "transferOwnership", to)
}

func (_TokenPriceOCR *TokenPriceOCRSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.TransferOwnership(&_TokenPriceOCR.TransactOpts, to)
}

func (_TokenPriceOCR *TokenPriceOCRTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.TransferOwnership(&_TokenPriceOCR.TransactOpts, to)
}

func (_TokenPriceOCR *TokenPriceOCRTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _TokenPriceOCR.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_TokenPriceOCR *TokenPriceOCRSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.Transmit(&_TokenPriceOCR.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_TokenPriceOCR *TokenPriceOCRTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.Transmit(&_TokenPriceOCR.TransactOpts, reportContext, report, rs, ss, rawVs)
}

type TokenPriceOCRConfigSetIterator struct {
	Event *TokenPriceOCRConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPriceOCRConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPriceOCRConfigSet)
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
		it.Event = new(TokenPriceOCRConfigSet)
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

func (it *TokenPriceOCRConfigSetIterator) Error() error {
	return it.fail
}

func (it *TokenPriceOCRConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPriceOCRConfigSet struct {
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

func (_TokenPriceOCR *TokenPriceOCRFilterer) FilterConfigSet(opts *bind.FilterOpts) (*TokenPriceOCRConfigSetIterator, error) {

	logs, sub, err := _TokenPriceOCR.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &TokenPriceOCRConfigSetIterator{contract: _TokenPriceOCR.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_TokenPriceOCR *TokenPriceOCRFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *TokenPriceOCRConfigSet) (event.Subscription, error) {

	logs, sub, err := _TokenPriceOCR.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPriceOCRConfigSet)
				if err := _TokenPriceOCR.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_TokenPriceOCR *TokenPriceOCRFilterer) ParseConfigSet(log types.Log) (*TokenPriceOCRConfigSet, error) {
	event := new(TokenPriceOCRConfigSet)
	if err := _TokenPriceOCR.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPriceOCROwnershipTransferRequestedIterator struct {
	Event *TokenPriceOCROwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPriceOCROwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPriceOCROwnershipTransferRequested)
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
		it.Event = new(TokenPriceOCROwnershipTransferRequested)
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

func (it *TokenPriceOCROwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *TokenPriceOCROwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPriceOCROwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_TokenPriceOCR *TokenPriceOCRFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenPriceOCROwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenPriceOCR.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenPriceOCROwnershipTransferRequestedIterator{contract: _TokenPriceOCR.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_TokenPriceOCR *TokenPriceOCRFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *TokenPriceOCROwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenPriceOCR.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPriceOCROwnershipTransferRequested)
				if err := _TokenPriceOCR.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_TokenPriceOCR *TokenPriceOCRFilterer) ParseOwnershipTransferRequested(log types.Log) (*TokenPriceOCROwnershipTransferRequested, error) {
	event := new(TokenPriceOCROwnershipTransferRequested)
	if err := _TokenPriceOCR.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPriceOCROwnershipTransferredIterator struct {
	Event *TokenPriceOCROwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPriceOCROwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPriceOCROwnershipTransferred)
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
		it.Event = new(TokenPriceOCROwnershipTransferred)
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

func (it *TokenPriceOCROwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *TokenPriceOCROwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPriceOCROwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_TokenPriceOCR *TokenPriceOCRFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenPriceOCROwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenPriceOCR.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenPriceOCROwnershipTransferredIterator{contract: _TokenPriceOCR.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_TokenPriceOCR *TokenPriceOCRFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenPriceOCROwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenPriceOCR.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPriceOCROwnershipTransferred)
				if err := _TokenPriceOCR.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_TokenPriceOCR *TokenPriceOCRFilterer) ParseOwnershipTransferred(log types.Log) (*TokenPriceOCROwnershipTransferred, error) {
	event := new(TokenPriceOCROwnershipTransferred)
	if err := _TokenPriceOCR.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPriceOCRTransmittedIterator struct {
	Event *TokenPriceOCRTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPriceOCRTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPriceOCRTransmitted)
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
		it.Event = new(TokenPriceOCRTransmitted)
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

func (it *TokenPriceOCRTransmittedIterator) Error() error {
	return it.fail
}

func (it *TokenPriceOCRTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPriceOCRTransmitted struct {
	ConfigDigest   [32]byte
	SequenceNumber uint64
	Raw            types.Log
}

func (_TokenPriceOCR *TokenPriceOCRFilterer) FilterTransmitted(opts *bind.FilterOpts) (*TokenPriceOCRTransmittedIterator, error) {

	logs, sub, err := _TokenPriceOCR.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &TokenPriceOCRTransmittedIterator{contract: _TokenPriceOCR.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_TokenPriceOCR *TokenPriceOCRFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *TokenPriceOCRTransmitted) (event.Subscription, error) {

	logs, sub, err := _TokenPriceOCR.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPriceOCRTransmitted)
				if err := _TokenPriceOCR.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_TokenPriceOCR *TokenPriceOCRFilterer) ParseTransmitted(log types.Log) (*TokenPriceOCRTransmitted, error) {
	event := new(TokenPriceOCRTransmitted)
	if err := _TokenPriceOCR.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetGasPriceUpdate struct {
	Value     *big.Int
	Timestamp uint32
}
type GetTokenPriceUpdate struct {
	Value     *big.Int
	Timestamp uint32
}
type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}
type LatestConfigDigestAndEpoch struct {
	ScanLogs       bool
	ConfigDigest   [32]byte
	SequenceNumber uint64
}

func (_TokenPriceOCR *TokenPriceOCR) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _TokenPriceOCR.abi.Events["ConfigSet"].ID:
		return _TokenPriceOCR.ParseConfigSet(log)
	case _TokenPriceOCR.abi.Events["OwnershipTransferRequested"].ID:
		return _TokenPriceOCR.ParseOwnershipTransferRequested(log)
	case _TokenPriceOCR.abi.Events["OwnershipTransferred"].ID:
		return _TokenPriceOCR.ParseOwnershipTransferred(log)
	case _TokenPriceOCR.abi.Events["Transmitted"].ID:
		return _TokenPriceOCR.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (TokenPriceOCRConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (TokenPriceOCROwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (TokenPriceOCROwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (TokenPriceOCRTransmitted) Topic() common.Hash {
	return common.HexToHash("0xe893c2681d327421d89e1cb54fbe64645b4dcea668d6826130b62cf4c6eefea2")
}

func (_TokenPriceOCR *TokenPriceOCR) Address() common.Address {
	return _TokenPriceOCR.address
}

type TokenPriceOCRInterface interface {
	ExposeForEncoding(opts *bind.CallOpts, report TokenPriceOCRReport) ([]byte, error)

	GetGasPriceUpdate(opts *bind.CallOpts, destChainSelector uint64) (GetGasPriceUpdate,

		error)

	GetTokenPriceUpdate(opts *bind.CallOpts, token common.Address) (GetTokenPriceUpdate,

		error)

	GetTransmitters(opts *bind.CallOpts) ([]common.Address, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetOCR3Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	SetPriceRegistry(opts *bind.TransactOpts, priceRegistry common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*TokenPriceOCRConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *TokenPriceOCRConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*TokenPriceOCRConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenPriceOCROwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *TokenPriceOCROwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*TokenPriceOCROwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenPriceOCROwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenPriceOCROwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*TokenPriceOCROwnershipTransferred, error)

	FilterTransmitted(opts *bind.FilterOpts) (*TokenPriceOCRTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *TokenPriceOCRTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*TokenPriceOCRTransmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
