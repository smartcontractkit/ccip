package model

import (
	"encoding/json"
	"fmt"
)

type ExecutePluginReport struct {
	ChainReports []ExecutionPluginReportSingleChain `json:"chainReports"`
}

type ExecutionPluginReportSingleChain struct {
	SourceChainSelector ChainSelector    `json:"sourceChainSelector"`
	Messages            []Evm2EvmMessage `json:"messages"`
	OffchainTokenData   [][][]byte       `json:"offchainTokenData"`
	Proofs              []Bytes32        `json:"proofs"`
	ProofFlagBits       BigInt           `json:"proofFlagBits"`
}

/////////////////////////
// Execute Observation //
/////////////////////////

type ExecutePluginObservation struct {
	CommitReports map[ChainSelector][]ExecutePluginCommitData `json:"commitReports"`
	Messages      map[ChainSelector][]ExecutePluginCCIPData   `json:"messages"`
}

type ExecutePluginCCIPData struct {
	SequenceNumber SeqNum  `json:"sequenceNumber"`
	Message        Bytes32 `json:"message"`
}

type ExecutePluginCommitData struct {
	MerkleRoot          Bytes32     `json:"merkleRoot"`
	SequenceNumberRange SeqNumRange `json:"sequenceNumberRange"`
	ExecutedMessages    []SeqNum    `json:"executed"`
}

func NewExecutePluginObservation(commitReports map[ChainSelector][]ExecutePluginCommitData, messages map[ChainSelector][]ExecutePluginCCIPData) ExecutePluginObservation {
	return ExecutePluginObservation{
		CommitReports: commitReports,
		Messages:      messages,
	}
}

func (obs ExecutePluginObservation) Encode() ([]byte, error) {
	return json.Marshal(obs)
}

func DecodeExecutePluginObservation(b []byte) (ExecutePluginObservation, error) {
	obs := ExecutePluginObservation{}
	err := json.Unmarshal(b, &obs)
	return obs, err
}

/////////////////////
// Execute Outcome //
/////////////////////

type ExecutePluginOutcome struct {
	NextCommits map[ChainSelector][]ExecutePluginCommitData `json:"nextCommits"`
	Messages    map[ChainSelector]map[SeqNum][]byte
}

func NewExecutePluginOutcome(
	nextCommits map[ChainSelector][]ExecutePluginCommitData,
	messages map[ChainSelector]map[SeqNum][]byte,
) ExecutePluginOutcome {
	return ExecutePluginOutcome{
		NextCommits: nextCommits,
		Messages:    messages,
	}
}

func (o ExecutePluginOutcome) Encode() ([]byte, error) {
	return json.Marshal(o)
}

func DecodeExecutePluginOutcome(b []byte) (ExecutePluginOutcome, error) {
	o := ExecutePluginOutcome{}
	err := json.Unmarshal(b, &o)
	return o, err
}

func (o ExecutePluginOutcome) String() string {
	return fmt.Sprintf("NextCommits: %v, Messages: %v", o.NextCommits, o.Messages)
}
