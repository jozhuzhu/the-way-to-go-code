package _error

import "fmt"

func badCall() {
	m := 22
	n := 0
	fmt.Println(m / n)
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n") // <-- would not reach
}

func TestMain() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}
