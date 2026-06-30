package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat('a')
	exp := "aaaaa"
	
	if got != exp {
		t.Errorf("Got %q want %q", got, exp)
	}
}
