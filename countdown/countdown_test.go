package countdown

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}


	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := "3\n2\n1\nGo!\n"
	
	if spySleeper.Calls != 4 {
		t.Errorf("Wanted 4 sleep calls got %d", spySleeper.Calls)
	}
	
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
