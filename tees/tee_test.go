package tees_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/pipelines"
	"github.com/fernandoocampo/concurrency/tees"
)

func TestTeeChannel(t *testing.T) {
	t.Parallel()
	// Given
	timeout := 2 * time.Second

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()

	want := []int{1, 2, 1, 2}

	// When
	tee1, tee2 := tees.Tee(
		ctx,
		pipelines.NewTakeGenerator(
			ctx, 4,
			pipelines.NewRepeatGenerator(
				ctx, 1, 2,
			),
		),
	)

	// Then
	for i := 0; i < len(want); i++ {
		for j := 0; j < 2; j++ {
			select {
			case <-ctx.Done():
				t.Fatalf("unexpected contex cancellation: %s", ctx.Err())
			case got := <-tee1:
				if got.(int) != want[i] {
					t.Errorf("want: %d, but got: %d in tee1", want[i], got)
				}
			case got := <-tee2:
				if got.(int) != want[i] {
					t.Errorf("want: %d, but got: %d in tee2", want[i], got)
				}
			}
		}
	}
}
