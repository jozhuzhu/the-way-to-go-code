package _goruntine

import (
	"fmt"
	"time"
)

func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {
	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			x := <-in
			a <- x
			b <- x
			c <- x
		}
	}()
	return a, b, c
}

func fib() <-chan int {
	in := make(chan int, 2)
	a, b, out := dup3(in)
	go func() {
		in <- 0
		in <- 1
		<-a
		for {
			in <- <-a + <-b
		}
	}()
	<-out
	return out
}

func TestGoFibonacci3() {
	start := time.Now()
	f := fib()
	for i := 0; i < 50; i++ {
		fmt.Printf("%d ", <-f)
	}
	fmt.Println()
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", duration)
}
