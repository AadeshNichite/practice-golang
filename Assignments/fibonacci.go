package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Println("Enter a number: ")

	fmt.Scanln(&number)

	first := 0
	second := 1

	fmt.Printf("%v %v", first, second)

	for i := 0; i < number; i++ {
		tmp := first + second
		fmt.Printf(" %v", tmp)
		first = second
		second = tmp
	}
	fmt.Println()
}
