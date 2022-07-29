package _goruntine

import (
	"fmt"
	"log"
	"time"
)

type Empty interface{}
type semaphore chan Empty

// acquire n resources
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resources
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

/* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}

func TestLock() {
	var sem = make(semaphore, 1)
	i := 1
	closure1 := func() {
		sem.Lock()
		defer sem.Unlock()

		log.Println("I'm closure1: ", i)
		i++
	}

	closure2 := func() {
		sem.Lock()
		defer sem.Unlock()

		log.Println("I'm closure2: ", i)
		i++
	}

	go closure2()
	go closure1()

	time.Sleep(1e9)

	fmt.Println(i)
}

//import (
//"fmt"
//)
//
//func sum(x, y int, c chan int) {
//	c <- x + y
//}
//
//func main() {
//	c := make(chan int)
//	go sum(12, 13, c)
//	fmt.Println(<-c) // 25
//}
