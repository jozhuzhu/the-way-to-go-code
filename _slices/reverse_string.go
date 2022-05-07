package _slices

import "fmt"

/*func main() {
	str := "Hello World!"
	r := _slices.Reverse(str)
	fmt.Println(r)
}*/

func Reverse(str string) (r string) {
	b := []byte(str)
	var tmp byte
	for i, j := 0, len(b)-1; i <= j; i, j = i+1, j-1 {
		tmp = b[i]
		b[i] = b[j]
		b[j] = tmp
		// b[i], b[j] = b[j], b[i]
	}

	fmt.Println(len(b))

	return string(b)
}

/*func main() {
	str := "Hello 世界!"
	r := _slices.Reverse(str)
	fmt.Println(r)
}*/

func ReverseUnicode(str string) (r string) {
	b := []rune(str)
	var tmp rune
	for i, j := 0, len(b)-1; i <= j; i, j = i+1, j-1 {
		tmp = b[i]
		b[i] = b[j]
		b[j] = tmp
		// b[i], b[j] = b[j], b[i]
	}

	fmt.Println(len(b))
	return string(b)
}
