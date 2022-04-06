package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 2500, 25, 12, 767, 21}

	//Initially we take max as first value and secondmax is -1.
	max := arr[0]
	secondMax := -1

	/* Will go for each array elements. If the value is greater than max then will assign
	previous max to secondMax and make new value max. If the value is less than max and greater
	than second max then will make it secondmax */
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			secondMax = max
			max = arr[i]
		} else if arr[i] > secondMax && arr[i] < max {
			secondMax = arr[i]
		}
	}

	fmt.Println(max, secondMax)
}
