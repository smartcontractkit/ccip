package launcher

import (
	"fmt"

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

func (c *ccipDeployment) StartBlue() error {
	var err error

	err = multierr.Append(err, c.commit.blue.Start())
	err = multierr.Append(err, c.exec.blue.Start())

	return err
}

func (c *ccipDeployment) ShutdownBlue() error {
	var err error

	err = multierr.Append(err, c.commit.blue.Shutdown())
	err = multierr.Append(err, c.exec.blue.Shutdown())

	return err
}

func (c *ccipDeployment) HandleBlueGreen(prevDeployment *ccipDeployment) error {
	if prevDeployment == nil {
		return fmt.Errorf("previous deployment is nil")
	}

	// two possible cases:
	// 1. both blue and green are present in prevDeployment, only blue is present in c.
	// this is a promotion of green to blue, so we need to shut down the blue deployment
	// and make green the new blue.
	// 2. only blue is present in prevDeployment, both blue and green are present in c
	var err error
	if prevDeployment.commit.green != nil && c.commit.green == nil {
		// case 1
		// green is already running so no need to start it.
		// shutdown blue.
		err = multierr.Append(err, prevDeployment.commit.blue.Shutdown())
	} else if prevDeployment.commit.green == nil && c.commit.green != nil {
		// case 2
		// blue is already running so no need to start it.
		// start green.
		err = multierr.Append(err, c.commit.green.Start())
	} else {
		return fmt.Errorf("invalid blue-green deployment transition")
	}

	if prevDeployment.exec.green != nil && c.exec.green == nil {
		// case 1
		// green is already running so no need to start it.
		// shutdown blue.
		err = multierr.Append(err, prevDeployment.exec.blue.Shutdown())
	} else if prevDeployment.exec.green == nil && c.exec.green != nil {
		// case 2
		// blue is already running so no need to start it.
		// start green.
		err = multierr.Append(err, c.exec.green.Start())
	} else {
		return fmt.Errorf("invalid blue-green deployment transition")
	}

	return err
}

// HasGreenCommitInstance returns true if and only if the green commit instance is not nil.
func (c *ccipDeployment) HasGreenCommitInstance() bool {
	return c.commit.green != nil
}

// HasGreenExecInstance returns true if and only if the green exec instance is not nil.
func (c *ccipDeployment) HasGreenExecInstance() bool {
	return c.exec.green != nil
}
