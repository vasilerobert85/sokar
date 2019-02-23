package sokar

// Run starts sokar
func (sk *Sokar) Run() {

	scaleEventChannel := make(chan ScaleEvent, 10)
	sk.scaleEventAggregator.Subscribe(scaleEventChannel)

	// main loop
	go func() {
		sk.logger.Info().Msg("Sokar main loop started")

	loop:
		for {
			select {
			case <-sk.stopChan:
				// send the stop message a second time to complete waiting join calls
				sk.stopChan <- struct{}{}
				break loop

			case se := <-scaleEventChannel:
				sk.handleScaleEvent(se)

			}
		}
		sk.logger.Info().Msg("Sokar main loop left")
	}()

}

// Stop tears down sokar
func (sk *Sokar) Stop() {
	sk.logger.Info().Msg("Teardown requested")
	// send the stop message
	sk.stopChan <- struct{}{}
}

// Join blocks/ waits until sokar has been stopped
func (sk *Sokar) Join() {
	<-sk.stopChan
}

func (sk *Sokar) handleScaleEvent(scaleEvent ScaleEvent) {

	sk.logger.Info().Msgf("SCALE-EVENT TRIGGERED: %v", scaleEvent)

	// plan
	plannedCount := sk.capacityPlanner.Plan(scaleEvent.ScaleFactor, 1)
	sk.scaler.ScaleBy(int(plannedCount))
}