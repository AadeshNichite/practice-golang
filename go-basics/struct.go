package main

import "fmt"

type Doctor struct {
	number     int
	first_name string
	last_name  string
}

type Person struct {
	number     int
	first_name string
	last_name  string
	names      []string
}

func main() {
	aDoctor := Doctor{
		number:     1,
		first_name: "ABC",
		last_name:  "DEf",
	}

	fmt.Println(aDoctor)

	pPerson := Person{
		number:     2,
		first_name: "ABCD",
		last_name:  "DEFG",
		names: []string{
			"abc",
			"bcs",
		},
	}
	fmt.Println(pPerson.names)
}
