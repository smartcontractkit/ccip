// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bank_erc20

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

var BankERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ccipRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"chainlinkOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"test_only_force_cross_chain_transfer\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MustBeChainlinkOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MustBeTrustedForwarder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawFailure\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCCIPRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainlinkOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTrustedForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"}],\"name\":\"metaTransfer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001d6238038062001d628339810160408190526200003491620003e3565b8383838333806000816200008f5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c257620000c28162000164565b5050600280546001600160a01b039687166001600160a01b03199182161790915560038054958716959091169490941790935560048054911515600160a01b026001600160a81b031990921692909416919091171790915550600862000129888262000537565b50600962000138878262000537565b5062000157620001506000546001600160a01b031690565b866200020f565b505050505050506200062a565b336001600160a01b03821603620001be5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000086565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b038216620002675760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640162000086565b80600760008282546200027b919062000603565b90915550506001600160a01b03821660009081526005602052604081208054839290620002aa90849062000603565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b505050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200032157600080fd5b81516001600160401b03808211156200033e576200033e620002f9565b604051601f8301601f19908116603f01168101908282118183101715620003695762000369620002f9565b816040528381526020925086838588010111156200038657600080fd5b600091505b83821015620003aa57858201830151818301840152908201906200038b565b83821115620003bc5760008385830101525b9695505050505050565b80516001600160a01b0381168114620003de57600080fd5b919050565b600080600080600080600060e0888a031215620003ff57600080fd5b87516001600160401b03808211156200041757600080fd5b620004258b838c016200030f565b985060208a01519150808211156200043c57600080fd5b506200044b8a828b016200030f565b965050604088015194506200046360608901620003c6565b93506200047360808901620003c6565b92506200048360a08901620003c6565b915060c088015180151581146200049957600080fd5b8091505092959891949750929550565b600181811c90821680620004be57607f821691505b602082108103620004df57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620002f457600081815260208120601f850160051c810160208610156200050e5750805b601f850160051c820191505b818110156200052f578281556001016200051a565b505050505050565b81516001600160401b03811115620005535762000553620002f9565b6200056b81620005648454620004a9565b84620004e5565b602080601f831160018114620005a357600084156200058a5750858301515b600019600386901b1c1916600185901b1785556200052f565b600085815260208120601f198616915b82811015620005d457888601518255948401946001909101908401620005b3565b5085821015620005f35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600082198211156200062557634e487b7160e01b600052601160045260246000fd5b500190565b611728806200063a6000396000f3fe60806040526004361061016e5760003560e01c806370a08231116100cb578063a457c2d71161007f578063ce1b815f11610059578063ce1b815f146103ab578063dd62ed3e14610441578063f2fde38b1461049457600080fd5b8063a457c2d7146103d6578063a9059cbb146103f6578063b0cd7a6f1461041657600080fd5b80638da5cb5b116100b05780638da5cb5b1461036b57806395d89b4114610396578063a0042526146103ab57600080fd5b806370a082311461031357806379ba50971461035657600080fd5b8063313ce5671161012257806350431ce41161010757806350431ce414610274578063572b6c051461028b578063588cbd0e146102c757600080fd5b8063313ce56714610238578063395093511461025457600080fd5b8063178293441161015357806317829344146101d557806318160ddd1461020357806323b872dd1461021857600080fd5b806306fdde031461017a578063095ea7b3146101a557600080fd5b3661017557005b600080fd5b34801561018657600080fd5b5061018f6104b4565b60405161019c91906113f2565b60405180910390f35b3480156101b157600080fd5b506101c56101c036600461142e565b610546565b604051901515815260200161019c565b3480156101e157600080fd5b506101f56101f0366004611458565b610563565b60405190815260200161019c565b34801561020f57600080fd5b506007546101f5565b34801561022457600080fd5b506101c56102333660046114a5565b610848565b34801561024457600080fd5b506040516012815260200161019c565b34801561026057600080fd5b506101c561026f36600461142e565b61096f565b34801561028057600080fd5b506102896109d0565b005b34801561029757600080fd5b506101c56102a63660046114e1565b60025473ffffffffffffffffffffffffffffffffffffffff91821691161490565b3480156102d357600080fd5b5060035473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019c565b34801561031f57600080fd5b506101f561032e3660046114e1565b73ffffffffffffffffffffffffffffffffffffffff1660009081526005602052604090205490565b34801561036257600080fd5b50610289610ae0565b34801561037757600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff166102ee565b3480156103a257600080fd5b5061018f610bdd565b3480156103b757600080fd5b5060025473ffffffffffffffffffffffffffffffffffffffff166102ee565b3480156103e257600080fd5b506101c56103f136600461142e565b610bec565b34801561040257600080fd5b506101c561041136600461142e565b610ce2565b34801561042257600080fd5b5060045473ffffffffffffffffffffffffffffffffffffffff166102ee565b34801561044d57600080fd5b506101f561045c3660046114fc565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260066020908152604080832093909416825291909152205490565b3480156104a057600080fd5b506102896104af3660046114e1565b610cf6565b6060600880546104c39061152f565b80601f01602080910402602001604051908101604052809291908181526020018280546104ef9061152f565b801561053c5780601f106105115761010080835404028352916020019161053c565b820191906000526020600020905b81548152906001019060200180831161051f57829003601f168201915b5050505050905090565b600061055a610553610d0a565b8484610d69565b50600192915050565b60025460009073ffffffffffffffffffffffffffffffffffffffff1633146105be576040517fa2f64cc50000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6105c782610f1c565b6105e5576105dd6105d6610d0a565b8585610f5b565b506000610841565b604080516001808252818301909252600091816020015b60408051808201909152600080825260208201528152602001906001900390816105fc57905050905060405180604001604052803073ffffffffffffffffffffffffffffffffffffffff168152602001858152508160008151811061066357610663611582565b6020908102919091018101919091526040805160a08101825273ffffffffffffffffffffffffffffffffffffffff88811660c0808401919091528351808403909101815260e0830184528252825180850184526000808252838601919091528284018690526060830181905283519485018452808552608083019490945260035492517f20487ded00000000000000000000000000000000000000000000000000000000815291939216906320487ded9061072490889086906004016115b1565b602060405180830381865afa158015610741573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076591906116c3565b9050610779610772610d0a565b3088610f5b565b60035461079e90309073ffffffffffffffffffffffffffffffffffffffff1688610d69565b6003546040517f96f4e9f900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116906396f4e9f99083906107f890899087906004016115b1565b60206040518083038185885af1158015610816573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061083b91906116c3565b93505050505b9392505050565b6000610855848484610f5b565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260066020526040812081610883610d0a565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610950576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206160448201527f6c6c6f77616e636500000000000000000000000000000000000000000000000060648201526084016105b5565b6109648561095c610d0a565b858403610d69565b506001949350505050565b600061055a61097c610d0a565b84846006600061098a610d0a565b73ffffffffffffffffffffffffffffffffffffffff908116825260208083019390935260409182016000908120918b16815292529020546109cb91906116dc565b610d69565b60045473ffffffffffffffffffffffffffffffffffffffff163314610a23576040517fed88db320000000000000000000000000000000000000000000000000000000081523360048201526024016105b5565b476000610a4560005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114610a9c576040519150601f19603f3d011682016040523d82523d6000602084013e610aa1565b606091505b5050905080610adc576040517f1a0263ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610b61576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016105b5565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600980546104c39061152f565b60008060066000610bfb610d0a565b73ffffffffffffffffffffffffffffffffffffffff90811682526020808301939093526040918201600090812091881681529252902054905082811015610cc4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084016105b5565b610cd8610ccf610d0a565b85858403610d69565b5060019392505050565b600061055a610cef610d0a565b8484610f5b565b610cfe61120f565b610d0781611292565b50565b600060143610801590610d34575060025473ffffffffffffffffffffffffffffffffffffffff1633145b15610d6457507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec36013560601c90565b503390565b73ffffffffffffffffffffffffffffffffffffffff8316610e0b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016105b5565b73ffffffffffffffffffffffffffffffffffffffff8216610eae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016105b5565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526006602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b60045460009074010000000000000000000000000000000000000000900460ff1615610f4a57506001919050565b5067ffffffffffffffff1646141590565b73ffffffffffffffffffffffffffffffffffffffff8316610ffe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016105b5565b73ffffffffffffffffffffffffffffffffffffffff82166110a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016105b5565b73ffffffffffffffffffffffffffffffffffffffff831660009081526005602052604090205481811015611157576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016105b5565b73ffffffffffffffffffffffffffffffffffffffff80851660009081526005602052604080822085850390559185168152908120805484929061119b9084906116dc565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161120191815260200190565b60405180910390a350505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314611290576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016105b5565b565b3373ffffffffffffffffffffffffffffffffffffffff821603611311576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016105b5565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000815180845260005b818110156113ad57602081850181015186830182015201611391565b818111156113bf576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006108416020830184611387565b803573ffffffffffffffffffffffffffffffffffffffff8116811461142957600080fd5b919050565b6000806040838503121561144157600080fd5b61144a83611405565b946020939093013593505050565b60008060006060848603121561146d57600080fd5b61147684611405565b925060208401359150604084013567ffffffffffffffff8116811461149a57600080fd5b809150509250925092565b6000806000606084860312156114ba57600080fd5b6114c384611405565b92506114d160208501611405565b9150604084013590509250925092565b6000602082840312156114f357600080fd5b61084182611405565b6000806040838503121561150f57600080fd5b61151883611405565b915061152660208401611405565b90509250929050565b600181811c9082168061154357607f821691505b60208210810361157c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000604067ffffffffffffffff8516835260208181850152845160a0838601526115de60e0860182611387565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808784030160608801526116198383611387565b88860151888203830160808a01528051808352908601945060009350908501905b80841015611679578451805173ffffffffffffffffffffffffffffffffffffffff1683528601518683015293850193600193909301929086019061163a565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a015295506116b58187611387565b9a9950505050505050505050565b6000602082840312156116d557600080fd5b5051919050565b60008219821115611716577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50019056fea164736f6c634300080f000a",
}

var BankERC20ABI = BankERC20MetaData.ABI

var BankERC20Bin = BankERC20MetaData.Bin

func DeployBankERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, totalSupply *big.Int, forwarder common.Address, ccipRouter common.Address, chainlinkOwner common.Address, test_only_force_cross_chain_transfer bool) (common.Address, *types.Transaction, *BankERC20, error) {
	parsed, err := BankERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankERC20Bin), backend, name, symbol, totalSupply, forwarder, ccipRouter, chainlinkOwner, test_only_force_cross_chain_transfer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BankERC20{BankERC20Caller: BankERC20Caller{contract: contract}, BankERC20Transactor: BankERC20Transactor{contract: contract}, BankERC20Filterer: BankERC20Filterer{contract: contract}}, nil
}

type BankERC20 struct {
	address common.Address
	abi     abi.ABI
	BankERC20Caller
	BankERC20Transactor
	BankERC20Filterer
}

type BankERC20Caller struct {
	contract *bind.BoundContract
}

type BankERC20Transactor struct {
	contract *bind.BoundContract
}

type BankERC20Filterer struct {
	contract *bind.BoundContract
}

type BankERC20Session struct {
	Contract     *BankERC20
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BankERC20CallerSession struct {
	Contract *BankERC20Caller
	CallOpts bind.CallOpts
}

type BankERC20TransactorSession struct {
	Contract     *BankERC20Transactor
	TransactOpts bind.TransactOpts
}

type BankERC20Raw struct {
	Contract *BankERC20
}

type BankERC20CallerRaw struct {
	Contract *BankERC20Caller
}

type BankERC20TransactorRaw struct {
	Contract *BankERC20Transactor
}

func NewBankERC20(address common.Address, backend bind.ContractBackend) (*BankERC20, error) {
	abi, err := abi.JSON(strings.NewReader(BankERC20ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBankERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BankERC20{address: address, abi: abi, BankERC20Caller: BankERC20Caller{contract: contract}, BankERC20Transactor: BankERC20Transactor{contract: contract}, BankERC20Filterer: BankERC20Filterer{contract: contract}}, nil
}

func NewBankERC20Caller(address common.Address, caller bind.ContractCaller) (*BankERC20Caller, error) {
	contract, err := bindBankERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankERC20Caller{contract: contract}, nil
}

func NewBankERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*BankERC20Transactor, error) {
	contract, err := bindBankERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankERC20Transactor{contract: contract}, nil
}

func NewBankERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*BankERC20Filterer, error) {
	contract, err := bindBankERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankERC20Filterer{contract: contract}, nil
}

func bindBankERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BankERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BankERC20 *BankERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BankERC20.Contract.BankERC20Caller.contract.Call(opts, result, method, params...)
}

func (_BankERC20 *BankERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BankERC20.Contract.BankERC20Transactor.contract.Transfer(opts)
}

func (_BankERC20 *BankERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BankERC20.Contract.BankERC20Transactor.contract.Transact(opts, method, params...)
}

func (_BankERC20 *BankERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BankERC20.Contract.contract.Call(opts, result, method, params...)
}

func (_BankERC20 *BankERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BankERC20.Contract.contract.Transfer(opts)
}

func (_BankERC20 *BankERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BankERC20.Contract.contract.Transact(opts, method, params...)
}

func (_BankERC20 *BankERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BankERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BankERC20 *BankERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BankERC20.Contract.Allowance(&_BankERC20.CallOpts, owner, spender)
}

func (_BankERC20 *BankERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BankERC20.Contract.Allowance(&_BankERC20.CallOpts, owner, spender)
}

func (_BankERC20 *BankERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BankERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BankERC20 *BankERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _BankERC20.Contract.BalanceOf(&_BankERC20.CallOpts, account)
}

func (_BankERC20 *BankERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BankERC20.Contract.BalanceOf(&_BankERC20.CallOpts, account)
}

func (_BankERC20 *BankERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BankERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BankERC20 *BankERC20Session) Decimals() (uint8, error) {
	return _BankERC20.Contract.Decimals(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20CallerSession) Decimals() (uint8, error) {
	return _BankERC20.Contract.Decimals(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20Caller) GetCCIPRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BankERC20.contract.Call(opts, &out, "getCCIPRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BankERC20 *BankERC20Session) GetCCIPRouter() (common.Address, error) {
	return _BankERC20.Contract.GetCCIPRouter(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20CallerSession) GetCCIPRouter() (common.Address, error) {
	return _BankERC20.Contract.GetCCIPRouter(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20Caller) GetChainlinkOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BankERC20.contract.Call(opts, &out, "getChainlinkOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BankERC20 *BankERC20Session) GetChainlinkOwner() (common.Address, error) {
	return _BankERC20.Contract.GetChainlinkOwner(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20CallerSession) GetChainlinkOwner() (common.Address, error) {
	return _BankERC20.Contract.GetChainlinkOwner(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20Caller) GetForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BankERC20.contract.Call(opts, &out, "getForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BankERC20 *BankERC20Session) GetForwarder() (common.Address, error) {
	return _BankERC20.Contract.GetForwarder(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20CallerSession) GetForwarder() (common.Address, error) {
	return _BankERC20.Contract.GetForwarder(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20Caller) GetTrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BankERC20.contract.Call(opts, &out, "getTrustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BankERC20 *BankERC20Session) GetTrustedForwarder() (common.Address, error) {
	return _BankERC20.Contract.GetTrustedForwarder(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20CallerSession) GetTrustedForwarder() (common.Address, error) {
	return _BankERC20.Contract.GetTrustedForwarder(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20Caller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _BankERC20.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BankERC20 *BankERC20Session) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _BankERC20.Contract.IsTrustedForwarder(&_BankERC20.CallOpts, forwarder)
}

func (_BankERC20 *BankERC20CallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _BankERC20.Contract.IsTrustedForwarder(&_BankERC20.CallOpts, forwarder)
}

func (_BankERC20 *BankERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BankERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BankERC20 *BankERC20Session) Name() (string, error) {
	return _BankERC20.Contract.Name(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20CallerSession) Name() (string, error) {
	return _BankERC20.Contract.Name(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BankERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BankERC20 *BankERC20Session) Owner() (common.Address, error) {
	return _BankERC20.Contract.Owner(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20CallerSession) Owner() (common.Address, error) {
	return _BankERC20.Contract.Owner(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BankERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BankERC20 *BankERC20Session) Symbol() (string, error) {
	return _BankERC20.Contract.Symbol(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20CallerSession) Symbol() (string, error) {
	return _BankERC20.Contract.Symbol(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BankERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BankERC20 *BankERC20Session) TotalSupply() (*big.Int, error) {
	return _BankERC20.Contract.TotalSupply(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _BankERC20.Contract.TotalSupply(&_BankERC20.CallOpts)
}

func (_BankERC20 *BankERC20Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BankERC20.contract.Transact(opts, "acceptOwnership")
}

func (_BankERC20 *BankERC20Session) AcceptOwnership() (*types.Transaction, error) {
	return _BankERC20.Contract.AcceptOwnership(&_BankERC20.TransactOpts)
}

func (_BankERC20 *BankERC20TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BankERC20.Contract.AcceptOwnership(&_BankERC20.TransactOpts)
}

func (_BankERC20 *BankERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BankERC20.contract.Transact(opts, "approve", spender, amount)
}

func (_BankERC20 *BankERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BankERC20.Contract.Approve(&_BankERC20.TransactOpts, spender, amount)
}

func (_BankERC20 *BankERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BankERC20.Contract.Approve(&_BankERC20.TransactOpts, spender, amount)
}

func (_BankERC20 *BankERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BankERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

func (_BankERC20 *BankERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BankERC20.Contract.DecreaseAllowance(&_BankERC20.TransactOpts, spender, subtractedValue)
}

func (_BankERC20 *BankERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BankERC20.Contract.DecreaseAllowance(&_BankERC20.TransactOpts, spender, subtractedValue)
}

func (_BankERC20 *BankERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BankERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

func (_BankERC20 *BankERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BankERC20.Contract.IncreaseAllowance(&_BankERC20.TransactOpts, spender, addedValue)
}

func (_BankERC20 *BankERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BankERC20.Contract.IncreaseAllowance(&_BankERC20.TransactOpts, spender, addedValue)
}

func (_BankERC20 *BankERC20Transactor) MetaTransfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, destinationChainId uint64) (*types.Transaction, error) {
	return _BankERC20.contract.Transact(opts, "metaTransfer", receiver, amount, destinationChainId)
}

func (_BankERC20 *BankERC20Session) MetaTransfer(receiver common.Address, amount *big.Int, destinationChainId uint64) (*types.Transaction, error) {
	return _BankERC20.Contract.MetaTransfer(&_BankERC20.TransactOpts, receiver, amount, destinationChainId)
}

func (_BankERC20 *BankERC20TransactorSession) MetaTransfer(receiver common.Address, amount *big.Int, destinationChainId uint64) (*types.Transaction, error) {
	return _BankERC20.Contract.MetaTransfer(&_BankERC20.TransactOpts, receiver, amount, destinationChainId)
}

func (_BankERC20 *BankERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BankERC20.contract.Transact(opts, "transfer", recipient, amount)
}

func (_BankERC20 *BankERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BankERC20.Contract.Transfer(&_BankERC20.TransactOpts, recipient, amount)
}

func (_BankERC20 *BankERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BankERC20.Contract.Transfer(&_BankERC20.TransactOpts, recipient, amount)
}

func (_BankERC20 *BankERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BankERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

func (_BankERC20 *BankERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BankERC20.Contract.TransferFrom(&_BankERC20.TransactOpts, sender, recipient, amount)
}

func (_BankERC20 *BankERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BankERC20.Contract.TransferFrom(&_BankERC20.TransactOpts, sender, recipient, amount)
}

func (_BankERC20 *BankERC20Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BankERC20.contract.Transact(opts, "transferOwnership", to)
}

func (_BankERC20 *BankERC20Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BankERC20.Contract.TransferOwnership(&_BankERC20.TransactOpts, to)
}

func (_BankERC20 *BankERC20TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BankERC20.Contract.TransferOwnership(&_BankERC20.TransactOpts, to)
}

func (_BankERC20 *BankERC20Transactor) WithdrawNative(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BankERC20.contract.Transact(opts, "withdrawNative")
}

func (_BankERC20 *BankERC20Session) WithdrawNative() (*types.Transaction, error) {
	return _BankERC20.Contract.WithdrawNative(&_BankERC20.TransactOpts)
}

func (_BankERC20 *BankERC20TransactorSession) WithdrawNative() (*types.Transaction, error) {
	return _BankERC20.Contract.WithdrawNative(&_BankERC20.TransactOpts)
}

func (_BankERC20 *BankERC20Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BankERC20.contract.RawTransact(opts, nil)
}

func (_BankERC20 *BankERC20Session) Receive() (*types.Transaction, error) {
	return _BankERC20.Contract.Receive(&_BankERC20.TransactOpts)
}

func (_BankERC20 *BankERC20TransactorSession) Receive() (*types.Transaction, error) {
	return _BankERC20.Contract.Receive(&_BankERC20.TransactOpts)
}

type BankERC20ApprovalIterator struct {
	Event *BankERC20Approval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BankERC20ApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankERC20Approval)
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
		it.Event = new(BankERC20Approval)
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

func (it *BankERC20ApprovalIterator) Error() error {
	return it.fail
}

func (it *BankERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BankERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_BankERC20 *BankERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BankERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BankERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BankERC20ApprovalIterator{contract: _BankERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_BankERC20 *BankERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BankERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BankERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BankERC20Approval)
				if err := _BankERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_BankERC20 *BankERC20Filterer) ParseApproval(log types.Log) (*BankERC20Approval, error) {
	event := new(BankERC20Approval)
	if err := _BankERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BankERC20OwnershipTransferRequestedIterator struct {
	Event *BankERC20OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BankERC20OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankERC20OwnershipTransferRequested)
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
		it.Event = new(BankERC20OwnershipTransferRequested)
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

func (it *BankERC20OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BankERC20OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BankERC20OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BankERC20 *BankERC20Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BankERC20OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BankERC20.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BankERC20OwnershipTransferRequestedIterator{contract: _BankERC20.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BankERC20 *BankERC20Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BankERC20OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BankERC20.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BankERC20OwnershipTransferRequested)
				if err := _BankERC20.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_BankERC20 *BankERC20Filterer) ParseOwnershipTransferRequested(log types.Log) (*BankERC20OwnershipTransferRequested, error) {
	event := new(BankERC20OwnershipTransferRequested)
	if err := _BankERC20.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BankERC20OwnershipTransferredIterator struct {
	Event *BankERC20OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BankERC20OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankERC20OwnershipTransferred)
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
		it.Event = new(BankERC20OwnershipTransferred)
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

func (it *BankERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BankERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BankERC20OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BankERC20 *BankERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BankERC20OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BankERC20.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BankERC20OwnershipTransferredIterator{contract: _BankERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BankERC20 *BankERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BankERC20OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BankERC20.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BankERC20OwnershipTransferred)
				if err := _BankERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_BankERC20 *BankERC20Filterer) ParseOwnershipTransferred(log types.Log) (*BankERC20OwnershipTransferred, error) {
	event := new(BankERC20OwnershipTransferred)
	if err := _BankERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BankERC20TransferIterator struct {
	Event *BankERC20Transfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BankERC20TransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankERC20Transfer)
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
		it.Event = new(BankERC20Transfer)
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

func (it *BankERC20TransferIterator) Error() error {
	return it.fail
}

func (it *BankERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BankERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_BankERC20 *BankERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BankERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BankERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BankERC20TransferIterator{contract: _BankERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_BankERC20 *BankERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BankERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BankERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BankERC20Transfer)
				if err := _BankERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_BankERC20 *BankERC20Filterer) ParseTransfer(log types.Log) (*BankERC20Transfer, error) {
	event := new(BankERC20Transfer)
	if err := _BankERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_BankERC20 *BankERC20) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BankERC20.abi.Events["Approval"].ID:
		return _BankERC20.ParseApproval(log)
	case _BankERC20.abi.Events["OwnershipTransferRequested"].ID:
		return _BankERC20.ParseOwnershipTransferRequested(log)
	case _BankERC20.abi.Events["OwnershipTransferred"].ID:
		return _BankERC20.ParseOwnershipTransferred(log)
	case _BankERC20.abi.Events["Transfer"].ID:
		return _BankERC20.ParseTransfer(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (BankERC20Approval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (BankERC20OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BankERC20OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BankERC20Transfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (_BankERC20 *BankERC20) Address() common.Address {
	return _BankERC20.address
}

type BankERC20Interface interface {
	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	GetCCIPRouter(opts *bind.CallOpts) (common.Address, error)

	GetChainlinkOwner(opts *bind.CallOpts) (common.Address, error)

	GetForwarder(opts *bind.CallOpts) (common.Address, error)

	GetTrustedForwarder(opts *bind.CallOpts) (common.Address, error)

	IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error)

	Name(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)

	DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error)

	IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error)

	MetaTransfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, destinationChainId uint64) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawNative(opts *bind.TransactOpts) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BankERC20ApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *BankERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*BankERC20Approval, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BankERC20OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BankERC20OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BankERC20OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BankERC20OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BankERC20OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BankERC20OwnershipTransferred, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BankERC20TransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *BankERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*BankERC20Transfer, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
