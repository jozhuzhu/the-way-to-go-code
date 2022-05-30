package _struct

import "fmt"

type stack [LIMIT]int

func (s *stack) push(e int) {
	length := len(s)
	s[length] = e
}

func (s *stack) pop() int {
	length := len(s)

	return s[length-1]
}

func (s *stack) String() string {
	return fmt.Sprintf("%#v", s)
}
