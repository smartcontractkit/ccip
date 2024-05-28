package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"math/big"
)

var _ commontypes.CCIPCommitProvider = (*SrcCommitProvider)(nil)
var _ commontypes.CCIPCommitProvider = (*DstCommitProvider)(nil)

type SrcCommitProvider struct {
	lggr                logger.Logger
	startBlock          uint64
	client              client.Client
	lp                  logpoller.LogPoller
	sourceChainSelector uint64
	destChainSelector   uint64
}

type DstCommitProvider struct {
	lggr          logger.Logger
	versionFinder ccip.VersionFinder
	startBlock    uint64
	client        client.Client
	lp            logpoller.LogPoller
	gasEstimator  gas.EvmFeeEstimator
	maxGasPrice   big.Int
}

func (P SrcCommitProvider) Name() string {
	return "CCIPCommitProvider.SourceRelayerProvider"
}

func (P SrcCommitProvider) Close() error {
	//TODO implement me
	panic("implement me")
}

func (P SrcCommitProvider) Ready() error {
	//TODO implement me
	panic("implement me")
}

func (P SrcCommitProvider) HealthReport() map[string]error {
	//TODO implement me
	panic("implement me")
}

func (P SrcCommitProvider) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	//TODO implement me
	panic("implement me")
}

func (P SrcCommitProvider) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	//TODO implement me
	panic("implement me")
}

func (P SrcCommitProvider) ContractTransmitter() ocrtypes.ContractTransmitter {
	//TODO implement me
	panic("implement me")
}

func (P SrcCommitProvider) ChainReader() commontypes.ChainReader {
	//TODO implement me
	panic("implement me")
}

func (P SrcCommitProvider) Codec() commontypes.Codec {
	//TODO implement me
	panic("implement me")
}

func (P DstCommitProvider) Name() string {
	//TODO implement me
	panic("implement me")
}

func (P DstCommitProvider) Close() error {
	//TODO implement me
	panic("implement me")
}

func (P DstCommitProvider) Ready() error {
	//TODO implement me
	panic("implement me")
}

func (P DstCommitProvider) HealthReport() map[string]error {
	//TODO implement me
	panic("implement me")
}

func (P DstCommitProvider) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	//TODO implement me
	panic("implement me")
}

func (P DstCommitProvider) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	//TODO implement me
	panic("implement me")
}

func (P DstCommitProvider) ContractTransmitter() ocrtypes.ContractTransmitter {
	//TODO implement me
	panic("implement me")
}

func (P DstCommitProvider) ChainReader() commontypes.ChainReader {
	//TODO implement me
	panic("implement me")
}

func (P DstCommitProvider) Codec() commontypes.Codec {
	//TODO implement me
	panic("implement me")
}

func NewSrcCommitProvider() commontypes.CCIPCommitProvider {
	panic("implement me")
}

func NewDstCommitProvider() commontypes.CCIPCommitProvider {
	panic("implement me")
}

func (P SrcCommitProvider) Start(ctx context.Context) error {
	if P.startBlock != 0 {
		P.lggr.Infow("start replaying src chain", "fromBlock", P.startBlock)
		return P.lp.Replay(ctx, int64(P.startBlock))
	}
	return nil
}

func (P DstCommitProvider) Start(ctx context.Context) error {
	if P.startBlock != 0 {
		P.lggr.Infow("start replaying dst chain", "fromBlock", P.startBlock)
		return P.lp.Replay(ctx, int64(P.startBlock))
	}
	return nil
}

func (P SrcCommitProvider) NewPriceGetter(ctx context.Context) (priceGetter cciptypes.PriceGetter, err error) {
	panic("Can't construct a price getter from one relayer.")
}

func (P DstCommitProvider) NewPriceGetter(ctx context.Context) (priceGetter cciptypes.PriceGetter, err error) {
	panic("Can't construct a price getter from one relayer.")
}

func (P SrcCommitProvider) NewCommitStoreReader(ctx context.Context, commitStoreAddress cciptypes.Address) (commitStoreReader cciptypes.CommitStoreReader, err error) {
	panic("Can't construct a commit store reader from one relayer.")
}

func (P DstCommitProvider) NewCommitStoreReader(ctx context.Context, commitStoreAddress cciptypes.Address) (commitStoreReader cciptypes.CommitStoreReader, err error) {
	panic("Can't construct a commit store reader from one relayer.")
}

func (P SrcCommitProvider) NewOnRampReader(ctx context.Context, onRampAddress cciptypes.Address, sourceChainSelector uint64, destChainSelector uint64) (onRampReader cciptypes.OnRampReader, err error) {
	versionFinder := ccip.NewEvmVersionFinder()
	onRampReader, err = ccip.NewOnRampReader(P.lggr, versionFinder, sourceChainSelector, destChainSelector, onRampAddress, P.lp, P.client)
	return
}

func (P DstCommitProvider) NewOnRampReader(ctx context.Context, onRampAddress cciptypes.Address, sourceChainSelector uint64, destChainSelector uint64) (onRampReader cciptypes.OnRampReader, err error) {
	panic("NewOnRampReader called for DstCommitProvider.NewOnRampReader should be called on SrcCommitProvider")
}

func (P SrcCommitProvider) NewOffRampReader(ctx context.Context, offRampAddr cciptypes.Address) (offRampReader cciptypes.OffRampReader, err error) {
	panic("Called NewOffRampReader for SrcCommitProvider. NewOffRampReader should be called on DstCommitProvider")
}

func (P DstCommitProvider) NewOffRampReader(ctx context.Context, offRampAddr cciptypes.Address) (offRampReader cciptypes.OffRampReader, err error) {
	offRampReader, err = ccip.NewOffRampReader(P.lggr, P.versionFinder, offRampAddr, P.client, P.lp, P.gasEstimator, &P.maxGasPrice, true)
	return
}

func (P SrcCommitProvider) NewPriceRegistryReader(ctx context.Context, addr cciptypes.Address) (priceRegistryReader cciptypes.PriceRegistryReader, err error) {
	panic("Called NewPriceRegistryReader for SrcCommitProvider. NewOffRampReader should be called on DstCommitProvider")
}

func (P DstCommitProvider) NewPriceRegistryReader(ctx context.Context, addr cciptypes.Address) (priceRegistryReader cciptypes.PriceRegistryReader, err error) {
	destPriceRegistry := ccip.NewEvmPriceRegistry(P.lp, P.client, P.lggr, ccip.CommitPluginLabel)
	priceRegistryReader, err = destPriceRegistry.NewPriceRegistryReader(ctx, addr)
	return
}

func (P SrcCommitProvider) SourceNativeToken(ctx context.Context, sourceRouterAddr cciptypes.Address) (cciptypes.Address, error) {
	sourceRouterAddrHex := sourceRouterAddr.ToCommonAddress()
	sourceRouter, err := router.NewRouter(sourceRouterAddrHex, P.client)
	if err != nil {
		return "", err
	}
	sourceNative, err := sourceRouter.GetWrappedNative(&bind.CallOpts{Context: ctx})
	if err != nil {
		return "", err
	}

	return cciptypes.FromCommonAddress(sourceNative), nil
}

func (P DstCommitProvider) SourceNativeToken(ctx context.Context, sourceRouterAddr cciptypes.Address) (cciptypes.Address, error) {
	panic("SourceNativeToken called for DstCommitProvider. SourceNativeToken should be called on SrcCommitProvider")
}
