package _interface

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
	Shape
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 {
	return 4 * sq.side
}

type Rectangle struct {
	length, width float32
	Shape
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radius float64
	Shape
}

func (c Circle) Area() float32 {
	return float32(math.Pi * math.Pow(c.radius, 2))
}

type Shape struct{}

func (s Shape) Area() float32 {
	return 0.0
}

func TestShape() {
	r := Rectangle{5, 3, Shape{}} // Area() of Rectangle needs a value
	q := &Square{5, Shape{}}      // Area() of Square needs a pointer
	c := Circle{5.4, Shape{}}     // Area() of Circle need a value
	shapes := []Shaper{r, q, c}
	fmt.Println("Looping through shapes for area ...")
	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
