package wait_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/wait"
)

func TestWaitForSum(t *testing.T) {
	// Given
	a := 1
	b := 2
	var result int
	expectedResult := 3
	timeout1Sec := 1 * time.Second
	ctx, cancel := context.WithTimeout(context.TODO(), timeout1Sec)
	defer cancel()

	// When
	resultStream := wait.Sum(a, b)
	select {
	case <-ctx.Done():
		t.Fatalf("it took longer than expected: %s", ctx.Err())
	case sum := <-resultStream:
		result = sum
	}

	// Then
	if result != expectedResult {
		t.Errorf("want: %d, but got: %d", expectedResult, result)
	}
}
