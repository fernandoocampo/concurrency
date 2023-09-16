package fanins

import (
	"context"
	"sync"
)

func FanIn(ctx context.Context, workers ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	multiplexedStream := make(chan int)

	multiplex := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			select {
			case <-ctx.Done():
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(workers))
	for _, c := range workers {
		go multiplex(c)
	}

	// wait for all the reads to complete
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func WorkerGenerator(ctx context.Context, values []int) []<-chan int {
	result := make([]<-chan int, 0, len(values))
	for _, value := range values {
		result = append(result, newProductWorker(ctx, value))
	}
	return result
}

func newProductWorker(ctx context.Context, value int) <-chan int {
	valueStream := make(chan int)
	go func() {
		defer close(valueStream)
		select {
		case <-ctx.Done():
			return
		case valueStream <- value * 5:
		}
	}()
	return valueStream
}
