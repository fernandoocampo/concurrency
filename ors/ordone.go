package ors

func OrDone(done <-chan struct{}, input <-chan any) <-chan any {
	resultStream := make(chan any)
	go func() {
		defer close(resultStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-input:
				if !ok {
					return
				}
				select {
				case resultStream <- v:
				case <-done:
				}
			}
		}
	}()
	return resultStream
}
