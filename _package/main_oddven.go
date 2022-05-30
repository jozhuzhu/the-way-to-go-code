package _package

import (
	"fmt"
	"the-way-to-go-code/_package/even"
)

func JudgeEvenBefore100() {
	evens := make([]int, 0)
	for i := 1; i <= 100; i++ {
		if even.IsEven(i) {
			evens = append(evens, i)
		}
	}

	fmt.Println(evens)
}
