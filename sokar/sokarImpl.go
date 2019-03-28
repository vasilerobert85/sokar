package sokar

import (
	"time"

	sokarIF "github.com/thomasobenaus/sokar/sokar/iface"
)

func (sk *Sokar) scaleEventProcessor(scaleEventChannel <-chan sokarIF.ScaleEvent) {
	sk.wg.Add(1)
	defer sk.wg.Done()

	sk.logger.Info().Msg("ScaleEventProcessor started.")

	for {
		select {
		case <-sk.stopChan:
			sk.logger.Info().Msg("ScaleEventProcessor stopped.")
			return
		case se := <-scaleEventChannel:
			sk.handleScaleEvent(se)
		}
	}
}

func (sk *Sokar) handleScaleEvent(scaleEvent sokarIF.ScaleEvent) {
	sk.logger.Info().Msgf("Scale Event received: %v", scaleEvent)

	sk.metrics.scaleEventsTotal.Inc()
	sk.metrics.scaleFactor.Set(float64(scaleEvent.ScaleFactor))

	preScaleJobCount, err := sk.scaler.GetCount()
	if err != nil {
		sk.metrics.failedScalingTotal.Inc()
		sk.logger.Error().Err(err).Msg("Scaling ignored. Failed to obtain current count.")
		return
	}
	sk.metrics.preScaleJobCount.Set(float64(preScaleJobCount))

	// Don't scale if sokar is in cool down mode
	if sk.capacityPlanner.IsCoolingDown(sk.lastScaleAction, scaleEvent.ScaleFactor) {
		sk.metrics.skippedScalingDuringCooldownTotal.Inc()
		sk.logger.Info().Msg("Skip scale event. Sokar is cooling down.")
		return
	}

	// plan
	plannedJobCount := sk.capacityPlanner.Plan(scaleEvent.ScaleFactor, preScaleJobCount)
	sk.metrics.plannedJobCount.Set(float64(plannedJobCount))
	err = sk.scaler.ScaleTo(plannedJobCount)

	// HACK: For now we ignore all rejected scaling tickets
	if err != nil {
		sk.metrics.failedScalingTotal.Inc()
		sk.logger.Error().Err(err).Msg("Failed to scale.")
	}

	sk.lastScaleAction = time.Now()
	sk.logger.Info().Uint("preScaleCnt", preScaleJobCount).Uint("plannedCnt", plannedJobCount).Msg("Scaling triggered. Scaler will apply the planned count.")
}
