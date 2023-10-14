package healing_test

import (
	"context"
	"log/slog"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/healing"
)

func TestProcessMessage(t *testing.T) {
	// Given
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	values := []string{"a", "b", "c", "d", "e", "f", "g"}
	want := []string{"A", "B", "C", "D", "E", "F", "G"}
	aQueueClient := healing.NewQueueClient(values, logger)
	aRepository := healing.NewRepository(logger)
	aWorker := healing.NewWorker(aQueueClient, aRepository, logger)
	aTimeout := 500 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.TODO(), aTimeout)
	defer cancel()
	// When
	_ = aWorker.ProcessMessages(ctx)
	// Then
	for i := 0; ; i++ {
		got := aRepository.Read(i)
		if got == "" {
			break
		}
		if !slices.Contains(want, got) {
			t.Errorf("on index: %d, want an expected value, but got: %q", i, got)
		}
	}
}
