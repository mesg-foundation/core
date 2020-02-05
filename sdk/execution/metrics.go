package executionsdk

import (
	"github.com/go-kit/kit/metrics"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var m *metric

type metric struct {
	Created    metrics.Counter
	InProgress metrics.Counter
	Updated    metrics.Counter
	Completed  metrics.Counter
}

func newMetric() *metric {
	return &metric{
		Created: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "mesg",
			Subsystem: "execution",
			Name:      "created",
			Help:      "executions created",
		}, []string{}),
		InProgress: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "mesg",
			Subsystem: "execution",
			Name:      "in_progress",
			Help:      "executions in progress",
		}, []string{}),
		Updated: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "mesg",
			Subsystem: "execution",
			Name:      "updated",
			Help:      "executions updated",
		}, []string{}),
		Completed: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "mesg",
			Subsystem: "execution",
			Name:      "completed",
			Help:      "executions completed",
		}, []string{}),
	}
}

func init() {
	m = newMetric()
}
