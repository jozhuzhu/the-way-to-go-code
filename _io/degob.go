package _io

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func ReadFromFile() {
	file, err := os.Open("vcard.gob")
	// 打开文件记得关文件
	defer file.Close()
	if err != nil {
		log.Fatalf("Error in open gob file: %s", err)
	}

	dec := gob.NewDecoder(file)

	var v VCard
	err = dec.Decode(&v)
	if err != nil {
		log.Fatalf("Error in decode gob file: %s", err)
	}

	fmt.Printf("%q,%q: {[{%q, %q, %q}, {%q, %q, %q}], %q}", v.FirstName, v.LastName,
		v.Addresses[0].Type, v.Addresses[0].City, v.Addresses[0].Country,
		v.Addresses[1].Type, v.Addresses[1].City, v.Addresses[1].Country,
		v.Remark)
}
