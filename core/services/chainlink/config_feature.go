package chainlink

import v2 "github.com/smartcontractkit/chainlink/v2/core/config/v2"

type featureConfig struct {
	c v2.Feature
}

func (f *featureConfig) CCIP() bool {
	return *f.c.CCIP
}

func (f *featureConfig) LegacyGasStation() bool {
	return *f.c.LegacyGasStation
}

func (f *featureConfig) FeedsManager() bool {
	return *f.c.FeedsManager
}

func (f *featureConfig) LogPoller() bool {
	return *f.c.LogPoller
}

func (f *featureConfig) UICSAKeys() bool {
	return *f.c.UICSAKeys
}
