package shapes

import "math"

func GetCircle(r float64) *Circle {
	return &Circle{r}
}

type Circle struct {
	radius float64
}

func (c *Circle) GetArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) GetPerimeter() float64 {
	return 2 * math.Pi * c.radius
}
