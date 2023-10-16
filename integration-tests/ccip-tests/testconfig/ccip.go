package testconfig

import (
	"time"
)

type CCIPTestConfig struct {
	KeepEnvAlive           bool          `toml:",omitempty"`
	MsgType                string        `toml:",omitempty"`
	PhaseTimeout           time.Duration `toml:",omitempty"`
	TestDuration           time.Duration `toml:",omitempty"`
	LocalCluster           bool          `toml:",omitempty"`
	ExistingDeployment     bool          `toml:",omitempty"`
	ExistingEnv            string        `toml:",omitempty"`
	ReuseContracts         bool          `toml:",omitempty"`
	SequentialLaneAddition bool          `toml:",omitempty"`
	NodeFunding            float64       `toml:",omitempty"`
}

func (C CCIPTestConfig) ApplyOverrides(from *Config) error {
	//TODO implement me
	panic("implement me")
}

func (C CCIPTestConfig) ReadSecrets() error {
	//TODO implement me
	panic("implement me")
}

func (C CCIPTestConfig) Validate() error {
	//TODO implement me
	panic("implement me")
}

type CCIPLoad struct {
}
