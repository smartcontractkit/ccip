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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"supplyAfterMint\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotBurner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotMinter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"BurnAccessGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"BurnAccessRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MintAccessGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MintAccessRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBurners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"grantBurnRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnAndMinter\",\"type\":\"address\"}],\"name\":\"grantMintAndBurnRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"grantMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"isBurner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"revokeBurnRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"revokeMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200229138038062002291833981016040819052620000349162000277565b338060008686818160036200004a838262000391565b50600462000059828262000391565b5050506001600160a01b0384169150620000bc90505760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600580546001600160a01b0319166001600160a01b0384811691909117909155811615620000ef57620000ef8162000106565b50505060ff90911660805260a052506200045d9050565b336001600160a01b03821603620001605760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000b3565b600680546001600160a01b0319166001600160a01b03838116918217909255600554604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001da57600080fd5b81516001600160401b0380821115620001f757620001f7620001b2565b604051601f8301601f19908116603f01168101908282118183101715620002225762000222620001b2565b816040528381526020925086838588010111156200023f57600080fd5b600091505b8382101562000263578582018301518183018401529082019062000244565b600093810190920192909252949350505050565b600080600080608085870312156200028e57600080fd5b84516001600160401b0380821115620002a657600080fd5b620002b488838901620001c8565b95506020870151915080821115620002cb57600080fd5b50620002da87828801620001c8565b935050604085015160ff81168114620002f257600080fd5b6060959095015193969295505050565b600181811c908216806200031757607f821691505b6020821081036200033857634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200038c57600081815260208120601f850160051c81016020861015620003675750805b601f850160051c820191505b81811015620003885782815560010162000373565b5050505b505050565b81516001600160401b03811115620003ad57620003ad620001b2565b620003c581620003be845462000302565b846200033e565b602080601f831160018114620003fd5760008415620003e45750858301515b600019600386901b1c1916600185901b17855562000388565b600085815260208120601f198616915b828110156200042e578886015182559484019460019091019084016200040d565b50858210156200044d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a051611e0062000491600039600081816104470152818161087701526108a1015260006102710152611e006000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c806379cc67901161010f578063c2e3273d116100a2578063d73dd62311610071578063d73dd6231461046b578063dd62ed3e1461047e578063f2fde38b146104c4578063f81094f3146104d757600080fd5b8063c2e3273d1461040c578063c630948d1461041f578063c64d0ebc14610432578063d5abeb011461044557600080fd5b80639dc29fac116100de5780639dc29fac146103c0578063a457c2d7146103d3578063a9059cbb146103e6578063aa271e1a146103f957600080fd5b806379cc67901461037557806386fe8b43146103885780638da5cb5b1461039057806395d89b41146103b857600080fd5b806340c10f19116101875780636618846311610156578063661884631461030f5780636b32810b1461032257806370a082311461033757806379ba50971461036d57600080fd5b806340c10f19146102c157806342966c68146102d65780634334614a146102e95780634f5632f8146102fc57600080fd5b806323b872dd116101c357806323b872dd14610257578063313ce5671461026a578063395093511461029b5780634000aea0146102ae57600080fd5b806301ffc9a7146101f557806306fdde031461021d578063095ea7b31461023257806318160ddd14610245575b600080fd5b61020861020336600461196d565b6104ea565b60405190151581526020015b60405180910390f35b6102256105cf565b6040516102149190611a13565b610208610240366004611a4f565b610661565b6002545b604051908152602001610214565b610208610265366004611a79565b610679565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610214565b6102086102a9366004611a4f565b61069d565b6102086102bc366004611ae4565b6106e9565b6102d46102cf366004611a4f565b61080c565b005b6102d46102e4366004611bcd565b610933565b6102086102f7366004611be6565b610980565b6102d461030a366004611be6565b61098d565b61020861031d366004611a4f565b6109e9565b61032a6109fc565b6040516102149190611c01565b610249610345366004611be6565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6102d4610a0d565b6102d4610383366004611a4f565b610b0e565b61032a610b5d565b60055460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610214565b610225610b69565b6102d46103ce366004611a4f565b610b78565b6102086103e1366004611a4f565b610b82565b6102086103f4366004611a4f565b610c53565b610208610407366004611be6565b610c61565b6102d461041a366004611be6565b610c6e565b6102d461042d366004611be6565b610cca565b6102d4610440366004611be6565b610cd8565b7f0000000000000000000000000000000000000000000000000000000000000000610249565b6102d4610479366004611a4f565b610d34565b61024961048c366004611c5b565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6102d46104d2366004611be6565b610d3e565b6102d46104e5366004611be6565b610d4f565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f4000aea000000000000000000000000000000000000000000000000000000000148061057d57507fffffffff0000000000000000000000000000000000000000000000000000000082167fe6599b4d00000000000000000000000000000000000000000000000000000000145b806105c957507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b6060600380546105de90611c8e565b80601f016020809104026020016040519081016040528092919081815260200182805461060a90611c8e565b80156106575780601f1061062c57610100808354040283529160200191610657565b820191906000526020600020905b81548152906001019060200180831161063a57829003601f168201915b5050505050905090565b60003361066f818585610dab565b5060019392505050565b600033610687858285610ddf565b610692858585610eb0565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061066f90829086906106e4908790611d10565b610dab565b60006106f58484610c53565b508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c168585604051610755929190611d23565b60405180910390a373ffffffffffffffffffffffffffffffffffffffff84163b1561066f576040517fa4c0ed3600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85169063a4c0ed36906107d090339087908790600401611d44565b600060405180830381600087803b1580156107ea57600080fd5b505af11580156107fe573d6000803e3d6000fd5b505050505060019392505050565b61081533610c61565b610852576040517fe2c8c9d50000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b813073ffffffffffffffffffffffffffffffffffffffff82160361087557600080fd5b7f0000000000000000000000000000000000000000000000000000000000000000158015906108d657507f0000000000000000000000000000000000000000000000000000000000000000826108ca60025490565b6108d49190611d10565b115b1561092457816108e560025490565b6108ef9190611d10565b6040517fcbbf111300000000000000000000000000000000000000000000000000000000815260040161084991815260200190565b61092e8383610ede565b505050565b61093c33610980565b610974576040517fc820b10b000000000000000000000000000000000000000000000000000000008152336004820152602401610849565b61097d81610fd1565b50565b60006105c9600983610fdb565b61099561100a565b6109a060098261108d565b1561097d5760405173ffffffffffffffffffffffffffffffffffffffff8216907f0a675452746933cefe3d74182e78db7afe57ba60eaa4234b5d85e9aa41b0610c90600090a250565b60006109f58383610b82565b9392505050565b6060610a0860076110af565b905090565b60065473ffffffffffffffffffffffffffffffffffffffff163314610a8e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610849565b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560068054909116905560405173ffffffffffffffffffffffffffffffffffffffff909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b610b1733610980565b610b4f576040517fc820b10b000000000000000000000000000000000000000000000000000000008152336004820152602401610849565b610b5982826110bc565b5050565b6060610a0860096110af565b6060600480546105de90611c8e565b610b598282610b0e565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610c46576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610849565b6106928286868403610dab565b60003361066f818585610eb0565b60006105c9600783610fdb565b610c7661100a565b610c816007826110d1565b1561097d5760405173ffffffffffffffffffffffffffffffffffffffff8216907fe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea90600090a250565b610cd381610c6e565b61097d815b610ce061100a565b610ceb6009826110d1565b1561097d5760405173ffffffffffffffffffffffffffffffffffffffff8216907f92308bb7573b2a3d17ddb868b39d8ebec433f3194421abc22d084f89658c9bad90600090a250565b61092e828261069d565b610d4661100a565b61097d816110f3565b610d5761100a565b610d6260078261108d565b1561097d5760405173ffffffffffffffffffffffffffffffffffffffff8216907fed998b960f6340d045f620c119730f7aa7995e7425c2401d3a5b64ff998a59e990600090a250565b813073ffffffffffffffffffffffffffffffffffffffff821603610dce57600080fd5b610dd98484846111e9565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610dd95781811015610ea3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610849565b610dd98484848403610dab565b813073ffffffffffffffffffffffffffffffffffffffff821603610ed357600080fd5b610dd984848461139c565b73ffffffffffffffffffffffffffffffffffffffff8216610f5b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610849565b8060026000828254610f6d9190611d10565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b61097d338261160b565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018301602052604081205415156109f5565b60055473ffffffffffffffffffffffffffffffffffffffff16331461108b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610849565b565b60006109f58373ffffffffffffffffffffffffffffffffffffffff84166117cf565b606060006109f5836118c2565b6110c7823383610ddf565b610b59828261160b565b60006109f58373ffffffffffffffffffffffffffffffffffffffff841661191e565b3373ffffffffffffffffffffffffffffffffffffffff821603611172576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610849565b600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600554604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b73ffffffffffffffffffffffffffffffffffffffff831661128b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610849565b73ffffffffffffffffffffffffffffffffffffffff821661132e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610849565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff831661143f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610849565b73ffffffffffffffffffffffffffffffffffffffff82166114e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610849565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015611598576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610849565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610dd9565b73ffffffffffffffffffffffffffffffffffffffff82166116ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610849565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015611764576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610849565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b600081815260018301602052604081205480156118b85760006117f3600183611d82565b855490915060009061180790600190611d82565b905081811461186c57600086600001828154811061182757611827611d95565b906000526020600020015490508087600001848154811061184a5761184a611d95565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061187d5761187d611dc4565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506105c9565b60009150506105c9565b60608160000180548060200260200160405190810160405280929190818152602001828054801561191257602002820191906000526020600020905b8154815260200190600101908083116118fe575b50505050509050919050565b6000818152600183016020526040812054611965575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556105c9565b5060006105c9565b60006020828403121561197f57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146109f557600080fd5b6000815180845260005b818110156119d5576020818501810151868301820152016119b9565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006109f560208301846119af565b803573ffffffffffffffffffffffffffffffffffffffff81168114611a4a57600080fd5b919050565b60008060408385031215611a6257600080fd5b611a6b83611a26565b946020939093013593505050565b600080600060608486031215611a8e57600080fd5b611a9784611a26565b9250611aa560208501611a26565b9150604084013590509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080600060608486031215611af957600080fd5b611b0284611a26565b925060208401359150604084013567ffffffffffffffff80821115611b2657600080fd5b818601915086601f830112611b3a57600080fd5b813581811115611b4c57611b4c611ab5565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611b9257611b92611ab5565b81604052828152896020848701011115611bab57600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b600060208284031215611bdf57600080fd5b5035919050565b600060208284031215611bf857600080fd5b6109f582611a26565b6020808252825182820181905260009190848201906040850190845b81811015611c4f57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101611c1d565b50909695505050505050565b60008060408385031215611c6e57600080fd5b611c7783611a26565b9150611c8560208401611a26565b90509250929050565b600181811c90821680611ca257607f821691505b602082108103611cdb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156105c9576105c9611ce1565b828152604060208201526000611d3c60408301846119af565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201526000611d7960608301846119af565b95945050505050565b818103818111156105c9576105c9611ce1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000813000a",
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

func (_BurnMintERC677 *BurnMintERC677Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintERC677.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC677 *BurnMintERC677Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC677.Contract.SupportsInterface(&_BurnMintERC677.CallOpts, interfaceId)
}

func (_BurnMintERC677 *BurnMintERC677CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC677.Contract.SupportsInterface(&_BurnMintERC677.CallOpts, interfaceId)
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

func (_BurnMintERC677 *BurnMintERC677Transactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.contract.Transact(opts, "burn0", account, amount)
}

func (_BurnMintERC677 *BurnMintERC677Session) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.Burn0(&_BurnMintERC677.TransactOpts, account, amount)
}

func (_BurnMintERC677 *BurnMintERC677TransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677.Contract.Burn0(&_BurnMintERC677.TransactOpts, account, amount)
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
