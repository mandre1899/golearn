package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known key", func (t *testing.T){
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown key", func(t *testing.T) {
		got, err := dictionary.Search("na")
		want := ""
		assertError(t, err, errNotFound.Error())
		assertStrings(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	
	dictionary.Add("test2", "wow it worked")
	

	t.Run("added key", func(t *testing.T) {
		got, err := dictionary.Search("test2")
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		want := "wow it worked"
		assertStrings(t, got, want)
	})
	t.Run("add existing key", func(t *testing.T) {
		err := dictionary.Add("test2", "wow")
		assertError(t, err, errKeyExists.Error())
	})
	
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func assertError(t testing.TB, err error, want string) {
	if err == nil {
		t.Fatal("expected error but got nil")
	}

	if err.Error() !=  want {
		t.Errorf("Got %q want %q", err.Error(), want)
	}
}
