package main

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

func Area(rec Rectangle) float64 {
	return rec.Width * rec.Height
}
