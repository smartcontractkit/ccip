// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package self_funded_ping_pong

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

type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

var SelfFundedPingPongMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"roundTripsBeforeFunding\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"countIncrBeforeFunding\",\"type\":\"uint8\"}],\"name\":\"CountIncrBeforeFundingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Funded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOutOfOrder\",\"type\":\"bool\"}],\"name\":\"OutOfOrderExecutionChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingPongCount\",\"type\":\"uint256\"}],\"name\":\"Ping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingPongCount\",\"type\":\"uint256\"}],\"name\":\"Pong\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pingPongCount\",\"type\":\"uint256\"}],\"name\":\"fundPingPong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCountIncrBeforeFunding\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounterpartAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounterpartChainSelector\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOutOfOrderExecution\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"countIncrBeforeFunding\",\"type\":\"uint8\"}],\"name\":\"setCountIncrBeforeFunding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"counterpartChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"counterpartAddress\",\"type\":\"address\"}],\"name\":\"setCounterpart\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setCounterpartAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"setCounterpartChainSelector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"outOfOrderExecution\",\"type\":\"bool\"}],\"name\":\"setOutOfOrderExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startPingPong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200190938038062001909833981016040819052620000349162000291565b828233806000846001600160a01b0381166200006b576040516335fdcccd60e21b8152600060048201526024015b60405180910390fd5b6001600160a01b039081166080528216620000c95760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f0000000000000000604482015260640162000062565b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000fc57620000fc81620001cd565b50506002805460ff60a01b1916905550600380546001600160a01b0319166001600160a01b0383811691821790925560405163095ea7b360e01b8152918416600483015260001960248301529063095ea7b3906044016020604051808303816000875af115801562000172573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001989190620002ea565b505050806002620001aa919062000315565b600360156101000a81548160ff021916908360ff16021790555050505062000347565b336001600160a01b03821603620002275760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000062565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b03811681146200028e57600080fd5b50565b600080600060608486031215620002a757600080fd5b8351620002b48162000278565b6020850151909350620002c78162000278565b604085015190925060ff81168114620002df57600080fd5b809150509250925092565b600060208284031215620002fd57600080fd5b815180151581146200030e57600080fd5b9392505050565b60ff81811683821602908116908181146200034057634e487b7160e01b600052601160045260246000fd5b5092915050565b60805161159162000378600039600081816102f301528181610728015281816108180152610d0901526115916000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80638f491cba116100d8578063b5a110111161008c578063e6c725f511610066578063e6c725f5146103a9578063ef686d8e146103da578063f2fde38b146103ed57600080fd5b8063b5a110111461033a578063bee518a41461034d578063ca709a251461038b57600080fd5b8063ae90de55116100bd578063ae90de55146102ce578063b0f479a1146102f1578063b187bd261461031757600080fd5b80638f491cba146102a85780639d2aede5146102bb57600080fd5b80632b6e5d631161012f57806379ba50971161011457806379ba50971461026f57806385572ffb146102775780638da5cb5b1461028a57600080fd5b80632b6e5d631461021d578063665ed5371461025c57600080fd5b8063181f5a7711610160578063181f5a77146101b95780631892b906146102025780632874d8bf1461021557600080fd5b806301ffc9a71461017c57806316c38b3c146101a4575b600080fd5b61018f61018a366004610f0b565b610400565b60405190151581526020015b60405180910390f35b6101b76101b2366004610f54565b610499565b005b6101f56040518060400160405280601881526020017f53656c6646756e64656450696e67506f6e6720312e352e30000000000000000081525081565b60405161019b9190610fda565b6101b761021036600461100a565b6104eb565b6101b7610546565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101b761026a366004610f54565b610582565b6101b761060e565b6101b7610285366004611025565b610710565b60005473ffffffffffffffffffffffffffffffffffffffff16610237565b6101b76102b6366004611060565b610795565b6101b76102c936600461109b565b610977565b60035474010000000000000000000000000000000000000000900460ff1661018f565b7f0000000000000000000000000000000000000000000000000000000000000000610237565b60025474010000000000000000000000000000000000000000900460ff1661018f565b6101b76103483660046110b8565b6109c6565b60015474010000000000000000000000000000000000000000900467ffffffffffffffff1660405167ffffffffffffffff909116815260200161019b565b60035473ffffffffffffffffffffffffffffffffffffffff16610237565b6003547501000000000000000000000000000000000000000000900460ff1660405160ff909116815260200161019b565b6101b76103e83660046110ef565b610a68565b6101b76103fb36600461109b565b610aeb565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f85572ffb00000000000000000000000000000000000000000000000000000000148061049357507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b6104a1610afc565b6002805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b6104f3610afc565b6001805467ffffffffffffffff90921674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff909216919091179055565b61054e610afc565b600280547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556105806001610b7d565b565b61058a610afc565b6003805482151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff9091161790556040517f05a3fef9935c9013a24c6193df2240d34fcf6b0ebf8786b85efe8401d696cdd99061060390831515815260200190565b60405180910390a150565b60015473ffffffffffffffffffffffffffffffffffffffff163314610694576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610781576040517fd7f7333400000000000000000000000000000000000000000000000000000000815233600482015260240161068b565b61079261078d82611317565b610dc0565b50565b6003547501000000000000000000000000000000000000000000900460ff1615806107dd57506003547501000000000000000000000000000000000000000000900460ff1681105b156107e55750565b600354600190610811907501000000000000000000000000000000000000000000900460ff16836113c4565b11610792577f00000000000000000000000000000000000000000000000000000000000000006001546040517fa8d87a3b0000000000000000000000000000000000000000000000000000000081527401000000000000000000000000000000000000000090910467ffffffffffffffff16600482015273ffffffffffffffffffffffffffffffffffffffff919091169063a8d87a3b90602401602060405180830381865afa1580156108c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ec91906113ff565b73ffffffffffffffffffffffffffffffffffffffff1663eff7cc486040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561093357600080fd5b505af1158015610947573d6000803e3d6000fd5b50506040517f302777af5d26fab9dd5120c5f1307c65193ebc51daf33244ada4365fab10602c925060009150a150565b61097f610afc565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6109ce610afc565b6001805467ffffffffffffffff90931674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff909316929092179091556002805473ffffffffffffffffffffffffffffffffffffffff9092167fffffffffffffffffffffffff0000000000000000000000000000000000000000909216919091179055565b610a70610afc565b600380547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff16750100000000000000000000000000000000000000000060ff8416908102919091179091556040519081527f4768dbf8645b24c54f2887651545d24f748c0d0d1d4c689eb810fb19f0befcf390602001610603565b610af3610afc565b61079281610e16565b60005473ffffffffffffffffffffffffffffffffffffffff163314610580576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161068b565b80600116600103610bc0576040518181527f48257dc961b6f792c2b78a080dacfed693b660960a702de21cee364e20270e2f9060200160405180910390a1610bf4565b6040518181527f58b69f57828e6962d216502094c54f6562f3bf082ba758966c3454f9e37b15259060200160405180910390a15b610bfd81610795565b6040805160a0810190915260025473ffffffffffffffffffffffffffffffffffffffff1660c08201526000908060e08101604051602081830303815290604052815260200183604051602001610c5591815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905281526020016000604051908082528060200260200182016040528015610ccf57816020015b6040805180820190915260008082526020820152815260200190600190039081610ca85790505b50815260035473ffffffffffffffffffffffffffffffffffffffff16602080830191909152604080519182018152600082529091015290507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166396f4e9f9600160149054906101000a900467ffffffffffffffff16836040518363ffffffff1660e01b8152600401610d7892919061141c565b6020604051808303816000875af1158015610d97573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dbb9190611531565b505050565b60008160600151806020019051810190610dda9190611531565b60025490915074010000000000000000000000000000000000000000900460ff16610e1257610e12610e0d82600161154a565b610b7d565b5050565b3373ffffffffffffffffffffffffffffffffffffffff821603610e95576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161068b565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215610f1d57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610f4d57600080fd5b9392505050565b600060208284031215610f6657600080fd5b81358015158114610f4d57600080fd5b6000815180845260005b81811015610f9c57602081850181015186830182015201610f80565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610f4d6020830184610f76565b803567ffffffffffffffff8116811461100557600080fd5b919050565b60006020828403121561101c57600080fd5b610f4d82610fed565b60006020828403121561103757600080fd5b813567ffffffffffffffff81111561104e57600080fd5b820160a08185031215610f4d57600080fd5b60006020828403121561107257600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461079257600080fd5b6000602082840312156110ad57600080fd5b8135610f4d81611079565b600080604083850312156110cb57600080fd5b6110d483610fed565b915060208301356110e481611079565b809150509250929050565b60006020828403121561110157600080fd5b813560ff81168114610f4d57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561116457611164611112565b60405290565b60405160a0810167ffffffffffffffff8111828210171561116457611164611112565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156111d4576111d4611112565b604052919050565b600082601f8301126111ed57600080fd5b813567ffffffffffffffff81111561120757611207611112565b61123860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161118d565b81815284602083860101111561124d57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261127b57600080fd5b8135602067ffffffffffffffff82111561129757611297611112565b6112a5818360051b0161118d565b82815260069290921b840181019181810190868411156112c457600080fd5b8286015b8481101561130c57604081890312156112e15760008081fd5b6112e9611141565b81356112f481611079565b815281850135858201528352918301916040016112c8565b509695505050505050565b600060a0823603121561132957600080fd5b61133161116a565b8235815261134160208401610fed565b6020820152604083013567ffffffffffffffff8082111561136157600080fd5b61136d368387016111dc565b6040840152606085013591508082111561138657600080fd5b611392368387016111dc565b606084015260808501359150808211156113ab57600080fd5b506113b83682860161126a565b60808301525092915050565b6000826113fa577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500690565b60006020828403121561141157600080fd5b8151610f4d81611079565b6000604067ffffffffffffffff851683526020604081850152845160a0604086015261144b60e0860182610f76565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808784030160608801526114868383610f76565b6040890151888203830160808a01528051808352908601945060009350908501905b808410156114e7578451805173ffffffffffffffffffffffffffffffffffffffff168352860151868301529385019360019390930192908601906114a8565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a015295506115238187610f76565b9a9950505050505050505050565b60006020828403121561154357600080fd5b5051919050565b80820180821115610493577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea164736f6c6343000818000a",
}

var SelfFundedPingPongABI = SelfFundedPingPongMetaData.ABI

var SelfFundedPingPongBin = SelfFundedPingPongMetaData.Bin

func DeploySelfFundedPingPong(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address, feeToken common.Address, roundTripsBeforeFunding uint8) (common.Address, *CustomTransaction, *SelfFundedPingPong, error) {
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
		return DeployZkSyncSelfFundedPingPong(auth, backend, router, feeToken, roundTripsBeforeFunding)
	}

	parsed, err := SelfFundedPingPongMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SelfFundedPingPongBin), backend, router, feeToken, roundTripsBeforeFunding)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, &CustomTransaction{Transaction: tx, CustomHash: tx.Hash()}, &SelfFundedPingPong{address: address, abi: *parsed, SelfFundedPingPongCaller: SelfFundedPingPongCaller{contract: contract}, SelfFundedPingPongTransactor: SelfFundedPingPongTransactor{contract: contract}, SelfFundedPingPongFilterer: SelfFundedPingPongFilterer{contract: contract}}, nil
}

type SelfFundedPingPong struct {
	address common.Address
	abi     abi.ABI
	SelfFundedPingPongCaller
	SelfFundedPingPongTransactor
	SelfFundedPingPongFilterer
}

type SelfFundedPingPongCaller struct {
	contract *bind.BoundContract
}

type SelfFundedPingPongTransactor struct {
	contract *bind.BoundContract
}

type SelfFundedPingPongFilterer struct {
	contract *bind.BoundContract
}

type SelfFundedPingPongSession struct {
	Contract     *SelfFundedPingPong
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type SelfFundedPingPongCallerSession struct {
	Contract *SelfFundedPingPongCaller
	CallOpts bind.CallOpts
}

type SelfFundedPingPongTransactorSession struct {
	Contract     *SelfFundedPingPongTransactor
	TransactOpts bind.TransactOpts
}

type SelfFundedPingPongRaw struct {
	Contract *SelfFundedPingPong
}

type SelfFundedPingPongCallerRaw struct {
	Contract *SelfFundedPingPongCaller
}

type SelfFundedPingPongTransactorRaw struct {
	Contract *SelfFundedPingPongTransactor
}

func NewSelfFundedPingPong(address common.Address, backend bind.ContractBackend) (*SelfFundedPingPong, error) {
	abi, err := abi.JSON(strings.NewReader(SelfFundedPingPongABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindSelfFundedPingPong(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPong{address: address, abi: abi, SelfFundedPingPongCaller: SelfFundedPingPongCaller{contract: contract}, SelfFundedPingPongTransactor: SelfFundedPingPongTransactor{contract: contract}, SelfFundedPingPongFilterer: SelfFundedPingPongFilterer{contract: contract}}, nil
}

func NewSelfFundedPingPongCaller(address common.Address, caller bind.ContractCaller) (*SelfFundedPingPongCaller, error) {
	contract, err := bindSelfFundedPingPong(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongCaller{contract: contract}, nil
}

func NewSelfFundedPingPongTransactor(address common.Address, transactor bind.ContractTransactor) (*SelfFundedPingPongTransactor, error) {
	contract, err := bindSelfFundedPingPong(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongTransactor{contract: contract}, nil
}

func NewSelfFundedPingPongFilterer(address common.Address, filterer bind.ContractFilterer) (*SelfFundedPingPongFilterer, error) {
	contract, err := bindSelfFundedPingPong(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongFilterer{contract: contract}, nil
}

func bindSelfFundedPingPong(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SelfFundedPingPongMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_SelfFundedPingPong *SelfFundedPingPongRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfFundedPingPong.Contract.SelfFundedPingPongCaller.contract.Call(opts, result, method, params...)
}

func (_SelfFundedPingPong *SelfFundedPingPongRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SelfFundedPingPongTransactor.contract.Transfer(opts)
}

func (_SelfFundedPingPong *SelfFundedPingPongRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SelfFundedPingPongTransactor.contract.Transact(opts, method, params...)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfFundedPingPong.Contract.contract.Call(opts, result, method, params...)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.contract.Transfer(opts)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.contract.Transact(opts, method, params...)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) GetCountIncrBeforeFunding(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "getCountIncrBeforeFunding")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) GetCountIncrBeforeFunding() (uint8, error) {
	return _SelfFundedPingPong.Contract.GetCountIncrBeforeFunding(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) GetCountIncrBeforeFunding() (uint8, error) {
	return _SelfFundedPingPong.Contract.GetCountIncrBeforeFunding(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) GetCounterpartAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "getCounterpartAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) GetCounterpartAddress() (common.Address, error) {
	return _SelfFundedPingPong.Contract.GetCounterpartAddress(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) GetCounterpartAddress() (common.Address, error) {
	return _SelfFundedPingPong.Contract.GetCounterpartAddress(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) GetCounterpartChainSelector(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "getCounterpartChainSelector")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) GetCounterpartChainSelector() (uint64, error) {
	return _SelfFundedPingPong.Contract.GetCounterpartChainSelector(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) GetCounterpartChainSelector() (uint64, error) {
	return _SelfFundedPingPong.Contract.GetCounterpartChainSelector(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) GetFeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "getFeeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) GetFeeToken() (common.Address, error) {
	return _SelfFundedPingPong.Contract.GetFeeToken(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) GetFeeToken() (common.Address, error) {
	return _SelfFundedPingPong.Contract.GetFeeToken(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) GetOutOfOrderExecution(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "getOutOfOrderExecution")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) GetOutOfOrderExecution() (bool, error) {
	return _SelfFundedPingPong.Contract.GetOutOfOrderExecution(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) GetOutOfOrderExecution() (bool, error) {
	return _SelfFundedPingPong.Contract.GetOutOfOrderExecution(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) GetRouter() (common.Address, error) {
	return _SelfFundedPingPong.Contract.GetRouter(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) GetRouter() (common.Address, error) {
	return _SelfFundedPingPong.Contract.GetRouter(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) IsPaused() (bool, error) {
	return _SelfFundedPingPong.Contract.IsPaused(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) IsPaused() (bool, error) {
	return _SelfFundedPingPong.Contract.IsPaused(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) Owner() (common.Address, error) {
	return _SelfFundedPingPong.Contract.Owner(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) Owner() (common.Address, error) {
	return _SelfFundedPingPong.Contract.Owner(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SelfFundedPingPong.Contract.SupportsInterface(&_SelfFundedPingPong.CallOpts, interfaceId)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SelfFundedPingPong.Contract.SupportsInterface(&_SelfFundedPingPong.CallOpts, interfaceId)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) TypeAndVersion() (string, error) {
	return _SelfFundedPingPong.Contract.TypeAndVersion(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) TypeAndVersion() (string, error) {
	return _SelfFundedPingPong.Contract.TypeAndVersion(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "acceptOwnership")
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) AcceptOwnership() (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.AcceptOwnership(&_SelfFundedPingPong.TransactOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.AcceptOwnership(&_SelfFundedPingPong.TransactOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "ccipReceive", message)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.CcipReceive(&_SelfFundedPingPong.TransactOpts, message)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.CcipReceive(&_SelfFundedPingPong.TransactOpts, message)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) FundPingPong(opts *bind.TransactOpts, pingPongCount *big.Int) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "fundPingPong", pingPongCount)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) FundPingPong(pingPongCount *big.Int) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.FundPingPong(&_SelfFundedPingPong.TransactOpts, pingPongCount)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) FundPingPong(pingPongCount *big.Int) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.FundPingPong(&_SelfFundedPingPong.TransactOpts, pingPongCount)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) SetCountIncrBeforeFunding(opts *bind.TransactOpts, countIncrBeforeFunding uint8) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "setCountIncrBeforeFunding", countIncrBeforeFunding)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) SetCountIncrBeforeFunding(countIncrBeforeFunding uint8) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetCountIncrBeforeFunding(&_SelfFundedPingPong.TransactOpts, countIncrBeforeFunding)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) SetCountIncrBeforeFunding(countIncrBeforeFunding uint8) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetCountIncrBeforeFunding(&_SelfFundedPingPong.TransactOpts, countIncrBeforeFunding)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) SetCounterpart(opts *bind.TransactOpts, counterpartChainSelector uint64, counterpartAddress common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "setCounterpart", counterpartChainSelector, counterpartAddress)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) SetCounterpart(counterpartChainSelector uint64, counterpartAddress common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetCounterpart(&_SelfFundedPingPong.TransactOpts, counterpartChainSelector, counterpartAddress)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) SetCounterpart(counterpartChainSelector uint64, counterpartAddress common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetCounterpart(&_SelfFundedPingPong.TransactOpts, counterpartChainSelector, counterpartAddress)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) SetCounterpartAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "setCounterpartAddress", addr)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) SetCounterpartAddress(addr common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetCounterpartAddress(&_SelfFundedPingPong.TransactOpts, addr)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) SetCounterpartAddress(addr common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetCounterpartAddress(&_SelfFundedPingPong.TransactOpts, addr)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) SetCounterpartChainSelector(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "setCounterpartChainSelector", chainSelector)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) SetCounterpartChainSelector(chainSelector uint64) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetCounterpartChainSelector(&_SelfFundedPingPong.TransactOpts, chainSelector)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) SetCounterpartChainSelector(chainSelector uint64) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetCounterpartChainSelector(&_SelfFundedPingPong.TransactOpts, chainSelector)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) SetOutOfOrderExecution(opts *bind.TransactOpts, outOfOrderExecution bool) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "setOutOfOrderExecution", outOfOrderExecution)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) SetOutOfOrderExecution(outOfOrderExecution bool) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetOutOfOrderExecution(&_SelfFundedPingPong.TransactOpts, outOfOrderExecution)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) SetOutOfOrderExecution(outOfOrderExecution bool) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetOutOfOrderExecution(&_SelfFundedPingPong.TransactOpts, outOfOrderExecution)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) SetPaused(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "setPaused", pause)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) SetPaused(pause bool) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetPaused(&_SelfFundedPingPong.TransactOpts, pause)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) SetPaused(pause bool) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetPaused(&_SelfFundedPingPong.TransactOpts, pause)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) StartPingPong(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "startPingPong")
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) StartPingPong() (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.StartPingPong(&_SelfFundedPingPong.TransactOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) StartPingPong() (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.StartPingPong(&_SelfFundedPingPong.TransactOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "transferOwnership", to)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.TransferOwnership(&_SelfFundedPingPong.TransactOpts, to)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.TransferOwnership(&_SelfFundedPingPong.TransactOpts, to)
}

type SelfFundedPingPongCountIncrBeforeFundingSetIterator struct {
	Event *SelfFundedPingPongCountIncrBeforeFundingSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SelfFundedPingPongCountIncrBeforeFundingSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SelfFundedPingPongCountIncrBeforeFundingSet)
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
		it.Event = new(SelfFundedPingPongCountIncrBeforeFundingSet)
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

func (it *SelfFundedPingPongCountIncrBeforeFundingSetIterator) Error() error {
	return it.fail
}

func (it *SelfFundedPingPongCountIncrBeforeFundingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SelfFundedPingPongCountIncrBeforeFundingSet struct {
	CountIncrBeforeFunding uint8
	Raw                    types.Log
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) FilterCountIncrBeforeFundingSet(opts *bind.FilterOpts) (*SelfFundedPingPongCountIncrBeforeFundingSetIterator, error) {

	logs, sub, err := _SelfFundedPingPong.contract.FilterLogs(opts, "CountIncrBeforeFundingSet")
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongCountIncrBeforeFundingSetIterator{contract: _SelfFundedPingPong.contract, event: "CountIncrBeforeFundingSet", logs: logs, sub: sub}, nil
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) WatchCountIncrBeforeFundingSet(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongCountIncrBeforeFundingSet) (event.Subscription, error) {

	logs, sub, err := _SelfFundedPingPong.contract.WatchLogs(opts, "CountIncrBeforeFundingSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SelfFundedPingPongCountIncrBeforeFundingSet)
				if err := _SelfFundedPingPong.contract.UnpackLog(event, "CountIncrBeforeFundingSet", log); err != nil {
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

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) ParseCountIncrBeforeFundingSet(log types.Log) (*SelfFundedPingPongCountIncrBeforeFundingSet, error) {
	event := new(SelfFundedPingPongCountIncrBeforeFundingSet)
	if err := _SelfFundedPingPong.contract.UnpackLog(event, "CountIncrBeforeFundingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SelfFundedPingPongFundedIterator struct {
	Event *SelfFundedPingPongFunded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SelfFundedPingPongFundedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SelfFundedPingPongFunded)
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
		it.Event = new(SelfFundedPingPongFunded)
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

func (it *SelfFundedPingPongFundedIterator) Error() error {
	return it.fail
}

func (it *SelfFundedPingPongFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SelfFundedPingPongFunded struct {
	Raw types.Log
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) FilterFunded(opts *bind.FilterOpts) (*SelfFundedPingPongFundedIterator, error) {

	logs, sub, err := _SelfFundedPingPong.contract.FilterLogs(opts, "Funded")
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongFundedIterator{contract: _SelfFundedPingPong.contract, event: "Funded", logs: logs, sub: sub}, nil
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) WatchFunded(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongFunded) (event.Subscription, error) {

	logs, sub, err := _SelfFundedPingPong.contract.WatchLogs(opts, "Funded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SelfFundedPingPongFunded)
				if err := _SelfFundedPingPong.contract.UnpackLog(event, "Funded", log); err != nil {
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

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) ParseFunded(log types.Log) (*SelfFundedPingPongFunded, error) {
	event := new(SelfFundedPingPongFunded)
	if err := _SelfFundedPingPong.contract.UnpackLog(event, "Funded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SelfFundedPingPongOutOfOrderExecutionChangeIterator struct {
	Event *SelfFundedPingPongOutOfOrderExecutionChange

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SelfFundedPingPongOutOfOrderExecutionChangeIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SelfFundedPingPongOutOfOrderExecutionChange)
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
		it.Event = new(SelfFundedPingPongOutOfOrderExecutionChange)
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

func (it *SelfFundedPingPongOutOfOrderExecutionChangeIterator) Error() error {
	return it.fail
}

func (it *SelfFundedPingPongOutOfOrderExecutionChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SelfFundedPingPongOutOfOrderExecutionChange struct {
	IsOutOfOrder bool
	Raw          types.Log
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) FilterOutOfOrderExecutionChange(opts *bind.FilterOpts) (*SelfFundedPingPongOutOfOrderExecutionChangeIterator, error) {

	logs, sub, err := _SelfFundedPingPong.contract.FilterLogs(opts, "OutOfOrderExecutionChange")
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongOutOfOrderExecutionChangeIterator{contract: _SelfFundedPingPong.contract, event: "OutOfOrderExecutionChange", logs: logs, sub: sub}, nil
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) WatchOutOfOrderExecutionChange(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongOutOfOrderExecutionChange) (event.Subscription, error) {

	logs, sub, err := _SelfFundedPingPong.contract.WatchLogs(opts, "OutOfOrderExecutionChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SelfFundedPingPongOutOfOrderExecutionChange)
				if err := _SelfFundedPingPong.contract.UnpackLog(event, "OutOfOrderExecutionChange", log); err != nil {
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

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) ParseOutOfOrderExecutionChange(log types.Log) (*SelfFundedPingPongOutOfOrderExecutionChange, error) {
	event := new(SelfFundedPingPongOutOfOrderExecutionChange)
	if err := _SelfFundedPingPong.contract.UnpackLog(event, "OutOfOrderExecutionChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SelfFundedPingPongOwnershipTransferRequestedIterator struct {
	Event *SelfFundedPingPongOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SelfFundedPingPongOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SelfFundedPingPongOwnershipTransferRequested)
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
		it.Event = new(SelfFundedPingPongOwnershipTransferRequested)
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

func (it *SelfFundedPingPongOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *SelfFundedPingPongOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SelfFundedPingPongOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SelfFundedPingPongOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SelfFundedPingPong.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongOwnershipTransferRequestedIterator{contract: _SelfFundedPingPong.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SelfFundedPingPong.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SelfFundedPingPongOwnershipTransferRequested)
				if err := _SelfFundedPingPong.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) ParseOwnershipTransferRequested(log types.Log) (*SelfFundedPingPongOwnershipTransferRequested, error) {
	event := new(SelfFundedPingPongOwnershipTransferRequested)
	if err := _SelfFundedPingPong.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SelfFundedPingPongOwnershipTransferredIterator struct {
	Event *SelfFundedPingPongOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SelfFundedPingPongOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SelfFundedPingPongOwnershipTransferred)
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
		it.Event = new(SelfFundedPingPongOwnershipTransferred)
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

func (it *SelfFundedPingPongOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *SelfFundedPingPongOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SelfFundedPingPongOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SelfFundedPingPongOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SelfFundedPingPong.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongOwnershipTransferredIterator{contract: _SelfFundedPingPong.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SelfFundedPingPong.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SelfFundedPingPongOwnershipTransferred)
				if err := _SelfFundedPingPong.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) ParseOwnershipTransferred(log types.Log) (*SelfFundedPingPongOwnershipTransferred, error) {
	event := new(SelfFundedPingPongOwnershipTransferred)
	if err := _SelfFundedPingPong.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SelfFundedPingPongPingIterator struct {
	Event *SelfFundedPingPongPing

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SelfFundedPingPongPingIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SelfFundedPingPongPing)
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
		it.Event = new(SelfFundedPingPongPing)
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

func (it *SelfFundedPingPongPingIterator) Error() error {
	return it.fail
}

func (it *SelfFundedPingPongPingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SelfFundedPingPongPing struct {
	PingPongCount *big.Int
	Raw           types.Log
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) FilterPing(opts *bind.FilterOpts) (*SelfFundedPingPongPingIterator, error) {

	logs, sub, err := _SelfFundedPingPong.contract.FilterLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongPingIterator{contract: _SelfFundedPingPong.contract, event: "Ping", logs: logs, sub: sub}, nil
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) WatchPing(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongPing) (event.Subscription, error) {

	logs, sub, err := _SelfFundedPingPong.contract.WatchLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SelfFundedPingPongPing)
				if err := _SelfFundedPingPong.contract.UnpackLog(event, "Ping", log); err != nil {
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

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) ParsePing(log types.Log) (*SelfFundedPingPongPing, error) {
	event := new(SelfFundedPingPongPing)
	if err := _SelfFundedPingPong.contract.UnpackLog(event, "Ping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SelfFundedPingPongPongIterator struct {
	Event *SelfFundedPingPongPong

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SelfFundedPingPongPongIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SelfFundedPingPongPong)
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
		it.Event = new(SelfFundedPingPongPong)
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

func (it *SelfFundedPingPongPongIterator) Error() error {
	return it.fail
}

func (it *SelfFundedPingPongPongIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SelfFundedPingPongPong struct {
	PingPongCount *big.Int
	Raw           types.Log
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) FilterPong(opts *bind.FilterOpts) (*SelfFundedPingPongPongIterator, error) {

	logs, sub, err := _SelfFundedPingPong.contract.FilterLogs(opts, "Pong")
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongPongIterator{contract: _SelfFundedPingPong.contract, event: "Pong", logs: logs, sub: sub}, nil
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) WatchPong(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongPong) (event.Subscription, error) {

	logs, sub, err := _SelfFundedPingPong.contract.WatchLogs(opts, "Pong")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SelfFundedPingPongPong)
				if err := _SelfFundedPingPong.contract.UnpackLog(event, "Pong", log); err != nil {
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

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) ParsePong(log types.Log) (*SelfFundedPingPongPong, error) {
	event := new(SelfFundedPingPongPong)
	if err := _SelfFundedPingPong.contract.UnpackLog(event, "Pong", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var SelfFundedPingPongZkBin string = ("0x0003000000000002000500000000000200020000000103550000006003100270000001ed0030019d000001ed0330019700000001002001900000001f0000c13d0000008002000039000000400020043f000000040030008c000003de0000413d000000000201043b000000e002200270000001fe0020009c000000560000a13d000001ff0020009c0000009f0000a13d000002000020009c000000b90000213d000002040020009c000001760000613d000002050020009c000001aa0000613d000002060020009c000003de0000c13d0000000001000416000000000001004b000003de0000c13d0000000301000039000003160000013d000000a004000039000000400040043f0000000002000416000000000002004b000003de0000c13d0000001f02300039000001ee02200197000000a002200039000000400020043f0000001f0530018f000001ef06300198000000a002600039000000310000613d000000000701034f000000007807043c0000000004840436000000000024004b0000002d0000c13d000000000005004b0000003e0000613d000000000161034f0000000304500210000000000502043300000000054501cf000000000545022f000000000101043b0000010004400089000000000141022f00000000014101cf000000000151019f0000000000120435000000600030008c000003de0000413d000000a00100043d000001f00010009c000003de0000213d000000c00200043d000001f00020009c000003de0000213d000000e00300043d000300000003001d000000ff0030008c000003de0000213d000001f0031001980000031b0000c13d000000400100043d000001fc02000041000000000021043500000004021000390000000000020435000001ed0010009c000001ed010080410000004001100210000001fd011001c7000007b3000104300000020d0020009c000000ac0000213d000002140020009c000000ed0000a13d000002150020009c000001460000613d000002160020009c000001b40000613d000002170020009c000003de0000c13d0000000001000416000000000001004b000003de0000c13d000000000100041a000001f0011001970000000002000411000000000012004b000003300000c13d0000000202000039000000000102041a000001f201100197000000000012041b0000000103000039000000800030043f0000000001000414000001ed0010009c000001ed01008041000000c0011002100000021f011001c70000800d020000390000022d0400004107b107a70000040f0000000100200190000003de0000613d0000000301000039000000000101041a0000021e01100197000002380010009c000004610000c13d0000000101000039000000000101041a000000400300043d00000239020000410000000000230435000000a001100270000001f501100197000300000003001d00000004023000390000000000120435000002250100004100000000001004430000000001000412000000040010044300000024000004430000000001000414000001ed0010009c000001ed01008041000000c00110021000000227011001c7000080050200003907b107ac0000040f0000000100200190000005240000613d000000000201043b0000000001000414000001f002200197000000040020008c000003ef0000c13d0000000103000031000000200030008c00000020040000390000000004034019000004180000013d000002070020009c000000d90000a13d000002080020009c0000012d0000613d000002090020009c000001380000613d0000020a0020009c000003de0000c13d0000000001000416000000000001004b000003de0000c13d0000000201000039000001310000013d0000020e0020009c000001060000a13d0000020f0020009c0000015a0000613d000002100020009c000001c90000613d000002110020009c000003de0000c13d0000000001000416000000000001004b000003de0000c13d000000000100041a000003170000013d000002010020009c000001960000613d000002020020009c000002dd0000613d000002030020009c000003de0000c13d000000240030008c000003de0000413d0000000002000416000000000002004b000003de0000c13d0000000401100370000000000601043b000001f00060009c000003de0000213d000000000100041a000001f0011001970000000005000411000000000015004b000003300000c13d000000000056004b000003ce0000c13d000001f901000041000000800010043f0000002001000039000000840010043f0000001701000039000000a40010043f0000021c01000041000000c40010043f0000021d01000041000007b3000104300000020b0020009c000001a00000613d0000020c0020009c000003de0000c13d000000240030008c000003de0000413d0000000002000416000000000002004b000003de0000c13d0000000401100370000000000101043b000300000001001d000001f00010009c000003de0000213d07b107760000040f0000000201000039000000000201041a000001f1022001970000000303000029000001920000013d000002180020009c000003010000613d000002190020009c000003de0000c13d000000240030008c000003de0000413d0000000002000416000000000002004b000003de0000c13d0000000401100370000000000201043b000000000002004b0000000001000039000000010100c039000300000002001d000000000012004b000003de0000c13d07b107760000040f000000030000006b0000000001000019000002350100c0410000000202000039000000000302041a000001f203300197000001c50000013d000002120020009c000003120000613d000002130020009c000003de0000c13d000000240030008c000003de0000413d0000000002000416000000000002004b000003de0000c13d0000000401100370000000000101043b000000000001004b0000000002000039000000010200c039000000000021004b000003de0000c13d000000000200041a000001f0022001970000000003000411000000000023004b000003300000c13d0000000302000039000000000302041a000001f203300197000000000001004b0000000004000019000002350400c041000000000343019f000000000032041b000000800010043f0000000001000414000001ed0010009c000001ed01008041000000c0011002100000021f011001c70000800d0200003900000001030000390000023604000041000002fc0000013d0000000001000416000000000001004b000003de0000c13d0000000301000039000000000101041a00000224001001980000000001000039000000010100c039000000800010043f0000022101000041000007b20001042e0000000001000416000000000001004b000003de0000c13d0000000001000412000500000001001d000400000000003d000080050100003900000044030000390000000004000415000000050440008a0000000504400210000002250200004107b1078e0000040f000003170000013d0000000001000416000000000001004b000003de0000c13d000000c001000039000000400010043f0000001801000039000000800010043f0000023f01000041000000a00010043f0000002001000039000000c00010043f0000008001000039000000e00200003907b106800000040f000000c00110008a000001ed0010009c000001ed01008041000000600110021000000240011001c7000007b20001042e0000000001000416000000000001004b000003de0000c13d0000000101000039000000000201041a000001f0032001970000000006000411000000000036004b000003550000c13d000000000300041a000001f104300197000000000464019f000000000040041b000001f102200197000000000021041b0000000001000414000001f005300197000001ed0010009c000001ed01008041000000c0011002100000021a011001c70000800d020000390000000303000039000002340400004107b107a70000040f0000000100200190000002ff0000c13d000003de0000013d000000440030008c000003de0000413d0000000002000416000000000002004b000003de0000c13d0000000402100370000000000202043b000300000002001d000001f50020009c000003de0000213d0000002401100370000000000101043b000200000001001d000001f00010009c000003de0000213d07b107760000040f0000000301000029000000a00110021000000222011001970000000102000039000000000302041a0000022303300197000000000113019f000000000012041b0000000201000039000000000201041a000001f1022001970000000203000029000000000232019f000000000021041b0000000001000019000007b20001042e0000000001000416000000000001004b000003de0000c13d0000000301000039000000000101041a000000a801100270000000ff0110018f000000800010043f0000022101000041000007b20001042e000000240030008c000003de0000413d0000000002000416000000000002004b000003de0000c13d0000000401100370000000000101043b07b106920000040f0000000001000019000007b20001042e0000000001000416000000000001004b000003de0000c13d0000000101000039000000000101041a000000a001100270000001f501100197000000800010043f0000022101000041000007b20001042e000000240030008c000003de0000413d0000000002000416000000000002004b000003de0000c13d0000000401100370000000000101043b000300000001001d000001f50010009c000003de0000213d07b107760000040f0000000301000029000000a00110021000000222011001970000000102000039000000000302041a0000022303300197000000000113019f000000000012041b0000000001000019000007b20001042e000000240030008c000003de0000413d0000000002000416000000000002004b000003de0000c13d0000000401100370000000000101043b000300000001001d000001f50010009c000003de0000213d000000030130006a000002260010009c000003de0000213d000000a40010008c000003de0000413d000002250100004100000000001004430000000001000412000000040010044300000024000004430000000001000414000001ed0010009c000001ed01008041000000c00110021000000227011001c7000080050200003907b107ac0000040f0000000100200190000005240000613d000000400200043d000000000101043b000001f0011001970000000003000411000000000013004b000003e00000c13d000002280020009c0000048b0000213d000000a001200039000000400010043f000000030500002900000004035000390000000201000367000000000331034f000000000303043b00000000043204360000002403500039000000000531034f000000000505043b000001f50050009c000003de0000213d00000000005404350000002005300039000000000351034f000000000303043b000001f50030009c000003de0000213d00000003093000290000002306900039000002290760019700000000030000310000022904300197000000000847013f000000000047004b00000000070000190000022907004041000000000036004b00000000060000190000022906008041000002290080009c000000000706c019000000000007004b000003de0000c13d000000040a9000390000000006a1034f000000000606043b000001f50060009c0000048b0000213d0000001f0760003900000244077001970000003f077000390000024408700197000000400700043d0000000008870019000000000078004b000000000b000039000000010b004039000001f50080009c0000048b0000213d0000000100b001900000048b0000c13d000000400080043f000000000867043600000000096900190000002409900039000000000039004b000003de0000213d0000002009a00039000000000a91034f000002440b6001980000001f0c60018f0000000009b80019000002350000613d000000000d0a034f000000000e08001900000000df0d043c000000000efe043600000000009e004b000002310000c13d00000000000c004b000002420000613d000000000aba034f000000030bc00210000000000c090433000000000cbc01cf000000000cbc022f000000000a0a043b000001000bb00089000000000aba022f000000000aba01cf000000000aca019f0000000000a9043500000000066800190000000000060435000000400620003900000000007604350000002006500039000000000561034f000000000505043b000001f50050009c000003de0000213d00000003095000290000002305900039000000000035004b000000000700001900000229070080410000022905500197000000000845013f000000000045004b00000000050000190000022905004041000002290080009c000000000507c019000000000005004b000003de0000c13d000000040a9000390000000005a1034f000000000505043b000001f50050009c0000048b0000213d0000001f0750003900000244077001970000003f077000390000024408700197000000400700043d0000000008870019000000000078004b000000000b000039000000010b004039000001f50080009c0000048b0000213d0000000100b001900000048b0000c13d000000400080043f000000000857043600000000095900190000002409900039000000000039004b000003de0000213d0000002009a00039000000000a91034f000002440b5001980000001f0c50018f0000000009b800190000027d0000613d000000000d0a034f000000000e08001900000000df0d043c000000000efe043600000000009e004b000002790000c13d00000000000c004b0000028a0000613d000000000aba034f000000030bc00210000000000c090433000000000cbc01cf000000000cbc022f000000000a0a043b000001000bb00089000000000aba022f000000000aba01cf000000000aca019f0000000000a9043500000000055800190000000000050435000000600520003900000000007504350000002006600039000000000661034f000000000606043b000001f50060009c000003de0000213d00000003066000290000002307600039000000000037004b000000000800001900000229080080410000022907700197000000000947013f000000000047004b00000000040000190000022904004041000002290090009c000000000408c019000000000004004b000003de0000c13d0000000404600039000000000441034f000000000804043b000001f50080009c0000048b0000213d00000005048002100000003f044000390000022a07400197000000400400043d0000000007740019000000000047004b00000000090000390000000109004039000001f50070009c0000048b0000213d00000001009001900000048b0000c13d000000400070043f0000000000840435000000240660003900000006078002100000000007670019000000000037004b000003de0000213d000000000008004b0000057e0000c13d0000008001200039000000000041043500000000010504330000000012010434000002260020009c000003de0000213d000000200020008c000003de0000413d0000000202000039000000000202041a0000022400200198000002ff0000c13d0000000001010433000000010010003a000003e90000413d0000000102100039000000400100043d0000000000210435000001ed0010009c000001ed010080410000004001100210000300000002001d0000000100200190000005980000c13d0000000002000414000001ed0020009c000001ed02008041000000c002200210000000000112019f0000022c011001c70000800d0200003900000001030000390000022e04000041000005a10000013d000000240030008c000003de0000413d0000000002000416000000000002004b000003de0000c13d0000000401100370000000000101043b000000ff0010008c000003de0000213d000000000200041a000001f0022001970000000003000411000000000023004b000003300000c13d000000ff0210018f000000a8011002100000021e011001970000000303000039000000000403041a000001f704400197000000000114019f000000000013041b000000800020043f0000000001000414000001ed0010009c000001ed01008041000000c0011002100000021f011001c70000800d020000390000000103000039000002200400004107b107a70000040f0000000100200190000003de0000613d0000000001000019000007b20001042e000000240030008c000003de0000413d0000000002000416000000000002004b000003de0000c13d0000000401100370000000000101043b0000024100100198000003de0000c13d000002420010009c00000000020000390000000102006039000002430010009c00000001022061bf000000800020043f0000022101000041000007b20001042e0000000001000416000000000001004b000003de0000c13d0000000201000039000000000101041a000001f001100197000000800010043f0000022101000041000007b20001042e000000800010043f000000400b00043d0000002401b000390000000404b000390000000005000411000000000005004b0000033a0000c13d000001f90200004100000000002b043500000020020000390000000000240435000000180200003900000000002104350000004401b00039000001fa020000410000000000210435000001ed00b0009c000001ed0b0080410000004001b00210000001fb011001c7000007b300010430000001f901000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f0000023701000041000000c40010043f0000021d01000041000007b300010430000001f002200197000000000600041a000001f106600197000000000556019f000000000050041b0000000205000039000000000605041a000001f206600197000000000065041b0000000306000039000000000506041a000001f105500197000000000525019f000000000056041b000001f30500004100000000005b04350000000000340435000000010300008a00000000003104350000000001000414000000040020008c0000035f0000c13d0000000103000031000000200030008c000000200400003900000000040340190000038a0000013d000001f901000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f0000023301000041000000c40010043f0000021d01000041000007b300010430000001ed00b0009c000001ed0300004100000000030b40190000004003300210000001ed0010009c000001ed01008041000000c001100210000000000131019f000001f4011001c700020000000b001d07b107a70000040f000000020b0000290000006003100270000001ed03300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b00190000037a0000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b000003760000c13d000000000006004b000003870000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000003b00000613d0000001f01400039000000600210018f0000000001b20019000000000021004b00000000020000390000000102004039000001f50010009c0000048b0000213d00000001002001900000048b0000c13d000000400010043f000000200030008c000003de0000413d00000000010b0433000000000001004b0000000002000039000000010200c039000000000021004b000003de0000c13d00000003010000290000008000100190000003e90000c13d000000a901100210000001f6011001970000000303000039000000000203041a000001f702200197000000000112019f000000000013041b000000800100043d000001400000044300000160001004430000002001000039000001000010044300000001010000390000012000100443000001f801000041000007b20001042e0000001f0530018f000001ef06300198000000400200043d0000000004620019000003bb0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000003b70000c13d000000000005004b000003c80000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000001ed0020009c000001ed020080410000004002200210000000000112019f000007b3000104300000000101000039000000000201041a000001f102200197000000000262019f000000000021041b0000000001000414000001ed0010009c000001ed01008041000000c0011002100000021a011001c70000800d0200003900000003030000390000021b0400004107b107a70000040f0000000100200190000002ff0000c13d0000000001000019000007b300010430000001fc01000041000000000012043500000004012000390000000000310435000001ed0020009c000001ed020080410000004001200210000001fd011001c7000007b3000104300000023201000041000000000010043f0000001101000039000000040010043f000001fd01000041000007b3000104300000000303000029000001ed0030009c000001ed030080410000004003300210000001ed0010009c000001ed01008041000000c001100210000000000131019f000001fd011001c707b107ac0000040f0000006003100270000001ed03300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000000305700029000004080000613d000000000801034f0000000309000029000000008a08043c0000000009a90436000000000059004b000004040000c13d000000000006004b000004150000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000004910000613d0000001f01400039000000600210018f0000000301200029000000000021004b00000000020000390000000102004039000001f50010009c0000048b0000213d00000001002001900000048b0000c13d000000400010043f000000200030008c000003de0000413d00000003010000290000000001010433000300000001001d000002350010009c000003de0000813d0000023a010000410000000000100443000000030100002900000004001004430000000001000414000001ed0010009c000001ed01008041000000c0011002100000023b011001c7000080020200003907b107ac0000040f0000000100200190000005240000613d000000000101043b000000000001004b000003de0000613d000000400200043d0000023c01000041000200000002001d000000000012043500000000010004140000000302000029000000040020008c000004510000613d0000000202000029000001ed0020009c000001ed020080410000004002200210000001ed0010009c000001ed01008041000000c001100210000000000121019f0000023d011001c7000000030200002907b107a70000040f0000006003100270000101ed0030019d0000000100200190000005710000613d0000000201000029000001f50010009c0000048b0000213d0000000201000029000000400010043f0000000001000414000001ed0010009c000001ed01008041000000c0011002100000021a011001c70000800d0200003900000001030000390000023e0400004107b107a70000040f0000000100200190000003de0000613d0000000201000039000000000101041a000001f001100197000000400200043d00000020032000390000000000130435000000200100003900000000001204350000022b0020009c0000048b0000213d0000004006200039000000400060043f00000060032000390000000104000039000000000043043500000000001604350000022f0020009c0000048b0000213d0000008004200039000000400040043f000002280020009c0000048b0000213d000000a001200039000000400010043f0000000000040435000000400500043d000002280050009c0000048b0000213d0000000301000039000000000101041a000000a003500039000000400030043f000001f003100197000000600150003900000000003104350000004003500039000000000043043500000000042504360000000000640435000000400600043d000002300060009c0000049d0000a13d0000023201000041000000000010043f0000004101000039000000040010043f000001fd01000041000007b3000104300000001f0530018f000001ef06300198000000400200043d0000000004620019000003bb0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000004980000c13d000003bb0000013d0000002002600039000000400020043f0000000000060435000000800250003900000000006204350000000106000039000000000606041a000000400900043d00000024079000390000004008000039000000000087043500000231070000410000000000790435000000a006600270000001f5066001970000000407900039000000000067043500000000050504330000004406900039000000a0070000390000000000760435000000e40890003900000000760504340000000000680435000300000009001d0000010405900039000000000006004b000004c10000613d00000000080000190000000009580019000000000a870019000000000a0a04330000000000a904350000002008800039000000000068004b000004ba0000413d000000000765001900000000000704350000001f066000390000024406600197000000000765001900000003060000290000000005670049000000440550008a00000000040404330000006406600039000000000056043500000000650404340000000004570436000000000005004b000004d80000613d000000000700001900000000084700190000000009760019000000000909043300000000009804350000002007700039000000000057004b000004d10000413d000000000654001900000000000604350000001f055000390000024405500197000000000654001900000003070000290000000004760049000000440540008a00000000040304330000008403700039000000000053043500000000050404330000000003560436000000000005004b000004f30000613d0000000006000019000000200440003900000000070404330000000087070434000001f00770019700000000077304360000000008080433000000000087043500000040033000390000000106600039000000000056004b000004e80000413d0000000001010433000001f0011001970000000305000029000000a40450003900000000001404350000000001530049000000440110008a0000000002020433000000c404500039000000000014043500000000160204340000000005630436000000000006004b000005090000613d000000000200001900000000035200190000000004210019000000000404043300000000004304350000002002200039000000000062004b000005020000413d000200000005001d000100000006001d00000000016500190000000000010435000002250100004100000000001004430000000001000412000000040010044300000024000004430000000001000414000001ed0010009c000001ed01008041000000c00110021000000227011001c7000080050200003907b107ac0000040f0000000100200190000005240000613d000000000201043b0000000001000414000001f002200197000000040020008c000005250000c13d0000000104000031000000200040008c0000002004008039000005570000013d000000000001042f00000001030000290000001f033000390000024403300197000000030500002900000000035300490000000203300029000001ed0030009c000001ed030080410000006003300210000001ed0050009c000001ed0400004100000000040540190000004004400210000000000343019f000001ed0010009c000001ed01008041000000c001100210000000000131019f07b107a70000040f0000006003100270000001ed03300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000000305700029000005470000613d000000000801034f0000000309000029000000008a08043c0000000009a90436000000000059004b000005430000c13d000000000006004b000005540000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000005650000613d0000001f01400039000000600210018f0000000301200029000000000021004b00000000020000390000000102004039000001f50010009c0000048b0000213d00000001002001900000048b0000c13d000000400010043f000000200040008c000002ff0000813d000003de0000013d0000001f0530018f000001ef06300198000000400200043d0000000004620019000003bb0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b0000056c0000c13d000003bb0000013d000001ed033001970000001f0530018f000001ef06300198000000400200043d0000000004620019000003bb0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000005790000c13d000003bb0000013d00000000080400190000000009630049000002260090009c000003de0000213d000000400090008c000003de0000413d000000400900043d0000022b0090009c0000048b0000213d000000400a9000390000004000a0043f000000000a61034f000000000a0a043b000001f000a0009c000003de0000213d0000002008800039000000000aa90436000000200b600039000000000bb1034f000000000b0b043b0000000000ba043500000000009804350000004006600039000000000076004b0000057f0000413d000002bb0000013d0000000002000414000001ed0020009c000001ed02008041000000c002200210000000000112019f0000022c011001c70000800d0200003900000001030000390000022d0400004107b107a70000040f0000000100200190000003de0000613d000000030100002907b106920000040f0000000201000039000000000101041a000001f001100197000000400200043d00000020032000390000000000130435000000200100003900000000001204350000022b0020009c0000048b0000213d0000004006200039000000400060043f00000060032000390000000304000029000000000043043500000000001604350000022f0020009c0000048b0000213d0000008004200039000000400040043f000002280020009c0000048b0000213d000000a001200039000000400010043f0000000000040435000000400500043d000002280050009c0000048b0000213d0000000301000039000000000101041a000000a003500039000000400030043f000001f003100197000000600150003900000000003104350000004003500039000000000043043500000000042504360000000000640435000000400600043d000002300060009c0000048b0000213d0000002002600039000000400020043f0000000000060435000000800250003900000000006204350000000106000039000000000606041a000000400900043d00000024079000390000004008000039000000000087043500000231070000410000000000790435000000a006600270000001f5066001970000000407900039000000000067043500000000050504330000004406900039000000a0070000390000000000760435000000e40690003900000000750504340000000000560435000300000009001d0000010406900039000000000005004b000005f40000613d00000000080000190000000009680019000000000a870019000000000a0a04330000000000a904350000002008800039000000000058004b000005ed0000413d000000000756001900000000000704350000001f055000390000024405500197000000000756001900000003060000290000000005670049000000440550008a00000000040404330000006406600039000000000056043500000000650404340000000004570436000000000005004b0000060b0000613d000000000700001900000000084700190000000009760019000000000909043300000000009804350000002007700039000000000057004b000006040000413d000000000654001900000000000604350000001f055000390000024405500197000000000654001900000003070000290000000004760049000000440540008a00000000040304330000008403700039000000000053043500000000050404330000000003560436000000000005004b000006260000613d0000000006000019000000200440003900000000070404330000000087070434000001f00770019700000000077304360000000008080433000000000087043500000040033000390000000106600039000000000056004b0000061b0000413d0000000001010433000001f0011001970000000305000029000000a40450003900000000001404350000000001530049000000440110008a0000000002020433000000c404500039000000000014043500000000420204340000000001230436000000000002004b0000063c0000613d000000000300001900000000051300190000000006340019000000000606043300000000006504350000002003300039000000000023004b000006350000413d0000000003210019000000000003043500000000030004140000000004000411000000040040008c000005200000613d0000001f022000390000024402200197000000030400002900000000024200490000000001120019000001ed0010009c000001ed010080410000006001100210000001ed0040009c000001ed0200004100000000020440190000004002200210000000000121019f000001ed0030009c000001ed03008041000000c002300210000000000112019f000000000200041107b107a70000040f0000006003100270000001ed03300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000000305700029000006640000613d000000000801034f0000000309000029000000008a08043c0000000009a90436000000000059004b000006600000c13d000000000006004b000006710000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000005570000c13d0000001f0530018f000001ef06300198000000400200043d0000000004620019000003bb0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b0000067b0000c13d000003bb0000013d00000000430104340000000001320436000000000003004b0000068c0000613d000000000200001900000000052100190000000006240019000000000606043300000000006504350000002002200039000000000032004b000006850000413d000000000231001900000000000204350000001f0230003900000244022001970000000001210019000000000001042d00010000000000020000000302000039000000000202041a000000a802200270000000ff0220018f000000010320008a000000000013004b0000072f0000813d00000000102100d9000000010010008c0000072f0000213d0000000101000039000000000101041a000000400300043d00000239020000410000000000230435000000a001100270000001f501100197000100000003001d00000004023000390000000000120435000002250100004100000000001004430000000001000412000000040010044300000024000004430000000001000414000001ed0010009c000001ed01008041000000c00110021000000227011001c7000080050200003907b107ac0000040f0000000100200190000007320000613d000000000201043b0000000001000414000001f002200197000000040020008c000006c00000c13d0000000103000031000000200030008c00000020040000390000000004034019000000010b000029000006ea0000013d0000000103000029000001ed0030009c000001ed030080410000004003300210000001ed0010009c000001ed01008041000000c001100210000000000131019f000001fd011001c707b107ac0000040f000000010b0000290000006003100270000001ed03300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b0019000006da0000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b000006d60000c13d000000000006004b000006e70000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000007390000613d0000001f01400039000000600210018f0000000001b20019000000000021004b00000000020000390000000102004039000001f50010009c000007330000213d0000000100200190000007330000c13d000000400010043f000000200030008c000007300000413d00000000020b0433000002350020009c000007300000813d0000023a010000410000000000100443000100000002001d00000004002004430000000001000414000001ed0010009c000001ed01008041000000c0011002100000023b011001c7000080020200003907b107ac0000040f0000000100200190000007320000613d000000000101043b000000000001004b0000000102000029000007300000613d000000400400043d0000023c0100004100000000001404350000000001000414000000040020008c000007210000613d000001ed0040009c000001ed0300004100000000030440190000004003300210000001ed0010009c000001ed01008041000000c001100210000000000131019f0000023d011001c7000100000004001d07b107a70000040f00000001040000290000006003100270000101ed0030019d0000000100200190000007570000613d000001f50040009c000007330000213d000000400040043f0000000001000414000001ed0010009c000001ed01008041000000c0011002100000021a011001c70000800d0200003900000001030000390000023e0400004107b107a70000040f0000000100200190000007300000613d000000000001042d0000000001000019000007b300010430000000000001042f0000023201000041000000000010043f0000004101000039000000040010043f000001fd01000041000007b3000104300000001f0530018f000001ef06300198000000400200043d0000000004620019000007440000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000007400000c13d000000000005004b000007510000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000001ed0020009c000001ed020080410000004002200210000000000121019f000007b300010430000001ed033001970000001f0530018f000001ef06300198000000400200043d0000000004620019000007630000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b0000075f0000c13d000000000005004b000007700000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000001ed0020009c000001ed020080410000004002200210000000000112019f000007b300010430000000000100041a000001f0011001970000000002000411000000000012004b0000077c0000c13d000000000001042d000000400100043d000000440210003900000237030000410000000000320435000000240210003900000016030000390000000000320435000001f9020000410000000000210435000000040210003900000020030000390000000000320435000001ed0010009c000001ed010080410000004001100210000001fb011001c7000007b300010430000000000001042f00000000050100190000000000200443000000040100003900000005024002700000000002020031000000000121043a0000002004400039000000000031004b000007910000413d000001ed0030009c000001ed0300804100000060013002100000000002000414000001ed0020009c000001ed02008041000000c002200210000000000112019f00000245011001c7000000000205001907b107ac0000040f0000000100200190000007a60000613d000000000101043b000000000001042d000000000001042f000007aa002104210000000102000039000000000001042d0000000002000019000000000001042d000007af002104230000000102000039000000000001042d0000000002000019000000000001042d000007b100000432000007b20001042e000007b30001043000000000000000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000000000000000000000000001ffffffe000000000000000000000000000000000000000000000000000000000ffffffe0000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000ffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff095ea7b3000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000044000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffff00000000000000000000fe000000000000000000000000000000000000000000ffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff000000020000000000000000000000000000008000000100000000000000000008c379a00000000000000000000000000000000000000000000000000000000043616e6e6f7420736574206f776e657220746f207a65726f00000000000000000000000000000000000000000000000000000064000000000000000000000000d7f73334000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000000000000000000000000000000000008f491cb900000000000000000000000000000000000000000000000000000000b5a1101000000000000000000000000000000000000000000000000000000000e6c725f400000000000000000000000000000000000000000000000000000000e6c725f500000000000000000000000000000000000000000000000000000000ef686d8e00000000000000000000000000000000000000000000000000000000f2fde38b00000000000000000000000000000000000000000000000000000000b5a1101100000000000000000000000000000000000000000000000000000000bee518a400000000000000000000000000000000000000000000000000000000ca709a2500000000000000000000000000000000000000000000000000000000ae90de5400000000000000000000000000000000000000000000000000000000ae90de5500000000000000000000000000000000000000000000000000000000b0f479a100000000000000000000000000000000000000000000000000000000b187bd26000000000000000000000000000000000000000000000000000000008f491cba000000000000000000000000000000000000000000000000000000009d2aede5000000000000000000000000000000000000000000000000000000002b6e5d620000000000000000000000000000000000000000000000000000000079ba50960000000000000000000000000000000000000000000000000000000079ba50970000000000000000000000000000000000000000000000000000000085572ffb000000000000000000000000000000000000000000000000000000008da5cb5b000000000000000000000000000000000000000000000000000000002b6e5d6300000000000000000000000000000000000000000000000000000000665ed53700000000000000000000000000000000000000000000000000000000181f5a7600000000000000000000000000000000000000000000000000000000181f5a77000000000000000000000000000000000000000000000000000000001892b906000000000000000000000000000000000000000000000000000000002874d8bf0000000000000000000000000000000000000000000000000000000001ffc9a70000000000000000000000000000000000000000000000000000000016c38b3c0200000000000000000000000000000000000000000000000000000000000000ed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127843616e6e6f74207472616e7366657220746f2073656c66000000000000000000000000000000000000000000000000000000006400000080000000000000000000000000000000000000ff00000000000000000000000000000000000000000002000000000000000000000000000000000000200000008000000000000000004768dbf8645b24c54f2887651545d24f748c0d0d1d4c689eb810fb19f0befcf3000000000000000000000000000000000000002000000080000000000000000000000000ffffffffffffffff0000000000000000000000000000000000000000ffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff0000000000000000000000ff0000000000000000000000000000000000000000310ab089e4439a4c15d089f94afb7896ff553aecb10793d0ab882de59d99a32e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0200000200000000000000000000000000000044000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffff5f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0000000000000000000000000000000000000000000000000ffffffffffffffbf020000000000000000000000000000000000002000000000000000000000000048257dc961b6f792c2b78a080dacfed693b660960a702de21cee364e20270e2f58b69f57828e6962d216502094c54f6562f3bf082ba758966c3454f9e37b1525000000000000000000000000000000000000000000000000ffffffffffffff7f000000000000000000000000000000000000000000000000ffffffffffffffdf96f4e9f9000000000000000000000000000000000000000000000000000000004e487b71000000000000000000000000000000000000000000000000000000004d7573742062652070726f706f736564206f776e6572000000000000000000008be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0000000000000000000000001000000000000000000000000000000000000000005a3fef9935c9013a24c6193df2240d34fcf6b0ebf8786b85efe8401d696cdd94f6e6c792063616c6c61626c65206279206f776e6572000000000000000000000000000000000000000001000000000000000000000000000000000000000000a8d87a3b000000000000000000000000000000000000000000000000000000001806aa1896bbf26568e884a7374b41e002500962caba6a15023a8d90e8508b830200000200000000000000000000000000000024000000000000000000000000eff7cc48000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000302777af5d26fab9dd5120c5f1307c65193ebc51daf33244ada4365fab10602c53656c6646756e64656450696e67506f6e6720312e352e3000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff01ffc9a70000000000000000000000000000000000000000000000000000000085572ffb00000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe002000002000000000000000000000000000000000000000000000000000000009b33ae408f9f4b66895c74dd45801bbc70862d957b3a053c842a32637109b989")

func (_SelfFundedPingPong *SelfFundedPingPong) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _SelfFundedPingPong.abi.Events["CountIncrBeforeFundingSet"].ID:
		return _SelfFundedPingPong.ParseCountIncrBeforeFundingSet(log)
	case _SelfFundedPingPong.abi.Events["Funded"].ID:
		return _SelfFundedPingPong.ParseFunded(log)
	case _SelfFundedPingPong.abi.Events["OutOfOrderExecutionChange"].ID:
		return _SelfFundedPingPong.ParseOutOfOrderExecutionChange(log)
	case _SelfFundedPingPong.abi.Events["OwnershipTransferRequested"].ID:
		return _SelfFundedPingPong.ParseOwnershipTransferRequested(log)
	case _SelfFundedPingPong.abi.Events["OwnershipTransferred"].ID:
		return _SelfFundedPingPong.ParseOwnershipTransferred(log)
	case _SelfFundedPingPong.abi.Events["Ping"].ID:
		return _SelfFundedPingPong.ParsePing(log)
	case _SelfFundedPingPong.abi.Events["Pong"].ID:
		return _SelfFundedPingPong.ParsePong(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (SelfFundedPingPongCountIncrBeforeFundingSet) Topic() common.Hash {
	return common.HexToHash("0x4768dbf8645b24c54f2887651545d24f748c0d0d1d4c689eb810fb19f0befcf3")
}

func (SelfFundedPingPongFunded) Topic() common.Hash {
	return common.HexToHash("0x302777af5d26fab9dd5120c5f1307c65193ebc51daf33244ada4365fab10602c")
}

func (SelfFundedPingPongOutOfOrderExecutionChange) Topic() common.Hash {
	return common.HexToHash("0x05a3fef9935c9013a24c6193df2240d34fcf6b0ebf8786b85efe8401d696cdd9")
}

func (SelfFundedPingPongOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (SelfFundedPingPongOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (SelfFundedPingPongPing) Topic() common.Hash {
	return common.HexToHash("0x48257dc961b6f792c2b78a080dacfed693b660960a702de21cee364e20270e2f")
}

func (SelfFundedPingPongPong) Topic() common.Hash {
	return common.HexToHash("0x58b69f57828e6962d216502094c54f6562f3bf082ba758966c3454f9e37b1525")
}

func (_SelfFundedPingPong *SelfFundedPingPong) Address() common.Address {
	return _SelfFundedPingPong.address
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

func DeployZkSyncSelfFundedPingPong(auth *bind.TransactOpts, backend bind.ContractBackend, params ...interface{}) (common.Address, *CustomTransaction, *SelfFundedPingPong, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	
	zksyncClient := zkSyncClient.NewClient(client.Client())
	
	wallet := auth.Context.Value("wallet").(*zkSyncAccounts.Wallet)
	
	
	decodedBytes := common.FromHex(SelfFundedPingPongZkBin)
	
	SelfFundedPingPongAbi, err := SelfFundedPingPongMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	constructor, err := SelfFundedPingPongAbi.Pack("", params...)
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
	
	receipt, err := zksyncClient.WaitMined(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	tx, _, err := zksyncClient.TransactionByHash(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	

	ethTx := ConvertToTransaction(*tx)
	address := receipt.ContractAddress

	parsed, err := SelfFundedPingPongMetaData.GetAbi()
	contractBind := bind.NewBoundContract(address, *parsed, backend, backend, backend)

	contractReturn := &SelfFundedPingPong{address: address, abi: *parsed, SelfFundedPingPongCaller: SelfFundedPingPongCaller{contract: contractBind}, SelfFundedPingPongTransactor: SelfFundedPingPongTransactor{contract: contractBind}, SelfFundedPingPongFilterer: SelfFundedPingPongFilterer{contract: contractBind}}

	return address, ethTx, contractReturn, err
}

type SelfFundedPingPongInterface interface {
	GetCountIncrBeforeFunding(opts *bind.CallOpts) (uint8, error)

	GetCounterpartAddress(opts *bind.CallOpts) (common.Address, error)

	GetCounterpartChainSelector(opts *bind.CallOpts) (uint64, error)

	GetFeeToken(opts *bind.CallOpts) (common.Address, error)

	GetOutOfOrderExecution(opts *bind.CallOpts) (bool, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IsPaused(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	FundPingPong(opts *bind.TransactOpts, pingPongCount *big.Int) (*types.Transaction, error)

	SetCountIncrBeforeFunding(opts *bind.TransactOpts, countIncrBeforeFunding uint8) (*types.Transaction, error)

	SetCounterpart(opts *bind.TransactOpts, counterpartChainSelector uint64, counterpartAddress common.Address) (*types.Transaction, error)

	SetCounterpartAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error)

	SetCounterpartChainSelector(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error)

	SetOutOfOrderExecution(opts *bind.TransactOpts, outOfOrderExecution bool) (*types.Transaction, error)

	SetPaused(opts *bind.TransactOpts, pause bool) (*types.Transaction, error)

	StartPingPong(opts *bind.TransactOpts) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterCountIncrBeforeFundingSet(opts *bind.FilterOpts) (*SelfFundedPingPongCountIncrBeforeFundingSetIterator, error)

	WatchCountIncrBeforeFundingSet(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongCountIncrBeforeFundingSet) (event.Subscription, error)

	ParseCountIncrBeforeFundingSet(log types.Log) (*SelfFundedPingPongCountIncrBeforeFundingSet, error)

	FilterFunded(opts *bind.FilterOpts) (*SelfFundedPingPongFundedIterator, error)

	WatchFunded(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongFunded) (event.Subscription, error)

	ParseFunded(log types.Log) (*SelfFundedPingPongFunded, error)

	FilterOutOfOrderExecutionChange(opts *bind.FilterOpts) (*SelfFundedPingPongOutOfOrderExecutionChangeIterator, error)

	WatchOutOfOrderExecutionChange(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongOutOfOrderExecutionChange) (event.Subscription, error)

	ParseOutOfOrderExecutionChange(log types.Log) (*SelfFundedPingPongOutOfOrderExecutionChange, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SelfFundedPingPongOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*SelfFundedPingPongOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SelfFundedPingPongOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*SelfFundedPingPongOwnershipTransferred, error)

	FilterPing(opts *bind.FilterOpts) (*SelfFundedPingPongPingIterator, error)

	WatchPing(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongPing) (event.Subscription, error)

	ParsePing(log types.Log) (*SelfFundedPingPongPing, error)

	FilterPong(opts *bind.FilterOpts) (*SelfFundedPingPongPongIterator, error)

	WatchPong(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongPong) (event.Subscription, error)

	ParsePong(log types.Log) (*SelfFundedPingPongPong, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
