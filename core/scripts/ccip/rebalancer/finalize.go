package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip/rebalancer/multienv"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	utilsbig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arb_node_interface"
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
	arg0Proof := getProof(env, l2ChainID, withdrawalInitiated.L2ToL1Id)
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
}

func encodeProofToHex(proof [][32]byte) []string {
	proofHex := make([]string, len(proof))
	for i, step := range proof {
		proofHex[i] = hexutil.Encode(step[:])
	}
	return proofHex
}

// Arg 0: proof. This takes multiple steps:
// 1. Get the latest L2 block
// 2. Call eth_getBlockByHash specifying the latest L2 block hash.
// 3. Get the `sendCount` field from the response.
// 4. Get the `l2ToL1Id` field from the `WithdrawalInitiated` log from the L2 withdrawal tx.
// 5. Call `constructOutboxProof` on the L2 node interface contract with the `sendCount` as the first argument and `l2ToL1Id` as the second argument.
// Note that the proof seems to change as the rollup accumulates more blocks, which makes sense.
func getProof(env multienv.Env, l2ChainID uint64, l2ToL1Id *big.Int) [][32]byte {
	l2Client := env.Clients[l2ChainID]
	latestL2Block, err := l2Client.HeaderByNumber(context.Background(), nil)
	helpers.PanicErr(err)
	l2Rpc := env.JRPCs[l2ChainID]
	type Response struct {
		SendCount *utilsbig.Big `json:"sendCount"`
	}
	response := new(Response)
	err = l2Rpc.Call(response, "eth_getBlockByHash", latestL2Block.Hash().Hex(), true)
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
