package _io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines int

func CountWordFromStdio() {
	// 答案的思路是 结束符单独一行，而我考虑的是 结束符可以是在某一个正常输入的末尾
	nrchars, nrwords, nrlines = 0, 0, 0
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input, type S to stop: ")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}
		if input == "S\r\n" { // Windows, on Linux it is "S\n"
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}
		Counters(input)
	}

	//var words []string
	//fmt.Println("please enter any words you want and end with 'S':")
	//inputReader := bufio.NewReader(os.Stdin)
	//input, err := inputReader.ReadString('S')
	//if err != nil {
	//	panic(err)
	//}
	//
	//var characterCount, wordCount, lineCount int
	//var lastIndex int
	//lineCount = 1
	//for ix, v := range input {
	//	switch v {
	//	case '\r', '\n':
	//		lineCount++
	//	case ' ', ',', '.':
	//		if lastIndex <= 0 {
	//			words = append(words, input[lastIndex:ix])
	//		} else {
	//			words = append(words, input[lastIndex + 1:ix])
	//		}
	//		lastIndex = ix
	//		wordCount++
	//	default:
	//		characterCount++
	//	}
	//}
	//
	//fmt.Printf("%d words, %d characters, %d lines\n", wordCount, characterCount, lineCount)
	//fmt.Println(words)
}

func Counters(input string) {
	nrchars += len(input) - 2 // -2 for \r\n
	// 通过默认空格进行分割，这种方式或许会省去其他标点符号的情况
	// 比较好的方式是使用 FieldsFun
	nrwords += len(strings.Fields(input))
	nrlines++
}
