package healing_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/fernandoocampo/concurrency/healing"
)

func TestSaveRepository(t *testing.T) {
	// Given
	aLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	aNewRepository := healing.NewRepository(aLogger)
	aRecord := "a"
	// When
	err := aNewRepository.Save(aRecord)
	// Then
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestReadRepository(t *testing.T) {
	// Given
	aLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	aRepository := healing.NewRepository(aLogger)
	aRecord := "a"
	err := aRepository.Save(aRecord)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	aRecordID := 0
	want := "a"
	// When
	got := aRepository.Read(aRecordID)
	// Then
	if want != got {
		t.Errorf("want: %q, but got: %q", want, got)
	}
}
