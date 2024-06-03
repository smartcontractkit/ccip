package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	gasEstimator                           gas.EvmFeeEstimator
	maxGasPrice                            big.Int
	jobID                                  string
	usdcReader                             ccip.USDCReaderImpl
	usdcAttestationAPI                     string
	usdcAttestationAPITimeoutSeconds       int
	usdcAttestationAPIIntervalMilliseconds int
}

func NewSrcExecProvider(
	lggr logger.Logger,
	versionFinder ccip.VersionFinder,
	client client.Client,
	lp logpoller.LogPoller,
	jobID string,
	usdcReader ccip.USDCReaderImpl,
	usdcAttestationAPI string,
	usdcAttestationAPITimeoutSeconds int,
	usdcAttestationAPIIntervalMilliseconds int,
	// gasEstimator gas.EvmFeeEstimator,
	// maxGasPrice big.Int,
) (commontypes.CCIPExecProvider, error) {

	return &SrcExecProvider{
		lggr:                                   lggr,
		versionFinder:                          versionFinder,
		client:                                 client,
		lp:                                     lp,
		jobID:                                  jobID,
		usdcReader:                             usdcReader,
		usdcAttestationAPI:                     usdcAttestationAPI,
		usdcAttestationAPITimeoutSeconds:       usdcAttestationAPITimeoutSeconds,
		usdcAttestationAPIIntervalMilliseconds: usdcAttestationAPIIntervalMilliseconds,
		//gasEstimator:  gasEstimator,
		//maxGasPrice:   maxGasPrice,
	}, nil
}

func (s SrcExecProvider) Name() string {
	//TODO implement me
	panic("implement me")
}

func (s SrcExecProvider) Start(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s SrcExecProvider) Close() error {
	//TODO implement me
	panic("implement me")
}

func (s SrcExecProvider) Ready() error {
	//TODO implement me
	panic("implement me")
}

func (s SrcExecProvider) HealthReport() map[string]error {
	//TODO implement me
	panic("implement me")
}

func (s SrcExecProvider) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	//TODO implement me
	panic("implement me")
}

func (s SrcExecProvider) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	//TODO implement me
	panic("implement me")
}

func (s SrcExecProvider) ContractTransmitter() ocrtypes.ContractTransmitter {
	//TODO implement me
	panic("implement me")
}

func (s SrcExecProvider) ChainReader() commontypes.ChainReader {
	//TODO implement me
	panic("implement me")
}

func (s SrcExecProvider) Codec() commontypes.Codec {
	//TODO implement me
	panic("implement me")
}

func (s SrcExecProvider) NewCommitStoreReader(ctx context.Context, addr cciptypes.Address) (cciptypes.CommitStoreReader, error) {
	//TODO implement me
	panic("implement me")
}

func (s SrcExecProvider) NewOffRampReader(ctx context.Context, addr cciptypes.Address) (cciptypes.OffRampReader, error) {
	panic("NewOffRampReader called on SrcExecProvider. Valid on DstExecProvider.")
}

func (s SrcExecProvider) NewOnRampReader(ctx context.Context, addr cciptypes.Address, sourceSelector uint64, destSelector uint64) (cciptypes.OnRampReader, error) {
	//TODO implement me
	panic("implement me")
}

func (s SrcExecProvider) NewPriceRegistryReader(ctx context.Context, addr cciptypes.Address) (cciptypes.PriceRegistryReader, error) {
	//TODO implement me
	panic("implement me")
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

func (s SrcExecProvider) NewTokenPoolBatchedReader(ctx context.Context, offRampAddr cciptypes.Address) (cciptypes.TokenPoolBatchedReader, error) {
	panic("don't do this")
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
	lggr                logger.Logger
	versionFinder       ccip.VersionFinder
	client              client.Client
	lp                  logpoller.LogPoller
	gasEstimator        gas.EvmFeeEstimator
	maxGasPrice         big.Int
	sourceChainSelector uint64
}

func NewDstExecProvider(
	lggr logger.Logger,
	versionFinder ccip.VersionFinder,
	client client.Client,
	lp logpoller.LogPoller,
	gasEstimator gas.EvmFeeEstimator,
	maxGasPrice big.Int,
	sourceChainSelector uint64,
) commontypes.CCIPExecProvider {
	return &DstExecProvider{
		lggr:                lggr,
		versionFinder:       versionFinder,
		client:              client,
		lp:                  lp,
		gasEstimator:        gasEstimator,
		maxGasPrice:         maxGasPrice,
		sourceChainSelector: sourceChainSelector,
	}
}

func (d DstExecProvider) Name() string {
	return "CCIP.DestRelayerExecProvider"
}

func (d DstExecProvider) Start(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) Close() error {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) Ready() error {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) HealthReport() map[string]error {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) ContractTransmitter() ocrtypes.ContractTransmitter {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) ChainReader() commontypes.ChainReader {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) Codec() commontypes.Codec {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) NewCommitStoreReader(ctx context.Context, addr cciptypes.Address) (cciptypes.CommitStoreReader, error) {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) NewOffRampReader(ctx context.Context, offRampAddress cciptypes.Address) (offRampReader cciptypes.OffRampReader, err error) {
	offRampReader, err = ccip.NewOffRampReader(d.lggr, d.versionFinder, offRampAddress, d.client, d.lp, d.gasEstimator, &d.maxGasPrice, true)
	return
}

func (d DstExecProvider) NewOnRampReader(ctx context.Context, addr cciptypes.Address, sourceSelector uint64, destSelector uint64) (cciptypes.OnRampReader, error) {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) NewPriceRegistryReader(ctx context.Context, addr cciptypes.Address) (cciptypes.PriceRegistryReader, error) {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) NewTokenDataReader(ctx context.Context, tokenAddress cciptypes.Address) (cciptypes.TokenDataReader, error) {
	//TODO implement me
	panic("implement me")
}

func (d DstExecProvider) NewTokenPoolBatchedReader(ctx context.Context, offRampAddress cciptypes.Address) (tokenPoolBatchedReader cciptypes.TokenPoolBatchedReader, err error) {
	batchCaller := ccip.NewDynamicLimitedBatchCaller(
		d.lggr,
		d.client,
		uint(ccip.DefaultRpcBatchSizeLimit),
		uint(ccip.DefaultRpcBatchBackOffMultiplier),
		uint(ccip.DefaultMaxParallelRpcCalls),
	)

	tokenPoolBatchedReader, err = ccip.NewEVMTokenPoolBatchedReader(d.lggr, d.sourceChainSelector, offRampAddress, batchCaller)
	if err != nil {
		return nil, fmt.Errorf("new token pool batched reader: %w", err)
	}
}

func (d DstExecProvider) SourceNativeToken(ctx context.Context, addr cciptypes.Address) (cciptypes.Address, error) {
	//TODO implement me
	panic("implement me")
}
