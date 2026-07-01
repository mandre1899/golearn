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
			t.Errorf("Got %g want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rec := Rectangle{12,6}
		want := 72.0

		checkArea(t, rec, want)
	})
	t.Run("circle", func(t *testing.T) {
		c := Circle{2}
		want := 12.566370614359172

		checkArea(t, c, want)
	})
}
