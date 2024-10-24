package main

import (
	"fmt"
)

func main() {
	res := add2(9, 10)
	fmt.Println("res:", res)

	res = add3(9, 10, 11)
	fmt.Println("res:", res)
}

func add2(a int, b int) int {
	return a + b
}

func add3(a, b, c int) int {
	return a + b + c
}
