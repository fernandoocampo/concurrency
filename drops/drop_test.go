package drops_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/drops"
)

func TestDropSomething(t *testing.T) {
	t.Parallel()
	// Given
	givenWork := 1000
	givenCapacity := 100
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()

	// When
	processedStream := drops.DropSomething(ctx, givenWork, givenCapacity)

	var got []string
	for v := range processedStream {
		got = append(got, v)
	}

	// Then
	if len(got) == 0 {
		t.Errorf("result cannot be empty")
	}
	if len(got) > givenWork {
		t.Errorf("wanted work should be less than %d, but got: %d", givenWork, len(got))
	}
}
