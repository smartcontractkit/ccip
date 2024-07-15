package ccipcommit

import (
	"context"
	"fmt"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

type MockCCIPCommitProvider struct{}

func (m MockCCIPCommitProvider) Name() string {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPCommitProvider) Start(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPCommitProvider) Close() error {
	return fmt.Errorf("expected err")
}

func (m MockCCIPCommitProvider) Ready() error {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPCommitProvider) HealthReport() map[string]error {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPCommitProvider) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPCommitProvider) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPCommitProvider) ContractTransmitter() ocrtypes.ContractTransmitter {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPCommitProvider) ChainReader() types.ChainReader {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPCommitProvider) Codec() types.Codec {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPCommitProvider) NewCommitStoreReader(ctx context.Context, addr cciptypes.Address) (cciptypes.CommitStoreReader, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPCommitProvider) NewOffRampReader(ctx context.Context, addr cciptypes.Address) (cciptypes.OffRampReader, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPCommitProvider) NewOnRampReader(ctx context.Context, addr cciptypes.Address, sourceSelector uint64, destSelector uint64) (cciptypes.OnRampReader, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPCommitProvider) NewPriceGetter(ctx context.Context) (cciptypes.PriceGetter, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPCommitProvider) NewPriceRegistryReader(ctx context.Context, addr cciptypes.Address) (cciptypes.PriceRegistryReader, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPCommitProvider) SourceNativeToken(ctx context.Context, addr cciptypes.Address) (cciptypes.Address, error) {
	//TODO implement me
	panic("implement me")
}

func TestGetCommitPluginFilterNamesFromSpec(t *testing.T) {
	testCases := []struct {
		description  string
		spec         *job.OCR2OracleSpec
		expectingErr bool
	}{
		{
			description:  "should not panic with nil spec",
			spec:         nil,
			expectingErr: true,
		},
		{
			description: "invalid config",
			spec: &job.OCR2OracleSpec{
				ContractID:   utils.ZeroAddress.String(),
				PluginConfig: map[string]interface{}{},
			},
			expectingErr: true,
		},
		{
			description: "invalid contract id",
			spec: &job.OCR2OracleSpec{
				ContractID: "whatever...",
			},
			expectingErr: true,
		},
		{
			description: "valid config",
			spec: &job.OCR2OracleSpec{
				ContractID:   utils.ZeroAddress.String(),
				PluginConfig: map[string]interface{}{},
				RelayConfig: map[string]interface{}{
					"chainID": 1234.0,
				},
			},
			expectingErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			err := UnregisterCommitPluginLpFilters(MockCCIPCommitProvider{}, MockCCIPCommitProvider{})
			if tc.expectingErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
