// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package native_token_pool

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

type TokenLimitsTokenBucket struct {
	Rate        *big.Int
	Capacity    *big.Int
	Tokens      *big.Int
	LastUpdated *big.Int
}

var NativeTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockBucketRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockBucketCapacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseBucketRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseBucketCapacity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"full\",\"type\":\"bool\"}],\"name\":\"NewLockBurnBucketConstructed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"full\",\"type\":\"bool\"}],\"name\":\"NewReleaseMintBucketConstructed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLockOrBurnBucket\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structTokenLimits.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReleaseOrMintBucket\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structTokenLimits.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"isOnRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockOrBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"releaseOrMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"full\",\"type\":\"bool\"}],\"name\":\"setLockOrBurnBucket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"permission\",\"type\":\"bool\"}],\"name\":\"setOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"permission\",\"type\":\"bool\"}],\"name\":\"setOnRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"full\",\"type\":\"bool\"}],\"name\":\"setReleaseOrMintBucket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200189f3803806200189f833981016040819052620000349162000274565b84848484843380600081620000905760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c357620000c38162000162565b50506001805460ff60a01b191681556001600160a01b038716608052620000fb9150859085906200020e602090811b620008e617901c565b80516004556020808201516005556040820151600655606090910151600755620001359083908390600190620008e66200020e821b17901c565b805160085560208101516009556040810151600a5560600151600b5550620002cb98505050505050505050565b6001600160a01b038116331415620001bd5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000087565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6200023a6040518060800160405280600081526020016000815260200160008152602001600081525090565b6000826200024a5760006200024c565b835b6040805160808101825296875260208701959095529385019390935250504260608301525090565b600080600080600060a086880312156200028d57600080fd5b85516001600160a01b0381168114620002a557600080fd5b602087015160408801516060890151608090990151929a91995097965090945092505050565b6080516115aa620002f560003960008181610170015281816106dd015261080801526115aa6000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c80638456cb59116100b2578063cc8af2e811610081578063eb54b3bf11610066578063eb54b3bf146102e3578063f0c6ff26146102f6578063f2fde38b146102fe57600080fd5b8063cc8af2e8146102bd578063ea6192a2146102d057600080fd5b80638456cb591461027157806384f52501146102795780638da5cb5b1461028c578063bd4612c4146102aa57600080fd5b80635c975abb116100ee5780635c975abb146101fa57806369e946d41461021d5780636f32b8721461023057806379ba50971461026957600080fd5b80631d7a74a01461012057806321df0da71461016e57806335e1e1e8146101b55780633f4ba83a146101f0575b600080fd5b61015961012e366004611383565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205460ff1690565b60405190151581526020015b60405180910390f35b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610165565b6101bd610311565b60405161016591908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6101f8610369565b005b60015474010000000000000000000000000000000000000000900460ff16610159565b6101f861022b3660046113ae565b61037b565b61015961023e366004611383565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460ff1690565b6101f86103d9565b6101f86104db565b6101f86102873660046113e7565b6104eb565b60005473ffffffffffffffffffffffffffffffffffffffff16610190565b6101f86102b83660046113e7565b610561565b6101f86102cb3660046113ae565b6105ce565b6101f86102de366004611420565b61062c565b6101f86102f1366004611420565b610757565b6101bd61087a565b6101f861030c366004611383565b6108d2565b61033c6040518060800160405280600081526020016000815260200160008152602001600081525090565b50604080516080810182526004548152600554602082015260065491810191909152600754606082015290565b61037161094f565b6103796109d0565b565b61038361094f565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260036020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b60015473ffffffffffffffffffffffffffffffffffffffff16331461045f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6104e361094f565b610379610ac9565b6104f361094f565b6104fe8383836108e6565b80516004556020808201516005556040808301516006556060928301516007558051868152918201859052831515908201527f4bab560bebbbda366298430a0e5cd55e85ea36dc92834dc281b93dc21b9683a291015b60405180910390a1505050565b61056961094f565b6105748383836108e6565b8051600855602080820151600955604080830151600a55606092830151600b558051868152918201859052831515908201527f8586e57eb58352bf12165ebd6104b5991cb1c22a183a9965a2961b7b462bc2df9101610554565b6105d661094f565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b60015474010000000000000000000000000000000000000000900460ff16156106b1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610456565b806106ba610bb5565b6106c381610c25565b61070473ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168484610c74565b60405182815273ffffffffffffffffffffffffffffffffffffffff84169033907f2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52906020015b60405180910390a3505050565b60015474010000000000000000000000000000000000000000900460ff16156107dc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610456565b806107e5610d4d565b6107ee81610d88565b61083073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016843085610dd7565b60405182815273ffffffffffffffffffffffffffffffffffffffff84169033907f989eaa915cbb416ea3d6f9a63b1a3de51770c7674b11fe21ecdf76b4e1d139109060200161074a565b6108a56040518060800160405280600081526020016000815260200160008152602001600081525090565b506040805160808101825260085481526009546020820152600a5491810191909152600b54606082015290565b6108da61094f565b6108e381610e3b565b50565b6109116040518060800160405280600081526020016000815260200160008152602001600081525090565b60008261091f576000610921565b835b90506040518060800160405280868152602001858152602001828152602001428152509150505b9392505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610379576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610456565b60015474010000000000000000000000000000000000000000900460ff16610a54576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610456565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60015474010000000000000000000000000000000000000000900460ff1615610b4e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610456565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610a9f3390565b60005473ffffffffffffffffffffffffffffffffffffffff163314801590610bee57503360009081526003602052604090205460ff165b155b15610379576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c30600882610f31565b6108e357600a546040517f331220f7000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610456565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610d489084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610f74565b505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314801590610bee57503360009081526002602052604090205460ff16610bec565b610d93600482610f31565b6108e3576006546040517f331220f7000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610456565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052610e359085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401610cc6565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8116331415610ebb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610456565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000610f3c83611080565b8183600201541015610f5057506000610f6e565b81836002016000828254610f64919061147b565b9091555060019150505b92915050565b6000610fd6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166111619092919063ffffffff16565b805190915015610d485780806020019051810190610ff49190611492565b610d48576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610456565b8060010154816002015411156110c2576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060010154816002015414156110d55750565b60038101544290811015611115576040517ff01f197500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482611127919061147b565b600184015484549192506111539161113f90846114af565b856002015461114e91906114ec565b611178565b600284015550600390910155565b6060611170848460008561118e565b949350505050565b60008183106111875781610948565b5090919050565b606082471015611220576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610456565b843b611288576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610456565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516112b19190611530565b60006040518083038185875af1925050503d80600081146112ee576040519150601f19603f3d011682016040523d82523d6000602084013e6112f3565b606091505b509150915061130382828661130e565b979650505050505050565b6060831561131d575081610948565b82511561132d5782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610456919061154c565b73ffffffffffffffffffffffffffffffffffffffff811681146108e357600080fd5b60006020828403121561139557600080fd5b813561094881611361565b80151581146108e357600080fd5b600080604083850312156113c157600080fd5b82356113cc81611361565b915060208301356113dc816113a0565b809150509250929050565b6000806000606084860312156113fc57600080fd5b83359250602084013591506040840135611415816113a0565b809150509250925092565b6000806040838503121561143357600080fd5b823561143e81611361565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561148d5761148d61144c565b500390565b6000602082840312156114a457600080fd5b8151610948816113a0565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156114e7576114e761144c565b500290565b600082198211156114ff576114ff61144c565b500190565b60005b8381101561151f578181015183820152602001611507565b83811115610e355750506000910152565b60008251611542818460208701611504565b9190910192915050565b602081526000825180602084015261156b816040850160208701611504565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fea164736f6c634300080c000a",
}

var NativeTokenPoolABI = NativeTokenPoolMetaData.ABI

var NativeTokenPoolBin = NativeTokenPoolMetaData.Bin

func DeployNativeTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, lockBucketRate *big.Int, lockBucketCapacity *big.Int, releaseBucketRate *big.Int, releaseBucketCapacity *big.Int) (common.Address, *types.Transaction, *NativeTokenPool, error) {
	parsed, err := NativeTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenPoolBin), backend, token, lockBucketRate, lockBucketCapacity, releaseBucketRate, releaseBucketCapacity)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenPool{NativeTokenPoolCaller: NativeTokenPoolCaller{contract: contract}, NativeTokenPoolTransactor: NativeTokenPoolTransactor{contract: contract}, NativeTokenPoolFilterer: NativeTokenPoolFilterer{contract: contract}}, nil
}

type NativeTokenPool struct {
	address common.Address
	abi     abi.ABI
	NativeTokenPoolCaller
	NativeTokenPoolTransactor
	NativeTokenPoolFilterer
}

type NativeTokenPoolCaller struct {
	contract *bind.BoundContract
}

type NativeTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type NativeTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type NativeTokenPoolSession struct {
	Contract     *NativeTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type NativeTokenPoolCallerSession struct {
	Contract *NativeTokenPoolCaller
	CallOpts bind.CallOpts
}

type NativeTokenPoolTransactorSession struct {
	Contract     *NativeTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type NativeTokenPoolRaw struct {
	Contract *NativeTokenPool
}

type NativeTokenPoolCallerRaw struct {
	Contract *NativeTokenPoolCaller
}

type NativeTokenPoolTransactorRaw struct {
	Contract *NativeTokenPoolTransactor
}

func NewNativeTokenPool(address common.Address, backend bind.ContractBackend) (*NativeTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(NativeTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindNativeTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPool{address: address, abi: abi, NativeTokenPoolCaller: NativeTokenPoolCaller{contract: contract}, NativeTokenPoolTransactor: NativeTokenPoolTransactor{contract: contract}, NativeTokenPoolFilterer: NativeTokenPoolFilterer{contract: contract}}, nil
}

func NewNativeTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenPoolCaller, error) {
	contract, err := bindNativeTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolCaller{contract: contract}, nil
}

func NewNativeTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenPoolTransactor, error) {
	contract, err := bindNativeTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolTransactor{contract: contract}, nil
}

func NewNativeTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenPoolFilterer, error) {
	contract, err := bindNativeTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolFilterer{contract: contract}, nil
}

func bindNativeTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NativeTokenPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_NativeTokenPool *NativeTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenPool.Contract.NativeTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_NativeTokenPool *NativeTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.NativeTokenPoolTransactor.contract.Transfer(opts)
}

func (_NativeTokenPool *NativeTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.NativeTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_NativeTokenPool *NativeTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_NativeTokenPool *NativeTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.contract.Transfer(opts)
}

func (_NativeTokenPool *NativeTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_NativeTokenPool *NativeTokenPoolCaller) GetLockOrBurnBucket(opts *bind.CallOpts) (TokenLimitsTokenBucket, error) {
	var out []interface{}
	err := _NativeTokenPool.contract.Call(opts, &out, "getLockOrBurnBucket")

	if err != nil {
		return *new(TokenLimitsTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenLimitsTokenBucket)).(*TokenLimitsTokenBucket)

	return out0, err

}

func (_NativeTokenPool *NativeTokenPoolSession) GetLockOrBurnBucket() (TokenLimitsTokenBucket, error) {
	return _NativeTokenPool.Contract.GetLockOrBurnBucket(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolCallerSession) GetLockOrBurnBucket() (TokenLimitsTokenBucket, error) {
	return _NativeTokenPool.Contract.GetLockOrBurnBucket(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolCaller) GetReleaseOrMintBucket(opts *bind.CallOpts) (TokenLimitsTokenBucket, error) {
	var out []interface{}
	err := _NativeTokenPool.contract.Call(opts, &out, "getReleaseOrMintBucket")

	if err != nil {
		return *new(TokenLimitsTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenLimitsTokenBucket)).(*TokenLimitsTokenBucket)

	return out0, err

}

func (_NativeTokenPool *NativeTokenPoolSession) GetReleaseOrMintBucket() (TokenLimitsTokenBucket, error) {
	return _NativeTokenPool.Contract.GetReleaseOrMintBucket(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolCallerSession) GetReleaseOrMintBucket() (TokenLimitsTokenBucket, error) {
	return _NativeTokenPool.Contract.GetReleaseOrMintBucket(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_NativeTokenPool *NativeTokenPoolSession) GetToken() (common.Address, error) {
	return _NativeTokenPool.Contract.GetToken(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _NativeTokenPool.Contract.GetToken(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolCaller) IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenPool.contract.Call(opts, &out, "isOffRamp", offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_NativeTokenPool *NativeTokenPoolSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _NativeTokenPool.Contract.IsOffRamp(&_NativeTokenPool.CallOpts, offRamp)
}

func (_NativeTokenPool *NativeTokenPoolCallerSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _NativeTokenPool.Contract.IsOffRamp(&_NativeTokenPool.CallOpts, offRamp)
}

func (_NativeTokenPool *NativeTokenPoolCaller) IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenPool.contract.Call(opts, &out, "isOnRamp", onRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_NativeTokenPool *NativeTokenPoolSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _NativeTokenPool.Contract.IsOnRamp(&_NativeTokenPool.CallOpts, onRamp)
}

func (_NativeTokenPool *NativeTokenPoolCallerSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _NativeTokenPool.Contract.IsOnRamp(&_NativeTokenPool.CallOpts, onRamp)
}

func (_NativeTokenPool *NativeTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_NativeTokenPool *NativeTokenPoolSession) Owner() (common.Address, error) {
	return _NativeTokenPool.Contract.Owner(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolCallerSession) Owner() (common.Address, error) {
	return _NativeTokenPool.Contract.Owner(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_NativeTokenPool *NativeTokenPoolSession) Paused() (bool, error) {
	return _NativeTokenPool.Contract.Paused(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolCallerSession) Paused() (bool, error) {
	return _NativeTokenPool.Contract.Paused(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_NativeTokenPool *NativeTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _NativeTokenPool.Contract.AcceptOwnership(&_NativeTokenPool.TransactOpts)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _NativeTokenPool.Contract.AcceptOwnership(&_NativeTokenPool.TransactOpts)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, depositor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "lockOrBurn", depositor, amount)
}

func (_NativeTokenPool *NativeTokenPoolSession) LockOrBurn(depositor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.LockOrBurn(&_NativeTokenPool.TransactOpts, depositor, amount)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) LockOrBurn(depositor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.LockOrBurn(&_NativeTokenPool.TransactOpts, depositor, amount)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "pause")
}

func (_NativeTokenPool *NativeTokenPoolSession) Pause() (*types.Transaction, error) {
	return _NativeTokenPool.Contract.Pause(&_NativeTokenPool.TransactOpts)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) Pause() (*types.Transaction, error) {
	return _NativeTokenPool.Contract.Pause(&_NativeTokenPool.TransactOpts)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "releaseOrMint", recipient, amount)
}

func (_NativeTokenPool *NativeTokenPoolSession) ReleaseOrMint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.ReleaseOrMint(&_NativeTokenPool.TransactOpts, recipient, amount)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) ReleaseOrMint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.ReleaseOrMint(&_NativeTokenPool.TransactOpts, recipient, amount)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) SetLockOrBurnBucket(opts *bind.TransactOpts, rate *big.Int, capacity *big.Int, full bool) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "setLockOrBurnBucket", rate, capacity, full)
}

func (_NativeTokenPool *NativeTokenPoolSession) SetLockOrBurnBucket(rate *big.Int, capacity *big.Int, full bool) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.SetLockOrBurnBucket(&_NativeTokenPool.TransactOpts, rate, capacity, full)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) SetLockOrBurnBucket(rate *big.Int, capacity *big.Int, full bool) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.SetLockOrBurnBucket(&_NativeTokenPool.TransactOpts, rate, capacity, full)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) SetOffRamp(opts *bind.TransactOpts, offRamp common.Address, permission bool) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "setOffRamp", offRamp, permission)
}

func (_NativeTokenPool *NativeTokenPoolSession) SetOffRamp(offRamp common.Address, permission bool) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.SetOffRamp(&_NativeTokenPool.TransactOpts, offRamp, permission)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) SetOffRamp(offRamp common.Address, permission bool) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.SetOffRamp(&_NativeTokenPool.TransactOpts, offRamp, permission)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) SetOnRamp(opts *bind.TransactOpts, onRamp common.Address, permission bool) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "setOnRamp", onRamp, permission)
}

func (_NativeTokenPool *NativeTokenPoolSession) SetOnRamp(onRamp common.Address, permission bool) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.SetOnRamp(&_NativeTokenPool.TransactOpts, onRamp, permission)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) SetOnRamp(onRamp common.Address, permission bool) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.SetOnRamp(&_NativeTokenPool.TransactOpts, onRamp, permission)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) SetReleaseOrMintBucket(opts *bind.TransactOpts, rate *big.Int, capacity *big.Int, full bool) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "setReleaseOrMintBucket", rate, capacity, full)
}

func (_NativeTokenPool *NativeTokenPoolSession) SetReleaseOrMintBucket(rate *big.Int, capacity *big.Int, full bool) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.SetReleaseOrMintBucket(&_NativeTokenPool.TransactOpts, rate, capacity, full)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) SetReleaseOrMintBucket(rate *big.Int, capacity *big.Int, full bool) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.SetReleaseOrMintBucket(&_NativeTokenPool.TransactOpts, rate, capacity, full)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_NativeTokenPool *NativeTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.TransferOwnership(&_NativeTokenPool.TransactOpts, to)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.TransferOwnership(&_NativeTokenPool.TransactOpts, to)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "unpause")
}

func (_NativeTokenPool *NativeTokenPoolSession) Unpause() (*types.Transaction, error) {
	return _NativeTokenPool.Contract.Unpause(&_NativeTokenPool.TransactOpts)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _NativeTokenPool.Contract.Unpause(&_NativeTokenPool.TransactOpts)
}

type NativeTokenPoolBurnedIterator struct {
	Event *NativeTokenPoolBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolBurned)
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
		it.Event = new(NativeTokenPoolBurned)
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

func (it *NativeTokenPoolBurnedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolBurned struct {
	Sender    common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address, depositor []common.Address) (*NativeTokenPoolBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "Burned", senderRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolBurnedIterator{contract: _NativeTokenPool.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolBurned, sender []common.Address, depositor []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "Burned", senderRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolBurned)
				if err := _NativeTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseBurned(log types.Log) (*NativeTokenPoolBurned, error) {
	event := new(NativeTokenPoolBurned)
	if err := _NativeTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolLockedIterator struct {
	Event *NativeTokenPoolLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolLocked)
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
		it.Event = new(NativeTokenPoolLocked)
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

func (it *NativeTokenPoolLockedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolLocked struct {
	Sender    common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address, depositor []common.Address) (*NativeTokenPoolLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "Locked", senderRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolLockedIterator{contract: _NativeTokenPool.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolLocked, sender []common.Address, depositor []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "Locked", senderRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolLocked)
				if err := _NativeTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseLocked(log types.Log) (*NativeTokenPoolLocked, error) {
	event := new(NativeTokenPoolLocked)
	if err := _NativeTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolMintedIterator struct {
	Event *NativeTokenPoolMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolMinted)
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
		it.Event = new(NativeTokenPoolMinted)
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

func (it *NativeTokenPoolMintedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*NativeTokenPoolMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolMintedIterator{contract: _NativeTokenPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolMinted)
				if err := _NativeTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseMinted(log types.Log) (*NativeTokenPoolMinted, error) {
	event := new(NativeTokenPoolMinted)
	if err := _NativeTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolNewLockBurnBucketConstructedIterator struct {
	Event *NativeTokenPoolNewLockBurnBucketConstructed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolNewLockBurnBucketConstructedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolNewLockBurnBucketConstructed)
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
		it.Event = new(NativeTokenPoolNewLockBurnBucketConstructed)
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

func (it *NativeTokenPoolNewLockBurnBucketConstructedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolNewLockBurnBucketConstructedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolNewLockBurnBucketConstructed struct {
	Rate     *big.Int
	Capacity *big.Int
	Full     bool
	Raw      types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterNewLockBurnBucketConstructed(opts *bind.FilterOpts) (*NativeTokenPoolNewLockBurnBucketConstructedIterator, error) {

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "NewLockBurnBucketConstructed")
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolNewLockBurnBucketConstructedIterator{contract: _NativeTokenPool.contract, event: "NewLockBurnBucketConstructed", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchNewLockBurnBucketConstructed(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolNewLockBurnBucketConstructed) (event.Subscription, error) {

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "NewLockBurnBucketConstructed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolNewLockBurnBucketConstructed)
				if err := _NativeTokenPool.contract.UnpackLog(event, "NewLockBurnBucketConstructed", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseNewLockBurnBucketConstructed(log types.Log) (*NativeTokenPoolNewLockBurnBucketConstructed, error) {
	event := new(NativeTokenPoolNewLockBurnBucketConstructed)
	if err := _NativeTokenPool.contract.UnpackLog(event, "NewLockBurnBucketConstructed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolNewReleaseMintBucketConstructedIterator struct {
	Event *NativeTokenPoolNewReleaseMintBucketConstructed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolNewReleaseMintBucketConstructedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolNewReleaseMintBucketConstructed)
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
		it.Event = new(NativeTokenPoolNewReleaseMintBucketConstructed)
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

func (it *NativeTokenPoolNewReleaseMintBucketConstructedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolNewReleaseMintBucketConstructedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolNewReleaseMintBucketConstructed struct {
	Rate     *big.Int
	Capacity *big.Int
	Full     bool
	Raw      types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterNewReleaseMintBucketConstructed(opts *bind.FilterOpts) (*NativeTokenPoolNewReleaseMintBucketConstructedIterator, error) {

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "NewReleaseMintBucketConstructed")
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolNewReleaseMintBucketConstructedIterator{contract: _NativeTokenPool.contract, event: "NewReleaseMintBucketConstructed", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchNewReleaseMintBucketConstructed(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolNewReleaseMintBucketConstructed) (event.Subscription, error) {

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "NewReleaseMintBucketConstructed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolNewReleaseMintBucketConstructed)
				if err := _NativeTokenPool.contract.UnpackLog(event, "NewReleaseMintBucketConstructed", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseNewReleaseMintBucketConstructed(log types.Log) (*NativeTokenPoolNewReleaseMintBucketConstructed, error) {
	event := new(NativeTokenPoolNewReleaseMintBucketConstructed)
	if err := _NativeTokenPool.contract.UnpackLog(event, "NewReleaseMintBucketConstructed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolOwnershipTransferRequestedIterator struct {
	Event *NativeTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolOwnershipTransferRequested)
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
		it.Event = new(NativeTokenPoolOwnershipTransferRequested)
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

func (it *NativeTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolOwnershipTransferRequestedIterator{contract: _NativeTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolOwnershipTransferRequested)
				if err := _NativeTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*NativeTokenPoolOwnershipTransferRequested, error) {
	event := new(NativeTokenPoolOwnershipTransferRequested)
	if err := _NativeTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolOwnershipTransferredIterator struct {
	Event *NativeTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolOwnershipTransferred)
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
		it.Event = new(NativeTokenPoolOwnershipTransferred)
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

func (it *NativeTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolOwnershipTransferredIterator{contract: _NativeTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolOwnershipTransferred)
				if err := _NativeTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*NativeTokenPoolOwnershipTransferred, error) {
	event := new(NativeTokenPoolOwnershipTransferred)
	if err := _NativeTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolPausedIterator struct {
	Event *NativeTokenPoolPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolPaused)
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
		it.Event = new(NativeTokenPoolPaused)
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

func (it *NativeTokenPoolPausedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*NativeTokenPoolPausedIterator, error) {

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolPausedIterator{contract: _NativeTokenPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolPaused) (event.Subscription, error) {

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolPaused)
				if err := _NativeTokenPool.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParsePaused(log types.Log) (*NativeTokenPoolPaused, error) {
	event := new(NativeTokenPoolPaused)
	if err := _NativeTokenPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolReleasedIterator struct {
	Event *NativeTokenPoolReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolReleased)
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
		it.Event = new(NativeTokenPoolReleased)
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

func (it *NativeTokenPoolReleasedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*NativeTokenPoolReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolReleasedIterator{contract: _NativeTokenPool.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolReleased)
				if err := _NativeTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseReleased(log types.Log) (*NativeTokenPoolReleased, error) {
	event := new(NativeTokenPoolReleased)
	if err := _NativeTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolUnpausedIterator struct {
	Event *NativeTokenPoolUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolUnpaused)
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
		it.Event = new(NativeTokenPoolUnpaused)
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

func (it *NativeTokenPoolUnpausedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NativeTokenPoolUnpausedIterator, error) {

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolUnpausedIterator{contract: _NativeTokenPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolUnpaused)
				if err := _NativeTokenPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseUnpaused(log types.Log) (*NativeTokenPoolUnpaused, error) {
	event := new(NativeTokenPoolUnpaused)
	if err := _NativeTokenPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_NativeTokenPool *NativeTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _NativeTokenPool.abi.Events["Burned"].ID:
		return _NativeTokenPool.ParseBurned(log)
	case _NativeTokenPool.abi.Events["Locked"].ID:
		return _NativeTokenPool.ParseLocked(log)
	case _NativeTokenPool.abi.Events["Minted"].ID:
		return _NativeTokenPool.ParseMinted(log)
	case _NativeTokenPool.abi.Events["NewLockBurnBucketConstructed"].ID:
		return _NativeTokenPool.ParseNewLockBurnBucketConstructed(log)
	case _NativeTokenPool.abi.Events["NewReleaseMintBucketConstructed"].ID:
		return _NativeTokenPool.ParseNewReleaseMintBucketConstructed(log)
	case _NativeTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _NativeTokenPool.ParseOwnershipTransferRequested(log)
	case _NativeTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _NativeTokenPool.ParseOwnershipTransferred(log)
	case _NativeTokenPool.abi.Events["Paused"].ID:
		return _NativeTokenPool.ParsePaused(log)
	case _NativeTokenPool.abi.Events["Released"].ID:
		return _NativeTokenPool.ParseReleased(log)
	case _NativeTokenPool.abi.Events["Unpaused"].ID:
		return _NativeTokenPool.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (NativeTokenPoolBurned) Topic() common.Hash {
	return common.HexToHash("0x6ab368f832c266c8eb942b84fbcaa20aedc24a699d2a05fae2568028733b1d09")
}

func (NativeTokenPoolLocked) Topic() common.Hash {
	return common.HexToHash("0x989eaa915cbb416ea3d6f9a63b1a3de51770c7674b11fe21ecdf76b4e1d13910")
}

func (NativeTokenPoolMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (NativeTokenPoolNewLockBurnBucketConstructed) Topic() common.Hash {
	return common.HexToHash("0x4bab560bebbbda366298430a0e5cd55e85ea36dc92834dc281b93dc21b9683a2")
}

func (NativeTokenPoolNewReleaseMintBucketConstructed) Topic() common.Hash {
	return common.HexToHash("0x8586e57eb58352bf12165ebd6104b5991cb1c22a183a9965a2961b7b462bc2df")
}

func (NativeTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (NativeTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (NativeTokenPoolPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (NativeTokenPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (NativeTokenPoolUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_NativeTokenPool *NativeTokenPool) Address() common.Address {
	return _NativeTokenPool.address
}

type NativeTokenPoolInterface interface {
	GetLockOrBurnBucket(opts *bind.CallOpts) (TokenLimitsTokenBucket, error)

	GetReleaseOrMintBucket(opts *bind.CallOpts) (TokenLimitsTokenBucket, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, depositor common.Address, amount *big.Int) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	SetLockOrBurnBucket(opts *bind.TransactOpts, rate *big.Int, capacity *big.Int, full bool) (*types.Transaction, error)

	SetOffRamp(opts *bind.TransactOpts, offRamp common.Address, permission bool) (*types.Transaction, error)

	SetOnRamp(opts *bind.TransactOpts, onRamp common.Address, permission bool) (*types.Transaction, error)

	SetReleaseOrMintBucket(opts *bind.TransactOpts, rate *big.Int, capacity *big.Int, full bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address, depositor []common.Address) (*NativeTokenPoolBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolBurned, sender []common.Address, depositor []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*NativeTokenPoolBurned, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address, depositor []common.Address) (*NativeTokenPoolLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolLocked, sender []common.Address, depositor []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*NativeTokenPoolLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*NativeTokenPoolMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*NativeTokenPoolMinted, error)

	FilterNewLockBurnBucketConstructed(opts *bind.FilterOpts) (*NativeTokenPoolNewLockBurnBucketConstructedIterator, error)

	WatchNewLockBurnBucketConstructed(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolNewLockBurnBucketConstructed) (event.Subscription, error)

	ParseNewLockBurnBucketConstructed(log types.Log) (*NativeTokenPoolNewLockBurnBucketConstructed, error)

	FilterNewReleaseMintBucketConstructed(opts *bind.FilterOpts) (*NativeTokenPoolNewReleaseMintBucketConstructedIterator, error)

	WatchNewReleaseMintBucketConstructed(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolNewReleaseMintBucketConstructed) (event.Subscription, error)

	ParseNewReleaseMintBucketConstructed(log types.Log) (*NativeTokenPoolNewReleaseMintBucketConstructed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*NativeTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*NativeTokenPoolOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*NativeTokenPoolPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*NativeTokenPoolPaused, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*NativeTokenPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*NativeTokenPoolReleased, error)

	FilterUnpaused(opts *bind.FilterOpts) (*NativeTokenPoolUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*NativeTokenPoolUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
