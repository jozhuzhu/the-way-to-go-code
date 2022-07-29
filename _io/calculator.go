package _io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	. "the-way-to-go-code/_io/stack"
)

// err 的情况特殊处理
// 其他情况包括数字的判断、结束符的判断都放在 switch-case 里面
func CLICalculator() {
	var token string
	operStack := new(OperatorStack)

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter numbers or operators to do calculation and end it with 'q'")

	token, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("Input error !")
		os.Exit(1)
	}
	token = token[0 : len(token)-2] // remove "\r\n"
	// fmt.Printf("--%s--\n",token)  // debug statement
	switch {
	case token == "q": // stop als invoer = "q"
		fmt.Println("Calculator stopped")
		return
	case token >= "0" && token <= "999999":
		i, _ := strconv.Atoi(token)
		operStack.Push(i)
	case token == "+":
		q, _ := operStack.Pop()
		p, _ := operStack.Pop()
		fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)+q.(int))
	case token == "-":
		q, _ := operStack.Pop()
		p, _ := operStack.Pop()
		fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)-q.(int))

	case token == "*":
		q, _ := operStack.Pop()
		p, _ := operStack.Pop()
		fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)*q.(int))

	case token == "/":
		q, _ := operStack.Pop()
		p, _ := operStack.Pop()
		fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)/q.(int))
	default:
		fmt.Println("No valid input")
	}

	//input, err := inputReader.ReadBytes('\n')
	//
	//for true {
	//	var n1, n2, result Obj
	//	if err != nil || input[0] == 'q' {
	//		break
	//	}
	//
	//	if input[0] > '0' && input[0] < '9' {
	//		intData, _ := strconv.Atoi(string(input[0]))
	//		operStack.Push(intData)
	//		input, err = inputReader.ReadBytes('\n')
	//		continue
	//	}
	//
	//	switch input[0] {
	//	case '+':
	//		n1, _ = operStack.Pop()
	//		n2, _ = operStack.Pop()
	//		result = n1.(int) + n2.(int)
	//	case '-':
	//		n1, _ = operStack.Pop()
	//		n2, _ = operStack.Pop()
	//		result = n1.(int) - n2.(int)
	//	case '*':
	//		n1, _ = operStack.Pop()
	//		n2, _ = operStack.Pop()
	//		result = n1.(int) * n2.(int)
	//	case '/':
	//		n1, _ = operStack.Pop()
	//		n2, _ = operStack.Pop()
	//		result = n1.(int) / n2.(int)
	//	}
	//
	//	fmt.Println(result)
	//	input, err = inputReader.ReadBytes('\n')
	//}

}
