package testhelpers

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lib/pq"

	"github.com/smartcontractkit/chainlink/integration-tests/client"

	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/relay"
	"github.com/smartcontractkit/chainlink/core/store/models"
)

type JobType string

const (
	Commit    JobType = "commit"
	Execution JobType = "exec"
	Boostrap  JobType = "bootstrap"
)

func JobName(jobType JobType, source string, destination string) string {
	return fmt.Sprintf("ccip-%s-%s-%s", jobType, source, destination)
}

type CCIPJobSpecParams struct {
	Name                   string
	OffRamp                common.Address
	OnRamp                 common.Address
	CommitStore            common.Address
	SourceChainName        string
	DestChainName          string
	SourceChainId          uint64
	DestChainId            uint64
	TokenPricesUSDPipeline string
	PollPeriod             time.Duration
	SourceStartBlock       uint64
	DestStartBlock         uint64
	RelayInflight          time.Duration
	ExecInflight           time.Duration
	RootSnooze             time.Duration
	P2PV2Bootstrappers     pq.StringArray
}

func (params CCIPJobSpecParams) Validate() error {
	if params.CommitStore == common.HexToAddress("0x0") {
		return fmt.Errorf("must set commit store address")
	}
	if params.SourceChainId == 0 {
		return fmt.Errorf("invalid source chain id")
	}
	if params.DestChainId == 0 {
		return fmt.Errorf("invalid destination chain id")
	}
	return nil
}

func (params CCIPJobSpecParams) ValidateCommitJobSpec() error {
	commonErr := params.Validate()
	if commonErr != nil {
		return commonErr
	}
	if params.OnRamp == common.HexToAddress("") {
		return fmt.Errorf("OnRampOnCommit cannot be empty. The commit job needs to set onRampID")
	}
	if params.OffRamp == common.HexToAddress("0x0") {
		return fmt.Errorf("OffRamp cannot be empty for execution job")
	}

	return nil
}

func (params CCIPJobSpecParams) ValidateExecJobSpec() error {
	commonErr := params.Validate()
	if commonErr != nil {
		return commonErr
	}
	if params.OnRamp == common.HexToAddress("0x0") {
		return fmt.Errorf("OnRampForExecution cannot be empty. The exec job needs to set onRampID")
	}
	if params.OffRamp == common.HexToAddress("0x0") {
		return fmt.Errorf("OffRamp cannot be empty for execution job")
	}
	return nil
}

// CommitJobSpec generates template for CCIP-relay job spec.
// OCRKeyBundleID,TransmitterID need to be set from the calling function
func (params CCIPJobSpecParams) CommitJobSpec() (*client.OCR2TaskJobSpec, error) {
	err := params.ValidateCommitJobSpec()
	if err != nil {
		return nil, err
	}
	ocrSpec := job.OCR2OracleSpec{
		Relay:                             relay.EVM,
		PluginType:                        job.CCIPCommit,
		ContractID:                        params.CommitStore.Hex(),
		ContractConfigConfirmations:       1,
		ContractConfigTrackerPollInterval: models.Interval(20 * time.Second),
		P2PV2Bootstrappers:                params.P2PV2Bootstrappers,
		PluginConfig: map[string]interface{}{
			"sourceChainID": params.SourceChainId,
			"onRampID":      fmt.Sprintf("\"%s\"", params.OnRamp.Hex()),
			"offRampID":     fmt.Sprintf("\"%s\"", params.OffRamp.Hex()),
			"pollPeriod":    `"1s"`,
			"tokenPricesUSDPipeline": fmt.Sprintf(`"""
%s
"""`, params.TokenPricesUSDPipeline),
		},
		RelayConfig: map[string]interface{}{
			"chainID": params.DestChainId,
		},
	}
	if params.RelayInflight.Seconds() > 0 {
		ocrSpec.PluginConfig["pollPeriod"] = fmt.Sprintf("\"%s\"", params.PollPeriod)
	}
	if params.RelayInflight.Seconds() > 0 {
		ocrSpec.PluginConfig["inflightCacheExpiry"] = fmt.Sprintf("\"%s\"", params.RelayInflight)
	}
	if params.DestStartBlock > 0 {
		ocrSpec.PluginConfig["destStartBlock"] = params.DestStartBlock
	}
	if params.SourceStartBlock > 0 {
		ocrSpec.PluginConfig["sourceStartBlock"] = params.SourceStartBlock
	}
	return &client.OCR2TaskJobSpec{
		OCR2OracleSpec: ocrSpec,
		JobType:        "offchainreporting2",
		Name:           JobName(Commit, params.SourceChainName, params.DestChainName),
	}, nil
}

// ExecutionJobSpec generates template for CCIP-execution job spec.
// OCRKeyBundleID,TransmitterID need to be set from the calling function
func (params CCIPJobSpecParams) ExecutionJobSpec() (*client.OCR2TaskJobSpec, error) {
	err := params.ValidateExecJobSpec()
	if err != nil {
		return nil, err
	}
	ocrSpec := job.OCR2OracleSpec{
		Relay:                             relay.EVM,
		PluginType:                        job.CCIPExecution,
		ContractID:                        params.OffRamp.Hex(),
		ContractConfigConfirmations:       1,
		ContractConfigTrackerPollInterval: models.Interval(20 * time.Second),
		P2PV2Bootstrappers:                params.P2PV2Bootstrappers,
		PluginConfig: map[string]interface{}{
			"sourceChainID": params.SourceChainId,
			"onRampID":      fmt.Sprintf("\"%s\"", params.OnRamp.Hex()),
			"commitStoreID": fmt.Sprintf("\"%s\"", params.CommitStore.Hex()),
			"tokenPricesUSDPipeline": fmt.Sprintf(`"""
%s
"""`, params.TokenPricesUSDPipeline),
		},
		RelayConfig: map[string]interface{}{
			"chainID": params.DestChainId,
		},
	}
	if params.ExecInflight.Seconds() > 0 {
		ocrSpec.PluginConfig["inflightCacheExpiry"] = fmt.Sprintf("\"%s\"", params.ExecInflight)
	}
	if params.RootSnooze.Seconds() > 0 {
		ocrSpec.PluginConfig["rootSnoozeTime"] = fmt.Sprintf("\"%s\"", params.RootSnooze)
	}
	if params.DestStartBlock > 0 {
		ocrSpec.PluginConfig["destStartBlock"] = params.DestStartBlock
	}
	if params.SourceStartBlock > 0 {
		ocrSpec.PluginConfig["sourceStartBlock"] = params.SourceStartBlock
	}
	return &client.OCR2TaskJobSpec{
		OCR2OracleSpec: ocrSpec,
		JobType:        "offchainreporting2",
		Name:           JobName(Execution, params.SourceChainName, params.DestChainName),
	}, err
}

func (params CCIPJobSpecParams) BootstrapJob(contractID string) *client.OCR2TaskJobSpec {
	return &client.OCR2TaskJobSpec{
		Name:    fmt.Sprintf("%s-%s", Boostrap, params.DestChainName),
		JobType: "bootstrap",
		OCR2OracleSpec: job.OCR2OracleSpec{
			ContractID:                        contractID,
			Relay:                             relay.EVM,
			ContractConfigConfirmations:       1,
			ContractConfigTrackerPollInterval: models.Interval(20 * time.Second),
			RelayConfig: map[string]interface{}{
				"chainID": params.DestChainId,
			},
		},
	}
}
