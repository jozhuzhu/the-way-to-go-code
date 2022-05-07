package _slices

/*func main() {
	d := []int{12,343,232,343,756,54,878,343,2112,343,98,56,5,565,87,676,456,22,3,898,452}
	d = _slices.BubbleSort(d)
	fmt.Println(d)
}*/

func BubbleSort(input []int) []int {
	for h := len(input) - 1; h > 0; h-- {
		for i := 0; i < h; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
			}
		}
	}

	return input
}
