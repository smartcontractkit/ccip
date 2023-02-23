// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package price_registry

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

type IPriceRegistryTimestampedUint128Value struct {
	Value     *big.Int
	Timestamp *big.Int
}

type InternalFeeTokenPriceUpdate struct {
	SourceFeeToken common.Address
	UsdPerFeeToken *big.Int
}

type InternalPriceUpdates struct {
	FeeTokenPriceUpdates []InternalFeeTokenPriceUpdate
	DestChainId          uint64
	UsdPerUnitGas        *big.Int
}

var PriceRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceFeeToken\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"usdPerFeeToken\",\"type\":\"uint128\"}],\"internalType\":\"structInternal.FeeTokenPriceUpdate[]\",\"name\":\"feeTokenPriceUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"usdPerUnitGas\",\"type\":\"uint128\"}],\"internalType\":\"structInternal.PriceUpdates\",\"name\":\"priceUpdates\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"priceUpdaters\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"stalenessThreshold\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chain\",\"type\":\"uint64\"}],\"name\":\"ChainNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStalenessThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByUpdaterOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePassed\",\"type\":\"uint256\"}],\"name\":\"StaleGasPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePassed\",\"type\":\"uint256\"}],\"name\":\"StaleTokenPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceUpdater\",\"type\":\"address\"}],\"name\":\"PriceUpdaterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceUpdater\",\"type\":\"address\"}],\"name\":\"PriceUpdaterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UsdPerFeeTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChain\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UsdPerUnitGasUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"priceUpdaters\",\"type\":\"address[]\"}],\"name\":\"addPriceUpdaters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"convertFeeTokenAmountToLinkAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"linkTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"name\":\"getDestinationChainGasPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"internalType\":\"structIPriceRegistry.TimestampedUint128Value\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"name\":\"getFeeTokenBaseUnitsPerUnitGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenBaseUnitsPerUnitGas\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeeTokenPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"internalType\":\"structIPriceRegistry.TimestampedUint128Value\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceUpdaters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"priceUpdaters\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStalenessThreshold\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"priceUpdaters\",\"type\":\"address[]\"}],\"name\":\"removePriceUpdaters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceFeeToken\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"usdPerFeeToken\",\"type\":\"uint128\"}],\"internalType\":\"structInternal.FeeTokenPriceUpdate[]\",\"name\":\"feeTokenPriceUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"usdPerUnitGas\",\"type\":\"uint128\"}],\"internalType\":\"structInternal.PriceUpdates\",\"name\":\"priceUpdates\",\"type\":\"tuple\"}],\"name\":\"updatePrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001f9038038062001f9083398101604081905262000034916200064d565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be8162000117565b505050620000d283620001c260201b60201c565b620000dd8262000246565b8063ffffffff166000036200010557604051631151410960e11b815260040160405180910390fd5b63ffffffff1660805250620007f19050565b336001600160a01b03821603620001715760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60005b81515181101562000214576200020182600001518281518110620001ed57620001ed620007b3565b6020026020010151620002f460201b60201c565b6200020c81620007c9565b9050620001c5565b5060208101516001600160401b031615620002435762000243816020015182604001516200039f60201b60201c565b50565b60005b8151811015620002f0576200028a8282815181106200026c576200026c620007b3565b602002602001015160046200043e60201b62000bb01790919060201c565b50818181518110620002a057620002a0620007b3565b60200260200101516001600160a01b03167f34a02290b7920078c19f58e94b78c77eb9cc10195b20676e19bd3b82085893b860405160405180910390a2620002e881620007c9565b905062000249565b5050565b604080518082018252602080840180516001600160801b0390811684524280821684860190815287516001600160a01b039081166000908152600390965294879020955190518316600160801b029216919091179093558451905193519116927f6ee3f6f5d4054ba011bcd5c323ed7b7fa0126b839aecb3f4e470fc81d5d4c4de9262000394926001600160801b03929092168252602082015260400190565b60405180910390a250565b6040805180820182526001600160801b0380841682524280821660208085019182526001600160401b0388166000818152600290925290869020945191518416600160801b02919093161790925591517fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e9162000432918591906001600160801b03929092168252602082015260400190565b60405180910390a25050565b600062000455836001600160a01b0384166200045e565b90505b92915050565b6000818152600183016020526040812054620004a75750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000458565b50600062000458565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715620004eb57620004eb620004b0565b60405290565b604080519081016001600160401b0381118282101715620004eb57620004eb620004b0565b604051601f8201601f191681016001600160401b0381118282101715620005415762000541620004b0565b604052919050565b60006001600160401b03821115620005655762000565620004b0565b5060051b60200190565b80516001600160a01b03811681146200058757600080fd5b919050565b80516001600160801b03811681146200058757600080fd5b80516001600160401b03811681146200058757600080fd5b600082601f830112620005ce57600080fd5b81516020620005e7620005e18362000549565b62000516565b82815260059290921b840181019181810190868411156200060757600080fd5b8286015b848110156200062d576200061f816200056f565b83529183019183016200060b565b509695505050505050565b805163ffffffff811681146200058757600080fd5b6000806000606084860312156200066357600080fd5b83516001600160401b03808211156200067b57600080fd5b90850190606082880312156200069057600080fd5b6200069a620004c6565b825182811115620006aa57600080fd5b8301601f81018913620006bc57600080fd5b80516020620006cf620005e18362000549565b82815260069290921b8301810191818101908c841115620006ef57600080fd5b938201935b838510156200074a576040858e0312156200070f5760008081fd5b62000719620004f1565b62000724866200056f565b8152620007338487016200058c565b8185015282526040949094019390820190620006f4565b8552506200075a868201620005a4565b818501526200076c604087016200058c565b60408501528901519297509193505050808211156200078a57600080fd5b506200079986828701620005bc565b925050620007aa6040850162000638565b90509250925092565b634e487b7160e01b600052603260045260246000fd5b600060018201620007ea57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805161174b6200084560003960008181610212015281816103d60152818161043f0152818161054e015281816105b601528181610864015281816108cc015281816109de0152610a46015261174b6000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c8063a5817a9511610081578063c39e61a11161005b578063c39e61a114610251578063d971819214610264578063f2fde38b146102de57600080fd5b8063a5817a95146101f5578063a6c94a7314610208578063bfcd45661461023c57600080fd5b806355346ac5116100b257806355346ac5146101bf57806379ba5097146101d25780638da5cb5b146101da57600080fd5b8063268e5d48146100d95780633c03ec3c146100ff578063514e8cff14610114575b600080fd5b6100ec6100e7366004611236565b6102f1565b6040519081526020015b60405180910390f35b61011261010d366004611377565b61062c565b005b61018f610122366004611493565b60408051808201909152600080825260208201525067ffffffffffffffff166000908152600260209081526040918290208251808401909352546fffffffffffffffffffffffffffffffff8082168452700100000000000000000000000000000000909104169082015290565b6040805182516fffffffffffffffffffffffffffffffff90811682526020938401511692810192909252016100f6565b6101126101cd3660046114ae565b610692565b6101126106a3565b6000546040516001600160a01b0390911681526020016100f6565b6100ec610203366004611546565b610786565b60405163ffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526020016100f6565b610244610adf565b6040516100f69190611582565b61011261025f3660046114ae565b610b8e565b61018f6102723660046115cf565b6040805180820190915260008082526020820152506001600160a01b03166000908152600360209081526040918290208251808401909352546fffffffffffffffffffffffffffffffff8082168452700100000000000000000000000000000000909104169082015290565b6101126102ec3660046115cf565b610b9f565b67ffffffffffffffff811660009081526002602090815260408083208151808301909252546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000090910416918101829052901580610366575080516fffffffffffffffffffffffffffffffff16155b156103ae576040517f2e59db3a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024015b60405180910390fd5b600081602001516fffffffffffffffffffffffffffffffff16426103d29190611619565b90507f000000000000000000000000000000000000000000000000000000000000000063ffffffff16811115610473576040517ff08bcb3e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8516600482015263ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166024820152604481018290526064016103a5565b6001600160a01b0385166000908152600360209081526040918290208251808401909352546fffffffffffffffffffffffffffffffff80821684527001000000000000000000000000000000009091041690820181905215806104e7575080516fffffffffffffffffffffffffffffffff16155b15610529576040517f06439c6b0000000000000000000000000000000000000000000000000000000081526001600160a01b03871660048201526024016103a5565b602081015161054a906fffffffffffffffffffffffffffffffff1642611619565b91507f000000000000000000000000000000000000000000000000000000000000000063ffffffff168211156105ea576040517fc65fdfca0000000000000000000000000000000000000000000000000000000081526001600160a01b038716600482015263ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166024820152604481018390526064016103a5565b805183516fffffffffffffffffffffffffffffffff918216916106169116670de0b6b3a7640000611630565b610620919061166d565b93505050505b92915050565b6000546001600160a01b0316331480159061064f575061064d600433610bcc565b155b15610686576040517f46f0815400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61068f81610bee565b50565b61069a610c58565b61068f81610cce565b6001546001600160a01b03163314610717576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016103a5565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6001600160a01b03821660009081526003602090815260408083208151808301909252546fffffffffffffffffffffffffffffffff8082168352700100000000000000000000000000000000909104169181018290529015806107fa575080516fffffffffffffffffffffffffffffffff16155b1561083c576040517f06439c6b0000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016103a5565b600081602001516fffffffffffffffffffffffffffffffff16426108609190611619565b90507f000000000000000000000000000000000000000000000000000000000000000063ffffffff16811115610900576040517fc65fdfca0000000000000000000000000000000000000000000000000000000081526001600160a01b038616600482015263ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166024820152604481018290526064016103a5565b6001600160a01b0386166000908152600360209081526040918290208251808401909352546fffffffffffffffffffffffffffffffff8082168452700100000000000000000000000000000000909104169082018190521580610974575080516fffffffffffffffffffffffffffffffff16155b156109b6576040517f06439c6b0000000000000000000000000000000000000000000000000000000081526001600160a01b03881660048201526024016103a5565b600081602001516fffffffffffffffffffffffffffffffff16426109da9190611619565b90507f000000000000000000000000000000000000000000000000000000000000000063ffffffff16811115610a7a576040517fc65fdfca0000000000000000000000000000000000000000000000000000000081526001600160a01b038916600482015263ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166024820152604481018290526064016103a5565b815184516000916fffffffffffffffffffffffffffffffff90811691610aa99116670de0b6b3a7640000611630565b610ab3919061166d565b9050670de0b6b3a7640000610ac88289611630565b610ad2919061166d565b9998505050505050505050565b6060610aeb6004610d70565b67ffffffffffffffff811115610b0357610b03611269565b604051908082528060200260200182016040528015610b2c578160200160208202803683370190505b50905060005b610b3c6004610d70565b811015610b8a57610b4e600482610d7a565b828281518110610b6057610b606116a8565b6001600160a01b039092166020928302919091019091015280610b82816116d7565b915050610b32565b5090565b610b96610c58565b61068f81610d86565b610ba7610c58565b61068f81610e1f565b6000610bc5836001600160a01b038416610efa565b9392505050565b6001600160a01b03811660009081526001830160205260408120541515610bc5565b60005b815151811015610c3157610c2182600001518281518110610c1457610c146116a8565b6020026020010151610f49565b610c2a816116d7565b9050610bf1565b50602081015167ffffffffffffffff161561068f5761068f81602001518260400151611012565b6000546001600160a01b03163314610ccc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016103a5565b565b60005b8151811015610d6c57610d07828281518110610cef57610cef6116a8565b602002602001015160046110d090919063ffffffff16565b15610d5c57818181518110610d1e57610d1e6116a8565b60200260200101516001600160a01b03167fff7dbb85c77ca68ca1f894d6498570e3d5095cd19466f07ee8d222b337e4068c60405160405180910390a25b610d65816116d7565b9050610cd1565b5050565b6000610626825490565b6000610bc583836110e5565b60005b8151811015610d6c57610dbf828281518110610da757610da76116a8565b60200260200101516004610bb090919063ffffffff16565b50818181518110610dd257610dd26116a8565b60200260200101516001600160a01b03167f34a02290b7920078c19f58e94b78c77eb9cc10195b20676e19bd3b82085893b860405160405180910390a2610e18816116d7565b9050610d89565b336001600160a01b03821603610e91576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016103a5565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000818152600183016020526040812054610f4157508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610626565b506000610626565b604080518082018252602080840180516fffffffffffffffffffffffffffffffff90811684524280821684860190815287516001600160a01b039081166000908152600390965294879020955190518316700100000000000000000000000000000000029216919091179093558451905193519116927f6ee3f6f5d4054ba011bcd5c323ed7b7fa0126b839aecb3f4e470fc81d5d4c4de92611007926fffffffffffffffffffffffffffffffff929092168252602082015260400190565b60405180910390a250565b6040805180820182526fffffffffffffffffffffffffffffffff808416825242808216602080850191825267ffffffffffffffff8816600081815260029092529086902094519151841670010000000000000000000000000000000002919093161790925591517fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e916110c4918591906fffffffffffffffffffffffffffffffff929092168252602082015260400190565b60405180910390a25050565b6000610bc5836001600160a01b03841661110f565b60008260000182815481106110fc576110fc6116a8565b9060005260206000200154905092915050565b600081815260018301602052604081205480156111f8576000611133600183611619565b855490915060009061114790600190611619565b90508181146111ac576000866000018281548110611167576111676116a8565b906000526020600020015490508087600001848154811061118a5761118a6116a8565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806111bd576111bd61170f565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610626565b6000915050610626565b80356001600160a01b038116811461121957600080fd5b919050565b803567ffffffffffffffff8116811461121957600080fd5b6000806040838503121561124957600080fd5b61125283611202565b91506112606020840161121e565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156112bb576112bb611269565b60405290565b6040805190810167ffffffffffffffff811182821017156112bb576112bb611269565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561132b5761132b611269565b604052919050565b600067ffffffffffffffff82111561134d5761134d611269565b5060051b60200190565b80356fffffffffffffffffffffffffffffffff8116811461121957600080fd5b6000602080838503121561138a57600080fd5b823567ffffffffffffffff808211156113a257600080fd5b90840190606082870312156113b657600080fd5b6113be611298565b8235828111156113cd57600080fd5b83019150601f820187136113e057600080fd5b81356113f36113ee82611333565b6112e4565b81815260069190911b8301850190858101908983111561141257600080fd5b938601935b82851015611464576040858b0312156114305760008081fd5b6114386112c1565b61144186611202565b815261144e888701611357565b8189015282526040949094019390860190611417565b835250611474905083850161121e565b8482015261148460408401611357565b60408201529695505050505050565b6000602082840312156114a557600080fd5b610bc58261121e565b600060208083850312156114c157600080fd5b823567ffffffffffffffff8111156114d857600080fd5b8301601f810185136114e957600080fd5b80356114f76113ee82611333565b81815260059190911b8201830190838101908783111561151657600080fd5b928401925b8284101561153b5761152c84611202565b8252928401929084019061151b565b979650505050505050565b60008060006060848603121561155b57600080fd5b61156484611202565b925061157260208501611202565b9150604084013590509250925092565b6020808252825182820181905260009190848201906040850190845b818110156115c35783516001600160a01b03168352928401929184019160010161159e565b50909695505050505050565b6000602082840312156115e157600080fd5b610bc582611202565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561162b5761162b6115ea565b500390565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611668576116686115ea565b500290565b6000826116a3577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611708576117086115ea565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c634300080f000a",
}

var PriceRegistryABI = PriceRegistryMetaData.ABI

var PriceRegistryBin = PriceRegistryMetaData.Bin

func DeployPriceRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, priceUpdates InternalPriceUpdates, priceUpdaters []common.Address, stalenessThreshold uint32) (common.Address, *types.Transaction, *PriceRegistry, error) {
	parsed, err := PriceRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PriceRegistryBin), backend, priceUpdates, priceUpdaters, stalenessThreshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceRegistry{PriceRegistryCaller: PriceRegistryCaller{contract: contract}, PriceRegistryTransactor: PriceRegistryTransactor{contract: contract}, PriceRegistryFilterer: PriceRegistryFilterer{contract: contract}}, nil
}

type PriceRegistry struct {
	address common.Address
	abi     abi.ABI
	PriceRegistryCaller
	PriceRegistryTransactor
	PriceRegistryFilterer
}

type PriceRegistryCaller struct {
	contract *bind.BoundContract
}

type PriceRegistryTransactor struct {
	contract *bind.BoundContract
}

type PriceRegistryFilterer struct {
	contract *bind.BoundContract
}

type PriceRegistrySession struct {
	Contract     *PriceRegistry
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type PriceRegistryCallerSession struct {
	Contract *PriceRegistryCaller
	CallOpts bind.CallOpts
}

type PriceRegistryTransactorSession struct {
	Contract     *PriceRegistryTransactor
	TransactOpts bind.TransactOpts
}

type PriceRegistryRaw struct {
	Contract *PriceRegistry
}

type PriceRegistryCallerRaw struct {
	Contract *PriceRegistryCaller
}

type PriceRegistryTransactorRaw struct {
	Contract *PriceRegistryTransactor
}

func NewPriceRegistry(address common.Address, backend bind.ContractBackend) (*PriceRegistry, error) {
	abi, err := abi.JSON(strings.NewReader(PriceRegistryABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindPriceRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceRegistry{address: address, abi: abi, PriceRegistryCaller: PriceRegistryCaller{contract: contract}, PriceRegistryTransactor: PriceRegistryTransactor{contract: contract}, PriceRegistryFilterer: PriceRegistryFilterer{contract: contract}}, nil
}

func NewPriceRegistryCaller(address common.Address, caller bind.ContractCaller) (*PriceRegistryCaller, error) {
	contract, err := bindPriceRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryCaller{contract: contract}, nil
}

func NewPriceRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceRegistryTransactor, error) {
	contract, err := bindPriceRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryTransactor{contract: contract}, nil
}

func NewPriceRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceRegistryFilterer, error) {
	contract, err := bindPriceRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryFilterer{contract: contract}, nil
}

func bindPriceRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_PriceRegistry *PriceRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceRegistry.Contract.PriceRegistryCaller.contract.Call(opts, result, method, params...)
}

func (_PriceRegistry *PriceRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRegistry.Contract.PriceRegistryTransactor.contract.Transfer(opts)
}

func (_PriceRegistry *PriceRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceRegistry.Contract.PriceRegistryTransactor.contract.Transact(opts, method, params...)
}

func (_PriceRegistry *PriceRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceRegistry.Contract.contract.Call(opts, result, method, params...)
}

func (_PriceRegistry *PriceRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRegistry.Contract.contract.Transfer(opts)
}

func (_PriceRegistry *PriceRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceRegistry.Contract.contract.Transact(opts, method, params...)
}

func (_PriceRegistry *PriceRegistryCaller) ConvertFeeTokenAmountToLinkAmount(opts *bind.CallOpts, linkToken common.Address, feeToken common.Address, feeTokenAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "convertFeeTokenAmountToLinkAmount", linkToken, feeToken, feeTokenAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) ConvertFeeTokenAmountToLinkAmount(linkToken common.Address, feeToken common.Address, feeTokenAmount *big.Int) (*big.Int, error) {
	return _PriceRegistry.Contract.ConvertFeeTokenAmountToLinkAmount(&_PriceRegistry.CallOpts, linkToken, feeToken, feeTokenAmount)
}

func (_PriceRegistry *PriceRegistryCallerSession) ConvertFeeTokenAmountToLinkAmount(linkToken common.Address, feeToken common.Address, feeTokenAmount *big.Int) (*big.Int, error) {
	return _PriceRegistry.Contract.ConvertFeeTokenAmountToLinkAmount(&_PriceRegistry.CallOpts, linkToken, feeToken, feeTokenAmount)
}

func (_PriceRegistry *PriceRegistryCaller) GetDestinationChainGasPrice(opts *bind.CallOpts, destChainId uint64) (IPriceRegistryTimestampedUint128Value, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getDestinationChainGasPrice", destChainId)

	if err != nil {
		return *new(IPriceRegistryTimestampedUint128Value), err
	}

	out0 := *abi.ConvertType(out[0], new(IPriceRegistryTimestampedUint128Value)).(*IPriceRegistryTimestampedUint128Value)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetDestinationChainGasPrice(destChainId uint64) (IPriceRegistryTimestampedUint128Value, error) {
	return _PriceRegistry.Contract.GetDestinationChainGasPrice(&_PriceRegistry.CallOpts, destChainId)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetDestinationChainGasPrice(destChainId uint64) (IPriceRegistryTimestampedUint128Value, error) {
	return _PriceRegistry.Contract.GetDestinationChainGasPrice(&_PriceRegistry.CallOpts, destChainId)
}

func (_PriceRegistry *PriceRegistryCaller) GetFeeTokenBaseUnitsPerUnitGas(opts *bind.CallOpts, token common.Address, destChainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getFeeTokenBaseUnitsPerUnitGas", token, destChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetFeeTokenBaseUnitsPerUnitGas(token common.Address, destChainId uint64) (*big.Int, error) {
	return _PriceRegistry.Contract.GetFeeTokenBaseUnitsPerUnitGas(&_PriceRegistry.CallOpts, token, destChainId)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetFeeTokenBaseUnitsPerUnitGas(token common.Address, destChainId uint64) (*big.Int, error) {
	return _PriceRegistry.Contract.GetFeeTokenBaseUnitsPerUnitGas(&_PriceRegistry.CallOpts, token, destChainId)
}

func (_PriceRegistry *PriceRegistryCaller) GetFeeTokenPrice(opts *bind.CallOpts, token common.Address) (IPriceRegistryTimestampedUint128Value, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getFeeTokenPrice", token)

	if err != nil {
		return *new(IPriceRegistryTimestampedUint128Value), err
	}

	out0 := *abi.ConvertType(out[0], new(IPriceRegistryTimestampedUint128Value)).(*IPriceRegistryTimestampedUint128Value)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetFeeTokenPrice(token common.Address) (IPriceRegistryTimestampedUint128Value, error) {
	return _PriceRegistry.Contract.GetFeeTokenPrice(&_PriceRegistry.CallOpts, token)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetFeeTokenPrice(token common.Address) (IPriceRegistryTimestampedUint128Value, error) {
	return _PriceRegistry.Contract.GetFeeTokenPrice(&_PriceRegistry.CallOpts, token)
}

func (_PriceRegistry *PriceRegistryCaller) GetPriceUpdaters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getPriceUpdaters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetPriceUpdaters() ([]common.Address, error) {
	return _PriceRegistry.Contract.GetPriceUpdaters(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetPriceUpdaters() ([]common.Address, error) {
	return _PriceRegistry.Contract.GetPriceUpdaters(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryCaller) GetStalenessThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getStalenessThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetStalenessThreshold() (*big.Int, error) {
	return _PriceRegistry.Contract.GetStalenessThreshold(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetStalenessThreshold() (*big.Int, error) {
	return _PriceRegistry.Contract.GetStalenessThreshold(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) Owner() (common.Address, error) {
	return _PriceRegistry.Contract.Owner(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryCallerSession) Owner() (common.Address, error) {
	return _PriceRegistry.Contract.Owner(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "acceptOwnership")
}

func (_PriceRegistry *PriceRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _PriceRegistry.Contract.AcceptOwnership(&_PriceRegistry.TransactOpts)
}

func (_PriceRegistry *PriceRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PriceRegistry.Contract.AcceptOwnership(&_PriceRegistry.TransactOpts)
}

func (_PriceRegistry *PriceRegistryTransactor) AddPriceUpdaters(opts *bind.TransactOpts, priceUpdaters []common.Address) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "addPriceUpdaters", priceUpdaters)
}

func (_PriceRegistry *PriceRegistrySession) AddPriceUpdaters(priceUpdaters []common.Address) (*types.Transaction, error) {
	return _PriceRegistry.Contract.AddPriceUpdaters(&_PriceRegistry.TransactOpts, priceUpdaters)
}

func (_PriceRegistry *PriceRegistryTransactorSession) AddPriceUpdaters(priceUpdaters []common.Address) (*types.Transaction, error) {
	return _PriceRegistry.Contract.AddPriceUpdaters(&_PriceRegistry.TransactOpts, priceUpdaters)
}

func (_PriceRegistry *PriceRegistryTransactor) RemovePriceUpdaters(opts *bind.TransactOpts, priceUpdaters []common.Address) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "removePriceUpdaters", priceUpdaters)
}

func (_PriceRegistry *PriceRegistrySession) RemovePriceUpdaters(priceUpdaters []common.Address) (*types.Transaction, error) {
	return _PriceRegistry.Contract.RemovePriceUpdaters(&_PriceRegistry.TransactOpts, priceUpdaters)
}

func (_PriceRegistry *PriceRegistryTransactorSession) RemovePriceUpdaters(priceUpdaters []common.Address) (*types.Transaction, error) {
	return _PriceRegistry.Contract.RemovePriceUpdaters(&_PriceRegistry.TransactOpts, priceUpdaters)
}

func (_PriceRegistry *PriceRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "transferOwnership", to)
}

func (_PriceRegistry *PriceRegistrySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _PriceRegistry.Contract.TransferOwnership(&_PriceRegistry.TransactOpts, to)
}

func (_PriceRegistry *PriceRegistryTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _PriceRegistry.Contract.TransferOwnership(&_PriceRegistry.TransactOpts, to)
}

func (_PriceRegistry *PriceRegistryTransactor) UpdatePrices(opts *bind.TransactOpts, priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "updatePrices", priceUpdates)
}

func (_PriceRegistry *PriceRegistrySession) UpdatePrices(priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _PriceRegistry.Contract.UpdatePrices(&_PriceRegistry.TransactOpts, priceUpdates)
}

func (_PriceRegistry *PriceRegistryTransactorSession) UpdatePrices(priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _PriceRegistry.Contract.UpdatePrices(&_PriceRegistry.TransactOpts, priceUpdates)
}

type PriceRegistryOwnershipTransferRequestedIterator struct {
	Event *PriceRegistryOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryOwnershipTransferRequested)
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
		it.Event = new(PriceRegistryOwnershipTransferRequested)
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

func (it *PriceRegistryOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PriceRegistryOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryOwnershipTransferRequestedIterator{contract: _PriceRegistry.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *PriceRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryOwnershipTransferRequested)
				if err := _PriceRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseOwnershipTransferRequested(log types.Log) (*PriceRegistryOwnershipTransferRequested, error) {
	event := new(PriceRegistryOwnershipTransferRequested)
	if err := _PriceRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryOwnershipTransferredIterator struct {
	Event *PriceRegistryOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryOwnershipTransferred)
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
		it.Event = new(PriceRegistryOwnershipTransferred)
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

func (it *PriceRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PriceRegistryOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryOwnershipTransferredIterator{contract: _PriceRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryOwnershipTransferred)
				if err := _PriceRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*PriceRegistryOwnershipTransferred, error) {
	event := new(PriceRegistryOwnershipTransferred)
	if err := _PriceRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryPriceUpdaterRemovedIterator struct {
	Event *PriceRegistryPriceUpdaterRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryPriceUpdaterRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryPriceUpdaterRemoved)
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
		it.Event = new(PriceRegistryPriceUpdaterRemoved)
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

func (it *PriceRegistryPriceUpdaterRemovedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryPriceUpdaterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryPriceUpdaterRemoved struct {
	PriceUpdater common.Address
	Raw          types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterPriceUpdaterRemoved(opts *bind.FilterOpts, priceUpdater []common.Address) (*PriceRegistryPriceUpdaterRemovedIterator, error) {

	var priceUpdaterRule []interface{}
	for _, priceUpdaterItem := range priceUpdater {
		priceUpdaterRule = append(priceUpdaterRule, priceUpdaterItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "PriceUpdaterRemoved", priceUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryPriceUpdaterRemovedIterator{contract: _PriceRegistry.contract, event: "PriceUpdaterRemoved", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchPriceUpdaterRemoved(opts *bind.WatchOpts, sink chan<- *PriceRegistryPriceUpdaterRemoved, priceUpdater []common.Address) (event.Subscription, error) {

	var priceUpdaterRule []interface{}
	for _, priceUpdaterItem := range priceUpdater {
		priceUpdaterRule = append(priceUpdaterRule, priceUpdaterItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "PriceUpdaterRemoved", priceUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryPriceUpdaterRemoved)
				if err := _PriceRegistry.contract.UnpackLog(event, "PriceUpdaterRemoved", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParsePriceUpdaterRemoved(log types.Log) (*PriceRegistryPriceUpdaterRemoved, error) {
	event := new(PriceRegistryPriceUpdaterRemoved)
	if err := _PriceRegistry.contract.UnpackLog(event, "PriceUpdaterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryPriceUpdaterSetIterator struct {
	Event *PriceRegistryPriceUpdaterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryPriceUpdaterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryPriceUpdaterSet)
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
		it.Event = new(PriceRegistryPriceUpdaterSet)
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

func (it *PriceRegistryPriceUpdaterSetIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryPriceUpdaterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryPriceUpdaterSet struct {
	PriceUpdater common.Address
	Raw          types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterPriceUpdaterSet(opts *bind.FilterOpts, priceUpdater []common.Address) (*PriceRegistryPriceUpdaterSetIterator, error) {

	var priceUpdaterRule []interface{}
	for _, priceUpdaterItem := range priceUpdater {
		priceUpdaterRule = append(priceUpdaterRule, priceUpdaterItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "PriceUpdaterSet", priceUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryPriceUpdaterSetIterator{contract: _PriceRegistry.contract, event: "PriceUpdaterSet", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchPriceUpdaterSet(opts *bind.WatchOpts, sink chan<- *PriceRegistryPriceUpdaterSet, priceUpdater []common.Address) (event.Subscription, error) {

	var priceUpdaterRule []interface{}
	for _, priceUpdaterItem := range priceUpdater {
		priceUpdaterRule = append(priceUpdaterRule, priceUpdaterItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "PriceUpdaterSet", priceUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryPriceUpdaterSet)
				if err := _PriceRegistry.contract.UnpackLog(event, "PriceUpdaterSet", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParsePriceUpdaterSet(log types.Log) (*PriceRegistryPriceUpdaterSet, error) {
	event := new(PriceRegistryPriceUpdaterSet)
	if err := _PriceRegistry.contract.UnpackLog(event, "PriceUpdaterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryUsdPerFeeTokenUpdatedIterator struct {
	Event *PriceRegistryUsdPerFeeTokenUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryUsdPerFeeTokenUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryUsdPerFeeTokenUpdated)
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
		it.Event = new(PriceRegistryUsdPerFeeTokenUpdated)
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

func (it *PriceRegistryUsdPerFeeTokenUpdatedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryUsdPerFeeTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryUsdPerFeeTokenUpdated struct {
	FeeToken  common.Address
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterUsdPerFeeTokenUpdated(opts *bind.FilterOpts, feeToken []common.Address) (*PriceRegistryUsdPerFeeTokenUpdatedIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "UsdPerFeeTokenUpdated", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryUsdPerFeeTokenUpdatedIterator{contract: _PriceRegistry.contract, event: "UsdPerFeeTokenUpdated", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchUsdPerFeeTokenUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryUsdPerFeeTokenUpdated, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "UsdPerFeeTokenUpdated", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryUsdPerFeeTokenUpdated)
				if err := _PriceRegistry.contract.UnpackLog(event, "UsdPerFeeTokenUpdated", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseUsdPerFeeTokenUpdated(log types.Log) (*PriceRegistryUsdPerFeeTokenUpdated, error) {
	event := new(PriceRegistryUsdPerFeeTokenUpdated)
	if err := _PriceRegistry.contract.UnpackLog(event, "UsdPerFeeTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryUsdPerUnitGasUpdatedIterator struct {
	Event *PriceRegistryUsdPerUnitGasUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryUsdPerUnitGasUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryUsdPerUnitGasUpdated)
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
		it.Event = new(PriceRegistryUsdPerUnitGasUpdated)
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

func (it *PriceRegistryUsdPerUnitGasUpdatedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryUsdPerUnitGasUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryUsdPerUnitGasUpdated struct {
	DestChain uint64
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterUsdPerUnitGasUpdated(opts *bind.FilterOpts, destChain []uint64) (*PriceRegistryUsdPerUnitGasUpdatedIterator, error) {

	var destChainRule []interface{}
	for _, destChainItem := range destChain {
		destChainRule = append(destChainRule, destChainItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "UsdPerUnitGasUpdated", destChainRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryUsdPerUnitGasUpdatedIterator{contract: _PriceRegistry.contract, event: "UsdPerUnitGasUpdated", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchUsdPerUnitGasUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryUsdPerUnitGasUpdated, destChain []uint64) (event.Subscription, error) {

	var destChainRule []interface{}
	for _, destChainItem := range destChain {
		destChainRule = append(destChainRule, destChainItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "UsdPerUnitGasUpdated", destChainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryUsdPerUnitGasUpdated)
				if err := _PriceRegistry.contract.UnpackLog(event, "UsdPerUnitGasUpdated", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseUsdPerUnitGasUpdated(log types.Log) (*PriceRegistryUsdPerUnitGasUpdated, error) {
	event := new(PriceRegistryUsdPerUnitGasUpdated)
	if err := _PriceRegistry.contract.UnpackLog(event, "UsdPerUnitGasUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_PriceRegistry *PriceRegistry) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _PriceRegistry.abi.Events["OwnershipTransferRequested"].ID:
		return _PriceRegistry.ParseOwnershipTransferRequested(log)
	case _PriceRegistry.abi.Events["OwnershipTransferred"].ID:
		return _PriceRegistry.ParseOwnershipTransferred(log)
	case _PriceRegistry.abi.Events["PriceUpdaterRemoved"].ID:
		return _PriceRegistry.ParsePriceUpdaterRemoved(log)
	case _PriceRegistry.abi.Events["PriceUpdaterSet"].ID:
		return _PriceRegistry.ParsePriceUpdaterSet(log)
	case _PriceRegistry.abi.Events["UsdPerFeeTokenUpdated"].ID:
		return _PriceRegistry.ParseUsdPerFeeTokenUpdated(log)
	case _PriceRegistry.abi.Events["UsdPerUnitGasUpdated"].ID:
		return _PriceRegistry.ParseUsdPerUnitGasUpdated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (PriceRegistryOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (PriceRegistryOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (PriceRegistryPriceUpdaterRemoved) Topic() common.Hash {
	return common.HexToHash("0xff7dbb85c77ca68ca1f894d6498570e3d5095cd19466f07ee8d222b337e4068c")
}

func (PriceRegistryPriceUpdaterSet) Topic() common.Hash {
	return common.HexToHash("0x34a02290b7920078c19f58e94b78c77eb9cc10195b20676e19bd3b82085893b8")
}

func (PriceRegistryUsdPerFeeTokenUpdated) Topic() common.Hash {
	return common.HexToHash("0x6ee3f6f5d4054ba011bcd5c323ed7b7fa0126b839aecb3f4e470fc81d5d4c4de")
}

func (PriceRegistryUsdPerUnitGasUpdated) Topic() common.Hash {
	return common.HexToHash("0xdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e")
}

func (_PriceRegistry *PriceRegistry) Address() common.Address {
	return _PriceRegistry.address
}

type PriceRegistryInterface interface {
	ConvertFeeTokenAmountToLinkAmount(opts *bind.CallOpts, linkToken common.Address, feeToken common.Address, feeTokenAmount *big.Int) (*big.Int, error)

	GetDestinationChainGasPrice(opts *bind.CallOpts, destChainId uint64) (IPriceRegistryTimestampedUint128Value, error)

	GetFeeTokenBaseUnitsPerUnitGas(opts *bind.CallOpts, token common.Address, destChainId uint64) (*big.Int, error)

	GetFeeTokenPrice(opts *bind.CallOpts, token common.Address) (IPriceRegistryTimestampedUint128Value, error)

	GetPriceUpdaters(opts *bind.CallOpts) ([]common.Address, error)

	GetStalenessThreshold(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPriceUpdaters(opts *bind.TransactOpts, priceUpdaters []common.Address) (*types.Transaction, error)

	RemovePriceUpdaters(opts *bind.TransactOpts, priceUpdaters []common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdatePrices(opts *bind.TransactOpts, priceUpdates InternalPriceUpdates) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PriceRegistryOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *PriceRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*PriceRegistryOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PriceRegistryOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*PriceRegistryOwnershipTransferred, error)

	FilterPriceUpdaterRemoved(opts *bind.FilterOpts, priceUpdater []common.Address) (*PriceRegistryPriceUpdaterRemovedIterator, error)

	WatchPriceUpdaterRemoved(opts *bind.WatchOpts, sink chan<- *PriceRegistryPriceUpdaterRemoved, priceUpdater []common.Address) (event.Subscription, error)

	ParsePriceUpdaterRemoved(log types.Log) (*PriceRegistryPriceUpdaterRemoved, error)

	FilterPriceUpdaterSet(opts *bind.FilterOpts, priceUpdater []common.Address) (*PriceRegistryPriceUpdaterSetIterator, error)

	WatchPriceUpdaterSet(opts *bind.WatchOpts, sink chan<- *PriceRegistryPriceUpdaterSet, priceUpdater []common.Address) (event.Subscription, error)

	ParsePriceUpdaterSet(log types.Log) (*PriceRegistryPriceUpdaterSet, error)

	FilterUsdPerFeeTokenUpdated(opts *bind.FilterOpts, feeToken []common.Address) (*PriceRegistryUsdPerFeeTokenUpdatedIterator, error)

	WatchUsdPerFeeTokenUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryUsdPerFeeTokenUpdated, feeToken []common.Address) (event.Subscription, error)

	ParseUsdPerFeeTokenUpdated(log types.Log) (*PriceRegistryUsdPerFeeTokenUpdated, error)

	FilterUsdPerUnitGasUpdated(opts *bind.FilterOpts, destChain []uint64) (*PriceRegistryUsdPerUnitGasUpdatedIterator, error)

	WatchUsdPerUnitGasUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryUsdPerUnitGasUpdated, destChain []uint64) (event.Subscription, error)

	ParseUsdPerUnitGasUpdated(log types.Log) (*PriceRegistryUsdPerUnitGasUpdated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
