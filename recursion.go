package main

import (
	"fmt"
)

func main() {
	fmt.Println("7! = ", factorialOf(7))
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println("The 7th element of the fibonacci sequence is", fib(7))
}

func factorialOf(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialOf(n-1)
}
