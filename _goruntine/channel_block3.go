package _goruntine

import (
	"fmt"
	"time"
)

func TestBlock() {
	ch1 := make(chan int, 1)
	go pump(ch1) // pump hangs
	go suck(ch1)
	time.Sleep(10e9)
	fmt.Println("main entrance ends.")
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("after 5s:")
			ch <- 10
		default:
			fmt.Println(<-ch)
		}
	}
}
