package queuing_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/pipelines"
	"github.com/fernandoocampo/concurrency/queuing"
)

func TestQueuing(t *testing.T) {
	// Given
	timeout := 20 * time.Second
	shortTime := 1 * time.Second
	longTime := 4 * time.Second
	bufferItems := 2
	want := make([]int, 3)
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()

	// When
	zeros := pipelines.NewTakeGenerator(ctx, 3, pipelines.NewRepeatGenerator(ctx, 0))
	short := queuing.Sleep(ctx, shortTime, zeros)
	buffer := queuing.Buffer(ctx, bufferItems, short)
	long := queuing.Sleep(ctx, longTime, buffer)

	// Then
	for _, v := range want {
		got := <-long
		if v != got {
			t.Errorf("want: %d, but %d", v, got)
		}
	}
}
