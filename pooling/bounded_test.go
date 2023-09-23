package pooling_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/pooling"
)

func TestBoundedWorkPooling(t *testing.T) {
	t.Parallel()
	// Given
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	someWork := []string{"work", "work", "work", 2000: "work"}
	wantedWorks := 4
	wantedWorkWord := "WORK"

	// When
	got := pooling.DoBoundedWork(ctx, someWork)

	// Then
	if len(got) != len(someWork) {
		t.Errorf("want: %d amount of work, but got: %d", len(someWork), len(got))
	}

	var gotWorkWithWork int
	for _, v := range got {
		if v == wantedWorkWord {
			gotWorkWithWork++
		}
	}

	if wantedWorks != gotWorkWithWork {
		t.Errorf("want: %d works with word %q, but got: %d", wantedWorks, wantedWorkWord, gotWorkWithWork)
	}
}
