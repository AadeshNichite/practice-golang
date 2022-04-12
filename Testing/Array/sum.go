package main

import "fmt"

func Sum(arr [5]int) int {

	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	// fmt.Println(sum)

	return sum
}

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println((numbers))
}
