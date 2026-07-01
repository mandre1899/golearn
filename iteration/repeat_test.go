package iteration

import (
	"fmt"
	"strings"
	"testing"
)

const repeatCount = 50

func TestRepeat(t *testing.T) {
	got := Repeat('a', repeatCount)
	exp := strings.Repeat("a", repeatCount)
	
	if got != exp {
		t.Errorf("Got %q want %q", got, exp)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat('s', 5)
	}
}

func ExampleRepeat() {
	got := Repeat('a', 5)
	fmt.Println(got)
	// Output: aaaaa
}
