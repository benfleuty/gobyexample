package main

import "fmt"

type person struct {
	name string
	age  int
}

type doge struct {
	name   string
	isGood bool
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// creating a new person
	fmt.Println(person{"Bob", 20})
	// fields can be named when creating
	fmt.Println(person{name: "Alice", age: 30})
	// if i don't set a field, it is zeroed
	fmt.Println(person{name: "Fred the ageless"})
	// pointy ann
	fmt.Println(&person{name: "Ann", age: 40})
	// idiomatic
	fmt.Println(newPerson("John"))

	// create sean and access him with dot
	s := person{name: "Sean of Bean", age: 50}
	fmt.Println(s.name)

	// point to sean and get his age
	sp := &s
	fmt.Println(sp.age)

	// happy birthday, sean
	sp.age = 51
	fmt.Println(sp.age)

	fmt.Println("sp and s are the same:", sp.age == s.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true, // :O trailing commas... nice!
	}
	fmt.Println(dog)

	if dog.isGood {
		fmt.Printf("%s is a good boy!", dog.name)
	} else {
		fmt.Printf("Lies, all dogs are good! %s is a good boy!", dog.name)
	}
	fmt.Println()

	var dogs []doge
	dogNames := []string{"Bella", "Henry", "Olive", "Bandit", "Dave"}
	for i := 0; i < len(dogNames); i++ {
		dogs = append(dogs, doge{name: dogNames[i], isGood: i%2 == 0})

	}

	for i, dog := range dogs {
		if dog.isGood {
			fmt.Printf("#%d: %s is a good dog!", i+1, dog.name)
		} else {
			fmt.Printf("#%d: Someone lied and said %s is not good! isGood should not be set to %t!", i+1, dog.name, dog.isGood)
		}
		fmt.Println()
	}
}
