package rpclib

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type EvmBatchCaller interface {
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
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

// EvmBatchCall will make a single batched rpc call for all the provided contract calls.
// It supports contract calls that return a single value of type T.
func EvmBatchCall(ctx context.Context, evmBatchCaller EvmBatchCaller, calls ...EvmCall) ([]DataAndErr, error) {
	if len(calls) == 0 {
		return nil, nil
	}

	bytesOutputs := make([]hexutil.Bytes, len(calls))
	rpcBatchCalls := make([]rpc.BatchElem, len(calls))

	for i, call := range calls {
		packedInputs, err := call.abi.Pack(call.methodName, call.args...)
		if err != nil {
			return nil, fmt.Errorf("pack %s(%+v): %w", call.methodName, call.args, err)
		}

		rpcBatchCalls[i] = rpc.BatchElem{
			Method: "eth_call",
			Args: []any{
				map[string]interface{}{
					"from": common.Address{},
					"to":   call.contractAddress,
					"data": hexutil.Bytes(packedInputs),
				},
				"latest",
			},
			Result: &bytesOutputs[i],
		}
	}

	err := evmBatchCaller.BatchCallContext(ctx, rpcBatchCalls)
	if err != nil {
		return nil, fmt.Errorf("batch call context: %w", err)
	}

	results := make([]DataAndErr, len(calls))
	for i, call := range calls {
		if rpcBatchCalls[i].Error != nil {
			results[i].Err = rpcBatchCalls[i].Error
			continue
		}

		unpackedOutputs, err := call.abi.Unpack(call.methodName, bytesOutputs[i])
		if err != nil {
			return nil, fmt.Errorf("unpack result %s(%+v): %w", call.methodName, call.args, err)
		}
		results[i].Outputs = unpackedOutputs
	}

	return results, nil
}

// EvmBatchCallWithLimit is similar to EvmBatchCall but splits the batches into sub-calls.
// For example if you want to make 100 calls and pass limit=50, then 2 rpc calls will be made.
func EvmBatchCallWithLimit(ctx context.Context, limit int, evmBatchCaller EvmBatchCaller, calls ...EvmCall) ([]DataAndErr, error) {
	if limit <= 0 {
		return EvmBatchCall(ctx, evmBatchCaller, calls...)
	}

	results := make([]DataAndErr, 0, len(calls))

	for i := 0; i < len(calls); i += limit {
		idxFrom := i
		idxTo := idxFrom + limit
		if idxTo > len(calls) {
			idxTo = len(calls)
		}

		subResults, err := EvmBatchCall(ctx, evmBatchCaller, calls[idxFrom:idxTo]...)
		if err != nil {
			return nil, err
		}
		results = append(results, subResults...)
	}

	return results, nil
}
