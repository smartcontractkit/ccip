// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_erc677_helper

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

var BurnMintERC677HelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"supplyAfterMint\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotBurner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotMinter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"BurnAccessGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"BurnAccessRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MintAccessGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MintAccessRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"drip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBurners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"grantBurnRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnAndMinter\",\"type\":\"address\"}],\"name\":\"grantMintAndBurnRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"grantMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"isBurner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"revokeBurnRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"revokeMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620022b7380380620022b783398101604081905262000034916200027e565b8181601260003380828686818160036200004f838262000377565b5060046200005e828262000377565b5050506001600160a01b0384169150620000c190505760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600580546001600160a01b0319166001600160a01b0384811691909117909155811615620000f457620000f4816200010d565b50505060ff90911660805260a052506200044392505050565b336001600160a01b03821603620001675760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000b8565b600680546001600160a01b0319166001600160a01b03838116918217909255600554604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001e157600080fd5b81516001600160401b0380821115620001fe57620001fe620001b9565b604051601f8301601f19908116603f01168101908282118183101715620002295762000229620001b9565b816040528381526020925086838588010111156200024657600080fd5b600091505b838210156200026a57858201830151818301840152908201906200024b565b600093810190920192909252949350505050565b600080604083850312156200029257600080fd5b82516001600160401b0380821115620002aa57600080fd5b620002b886838701620001cf565b93506020850151915080821115620002cf57600080fd5b50620002de85828601620001cf565b9150509250929050565b600181811c90821680620002fd57607f821691505b6020821081036200031e57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200037257600081815260208120601f850160051c810160208610156200034d5750805b601f850160051c820191505b818110156200036e5782815560010162000359565b5050505b505050565b81516001600160401b03811115620003935762000393620001b9565b620003ab81620003a48454620002e8565b8462000324565b602080601f831160018114620003e35760008415620003ca5750858301515b600019600386901b1c1916600185901b1785556200036e565b600085815260208120601f198616915b828110156200041457888601518255948401946001909101908401620003f3565b5085821015620004335787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a051611e406200047760003960008181610475015281816108a501526108cf0152600061028c0152611e406000f3fe608060405234801561001057600080fd5b506004361061020b5760003560e01c806379ba50971161012a578063aa271e1a116100bd578063d5abeb011161008c578063dd62ed3e11610071578063dd62ed3e146104ac578063f2fde38b146104f2578063f81094f31461050557600080fd5b8063d5abeb0114610473578063d73dd6231461049957600080fd5b8063aa271e1a14610427578063c2e3273d1461043a578063c630948d1461044d578063c64d0ebc1461046057600080fd5b806395d89b41116100f957806395d89b41146103e65780639dc29fac146103ee578063a457c2d714610401578063a9059cbb1461041457600080fd5b806379ba50971461039b57806379cc6790146103a357806386fe8b43146103b65780638da5cb5b146103be57600080fd5b806340c10f19116101a25780636618846311610171578063661884631461032a57806367a5cd061461033d5780636b32810b1461035057806370a082311461036557600080fd5b806340c10f19146102dc57806342966c68146102f15780634334614a146103045780634f5632f81461031757600080fd5b806323b872dd116101de57806323b872dd14610272578063313ce5671461028557806339509351146102b65780634000aea0146102c957600080fd5b806301ffc9a71461021057806306fdde0314610238578063095ea7b31461024d57806318160ddd14610260575b600080fd5b61022361021e3660046119ad565b610518565b60405190151581526020015b60405180910390f35b6102406105fd565b60405161022f9190611a53565b61022361025b366004611a8f565b61068f565b6002545b60405190815260200161022f565b610223610280366004611ab9565b6106a7565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016815260200161022f565b6102236102c4366004611a8f565b6106cb565b6102236102d7366004611b24565b610717565b6102ef6102ea366004611a8f565b61083a565b005b6102ef6102ff366004611c0d565b610961565b610223610312366004611c26565b6109ae565b6102ef610325366004611c26565b6109bb565b610223610338366004611a8f565b610a17565b6102ef61034b366004611c26565b610a2a565b610358610a3c565b60405161022f9190611c41565b610264610373366004611c26565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6102ef610a4d565b6102ef6103b1366004611a8f565b610b4e565b610358610b9d565b60055460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161022f565b610240610ba9565b6102ef6103fc366004611a8f565b610bb8565b61022361040f366004611a8f565b610bc2565b610223610422366004611a8f565b610c93565b610223610435366004611c26565b610ca1565b6102ef610448366004611c26565b610cae565b6102ef61045b366004611c26565b610d0a565b6102ef61046e366004611c26565b610d18565b7f0000000000000000000000000000000000000000000000000000000000000000610264565b6102ef6104a7366004611a8f565b610d74565b6102646104ba366004611c9b565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6102ef610500366004611c26565b610d7e565b6102ef610513366004611c26565b610d8f565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f4000aea00000000000000000000000000000000000000000000000000000000014806105ab57507fffffffff0000000000000000000000000000000000000000000000000000000082167fe6599b4d00000000000000000000000000000000000000000000000000000000145b806105f757507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b60606003805461060c90611cce565b80601f016020809104026020016040519081016040528092919081815260200182805461063890611cce565b80156106855780601f1061065a57610100808354040283529160200191610685565b820191906000526020600020905b81548152906001019060200180831161066857829003601f168201915b5050505050905090565b60003361069d818585610deb565b5060019392505050565b6000336106b5858285610e1f565b6106c0858585610ef0565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061069d9082908690610712908790611d50565b610deb565b60006107238484610c93565b508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c168585604051610783929190611d63565b60405180910390a373ffffffffffffffffffffffffffffffffffffffff84163b1561069d576040517fa4c0ed3600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85169063a4c0ed36906107fe90339087908790600401611d84565b600060405180830381600087803b15801561081857600080fd5b505af115801561082c573d6000803e3d6000fd5b505050505060019392505050565b61084333610ca1565b610880576040517fe2c8c9d50000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b813073ffffffffffffffffffffffffffffffffffffffff8216036108a357600080fd5b7f00000000000000000000000000000000000000000000000000000000000000001580159061090457507f0000000000000000000000000000000000000000000000000000000000000000826108f860025490565b6109029190611d50565b115b15610952578161091360025490565b61091d9190611d50565b6040517fcbbf111300000000000000000000000000000000000000000000000000000000815260040161087791815260200190565b61095c8383610f1e565b505050565b61096a336109ae565b6109a2576040517fc820b10b000000000000000000000000000000000000000000000000000000008152336004820152602401610877565b6109ab81611011565b50565b60006105f760098361101b565b6109c361104a565b6109ce6009826110cd565b156109ab5760405173ffffffffffffffffffffffffffffffffffffffff8216907f0a675452746933cefe3d74182e78db7afe57ba60eaa4234b5d85e9aa41b0610c90600090a250565b6000610a238383610bc2565b9392505050565b6109ab81670de0b6b3a7640000610f1e565b6060610a4860076110ef565b905090565b60065473ffffffffffffffffffffffffffffffffffffffff163314610ace576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610877565b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560068054909116905560405173ffffffffffffffffffffffffffffffffffffffff909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b610b57336109ae565b610b8f576040517fc820b10b000000000000000000000000000000000000000000000000000000008152336004820152602401610877565b610b9982826110fc565b5050565b6060610a4860096110ef565b60606004805461060c90611cce565b610b998282610b4e565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610c86576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610877565b6106c08286868403610deb565b60003361069d818585610ef0565b60006105f760078361101b565b610cb661104a565b610cc1600782611111565b156109ab5760405173ffffffffffffffffffffffffffffffffffffffff8216907fe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea90600090a250565b610d1381610cae565b6109ab815b610d2061104a565b610d2b600982611111565b156109ab5760405173ffffffffffffffffffffffffffffffffffffffff8216907f92308bb7573b2a3d17ddb868b39d8ebec433f3194421abc22d084f89658c9bad90600090a250565b61095c82826106cb565b610d8661104a565b6109ab81611133565b610d9761104a565b610da26007826110cd565b156109ab5760405173ffffffffffffffffffffffffffffffffffffffff8216907fed998b960f6340d045f620c119730f7aa7995e7425c2401d3a5b64ff998a59e990600090a250565b813073ffffffffffffffffffffffffffffffffffffffff821603610e0e57600080fd5b610e19848484611229565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610e195781811015610ee3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610877565b610e198484848403610deb565b813073ffffffffffffffffffffffffffffffffffffffff821603610f1357600080fd5b610e198484846113dc565b73ffffffffffffffffffffffffffffffffffffffff8216610f9b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610877565b8060026000828254610fad9190611d50565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6109ab338261164b565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515610a23565b60055473ffffffffffffffffffffffffffffffffffffffff1633146110cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610877565b565b6000610a238373ffffffffffffffffffffffffffffffffffffffff841661180f565b60606000610a2383611902565b611107823383610e1f565b610b99828261164b565b6000610a238373ffffffffffffffffffffffffffffffffffffffff841661195e565b3373ffffffffffffffffffffffffffffffffffffffff8216036111b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610877565b600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600554604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b73ffffffffffffffffffffffffffffffffffffffff83166112cb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610877565b73ffffffffffffffffffffffffffffffffffffffff821661136e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610877565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff831661147f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610877565b73ffffffffffffffffffffffffffffffffffffffff8216611522576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610877565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156115d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610877565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610e19565b73ffffffffffffffffffffffffffffffffffffffff82166116ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610877565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260208190526040902054818110156117a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610877565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b600081815260018301602052604081205480156118f8576000611833600183611dc2565b855490915060009061184790600190611dc2565b90508181146118ac57600086600001828154811061186757611867611dd5565b906000526020600020015490508087600001848154811061188a5761188a611dd5565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806118bd576118bd611e04565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506105f7565b60009150506105f7565b60608160000180548060200260200160405190810160405280929190818152602001828054801561195257602002820191906000526020600020905b81548152602001906001019080831161193e575b50505050509050919050565b60008181526001830160205260408120546119a5575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556105f7565b5060006105f7565b6000602082840312156119bf57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610a2357600080fd5b6000815180845260005b81811015611a15576020818501810151868301820152016119f9565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610a2360208301846119ef565b803573ffffffffffffffffffffffffffffffffffffffff81168114611a8a57600080fd5b919050565b60008060408385031215611aa257600080fd5b611aab83611a66565b946020939093013593505050565b600080600060608486031215611ace57600080fd5b611ad784611a66565b9250611ae560208501611a66565b9150604084013590509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080600060608486031215611b3957600080fd5b611b4284611a66565b925060208401359150604084013567ffffffffffffffff80821115611b6657600080fd5b818601915086601f830112611b7a57600080fd5b813581811115611b8c57611b8c611af5565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611bd257611bd2611af5565b81604052828152896020848701011115611beb57600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b600060208284031215611c1f57600080fd5b5035919050565b600060208284031215611c3857600080fd5b610a2382611a66565b6020808252825182820181905260009190848201906040850190845b81811015611c8f57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101611c5d565b50909695505050505050565b60008060408385031215611cae57600080fd5b611cb783611a66565b9150611cc560208401611a66565b90509250929050565b600181811c90821680611ce257607f821691505b602082108103611d1b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156105f7576105f7611d21565b828152604060208201526000611d7c60408301846119ef565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201526000611db960608301846119ef565b95945050505050565b818103818111156105f7576105f7611d21565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000813000a",
}

var BurnMintERC677HelperABI = BurnMintERC677HelperMetaData.ABI

var BurnMintERC677HelperBin = BurnMintERC677HelperMetaData.Bin

func DeployBurnMintERC677Helper(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *BurnMintERC677Helper, error) {
	parsed, err := BurnMintERC677HelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintERC677HelperBin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintERC677Helper{BurnMintERC677HelperCaller: BurnMintERC677HelperCaller{contract: contract}, BurnMintERC677HelperTransactor: BurnMintERC677HelperTransactor{contract: contract}, BurnMintERC677HelperFilterer: BurnMintERC677HelperFilterer{contract: contract}}, nil
}

type BurnMintERC677Helper struct {
	address common.Address
	abi     abi.ABI
	BurnMintERC677HelperCaller
	BurnMintERC677HelperTransactor
	BurnMintERC677HelperFilterer
}

type BurnMintERC677HelperCaller struct {
	contract *bind.BoundContract
}

type BurnMintERC677HelperTransactor struct {
	contract *bind.BoundContract
}

type BurnMintERC677HelperFilterer struct {
	contract *bind.BoundContract
}

type BurnMintERC677HelperSession struct {
	Contract     *BurnMintERC677Helper
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintERC677HelperCallerSession struct {
	Contract *BurnMintERC677HelperCaller
	CallOpts bind.CallOpts
}

type BurnMintERC677HelperTransactorSession struct {
	Contract     *BurnMintERC677HelperTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintERC677HelperRaw struct {
	Contract *BurnMintERC677Helper
}

type BurnMintERC677HelperCallerRaw struct {
	Contract *BurnMintERC677HelperCaller
}

type BurnMintERC677HelperTransactorRaw struct {
	Contract *BurnMintERC677HelperTransactor
}

func NewBurnMintERC677Helper(address common.Address, backend bind.ContractBackend) (*BurnMintERC677Helper, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintERC677HelperABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintERC677Helper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677Helper{address: address, abi: abi, BurnMintERC677HelperCaller: BurnMintERC677HelperCaller{contract: contract}, BurnMintERC677HelperTransactor: BurnMintERC677HelperTransactor{contract: contract}, BurnMintERC677HelperFilterer: BurnMintERC677HelperFilterer{contract: contract}}, nil
}

func NewBurnMintERC677HelperCaller(address common.Address, caller bind.ContractCaller) (*BurnMintERC677HelperCaller, error) {
	contract, err := bindBurnMintERC677Helper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperCaller{contract: contract}, nil
}

func NewBurnMintERC677HelperTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintERC677HelperTransactor, error) {
	contract, err := bindBurnMintERC677Helper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperTransactor{contract: contract}, nil
}

func NewBurnMintERC677HelperFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintERC677HelperFilterer, error) {
	contract, err := bindBurnMintERC677Helper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperFilterer{contract: contract}, nil
}

func bindBurnMintERC677Helper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintERC677HelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC677Helper.Contract.BurnMintERC677HelperCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.BurnMintERC677HelperTransactor.contract.Transfer(opts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.BurnMintERC677HelperTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC677Helper.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.contract.Transfer(opts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.Allowance(&_BurnMintERC677Helper.CallOpts, owner, spender)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.Allowance(&_BurnMintERC677Helper.CallOpts, owner, spender)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.BalanceOf(&_BurnMintERC677Helper.CallOpts, account)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.BalanceOf(&_BurnMintERC677Helper.CallOpts, account)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Decimals() (uint8, error) {
	return _BurnMintERC677Helper.Contract.Decimals(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) Decimals() (uint8, error) {
	return _BurnMintERC677Helper.Contract.Decimals(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) GetBurners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "getBurners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) GetBurners() ([]common.Address, error) {
	return _BurnMintERC677Helper.Contract.GetBurners(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) GetBurners() ([]common.Address, error) {
	return _BurnMintERC677Helper.Contract.GetBurners(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) GetMinters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "getMinters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) GetMinters() ([]common.Address, error) {
	return _BurnMintERC677Helper.Contract.GetMinters(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) GetMinters() ([]common.Address, error) {
	return _BurnMintERC677Helper.Contract.GetMinters(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) IsBurner(opts *bind.CallOpts, burner common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "isBurner", burner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) IsBurner(burner common.Address) (bool, error) {
	return _BurnMintERC677Helper.Contract.IsBurner(&_BurnMintERC677Helper.CallOpts, burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) IsBurner(burner common.Address) (bool, error) {
	return _BurnMintERC677Helper.Contract.IsBurner(&_BurnMintERC677Helper.CallOpts, burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) IsMinter(opts *bind.CallOpts, minter common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "isMinter", minter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) IsMinter(minter common.Address) (bool, error) {
	return _BurnMintERC677Helper.Contract.IsMinter(&_BurnMintERC677Helper.CallOpts, minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) IsMinter(minter common.Address) (bool, error) {
	return _BurnMintERC677Helper.Contract.IsMinter(&_BurnMintERC677Helper.CallOpts, minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.MaxSupply(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.MaxSupply(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Name() (string, error) {
	return _BurnMintERC677Helper.Contract.Name(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) Name() (string, error) {
	return _BurnMintERC677Helper.Contract.Name(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Owner() (common.Address, error) {
	return _BurnMintERC677Helper.Contract.Owner(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) Owner() (common.Address, error) {
	return _BurnMintERC677Helper.Contract.Owner(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC677Helper.Contract.SupportsInterface(&_BurnMintERC677Helper.CallOpts, interfaceId)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC677Helper.Contract.SupportsInterface(&_BurnMintERC677Helper.CallOpts, interfaceId)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Symbol() (string, error) {
	return _BurnMintERC677Helper.Contract.Symbol(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) Symbol() (string, error) {
	return _BurnMintERC677Helper.Contract.Symbol(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.TotalSupply(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.TotalSupply(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "acceptOwnership")
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.AcceptOwnership(&_BurnMintERC677Helper.TransactOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.AcceptOwnership(&_BurnMintERC677Helper.TransactOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "approve", spender, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Approve(&_BurnMintERC677Helper.TransactOpts, spender, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Approve(&_BurnMintERC677Helper.TransactOpts, spender, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "burn", amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Burn(&_BurnMintERC677Helper.TransactOpts, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Burn(&_BurnMintERC677Helper.TransactOpts, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "burn0", account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Burn0(&_BurnMintERC677Helper.TransactOpts, account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Burn0(&_BurnMintERC677Helper.TransactOpts, account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "burnFrom", account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.BurnFrom(&_BurnMintERC677Helper.TransactOpts, account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.BurnFrom(&_BurnMintERC677Helper.TransactOpts, account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.DecreaseAllowance(&_BurnMintERC677Helper.TransactOpts, spender, subtractedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.DecreaseAllowance(&_BurnMintERC677Helper.TransactOpts, spender, subtractedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "decreaseApproval", spender, subtractedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.DecreaseApproval(&_BurnMintERC677Helper.TransactOpts, spender, subtractedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.DecreaseApproval(&_BurnMintERC677Helper.TransactOpts, spender, subtractedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) Drip(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "drip", to)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Drip(to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Drip(&_BurnMintERC677Helper.TransactOpts, to)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) Drip(to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Drip(&_BurnMintERC677Helper.TransactOpts, to)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) GrantBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "grantBurnRole", burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) GrantBurnRole(burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.GrantBurnRole(&_BurnMintERC677Helper.TransactOpts, burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) GrantBurnRole(burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.GrantBurnRole(&_BurnMintERC677Helper.TransactOpts, burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.GrantMintAndBurnRoles(&_BurnMintERC677Helper.TransactOpts, burnAndMinter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.GrantMintAndBurnRoles(&_BurnMintERC677Helper.TransactOpts, burnAndMinter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) GrantMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "grantMintRole", minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) GrantMintRole(minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.GrantMintRole(&_BurnMintERC677Helper.TransactOpts, minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) GrantMintRole(minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.GrantMintRole(&_BurnMintERC677Helper.TransactOpts, minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.IncreaseAllowance(&_BurnMintERC677Helper.TransactOpts, spender, addedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.IncreaseAllowance(&_BurnMintERC677Helper.TransactOpts, spender, addedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "increaseApproval", spender, addedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.IncreaseApproval(&_BurnMintERC677Helper.TransactOpts, spender, addedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.IncreaseApproval(&_BurnMintERC677Helper.TransactOpts, spender, addedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "mint", account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Mint(&_BurnMintERC677Helper.TransactOpts, account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Mint(&_BurnMintERC677Helper.TransactOpts, account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) RevokeBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "revokeBurnRole", burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) RevokeBurnRole(burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.RevokeBurnRole(&_BurnMintERC677Helper.TransactOpts, burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) RevokeBurnRole(burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.RevokeBurnRole(&_BurnMintERC677Helper.TransactOpts, burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) RevokeMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "revokeMintRole", minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) RevokeMintRole(minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.RevokeMintRole(&_BurnMintERC677Helper.TransactOpts, minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) RevokeMintRole(minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.RevokeMintRole(&_BurnMintERC677Helper.TransactOpts, minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "transfer", to, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Transfer(&_BurnMintERC677Helper.TransactOpts, to, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Transfer(&_BurnMintERC677Helper.TransactOpts, to, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "transferAndCall", to, amount, data)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) TransferAndCall(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.TransferAndCall(&_BurnMintERC677Helper.TransactOpts, to, amount, data)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) TransferAndCall(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.TransferAndCall(&_BurnMintERC677Helper.TransactOpts, to, amount, data)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "transferFrom", from, to, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.TransferFrom(&_BurnMintERC677Helper.TransactOpts, from, to, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.TransferFrom(&_BurnMintERC677Helper.TransactOpts, from, to, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "transferOwnership", to)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.TransferOwnership(&_BurnMintERC677Helper.TransactOpts, to)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.TransferOwnership(&_BurnMintERC677Helper.TransactOpts, to)
}

type BurnMintERC677HelperApprovalIterator struct {
	Event *BurnMintERC677HelperApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperApproval)
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
		it.Event = new(BurnMintERC677HelperApproval)
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

func (it *BurnMintERC677HelperApprovalIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC677HelperApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperApprovalIterator{contract: _BurnMintERC677Helper.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperApproval)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseApproval(log types.Log) (*BurnMintERC677HelperApproval, error) {
	event := new(BurnMintERC677HelperApproval)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperBurnAccessGrantedIterator struct {
	Event *BurnMintERC677HelperBurnAccessGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperBurnAccessGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperBurnAccessGranted)
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
		it.Event = new(BurnMintERC677HelperBurnAccessGranted)
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

func (it *BurnMintERC677HelperBurnAccessGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperBurnAccessGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperBurnAccessGranted struct {
	Burner common.Address
	Raw    types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterBurnAccessGranted(opts *bind.FilterOpts, burner []common.Address) (*BurnMintERC677HelperBurnAccessGrantedIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "BurnAccessGranted", burnerRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperBurnAccessGrantedIterator{contract: _BurnMintERC677Helper.contract, event: "BurnAccessGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchBurnAccessGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperBurnAccessGranted, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "BurnAccessGranted", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperBurnAccessGranted)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "BurnAccessGranted", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseBurnAccessGranted(log types.Log) (*BurnMintERC677HelperBurnAccessGranted, error) {
	event := new(BurnMintERC677HelperBurnAccessGranted)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "BurnAccessGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperBurnAccessRevokedIterator struct {
	Event *BurnMintERC677HelperBurnAccessRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperBurnAccessRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperBurnAccessRevoked)
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
		it.Event = new(BurnMintERC677HelperBurnAccessRevoked)
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

func (it *BurnMintERC677HelperBurnAccessRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperBurnAccessRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperBurnAccessRevoked struct {
	Burner common.Address
	Raw    types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterBurnAccessRevoked(opts *bind.FilterOpts, burner []common.Address) (*BurnMintERC677HelperBurnAccessRevokedIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "BurnAccessRevoked", burnerRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperBurnAccessRevokedIterator{contract: _BurnMintERC677Helper.contract, event: "BurnAccessRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchBurnAccessRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperBurnAccessRevoked, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "BurnAccessRevoked", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperBurnAccessRevoked)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "BurnAccessRevoked", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseBurnAccessRevoked(log types.Log) (*BurnMintERC677HelperBurnAccessRevoked, error) {
	event := new(BurnMintERC677HelperBurnAccessRevoked)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "BurnAccessRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperMintAccessGrantedIterator struct {
	Event *BurnMintERC677HelperMintAccessGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperMintAccessGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperMintAccessGranted)
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
		it.Event = new(BurnMintERC677HelperMintAccessGranted)
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

func (it *BurnMintERC677HelperMintAccessGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperMintAccessGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperMintAccessGranted struct {
	Minter common.Address
	Raw    types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterMintAccessGranted(opts *bind.FilterOpts, minter []common.Address) (*BurnMintERC677HelperMintAccessGrantedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "MintAccessGranted", minterRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperMintAccessGrantedIterator{contract: _BurnMintERC677Helper.contract, event: "MintAccessGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchMintAccessGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperMintAccessGranted, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "MintAccessGranted", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperMintAccessGranted)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "MintAccessGranted", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseMintAccessGranted(log types.Log) (*BurnMintERC677HelperMintAccessGranted, error) {
	event := new(BurnMintERC677HelperMintAccessGranted)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "MintAccessGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperMintAccessRevokedIterator struct {
	Event *BurnMintERC677HelperMintAccessRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperMintAccessRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperMintAccessRevoked)
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
		it.Event = new(BurnMintERC677HelperMintAccessRevoked)
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

func (it *BurnMintERC677HelperMintAccessRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperMintAccessRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperMintAccessRevoked struct {
	Minter common.Address
	Raw    types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterMintAccessRevoked(opts *bind.FilterOpts, minter []common.Address) (*BurnMintERC677HelperMintAccessRevokedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "MintAccessRevoked", minterRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperMintAccessRevokedIterator{contract: _BurnMintERC677Helper.contract, event: "MintAccessRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchMintAccessRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperMintAccessRevoked, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "MintAccessRevoked", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperMintAccessRevoked)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "MintAccessRevoked", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseMintAccessRevoked(log types.Log) (*BurnMintERC677HelperMintAccessRevoked, error) {
	event := new(BurnMintERC677HelperMintAccessRevoked)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "MintAccessRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperOwnershipTransferRequestedIterator struct {
	Event *BurnMintERC677HelperOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperOwnershipTransferRequested)
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
		it.Event = new(BurnMintERC677HelperOwnershipTransferRequested)
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

func (it *BurnMintERC677HelperOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperOwnershipTransferRequestedIterator{contract: _BurnMintERC677Helper.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperOwnershipTransferRequested)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseOwnershipTransferRequested(log types.Log) (*BurnMintERC677HelperOwnershipTransferRequested, error) {
	event := new(BurnMintERC677HelperOwnershipTransferRequested)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperOwnershipTransferredIterator struct {
	Event *BurnMintERC677HelperOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperOwnershipTransferred)
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
		it.Event = new(BurnMintERC677HelperOwnershipTransferred)
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

func (it *BurnMintERC677HelperOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperOwnershipTransferredIterator{contract: _BurnMintERC677Helper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperOwnershipTransferred)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseOwnershipTransferred(log types.Log) (*BurnMintERC677HelperOwnershipTransferred, error) {
	event := new(BurnMintERC677HelperOwnershipTransferred)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperTransferIterator struct {
	Event *BurnMintERC677HelperTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperTransfer)
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
		it.Event = new(BurnMintERC677HelperTransfer)
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

func (it *BurnMintERC677HelperTransferIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperTransferIterator{contract: _BurnMintERC677Helper.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperTransfer)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseTransfer(log types.Log) (*BurnMintERC677HelperTransfer, error) {
	event := new(BurnMintERC677HelperTransfer)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperTransfer0Iterator struct {
	Event *BurnMintERC677HelperTransfer0

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperTransfer0Iterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperTransfer0)
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
		it.Event = new(BurnMintERC677HelperTransfer0)
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

func (it *BurnMintERC677HelperTransfer0Iterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperTransfer0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperTransfer0 struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterTransfer0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperTransfer0Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperTransfer0Iterator{contract: _BurnMintERC677Helper.contract, event: "Transfer0", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchTransfer0(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperTransfer0, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperTransfer0)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "Transfer0", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseTransfer0(log types.Log) (*BurnMintERC677HelperTransfer0, error) {
	event := new(BurnMintERC677HelperTransfer0)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "Transfer0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_BurnMintERC677Helper *BurnMintERC677Helper) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BurnMintERC677Helper.abi.Events["Approval"].ID:
		return _BurnMintERC677Helper.ParseApproval(log)
	case _BurnMintERC677Helper.abi.Events["BurnAccessGranted"].ID:
		return _BurnMintERC677Helper.ParseBurnAccessGranted(log)
	case _BurnMintERC677Helper.abi.Events["BurnAccessRevoked"].ID:
		return _BurnMintERC677Helper.ParseBurnAccessRevoked(log)
	case _BurnMintERC677Helper.abi.Events["MintAccessGranted"].ID:
		return _BurnMintERC677Helper.ParseMintAccessGranted(log)
	case _BurnMintERC677Helper.abi.Events["MintAccessRevoked"].ID:
		return _BurnMintERC677Helper.ParseMintAccessRevoked(log)
	case _BurnMintERC677Helper.abi.Events["OwnershipTransferRequested"].ID:
		return _BurnMintERC677Helper.ParseOwnershipTransferRequested(log)
	case _BurnMintERC677Helper.abi.Events["OwnershipTransferred"].ID:
		return _BurnMintERC677Helper.ParseOwnershipTransferred(log)
	case _BurnMintERC677Helper.abi.Events["Transfer"].ID:
		return _BurnMintERC677Helper.ParseTransfer(log)
	case _BurnMintERC677Helper.abi.Events["Transfer0"].ID:
		return _BurnMintERC677Helper.ParseTransfer0(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (BurnMintERC677HelperApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (BurnMintERC677HelperBurnAccessGranted) Topic() common.Hash {
	return common.HexToHash("0x92308bb7573b2a3d17ddb868b39d8ebec433f3194421abc22d084f89658c9bad")
}

func (BurnMintERC677HelperBurnAccessRevoked) Topic() common.Hash {
	return common.HexToHash("0x0a675452746933cefe3d74182e78db7afe57ba60eaa4234b5d85e9aa41b0610c")
}

func (BurnMintERC677HelperMintAccessGranted) Topic() common.Hash {
	return common.HexToHash("0xe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea")
}

func (BurnMintERC677HelperMintAccessRevoked) Topic() common.Hash {
	return common.HexToHash("0xed998b960f6340d045f620c119730f7aa7995e7425c2401d3a5b64ff998a59e9")
}

func (BurnMintERC677HelperOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BurnMintERC677HelperOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BurnMintERC677HelperTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (BurnMintERC677HelperTransfer0) Topic() common.Hash {
	return common.HexToHash("0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16")
}

func (_BurnMintERC677Helper *BurnMintERC677Helper) Address() common.Address {
	return _BurnMintERC677Helper.address
}

type BurnMintERC677HelperInterface interface {
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

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)

	Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error)

	DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error)

	Drip(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

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

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC677HelperApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*BurnMintERC677HelperApproval, error)

	FilterBurnAccessGranted(opts *bind.FilterOpts, burner []common.Address) (*BurnMintERC677HelperBurnAccessGrantedIterator, error)

	WatchBurnAccessGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperBurnAccessGranted, burner []common.Address) (event.Subscription, error)

	ParseBurnAccessGranted(log types.Log) (*BurnMintERC677HelperBurnAccessGranted, error)

	FilterBurnAccessRevoked(opts *bind.FilterOpts, burner []common.Address) (*BurnMintERC677HelperBurnAccessRevokedIterator, error)

	WatchBurnAccessRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperBurnAccessRevoked, burner []common.Address) (event.Subscription, error)

	ParseBurnAccessRevoked(log types.Log) (*BurnMintERC677HelperBurnAccessRevoked, error)

	FilterMintAccessGranted(opts *bind.FilterOpts, minter []common.Address) (*BurnMintERC677HelperMintAccessGrantedIterator, error)

	WatchMintAccessGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperMintAccessGranted, minter []common.Address) (event.Subscription, error)

	ParseMintAccessGranted(log types.Log) (*BurnMintERC677HelperMintAccessGranted, error)

	FilterMintAccessRevoked(opts *bind.FilterOpts, minter []common.Address) (*BurnMintERC677HelperMintAccessRevokedIterator, error)

	WatchMintAccessRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperMintAccessRevoked, minter []common.Address) (event.Subscription, error)

	ParseMintAccessRevoked(log types.Log) (*BurnMintERC677HelperMintAccessRevoked, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnMintERC677HelperOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnMintERC677HelperOwnershipTransferred, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*BurnMintERC677HelperTransfer, error)

	FilterTransfer0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperTransfer0Iterator, error)

	WatchTransfer0(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperTransfer0, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer0(log types.Log) (*BurnMintERC677HelperTransfer0, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
