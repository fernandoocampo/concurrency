package confinement_test

import (
	"slices"
	"testing"

	"github.com/fernandoocampo/concurrency/confinement"
)

func TestGenerateData(t *testing.T) {
	t.Parallel()
	// Given
	want := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	// When
	data := confinement.GenerateData()
	// Then
	got := make([]string, 0)
	for v := range data {
		got = append(got, v)
	}

	if len(got) != len(want) {
		t.Errorf("want: %+v, but got: %+v", want, got)
	}

	if !slices.Equal(want, got) {
		t.Errorf("want: %+v, but got: %+v", want, got)
	}
}
