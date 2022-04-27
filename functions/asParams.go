package functions

import (
	"strings"
	"unicode/utf8"
)

/*func main() {
	in := "hello, world!, I wanna convert 世界 to ' '."

	out := functions.NonasciiToQuestionMark(in)

	fmt.Println(out)

	// output:
    // hello, world!, I wanna convert ?? to ' '.
}*/
func NonasciiToQuestionMark(in string) (out string) {
	convert := func(r rune) rune {
		// r is not a ASCII rune.
		if r >= utf8.RuneSelf {
			return '?'
		} else {
			return r
		}
	}

	return strings.Map(convert, in)
}
