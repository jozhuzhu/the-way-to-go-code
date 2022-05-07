package _slices

/*func main() {
	str := "Hello World!"
	former, latter := _slices.Split(str, 5)
	fmt.Println(former)
	fmt.Println(latter)
}*/

func Split(str string, idx int) (former, latter string) {
	b := []byte(str)
	former = string(b[:idx])
	latter = string(b[idx:])

	return
}
