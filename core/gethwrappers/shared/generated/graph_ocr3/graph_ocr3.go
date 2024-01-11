// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package graph_ocr3

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

type GraphOCR3Neighbor struct {
	ChainId         *big.Int
	ContractAddress common.Address
}

var GraphOCR3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"ForkedChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"latestSequenceNumber\",\"type\":\"uint64\"}],\"name\":\"NonIncreasingSequenceNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"NeighborAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"addNeighbor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNeighbors\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structGraphOCR3.Neighbor[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR3Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b503380600081620000695760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156200009c576200009c81620000a9565b5050466080525062000154565b336001600160a01b03821603620001035760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000060565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b608051611fab6200017760003960008181610d6f0152610dbb0152611fab6000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806381ff704811610081578063b1dc65a41161005b578063b1dc65a4146101f8578063c96682b61461020b578063f2fde38b1461021e57600080fd5b806381ff7048146101675780638da5cb5b14610197578063afcb95d7146101bf57600080fd5b8063666cab8d116100b2578063666cab8d146101355780636a11ee901461014a57806379ba50971461015f57600080fd5b8063181f5a77146100ce5780634af4fdd314610120575b600080fd5b61010a6040518060400160405280600f81526020017f47726170684f43523320312e302e30000000000000000000000000000000000081525081565b6040516101179190611811565b60405180910390f35b610128610231565b604051610117919061182b565b61013d6102b3565b60405161011791906118e1565b61015d610158366004611ad9565b610322565b005b61015d610b3a565b6004546002546040805163ffffffff80851682526401000000009094049093166020840152820152606001610117565b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610117565b600254600454604080516001815260208101939093526801000000000000000090910467ffffffffffffffff1690820152606001610117565b61015d610206366004611bf2565b610c37565b61015d610219366004611cd7565b6112a6565b61015d61022c366004611d03565b6114b8565b60606008805480602002602001604051908101604052809291908181526020016000905b828210156102aa57600084815260209081902060408051808201909152600285029091018054825260019081015473ffffffffffffffffffffffffffffffffffffffff16828401529083529092019101610255565b50505050905090565b6060600780548060200260200160405190810160405280929190818152602001828054801561031857602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116102ed575b5050505050905090565b855185518560ff16601f83111561039a576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064015b60405180910390fd5b80600003610404576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610391565b818314610492576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610391565b61049d816003611d4d565b8311610505576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610391565b61050d6114cc565b60065460005b8181101561060957600560006006838154811061053257610532611d6a565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000169055600780546005929190849081106105a2576105a2611d6a565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905561060281611d99565b9050610513565b50895160005b818110156109e25760008c828151811061062b5761062b611d6a565b602002602001015190506000600281111561064857610648611dd1565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260056020526040902054610100900460ff16600281111561068757610687611dd1565b146106ee576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610391565b73ffffffffffffffffffffffffffffffffffffffff811661073b576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260ff83168152602081016001905273ffffffffffffffffffffffffffffffffffffffff821660009081526005602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156107eb576107eb611dd1565b021790555090505060008c838151811061080757610807611d6a565b602002602001015190506000600281111561082457610824611dd1565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260056020526040902054610100900460ff16600281111561086357610863611dd1565b146108ca576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610391565b73ffffffffffffffffffffffffffffffffffffffff8116610917576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260ff84168152602081016002905273ffffffffffffffffffffffffffffffffffffffff821660009081526005602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156109c7576109c7611dd1565b02179055509050505050806109db90611d99565b905061060f565b508a516109f69060069060208e01906116ef565b508951610a0a9060079060208d01906116ef565b506003805460ff838116610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000909216908c161717905560048054610a90914691309190600090610a629063ffffffff16611e00565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168e8e8e8e8e8e61154f565b600260000181905550600060048054906101000a900463ffffffff169050436004806101000a81548163ffffffff021916908363ffffffff1602179055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0581600260000154600460009054906101000a900463ffffffff168f8f8f8f8f8f604051610b2499989796959493929190611e23565b60405180910390a1505050505050505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610bbb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610391565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60045460208901359067ffffffffffffffff68010000000000000000909104811690821611610cba57600480546040517f6e376b6600000000000000000000000000000000000000000000000000000000815267ffffffffffffffff80851693820193909352680100000000000000009091049091166024820152604401610391565b600480547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8416021790556040805160608101825260025480825260035460ff808216602085015261010090910416928201929092528a35918214610d6c5780516040517f93df584c000000000000000000000000000000000000000000000000000000008152600481019190915260248101839052604401610391565b467f000000000000000000000000000000000000000000000000000000000000000014610ded576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006004820152466024820152604401610391565b6040805183815267ffffffffffffffff851660208201527fe893c2681d327421d89e1cb54fbe64645b4dcea668d6826130b62cf4c6eefea2910160405180910390a16020810151610e3f906001611eb9565b60ff168714610e7a576040517f71253a2500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b868514610eb3576040517fa75d88af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526005602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115610ef657610ef6611dd1565b6002811115610f0757610f07611dd1565b9052509050600281602001516002811115610f2457610f24611dd1565b148015610f6b57506007816000015160ff1681548110610f4657610f46611d6a565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b610fa1576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506000610faf866020611d4d565b610fba896020611d4d565b610fc68c610144611ed2565b610fd09190611ed2565b610fda9190611ed2565b905036811461101e576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401610391565b5060008a8a604051611031929190611ee5565b604051908190038120611048918e90602001611ef5565b604051602081830303815290604052805190602001209050611068611779565b8860005b818110156112955760006001858a846020811061108b5761108b611d6a565b61109891901a601b611eb9565b8f8f868181106110aa576110aa611d6a565b905060200201358e8e878181106110c3576110c3611d6a565b9050602002013560405160008152602001604052604051611100949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611122573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff8116600090815260056020908152848220848601909552845460ff80821686529397509195509293928401916101009091041660028111156111a5576111a5611dd1565b60028111156111b6576111b6611dd1565b90525090506001816020015160028111156111d3576111d3611dd1565b1461120a576040517fca31867a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051859060ff16601f811061122157611221611d6a565b60200201511561125d576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600185826000015160ff16601f811061127857611278611d6a565b911515602090920201525061128e905081611d99565b905061106c565b505050505050505050505050505050565b600082826040516020016112da92919091825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600990935291205490915060ff161561138a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f616c7265616479206164646564000000000000000000000000000000000000006044820152606401610391565b60408051808201825284815273ffffffffffffffffffffffffffffffffffffffff84811660208084018281526008805460018082018355600092835296517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee360029092029182015591517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee490920180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169290951691909117909355858352600981529184902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558251868152908101919091527fbbfdc018e8b91f5b758da7f1c7fdc525ce2fa1e37315698b8b6081718a5c2c81910160405180910390a1505050565b6114c06114cc565b6114c9816115fa565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461154d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610391565b565b6000808a8a8a8a8a8a8a8a8a60405160200161157399989796959493929190611f09565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603611679576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610391565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215611769579160200282015b8281111561176957825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90911617825560209092019160019091019061170f565b50611775929150611798565b5090565b604051806103e00160405280601f906020820280368337509192915050565b5b808211156117755760008155600101611799565b6000815180845260005b818110156117d3576020818501810151868301820152016117b7565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061182460208301846117ad565b9392505050565b602080825282518282018190526000919060409081850190868401855b828110156118835781518051855286015173ffffffffffffffffffffffffffffffffffffffff16868501529284019290850190600101611848565b5091979650505050505050565b600081518084526020808501945080840160005b838110156118d657815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016118a4565b509495945050505050565b6020815260006118246020830184611890565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561196a5761196a6118f4565b604052919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461199657600080fd5b919050565b600082601f8301126119ac57600080fd5b8135602067ffffffffffffffff8211156119c8576119c86118f4565b8160051b6119d7828201611923565b92835284810182019282810190878511156119f157600080fd5b83870192505b84831015611a1757611a0883611972565b825291830191908301906119f7565b979650505050505050565b803560ff8116811461199657600080fd5b600082601f830112611a4457600080fd5b813567ffffffffffffffff811115611a5e57611a5e6118f4565b611a8f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611923565b818152846020838601011115611aa457600080fd5b816020850160208301376000918101602001919091529392505050565b803567ffffffffffffffff8116811461199657600080fd5b60008060008060008060c08789031215611af257600080fd5b863567ffffffffffffffff80821115611b0a57600080fd5b611b168a838b0161199b565b97506020890135915080821115611b2c57600080fd5b611b388a838b0161199b565b9650611b4660408a01611a22565b95506060890135915080821115611b5c57600080fd5b611b688a838b01611a33565b9450611b7660808a01611ac1565b935060a0890135915080821115611b8c57600080fd5b50611b9989828a01611a33565b9150509295509295509295565b60008083601f840112611bb857600080fd5b50813567ffffffffffffffff811115611bd057600080fd5b6020830191508360208260051b8501011115611beb57600080fd5b9250929050565b60008060008060008060008060e0898b031215611c0e57600080fd5b606089018a811115611c1f57600080fd5b8998503567ffffffffffffffff80821115611c3957600080fd5b818b0191508b601f830112611c4d57600080fd5b813581811115611c5c57600080fd5b8c6020828501011115611c6e57600080fd5b6020830199508098505060808b0135915080821115611c8c57600080fd5b611c988c838d01611ba6565b909750955060a08b0135915080821115611cb157600080fd5b50611cbe8b828c01611ba6565b999c989b50969995989497949560c00135949350505050565b60008060408385031215611cea57600080fd5b82359150611cfa60208401611972565b90509250929050565b600060208284031215611d1557600080fd5b61182482611972565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082028115828204841417611d6457611d64611d1e565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611dca57611dca611d1e565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600063ffffffff808316818103611e1957611e19611d1e565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152611e538184018a611890565b90508281036080840152611e678189611890565b905060ff871660a084015282810360c0840152611e8481876117ad565b905067ffffffffffffffff851660e0840152828103610100840152611ea981856117ad565b9c9b505050505050505050505050565b60ff8181168382160190811115611d6457611d64611d1e565b80820180821115611d6457611d64611d1e565b8183823760009101908152919050565b828152606082602083013760800192915050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b166040850152816060850152611f508285018b611890565b91508382036080850152611f64828a611890565b915060ff881660a085015283820360c0850152611f8182886117ad565b90861660e08501528381036101008501529050611ea981856117ad56fea164736f6c6343000813000a",
}

var GraphOCR3ABI = GraphOCR3MetaData.ABI

var GraphOCR3Bin = GraphOCR3MetaData.Bin

func DeployGraphOCR3(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GraphOCR3, error) {
	parsed, err := GraphOCR3MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GraphOCR3Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GraphOCR3{address: address, abi: *parsed, GraphOCR3Caller: GraphOCR3Caller{contract: contract}, GraphOCR3Transactor: GraphOCR3Transactor{contract: contract}, GraphOCR3Filterer: GraphOCR3Filterer{contract: contract}}, nil
}

type GraphOCR3 struct {
	address common.Address
	abi     abi.ABI
	GraphOCR3Caller
	GraphOCR3Transactor
	GraphOCR3Filterer
}

type GraphOCR3Caller struct {
	contract *bind.BoundContract
}

type GraphOCR3Transactor struct {
	contract *bind.BoundContract
}

type GraphOCR3Filterer struct {
	contract *bind.BoundContract
}

type GraphOCR3Session struct {
	Contract     *GraphOCR3
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type GraphOCR3CallerSession struct {
	Contract *GraphOCR3Caller
	CallOpts bind.CallOpts
}

type GraphOCR3TransactorSession struct {
	Contract     *GraphOCR3Transactor
	TransactOpts bind.TransactOpts
}

type GraphOCR3Raw struct {
	Contract *GraphOCR3
}

type GraphOCR3CallerRaw struct {
	Contract *GraphOCR3Caller
}

type GraphOCR3TransactorRaw struct {
	Contract *GraphOCR3Transactor
}

func NewGraphOCR3(address common.Address, backend bind.ContractBackend) (*GraphOCR3, error) {
	abi, err := abi.JSON(strings.NewReader(GraphOCR3ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindGraphOCR3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GraphOCR3{address: address, abi: abi, GraphOCR3Caller: GraphOCR3Caller{contract: contract}, GraphOCR3Transactor: GraphOCR3Transactor{contract: contract}, GraphOCR3Filterer: GraphOCR3Filterer{contract: contract}}, nil
}

func NewGraphOCR3Caller(address common.Address, caller bind.ContractCaller) (*GraphOCR3Caller, error) {
	contract, err := bindGraphOCR3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GraphOCR3Caller{contract: contract}, nil
}

func NewGraphOCR3Transactor(address common.Address, transactor bind.ContractTransactor) (*GraphOCR3Transactor, error) {
	contract, err := bindGraphOCR3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GraphOCR3Transactor{contract: contract}, nil
}

func NewGraphOCR3Filterer(address common.Address, filterer bind.ContractFilterer) (*GraphOCR3Filterer, error) {
	contract, err := bindGraphOCR3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GraphOCR3Filterer{contract: contract}, nil
}

func bindGraphOCR3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GraphOCR3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_GraphOCR3 *GraphOCR3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GraphOCR3.Contract.GraphOCR3Caller.contract.Call(opts, result, method, params...)
}

func (_GraphOCR3 *GraphOCR3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GraphOCR3.Contract.GraphOCR3Transactor.contract.Transfer(opts)
}

func (_GraphOCR3 *GraphOCR3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GraphOCR3.Contract.GraphOCR3Transactor.contract.Transact(opts, method, params...)
}

func (_GraphOCR3 *GraphOCR3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GraphOCR3.Contract.contract.Call(opts, result, method, params...)
}

func (_GraphOCR3 *GraphOCR3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GraphOCR3.Contract.contract.Transfer(opts)
}

func (_GraphOCR3 *GraphOCR3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GraphOCR3.Contract.contract.Transact(opts, method, params...)
}

func (_GraphOCR3 *GraphOCR3Caller) GetNeighbors(opts *bind.CallOpts) ([]GraphOCR3Neighbor, error) {
	var out []interface{}
	err := _GraphOCR3.contract.Call(opts, &out, "getNeighbors")

	if err != nil {
		return *new([]GraphOCR3Neighbor), err
	}

	out0 := *abi.ConvertType(out[0], new([]GraphOCR3Neighbor)).(*[]GraphOCR3Neighbor)

	return out0, err

}

func (_GraphOCR3 *GraphOCR3Session) GetNeighbors() ([]GraphOCR3Neighbor, error) {
	return _GraphOCR3.Contract.GetNeighbors(&_GraphOCR3.CallOpts)
}

func (_GraphOCR3 *GraphOCR3CallerSession) GetNeighbors() ([]GraphOCR3Neighbor, error) {
	return _GraphOCR3.Contract.GetNeighbors(&_GraphOCR3.CallOpts)
}

func (_GraphOCR3 *GraphOCR3Caller) GetTransmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GraphOCR3.contract.Call(opts, &out, "getTransmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_GraphOCR3 *GraphOCR3Session) GetTransmitters() ([]common.Address, error) {
	return _GraphOCR3.Contract.GetTransmitters(&_GraphOCR3.CallOpts)
}

func (_GraphOCR3 *GraphOCR3CallerSession) GetTransmitters() ([]common.Address, error) {
	return _GraphOCR3.Contract.GetTransmitters(&_GraphOCR3.CallOpts)
}

func (_GraphOCR3 *GraphOCR3Caller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _GraphOCR3.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_GraphOCR3 *GraphOCR3Session) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _GraphOCR3.Contract.LatestConfigDetails(&_GraphOCR3.CallOpts)
}

func (_GraphOCR3 *GraphOCR3CallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _GraphOCR3.Contract.LatestConfigDetails(&_GraphOCR3.CallOpts)
}

func (_GraphOCR3 *GraphOCR3Caller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _GraphOCR3.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.SequenceNumber = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

func (_GraphOCR3 *GraphOCR3Session) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _GraphOCR3.Contract.LatestConfigDigestAndEpoch(&_GraphOCR3.CallOpts)
}

func (_GraphOCR3 *GraphOCR3CallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _GraphOCR3.Contract.LatestConfigDigestAndEpoch(&_GraphOCR3.CallOpts)
}

func (_GraphOCR3 *GraphOCR3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GraphOCR3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_GraphOCR3 *GraphOCR3Session) Owner() (common.Address, error) {
	return _GraphOCR3.Contract.Owner(&_GraphOCR3.CallOpts)
}

func (_GraphOCR3 *GraphOCR3CallerSession) Owner() (common.Address, error) {
	return _GraphOCR3.Contract.Owner(&_GraphOCR3.CallOpts)
}

func (_GraphOCR3 *GraphOCR3Caller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GraphOCR3.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_GraphOCR3 *GraphOCR3Session) TypeAndVersion() (string, error) {
	return _GraphOCR3.Contract.TypeAndVersion(&_GraphOCR3.CallOpts)
}

func (_GraphOCR3 *GraphOCR3CallerSession) TypeAndVersion() (string, error) {
	return _GraphOCR3.Contract.TypeAndVersion(&_GraphOCR3.CallOpts)
}

func (_GraphOCR3 *GraphOCR3Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GraphOCR3.contract.Transact(opts, "acceptOwnership")
}

func (_GraphOCR3 *GraphOCR3Session) AcceptOwnership() (*types.Transaction, error) {
	return _GraphOCR3.Contract.AcceptOwnership(&_GraphOCR3.TransactOpts)
}

func (_GraphOCR3 *GraphOCR3TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _GraphOCR3.Contract.AcceptOwnership(&_GraphOCR3.TransactOpts)
}

func (_GraphOCR3 *GraphOCR3Transactor) AddNeighbor(opts *bind.TransactOpts, chainId *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _GraphOCR3.contract.Transact(opts, "addNeighbor", chainId, contractAddress)
}

func (_GraphOCR3 *GraphOCR3Session) AddNeighbor(chainId *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _GraphOCR3.Contract.AddNeighbor(&_GraphOCR3.TransactOpts, chainId, contractAddress)
}

func (_GraphOCR3 *GraphOCR3TransactorSession) AddNeighbor(chainId *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _GraphOCR3.Contract.AddNeighbor(&_GraphOCR3.TransactOpts, chainId, contractAddress)
}

func (_GraphOCR3 *GraphOCR3Transactor) SetOCR3Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _GraphOCR3.contract.Transact(opts, "setOCR3Config", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_GraphOCR3 *GraphOCR3Session) SetOCR3Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _GraphOCR3.Contract.SetOCR3Config(&_GraphOCR3.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_GraphOCR3 *GraphOCR3TransactorSession) SetOCR3Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _GraphOCR3.Contract.SetOCR3Config(&_GraphOCR3.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_GraphOCR3 *GraphOCR3Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _GraphOCR3.contract.Transact(opts, "transferOwnership", to)
}

func (_GraphOCR3 *GraphOCR3Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _GraphOCR3.Contract.TransferOwnership(&_GraphOCR3.TransactOpts, to)
}

func (_GraphOCR3 *GraphOCR3TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _GraphOCR3.Contract.TransferOwnership(&_GraphOCR3.TransactOpts, to)
}

func (_GraphOCR3 *GraphOCR3Transactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _GraphOCR3.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_GraphOCR3 *GraphOCR3Session) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _GraphOCR3.Contract.Transmit(&_GraphOCR3.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_GraphOCR3 *GraphOCR3TransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _GraphOCR3.Contract.Transmit(&_GraphOCR3.TransactOpts, reportContext, report, rs, ss, rawVs)
}

type GraphOCR3ConfigSetIterator struct {
	Event *GraphOCR3ConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GraphOCR3ConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphOCR3ConfigSet)
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
		it.Event = new(GraphOCR3ConfigSet)
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

func (it *GraphOCR3ConfigSetIterator) Error() error {
	return it.fail
}

func (it *GraphOCR3ConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GraphOCR3ConfigSet struct {
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

func (_GraphOCR3 *GraphOCR3Filterer) FilterConfigSet(opts *bind.FilterOpts) (*GraphOCR3ConfigSetIterator, error) {

	logs, sub, err := _GraphOCR3.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &GraphOCR3ConfigSetIterator{contract: _GraphOCR3.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_GraphOCR3 *GraphOCR3Filterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *GraphOCR3ConfigSet) (event.Subscription, error) {

	logs, sub, err := _GraphOCR3.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GraphOCR3ConfigSet)
				if err := _GraphOCR3.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_GraphOCR3 *GraphOCR3Filterer) ParseConfigSet(log types.Log) (*GraphOCR3ConfigSet, error) {
	event := new(GraphOCR3ConfigSet)
	if err := _GraphOCR3.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GraphOCR3NeighborAddedIterator struct {
	Event *GraphOCR3NeighborAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GraphOCR3NeighborAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphOCR3NeighborAdded)
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
		it.Event = new(GraphOCR3NeighborAdded)
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

func (it *GraphOCR3NeighborAddedIterator) Error() error {
	return it.fail
}

func (it *GraphOCR3NeighborAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GraphOCR3NeighborAdded struct {
	ChainId         *big.Int
	ContractAddress common.Address
	Raw             types.Log
}

func (_GraphOCR3 *GraphOCR3Filterer) FilterNeighborAdded(opts *bind.FilterOpts) (*GraphOCR3NeighborAddedIterator, error) {

	logs, sub, err := _GraphOCR3.contract.FilterLogs(opts, "NeighborAdded")
	if err != nil {
		return nil, err
	}
	return &GraphOCR3NeighborAddedIterator{contract: _GraphOCR3.contract, event: "NeighborAdded", logs: logs, sub: sub}, nil
}

func (_GraphOCR3 *GraphOCR3Filterer) WatchNeighborAdded(opts *bind.WatchOpts, sink chan<- *GraphOCR3NeighborAdded) (event.Subscription, error) {

	logs, sub, err := _GraphOCR3.contract.WatchLogs(opts, "NeighborAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GraphOCR3NeighborAdded)
				if err := _GraphOCR3.contract.UnpackLog(event, "NeighborAdded", log); err != nil {
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

func (_GraphOCR3 *GraphOCR3Filterer) ParseNeighborAdded(log types.Log) (*GraphOCR3NeighborAdded, error) {
	event := new(GraphOCR3NeighborAdded)
	if err := _GraphOCR3.contract.UnpackLog(event, "NeighborAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GraphOCR3OwnershipTransferRequestedIterator struct {
	Event *GraphOCR3OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GraphOCR3OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphOCR3OwnershipTransferRequested)
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
		it.Event = new(GraphOCR3OwnershipTransferRequested)
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

func (it *GraphOCR3OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *GraphOCR3OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GraphOCR3OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_GraphOCR3 *GraphOCR3Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GraphOCR3OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GraphOCR3.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GraphOCR3OwnershipTransferRequestedIterator{contract: _GraphOCR3.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_GraphOCR3 *GraphOCR3Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *GraphOCR3OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GraphOCR3.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GraphOCR3OwnershipTransferRequested)
				if err := _GraphOCR3.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_GraphOCR3 *GraphOCR3Filterer) ParseOwnershipTransferRequested(log types.Log) (*GraphOCR3OwnershipTransferRequested, error) {
	event := new(GraphOCR3OwnershipTransferRequested)
	if err := _GraphOCR3.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GraphOCR3OwnershipTransferredIterator struct {
	Event *GraphOCR3OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GraphOCR3OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphOCR3OwnershipTransferred)
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
		it.Event = new(GraphOCR3OwnershipTransferred)
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

func (it *GraphOCR3OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *GraphOCR3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GraphOCR3OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_GraphOCR3 *GraphOCR3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GraphOCR3OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GraphOCR3.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GraphOCR3OwnershipTransferredIterator{contract: _GraphOCR3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_GraphOCR3 *GraphOCR3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GraphOCR3OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GraphOCR3.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GraphOCR3OwnershipTransferred)
				if err := _GraphOCR3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_GraphOCR3 *GraphOCR3Filterer) ParseOwnershipTransferred(log types.Log) (*GraphOCR3OwnershipTransferred, error) {
	event := new(GraphOCR3OwnershipTransferred)
	if err := _GraphOCR3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GraphOCR3TransmittedIterator struct {
	Event *GraphOCR3Transmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GraphOCR3TransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphOCR3Transmitted)
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
		it.Event = new(GraphOCR3Transmitted)
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

func (it *GraphOCR3TransmittedIterator) Error() error {
	return it.fail
}

func (it *GraphOCR3TransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GraphOCR3Transmitted struct {
	ConfigDigest   [32]byte
	SequenceNumber uint64
	Raw            types.Log
}

func (_GraphOCR3 *GraphOCR3Filterer) FilterTransmitted(opts *bind.FilterOpts) (*GraphOCR3TransmittedIterator, error) {

	logs, sub, err := _GraphOCR3.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &GraphOCR3TransmittedIterator{contract: _GraphOCR3.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_GraphOCR3 *GraphOCR3Filterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *GraphOCR3Transmitted) (event.Subscription, error) {

	logs, sub, err := _GraphOCR3.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GraphOCR3Transmitted)
				if err := _GraphOCR3.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_GraphOCR3 *GraphOCR3Filterer) ParseTransmitted(log types.Log) (*GraphOCR3Transmitted, error) {
	event := new(GraphOCR3Transmitted)
	if err := _GraphOCR3.contract.UnpackLog(event, "Transmitted", log); err != nil {
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
	ScanLogs       bool
	ConfigDigest   [32]byte
	SequenceNumber uint64
}

func (_GraphOCR3 *GraphOCR3) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _GraphOCR3.abi.Events["ConfigSet"].ID:
		return _GraphOCR3.ParseConfigSet(log)
	case _GraphOCR3.abi.Events["NeighborAdded"].ID:
		return _GraphOCR3.ParseNeighborAdded(log)
	case _GraphOCR3.abi.Events["OwnershipTransferRequested"].ID:
		return _GraphOCR3.ParseOwnershipTransferRequested(log)
	case _GraphOCR3.abi.Events["OwnershipTransferred"].ID:
		return _GraphOCR3.ParseOwnershipTransferred(log)
	case _GraphOCR3.abi.Events["Transmitted"].ID:
		return _GraphOCR3.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (GraphOCR3ConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (GraphOCR3NeighborAdded) Topic() common.Hash {
	return common.HexToHash("0xbbfdc018e8b91f5b758da7f1c7fdc525ce2fa1e37315698b8b6081718a5c2c81")
}

func (GraphOCR3OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (GraphOCR3OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (GraphOCR3Transmitted) Topic() common.Hash {
	return common.HexToHash("0xe893c2681d327421d89e1cb54fbe64645b4dcea668d6826130b62cf4c6eefea2")
}

func (_GraphOCR3 *GraphOCR3) Address() common.Address {
	return _GraphOCR3.address
}

type GraphOCR3Interface interface {
	GetNeighbors(opts *bind.CallOpts) ([]GraphOCR3Neighbor, error)

	GetTransmitters(opts *bind.CallOpts) ([]common.Address, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddNeighbor(opts *bind.TransactOpts, chainId *big.Int, contractAddress common.Address) (*types.Transaction, error)

	SetOCR3Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*GraphOCR3ConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *GraphOCR3ConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*GraphOCR3ConfigSet, error)

	FilterNeighborAdded(opts *bind.FilterOpts) (*GraphOCR3NeighborAddedIterator, error)

	WatchNeighborAdded(opts *bind.WatchOpts, sink chan<- *GraphOCR3NeighborAdded) (event.Subscription, error)

	ParseNeighborAdded(log types.Log) (*GraphOCR3NeighborAdded, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GraphOCR3OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *GraphOCR3OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*GraphOCR3OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GraphOCR3OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GraphOCR3OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*GraphOCR3OwnershipTransferred, error)

	FilterTransmitted(opts *bind.FilterOpts) (*GraphOCR3TransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *GraphOCR3Transmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*GraphOCR3Transmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
