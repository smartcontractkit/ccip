package model

import "encoding/json"

// ExecutePluginReport is placed here for reference of shared readers structure.
type ExecutePluginReport struct{}

type ExecutePluginObservation struct {
	NodeID  NodeID               `json:"nodeID"`
	NewMsgs []CCIPMsgBaseDetails `json:"newMsgs"`
}

func NewExecutePluginObservation(nodeID NodeID, newMsgs []CCIPMsgBaseDetails) ExecutePluginObservation {
	return ExecutePluginObservation{
		NodeID:  nodeID,
		NewMsgs: newMsgs,
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
