package _goruntine

import "fmt"

func TestCommunicationInGoruntine() {
	signal := make(chan interface{}, 2)
	producer := func(c chan int) {
		for i := 0; i < 10; i++ {
			c <- i * 10
		}
		// It's not necessary for producer
		//signal <- new(interface{})
		// the close will execute until the chan is empty
		close(c)
	}

	comsumer := func(c chan int) {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		signal <- new(interface{})
	}

	c := make(chan int, 1)

	go producer(c)
	go comsumer(c)

	<-signal
}

//import "fmt"
//
//// integer producer:
//func numGen(start, count int, out chan<- int) {
//	for i := 0; i < count; i++ {
//		out <- start
//		start = start + count
//	}
//	close(out)
//}
//
//// integer consumer:
//func numEchoRange(in <-chan int, done chan<- bool) {
//	for num := range in {
//		fmt.Printf("%d\n", num)
//	}
//	done <- true
//}
//
//func main() {
//	numChan := make(chan int)
//	done := make(chan bool)
//	go numGen(0, 10, numChan)
//	go numEchoRange(numChan, done)
//
//	<-done
//}
