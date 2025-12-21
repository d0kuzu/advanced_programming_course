package shapes

import "math"

func GetTriangle(a, b, c float64) *Triangle {
	return &Triangle{a, b, c}
}

type Triangle struct {
	a, b, c float64
}

func (t *Triangle) GetArea() float64 {
	halfP := t.GetPerimeter() / 2
	return math.Sqrt(halfP * (halfP - t.a) * (halfP - t.b) * (halfP - t.c))
}

func (t *Triangle) GetPerimeter() float64 {
	return t.a + t.b + t.c
}
