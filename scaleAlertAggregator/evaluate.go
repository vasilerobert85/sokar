package scaleAlertAggregator

import (
	"fmt"
	"time"
)

// isScalingNeeded returns true if the current scaleCounter violates either the upScaling-
// or downScaling threshold
func (sc *ScaleAlertAggregator) isScalingNeeded() bool {
	scaleUpNeeded := sc.scaleCounter > sc.upScalingThreshold
	scaleDownNeeded := sc.scaleCounter < sc.downScalingThreshold

	return scaleDownNeeded || scaleUpNeeded
}

// gradientToScaleDir returns the scaling direction based on the given gradient
func gradientToScaleDir(gradient float32) string {
	result := "NO"

	if gradient < 0 {
		result = "DOWN"
	}
	if gradient > 0 {
		result = "UP"
	}
	return result
}

func (sc *ScaleAlertAggregator) evaluate() float32 {
	sc.evaluationCounter++
	now := time.Now()

	gradientRefreshCause := fmt.Sprintf("Evaluation period (%fs) exceeded.", float64(sc.evaluationPeriodFactor)*sc.evaluationCycle.Seconds())
	var gradient float32
	if sc.isScalingNeeded() {
		gradient = sc.scaleCounterGradient.UpdateAndGet(sc.scaleCounter, now)
		scaleDir := gradientToScaleDir(gradient)
		sc.logger.Info().Str("sDir", scaleDir).Float32("sCnt", sc.scaleCounter).Float32("upThrs", sc.upScalingThreshold).Float32("downTrhs", sc.downScalingThreshold).Float32("grad", gradient).Msgf("Scale %s.", scaleDir)

		// reset the scaleCounter
		sc.scaleCounter = 0
		gradientRefreshCause = "Scale needed."

		// restart evaluation counter to force a reset of the gradient
		sc.evaluationCounter = 0

		sc.metrics.scaleEventCounter.WithLabelValues(scaleDir).Inc()
	} else {
		gr := sc.scaleCounterGradient.Get(sc.scaleCounter, now)
		sc.logger.Debug().Float32("sCnt", sc.scaleCounter).Float32("upThrs", sc.upScalingThreshold).Float32("downTrhs", sc.downScalingThreshold).Float32("grad", gr).Msg("No scale needed.")
	}

	// Reset the gradient if the evaluation was exceeded.
	if sc.evaluationCounter%sc.evaluationPeriodFactor == 0 {
		gr := sc.scaleCounterGradient.UpdateAndGet(sc.scaleCounter, now)
		sc.logger.Debug().Msgf("Refresh gradient %f. %s", gr, gradientRefreshCause)
	}

	sc.metrics.scaleFactor.Set(float64(gradient))
	return gradient
}
