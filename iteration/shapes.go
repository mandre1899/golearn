package iteration


type Rectangle struct {
	width	float64
	heigth	float64
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.heigth + rec.width)
}

func Area(rec Rectangle) float64 {
	return rec.heigth * rec.width
}
