package main

import (
	"fmt"
)

func main() {

	// var arr[3] int  - way 1
	arr := [3]int{1, 2, 3} // way 2
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v ", arr[i])
	}
	cpp := [3]int{}
	cpp = arr
	cpp[1] = 3

	fmt.Printf("%v, %T", arr, arr)
	fmt.Println(cpp)
	fmt.Println()

	slices := make([]int, 0) //way1
	slices = append(slices, 1, 2, 3)
	slice2 := slices

	slice3 := arr[2:]
	fmt.Println(slice3)
	slice2[1] = 4
	fmt.Println(slice2)
	fmt.Println(slices[:len(slices)-1])
	fmt.Printf("%v, %T", slices, slices)
	fmt.Println()
}
