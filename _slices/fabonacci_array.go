package _slices

import (
	"fmt"
)

/*func main() {
	var start, end time.Time

	start = time.Now()
	_slices.FabonacciUsingClosure(40)
	end = time.Now()
	fmt.Printf("longCalculation took this amount of time: %s\n", end.Sub(start))

	start = time.Now()
	_slices.FabonacciUsingArray(40)
	end = time.Now()
	fmt.Printf("longCalculation took this amount of time: %s\n", end.Sub(start))
}*/

var res [50]int

func FabonacciUsingClosure(in int) {
	var f1, f2 = 0, 1
	iterator := func(n int) int {
		if n <= 1 {
			return 1
		}

		f1, f2 = f2, f1+f2
		return f2
	}

	for i := 0; i < in; i++ {
		fmt.Printf("[%d] ==> %d\n", i, iterator(i))
	}
}

func FabonacciUsingArray(in int) {
	for i := 0; i < in; i++ {
		fmt.Printf("[%d] ==> %d\n", i, fabonacci(i))
	}
}

func FabonacciReturnArray(in int) (out []int) {
	if res[in] != 0 {
		return res[:in]
	}

	fabonacci(in)
	return res[:in]
}

// 该实现版本所得到的fabonacci是从 0 开始的
func fabonacci(in int) (out int) {
	if res[in] != 0 {
		out = res[in]
		return
	}

	if in <= 1 {
		out = 1
	} else {
		out = fabonacci(in-1) + fabonacci(in-2)
	}
	res[in] = out

	return
}
