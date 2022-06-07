package _interface

import "fmt"

type AreaInterface interface {
	Area() float64
}

type PeriInterface interface {
	Perimeter() float32
}

type Triangle struct {
	bottom, high float64
}

func (t *Triangle) Area() float64 {
	return t.bottom * t.high / 2
}

func TestTriangle() {
	triangle := &Triangle{
		3.0,
		4.0,
	}

	fmt.Printf("The area of a triangle '%v' is '%2.2f'\n", triangle, triangle.Area())

	q := &Square{
		side: 5,
	}
	fmt.Printf("The perimeter of a square '%v' is '%2.2f'\n", q, q.Perimeter())
}
