package replicates_test

import (
	"sync"
	"testing"

	"github.com/fernandoocampo/concurrency/replicates"
)

func TestReplicatedRequests(t *testing.T) {
	t.Parallel()
	// Given
	done := make(chan any)
	result := make(chan int)
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		// When
		go replicates.Request(done, i, &wg, result)
	}

	firstReturned := <-result
	close(done)
	wg.Wait()
	// Then
	if firstReturned < 0 || firstReturned > 9 {
		t.Errorf("want: %d, but got: %d", 11, firstReturned)
	}
}
