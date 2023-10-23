package testconfig

import (
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

type CCIPTestConfig struct {
	KeepEnvAlive               *bool              `toml:",omitempty"`
	MsgType                    string             `toml:",omitempty"`
	PhaseTimeout               *models.Duration   `toml:",omitempty"`
	TestDuration               *models.Duration   `toml:",omitempty"`
	LocalCluster               *bool              `toml:",omitempty"`
	ExistingDeployment         *bool              `toml:",omitempty"`
	ExistingEnv                string             `toml:",omitempty"`
	ReuseContracts             *bool              `toml:",omitempty"`
	SequentialLaneAddition     *bool              `toml:",omitempty"`
	NodeFunding                float64            `toml:",omitempty"`
	RequestPerUnitTime         []int64            `toml:",omitempty"`
	TimeUnit                   *models.Duration   `toml:",omitempty"`
	StepDuration               []*models.Duration `toml:",omitempty"`
	WaitBetweenChaosDuringLoad *models.Duration   `toml:",omitempty"`
	NetworkPairs               []string           `toml:",omitempty"`
	NoOfNetworks               int                `toml:",omitempty"`
	NoOfRoutersPerPair         int                `toml:",omitempty"`
	Blockscout                 bool               `toml:",omitempty"`
	NoOfTokensPerChain         int                `toml:",omitempty"`
	NoOfTokensInMsg            int                `toml:",omitempty"`
	AmountPerToken             int64              `toml:",omitempty"`
	MaxNoOfLanes               int                `toml:",omitempty"`
}

func (c CCIPTestConfig) ApplyOverrides(from Group) error {
	fromCfg, ok := from.(CCIPTestConfig)
	if !ok {
		return errors.Errorf("invalid CCIP test config type %T", from)
	}
	if fromCfg.KeepEnvAlive != nil {
		c.KeepEnvAlive = fromCfg.KeepEnvAlive
	}
	if fromCfg.MsgType != "" {
		c.MsgType = fromCfg.MsgType
	}
	if fromCfg.PhaseTimeout != nil {
		c.PhaseTimeout = fromCfg.PhaseTimeout
	}
	if fromCfg.TestDuration != nil {
		c.TestDuration = fromCfg.TestDuration
	}
	if fromCfg.LocalCluster != nil {
		c.LocalCluster = fromCfg.LocalCluster
	}
	if fromCfg.ExistingDeployment != nil {
		c.ExistingDeployment = fromCfg.ExistingDeployment
	}
	if fromCfg.ExistingEnv != "" {
		c.ExistingEnv = fromCfg.ExistingEnv
	}
	if fromCfg.ReuseContracts != nil {
		c.ReuseContracts = fromCfg.ReuseContracts
	}
	if fromCfg.SequentialLaneAddition != nil {
		c.SequentialLaneAddition = fromCfg.SequentialLaneAddition
	}
	if fromCfg.NodeFunding != 0 {
		c.NodeFunding = fromCfg.NodeFunding
	}
	if len(fromCfg.RequestPerUnitTime) != 0 {
		c.RequestPerUnitTime = fromCfg.RequestPerUnitTime
	}
	if fromCfg.TimeUnit != nil {
		c.TimeUnit = fromCfg.TimeUnit
	}
	if len(fromCfg.StepDuration) != 0 {
		c.StepDuration = fromCfg.StepDuration
	}
	if fromCfg.WaitBetweenChaosDuringLoad != nil {
		c.WaitBetweenChaosDuringLoad = fromCfg.WaitBetweenChaosDuringLoad
	}
	if len(fromCfg.NetworkPairs) != 0 {
		c.NetworkPairs = fromCfg.NetworkPairs
	}
	if fromCfg.NoOfNetworks != 0 {
		c.NoOfNetworks = fromCfg.NoOfNetworks
	}
	if fromCfg.NoOfRoutersPerPair != 0 {
		c.NoOfRoutersPerPair = fromCfg.NoOfRoutersPerPair
	}
	if fromCfg.Blockscout {
		c.Blockscout = fromCfg.Blockscout
	}
	if fromCfg.NoOfTokensPerChain != 0 {
		c.NoOfTokensPerChain = fromCfg.NoOfTokensPerChain
	}
	if fromCfg.NoOfTokensInMsg != 0 {
		c.NoOfTokensInMsg = fromCfg.NoOfTokensInMsg
	}
	if fromCfg.MaxNoOfLanes != 0 {
		c.MaxNoOfLanes = fromCfg.MaxNoOfLanes
	}
	if fromCfg.AmountPerToken != 0 {
		c.AmountPerToken = fromCfg.AmountPerToken
	}
	return nil
}

func (c CCIPTestConfig) ReadSecrets() error {
	// no secrets to read
	return nil
}

func (c CCIPTestConfig) Validate() error {
	if c.PhaseTimeout.Duration().Minutes() < 1 || c.PhaseTimeout.Duration().Minutes() > 50 {
		return errors.Errorf("phase timeout should be between 1 and 50 minutes")
	}
	if c.TestDuration.Duration().Minutes() < 1 {
		return errors.Errorf("test duration should be greater than 1 minute")
	}
	if c.MsgType != "WithoutToken" && c.MsgType != "WithToken" {
		return errors.Errorf("msg type should be either WithoutToken or WithToken")
	}

	if c.MsgType == "WithToken" {
		if c.AmountPerToken == 0 {
			return errors.Errorf("token amount should be greater than 0")
		}
		if c.NoOfTokensPerChain == 0 {
			return errors.Errorf("number of tokens per chain should be greater than 0")
		}
		if c.NoOfTokensInMsg == 0 {
			return errors.Errorf("number of tokens in msg should be greater than 0")
		}
	}

	return nil
}
