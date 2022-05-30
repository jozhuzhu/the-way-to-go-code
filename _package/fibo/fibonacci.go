package fibo

// 题目理解错误，以为操作符是针对fibonacci数列的累加或累乘
//func Fibonacci(num int) (out int){
//	var result []int = make([]int, 100)
//	if result[num] != 0 {
//		return result[num]
//	}
//	if num <= 1 {
//		out = 1
//	} else {
//		out = Fibonacci(num - 1) + Fibonacci(num - 2)
//	}
//	result[num] = out
//	return
//}
//
//func CalcFibo(num int, operator string) (out int) {
//	var outs []int = make([]int, num)
//	for i := 0; i < num; i++ {
//		outs[i] = Fibonacci(i)
//	}
//
//	fmt.Println(outs)
//	switch operator {
//	case "+" :
//		for i := 0; i < num; i++ {
//			out += outs[i]
//		}
//	case "*" :
//		out = 1
//		for i := 0; i < num; i++ {
//			out *= outs[i]
//		}
//	}
//
//	return
//}

// fibonacci 默认返回的是相应位置的值
func Fibonacci(op string, n int) (res int) {
	if n <= 1 {
		switch op {
		case "+":
			res = 1
		case "*":
			res = 2
		default:
			res = 0
		}
	} else {
		switch op {
		case "+":
			res = Fibonacci(op, n-1) + Fibonacci(op, n-2)
		case "*":
			res = Fibonacci(op, n-1) * Fibonacci(op, n-2)
		default:
			res = 0
		}
	}
	return
}
