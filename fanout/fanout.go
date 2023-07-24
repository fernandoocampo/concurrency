package fanout

import (
	"context"
	"fmt"
	"runtime"
)

// Process fan out multiple values from different workers.
func Process(ctx context.Context, numberOfWorkers int) ([]string, error) {
	workers := make([]<-chan string, numberOfWorkers)
	// starting many copies as we have CPUs.
	for i := 0; i < numberOfWorkers; i++ {
		workers[i] = generateData(ctx)
	}

	result := make([]string, 0, numberOfWorkers)
	for indx := range workers {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("reading workers data: %w", ctx.Err())
		case value := <-workers[indx]:
			result = append(result, value)
		}
	}

	return result, nil
}

func ProcessWithSemaphore(ctx context.Context, numberOfWorkers int) ([]string, error) {
	workers := make([]<-chan string, numberOfWorkers)
	semaphore := newSemaphore()

	// starting many copies as number of workers we want.
	for i := 0; i < numberOfWorkers; i++ {
		workers[i] = generateDataWithSemaphore(ctx, semaphore)
	}

	result := make([]string, 0, numberOfWorkers)
	for indx := range workers {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("reading workers data: %w", ctx.Err())
		case value := <-workers[indx]:
			result = append(result, value)
		}
	}

	return result, nil
}

func generateData(ctx context.Context) chan string {
	valueStream := make(chan string)
	go func(ctx context.Context) {
		defer close(valueStream)
		select {
		case <-ctx.Done():
			return
		case valueStream <- "value":
		}
	}(ctx)
	return valueStream
}

func generateDataWithSemaphore(ctx context.Context, semaphore chan bool) chan string {
	valueStream := make(chan string, 1)
	go func(ctx context.Context) {
		defer close(valueStream)
		semaphore <- true
		{
			select {
			case <-ctx.Done():
				return
			case valueStream <- "value":
			}
		}
		<-semaphore
	}(ctx)
	return valueStream
}

func newSemaphore() chan bool {
	workersAvailable := runtime.GOMAXPROCS(0)
	return make(chan bool, workersAvailable)
}
