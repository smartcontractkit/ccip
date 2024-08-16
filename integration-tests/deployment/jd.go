package deployment

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type JDConfig struct {
	Server string
	// TODO: add auth details
}

func NewClientConnection(cfg JDConfig) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	// TODO: add auth details
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(cfg.Server, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect Job Distributor service. Err: %w", err)
	}

	return conn, nil
}
