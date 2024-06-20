package ccipcapability

import (
	"fmt"

	"github.com/pelletier/go-toml"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

func ValidatedCCIPSpec(tomlString string) (jb job.Job, err error) {
	var spec job.CCIPSpec
	tree, err := toml.Load(tomlString)
	if err != nil {
		return job.Job{}, fmt.Errorf("toml error on load: %w", err)
	}
	// Note this validates all the fields which implement an UnmarshalText
	err = tree.Unmarshal(&spec)
	if err != nil {
		return job.Job{}, fmt.Errorf("toml unmarshal error on spec: %w", err)
	}
	err = tree.Unmarshal(&jb)
	if err != nil {
		return job.Job{}, fmt.Errorf("toml unmarshal error on job: %w", err)
	}
	jb.CCIPSpec = &spec

	if jb.Type != job.CCIP {
		return job.Job{}, fmt.Errorf("the only supported type is currently 'ccip', got %s", jb.Type)
	}
	if jb.CCIPSpec.CapabilityLabelledName == "" {
		return job.Job{}, fmt.Errorf("capabilityLabelledName must be set")
	}
	if jb.CCIPSpec.CapabilityVersion == "" {
		return job.Job{}, fmt.Errorf("capabilityVersion must be set")
	}
	if jb.CCIPSpec.P2PKeyID == "" {
		return job.Job{}, fmt.Errorf("p2pKeyID must be set")
	}

	return jb, nil
}

func ValidatedCCIPBootstrapSpec(tomlString string) (jb job.Job, err error) {
	var spec job.CCIPBootstrapSpec
	tree, err := toml.Load(tomlString)
	if err != nil {
		return job.Job{}, fmt.Errorf("toml error on load: %w", err)
	}
	// Note this validates all the fields which implement an UnmarshalText
	err = tree.Unmarshal(&spec)
	if err != nil {
		return job.Job{}, fmt.Errorf("toml unmarshal error on spec: %w", err)
	}
	err = tree.Unmarshal(&jb)
	if err != nil {
		return job.Job{}, fmt.Errorf("toml unmarshal error on job: %w", err)
	}
	jb.CCIPBootstrapSpec = &spec

	if jb.Type != job.CCIPBootstrap {
		return job.Job{}, fmt.Errorf("the only supported type is currently 'ccipbootstrap', got %s", jb.Type)
	}
	if jb.CCIPBootstrapSpec.CapabilityLabelledName == "" {
		return job.Job{}, fmt.Errorf("capabilityLabelledName must be set")
	}
	if jb.CCIPBootstrapSpec.CapabilityVersion == "" {
		return job.Job{}, fmt.Errorf("capabilityVersion must be set")
	}

	return jb, nil
}
