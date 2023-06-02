// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_erc677

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

var BurnMintERC677MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"supplyAfterMint\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotBurner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotMinter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"BurnAccessGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"BurnAccessRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MintAccessGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MintAccessRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBurners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"grantBurnRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnAndMinter\",\"type\":\"address\"}],\"name\":\"grantMintAndBurnRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"grantMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"isBurner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"revokeBurnRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"revokeMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200212b3803806200212b833981016040819052620000349162000277565b338060008686818160036200004a838262000391565b50600462000059828262000391565b5050506001600160a01b0384169150620000bc90505760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600580546001600160a01b0319166001600160a01b0384811691909117909155811615620000ef57620000ef8162000106565b50505060ff90911660805260a052506200045d9050565b336001600160a01b03821603620001605760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000b3565b600680546001600160a01b0319166001600160a01b03838116918217909255600554604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001da57600080fd5b81516001600160401b0380821115620001f757620001f7620001b2565b604051601f8301601f19908116603f01168101908282118183101715620002225762000222620001b2565b816040528381526020925086838588010111156200023f57600080fd5b600091505b8382101562000263578582018301518183018401529082019062000244565b600093810190920192909252949350505050565b600080600080608085870312156200028e57600080fd5b84516001600160401b0380821115620002a657600080fd5b620002b488838901620001c8565b95506020870151915080821115620002cb57600080fd5b50620002da87828801620001c8565b935050604085015160ff81168114620002f257600080fd5b6060959095015193969295505050565b600181811c908216806200031757607f821691505b6020821081036200033857634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200038c57600081815260208120601f850160051c81016020861015620003675750805b601f850160051c820191505b81811015620003885782815560010162000373565b5050505b505050565b81516001600160401b03811115620003ad57620003ad620001b2565b620003c581620003be845462000302565b846200033e565b602080601f831160018114620003fd5760008415620003e45750858301515b600019600386901b1c1916600185901b17855562000388565b600085815260208120601f198616915b828110156200042e578886015182559484019460019091019084016200040d565b50858210156200044d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a051611c9a620004916000396000818161040f0152818161075d01526107870152600061024c0152611c9a6000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c806379cc679011610104578063c2e3273d116100a2578063d73dd62311610071578063d73dd62314610433578063dd62ed3e14610446578063f2fde38b1461048c578063f81094f31461049f57600080fd5b8063c2e3273d146103d4578063c630948d146103e7578063c64d0ebc146103fa578063d5abeb011461040d57600080fd5b806395d89b41116100de57806395d89b4114610393578063a457c2d71461039b578063a9059cbb146103ae578063aa271e1a146103c157600080fd5b806379cc67901461035057806386fe8b43146103635780638da5cb5b1461036b57600080fd5b806340c10f191161017c578063661884631161014b57806366188463146102ea5780636b32810b146102fd57806370a082311461031257806379ba50971461034857600080fd5b806340c10f191461029c57806342966c68146102b15780634334614a146102c45780634f5632f8146102d757600080fd5b806323b872dd116101b857806323b872dd14610232578063313ce5671461024557806339509351146102765780634000aea01461028957600080fd5b806306fdde03146101df578063095ea7b3146101fd57806318160ddd14610220575b600080fd5b6101e76104b2565b6040516101f491906118ad565b60405180910390f35b61021061020b3660046118e9565b610544565b60405190151581526020016101f4565b6002545b6040519081526020016101f4565b610210610240366004611913565b61055e565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101f4565b6102106102843660046118e9565b610582565b61021061029736600461197e565b6105ce565b6102af6102aa3660046118e9565b6106f2565b005b6102af6102bf366004611a67565b610819565b6102106102d2366004611a80565b610866565b6102af6102e5366004611a80565b610873565b6102106102f83660046118e9565b6108cf565b6103056108e2565b6040516101f49190611a9b565b610224610320366004611a80565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6102af6108f3565b6102af61035e3660046118e9565b6109f4565b610305610a43565b60055460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101f4565b6101e7610a4f565b6102106103a93660046118e9565b610a5e565b6102106103bc3660046118e9565b610b2f565b6102106103cf366004611a80565b610b3d565b6102af6103e2366004611a80565b610b4a565b6102af6103f5366004611a80565b610ba6565b6102af610408366004611a80565b610bb4565b7f0000000000000000000000000000000000000000000000000000000000000000610224565b6102af6104413660046118e9565b610c10565b610224610454366004611af5565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6102af61049a366004611a80565b610c1a565b6102af6104ad366004611a80565b610c2b565b6060600380546104c190611b28565b80601f01602080910402602001604051908101604052809291908181526020018280546104ed90611b28565b801561053a5780601f1061050f5761010080835404028352916020019161053a565b820191906000526020600020905b81548152906001019060200180831161051d57829003601f168201915b5050505050905090565b600033610552818585610c87565b60019150505b92915050565b60003361056c858285610cbb565b610577858585610d8c565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061055290829086906105c9908790611baa565b610c87565b60006105da8484610b2f565b508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16858560405161063a929190611bbd565b60405180910390a373ffffffffffffffffffffffffffffffffffffffff84163b156106e8576040517fa4c0ed3600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85169063a4c0ed36906106b590339087908790600401611bde565b600060405180830381600087803b1580156106cf57600080fd5b505af11580156106e3573d6000803e3d6000fd5b505050505b5060019392505050565b6106fb33610b3d565b610738576040517fe2c8c9d50000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b813073ffffffffffffffffffffffffffffffffffffffff82160361075b57600080fd5b7f0000000000000000000000000000000000000000000000000000000000000000158015906107bc57507f0000000000000000000000000000000000000000000000000000000000000000826107b060025490565b6107ba9190611baa565b115b1561080a57816107cb60025490565b6107d59190611baa565b6040517fcbbf111300000000000000000000000000000000000000000000000000000000815260040161072f91815260200190565b6108148383610dba565b505050565b61082233610866565b61085a576040517fc820b10b00000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b61086381610ead565b50565b6000610558600983610eb7565b61087b610ee6565b610886600982610f69565b156108635760405173ffffffffffffffffffffffffffffffffffffffff8216907f0a675452746933cefe3d74182e78db7afe57ba60eaa4234b5d85e9aa41b0610c90600090a250565b60006108db8383610a5e565b9392505050565b60606108ee6007610f8b565b905090565b60065473ffffffffffffffffffffffffffffffffffffffff163314610974576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161072f565b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560068054909116905560405173ffffffffffffffffffffffffffffffffffffffff909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b6109fd33610866565b610a35576040517fc820b10b00000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b610a3f8282610f98565b5050565b60606108ee6009610f8b565b6060600480546104c190611b28565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610b22576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f000000000000000000000000000000000000000000000000000000606482015260840161072f565b6105778286868403610c87565b600033610552818585610d8c565b6000610558600783610eb7565b610b52610ee6565b610b5d600782610fad565b156108635760405173ffffffffffffffffffffffffffffffffffffffff8216907fe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea90600090a250565b610baf81610b4a565b610863815b610bbc610ee6565b610bc7600982610fad565b156108635760405173ffffffffffffffffffffffffffffffffffffffff8216907f92308bb7573b2a3d17ddb868b39d8ebec433f3194421abc22d084f89658c9bad90600090a250565b6108148282610582565b610c22610ee6565b61086381610fcf565b610c33610ee6565b610c3e600782610f69565b156108635760405173ffffffffffffffffffffffffffffffffffffffff8216907fed998b960f6340d045f620c119730f7aa7995e7425c2401d3a5b64ff998a59e990600090a250565b813073ffffffffffffffffffffffffffffffffffffffff821603610caa57600080fd5b610cb58484846110c5565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610cb55781811015610d7f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161072f565b610cb58484848403610c87565b813073ffffffffffffffffffffffffffffffffffffffff821603610daf57600080fd5b610cb5848484611278565b73ffffffffffffffffffffffffffffffffffffffff8216610e37576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161072f565b8060026000828254610e499190611baa565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b61086333826114e7565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018301602052604081205415156108db565b60055473ffffffffffffffffffffffffffffffffffffffff163314610f67576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161072f565b565b60006108db8373ffffffffffffffffffffffffffffffffffffffff84166116ab565b606060006108db8361179e565b610fa3823383610cbb565b610a3f82826114e7565b60006108db8373ffffffffffffffffffffffffffffffffffffffff84166117fa565b3373ffffffffffffffffffffffffffffffffffffffff82160361104e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161072f565b600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600554604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b73ffffffffffffffffffffffffffffffffffffffff8316611167576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f7265737300000000000000000000000000000000000000000000000000000000606482015260840161072f565b73ffffffffffffffffffffffffffffffffffffffff821661120a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f7373000000000000000000000000000000000000000000000000000000000000606482015260840161072f565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff831661131b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015260840161072f565b73ffffffffffffffffffffffffffffffffffffffff82166113be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015260840161072f565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015611474576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e63650000000000000000000000000000000000000000000000000000606482015260840161072f565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610cb5565b73ffffffffffffffffffffffffffffffffffffffff821661158a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161072f565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015611640576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f6365000000000000000000000000000000000000000000000000000000000000606482015260840161072f565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b600081815260018301602052604081205480156117945760006116cf600183611c1c565b85549091506000906116e390600190611c1c565b905081811461174857600086600001828154811061170357611703611c2f565b906000526020600020015490508087600001848154811061172657611726611c2f565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061175957611759611c5e565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610558565b6000915050610558565b6060816000018054806020026020016040519081016040528092919081815260200182805480156117ee57602002820191906000526020600020905b8154815260200190600101908083116117da575b50505050509050919050565b600081815260018301602052604081205461184157508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610558565b506000610558565b6000815180845260005b8181101561186f57602081850181015186830182015201611853565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006108db6020830184611849565b803573ffffffffffffffffffffffffffffffffffffffff811681146118e457600080fd5b919050565b600080604083850312156118fc57600080fd5b611905836118c0565b946020939093013593505050565b60008060006060848603121561192857600080fd5b611931846118c0565b925061193f602085016118c0565b9150604084013590509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060006060848603121561199357600080fd5b61199c846118c0565b925060208401359150604084013567ffffffffffffffff808211156119c057600080fd5b818601915086601f8301126119d457600080fd5b8135818111156119e6576119e661194f565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611a2c57611a2c61194f565b81604052828152896020848701011115611a4557600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b600060208284031215611a7957600080fd5b5035919050565b600060208284031215611a9257600080fd5b6108db826118c0565b6020808252825182820181905260009190848201906040850190845b81811015611ae957835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101611ab7565b50909695505050505050565b60008060408385031215611b0857600080fd5b611b11836118c0565b9150611b1f602084016118c0565b90509250929050565b600181811c90821680611b3c57607f821691505b602082108103611b75577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561055857610558611b7b565b828152604060208201526000611bd66040830184611849565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201526000611c136060830184611849565b95945050505050565b8181038181111561055857610558611b7b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000813000a",
}

var BurnMintERC677ABI = BurnMintERC677MetaData.ABI

var BurnMintERC677Bin = BurnMintERC677MetaData.Bin

func DeployBurnMintERC677(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals_ uint8, maxSupply_ *big.Int) (common.Address, *types.Transaction, *BurnMintERC677, error) {
	parsed, err := BurnMintERC677MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintERC677Bin), backend, name, symbol, decimals_, maxSupply_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintERC677{BurnMintERC677Caller: BurnMintERC677Caller{contract: contract}, BurnMintERC677Transactor: BurnMintERC677Transactor{contract: contract}, BurnMintERC677Filterer: BurnMintERC677Filterer{contract: contract}}, nil
}

type BurnMintERC677 struct {
	address common.Address
	abi     abi.ABI
	BurnMintERC677Caller
	BurnMintERC677Transactor
	BurnMintERC677Filterer
}

type BurnMintERC677Caller struct {
	contract *bind.BoundContract
}

type BurnMintERC677Transactor struct {
	contract *bind.BoundContract
}

type BurnMintERC677Filterer struct {
	contract *bind.BoundContract
}

type BurnMintERC677Session struct {
	Contract     *BurnMintERC677
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintERC677CallerSession struct {
	Contract *BurnMintERC677Caller
	CallOpts bind.CallOpts
}

type BurnMintERC677TransactorSession struct {
	Contract     *BurnMintERC677Transactor
	TransactOpts bind.TransactOpts
}

type BurnMintERC677Raw struct {
	Contract *BurnMintERC677
}

type BurnMintERC677CallerRaw struct {
	Contract *BurnMintERC677Caller
}

type BurnMintERC677TransactorRaw struct {
	Contract *BurnMintERC677Transactor
}

func NewBurnMintERC677(address common.Address, backend bind.ContractBackend) (*BurnMintERC677, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintERC677ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintERC677(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677{address: address, abi: abi, BurnMintERC677Caller: BurnMintERC677Caller{contract: contract}, BurnMintERC677Transactor: BurnMintERC677Transactor{contract: contract}, BurnMintERC677Filterer: BurnMintERC677Filterer{contract: contract}}, nil
}

func NewBurnMintERC677Caller(address common.Address, caller bind.ContractCaller) (*BurnMintERC677Caller, error) {
	contract, err := bindBurnMintERC677(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677Caller{contract: contract}, nil
}

func NewBurnMintERC677Transactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintERC677Transactor, error) {
	contract, err := bindBurnMintERC677(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677Transactor{contract: contract}, nil
}

func NewBurnMintERC677Filterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintERC677Filterer, error) {
	contract, err := bindBurnMintERC677(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677Filterer{contract: contract}, nil
}

func bindBurnMintERC677(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintERC677MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintERC677 *BurnMintERC677Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC677.Contract.BurnMintERC677Caller.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC677 *BurnMintERC677Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.BurnMintERC677Transactor.contract.Transfer(opts)
}

func (_BurnMintERC677 *BurnMintERC677Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.BurnMintERC677Transactor.contract.Transact(opts, method, params...)
}

func (_BurnMintERC677 *BurnMintERC677CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC677.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC677 *BurnMintERC677TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.contract.Transfer(opts)
}

func (_BurnMintERC677 *BurnMintERC677TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintERC677 *BurnMintERC677Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC677.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC677 *BurnMintERC677Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC677.Contract.Allowance(&_BurnMintERC677.CallOpts, owner, spender)
}

func (_BurnMintERC677 *BurnMintERC677CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC677.Contract.Allowance(&_BurnMintERC677.CallOpts, owner, spender)
}

func (_BurnMintERC677 *BurnMintERC677Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC677.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC677 *BurnMintERC677Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC677.Contract.BalanceOf(&_BurnMintERC677.CallOpts, account)
}

func (_BurnMintERC677 *BurnMintERC677CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC677.Contract.BalanceOf(&_BurnMintERC677.CallOpts, account)
}

func (_BurnMintERC677 *BurnMintERC677Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintERC677.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintERC677 *BurnMintERC677Session) Decimals() (uint8, error) {
	return _BurnMintERC677.Contract.Decimals(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677CallerSession) Decimals() (uint8, error) {
	return _BurnMintERC677.Contract.Decimals(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677Caller) GetBurners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintERC677.contract.Call(opts, &out, "getBurners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintERC677 *BurnMintERC677Session) GetBurners() ([]common.Address, error) {
	return _BurnMintERC677.Contract.GetBurners(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677CallerSession) GetBurners() ([]common.Address, error) {
	return _BurnMintERC677.Contract.GetBurners(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677Caller) GetMinters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintERC677.contract.Call(opts, &out, "getMinters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintERC677 *BurnMintERC677Session) GetMinters() ([]common.Address, error) {
	return _BurnMintERC677.Contract.GetMinters(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677CallerSession) GetMinters() ([]common.Address, error) {
	return _BurnMintERC677.Contract.GetMinters(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677Caller) IsBurner(opts *bind.CallOpts, burner common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC677.contract.Call(opts, &out, "isBurner", burner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC677 *BurnMintERC677Session) IsBurner(burner common.Address) (bool, error) {
	return _BurnMintERC677.Contract.IsBurner(&_BurnMintERC677.CallOpts, burner)
}

func (_BurnMintERC677 *BurnMintERC677CallerSession) IsBurner(burner common.Address) (bool, error) {
	return _BurnMintERC677.Contract.IsBurner(&_BurnMintERC677.CallOpts, burner)
}

func (_BurnMintERC677 *BurnMintERC677Caller) IsMinter(opts *bind.CallOpts, minter common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC677.contract.Call(opts, &out, "isMinter", minter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC677 *BurnMintERC677Session) IsMinter(minter common.Address) (bool, error) {
	return _BurnMintERC677.Contract.IsMinter(&_BurnMintERC677.CallOpts, minter)
}

func (_BurnMintERC677 *BurnMintERC677CallerSession) IsMinter(minter common.Address) (bool, error) {
	return _BurnMintERC677.Contract.IsMinter(&_BurnMintERC677.CallOpts, minter)
}

func (_BurnMintERC677 *BurnMintERC677Caller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC677.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC677 *BurnMintERC677Session) MaxSupply() (*big.Int, error) {
	return _BurnMintERC677.Contract.MaxSupply(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677CallerSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC677.Contract.MaxSupply(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC677.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC677 *BurnMintERC677Session) Name() (string, error) {
	return _BurnMintERC677.Contract.Name(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677CallerSession) Name() (string, error) {
	return _BurnMintERC677.Contract.Name(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC677.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC677 *BurnMintERC677Session) Owner() (common.Address, error) {
	return _BurnMintERC677.Contract.Owner(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677CallerSession) Owner() (common.Address, error) {
	return _BurnMintERC677.Contract.Owner(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC677.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC677 *BurnMintERC677Session) Symbol() (string, error) {
	return _BurnMintERC677.Contract.Symbol(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677CallerSession) Symbol() (string, error) {
	return _BurnMintERC677.Contract.Symbol(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC677.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC677 *BurnMintERC677Session) TotalSupply() (*big.Int, error) {
	return _BurnMintERC677.Contract.TotalSupply(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677CallerSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC677.Contract.TotalSupply(&_BurnMintERC677.CallOpts)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "acceptOwnership")
}

func (_BurnMintERC677 *BurnMintERC677Session) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintERC677.Contract.AcceptOwnership(&_BurnMintERC677.TransactOpts)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintERC677.Contract.AcceptOwnership(&_BurnMintERC677.TransactOpts)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "approve", spender, amount)
}

func (_BurnMintERC677 *BurnMintERC677Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.Approve(&_BurnMintERC677.TransactOpts, spender, amount)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.Approve(&_BurnMintERC677.TransactOpts, spender, amount)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "burn", amount)
}

func (_BurnMintERC677 *BurnMintERC677Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.Burn(&_BurnMintERC677.TransactOpts, amount)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.Burn(&_BurnMintERC677.TransactOpts, amount)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "burnFrom", account, amount)
}

func (_BurnMintERC677 *BurnMintERC677Session) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.BurnFrom(&_BurnMintERC677.TransactOpts, account, amount)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.BurnFrom(&_BurnMintERC677.TransactOpts, account, amount)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

func (_BurnMintERC677 *BurnMintERC677Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.DecreaseAllowance(&_BurnMintERC677.TransactOpts, spender, subtractedValue)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.DecreaseAllowance(&_BurnMintERC677.TransactOpts, spender, subtractedValue)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "decreaseApproval", spender, subtractedValue)
}

func (_BurnMintERC677 *BurnMintERC677Session) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.DecreaseApproval(&_BurnMintERC677.TransactOpts, spender, subtractedValue)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.DecreaseApproval(&_BurnMintERC677.TransactOpts, spender, subtractedValue)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) GrantBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "grantBurnRole", burner)
}

func (_BurnMintERC677 *BurnMintERC677Session) GrantBurnRole(burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.GrantBurnRole(&_BurnMintERC677.TransactOpts, burner)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) GrantBurnRole(burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.GrantBurnRole(&_BurnMintERC677.TransactOpts, burner)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (_BurnMintERC677 *BurnMintERC677Session) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.GrantMintAndBurnRoles(&_BurnMintERC677.TransactOpts, burnAndMinter)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.GrantMintAndBurnRoles(&_BurnMintERC677.TransactOpts, burnAndMinter)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) GrantMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "grantMintRole", minter)
}

func (_BurnMintERC677 *BurnMintERC677Session) GrantMintRole(minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.GrantMintRole(&_BurnMintERC677.TransactOpts, minter)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) GrantMintRole(minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.GrantMintRole(&_BurnMintERC677.TransactOpts, minter)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

func (_BurnMintERC677 *BurnMintERC677Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.IncreaseAllowance(&_BurnMintERC677.TransactOpts, spender, addedValue)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.IncreaseAllowance(&_BurnMintERC677.TransactOpts, spender, addedValue)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "increaseApproval", spender, addedValue)
}

func (_BurnMintERC677 *BurnMintERC677Session) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.IncreaseApproval(&_BurnMintERC677.TransactOpts, spender, addedValue)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.IncreaseApproval(&_BurnMintERC677.TransactOpts, spender, addedValue)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "mint", account, amount)
}

func (_BurnMintERC677 *BurnMintERC677Session) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.Mint(&_BurnMintERC677.TransactOpts, account, amount)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.Mint(&_BurnMintERC677.TransactOpts, account, amount)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) RevokeBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "revokeBurnRole", burner)
}

func (_BurnMintERC677 *BurnMintERC677Session) RevokeBurnRole(burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.RevokeBurnRole(&_BurnMintERC677.TransactOpts, burner)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) RevokeBurnRole(burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.RevokeBurnRole(&_BurnMintERC677.TransactOpts, burner)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) RevokeMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "revokeMintRole", minter)
}

func (_BurnMintERC677 *BurnMintERC677Session) RevokeMintRole(minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.RevokeMintRole(&_BurnMintERC677.TransactOpts, minter)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) RevokeMintRole(minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.RevokeMintRole(&_BurnMintERC677.TransactOpts, minter)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "transfer", to, amount)
}

func (_BurnMintERC677 *BurnMintERC677Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.Transfer(&_BurnMintERC677.TransactOpts, to, amount)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.Transfer(&_BurnMintERC677.TransactOpts, to, amount)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "transferAndCall", to, amount, data)
}

func (_BurnMintERC677 *BurnMintERC677Session) TransferAndCall(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.TransferAndCall(&_BurnMintERC677.TransactOpts, to, amount, data)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) TransferAndCall(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.TransferAndCall(&_BurnMintERC677.TransactOpts, to, amount, data)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "transferFrom", from, to, amount)
}

func (_BurnMintERC677 *BurnMintERC677Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.TransferFrom(&_BurnMintERC677.TransactOpts, from, to, amount)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.TransferFrom(&_BurnMintERC677.TransactOpts, from, to, amount)
}

func (_BurnMintERC677 *BurnMintERC677Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "transferOwnership", to)
}

func (_BurnMintERC677 *BurnMintERC677Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.TransferOwnership(&_BurnMintERC677.TransactOpts, to)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.TransferOwnership(&_BurnMintERC677.TransactOpts, to)
}

type BurnMintERC677ApprovalIterator struct {
	Event *BurnMintERC677Approval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677ApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677Approval)
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
		it.Event = new(BurnMintERC677Approval)
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

func (it *BurnMintERC677ApprovalIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_BurnMintERC677 *BurnMintERC677Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC677ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC677.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677ApprovalIterator{contract: _BurnMintERC677.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677 *BurnMintERC677Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC677Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC677.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677Approval)
				if err := _BurnMintERC677.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_BurnMintERC677 *BurnMintERC677Filterer) ParseApproval(log types.Log) (*BurnMintERC677Approval, error) {
	event := new(BurnMintERC677Approval)
	if err := _BurnMintERC677.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677BurnAccessGrantedIterator struct {
	Event *BurnMintERC677BurnAccessGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677BurnAccessGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677BurnAccessGranted)
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
		it.Event = new(BurnMintERC677BurnAccessGranted)
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

func (it *BurnMintERC677BurnAccessGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677BurnAccessGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677BurnAccessGranted struct {
	Burner common.Address
	Raw    types.Log
}

func (_BurnMintERC677 *BurnMintERC677Filterer) FilterBurnAccessGranted(opts *bind.FilterOpts, burner []common.Address) (*BurnMintERC677BurnAccessGrantedIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnMintERC677.contract.FilterLogs(opts, "BurnAccessGranted", burnerRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677BurnAccessGrantedIterator{contract: _BurnMintERC677.contract, event: "BurnAccessGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677 *BurnMintERC677Filterer) WatchBurnAccessGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC677BurnAccessGranted, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnMintERC677.contract.WatchLogs(opts, "BurnAccessGranted", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677BurnAccessGranted)
				if err := _BurnMintERC677.contract.UnpackLog(event, "BurnAccessGranted", log); err != nil {
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

func (_BurnMintERC677 *BurnMintERC677Filterer) ParseBurnAccessGranted(log types.Log) (*BurnMintERC677BurnAccessGranted, error) {
	event := new(BurnMintERC677BurnAccessGranted)
	if err := _BurnMintERC677.contract.UnpackLog(event, "BurnAccessGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677BurnAccessRevokedIterator struct {
	Event *BurnMintERC677BurnAccessRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677BurnAccessRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677BurnAccessRevoked)
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
		it.Event = new(BurnMintERC677BurnAccessRevoked)
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

func (it *BurnMintERC677BurnAccessRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677BurnAccessRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677BurnAccessRevoked struct {
	Burner common.Address
	Raw    types.Log
}

func (_BurnMintERC677 *BurnMintERC677Filterer) FilterBurnAccessRevoked(opts *bind.FilterOpts, burner []common.Address) (*BurnMintERC677BurnAccessRevokedIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnMintERC677.contract.FilterLogs(opts, "BurnAccessRevoked", burnerRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677BurnAccessRevokedIterator{contract: _BurnMintERC677.contract, event: "BurnAccessRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677 *BurnMintERC677Filterer) WatchBurnAccessRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC677BurnAccessRevoked, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnMintERC677.contract.WatchLogs(opts, "BurnAccessRevoked", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677BurnAccessRevoked)
				if err := _BurnMintERC677.contract.UnpackLog(event, "BurnAccessRevoked", log); err != nil {
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

func (_BurnMintERC677 *BurnMintERC677Filterer) ParseBurnAccessRevoked(log types.Log) (*BurnMintERC677BurnAccessRevoked, error) {
	event := new(BurnMintERC677BurnAccessRevoked)
	if err := _BurnMintERC677.contract.UnpackLog(event, "BurnAccessRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677MintAccessGrantedIterator struct {
	Event *BurnMintERC677MintAccessGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677MintAccessGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677MintAccessGranted)
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
		it.Event = new(BurnMintERC677MintAccessGranted)
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

func (it *BurnMintERC677MintAccessGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677MintAccessGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677MintAccessGranted struct {
	Minter common.Address
	Raw    types.Log
}

func (_BurnMintERC677 *BurnMintERC677Filterer) FilterMintAccessGranted(opts *bind.FilterOpts, minter []common.Address) (*BurnMintERC677MintAccessGrantedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BurnMintERC677.contract.FilterLogs(opts, "MintAccessGranted", minterRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677MintAccessGrantedIterator{contract: _BurnMintERC677.contract, event: "MintAccessGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677 *BurnMintERC677Filterer) WatchMintAccessGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC677MintAccessGranted, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BurnMintERC677.contract.WatchLogs(opts, "MintAccessGranted", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677MintAccessGranted)
				if err := _BurnMintERC677.contract.UnpackLog(event, "MintAccessGranted", log); err != nil {
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

func (_BurnMintERC677 *BurnMintERC677Filterer) ParseMintAccessGranted(log types.Log) (*BurnMintERC677MintAccessGranted, error) {
	event := new(BurnMintERC677MintAccessGranted)
	if err := _BurnMintERC677.contract.UnpackLog(event, "MintAccessGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677MintAccessRevokedIterator struct {
	Event *BurnMintERC677MintAccessRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677MintAccessRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677MintAccessRevoked)
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
		it.Event = new(BurnMintERC677MintAccessRevoked)
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

func (it *BurnMintERC677MintAccessRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677MintAccessRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677MintAccessRevoked struct {
	Minter common.Address
	Raw    types.Log
}

func (_BurnMintERC677 *BurnMintERC677Filterer) FilterMintAccessRevoked(opts *bind.FilterOpts, minter []common.Address) (*BurnMintERC677MintAccessRevokedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BurnMintERC677.contract.FilterLogs(opts, "MintAccessRevoked", minterRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677MintAccessRevokedIterator{contract: _BurnMintERC677.contract, event: "MintAccessRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677 *BurnMintERC677Filterer) WatchMintAccessRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC677MintAccessRevoked, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BurnMintERC677.contract.WatchLogs(opts, "MintAccessRevoked", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677MintAccessRevoked)
				if err := _BurnMintERC677.contract.UnpackLog(event, "MintAccessRevoked", log); err != nil {
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

func (_BurnMintERC677 *BurnMintERC677Filterer) ParseMintAccessRevoked(log types.Log) (*BurnMintERC677MintAccessRevoked, error) {
	event := new(BurnMintERC677MintAccessRevoked)
	if err := _BurnMintERC677.contract.UnpackLog(event, "MintAccessRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677OwnershipTransferRequestedIterator struct {
	Event *BurnMintERC677OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677OwnershipTransferRequested)
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
		it.Event = new(BurnMintERC677OwnershipTransferRequested)
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

func (it *BurnMintERC677OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintERC677 *BurnMintERC677Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677OwnershipTransferRequestedIterator{contract: _BurnMintERC677.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677 *BurnMintERC677Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintERC677OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677OwnershipTransferRequested)
				if err := _BurnMintERC677.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_BurnMintERC677 *BurnMintERC677Filterer) ParseOwnershipTransferRequested(log types.Log) (*BurnMintERC677OwnershipTransferRequested, error) {
	event := new(BurnMintERC677OwnershipTransferRequested)
	if err := _BurnMintERC677.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677OwnershipTransferredIterator struct {
	Event *BurnMintERC677OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677OwnershipTransferred)
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
		it.Event = new(BurnMintERC677OwnershipTransferred)
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

func (it *BurnMintERC677OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintERC677 *BurnMintERC677Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677OwnershipTransferredIterator{contract: _BurnMintERC677.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677 *BurnMintERC677Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC677OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677OwnershipTransferred)
				if err := _BurnMintERC677.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_BurnMintERC677 *BurnMintERC677Filterer) ParseOwnershipTransferred(log types.Log) (*BurnMintERC677OwnershipTransferred, error) {
	event := new(BurnMintERC677OwnershipTransferred)
	if err := _BurnMintERC677.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677TransferIterator struct {
	Event *BurnMintERC677Transfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677TransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677Transfer)
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
		it.Event = new(BurnMintERC677Transfer)
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

func (it *BurnMintERC677TransferIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_BurnMintERC677 *BurnMintERC677Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677TransferIterator{contract: _BurnMintERC677.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677 *BurnMintERC677Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC677Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677Transfer)
				if err := _BurnMintERC677.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_BurnMintERC677 *BurnMintERC677Filterer) ParseTransfer(log types.Log) (*BurnMintERC677Transfer, error) {
	event := new(BurnMintERC677Transfer)
	if err := _BurnMintERC677.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677Transfer0Iterator struct {
	Event *BurnMintERC677Transfer0

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677Transfer0Iterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677Transfer0)
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
		it.Event = new(BurnMintERC677Transfer0)
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

func (it *BurnMintERC677Transfer0Iterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677Transfer0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677Transfer0 struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log
}

func (_BurnMintERC677 *BurnMintERC677Filterer) FilterTransfer0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677Transfer0Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677.contract.FilterLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677Transfer0Iterator{contract: _BurnMintERC677.contract, event: "Transfer0", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677 *BurnMintERC677Filterer) WatchTransfer0(opts *bind.WatchOpts, sink chan<- *BurnMintERC677Transfer0, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677.contract.WatchLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677Transfer0)
				if err := _BurnMintERC677.contract.UnpackLog(event, "Transfer0", log); err != nil {
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

func (_BurnMintERC677 *BurnMintERC677Filterer) ParseTransfer0(log types.Log) (*BurnMintERC677Transfer0, error) {
	event := new(BurnMintERC677Transfer0)
	if err := _BurnMintERC677.contract.UnpackLog(event, "Transfer0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_BurnMintERC677 *BurnMintERC677) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BurnMintERC677.abi.Events["Approval"].ID:
		return _BurnMintERC677.ParseApproval(log)
	case _BurnMintERC677.abi.Events["BurnAccessGranted"].ID:
		return _BurnMintERC677.ParseBurnAccessGranted(log)
	case _BurnMintERC677.abi.Events["BurnAccessRevoked"].ID:
		return _BurnMintERC677.ParseBurnAccessRevoked(log)
	case _BurnMintERC677.abi.Events["MintAccessGranted"].ID:
		return _BurnMintERC677.ParseMintAccessGranted(log)
	case _BurnMintERC677.abi.Events["MintAccessRevoked"].ID:
		return _BurnMintERC677.ParseMintAccessRevoked(log)
	case _BurnMintERC677.abi.Events["OwnershipTransferRequested"].ID:
		return _BurnMintERC677.ParseOwnershipTransferRequested(log)
	case _BurnMintERC677.abi.Events["OwnershipTransferred"].ID:
		return _BurnMintERC677.ParseOwnershipTransferred(log)
	case _BurnMintERC677.abi.Events["Transfer"].ID:
		return _BurnMintERC677.ParseTransfer(log)
	case _BurnMintERC677.abi.Events["Transfer0"].ID:
		return _BurnMintERC677.ParseTransfer0(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (BurnMintERC677Approval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (BurnMintERC677BurnAccessGranted) Topic() common.Hash {
	return common.HexToHash("0x92308bb7573b2a3d17ddb868b39d8ebec433f3194421abc22d084f89658c9bad")
}

func (BurnMintERC677BurnAccessRevoked) Topic() common.Hash {
	return common.HexToHash("0x0a675452746933cefe3d74182e78db7afe57ba60eaa4234b5d85e9aa41b0610c")
}

func (BurnMintERC677MintAccessGranted) Topic() common.Hash {
	return common.HexToHash("0xe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea")
}

func (BurnMintERC677MintAccessRevoked) Topic() common.Hash {
	return common.HexToHash("0xed998b960f6340d045f620c119730f7aa7995e7425c2401d3a5b64ff998a59e9")
}

func (BurnMintERC677OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BurnMintERC677OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BurnMintERC677Transfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (BurnMintERC677Transfer0) Topic() common.Hash {
	return common.HexToHash("0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16")
}

func (_BurnMintERC677 *BurnMintERC677) Address() common.Address {
	return _BurnMintERC677.address
}

type BurnMintERC677Interface interface {
	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	GetBurners(opts *bind.CallOpts) ([]common.Address, error)

	GetMinters(opts *bind.CallOpts) ([]common.Address, error)

	IsBurner(opts *bind.CallOpts, burner common.Address) (bool, error)

	IsMinter(opts *bind.CallOpts, minter common.Address) (bool, error)

	MaxSupply(opts *bind.CallOpts) (*big.Int, error)

	Name(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)

	Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error)

	DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error)

	GrantBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error)

	GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error)

	GrantMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error)

	IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error)

	IncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	RevokeBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error)

	RevokeMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)

	TransferAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC677ApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC677Approval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*BurnMintERC677Approval, error)

	FilterBurnAccessGranted(opts *bind.FilterOpts, burner []common.Address) (*BurnMintERC677BurnAccessGrantedIterator, error)

	WatchBurnAccessGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC677BurnAccessGranted, burner []common.Address) (event.Subscription, error)

	ParseBurnAccessGranted(log types.Log) (*BurnMintERC677BurnAccessGranted, error)

	FilterBurnAccessRevoked(opts *bind.FilterOpts, burner []common.Address) (*BurnMintERC677BurnAccessRevokedIterator, error)

	WatchBurnAccessRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC677BurnAccessRevoked, burner []common.Address) (event.Subscription, error)

	ParseBurnAccessRevoked(log types.Log) (*BurnMintERC677BurnAccessRevoked, error)

	FilterMintAccessGranted(opts *bind.FilterOpts, minter []common.Address) (*BurnMintERC677MintAccessGrantedIterator, error)

	WatchMintAccessGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC677MintAccessGranted, minter []common.Address) (event.Subscription, error)

	ParseMintAccessGranted(log types.Log) (*BurnMintERC677MintAccessGranted, error)

	FilterMintAccessRevoked(opts *bind.FilterOpts, minter []common.Address) (*BurnMintERC677MintAccessRevokedIterator, error)

	WatchMintAccessRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC677MintAccessRevoked, minter []common.Address) (event.Subscription, error)

	ParseMintAccessRevoked(log types.Log) (*BurnMintERC677MintAccessRevoked, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintERC677OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnMintERC677OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC677OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnMintERC677OwnershipTransferred, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677TransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC677Transfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*BurnMintERC677Transfer, error)

	FilterTransfer0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677Transfer0Iterator, error)

	WatchTransfer0(opts *bind.WatchOpts, sink chan<- *BurnMintERC677Transfer0, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer0(log types.Log) (*BurnMintERC677Transfer0, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
