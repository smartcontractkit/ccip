package ccip

import (
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

type executionState struct {
	executed  bool
	finalized bool
}

type internalCommitStoreReport struct {
	blockTimestamp time.Time
	blockNumber    uint64
	seqNrMin       uint64
	seqNrMax       uint64
}

type CommitStoreRoot struct {
	Root     common.Hash
	MinSeqNr uint64
	MaxSeqNr uint64
	Executed map[uint64]bool
}

type CommitReportsStash struct {
	permissionlessExecThreshold time.Duration
	lastFinalizedBlockNumber    uint64
	commitReportsRoots          []common.Hash
	commitStoreReports          map[common.Hash]internalCommitStoreReport
}

// RootsEligibleForExecution Usage in the exec plugin
//
//	 func NewReportingPlugin() {
//			p.commitRoots := CommitReportsStash{}
//	 }
//
//		func Observation() {
//			roots := p.commitRoots.RootsEligibleForExecution()
//		 	// some initial processing
//		 	for _, root := range roots {
//		 		rootWithMgs := FetchMessages(root)
//					return buildBatch(rootWithMsgs)
//		 	}
//		}
func (c *CommitReportsStash) RootsEligibleForExecution() []CommitStoreRoot {
	// 1. Fetch only the logs after the last finalized block number
	newLogs, lastFinalizedBlock := c.fetchNewLogs()
	// 2. Add logs to the internal state
	c.updateStateWithNewLogs(newLogs, lastFinalizedBlock)
	// 3. Fetch state of the execution
	seqNrs := c.fetchExecutedSeqNrs(lastFinalizedBlock)
	// 4. Remove fully executed roots
	c.removeFullyExecutedRoots(seqNrs)
	// 5. Return roots eligible for execution
	return c.externalRepresentationRoots(seqNrs)
}

func (c *CommitReportsStash) IsSnoozed(hash common.Hash) bool {
	return false
}

func (c *CommitReportsStash) Snooze(hash common.Hash) {

}

func (c *CommitReportsStash) fetchExecutedSeqNrs(lastFinalizedBlock uint64) map[uint64]executionState {
	oldestRoot, newestRoot := c.commitReportsRoots[0], c.commitReportsRoots[len(c.commitReportsRoots)-1]
	minSeqNr, maxSeqNr := c.commitStoreReports[oldestRoot].seqNrMin, c.commitStoreReports[newestRoot].seqNrMax
	executedSeqNrs := fetchExecutedSeqNrs(minSeqNr, maxSeqNr)

	seqNrs := map[uint64]executionState{}
	for _, seqNr := range executedSeqNrs {
		seqNrs[seqNr] = executionState{
			executed:  true,
			finalized: false, // fixme
		}
	}

	return seqNrs
}

func (c *CommitReportsStash) fetchNewLogs() ([]ccipdata.CommitStoreReport, uint64) {
	if c.lastFinalizedBlockNumber == 0 {
		// this is initial fetch, might require paging
		// select * from evm.logs where block_timestamp > now() - permissionless_execution_threshold
	} else {
		// This will fetch rather small number of logs
		// select * from evm.logs where block_number > last_finalized_block_number
	}
	return []ccipdata.CommitStoreReport{}, 0
}

func (c *CommitReportsStash) updateStateWithNewLogs(logs []ccipdata.CommitStoreReport, lastFinalizedBlock uint64) {
	// Handle eviction here
	//i := 0
	//for j, root := range c.commitReportsRoots {
	//	 if c.commitStoreReports[root].blockTimestamp < time.Now().Add(-c.permissionlessExecThreshold) {
	//		 i++
	//	 } else {
	//		 break
	//	 }
	//}
	// c.commitReportsRoots = c.commitReportsRoots[i:]

	for _, log := range logs {
		c.commitReportsRoots = append(c.commitReportsRoots, log.MerkleRoot)
		c.commitStoreReports[log.MerkleRoot] = internalCommitStoreReport{
			seqNrMin:    log.Interval.Min,
			seqNrMax:    log.Interval.Max,
			blockNumber: 0,
		}
	}
	c.lastFinalizedBlockNumber = lastFinalizedBlock
}

func (c *CommitReportsStash) removeFullyExecutedRoots(seqNrs map[uint64]executionState) {
	for _, seqNr := range seqNrs {
		// This is similar to snooze forever
		if seqNr.executed && seqNr.finalized {
			// find commit root based on the seqNr
			commitRoot := common.Hash{}
			delete(c.commitStoreReports, commitRoot)
			// find commit root and remove it from slice
			c.commitReportsRoots = c.commitReportsRoots // (- commitRoot)
		}
	}
}

func (c *CommitReportsStash) externalRepresentationRoots(seqNrs map[uint64]executionState) []CommitStoreRoot {
	result := make([]CommitStoreRoot, 0, len(c.commitReportsRoots))
	for _, root := range c.commitReportsRoots {
		report := c.commitStoreReports[root]
		states := make(map[uint64]bool)

		for i := report.seqNrMin; i <= report.seqNrMax; i++ {
			states[i] = seqNrs[i].executed
		}

		result = append(result, CommitStoreRoot{
			Root:     root,
			MinSeqNr: report.seqNrMin,
			MaxSeqNr: report.seqNrMax,
			Executed: states,
		})
	}
	return result
}

func fetchExecutedSeqNrs(min uint64, max2 uint64) []uint64 {
	//logs := `SELECT topics[$4], block_number FROM evm.logs
	//				WHERE evm.logs.evm_chain_id = $1
	//				AND address = $2 AND event_sig = $3
	//				AND topics[$4] >= $5
	//				AND topics[$4] <= $6
	//              AND topics[$4] not in (executed and finalized) -- improvement, either passed from memory or as a nested query
	//				AND block_number <= confs`
	return []uint64{}
}
