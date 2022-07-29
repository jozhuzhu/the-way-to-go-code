package _io

import (
	"fmt"
	"os"
	"strings"
)

func TESTCLI() {
	people := strings.Join(os.Args[1:], " ")

	fmt.Printf("Hello %s!", people)
}
