package ors_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/ors"
)

func TestOrChannel(t *testing.T) {
	t.Parallel()
	// Given
	aFunc := func(name string, after time.Duration) <-chan any {
		aStream := make(chan any)
		go func() {
			defer close(aStream)
			// some work
			time.Sleep(after)
			fmt.Println("g:", name)
		}()
		return aStream
	}
	start := time.Now()
	// When
	<-ors.Or(
		aFunc("a", 2*time.Hour),
		aFunc("b", 5*time.Minute),
		aFunc("c", 1*time.Second),
		aFunc("d", 1*time.Hour),
		aFunc("e", 1*time.Minute),
		aFunc("f", 3*time.Hour),
		aFunc("h", 5*time.Minute),
		aFunc("i", 10*time.Millisecond),
	)
	got := time.Since(start)
	// Then
	if got <= 1 {
		t.Errorf("wanted an elapsed time greater than 1, but got: %d", got)
	}
}
