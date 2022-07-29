package _goruntine

import (
	"fmt"
	"os"
)

func tel(ch chan<- int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch)
}

func TestTel() {
	ch := make(chan int)
	go tel(ch)

	for v := range ch {
		fmt.Printf("%d ", v)
	}
}

func tel2(ch chan int, quit chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	quit <- true
}

func TestTel2() {
	var ok = true
	ch := make(chan int)
	quit := make(chan bool)

	go tel2(ch, quit)
	for ok {
		select {
		case i := <-ch:
			fmt.Printf("The counter is at %d\n", i)
		case <-quit:
			os.Exit(0)
		}
	}
}
