package _io

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var lineNumber = flag.Bool("n", false, "print line number.")

func cat(r *bufio.Reader) {
	nline := 1
	for {
		r, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if *lineNumber {
			fmt.Fprintf(os.Stdout, "%d %s", nline, r)
			nline++
		} else {
			fmt.Fprintf(os.Stdout, "%s", r)
		}
	}
}

func TestReadByBuf() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}
}
