package model

import "encoding/json"

// CommitPluginReport is placed here for reference of shared readers structure.
type CommitPluginReport struct{}

type CommitPluginObservation struct {
	NodeID      NodeID               `json:"nodeID"`
	NewMsgs     []CCIPMsgBaseDetails `json:"newMsgs"`
	GasPrices   []GasPriceChain      `json:"gasPrices,string"`
	TokenPrices []TokenPrice         `json:"tokenPrices"`
}

func NewCommitPluginObservation(
	nodeID NodeID,
	newMsgs []CCIPMsgBaseDetails,
	gasPrices []GasPriceChain,
	tokenPrices []TokenPrice,
) CommitPluginObservation {
	return CommitPluginObservation{
		NodeID:      nodeID,
		NewMsgs:     newMsgs,
		GasPrices:   gasPrices,
		TokenPrices: tokenPrices,
	}
}

func (obs CommitPluginObservation) Encode() ([]byte, error) {
	return json.Marshal(obs)
}

func DecodeCommitPluginObservation(b []byte) (CommitPluginObservation, error) {
	obs := CommitPluginObservation{}
	err := json.Unmarshal(b, &obs)
	return obs, err
}

type CommitPluginOutcome struct {
	SequenceNumbers []SeqNumChain `json:"sequenceNumbers"`
}

func NewCommitPluginOutcome() CommitPluginOutcome {
	return CommitPluginOutcome{}
}

func (o CommitPluginOutcome) Encode() ([]byte, error) {
	return json.Marshal(o)
}

func DecodeCommitPluginOutcome(b []byte) (CommitPluginOutcome, error) {
	o := CommitPluginOutcome{}
	err := json.Unmarshal(b, &o)
	return o, err
}

type SeqNumChain struct {
	ChainSel ChainSelector
	SeqNum   SeqNum
}
