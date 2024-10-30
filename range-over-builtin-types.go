package main

import (
	"fmt"
	"strconv"
)

func posInListAsOrdinal(idx int) string {
	pos := idx + 1
	switch pos % 10 {
	case 1:
		return strconv.Itoa(pos) + "st"
	case 2:
		return strconv.Itoa(pos) + "nd"
	case 3:
		return strconv.Itoa(pos) + "rd"
	default:
		return strconv.Itoa(pos) + "th"
	}
}

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
			fmt.Println(posInListAsOrdinal(i), "item in the list")
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println("idx", i, "of go is char", c)
	}

}
