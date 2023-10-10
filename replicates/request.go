package replicates

import (
	"math/rand"
	"sync"
	"time"
)

func Request(done <-chan any, id int, wg *sync.WaitGroup, result chan<- int) {
	started := time.Now()
	defer wg.Done()

	simulatedLoadTime := time.Duration(1+rand.Intn(10)) * time.Millisecond
	select {
	case <-done:
	case <-time.After(simulatedLoadTime):
	}

	select {
	case <-done:
	case result <- id:
	}

	took := time.Since(started)

	if took < simulatedLoadTime {
		took = simulatedLoadTime
	}
}
