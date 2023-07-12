package pooling_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/pooling"
)

func TestSumInput(t *testing.T) {
	// Given
	want := uint32(5050)
	timeout := 2 * time.Second
	dataAmount := 100
	inputStream := make(chan uint32)

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()

	go func() {
		defer close(inputStream)
		for i := 1; i <= dataAmount; i++ {
			select {
			case <-ctx.Done():
				return
			case inputStream <- uint32(i):
			}
		}
	}()

	// When
	got := pooling.SumInput(ctx, inputStream)

	// Then
	if got != want {
		t.Errorf("want: %d, but got: %d", want, got)
	}
}
