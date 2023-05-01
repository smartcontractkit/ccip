package ccip

import (
	"context"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

const (
	MessageStateUntouched = iota
	MessageStateInProgress
	MessageStateSuccess
	MessageStateFailure
)

var ErrCommitStoreIsDown = errors.New("commitStore is down")

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
	seqNums          []uint64
	encMsgs          [][]byte
	tokenData        [][][]byte
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
	ctx context.Context,
	lggr logger.Logger,
	onRamp common.Address,
	eventSigs EventSignatures,
	seqParser func(logpoller.Log) (uint64, error),
	interval commit_store.CommitStoreInterval,
	srcLogPoller logpoller.LogPoller,
	hasher LeafHasherInterface[[32]byte],
	confs int,
) ([][32]byte, error) {
	// Logs are guaranteed to be in order of seq num, since these are finalized logs only
	// and the contract's seq num is auto-incrementing.
	logs, err := srcLogPoller.LogsDataWordRange(
		eventSigs.SendRequested,
		onRamp,
		eventSigs.SendRequestedSequenceNumberIndex,
		logpoller.EvmWord(interval.Min),
		logpoller.EvmWord(interval.Max),
		confs,
		pg.WithParentCtx(ctx))
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

func aggregateTokenValue(destTokenPricesUSD map[common.Address]*big.Int, srcToDst map[common.Address]common.Address, tokens []common.Address, amounts []*big.Int) (*big.Int, error) {
	sum := big.NewInt(0)
	for i := 0; i < len(tokens); i++ {
		price, ok := destTokenPricesUSD[srcToDst[tokens[i]]]
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

func commitReport(dstLogPoller logpoller.LogPoller, commitStore *commit_store.CommitStore, seqNr uint64) (commit_store.CommitStoreCommitReport, error) {
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
		reportAccepted, err := commitStore.ParseReportAccepted(log.GetGethLog())
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
	serviceCtx context.Context,
	lggr logger.Logger,
	source, dest logpoller.LogPoller,
	onRampAddress common.Address,
	observedMessages []ObservedMessage,
	commitStore *commit_store.CommitStore,
	confs int,
	eventSignatures EventSignatures,
	seqNumFromLog func(log logpoller.Log) (uint64, error),
	hashLeaf LeafHasherInterface[[32]byte],
) (*MessageExecution, error) {
	nextMin, err := commitStore.GetExpectedNextSequenceNumber(&bind.CallOpts{Context: serviceCtx})
	if err != nil {
		return nil, err
	}
	maxSeqNumInBatch := uint64(0)
	for _, seqNum := range observedMessages {
		if seqNum.SeqNr > maxSeqNumInBatch {
			maxSeqNumInBatch = seqNum.SeqNr
		}
	}

	if maxSeqNumInBatch >= nextMin {
		return nil, errors.Errorf("Cannot execute uncommitted seq num. nextMin %v, seqNums %v", nextMin, observedMessages)
	}
	rep, err := commitReport(dest, commitStore, observedMessages[0].SeqNr)
	if err != nil {
		return nil, err
	}
	lggr.Infow("Building execution report", "observations", observedMessages, "merkleRoot", hexutil.Encode(rep.MerkleRoot[:]), "report", rep)

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
	leaves, err := leavesFromIntervals(serviceCtx, lggr, onRampAddress, eventSignatures, seqNumFromLog, rep.Interval, source, hashLeaf, confs)
	if err != nil {
		return nil, err
	}
	ctx := hasher.NewKeccakCtx()
	tree, err := merklemulti.NewTree(ctx, leaves)
	if err != nil {
		return nil, err
	}

	var batch MessageExecution
	var innerIdxs []int
	var hashes [][32]byte
	for _, observedMessage := range observedMessages {
		if observedMessage.SeqNr < rep.Interval.Min || observedMessage.SeqNr > rep.Interval.Max {
			// We only return messages from a single root (the root of the first message).
			continue
		}
		innerIdx := int(observedMessage.SeqNr - rep.Interval.Min)

		batch.seqNums = append(batch.seqNums, observedMessage.SeqNr)
		batch.encMsgs = append(batch.encMsgs, msgsInRoot[innerIdx].Data)
		batch.tokenData = append(batch.tokenData, observedMessage.TokenData)

		innerIdxs = append(innerIdxs, innerIdx)
		hash, err2 := hashLeaf.HashLeaf(msgsInRoot[innerIdx].ToGethLog())
		if err2 != nil {
			return nil, err2
		}
		hashes = append(hashes, hash)
	}
	merkleProof := tree.Prove(innerIdxs)
	// Double check this verifies before sending.
	res, err := commitStore.Verify(&bind.CallOpts{Context: serviceCtx}, hashes, merkleProof.Hashes, ProofFlagsToBits(merkleProof.SourceFlags))
	if err != nil {
		lggr.Errorw("Unable to call verify", "observations", observedMessages, "indices", innerIdxs, "root", rep.MerkleRoot[:], "seqRange", rep.Interval, "err", err)
		return nil, err
	}
	// No timestamp, means failed to verify root.
	if res.Cmp(big.NewInt(0)) == 0 {
		ir := tree.Root()
		lggr.Errorf("Root does not verify for messages: %v (indices %v) our inner root %x contract",
			observedMessages, innerIdxs, ir[:])
		return nil, errors.New("root does not verify")
	}

	batch.proofs = merkleProof.Hashes
	batch.proofSourceFlags = merkleProof.SourceFlags

	return &batch, nil
}

func isCommitStoreDownNow(ctx context.Context, lggr logger.Logger, commitStore *commit_store.CommitStore) bool {
	paused, err := commitStore.Paused(&bind.CallOpts{Context: ctx})
	if err != nil {
		// Air on side of caution by halting if we cannot read the state?
		lggr.Errorw("Unable to read CommitStore paused", "err", err)
		return true
	}
	healthy, err := commitStore.IsAFNHealthy(&bind.CallOpts{Context: ctx})
	if err != nil {
		lggr.Errorw("Unable to read CommitStore AFN state", "err", err)
		return true
	}
	return paused || !healthy
}
