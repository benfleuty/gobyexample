package main

import (
	"fmt"
	"maps"
	//"maps"
)

func main() {
	m := make(map[string]int)
	m["key1"] = 7
	m["key2"] = 13

	fmt.Println("map:", m)

	val1 := m["key1"]
	fmt.Println("val1:", val1)

	val3 := m["key3"]
	fmt.Println("val3:", val3)

	fmt.Println("len:", len(m))

	delete(m, "key2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["key2"]
	fmt.Println("was key present?", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n2)

	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	} else {
		fmt.Println("n != n2")
	}
}
