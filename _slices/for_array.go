package _slices

import "fmt"

func PrintArray() {
	var arr [15]int
	for i := 0; i < 15; i++ {
		arr[i] = i * i
	}

	for i, v := range arr {
		fmt.Printf("arr[%d] = %d\n", i, v)
	}
}
