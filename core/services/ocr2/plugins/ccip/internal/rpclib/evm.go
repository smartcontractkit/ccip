package rpclib

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

//go:generate mockery --quiet --name EvmBatchCaller --output ./rpclibmocks --outpkg rpclibmocks --filename evm_mock.go --case=underscore
type EvmBatchCaller interface {
	// BatchCall executes all the provided EvmCall and returns the results in the same order
	// of the calls.
	BatchCall(ctx context.Context, blockNumber uint64, calls []EvmCall) ([]DataAndErr, error)
}

const (
	// DefaultRpcBatchSizeLimit defines the maximum number of rpc requests to be included in a batch.
	DefaultRpcBatchSizeLimit = 100

	// DefaultRpcBatchBackOffMultiplier defines the rate of reducing the batch size limit for retried calls.
	// For example if limit is 20 and multiplier is 4:
	// 1.        20
	// 2. 20/4 = 5
	// 3. 5/4  = 1
	DefaultRpcBatchBackOffMultiplier = 5
)

// DynamicLimitedBatchCaller makes batched rpc calls and perform retries by reducing the batch size on each retry.
type DynamicLimitedBatchCaller struct {
	bc *defaultEvmBatchCaller
}

func NewDynamicLimitedBatchCaller(lggr logger.Logger, batchSender client.BatchSender, batchSizeLimit, backOffMultiplier uint) *DynamicLimitedBatchCaller {
	return &DynamicLimitedBatchCaller{
		bc: newDefaultEvmBatchCaller(lggr, batchSender, batchSizeLimit, backOffMultiplier),
	}
}

func (c *DynamicLimitedBatchCaller) BatchCall(ctx context.Context, blockNumber uint64, calls []EvmCall) ([]DataAndErr, error) {
	return c.bc.batchCallDynamicLimitRetries(ctx, blockNumber, calls)
}

type defaultEvmBatchCaller struct {
	lggr              logger.Logger
	batchSender       client.BatchSender
	batchSizeLimit    uint
	backOffMultiplier uint
}

// NewDefaultEvmBatchCaller returns a new batch caller instance.
// batchCallLimit defines the maximum number of calls for BatchCallLimit method, pass 0 to keep the default.
// backOffMultiplier defines the back-off strategy for retries on BatchCallDynamicLimitRetries method, pass 0 to keep the default.
func newDefaultEvmBatchCaller(lggr logger.Logger, batchSender client.BatchSender, batchSizeLimit, backOffMultiplier uint) *defaultEvmBatchCaller {
	batchSize := uint(DefaultRpcBatchSizeLimit)
	if batchSizeLimit > 0 {
		batchSize = batchSizeLimit
	}

	multiplier := uint(DefaultRpcBatchBackOffMultiplier)
	if backOffMultiplier > 0 {
		multiplier = backOffMultiplier
	}

	return &defaultEvmBatchCaller{
		lggr:              lggr,
		batchSender:       batchSender,
		batchSizeLimit:    batchSize,
		backOffMultiplier: multiplier,
	}
}

func (c *defaultEvmBatchCaller) batchCall(ctx context.Context, blockNumber uint64, calls []EvmCall) ([]DataAndErr, error) {
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

	err := c.batchSender.BatchCallContext(ctx, rpcBatchCalls)
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

func (c *defaultEvmBatchCaller) batchCallDynamicLimitRetries(ctx context.Context, blockNumber uint64, calls []EvmCall) ([]DataAndErr, error) {
	lim := c.batchSizeLimit
	for {
		results, err := c.batchCallLimit(ctx, blockNumber, calls, lim)
		if err == nil {
			return results, nil
		}

		if lim <= 1 {
			return nil, err
		}

		newLim := lim / c.backOffMultiplier
		if newLim == 0 || newLim == lim {
			newLim = 1
		}
		lim = newLim
		c.lggr.Errorf("retrying batch call with %d calls and %d limit that failed with error=%s",
			len(calls), lim, err)
	}
}

func (c *defaultEvmBatchCaller) batchCallLimit(ctx context.Context, blockNumber uint64, calls []EvmCall, batchSizeLimit uint) ([]DataAndErr, error) {
	if batchSizeLimit <= 0 {
		return c.batchCall(ctx, blockNumber, calls)
	}

	results := make([]DataAndErr, 0, len(calls))

	for i := 0; i < len(calls); i += int(batchSizeLimit) {
		idxFrom := i
		idxTo := idxFrom + int(batchSizeLimit)
		if idxTo > len(calls) {
			idxTo = len(calls)
		}

		subResults, err := c.batchCall(ctx, blockNumber, calls[idxFrom:idxTo])
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

type DataAndErr struct {
	Outputs []any
	Err     error
}

func ParseOutputs[T any](results []DataAndErr, parseFunc func(d DataAndErr) (T, error)) ([]T, error) {
	parsed := make([]T, 0, len(results))

	for _, res := range results {
		v, err := parseFunc(res)
		if err != nil {
			return nil, fmt.Errorf("parse contract output: %w", err)
		}
		parsed = append(parsed, v)
	}

	return parsed, nil
}

func ParseOutput[T any](dataAndErr DataAndErr, idx int) (T, error) {
	var parsed T

	if dataAndErr.Err != nil {
		return parsed, fmt.Errorf("rpc call error: %w", dataAndErr.Err)
	}

	if idx < 0 || idx >= len(dataAndErr.Outputs) {
		return parsed, fmt.Errorf("idx %d is out of bounds for %d outputs", idx, len(dataAndErr.Outputs))
	}

	res, is := dataAndErr.Outputs[idx].(T)
	if !is {
		// some rpc types are not strictly defined
		// for that reason we try to manually map the fields using json encoding
		b, err := json.Marshal(dataAndErr.Outputs[idx])
		if err == nil {
			var empty T
			if err := json.Unmarshal(b, &parsed); err == nil && !reflect.DeepEqual(parsed, empty) {
				return parsed, nil
			}
		}

		return parsed, fmt.Errorf("the result type is: %T, expected: %T", dataAndErr.Outputs[idx], parsed)
	}

	return res, nil
}
