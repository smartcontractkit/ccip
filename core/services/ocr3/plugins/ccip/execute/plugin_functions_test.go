package commit

import (
	"testing"

	"github.com/smartcontractkit/ccipocr3/internal/model"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/stretchr/testify/assert"
)

func Test_validateObserverReadingEligibility(t *testing.T) {
	tests := []struct {
		name         string
		observer     commontypes.OracleID
		observerCfg  map[commontypes.OracleID]model.ObserverInfo
		observedMsgs map[model.ChainSelector][]model.ExecutePluginCCIPData
		expErr       bool
	}{
		{
			name:     "ValidObserverAndMessages",
			observer: commontypes.OracleID(1),
			observerCfg: map[commontypes.OracleID]model.ObserverInfo{
				1: {Reads: []model.ChainSelector{1, 2}},
			},
			observedMsgs: map[model.ChainSelector][]model.ExecutePluginCCIPData{
				1: {{}, {}},
				2: {{}, {}},
			},
		},
		{
			name:     "ObserverNotFound",
			observer: commontypes.OracleID(1),
			observerCfg: map[commontypes.OracleID]model.ObserverInfo{
				2: {Reads: []model.ChainSelector{1, 2}},
			},
			observedMsgs: map[model.ChainSelector][]model.ExecutePluginCCIPData{
				1: {{}, {}},
			},
			expErr: true,
		},
		{
			name:     "ObserverNotAllowedToReadChain",
			observer: commontypes.OracleID(1),
			observerCfg: map[commontypes.OracleID]model.ObserverInfo{
				1: {Reads: []model.ChainSelector{1}},
			},
			observedMsgs: map[model.ChainSelector][]model.ExecutePluginCCIPData{
				2: {{}, {}},
			},
			expErr: true,
		},
		{
			name:     "NoMessagesObserved",
			observer: commontypes.OracleID(1),
			observerCfg: map[commontypes.OracleID]model.ObserverInfo{
				1: {Reads: []model.ChainSelector{1, 2}},
			},
			observedMsgs: map[model.ChainSelector][]model.ExecutePluginCCIPData{},
		},
		{
			name:     "EmptyMessagesInChain",
			observer: commontypes.OracleID(1),
			observerCfg: map[commontypes.OracleID]model.ObserverInfo{
				1: {Reads: []model.ChainSelector{1, 2}},
			},
			observedMsgs: map[model.ChainSelector][]model.ExecutePluginCCIPData{
				1: {},
				2: {{}, {}},
			},
		},
		{
			name:     "AllMessagesEmpty",
			observer: commontypes.OracleID(1),
			observerCfg: map[commontypes.OracleID]model.ObserverInfo{
				1: {Reads: []model.ChainSelector{1, 2}},
			},
			observedMsgs: map[model.ChainSelector][]model.ExecutePluginCCIPData{
				1: {},
				2: {},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObserverReadingEligibility(tc.observer, tc.observerCfg, tc.observedMsgs)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_validateObservedSequenceNumbers(t *testing.T) {
	testCases := []struct {
		name         string
		observedData map[model.ChainSelector][]model.ExecutePluginCommitData
		expErr       bool
	}{
		{
			name: "ValidData",
			observedData: map[model.ChainSelector][]model.ExecutePluginCommitData{
				1: {
					{
						MerkleRoot:          model.Bytes32{1},
						SequenceNumberRange: model.SeqNumRange{1, 10},
						ExecutedMessages:    []model.SeqNum{1, 2, 3},
					},
				},
				2: {
					{
						MerkleRoot:          model.Bytes32{2},
						SequenceNumberRange: model.SeqNumRange{11, 20},
						ExecutedMessages:    []model.SeqNum{11, 12, 13},
					},
				},
			},
		},
		{
			name: "DuplicateMerkleRoot",
			observedData: map[model.ChainSelector][]model.ExecutePluginCommitData{
				1: {
					{
						MerkleRoot:          model.Bytes32{1},
						SequenceNumberRange: model.SeqNumRange{1, 10},
						ExecutedMessages:    []model.SeqNum{1, 2, 3},
					},
					{
						MerkleRoot:          model.Bytes32{1},
						SequenceNumberRange: model.SeqNumRange{11, 20},
						ExecutedMessages:    []model.SeqNum{11, 12, 13},
					},
				},
			},
			expErr: true,
		},
		{
			name: "OverlappingSequenceNumberRange",
			observedData: map[model.ChainSelector][]model.ExecutePluginCommitData{
				1: {
					{
						MerkleRoot:          model.Bytes32{1},
						SequenceNumberRange: model.SeqNumRange{1, 10},
						ExecutedMessages:    []model.SeqNum{1, 2, 3},
					},
					{
						MerkleRoot:          model.Bytes32{2},
						SequenceNumberRange: model.SeqNumRange{5, 15},
						ExecutedMessages:    []model.SeqNum{6, 7, 8},
					},
				},
			},
			expErr: true,
		},
		{
			name: "ExecutedMessageOutsideObservedRange",
			observedData: map[model.ChainSelector][]model.ExecutePluginCommitData{
				1: {
					{
						MerkleRoot:          model.Bytes32{1},
						SequenceNumberRange: model.SeqNumRange{1, 10},
						ExecutedMessages:    []model.SeqNum{1, 2, 11},
					},
				},
			},
			expErr: true,
		},
		{
			name: "NoCommitData",
			observedData: map[model.ChainSelector][]model.ExecutePluginCommitData{
				1: {},
			},
		},
		{
			name:         "EmptyObservedData",
			observedData: map[model.ChainSelector][]model.ExecutePluginCommitData{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObservedSequenceNumbers(tc.observedData)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
