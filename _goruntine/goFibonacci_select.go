package _goruntine

import "fmt"

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			// break 这样会导致继续报错：all goroutines are asleep-deadlock!
			return
		}
	}
}

func TestGoFibonacci2() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 25; i++ {
			fmt.Printf("%d ", <-c)
		}
		quit <- 0
	}()

	fibonacci2(c, quit)
}
