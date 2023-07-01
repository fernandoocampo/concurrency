package wait_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/wait"
)

func TestWaitForTask(t *testing.T) {
	// Given
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	givenValue := 2
	givenValueStream := createValueStream(ctx, givenValue)

	want := 8

	// When
	got, err := wait.PowerOfThree(ctx, givenValueStream)

	// Then
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if got != want {
		t.Errorf("want: %d, but got: %d", want, got)
	}
}

func createValueStream(ctx context.Context, value int) <-chan int {
	givenValueStream := make(chan int)

	go func() {
		defer close(givenValueStream)
		select {
		case <-ctx.Done():
			return
		case givenValueStream <- value:
		}
	}()

	return givenValueStream
}
