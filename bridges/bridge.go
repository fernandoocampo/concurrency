package bridges

func GenValues(done <-chan any, values []int) <-chan (<-chan any) {
	chanStream := make(chan (<-chan any))
	go func() {
		defer close(chanStream)

		for _, v := range values {
			valueStream := make(chan any, 1)
			valueStream <- v
			close(valueStream)
			chanStream <- valueStream
		}
	}()
	return chanStream
}

func Bridge(done <-chan any, chanStream <-chan (<-chan any)) <-chan any {
	bridgeStream := make(chan any)
	go func() {
		defer close(bridgeStream)

		for {
			var stream <-chan any

			select {
			case <-done:
				return
			case maybeStream, ok := <-chanStream:
				if !ok {
					return
				}

				stream = maybeStream
			}
			// read values off stream and add them to bridge stream
			// once the stream is closed we continue with the other
			// channels.
			for val := range stream {
				select {
				case <-done:
				case bridgeStream <- val:
				}
			}
		}
	}()
	return bridgeStream
}
