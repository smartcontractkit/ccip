// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ccip_capability_configuration

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

type CCIPCapabilityConfigurationChainConfig struct {
	Readers [][32]byte
	Config  []byte
}

type CCIPCapabilityConfigurationChainConfigUpdate struct {
	ChainSelector uint64
	ChainConfig   CCIPCapabilityConfigurationChainConfig
}

type CCIPCapabilityConfigurationOCR3Config struct {
	PluginId              uint8
	ChainSelector         uint64
	Signers               [][][]byte
	Transmitters          [][][]byte
	F                     uint8
	OnchainConfig         []byte
	OffchainConfigVersion uint64
	OffchainConfig        []byte
}

var CCIPCapabilityConfigurationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"capabilityRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainSelectorNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainSelectorNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfigOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"donId\",\"type\":\"uint32\"}],\"name\":\"NoCapabilityConfigurationSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"p2pId\",\"type\":\"bytes32\"}],\"name\":\"NodeNotInRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCapabilityRegistryCanCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gotLength\",\"type\":\"uint256\"}],\"name\":\"SignerP2PIdPairMustBeLengthTwo\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CapabilityConfigurationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainConfigRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"readers\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structCCIPCapabilityConfiguration.ChainConfig\",\"name\":\"chainConfig\",\"type\":\"tuple\"}],\"name\":\"ChainConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"readers\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPCapabilityConfiguration.ChainConfig\",\"name\":\"chainConfig\",\"type\":\"tuple\"}],\"internalType\":\"structCCIPCapabilityConfiguration.ChainConfigUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"readers\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPCapabilityConfiguration.ChainConfig\",\"name\":\"chainConfig\",\"type\":\"tuple\"}],\"internalType\":\"structCCIPCapabilityConfiguration.ChainConfigUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"}],\"name\":\"applyChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"nodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"donId\",\"type\":\"uint32\"}],\"name\":\"beforeCapabilityConfigSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllChainConfigs\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"readers\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPCapabilityConfiguration.ChainConfig[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllOCRConfigs\",\"outputs\":[{\"components\":[{\"internalType\":\"enumCCIPCapabilityConfiguration.PluginId\",\"name\":\"pluginId\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes[][]\",\"name\":\"signers\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes[][]\",\"name\":\"transmitters\",\"type\":\"bytes[][]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"internalType\":\"structCCIPCapabilityConfiguration.OCR3Config[][]\",\"name\":\"\",\"type\":\"tuple[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getCapabilityConfiguration\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"configuration\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200283d3803806200283d83398101604081905262000034916200017e565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000d3565b5050506001600160a01b0316608052620001b0565b336001600160a01b038216036200012d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000602082840312156200019157600080fd5b81516001600160a01b0381168114620001a957600080fd5b9392505050565b60805161266a620001d3600039600081816112bc0152611453015261266a6000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063c97b8f551161005b578063c97b8f5514610109578063ddc042a81461011e578063f2fde38b14610133578063fba64a7c1461014657600080fd5b80635bc6cf881461008d57806379ba5097146100a2578063884efe61146100aa5780638da5cb5b146100e1575b600080fd5b6100a061009b366004611878565b610159565b005b6100a0610537565b6100cb6100b83660046118e4565b5060408051602081019091526000815290565b6040516100d89190611961565b60405180910390f35b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100d8565b610111610634565b6040516100d89190611a98565b6101266110d5565b6040516100d89190611c37565b6100a0610141366004611d1a565b611290565b6100a0610154366004611d78565b6112a4565b610161611375565b60005b83811015610359576101b385858381811061018157610181611e35565b90506020028101906101939190611e64565b6101a1906020810190611ea2565b60039067ffffffffffffffff166113f8565b61022d578484828181106101c9576101c9611e35565b90506020028101906101db9190611e64565b6101e9906020810190611ea2565b6040517f1bd4d2d200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024015b60405180910390fd5b6002600086868481811061024357610243611e35565b90506020028101906102559190611e64565b610263906020810190611ea2565b67ffffffffffffffff168152602081019190915260400160009081209061028a82826117bf565b6102986001830160006117dd565b50506102e18585838181106102af576102af611e35565b90506020028101906102c19190611e64565b6102cf906020810190611ea2565b60039067ffffffffffffffff16611415565b507f2a680691fef3b2d105196805935232c661ce703e92d464ef0b94a7bc62d714f085858381811061031557610315611e35565b90506020028101906103279190611e64565b610335906020810190611ea2565b60405167ffffffffffffffff909116815260200160405180910390a1600101610164565b5060005b8181101561053057600083838381811061037957610379611e35565b905060200281019061038b9190611e64565b610399906020810190611e64565b6103a39080611ebd565b8080602002602001604051908101604052809392919081815260200183836020028082843760009201829052509394508792508691508590508181106103eb576103eb611e35565b90506020028101906103fd9190611e64565b61040b906020810190611ea2565b905060005b825184101561044d5761043b83828151811061042e5761042e611e35565b6020026020010151611421565b8061044581611f54565b915050610410565b5084848481811061046057610460611e35565b90506020028101906104729190611e64565b610480906020810190611e64565b67ffffffffffffffff821660009081526002602052604090206104a382826121d8565b506104bb9050600367ffffffffffffffff831661153c565b507f0a93b217b9314a2a41b31bcc6cc8543471c1a3b3c4b9214afc1cf8f9987d2a50818686868181106104f0576104f0611e35565b90506020028101906105029190611e64565b610510906020810190611e64565b60405161051e929190612386565b60405180910390a1505060010161035d565b5050505050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146105b8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610224565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b606060006106426006611548565b67ffffffffffffffff81111561065a5761065a611f8c565b60405190808252806020026020018201604052801561068d57816020015b60608152602001906001900390816106785790505b50905060005b61069d6006611548565b8110156110cf5760006106b1600683611552565b63ffffffff8116600090815260056020908152604080832060018452909152808220548280529120549192506106e69161246e565b67ffffffffffffffff8111156106fe576106fe611f8c565b60405190808252806020026020018201604052801561076f57816020015b604080516101008101825260008082526020820181905260609282018390528282018390526080820181905260a0820183905260c082015260e081019190915281526020019060019003908161071c5790505b5083838151811061078257610782611e35565b602002602001018190525060005b63ffffffff82166000908152600560209081526040808320838052909152902054811015610c115763ffffffff8216600090815260056020908152604080832083805290915290208054829081106107ea576107ea611e35565b600091825260209091206040805161010081019091526007909202018054829060ff16600181111561081e5761081e611974565b600181111561082f5761082f611974565b81528154610100900467ffffffffffffffff16602080830191909152600183018054604080518285028101850182528281529401939260009084015b8282101561095e57838290600052602060002001805480602002602001604051908101604052809291908181526020016000905b8282101561094b5783829060005260206000200180546108be90612020565b80601f01602080910402602001604051908101604052809291908181526020018280546108ea90612020565b80156109375780601f1061090c57610100808354040283529160200191610937565b820191906000526020600020905b81548152906001019060200180831161091a57829003601f168201915b50505050508152602001906001019061089f565b505050508152602001906001019061086b565b50505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b82821015610a7e57838290600052602060002001805480602002602001604051908101604052809291908181526020016000905b82821015610a6b5783829060005260206000200180546109de90612020565b80601f0160208091040260200160405190810160405280929190818152602001828054610a0a90612020565b8015610a575780601f10610a2c57610100808354040283529160200191610a57565b820191906000526020600020905b815481529060010190602001808311610a3a57829003601f168201915b5050505050815260200190600101906109bf565b505050508152602001906001019061098b565b50505090825250600382015460ff166020820152600482018054604090920191610aa790612020565b80601f0160208091040260200160405190810160405280929190818152602001828054610ad390612020565b8015610b205780601f10610af557610100808354040283529160200191610b20565b820191906000526020600020905b815481529060010190602001808311610b0357829003601f168201915b5050509183525050600582015467ffffffffffffffff166020820152600682018054604090920191610b5190612020565b80601f0160208091040260200160405190810160405280929190818152602001828054610b7d90612020565b8015610bca5780601f10610b9f57610100808354040283529160200191610bca565b820191906000526020600020905b815481529060010190602001808311610bad57829003601f168201915b505050505081525050848481518110610be557610be5611e35565b60200260200101518281518110610bfe57610bfe611e35565b6020908102919091010152600101610790565b5060005b63ffffffff82166000908152600560209081526040808320600184529091529020548110156110c55763ffffffff82166000908152600560209081526040808320600184529091529020805482908110610c7157610c71611e35565b600091825260209091206040805161010081019091526007909202018054829060ff166001811115610ca557610ca5611974565b6001811115610cb657610cb6611974565b81528154610100900467ffffffffffffffff16602080830191909152600183018054604080518285028101850182528281529401939260009084015b82821015610de557838290600052602060002001805480602002602001604051908101604052809291908181526020016000905b82821015610dd2578382906000526020600020018054610d4590612020565b80601f0160208091040260200160405190810160405280929190818152602001828054610d7190612020565b8015610dbe5780601f10610d9357610100808354040283529160200191610dbe565b820191906000526020600020905b815481529060010190602001808311610da157829003601f168201915b505050505081526020019060010190610d26565b5050505081526020019060010190610cf2565b50505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b82821015610f0557838290600052602060002001805480602002602001604051908101604052809291908181526020016000905b82821015610ef2578382906000526020600020018054610e6590612020565b80601f0160208091040260200160405190810160405280929190818152602001828054610e9190612020565b8015610ede5780601f10610eb357610100808354040283529160200191610ede565b820191906000526020600020905b815481529060010190602001808311610ec157829003601f168201915b505050505081526020019060010190610e46565b5050505081526020019060010190610e12565b50505090825250600382015460ff166020820152600482018054604090920191610f2e90612020565b80601f0160208091040260200160405190810160405280929190818152602001828054610f5a90612020565b8015610fa75780601f10610f7c57610100808354040283529160200191610fa7565b820191906000526020600020905b815481529060010190602001808311610f8a57829003601f168201915b5050509183525050600582015467ffffffffffffffff166020820152600682018054604090920191610fd890612020565b80601f016020809104026020016040519081016040528092919081815260200182805461100490612020565b80156110515780601f1061102657610100808354040283529160200191611051565b820191906000526020600020905b81548152906001019060200180831161103457829003601f168201915b50505050508152505084848151811061106c5761106c611e35565b60209081029190910181015163ffffffff8516600090815260058352604080822082805290935291909120546110a2908461246e565b815181106110b2576110b2611e35565b6020908102919091010152600101610c15565b5050600101610693565b50919050565b606060006110e36003611548565b67ffffffffffffffff8111156110fb576110fb611f8c565b60405190808252806020026020018201604052801561114057816020015b60408051808201909152606080825260208201528152602001906001900390816111195790505b50905060005b6111506003611548565b8110156110cf5760026000611166600384611552565b67ffffffffffffffff16815260208082019290925260409081016000208151815460609481028201850184529281018381529093919284928491908401828280156111d057602002820191906000526020600020905b8154815260200190600101908083116111bc575b505050505081526020016001820180546111e990612020565b80601f016020809104026020016040519081016040528092919081815260200182805461121590612020565b80156112625780601f1061123757610100808354040283529160200191611262565b820191906000526020600020905b81548152906001019060200180831161124557829003601f168201915b50505050508152505082828151811061127d5761127d611e35565b6020908102919091010152600101611146565b611298611375565b6112a18161155e565b50565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614611313576040517f7b2485a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f756e696d706c656d656e746564000000000000000000000000000000000000006044820152606401610224565b60005473ffffffffffffffffffffffffffffffffffffffff1633146113f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610224565b565b600081815260018301602052604081205415155b90505b92915050565b600061140c8383611653565b6040517f50c946fe000000000000000000000000000000000000000000000000000000008152600481018290526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906350c946fe90602401600060405180830381865afa1580156114af573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526114f59190810190612504565b50905081816040015114611538576040517f8907a4fa00000000000000000000000000000000000000000000000000000000815260048101839052602401610224565b5050565b600061140c8383611746565b600061140f825490565b600061140c8383611795565b3373ffffffffffffffffffffffffffffffffffffffff8216036115dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610224565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000818152600183016020526040812054801561173c57600061167760018361261b565b855490915060009061168b9060019061261b565b90508181146116f05760008660000182815481106116ab576116ab611e35565b90600052602060002001549050808760000184815481106116ce576116ce611e35565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806117015761170161262e565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061140f565b600091505061140f565b600081815260018301602052604081205461178d5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561140f565b50600061140f565b60008260000182815481106117ac576117ac611e35565b9060005260206000200154905092915050565b50805460008255906000526020600020908101906112a19190611813565b5080546117e990612020565b6000825580601f106117f9575050565b601f0160209004906000526020600020908101906112a191905b5b808211156118285760008155600101611814565b5090565b60008083601f84011261183e57600080fd5b50813567ffffffffffffffff81111561185657600080fd5b6020830191508360208260051b850101111561187157600080fd5b9250929050565b6000806000806040858703121561188e57600080fd5b843567ffffffffffffffff808211156118a657600080fd5b6118b28883890161182c565b909650945060208701359150808211156118cb57600080fd5b506118d88782880161182c565b95989497509550505050565b6000602082840312156118f657600080fd5b5035919050565b6000815180845260005b8181101561192357602081850181015186830182015201611907565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061140c60208301846118fd565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600281106119da577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b6000828251808552602080860195506005818360051b8501018287016000805b86811015611a89577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe088850381018c5283518051808752908801908887019080891b88018a01865b82811015611a7257858a8303018452611a608286516118fd565b948c0194938c01939150600101611a46565b509e8a019e975050509387019350506001016119fe565b50919998505050505050505050565b600060208083018184528085518083526040925060408601915060408160051b87010184880160005b83811015611c29578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018552815180518085529088019088850190600581901b86018a0160005b82811015611c13577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08883030184528451610100611b4a8483516119a3565b8d820151611b638f86018267ffffffffffffffff169052565b508c820151818e860152611b79828601826119de565b91505060608083015185830382870152611b9383826119de565b92505050608080830151611bab8287018260ff169052565b505060a08083015185830382870152611bc483826118fd565b9250505060c080830151611be38287018267ffffffffffffffff169052565b505060e08083015192508482038186015250611bff81836118fd565b968e0196958e019593505050600101611b0b565b50978a0197955050509187019150600101611ac1565b509098975050505050505050565b600060208083018184528085518083526040925060408601915060408160051b8701018488016000805b84811015611a89578984037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018652825180518886528051898701819052908a0190849060608801905b80831015611ccb5783518252928c019260019290920191908c0190611cab565b50928b0151878403888d015292611ce281856118fd565b998c019997505050938901935050600101611c61565b73ffffffffffffffffffffffffffffffffffffffff811681146112a157600080fd5b600060208284031215611d2c57600080fd5b8135611d3781611cf8565b9392505050565b803567ffffffffffffffff81168114611d5657600080fd5b919050565b63ffffffff811681146112a157600080fd5b8035611d5681611d5b565b60008060008060008060808789031215611d9157600080fd5b863567ffffffffffffffff80821115611da957600080fd5b611db58a838b0161182c565b90985096506020890135915080821115611dce57600080fd5b818901915089601f830112611de257600080fd5b813581811115611df157600080fd5b8a6020828501011115611e0357600080fd5b602083019650809550505050611e1b60408801611d3e565b9150611e2960608801611d6d565b90509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc1833603018112611e9857600080fd5b9190910192915050565b600060208284031215611eb457600080fd5b61140c82611d3e565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611ef257600080fd5b83018035915067ffffffffffffffff821115611f0d57600080fd5b6020019150600581901b360382131561187157600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611f8557611f85611f25565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611ff057600080fd5b83018035915067ffffffffffffffff82111561200b57600080fd5b60200191503681900382131561187157600080fd5b600181811c9082168061203457607f821691505b6020821081036110cf577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f8211156120b9576000816000526020600020601f850160051c810160208610156120965750805b601f850160051c820191505b818110156120b5578281556001016120a2565b5050505b505050565b67ffffffffffffffff8311156120d6576120d6611f8c565b6120ea836120e48354612020565b8361206d565b6000601f84116001811461213c57600085156121065750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355610530565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b8281101561218b578685013582556020948501946001909201910161216b565b50868210156121c6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b81357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe183360301811261220a57600080fd5b8201803567ffffffffffffffff81111561222357600080fd5b602080830192508160051b360383131561223c57600080fd5b6801000000000000000082111561225557612255611f8c565b83548285558083101561228c576000858152602081208481019083015b8082101561228857828255600182019150612272565b5050505b5060008481526020902060005b838110156122b35784358282015593820193600101612299565b50505050506122c56020830183611fbb565b6122d38183600186016120be565b50505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261230e57600080fd5b830160208101925035905067ffffffffffffffff81111561232e57600080fd5b80360382131561187157600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b600067ffffffffffffffff80851683526040602084015283357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18536030181126123cf57600080fd5b8401602081019035828111156123e457600080fd5b8060051b92508236038213156123f957600080fd5b6040808601528060808601527f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81111561243257600080fd5b50818160a086013750820161244a60208501856122d9565b606085840301606086015261246360a08401828461233d565b979650505050505050565b8082018082111561140f5761140f611f25565b6040516080810167ffffffffffffffff811182821017156124a4576124a4611f8c565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156124f1576124f1611f8c565b604052919050565b8051611d5681611d5b565b6000806040838503121561251757600080fd5b825167ffffffffffffffff8082111561252f57600080fd5b908401906080828703121561254357600080fd5b61254b612481565b825161255681611d5b565b815260208381015161256781611cf8565b828201526040848101519083015260608401518381111561258757600080fd5b80850194505087601f85011261259c57600080fd5b8351838111156125ae576125ae611f8c565b8060051b93506125bf8285016124aa565b818152938501820193828101908a8611156125d957600080fd5b958301955b858710156125f7578651825295830195908301906125de565b80606086015250505081955061260e8188016124f9565b9450505050509250929050565b8181038181111561140f5761140f611f25565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000818000a",
}

var CCIPCapabilityConfigurationABI = CCIPCapabilityConfigurationMetaData.ABI

var CCIPCapabilityConfigurationBin = CCIPCapabilityConfigurationMetaData.Bin

func DeployCCIPCapabilityConfiguration(auth *bind.TransactOpts, backend bind.ContractBackend, capabilityRegistry common.Address) (common.Address, *types.Transaction, *CCIPCapabilityConfiguration, error) {
	parsed, err := CCIPCapabilityConfigurationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCIPCapabilityConfigurationBin), backend, capabilityRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCIPCapabilityConfiguration{address: address, abi: *parsed, CCIPCapabilityConfigurationCaller: CCIPCapabilityConfigurationCaller{contract: contract}, CCIPCapabilityConfigurationTransactor: CCIPCapabilityConfigurationTransactor{contract: contract}, CCIPCapabilityConfigurationFilterer: CCIPCapabilityConfigurationFilterer{contract: contract}}, nil
}

type CCIPCapabilityConfiguration struct {
	address common.Address
	abi     abi.ABI
	CCIPCapabilityConfigurationCaller
	CCIPCapabilityConfigurationTransactor
	CCIPCapabilityConfigurationFilterer
}

type CCIPCapabilityConfigurationCaller struct {
	contract *bind.BoundContract
}

type CCIPCapabilityConfigurationTransactor struct {
	contract *bind.BoundContract
}

type CCIPCapabilityConfigurationFilterer struct {
	contract *bind.BoundContract
}

type CCIPCapabilityConfigurationSession struct {
	Contract     *CCIPCapabilityConfiguration
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CCIPCapabilityConfigurationCallerSession struct {
	Contract *CCIPCapabilityConfigurationCaller
	CallOpts bind.CallOpts
}

type CCIPCapabilityConfigurationTransactorSession struct {
	Contract     *CCIPCapabilityConfigurationTransactor
	TransactOpts bind.TransactOpts
}

type CCIPCapabilityConfigurationRaw struct {
	Contract *CCIPCapabilityConfiguration
}

type CCIPCapabilityConfigurationCallerRaw struct {
	Contract *CCIPCapabilityConfigurationCaller
}

type CCIPCapabilityConfigurationTransactorRaw struct {
	Contract *CCIPCapabilityConfigurationTransactor
}

func NewCCIPCapabilityConfiguration(address common.Address, backend bind.ContractBackend) (*CCIPCapabilityConfiguration, error) {
	abi, err := abi.JSON(strings.NewReader(CCIPCapabilityConfigurationABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCCIPCapabilityConfiguration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCIPCapabilityConfiguration{address: address, abi: abi, CCIPCapabilityConfigurationCaller: CCIPCapabilityConfigurationCaller{contract: contract}, CCIPCapabilityConfigurationTransactor: CCIPCapabilityConfigurationTransactor{contract: contract}, CCIPCapabilityConfigurationFilterer: CCIPCapabilityConfigurationFilterer{contract: contract}}, nil
}

func NewCCIPCapabilityConfigurationCaller(address common.Address, caller bind.ContractCaller) (*CCIPCapabilityConfigurationCaller, error) {
	contract, err := bindCCIPCapabilityConfiguration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCIPCapabilityConfigurationCaller{contract: contract}, nil
}

func NewCCIPCapabilityConfigurationTransactor(address common.Address, transactor bind.ContractTransactor) (*CCIPCapabilityConfigurationTransactor, error) {
	contract, err := bindCCIPCapabilityConfiguration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCIPCapabilityConfigurationTransactor{contract: contract}, nil
}

func NewCCIPCapabilityConfigurationFilterer(address common.Address, filterer bind.ContractFilterer) (*CCIPCapabilityConfigurationFilterer, error) {
	contract, err := bindCCIPCapabilityConfiguration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCIPCapabilityConfigurationFilterer{contract: contract}, nil
}

func bindCCIPCapabilityConfiguration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCIPCapabilityConfigurationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCIPCapabilityConfiguration.Contract.CCIPCapabilityConfigurationCaller.contract.Call(opts, result, method, params...)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.Contract.CCIPCapabilityConfigurationTransactor.contract.Transfer(opts)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.Contract.CCIPCapabilityConfigurationTransactor.contract.Transact(opts, method, params...)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCIPCapabilityConfiguration.Contract.contract.Call(opts, result, method, params...)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.Contract.contract.Transfer(opts)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.Contract.contract.Transact(opts, method, params...)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationCaller) GetAllChainConfigs(opts *bind.CallOpts) ([]CCIPCapabilityConfigurationChainConfig, error) {
	var out []interface{}
	err := _CCIPCapabilityConfiguration.contract.Call(opts, &out, "getAllChainConfigs")

	if err != nil {
		return *new([]CCIPCapabilityConfigurationChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]CCIPCapabilityConfigurationChainConfig)).(*[]CCIPCapabilityConfigurationChainConfig)

	return out0, err

}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationSession) GetAllChainConfigs() ([]CCIPCapabilityConfigurationChainConfig, error) {
	return _CCIPCapabilityConfiguration.Contract.GetAllChainConfigs(&_CCIPCapabilityConfiguration.CallOpts)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationCallerSession) GetAllChainConfigs() ([]CCIPCapabilityConfigurationChainConfig, error) {
	return _CCIPCapabilityConfiguration.Contract.GetAllChainConfigs(&_CCIPCapabilityConfiguration.CallOpts)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationCaller) GetAllOCRConfigs(opts *bind.CallOpts) ([][]CCIPCapabilityConfigurationOCR3Config, error) {
	var out []interface{}
	err := _CCIPCapabilityConfiguration.contract.Call(opts, &out, "getAllOCRConfigs")

	if err != nil {
		return *new([][]CCIPCapabilityConfigurationOCR3Config), err
	}

	out0 := *abi.ConvertType(out[0], new([][]CCIPCapabilityConfigurationOCR3Config)).(*[][]CCIPCapabilityConfigurationOCR3Config)

	return out0, err

}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationSession) GetAllOCRConfigs() ([][]CCIPCapabilityConfigurationOCR3Config, error) {
	return _CCIPCapabilityConfiguration.Contract.GetAllOCRConfigs(&_CCIPCapabilityConfiguration.CallOpts)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationCallerSession) GetAllOCRConfigs() ([][]CCIPCapabilityConfigurationOCR3Config, error) {
	return _CCIPCapabilityConfiguration.Contract.GetAllOCRConfigs(&_CCIPCapabilityConfiguration.CallOpts)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationCaller) GetCapabilityConfiguration(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _CCIPCapabilityConfiguration.contract.Call(opts, &out, "getCapabilityConfiguration", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationSession) GetCapabilityConfiguration(arg0 *big.Int) ([]byte, error) {
	return _CCIPCapabilityConfiguration.Contract.GetCapabilityConfiguration(&_CCIPCapabilityConfiguration.CallOpts, arg0)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationCallerSession) GetCapabilityConfiguration(arg0 *big.Int) ([]byte, error) {
	return _CCIPCapabilityConfiguration.Contract.GetCapabilityConfiguration(&_CCIPCapabilityConfiguration.CallOpts, arg0)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCIPCapabilityConfiguration.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationSession) Owner() (common.Address, error) {
	return _CCIPCapabilityConfiguration.Contract.Owner(&_CCIPCapabilityConfiguration.CallOpts)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationCallerSession) Owner() (common.Address, error) {
	return _CCIPCapabilityConfiguration.Contract.Owner(&_CCIPCapabilityConfiguration.CallOpts)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.contract.Transact(opts, "acceptOwnership")
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.Contract.AcceptOwnership(&_CCIPCapabilityConfiguration.TransactOpts)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.Contract.AcceptOwnership(&_CCIPCapabilityConfiguration.TransactOpts)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationTransactor) ApplyChainConfigUpdates(opts *bind.TransactOpts, removes []CCIPCapabilityConfigurationChainConfigUpdate, adds []CCIPCapabilityConfigurationChainConfigUpdate) (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.contract.Transact(opts, "applyChainConfigUpdates", removes, adds)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationSession) ApplyChainConfigUpdates(removes []CCIPCapabilityConfigurationChainConfigUpdate, adds []CCIPCapabilityConfigurationChainConfigUpdate) (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.Contract.ApplyChainConfigUpdates(&_CCIPCapabilityConfiguration.TransactOpts, removes, adds)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationTransactorSession) ApplyChainConfigUpdates(removes []CCIPCapabilityConfigurationChainConfigUpdate, adds []CCIPCapabilityConfigurationChainConfigUpdate) (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.Contract.ApplyChainConfigUpdates(&_CCIPCapabilityConfiguration.TransactOpts, removes, adds)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationTransactor) BeforeCapabilityConfigSet(opts *bind.TransactOpts, nodes [][32]byte, config []byte, configCount uint64, donId uint32) (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.contract.Transact(opts, "beforeCapabilityConfigSet", nodes, config, configCount, donId)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationSession) BeforeCapabilityConfigSet(nodes [][32]byte, config []byte, configCount uint64, donId uint32) (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.Contract.BeforeCapabilityConfigSet(&_CCIPCapabilityConfiguration.TransactOpts, nodes, config, configCount, donId)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationTransactorSession) BeforeCapabilityConfigSet(nodes [][32]byte, config []byte, configCount uint64, donId uint32) (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.Contract.BeforeCapabilityConfigSet(&_CCIPCapabilityConfiguration.TransactOpts, nodes, config, configCount, donId)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.contract.Transact(opts, "transferOwnership", to)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.Contract.TransferOwnership(&_CCIPCapabilityConfiguration.TransactOpts, to)
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCIPCapabilityConfiguration.Contract.TransferOwnership(&_CCIPCapabilityConfiguration.TransactOpts, to)
}

type CCIPCapabilityConfigurationCapabilityConfigurationSetIterator struct {
	Event *CCIPCapabilityConfigurationCapabilityConfigurationSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPCapabilityConfigurationCapabilityConfigurationSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPCapabilityConfigurationCapabilityConfigurationSet)
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
		it.Event = new(CCIPCapabilityConfigurationCapabilityConfigurationSet)
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

func (it *CCIPCapabilityConfigurationCapabilityConfigurationSetIterator) Error() error {
	return it.fail
}

func (it *CCIPCapabilityConfigurationCapabilityConfigurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPCapabilityConfigurationCapabilityConfigurationSet struct {
	Raw types.Log
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) FilterCapabilityConfigurationSet(opts *bind.FilterOpts) (*CCIPCapabilityConfigurationCapabilityConfigurationSetIterator, error) {

	logs, sub, err := _CCIPCapabilityConfiguration.contract.FilterLogs(opts, "CapabilityConfigurationSet")
	if err != nil {
		return nil, err
	}
	return &CCIPCapabilityConfigurationCapabilityConfigurationSetIterator{contract: _CCIPCapabilityConfiguration.contract, event: "CapabilityConfigurationSet", logs: logs, sub: sub}, nil
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) WatchCapabilityConfigurationSet(opts *bind.WatchOpts, sink chan<- *CCIPCapabilityConfigurationCapabilityConfigurationSet) (event.Subscription, error) {

	logs, sub, err := _CCIPCapabilityConfiguration.contract.WatchLogs(opts, "CapabilityConfigurationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPCapabilityConfigurationCapabilityConfigurationSet)
				if err := _CCIPCapabilityConfiguration.contract.UnpackLog(event, "CapabilityConfigurationSet", log); err != nil {
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

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) ParseCapabilityConfigurationSet(log types.Log) (*CCIPCapabilityConfigurationCapabilityConfigurationSet, error) {
	event := new(CCIPCapabilityConfigurationCapabilityConfigurationSet)
	if err := _CCIPCapabilityConfiguration.contract.UnpackLog(event, "CapabilityConfigurationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPCapabilityConfigurationChainConfigRemovedIterator struct {
	Event *CCIPCapabilityConfigurationChainConfigRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPCapabilityConfigurationChainConfigRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPCapabilityConfigurationChainConfigRemoved)
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
		it.Event = new(CCIPCapabilityConfigurationChainConfigRemoved)
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

func (it *CCIPCapabilityConfigurationChainConfigRemovedIterator) Error() error {
	return it.fail
}

func (it *CCIPCapabilityConfigurationChainConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPCapabilityConfigurationChainConfigRemoved struct {
	ChainSelector uint64
	Raw           types.Log
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) FilterChainConfigRemoved(opts *bind.FilterOpts) (*CCIPCapabilityConfigurationChainConfigRemovedIterator, error) {

	logs, sub, err := _CCIPCapabilityConfiguration.contract.FilterLogs(opts, "ChainConfigRemoved")
	if err != nil {
		return nil, err
	}
	return &CCIPCapabilityConfigurationChainConfigRemovedIterator{contract: _CCIPCapabilityConfiguration.contract, event: "ChainConfigRemoved", logs: logs, sub: sub}, nil
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) WatchChainConfigRemoved(opts *bind.WatchOpts, sink chan<- *CCIPCapabilityConfigurationChainConfigRemoved) (event.Subscription, error) {

	logs, sub, err := _CCIPCapabilityConfiguration.contract.WatchLogs(opts, "ChainConfigRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPCapabilityConfigurationChainConfigRemoved)
				if err := _CCIPCapabilityConfiguration.contract.UnpackLog(event, "ChainConfigRemoved", log); err != nil {
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

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) ParseChainConfigRemoved(log types.Log) (*CCIPCapabilityConfigurationChainConfigRemoved, error) {
	event := new(CCIPCapabilityConfigurationChainConfigRemoved)
	if err := _CCIPCapabilityConfiguration.contract.UnpackLog(event, "ChainConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPCapabilityConfigurationChainConfigSetIterator struct {
	Event *CCIPCapabilityConfigurationChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPCapabilityConfigurationChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPCapabilityConfigurationChainConfigSet)
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
		it.Event = new(CCIPCapabilityConfigurationChainConfigSet)
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

func (it *CCIPCapabilityConfigurationChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCIPCapabilityConfigurationChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPCapabilityConfigurationChainConfigSet struct {
	ChainSelector uint64
	ChainConfig   CCIPCapabilityConfigurationChainConfig
	Raw           types.Log
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) FilterChainConfigSet(opts *bind.FilterOpts) (*CCIPCapabilityConfigurationChainConfigSetIterator, error) {

	logs, sub, err := _CCIPCapabilityConfiguration.contract.FilterLogs(opts, "ChainConfigSet")
	if err != nil {
		return nil, err
	}
	return &CCIPCapabilityConfigurationChainConfigSetIterator{contract: _CCIPCapabilityConfiguration.contract, event: "ChainConfigSet", logs: logs, sub: sub}, nil
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) WatchChainConfigSet(opts *bind.WatchOpts, sink chan<- *CCIPCapabilityConfigurationChainConfigSet) (event.Subscription, error) {

	logs, sub, err := _CCIPCapabilityConfiguration.contract.WatchLogs(opts, "ChainConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPCapabilityConfigurationChainConfigSet)
				if err := _CCIPCapabilityConfiguration.contract.UnpackLog(event, "ChainConfigSet", log); err != nil {
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

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) ParseChainConfigSet(log types.Log) (*CCIPCapabilityConfigurationChainConfigSet, error) {
	event := new(CCIPCapabilityConfigurationChainConfigSet)
	if err := _CCIPCapabilityConfiguration.contract.UnpackLog(event, "ChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPCapabilityConfigurationOwnershipTransferRequestedIterator struct {
	Event *CCIPCapabilityConfigurationOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPCapabilityConfigurationOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPCapabilityConfigurationOwnershipTransferRequested)
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
		it.Event = new(CCIPCapabilityConfigurationOwnershipTransferRequested)
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

func (it *CCIPCapabilityConfigurationOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CCIPCapabilityConfigurationOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPCapabilityConfigurationOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCIPCapabilityConfigurationOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCIPCapabilityConfiguration.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCIPCapabilityConfigurationOwnershipTransferRequestedIterator{contract: _CCIPCapabilityConfiguration.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCIPCapabilityConfigurationOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCIPCapabilityConfiguration.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPCapabilityConfigurationOwnershipTransferRequested)
				if err := _CCIPCapabilityConfiguration.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) ParseOwnershipTransferRequested(log types.Log) (*CCIPCapabilityConfigurationOwnershipTransferRequested, error) {
	event := new(CCIPCapabilityConfigurationOwnershipTransferRequested)
	if err := _CCIPCapabilityConfiguration.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCIPCapabilityConfigurationOwnershipTransferredIterator struct {
	Event *CCIPCapabilityConfigurationOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCIPCapabilityConfigurationOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCIPCapabilityConfigurationOwnershipTransferred)
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
		it.Event = new(CCIPCapabilityConfigurationOwnershipTransferred)
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

func (it *CCIPCapabilityConfigurationOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CCIPCapabilityConfigurationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCIPCapabilityConfigurationOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCIPCapabilityConfigurationOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCIPCapabilityConfiguration.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCIPCapabilityConfigurationOwnershipTransferredIterator{contract: _CCIPCapabilityConfiguration.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCIPCapabilityConfigurationOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCIPCapabilityConfiguration.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCIPCapabilityConfigurationOwnershipTransferred)
				if err := _CCIPCapabilityConfiguration.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfigurationFilterer) ParseOwnershipTransferred(log types.Log) (*CCIPCapabilityConfigurationOwnershipTransferred, error) {
	event := new(CCIPCapabilityConfigurationOwnershipTransferred)
	if err := _CCIPCapabilityConfiguration.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfiguration) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CCIPCapabilityConfiguration.abi.Events["CapabilityConfigurationSet"].ID:
		return _CCIPCapabilityConfiguration.ParseCapabilityConfigurationSet(log)
	case _CCIPCapabilityConfiguration.abi.Events["ChainConfigRemoved"].ID:
		return _CCIPCapabilityConfiguration.ParseChainConfigRemoved(log)
	case _CCIPCapabilityConfiguration.abi.Events["ChainConfigSet"].ID:
		return _CCIPCapabilityConfiguration.ParseChainConfigSet(log)
	case _CCIPCapabilityConfiguration.abi.Events["OwnershipTransferRequested"].ID:
		return _CCIPCapabilityConfiguration.ParseOwnershipTransferRequested(log)
	case _CCIPCapabilityConfiguration.abi.Events["OwnershipTransferred"].ID:
		return _CCIPCapabilityConfiguration.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CCIPCapabilityConfigurationCapabilityConfigurationSet) Topic() common.Hash {
	return common.HexToHash("0x84ad7751b744c9e2ee77da1d902b428aec7f0a343d67a24bbe2142e6f58a8d0f")
}

func (CCIPCapabilityConfigurationChainConfigRemoved) Topic() common.Hash {
	return common.HexToHash("0x2a680691fef3b2d105196805935232c661ce703e92d464ef0b94a7bc62d714f0")
}

func (CCIPCapabilityConfigurationChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x0a93b217b9314a2a41b31bcc6cc8543471c1a3b3c4b9214afc1cf8f9987d2a50")
}

func (CCIPCapabilityConfigurationOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCIPCapabilityConfigurationOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_CCIPCapabilityConfiguration *CCIPCapabilityConfiguration) Address() common.Address {
	return _CCIPCapabilityConfiguration.address
}

type CCIPCapabilityConfigurationInterface interface {
	GetAllChainConfigs(opts *bind.CallOpts) ([]CCIPCapabilityConfigurationChainConfig, error)

	GetAllOCRConfigs(opts *bind.CallOpts) ([][]CCIPCapabilityConfigurationOCR3Config, error)

	GetCapabilityConfiguration(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyChainConfigUpdates(opts *bind.TransactOpts, removes []CCIPCapabilityConfigurationChainConfigUpdate, adds []CCIPCapabilityConfigurationChainConfigUpdate) (*types.Transaction, error)

	BeforeCapabilityConfigSet(opts *bind.TransactOpts, nodes [][32]byte, config []byte, configCount uint64, donId uint32) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterCapabilityConfigurationSet(opts *bind.FilterOpts) (*CCIPCapabilityConfigurationCapabilityConfigurationSetIterator, error)

	WatchCapabilityConfigurationSet(opts *bind.WatchOpts, sink chan<- *CCIPCapabilityConfigurationCapabilityConfigurationSet) (event.Subscription, error)

	ParseCapabilityConfigurationSet(log types.Log) (*CCIPCapabilityConfigurationCapabilityConfigurationSet, error)

	FilterChainConfigRemoved(opts *bind.FilterOpts) (*CCIPCapabilityConfigurationChainConfigRemovedIterator, error)

	WatchChainConfigRemoved(opts *bind.WatchOpts, sink chan<- *CCIPCapabilityConfigurationChainConfigRemoved) (event.Subscription, error)

	ParseChainConfigRemoved(log types.Log) (*CCIPCapabilityConfigurationChainConfigRemoved, error)

	FilterChainConfigSet(opts *bind.FilterOpts) (*CCIPCapabilityConfigurationChainConfigSetIterator, error)

	WatchChainConfigSet(opts *bind.WatchOpts, sink chan<- *CCIPCapabilityConfigurationChainConfigSet) (event.Subscription, error)

	ParseChainConfigSet(log types.Log) (*CCIPCapabilityConfigurationChainConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCIPCapabilityConfigurationOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCIPCapabilityConfigurationOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCIPCapabilityConfigurationOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCIPCapabilityConfigurationOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCIPCapabilityConfigurationOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCIPCapabilityConfigurationOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
