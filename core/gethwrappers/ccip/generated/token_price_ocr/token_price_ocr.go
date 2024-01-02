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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPriceRegistry\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"ForkedChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint224\",\"name\":\"usdPerToken\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"name\":\"priceUpdates\",\"type\":\"tuple[]\"}],\"internalType\":\"structTokenPriceOCR.Report\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"exposeForEncoding\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR2Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620022e3380380620022e383398101604081905262000034916200019d565b600133806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c081620000f2565b50505015156080524660a052600880546001600160a01b0319166001600160a01b0392909216919091179055620001cf565b336001600160a01b038216036200014c5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001b057600080fd5b81516001600160a01b0381168114620001c857600080fd5b9392505050565b60805160a0516120e7620001fc60003960008181610c2a0152610c7601526000610cf101526120e76000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806381ff704811610076578063afcb95d71161005b578063afcb95d7146101b2578063b1dc65a4146101d2578063f2fde38b146101e557600080fd5b806381ff70481461015a5780638da5cb5b1461018a57600080fd5b806348dab751116100a757806348dab7511461012a578063666cab8d1461013d57806379ba50971461015257600080fd5b8063181f5a77146100c35780631ef3817414610115575b600080fd5b6100ff6040518060400160405280601381526020017f546f6b656e50726963654f435220312e302e300000000000000000000000000081525081565b60405161010c91906115bb565b60405180910390f35b6101286101233660046117f2565b6101f8565b005b6100ff610138366004611978565b610a10565b610145610a39565b60405161010c9190611a4c565b610128610aa8565b6004546002546040805163ffffffff8085168252640100000000909404909316602084015282015260600161010c565b60005460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010c565b60408051600181526000602082018190529181019190915260600161010c565b6101286101e0366004611aab565b610ba5565b6101286101f3366004611b90565b6111c5565b855185518560ff16601f831115610270576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064015b60405180910390fd5b806000036102da576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610267565b818314610368576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610267565b610373816003611bda565b83116103db576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610267565b6103e36111d9565b60065460005b818110156104df57600560006006838154811061040857610408611bf7565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001690556007805460059291908490811061047857610478611bf7565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001690556104d881611c26565b90506103e9565b50895160005b818110156108b85760008c828151811061050157610501611bf7565b602002602001015190506000600281111561051e5761051e611c5e565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260056020526040902054610100900460ff16600281111561055d5761055d611c5e565b146105c4576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610267565b73ffffffffffffffffffffffffffffffffffffffff8116610611576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260ff83168152602081016001905273ffffffffffffffffffffffffffffffffffffffff821660009081526005602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156106c1576106c1611c5e565b021790555090505060008c83815181106106dd576106dd611bf7565b60200260200101519050600060028111156106fa576106fa611c5e565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260056020526040902054610100900460ff16600281111561073957610739611c5e565b146107a0576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610267565b73ffffffffffffffffffffffffffffffffffffffff81166107ed576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260ff84168152602081016002905273ffffffffffffffffffffffffffffffffffffffff821660009081526005602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561089d5761089d611c5e565b02179055509050505050806108b190611c26565b90506104e5565b508a516108cc9060069060208e0190611499565b5089516108e09060079060208d0190611499565b506003805460ff838116610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000909216908c1617179055600480546109669146913091906000906109389063ffffffff16611c8d565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168e8e8e8e8e8e61125c565b600260000181905550600060048054906101000a900463ffffffff169050436004806101000a81548163ffffffff021916908363ffffffff1602179055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0581600260000154600460009054906101000a900463ffffffff168f8f8f8f8f8f6040516109fa99989796959493929190611cb0565b60405180910390a1505050505050505050505050565b606081604051602001610a239190611db5565b6040516020818303038152906040529050919050565b60606007805480602002602001604051908101604052809291908181526020018280548015610a9e57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610a73575b5050505050905090565b60015473ffffffffffffffffffffffffffffffffffffffff163314610b29576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610267565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610bb4878760208b0135611307565b6040805160608101825260025480825260035460ff808216602085015261010090910416928201929092528935918214610c275780516040517f93df584c000000000000000000000000000000000000000000000000000000008152600481019190915260248101839052604401610267565b467f000000000000000000000000000000000000000000000000000000000000000014610ca8576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006004820152466024820152604401610267565b6040805183815260208c81013560081c63ffffffff16908201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160007f000000000000000000000000000000000000000000000000000000000000000015610d4a57600282602001518360400151610d2b9190611dd8565b610d359190611df1565b610d40906001611dd8565b60ff169050610d60565b6020820151610d5a906001611dd8565b60ff1690505b868114610d99576040517f71253a2500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b868514610dd2576040517fa75d88af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526005602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115610e1557610e15611c5e565b6002811115610e2657610e26611c5e565b9052509050600281602001516002811115610e4357610e43611c5e565b148015610e8a57506007816000015160ff1681548110610e6557610e65611bf7565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b610ec0576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506000610ece866020611bda565b610ed9896020611bda565b610ee58c610144611e3a565b610eef9190611e3a565b610ef99190611e3a565b9050368114610f3d576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401610267565b5060008a8a604051610f50929190611e4d565b604051908190038120610f67918e90602001611e5d565b604051602081830303815290604052805190602001209050610f87611523565b8860005b818110156111b45760006001858a8460208110610faa57610faa611bf7565b610fb791901a601b611dd8565b8f8f86818110610fc957610fc9611bf7565b905060200201358e8e87818110610fe257610fe2611bf7565b905060200201356040516000815260200160405260405161101f949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611041573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff8116600090815260056020908152848220848601909552845460ff80821686529397509195509293928401916101009091041660028111156110c4576110c4611c5e565b60028111156110d5576110d5611c5e565b90525090506001816020015160028111156110f2576110f2611c5e565b14611129576040517fca31867a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051859060ff16601f811061114057611140611bf7565b60200201511561117c576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600185826000015160ff16601f811061119757611197611bf7565b91151560209092020152506111ad905081611c26565b9050610f8b565b505050505050505050505050505050565b6111cd6111d9565b6111d6816113a4565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461125a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610267565b565b6000808a8a8a8a8a8a8a8a8a60405160200161128099989796959493929190611e71565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b600061131583850185611f06565b6008546040517f3937306f00000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff1690633937306f9061136c908490600401612022565b600060405180830381600087803b15801561138657600080fd5b505af115801561139a573d6000803e3d6000fd5b5050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603611423576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610267565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215611513579160200282015b8281111561151357825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9091161782556020909201916001909101906114b9565b5061151f929150611542565b5090565b604051806103e00160405280601f906020820280368337509192915050565b5b8082111561151f5760008155600101611543565b6000815180845260005b8181101561157d57602081850181015186830182015201611561565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006115ce6020830184611557565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715611627576116276115d5565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611674576116746115d5565b604052919050565b600067ffffffffffffffff821115611696576116966115d5565b5060051b60200190565b803573ffffffffffffffffffffffffffffffffffffffff811681146116c457600080fd5b919050565b600082601f8301126116da57600080fd5b813560206116ef6116ea8361167c565b61162d565b82815260059290921b8401810191818101908684111561170e57600080fd5b8286015b8481101561173057611723816116a0565b8352918301918301611712565b509695505050505050565b803560ff811681146116c457600080fd5b600082601f83011261175d57600080fd5b813567ffffffffffffffff811115611777576117776115d5565b6117a860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161162d565b8181528460208386010111156117bd57600080fd5b816020850160208301376000918101602001919091529392505050565b803567ffffffffffffffff811681146116c457600080fd5b60008060008060008060c0878903121561180b57600080fd5b863567ffffffffffffffff8082111561182357600080fd5b61182f8a838b016116c9565b9750602089013591508082111561184557600080fd5b6118518a838b016116c9565b965061185f60408a0161173b565b9550606089013591508082111561187557600080fd5b6118818a838b0161174c565b945061188f60808a016117da565b935060a08901359150808211156118a557600080fd5b506118b289828a0161174c565b9150509295509295509295565b80357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811681146116c457600080fd5b600082601f8301126118fc57600080fd5b8135602061190c6116ea8361167c565b82815260069290921b8401810191818101908684111561192b57600080fd5b8286015b8481101561173057604081890312156119485760008081fd5b611950611604565b611959826116a0565b81526119668583016118bf565b8186015283529183019160400161192f565b60006020828403121561198a57600080fd5b813567ffffffffffffffff808211156119a257600080fd5b90830190602082860312156119b657600080fd5b6040516020810181811083821117156119d1576119d16115d5565b6040528235828111156119e357600080fd5b6119ef878286016118eb565b82525095945050505050565b600081518084526020808501945080840160005b83811015611a4157815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611a0f565b509495945050505050565b6020815260006115ce60208301846119fb565b60008083601f840112611a7157600080fd5b50813567ffffffffffffffff811115611a8957600080fd5b6020830191508360208260051b8501011115611aa457600080fd5b9250929050565b60008060008060008060008060e0898b031215611ac757600080fd5b606089018a811115611ad857600080fd5b8998503567ffffffffffffffff80821115611af257600080fd5b818b0191508b601f830112611b0657600080fd5b813581811115611b1557600080fd5b8c6020828501011115611b2757600080fd5b6020830199508098505060808b0135915080821115611b4557600080fd5b611b518c838d01611a5f565b909750955060a08b0135915080821115611b6a57600080fd5b50611b778b828c01611a5f565b999c989b50969995989497949560c00135949350505050565b600060208284031215611ba257600080fd5b6115ce826116a0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082028115828204841417611bf157611bf1611bab565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c5757611c57611bab565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600063ffffffff808316818103611ca657611ca6611bab565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152611ce08184018a6119fb565b90508281036080840152611cf481896119fb565b905060ff871660a084015282810360c0840152611d118187611557565b905067ffffffffffffffff851660e0840152828103610100840152611d368185611557565b9c9b505050505050505050505050565b600081518084526020808501945080840160005b83811015611a41578151805173ffffffffffffffffffffffffffffffffffffffff1688528301517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168388015260409096019590820190600101611d5a565b6020815260008251602080840152611dd06040840182611d46565b949350505050565b60ff8181168382160190811115611bf157611bf1611bab565b600060ff831680611e2b577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8060ff84160491505092915050565b80820180821115611bf157611bf1611bab565b8183823760009101908152919050565b828152606082602083013760800192915050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b166040850152816060850152611eb88285018b6119fb565b91508382036080850152611ecc828a6119fb565b915060ff881660a085015283820360c0850152611ee98288611557565b90861660e08501528381036101008501529050611d368185611557565b60006020808385031215611f1957600080fd5b823567ffffffffffffffff80821115611f3157600080fd5b81850191506040808388031215611f4757600080fd5b611f4f611604565b833583811115611f5e57600080fd5b611f6a898287016118eb565b8252508484013583811115611f7e57600080fd5b80850194505087601f850112611f9357600080fd5b83359250611fa36116ea8461167c565b83815260069390931b84018501928581019089851115611fc257600080fd5b948601945b848610156120105783868b031215611fdf5760008081fd5b611fe7611604565b611ff0876117da565b8152611ffd8888016118bf565b8189015282529483019490860190611fc7565b95820195909552979650505050505050565b6000602080835283516040808386015261203f6060860183611d46565b868401518682037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00183880152805180835290850193506000918501905b808310156120ce578451805167ffffffffffffffff1683528601517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168683015293850193600192909201919083019061207d565b5097965050505050505056fea164736f6c6343000813000a",
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
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

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

func (_TokenPriceOCR *TokenPriceOCRTransactor) SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _TokenPriceOCR.contract.Transact(opts, "setOCR2Config", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_TokenPriceOCR *TokenPriceOCRSession) SetOCR2Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.SetOCR2Config(&_TokenPriceOCR.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_TokenPriceOCR *TokenPriceOCRTransactorSession) SetOCR2Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _TokenPriceOCR.Contract.SetOCR2Config(&_TokenPriceOCR.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
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
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
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

type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}
type LatestConfigDigestAndEpoch struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
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
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (_TokenPriceOCR *TokenPriceOCR) Address() common.Address {
	return _TokenPriceOCR.address
}

type TokenPriceOCRInterface interface {
	ExposeForEncoding(opts *bind.CallOpts, report TokenPriceOCRReport) ([]byte, error)

	GetTransmitters(opts *bind.CallOpts) ([]common.Address, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

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
