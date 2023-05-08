// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package custom_token_pool

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

type RateLimiterConfig struct {
	IsEnabled bool
	Rate      *big.Int
	Capacity  *big.Int
}

type RateLimiterTokenBucket struct {
	Capacity    *big.Int
	Tokens      *big.Int
	Rate        *big.Int
	LastUpdated *big.Int
	IsEnabled   bool
}

type TokenPoolRampUpdate struct {
	Ramp    common.Address
	Allowed bool
}

var CustomTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionsError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"OffRampAllowanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"OnRampAllowanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SynthBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SynthMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"ramp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"internalType\":\"structTokenPool.RampUpdate[]\",\"name\":\"onRamps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ramp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"internalType\":\"structTokenPool.RampUpdate[]\",\"name\":\"offRamps\",\"type\":\"tuple[]\"}],\"name\":\"applyRampUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdated\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"isOnRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"lockOrBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"releaseOrMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001b2938038062001b29833981016040819052620000349162000275565b818133806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c081620001ad565b50506001805460ff60a01b19169055506001600160a01b038216620000f857604051634655efd160e11b815260040160405180910390fd5b6040805160a0810182528282018051808352905160208084018290528501516001600160d01b03169383018490524264ffffffffff1660608401819052945115156080938401819052600692909255600755600880547fff0000000000000000000000000000000000000000000000000000000000000016909317600160d01b909402939093176001600160f81b0316600160f81b9093029290921790556001600160a01b0391909116905250620003289050565b336001600160a01b03821603620002075760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160d01b03811681146200027057600080fd5b919050565b60008082840360808112156200028a57600080fd5b83516001600160a01b0381168114620002a257600080fd5b92506060601f1982011215620002b757600080fd5b50604051606081016001600160401b0381118282101715620002e957634e487b7160e01b600052604160045260246000fd5b604052602084015180151581146200030057600080fd5b8152620003106040850162000258565b60208201526060939093015160408401525092909150565b6080516117e562000344600039600061013701526117e56000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80636f32b872116100975780638da5cb5b116100665780638da5cb5b14610261578063968754451461027f578063af51911214610292578063f2fde38b146102a557600080fd5b80636f32b8721461022b57806379ba50971461023e5780638456cb59146102465780638627fad61461024e57600080fd5b80633091aee7116100d35780633091aee71461017c5780633f4ba83a14610191578063546719cd146101995780635c975abb1461020857600080fd5b806301ffc9a7146100fa5780631d7a74a01461012257806321df0da714610135575b600080fd5b61010d6101083660046111df565b6102b8565b60405190151581526020015b60405180910390f35b61010d61013036600461124a565b610351565b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610119565b61018f61018a36600461131c565b61035e565b005b61018f610374565b6101a1610386565b6040516101199190600060a082019050825182526020830151602083015279ffffffffffffffffffffffffffffffffffffffffffffffffffff604084015116604083015264ffffffffff606084015116606083015260808301511515608083015292915050565b60015474010000000000000000000000000000000000000000900460ff1661010d565b61010d61023936600461124a565b610455565b61018f610462565b61018f610564565b61018f61025c36600461144a565b610574565b60005473ffffffffffffffffffffffffffffffffffffffff16610157565b61018f61028d366004611522565b610672565b61018f6102a0366004611672565b610772565b61018f6102b336600461124a565b610982565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f317fa33400000000000000000000000000000000000000000000000000000000148061034b57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b600061034b600483610993565b6103666109c5565b610371600682610a46565b50565b61037c6109c5565b610384610b72565b565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526006548152600754602082015260085479ffffffffffffffffffffffffffffffffffffffffffffffffffff8116928201929092527a010000000000000000000000000000000000000000000000000000820464ffffffffff1660608201527f010000000000000000000000000000000000000000000000000000000000000090910460ff161515608082015261045090610c6b565b905090565b600061034b600283610993565b60015473ffffffffffffffffffffffffffffffffffffffff1633146104e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61056c6109c5565b610384610d0e565b60015474010000000000000000000000000000000000000000900460ff16156105f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016104df565b61060233610351565b610638576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518381527fbb0b72e5f44e331506684da008a30e10d50658c29d8159f6c6ab40bf1e52e6009060200160405180910390a15050505050565b60015474010000000000000000000000000000000000000000900460ff16156106f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016104df565b61070033610455565b610736576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518481527f02992093bca69a36949677658a77d359b510dc6232c68f9f118f7c0127a1b1479060200160405180910390a150505050505050565b61077a6109c5565b60005b825181101561087b57600083828151811061079a5761079a6116d6565b6020026020010151905080602001516107c05780516107bb90600290610dfa565b6107ce565b80516107ce90600290610e1c565b1561086a577fbceff8f229c6dfcbf8bdcfb18726b84b0fd249b4803deb3948ff34d904013662848381518110610806576108066116d6565b602002602001015160000151858481518110610824576108246116d6565b60200260200101516020015160405161086192919073ffffffffffffffffffffffffffffffffffffffff9290921682521515602082015260400190565b60405180910390a15b5061087481611734565b905061077d565b5060005b815181101561097d57600082828151811061089c5761089c6116d6565b6020026020010151905080602001516108c25780516108bd90600490610dfa565b6108d0565b80516108d090600490610e1c565b1561096c577fd8c3333ded377884ced3869cd0bcb9be54ea664076df1f5d39c4689120313648838381518110610908576109086116d6565b602002602001015160000151848481518110610926576109266116d6565b60200260200101516020015160405161096392919073ffffffffffffffffffffffffffffffffffffffff9290921682521515602082015260400190565b60405180910390a15b5061097681611734565b905061087f565b505050565b61098a6109c5565b61037181610e3e565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018301602052604081205415155b9392505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610384576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016104df565b610a4f82610f33565b8051600283018054604084015180865560208501517effffffffff00000000000000000000000000000000000000000000000000009092167f0100000000000000000000000000000000000000000000000000000000000000941515949094027fffffffffffff0000000000000000000000000000000000000000000000000000169390931779ffffffffffffffffffffffffffffffffffffffffffffffffffff9091161790556001830154610b059190611087565b60018301556040805182511515815260208084015179ffffffffffffffffffffffffffffffffffffffffffffffffffff1690820152828201518183015290517f44a2350342338075ac038f37b8d9e49e696e360492cb44cc6bc37fc117f19df89181900360600190a15050565b60015474010000000000000000000000000000000000000000900460ff16610bf6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016104df565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526000826060015164ffffffffff1642610caf919061176c565b83516040850151919250610cf791610ce39079ffffffffffffffffffffffffffffffffffffffffffffffffffff168461177f565b8560200151610cf29190611796565b611087565b6020840152505064ffffffffff4216606082015290565b60015474010000000000000000000000000000000000000000900460ff1615610d93576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016104df565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610c413390565b60006109be8373ffffffffffffffffffffffffffffffffffffffff841661109d565b60006109be8373ffffffffffffffffffffffffffffffffffffffff8416611190565b3373ffffffffffffffffffffffffffffffffffffffff821603610ebd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016104df565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b805460018201541480610f6f5750600281015464ffffffffff7a0100000000000000000000000000000000000000000000000000009091041642145b15610f775750565b805460018201541115610fb6576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002810154600090610fed907a010000000000000000000000000000000000000000000000000000900464ffffffffff164261176c565b82546002840154919250611030916110219079ffffffffffffffffffffffffffffffffffffffffffffffffffff168461177f565b8460010154610cf29190611796565b60018301555060020180547fff0000000000ffffffffffffffffffffffffffffffffffffffffffffffffffff167a0100000000000000000000000000000000000000000000000000004264ffffffffff1602179055565b600081831061109657816109be565b5090919050565b600081815260018301602052604081205480156111865760006110c160018361176c565b85549091506000906110d59060019061176c565b905081811461113a5760008660000182815481106110f5576110f56116d6565b9060005260206000200154905080876000018481548110611118576111186116d6565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061114b5761114b6117a9565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061034b565b600091505061034b565b60008181526001830160205260408120546111d75750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561034b565b50600061034b565b6000602082840312156111f157600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146109be57600080fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461124557600080fd5b919050565b60006020828403121561125c57600080fd5b6109be82611221565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156112b7576112b7611265565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561130457611304611265565b604052919050565b8035801515811461124557600080fd5b60006060828403121561132e57600080fd5b6040516060810181811067ffffffffffffffff8211171561135157611351611265565b60405261135d8361130c565b8152602083013579ffffffffffffffffffffffffffffffffffffffffffffffffffff8116811461138c57600080fd5b60208201526040928301359281019290925250919050565b600082601f8301126113b557600080fd5b813567ffffffffffffffff8111156113cf576113cf611265565b61140060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016112bd565b81815284602083860101111561141557600080fd5b816020850160208301376000918101602001919091529392505050565b803567ffffffffffffffff8116811461124557600080fd5b600080600080600060a0868803121561146257600080fd5b853567ffffffffffffffff8082111561147a57600080fd5b61148689838a016113a4565b965061149460208901611221565b9550604088013594506114a960608901611432565b935060808801359150808211156114bf57600080fd5b506114cc888289016113a4565b9150509295509295909350565b60008083601f8401126114eb57600080fd5b50813567ffffffffffffffff81111561150357600080fd5b60208301915083602082850101111561151b57600080fd5b9250929050565b600080600080600080600060a0888a03121561153d57600080fd5b61154688611221565b9650602088013567ffffffffffffffff8082111561156357600080fd5b61156f8b838c016114d9565b909850965060408a0135955086915061158a60608b01611432565b945060808a01359150808211156115a057600080fd5b506115ad8a828b016114d9565b989b979a50959850939692959293505050565b600082601f8301126115d157600080fd5b8135602067ffffffffffffffff8211156115ed576115ed611265565b6115fb818360051b016112bd565b82815260069290921b8401810191818101908684111561161a57600080fd5b8286015b8481101561166757604081890312156116375760008081fd5b61163f611294565b61164882611221565b815261165585830161130c565b8186015283529183019160400161161e565b509695505050505050565b6000806040838503121561168557600080fd5b823567ffffffffffffffff8082111561169d57600080fd5b6116a9868387016115c0565b935060208501359150808211156116bf57600080fd5b506116cc858286016115c0565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361176557611765611705565b5060010190565b8181038181111561034b5761034b611705565b808202811582820484141761034b5761034b611705565b8082018082111561034b5761034b611705565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000813000a",
}

var CustomTokenPoolABI = CustomTokenPoolMetaData.ABI

var CustomTokenPoolBin = CustomTokenPoolMetaData.Bin

func DeployCustomTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, rateLimiterConfig RateLimiterConfig) (common.Address, *types.Transaction, *CustomTokenPool, error) {
	parsed, err := CustomTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CustomTokenPoolBin), backend, token, rateLimiterConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CustomTokenPool{CustomTokenPoolCaller: CustomTokenPoolCaller{contract: contract}, CustomTokenPoolTransactor: CustomTokenPoolTransactor{contract: contract}, CustomTokenPoolFilterer: CustomTokenPoolFilterer{contract: contract}}, nil
}

type CustomTokenPool struct {
	address common.Address
	abi     abi.ABI
	CustomTokenPoolCaller
	CustomTokenPoolTransactor
	CustomTokenPoolFilterer
}

type CustomTokenPoolCaller struct {
	contract *bind.BoundContract
}

type CustomTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type CustomTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type CustomTokenPoolSession struct {
	Contract     *CustomTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CustomTokenPoolCallerSession struct {
	Contract *CustomTokenPoolCaller
	CallOpts bind.CallOpts
}

type CustomTokenPoolTransactorSession struct {
	Contract     *CustomTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type CustomTokenPoolRaw struct {
	Contract *CustomTokenPool
}

type CustomTokenPoolCallerRaw struct {
	Contract *CustomTokenPoolCaller
}

type CustomTokenPoolTransactorRaw struct {
	Contract *CustomTokenPoolTransactor
}

func NewCustomTokenPool(address common.Address, backend bind.ContractBackend) (*CustomTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(CustomTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCustomTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPool{address: address, abi: abi, CustomTokenPoolCaller: CustomTokenPoolCaller{contract: contract}, CustomTokenPoolTransactor: CustomTokenPoolTransactor{contract: contract}, CustomTokenPoolFilterer: CustomTokenPoolFilterer{contract: contract}}, nil
}

func NewCustomTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*CustomTokenPoolCaller, error) {
	contract, err := bindCustomTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolCaller{contract: contract}, nil
}

func NewCustomTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*CustomTokenPoolTransactor, error) {
	contract, err := bindCustomTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolTransactor{contract: contract}, nil
}

func NewCustomTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*CustomTokenPoolFilterer, error) {
	contract, err := bindCustomTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolFilterer{contract: contract}, nil
}

func bindCustomTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CustomTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CustomTokenPool *CustomTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustomTokenPool.Contract.CustomTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_CustomTokenPool *CustomTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.CustomTokenPoolTransactor.contract.Transfer(opts)
}

func (_CustomTokenPool *CustomTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.CustomTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_CustomTokenPool *CustomTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustomTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_CustomTokenPool *CustomTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.contract.Transfer(opts)
}

func (_CustomTokenPool *CustomTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_CustomTokenPool *CustomTokenPoolCaller) CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "currentRateLimiterState")

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _CustomTokenPool.Contract.CurrentRateLimiterState(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _CustomTokenPool.Contract.CurrentRateLimiterState(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) GetToken() (common.Address, error) {
	return _CustomTokenPool.Contract.GetToken(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _CustomTokenPool.Contract.GetToken(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCaller) IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "isOffRamp", offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _CustomTokenPool.Contract.IsOffRamp(&_CustomTokenPool.CallOpts, offRamp)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _CustomTokenPool.Contract.IsOffRamp(&_CustomTokenPool.CallOpts, offRamp)
}

func (_CustomTokenPool *CustomTokenPoolCaller) IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "isOnRamp", onRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _CustomTokenPool.Contract.IsOnRamp(&_CustomTokenPool.CallOpts, onRamp)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _CustomTokenPool.Contract.IsOnRamp(&_CustomTokenPool.CallOpts, onRamp)
}

func (_CustomTokenPool *CustomTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) Owner() (common.Address, error) {
	return _CustomTokenPool.Contract.Owner(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) Owner() (common.Address, error) {
	return _CustomTokenPool.Contract.Owner(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) Paused() (bool, error) {
	return _CustomTokenPool.Contract.Paused(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) Paused() (bool, error) {
	return _CustomTokenPool.Contract.Paused(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CustomTokenPool.Contract.SupportsInterface(&_CustomTokenPool.CallOpts, interfaceId)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CustomTokenPool.Contract.SupportsInterface(&_CustomTokenPool.CallOpts, interfaceId)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_CustomTokenPool *CustomTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.AcceptOwnership(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.AcceptOwnership(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) ApplyRampUpdates(opts *bind.TransactOpts, onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "applyRampUpdates", onRamps, offRamps)
}

func (_CustomTokenPool *CustomTokenPoolSession) ApplyRampUpdates(onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.ApplyRampUpdates(&_CustomTokenPool.TransactOpts, onRamps, offRamps)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) ApplyRampUpdates(onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.ApplyRampUpdates(&_CustomTokenPool.TransactOpts, onRamps, offRamps)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, arg0 common.Address, arg1 []byte, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "lockOrBurn", arg0, arg1, amount, arg3, arg4)
}

func (_CustomTokenPool *CustomTokenPoolSession) LockOrBurn(arg0 common.Address, arg1 []byte, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.LockOrBurn(&_CustomTokenPool.TransactOpts, arg0, arg1, amount, arg3, arg4)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) LockOrBurn(arg0 common.Address, arg1 []byte, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.LockOrBurn(&_CustomTokenPool.TransactOpts, arg0, arg1, amount, arg3, arg4)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "pause")
}

func (_CustomTokenPool *CustomTokenPoolSession) Pause() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.Pause(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) Pause() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.Pause(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, arg0 []byte, arg1 common.Address, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "releaseOrMint", arg0, arg1, amount, arg3, arg4)
}

func (_CustomTokenPool *CustomTokenPoolSession) ReleaseOrMint(arg0 []byte, arg1 common.Address, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.ReleaseOrMint(&_CustomTokenPool.TransactOpts, arg0, arg1, amount, arg3, arg4)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) ReleaseOrMint(arg0 []byte, arg1 common.Address, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.ReleaseOrMint(&_CustomTokenPool.TransactOpts, arg0, arg1, amount, arg3, arg4)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_CustomTokenPool *CustomTokenPoolSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.SetRateLimiterConfig(&_CustomTokenPool.TransactOpts, config)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.SetRateLimiterConfig(&_CustomTokenPool.TransactOpts, config)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_CustomTokenPool *CustomTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.TransferOwnership(&_CustomTokenPool.TransactOpts, to)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.TransferOwnership(&_CustomTokenPool.TransactOpts, to)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "unpause")
}

func (_CustomTokenPool *CustomTokenPoolSession) Unpause() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.Unpause(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.Unpause(&_CustomTokenPool.TransactOpts)
}

type CustomTokenPoolBurnedIterator struct {
	Event *CustomTokenPoolBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolBurned)
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
		it.Event = new(CustomTokenPoolBurned)
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

func (it *CustomTokenPoolBurnedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*CustomTokenPoolBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolBurnedIterator{contract: _CustomTokenPool.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolBurned)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseBurned(log types.Log) (*CustomTokenPoolBurned, error) {
	event := new(CustomTokenPoolBurned)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolLockedIterator struct {
	Event *CustomTokenPoolLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolLocked)
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
		it.Event = new(CustomTokenPoolLocked)
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

func (it *CustomTokenPoolLockedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*CustomTokenPoolLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolLockedIterator{contract: _CustomTokenPool.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolLocked)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseLocked(log types.Log) (*CustomTokenPoolLocked, error) {
	event := new(CustomTokenPoolLocked)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolMintedIterator struct {
	Event *CustomTokenPoolMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolMinted)
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
		it.Event = new(CustomTokenPoolMinted)
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

func (it *CustomTokenPoolMintedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*CustomTokenPoolMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolMintedIterator{contract: _CustomTokenPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolMinted)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseMinted(log types.Log) (*CustomTokenPoolMinted, error) {
	event := new(CustomTokenPoolMinted)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolOffRampAllowanceSetIterator struct {
	Event *CustomTokenPoolOffRampAllowanceSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolOffRampAllowanceSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolOffRampAllowanceSet)
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
		it.Event = new(CustomTokenPoolOffRampAllowanceSet)
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

func (it *CustomTokenPoolOffRampAllowanceSetIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolOffRampAllowanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolOffRampAllowanceSet struct {
	OnRamp  common.Address
	Allowed bool
	Raw     types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterOffRampAllowanceSet(opts *bind.FilterOpts) (*CustomTokenPoolOffRampAllowanceSetIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "OffRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolOffRampAllowanceSetIterator{contract: _CustomTokenPool.contract, event: "OffRampAllowanceSet", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchOffRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOffRampAllowanceSet) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "OffRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolOffRampAllowanceSet)
				if err := _CustomTokenPool.contract.UnpackLog(event, "OffRampAllowanceSet", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseOffRampAllowanceSet(log types.Log) (*CustomTokenPoolOffRampAllowanceSet, error) {
	event := new(CustomTokenPoolOffRampAllowanceSet)
	if err := _CustomTokenPool.contract.UnpackLog(event, "OffRampAllowanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolOnRampAllowanceSetIterator struct {
	Event *CustomTokenPoolOnRampAllowanceSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolOnRampAllowanceSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolOnRampAllowanceSet)
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
		it.Event = new(CustomTokenPoolOnRampAllowanceSet)
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

func (it *CustomTokenPoolOnRampAllowanceSetIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolOnRampAllowanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolOnRampAllowanceSet struct {
	OnRamp  common.Address
	Allowed bool
	Raw     types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterOnRampAllowanceSet(opts *bind.FilterOpts) (*CustomTokenPoolOnRampAllowanceSetIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "OnRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolOnRampAllowanceSetIterator{contract: _CustomTokenPool.contract, event: "OnRampAllowanceSet", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchOnRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOnRampAllowanceSet) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "OnRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolOnRampAllowanceSet)
				if err := _CustomTokenPool.contract.UnpackLog(event, "OnRampAllowanceSet", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseOnRampAllowanceSet(log types.Log) (*CustomTokenPoolOnRampAllowanceSet, error) {
	event := new(CustomTokenPoolOnRampAllowanceSet)
	if err := _CustomTokenPool.contract.UnpackLog(event, "OnRampAllowanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolOwnershipTransferRequestedIterator struct {
	Event *CustomTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolOwnershipTransferRequested)
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
		it.Event = new(CustomTokenPoolOwnershipTransferRequested)
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

func (it *CustomTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CustomTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolOwnershipTransferRequestedIterator{contract: _CustomTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolOwnershipTransferRequested)
				if err := _CustomTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*CustomTokenPoolOwnershipTransferRequested, error) {
	event := new(CustomTokenPoolOwnershipTransferRequested)
	if err := _CustomTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolOwnershipTransferredIterator struct {
	Event *CustomTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolOwnershipTransferred)
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
		it.Event = new(CustomTokenPoolOwnershipTransferred)
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

func (it *CustomTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CustomTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolOwnershipTransferredIterator{contract: _CustomTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolOwnershipTransferred)
				if err := _CustomTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*CustomTokenPoolOwnershipTransferred, error) {
	event := new(CustomTokenPoolOwnershipTransferred)
	if err := _CustomTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolPausedIterator struct {
	Event *CustomTokenPoolPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolPaused)
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
		it.Event = new(CustomTokenPoolPaused)
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

func (it *CustomTokenPoolPausedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*CustomTokenPoolPausedIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolPausedIterator{contract: _CustomTokenPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolPaused) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolPaused)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParsePaused(log types.Log) (*CustomTokenPoolPaused, error) {
	event := new(CustomTokenPoolPaused)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolReleasedIterator struct {
	Event *CustomTokenPoolReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolReleased)
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
		it.Event = new(CustomTokenPoolReleased)
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

func (it *CustomTokenPoolReleasedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*CustomTokenPoolReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolReleasedIterator{contract: _CustomTokenPool.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolReleased)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseReleased(log types.Log) (*CustomTokenPoolReleased, error) {
	event := new(CustomTokenPoolReleased)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolSynthBurnedIterator struct {
	Event *CustomTokenPoolSynthBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolSynthBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolSynthBurned)
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
		it.Event = new(CustomTokenPoolSynthBurned)
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

func (it *CustomTokenPoolSynthBurnedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolSynthBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolSynthBurned struct {
	Amount *big.Int
	Raw    types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterSynthBurned(opts *bind.FilterOpts) (*CustomTokenPoolSynthBurnedIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "SynthBurned")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolSynthBurnedIterator{contract: _CustomTokenPool.contract, event: "SynthBurned", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchSynthBurned(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolSynthBurned) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "SynthBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolSynthBurned)
				if err := _CustomTokenPool.contract.UnpackLog(event, "SynthBurned", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseSynthBurned(log types.Log) (*CustomTokenPoolSynthBurned, error) {
	event := new(CustomTokenPoolSynthBurned)
	if err := _CustomTokenPool.contract.UnpackLog(event, "SynthBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolSynthMintedIterator struct {
	Event *CustomTokenPoolSynthMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolSynthMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolSynthMinted)
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
		it.Event = new(CustomTokenPoolSynthMinted)
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

func (it *CustomTokenPoolSynthMintedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolSynthMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolSynthMinted struct {
	Amount *big.Int
	Raw    types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterSynthMinted(opts *bind.FilterOpts) (*CustomTokenPoolSynthMintedIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "SynthMinted")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolSynthMintedIterator{contract: _CustomTokenPool.contract, event: "SynthMinted", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchSynthMinted(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolSynthMinted) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "SynthMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolSynthMinted)
				if err := _CustomTokenPool.contract.UnpackLog(event, "SynthMinted", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseSynthMinted(log types.Log) (*CustomTokenPoolSynthMinted, error) {
	event := new(CustomTokenPoolSynthMinted)
	if err := _CustomTokenPool.contract.UnpackLog(event, "SynthMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolUnpausedIterator struct {
	Event *CustomTokenPoolUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolUnpaused)
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
		it.Event = new(CustomTokenPoolUnpaused)
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

func (it *CustomTokenPoolUnpausedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CustomTokenPoolUnpausedIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolUnpausedIterator{contract: _CustomTokenPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolUnpaused)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseUnpaused(log types.Log) (*CustomTokenPoolUnpaused, error) {
	event := new(CustomTokenPoolUnpaused)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_CustomTokenPool *CustomTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CustomTokenPool.abi.Events["Burned"].ID:
		return _CustomTokenPool.ParseBurned(log)
	case _CustomTokenPool.abi.Events["Locked"].ID:
		return _CustomTokenPool.ParseLocked(log)
	case _CustomTokenPool.abi.Events["Minted"].ID:
		return _CustomTokenPool.ParseMinted(log)
	case _CustomTokenPool.abi.Events["OffRampAllowanceSet"].ID:
		return _CustomTokenPool.ParseOffRampAllowanceSet(log)
	case _CustomTokenPool.abi.Events["OnRampAllowanceSet"].ID:
		return _CustomTokenPool.ParseOnRampAllowanceSet(log)
	case _CustomTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _CustomTokenPool.ParseOwnershipTransferRequested(log)
	case _CustomTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _CustomTokenPool.ParseOwnershipTransferred(log)
	case _CustomTokenPool.abi.Events["Paused"].ID:
		return _CustomTokenPool.ParsePaused(log)
	case _CustomTokenPool.abi.Events["Released"].ID:
		return _CustomTokenPool.ParseReleased(log)
	case _CustomTokenPool.abi.Events["SynthBurned"].ID:
		return _CustomTokenPool.ParseSynthBurned(log)
	case _CustomTokenPool.abi.Events["SynthMinted"].ID:
		return _CustomTokenPool.ParseSynthMinted(log)
	case _CustomTokenPool.abi.Events["Unpaused"].ID:
		return _CustomTokenPool.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CustomTokenPoolBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
}

func (CustomTokenPoolLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (CustomTokenPoolMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (CustomTokenPoolOffRampAllowanceSet) Topic() common.Hash {
	return common.HexToHash("0xd8c3333ded377884ced3869cd0bcb9be54ea664076df1f5d39c4689120313648")
}

func (CustomTokenPoolOnRampAllowanceSet) Topic() common.Hash {
	return common.HexToHash("0xbceff8f229c6dfcbf8bdcfb18726b84b0fd249b4803deb3948ff34d904013662")
}

func (CustomTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CustomTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CustomTokenPoolPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (CustomTokenPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (CustomTokenPoolSynthBurned) Topic() common.Hash {
	return common.HexToHash("0x02992093bca69a36949677658a77d359b510dc6232c68f9f118f7c0127a1b147")
}

func (CustomTokenPoolSynthMinted) Topic() common.Hash {
	return common.HexToHash("0xbb0b72e5f44e331506684da008a30e10d50658c29d8159f6c6ab40bf1e52e600")
}

func (CustomTokenPoolUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_CustomTokenPool *CustomTokenPool) Address() common.Address {
	return _CustomTokenPool.address
}

type CustomTokenPoolInterface interface {
	CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyRampUpdates(opts *bind.TransactOpts, onRamps []TokenPoolRampUpdate, offRamps []TokenPoolRampUpdate) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, arg0 common.Address, arg1 []byte, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, arg0 []byte, arg1 common.Address, amount *big.Int, arg3 uint64, arg4 []byte) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*CustomTokenPoolBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*CustomTokenPoolBurned, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*CustomTokenPoolLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*CustomTokenPoolLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*CustomTokenPoolMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*CustomTokenPoolMinted, error)

	FilterOffRampAllowanceSet(opts *bind.FilterOpts) (*CustomTokenPoolOffRampAllowanceSetIterator, error)

	WatchOffRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOffRampAllowanceSet) (event.Subscription, error)

	ParseOffRampAllowanceSet(log types.Log) (*CustomTokenPoolOffRampAllowanceSet, error)

	FilterOnRampAllowanceSet(opts *bind.FilterOpts) (*CustomTokenPoolOnRampAllowanceSetIterator, error)

	WatchOnRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOnRampAllowanceSet) (event.Subscription, error)

	ParseOnRampAllowanceSet(log types.Log) (*CustomTokenPoolOnRampAllowanceSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CustomTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CustomTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CustomTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CustomTokenPoolOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*CustomTokenPoolPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*CustomTokenPoolPaused, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*CustomTokenPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*CustomTokenPoolReleased, error)

	FilterSynthBurned(opts *bind.FilterOpts) (*CustomTokenPoolSynthBurnedIterator, error)

	WatchSynthBurned(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolSynthBurned) (event.Subscription, error)

	ParseSynthBurned(log types.Log) (*CustomTokenPoolSynthBurned, error)

	FilterSynthMinted(opts *bind.FilterOpts) (*CustomTokenPoolSynthMintedIterator, error)

	WatchSynthMinted(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolSynthMinted) (event.Subscription, error)

	ParseSynthMinted(log types.Log) (*CustomTokenPoolSynthMinted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*CustomTokenPoolUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*CustomTokenPoolUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
