package main

import (
	"fmt"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func main() {
	s := "()[]{}"

	if len(s)%2 != 0 {
		fmt.Println("No Matching Parenthesis Available")
		return
	}

	var stack Stack

	indicator := true

	for i := 0; i < len(s); i++ {
		if s[i] == ')' && !stack.IsEmpty() {
			val, flag := stack.Pop()
			if !(flag == true && val == "(") {
				indicator = false
				break
			}
		} else if s[i] == '}' && !stack.IsEmpty() {
			val, flag := stack.Pop()
			if !(flag == true && val == "{") {
				indicator = false
				break
			}
		} else if s[i] == ']' && !stack.IsEmpty() {
			val, flag := stack.Pop()
			if !(flag == true && val == "[") {
				indicator = false
				break
			}
		} else {
			stack.Push((string(s[i])))
		}
	}

	if indicator == true {
		fmt.Println("Matching Parenthesis Available")
	} else {
		fmt.Println("No Matching Parenthesis Available")
	}

}
