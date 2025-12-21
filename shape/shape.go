package shape

import (
	. "advance/shape/shapes"
	"fmt"
	"reflect"
)

type Shape interface {
	GetArea() float64
	GetPerimeter() float64
}

func Start() {
	var shapes []Shape

	shapes = append(shapes, GetRectangle(5, 6))
	shapes = append(shapes, GetCircle(7))
	shapes = append(shapes, GetSquare(8))
	shapes = append(shapes, GetTriangle(5, 6, 7))

	for _, shape := range shapes {
		fmt.Println("Shape is", reflect.TypeOf(shape).Elem().Name()) //Elem() - cause we using pointer
		fmt.Println("Area is", shape.GetArea())
		fmt.Println("Perimeter is", shape.GetPerimeter())
		fmt.Println("")
	}
}
