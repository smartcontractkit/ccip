package types

// ChainSide represents the two sides of a CCIP plugin,
// the source and the destination chains.
type ChainSide string

const (
	// ChainSideSource represents the source chain of a CCIP plugin.
	ChainSideSource ChainSide = "source"
	// ChainSideDest represents the destination chain of a CCIP plugin.
	ChainSideDest ChainSide = "dest"
)
