package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Print("one")
	case 2:
		fmt.Print("two")
	case 3:
		fmt.Print("three")
	default:
		fmt.Print("error this number is not supported!")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	case time.Friday:
		fmt.Println("It's Friday, Friday, gotta write go on Friday")
	default:
		fmt.Println("Just another weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 10:
		fmt.Println("Time for breakfast")
	case t.Hour() < 12:
		fmt.Println("Time for second breakfast")
	case t.Hour() > 12 && t.Hour() < 17:
		fmt.Println("Time for lunch")
	case t.Hour() < 21:
		fmt.Println("Tea Time!")
	case t.Hour() <= 24:
		fmt.Println("Isn't it a bit late for food?")
	default:
		fmt.Println("You broke time. ggwp")
	}

	typeof := func(i interface{}) {
		switch t := i.(type) {
		case string:
			fmt.Println("typeof string")
		case int:
			fmt.Println("typeof int")
		// case float64, float32:
		// 	fmt.Println("typeof float32/64")
		// case complex64, complex128:
		// 	fmt.Println("typeof complex64/128")
		case bool:
			fmt.Println("typeof bool")
		default:
			fmt.Println("typeof cannot determine type for:", t)
		}
	}

	typeof(true)
	typeof(1)
	typeof("hey")
	typeof(1e0)
}
