package ccipexec

import (
	"context"
	"fmt"
	"testing"

	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

type MockCCIPExecProvider struct{}

func (m MockCCIPExecProvider) Name() string {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) Start(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) Close() error {
	return fmt.Errorf("expected err")
}

func (m MockCCIPExecProvider) Ready() error {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) HealthReport() map[string]error {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) ContractTransmitter() ocrtypes.ContractTransmitter {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) ChainReader() types.ChainReader {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) Codec() types.Codec {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) NewCommitStoreReader(ctx context.Context, addr cciptypes.Address) (cciptypes.CommitStoreReader, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) NewOffRampReader(ctx context.Context, addr cciptypes.Address) (cciptypes.OffRampReader, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) NewOnRampReader(ctx context.Context, addr cciptypes.Address, sourceSelector uint64, destSelector uint64) (cciptypes.OnRampReader, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) NewPriceRegistryReader(ctx context.Context, addr cciptypes.Address) (cciptypes.PriceRegistryReader, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) NewTokenDataReader(ctx context.Context, tokenAddress cciptypes.Address) (cciptypes.TokenDataReader, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) NewTokenPoolBatchedReader(ctx context.Context, offRampAddress cciptypes.Address, sourceSelector uint64) (cciptypes.TokenPoolBatchedReader, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockCCIPExecProvider) SourceNativeToken(ctx context.Context, addr cciptypes.Address) (cciptypes.Address, error) {
	//TODO implement me
	panic("implement me")
}

func TestGetExecutionPluginFilterNamesFromSpec(t *testing.T) {
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
				PluginConfig: map[string]interface{}{},
			},
			expectingErr: true,
		},
		{
			description: "invalid off ramp address",
			spec: &job.OCR2OracleSpec{
				PluginConfig: map[string]interface{}{"offRamp": "123"},
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
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			err := UnregisterExecPluginLpFilters(MockCCIPExecProvider{}, MockCCIPExecProvider{})
			if tc.expectingErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
