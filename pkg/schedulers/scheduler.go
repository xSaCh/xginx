package schedulers

import (
	"github.com/xSaCh/xginx/pkg"
)

type Schedulers string

const (
	SCHEDULER_ROUND_ROBIN Schedulers = "ROUND_ROBIN"
)

type Scheduler interface {
	GetNextBackend() *pkg.Backend
	T() *pkg.ServerPool
}

func NewScheduler(s Schedulers, pool *pkg.ServerPool) Scheduler {
	switch s {
	case SCHEDULER_ROUND_ROBIN:
		return NewRoundRobin(pool)
	default:
		return nil
	}
}
