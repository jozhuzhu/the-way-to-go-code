package _slices

func extendSlice(s []int, factor int) []int {
	m := len(s)
	n := cap(s)

	if n >= m*factor {
		s = s[:m*factor]
	} else {
		newSlice := make([]int, m*factor)
		copy(newSlice, s[0:m])
		s = newSlice
	}

	return s
}
