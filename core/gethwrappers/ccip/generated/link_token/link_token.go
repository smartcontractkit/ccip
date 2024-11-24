// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package link_token

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
	zkSyncAccounts "github.com/zksync-sdk/zksync2-go/accounts"
	zkSyncClient "github.com/zksync-sdk/zksync2-go/clients"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
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

var LinkTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"supplyAfterMint\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotBurner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotMinter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"BurnAccessGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"BurnAccessRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MintAccessGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MintAccessRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBurners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"grantBurnRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnAndMinter\",\"type\":\"address\"}],\"name\":\"grantMintAndBurnRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"grantMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"isBurner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"revokeBurnRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"revokeMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040518060400160405280600f81526020016e21b430b4b72634b735902a37b5b2b760891b815250604051806040016040528060048152602001634c494e4b60e01b81525060126b033b2e3c9fd0803ce8000000338060008686818181600390816200007f91906200028c565b5060046200008e82826200028c565b5050506001600160a01b0384169150620000f190505760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600580546001600160a01b0319166001600160a01b0384811691909117909155811615620001245762000124816200013b565b50505060ff90911660805260a05250620003589050565b336001600160a01b03821603620001955760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000e8565b600680546001600160a01b0319166001600160a01b03838116918217909255600554604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200021257607f821691505b6020821081036200023357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200028757600081815260208120601f850160051c81016020861015620002625750805b601f850160051c820191505b8181101562000283578281556001016200026e565b5050505b505050565b81516001600160401b03811115620002a857620002a8620001e7565b620002c081620002b98454620001fd565b8462000239565b602080601f831160018114620002f85760008415620002df5750858301515b600019600386901b1c1916600185901b17855562000283565b600085815260208120601f198616915b82811015620003295788860151825594840194600190910190840162000308565b5085821015620003485787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a051611e4c6200038c60003960008181610447015281816108c301526108ed015260006102710152611e4c6000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c806379cc67901161010f578063c2e3273d116100a2578063d73dd62311610071578063d73dd6231461046b578063dd62ed3e1461047e578063f2fde38b146104c4578063f81094f3146104d757600080fd5b8063c2e3273d1461040c578063c630948d1461041f578063c64d0ebc14610432578063d5abeb011461044557600080fd5b80639dc29fac116100de5780639dc29fac146103c0578063a457c2d7146103d3578063a9059cbb146103e6578063aa271e1a146103f957600080fd5b806379cc67901461037557806386fe8b43146103885780638da5cb5b1461039057806395d89b41146103b857600080fd5b806340c10f19116101875780636618846311610156578063661884631461030f5780636b32810b1461032257806370a082311461033757806379ba50971461036d57600080fd5b806340c10f19146102c157806342966c68146102d65780634334614a146102e95780634f5632f8146102fc57600080fd5b806323b872dd116101c357806323b872dd14610257578063313ce5671461026a578063395093511461029b5780634000aea0146102ae57600080fd5b806301ffc9a7146101f557806306fdde031461021d578063095ea7b31461023257806318160ddd14610245575b600080fd5b6102086102033660046119b9565b6104ea565b60405190151581526020015b60405180910390f35b61022561061b565b6040516102149190611a5f565b610208610240366004611a9b565b6106ad565b6002545b604051908152602001610214565b610208610265366004611ac5565b6106c5565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610214565b6102086102a9366004611a9b565b6106e9565b6102086102bc366004611b30565b610735565b6102d46102cf366004611a9b565b610858565b005b6102d46102e4366004611c19565b61097f565b6102086102f7366004611c32565b6109cc565b6102d461030a366004611c32565b6109d9565b61020861031d366004611a9b565b610a35565b61032a610a48565b6040516102149190611c4d565b610249610345366004611c32565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6102d4610a59565b6102d4610383366004611a9b565b610b5a565b61032a610ba9565b60055460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610214565b610225610bb5565b6102d46103ce366004611a9b565b610bc4565b6102086103e1366004611a9b565b610bce565b6102086103f4366004611a9b565b610c9f565b610208610407366004611c32565b610cad565b6102d461041a366004611c32565b610cba565b6102d461042d366004611c32565b610d16565b6102d4610440366004611c32565b610d24565b7f0000000000000000000000000000000000000000000000000000000000000000610249565b6102d4610479366004611a9b565b610d80565b61024961048c366004611ca7565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6102d46104d2366004611c32565b610d8a565b6102d46104e5366004611c32565b610d9b565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b0700000000000000000000000000000000000000000000000000000000148061057d57507fffffffff0000000000000000000000000000000000000000000000000000000082167f4000aea000000000000000000000000000000000000000000000000000000000145b806105c957507fffffffff0000000000000000000000000000000000000000000000000000000082167fe6599b4d00000000000000000000000000000000000000000000000000000000145b8061061557507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b60606003805461062a90611cda565b80601f016020809104026020016040519081016040528092919081815260200182805461065690611cda565b80156106a35780601f10610678576101008083540402835291602001916106a3565b820191906000526020600020905b81548152906001019060200180831161068657829003601f168201915b5050505050905090565b6000336106bb818585610df7565b5060019392505050565b6000336106d3858285610e2b565b6106de858585610efc565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906106bb9082908690610730908790611d5c565b610df7565b60006107418484610c9f565b508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1685856040516107a1929190611d6f565b60405180910390a373ffffffffffffffffffffffffffffffffffffffff84163b156106bb576040517fa4c0ed3600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85169063a4c0ed369061081c90339087908790600401611d90565b600060405180830381600087803b15801561083657600080fd5b505af115801561084a573d6000803e3d6000fd5b505050505060019392505050565b61086133610cad565b61089e576040517fe2c8c9d50000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b813073ffffffffffffffffffffffffffffffffffffffff8216036108c157600080fd5b7f00000000000000000000000000000000000000000000000000000000000000001580159061092257507f00000000000000000000000000000000000000000000000000000000000000008261091660025490565b6109209190611d5c565b115b15610970578161093160025490565b61093b9190611d5c565b6040517fcbbf111300000000000000000000000000000000000000000000000000000000815260040161089591815260200190565b61097a8383610f2a565b505050565b610988336109cc565b6109c0576040517fc820b10b000000000000000000000000000000000000000000000000000000008152336004820152602401610895565b6109c98161101d565b50565b6000610615600983611027565b6109e1611056565b6109ec6009826110d9565b156109c95760405173ffffffffffffffffffffffffffffffffffffffff8216907f0a675452746933cefe3d74182e78db7afe57ba60eaa4234b5d85e9aa41b0610c90600090a250565b6000610a418383610bce565b9392505050565b6060610a5460076110fb565b905090565b60065473ffffffffffffffffffffffffffffffffffffffff163314610ada576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610895565b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560068054909116905560405173ffffffffffffffffffffffffffffffffffffffff909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b610b63336109cc565b610b9b576040517fc820b10b000000000000000000000000000000000000000000000000000000008152336004820152602401610895565b610ba58282611108565b5050565b6060610a5460096110fb565b60606004805461062a90611cda565b610ba58282610b5a565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610c92576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610895565b6106de8286868403610df7565b6000336106bb818585610efc565b6000610615600783611027565b610cc2611056565b610ccd60078261111d565b156109c95760405173ffffffffffffffffffffffffffffffffffffffff8216907fe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea90600090a250565b610d1f81610cba565b6109c9815b610d2c611056565b610d3760098261111d565b156109c95760405173ffffffffffffffffffffffffffffffffffffffff8216907f92308bb7573b2a3d17ddb868b39d8ebec433f3194421abc22d084f89658c9bad90600090a250565b61097a82826106e9565b610d92611056565b6109c98161113f565b610da3611056565b610dae6007826110d9565b156109c95760405173ffffffffffffffffffffffffffffffffffffffff8216907fed998b960f6340d045f620c119730f7aa7995e7425c2401d3a5b64ff998a59e990600090a250565b813073ffffffffffffffffffffffffffffffffffffffff821603610e1a57600080fd5b610e25848484611235565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610e255781811015610eef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610895565b610e258484848403610df7565b813073ffffffffffffffffffffffffffffffffffffffff821603610f1f57600080fd5b610e258484846113e8565b73ffffffffffffffffffffffffffffffffffffffff8216610fa7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610895565b8060026000828254610fb99190611d5c565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6109c93382611657565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515610a41565b60055473ffffffffffffffffffffffffffffffffffffffff1633146110d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610895565b565b6000610a418373ffffffffffffffffffffffffffffffffffffffff841661181b565b60606000610a418361190e565b611113823383610e2b565b610ba58282611657565b6000610a418373ffffffffffffffffffffffffffffffffffffffff841661196a565b3373ffffffffffffffffffffffffffffffffffffffff8216036111be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610895565b600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600554604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b73ffffffffffffffffffffffffffffffffffffffff83166112d7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610895565b73ffffffffffffffffffffffffffffffffffffffff821661137a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610895565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff831661148b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610895565b73ffffffffffffffffffffffffffffffffffffffff821661152e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610895565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156115e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610895565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610e25565b73ffffffffffffffffffffffffffffffffffffffff82166116fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610895565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260208190526040902054818110156117b0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610895565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b6000818152600183016020526040812054801561190457600061183f600183611dce565b855490915060009061185390600190611dce565b90508181146118b857600086600001828154811061187357611873611de1565b906000526020600020015490508087600001848154811061189657611896611de1565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806118c9576118c9611e10565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610615565b6000915050610615565b60608160000180548060200260200160405190810160405280929190818152602001828054801561195e57602002820191906000526020600020905b81548152602001906001019080831161194a575b50505050509050919050565b60008181526001830160205260408120546119b157508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610615565b506000610615565b6000602082840312156119cb57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610a4157600080fd5b6000815180845260005b81811015611a2157602081850181015186830182015201611a05565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610a4160208301846119fb565b803573ffffffffffffffffffffffffffffffffffffffff81168114611a9657600080fd5b919050565b60008060408385031215611aae57600080fd5b611ab783611a72565b946020939093013593505050565b600080600060608486031215611ada57600080fd5b611ae384611a72565b9250611af160208501611a72565b9150604084013590509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080600060608486031215611b4557600080fd5b611b4e84611a72565b925060208401359150604084013567ffffffffffffffff80821115611b7257600080fd5b818601915086601f830112611b8657600080fd5b813581811115611b9857611b98611b01565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611bde57611bde611b01565b81604052828152896020848701011115611bf757600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b600060208284031215611c2b57600080fd5b5035919050565b600060208284031215611c4457600080fd5b610a4182611a72565b6020808252825182820181905260009190848201906040850190845b81811015611c9b57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101611c69565b50909695505050505050565b60008060408385031215611cba57600080fd5b611cc383611a72565b9150611cd160208401611a72565b90509250929050565b600181811c90821680611cee57607f821691505b602082108103611d27577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561061557610615611d2d565b828152604060208201526000611d8860408301846119fb565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201526000611dc560608301846119fb565b95945050505050565b8181038181111561061557610615611d2d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000813000a",
}

var LinkTokenABI = LinkTokenMetaData.ABI

var LinkTokenBin = LinkTokenMetaData.Bin

func DeployLinkToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *CustomTransaction, *LinkToken, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	switch chainId.Uint64() {

	case 324, 280, 300:
		return DeployZkSyncLinkToken(auth, backend)
	}

	parsed, err := LinkTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LinkTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, &CustomTransaction{Transaction: tx, CustomHash: tx.Hash()}, &LinkToken{address: address, abi: *parsed, LinkTokenCaller: LinkTokenCaller{contract: contract}, LinkTokenTransactor: LinkTokenTransactor{contract: contract}, LinkTokenFilterer: LinkTokenFilterer{contract: contract}}, nil
}

type LinkToken struct {
	address common.Address
	abi     abi.ABI
	LinkTokenCaller
	LinkTokenTransactor
	LinkTokenFilterer
}

type LinkTokenCaller struct {
	contract *bind.BoundContract
}

type LinkTokenTransactor struct {
	contract *bind.BoundContract
}

type LinkTokenFilterer struct {
	contract *bind.BoundContract
}

type LinkTokenSession struct {
	Contract     *LinkToken
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type LinkTokenCallerSession struct {
	Contract *LinkTokenCaller
	CallOpts bind.CallOpts
}

type LinkTokenTransactorSession struct {
	Contract     *LinkTokenTransactor
	TransactOpts bind.TransactOpts
}

type LinkTokenRaw struct {
	Contract *LinkToken
}

type LinkTokenCallerRaw struct {
	Contract *LinkTokenCaller
}

type LinkTokenTransactorRaw struct {
	Contract *LinkTokenTransactor
}

func NewLinkToken(address common.Address, backend bind.ContractBackend) (*LinkToken, error) {
	abi, err := abi.JSON(strings.NewReader(LinkTokenABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindLinkToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LinkToken{address: address, abi: abi, LinkTokenCaller: LinkTokenCaller{contract: contract}, LinkTokenTransactor: LinkTokenTransactor{contract: contract}, LinkTokenFilterer: LinkTokenFilterer{contract: contract}}, nil
}

func NewLinkTokenCaller(address common.Address, caller bind.ContractCaller) (*LinkTokenCaller, error) {
	contract, err := bindLinkToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LinkTokenCaller{contract: contract}, nil
}

func NewLinkTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*LinkTokenTransactor, error) {
	contract, err := bindLinkToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LinkTokenTransactor{contract: contract}, nil
}

func NewLinkTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*LinkTokenFilterer, error) {
	contract, err := bindLinkToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LinkTokenFilterer{contract: contract}, nil
}

func bindLinkToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LinkTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_LinkToken *LinkTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LinkToken.Contract.LinkTokenCaller.contract.Call(opts, result, method, params...)
}

func (_LinkToken *LinkTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkToken.Contract.LinkTokenTransactor.contract.Transfer(opts)
}

func (_LinkToken *LinkTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkToken.Contract.LinkTokenTransactor.contract.Transact(opts, method, params...)
}

func (_LinkToken *LinkTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LinkToken.Contract.contract.Call(opts, result, method, params...)
}

func (_LinkToken *LinkTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkToken.Contract.contract.Transfer(opts)
}

func (_LinkToken *LinkTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkToken.Contract.contract.Transact(opts, method, params...)
}

func (_LinkToken *LinkTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LinkToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_LinkToken *LinkTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LinkToken.Contract.Allowance(&_LinkToken.CallOpts, owner, spender)
}

func (_LinkToken *LinkTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LinkToken.Contract.Allowance(&_LinkToken.CallOpts, owner, spender)
}

func (_LinkToken *LinkTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LinkToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_LinkToken *LinkTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LinkToken.Contract.BalanceOf(&_LinkToken.CallOpts, account)
}

func (_LinkToken *LinkTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LinkToken.Contract.BalanceOf(&_LinkToken.CallOpts, account)
}

func (_LinkToken *LinkTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LinkToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_LinkToken *LinkTokenSession) Decimals() (uint8, error) {
	return _LinkToken.Contract.Decimals(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCallerSession) Decimals() (uint8, error) {
	return _LinkToken.Contract.Decimals(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCaller) GetBurners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LinkToken.contract.Call(opts, &out, "getBurners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_LinkToken *LinkTokenSession) GetBurners() ([]common.Address, error) {
	return _LinkToken.Contract.GetBurners(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCallerSession) GetBurners() ([]common.Address, error) {
	return _LinkToken.Contract.GetBurners(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCaller) GetMinters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LinkToken.contract.Call(opts, &out, "getMinters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_LinkToken *LinkTokenSession) GetMinters() ([]common.Address, error) {
	return _LinkToken.Contract.GetMinters(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCallerSession) GetMinters() ([]common.Address, error) {
	return _LinkToken.Contract.GetMinters(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCaller) IsBurner(opts *bind.CallOpts, burner common.Address) (bool, error) {
	var out []interface{}
	err := _LinkToken.contract.Call(opts, &out, "isBurner", burner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LinkToken *LinkTokenSession) IsBurner(burner common.Address) (bool, error) {
	return _LinkToken.Contract.IsBurner(&_LinkToken.CallOpts, burner)
}

func (_LinkToken *LinkTokenCallerSession) IsBurner(burner common.Address) (bool, error) {
	return _LinkToken.Contract.IsBurner(&_LinkToken.CallOpts, burner)
}

func (_LinkToken *LinkTokenCaller) IsMinter(opts *bind.CallOpts, minter common.Address) (bool, error) {
	var out []interface{}
	err := _LinkToken.contract.Call(opts, &out, "isMinter", minter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LinkToken *LinkTokenSession) IsMinter(minter common.Address) (bool, error) {
	return _LinkToken.Contract.IsMinter(&_LinkToken.CallOpts, minter)
}

func (_LinkToken *LinkTokenCallerSession) IsMinter(minter common.Address) (bool, error) {
	return _LinkToken.Contract.IsMinter(&_LinkToken.CallOpts, minter)
}

func (_LinkToken *LinkTokenCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LinkToken.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_LinkToken *LinkTokenSession) MaxSupply() (*big.Int, error) {
	return _LinkToken.Contract.MaxSupply(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCallerSession) MaxSupply() (*big.Int, error) {
	return _LinkToken.Contract.MaxSupply(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LinkToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_LinkToken *LinkTokenSession) Name() (string, error) {
	return _LinkToken.Contract.Name(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCallerSession) Name() (string, error) {
	return _LinkToken.Contract.Name(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LinkToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LinkToken *LinkTokenSession) Owner() (common.Address, error) {
	return _LinkToken.Contract.Owner(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCallerSession) Owner() (common.Address, error) {
	return _LinkToken.Contract.Owner(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LinkToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LinkToken *LinkTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LinkToken.Contract.SupportsInterface(&_LinkToken.CallOpts, interfaceId)
}

func (_LinkToken *LinkTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LinkToken.Contract.SupportsInterface(&_LinkToken.CallOpts, interfaceId)
}

func (_LinkToken *LinkTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LinkToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_LinkToken *LinkTokenSession) Symbol() (string, error) {
	return _LinkToken.Contract.Symbol(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCallerSession) Symbol() (string, error) {
	return _LinkToken.Contract.Symbol(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LinkToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_LinkToken *LinkTokenSession) TotalSupply() (*big.Int, error) {
	return _LinkToken.Contract.TotalSupply(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _LinkToken.Contract.TotalSupply(&_LinkToken.CallOpts)
}

func (_LinkToken *LinkTokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "acceptOwnership")
}

func (_LinkToken *LinkTokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _LinkToken.Contract.AcceptOwnership(&_LinkToken.TransactOpts)
}

func (_LinkToken *LinkTokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LinkToken.Contract.AcceptOwnership(&_LinkToken.TransactOpts)
}

func (_LinkToken *LinkTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "approve", spender, amount)
}

func (_LinkToken *LinkTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.Approve(&_LinkToken.TransactOpts, spender, amount)
}

func (_LinkToken *LinkTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.Approve(&_LinkToken.TransactOpts, spender, amount)
}

func (_LinkToken *LinkTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "burn", amount)
}

func (_LinkToken *LinkTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.Burn(&_LinkToken.TransactOpts, amount)
}

func (_LinkToken *LinkTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.Burn(&_LinkToken.TransactOpts, amount)
}

func (_LinkToken *LinkTokenTransactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "burn0", account, amount)
}

func (_LinkToken *LinkTokenSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.Burn0(&_LinkToken.TransactOpts, account, amount)
}

func (_LinkToken *LinkTokenTransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.Burn0(&_LinkToken.TransactOpts, account, amount)
}

func (_LinkToken *LinkTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "burnFrom", account, amount)
}

func (_LinkToken *LinkTokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.BurnFrom(&_LinkToken.TransactOpts, account, amount)
}

func (_LinkToken *LinkTokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.BurnFrom(&_LinkToken.TransactOpts, account, amount)
}

func (_LinkToken *LinkTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

func (_LinkToken *LinkTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.DecreaseAllowance(&_LinkToken.TransactOpts, spender, subtractedValue)
}

func (_LinkToken *LinkTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.DecreaseAllowance(&_LinkToken.TransactOpts, spender, subtractedValue)
}

func (_LinkToken *LinkTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "decreaseApproval", spender, subtractedValue)
}

func (_LinkToken *LinkTokenSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.DecreaseApproval(&_LinkToken.TransactOpts, spender, subtractedValue)
}

func (_LinkToken *LinkTokenTransactorSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.DecreaseApproval(&_LinkToken.TransactOpts, spender, subtractedValue)
}

func (_LinkToken *LinkTokenTransactor) GrantBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "grantBurnRole", burner)
}

func (_LinkToken *LinkTokenSession) GrantBurnRole(burner common.Address) (*types.Transaction, error) {
	return _LinkToken.Contract.GrantBurnRole(&_LinkToken.TransactOpts, burner)
}

func (_LinkToken *LinkTokenTransactorSession) GrantBurnRole(burner common.Address) (*types.Transaction, error) {
	return _LinkToken.Contract.GrantBurnRole(&_LinkToken.TransactOpts, burner)
}

func (_LinkToken *LinkTokenTransactor) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (_LinkToken *LinkTokenSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _LinkToken.Contract.GrantMintAndBurnRoles(&_LinkToken.TransactOpts, burnAndMinter)
}

func (_LinkToken *LinkTokenTransactorSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _LinkToken.Contract.GrantMintAndBurnRoles(&_LinkToken.TransactOpts, burnAndMinter)
}

func (_LinkToken *LinkTokenTransactor) GrantMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "grantMintRole", minter)
}

func (_LinkToken *LinkTokenSession) GrantMintRole(minter common.Address) (*types.Transaction, error) {
	return _LinkToken.Contract.GrantMintRole(&_LinkToken.TransactOpts, minter)
}

func (_LinkToken *LinkTokenTransactorSession) GrantMintRole(minter common.Address) (*types.Transaction, error) {
	return _LinkToken.Contract.GrantMintRole(&_LinkToken.TransactOpts, minter)
}

func (_LinkToken *LinkTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

func (_LinkToken *LinkTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.IncreaseAllowance(&_LinkToken.TransactOpts, spender, addedValue)
}

func (_LinkToken *LinkTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.IncreaseAllowance(&_LinkToken.TransactOpts, spender, addedValue)
}

func (_LinkToken *LinkTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "increaseApproval", spender, addedValue)
}

func (_LinkToken *LinkTokenSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.IncreaseApproval(&_LinkToken.TransactOpts, spender, addedValue)
}

func (_LinkToken *LinkTokenTransactorSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.IncreaseApproval(&_LinkToken.TransactOpts, spender, addedValue)
}

func (_LinkToken *LinkTokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "mint", account, amount)
}

func (_LinkToken *LinkTokenSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.Mint(&_LinkToken.TransactOpts, account, amount)
}

func (_LinkToken *LinkTokenTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.Mint(&_LinkToken.TransactOpts, account, amount)
}

func (_LinkToken *LinkTokenTransactor) RevokeBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "revokeBurnRole", burner)
}

func (_LinkToken *LinkTokenSession) RevokeBurnRole(burner common.Address) (*types.Transaction, error) {
	return _LinkToken.Contract.RevokeBurnRole(&_LinkToken.TransactOpts, burner)
}

func (_LinkToken *LinkTokenTransactorSession) RevokeBurnRole(burner common.Address) (*types.Transaction, error) {
	return _LinkToken.Contract.RevokeBurnRole(&_LinkToken.TransactOpts, burner)
}

func (_LinkToken *LinkTokenTransactor) RevokeMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "revokeMintRole", minter)
}

func (_LinkToken *LinkTokenSession) RevokeMintRole(minter common.Address) (*types.Transaction, error) {
	return _LinkToken.Contract.RevokeMintRole(&_LinkToken.TransactOpts, minter)
}

func (_LinkToken *LinkTokenTransactorSession) RevokeMintRole(minter common.Address) (*types.Transaction, error) {
	return _LinkToken.Contract.RevokeMintRole(&_LinkToken.TransactOpts, minter)
}

func (_LinkToken *LinkTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "transfer", to, amount)
}

func (_LinkToken *LinkTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.Transfer(&_LinkToken.TransactOpts, to, amount)
}

func (_LinkToken *LinkTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.Transfer(&_LinkToken.TransactOpts, to, amount)
}

func (_LinkToken *LinkTokenTransactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "transferAndCall", to, amount, data)
}

func (_LinkToken *LinkTokenSession) TransferAndCall(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _LinkToken.Contract.TransferAndCall(&_LinkToken.TransactOpts, to, amount, data)
}

func (_LinkToken *LinkTokenTransactorSession) TransferAndCall(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _LinkToken.Contract.TransferAndCall(&_LinkToken.TransactOpts, to, amount, data)
}

func (_LinkToken *LinkTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

func (_LinkToken *LinkTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.TransferFrom(&_LinkToken.TransactOpts, from, to, amount)
}

func (_LinkToken *LinkTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LinkToken.Contract.TransferFrom(&_LinkToken.TransactOpts, from, to, amount)
}

func (_LinkToken *LinkTokenTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _LinkToken.contract.Transact(opts, "transferOwnership", to)
}

func (_LinkToken *LinkTokenSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LinkToken.Contract.TransferOwnership(&_LinkToken.TransactOpts, to)
}

func (_LinkToken *LinkTokenTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LinkToken.Contract.TransferOwnership(&_LinkToken.TransactOpts, to)
}

type LinkTokenApprovalIterator struct {
	Event *LinkTokenApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LinkTokenApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkTokenApproval)
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
		it.Event = new(LinkTokenApproval)
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

func (it *LinkTokenApprovalIterator) Error() error {
	return it.fail
}

func (it *LinkTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LinkTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_LinkToken *LinkTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LinkTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LinkToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LinkTokenApprovalIterator{contract: _LinkToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_LinkToken *LinkTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LinkTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LinkToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LinkTokenApproval)
				if err := _LinkToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_LinkToken *LinkTokenFilterer) ParseApproval(log types.Log) (*LinkTokenApproval, error) {
	event := new(LinkTokenApproval)
	if err := _LinkToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LinkTokenBurnAccessGrantedIterator struct {
	Event *LinkTokenBurnAccessGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LinkTokenBurnAccessGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkTokenBurnAccessGranted)
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
		it.Event = new(LinkTokenBurnAccessGranted)
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

func (it *LinkTokenBurnAccessGrantedIterator) Error() error {
	return it.fail
}

func (it *LinkTokenBurnAccessGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LinkTokenBurnAccessGranted struct {
	Burner common.Address
	Raw    types.Log
}

func (_LinkToken *LinkTokenFilterer) FilterBurnAccessGranted(opts *bind.FilterOpts, burner []common.Address) (*LinkTokenBurnAccessGrantedIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _LinkToken.contract.FilterLogs(opts, "BurnAccessGranted", burnerRule)
	if err != nil {
		return nil, err
	}
	return &LinkTokenBurnAccessGrantedIterator{contract: _LinkToken.contract, event: "BurnAccessGranted", logs: logs, sub: sub}, nil
}

func (_LinkToken *LinkTokenFilterer) WatchBurnAccessGranted(opts *bind.WatchOpts, sink chan<- *LinkTokenBurnAccessGranted, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _LinkToken.contract.WatchLogs(opts, "BurnAccessGranted", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LinkTokenBurnAccessGranted)
				if err := _LinkToken.contract.UnpackLog(event, "BurnAccessGranted", log); err != nil {
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

func (_LinkToken *LinkTokenFilterer) ParseBurnAccessGranted(log types.Log) (*LinkTokenBurnAccessGranted, error) {
	event := new(LinkTokenBurnAccessGranted)
	if err := _LinkToken.contract.UnpackLog(event, "BurnAccessGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LinkTokenBurnAccessRevokedIterator struct {
	Event *LinkTokenBurnAccessRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LinkTokenBurnAccessRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkTokenBurnAccessRevoked)
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
		it.Event = new(LinkTokenBurnAccessRevoked)
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

func (it *LinkTokenBurnAccessRevokedIterator) Error() error {
	return it.fail
}

func (it *LinkTokenBurnAccessRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LinkTokenBurnAccessRevoked struct {
	Burner common.Address
	Raw    types.Log
}

func (_LinkToken *LinkTokenFilterer) FilterBurnAccessRevoked(opts *bind.FilterOpts, burner []common.Address) (*LinkTokenBurnAccessRevokedIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _LinkToken.contract.FilterLogs(opts, "BurnAccessRevoked", burnerRule)
	if err != nil {
		return nil, err
	}
	return &LinkTokenBurnAccessRevokedIterator{contract: _LinkToken.contract, event: "BurnAccessRevoked", logs: logs, sub: sub}, nil
}

func (_LinkToken *LinkTokenFilterer) WatchBurnAccessRevoked(opts *bind.WatchOpts, sink chan<- *LinkTokenBurnAccessRevoked, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _LinkToken.contract.WatchLogs(opts, "BurnAccessRevoked", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LinkTokenBurnAccessRevoked)
				if err := _LinkToken.contract.UnpackLog(event, "BurnAccessRevoked", log); err != nil {
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

func (_LinkToken *LinkTokenFilterer) ParseBurnAccessRevoked(log types.Log) (*LinkTokenBurnAccessRevoked, error) {
	event := new(LinkTokenBurnAccessRevoked)
	if err := _LinkToken.contract.UnpackLog(event, "BurnAccessRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LinkTokenMintAccessGrantedIterator struct {
	Event *LinkTokenMintAccessGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LinkTokenMintAccessGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkTokenMintAccessGranted)
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
		it.Event = new(LinkTokenMintAccessGranted)
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

func (it *LinkTokenMintAccessGrantedIterator) Error() error {
	return it.fail
}

func (it *LinkTokenMintAccessGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LinkTokenMintAccessGranted struct {
	Minter common.Address
	Raw    types.Log
}

func (_LinkToken *LinkTokenFilterer) FilterMintAccessGranted(opts *bind.FilterOpts, minter []common.Address) (*LinkTokenMintAccessGrantedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _LinkToken.contract.FilterLogs(opts, "MintAccessGranted", minterRule)
	if err != nil {
		return nil, err
	}
	return &LinkTokenMintAccessGrantedIterator{contract: _LinkToken.contract, event: "MintAccessGranted", logs: logs, sub: sub}, nil
}

func (_LinkToken *LinkTokenFilterer) WatchMintAccessGranted(opts *bind.WatchOpts, sink chan<- *LinkTokenMintAccessGranted, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _LinkToken.contract.WatchLogs(opts, "MintAccessGranted", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LinkTokenMintAccessGranted)
				if err := _LinkToken.contract.UnpackLog(event, "MintAccessGranted", log); err != nil {
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

func (_LinkToken *LinkTokenFilterer) ParseMintAccessGranted(log types.Log) (*LinkTokenMintAccessGranted, error) {
	event := new(LinkTokenMintAccessGranted)
	if err := _LinkToken.contract.UnpackLog(event, "MintAccessGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LinkTokenMintAccessRevokedIterator struct {
	Event *LinkTokenMintAccessRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LinkTokenMintAccessRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkTokenMintAccessRevoked)
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
		it.Event = new(LinkTokenMintAccessRevoked)
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

func (it *LinkTokenMintAccessRevokedIterator) Error() error {
	return it.fail
}

func (it *LinkTokenMintAccessRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LinkTokenMintAccessRevoked struct {
	Minter common.Address
	Raw    types.Log
}

func (_LinkToken *LinkTokenFilterer) FilterMintAccessRevoked(opts *bind.FilterOpts, minter []common.Address) (*LinkTokenMintAccessRevokedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _LinkToken.contract.FilterLogs(opts, "MintAccessRevoked", minterRule)
	if err != nil {
		return nil, err
	}
	return &LinkTokenMintAccessRevokedIterator{contract: _LinkToken.contract, event: "MintAccessRevoked", logs: logs, sub: sub}, nil
}

func (_LinkToken *LinkTokenFilterer) WatchMintAccessRevoked(opts *bind.WatchOpts, sink chan<- *LinkTokenMintAccessRevoked, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _LinkToken.contract.WatchLogs(opts, "MintAccessRevoked", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LinkTokenMintAccessRevoked)
				if err := _LinkToken.contract.UnpackLog(event, "MintAccessRevoked", log); err != nil {
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

func (_LinkToken *LinkTokenFilterer) ParseMintAccessRevoked(log types.Log) (*LinkTokenMintAccessRevoked, error) {
	event := new(LinkTokenMintAccessRevoked)
	if err := _LinkToken.contract.UnpackLog(event, "MintAccessRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LinkTokenOwnershipTransferRequestedIterator struct {
	Event *LinkTokenOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LinkTokenOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkTokenOwnershipTransferRequested)
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
		it.Event = new(LinkTokenOwnershipTransferRequested)
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

func (it *LinkTokenOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *LinkTokenOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LinkTokenOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LinkToken *LinkTokenFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LinkTokenOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LinkToken.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LinkTokenOwnershipTransferRequestedIterator{contract: _LinkToken.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_LinkToken *LinkTokenFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LinkTokenOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LinkToken.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LinkTokenOwnershipTransferRequested)
				if err := _LinkToken.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_LinkToken *LinkTokenFilterer) ParseOwnershipTransferRequested(log types.Log) (*LinkTokenOwnershipTransferRequested, error) {
	event := new(LinkTokenOwnershipTransferRequested)
	if err := _LinkToken.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LinkTokenOwnershipTransferredIterator struct {
	Event *LinkTokenOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LinkTokenOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkTokenOwnershipTransferred)
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
		it.Event = new(LinkTokenOwnershipTransferred)
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

func (it *LinkTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *LinkTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LinkTokenOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LinkToken *LinkTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LinkTokenOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LinkToken.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LinkTokenOwnershipTransferredIterator{contract: _LinkToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_LinkToken *LinkTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LinkTokenOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LinkToken.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LinkTokenOwnershipTransferred)
				if err := _LinkToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_LinkToken *LinkTokenFilterer) ParseOwnershipTransferred(log types.Log) (*LinkTokenOwnershipTransferred, error) {
	event := new(LinkTokenOwnershipTransferred)
	if err := _LinkToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LinkTokenTransferIterator struct {
	Event *LinkTokenTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LinkTokenTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkTokenTransfer)
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
		it.Event = new(LinkTokenTransfer)
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

func (it *LinkTokenTransferIterator) Error() error {
	return it.fail
}

func (it *LinkTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LinkTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_LinkToken *LinkTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LinkTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LinkToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LinkTokenTransferIterator{contract: _LinkToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_LinkToken *LinkTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LinkTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LinkToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LinkTokenTransfer)
				if err := _LinkToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_LinkToken *LinkTokenFilterer) ParseTransfer(log types.Log) (*LinkTokenTransfer, error) {
	event := new(LinkTokenTransfer)
	if err := _LinkToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LinkTokenTransfer0Iterator struct {
	Event *LinkTokenTransfer0

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LinkTokenTransfer0Iterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkTokenTransfer0)
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
		it.Event = new(LinkTokenTransfer0)
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

func (it *LinkTokenTransfer0Iterator) Error() error {
	return it.fail
}

func (it *LinkTokenTransfer0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LinkTokenTransfer0 struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log
}

func (_LinkToken *LinkTokenFilterer) FilterTransfer0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LinkTokenTransfer0Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LinkToken.contract.FilterLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LinkTokenTransfer0Iterator{contract: _LinkToken.contract, event: "Transfer0", logs: logs, sub: sub}, nil
}

func (_LinkToken *LinkTokenFilterer) WatchTransfer0(opts *bind.WatchOpts, sink chan<- *LinkTokenTransfer0, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LinkToken.contract.WatchLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LinkTokenTransfer0)
				if err := _LinkToken.contract.UnpackLog(event, "Transfer0", log); err != nil {
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

func (_LinkToken *LinkTokenFilterer) ParseTransfer0(log types.Log) (*LinkTokenTransfer0, error) {
	event := new(LinkTokenTransfer0)
	if err := _LinkToken.contract.UnpackLog(event, "Transfer0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var LinkTokenZkBin string = ("0x0002000000000002000800000000000200010000000103550000006003100270000002aa0030019d00000001002001900000003a0000c13d000002aa023001970000008003000039000000400030043f000000040020008c000004ee0000413d000000000301043b000000e003300270000002b60030009c000000570000213d000002ce0030009c0000009c0000213d000002da0030009c000001810000213d000002e00030009c000001d00000213d000002e30030009c0000034f0000613d000002e40030009c000004ee0000c13d0000000001000416000000000001004b000004ee0000c13d0000000303000039000000000203041a000000010520019000000001012002700000007f0410018f00000000010460190000001f0010008c00000000060000390000000106002039000000000662013f0000000100600190000000510000c13d000000800010043f000000000005004b000001f20000613d000000000030043f000000020020008c0000056d0000413d000003180200004100000000040000190000000003040019000000000402041a000000a005300039000000000045043500000001022000390000002004300039000000000014004b000000310000413d000005cb0000013d0000000001000416000000000001004b000004ee0000c13d0000000f01000039000000c00010043f000002ab01000041000000e00010043f0000014001000039000000400010043f0000000406000039000001000060043f000002ac01000041000001200010043f0000000303000039000000000103041a000000010210019000000001041002700000007f0440618f0000001f0040008c00000000010000390000000101002039000000000012004b000000750000613d000002fd01000041000000000010043f0000002201000039000000040010043f000002f70100004100000aa700010430000002b70030009c000000c10000213d000002c30030009c000001980000213d000002c90030009c000001dc0000213d000002cc0030009c000003620000613d000002cd0030009c000004ee0000c13d0000000001000416000000000001004b000004ee0000c13d0000000902000039000000000102041a000000800010043f000000000020043f000000000001004b000001ae0000613d000000a004000039000002f90200004100000000030000190000000005040019000000000402041a000000000445043600000001022000390000000103300039000000000013004b0000006d0000413d0000053c0000013d000000200040008c0000008f0000413d000400000004001d000000000030043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002ad011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b00000004020000290000001f0220003900000005022002700000000002210019000000000021004b000000040600003900000003030000390000008f0000813d000000000001041b0000000101100039000000000021004b0000008b0000413d000000e00100043d000002ae011001970000001e011001bf000000000013041b000001000400043d000002af0040009c000001430000413d000002fd01000041000000000010043f0000004101000039000000040010043f000002f70100004100000aa700010430000002cf0030009c0000019f0000213d000002d50030009c000001f80000213d000002d80030009c000003750000613d000002d90030009c000004ee0000c13d000000240020008c000004ee0000413d0000000001000416000000000001004b000004ee0000c13d0000000001000411000000000010043f0000000a01000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b000000000101041a000000000001004b0000041c0000613d00000004010000390000000101100367000000000201043b00000000010004110aa509800000040f000000000100001900000aa60001042e000002b80030009c000001b00000213d000002be0030009c000002430000213d000002c10030009c000003960000613d000002c20030009c000004ee0000c13d000000240020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000401100370000000000101043b000400000001001d000002e50010009c000004ee0000213d0000000501000039000000000101041a000002e5011001970000000002000411000000000012004b000005600000c13d0000000401000029000000000010043f0000000801000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b000000000101041a000000000001004b000001100000c13d0000000701000039000000000201041a000002ef0020009c000000960000213d0000000103200039000000000031041b000002f20220009a0000000403000029000000000032041b000000000101041a000300000001001d000000000030043f0000000801000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b0000000302000029000000000021041b0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e8011001c70000800d020000390000000203000039000002f30400004100000004050000290aa50a9b0000040f0000000100200190000004ee0000613d0000000501000039000000000101041a000002e5011001970000000002000411000000000012004b000006220000c13d0000000401000029000000000010043f0000000a01000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b000000000101041a000000000001004b0000064d0000c13d0000000901000039000000000201041a000002ef0020009c000000960000213d0000000103200039000000000031041b000002f00220009a0000000403000029000000000032041b000000000101041a000300000001001d000000000030043f0000000a01000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b0000000302000029000000000021041b0000000001000414000004a80000013d000000000106041a000000010010019000000001031002700000007f0330618f0000001f0030008c00000000020000390000000102002039000000000121013f0000000100100190000000510000c13d000000200030008c0000016d0000413d000300000003001d000400000004001d000000000060043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002ad011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d00000004040000290000001f024000390000000502200270000000200040008c0000000002004019000000000301043b00000003010000290000001f01100039000000050110027000000000011300190000000002230019000000000012004b00000004060000390000016d0000813d000000000002041b0000000102200039000000000012004b000001690000413d0000001f0040008c000003440000a13d000400000004001d000000000060043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002ad011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d00000004070000290000031f02700198000000000101043b000004f00000c13d00000120030000390000000406000039000004ff0000013d000002db0030009c000002570000213d000002de0030009c000003da0000613d000002df0030009c000004ee0000c13d0000000001000416000000000001004b000004ee0000c13d0000000001000412000800000001001d000700000000003d000080050100003900000044030000390000000004000415000000080440008a0000000504400210000002ee020000410aa50a820000040f000000ff0110018f000000800010043f000002ed0100004100000aa60001042e000002c40030009c000002e30000213d000002c70030009c000003fe0000613d000002c80030009c000003430000613d000004ee0000013d000002d00030009c000002f20000213d000002d30030009c000003430000613d000002d40030009c000004ee0000c13d0000000001000416000000000001004b000004ee0000c13d0000000702000039000000000102041a000000800010043f000000000020043f000000000001004b000005320000c13d00000020020000390000053d0000013d000002b90030009c000003130000213d000002bc0030009c000004270000613d000002bd0030009c000004ee0000c13d000000440020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000402100370000000000202043b000002e50020009c000004ee0000213d0000002401100370000000000101043b000400000001001d000002e50010009c000004ee0000213d000000000020043f0000000101000039000000200010043f000000400200003900000000010000190aa50a6d0000040f0000000402000029000000000020043f000000200010043f00000000010000190000004002000039000004e00000013d000002e10030009c000004360000613d000002e20030009c000004ee0000c13d0000000001000416000000000001004b000004ee0000c13d0000000201000039000000000101041a000000800010043f000002ed0100004100000aa60001042e000002ca0030009c000004550000613d000002cb0030009c000004ee0000c13d0000000001000416000000000001004b000004ee0000c13d0000000403000039000000000203041a000000010520019000000001012002700000007f0410018f00000000010460190000001f0010008c00000000060000390000000106002039000000000662013f0000000100600190000000510000c13d000000800010043f000000000005004b0000056a0000c13d0000032001200197000000a00010043f000000000004004b000000c001000039000000a001006039000005cc0000013d000002d60030009c0000045e0000613d000002d70030009c000004ee0000c13d000000240020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000401100370000000000101043b000400000001001d000002e50010009c000004ee0000213d0000000501000039000000000101041a000002e5011001970000000002000411000000000012004b000005600000c13d0000000401000029000000000010043f0000000a01000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b000000000101041a000300000001001d000000000001004b0000064d0000613d0000000901000039000000000201041a000000000002004b0000033d0000613d0000000303000029000000010130008a000000000032004b0000067c0000c13d0000000302000029000002fe0220009a000000000002041b0000000902000039000000000012041b0000000401000029000000000010043f0000000a01000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b000000000001041b0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e8011001c70000800d020000390000000203000039000002ff04000041000006490000013d000002bf0030009c0000046c0000613d000002c00030009c000004ee0000c13d0000000001000416000000000001004b000004ee0000c13d0000000001000412000600000001001d000500200000003d000080050100003900000044030000390000000004000415000000060440008a0000000504400210000002ee020000410aa50a820000040f000000800010043f000002ed0100004100000aa60001042e000002dc0030009c000004b00000613d000002dd0030009c000004ee0000c13d000000640020008c000004ee0000413d0000000003000416000000000003004b000004ee0000c13d0000000403100370000000000303043b000400000003001d000002e50030009c000004ee0000213d0000002403100370000000000303043b000300000003001d0000004403100370000000000403043b000002ef0040009c000004ee0000213d0000002303400039000000000023004b000004ee0000813d0000000405400039000000000351034f000000000303043b000002ef0030009c000000960000213d0000001f063000390000031f066001970000003f066000390000031f06600197000003050060009c000000960000213d0000008006600039000000400060043f000000800030043f00000000043400190000002404400039000000000024004b000004ee0000213d0000002002500039000000000221034f0000031f043001980000001f0530018f000000a0014000390000028d0000613d000000a006000039000000000702034f000000007807043c0000000006860436000000000016004b000002890000c13d000000000005004b0000029a0000613d000000000242034f0000000304500210000000000501043300000000054501cf000000000545022f000000000202043b0000010004400089000000000242022f00000000024201cf000000000252019f0000000000210435000000a0013000390000000000010435000000040100002900000003020000290aa508850000040f000000400100043d000000200210003900000040030000390000000000320435000000030200002900000000002104350000004003100039000000800200043d000000000023043500000000070004110000006003100039000000000002004b000002b40000613d00000000040000190000000005340019000000a006400039000000000606043300000000006504350000002004400039000000000024004b000002ad0000413d0000001f042000390000031f04400197000000000232001900000000000204350000006002400039000002aa0020009c000002aa020080410000006002200210000002aa0010009c000002aa010080410000004001100210000000000112019f0000000002000414000002aa0020009c000002aa02008041000000c002200210000000000112019f000002e8011001c70000800d0200003900000003030000390000030604000041000000000507001900000004060000290aa50a9b0000040f0000000100200190000004ee0000613d00000307010000410000000000100443000000040100002900000004001004430000000001000414000002aa0010009c000002aa01008041000000c00110021000000308011001c700008002020000390aa50aa00000040f00000001002001900000076a0000613d000000000101043b000000000001004b000007200000c13d000000400100043d000300000001001d00000001010000390000000302000029000004cd0000013d000002c50030009c000004be0000613d000002c60030009c000004ee0000c13d000000240020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000401100370000000000101043b000002e50010009c000004ee0000213d0aa5085d0000040f000004680000013d000002d10030009c000004d30000613d000002d20030009c000004ee0000c13d0000000001000416000000000001004b000004ee0000c13d0000000601000039000000000201041a000002e5032001970000000006000411000000000036004b000005440000c13d0000000503000039000000000403041a000002b005400197000000000565019f000000000053041b000002b002200197000000000021041b0000000001000414000002e505400197000002aa0010009c000002aa01008041000000c001100210000002e8011001c70000800d020000390000000303000039000002fb040000410aa50a9b0000040f00000001002001900000064d0000c13d000004ee0000013d000002ba0030009c000004e50000613d000002bb0030009c000004ee0000c13d000000240020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000401100370000000000101043b000400000001001d000002e50010009c000004ee0000213d0000000501000039000000000101041a000002e5011001970000000002000411000000000012004b000005600000c13d0000000401000029000000000010043f0000000801000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b000000000101041a000300000001001d000000000001004b0000064d0000613d0000000701000039000000000201041a000000000002004b000006290000c13d000002fd01000041000000000010043f0000001101000039000000040010043f000002f70100004100000aa7000104300aa507b20000040f000000000004004b0000000001000019000003480000613d000001200100043d0000000302400210000003210220027f0000032102200167000000000121016f0000000102400210000000000121019f0000050a0000013d000000240020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000401100370000000000201043b0000031900200198000004ee0000c13d00000001010000390000031a0020009c0000056f0000213d0000031d0020009c000004e20000613d0000031e0020009c000000000100c019000000800010043f000002ed0100004100000aa60001042e000000440020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000402100370000000000202043b000400000002001d000002e50020009c000004ee0000213d0000002401100370000000000101043b000300000001001d0000000001000411000000000010043f0000000a01000039000000200010043f0000000001000414000004100000013d000000440020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000402100370000000000202043b000400000002001d000002e50020009c000004ee0000213d0000002401100370000000000101043b000300000001001d0000000001000411000000000010043f0000000801000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b000000000101041a000000000001004b0000057f0000c13d000000400100043d00000304020000410000041e0000013d000000240020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000401100370000000000101043b000400000001001d000002e50010009c000004ee0000213d0000000501000039000000000101041a000002e5011001970000000002000411000000000012004b000005600000c13d0000000401000029000000000010043f0000000801000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b000000000101041a000000000001004b0000064d0000c13d0000000701000039000000000201041a000002ef0020009c000000960000213d0000000103200039000000000031041b000002f20220009a0000000403000029000000000032041b000000000101041a000300000001001d000000000030043f0000000801000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b0000000302000029000000000021041b0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e8011001c70000800d020000390000000203000039000002f304000041000006490000013d000000640020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000402100370000000000202043b000400000002001d000002e50020009c000004ee0000213d0000002402100370000000000202043b000300000002001d000002e50020009c000004ee0000213d0000004401100370000000000301043b00000000020004110000000401000029000200000003001d0aa509dc0000040f0000000001000410000000030010006b000004ee0000613d000000040000006b000005df0000c13d000000400100043d00000064021000390000031003000041000000000032043500000044021000390000031103000041000000000032043500000024021000390000002503000039000006680000013d000000440020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000402100370000000000202043b000400000002001d000002e50020009c000004ee0000213d0000002401100370000000000101043b000300000001001d0000000001000411000000000010043f0000000a01000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b000000000101041a000000000001004b000005760000c13d000000400100043d000002f6020000410000000000210435000000040210003900000000030004110000000000320435000002aa0010009c000002aa010080410000004001100210000002f7011001c700000aa700010430000000440020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000402100370000000000302043b000002e50030009c000004ee0000213d0000002401100370000000000201043b00000000010300190aa508fc0000040f000000000100001900000aa60001042e000000440020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000402100370000000000202043b000400000002001d000002e50020009c000004ee0000213d0000002401100370000000000101043b000300000001001d0000000001000410000000040010006b000004ee0000613d0000000002000411000000000002004b000005a20000c13d000002b401000041000000800010043f0000002001000039000000840010043f0000002401000039000000a40010043f0000031601000041000000c40010043f0000031701000041000000e40010043f000003150100004100000aa7000104300000000001000416000000000001004b000004ee0000c13d0000000501000039000000000101041a000002e501100197000000800010043f000002ed0100004100000aa60001042e000000240020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000401100370000000000101043b000002e50010009c000004ee0000213d0aa508710000040f000000000001004b0000000001000039000000010100c039000004cc0000013d000000240020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000401100370000000000101043b000400000001001d000002e50010009c000004ee0000213d0000000501000039000000000101041a000002e5011001970000000002000411000000000012004b000005600000c13d0000000401000029000000000010043f0000000a01000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b000000000101041a000000000001004b0000064d0000c13d0000000901000039000000000201041a000002ef0020009c000000960000213d0000000103200039000000000031041b000002f00220009a0000000403000029000000000032041b000000000101041a000300000001001d000000000030043f0000000a01000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b0000000302000029000000000021041b0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e8011001c70000800d020000390000000203000039000002f104000041000006490000013d000000440020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000402100370000000000302043b000002e50030009c000004ee0000213d0000002401100370000000000201043b00000000010300190aa508fc0000040f000004cb0000013d000000440020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000402100370000000000302043b000002e50030009c000004ee0000213d0000002401100370000000000201043b00000000010300190aa508850000040f0000000101000039000000400200043d0000000000120435000002aa0020009c000002aa020080410000004001200210000002f5011001c700000aa60001042e000000240020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000401100370000000000101043b000002e50010009c000004ee0000213d000000000010043f000000200000043f000000400200003900000000010000190aa50a6d0000040f000000000101041a000000800010043f000002ed0100004100000aa60001042e000000240020008c000004ee0000413d0000000002000416000000000002004b000004ee0000c13d0000000401100370000000000601043b000002e50060009c0000054e0000a13d000000000100001900000aa700010430000000010320008a00000005033002700000000003310019000000200400003900000001033000390000000406000039000000000504001900000100044000390000000004040433000000000041041b00000020045000390000000101100039000000000031004b000004f60000c13d0000012003500039000000000072004b000005080000813d0000000302700210000000f80220018f000003210220027f00000321022001670000000003030433000000000223016f000000000021041b000000010170021000000001011001bf000000000016041b0000000001000411000000000001004b0000051f0000c13d000000400100043d0000004402100039000002b3030000410000000000320435000000240210003900000018030000390000000000320435000002b4020000410000000000210435000000040210003900000020030000390000000000320435000002aa0010009c000002aa010080410000004001100210000002b5011001c700000aa7000104300000000502000039000000000302041a000002b003300197000000000113019f000000000012041b0000001201000039000000800010043f000002b102000041000000a00020043f0000014000000443000001600010044300000020010000390000018000100443000001a000200443000001000010044300000002010000390000012000100443000002b20100004100000aa60001042e000000a004000039000002fc0200004100000000030000190000000005040019000000000402041a000000000445043600000001022000390000000103300039000000000013004b000005350000413d000000600250008a00000080010000390aa507a00000040f000000400100043d000400000001001d00000080020000390aa5084d0000040f000005d50000013d000002b401000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f000002fa01000041000000c40010043f000002ec0100004100000aa7000104300000000501000039000000000101041a000002e5011001970000000005000411000000000015004b000005600000c13d000000000056004b000005b00000c13d000002b401000041000000800010043f0000002001000039000000840010043f0000001701000039000000a40010043f000002eb01000041000000c40010043f000002ec0100004100000aa700010430000002b401000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f000002f401000041000000c40010043f000002ec0100004100000aa700010430000000000030043f000000020020008c000005c10000813d000000a001000039000005cc0000013d0000031b0020009c000004e20000613d0000031c0020009c000004e20000613d000000800000043f000002ed0100004100000aa60001042e0000000401000029000000000200041100000003030000290aa509dc0000040f000000040100002900000003020000290aa509800000040f000000000100001900000aa60001042e0000000001000410000000040010006b000004ee0000613d000002ee01000041000000000010044300000000010004120000000400100443000000200100003900000024001004430000000001000414000002aa0010009c000002aa01008041000000c00110021000000300011001c700008005020000390aa50aa00000040f00000001002001900000076a0000613d000000000101043b000000000001004b000006730000613d0000000202000039000000000202041a0000000303000029000000000032001a0000033d0000413d0000000002320019000000000012004b000006730000a13d000000400100043d0000030103000041000000000031043500000004031000390000000000230435000004220000013d000000040000006b000005eb0000c13d000002b401000041000000800010043f0000002001000039000000840010043f0000002201000039000000a40010043f0000031301000041000000c40010043f0000031401000041000000e40010043f000003150100004100000aa7000104300000000601000039000000000201041a000002b002200197000000000262019f000000000021041b0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e8011001c70000800d020000390000000303000039000002ea040000410aa50a9b0000040f00000001002001900000064d0000c13d000004ee0000013d000002f80200004100000000040000190000000003040019000000000402041a000000a005300039000000000045043500000001022000390000002004300039000000000014004b000005c30000413d000000c001300039000000800210008a00000080010000390aa507a00000040f0000002001000039000000400200043d000400000002001d000000000212043600000080010000390aa5078e0000040f00000004020000290000000001210049000002aa0010009c000002aa010080410000006001100210000002aa0020009c000002aa020080410000004002200210000000000121019f00000aa60001042e000000030000006b0000064f0000c13d000000400100043d00000064021000390000030e03000041000000000032043500000044021000390000030f03000041000000000032043500000024021000390000002303000039000006680000013d000000000020043f0000000101000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b0000000402000029000000000020043f000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b0000000302000029000000000021041b000000400100043d0000000000210435000002aa0010009c000002aa0100804100000040011002100000000002000414000002aa0020009c000002aa02008041000000c002200210000000000112019f000002ad011001c70000800d0200003900000003030000390000031204000041000000000500041100000004060000290aa50a9b0000040f0000000100200190000004ee0000613d000000400100043d00000001020000390000000000210435000002aa0010009c000002aa010080410000004001100210000002f5011001c700000aa60001042e000000400100043d0000004402100039000002f403000041000000000032043500000024021000390000001603000039000005140000013d0000000303000029000000010130008a000000000032004b0000069a0000c13d0000000302000029000002e70220009a000000000002041b0000000702000039000000000012041b0000000401000029000000000010043f0000000801000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b000000000001041b0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e8011001c70000800d020000390000000203000039000002e90400004100000004050000290aa50a9b0000040f0000000100200190000004ee0000613d000000000100001900000aa60001042e0000000401000029000000000010043f000000200000043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b000000000101041a0001000200100074000006ed0000813d000000400100043d00000064021000390000030b03000041000000000032043500000044021000390000030c030000410000000000320435000000240210003900000026030000390000000000320435000002b4020000410000000000210435000000040210003900000020030000390000000000320435000002aa0010009c000002aa0100804100000040011002100000030d011001c700000aa700010430000000040000006b000006c10000c13d000000400100043d00000044021000390000030303000041000000000032043500000024021000390000001f03000039000005140000013d000000000012004b000006bb0000a13d0000000301000029000002fe0110009a000002fe0220009a000000000202041a000000000021041b000000000020043f0000000a01000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b0000000302000029000000000021041b0000000901000039000000000101041a000300000001001d000000000001004b000006b50000613d0000000301000029000000010110008a000002260000013d000000000012004b000006bb0000a13d0000000301000029000002e70110009a000002e70220009a000000000202041a000000000021041b000000000020043f0000000801000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b0000000302000029000000000021041b0000000701000039000000000101041a000300000001001d000000000001004b0000076b0000c13d000002fd01000041000000000010043f0000003101000039000000040010043f000002f70100004100000aa700010430000002fd01000041000000000010043f0000003201000039000000040010043f000002f70100004100000aa7000104300000000201000039000000000201041a0000000303000029000000000032001a0000033d0000413d0000000002320019000000000021041b0000000401000029000000000010043f000000200000043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b000000000201041a00000003030000290000000002320019000000000021041b000000400100043d0000000000310435000002aa0010009c000002aa0100804100000040011002100000000002000414000002aa0020009c000002aa02008041000000c002200210000000000112019f000002ad011001c70000800d0200003900000003030000390000030204000041000000000500001900000004060000290aa50a9b0000040f00000001002001900000064d0000c13d000004ee0000013d0000000401000029000000000010043f000000200000043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b0000000102000029000000000021041b0000000301000029000000000010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000004ee0000613d000000000101043b000000000201041a00000002030000290000000002320019000000000021041b000000400100043d0000000000310435000002aa0010009c000002aa0100804100000040011002100000000002000414000002aa0020009c000002aa02008041000000c002200210000000000112019f000002ad011001c70000800d0200003900000003030000390000030204000041000000040500002900000003060000290aa50a9b0000040f00000001002001900000061a0000c13d000004ee0000013d00000307010000410000000000100443000000040100002900000004001004430000000001000414000002aa0010009c000002aa01008041000000c00110021000000308011001c700008002020000390aa50aa00000040f00000001002001900000076a0000613d000000000101043b000000000001004b000004ee0000613d000000400300043d000000440130003900000060020000390000000000210435000000240130003900000003020000290000000000210435000003090100004100000000001304350000000401300039000000000200041100000000002104350000006402300039000000800100043d0000000000120435000300000003001d0000008402300039000000000001004b0000074b0000613d00000000030000190000000004230019000000a005300039000000000505043300000000005404350000002003300039000000000013004b000007440000413d0000000002210019000000000002043500000000020004140000000403000029000000040030008c000007640000613d0000001f011000390000031f011001970000008401100039000002aa0010009c000002aa0100804100000060011002100000000303000029000002aa0030009c000002aa030080410000004003300210000000000131019f000002aa0020009c000002aa02008041000000c002200210000000000112019f00000004020000290aa50a9b0000040f00000001002001900000076e0000613d0000000301000029000002ef0010009c000000960000213d0000000301000029000000400010043f000002e00000013d000000000001042f0000000301000029000000010110008a0000062d0000013d00000060061002700000001f0460018f0000030a05600198000000400200043d00000000035200190000077a0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000038004b000007760000c13d000002aa06600197000000000004004b000007880000613d000000000151034f0000000304400210000000000503043300000000054501cf000000000545022f000000000101043b0000010004400089000000000141022f00000000014101cf000000000151019f00000000001304350000006001600210000002aa0020009c000002aa020080410000004002200210000000000112019f00000aa70001043000000000430104340000000001320436000000000003004b0000079a0000613d000000000200001900000000052100190000000006240019000000000606043300000000006504350000002002200039000000000032004b000007930000413d000000000231001900000000000204350000001f023000390000031f022001970000000001210019000000000001042d0000001f022000390000031f022001970000000001120019000000000021004b00000000020000390000000102004039000002ef0010009c000007ac0000213d0000000100200190000007ac0000c13d000000400010043f000000000001042d000002fd01000041000000000010043f0000004101000039000000040010043f000002f70100004100000aa70001043000020000000000020000000001000416000000000001004b000007f90000c13d0000000001000031000003220010009c000007f90000213d000000430010008c000007f90000a13d00000001010003670000000402100370000000000202043b000200000002001d000002e50020009c000007f90000213d0000002401100370000000000101043b000100000001001d0000000001000411000000000010043f0000000101000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000007f90000613d000000000101043b0000000202000029000000000020043f000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000007f90000613d000000000101043b000000000101041a000000010310006c000007f60000813d000000400100043d000000640210003900000323030000410000000000320435000000440210003900000324030000410000000000320435000000240210003900000025030000390000000000320435000002b4020000410000000000210435000000040210003900000020030000390000000000320435000002aa0010009c000002aa0100804100000040011002100000030d011001c700000aa7000104300000000001000410000000020010006b000007fb0000c13d000000000100001900000aa7000104300000000001000411000000000001004b000008080000c13d000000400100043d00000064021000390000031703000041000000000032043500000044021000390000031603000041000000000032043500000024021000390000002403000039000007eb0000013d000000020000006b000008140000c13d000000400100043d00000064021000390000031403000041000000000032043500000044021000390000031303000041000000000032043500000024021000390000002203000039000007eb0000013d0000000001000411000000000010043f0000000101000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c70000801002000039000100000003001d0aa50aa00000040f0000000100200190000007f90000613d000000000101043b0000000202000029000000000020043f000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f00000001030000290000000100200190000007f90000613d000000000101043b000000000031041b000000400100043d0000000000310435000002aa0010009c000002aa0100804100000040011002100000000002000414000002aa0020009c000002aa02008041000000c002200210000000000112019f000002ad011001c70000800d0200003900000003030000390000031204000041000000000500041100000002060000290aa50a9b0000040f0000000100200190000007f90000613d000000400100043d00000001020000390000000000210435000002aa0010009c000002aa010080410000004001100210000002f5011001c700000aa60001042e00000020030000390000000004310436000000000302043300000000003404350000004001100039000000000003004b0000085c0000613d000000000400001900000020022000390000000005020433000002e50550019700000000015104360000000104400039000000000034004b000008550000413d000000000001042d000000000010043f0000000801000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f00000001002001900000086f0000613d000000000101043b000000000101041a000000000001004b0000000001000039000000010100c039000000000001042d000000000100001900000aa700010430000000000010043f0000000a01000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000008830000613d000000000101043b000000000101041a000000000001004b0000000001000039000000010100c039000000000001042d000000000100001900000aa7000104300003000000000002000200000002001d000302e50010019b0000000001000410000000030010006b000008d20000613d0000000002000411000000000002004b000008d40000613d000000030000006b000008de0000613d000000000020043f000000200000043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000008d20000613d000000000101043b000000000101041a0001000200100074000008e80000413d0000000001000411000000000010043f000000200000043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000008d20000613d000000000101043b0000000102000029000000000021041b0000000301000029000000000010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000008d20000613d000000000101043b000000000201041a00000002030000290000000002320019000000000021041b000000400100043d0000000000310435000002aa0010009c000002aa0100804100000040011002100000000002000414000002aa0020009c000002aa02008041000000c002200210000000000112019f000002ad011001c70000800d0200003900000003030000390000030204000041000000000500041100000003060000290aa50a9b0000040f0000000100200190000008d20000613d000000000001042d000000000100001900000aa700010430000000400100043d00000064021000390000031003000041000000000032043500000044021000390000031103000041000000000032043500000024021000390000002503000039000008f10000013d000000400100043d00000064021000390000030e03000041000000000032043500000044021000390000030f03000041000000000032043500000024021000390000002303000039000008f10000013d000000400100043d00000064021000390000030b03000041000000000032043500000044021000390000030c030000410000000000320435000000240210003900000026030000390000000000320435000002b4020000410000000000210435000000040210003900000020030000390000000000320435000002aa0010009c000002aa0100804100000040011002100000030d011001c700000aa7000104300003000000000002000200000002001d000300000001001d0000000001000411000000000010043f0000000101000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f00000001002001900000095a0000613d000000000101043b0000000302000029000002e502200197000300000002001d000000000020043f000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f00000001002001900000095a0000613d000000000101043b000000000201041a000000020020002a0000097a0000413d000100000002001d0000000001000410000000030010006b0000095a0000613d0000000001000411000000000001004b0000095c0000613d000000030000006b000009660000613d0000000001000411000000000010043f0000000101000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f00000001002001900000095a0000613d000000000101043b0000000302000029000000000020043f000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f00000001002001900000095a0000613d00000001030000290000000202300029000000000101043b000000000021041b000000400100043d0000000000210435000002aa0010009c000002aa0100804100000040011002100000000002000414000002aa0020009c000002aa02008041000000c002200210000000000112019f000002ad011001c70000800d0200003900000003030000390000031204000041000000000500041100000003060000290aa50a9b0000040f00000001002001900000095a0000613d000000000001042d000000000100001900000aa700010430000000400100043d000000640210003900000317030000410000000000320435000000440210003900000316030000410000000000320435000000240210003900000024030000390000096f0000013d000000400100043d000000640210003900000314030000410000000000320435000000440210003900000313030000410000000000320435000000240210003900000022030000390000000000320435000002b4020000410000000000210435000000040210003900000020030000390000000000320435000002aa0010009c000002aa0100804100000040011002100000030d011001c700000aa700010430000002fd01000041000000000010043f0000001101000039000000040010043f000002f70100004100000aa7000104300003000000000002000300000002001d000002e503100198000009be0000613d000000000030043f000000200000043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c70000801002000039000200000003001d0aa50aa00000040f0000000100200190000009bc0000613d0000000202000029000000000101043b000000000101041a0001000300100074000009c80000413d000000000020043f000000200000043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f0000000100200190000009bc0000613d000000000101043b0000000102000029000000000021041b0000000201000039000000000201041a00000003030000290000000002320049000000000021041b000000400100043d0000000000310435000002aa0010009c000002aa0100804100000040011002100000000002000414000002aa0020009c000002aa02008041000000c002200210000000000112019f000002ad011001c70000800d0200003900000003030000390000030204000041000000020500002900000000060000190aa50a9b0000040f0000000100200190000009bc0000613d000000000001042d000000000100001900000aa700010430000000400100043d00000064021000390000032703000041000000000032043500000044021000390000032803000041000000000032043500000024021000390000002103000039000009d10000013d000000400100043d000000640210003900000325030000410000000000320435000000440210003900000326030000410000000000320435000000240210003900000022030000390000000000320435000002b4020000410000000000210435000000040210003900000020030000390000000000320435000002aa0010009c000002aa0100804100000040011002100000030d011001c700000aa7000104300003000000000002000100000003001d000300000002001d000002e501100197000200000001001d000000000010043f0000000101000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f000000010020019000000a3b0000613d000000000101043b0000000302000029000002e502200197000300000002001d000000000020043f000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f000000010020019000000a3b0000613d000000000101043b000000000101041a000003210010009c00000a3a0000613d000000010110006c00000a3d0000413d000100000001001d0000000001000410000000030010006b00000a3b0000613d000000020000006b00000a4e0000613d000000030000006b00000a580000613d0000000201000029000000000010043f0000000101000039000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f000000010020019000000a3b0000613d000000000101043b0000000302000029000000000020043f000000200010043f0000000001000414000002aa0010009c000002aa01008041000000c001100210000002e6011001c700008010020000390aa50aa00000040f000000010020019000000a3b0000613d000000000101043b0000000102000029000000000021041b000000400100043d0000000000210435000002aa0010009c000002aa0100804100000040011002100000000002000414000002aa0020009c000002aa02008041000000c002200210000000000112019f000002ad011001c70000800d0200003900000003030000390000031204000041000000020500002900000003060000290aa50a9b0000040f000000010020019000000a3b0000613d000000000001042d000000000100001900000aa700010430000000400100043d00000044021000390000032903000041000000000032043500000024021000390000001d030000390000000000320435000002b4020000410000000000210435000000040210003900000020030000390000000000320435000002aa0010009c000002aa010080410000004001100210000002b5011001c700000aa700010430000000400100043d0000006402100039000003170300004100000000003204350000004402100039000003160300004100000000003204350000002402100039000000240300003900000a610000013d000000400100043d000000640210003900000314030000410000000000320435000000440210003900000313030000410000000000320435000000240210003900000022030000390000000000320435000002b4020000410000000000210435000000040210003900000020030000390000000000320435000002aa0010009c000002aa0100804100000040011002100000030d011001c700000aa700010430000000000001042f000002aa0010009c000002aa010080410000004001100210000002aa0020009c000002aa020080410000006002200210000000000112019f0000000002000414000002aa0020009c000002aa02008041000000c002200210000000000112019f000002e8011001c700008010020000390aa50aa00000040f000000010020019000000a800000613d000000000101043b000000000001042d000000000100001900000aa70001043000000000050100190000000000200443000000040100003900000005024002700000000002020031000000000121043a0000002004400039000000000031004b00000a850000413d000002aa0030009c000002aa0300804100000060013002100000000002000414000002aa0020009c000002aa02008041000000c002200210000000000112019f0000032a011001c700000000020500190aa50aa00000040f000000010020019000000a9a0000613d000000000101043b000000000001042d000000000001042f00000a9e002104210000000102000039000000000001042d0000000002000019000000000001042d00000aa3002104230000000102000039000000000001042d0000000002000019000000000001042d00000aa50000043200000aa60001042e00000aa70001043000000000000000000000000000000000000000000000000000000000ffffffff436861696e4c696e6b20546f6b656e00000000000000000000000000000000004c494e4b000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000020000000000000000000000000ffffffffffffffffffffffffffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000ffffffffffffffffffffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000033b2e3c9fd0803ce800000000000002000000000000000000000000000000c000000100000000000000000043616e6e6f7420736574206f776e657220746f207a65726f000000000000000008c379a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000000000000000000000079cc678f00000000000000000000000000000000000000000000000000000000c2e3273c00000000000000000000000000000000000000000000000000000000d73dd62200000000000000000000000000000000000000000000000000000000f2fde38a00000000000000000000000000000000000000000000000000000000f2fde38b00000000000000000000000000000000000000000000000000000000f81094f300000000000000000000000000000000000000000000000000000000d73dd62300000000000000000000000000000000000000000000000000000000dd62ed3e00000000000000000000000000000000000000000000000000000000c64d0ebb00000000000000000000000000000000000000000000000000000000c64d0ebc00000000000000000000000000000000000000000000000000000000d5abeb0100000000000000000000000000000000000000000000000000000000c2e3273d00000000000000000000000000000000000000000000000000000000c630948d000000000000000000000000000000000000000000000000000000009dc29fab00000000000000000000000000000000000000000000000000000000a9059cba00000000000000000000000000000000000000000000000000000000a9059cbb00000000000000000000000000000000000000000000000000000000aa271e1a000000000000000000000000000000000000000000000000000000009dc29fac00000000000000000000000000000000000000000000000000000000a457c2d7000000000000000000000000000000000000000000000000000000008da5cb5a000000000000000000000000000000000000000000000000000000008da5cb5b0000000000000000000000000000000000000000000000000000000095d89b410000000000000000000000000000000000000000000000000000000079cc67900000000000000000000000000000000000000000000000000000000086fe8b430000000000000000000000000000000000000000000000000000000040c10f1800000000000000000000000000000000000000000000000000000000661884620000000000000000000000000000000000000000000000000000000070a082300000000000000000000000000000000000000000000000000000000070a082310000000000000000000000000000000000000000000000000000000079ba50970000000000000000000000000000000000000000000000000000000066188463000000000000000000000000000000000000000000000000000000006b32810b0000000000000000000000000000000000000000000000000000000043346149000000000000000000000000000000000000000000000000000000004334614a000000000000000000000000000000000000000000000000000000004f5632f80000000000000000000000000000000000000000000000000000000040c10f190000000000000000000000000000000000000000000000000000000042966c680000000000000000000000000000000000000000000000000000000023b872dc00000000000000000000000000000000000000000000000000000000395093500000000000000000000000000000000000000000000000000000000039509351000000000000000000000000000000000000000000000000000000004000aea00000000000000000000000000000000000000000000000000000000023b872dd00000000000000000000000000000000000000000000000000000000313ce56700000000000000000000000000000000000000000000000000000000095ea7b200000000000000000000000000000000000000000000000000000000095ea7b30000000000000000000000000000000000000000000000000000000018160ddd0000000000000000000000000000000000000000000000000000000001ffc9a70000000000000000000000000000000000000000000000000000000006fdde03000000000000000000000000ffffffffffffffffffffffffffffffffffffffff0200000000000000000000000000000000000040000000000000000000000000599336d74a1247d50642b66dd6abeaa5484f6bd96b415b31bb99e26578c939790200000000000000000000000000000000000000000000000000000000000000ed998b960f6340d045f620c119730f7aa7995e7425c2401d3a5b64ff998a59e9ed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127843616e6e6f74207472616e7366657220746f2073656c6600000000000000000000000000000000000000000000000000000000640000008000000000000000000000000000000000000000000000000000000020000000800000000000000000310ab089e4439a4c15d089f94afb7896ff553aecb10793d0ab882de59d99a32e000000000000000000000000000000000000000000000000ffffffffffffffff91eabfe8e493f369f48e58fdf2609ff8809506ce57440a6f25fddc25308a385192308bb7573b2a3d17ddb868b39d8ebec433f3194421abc22d084f89658c9bad599336d74a1247d50642b66dd6abeaa5484f6bd96b415b31bb99e26578c93978e46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000000000000000000000000000000000000000000020000000000000000000000000c820b10b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000240000000000000000000000008a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af4d7573742062652070726f706f736564206f776e6572000000000000000000008be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6884e487b710000000000000000000000000000000000000000000000000000000091eabfe8e493f369f48e58fdf2609ff8809506ce57440a6f25fddc25308a38520a675452746933cefe3d74182e78db7afe57ba60eaa4234b5d85e9aa41b0610c0200000200000000000000000000000000000044000000000000000000000000cbbf111300000000000000000000000000000000000000000000000000000000ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef45524332303a206d696e7420746f20746865207a65726f206164647265737300e2c8c9d500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffff7fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c161806aa1896bbf26568e884a7374b41e002500962caba6a15023a8d90e8508b830200000200000000000000000000000000000024000000000000000000000000a4c0ed360000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffe0616c616e6365000000000000000000000000000000000000000000000000000045524332303a207472616e7366657220616d6f756e74206578636565647320620000000000000000000000000000000000000084000000000000000000000000657373000000000000000000000000000000000000000000000000000000000045524332303a207472616e7366657220746f20746865207a65726f2061646472647265737300000000000000000000000000000000000000000000000000000045524332303a207472616e736665722066726f6d20746865207a65726f2061648c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92545524332303a20617070726f766520746f20746865207a65726f2061646472657373000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008400000080000000000000000045524332303a20617070726f76652066726f6d20746865207a65726f206164647265737300000000000000000000000000000000000000000000000000000000c2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b00000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff4000ae9fffffffffffffffffffffffffffffffffffffffffffffffffffffffff4000aea000000000000000000000000000000000000000000000000000000000e6599b4d0000000000000000000000000000000000000000000000000000000001ffc9a70000000000000000000000000000000000000000000000000000000036372b0700000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff207a65726f00000000000000000000000000000000000000000000000000000045524332303a2064656372656173656420616c6c6f77616e63652062656c6f77636500000000000000000000000000000000000000000000000000000000000045524332303a206275726e20616d6f756e7420657863656564732062616c616e730000000000000000000000000000000000000000000000000000000000000045524332303a206275726e2066726f6d20746865207a65726f2061646472657345524332303a20696e73756666696369656e7420616c6c6f77616e636500000002000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000eb5af27c73011fcd42467775e31007ee225265fad950ec80e4551fd770a5b3bc")

func (_LinkToken *LinkToken) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _LinkToken.abi.Events["Approval"].ID:
		return _LinkToken.ParseApproval(log)
	case _LinkToken.abi.Events["BurnAccessGranted"].ID:
		return _LinkToken.ParseBurnAccessGranted(log)
	case _LinkToken.abi.Events["BurnAccessRevoked"].ID:
		return _LinkToken.ParseBurnAccessRevoked(log)
	case _LinkToken.abi.Events["MintAccessGranted"].ID:
		return _LinkToken.ParseMintAccessGranted(log)
	case _LinkToken.abi.Events["MintAccessRevoked"].ID:
		return _LinkToken.ParseMintAccessRevoked(log)
	case _LinkToken.abi.Events["OwnershipTransferRequested"].ID:
		return _LinkToken.ParseOwnershipTransferRequested(log)
	case _LinkToken.abi.Events["OwnershipTransferred"].ID:
		return _LinkToken.ParseOwnershipTransferred(log)
	case _LinkToken.abi.Events["Transfer"].ID:
		return _LinkToken.ParseTransfer(log)
	case _LinkToken.abi.Events["Transfer0"].ID:
		return _LinkToken.ParseTransfer0(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (LinkTokenApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (LinkTokenBurnAccessGranted) Topic() common.Hash {
	return common.HexToHash("0x92308bb7573b2a3d17ddb868b39d8ebec433f3194421abc22d084f89658c9bad")
}

func (LinkTokenBurnAccessRevoked) Topic() common.Hash {
	return common.HexToHash("0x0a675452746933cefe3d74182e78db7afe57ba60eaa4234b5d85e9aa41b0610c")
}

func (LinkTokenMintAccessGranted) Topic() common.Hash {
	return common.HexToHash("0xe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea")
}

func (LinkTokenMintAccessRevoked) Topic() common.Hash {
	return common.HexToHash("0xed998b960f6340d045f620c119730f7aa7995e7425c2401d3a5b64ff998a59e9")
}

func (LinkTokenOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (LinkTokenOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (LinkTokenTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (LinkTokenTransfer0) Topic() common.Hash {
	return common.HexToHash("0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16")
}

func (_LinkToken *LinkToken) Address() common.Address {
	return _LinkToken.address
}

type CustomTransaction struct {
	*types.Transaction
	CustomHash common.Hash
}

func (tx *CustomTransaction) Hash() common.Hash {
	return tx.CustomHash
}

func ConvertToTransaction(resp zktypes.TransactionResponse) *CustomTransaction {
	dtx := &types.DynamicFeeTx{
		ChainID:   resp.ChainID.ToInt(),
		Nonce:     uint64(resp.Nonce),
		GasTipCap: resp.MaxPriorityFeePerGas.ToInt(),
		GasFeeCap: resp.MaxFeePerGas.ToInt(),
		To:        &resp.To,
		Value:     resp.Value.ToInt(),
		Data:      resp.Data,
		Gas:       uint64(resp.Gas),
	}

	tx := types.NewTx(dtx)
	customTransaction := CustomTransaction{Transaction: tx, CustomHash: resp.Hash}
	return &customTransaction
}

func DeployZkSyncLinkToken(auth *bind.TransactOpts, backend bind.ContractBackend, params ...interface{}) (common.Address, *CustomTransaction, *LinkToken, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	fmt.Println("Deploying zksync contract")
	zksyncClient := zkSyncClient.NewClient(client.Client())
	fmt.Println("getting wallet")
	wallet := auth.Context.Value("wallet").(*zkSyncAccounts.Wallet)
	fmt.Println("got wallet")
	fmt.Println("getting bytes")
	decodedBytes := common.FromHex(LinkTokenZkBin)
	fmt.Println("deploying")
	LinkTokenAbi, err := LinkTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	constructor, err := LinkTokenAbi.Pack("", params...)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	hash, err := wallet.DeployWithCreate(nil, zkSyncAccounts.CreateTransaction{
		Bytecode: decodedBytes,
		Calldata: constructor,
	})
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	fmt.Println("hash of tx", hash)
	receipt, err := zksyncClient.WaitMined(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	tx, _, err := zksyncClient.TransactionByHash(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	fmt.Println("tx hash", tx.Hash)

	ethTx := ConvertToTransaction(*tx)
	address := receipt.ContractAddress

	parsed, err := LinkTokenMetaData.GetAbi()
	contractBind := bind.NewBoundContract(address, *parsed, backend, backend, backend)

	contractReturn := &LinkToken{address: address, abi: *parsed, LinkTokenCaller: LinkTokenCaller{contract: contractBind}, LinkTokenTransactor: LinkTokenTransactor{contract: contractBind}, LinkTokenFilterer: LinkTokenFilterer{contract: contractBind}}

	return address, ethTx, contractReturn, err
}

type LinkTokenInterface interface {
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

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LinkTokenApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *LinkTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*LinkTokenApproval, error)

	FilterBurnAccessGranted(opts *bind.FilterOpts, burner []common.Address) (*LinkTokenBurnAccessGrantedIterator, error)

	WatchBurnAccessGranted(opts *bind.WatchOpts, sink chan<- *LinkTokenBurnAccessGranted, burner []common.Address) (event.Subscription, error)

	ParseBurnAccessGranted(log types.Log) (*LinkTokenBurnAccessGranted, error)

	FilterBurnAccessRevoked(opts *bind.FilterOpts, burner []common.Address) (*LinkTokenBurnAccessRevokedIterator, error)

	WatchBurnAccessRevoked(opts *bind.WatchOpts, sink chan<- *LinkTokenBurnAccessRevoked, burner []common.Address) (event.Subscription, error)

	ParseBurnAccessRevoked(log types.Log) (*LinkTokenBurnAccessRevoked, error)

	FilterMintAccessGranted(opts *bind.FilterOpts, minter []common.Address) (*LinkTokenMintAccessGrantedIterator, error)

	WatchMintAccessGranted(opts *bind.WatchOpts, sink chan<- *LinkTokenMintAccessGranted, minter []common.Address) (event.Subscription, error)

	ParseMintAccessGranted(log types.Log) (*LinkTokenMintAccessGranted, error)

	FilterMintAccessRevoked(opts *bind.FilterOpts, minter []common.Address) (*LinkTokenMintAccessRevokedIterator, error)

	WatchMintAccessRevoked(opts *bind.WatchOpts, sink chan<- *LinkTokenMintAccessRevoked, minter []common.Address) (event.Subscription, error)

	ParseMintAccessRevoked(log types.Log) (*LinkTokenMintAccessRevoked, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LinkTokenOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LinkTokenOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*LinkTokenOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LinkTokenOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LinkTokenOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*LinkTokenOwnershipTransferred, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LinkTokenTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *LinkTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*LinkTokenTransfer, error)

	FilterTransfer0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LinkTokenTransfer0Iterator, error)

	WatchTransfer0(opts *bind.WatchOpts, sink chan<- *LinkTokenTransfer0, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer0(log types.Log) (*LinkTokenTransfer0, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
