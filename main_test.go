package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mateo")
	want := "Hello, Mateo"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
