package ors_test

import (
	"testing"

	"github.com/fernandoocampo/concurrency/ors"
)

func TestOrDone(t *testing.T) {
	// Given
	done := make(chan struct{})
	defer close(done)
	inputStream := make(chan any)
	go func() {
		defer close(inputStream)
		select {
		case <-done:
			return
		case inputStream <- 1:
		}
	}()
	want := 1
	// When
	got := ors.OrDone(done, inputStream)
	// Then
	if want != <-got {
		t.Errorf("want: %d, but got: %d", want, got)
	}
}
