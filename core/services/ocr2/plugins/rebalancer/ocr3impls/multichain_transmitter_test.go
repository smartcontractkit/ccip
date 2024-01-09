package ocr3impls_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/ocr3impls"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/require"
)

func TestMultichainTransmitter(t *testing.T) {
	t.Parallel()
	// create many separate transmitters and separate chains
	numChains := 4
	unis := make(map[int]testUniverse[multichainMeta])
	for i := 0; i < numChains; i++ {
		unis[i] = newTestUniverse[multichainMeta](t, nil)
	}

	mct, err := ocr3impls.NewMultichainTransmitterOCR3[multichainMeta](
		map[relay.ID]ocr3types.ContractTransmitter[multichainMeta]{
			relay.NewID(relay.EVM, "0"): unis[0].ocr3Transmitter,
			relay.NewID(relay.EVM, "1"): unis[1].ocr3Transmitter,
			relay.NewID(relay.EVM, "2"): unis[2].ocr3Transmitter,
			relay.NewID(relay.EVM, "3"): unis[3].ocr3Transmitter,
		},
		logger.TestLogger(t),
	)
	require.NoError(t, err)

	expectedTransmitters := []string{
		unis[0].transmitters[0].From.String(),
		unis[1].transmitters[0].From.String(),
		unis[2].transmitters[0].From.String(),
		unis[3].transmitters[0].From.String(),
	}
	slices.Sort(expectedTransmitters)
	expectedFromAccount := strings.Join(expectedTransmitters, ",")
	fromAccount, err := mct.FromAccount()
	require.NoError(t, err)
	require.Equal(t, expectedFromAccount, string(fromAccount))

	var configDigests []ocrtypes.ConfigDigest
	for _, uni := range unis {
		c, err2 := uni.wrapper.LatestConfigDigestAndEpoch(nil)
		require.NoError(t, err2)
		configDigests = append(configDigests, c.ConfigDigest)
	}

	// generate a report for each chain and sign it
	// note that in this test each chain has a different set of signers
	// this is okay because it's just a test
	// in actuality the same signers will be used across all chains
	var reports []ocr3types.ReportWithInfo[multichainMeta]
	for i := 0; i < numChains; i++ {
		report := ocr3types.ReportWithInfo[multichainMeta]{
			Info:   multichainMeta{destChainIndex: i},
			Report: []byte{},
		}
		reports = append(reports, report)
	}
	seqNum := uint64(1)
	for i := range reports {
		attributedSigs := unis[i].SignReport(t, configDigests[i], reports[i], seqNum)
		err = mct.Transmit(testutils.Context(t), configDigests[i], seqNum, reports[i], attributedSigs)
		require.NoError(t, err)
		// TODO: for some reason this event isn't being emitted in the simulated backend
		// events := unis[i].TransmittedEvents(t)
		// require.Len(t, events, 1)
		// require.Equal(t, configDigests[i], events[0].ConfigDigest, "config digest mismatch")
		// require.Equal(t, seqNum, events[0].SequenceNumber, "sequence number mismatch")
		// increment sequence number so that each chain gets a unique one for this test
		seqNum++
	}
}

type multichainMeta struct {
	destChainIndex int
}

func (m multichainMeta) GetDestinationChain() relay.ID {
	return relay.NewID(relay.EVM, fmt.Sprintf("%d", m.destChainIndex))
}
