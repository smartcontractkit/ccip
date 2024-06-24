package ccipcapability

import (
	"fmt"
	"strconv"
	"strings"

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

	// ensure that the P2PV2Bootstrappers is in the right format.
	for _, bootstrapperLocator := range jb.CCIPSpec.P2PV2Bootstrappers {
		// needs to be of the form <peer_id>@<ip-address>:<port>
		_, _, _, err := parseBootstrapperLocator(bootstrapperLocator)
		if err != nil {
			return job.Job{}, fmt.Errorf("p2p v2 bootstrapper locator %s is not in the correct format: %w", bootstrapperLocator, err)
		}
	}

	return jb, nil
}

func parseBootstrapperLocator(locator string) (peerID, ipAddress string, port int, err error) {
	// needs to be of the form <peer_id>@<ip-address>:<port>
	parts := strings.Split(locator, "@")
	if len(parts) != 2 {
		return "", "", 0, fmt.Errorf("expected 2 parts after splitting on '@', got %d", len(parts))
	}
	ipAndPort := strings.Split(parts[1], ":")
	if len(ipAndPort) != 2 {
		return "", "", 0, fmt.Errorf("expected 2 parts after splitting on ':', got %d", len(ipAndPort))
	}
	// validate port is an integer
	p, err := strconv.ParseInt(ipAndPort[1], 10, 64)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to parse port as an integer: %w", err)
	}
	return parts[0], ipAndPort[0], int(p), nil
}
