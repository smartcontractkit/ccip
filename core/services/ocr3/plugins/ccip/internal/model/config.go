package model

import "time"

type HomeChainConfig struct{}

type CommitPluginConfig struct {
	Writer             bool
	Reads              []ChainSelector
	NewMsgScanDuration time.Duration
	NewMsgScanLimit    int
}
