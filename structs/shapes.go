package structs

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Height + rect.Width)
}

func Area(rect Rectangle) float64 {
	return rect.Height * rect.Width
}
