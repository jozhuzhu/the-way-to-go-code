package _package

import (
	"fmt"
	"regexp"
	"strconv"
)

func Match() {
	//目标字符串
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "\\d+.\\d+"

	if ok, _ := regexp.MatchString(pat, searchIn); ok {
		fmt.Println("Match string is found")
	}

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)

	str = re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str)
}
