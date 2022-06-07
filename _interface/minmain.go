package _interface

import "fmt"

func ints() {
	data := IntArray{1, 34, 342, 34, 32, 7, 54, 576, 232, 45, 56, 565, -34, -2134, -9439}
	min := Min(data)
	fmt.Println(data)
	fmt.Printf("The minimum of the array is %d\n", min)
}

func strings() {
	data := StringArray{"jkekjf", "lksdfk", "jkdjfka", "dafdaj", "bncadfa", "adkfjakf", "3949348"}
	min := Min(data)
	fmt.Println(data)
	fmt.Printf("The minimum of the array is %s\n", min)
}

func TestMin() {
	ints()
	strings()
}
