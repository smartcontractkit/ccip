package launcher

import (
	"go.uber.org/multierr"

	cctypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
)

// blueGreenDeployment represents a blue-green deployment of OCR instances.
type blueGreenDeployment struct {
	// blue is the blue OCR instance.
	// blue must always be present.
	blue cctypes.CCIPOracle

	// green is the green OCR instance.
	// green may or may not be present.
	// green must never be present if blue is not present.
	// TODO: should we enforce this invariant somehow?
	green cctypes.CCIPOracle
}

// ccipDeployment represents blue-green deployments of both commit and exec
// OCR instances.
type ccipDeployment struct {
	commit blueGreenDeployment
	exec   blueGreenDeployment
}

// Shutdown shuts down all OCR instances in the deployment.
func (c *ccipDeployment) Shutdown() error {
	var err error

	err = multierr.Append(err, c.commit.blue.Shutdown())
	if c.commit.green != nil {
		err = multierr.Append(err, c.commit.green.Shutdown())
	}

	err = multierr.Append(err, c.exec.blue.Shutdown())
	if c.exec.green != nil {
		err = multierr.Append(err, c.exec.green.Shutdown())
	}
	return err
}

// NumCommitInstances returns the number of commit OCR instances in the deployment.
func (c *ccipDeployment) NumCommitInstances() int {
	if c.commit.green != nil {
		return 2
	}
	return 1
}

// NumExecInstances returns the number of exec OCR instances in the deployment.
func (c *ccipDeployment) NumExecInstances() int {
	if c.exec.green != nil {
		return 2
	}
	return 1
}
