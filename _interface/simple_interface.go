package _interface

import "fmt"

type Simpler interface {
	Get() int
	Set(d int)
}

type Simple struct {
	data int
}

func (s *Simple) Get() int {
	return s.data
}

func (s *Simple) Set(d int) {
	s.data = d
}

func TestSimple() {
	var s Simpler

	s = &Simple{data: 10000}

	fmt.Println(s.Get())

	s.Set(50000)

	fmt.Println(s.Get())
}
