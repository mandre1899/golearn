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
	t.Run("rectangle", func(t *testing.T) {
		rec := Rectangle{12,6}
		got := rec.Area()
		want := 72.0

		if got != want {
			t.Errorf("Got %.2f want %.2f", got, want)
		}
	})
	t.Run("circle", func(t *testing.T) {
		c := Circle{2}
		got := c.Area()
		want := 12.566370614359172

		if got != want {
			t.Errorf("Got %g want %g", got, want)
		}
	})
}
