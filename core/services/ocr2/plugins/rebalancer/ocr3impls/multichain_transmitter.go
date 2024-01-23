package ocr3impls

import (
	"context"
	"fmt"
	"strconv"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay"
)

type MultichainMeta interface {
	GetDestinationChain() relay.ID
	GetDestinationConfigDigest() types.ConfigDigest
}

// multichainTransmitterOCR3 is a transmitter that can transmit to multiple chains.
// It uses the information in the MultichainMeta to determine which chain to transmit to.
// Note that this would only work with the appropriate multi-chain config tracker implementation.
type multichainTransmitterOCR3[RI MultichainMeta] struct {
	transmitters map[relay.ID]ocr3types.ContractTransmitter[RI]
	lggr         logger.Logger
}

func NewMultichainTransmitterOCR3[RI MultichainMeta](
	transmitters map[relay.ID]ocr3types.ContractTransmitter[RI],
	lggr logger.Logger,
) (*multichainTransmitterOCR3[RI], error) {
	return &multichainTransmitterOCR3[RI]{
		transmitters: transmitters,
		lggr:         lggr,
	}, nil
}

// FromAccount implements ocr3types.ContractTransmitter.
func (m *multichainTransmitterOCR3[RI]) FromAccount() (types.Account, error) {
	var accounts []string
	for relayID, t := range m.transmitters {
		account, err := t.FromAccount()
		if err != nil {
			return "", err
		}
		accounts = append(accounts, EncodeTransmitter(relayID, account))
	}
	return types.Account(JoinTransmitters(accounts)), nil
}

// Transmit implements ocr3types.ContractTransmitter.
func (m *multichainTransmitterOCR3[RI]) Transmit(ctx context.Context, configDigest types.ConfigDigest, seqNr uint64, rwi ocr3types.ReportWithInfo[RI], sigs []types.AttributedOnchainSignature) error {
	destChain := rwi.Info.GetDestinationChain()

	// get transmitter by evm chain id
	transmitter, ok := m.transmitters[destChain]

	// get transmitter by chain selector
	chainSel, err := strconv.ParseUint(destChain.ChainID, 10, 64)
	if err == nil {
		ch, exists := chainsel.ChainBySelector(chainSel)
		if exists {
			transmitter, ok = m.transmitters[relay.NewID(relay.EVM, strconv.FormatUint(ch.EvmChainID, 10))]
		}
	}

	if !ok {
		return fmt.Errorf("no transmitter for chain %s", destChain)
	}
	m.lggr.Infow("multichain transmitter: transmitting to chain", "destChain", destChain.String(), "configDigest", rwi.Info.GetDestinationConfigDigest().Hex())
	return transmitter.Transmit(ctx, rwi.Info.GetDestinationConfigDigest(), seqNr, rwi, sigs)
}
