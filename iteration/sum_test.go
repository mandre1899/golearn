package iteration

import (
	"testing"
)


func TestSum(t *testing.T) {
	nbrs := [5]int{1,2,3,4,5}

	got := Sum(nbrs)
	want := 15
	
	if got != want {
		t.Errorf("Got %d want %d given, %v", got, want, nbrs)
	}
}
