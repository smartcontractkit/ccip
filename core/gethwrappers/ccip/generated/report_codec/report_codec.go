// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package report_codec

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

type EVM2EVMMultiOffRampCommitReport struct {
	PriceUpdates InternalPriceUpdates
	MerkleRoots  []EVM2EVMMultiOffRampMerkleRoot
}

type EVM2EVMMultiOffRampInterval struct {
	Min uint64
	Max uint64
}

type EVM2EVMMultiOffRampMerkleRoot struct {
	SourceChainSelector uint64
	Interval            EVM2EVMMultiOffRampInterval
	MerkleRoot          [32]byte
}

type InternalAny2EVMRampMessage struct {
	Header       InternalRampMessageHeader
	Sender       []byte
	Data         []byte
	Receiver     common.Address
	GasLimit     *big.Int
	TokenAmounts []InternalRampTokenAmount
}

type InternalExecutionReportSingleChain struct {
	SourceChainSelector uint64
	Messages            []InternalAny2EVMRampMessage
	OffchainTokenData   [][][]byte
	Proofs              [][32]byte
	ProofFlagBits       *big.Int
}

type InternalGasPriceUpdate struct {
	DestChainSelector uint64
	UsdPerUnitGas     *big.Int
}

type InternalPriceUpdates struct {
	TokenPriceUpdates []InternalTokenPriceUpdate
	GasPriceUpdates   []InternalGasPriceUpdate
}

type InternalRampMessageHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	Nonce               uint64
}

type InternalRampTokenAmount struct {
	SourcePoolAddress []byte
	DestTokenAddress  []byte
	ExtraData         []byte
	Amount            *big.Int
	DestGasAmount     uint32
}

type InternalTokenPriceUpdate struct {
	SourceToken common.Address
	UsdPerToken *big.Int
}

var ReportCodecMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint224\",\"name\":\"usdPerToken\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint224\",\"name\":\"usdPerUnitGas\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.GasPriceUpdate[]\",\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.PriceUpdates\",\"name\":\"priceUpdates\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.Interval\",\"name\":\"interval\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.MerkleRoot[]\",\"name\":\"merkleRoots\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.CommitReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"CommitReportDecoded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structInternal.RampMessageHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"destGasAmount\",\"type\":\"uint32\"}],\"internalType\":\"structInternal.RampTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.Any2EVMRampMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structInternal.ExecutionReportSingleChain[]\",\"name\":\"report\",\"type\":\"tuple[]\"}],\"name\":\"ExecuteReportDecoded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"decodeCommitReport\",\"outputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint224\",\"name\":\"usdPerToken\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint224\",\"name\":\"usdPerUnitGas\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.GasPriceUpdate[]\",\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.PriceUpdates\",\"name\":\"priceUpdates\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.Interval\",\"name\":\"interval\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.MerkleRoot[]\",\"name\":\"merkleRoots\",\"type\":\"tuple[]\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.CommitReport\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"decodeExecuteReport\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structInternal.RampMessageHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"destGasAmount\",\"type\":\"uint32\"}],\"internalType\":\"structInternal.RampTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.Any2EVMRampMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReportSingleChain[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611261806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636fb349561461003b578063f816ec6014610064575b600080fd5b61004e61004936600461022c565b610084565b60405161005b91906104e1565b60405180910390f35b61007761007236600461022c565b6100a0565b60405161005b919061079a565b60608180602001905181019061009a9190610dd5565b92915050565b604080516080810182526060918101828152828201839052815260208101919091528180602001905181019061009a91906110eb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff81118282101715610128576101286100d6565b60405290565b60405160c0810167ffffffffffffffff81118282101715610128576101286100d6565b6040805190810167ffffffffffffffff81118282101715610128576101286100d6565b6040516060810167ffffffffffffffff81118282101715610128576101286100d6565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156101de576101de6100d6565b604052919050565b600067ffffffffffffffff821115610200576102006100d6565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60006020828403121561023e57600080fd5b813567ffffffffffffffff81111561025557600080fd5b8201601f8101841361026657600080fd5b8035610279610274826101e6565b610197565b81815285602083850101111561028e57600080fd5b81602084016020830137600091810160200191909152949350505050565b60005b838110156102c75781810151838201526020016102af565b50506000910152565b600081518084526102e88160208601602086016102ac565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208086019550808260051b84010181860160005b848110156103de577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868403018952815160a0815181865261037b828701826102d0565b915050858201518582038787015261039382826102d0565b915050604080830151868303828801526103ad83826102d0565b6060858101519089015260809485015163ffffffff1694909701939093525050509783019790830190600101610337565b5090979650505050505050565b6000828251808552602080860195506005818360051b8501018287016000805b86811015610496577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe088850381018c5283518051808752908801908887019080891b88018a01865b8281101561047f57858a830301845261046d8286516102d0565b948c0194938c01939150600101610453565b509e8a019e9750505093870193505060010161040b565b50919998505050505050505050565b60008151808452602080850194506020840160005b838110156104d6578151875295820195908201906001016104ba565b509495945050505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156106c9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0888603018452815160a0860167ffffffffffffffff8083511688528883015160a08a8a015282815180855260c08b01915060c08160051b8c010194508b8301925060005b81811015610672577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408c87030183528351805180518852868f820151168f890152866040820151166040890152866060820151166060890152866080820151166080890152508d81015161014060a08901526105f56101408901826102d0565b9050604082015188820360c08a015261060e82826102d0565b915050606082015161063860e08a018273ffffffffffffffffffffffffffffffffffffffff169052565b50608082015161010089015260a0820151915087810361012089015261065e818361031a565b97505050928c0192918c0191600101610575565b50505050506040820151878203604089015261068e82826103eb565b915050606082015187820360608901526106a882826104a5565b60809384015198909301979097525094509285019290850190600101610508565b5092979650505050505050565b60008151808452602080850194506020840160005b838110156104d6578151805167ffffffffffffffff1688528301517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1683880152604090960195908201906001016106eb565b600081518084526020808501945080840160005b838110156104d6578151805167ffffffffffffffff90811689528482015180518216868b0152850151166040808a0191909152015160608801526080909601959082019060010161074e565b6000602080835283516040808386015260a0850182516040606088015281815180845260c0890191508683019350600092505b8083101561082a578351805173ffffffffffffffffffffffffffffffffffffffff1683528701517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16878301529286019260019290920191908401906107cd565b50938501518785037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa00160808901529361086481866106d6565b9450505050508185015191507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08482030160408501526108a4818361073a565b95945050505050565b600067ffffffffffffffff8211156108c7576108c76100d6565b5060051b60200190565b805167ffffffffffffffff811681146108e957600080fd5b919050565b600060a0828403121561090057600080fd5b610908610105565b90508151815261091a602083016108d1565b602082015261092b604083016108d1565b604082015261093c606083016108d1565b606082015261094d608083016108d1565b608082015292915050565b600082601f83011261096957600080fd5b8151610977610274826101e6565b81815284602083860101111561098c57600080fd5b61099d8260208301602087016102ac565b949350505050565b805173ffffffffffffffffffffffffffffffffffffffff811681146108e957600080fd5b805163ffffffff811681146108e957600080fd5b600082601f8301126109ee57600080fd5b815160206109fe610274836108ad565b82815260059290921b84018101918181019086841115610a1d57600080fd5b8286015b84811015610b2357805167ffffffffffffffff80821115610a425760008081fd5b818901915060a0807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848d03011215610a7b5760008081fd5b610a83610105565b8784015183811115610a955760008081fd5b610aa38d8a83880101610958565b82525060408085015184811115610aba5760008081fd5b610ac88e8b83890101610958565b8a8401525060608086015185811115610ae15760008081fd5b610aef8f8c838a0101610958565b83850152506080945084860151818401525050610b0d8285016109c9565b9281019290925250845250918301918301610a21565b509695505050505050565b600082601f830112610b3f57600080fd5b81516020610b4f610274836108ad565b82815260059290921b84018101918181019086841115610b6e57600080fd5b8286015b84811015610b2357805167ffffffffffffffff80821115610b935760008081fd5b8189019150610140807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848d03011215610bcd5760008081fd5b610bd561012e565b610be18c8986016108ee565b815260c084015183811115610bf65760008081fd5b610c048d8a83880101610958565b898301525060e084015183811115610c1c5760008081fd5b610c2a8d8a83880101610958565b604083015250610c3d61010085016109a5565b60608201526101208401516080820152908301519082821115610c605760008081fd5b610c6e8c89848701016109dd565b60a08201528652505050918301918301610b72565b600082601f830112610c9457600080fd5b81516020610ca4610274836108ad565b82815260059290921b84018101918181019086841115610cc357600080fd5b8286015b84811015610b2357805167ffffffffffffffff80821115610ce757600080fd5b818901915089603f830112610cfb57600080fd5b85820151610d0b610274826108ad565b81815260059190911b830160400190878101908c831115610d2b57600080fd5b604085015b83811015610d6457805185811115610d4757600080fd5b610d568f6040838a0101610958565b845250918901918901610d30565b50875250505092840192508301610cc7565b600082601f830112610d8757600080fd5b81516020610d97610274836108ad565b8083825260208201915060208460051b870101935086841115610db957600080fd5b602086015b84811015610b235780518352918301918301610dbe565b60006020808385031215610de857600080fd5b825167ffffffffffffffff80821115610e0057600080fd5b818501915085601f830112610e1457600080fd5b8151610e22610274826108ad565b81815260059190911b83018401908481019088831115610e4157600080fd5b8585015b83811015610f3b57805185811115610e5c57600080fd5b860160a0818c037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0011215610e915760008081fd5b610e99610105565b610ea48983016108d1565b815260408083015188811115610eba5760008081fd5b610ec88e8c83870101610b2e565b8b8401525060608084015189811115610ee15760008081fd5b610eef8f8d83880101610c83565b8385015250608091508184015189811115610f0a5760008081fd5b610f188f8d83880101610d76565b918401919091525060a09290920151918101919091528352918601918601610e45565b5098975050505050505050565b80517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811681146108e957600080fd5b600082601f830112610f8557600080fd5b81516020610f95610274836108ad565b82815260069290921b84018101918181019086841115610fb457600080fd5b8286015b84811015610b235760408189031215610fd15760008081fd5b610fd9610151565b610fe2826108d1565b8152610fef858301610f48565b81860152835291830191604001610fb8565b600082601f83011261101257600080fd5b81516020611022610274836108ad565b82815260079290921b8401810191818101908684111561104157600080fd5b8286015b84811015610b2357808803608081121561105f5760008081fd5b611067610174565b611070836108d1565b81526040807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0840112156110a45760008081fd5b6110ac610151565b92506110b98785016108d1565b83526110c68185016108d1565b8388015281870192909252606083015191810191909152835291830191608001611045565b600060208083850312156110fe57600080fd5b825167ffffffffffffffff8082111561111657600080fd5b8185019150604080838803121561112c57600080fd5b611134610151565b83518381111561114357600080fd5b84016040818a03121561115557600080fd5b61115d610151565b81518581111561116c57600080fd5b8201601f81018b1361117d57600080fd5b805161118b610274826108ad565b81815260069190911b8201890190898101908d8311156111aa57600080fd5b928a01925b828410156111f85787848f0312156111c75760008081fd5b6111cf610151565b6111d8856109a5565b81526111e58c8601610f48565b818d0152825292870192908a01906111af565b84525050508187015193508484111561121057600080fd5b61121c8a858401610f74565b818801528252508385015191508282111561123657600080fd5b61124288838601611001565b8582015280955050505050509291505056fea164736f6c6343000818000a",
}

var ReportCodecABI = ReportCodecMetaData.ABI

var ReportCodecBin = ReportCodecMetaData.Bin

func DeployReportCodec(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReportCodec, error) {
	parsed, err := ReportCodecMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReportCodecBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReportCodec{address: address, abi: *parsed, ReportCodecCaller: ReportCodecCaller{contract: contract}, ReportCodecTransactor: ReportCodecTransactor{contract: contract}, ReportCodecFilterer: ReportCodecFilterer{contract: contract}}, nil
}

type ReportCodec struct {
	address common.Address
	abi     abi.ABI
	ReportCodecCaller
	ReportCodecTransactor
	ReportCodecFilterer
}

type ReportCodecCaller struct {
	contract *bind.BoundContract
}

type ReportCodecTransactor struct {
	contract *bind.BoundContract
}

type ReportCodecFilterer struct {
	contract *bind.BoundContract
}

type ReportCodecSession struct {
	Contract     *ReportCodec
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ReportCodecCallerSession struct {
	Contract *ReportCodecCaller
	CallOpts bind.CallOpts
}

type ReportCodecTransactorSession struct {
	Contract     *ReportCodecTransactor
	TransactOpts bind.TransactOpts
}

type ReportCodecRaw struct {
	Contract *ReportCodec
}

type ReportCodecCallerRaw struct {
	Contract *ReportCodecCaller
}

type ReportCodecTransactorRaw struct {
	Contract *ReportCodecTransactor
}

func NewReportCodec(address common.Address, backend bind.ContractBackend) (*ReportCodec, error) {
	abi, err := abi.JSON(strings.NewReader(ReportCodecABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindReportCodec(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReportCodec{address: address, abi: abi, ReportCodecCaller: ReportCodecCaller{contract: contract}, ReportCodecTransactor: ReportCodecTransactor{contract: contract}, ReportCodecFilterer: ReportCodecFilterer{contract: contract}}, nil
}

func NewReportCodecCaller(address common.Address, caller bind.ContractCaller) (*ReportCodecCaller, error) {
	contract, err := bindReportCodec(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReportCodecCaller{contract: contract}, nil
}

func NewReportCodecTransactor(address common.Address, transactor bind.ContractTransactor) (*ReportCodecTransactor, error) {
	contract, err := bindReportCodec(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReportCodecTransactor{contract: contract}, nil
}

func NewReportCodecFilterer(address common.Address, filterer bind.ContractFilterer) (*ReportCodecFilterer, error) {
	contract, err := bindReportCodec(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReportCodecFilterer{contract: contract}, nil
}

func bindReportCodec(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReportCodecMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ReportCodec *ReportCodecRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReportCodec.Contract.ReportCodecCaller.contract.Call(opts, result, method, params...)
}

func (_ReportCodec *ReportCodecRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReportCodec.Contract.ReportCodecTransactor.contract.Transfer(opts)
}

func (_ReportCodec *ReportCodecRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReportCodec.Contract.ReportCodecTransactor.contract.Transact(opts, method, params...)
}

func (_ReportCodec *ReportCodecCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReportCodec.Contract.contract.Call(opts, result, method, params...)
}

func (_ReportCodec *ReportCodecTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReportCodec.Contract.contract.Transfer(opts)
}

func (_ReportCodec *ReportCodecTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReportCodec.Contract.contract.Transact(opts, method, params...)
}

func (_ReportCodec *ReportCodecCaller) DecodeCommitReport(opts *bind.CallOpts, report []byte) (EVM2EVMMultiOffRampCommitReport, error) {
	var out []interface{}
	err := _ReportCodec.contract.Call(opts, &out, "decodeCommitReport", report)

	if err != nil {
		return *new(EVM2EVMMultiOffRampCommitReport), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMMultiOffRampCommitReport)).(*EVM2EVMMultiOffRampCommitReport)

	return out0, err

}

func (_ReportCodec *ReportCodecSession) DecodeCommitReport(report []byte) (EVM2EVMMultiOffRampCommitReport, error) {
	return _ReportCodec.Contract.DecodeCommitReport(&_ReportCodec.CallOpts, report)
}

func (_ReportCodec *ReportCodecCallerSession) DecodeCommitReport(report []byte) (EVM2EVMMultiOffRampCommitReport, error) {
	return _ReportCodec.Contract.DecodeCommitReport(&_ReportCodec.CallOpts, report)
}

func (_ReportCodec *ReportCodecCaller) DecodeExecuteReport(opts *bind.CallOpts, report []byte) ([]InternalExecutionReportSingleChain, error) {
	var out []interface{}
	err := _ReportCodec.contract.Call(opts, &out, "decodeExecuteReport", report)

	if err != nil {
		return *new([]InternalExecutionReportSingleChain), err
	}

	out0 := *abi.ConvertType(out[0], new([]InternalExecutionReportSingleChain)).(*[]InternalExecutionReportSingleChain)

	return out0, err

}

func (_ReportCodec *ReportCodecSession) DecodeExecuteReport(report []byte) ([]InternalExecutionReportSingleChain, error) {
	return _ReportCodec.Contract.DecodeExecuteReport(&_ReportCodec.CallOpts, report)
}

func (_ReportCodec *ReportCodecCallerSession) DecodeExecuteReport(report []byte) ([]InternalExecutionReportSingleChain, error) {
	return _ReportCodec.Contract.DecodeExecuteReport(&_ReportCodec.CallOpts, report)
}

type ReportCodecCommitReportDecodedIterator struct {
	Event *ReportCodecCommitReportDecoded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ReportCodecCommitReportDecodedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReportCodecCommitReportDecoded)
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
		it.Event = new(ReportCodecCommitReportDecoded)
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

func (it *ReportCodecCommitReportDecodedIterator) Error() error {
	return it.fail
}

func (it *ReportCodecCommitReportDecodedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ReportCodecCommitReportDecoded struct {
	Report EVM2EVMMultiOffRampCommitReport
	Raw    types.Log
}

func (_ReportCodec *ReportCodecFilterer) FilterCommitReportDecoded(opts *bind.FilterOpts) (*ReportCodecCommitReportDecodedIterator, error) {

	logs, sub, err := _ReportCodec.contract.FilterLogs(opts, "CommitReportDecoded")
	if err != nil {
		return nil, err
	}
	return &ReportCodecCommitReportDecodedIterator{contract: _ReportCodec.contract, event: "CommitReportDecoded", logs: logs, sub: sub}, nil
}

func (_ReportCodec *ReportCodecFilterer) WatchCommitReportDecoded(opts *bind.WatchOpts, sink chan<- *ReportCodecCommitReportDecoded) (event.Subscription, error) {

	logs, sub, err := _ReportCodec.contract.WatchLogs(opts, "CommitReportDecoded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ReportCodecCommitReportDecoded)
				if err := _ReportCodec.contract.UnpackLog(event, "CommitReportDecoded", log); err != nil {
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

func (_ReportCodec *ReportCodecFilterer) ParseCommitReportDecoded(log types.Log) (*ReportCodecCommitReportDecoded, error) {
	event := new(ReportCodecCommitReportDecoded)
	if err := _ReportCodec.contract.UnpackLog(event, "CommitReportDecoded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ReportCodecExecuteReportDecodedIterator struct {
	Event *ReportCodecExecuteReportDecoded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ReportCodecExecuteReportDecodedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReportCodecExecuteReportDecoded)
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
		it.Event = new(ReportCodecExecuteReportDecoded)
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

func (it *ReportCodecExecuteReportDecodedIterator) Error() error {
	return it.fail
}

func (it *ReportCodecExecuteReportDecodedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ReportCodecExecuteReportDecoded struct {
	Report []InternalExecutionReportSingleChain
	Raw    types.Log
}

func (_ReportCodec *ReportCodecFilterer) FilterExecuteReportDecoded(opts *bind.FilterOpts) (*ReportCodecExecuteReportDecodedIterator, error) {

	logs, sub, err := _ReportCodec.contract.FilterLogs(opts, "ExecuteReportDecoded")
	if err != nil {
		return nil, err
	}
	return &ReportCodecExecuteReportDecodedIterator{contract: _ReportCodec.contract, event: "ExecuteReportDecoded", logs: logs, sub: sub}, nil
}

func (_ReportCodec *ReportCodecFilterer) WatchExecuteReportDecoded(opts *bind.WatchOpts, sink chan<- *ReportCodecExecuteReportDecoded) (event.Subscription, error) {

	logs, sub, err := _ReportCodec.contract.WatchLogs(opts, "ExecuteReportDecoded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ReportCodecExecuteReportDecoded)
				if err := _ReportCodec.contract.UnpackLog(event, "ExecuteReportDecoded", log); err != nil {
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

func (_ReportCodec *ReportCodecFilterer) ParseExecuteReportDecoded(log types.Log) (*ReportCodecExecuteReportDecoded, error) {
	event := new(ReportCodecExecuteReportDecoded)
	if err := _ReportCodec.contract.UnpackLog(event, "ExecuteReportDecoded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_ReportCodec *ReportCodec) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _ReportCodec.abi.Events["CommitReportDecoded"].ID:
		return _ReportCodec.ParseCommitReportDecoded(log)
	case _ReportCodec.abi.Events["ExecuteReportDecoded"].ID:
		return _ReportCodec.ParseExecuteReportDecoded(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ReportCodecCommitReportDecoded) Topic() common.Hash {
	return common.HexToHash("0x1b2cb5e9d31bdaabb2ae07532436ae669406f84003ca27179b4dfb72f127f7dc")
}

func (ReportCodecExecuteReportDecoded) Topic() common.Hash {
	return common.HexToHash("0x2e49eb9d5a71b8aa2af6c92c70eefbb46b7032fc54e580d002b124f14d4f993d")
}

func (_ReportCodec *ReportCodec) Address() common.Address {
	return _ReportCodec.address
}

type ReportCodecInterface interface {
	DecodeCommitReport(opts *bind.CallOpts, report []byte) (EVM2EVMMultiOffRampCommitReport, error)

	DecodeExecuteReport(opts *bind.CallOpts, report []byte) ([]InternalExecutionReportSingleChain, error)

	FilterCommitReportDecoded(opts *bind.FilterOpts) (*ReportCodecCommitReportDecodedIterator, error)

	WatchCommitReportDecoded(opts *bind.WatchOpts, sink chan<- *ReportCodecCommitReportDecoded) (event.Subscription, error)

	ParseCommitReportDecoded(log types.Log) (*ReportCodecCommitReportDecoded, error)

	FilterExecuteReportDecoded(opts *bind.FilterOpts) (*ReportCodecExecuteReportDecodedIterator, error)

	WatchExecuteReportDecoded(opts *bind.WatchOpts, sink chan<- *ReportCodecExecuteReportDecoded) (event.Subscription, error)

	ParseExecuteReportDecoded(log types.Log) (*ReportCodecExecuteReportDecoded, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
