package _interface

import (
	"fmt"
	stack2 "the-way-to-go-code/_interface/stack"
)

func TestStack() {
	stack := stack2.NewStack(20)
	fmt.Println("The length of STACK is ", stack.Len())
	fmt.Println("the stack is empty? ", stack.IsEmpty())

	stack.Push(1)
	stack.Push("love")
	stack.Push("human")
	stack.Push("reminisce")
	fmt.Println("The length of STACK is ", stack.Len())

	x, _ := stack.Pop()
	fmt.Println("the top element of stack is ", x)
	fmt.Println("The length of STACK is ", stack.Len())

	fmt.Println("the stack is empty? ", stack.IsEmpty())
}
