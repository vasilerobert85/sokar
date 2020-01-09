package sokar

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	m "github.com/thomasobenaus/sokar/metrics"
)

// Metrics represents the collection of metrics internally set by sokar.
type Metrics struct {
	scheduledScaleBounds              m.GaugeVec
	scaleEventsTotal                  m.Counter
	failedScalingTotal                m.Counter
	skippedScalingDuringCooldownTotal m.Counter
	preScaleJobCount                  m.Gauge
	plannedJobCount                   m.Gauge
	scaleFactor                       m.Gauge
}

// NewMetrics returns the metrics collection needed for the SAA.
func NewMetrics() Metrics {

	bound := []string{"bound"}
	scheduledScaleBounds := m.NewWrappedGaugeVec(prometheus.GaugeOpts{
		Namespace: "sokar",
		Subsystem: "cap",
		Name:      "scheduled_scale_bounds",
		Help:      "Shows the min and max scale value of the currently active scale schedule. In case no schedule is active both values are 0.",
	}, bound)

	scaleEventsTotal := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "sokar",
		Name:      "scale_events_total",
		Help:      "Number of received ScaleEvents in total.",
	})

	failedScalingTotal := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "sokar",
		Name:      "failed_scaling_total",
		Help:      "Number of failed scaling actions in total.",
	})

	skippedScalingDuringCooldownTotal := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "sokar",
		Subsystem: "cap",
		Name:      "skipped_scaling_during_cooldown_total",
		Help:      "Number of scalings that where not executed since sokar was cooling down.",
	})

	preScaleJobCount := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "sokar",
		Subsystem: "cap",
		Name:      "pre_scale_job_count",
		Help:      "The scaling-object count before the scaling action. Based on this count sokar does the planning.",
	})

	plannedJobCount := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "sokar",
		Subsystem: "cap",
		Name:      "planned_job_count",
		Help:      "The count planned by the CapacityPlanner for the current scale action.",
	})

	scaleFactor := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "sokar",
		Subsystem: "cap",
		Name:      "scale_factor",
		Help:      "The scale factor (gradient) as it was received with a ScalingEvent.",
	})

	return Metrics{
		scheduledScaleBounds:              scheduledScaleBounds,
		scaleEventsTotal:                  scaleEventsTotal,
		failedScalingTotal:                failedScalingTotal,
		skippedScalingDuringCooldownTotal: skippedScalingDuringCooldownTotal,
		preScaleJobCount:                  preScaleJobCount,
		plannedJobCount:                   plannedJobCount,
		scaleFactor:                       scaleFactor,
	}
}
