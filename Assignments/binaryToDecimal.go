package main

import (
	"fmt"
)

func main() {

	//Provide Number here to print a decimal of it
	number := 10110

	tmp := number

	powerOfTwo := 1

	result := 0

	for tmp > 0 {
		lastDigit := tmp % 10

		result = result + (powerOfTwo * lastDigit)

		powerOfTwo = powerOfTwo * 2

		tmp = tmp / 10
	}

	//Result it here
	fmt.Println(result)
}
