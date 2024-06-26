package types

import (
	ocr3reader "github.com/smartcontractkit/ccipocr3/pkg/reader"
)

type OCR3ConfigWithMeta ocr3reader.OCR3ConfigWithMeta

type PluginType uint8

const (
	PluginTypeCCIPCommit PluginType = 0
	PluginTypeCCIPExec   PluginType = 1
)

// CCIPOracle represents either a CCIP commit or exec oracle or a bootstrap node.
//
//go:generate mockery --name CCIPOracle --output ./mocks/ --case underscore
type CCIPOracle interface {
	Close() error
	Start() error
}

// OracleCreator is an interface for creating CCIP oracles.
// Whether the oracle uses a LOOPP or not is an implementation detail.
//
//go:generate mockery --name OracleCreator --output ./mocks/ --case underscore
type OracleCreator interface {
	// CreateCommitOracle creates a new oracle that will run the CCIP commit plugin.
	// The oracle must be returned unstarted.
	CreateCommitOracle(config OCR3ConfigWithMeta) (CCIPOracle, error)

	// CreateExecOracle creates a new oracle that will run the CCIP exec plugin.
	// The oracle must be returned unstarted.
	CreateExecOracle(config OCR3ConfigWithMeta) (CCIPOracle, error)

	// CreateBootstrapOracle creates a new bootstrap node with the given OCR config.
	// The oracle must be returned unstarted.
	CreateBootstrapOracle(config OCR3ConfigWithMeta) (CCIPOracle, error)
}
