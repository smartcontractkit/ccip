package config

// CommitPluginConfig contains the plugin specific variables for the ccip.CCIPCommit plugin.
// We use ID here to keep it as general as possible, e.g. abstracting for chains which don't have an address concept.
type CommitPluginConfig struct {
	SourceStartBlock, DestStartBlock int64 // Only for first time job add.
	// TokenPricesUSDPipeline should contain a token price pipeline for the following tokens:
	//		The SOURCE chain wrapped native
	// 		The DESTINATION fee tokens as defined in the destination chain PriceRegistry.
	TokenPricesUSDPipeline string `json:"tokenPricesUSDPipeline"`
}

// ExecutionPluginConfig contains the plugin specific variables for the ccip.CCIPExecution plugin.
type ExecutionPluginConfig struct {
	SourceStartBlock, DestStartBlock int64 // Only for first time job add.
}
