package iteration

import (
	"testing"
)

func Test(t *testing.T) {
	rec := Rectangle{10,10}
	got := rec.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	
	checkArea := func(t testing.TB, s Shape, want float64) {
		t.Helper()
		got := s.Area()
		if got != want {
			t.Errorf("%#v got %g want %g", s, got, want)
		}
	}

	tests := []struct {
		s		Shape
		want	float64
	} {
		{Rectangle{12,6}, 72.0},
		{Circle{2}, 12.566370614359172},
		{Triangle{12, 6}, 36.0},
	}

	for _,tc := range tests {
		t.Run("test", func(t *testing.T) {
			checkArea(t, tc.s, tc.want)
		})
	}
}
