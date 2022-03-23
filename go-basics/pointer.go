package main

import "fmt"

type myStruct struct {
	name string
}

func main() {

	a := 42
	var b *int = &a
	fmt.Println(*b)

	ms := myStruct{
		name: "b",
	}

	p := &ms

	fmt.Println(*&p.name)
}
