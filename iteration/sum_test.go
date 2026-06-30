package iteration

import (
	"slices"
	"testing"
)


func TestSum(t *testing.T) {
	nbrs := []int{1,2,3,4,5}

	got := Sum(nbrs)
	want := 15
	
	if got != want {
		t.Errorf("Got %d want %d given, %v", got, want, nbrs)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAllTails([]int{4, 5, 6}, []int{1,2,3})
	want := []int{11, 5}

	if !slices.Equal(got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}
