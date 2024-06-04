package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata/usdc"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"math/big"
	"net/url"
	"time"
)

type SrcExecProvider struct {
	lggr                                   logger.Logger
	versionFinder                          ccip.VersionFinder
	client                                 client.Client
	lp                                     logpoller.LogPoller
	startBlock                             uint64
	contractTransmitter                    *contractTransmitter
	configWatcher                          *configWatcher
	gasEstimator                           gas.EvmFeeEstimator
	maxGasPrice                            big.Int
	jobID                                  string
	usdcReader                             ccip.USDCReaderImpl
	usdcAttestationAPI                     string
	usdcAttestationAPITimeoutSeconds       int
	usdcAttestationAPIIntervalMilliseconds int
	usdcSrcMsgTransmitterAddr              common.Address
}

func NewSrcExecProvider(
	lggr logger.Logger,
	versionFinder ccip.VersionFinder,
	client client.Client,
	lp logpoller.LogPoller,
	startBlock uint64,
	contractTransmitter *contractTransmitter,
	configWatcher *configWatcher,
	jobID string,
	usdcAttestationAPI string,
	usdcAttestationAPITimeoutSeconds int,
	usdcAttestationAPIIntervalMilliseconds int,
	usdcSrcMsgTransmitterAddr common.Address,
	// gasEstimator gas.EvmFeeEstimator,
	// maxGasPrice big.Int,
) (commontypes.CCIPExecProvider, error) {

	usdcReader, err := ccip.NewUSDCReader(lggr, jobID, usdcSrcMsgTransmitterAddr, lp, true)
	if err != nil {
		return nil, errors.Wrap(err, "new usdc reader")
	}

	return &SrcExecProvider{
		lggr:                                   lggr,
		versionFinder:                          versionFinder,
		client:                                 client,
		lp:                                     lp,
		startBlock:                             startBlock,
		contractTransmitter:                    contractTransmitter,
		configWatcher:                          configWatcher,
		usdcReader:                             *usdcReader,
		usdcAttestationAPI:                     usdcAttestationAPI,
		usdcAttestationAPITimeoutSeconds:       usdcAttestationAPITimeoutSeconds,
		usdcAttestationAPIIntervalMilliseconds: usdcAttestationAPIIntervalMilliseconds,
		usdcSrcMsgTransmitterAddr:              usdcSrcMsgTransmitterAddr,
		//gasEstimator:  gasEstimator,
		//maxGasPrice:   maxGasPrice,
	}, nil
}

func (s SrcExecProvider) Name() string {
	//TODO implement me
	panic("implement me")
}

func (s SrcExecProvider) Start(ctx context.Context) error {
	if s.startBlock != 0 {
		s.lggr.Infow("start replaying src chain", "fromBlock", s.startBlock)
		return s.lp.Replay(ctx, int64(s.startBlock))
	}
	return nil
}

func (s SrcExecProvider) Close() error {
	return nil
}

func (s SrcExecProvider) Ready() error {
	return nil
}

func (s SrcExecProvider) HealthReport() map[string]error {
	return nil
}

func (s SrcExecProvider) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	return s.configWatcher.OffchainConfigDigester()
}

func (s SrcExecProvider) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	return s.configWatcher.ContractConfigTracker()
}

func (s SrcExecProvider) ContractTransmitter() ocrtypes.ContractTransmitter {
	return s.contractTransmitter
}

func (s SrcExecProvider) ChainReader() commontypes.ChainReader {
	return nil
}

func (s SrcExecProvider) Codec() commontypes.Codec {
	return nil
}

func (s SrcExecProvider) NewCommitStoreReader(ctx context.Context, addr cciptypes.Address) (cciptypes.CommitStoreReader, error) {
	//TODO implement me
	panic("implement me")
}

func (s SrcExecProvider) NewOffRampReader(ctx context.Context, addr cciptypes.Address) (cciptypes.OffRampReader, error) {
	panic("NewOffRampReader called on SrcExecProvider. Valid on DstExecProvider.")
}

func (s SrcExecProvider) NewOnRampReader(ctx context.Context, onRampAddress cciptypes.Address, sourceChainSelector uint64, destChainSelector uint64) (onRampReader cciptypes.OnRampReader, err error) {
	versionFinder := ccip.NewEvmVersionFinder()
	onRampReader, err = ccip.NewOnRampReader(s.lggr, versionFinder, sourceChainSelector, destChainSelector, onRampAddress, s.lp, s.client)
	return
}

func (s SrcExecProvider) NewPriceRegistryReader(ctx context.Context, addr cciptypes.Address) (priceRegistryReader cciptypes.PriceRegistryReader, err error) {
	srcPriceRegistry := ccip.NewEvmPriceRegistry(s.lp, s.client, s.lggr, ccip.ExecPluginLabel)
	priceRegistryReader, err = srcPriceRegistry.NewPriceRegistryReader(ctx, addr)
	return
}

func (s SrcExecProvider) NewTokenDataReader(ctx context.Context, tokenAddress cciptypes.Address) (tokenDataReader cciptypes.TokenDataReader, err error) {
	attestationURI, err := url.ParseRequestURI(s.usdcAttestationAPI)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse USDC attestation API")
	}

	tokenAddr, err := ccip.GenericAddrToEvm(tokenAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse token address")
	}
	tokenDataReader = usdc.NewUSDCTokenDataReader(
		s.lggr,
		&s.usdcReader,
		attestationURI,
		s.usdcAttestationAPITimeoutSeconds,
		tokenAddr,
		time.Duration(s.usdcAttestationAPIIntervalMilliseconds)*time.Millisecond,
	)
	return tokenDataReader, nil
}

func (s SrcExecProvider) NewTokenPoolBatchedReader(ctx context.Context, offRampAddr cciptypes.Address, sourceChainSelector uint64) (cciptypes.TokenPoolBatchedReader, error) {
	panic("NewTokenPoolBatchedReader called on SrcExecProvider. It should only be called on DstExecProvdier")
}

func (s SrcExecProvider) SourceNativeToken(ctx context.Context, sourceRouterAddr cciptypes.Address) (cciptypes.Address, error) {
	sourceRouterAddrHex, err := ccip.GenericAddrToEvm(sourceRouterAddr)
	if err != nil {
		return "", err
	}
	sourceRouter, err := router.NewRouter(sourceRouterAddrHex, s.client)
	if err != nil {
		return "", err
	}
	sourceNative, err := sourceRouter.GetWrappedNative(&bind.CallOpts{Context: ctx})
	if err != nil {
		return "", err
	}

	return ccip.EvmAddrToGeneric(sourceNative), nil
}

type DstExecProvider struct {
	lggr          logger.Logger
	versionFinder ccip.VersionFinder
	client        client.Client
	lp            logpoller.LogPoller
	startBlock    uint64
	gasEstimator  gas.EvmFeeEstimator
	maxGasPrice   big.Int
}

func NewDstExecProvider(
	lggr logger.Logger,
	versionFinder ccip.VersionFinder,
	client client.Client,
	lp logpoller.LogPoller,
	startBlock uint64,
	gasEstimator gas.EvmFeeEstimator,
	maxGasPrice big.Int,
) (commontypes.CCIPExecProvider, error) {
	return &DstExecProvider{
		lggr:          lggr,
		versionFinder: versionFinder,
		client:        client,
		lp:            lp,
		startBlock:    startBlock,
		gasEstimator:  gasEstimator,
		maxGasPrice:   maxGasPrice,
	}, nil
}

func (d DstExecProvider) Name() string {
	return "CCIP.DestRelayerExecProvider"
}

func (d DstExecProvider) Start(ctx context.Context) error {
	if d.startBlock != 0 {
		d.lggr.Infow("start replaying dst chain", "fromBlock", d.startBlock)
		return d.lp.Replay(ctx, int64(d.startBlock))
	}
	return nil
}

func (d DstExecProvider) Close() error {
	return nil
}

func (d DstExecProvider) Ready() error {
	return nil
}

func (d DstExecProvider) HealthReport() map[string]error {
	return nil
}

func (d DstExecProvider) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	panic("OffchainConfigDigester called on DstExecProvider. It should only be called on SrcExecProvider.")
}

func (d DstExecProvider) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	panic("ContractConfigTracker called on DstExecProvider. It should only be called on SrcExecProvider.")
}

func (d DstExecProvider) ContractTransmitter() ocrtypes.ContractTransmitter {
	panic("ContractTransmitter called on DstExecProvider. It should only be called on SrcExecProvider.")
}

func (d DstExecProvider) ChainReader() commontypes.ChainReader {
	return nil
}

func (d DstExecProvider) Codec() commontypes.Codec {
	return nil
}

func (d DstExecProvider) NewCommitStoreReader(ctx context.Context, addr cciptypes.Address) (cciptypes.CommitStoreReader, error) {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) NewOffRampReader(ctx context.Context, offRampAddress cciptypes.Address) (offRampReader cciptypes.OffRampReader, err error) {
	offRampReader, err = ccip.NewOffRampReader(d.lggr, d.versionFinder, offRampAddress, d.client, d.lp, d.gasEstimator, &d.maxGasPrice, true)
	return
}

func (d DstExecProvider) NewOnRampReader(ctx context.Context, addr cciptypes.Address, sourceChainSelector uint64, destChainSelector uint64) (cciptypes.OnRampReader, error) {
	panic("NewOnRampReader called on DstExecProvider. It should only be called on SrcExecProvider")
}

func (d DstExecProvider) NewPriceRegistryReader(ctx context.Context, addr cciptypes.Address) (priceRegistryReader cciptypes.PriceRegistryReader, err error) {
	destPriceRegistry := ccip.NewEvmPriceRegistry(d.lp, d.client, d.lggr, ccip.ExecPluginLabel)
	priceRegistryReader, err = destPriceRegistry.NewPriceRegistryReader(ctx, addr)
	return
}

func (d DstExecProvider) NewTokenDataReader(ctx context.Context, tokenAddress cciptypes.Address) (cciptypes.TokenDataReader, error) {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) NewTokenPoolBatchedReader(ctx context.Context, offRampAddress cciptypes.Address, sourceChainSelector uint64) (tokenPoolBatchedReader cciptypes.TokenPoolBatchedReader, err error) {
	batchCaller := ccip.NewDynamicLimitedBatchCaller(
		d.lggr,
		d.client,
		uint(ccip.DefaultRpcBatchSizeLimit),
		uint(ccip.DefaultRpcBatchBackOffMultiplier),
		uint(ccip.DefaultMaxParallelRpcCalls),
	)

	tokenPoolBatchedReader, err = ccip.NewEVMTokenPoolBatchedReader(d.lggr, sourceChainSelector, offRampAddress, batchCaller)
	if err != nil {
		return nil, fmt.Errorf("new token pool batched reader: %w", err)
	}
	return
}

func (d DstExecProvider) SourceNativeToken(ctx context.Context, addr cciptypes.Address) (cciptypes.Address, error) {
	//TODO implement me
	panic("implement me")
}
