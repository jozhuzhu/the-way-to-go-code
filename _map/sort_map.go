package _map

import (
	"bytes"
	"fmt"
	"sort"
)

func PrintDrink() {
	drinks := map[string]string{
		"pepsi":         "百事可乐",
		"hedy":          "七喜",
		"apple vinegar": "苹果醋",
		"orangeade":     "橙汁",
	}

	fmt.Println("unsorted:")
	for key, value := range drinks {
		fmt.Printf("%s: %s\n", key, value)
	}

	var drinkNames []string
	for key, _ := range drinks {
		drinkNames = append(drinkNames, key)
	}

	sort.Strings(drinkNames)
	fmt.Println("sorted:")
	for _, key := range drinkNames {
		fmt.Printf("%s: %s\n", key, drinks[key])
	}
}

type drink struct {
	key   string
	value string
}

type drinks []drink

func (d drinks) Len() int {
	return len(d)
}

func (d drinks) Less(i, j int) bool {
	return bytes.Compare([]byte(d[i].key), []byte(d[j].key)) < 0
}

func (d drinks) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// 结构体切片可以更有效生成排序列表
func PrintDrink_1() {
	d := drinks{
		{
			"pepsi", "百事可乐",
		},
		{
			"hedy", "七喜",
		},
		{
			"apple vinegar", "苹果醋",
		},
		{
			"orangeade", "橙汁",
		},
	}

	fmt.Println("unsorted: ")
	for _, d := range d {
		fmt.Printf("%s: %s\n", d.key, d.value)
	}

	fmt.Println("sorted: ")
	sort.Sort(d)
	for _, d := range d {
		fmt.Printf("%s: %s\n", d.key, d.value)
	}
}
