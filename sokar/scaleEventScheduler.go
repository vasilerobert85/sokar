package sokar

import (
	"time"

	"github.com/thomasobenaus/sokar/helper"
	sokarIF "github.com/thomasobenaus/sokar/sokar/iface"
)

func (sk *Sokar) scaleEventScheduler(scaleEventChannel chan sokarIF.ScaleEvent) {
	sk.wg.Add(1)
	defer sk.wg.Done()

	evaluationTicker := time.NewTicker(sk.scheduledScaleEventCycle)

	sk.logger.Info().Msg("ScaleEventScheduler started.")

	for {
		select {
		case <-sk.stopChan:
			sk.logger.Info().Msg("ScaleEventScheduler stopped.")
			return
		case <-evaluationTicker.C:
			if sk.shouldFireScaleEvent(time.Now()) {
				sk.logger.Info().Msg("Trigger a scheduled scale event. Forces to evaluate if the current scale meets the schedule.")
				scaleEventChannel <- sokarIF.NewScheduledScaleEvent()
			} else {
				sk.logger.Debug().Msg("No need to fire a scheduled scale event. Currently there is no active schedule.")
			}
		}
	}
}

// shouldFireScaleEvent checks if, according to the current Schedule, a ScaleEvent shall be fired
func (sk *Sokar) shouldFireScaleEvent(now time.Time) bool {
	day := now.Weekday()
	hour := now.Hour()
	minute := now.Minute()
	at, err := helper.NewTime(uint(hour), uint(minute))

	if err != nil {
		sk.logger.Warn().Msgf("Could not evaluate scaling schedule: %s", err.Error())
		return false
	}

	// Just for metrics
	minScale, maxScale, err := sk.schedule.ScaleRangeAt(day, at)
	if err != nil {
		minScale = 0
		maxScale = 0
	}
	sk.metrics.scheduledScaleBounds.WithLabelValues("min").Set(float64(minScale))
	sk.metrics.scheduledScaleBounds.WithLabelValues("max").Set(float64(maxScale))

	return sk.schedule.IsActiveAt(day, at)
}
