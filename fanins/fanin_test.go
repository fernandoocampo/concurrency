package fanins_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/fanins"
)

func TestFanIn(t *testing.T) {
	t.Parallel()
	// Given
	values := []int{1, 2, 3, 4, 5}
	want := map[int]int{5: 1, 10: 2, 15: 3, 20: 4, 25: 5}
	timeout := 1 * time.Second
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	// When
	got := fanins.FanIn(ctx, fanins.WorkerGenerator(ctx, values)...)
	// Then
	var count int
	for gotValue := range got {
		count++
		_, ok := want[gotValue]
		if !ok {
			t.Errorf("unknown result: %d", gotValue)
			continue
		}
	}

	if count != len(want) {
		t.Errorf("want %d elements, but got %d", len(want), count)
	}
}
