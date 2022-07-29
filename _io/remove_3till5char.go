package _io

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Remove_3till5Char() {
	inputFile, _ := os.Open("goprogram")
	outputFile, _ := os.OpenFile("goprogramT", os.O_WRONLY|os.O_CREATE, 0666)
	defer outputFile.Close()
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)

	var outputString string
	for {
		inputString, _, readError := inputReader.ReadLine()
		if readError == io.EOF {
			fmt.Println("EOF")
			return
		}

		//outputString := string(inputString[3:6]) + "\r\n"
		// 需要考虑好文件里面可能不存在对应长度的字节
		//fmt.Printf("The input was: --%s--", inputString)
		if len(inputString) < 3 {
			outputString = "\r\n"
		} else if len(inputString) < 5 {
			outputString = string([]byte(inputString)[2:len(inputString)]) + "\r\n"
		} else {
			outputString = string([]byte(inputString)[2:5]) + "\r\n"
		}

		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	outputWriter.Flush()

	fmt.Println("Conversion done")
}
