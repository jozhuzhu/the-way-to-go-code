package _slices

import "fmt"

func ArrayAssign() {
	var arr1 = new([5]int)
	var arr2 [5]int
	arr2 = *arr1

	fmt.Printf("Address of arr1: %p\n", arr1)
	fmt.Printf("Address of arr2: %p\n", &arr2)
}
