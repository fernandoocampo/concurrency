package heartbeats_test

import (
	"slices"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/heartbeats"
)

func TestBasicHearbeat(t *testing.T) {
	t.Parallel()
	// Given
	input := newRadioAlphabet()
	timeout := 2 * time.Second
	done := make(chan any)
	time.AfterFunc(timeout, func() { close(done) })
	timeInterval := 20 * time.Millisecond
	// When
	got, pulses := heartbeats.DoBasicProcess(done, input, timeInterval)
	// Then
	for _, pulse := range pulses {
		if pulse != heartbeats.APulse {
			t.Fatalf("want: %q, but got: %q", heartbeats.APulse, pulse)
		}
	}
	if len(got) != len(input) {
		t.Fatalf("want this number of items: %d, but got: %d", len(input), len(got))
	}
	for _, word := range input {
		if !slices.Contains(got, word) {
			t.Errorf("want this word: %q, but doesn't exist", word)
		}
	}

}

func newRadioAlphabet() []string {
	return []string{
		"Alfa",
		"Bravo",
		"Charlie",
		"Delta",
		"Echo",
		"Foxtrot",
		"Golf",
		"Hotel",
		"India",
		"Juliett",
		"Kilo",
		"Lima",
		"Mike",
		"November",
		"Oscar",
		"Papa",
		"Quebec",
		"Romeo",
		"Sierra",
		"Tango",
		"Uniform",
		"Victor",
		"Whiskey",
		"X-Ray",
		"Yankee",
		"Zulu",
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Zero",
		"Hundred",
		"Thousand",
		"Decimal",
		"Stop",
	}
}
