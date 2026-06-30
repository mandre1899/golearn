package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		checkTest(t, got, want)
	})
	t.Run("testing default fallback when no name provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		checkTest(t, got, want)
	})
}

func checkTest(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
