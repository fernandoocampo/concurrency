package bridges_test

import (
	"slices"
	"testing"

	"github.com/fernandoocampo/concurrency/bridges"
)

func TestBridge(t *testing.T) {
	t.Parallel()
	// Given
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	want := []any{1, 2, 3, 4, 5, 6, 7, 8, 9}
	done := make(chan any)
	defer close(done)
	// Then
	bridgeResult := bridges.Bridge(done, bridges.GenValues(done, values))
	// When
	got := make([]any, 0)
	for value := range bridgeResult {
		got = append(got, value)
	}

	if !slices.Equal(want, got) {
		t.Errorf("want: %+v, but got: %+v", want, got)
	}
}
