package ccip

import (
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/services/job"
)

func ValidatedCCIPSpec(tomlString string) (job.Job, error) {
	var jb = job.Job{}
	tree, err := toml.Load(tomlString)
	if err != nil {
		return jb, err
	}
	err = tree.Unmarshal(&jb)
	if err != nil {
		return jb, err
	}
	switch jb.Type {
	case job.CCIPRelay:
		var spec job.CCIPRelaySpec
		err = tree.Unmarshal(&spec)
		if err != nil {
			return jb, err
		}
		jb.CCIPRelaySpec = &spec
	case job.CCIPExecution:
		var spec job.CCIPExecutionSpec
		err = tree.Unmarshal(&spec)
		if err != nil {
			return jb, err
		}
		jb.CCIPExecutionSpec = &spec
	default:
		return jb, errors.Errorf("unsupported type %s", jb.Type)
	}

	return jb, nil
}
