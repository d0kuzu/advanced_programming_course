package shapes

func GetSquare(a float64) *Square {
	return &Square{a}
}

type Square struct {
	a float64
}

func (s *Square) GetArea() float64 {
	return s.a * s.a
}

func (s *Square) GetPerimeter() float64 {
	return 4 * s.a
}
