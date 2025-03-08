package schedulers

import (
	"github.com/xSaCh/xginx/pkg"
)

type SchedulerAlgorithm string

const (
	SCHEDULER_ROUND_ROBIN SchedulerAlgorithm = "round_robin"
)

type Scheduler interface {
	GetNextBackend() *pkg.Backend
	T() *pkg.ServerPool
}

func NewScheduler(s SchedulerAlgorithm, pool *pkg.ServerPool) Scheduler {
	switch s {
	case SCHEDULER_ROUND_ROBIN:
		return NewRoundRobin(pool)
	default:
		return nil
	}
}
