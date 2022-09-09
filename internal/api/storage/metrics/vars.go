package metrics

import "expvar"

var (
	CacheHitCounter  = expvar.NewInt("storage.cacheHit")
	CacheMissCounter = expvar.NewInt("storage.cacheMiss")
)
