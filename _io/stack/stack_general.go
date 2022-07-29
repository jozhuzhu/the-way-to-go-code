package stack

import (
	"errors"
)

type Obj interface{}

type Stack interface {
	Len() int
	IsEmpty() bool
	Push(x Obj)
	Pop() (Obj, error)
}

type OperatorStack []Obj

func (o OperatorStack) Len() int {
	return len(o)
}

func (o OperatorStack) IsEmpty() bool {
	return o.Len() <= 0
}

func (o *OperatorStack) Push(x Obj) {
	*o = append(*o, x)
}

func (o *OperatorStack) Pop() (Obj, error) {
	if o.IsEmpty() {
		return nil, errors.New("operation stack is empty")
	}
	stack := *o
	result := stack[o.Len()-1]
	*o = stack[:o.Len()-1]
	return result, nil
}
