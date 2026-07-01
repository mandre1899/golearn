package iteration

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width	float64
	heigth	float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.heigth + r.width)
}

func (r Rectangle) Area() float64 {
	return r.heigth * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}
