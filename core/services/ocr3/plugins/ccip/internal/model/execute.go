package model

import (
	"encoding/json"
	"fmt"
)

// ExecutePluginReport is placed here for reference of shared readers structure.
type ExecutePluginReport struct{}

/////////////////////////
// Execute Observation //
/////////////////////////

type ExecutePluginObservation struct {
	CommitReports map[ChainSelector][]CommitReport `json:"CommitReports"`
	Messages      map[ChainSelector][]CCIPMessage  `json:"messages"`
}

type CCIPMessage struct {
	SequenceNumber SeqNum `json:"sequenceNumber"`
	Message        []byte `json:"message"`
}
type CommitReport struct {
	MerkleRoot          []byte      `json:"merkleRoot"`
	SequenceNumberRange SeqNumRange `json:"sequenceNumberRange"`
	ExecutedMessages    []SeqNum    `json:"executed"`
}

func NewExecutePluginObservation(commitReports map[ChainSelector][]CommitReport, messages map[ChainSelector][]CCIPMessage) ExecutePluginObservation {
	return ExecutePluginObservation{
		CommitReports: commitReports,
		Messages:      messages,
	}
}

func (obs ExecutePluginObservation) Encode() ([]byte, error) {
	return json.Marshal(obs)
}

func DecodeExecutePluginObservation(b []byte) (CommitPluginObservation, error) {
	obs := CommitPluginObservation{}
	err := json.Unmarshal(b, &obs)
	return obs, err
}

/////////////////////
// Execute Outcome //
/////////////////////

type ExecutePluginOutcome struct {
	NextCommits map[ChainSelector][]CommitReport `json:"nextCommits"`
	Messages    map[ChainSelector]map[SeqNum][]byte
}

func NewExecutePluginOutcome(
	nextCommits map[ChainSelector][]CommitReport,
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
