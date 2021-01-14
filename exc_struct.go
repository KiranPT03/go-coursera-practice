package main

import "fmt"

func main() {
	type person struct {
		name    string
		surname string
	}

	p1 := person{
		"Kiran",
		"Tavadare",
	}

	p2 := new(person)

	p2.name = "Amit"
	p2.surname = "Tavadare"

	p3 := person{
		name:    "Prakash",
		surname: "Tavadare",
	}

	fmt.Println(p3)
	fmt.Println(*p2)
	fmt.Println(p1)
}
