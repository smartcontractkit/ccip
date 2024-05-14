package commit

import (
	"context"
	"reflect"
	"slices"
	"testing"

	"github.com/smartcontractkit/ccipocr3/internal/mocks"
	"github.com/smartcontractkit/ccipocr3/internal/model"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

func TestPlugin_observeMaxSeqNumsPerChain(t *testing.T) {
	testCases := []struct {
		name           string
		prevOutcome    model.CommitPluginOutcome
		onChainSeqNums map[model.ChainSelector]model.SeqNum
		readChains     []model.ChainSelector
		destChain      model.ChainSelector
		expErr         bool
		expMaxSeqNums  []model.SeqNumChain
	}{
		{
			name:        "report on chain seq num when no previous outcome and can read dest",
			prevOutcome: model.CommitPluginOutcome{},
			onChainSeqNums: map[model.ChainSelector]model.SeqNum{
				1: 10,
				2: 20,
			},
			readChains: []model.ChainSelector{1, 2, 3},
			destChain:  3,
			expErr:     false,
			expMaxSeqNums: []model.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
		},
		{
			name:        "nothing to report when there is no previous outcome and cannot read dest",
			prevOutcome: model.CommitPluginOutcome{},
			onChainSeqNums: map[model.ChainSelector]model.SeqNum{
				1: 10,
				2: 20,
			},
			readChains:    []model.ChainSelector{1, 2},
			destChain:     3,
			expErr:        false,
			expMaxSeqNums: []model.SeqNumChain{},
		},
		{
			name: "report previous outcome seq nums and override when on chain is higher if can read dest",
			prevOutcome: model.CommitPluginOutcome{
				MaxSeqNums: []model.SeqNumChain{
					{ChainSel: 1, SeqNum: 11}, // for chain 1 previous outcome is higher than on-chain state
					{ChainSel: 2, SeqNum: 19}, // for chain 2 previous outcome is behind on-chain state
				},
			},
			onChainSeqNums: map[model.ChainSelector]model.SeqNum{
				1: 10,
				2: 20,
			},
			readChains: []model.ChainSelector{1, 2, 3},
			destChain:  3,
			expErr:     false,
			expMaxSeqNums: []model.SeqNumChain{
				{ChainSel: 1, SeqNum: 11},
				{ChainSel: 2, SeqNum: 20},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockReader := mocks.NewCCIPReader()
			lggr := logger.Test(t)

			p := NewPlugin(
				ctx,
				commontypes.OracleID(123),
				model.CommitPluginConfig{
					Writer:              false,
					Reads:               tc.readChains,
					DestChain:           tc.destChain,
					FChain:              map[model.ChainSelector]int{},
					ObserverInfo:        nil,
					NewMsgScanDuration:  0,
					NewMsgScanLimit:     0,
					NewMsgScanBatchSize: 0,
				},
				mockReader,
				nil,
				lggr,
			)

			var b []byte
			var err error
			if !reflect.DeepEqual(tc.prevOutcome, model.CommitPluginOutcome{}) {
				b, err = tc.prevOutcome.Encode()
				assert.NoError(t, err)
			}

			knownChainsSlice := p.knownSourceChainsSlice()
			onChainSeqNums := make([]model.SeqNum, 0)
			for _, chain := range knownChainsSlice {
				if v, ok := tc.onChainSeqNums[chain]; !ok {
					t.Fatalf("invalid test case missing on chain seq num expectation for %d", chain)
				} else {
					onChainSeqNums = append(onChainSeqNums, v)
				}
			}
			mockReader.On("NextSeqNum", ctx, knownChainsSlice).Return(onChainSeqNums, nil)

			seqNums, err := p.observeMaxSeqNumsPerChain(ctx, b)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expMaxSeqNums, seqNums)
		})
	}
}

func TestPlugin_observeNewMsgs(t *testing.T) {
	testCases := []struct {
		name               string
		maxSeqNumsPerChain []model.SeqNumChain
		readChains         []model.ChainSelector
		destChain          model.ChainSelector
		msgScanBatchSize   int
		newMsgs            map[model.ChainSelector][]model.CCIPMsg
		expMsgs            []model.CCIPMsgBaseDetails
		expErr             bool
	}{
		{
			name: "no new messages",
			maxSeqNumsPerChain: []model.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			readChains:       []model.ChainSelector{1, 2},
			msgScanBatchSize: 256,
			newMsgs: map[model.ChainSelector][]model.CCIPMsg{
				1: {},
				2: {},
			},
			expMsgs: []model.CCIPMsgBaseDetails{},
			expErr:  false,
		},
		{
			name: "new messages",
			maxSeqNumsPerChain: []model.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			readChains:       []model.ChainSelector{1, 2},
			msgScanBatchSize: 256,
			newMsgs: map[model.ChainSelector][]model.CCIPMsg{
				1: {
					{CCIPMsgBaseDetails: model.CCIPMsgBaseDetails{ID: [32]byte{1}, SourceChain: 1, SeqNum: 11}},
				},
				2: {
					{CCIPMsgBaseDetails: model.CCIPMsgBaseDetails{ID: [32]byte{2}, SourceChain: 2, SeqNum: 21}},
					{CCIPMsgBaseDetails: model.CCIPMsgBaseDetails{ID: [32]byte{3}, SourceChain: 2, SeqNum: 22}},
				},
			},
			expMsgs: []model.CCIPMsgBaseDetails{
				{ID: [32]byte{1}, SourceChain: 1, SeqNum: 11},
				{ID: [32]byte{2}, SourceChain: 2, SeqNum: 21},
				{ID: [32]byte{3}, SourceChain: 2, SeqNum: 22},
			},
			expErr: false,
		},
		{
			name: "new messages but one chain is not readable",
			maxSeqNumsPerChain: []model.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			readChains:       []model.ChainSelector{2},
			msgScanBatchSize: 256,
			newMsgs: map[model.ChainSelector][]model.CCIPMsg{
				2: {
					{CCIPMsgBaseDetails: model.CCIPMsgBaseDetails{ID: [32]byte{2}, SourceChain: 2, SeqNum: 21}},
					{CCIPMsgBaseDetails: model.CCIPMsgBaseDetails{ID: [32]byte{3}, SourceChain: 2, SeqNum: 22}},
				},
			},
			expMsgs: []model.CCIPMsgBaseDetails{
				{ID: [32]byte{2}, SourceChain: 2, SeqNum: 21},
				{ID: [32]byte{3}, SourceChain: 2, SeqNum: 22},
			},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockReader := mocks.NewCCIPReader()
			lggr := logger.Test(t)

			p := NewPlugin(
				ctx,
				commontypes.OracleID(123),
				model.CommitPluginConfig{
					Writer:              false,
					Reads:               tc.readChains,
					DestChain:           tc.destChain,
					FChain:              map[model.ChainSelector]int{},
					ObserverInfo:        nil,
					NewMsgScanDuration:  0,
					NewMsgScanLimit:     0,
					NewMsgScanBatchSize: tc.msgScanBatchSize,
				},
				mockReader,
				nil,
				lggr,
			)

			for _, seqNumChain := range tc.maxSeqNumsPerChain {
				if slices.Contains(tc.readChains, seqNumChain.ChainSel) {
					mockReader.On(
						"MsgsBetweenSeqNums",
						ctx,
						[]model.ChainSelector{seqNumChain.ChainSel},
						model.NewSeqNumRange(seqNumChain.SeqNum+1, seqNumChain.SeqNum+model.SeqNum(1+tc.msgScanBatchSize)),
					).Return(tc.newMsgs[seqNumChain.ChainSel], nil)
				}
			}

			msgs, err := p.observeNewMsgs(ctx, tc.maxSeqNumsPerChain)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expMsgs, msgs)
			mockReader.AssertExpectations(t)
		})
	}
}
