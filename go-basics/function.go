package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous fn call")
	}()

	// values := [...]int {1,2,3,3,3}
	sum := sum(1, 2, 3, 4, 4)
	fmt.Printf("%v", sum)
	fmt.Println()

}

func sum(values ...int) int {

	sum := 0
	for _, v := range values {
		sum = sum + v
	}
	return sum
}
