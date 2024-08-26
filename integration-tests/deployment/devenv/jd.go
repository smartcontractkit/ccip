package devenv

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/smartcontractkit/ccip/integration-tests/deployment"
	csav1 "github.com/smartcontractkit/ccip/integration-tests/deployment/jd/csa/v1"
	jobv1 "github.com/smartcontractkit/ccip/integration-tests/deployment/jd/job/v1"
	nodev1 "github.com/smartcontractkit/ccip/integration-tests/deployment/jd/node/v1"
)

type JDConfig struct {
	Server string
	Creds  credentials.TransportCredentials
}

func NewClientConnection(cfg JDConfig) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	// TODO: add auth details
	if cfg.Creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(cfg.Creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	}

	conn, err := grpc.NewClient(cfg.Server, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect Job Distributor service. Err: %w", err)
	}

	return conn, nil
}

type JDClient struct {
	nodev1.NodeServiceClient
	jobv1.JobServiceClient
	csav1.CSAServiceClient
}

func NewJDClient(cfg JDConfig) (deployment.OffchainClient, error) {
	conn, err := NewClientConnection(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect Job Distributor service. Err: %w", err)
	}
	return JDClient{
		nodev1.NewNodeServiceClient(conn),
		jobv1.NewJobServiceClient(conn),
		csav1.NewCSAServiceClient(conn),
	}, err
}
