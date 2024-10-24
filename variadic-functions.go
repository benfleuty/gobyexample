package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	res := sum(nums...)

	fmt.Println("Sum of ", nums, " is ", res)
}

func sum(nums ...int) int {
	fmt.Println("Incoming values: ", nums)
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}
