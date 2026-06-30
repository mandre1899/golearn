package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat('a')
	exp := strings.Repeat("a", 5)
	
	if got != exp {
		t.Errorf("Got %q want %q", got, exp)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat('s')
	}
}

func ExampleRepeat() {
	got := Repeat('a')
	fmt.Println(got)
	// Output: aaaaa
}
