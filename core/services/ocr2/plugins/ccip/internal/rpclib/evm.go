package rpclib

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
)

//go:generate mockery --quiet --name EvmBatchCaller --output . --filename evm_mock.go --inpackage --case=underscore
type EvmBatchCaller interface {
	BatchCall(ctx context.Context, batchSender client.BatchSender, blockNumber uint64, calls []EvmCall) ([]DataAndErr, error)
	BatchCallLimit(ctx context.Context, limit int, batchSender client.BatchSender, blockNumber uint64, calls []EvmCall) ([]DataAndErr, error)
}

type DefaultEvmBatchCaller struct{}

func NewDefaultEvmBatchCaller() *DefaultEvmBatchCaller {
	return &DefaultEvmBatchCaller{}
}

// BatchCall will make a single batched rpc call for all the provided contract calls.
// It supports contract calls that return a single value of type T.
func (c *DefaultEvmBatchCaller) BatchCall(ctx context.Context, batchSender client.BatchSender, blockNumber uint64, calls []EvmCall) ([]DataAndErr, error) {
	if len(calls) == 0 {
		return nil, nil
	}

	packedOutputs := make([]string, len(calls))
	rpcBatchCalls := make([]rpc.BatchElem, len(calls))

	for i, call := range calls {
		packedInputs, err := call.abi.Pack(call.methodName, call.args...)
		if err != nil {
			return nil, fmt.Errorf("pack %s(%+v): %w", call.methodName, call.args, err)
		}

		bn := big.NewInt(0).SetUint64(blockNumber)

		rpcBatchCalls[i] = rpc.BatchElem{
			Method: "eth_call",
			Args: []any{
				map[string]interface{}{
					"from": common.Address{},
					"to":   call.contractAddress,
					"data": hexutil.Bytes(packedInputs),
				},
				hexutil.EncodeBig(bn),
			},
			Result: &packedOutputs[i],
		}
	}

	err := batchSender.BatchCallContext(ctx, rpcBatchCalls)
	if err != nil {
		return nil, fmt.Errorf("batch call context: %w", err)
	}

	results := make([]DataAndErr, len(calls))
	for i, call := range calls {
		if rpcBatchCalls[i].Error != nil {
			results[i].Err = rpcBatchCalls[i].Error
			continue
		}

		b, err := hexutil.Decode(packedOutputs[i])
		if err != nil {
			return nil, err
		}

		unpackedOutputs, err := call.abi.Unpack(call.methodName, b)
		if err != nil {
			return nil, fmt.Errorf("unpack result %s(%+v): %w", call.methodName, call.args, err)
		}
		results[i].Outputs = unpackedOutputs
	}

	return results, nil
}

// BatchCallLimit is similar to EvmBatchCall but splits the batches into sub-calls.
// For example if you want to make 100 calls and pass limit=50, then 2 rpc calls will be made.
func (c *DefaultEvmBatchCaller) BatchCallLimit(ctx context.Context, limit int, batchSender client.BatchSender, blockNumber uint64, calls []EvmCall) ([]DataAndErr, error) {
	if limit <= 0 {
		return c.BatchCall(ctx, batchSender, blockNumber, calls)
	}

	results := make([]DataAndErr, 0, len(calls))

	for i := 0; i < len(calls); i += limit {
		idxFrom := i
		idxTo := idxFrom + limit
		if idxTo > len(calls) {
			idxTo = len(calls)
		}

		subResults, err := c.BatchCall(ctx, batchSender, blockNumber, calls[idxFrom:idxTo])
		if err != nil {
			return nil, err
		}
		results = append(results, subResults...)
	}

	return results, nil
}

type AbiPackerUnpacker interface {
	Pack(name string, args ...interface{}) ([]byte, error)
	Unpack(name string, data []byte) ([]interface{}, error)
}

type EvmCall struct {
	abi             AbiPackerUnpacker
	methodName      string
	contractAddress common.Address
	args            []any
}

func NewEvmCall(abi AbiPackerUnpacker, methodName string, contractAddress common.Address, args ...any) EvmCall {
	return EvmCall{
		abi:             abi,
		methodName:      methodName,
		contractAddress: contractAddress,
		args:            args,
	}
}

type EvmCallResponse struct {
	Data []byte
	Err  error
}

type DataAndErr struct {
	Outputs []any
	Err     error
}
