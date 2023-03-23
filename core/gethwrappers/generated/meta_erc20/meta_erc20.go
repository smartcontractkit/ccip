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

var MetaERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MustBeTrustedForwarder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTrustedForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"}],\"name\":\"metaTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"setFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIForwarder\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"setForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001c1f38038062001c1f8339810160408190526200003491620001c5565b33806000836001600160a01b03811662000069576040516335fdcccd60e21b8152600060048201526024015b60405180910390fd5b6001600160a01b039081166080528216620000c75760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f0000000000000000604482015260640162000060565b600180546001600160a01b0319166001600160a01b0384811691909117909155811615620000fa57620000fa8162000119565b5050506003829055503360009081526004602052604090205562000204565b336001600160a01b03821603620001735760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000060565b600280546001600160a01b0319166001600160a01b03838116918217909255600154604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b60008060408385031215620001d957600080fd5b825160208401519092506001600160a01b0381168114620001f957600080fd5b809150509250929050565b6080516119ea62000235600039600081816103660152818161097101528181610a0c0152610bbb01526119ea6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806382b02281116100d8578063b0f479a11161008c578063ce1b815f11610066578063ce1b815f146103b0578063dd62ed3e146103ce578063f2fde38b146103f957600080fd5b8063b0f479a114610364578063b9998a241461038a578063ca1d209d1461039d57600080fd5b80638da5cb5b116100bd5780638da5cb5b146102d657806395d89b4114610315578063a9059cbb1461035157600080fd5b806382b02281146102b057806385572ffb146102c357600080fd5b806323b872dd1161012f578063572b6c0511610114578063572b6c051461025957806370a082311461028857806379ba5097146102a857600080fd5b806323b872dd1461022c578063313ce5671461023f57600080fd5b8063095ea7b311610160578063095ea7b3146101ed57806315cce2241461020057806318160ddd1461021557600080fd5b806301ffc9a71461017c57806306fdde03146101a4575b600080fd5b61018f61018a3660046112eb565b61040c565b60405190151581526020015b60405180910390f35b6101e06040518060400160405280600981526020017f42616e6b546f6b656e000000000000000000000000000000000000000000000081525081565b60405161019b9190611398565b61018f6101fb3660046113cd565b6104a5565b61021361020e3660046113f9565b6104c2565b005b61021e60035481565b60405190815260200161019b565b61018f61023a366004611416565b610511565b610247601281565b60405160ff909116815260200161019b565b61018f6102673660046113f9565b60005473ffffffffffffffffffffffffffffffffffffffff91821691161490565b61021e6102963660046113f9565b60046020526000908152604090205481565b61021361066d565b6102136102be366004611474565b610773565b6102136102d13660046114c5565b6109f4565b60015473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101e06040518060400160405280600981526020017f42414e4b544f4b454e000000000000000000000000000000000000000000000081525081565b61018f61035f3660046113cd565b610a79565b7f00000000000000000000000000000000000000000000000000000000000000006102f0565b6102136103983660046113f9565b610a8d565b6102136103ab366004611500565b610ad8565b60005473ffffffffffffffffffffffffffffffffffffffff166102f0565b61021e6103dc366004611519565b600560209081526000928352604080842090915290825290205481565b6102136104073660046113f9565b610c74565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f85572ffb00000000000000000000000000000000000000000000000000000000148061049f57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b60006104b96104b2610c85565b8484610ce4565b50600192915050565b6104ca610d53565b600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b73ffffffffffffffffffffffffffffffffffffffff831660009081526005602052604081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9082610561610c85565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146106585773ffffffffffffffffffffffffffffffffffffffff84166000908152600560205260408120610602918491906105d5610c85565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040016000205490610dd6565b73ffffffffffffffffffffffffffffffffffffffff8516600090815260056020526040812090610630610c85565b73ffffffffffffffffffffffffffffffffffffffff1681526020810191909152604001600020555b610663848484610de9565b5060019392505050565b60025473ffffffffffffffffffffffffffffffffffffffff1633146106f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560028054909116905560405173ffffffffffffffffffffffffffffffffffffffff909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b60005473ffffffffffffffffffffffffffffffffffffffff1633146107c6576040517fa2f64cc50000000000000000000000000000000000000000000000000000000081523360048201526024016106ea565b6040805160a0810190915273ffffffffffffffffffffffffffffffffffffffff851660c08201526060906000908060e081016040516020818303038152906040528152602001868660405160200161084092919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152908252602080830186905260065473ffffffffffffffffffffffffffffffffffffffff1683830152815180830183526207a12080825260009183019182528351602481019190915290511515604480830191909152835180830390910181526064909101909252810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f97a657c900000000000000000000000000000000000000000000000000000000179052606090910152905061093461092e610c85565b85610eb6565b6040517f96f4e9f900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906396f4e9f9906109a89086908590600401611552565b6020604051808303816000875af11580156109c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109eb9190611664565b50505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610a65576040517fd7f733340000000000000000000000000000000000000000000000000000000081523360048201526024016106ea565b610a76610a7182611882565b61109b565b50565b60006104b9610a86610c85565b8484610de9565b610a95610d53565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff831617905550565b60065473ffffffffffffffffffffffffffffffffffffffff166323b872dd610afe610c85565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152306024820152604481018490526064016020604051808303816000875af1158015610b76573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9a919061192f565b5060065473ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000006040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602481018490526044016020604051808303816000875af1158015610c4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c70919061192f565b5050565b610c7c610d53565b610a76816110c9565b600060143610801590610caf575060005473ffffffffffffffffffffffffffffffffffffffff1633145b15610cdf57507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec36013560601c90565b503390565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526005602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610dd4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016106ea565b565b6000610de28284611980565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260046020526040902054610e199082610dd6565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600460205260408082209390935590841681522054610e5590826111bf565b73ffffffffffffffffffffffffffffffffffffffff80841660008181526004602052604090819020939093559151908516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90610d469085815260200190565b73ffffffffffffffffffffffffffffffffffffffff8216610f59576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f730000000000000000000000000000000000000000000000000000000000000060648201526084016106ea565b73ffffffffffffffffffffffffffffffffffffffff82166000908152600460205260409020548181101561100f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f636500000000000000000000000000000000000000000000000000000000000060648201526084016106ea565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260046020526040812083830390556003805484929061104b908490611980565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610d46565b60008082606001518060200190518101906110b69190611997565b915091506110c482826111cb565b505050565b3373ffffffffffffffffffffffffffffffffffffffff821603611148576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016106ea565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600154604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b6000610de282846119c5565b73ffffffffffffffffffffffffffffffffffffffff8216611248576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016106ea565b806003600082825461125a91906119c5565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260046020526040812080548392906112949084906119c5565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000602082840312156112fd57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610de257600080fd5b6000815180845260005b8181101561135357602081850181015186830182015201611337565b81811115611365576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610de2602083018461132d565b73ffffffffffffffffffffffffffffffffffffffff81168114610a7657600080fd5b600080604083850312156113e057600080fd5b82356113eb816113ab565b946020939093013593505050565b60006020828403121561140b57600080fd5b8135610de2816113ab565b60008060006060848603121561142b57600080fd5b8335611436816113ab565b92506020840135611446816113ab565b929592945050506040919091013590565b803567ffffffffffffffff8116811461146f57600080fd5b919050565b6000806000806080858703121561148a57600080fd5b8435611495816113ab565b935060208501356114a5816113ab565b9250604085013591506114ba60608601611457565b905092959194509250565b6000602082840312156114d757600080fd5b813567ffffffffffffffff8111156114ee57600080fd5b820160a08185031215610de257600080fd5b60006020828403121561151257600080fd5b5035919050565b6000806040838503121561152c57600080fd5b8235611537816113ab565b91506020830135611547816113ab565b809150509250929050565b6000604067ffffffffffffffff8516835260208181850152845160a08386015261157f60e086018261132d565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808784030160608801526115ba838361132d565b88860151888203830160808a01528051808352908601945060009350908501905b8084101561161a578451805173ffffffffffffffffffffffffffffffffffffffff168352860151868301529385019360019390930192908601906115db565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a01529550611656818761132d565b9a9950505050505050505050565b60006020828403121561167657600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156116cf576116cf61167d565b60405290565b60405160a0810167ffffffffffffffff811182821017156116cf576116cf61167d565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561173f5761173f61167d565b604052919050565b600082601f83011261175857600080fd5b813567ffffffffffffffff8111156117725761177261167d565b6117a360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016116f8565b8181528460208386010111156117b857600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126117e657600080fd5b8135602067ffffffffffffffff8211156118025761180261167d565b611810818360051b016116f8565b82815260069290921b8401810191818101908684111561182f57600080fd5b8286015b84811015611877576040818903121561184c5760008081fd5b6118546116ac565b813561185f816113ab565b81528185013585820152835291830191604001611833565b509695505050505050565b600060a0823603121561189457600080fd5b61189c6116d5565b823581526118ac60208401611457565b6020820152604083013567ffffffffffffffff808211156118cc57600080fd5b6118d836838701611747565b604084015260608501359150808211156118f157600080fd5b6118fd36838701611747565b6060840152608085013591508082111561191657600080fd5b50611923368286016117d5565b60808301525092915050565b60006020828403121561194157600080fd5b81518015158114610de257600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561199257611992611951565b500390565b600080604083850312156119aa57600080fd5b82516119b5816113ab565b6020939093015192949293505050565b600082198211156119d8576119d8611951565b50019056fea164736f6c634300080f000a",
}

var MetaERC20ABI = MetaERC20MetaData.ABI

var MetaERC20Bin = MetaERC20MetaData.Bin

func DeployMetaERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _totalSupply *big.Int, router common.Address) (common.Address, *types.Transaction, *MetaERC20, error) {
	parsed, err := MetaERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MetaERC20Bin), backend, _totalSupply, router)
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

func (_MetaERC20 *MetaERC20Caller) GetTrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "getTrustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) GetTrustedForwarder() (common.Address, error) {
	return _MetaERC20.Contract.GetTrustedForwarder(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) GetTrustedForwarder() (common.Address, error) {
	return _MetaERC20.Contract.GetTrustedForwarder(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20Caller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _MetaERC20.Contract.IsTrustedForwarder(&_MetaERC20.CallOpts, forwarder)
}

func (_MetaERC20 *MetaERC20CallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _MetaERC20.Contract.IsTrustedForwarder(&_MetaERC20.CallOpts, forwarder)
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

func (_MetaERC20 *MetaERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MetaERC20 *MetaERC20Session) Owner() (common.Address, error) {
	return _MetaERC20.Contract.Owner(&_MetaERC20.CallOpts)
}

func (_MetaERC20 *MetaERC20CallerSession) Owner() (common.Address, error) {
	return _MetaERC20.Contract.Owner(&_MetaERC20.CallOpts)
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

func (_MetaERC20 *MetaERC20Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "acceptOwnership")
}

func (_MetaERC20 *MetaERC20Session) AcceptOwnership() (*types.Transaction, error) {
	return _MetaERC20.Contract.AcceptOwnership(&_MetaERC20.TransactOpts)
}

func (_MetaERC20 *MetaERC20TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MetaERC20.Contract.AcceptOwnership(&_MetaERC20.TransactOpts)
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

func (_MetaERC20 *MetaERC20Transactor) MetaTransfer(opts *bind.TransactOpts, destinationTokenAddress common.Address, recipientAddress common.Address, amount *big.Int, destinationChainId uint64) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "metaTransfer", destinationTokenAddress, recipientAddress, amount, destinationChainId)
}

func (_MetaERC20 *MetaERC20Session) MetaTransfer(destinationTokenAddress common.Address, recipientAddress common.Address, amount *big.Int, destinationChainId uint64) (*types.Transaction, error) {
	return _MetaERC20.Contract.MetaTransfer(&_MetaERC20.TransactOpts, destinationTokenAddress, recipientAddress, amount, destinationChainId)
}

func (_MetaERC20 *MetaERC20TransactorSession) MetaTransfer(destinationTokenAddress common.Address, recipientAddress common.Address, amount *big.Int, destinationChainId uint64) (*types.Transaction, error) {
	return _MetaERC20.Contract.MetaTransfer(&_MetaERC20.TransactOpts, destinationTokenAddress, recipientAddress, amount, destinationChainId)
}

func (_MetaERC20 *MetaERC20Transactor) SetFeeToken(opts *bind.TransactOpts, feeToken common.Address) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "setFeeToken", feeToken)
}

func (_MetaERC20 *MetaERC20Session) SetFeeToken(feeToken common.Address) (*types.Transaction, error) {
	return _MetaERC20.Contract.SetFeeToken(&_MetaERC20.TransactOpts, feeToken)
}

func (_MetaERC20 *MetaERC20TransactorSession) SetFeeToken(feeToken common.Address) (*types.Transaction, error) {
	return _MetaERC20.Contract.SetFeeToken(&_MetaERC20.TransactOpts, feeToken)
}

func (_MetaERC20 *MetaERC20Transactor) SetForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "setForwarder", forwarder)
}

func (_MetaERC20 *MetaERC20Session) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _MetaERC20.Contract.SetForwarder(&_MetaERC20.TransactOpts, forwarder)
}

func (_MetaERC20 *MetaERC20TransactorSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _MetaERC20.Contract.SetForwarder(&_MetaERC20.TransactOpts, forwarder)
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

func (_MetaERC20 *MetaERC20Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MetaERC20.contract.Transact(opts, "transferOwnership", to)
}

func (_MetaERC20 *MetaERC20Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MetaERC20.Contract.TransferOwnership(&_MetaERC20.TransactOpts, to)
}

func (_MetaERC20 *MetaERC20TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MetaERC20.Contract.TransferOwnership(&_MetaERC20.TransactOpts, to)
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

type MetaERC20OwnershipTransferRequestedIterator struct {
	Event *MetaERC20OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MetaERC20OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaERC20OwnershipTransferRequested)
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
		it.Event = new(MetaERC20OwnershipTransferRequested)
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

func (it *MetaERC20OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *MetaERC20OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MetaERC20OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MetaERC20 *MetaERC20Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MetaERC20OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetaERC20.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MetaERC20OwnershipTransferRequestedIterator{contract: _MetaERC20.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_MetaERC20 *MetaERC20Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MetaERC20OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetaERC20.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MetaERC20OwnershipTransferRequested)
				if err := _MetaERC20.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_MetaERC20 *MetaERC20Filterer) ParseOwnershipTransferRequested(log types.Log) (*MetaERC20OwnershipTransferRequested, error) {
	event := new(MetaERC20OwnershipTransferRequested)
	if err := _MetaERC20.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MetaERC20OwnershipTransferredIterator struct {
	Event *MetaERC20OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MetaERC20OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaERC20OwnershipTransferred)
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
		it.Event = new(MetaERC20OwnershipTransferred)
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

func (it *MetaERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *MetaERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MetaERC20OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MetaERC20 *MetaERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MetaERC20OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetaERC20.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MetaERC20OwnershipTransferredIterator{contract: _MetaERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_MetaERC20 *MetaERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MetaERC20OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetaERC20.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MetaERC20OwnershipTransferred)
				if err := _MetaERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_MetaERC20 *MetaERC20Filterer) ParseOwnershipTransferred(log types.Log) (*MetaERC20OwnershipTransferred, error) {
	event := new(MetaERC20OwnershipTransferred)
	if err := _MetaERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
	case _MetaERC20.abi.Events["OwnershipTransferRequested"].ID:
		return _MetaERC20.ParseOwnershipTransferRequested(log)
	case _MetaERC20.abi.Events["OwnershipTransferred"].ID:
		return _MetaERC20.ParseOwnershipTransferred(log)
	case _MetaERC20.abi.Events["Transfer"].ID:
		return _MetaERC20.ParseTransfer(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MetaERC20Approval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (MetaERC20OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (MetaERC20OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (MetaERC20Transfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (_MetaERC20 *MetaERC20) Address() common.Address {
	return _MetaERC20.address
}

type MetaERC20Interface interface {
	Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetTrustedForwarder(opts *bind.CallOpts) (common.Address, error)

	IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error)

	Name(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	Fund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	MetaTransfer(opts *bind.TransactOpts, destinationTokenAddress common.Address, recipientAddress common.Address, amount *big.Int, destinationChainId uint64) (*types.Transaction, error)

	SetFeeToken(opts *bind.TransactOpts, feeToken common.Address) (*types.Transaction, error)

	SetForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MetaERC20ApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *MetaERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*MetaERC20Approval, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MetaERC20OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MetaERC20OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*MetaERC20OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MetaERC20OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MetaERC20OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*MetaERC20OwnershipTransferred, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MetaERC20TransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *MetaERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*MetaERC20Transfer, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
