package model

import (
	"fmt"

	chainselectors "github.com/smartcontractkit/chain-selectors"
)

type SeqNum uint64

type SeqNumRange [2]SeqNum

func (s SeqNumRange) String() string {
	return fmt.Sprintf("[%d -> %d]", s[0], s[1])
}

type ChainSelector uint64

func (c ChainSelector) String() string {
	ch, exists := chainselectors.ChainBySelector(uint64(c))
	if !exists || ch.Name == "" {
		return fmt.Sprintf("ChainSelector(%d)", c)
	}
	return fmt.Sprintf("%d (%s)", c, ch.Name)
}

type NodeID string

type CCIPMsg struct {
	CCIPMsgBaseDetails
}

func (c CCIPMsg) String() string {
	return fmt.Sprintf("%#v", c)
}

type CCIPMsgBaseDetails struct {
	SourceChain ChainSelector
	SeqNum      SeqNum
}
