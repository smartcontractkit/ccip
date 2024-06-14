package evm

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	pkgerrors "github.com/pkg/errors"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"

	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"

	txmgrcommon "github.com/smartcontractkit/chainlink/v2/common/txmgr"
	txm "github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/ccipcommit"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/ccipexec"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	cciptransmitter "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/transmitter"
	statuschecker "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/statuschecker"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

// CCIPCommitProvider provides all components needed for a CCIP Relay OCR2 plugin.
type CCIPCommitProvider interface {
	commontypes.Plugin
}

// CCIPExecutionProvider provides all components needed for a CCIP Execution OCR2 plugin.
type CCIPExecutionProvider interface {
	commontypes.Plugin
}

type ccipCommitProvider struct {
	*configWatcher
	contractTransmitter *contractTransmitter
}

func chainToUUID(chainID *big.Int) uuid.UUID {
	// See https://www.rfc-editor.org/rfc/rfc4122.html#section-4.1.3 for the list of supported versions.
	const VersionSHA1 = 5
	var buf bytes.Buffer
	buf.WriteString("CCIP:")
	buf.Write(chainID.Bytes())
	// We use SHA-256 instead of SHA-1 because the former has better collision resistance.
	// The UUID will contain only the first 16 bytes of the hash.
	// You can't say which algorithms was used just by looking at the UUID bytes.
	return uuid.NewHash(sha256.New(), uuid.NameSpaceOID, buf.Bytes(), VersionSHA1)
}

func NewCCIPCommitProvider(ctx context.Context, lggr logger.Logger, chainSet legacyevm.Chain, rargs commontypes.RelayArgs, transmitterID string, ks keystore.Eth) (CCIPCommitProvider, error) {
	relayOpts := types.NewRelayOpts(rargs)
	configWatcher, err := newStandardConfigProvider(ctx, lggr, chainSet, relayOpts)
	if err != nil {
		return nil, err
	}
	address := common.HexToAddress(relayOpts.ContractID)
	typ, ver, err := ccipconfig.TypeAndVersion(address, chainSet.Client())
	if err != nil {
		return nil, err
	}
	fn, err := ccipcommit.CommitReportToEthTxMeta(typ, ver)
	if err != nil {
		return nil, err
	}
	subjectID := chainToUUID(configWatcher.chain.ID())
	contractTransmitter, err := newCCIPOnChainContractTransmitter(ctx, lggr, rargs, transmitterID, ks, configWatcher, configTransmitterOpts{
		subjectID: &subjectID,
	}, OCR2AggregatorTransmissionContractABI, fn, 0)
	if err != nil {
		return nil, err
	}
	return &ccipCommitProvider{
		configWatcher:       configWatcher,
		contractTransmitter: contractTransmitter,
	}, nil
}

func (c *ccipCommitProvider) ContractTransmitter() ocrtypes.ContractTransmitter {
	return c.contractTransmitter
}

func (c *ccipCommitProvider) ChainReader() commontypes.ContractReader {
	return nil
}

func (c *ccipCommitProvider) Codec() commontypes.Codec {
	return nil
}

type ccipExecutionProvider struct {
	*configWatcher
	contractTransmitter *contractTransmitter
}

var _ commontypes.Plugin = (*ccipExecutionProvider)(nil)

func NewCCIPExecutionProvider(ctx context.Context, lggr logger.Logger, chainSet legacyevm.Chain, rargs commontypes.RelayArgs, transmitterID string, ks keystore.Eth) (CCIPExecutionProvider, error) {
	relayOpts := types.NewRelayOpts(rargs)

	configWatcher, err := newStandardConfigProvider(ctx, lggr, chainSet, relayOpts)
	if err != nil {
		return nil, err
	}
	address := common.HexToAddress(relayOpts.ContractID)
	typ, ver, err := ccipconfig.TypeAndVersion(address, chainSet.Client())
	if err != nil {
		return nil, err
	}
	fn, err := ccipexec.ExecReportToEthTxMeta(ctx, typ, ver)
	if err != nil {
		return nil, err
	}
	subjectID := chainToUUID(configWatcher.chain.ID())
	contractTransmitter, err := newCCIPOnChainContractTransmitter(ctx, lggr, rargs, transmitterID, ks, configWatcher, configTransmitterOpts{
		subjectID: &subjectID,
	}, OCR2AggregatorTransmissionContractABI, fn, 0)
	if err != nil {
		return nil, err
	}
	return &ccipExecutionProvider{
		configWatcher:       configWatcher,
		contractTransmitter: contractTransmitter,
	}, nil
}

func newCCIPOnChainContractTransmitter(ctx context.Context, lggr logger.Logger, rargs commontypes.RelayArgs, transmitterID string, ethKeystore keystore.Eth, configWatcher *configWatcher, opts configTransmitterOpts, transmissionContractABI abi.ABI, reportToEvmTxMeta ReportToEthMetadata, transmissionContractRetention time.Duration) (*contractTransmitter, error) {
	var relayConfig types.RelayConfig
	if err := json.Unmarshal(rargs.RelayConfig, &relayConfig); err != nil {
		return nil, err
	}
	var fromAddresses []common.Address
	sendingKeys := relayConfig.SendingKeys
	if !relayConfig.EffectiveTransmitterID.Valid {
		return nil, pkgerrors.New("EffectiveTransmitterID must be specified")
	}
	effectiveTransmitterAddress := common.HexToAddress(relayConfig.EffectiveTransmitterID.String)

	sendingKeysLength := len(sendingKeys)
	if sendingKeysLength == 0 {
		return nil, pkgerrors.New("no sending keys provided")
	}

	// If we are using multiple sending keys, then a forwarder is needed to rotate transmissions.
	// Ensure that this forwarder is not set to a local sending key, and ensure our sending keys are enabled.
	for _, s := range sendingKeys {
		if sendingKeysLength > 1 && s == effectiveTransmitterAddress.String() {
			return nil, pkgerrors.New("the transmitter is a local sending key with transaction forwarding enabled")
		}
		if err := ethKeystore.CheckEnabled(ctx, common.HexToAddress(s), configWatcher.chain.Config().EVM().ChainID()); err != nil {
			return nil, pkgerrors.Wrap(err, "one of the sending keys given is not enabled")
		}
		fromAddresses = append(fromAddresses, common.HexToAddress(s))
	}

	subject := rargs.ExternalJobID
	if opts.subjectID != nil {
		subject = *opts.subjectID
	}
	strategy := txmgrcommon.NewQueueingTxStrategy(subject, relayConfig.DefaultTransactionQueueDepth)

	var checker txm.TransmitCheckerSpec
	if relayConfig.SimulateTransactions {
		checker.CheckerType = txm.TransmitCheckerTypeSimulate
	}

	gasLimit := configWatcher.chain.Config().EVM().GasEstimator().LimitDefault()
	ocr2Limit := configWatcher.chain.Config().EVM().GasEstimator().LimitJobType().OCR2()
	if ocr2Limit != nil {
		gasLimit = uint64(*ocr2Limit)
	}
	if opts.pluginGasLimit != nil {
		gasLimit = uint64(*opts.pluginGasLimit)
	}

	var transmitter Transmitter
	var err error

	switch commontypes.OCR2PluginType(rargs.ProviderType) {
	case commontypes.CCIPExecution:
		transmitter, err = cciptransmitter.NewTransmitterWithStatusChecker(
			configWatcher.chain.TxManager(),
			fromAddresses,
			gasLimit,
			effectiveTransmitterAddress,
			strategy,
			checker,
			configWatcher.chain.ID(),
			ethKeystore,
			statuschecker.NewTxmStatusChecker(nil), // TODO: remove after TXM changes are merged
			// statuschecker.NewTransactionStatusChecker(configWatcher.chain.TxManager()), // TODO: uncomment after TXM changes are merged
		)
	default:
		transmitter, err = cciptransmitter.NewTransmitter(
			configWatcher.chain.TxManager(),
			fromAddresses,
			gasLimit,
			effectiveTransmitterAddress,
			strategy,
			checker,
			configWatcher.chain.ID(),
			ethKeystore,
		)
	}
	if err != nil {
		return nil, pkgerrors.Wrap(err, "failed to create transmitter")
	}

	return NewOCRContractTransmitterWithRetention(
		ctx,
		configWatcher.contractAddress,
		configWatcher.chain.Client(),
		transmissionContractABI,
		transmitter,
		configWatcher.chain.LogPoller(),
		lggr,
		reportToEvmTxMeta,
		transmissionContractRetention,
	)
}

func (c *ccipExecutionProvider) ContractTransmitter() ocrtypes.ContractTransmitter {
	return c.contractTransmitter
}

func (c *ccipExecutionProvider) ChainReader() commontypes.ContractReader {
	return nil
}

func (c *ccipExecutionProvider) Codec() commontypes.Codec {
	return nil
}
