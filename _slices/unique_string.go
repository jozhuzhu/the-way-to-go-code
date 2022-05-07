package _slices

func Uniq(str string) (u string) {
	var result []rune
	b := []rune(str)

	for i := 0; i < len(b)-1; i++ {
		if b[i] != b[i+1] {
			result = append(result, b[i])
		}
	}

	return string(result)
}
