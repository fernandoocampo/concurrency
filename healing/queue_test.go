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

func TestQueueClient(t *testing.T) {
	// Given
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	values := []string{"a", "b", "c", "d", "e", "f", "g"}
	want := []string{"a", "b", "c", "d", "e", "f", "g"}
	aQueueClient := healing.NewQueueClient(values, logger)
	aTimeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.TODO(), aTimeout)
	defer cancel()
	// When
	got := aQueueClient.RetrieveMessages(ctx)
	// Then
	if !slices.Equal(want, got) {
		t.Errorf("want: %+v, but got: %+v", want, got)
	}
}
