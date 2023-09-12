package pipelines_test

import (
	"context"
	"slices"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/pipelines"
)

func TestReverseSlice(t *testing.T) {
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
