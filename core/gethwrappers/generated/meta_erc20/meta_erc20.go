// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package meta_erc20

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
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated"
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

type ClientAny2EVMMessage struct {
	MessageId        [32]byte
	SourceChainId    uint64
	Sender           []byte
	Data             []byte
	DestTokenAmounts []ClientEVMTokenAmount
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type MetaERC20MetaTransferMessage struct {
	Owner    common.Address
	To       common.Address
	Amount   *big.Int
	Deadline *big.Int
	ChainId  uint64
	GasLimit uint64
}

var MetaERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"ccipFeeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"META_TRANSFER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structMetaERC20.MetaTransferMessage\",\"name\":\"m\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"metaTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200175d3803806200175d833981016040819052620000349162000194565b806001600160a01b03811662000064576040516335fdcccd60e21b81526000600482015260240160405180910390fd5b6001600160a01b039081166080908152600085815533815260016020818152604092839020889055600380546001600160a01b031916958816959095179094558151808301835260098152682130b735aa37b5b2b760b91b9085015281518083018352908152603160f81b9084015280517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f938101939093527f067a6a9de267623b53ea39bb8e464e087daccddeb12a976568b2c98febfe6d6b908301527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6606083015246908201523060a082015260c00160405160208183030381529060405280519060200120600481905550505050620001dc565b6001600160a01b03811681146200019157600080fd5b50565b600080600060608486031215620001aa57600080fd5b835192506020840151620001be816200017b565b6040850151909250620001d1816200017b565b809150509250925092565b6080516115506200020d600039600081816102e901528181610a1f01528181610aed0152610c2401526115506000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806376401ad3116100b2578063a9059cbb11610081578063b0f479a111610066578063b0f479a1146102cf578063ca1d209d14610313578063dd62ed3e1461032657600080fd5b8063a9059cbb14610295578063ac302bc9146102a857600080fd5b806376401ad3146102115780637ecebe001461022657806385572ffb1461024657806395d89b411461025957600080fd5b806323b872dd116100ee57806323b872dd146101bb578063313ce567146101ce5780633644e515146101e857806370a08231146101f157600080fd5b806301ffc9a71461012057806306fdde0314610148578063095ea7b31461019157806318160ddd146101a4575b600080fd5b61013361012e366004610e38565b610351565b60405190151581526020015b60405180910390f35b6101846040518060400160405280600981526020017f42616e6b546f6b656e000000000000000000000000000000000000000000000081525081565b60405161013f9190610ee5565b61013361019f366004610f21565b6103ea565b6101ad60005481565b60405190815260200161013f565b6101336101c9366004610f4b565b610400565b6101d6601281565b60405160ff909116815260200161013f565b6101ad60045481565b6101ad6101ff366004610f87565b60016020526000908152604090205481565b61022461021f366004610fa2565b6104d9565b005b6101ad610234366004610f87565b60056020526000908152604090205481565b610224610254366004610ff6565b610ad5565b6101846040518060400160405280600981526020017f42414e4b544f4b454e000000000000000000000000000000000000000000000081525081565b6101336102a3366004610f21565b610b55565b6101ad7ffa015aba6914852681b8c40c25589bc710e6970e3bc80c5140186542aec3299a81565b60405173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815260200161013f565b610224610321366004611031565b610b62565b6101ad61033436600461104a565b600260209081526000928352604080842090915290825290205481565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f85572ffb0000000000000000000000000000000000000000000000000000000014806103e457507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b60006103f7338484610cdd565b50600192915050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526002602090815260408083203384529091528120547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff146104c45773ffffffffffffffffffffffffffffffffffffffff841660009081526002602090815260408083203384529091529020546104929083610d4c565b73ffffffffffffffffffffffffffffffffffffffff851660009081526002602090815260408083203384529091529020555b6104cf848484610d5f565b5060019392505050565b428460600135101561054c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f455850495245440000000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b6004546000907ffa015aba6914852681b8c40c25589bc710e6970e3bc80c5140186542aec3299a6105806020880188610f87565b6105906040890160208a01610f87565b6040890135600560006105a660208d018d610f87565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160009081208054916105da836110ac565b909155506105ee60a08c0160808d016110fc565b6105fe60c08d0160a08e016110fc565b60408051602081019890985273ffffffffffffffffffffffffffffffffffffffff9687169088015294909316606086810191909152608086019290925260a085015267ffffffffffffffff91821660c0850152911660e083015287013561010082015261012001604051602081830303815290604052805190602001206040516020016106bd9291907f190100000000000000000000000000000000000000000000000000000000000081526002810192909252602282015260420190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600080855291840180845281905260ff88169284019290925260608301869052608083018590529092509060019060a0016020604051602081039080840390855afa158015610746573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116158015906107cd575061079e6020870187610f87565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610833576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f494e56414c49445f5349474e41545552450000000000000000000000000000006044820152606401610543565b4661084460a08801608089016110fc565b67ffffffffffffffff16036108825761087d6108636020880188610f87565b6108736040890160208a01610f87565b8860400135610d5f565b610acd565b60006040518060a001604052808860200160208101906108a29190610f87565b6040805173ffffffffffffffffffffffffffffffffffffffff9092166020830152016040516020818303038152906040528152602001604051806020016040528060008152508152602001886040013567ffffffffffffffff81111561090a5761090a611117565b60405190808252806020026020018201604052801561094f57816020015b60408051808201909152600080825260208201528152602001906001900390816109285790505b50815260035473ffffffffffffffffffffffffffffffffffffffff166020820152604080518082018252910190610a19908061099160c08d0160a08e016110fc565b67ffffffffffffffff168152600060209182015260408051835160248201529282015115156044808501919091528151808503909101815260649093019052810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f97a657c90000000000000000000000000000000000000000000000000000000017905290565b905290507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166396f4e9f9610a6960a08a0160808b016110fc565b836040518363ffffffff1660e01b8152600401610a87929190611146565b6020604051808303816000875af1158015610aa6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aca9190611258565b50505b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610b46576040517fd7f73334000000000000000000000000000000000000000000000000000000008152336004820152602401610543565b610b52610b5282611445565b50565b60006103f7338484610d5f565b6003546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810183905273ffffffffffffffffffffffffffffffffffffffff909116906323b872dd906064016020604051808303816000875af1158015610bdf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0391906114f2565b5060035473ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000006040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602481018490526044016020604051808303816000875af1158015610cb5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd991906114f2565b5050565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526002602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6000610d588284611514565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260016020526040902054610d8f9082610d4c565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600160205260408082209390935590841681522054610dcb9082610e2c565b73ffffffffffffffffffffffffffffffffffffffff80841660008181526001602052604090819020939093559151908516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90610d3f9085815260200190565b6000610d58828461152b565b600060208284031215610e4a57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610d5857600080fd5b6000815180845260005b81811015610ea057602081850181015186830182015201610e84565b81811115610eb2576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610d586020830184610e7a565b803573ffffffffffffffffffffffffffffffffffffffff81168114610f1c57600080fd5b919050565b60008060408385031215610f3457600080fd5b610f3d83610ef8565b946020939093013593505050565b600080600060608486031215610f6057600080fd5b610f6984610ef8565b9250610f7760208501610ef8565b9150604084013590509250925092565b600060208284031215610f9957600080fd5b610d5882610ef8565b600080600080848603610120811215610fba57600080fd5b60c0811215610fc857600080fd5b5084935060c085013560ff81168114610fe057600080fd5b939693955050505060e082013591610100013590565b60006020828403121561100857600080fd5b813567ffffffffffffffff81111561101f57600080fd5b820160a08185031215610d5857600080fd5b60006020828403121561104357600080fd5b5035919050565b6000806040838503121561105d57600080fd5b61106683610ef8565b915061107460208401610ef8565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110dd576110dd61107d565b5060010190565b803567ffffffffffffffff81168114610f1c57600080fd5b60006020828403121561110e57600080fd5b610d58826110e4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000604067ffffffffffffffff8516835260208181850152845160a08386015261117360e0860182610e7a565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808784030160608801526111ae8383610e7a565b88860151888203830160808a01528051808352908601945060009350908501905b8084101561120e578451805173ffffffffffffffffffffffffffffffffffffffff168352860151868301529385019360019390930192908601906111cf565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a0152955061124a8187610e7a565b9a9950505050505050505050565b60006020828403121561126a57600080fd5b5051919050565b6040805190810167ffffffffffffffff8111828210171561129457611294611117565b60405290565b60405160a0810167ffffffffffffffff8111828210171561129457611294611117565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561130457611304611117565b604052919050565b600082601f83011261131d57600080fd5b813567ffffffffffffffff81111561133757611337611117565b61136860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016112bd565b81815284602083860101111561137d57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126113ab57600080fd5b8135602067ffffffffffffffff8211156113c7576113c7611117565b6113d5818360051b016112bd565b82815260069290921b840181019181810190868411156113f457600080fd5b8286015b8481101561143a57604081890312156114115760008081fd5b611419611271565b61142282610ef8565b815281850135858201528352918301916040016113f8565b509695505050505050565b600060a0823603121561145757600080fd5b61145f61129a565b8235815261146f602084016110e4565b6020820152604083013567ffffffffffffffff8082111561148f57600080fd5b61149b3683870161130c565b604084015260608501359150808211156114b457600080fd5b6114c03683870161130c565b606084015260808501359150808211156114d957600080fd5b506114e63682860161139a565b60808301525092915050565b60006020828403121561150457600080fd5b81518015158114610d5857600080fd5b6000828210156115265761152661107d565b500390565b6000821982111561153e5761153e61107d565b50019056fea164736f6c634300080f000a",
}

var MetaERC20ABI = MetaERC20MetaData.ABI

var MetaERC20Bin = MetaERC20MetaData.Bin

func DeployMetaERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _totalSupply *big.Int, ccipFeeToken common.Address, router common.Address) (common.Address, *types.Transaction, *MetaERC20, error) {
	parsed, err := MetaERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MetaERC20Bin), backend, _totalSupply, ccipFeeToken, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MetaERC20{MetaERC20Caller: MetaERC20Caller{contract: contract}, MetaERC20Transactor: MetaERC20Transactor{contract: contract}, MetaERC20Filterer: MetaERC20Filterer{contract: contract}}, nil
}

type MetaERC20 struct {
	address common.Address
	abi     abi.ABI
	MetaERC20Caller
	MetaERC20Transactor
	MetaERC20Filterer
}

type MetaERC20Caller struct {
	contract *bind.BoundContract
}

type MetaERC20Transactor struct {
	contract *bind.BoundContract
}

type MetaERC20Filterer struct {
	contract *bind.BoundContract
}

type MetaERC20Session struct {
	Contract     *MetaERC20
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MetaERC20CallerSession struct {
	Contract *MetaERC20Caller
	CallOpts bind.CallOpts
}

type MetaERC20TransactorSession struct {
	Contract     *MetaERC20Transactor
	TransactOpts bind.TransactOpts
}

type MetaERC20Raw struct {
	Contract *MetaERC20
}

type MetaERC20CallerRaw struct {
	Contract *MetaERC20Caller
}

type MetaERC20TransactorRaw struct {
	Contract *MetaERC20Transactor
}

func NewMetaERC20(address common.Address, backend bind.ContractBackend) (*MetaERC20, error) {
	abi, err := abi.JSON(strings.NewReader(MetaERC20ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMetaERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaERC20{address: address, abi: abi, MetaERC20Caller: MetaERC20Caller{contract: contract}, MetaERC20Transactor: MetaERC20Transactor{contract: contract}, MetaERC20Filterer: MetaERC20Filterer{contract: contract}}, nil
}

func NewMetaERC20Caller(address common.Address, caller bind.ContractCaller) (*MetaERC20Caller, error) {
	contract, err := bindMetaERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaERC20Caller{contract: contract}, nil
}

func NewMetaERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*MetaERC20Transactor, error) {
	contract, err := bindMetaERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaERC20Transactor{contract: contract}, nil
}

func NewMetaERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*MetaERC20Filterer, error) {
	contract, err := bindMetaERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaERC20Filterer{contract: contract}, nil
}

func bindMetaERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_MetaERC20 *MetaERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaERC20.Contract.MetaERC20Caller.contract.Call(opts, result, method, params...)
}

func (_MetaERC20 *MetaERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaERC20.Contract.MetaERC20Transactor.contract.Transfer(opts)
}

func (_MetaERC20 *MetaERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaERC20.Contract.MetaERC20Transactor.contract.Transact(opts, method, params...)
}

func (_MetaERC20 *MetaERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaERC20.Contract.contract.Call(opts, result, method, params...)
}

func (_MetaERC20 *MetaERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaERC20.Contract.contract.Transfer(opts)
}

func (_MetaERC20 *MetaERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaERC20.Contract.contract.Transact(opts, method, params...)
}

func (_MetaERC20 *MetaERC20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _MetaERC20.Contract.DOMAINSEPARATOR(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MetaERC20.Contract.DOMAINSEPARATOR(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20Caller) METATRANSFERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "META_TRANSFER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) METATRANSFERTYPEHASH() ([32]byte, error) {
	return _MetaERC20.Contract.METATRANSFERTYPEHASH(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) METATRANSFERTYPEHASH() ([32]byte, error) {
	return _MetaERC20.Contract.METATRANSFERTYPEHASH(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MetaERC20.Contract.Allowance(&_MetaERC20.CallOpts, arg0, arg1)
}

func (_MetaERC20 *MetaERC20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MetaERC20.Contract.Allowance(&_MetaERC20.CallOpts, arg0, arg1)
}

func (_MetaERC20 *MetaERC20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MetaERC20.Contract.BalanceOf(&_MetaERC20.CallOpts, arg0)
}

func (_MetaERC20 *MetaERC20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MetaERC20.Contract.BalanceOf(&_MetaERC20.CallOpts, arg0)
}

func (_MetaERC20 *MetaERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) Decimals() (uint8, error) {
	return _MetaERC20.Contract.Decimals(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) Decimals() (uint8, error) {
	return _MetaERC20.Contract.Decimals(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20Caller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) GetRouter() (common.Address, error) {
	return _MetaERC20.Contract.GetRouter(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) GetRouter() (common.Address, error) {
	return _MetaERC20.Contract.GetRouter(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) Name() (string, error) {
	return _MetaERC20.Contract.Name(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) Name() (string, error) {
	return _MetaERC20.Contract.Name(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _MetaERC20.Contract.Nonces(&_MetaERC20.CallOpts, arg0)
}

func (_MetaERC20 *MetaERC20CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _MetaERC20.Contract.Nonces(&_MetaERC20.CallOpts, arg0)
}

func (_MetaERC20 *MetaERC20Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetaERC20.Contract.SupportsInterface(&_MetaERC20.CallOpts, interfaceId)
}

func (_MetaERC20 *MetaERC20CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetaERC20.Contract.SupportsInterface(&_MetaERC20.CallOpts, interfaceId)
}

func (_MetaERC20 *MetaERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) Symbol() (string, error) {
	return _MetaERC20.Contract.Symbol(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) Symbol() (string, error) {
	return _MetaERC20.Contract.Symbol(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) TotalSupply() (*big.Int, error) {
	return _MetaERC20.Contract.TotalSupply(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _MetaERC20.Contract.TotalSupply(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "approve", spender, amount)
}

func (_MetaERC20 *MetaERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Approve(&_MetaERC20.TransactOpts, spender, amount)
}

func (_MetaERC20 *MetaERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Approve(&_MetaERC20.TransactOpts, spender, amount)
}

func (_MetaERC20 *MetaERC20Transactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "ccipReceive", message)
}

func (_MetaERC20 *MetaERC20Session) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _MetaERC20.Contract.CcipReceive(&_MetaERC20.TransactOpts, message)
}

func (_MetaERC20 *MetaERC20TransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _MetaERC20.Contract.CcipReceive(&_MetaERC20.TransactOpts, message)
}

func (_MetaERC20 *MetaERC20Transactor) Fund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "fund", amount)
}

func (_MetaERC20 *MetaERC20Session) Fund(amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Fund(&_MetaERC20.TransactOpts, amount)
}

func (_MetaERC20 *MetaERC20TransactorSession) Fund(amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Fund(&_MetaERC20.TransactOpts, amount)
}

func (_MetaERC20 *MetaERC20Transactor) MetaTransfer(opts *bind.TransactOpts, m MetaERC20MetaTransferMessage, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "metaTransfer", m, v, r, s)
}

func (_MetaERC20 *MetaERC20Session) MetaTransfer(m MetaERC20MetaTransferMessage, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MetaERC20.Contract.MetaTransfer(&_MetaERC20.TransactOpts, m, v, r, s)
}

func (_MetaERC20 *MetaERC20TransactorSession) MetaTransfer(m MetaERC20MetaTransferMessage, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MetaERC20.Contract.MetaTransfer(&_MetaERC20.TransactOpts, m, v, r, s)
}

func (_MetaERC20 *MetaERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "transfer", to, amount)
}

func (_MetaERC20 *MetaERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Transfer(&_MetaERC20.TransactOpts, to, amount)
}

func (_MetaERC20 *MetaERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Transfer(&_MetaERC20.TransactOpts, to, amount)
}

func (_MetaERC20 *MetaERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

func (_MetaERC20 *MetaERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.TransferFrom(&_MetaERC20.TransactOpts, from, to, amount)
}

func (_MetaERC20 *MetaERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.TransferFrom(&_MetaERC20.TransactOpts, from, to, amount)
}

type MetaERC20ApprovalIterator struct {
	Event *MetaERC20Approval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MetaERC20ApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaERC20Approval)
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
		it.Event = new(MetaERC20Approval)
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

func (it *MetaERC20ApprovalIterator) Error() error {
	return it.fail
}

func (it *MetaERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MetaERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_MetaERC20 *MetaERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MetaERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MetaERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MetaERC20ApprovalIterator{contract: _MetaERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_MetaERC20 *MetaERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MetaERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MetaERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MetaERC20Approval)
				if err := _MetaERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_MetaERC20 *MetaERC20Filterer) ParseApproval(log types.Log) (*MetaERC20Approval, error) {
	event := new(MetaERC20Approval)
	if err := _MetaERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MetaERC20TransferIterator struct {
	Event *MetaERC20Transfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MetaERC20TransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaERC20Transfer)
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
		it.Event = new(MetaERC20Transfer)
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

func (it *MetaERC20TransferIterator) Error() error {
	return it.fail
}

func (it *MetaERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MetaERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_MetaERC20 *MetaERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MetaERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetaERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MetaERC20TransferIterator{contract: _MetaERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_MetaERC20 *MetaERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MetaERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetaERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MetaERC20Transfer)
				if err := _MetaERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_MetaERC20 *MetaERC20Filterer) ParseTransfer(log types.Log) (*MetaERC20Transfer, error) {
	event := new(MetaERC20Transfer)
	if err := _MetaERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_MetaERC20 *MetaERC20) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MetaERC20.abi.Events["Approval"].ID:
		return _MetaERC20.ParseApproval(log)
	case _MetaERC20.abi.Events["Transfer"].ID:
		return _MetaERC20.ParseTransfer(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MetaERC20Approval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (MetaERC20Transfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (_MetaERC20 *MetaERC20) Address() common.Address {
	return _MetaERC20.address
}

type MetaERC20Interface interface {
	DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error)

	METATRANSFERTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	Name(opts *bind.CallOpts) (string, error)

	Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	Fund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	MetaTransfer(opts *bind.TransactOpts, m MetaERC20MetaTransferMessage, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MetaERC20ApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *MetaERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*MetaERC20Approval, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MetaERC20TransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *MetaERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*MetaERC20Transfer, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
