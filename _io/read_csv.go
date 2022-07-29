package _io

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type priceChart struct {
	title  string
	price  string
	number string
}

func ReadCSV() {
	datas := make([]*priceChart, 0)

	//filePath := filepath.Clean(".\\_io\\testcase\\products.txt")

	file, err := os.Open(".\\src\\the-way-to-go-code\\_io\\testcase\\products.txt")
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(file)
	r.Comma = ';'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		data := new(priceChart)
		data.title = record[0]
		data.price = record[1]
		data.number = record[2]
		datas = append(datas, data)
	}

	for _, v := range datas {
		fmt.Printf("%+v\n", v)
	}
}
