package shapes

func GetRectangle(a, b float64) *Rectangle {
	return &Rectangle{a, b}
}

type Rectangle struct {
	a, b float64
}

func (r *Rectangle) GetArea() float64 {
	return r.a * r.b
}

func (r *Rectangle) GetPerimeter() float64 {
	return 2 * (r.a + r.b)
}
