package ccip

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/merklemulti"
)

func buildExecutionReport(
	ctx context.Context,
	lggr logger.Logger,
	destLP logpoller.LogPoller,
	observedMessages []ObservedMessage,
	commitStore commit_store.CommitStoreInterface,
	seqNumFromLog func(log logpoller.Log) (uint64, error),
	hashLeaf hasher.LeafHasherInterface[[32]byte],
	getMsgLogs func(min uint64, max uint64) ([]logpoller.Log, error),
) (evm_2_evm_offramp.InternalExecutionReport, error) {
	if err := validateSeqNumbers(ctx, commitStore, observedMessages); err != nil {
		return evm_2_evm_offramp.InternalExecutionReport{}, err
	}
	commitReport, err := getCommitReportForSeqNum(destLP, commitStore, observedMessages[0].SeqNr)
	if err != nil {
		return evm_2_evm_offramp.InternalExecutionReport{}, err
	}
	lggr.Infow("Building execution report", "observations", observedMessages, "merkleRoot", hexutil.Encode(commitReport.MerkleRoot[:]), "report", commitReport)

	msgsInRoot, err := getMsgLogs(commitReport.Interval.Min, commitReport.Interval.Max)
	if err != nil {
		return evm_2_evm_offramp.InternalExecutionReport{}, err
	}
	if len(msgsInRoot) != int(commitReport.Interval.Max-commitReport.Interval.Min+1) {
		return evm_2_evm_offramp.InternalExecutionReport{}, errors.Errorf("unexpected missing msgs in committed root %x have %d want %d", commitReport.MerkleRoot, len(msgsInRoot), int(commitReport.Interval.Max-commitReport.Interval.Min+1))
	}
	leaves, err := leavesFromIntervals(lggr, seqNumFromLog, commitReport.Interval, hashLeaf, msgsInRoot)
	if err != nil {
		return evm_2_evm_offramp.InternalExecutionReport{}, err
	}
	tree, err := merklemulti.NewTree(hasher.NewKeccakCtx(), leaves)
	if err != nil {
		return evm_2_evm_offramp.InternalExecutionReport{}, err
	}

	var execReport evm_2_evm_offramp.InternalExecutionReport
	var innerIdxs []int
	var hashes [][32]byte
	for _, observedMessage := range observedMessages {
		if observedMessage.SeqNr < commitReport.Interval.Min || observedMessage.SeqNr > commitReport.Interval.Max {
			// We only return messages from a single root (the root of the first message).
			continue
		}
		innerIdx := int(observedMessage.SeqNr - commitReport.Interval.Min)

		execReport.SequenceNumbers = append(execReport.SequenceNumbers, observedMessage.SeqNr)
		execReport.EncodedMessages = append(execReport.EncodedMessages, msgsInRoot[innerIdx].Data)
		execReport.OffchainTokenData = append(execReport.OffchainTokenData, observedMessage.TokenData)

		innerIdxs = append(innerIdxs, innerIdx)
		hash, err2 := hashLeaf.HashLeaf(msgsInRoot[innerIdx].ToGethLog())
		if err2 != nil {
			return execReport, err2
		}
		hashes = append(hashes, hash)
	}
	merkleProof := tree.Prove(innerIdxs)
	// Double check this verifies before sending.
	res, err := commitStore.Verify(&bind.CallOpts{Context: ctx}, hashes, merkleProof.Hashes, abihelpers.ProofFlagsToBits(merkleProof.SourceFlags))
	if err != nil {
		lggr.Errorw("Unable to call verify", "observations", observedMessages, "indices", innerIdxs, "root", commitReport.MerkleRoot[:], "seqRange", commitReport.Interval, "err", err)
		return execReport, err
	}
	// No timestamp, means failed to verify root.
	if res.Cmp(big.NewInt(0)) == 0 {
		root := tree.Root()
		lggr.Errorf("Root does not verify for messages: %v (indices %v) our inner root %x contract",
			observedMessages, innerIdxs, root[:])
		return execReport, errors.New("root does not verify")
	}

	execReport.Proofs = merkleProof.Hashes
	execReport.ProofFlagBits = abihelpers.ProofFlagsToBits(merkleProof.SourceFlags)

	return execReport, nil
}

// Validates the given message observations do not exceed the committed sequence numbers
// in the commitStore.
func validateSeqNumbers(serviceCtx context.Context, commitStore commit_store.CommitStoreInterface, observedMessages []ObservedMessage) error {
	nextMin, err := commitStore.GetExpectedNextSequenceNumber(&bind.CallOpts{Context: serviceCtx})
	if err != nil {
		return err
	}
	// observedMessages are always sorted by SeqNr and never empty, so it's safe to take last element
	maxSeqNumInBatch := observedMessages[len(observedMessages)-1].SeqNr

	if maxSeqNumInBatch >= nextMin {
		return errors.Errorf("Cannot execute uncommitted seq num. nextMin %v, seqNums %v", nextMin, observedMessages)
	}
	return nil
}

// Gets the commit report from the saved logs for a given sequence number.
func getCommitReportForSeqNum(dstLogPoller logpoller.LogPoller, commitStore commit_store.CommitStoreInterface, seqNr uint64) (commit_store.CommitStoreCommitReport, error) {
	// fetch commitReports which report.Interval.Max >= seqNr
	logs, err := dstLogPoller.LogsDataWordGreaterThan(
		abihelpers.EventSignatures.ReportAccepted,
		commitStore.Address(),
		abihelpers.EventSignatures.ReportAcceptedMaxSequenceNumberWord,
		logpoller.EvmWord(seqNr),
		0,
	)
	if err != nil {
		return commit_store.CommitStoreCommitReport{}, err
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
