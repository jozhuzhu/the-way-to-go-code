package _struct

import "fmt"

type Rectangle struct {
	Length int
	Width  int
}

func (r *Rectangle) Area() int {
	return r.Length * r.Width
}

func (r *Rectangle) Perimeter() int {
	return (r.Length + r.Width) * 2
}

func TestRectangle() {
	rec := Rectangle{
		Length: 6,
		Width:  8,
	}

	fmt.Printf("The area of rectangle = %d\n", rec.Area())
	fmt.Printf("The perimeter of rectangle = %d\n", rec.Perimeter())
	fmt.Printf("%v", rec)
}
