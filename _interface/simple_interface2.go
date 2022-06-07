package _interface

import "fmt"

type RSimple struct {
	radius int
}

func (r *RSimple) Get() int {
	return r.radius
}

func (r *RSimple) Set(d int) {
	r.radius = d
}

func fi(s Simpler) {
	switch s.(type) {
	case *Simple:
		fmt.Println("This is Simple instance")
	case *RSimple:
		fmt.Println("This is RSimple instance")
	}
}

func TestType() {
	var s Simpler
	s = &Simple{1}
	gI(s)

	s = &RSimple{2}
	gI(s)
}
