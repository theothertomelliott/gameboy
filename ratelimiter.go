package gameboy

import (
	"time"
)

const DefaultCyclesPerSecond = 4194304

func NewDefaultRateLimiter() *RateLimiter {
	syncPeriod := 10 * time.Millisecond
	cyclesPerPeriod := float64(DefaultCyclesPerSecond) * (float64(syncPeriod) / float64(time.Second))

	return &RateLimiter{
		cyclesPerPeriod: cyclesPerPeriod,
		syncPeriod:      syncPeriod,
		lastSync:        time.Now(),
	}
}

type RateLimiter struct {
	syncPeriod      time.Duration
	cyclesPerPeriod float64
	cycleCount      int
	lastSync        time.Time
}

func (r *RateLimiter) Increment(cycles int) {
	now := time.Now()
	r.cycleCount += cycles

	if float64(r.cycleCount) < r.cyclesPerPeriod {
		return
	}

	defer func() {
		r.lastSync = now
		r.cycleCount = 0
	}()

	delay := r.syncPeriod - now.Sub(r.lastSync)
	if delay < 0 {
		return
	}
	time.Sleep(delay)

}
