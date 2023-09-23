package ctxcancel_test

import (
	"testing"

	"github.com/fernandoocampo/concurrency/ctxcancel"
)

func TestDo(t *testing.T) {
	t.Parallel()
	// Given
	done := make(chan any)
	want := "work not completed"
	resultStream := ctxcancel.Do(done)
	close(done)
	got := <-resultStream
	// Then
	if got != want {
		t.Errorf("want: %s, but got: %s", want, got)
	}
}
