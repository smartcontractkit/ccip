package ccipdata

type DAConfigCacheWriter interface {
	Set(destDataAvailabilityOverheadGas, destGasPerDataAvailabilityByte, destDataAvailabilityMultiplierBps int64)
}

type DAConfigCacheReader interface {
	Get() (destDataAvailabilityOverheadGas, destGasPerDataAvailabilityByte, destDataAvailabilityMultiplierBps int64)
}
