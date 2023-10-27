package rpclib

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

//go:generate mockery --quiet --name EvmBatchCaller --output . --filename evm_mock.go --inpackage --case=underscore
type EvmBatchCaller interface {
	// BatchCall will make a single batched rpc call for all the provided contract calls.
	BatchCall(ctx context.Context, blockNumber uint64, calls []EvmCall) ([]DataAndErr, error)

	// BatchCallLimit is similar to EvmBatchCall but splits the batches into sub-calls based on the defined batch size limit.
	// This method should be preferred over BatchCall for large payloads.
	BatchCallLimit(ctx context.Context, blockNumber uint64, calls []EvmCall) ([]DataAndErr, error)

	// BatchCallDynamicLimitRetries is similar to BatchCallLimit but will perform retries by reducing the batch size on each retry.
	// It is preferred when the RPC behavior/limits are unknown or when retrying is required.
	BatchCallDynamicLimitRetries(ctx context.Context, blockNumber uint64, calls []EvmCall) ([]DataAndErr, error)
}

const (
	// DefaultRpcBatchSizeLimit defines the maximum number of rpc requests to be included in a batch.
	DefaultRpcBatchSizeLimit = 20

	// DefaultRpcBatchBackOffMultiplier defines the rate of reducing the batch size limit for retried calls.
	// For example if limit is 20 and multiplier is 4:
	// 1.        20
	// 2. 20/4 = 5
	// 3. 5/4  = 1
	DefaultRpcBatchBackOffMultiplier = 4
)

type DefaultEvmBatchCaller struct {
	lggr              logger.Logger
	batchSender       client.BatchSender
	batchSizeLimit    int
	backOffMultiplier int
}

// NewDefaultEvmBatchCaller returns a new batch caller instance.
// batchCallLimit defines the maximum number of calls for BatchCallLimit method, pass 0 to keep the default.
// backOffMultiplier defines the back-off strategy for retries on BatchCallDynamicLimitRetries method, pass 0 to keep the default.
func NewDefaultEvmBatchCaller(lggr logger.Logger, batchSender client.BatchSender, batchSizeLimit, backOffMultiplier int) *DefaultEvmBatchCaller {
	batchSize := DefaultRpcBatchSizeLimit
	if batchSizeLimit > 0 {
		batchSize = batchSizeLimit
	}

	multiplier := DefaultRpcBatchBackOffMultiplier
	if backOffMultiplier > 0 {
		multiplier = backOffMultiplier
	}

	return &DefaultEvmBatchCaller{
		lggr:              lggr,
		batchSender:       batchSender,
		batchSizeLimit:    batchSize,
		backOffMultiplier: multiplier,
	}
}

func (c *DefaultEvmBatchCaller) BatchCall(ctx context.Context, blockNumber uint64, calls []EvmCall) ([]DataAndErr, error) {
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

func (c *DefaultEvmBatchCaller) BatchCallDynamicLimitRetries(ctx context.Context, blockNumber uint64, calls []EvmCall) ([]DataAndErr, error) {
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

func (c *DefaultEvmBatchCaller) BatchCallLimit(ctx context.Context, blockNumber uint64, calls []EvmCall) ([]DataAndErr, error) {
	return c.batchCallLimit(ctx, blockNumber, calls, c.batchSizeLimit)
}

func (c *DefaultEvmBatchCaller) batchCallLimit(ctx context.Context, blockNumber uint64, calls []EvmCall, batchSizeLimit int) ([]DataAndErr, error) {
	if batchSizeLimit <= 0 {
		return c.BatchCall(ctx, blockNumber, calls)
	}

	results := make([]DataAndErr, 0, len(calls))

	for i := 0; i < len(calls); i += batchSizeLimit {
		idxFrom := i
		idxTo := idxFrom + batchSizeLimit
		if idxTo > len(calls) {
			idxTo = len(calls)
		}

		subResults, err := c.BatchCall(ctx, blockNumber, calls[idxFrom:idxTo])
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
