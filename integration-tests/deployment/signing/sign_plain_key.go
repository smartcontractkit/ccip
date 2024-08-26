package signing

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	// NOTE MUST BE > 1.14 for this fix
	// https://github.com/ethereum/go-ethereum/pull/28945

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment/executable"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/managed"
)

// Just run this locally to sign from the ledger.
func signPlainKey(privateKeyHex string) {
	// Load file
	proposal, _ := ProposalFromFile(managed.MCMSProposalTypeMap[os.Args[0]], os.Args[1])
	err := proposal.Validate()
	if err != nil {
		fmt.Println(err)
		return
	}

	executableProposal, err := proposal.ToExecutableMCMSProposal()
	if err != nil {
		fmt.Println(err)
		return
	}

	executor, err := executableProposal.ToExecutor(make(map[string]executable.ContractDeployBackend)) // TODO: pass in a real backend
	if err != nil {
		log.Fatal(err)
	}

	// Get the signing hash
	payload, err := executor.SigningHash()
	if err != nil {
		log.Fatal(err)
	}

	// Load private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	// Sign the payload
	sig, err := crypto.Sign(payload.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Sign the payload
	unmarshalledSig := executable.Signature{}
	err = json.Unmarshal(sig, &unmarshalledSig)
	if err != nil {
		log.Fatal(err)
	}

	// Add signature to proposal
	proposal.AddSignature(unmarshalledSig)

	// Write proposal to file
	WriteProposalToFile(proposal, os.Args[0])
}
