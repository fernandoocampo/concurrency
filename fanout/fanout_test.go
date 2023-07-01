package fanout_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/fanout"
)

var defaultTimeout = 5 * time.Second

func TestFanoutData(t *testing.T) {
	// Given
	givenNumberOfResults := 100
	expectedNumberOfResults := 100
	expectedValuePerResult := "value"

	ctx, cancel := context.WithTimeout(context.TODO(), defaultTimeout)
	defer cancel()

	// When
	got, err := fanout.Process(ctx, givenNumberOfResults)

	// Then
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if len(got) != expectedNumberOfResults {
		t.Errorf("want: %d number of resuts, but got: %d", expectedNumberOfResults, len(got))
	}

	for _, v := range got {
		value := v
		if v != expectedValuePerResult {
			t.Errorf("want always: %q, but one of the values is: %q", expectedValuePerResult, value)
		}
	}
}
