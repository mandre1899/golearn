package main

import "testing"

func TestAdd(t *testing.T) {
	sum := add(2, 2)
	exp := 4

	if sum != exp {
		t.Errorf("failed: expected %d got %d", exp, sum)
	}

}
