package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("Got '%d' want '%d'", got, want)
	}
}

func ExampleAdd() {
	sum := Add(4, 4)
	fmt.Println(sum)
	// Output: 8
}
