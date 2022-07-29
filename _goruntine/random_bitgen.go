package _goruntine

import (
	"fmt"
	"math/rand"
	"time"
)

func randombitgen() <-chan string {
	out := make(chan string)
	go func() {
		for {
			rand.Seed(time.Now().Unix())
			r := rand.Int()
			out <- intToBinary(r)
		}
	}()

	return out
}

func intToBinary(n int) string {
	var b []byte
	for n != 0 {
		k := n % 2
		b = append(b, byte(k+48))
		n = n / 2
	}

	//fmt.Println(string(b))
	return string(b)
}

func TestRandomBitGen() {
	//c := randombitgen()
	//
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-c)
	//}

	c := make(chan int)
	// consumer:
	go func() {
		for {
			fmt.Print(<-c, " ")
		}
	}()
	// producer:
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}
