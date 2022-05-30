package _struct

import "fmt"

const LIMIT = 4

type Stack struct {
	data []int
	n    int
}

func (t *Stack) push(e int) {
	if cap(t.data) >= t.n+1 {
		t.data[t.n] = e
	} else {
		t.data = append(t.data, e)
	}
	t.n++
}

func (t *Stack) pop() int {
	t.n--
	return t.data[t.n]
}

func (t *Stack) String() string {
	return fmt.Sprintf("%v", t.data)
}
