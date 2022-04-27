package functions

import "fmt"

func NewFabonacci_1(n int) (index, value int) {
	index = n
	if n <= 1 {
		value = 1
	}
	_, lastOne := NewFabonacci_1(n - 1)
	_, lastTwo := NewFabonacci_1(n - 2)
	value = lastOne + lastTwo
	return
}

func NewFabonacci_2(n int) (value int) {
	func(n int) {
		f1, f2 := 1, 1
		if n <= 2 {
			value = f2
		} else {
			for i := 3; i <= n; i++ {
				value = f1 + f2
				f1 = f2
				f2 = value
			}
		}
	}(n)

	return
}

/*func main() {
	f := functions.NewFabonacci_3() //返回一个闭包函数
	var array [40]int
	for i := 0; i < 40; i++ {
		array[i] = f()
	}
	fmt.Println(array)
}*/

func NewFabonacci_3() func() int {
	back1, back2 := 0, 1
	return func() int {
		// 重新赋值
		back1, back2 = back2, back1+back2
		return back1
	}
}

func From10To1() {
	printNum(10)
}

func printNum(n int) {
	if n <= 0 {
		return
	}
	fmt.Println(n)
	printNum(n - 1)
}
