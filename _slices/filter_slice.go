package _slices

/*func main() {
	isEven := func(i int) bool {
		return i % 2 == 0
	}

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s = _slices.Filter(s, isEven)
	for _, item := range s {
		fmt.Println(item)
	}
}*/

func Filter(s []int, fn func(int) bool) []int {
	newSlice := make([]int, 0, len(s))
	// var newSlice []int // == nil
	for _, item := range s {
		if fn(item) {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice
}
