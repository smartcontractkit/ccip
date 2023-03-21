package ccip

import (
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	ccipconfig "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/core/utils/mathutil"
)

const (
	MessageStateUntouched = iota
	MessageStateInProgress
	MessageStateSuccess
	MessageStateFailure
)

const (
	BatchGasLimit = 5_000_000                 // TODO: think if a good value for this
	GasLimitPerTx = BatchGasLimit - 1_000_000 // Leave a buffer for overhead.
	MaxGasPrice   = int64(200e9)              // 200 gwei. TODO: probably want this to be some dynamic value, a multiplier of the current gas price.
)

var (
	ErrCommitStoreIsDown = errors.New("commitStore is down")
)

func LoadOnRamp(onRampAddress common.Address, client client.Client) (*evm_2_evm_onramp.EVM2EVMOnRamp, error) {
	err := ccipconfig.VerifyTypeAndVersion(onRampAddress, client, ccipconfig.EVM2EVMOnRamp)
	if err != nil {
		return nil, errors.Wrap(err, "Invalid onRamp contract")
	}
	return evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddress, client)
}

func LoadOffRamp(offRampAddress common.Address, client client.Client) (*evm_2_evm_offramp.EVM2EVMOffRamp, error) {
	err := ccipconfig.VerifyTypeAndVersion(offRampAddress, client, ccipconfig.EVM2EVMOffRamp)
	if err != nil {
		return nil, errors.Wrap(err, "Invalid offRamp contract")
	}
	return evm_2_evm_offramp.NewEVM2EVMOffRamp(offRampAddress, client)
}

func LoadCommitStore(commitStoreAddress common.Address, client client.Client) (*commit_store.CommitStore, error) {
	err := ccipconfig.VerifyTypeAndVersion(commitStoreAddress, client, ccipconfig.CommitStore)
	if err != nil {
		return nil, errors.Wrap(err, "Invalid commitStore contract")
	}
	return commit_store.NewCommitStore(commitStoreAddress, client)
}

func median(vals []*big.Int) *big.Int {
	valsCopy := make([]*big.Int, len(vals))
	copy(valsCopy[:], vals[:])
	sort.Slice(valsCopy, func(i, j int) bool {
		return valsCopy[i].Cmp(valsCopy[j]) == -1
	})
	return valsCopy[len(valsCopy)/2]
}

type MessageExecution struct {
	encMsgs          [][]byte
	proofs           [][32]byte
	proofSourceFlags []bool
}

func contiguousReqs(lggr logger.Logger, min, max uint64, seqNrs []uint64) bool {
	for i, j := min, 0; i < max && j < len(seqNrs); i, j = i+1, j+1 {
		if seqNrs[j] != i {
			lggr.Errorw("unexpected gap in seq nums", "seq", i)
			return false
		}
	}
	return true
}

func leavesFromIntervals(
	lggr logger.Logger,
	onRamp common.Address,
	eventSigs EventSignatures,
	seqParser func(logpoller.Log) (uint64, error),
	interval commit_store.CommitStoreInterval,
	srcLogPoller logpoller.LogPoller,
	hasher LeafHasherInterface[[32]byte],
	confs int) ([][32]byte, error) {
	// Logs are guaranteed to be in order of seq num, since these are finalized logs only
	// and the contract's seq num is auto-incrementing.
	logs, err := srcLogPoller.LogsDataWordRange(
		eventSigs.SendRequested,
		onRamp,
		eventSigs.SendRequestedSequenceNumberIndex,
		logpoller.EvmWord(interval.Min),
		logpoller.EvmWord(interval.Max),
		confs)
	if err != nil {
		return nil, err
	}
	var seqNrs []uint64
	for _, log := range logs {
		seqNr, err2 := seqParser(log)
		if err2 != nil {
			return nil, err2
		}
		seqNrs = append(seqNrs, seqNr)
	}
	if !contiguousReqs(lggr, interval.Min, interval.Max, seqNrs) {
		return nil, errors.Errorf("do not have full range [%v, %v] have %v", interval.Min, interval.Max, seqNrs)
	}
	var leaves [][32]byte
	for _, log := range logs {
		hash, err2 := hasher.HashLeaf(log.ToGethLog())
		if err2 != nil {
			return nil, err2
		}
		leaves = append(leaves, hash)
	}

	return leaves, nil
}

func aggregateTokenValue(tokenLimitPrices map[common.Address]*big.Int, srcToDst map[common.Address]common.Address, tokens []common.Address, amounts []*big.Int) (*big.Int, error) {
	sum := big.NewInt(0)
	for i := 0; i < len(tokens); i++ {
		price, ok := tokenLimitPrices[srcToDst[tokens[i]]]
		if !ok {
			return nil, errors.Errorf("do not have price for src token %x", tokens[i])
		}
		sum.Add(sum, new(big.Int).Mul(price, amounts[i]))
	}
	return sum, nil
}

// EventSignatures contain pluginType specific signatures and indexes.
// Indexes are zero indexed
type EventSignatures struct {
	SendRequested                            common.Hash
	SendRequestedSequenceNumberIndex         int
	ExecutionStateChanged                    common.Hash
	ExecutionStateChangedSequenceNumberIndex int
}

func commitReport(dstLogPoller logpoller.LogPoller, onRamp common.Address, commitStore *commit_store.CommitStore, seqNr uint64) (commit_store.CommitStoreCommitReport, error) {
	latest, err := dstLogPoller.LatestBlock()
	if err != nil {
		return commit_store.CommitStoreCommitReport{}, err
	}
	logs, err := dstLogPoller.Logs(1, latest, ReportAccepted, commitStore.Address())
	if err != nil {
		return commit_store.CommitStoreCommitReport{}, err
	}
	if len(logs) == 0 {
		return commit_store.CommitStoreCommitReport{}, errors.Errorf("seq number not committed, nothing committed")
	}
	for _, log := range logs {
		reportAccepted, err := commitStore.ParseReportAccepted(types.Log{
			Topics: log.GetTopics(),
			Data:   log.Data,
		})
		if err != nil {
			return commit_store.CommitStoreCommitReport{}, err
		}
		if reportAccepted.Report.Interval.Min <= seqNr && seqNr <= reportAccepted.Report.Interval.Max {
			return reportAccepted.Report, nil
		}
	}
	return commit_store.CommitStoreCommitReport{}, errors.Errorf("seq number not committed")
}

func buildExecution(
	lggr logger.Logger,
	source,
	dest logpoller.LogPoller,
	onRampAddress common.Address,
	finalSeqNums []uint64,
	commitStore *commit_store.CommitStore,
	confs int,
	eventSignatures EventSignatures,
	seqNumFromLog func(log logpoller.Log) (uint64, error),
	hashLeaf LeafHasherInterface[[32]byte],
) (*MessageExecution, error) {
	nextMin, err := commitStore.GetExpectedNextSequenceNumber(nil)
	if err != nil {
		return nil, err
	}
	if mathutil.Max(finalSeqNums[0], finalSeqNums[1:]...) >= nextMin {
		return nil, errors.Errorf("Cannot execute uncommitted seq num. nextMin %v, seqNums %v", nextMin, finalSeqNums)
	}
	rep, err := commitReport(dest, onRampAddress, commitStore, finalSeqNums[0])
	if err != nil {
		return nil, err
	}
	lggr.Infow("Building execution report", "finalSeqNums", finalSeqNums, "merkleRoot", hexutil.Encode(rep.MerkleRoot[:]), "report", rep)

	msgsInRoot, err := source.LogsDataWordRange(
		eventSignatures.SendRequested,
		onRampAddress,
		eventSignatures.SendRequestedSequenceNumberIndex,
		EvmWord(rep.Interval.Min), EvmWord(rep.Interval.Max), confs)
	if err != nil {
		return nil, err
	}
	if len(msgsInRoot) != int(rep.Interval.Max-rep.Interval.Min+1) {
		return nil, errors.Errorf("unexpected missing msgs in committed root %x have %d want %d", rep.MerkleRoot, len(msgsInRoot), int(rep.Interval.Max-rep.Interval.Min+1))
	}
	leaves, err := leavesFromIntervals(lggr, onRampAddress, eventSignatures, seqNumFromLog, rep.Interval, source, hashLeaf, confs)
	if err != nil {
		return nil, err
	}
	ctx := hasher.NewKeccakCtx()
	tree, err := merklemulti.NewTree[[32]byte](ctx, leaves)
	if err != nil {
		return nil, err
	}

	var innerIdxs []int
	var encMsgs [][]byte
	var hashes [][32]byte
	for _, seqNum := range finalSeqNums {
		if seqNum < rep.Interval.Min || seqNum > rep.Interval.Max {
			// We only return messages from a single root (the root of the first message).
			continue
		}
		innerIdx := int(seqNum - rep.Interval.Min)
		innerIdxs = append(innerIdxs, innerIdx)
		encMsgs = append(encMsgs, msgsInRoot[innerIdx].Data)
		hash, err2 := hashLeaf.HashLeaf(msgsInRoot[innerIdx].ToGethLog())
		if err2 != nil {
			return nil, err2
		}
		hashes = append(hashes, hash)
	}
	merkleProof := tree.Prove(innerIdxs)
	// Double check this verifies before sending.
	res, err := commitStore.Verify(nil, hashes, merkleProof.Hashes, ProofFlagsToBits(merkleProof.SourceFlags))
	if err != nil {
		lggr.Errorw("Unable to call verify", "seqNums", finalSeqNums, "indices", innerIdxs, "root", rep.MerkleRoot[:], "seqRange", rep.Interval, "err", err)
		return nil, err
	}
	// No timestamp, means failed to verify root.
	if res.Cmp(big.NewInt(0)) == 0 {
		ir := tree.Root()
		lggr.Errorf("Root does not verify for messages: %v (indices %v) our inner root %x contract",
			finalSeqNums, innerIdxs, ir[:])
		return nil, errors.New("root does not verify")
	}
	return &MessageExecution{
		encMsgs:          encMsgs,
		proofs:           merkleProof.Hashes,
		proofSourceFlags: merkleProof.SourceFlags,
	}, nil
}
