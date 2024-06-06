package validation

import (
	"fmt"

	"golang.org/x/crypto/sha3"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type reportCounter struct {
	data  cciptypes.ExecutePluginCommitData
	count int
}

type CommitReportValidator interface {
	AddReport(data cciptypes.ExecutePluginCommitData) error
	GetValidatedReports() ([]cciptypes.ExecutePluginCommitData, error)
}

// commitValidator is a helper object to validate reports for a single chain.
// It keeps track of all reports and determines if they observations are consistent
// with one another and whether they meet the required fChain threshold.
type commitReportValidator struct {
	minObservation int
	cache          map[cciptypes.Bytes32]*reportCounter
}

func NewCommitReportValidator(min int) *commitReportValidator {
	return &commitReportValidator{
		minObservation: min,
		cache:          make(map[cciptypes.Bytes32]*reportCounter),
	}
}

func (cv *commitReportValidator) AddReport(data cciptypes.ExecutePluginCommitData) error {
	//id := sha3.Sum256(data.ToBytes())
	id := sha3.Sum256([]byte(fmt.Sprintf("%v", data)))
	if _, ok := cv.cache[id]; ok {
		cv.cache[id].count++
	} else {
		cv.cache[id] = &reportCounter{data: data, count: 1}
	}
	return nil
}

func (cv *commitReportValidator) GetValidatedReports() ([]cciptypes.ExecutePluginCommitData, error) {
	var validated []cciptypes.ExecutePluginCommitData
	for _, rc := range cv.cache {
		if rc.count >= cv.minObservation {
			validated = append(validated, rc.data)
		}
	}
	return validated, nil
}
