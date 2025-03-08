package schedulers

import (
	"sync/atomic"

	"github.com/xSaCh/xginx/pkg"
)

type RoundRobin struct {
	pool           *pkg.ServerPool
	currentBackend uint64
}

func NewRoundRobin(pool *pkg.ServerPool) *RoundRobin {
	return &RoundRobin{pool: pool}
}

func (rr *RoundRobin) GetNextBackend() *pkg.Backend {
	n := atomic.AddUint64(&rr.currentBackend, 1) % uint64(len(rr.pool.Servers))
	l := uint64(len(rr.pool.Servers)) + n
	for ; n < l; n++ {
		idx := n % uint64(len(rr.pool.Servers))
		if rr.pool.Servers[idx].IsAlive() {
			atomic.StoreUint64(&rr.currentBackend, idx)
			return rr.pool.Servers[idx]
		}
	}
	return nil
}
