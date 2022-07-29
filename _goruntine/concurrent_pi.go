package _goruntine

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func compution(k int) <-chan float64 {
	out := make(chan float64)
	go func() {
		for i := 0; i < k; i++ {
			out <- 4 * math.Pow(-1, float64(i)) / (2*float64(i) + 1)
		}
		close(out)
	}()
	return out
}

func ComputePI() {
	runtime.GOMAXPROCS(1)
	start := time.Now()
	e := compution(10000)
	var sum float64 = 0
	for v := range e {
		sum += v
	}

	fmt.Println("pi=", sum)
	fmt.Println("real pi=", math.Pi)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("time cost in computing pi: %s", delta)
}

const NCPU = 2

func ComputePI2() {
	start := time.Now()
	runtime.GOMAXPROCS(2)
	fmt.Println(CalculatePi(12000))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func CalculatePi(end int) float64 {
	ch := make(chan float64)
	for i := 0; i < NCPU; i++ {
		go term(ch, i*end/NCPU, (i+1)*end/NCPU)
	}
	result := 0.0
	for i := 0; i < NCPU; i++ {
		result += <-ch
	}
	return result
}

func term(ch chan float64, start, end int) {
	result := 0.0
	for i := start; i < end; i++ {
		x := float64(i)
		result += 4 * (math.Pow(-1, x) / (2.0*x + 1.0))
	}
	ch <- result
}
