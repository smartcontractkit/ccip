// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package single_token_onramp

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

type CCIPMessage struct {
	SequenceNumber     *big.Int
	SourceChainId      *big.Int
	DestinationChainId *big.Int
	Sender             common.Address
	Payload            CCIPMessagePayload
}

type CCIPMessagePayload struct {
	Receiver common.Address
	Data     []byte
	Tokens   []common.Address
	Amounts  []*big.Int
	Executor common.Address
	Options  []byte
}

type TokenLimitsTokenBucket struct {
	Rate        *big.Int
	Capacity    *big.Int
	Tokens      *big.Int
	LastUpdated *big.Int
}

var SingleTokenOnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"sourcePool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destinationToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"enableAllowlist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tokenBucketRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBucketCapacity\",\"type\":\"uint256\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AllowlistEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"AllowlistSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CrossChainSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"full\",\"type\":\"bool\"}],\"name\":\"NewTokenBucketConstructed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DESTINATION_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DESTINATION_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"full\",\"type\":\"bool\"}],\"name\":\"configureTokenBucket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenBucket\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structTokenLimits.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"requestCrossChainSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"setAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setAllowlistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b50604051620026903803806200269083398101604081905262000035916200056f565b6000805460ff191681558290829033908190816200009a5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000d457620000d481620002d9565b5050506001600160a01b0382161580620000ec575080155b156200010b57604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03938416179055600355604080516321df0da760e01b815290518c8316928c16916321df0da7916004808301926020929190829003018186803b1580156200016757600080fd5b505afa1580156200017c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001a2919062000548565b6001600160a01b031614620001ca5760405163936bb5ad60e01b815260040160405180910390fd5b60c08b90526001600160601b031960608b811b8216610100528a811b821660e05260808a905288901b1660a05260016007556004805486151560ff19909116179055855162000221906006906020890190620003f1565b5060005b865181101562000291576001600560008984815181106200024a576200024a62000671565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191691151591909117905580620002888162000647565b91505062000225565b50620002ab848460016200038b60201b620011511760201c565b805160085560208101516009556040810151600a5560600151600b5550620006b69950505050505050505050565b6001600160a01b038116331415620003345760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000091565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b620003b76040518060800160405280600081526020016000815260200160008152602001600081525090565b600082620003c7576000620003c9565b835b6040805160808101825296875260208701959095529385019390935250504260608301525090565b82805482825590600052602060002090810192821562000449579160200282015b828111156200044957825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000412565b50620004579291506200045b565b5090565b5b808211156200045757600081556001016200045c565b80516200047f816200069d565b919050565b600082601f8301126200049657600080fd5b815160206001600160401b0380831115620004b557620004b562000687565b8260051b604051601f19603f83011681018181108482111715620004dd57620004dd62000687565b60405284815283810192508684018288018501891015620004fd57600080fd5b600092505b858310156200052b57620005168162000472565b84529284019260019290920191840162000502565b50979650505050505050565b805180151581146200047f57600080fd5b6000602082840312156200055b57600080fd5b815162000568816200069d565b9392505050565b60008060008060008060008060008060006101608c8e0312156200059257600080fd5b8b519a5060208c0151620005a6816200069d565b60408d0151909a50620005b9816200069d565b60608d01519099509750620005d160808d0162000472565b60a08d01519097506001600160401b03811115620005ee57600080fd5b620005fc8e828f0162000484565b9650506200060d60c08d0162000537565b945060e08c015193506101008c015192506200062d6101208d0162000472565b91506101408c015190509295989b509295989b9093969950565b60006000198214156200066a57634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114620006b357600080fd5b50565b60805160a05160601c60c05160e05160601c6101005160601c611f626200072e600039600081816102eb0152818161073b01526107ac0152600081816102bc015261096f01526000818161031a0152610a500152600081816103e201526108640152600081816102630152610a760152611f626000f3fe608060405234801561001057600080fd5b506004361061018d5760003560e01c806382bfefc8116100e3578063b6608c3b1161008c578063e3a604f811610066578063e3a604f8146103dd578063f2fde38b14610404578063f78faa321461041757600080fd5b8063b6608c3b146103a2578063c5eff3d0146103b5578063d7644ba2146103ca57600080fd5b80638da5cb5b116100bd5780638da5cb5b1461033c5780639504c5191461035f578063b034909c1461039a57600080fd5b806382bfefc8146102e65780638456cb591461030d57806385e1f4d01461031557600080fd5b80632ea02369116101455780635c975abb1161011f5780635c975abb146102a05780637535d246146102b757806379ba5097146102de57600080fd5b80632ea023691461025e5780633f4ba83a14610285578063552b818b1461028d57600080fd5b8063108ee5fc11610176578063108ee5fc146101cd578063181f5a77146101e05780632222dd421461021f57600080fd5b80630649d292146101925780630b514037146101b8575b600080fd5b6101a56101a0366004611951565b610422565b6040519081526020015b60405180910390f35b6101cb6101c6366004611a65565b610b12565b005b6101cb6101db366004611829565b610b88565b604080518082018252601781527f53696e676c65546f6b656e4f6e52616d7020312e312e30000000000000000000602082015290516101af9190611bf9565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101af565b6101a57f000000000000000000000000000000000000000000000000000000000000000081565b6101cb610c64565b6101cb61029b366004611846565b610c76565b60005460ff165b60405190151581526020016101af565b6102397f000000000000000000000000000000000000000000000000000000000000000081565b6101cb610e5a565b6102397f000000000000000000000000000000000000000000000000000000000000000081565b6101cb610f7c565b6101a57f000000000000000000000000000000000000000000000000000000000000000081565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610239565b610367610f8c565b6040516101af91908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6003546101a5565b6101cb6103b0366004611a4c565b610fe4565b6103bd611061565b6040516101af9190611b9f565b6101cb6103d83660046118bb565b6110d0565b6102397f000000000000000000000000000000000000000000000000000000000000000081565b6101cb610412366004611829565b61113d565b60045460ff166102a7565b6000805460ff1615610495576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156104ff57600080fd5b505af1158015610513573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053791906118d8565b1561056d576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b491600480830192606092919082900301818787803b1580156105d957600080fd5b505af11580156105ed573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061191906118f5565b90506003548160200151426106269190611e48565b111561065e576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600454339060ff168015610698575073ffffffffffffffffffffffffffffffffffffffff811660009081526005602052604090205460ff16155b156106e7576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260240161048c565b83604001515160011415806107025750836060015151600114155b15610739576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16846040015160008151811061078757610787611ec7565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614610840577f000000000000000000000000000000000000000000000000000000000000000084604001516000815181106107e2576107e2611ec7565b60200260200101516040517f7626e29600000000000000000000000000000000000000000000000000000000815260040161048c92919073ffffffffffffffffffffffffffffffffffffffff92831681529116602082015260400190565b604080516001808252818301909252600091602080830190803683370190505090507f00000000000000000000000000000000000000000000000000000000000000008160008151811061089657610896611ec7565b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080856040018190525061090b85606001516000815181106108f3576108f3611ec7565b602002602001015160086111b490919063ffffffff16565b61096d57600a546060860151805160009061092857610928611ec7565b60200260200101516040517f331220f700000000000000000000000000000000000000000000000000000000815260040161048c929190918252602082015260400190565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663eb54b3bf8387606001516000815181106109c1576109c1611ec7565b60200260200101516040518363ffffffff1660e01b8152600401610a0792919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b600060405180830381600087803b158015610a2157600080fd5b505af1158015610a35573d6000803e3d6000fd5b5050505060006040518060a0016040528060075481526020017f000000000000000000000000000000000000000000000000000000000000000081526020017f000000000000000000000000000000000000000000000000000000000000000081526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018781525090507f62a59f5eba2b77be4874d2e6dd09a2749d00338373ed017034e6184899c7eb5981604051610aeb9190611c0c565b60405180910390a160078054906000610b0383611e5f565b90915550505195945050505050565b610b1a6111f7565b610b25838383611151565b8051600855602080820151600955604080830151600a55606092830151600b558051868152918201859052831515908201527ffaf3310019e551542b5c6014c1ae13e2a8d3943d7611d779c4df9b36c111924f91015b60405180910390a1505050565b610b906111f7565b73ffffffffffffffffffffffffffffffffffffffff8116610bdd576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b610c6c6111f7565b610c7461127d565b565b610c7e6111f7565b60006006805480602002602001604051908101604052809291908181526020018280548015610ce357602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610cb8575b5050505050905060005b8151811015610d7d57600060056000848481518110610d0e57610d0e611ec7565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905580610d7581611e5f565b915050610ced565b50610d8a60068484611613565b5060005b82811015610e2857600160056000868685818110610dae57610dae611ec7565b9050602002016020810190610dc39190611829565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905580610e2081611e5f565b915050610d8e565b507f27f242de1bc4ed72c4329591ffff7d223b5f025e3514a07e05afec6d4eb889cf8383604051610b7b929190611b44565b60015473ffffffffffffffffffffffffffffffffffffffff163314610edb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161048c565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610f846111f7565b610c7461135e565b610fb76040518060800160405280600081526020016000815260200160008152602001600081525090565b506040805160808101825260085481526009546020820152600a5491810191909152600b54606082015290565b610fec6111f7565b80611023576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c2519101610c58565b606060068054806020026020016040519081016040528092919081815260200182805480156110c657602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161109b575b5050505050905090565b6110d86111f7565b600480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168215159081179091556040519081527fa1bf86c493917580dec207969ef59976f0c378f10ece581237f19acfbd858f1c9060200160405180910390a150565b6111456111f7565b61114e8161141e565b50565b61117c6040518060800160405280600081526020016000815260200160008152602001600081525090565b60008261118a57600061118c565b835b6040805160808101825296875260208701959095529385019390935250504260608301525090565b60006111bf8361151a565b81836002015410156111d3575060006111f1565b818360020160008282546111e79190611e48565b9091555060019150505b92915050565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610c74576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161048c565b60005460ff166112e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161048c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60005460ff16156113cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161048c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586113343390565b73ffffffffffffffffffffffffffffffffffffffff811633141561149e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161048c565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80600101548160020154111561155c576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600101548160020154141561156f5750565b600381015442908110156115af576040517ff01f197500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260030154826115c19190611e48565b600184015484549192506115ed916115d99084611e0b565b85600201546115e89190611df3565b6115fb565b600284015550600390910155565b600081831061160a578161160c565b825b9392505050565b82805482825590600052602060002090810192821561168b579160200282015b8281111561168b5781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff843516178255602090920191600190910190611633565b5061169792915061169b565b5090565b5b80821115611697576000815560010161169c565b80356116bb81611f25565b919050565b600082601f8301126116d157600080fd5b813560206116e66116e183611dcf565b611d80565b80838252828201915082860187848660051b890101111561170657600080fd5b60005b8581101561172e57813561171c81611f25565b84529284019290840190600101611709565b5090979650505050505050565b600082601f83011261174c57600080fd5b8135602061175c6116e183611dcf565b80838252828201915082860187848660051b890101111561177c57600080fd5b60005b8581101561172e5781358452928401929084019060010161177f565b600082601f8301126117ac57600080fd5b813567ffffffffffffffff8111156117c6576117c6611ef6565b6117f760207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611d80565b81815284602083860101111561180c57600080fd5b816020850160208301376000918101602001919091529392505050565b60006020828403121561183b57600080fd5b813561160c81611f25565b6000806020838503121561185957600080fd5b823567ffffffffffffffff8082111561187157600080fd5b818501915085601f83011261188557600080fd5b81358181111561189457600080fd5b8660208260051b85010111156118a957600080fd5b60209290920196919550909350505050565b6000602082840312156118cd57600080fd5b813561160c81611f47565b6000602082840312156118ea57600080fd5b815161160c81611f47565b60006060828403121561190757600080fd5b6040516060810181811067ffffffffffffffff8211171561192a5761192a611ef6565b80604052508251815260208301516020820152604083015160408201528091505092915050565b60006020828403121561196357600080fd5b813567ffffffffffffffff8082111561197b57600080fd5b9083019060c0828603121561198f57600080fd5b611997611d57565b6119a0836116b0565b81526020830135828111156119b457600080fd5b6119c08782860161179b565b6020830152506040830135828111156119d857600080fd5b6119e4878286016116c0565b6040830152506060830135828111156119fc57600080fd5b611a088782860161173b565b606083015250611a1a608084016116b0565b608082015260a083013582811115611a3157600080fd5b611a3d8782860161179b565b60a08301525095945050505050565b600060208284031215611a5e57600080fd5b5035919050565b600080600060608486031215611a7a57600080fd5b83359250602084013591506040840135611a9381611f47565b809150509250925092565b600081518084526020808501945080840160005b83811015611ace57815187529582019590820190600101611ab2565b509495945050505050565b6000815180845260005b81811015611aff57602081850181015186830182015201611ae3565b81811115611b11576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208082528181018390526000908460408401835b86811015611b94578235611b6c81611f25565b73ffffffffffffffffffffffffffffffffffffffff1682529183019190830190600101611b59565b509695505050505050565b6020808252825182820181905260009190848201906040850190845b81811015611bed57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101611bbb565b50909695505050505050565b60208152600061160c6020830184611ad9565b6000602080835283518184015280840151604084015260408401516060840152606084015173ffffffffffffffffffffffffffffffffffffffff80821660808601526080860151915060a0808601528082511660c08601528282015160c060e0870152611c7d610180870182611ad9565b60408401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4088830381016101008a0152815180845291870193506000929091908701905b80841015611ce457845186168252938701936001939093019290870190611cc2565b506060860151965081898203016101208a0152611d018188611a9e565b96505060808501519350611d2e61014089018573ffffffffffffffffffffffffffffffffffffffff169052565b60a08501519450808887030161016089015250505050611d4e8282611ad9565b95945050505050565b60405160c0810167ffffffffffffffff81118282101715611d7a57611d7a611ef6565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611dc757611dc7611ef6565b604052919050565b600067ffffffffffffffff821115611de957611de9611ef6565b5060051b60200190565b60008219821115611e0657611e06611e98565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611e4357611e43611e98565b500290565b600082821015611e5a57611e5a611e98565b500390565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611e9157611e91611e98565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461114e57600080fd5b801515811461114e57600080fdfea164736f6c6343000806000a",
}

var SingleTokenOnRampABI = SingleTokenOnRampMetaData.ABI

var SingleTokenOnRampBin = SingleTokenOnRampMetaData.Bin

func DeploySingleTokenOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, sourceToken common.Address, sourcePool common.Address, destinationChainId *big.Int, destinationToken common.Address, allowlist []common.Address, enableAllowlist bool, tokenBucketRate *big.Int, tokenBucketCapacity *big.Int, afn common.Address, maxTimeWithoutAFNSignal *big.Int) (common.Address, *types.Transaction, *SingleTokenOnRamp, error) {
	parsed, err := SingleTokenOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SingleTokenOnRampBin), backend, sourceChainId, sourceToken, sourcePool, destinationChainId, destinationToken, allowlist, enableAllowlist, tokenBucketRate, tokenBucketCapacity, afn, maxTimeWithoutAFNSignal)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SingleTokenOnRamp{SingleTokenOnRampCaller: SingleTokenOnRampCaller{contract: contract}, SingleTokenOnRampTransactor: SingleTokenOnRampTransactor{contract: contract}, SingleTokenOnRampFilterer: SingleTokenOnRampFilterer{contract: contract}}, nil
}

type SingleTokenOnRamp struct {
	address common.Address
	abi     abi.ABI
	SingleTokenOnRampCaller
	SingleTokenOnRampTransactor
	SingleTokenOnRampFilterer
}

type SingleTokenOnRampCaller struct {
	contract *bind.BoundContract
}

type SingleTokenOnRampTransactor struct {
	contract *bind.BoundContract
}

type SingleTokenOnRampFilterer struct {
	contract *bind.BoundContract
}

type SingleTokenOnRampSession struct {
	Contract     *SingleTokenOnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type SingleTokenOnRampCallerSession struct {
	Contract *SingleTokenOnRampCaller
	CallOpts bind.CallOpts
}

type SingleTokenOnRampTransactorSession struct {
	Contract     *SingleTokenOnRampTransactor
	TransactOpts bind.TransactOpts
}

type SingleTokenOnRampRaw struct {
	Contract *SingleTokenOnRamp
}

type SingleTokenOnRampCallerRaw struct {
	Contract *SingleTokenOnRampCaller
}

type SingleTokenOnRampTransactorRaw struct {
	Contract *SingleTokenOnRampTransactor
}

func NewSingleTokenOnRamp(address common.Address, backend bind.ContractBackend) (*SingleTokenOnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(SingleTokenOnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindSingleTokenOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SingleTokenOnRamp{address: address, abi: abi, SingleTokenOnRampCaller: SingleTokenOnRampCaller{contract: contract}, SingleTokenOnRampTransactor: SingleTokenOnRampTransactor{contract: contract}, SingleTokenOnRampFilterer: SingleTokenOnRampFilterer{contract: contract}}, nil
}

func NewSingleTokenOnRampCaller(address common.Address, caller bind.ContractCaller) (*SingleTokenOnRampCaller, error) {
	contract, err := bindSingleTokenOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SingleTokenOnRampCaller{contract: contract}, nil
}

func NewSingleTokenOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*SingleTokenOnRampTransactor, error) {
	contract, err := bindSingleTokenOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SingleTokenOnRampTransactor{contract: contract}, nil
}

func NewSingleTokenOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*SingleTokenOnRampFilterer, error) {
	contract, err := bindSingleTokenOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SingleTokenOnRampFilterer{contract: contract}, nil
}

func bindSingleTokenOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SingleTokenOnRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_SingleTokenOnRamp *SingleTokenOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SingleTokenOnRamp.Contract.SingleTokenOnRampCaller.contract.Call(opts, result, method, params...)
}

func (_SingleTokenOnRamp *SingleTokenOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.SingleTokenOnRampTransactor.contract.Transfer(opts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.SingleTokenOnRampTransactor.contract.Transact(opts, method, params...)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SingleTokenOnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.contract.Transfer(opts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingleTokenOnRamp.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) CHAINID() (*big.Int, error) {
	return _SingleTokenOnRamp.Contract.CHAINID(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCallerSession) CHAINID() (*big.Int, error) {
	return _SingleTokenOnRamp.Contract.CHAINID(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCaller) DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingleTokenOnRamp.contract.Call(opts, &out, "DESTINATION_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _SingleTokenOnRamp.Contract.DESTINATIONCHAINID(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCallerSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _SingleTokenOnRamp.Contract.DESTINATIONCHAINID(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCaller) DESTINATIONTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleTokenOnRamp.contract.Call(opts, &out, "DESTINATION_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) DESTINATIONTOKEN() (common.Address, error) {
	return _SingleTokenOnRamp.Contract.DESTINATIONTOKEN(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCallerSession) DESTINATIONTOKEN() (common.Address, error) {
	return _SingleTokenOnRamp.Contract.DESTINATIONTOKEN(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCaller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleTokenOnRamp.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) POOL() (common.Address, error) {
	return _SingleTokenOnRamp.Contract.POOL(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCallerSession) POOL() (common.Address, error) {
	return _SingleTokenOnRamp.Contract.POOL(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCaller) TOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleTokenOnRamp.contract.Call(opts, &out, "TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) TOKEN() (common.Address, error) {
	return _SingleTokenOnRamp.Contract.TOKEN(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCallerSession) TOKEN() (common.Address, error) {
	return _SingleTokenOnRamp.Contract.TOKEN(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleTokenOnRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) GetAFN() (common.Address, error) {
	return _SingleTokenOnRamp.Contract.GetAFN(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCallerSession) GetAFN() (common.Address, error) {
	return _SingleTokenOnRamp.Contract.GetAFN(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCaller) GetAllowlist(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SingleTokenOnRamp.contract.Call(opts, &out, "getAllowlist")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) GetAllowlist() ([]common.Address, error) {
	return _SingleTokenOnRamp.Contract.GetAllowlist(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCallerSession) GetAllowlist() ([]common.Address, error) {
	return _SingleTokenOnRamp.Contract.GetAllowlist(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCaller) GetAllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SingleTokenOnRamp.contract.Call(opts, &out, "getAllowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) GetAllowlistEnabled() (bool, error) {
	return _SingleTokenOnRamp.Contract.GetAllowlistEnabled(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCallerSession) GetAllowlistEnabled() (bool, error) {
	return _SingleTokenOnRamp.Contract.GetAllowlistEnabled(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCaller) GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingleTokenOnRamp.contract.Call(opts, &out, "getMaxSecondsWithoutAFNHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _SingleTokenOnRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCallerSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _SingleTokenOnRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCaller) GetTokenBucket(opts *bind.CallOpts) (TokenLimitsTokenBucket, error) {
	var out []interface{}
	err := _SingleTokenOnRamp.contract.Call(opts, &out, "getTokenBucket")

	if err != nil {
		return *new(TokenLimitsTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenLimitsTokenBucket)).(*TokenLimitsTokenBucket)

	return out0, err

}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) GetTokenBucket() (TokenLimitsTokenBucket, error) {
	return _SingleTokenOnRamp.Contract.GetTokenBucket(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCallerSession) GetTokenBucket() (TokenLimitsTokenBucket, error) {
	return _SingleTokenOnRamp.Contract.GetTokenBucket(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleTokenOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) Owner() (common.Address, error) {
	return _SingleTokenOnRamp.Contract.Owner(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCallerSession) Owner() (common.Address, error) {
	return _SingleTokenOnRamp.Contract.Owner(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SingleTokenOnRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) Paused() (bool, error) {
	return _SingleTokenOnRamp.Contract.Paused(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCallerSession) Paused() (bool, error) {
	return _SingleTokenOnRamp.Contract.Paused(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SingleTokenOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) TypeAndVersion() (string, error) {
	return _SingleTokenOnRamp.Contract.TypeAndVersion(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampCallerSession) TypeAndVersion() (string, error) {
	return _SingleTokenOnRamp.Contract.TypeAndVersion(&_SingleTokenOnRamp.CallOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleTokenOnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.AcceptOwnership(&_SingleTokenOnRamp.TransactOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.AcceptOwnership(&_SingleTokenOnRamp.TransactOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactor) ConfigureTokenBucket(opts *bind.TransactOpts, rate *big.Int, capacity *big.Int, full bool) (*types.Transaction, error) {
	return _SingleTokenOnRamp.contract.Transact(opts, "configureTokenBucket", rate, capacity, full)
}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) ConfigureTokenBucket(rate *big.Int, capacity *big.Int, full bool) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.ConfigureTokenBucket(&_SingleTokenOnRamp.TransactOpts, rate, capacity, full)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactorSession) ConfigureTokenBucket(rate *big.Int, capacity *big.Int, full bool) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.ConfigureTokenBucket(&_SingleTokenOnRamp.TransactOpts, rate, capacity, full)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleTokenOnRamp.contract.Transact(opts, "pause")
}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) Pause() (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.Pause(&_SingleTokenOnRamp.TransactOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactorSession) Pause() (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.Pause(&_SingleTokenOnRamp.TransactOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactor) RequestCrossChainSend(opts *bind.TransactOpts, payload CCIPMessagePayload) (*types.Transaction, error) {
	return _SingleTokenOnRamp.contract.Transact(opts, "requestCrossChainSend", payload)
}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) RequestCrossChainSend(payload CCIPMessagePayload) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.RequestCrossChainSend(&_SingleTokenOnRamp.TransactOpts, payload)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactorSession) RequestCrossChainSend(payload CCIPMessagePayload) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.RequestCrossChainSend(&_SingleTokenOnRamp.TransactOpts, payload)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _SingleTokenOnRamp.contract.Transact(opts, "setAFN", afn)
}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.SetAFN(&_SingleTokenOnRamp.TransactOpts, afn)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.SetAFN(&_SingleTokenOnRamp.TransactOpts, afn)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactor) SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error) {
	return _SingleTokenOnRamp.contract.Transact(opts, "setAllowlist", allowlist)
}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.SetAllowlist(&_SingleTokenOnRamp.TransactOpts, allowlist)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactorSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.SetAllowlist(&_SingleTokenOnRamp.TransactOpts, allowlist)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactor) SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _SingleTokenOnRamp.contract.Transact(opts, "setAllowlistEnabled", enabled)
}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.SetAllowlistEnabled(&_SingleTokenOnRamp.TransactOpts, enabled)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactorSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.SetAllowlistEnabled(&_SingleTokenOnRamp.TransactOpts, enabled)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _SingleTokenOnRamp.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_SingleTokenOnRamp.TransactOpts, newTime)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_SingleTokenOnRamp.TransactOpts, newTime)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SingleTokenOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.TransferOwnership(&_SingleTokenOnRamp.TransactOpts, to)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.TransferOwnership(&_SingleTokenOnRamp.TransactOpts, to)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleTokenOnRamp.contract.Transact(opts, "unpause")
}

func (_SingleTokenOnRamp *SingleTokenOnRampSession) Unpause() (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.Unpause(&_SingleTokenOnRamp.TransactOpts)
}

func (_SingleTokenOnRamp *SingleTokenOnRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _SingleTokenOnRamp.Contract.Unpause(&_SingleTokenOnRamp.TransactOpts)
}

type SingleTokenOnRampAFNMaxHeartbeatTimeSetIterator struct {
	Event *SingleTokenOnRampAFNMaxHeartbeatTimeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SingleTokenOnRampAFNMaxHeartbeatTimeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleTokenOnRampAFNMaxHeartbeatTimeSet)
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
		it.Event = new(SingleTokenOnRampAFNMaxHeartbeatTimeSet)
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

func (it *SingleTokenOnRampAFNMaxHeartbeatTimeSetIterator) Error() error {
	return it.fail
}

func (it *SingleTokenOnRampAFNMaxHeartbeatTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SingleTokenOnRampAFNMaxHeartbeatTimeSet struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*SingleTokenOnRampAFNMaxHeartbeatTimeSetIterator, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.FilterLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return &SingleTokenOnRampAFNMaxHeartbeatTimeSetIterator{contract: _SingleTokenOnRamp.contract, event: "AFNMaxHeartbeatTimeSet", logs: logs, sub: sub}, nil
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampAFNMaxHeartbeatTimeSet) (event.Subscription, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.WatchLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SingleTokenOnRampAFNMaxHeartbeatTimeSet)
				if err := _SingleTokenOnRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
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

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) ParseAFNMaxHeartbeatTimeSet(log types.Log) (*SingleTokenOnRampAFNMaxHeartbeatTimeSet, error) {
	event := new(SingleTokenOnRampAFNMaxHeartbeatTimeSet)
	if err := _SingleTokenOnRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SingleTokenOnRampAFNSetIterator struct {
	Event *SingleTokenOnRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SingleTokenOnRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleTokenOnRampAFNSet)
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
		it.Event = new(SingleTokenOnRampAFNSet)
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

func (it *SingleTokenOnRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *SingleTokenOnRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SingleTokenOnRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*SingleTokenOnRampAFNSetIterator, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &SingleTokenOnRampAFNSetIterator{contract: _SingleTokenOnRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SingleTokenOnRampAFNSet)
				if err := _SingleTokenOnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) ParseAFNSet(log types.Log) (*SingleTokenOnRampAFNSet, error) {
	event := new(SingleTokenOnRampAFNSet)
	if err := _SingleTokenOnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SingleTokenOnRampAllowlistEnabledSetIterator struct {
	Event *SingleTokenOnRampAllowlistEnabledSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SingleTokenOnRampAllowlistEnabledSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleTokenOnRampAllowlistEnabledSet)
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
		it.Event = new(SingleTokenOnRampAllowlistEnabledSet)
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

func (it *SingleTokenOnRampAllowlistEnabledSetIterator) Error() error {
	return it.fail
}

func (it *SingleTokenOnRampAllowlistEnabledSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SingleTokenOnRampAllowlistEnabledSet struct {
	Enabled bool
	Raw     types.Log
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) FilterAllowlistEnabledSet(opts *bind.FilterOpts) (*SingleTokenOnRampAllowlistEnabledSetIterator, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.FilterLogs(opts, "AllowlistEnabledSet")
	if err != nil {
		return nil, err
	}
	return &SingleTokenOnRampAllowlistEnabledSetIterator{contract: _SingleTokenOnRamp.contract, event: "AllowlistEnabledSet", logs: logs, sub: sub}, nil
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) WatchAllowlistEnabledSet(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampAllowlistEnabledSet) (event.Subscription, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.WatchLogs(opts, "AllowlistEnabledSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SingleTokenOnRampAllowlistEnabledSet)
				if err := _SingleTokenOnRamp.contract.UnpackLog(event, "AllowlistEnabledSet", log); err != nil {
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

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) ParseAllowlistEnabledSet(log types.Log) (*SingleTokenOnRampAllowlistEnabledSet, error) {
	event := new(SingleTokenOnRampAllowlistEnabledSet)
	if err := _SingleTokenOnRamp.contract.UnpackLog(event, "AllowlistEnabledSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SingleTokenOnRampAllowlistSetIterator struct {
	Event *SingleTokenOnRampAllowlistSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SingleTokenOnRampAllowlistSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleTokenOnRampAllowlistSet)
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
		it.Event = new(SingleTokenOnRampAllowlistSet)
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

func (it *SingleTokenOnRampAllowlistSetIterator) Error() error {
	return it.fail
}

func (it *SingleTokenOnRampAllowlistSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SingleTokenOnRampAllowlistSet struct {
	Allowlist []common.Address
	Raw       types.Log
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) FilterAllowlistSet(opts *bind.FilterOpts) (*SingleTokenOnRampAllowlistSetIterator, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.FilterLogs(opts, "AllowlistSet")
	if err != nil {
		return nil, err
	}
	return &SingleTokenOnRampAllowlistSetIterator{contract: _SingleTokenOnRamp.contract, event: "AllowlistSet", logs: logs, sub: sub}, nil
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) WatchAllowlistSet(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampAllowlistSet) (event.Subscription, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.WatchLogs(opts, "AllowlistSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SingleTokenOnRampAllowlistSet)
				if err := _SingleTokenOnRamp.contract.UnpackLog(event, "AllowlistSet", log); err != nil {
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

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) ParseAllowlistSet(log types.Log) (*SingleTokenOnRampAllowlistSet, error) {
	event := new(SingleTokenOnRampAllowlistSet)
	if err := _SingleTokenOnRamp.contract.UnpackLog(event, "AllowlistSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SingleTokenOnRampCrossChainSendRequestedIterator struct {
	Event *SingleTokenOnRampCrossChainSendRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SingleTokenOnRampCrossChainSendRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleTokenOnRampCrossChainSendRequested)
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
		it.Event = new(SingleTokenOnRampCrossChainSendRequested)
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

func (it *SingleTokenOnRampCrossChainSendRequestedIterator) Error() error {
	return it.fail
}

func (it *SingleTokenOnRampCrossChainSendRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SingleTokenOnRampCrossChainSendRequested struct {
	Message CCIPMessage
	Raw     types.Log
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) FilterCrossChainSendRequested(opts *bind.FilterOpts) (*SingleTokenOnRampCrossChainSendRequestedIterator, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.FilterLogs(opts, "CrossChainSendRequested")
	if err != nil {
		return nil, err
	}
	return &SingleTokenOnRampCrossChainSendRequestedIterator{contract: _SingleTokenOnRamp.contract, event: "CrossChainSendRequested", logs: logs, sub: sub}, nil
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) WatchCrossChainSendRequested(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampCrossChainSendRequested) (event.Subscription, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.WatchLogs(opts, "CrossChainSendRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SingleTokenOnRampCrossChainSendRequested)
				if err := _SingleTokenOnRamp.contract.UnpackLog(event, "CrossChainSendRequested", log); err != nil {
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

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) ParseCrossChainSendRequested(log types.Log) (*SingleTokenOnRampCrossChainSendRequested, error) {
	event := new(SingleTokenOnRampCrossChainSendRequested)
	if err := _SingleTokenOnRamp.contract.UnpackLog(event, "CrossChainSendRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SingleTokenOnRampNewTokenBucketConstructedIterator struct {
	Event *SingleTokenOnRampNewTokenBucketConstructed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SingleTokenOnRampNewTokenBucketConstructedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleTokenOnRampNewTokenBucketConstructed)
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
		it.Event = new(SingleTokenOnRampNewTokenBucketConstructed)
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

func (it *SingleTokenOnRampNewTokenBucketConstructedIterator) Error() error {
	return it.fail
}

func (it *SingleTokenOnRampNewTokenBucketConstructedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SingleTokenOnRampNewTokenBucketConstructed struct {
	Rate     *big.Int
	Capacity *big.Int
	Full     bool
	Raw      types.Log
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) FilterNewTokenBucketConstructed(opts *bind.FilterOpts) (*SingleTokenOnRampNewTokenBucketConstructedIterator, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.FilterLogs(opts, "NewTokenBucketConstructed")
	if err != nil {
		return nil, err
	}
	return &SingleTokenOnRampNewTokenBucketConstructedIterator{contract: _SingleTokenOnRamp.contract, event: "NewTokenBucketConstructed", logs: logs, sub: sub}, nil
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) WatchNewTokenBucketConstructed(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampNewTokenBucketConstructed) (event.Subscription, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.WatchLogs(opts, "NewTokenBucketConstructed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SingleTokenOnRampNewTokenBucketConstructed)
				if err := _SingleTokenOnRamp.contract.UnpackLog(event, "NewTokenBucketConstructed", log); err != nil {
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

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) ParseNewTokenBucketConstructed(log types.Log) (*SingleTokenOnRampNewTokenBucketConstructed, error) {
	event := new(SingleTokenOnRampNewTokenBucketConstructed)
	if err := _SingleTokenOnRamp.contract.UnpackLog(event, "NewTokenBucketConstructed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SingleTokenOnRampOwnershipTransferRequestedIterator struct {
	Event *SingleTokenOnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SingleTokenOnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleTokenOnRampOwnershipTransferRequested)
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
		it.Event = new(SingleTokenOnRampOwnershipTransferRequested)
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

func (it *SingleTokenOnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *SingleTokenOnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SingleTokenOnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SingleTokenOnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SingleTokenOnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SingleTokenOnRampOwnershipTransferRequestedIterator{contract: _SingleTokenOnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SingleTokenOnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SingleTokenOnRampOwnershipTransferRequested)
				if err := _SingleTokenOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*SingleTokenOnRampOwnershipTransferRequested, error) {
	event := new(SingleTokenOnRampOwnershipTransferRequested)
	if err := _SingleTokenOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SingleTokenOnRampOwnershipTransferredIterator struct {
	Event *SingleTokenOnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SingleTokenOnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleTokenOnRampOwnershipTransferred)
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
		it.Event = new(SingleTokenOnRampOwnershipTransferred)
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

func (it *SingleTokenOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *SingleTokenOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SingleTokenOnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SingleTokenOnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SingleTokenOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SingleTokenOnRampOwnershipTransferredIterator{contract: _SingleTokenOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SingleTokenOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SingleTokenOnRampOwnershipTransferred)
				if err := _SingleTokenOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*SingleTokenOnRampOwnershipTransferred, error) {
	event := new(SingleTokenOnRampOwnershipTransferred)
	if err := _SingleTokenOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SingleTokenOnRampPausedIterator struct {
	Event *SingleTokenOnRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SingleTokenOnRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleTokenOnRampPaused)
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
		it.Event = new(SingleTokenOnRampPaused)
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

func (it *SingleTokenOnRampPausedIterator) Error() error {
	return it.fail
}

func (it *SingleTokenOnRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SingleTokenOnRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) FilterPaused(opts *bind.FilterOpts) (*SingleTokenOnRampPausedIterator, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SingleTokenOnRampPausedIterator{contract: _SingleTokenOnRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampPaused) (event.Subscription, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SingleTokenOnRampPaused)
				if err := _SingleTokenOnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) ParsePaused(log types.Log) (*SingleTokenOnRampPaused, error) {
	event := new(SingleTokenOnRampPaused)
	if err := _SingleTokenOnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SingleTokenOnRampUnpausedIterator struct {
	Event *SingleTokenOnRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SingleTokenOnRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleTokenOnRampUnpaused)
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
		it.Event = new(SingleTokenOnRampUnpaused)
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

func (it *SingleTokenOnRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *SingleTokenOnRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SingleTokenOnRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SingleTokenOnRampUnpausedIterator, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SingleTokenOnRampUnpausedIterator{contract: _SingleTokenOnRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _SingleTokenOnRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SingleTokenOnRampUnpaused)
				if err := _SingleTokenOnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_SingleTokenOnRamp *SingleTokenOnRampFilterer) ParseUnpaused(log types.Log) (*SingleTokenOnRampUnpaused, error) {
	event := new(SingleTokenOnRampUnpaused)
	if err := _SingleTokenOnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_SingleTokenOnRamp *SingleTokenOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _SingleTokenOnRamp.abi.Events["AFNMaxHeartbeatTimeSet"].ID:
		return _SingleTokenOnRamp.ParseAFNMaxHeartbeatTimeSet(log)
	case _SingleTokenOnRamp.abi.Events["AFNSet"].ID:
		return _SingleTokenOnRamp.ParseAFNSet(log)
	case _SingleTokenOnRamp.abi.Events["AllowlistEnabledSet"].ID:
		return _SingleTokenOnRamp.ParseAllowlistEnabledSet(log)
	case _SingleTokenOnRamp.abi.Events["AllowlistSet"].ID:
		return _SingleTokenOnRamp.ParseAllowlistSet(log)
	case _SingleTokenOnRamp.abi.Events["CrossChainSendRequested"].ID:
		return _SingleTokenOnRamp.ParseCrossChainSendRequested(log)
	case _SingleTokenOnRamp.abi.Events["NewTokenBucketConstructed"].ID:
		return _SingleTokenOnRamp.ParseNewTokenBucketConstructed(log)
	case _SingleTokenOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _SingleTokenOnRamp.ParseOwnershipTransferRequested(log)
	case _SingleTokenOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _SingleTokenOnRamp.ParseOwnershipTransferred(log)
	case _SingleTokenOnRamp.abi.Events["Paused"].ID:
		return _SingleTokenOnRamp.ParsePaused(log)
	case _SingleTokenOnRamp.abi.Events["Unpaused"].ID:
		return _SingleTokenOnRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (SingleTokenOnRampAFNMaxHeartbeatTimeSet) Topic() common.Hash {
	return common.HexToHash("0x72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251")
}

func (SingleTokenOnRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (SingleTokenOnRampAllowlistEnabledSet) Topic() common.Hash {
	return common.HexToHash("0xa1bf86c493917580dec207969ef59976f0c378f10ece581237f19acfbd858f1c")
}

func (SingleTokenOnRampAllowlistSet) Topic() common.Hash {
	return common.HexToHash("0x27f242de1bc4ed72c4329591ffff7d223b5f025e3514a07e05afec6d4eb889cf")
}

func (SingleTokenOnRampCrossChainSendRequested) Topic() common.Hash {
	return common.HexToHash("0x62a59f5eba2b77be4874d2e6dd09a2749d00338373ed017034e6184899c7eb59")
}

func (SingleTokenOnRampNewTokenBucketConstructed) Topic() common.Hash {
	return common.HexToHash("0xfaf3310019e551542b5c6014c1ae13e2a8d3943d7611d779c4df9b36c111924f")
}

func (SingleTokenOnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (SingleTokenOnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (SingleTokenOnRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (SingleTokenOnRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_SingleTokenOnRamp *SingleTokenOnRamp) Address() common.Address {
	return _SingleTokenOnRamp.address
}

type SingleTokenOnRampInterface interface {
	CHAINID(opts *bind.CallOpts) (*big.Int, error)

	DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error)

	DESTINATIONTOKEN(opts *bind.CallOpts) (common.Address, error)

	POOL(opts *bind.CallOpts) (common.Address, error)

	TOKEN(opts *bind.CallOpts) (common.Address, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetAllowlist(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowlistEnabled(opts *bind.CallOpts) (bool, error)

	GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error)

	GetTokenBucket(opts *bind.CallOpts) (TokenLimitsTokenBucket, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ConfigureTokenBucket(opts *bind.TransactOpts, rate *big.Int, capacity *big.Int, full bool) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RequestCrossChainSend(opts *bind.TransactOpts, payload CCIPMessagePayload) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error)

	SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error)

	SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*SingleTokenOnRampAFNMaxHeartbeatTimeSetIterator, error)

	WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampAFNMaxHeartbeatTimeSet) (event.Subscription, error)

	ParseAFNMaxHeartbeatTimeSet(log types.Log) (*SingleTokenOnRampAFNMaxHeartbeatTimeSet, error)

	FilterAFNSet(opts *bind.FilterOpts) (*SingleTokenOnRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*SingleTokenOnRampAFNSet, error)

	FilterAllowlistEnabledSet(opts *bind.FilterOpts) (*SingleTokenOnRampAllowlistEnabledSetIterator, error)

	WatchAllowlistEnabledSet(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampAllowlistEnabledSet) (event.Subscription, error)

	ParseAllowlistEnabledSet(log types.Log) (*SingleTokenOnRampAllowlistEnabledSet, error)

	FilterAllowlistSet(opts *bind.FilterOpts) (*SingleTokenOnRampAllowlistSetIterator, error)

	WatchAllowlistSet(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampAllowlistSet) (event.Subscription, error)

	ParseAllowlistSet(log types.Log) (*SingleTokenOnRampAllowlistSet, error)

	FilterCrossChainSendRequested(opts *bind.FilterOpts) (*SingleTokenOnRampCrossChainSendRequestedIterator, error)

	WatchCrossChainSendRequested(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampCrossChainSendRequested) (event.Subscription, error)

	ParseCrossChainSendRequested(log types.Log) (*SingleTokenOnRampCrossChainSendRequested, error)

	FilterNewTokenBucketConstructed(opts *bind.FilterOpts) (*SingleTokenOnRampNewTokenBucketConstructedIterator, error)

	WatchNewTokenBucketConstructed(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampNewTokenBucketConstructed) (event.Subscription, error)

	ParseNewTokenBucketConstructed(log types.Log) (*SingleTokenOnRampNewTokenBucketConstructed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SingleTokenOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*SingleTokenOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SingleTokenOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*SingleTokenOnRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*SingleTokenOnRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*SingleTokenOnRampPaused, error)

	FilterUnpaused(opts *bind.FilterOpts) (*SingleTokenOnRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SingleTokenOnRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*SingleTokenOnRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
