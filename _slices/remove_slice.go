package _slices

/*func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7 ,8, 9, 10}

	s := _slices.RemoveStringSlice(s1, 0, 3)
	fmt.Println(s)

	s = _slices.RemoveStringSlice(s1, 3, 6)
	fmt.Println(s)

	s = _slices.RemoveStringSlice(s1, 8, 9)
	fmt.Println(s)
}*/

func RemoveStringSlice(base []int, start, end int) []int {
	bLen := len(base)
	//if start == 0 {
	//	base = base[end:bLen]
	//} else if end == bLen - 1 {
	//	base = base[:start]
	//} else {
	newSlice := make([]int, bLen-end+start)
	at := copy(newSlice, base[:start])
	at += copy(newSlice[at:], base[end:])
	base = newSlice
	//}

	return base
}
