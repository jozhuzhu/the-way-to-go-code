package _goruntine

import (
	"fmt"
	"os"
	"time"
)

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func goFibonacci(n int) <-chan int {
	out := make(chan int)
	f := func(n int) {
		for i := 0; i < n; i++ {
			out <- fibonacci(i)
		}
		defer close(out)
	}

	go f(n)
	return out
}

func goFibonacci2(n int, out chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		out <- x
		x, y = y, x+y
	}
}

// 通过方法参数将通道进行传递
func fibnterms(term int, c chan int) {
	for i := 0; i <= term; i++ {
		c <- fibonacci(i)
	}
	close(c)
}

func TestGoFibonacci() {
	//start := time.Now()
	//f := goFibonacci(40)
	//for vf := range f {
	//	fmt.Printf("%d ", vf)
	//}
	////for i := 0; i < 40; i++ {
	////	fmt.Printf("%d ", fibonacci(i))
	////}
	//fmt.Println()
	//end := time.Now()
	//delta := end.Sub(start)
	//fmt.Printf("longCalculation took this amount of time: %s\n", delta)

	term := 25
	i := 0
	c := make(chan int)
	start := time.Now()

	go fibnterms(term, c)
	for {
		if result, ok := <-c; ok {
			fmt.Printf("fibonacci(%d) is: %d\n", i, result)
			i++
		} else {
			end := time.Now()
			delta := end.Sub(start)
			fmt.Printf("longCalculation took this amount of time: %s\n", delta)
			os.Exit(0)
		}
	}
}
