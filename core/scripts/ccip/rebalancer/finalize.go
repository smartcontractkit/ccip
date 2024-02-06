package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip/rebalancer/arb"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip/rebalancer/multienv"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	utilsbig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arb_node_interface"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbsys"
)

var (
	// Events emitted on L2
	TxToL1Topic              = common.HexToHash("0x2b986d32a0536b7e19baa48ab949fec7b903b7fad7730820b20632d100cc3a68")
	WithdrawalInitiatedTopic = common.HexToHash("0x3073a74ecb728d10be779fe19a74a1428e20468f5b4d167bf9c73d9067847d73")
	L2ToL1TxTopic            = common.HexToHash("0x3e7aafa77dbf186b7fd488006beff893744caa3c4f6f299e8a709fa2087374fc")
	TransferRoutedTopic      = common.HexToHash("0x85291dff2161a93c2f12c819d31889c96c63042116f5bc5a205aa701c2c429f5")

	// Important addresses on L2
	NodeInterfaceAddress = common.HexToAddress("0x00000000000000000000000000000000000000c8")
	ArbSysAddress        = common.HexToAddress("0x0000000000000000000000000000000000000064")

	// Events emitted on L1
	NodeConfirmedTopic = common.HexToHash("0x22ef0479a7ff660660d1c2fe35f1b632cf31675c2d9378db8cec95b00d8ffa3c")
)

// function executeTransaction(
//
//	  bytes32[] calldata proof,
//	  uint256 index,
//	  address l2Sender,
//	  address to,
//	  uint256 l2Block,
//	  uint256 l1Block,
//	  uint256 l2Timestamp,
//	  uint256 value,
//	  bytes calldata data
//	) external;
//
// Arg 0: proof. This takes multiple steps:
// 1. Get the latest L2 block
// 2. Call eth_getBlockByHash specifying the latest L2 block hash.
// 3. Get the `sendCount` field from the response.
// 4. Get the `l2ToL1Id` field from the `WithdrawalInitiated` log from the L2 withdrawal tx.
// 5. Call `constructOutboxProof` on the L2 node interface contract with the `sendCount` as the first argument and `l2ToL1Id` as the second argument.
// Arg 1: index. Fetch the index from the TxToL1 log in the L2 tx.
// Arg 2: l2Sender. Fetch the source of the WithdrawalInitiated log in the L2 tx.
// Arg 3: to. Fetch the `to` field of the WithdrawalInitiated log in the L2 tx.
// Arg 4: l1Block. Fetch the `l1BlockNumber` field of the JSON-RPC response to eth_getTransactionReceipt
// passing in the L2 tx hash as the param.
// Arg 5: l2Block. This is the l2 block number in which the withdrawal tx was included.
// Arg 6: l2Timestamp. Get the `timestamp` field from the L2ToL1Tx event emitted by ArbSys (0x64).
// Arg 7: value. Fetch the `value` field from the WithdrawalInitiated log in the L2 tx.
// Arg 8: data. Fetch the `data` field from the TxToL1 log in the L2 tx.
func arbFinalizeL1(
	env multienv.Env,
	l1ChainID uint64,
	l2ChainID uint64,
	l1BridgeAdapterAddress common.Address,
	l2TxHash common.Hash,
) {
	// get the logs we care about from the L2 tx:
	// 1. L2ToL1Tx
	// 2. WithdrawalInitiated
	// 3. TxToL1
	l2Client := env.Clients[l2ChainID]
	receipt, err := l2Client.TransactionReceipt(context.Background(), l2TxHash)
	helpers.PanicErr(err)
	var (
		l2ToL1TxLog, withdrawalInitiatedLog, txToL1Log *types.Log
	)
	for _, lg := range receipt.Logs {
		if lg.Topics[0] == L2ToL1TxTopic {
			l2ToL1TxLog = lg
		} else if lg.Topics[0] == WithdrawalInitiatedTopic {
			withdrawalInitiatedLog = lg
		} else if lg.Topics[0] == TxToL1Topic {
			txToL1Log = lg
		}
	}
	if l2ToL1TxLog == nil || withdrawalInitiatedLog == nil || txToL1Log == nil {
		helpers.PanicErr(fmt.Errorf("missing logs in L2 tx %s", l2TxHash.String()))
		return
	}
	arbSys, err := arbsys.NewArbSys(ArbSysAddress, env.Clients[l2ChainID])
	helpers.PanicErr(err)
	// parse logs
	l2ToL1Tx, err := arbSys.ParseL2ToL1Tx(*l2ToL1TxLog)
	helpers.PanicErr(err)
	withdrawalInitiated := parseWithdrawalInitiated(withdrawalInitiatedABI(), withdrawalInitiatedLog)
	txToL1 := parseTxToL1(txToL1ABI(), txToL1Log)
	// get the proof
	arg0Proof := getProof(env, l1ChainID, l2ChainID, withdrawalInitiated.L2ToL1Id)
	// argument 1: index
	arg1Index := withdrawalInitiated.L2ToL1Id
	// argument 2: l2Sender
	arg2L2Sender := withdrawalInitiatedLog.Address
	// argument 3: to
	arg3To := txToL1.To
	// argument 4: l1Block
	arg4L1Block := getL1BlockFromRPC(env, l2ChainID, l2TxHash)
	// argument 5: l2Block
	arg5L2Block := receipt.BlockNumber
	// argument 6: l2Timestamp
	arg6L2Timestamp := l2ToL1Tx.Timestamp
	// argument 7: value
	arg7Value := withdrawalInitiated.Amount
	// argument 8: data
	arg8Data := txToL1.Data

	// print the arguments for the executeTransaction call
	fmt.Println("proof:", encodeProofToHex(arg0Proof), "\n",
		"index:", arg1Index, "\n",
		"l2Sender:", arg2L2Sender, "\n",
		"to:", arg3To, "\n",
		"l1Block:", arg4L1Block, "\n",
		"l2Block:", arg5L2Block, "\n",
		"l2Timestamp:", arg6L2Timestamp, "\n",
		"value:", arg7Value, "\n",
		"data:", hexutil.Encode(arg8Data))

	// execute the transaction
	fmt.Println("executing transaction on the bridge adapter with the above data")

	adapter, err := arbitrum_l1_bridge_adapter.NewArbitrumL1BridgeAdapter(l1BridgeAdapterAddress, env.Clients[l1ChainID])
	helpers.PanicErr(err)

	adapterABI, err := arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterMetaData.GetAbi()
	helpers.PanicErr(err)

	finalizationPayload, err := adapterABI.Pack("exposeForEncoding", arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumFinalizationPayload{
		Proof:       arg0Proof,
		Index:       arg1Index,
		L2Sender:    arg2L2Sender,
		To:          arg3To,
		L1Block:     arg4L1Block,
		L2Block:     arg5L2Block,
		L2Timestamp: arg6L2Timestamp,
		Value:       arg7Value,
		Data:        arg8Data,
	})
	helpers.PanicErr(err)

	// packed, err := adapterABI.Pack("finalizeWithdrawERC20", common.HexToAddress("0x0"), common.HexToAddress("0x0"), finalizationPayload)
	// helpers.PanicErr(err)

	// nonce, err := env.Clients[l1ChainID].PendingNonceAt(context.Background(), env.Transactors[l1ChainID].From)
	// helpers.PanicErr(err)

	// gasPrice, err := env.Clients[l1ChainID].SuggestGasPrice(context.Background())
	// helpers.PanicErr(err)

	// rawTx := types.NewTx(&types.LegacyTx{
	// 	Nonce:    nonce,
	// 	To:       &l1BridgeAdapterAddress,
	// 	Value:    big.NewInt(0),
	// 	GasPrice: gasPrice,
	// 	Gas:      1e6,
	// 	Data:     packed,
	// })
	// signedTx, err := env.Transactors[l1ChainID].Signer(env.Transactors[l1ChainID].From, rawTx)
	// helpers.PanicErr(err)

	// err = env.Clients[l1ChainID].SendTransaction(context.Background(), signedTx)
	tx, err := adapter.FinalizeWithdrawERC20(env.Transactors[l1ChainID], common.HexToAddress("0x0"), common.HexToAddress("0x0"), finalizationPayload)
	helpers.PanicErr(err)
	helpers.ConfirmTXMined(context.Background(), env.Clients[l1ChainID], tx, int64(l1ChainID))
}

func encodeProofToHex(proof [][32]byte) []string {
	proofHex := make([]string, len(proof))
	for i, step := range proof {
		proofHex[i] = hexutil.Encode(step[:])
	}
	return proofHex
}

// Arg 0: proof. This takes multiple steps:
// 1. Get the latest NodeConfirmed event on L1, which indicates the latest node that was confirmed by the rollup.
// 2. Call eth_getBlockByHash on L2 specifying the L2 block hash in the NodeConfirmed event.
// 3. Get the `sendCount` field from the response.
// 4. Get the `l2ToL1Id` field from the `WithdrawalInitiated` log from the L2 withdrawal tx.
// 5. Call `constructOutboxProof` on the L2 node interface contract with the `sendCount` as the first argument and `l2ToL1Id` as the second argument.
// Note that the proof seems to change as the rollup accumulates more blocks, which makes sense. All that matters is that the root is committed
// to the rollup prior to posting the proof to L1.
func getProof(env multienv.Env, l1ChainID, l2ChainID uint64, l2ToL1Id *big.Int) [][32]byte {
	l1Client := env.Clients[l1ChainID]
	latestHeader, err := l1Client.HeaderByNumber(context.Background(), nil)
	helpers.PanicErr(err)
	// start two hours back in terms of blocks
	// 12 seconds per block => 5 * 120 = 600 blocks
	startBlock := big.NewInt(0).Sub(latestHeader.Number, big.NewInt(600))
	lgs, err := l1Client.FilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: []common.Address{arb.ArbitrumContracts[l1ChainID]["Rollup"]},
		Topics: [][]common.Hash{{
			NodeConfirmedTopic,
		}},
		FromBlock: startBlock,
	})
	helpers.PanicErr(err)
	var latestNodeConfirmed *types.Log
	for _, lg := range lgs {
		if latestNodeConfirmed == nil || lg.BlockNumber > latestNodeConfirmed.BlockNumber {
			latestNodeConfirmed = &lg
		}
	}
	if latestNodeConfirmed == nil {
		helpers.PanicErr(fmt.Errorf("no node confirmed event found"))
	}
	// parse latest nodeconfirmed event
	nodeConfirmed := parseNodeConfirmed(nodeConfirmedABI(), latestNodeConfirmed)
	fmt.Println("latest node confirmed:", nodeConfirmed)
	type Response struct {
		SendCount *utilsbig.Big `json:"sendCount"`
	}
	response := new(Response)
	l2Rpc := env.JRPCs[l2ChainID]
	err = l2Rpc.Call(response, "eth_getBlockByHash", hexutil.Encode(nodeConfirmed.BlockHash[:]), false)
	helpers.PanicErr(err)
	nodeInterface, err := arb_node_interface.NewNodeInterface(NodeInterfaceAddress, env.Clients[l2ChainID])
	helpers.PanicErr(err)
	fmt.Println("send count:", response.SendCount, "l2 to l1 id:", l2ToL1Id)
	outboxProof, err := nodeInterface.ConstructOutboxProof(nil, response.SendCount.ToInt().Uint64(), l2ToL1Id.Uint64())
	helpers.PanicErr(err)
	return outboxProof.Proof
}

func getL1BlockFromRPC(env multienv.Env, l2ChainID uint64, l2TxHash common.Hash) *big.Int {
	l1Rpc := env.JRPCs[l2ChainID]
	type Response struct {
		L1BlockNumber hexutil.Big `json:"l1BlockNumber"`
	}
	response := new(Response)
	err := l1Rpc.Call(response, "eth_getTransactionReceipt", l2TxHash.String())
	helpers.PanicErr(err)
	return response.L1BlockNumber.ToInt()
}

func withdrawalInitiatedABI() abi.ABI {
	jsonABI := `[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"l1Token","type":"address"},{"indexed":true,"internalType":"address","name":"_from","type":"address"},{"indexed":true,"internalType":"address","name":"_to","type":"address"},{"indexed":true,"internalType":"uint256","name":"_l2ToL1Id","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"_exitNum","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"_amount","type":"uint256"}],"name":"WithdrawalInitiated","type":"event"}]`
	abi, err := abi.JSON(strings.NewReader(jsonABI))
	helpers.PanicErr(err)
	return abi
}

type WithdrawalInitiated struct {
	L1Token  common.Address
	From     common.Address
	To       common.Address
	L2ToL1Id *big.Int
	ExitNum  *big.Int
	Amount   *big.Int
}

func parseWithdrawalInitiated(tabi abi.ABI, lg *types.Log) *WithdrawalInitiated {
	event := new(WithdrawalInitiated)
	if err := UnpackLog(tabi, event, "WithdrawalInitiated", *lg); err != nil {
		helpers.PanicErr(err)
	}
	return event
}

func txToL1ABI() abi.ABI {
	jsonABI := `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"_from","type":"address"},{"indexed":true,"internalType":"address","name":"_to","type":"address"},{"indexed":true,"internalType":"uint256","name":"_id","type":"uint256"},{"indexed":false,"internalType":"bytes","name":"_data","type":"bytes"}],"name":"TxToL1","type":"event"}]`
	abi, err := abi.JSON(strings.NewReader(jsonABI))
	helpers.PanicErr(err)
	return abi
}

type TxToL1 struct {
	From, To common.Address
	Id       *big.Int
	Data     []byte
}

func parseTxToL1(tabi abi.ABI, lg *types.Log) *TxToL1 {
	event := new(TxToL1)
	if err := UnpackLog(tabi, event, "TxToL1", *lg); err != nil {
		helpers.PanicErr(err)
	}
	return event
}

func nodeConfirmedABI() abi.ABI {
	jsonABI := `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint64","name":"nodeNum","type":"uint64"},{"indexed":false,"internalType":"bytes32","name":"blockHash","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"sendRoot","type":"bytes32"}],"name":"NodeConfirmed","type":"event"}]`
	abi, err := abi.JSON(strings.NewReader(jsonABI))
	helpers.PanicErr(err)
	return abi
}

type NodeConfirmed struct {
	NodeNum             uint64
	BlockHash, SendRoot [32]byte
}

func (n NodeConfirmed) String() string {
	return fmt.Sprintf("{NodeNum: %d, BlockHash: %s, SendRoot: %s}", n.NodeNum, hexutil.Encode(n.BlockHash[:]), hexutil.Encode(n.SendRoot[:]))
}

func parseNodeConfirmed(tabi abi.ABI, lg *types.Log) *NodeConfirmed {
	event := new(NodeConfirmed)
	if err := UnpackLog(tabi, event, "NodeConfirmed", *lg); err != nil {
		helpers.PanicErr(err)
	}
	return event
}

func UnpackLog(tabi abi.ABI, out interface{}, event string, log types.Log) error {
	// Anonymous events are not supported.
	if len(log.Topics) == 0 {
		return fmt.Errorf("no event signature")
	}
	if log.Topics[0] != tabi.Events[event].ID {
		return fmt.Errorf("event signature mismatch")
	}
	if len(log.Data) > 0 {
		if err := tabi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tabi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}
