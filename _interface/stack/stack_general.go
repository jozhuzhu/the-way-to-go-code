package stack

import "errors"

type Stack interface {
	Len() int
	IsEmpty() bool
	Push(x interface{})
	Pop() (interface{}, error)
}

func NewStack(length int) Stack {
	data := make([]interface{}, length)
	gs := &GeneralStack{
		data,
		0,
	}

	return gs
}

type GeneralStack struct {
	data []interface{}
	n    int
}

func (g *GeneralStack) Len() int {
	return g.n
}

func (g *GeneralStack) IsEmpty() bool {
	return g.n <= 0
}

func (g *GeneralStack) Push(x interface{}) {
	g.data[g.n] = x
	g.n++
}

func (g *GeneralStack) Pop() (interface{}, error) {
	if g.IsEmpty() {
		return nil, errors.New("the stack is empty")
	}
	g.n--
	result := g.data[g.n]
	return result, nil
}
