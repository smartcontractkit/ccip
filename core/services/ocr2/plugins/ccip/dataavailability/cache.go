package dataavailability

type DAConfigCache struct {
	destDataAvailabilityOverheadGas   int64 //    Extra data availability gas charged on top of the message, e.g. for OCR
	destGasPerDataAvailabilityByte    int64 //     Amount of gas to charge per byte of message data that needs availability
	destDataAvailabilityMultiplierBps int64 // Multiplier for data availability gas, multiples of bps, or 0.0001
}

func NewDAConfigCache() *DAConfigCache {
	return &DAConfigCache{}
}

func (c *DAConfigCache) Set(destDataAvailabilityOverheadGas, destGasPerDataAvailabilityByte, destDataAvailabilityMultiplierBps int64) {
	c.destGasPerDataAvailabilityByte = destGasPerDataAvailabilityByte
	c.destDataAvailabilityOverheadGas = destDataAvailabilityOverheadGas
	c.destDataAvailabilityMultiplierBps = destDataAvailabilityMultiplierBps
}

func (c *DAConfigCache) Get() (destDataAvailabilityOverheadGas, destGasPerDataAvailabilityByte, destDataAvailabilityMultiplierBps int64) {
	return c.destDataAvailabilityOverheadGas, c.destGasPerDataAvailabilityByte, c.destDataAvailabilityMultiplierBps
}
