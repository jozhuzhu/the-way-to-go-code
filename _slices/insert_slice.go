package _slices

/*func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7 ,8, 9, 10}

	s := _slices.InsertStringSlice(s1, 2, s2)
	fmt.Println(s)
}*/

func InsertStringSlice(base []int, idx int, elems []int) []int {
	bLen := len(base)
	eLen := len(elems)
	var tmp []int = make([]int, bLen-idx)

	if cap(base) >= bLen+eLen {
		base = base[:bLen+eLen]
	} else {
		newSlice := make([]int, bLen+eLen)
		copy(newSlice, base[:bLen])
		base = newSlice
	}

	copy(tmp, base[idx:bLen])
	copy(base[idx:], elems)
	copy(base[idx+eLen:], tmp)
	return base
}

func InsertStringSlice_1(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}
