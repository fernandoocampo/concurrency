package pooling

import (
	"context"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

func SumInput(ctx context.Context, numbers <-chan uint32) uint32 {
	var result uint32
	maxProcs := runtime.GOMAXPROCS(0)

	var wg sync.WaitGroup
	wg.Add(maxProcs)

	for c := 0; c < maxProcs; c++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-numbers:
					if !ok {
						return
					}
					atomic.AddUint32(&result, v)
				}
			}
		}()
	}

	wg.Wait()

	return result
}

func DoBoundedWork(ctx context.Context, work []string) []string {
	goroutines := runtime.GOMAXPROCS(0)
	wg := sync.WaitGroup{}
	wg.Add(goroutines)
	inputStream := make(chan string, goroutines)
	resultStream := make(chan string, len(work))

	for i := 0; i < goroutines; i++ {
		go func(subctx context.Context) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case value, ok := <-inputStream:
					if !ok {
						return
					}

					resultStream <- strings.ToUpper(value)
				}
			}
		}(ctx)
	}

	for _, v := range work {
		word := v
		inputStream <- word
	}

	close(inputStream)

	wg.Wait()

	close(resultStream)

	result := make([]string, 0, len(work))
	for v := range resultStream {
		result = append(result, v)
	}

	return result
}
