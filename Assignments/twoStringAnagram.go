package main

import (
	"fmt"
	"strings"
)

func main() {
	flag := true

	s1 := "Listen"
	s2 := "Silent"

	//Will make here both string uppercase.
	s1 = strings.ToUpper(s1)
	s2 = strings.ToUpper(s2)

	arr := [26]int{0}

	for i := 0; i < len(s1); i++ {
		arr[s1[i]-'A']++
	}
	for i := 0; i < len(s2); i++ {
		arr[s2[i]-'A']--
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			flag = false
			break
		}
	}

	if flag == true {
		fmt.Println("Both Strigs are Anagram to each other")
	} else {
		fmt.Println("Both Strigs are not Anagram to each other")
	}
}
