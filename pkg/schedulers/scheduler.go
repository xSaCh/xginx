package schedulers

import (
	"fmt"

	"github.com/xSaCh/xginx/pkg"
	"gopkg.in/yaml.v3"
)

type SchedulerAlgorithm string

const (
	SCHEDULER_ROUND_ROBIN SchedulerAlgorithm = "round_robin"
)

func ToSchedulerAlogrithm(raw string) (SchedulerAlgorithm, error) {
	switch raw {
	case "round_robin":
		return SCHEDULER_ROUND_ROBIN, nil
	default:
		return "", fmt.Errorf("invalid Scheduler value: %s", raw)
	}
}

func (a *SchedulerAlgorithm) UnmarshalYAML(value *yaml.Node) error {
	var algoStr string
	if err := value.Decode(&algoStr); err != nil {
		return err
	}

	sch, err := ToSchedulerAlogrithm(algoStr)
	if err != nil {
		return err
	}

	*a = sch
	return nil
}

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
