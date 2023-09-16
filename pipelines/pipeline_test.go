package pipelines_test

import (
	"context"
	"slices"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/pipelines"
)

func TestReverseSlice(t *testing.T) {
	t.Parallel()
	// Given
	oneSecond := 1 * time.Second
	ctx, cancel := context.WithTimeout(context.TODO(), oneSecond)
	defer cancel()
	want := []string{"Zulu", "Yankee", "X-Ray", "Whiskey", "Victor", "Uniform", "Tango", "Sierra", "Romeo", "Quebec", "Papa", "Oscar", "November", "Mike", "Lima", "Kilo", "Juliett", "India", "Hotel", "Golf", "Foxtrot", "Echo", "Delta", "Charlie", "Bravo", "Alfa"}
	words := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-Ray", "Yankee", "Zulu"}
	// When
	got := pipelines.ReverseSlice(ctx, words)
	// Then
	if !slices.Equal(want, got) {
		t.Errorf("want: %+v, but got: %+v", want, got)
	}
}

func TestMultipleOperations(t *testing.T) {
	t.Parallel()
	// Given
	values := []int{1, 2, 3, 4, 5, 6}
	want := []int{4, 16, 36, 64, 100, 144}
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()
	// When
	got := pipelines.IntCollector(
		ctx,
		pipelines.ProductData(
			ctx,
			pipelines.SumData(
				ctx,
				pipelines.IntGenerator(ctx, values),
			),
		),
	)
	// Then
	if !slices.Equal(want, got) {
		t.Errorf("want: %+v, but got: %+v", want, got)
	}
}

func TestRepeatGenerator(t *testing.T) {
	t.Parallel()
	// Given
	values := []any{1, 2, 3}
	timeout := 10 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	// When
	got := pipelines.NewRepeatGenerator(ctx, values...)
	// Then
	i := 0
	for v := range got {
		if v != values[i] {
			t.Fatalf("want: %d, but got: %d", values[i], v)
		}
		if i == 2 {
			i = 0
			continue
		}
		i++
	}
}

func TestTakeGenerator(t *testing.T) {
	t.Parallel()
	// Given
	values := []any{1, 2, 3}
	takeNumber := 12
	want := []any{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3}
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	// When
	got := pipelines.NewTakeGenerator(
		ctx,
		takeNumber,
		pipelines.NewRepeatGenerator(ctx, values...),
	)
	// Then
	for indx, wantValue := range want {
		gotValue := <-got
		if wantValue != gotValue {
			t.Errorf("want: [%d]%v, but got: [%d]%v", indx, wantValue, indx, gotValue)
		}
	}
}
