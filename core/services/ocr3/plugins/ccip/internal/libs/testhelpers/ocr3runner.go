package testhelpers

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"

	"github.com/smartcontractkit/ccipocr3/internal/libs/slicelib"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

var (
	ErrQuery                        = errors.New("error in query phase")
	ErrObservation                  = errors.New("error in observation phase")
	ErrValidateObservation          = errors.New("error in validate observation phase")
	ErrOutcome                      = errors.New("error in outcome phase")
	ErrReports                      = errors.New("error in reports phase")
	ErrShouldAcceptAttestedReport   = errors.New("error in should accept attested report phase")
	ErrShouldTransmitAcceptedReport = errors.New("error in should transmit accepted report phase")
)

// OCR3Runner is a simple runner for OCR3.
//
// TODO: move to a shared repository.
type OCR3Runner[RI any] struct {
	nodes           []ocr3types.ReportingPlugin[RI]
	round           int
	previousOutcome ocr3types.Outcome
}

func NewOCR3Runner[RI any](nodes []ocr3types.ReportingPlugin[RI]) *OCR3Runner[RI] {
	return &OCR3Runner[RI]{
		nodes: nodes,
		round: 0,
	}
}

// RunRound will run some basic steps of an OCR3 flow.
// This is not a full OCR3 round but only the bare minimum.
func (r *OCR3Runner[RI]) RunRound(ctx context.Context) (result RoundResult[RI], err error) {
	r.round++
	seqNr := uint64(r.round)

	leaderNode := r.selectLeader()

	outcomeCtx := ocr3types.OutcomeContext{SeqNr: seqNr, PreviousOutcome: r.previousOutcome}

	q, err := leaderNode.Query(ctx, outcomeCtx)
	if err != nil {
		return RoundResult[RI]{}, fmt.Errorf("%s: %w", err, ErrQuery)
	}

	attributedObservations := make([]types.AttributedObservation, len(r.nodes))
	for i, n := range r.nodes {
		obs, err2 := n.Observation(ctx, outcomeCtx, q)
		if err2 != nil {
			return RoundResult[RI]{}, fmt.Errorf("%s: %w", err2, ErrObservation)
		}

		attrObs := types.AttributedObservation{Observation: obs, Observer: commontypes.OracleID(i)}
		err = leaderNode.ValidateObservation(outcomeCtx, q, attrObs)
		if err != nil {
			return RoundResult[RI]{}, fmt.Errorf("%s: %w", err, ErrValidateObservation)
		}

		attributedObservations[i] = attrObs
	}

	outcomes := make([]ocr3types.Outcome, len(r.nodes))
	for i, n := range r.nodes {
		outcome, err2 := n.Outcome(outcomeCtx, q, attributedObservations)
		if err2 != nil {
			return RoundResult[RI]{}, fmt.Errorf("%s: %w", err2, ErrOutcome)
		}

		outcomes[i] = outcome
	}

	// check that all the outcomes are the same.
	if countUniqueOutcomes(outcomes) > 1 {
		return RoundResult[RI]{}, fmt.Errorf("outcomes are not equal")
	}

	r.previousOutcome = outcomes[0]

	allReports := make([][]ocr3types.ReportWithInfo[RI], len(r.nodes))
	for i, n := range r.nodes {
		reportsWithInfo, err2 := n.Reports(seqNr, outcomes[0])
		if err2 != nil {
			return RoundResult[RI]{}, fmt.Errorf("%s: %w", err2, ErrReports)
		}

		allReports[i] = reportsWithInfo
	}

	// check that all the reports are the same.
	if countUniqueReports(slicelib.Flatten(allReports)) > 1 {
		return RoundResult[RI]{}, fmt.Errorf("reports are not equal")
	}

	transmitted := make([]ocr3types.ReportWithInfo[RI], 0)
	notAccepted := make([]ocr3types.ReportWithInfo[RI], 0)
	notTransmitted := make([]ocr3types.ReportWithInfo[RI], 0)

	for _, report := range allReports[0] {
		allShouldAccept := make([]bool, len(r.nodes))
		for i, n := range r.nodes {
			shouldAccept, err2 := n.ShouldAcceptAttestedReport(ctx, seqNr, report)
			if err2 != nil {
				return RoundResult[RI]{}, fmt.Errorf("%s: %w", err2, ErrShouldAcceptAttestedReport)
			}

			allShouldAccept[i] = shouldAccept
		}
		if slicelib.CountUnique(allShouldAccept) > 1 {
			return RoundResult[RI]{}, fmt.Errorf("should accept attested report from all oracles is not equal")
		}

		if !allShouldAccept[0] {
			notAccepted = append(notAccepted, report)
			continue
		}

		allShouldTransmit := make([]bool, len(r.nodes))
		for i, n := range r.nodes {
			shouldTransmit, err2 := n.ShouldTransmitAcceptedReport(ctx, seqNr, report)
			if err2 != nil {
				return RoundResult[RI]{}, fmt.Errorf("%s: %w", err2, ErrShouldTransmitAcceptedReport)
			}

			allShouldTransmit[i] = shouldTransmit
		}
		if slicelib.CountUnique(allShouldTransmit) > 1 {
			return RoundResult[RI]{}, fmt.Errorf("should transmit accepted report from all oracles is not equal")
		}

		if !allShouldTransmit[0] {
			notTransmitted = append(notTransmitted, report)
			continue
		}

		transmitted = append(transmitted, report)
	}

	return RoundResult[RI]{
		Transmitted:    transmitted,
		NotAccepted:    notAccepted,
		NotTransmitted: notTransmitted,
		Outcome:        outcomes[0],
	}, nil
}

func (r *OCR3Runner[RI]) selectLeader() ocr3types.ReportingPlugin[RI] {
	numNodes := len(r.nodes)
	if numNodes == 0 {
		return nil
	}
	return r.nodes[rand.Intn(numNodes)]
}

type RoundResult[RI any] struct {
	Transmitted    []ocr3types.ReportWithInfo[RI]
	NotAccepted    []ocr3types.ReportWithInfo[RI]
	NotTransmitted []ocr3types.ReportWithInfo[RI]
	Outcome        []byte
}

func countUniqueOutcomes(outcomes []ocr3types.Outcome) int {
	flattenedHashes := make([]string, 0, len(outcomes))
	for _, o := range outcomes {
		h := sha1.New()
		h.Write(o)
		flattenedHashes = append(flattenedHashes, hex.EncodeToString(h.Sum(nil)))
	}
	return slicelib.CountUnique(flattenedHashes)
}

func countUniqueReports[RI any](reports []ocr3types.ReportWithInfo[RI]) int {
	flattenedHashes := make([]string, 0, len(reports))
	for _, report := range reports {
		h := sha1.New()
		h.Write(report.Report)
		flattenedHashes = append(flattenedHashes, hex.EncodeToString(h.Sum(nil)))
	}
	return slicelib.CountUnique(flattenedHashes)
}
