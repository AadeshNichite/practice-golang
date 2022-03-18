package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Println("Enter a string: ")

	//Takign string as an input
	fmt.Scanln(&s)

	if isPalindrom(s) {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Not an Palindrom")
	}
}

// Function which checks a string is palindrom or not and return accordingly
func isPalindrom(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
