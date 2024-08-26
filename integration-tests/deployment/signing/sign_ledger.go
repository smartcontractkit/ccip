package signing

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	// NOTE MUST BE > 1.14 for this fix
	// https://github.com/ethereum/go-ethereum/pull/28945

	"github.com/ethereum/go-ethereum/accounts/usbwallet"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment/executable"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/managed"
)

// Just run this locally to sign from the ledger.
func signLedger() {
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

	// Load ledger
	ledgerhub, _ := usbwallet.NewLedgerHub()
	wallets := ledgerhub.Wallets()
	wallet := wallets[0]

	// Open the ledger.
	_ = wallet.Open("")

	// Load account.
	// BIP44 derivation path used in ledger.
	// Could pass this in as an argument as well.
	var derivationPath = []uint32{
		44 | 0x80000000,
		60 | 0x80000000,
		0 | 0x80000000,
		0,
		uint32(0), // Account 0
	}
	account, _ := wallet.Derive(derivationPath, true)

	executor, err := executableProposal.ToExecutor(make(map[string]executable.ContractDeployBackend)) // TODO: pass in a real backend
	if err != nil {
		log.Fatal(err)
	}

	// Get the signing hash
	payload, err := executor.SigningHash()
	if err != nil {
		log.Fatal(err)
	}

	// Sign the payload
	sig, _ := wallet.SignData(account, "", payload.Bytes())
	unmarshalledSig := executable.Signature{}
	err = json.Unmarshal(sig, &unmarshalledSig)
	if err != nil {
		log.Fatal(err)
	}

	// Add signature to proposal
	proposal.AddSignature(unmarshalledSig)

	// Write proposal to file
	WriteProposalToFile(proposal, os.Args[0])

	// Close wallet
	_ = wallet.Close()
}
