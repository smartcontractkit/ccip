package devenv

import (
	"github.com/smartcontractkit/chainlink/integration-tests/web/sdk/client"
)

type DON struct {
	Bootstrap NodeDetail
	Nodes     []NodeDetail
}

type NodeDetail struct {
	FMS client.Client
}

// For Each Node Operations
//  - Create Node Connection :
//      1. Instantiate Node Connection with graphQL client with Node URL and creds
//  - Create Feeds Manager Connection :
//      1. Create Feeds Manager Connection with Node's graphQL request

// JD operations:
//	- Create JD connection :
//      1. create JD connection with JD URL and creds
//      2. With the JD connection, create CSAService Client, NodeService Client and JobService Client
//      3. Open Qs -
//  	  - Where would the definition for CSAService Client, NodeService Client and JobService Client live?
//  - Get CSA Keys :
// 	    1.  Get CSA Keys of JD with CSAService Client
//  - Register Nodes :
//      1. Register Each Node with JD's NodeService Client
//      2. Store Node ID returned by JD's Register Node response

// Job Proposal Operations
//  - Propose Job by JD to Node :
//      1. Propose Job to Node with JD's JobService Client with Node ID
// Node Operations
//  - Approve Job by Node :
//      1. Approve Job with Node's FeedsManagerClient ( How to create the FeedsManagerClient client to approve the job?)
