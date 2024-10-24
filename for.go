package main

import (
	"fmt"
)

func main() {
	i := 1
	for i < 3 {
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop until break is called")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	for x := 5; x < 22; x *= 2 {
		fmt.Println(x)
	}
}
