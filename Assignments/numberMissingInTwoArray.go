package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{0, 2, 3, 5, 4}

	res := -1

	//Loop through both arrays and check which is missing using two loops.
	for i := 0; i < len(arr1); i++ {
		flag := false
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				flag = true
			}
		}
		if flag == false {
			res = arr1[i]
			break
		}
	}

	if res != -1 {
		fmt.Printf("Missing Number is %d", res)
		fmt.Println()
	} else {
		fmt.Println("Both Arrays are same")
	}

}
