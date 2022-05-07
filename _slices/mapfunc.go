package _slices

/*func main() {
	d := []int{12,343,232,343,756,54,878,343,2112,343,98,56,5,565,87,676,456,22,3,898,452}
	f := func(in int) int {
		return 10 * in
	}

	d = _slices.MapFunc(f, d)
	fmt.Println(d)
}*/

func MapFunc(f func(int) int, in []int) []int {
	for idx, item := range in {
		in[idx] = f(item)
	}

	return in
}
