package ctxcancel

import "time"

func Do(done <-chan any) <-chan string {
	result := make(chan string)
	go func() {
		defer close(result)

		select {
		case <-done:
			result <- "work not completed"
		case <-time.After(3 * time.Second):
			result <- "work completed"
		}
	}()
	return result
}
