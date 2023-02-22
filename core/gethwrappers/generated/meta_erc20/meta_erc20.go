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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"ccipFeeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReachHere\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"META_TRANSFER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structMetaERC20.MetaTransferMessage\",\"name\":\"m\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"metaTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001b4538038062001b45833981016040819052620000349162000194565b806001600160a01b03811662000064576040516335fdcccd60e21b81526000600482015260240160405180910390fd5b6001600160a01b039081166080908152600085815533815260016020818152604092839020889055600380546001600160a01b031916958816959095179094558151808301835260098152682130b735aa37b5b2b760b91b9085015281518083018352908152603160f81b9084015280517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f938101939093527f067a6a9de267623b53ea39bb8e464e087daccddeb12a976568b2c98febfe6d6b908301527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6606083015246908201523060a082015260c00160405160208183030381529060405280519060200120600481905550505050620001dc565b6001600160a01b03811681146200019157600080fd5b50565b600080600060608486031215620001aa57600080fd5b835192506020840151620001be816200017b565b6040850151909250620001d1816200017b565b809150509250925092565b608051611931620002146000396000818161037301528181610bc101528181610c8f01528181610eeb0152610fdf01526119316000f3fe608060405234801561001057600080fd5b506004361061016c5760003560e01c80637ecebe00116100cd578063a9059cbb11610081578063b0f479a111610066578063b0f479a114610359578063ca1d209d1461039d578063dd62ed3e146103b057600080fd5b8063a9059cbb1461031f578063ac302bc91461033257600080fd5b806395d89b41116100b257806395d89b41146102bd5780639dc29fac146102f9578063a35ff0a81461030c57600080fd5b80637ecebe001461028a57806385572ffb146102aa57600080fd5b8063313ce5671161012457806340c10f191161010957806340c10f191461024257806370a082311461025757806376401ad31461027757600080fd5b8063313ce5671461021f5780633644e5151461023957600080fd5b8063095ea7b311610155578063095ea7b3146101e257806318160ddd146101f557806323b872dd1461020c57600080fd5b806301ffc9a71461017157806306fdde0314610199575b600080fd5b61018461017f3660046111ea565b6103db565b60405190151581526020015b60405180910390f35b6101d56040518060400160405280600981526020017f42616e6b546f6b656e000000000000000000000000000000000000000000000081525081565b6040516101909190611297565b6101846101f03660046112d3565b610474565b6101fe60005481565b604051908152602001610190565b61018461021a3660046112fd565b61048a565b610227601281565b60405160ff9091168152602001610190565b6101fe60045481565b6102556102503660046112d3565b610563565b005b6101fe610265366004611339565b60016020526000908152604090205481565b610255610285366004611354565b610687565b6101fe610298366004611339565b60056020526000908152604090205481565b6102556102b83660046113a8565b610c77565b6101d56040518060400160405280600981526020017f42414e4b544f4b454e000000000000000000000000000000000000000000000081525081565b6102556103073660046112d3565b610cf7565b61025561031a3660046113e3565b610ee5565b61018461032d3660046112d3565b610f10565b6101fe7ffa015aba6914852681b8c40c25589bc710e6970e3bc80c5140186542aec3299a81565b60405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610190565b6102556103ab3660046113e3565b610f1d565b6101fe6103be3660046113fc565b600260209081526000928352604080842090915290825290205481565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f85572ffb00000000000000000000000000000000000000000000000000000000148061046e57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b6000610481338484611098565b50600192915050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526002602090815260408083203384529091528120547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1461054e5773ffffffffffffffffffffffffffffffffffffffff8416600090815260026020908152604080832033845290915290205461051c90836110fe565b73ffffffffffffffffffffffffffffffffffffffff851660009081526002602090815260408083203384529091529020555b610559848484611111565b5060019392505050565b73ffffffffffffffffffffffffffffffffffffffff82166105e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064015b60405180910390fd5b806000808282546105f6919061145e565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000908152600160205260408120805483929061063090849061145e565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b42846060013510156106f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f455850495245440000000000000000000000000000000000000000000000000060448201526064016105dc565b6004546000907ffa015aba6914852681b8c40c25589bc710e6970e3bc80c5140186542aec3299a6107296020880188611339565b6107396040890160208a01611339565b60408901356005600061074f60208d018d611339565b73ffffffffffffffffffffffffffffffffffffffff1681526020810191909152604001600090812080549161078383611476565b9091555061079760a08c0160808d016114c6565b6107a760c08d0160a08e016114c6565b60408051602081019890985273ffffffffffffffffffffffffffffffffffffffff9687169088015294909316606086810191909152608086019290925260a085015267ffffffffffffffff91821660c0850152911660e083015287013561010082015261012001604051602081830303815290604052805190602001206040516020016108669291907f190100000000000000000000000000000000000000000000000000000000000081526002810192909252602282015260420190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600080855291840180845281905260ff88169284019290925260608301869052608083018590529092509060019060a0016020604051602081039080840390855afa1580156108ef573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81161580159061097657506109476020870187611339565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b6109dc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f494e56414c49445f5349474e415455524500000000000000000000000000000060448201526064016105dc565b604080516001808252818301909252600091816020015b60408051808201909152600080825260208201528152602001906001900390816109f357905050905060405180604001604052803073ffffffffffffffffffffffffffffffffffffffff168152602001886040013581525081600081518110610a5e57610a5e611510565b602002602001018190525060006040518060a00160405280896020016020810190610a899190611339565b6040805173ffffffffffffffffffffffffffffffffffffffff909216602083015201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181529082528051602081810183526000825283015281810185905260035473ffffffffffffffffffffffffffffffffffffffff1660608301528051808201909152608090910190610bbb9080610b3360c08e0160a08f016114c6565b67ffffffffffffffff168152600060209182015260408051835160248201529282015115156044808501919091528151808503909101815260649093019052810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f97a657c90000000000000000000000000000000000000000000000000000000017905290565b905290507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166396f4e9f9610c0b60a08b0160808c016114c6565b836040518363ffffffff1660e01b8152600401610c2992919061153f565b6020604051808303816000875af1158015610c48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c6c9190611651565b505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610ce8576040517fd7f733340000000000000000000000000000000000000000000000000000000081523360048201526024016105dc565b610cf4610cf48261183e565b50565b73ffffffffffffffffffffffffffffffffffffffff8216610d9a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f730000000000000000000000000000000000000000000000000000000000000060648201526084016105dc565b73ffffffffffffffffffffffffffffffffffffffff821660009081526001602052604090205481811015610e50576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f636500000000000000000000000000000000000000000000000000000000000060648201526084016105dc565b73ffffffffffffffffffffffffffffffffffffffff83166000908152600160205260408120838303905580548391908190610e8c9084906118eb565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020015b60405180910390a3505050565b610cf4307f000000000000000000000000000000000000000000000000000000000000000083611098565b6000610481338484611111565b6003546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810183905273ffffffffffffffffffffffffffffffffffffffff909116906323b872dd906064016020604051808303816000875af1158015610f9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fbe9190611902565b5060035473ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000006040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602481018490526044016020604051808303816000875af1158015611070573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110949190611902565b5050565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526002602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259101610ed8565b600061110a82846118eb565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526001602052604090205461114190826110fe565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020526040808220939093559084168152205461117d90826111de565b73ffffffffffffffffffffffffffffffffffffffff80841660008181526001602052604090819020939093559151908516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90610ed89085815260200190565b600061110a828461145e565b6000602082840312156111fc57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461110a57600080fd5b6000815180845260005b8181101561125257602081850181015186830182015201611236565b81811115611264576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061110a602083018461122c565b803573ffffffffffffffffffffffffffffffffffffffff811681146112ce57600080fd5b919050565b600080604083850312156112e657600080fd5b6112ef836112aa565b946020939093013593505050565b60008060006060848603121561131257600080fd5b61131b846112aa565b9250611329602085016112aa565b9150604084013590509250925092565b60006020828403121561134b57600080fd5b61110a826112aa565b60008060008084860361012081121561136c57600080fd5b60c081121561137a57600080fd5b5084935060c085013560ff8116811461139257600080fd5b939693955050505060e082013591610100013590565b6000602082840312156113ba57600080fd5b813567ffffffffffffffff8111156113d157600080fd5b820160a0818503121561110a57600080fd5b6000602082840312156113f557600080fd5b5035919050565b6000806040838503121561140f57600080fd5b611418836112aa565b9150611426602084016112aa565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156114715761147161142f565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036114a7576114a761142f565b5060010190565b803567ffffffffffffffff811681146112ce57600080fd5b6000602082840312156114d857600080fd5b61110a826114ae565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000604067ffffffffffffffff8516835260208181850152845160a08386015261156c60e086018261122c565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808784030160608801526115a7838361122c565b88860151888203830160808a01528051808352908601945060009350908501905b80841015611607578451805173ffffffffffffffffffffffffffffffffffffffff168352860151868301529385019360019390930192908601906115c8565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a01529550611643818761122c565b9a9950505050505050505050565b60006020828403121561166357600080fd5b5051919050565b6040805190810167ffffffffffffffff8111828210171561168d5761168d6114e1565b60405290565b60405160a0810167ffffffffffffffff8111828210171561168d5761168d6114e1565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156116fd576116fd6114e1565b604052919050565b600082601f83011261171657600080fd5b813567ffffffffffffffff811115611730576117306114e1565b61176160207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016116b6565b81815284602083860101111561177657600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126117a457600080fd5b8135602067ffffffffffffffff8211156117c0576117c06114e1565b6117ce818360051b016116b6565b82815260069290921b840181019181810190868411156117ed57600080fd5b8286015b84811015611833576040818903121561180a5760008081fd5b61181261166a565b61181b826112aa565b815281850135858201528352918301916040016117f1565b509695505050505050565b600060a0823603121561185057600080fd5b611858611693565b82358152611868602084016114ae565b6020820152604083013567ffffffffffffffff8082111561188857600080fd5b61189436838701611705565b604084015260608501359150808211156118ad57600080fd5b6118b936838701611705565b606084015260808501359150808211156118d257600080fd5b506118df36828601611793565b60808301525092915050565b6000828210156118fd576118fd61142f565b500390565b60006020828403121561191457600080fd5b8151801515811461110a57600080fdfea164736f6c634300080f000a",
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

func (_MetaERC20 *MetaERC20Transactor) ApproveRouter(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "approveRouter", amount)
}

func (_MetaERC20 *MetaERC20Session) ApproveRouter(amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.ApproveRouter(&_MetaERC20.TransactOpts, amount)
}

func (_MetaERC20 *MetaERC20TransactorSession) ApproveRouter(amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.ApproveRouter(&_MetaERC20.TransactOpts, amount)
}

func (_MetaERC20 *MetaERC20Transactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "burn", account, amount)
}

func (_MetaERC20 *MetaERC20Session) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Burn(&_MetaERC20.TransactOpts, account, amount)
}

func (_MetaERC20 *MetaERC20TransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Burn(&_MetaERC20.TransactOpts, account, amount)
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

func (_MetaERC20 *MetaERC20Transactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "mint", account, amount)
}

func (_MetaERC20 *MetaERC20Session) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Mint(&_MetaERC20.TransactOpts, account, amount)
}

func (_MetaERC20 *MetaERC20TransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaERC20.Contract.Mint(&_MetaERC20.TransactOpts, account, amount)
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

	ApproveRouter(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	Fund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	MetaTransfer(opts *bind.TransactOpts, m MetaERC20MetaTransferMessage, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

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
