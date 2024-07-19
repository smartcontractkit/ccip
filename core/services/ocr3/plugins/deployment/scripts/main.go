package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	deployments "github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment/jobdistributor"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"os"
)

var (
	jobDistributorURL string
	chainIDs          []int64
	nodeIDs           []string
	addressBookFile   string
)

func newBackend(chainID int64) bind.ContractBackend {
	// probably from env/config file
	return nil
}

func clientConn(url string) grpc.ClientConnInterface {
	return nil
}

func main() {
	// Example of command to deploy to testnet, using same stack that was used for in memory test.
	// Note that eventually this could run in CI even, but we'd need a way to
	// manually recover from partial deployments.
	var testnetDeploy = &cobra.Command{
		Use: "deploy fresh CCIP to existing DON",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: validate args.
			chains := make(map[uint64]bind.ContractBackend)
			for _, chainID := range chainIDs {
				chains[uint64(chainID)] = newBackend(chainID)
			}
			deployments.DeployNewCCIPToExistingDON(
				deployments.NewFileAddressBook(addressBookFile),
				nodeIDs,
				chains,
				nil,
				jobdistributor.NewJobServiceClient(clientConn(jobDistributorURL)))
		},
	}

	testnetDeploy.Flags().StringVar(&jobDistributorURL, "jdURL", "", "")
	testnetDeploy.Flags().StringVar(&addressBookFile, "addressBookFile", "", "where to save addresses")
	// Assume chain->RPC mapping stores somewhere (like CCIP scripts)
	testnetDeploy.Flags().Int64SliceVar(&chainIDs, "chainIDs", []int64{}, "list of chainIDs to deploy to")
	testnetDeploy.Flags().StringArrayVar(&nodeIDs, "nodeIds", []string{}, "list of nodes to deploy to")

	if err := testnetDeploy.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
