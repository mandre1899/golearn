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
	got := SumAllTails([]int{4, 5, 6, 1}, []int{1,2,3}, []int{100,0}, []int{})
	want := []int{12, 5, 0, 0}

	if !slices.Equal(got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}
