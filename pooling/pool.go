package pooling

import (
	"context"
	"runtime"
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
