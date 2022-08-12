package ccip

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
)

const (
	MAX_OVERHEAD_GAS_TOLL = 0 // TODO: Once contracts stable, add the worst case overhead gas outside of user callback for toll offramp.
)

type TollExecutionBatch struct {
	remainingGasLimit    uint64
	maxGasPrice          uint64
	srcToDstToken        map[common.Address]common.Address
	tollTokensPerFeeCoin map[common.Address]*big.Int
	seqNrs               []uint64
}

func NewTollExecutionBatch(maxGasLimit uint64, maxGasPrice uint64, srcToDstToken map[common.Address]common.Address, tokensPerFeeCoin map[common.Address]*big.Int) *TollExecutionBatch {
	return &TollExecutionBatch{
		remainingGasLimit:    maxGasLimit,
		maxGasPrice:          maxGasPrice,
		srcToDstToken:        srcToDstToken,
		tollTokensPerFeeCoin: tokensPerFeeCoin,
	}
}

func (teb *TollExecutionBatch) Add(msg Message) bool {
	if big.NewInt(int64(teb.remainingGasLimit)).Cmp(msg.GasLimit) < 1 {
		return false
	}
	dstToken, present := teb.srcToDstToken[msg.FeeToken]
	if !present {
		// TODO: error?
		return false
	}
	tollTokensPerFeeCoin := teb.tollTokensPerFeeCoin[dstToken]
	if tollTokensPerFeeCoin == nil {
		// TODO: error?
		return false
	}

	maxCostInFeeCoin := new(big.Int).Div(new(big.Int).Mul(new(big.Int).Mul(big.NewInt(MaxGasPrice), msg.GasLimit), tollTokensPerFeeCoin), big.NewInt(1e18))
	if msg.FeeTokenAmount.Cmp(maxCostInFeeCoin) == -1 {
		return false
	}
	teb.remainingGasLimit -= msg.GasLimit.Uint64() // TODO: this should probably include overhead?
	teb.seqNrs = append(teb.seqNrs, msg.SequenceNumber)
	return true
}

func (teb *TollExecutionBatch) SeqNrs() []uint64 {
	return teb.seqNrs
}

type TollBatchBuilder struct {
	tollABI abi.ABI
	lggr    logger.Logger
}

func NewTollBatchBuilder(lggr logger.Logger) *TollBatchBuilder {
	tollABI, _ := abi.JSON(strings.NewReader(evm_2_evm_toll_onramp.EVM2EVMTollOnRampABI))
	return &TollBatchBuilder{
		tollABI: tollABI,
		lggr:    lggr,
	}
}

func (tb *TollBatchBuilder) parseLog(log types.Log) (*evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested, error) {
	event := new(evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested)
	err := bind.NewBoundContract(common.Address{}, tb.tollABI, nil, nil, nil).UnpackLog(event, "CCIPSendRequested", log)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (tb *TollBatchBuilder) BuildBatch(srcToDst map[common.Address]common.Address, msgs []logpoller.Log, executed map[uint64]struct{}, gasLimit uint64, gasPrice uint64, tollTokensPerFeeCoin map[common.Address]*big.Int, inflight []InflightExecutionReport) []uint64 {
	inflightSeqNrs := tb.inflightSeqNrs(inflight)
	tollBatch := NewTollExecutionBatch(gasLimit, gasPrice, srcToDst, tollTokensPerFeeCoin)
	haveOne := false
	for _, msg := range msgs {
		tollMsgEvent, err := tb.parseLog(types.Log{
			Topics: msg.GetTopics(),
			Data:   msg.Data,
		})
		if err != nil {
			tb.lggr.Errorw("unable to parse message", "err", err, "msg", msg)
			continue
		}
		tollMsg := EVM2EVMTollEventToMessage(tollMsgEvent.Message)
		if _, inflight := inflightSeqNrs[tollMsg.SequenceNumber]; inflight {
			continue
		}
		if _, executed := executed[tollMsg.SequenceNumber]; executed {
			continue
		}
		added := tollBatch.Add(tollMsg)
		if !added && haveOne {
			break
		}
		if added && !haveOne {
			haveOne = true
		}
	}
	return tollBatch.SeqNrs()
}

func (tb *TollBatchBuilder) inflightSeqNrs(inflight []InflightExecutionReport) map[uint64]struct{} {
	inflightSeqNrs := make(map[uint64]struct{})
	for _, rep := range inflight {
		for _, seqNr := range rep.report.SequenceNumbers {
			inflightSeqNrs[seqNr] = struct{}{}
		}
	}
	return inflightSeqNrs
}
