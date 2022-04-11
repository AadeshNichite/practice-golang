package main

import "fmt"

type Stack []string

//IsEmpty: check if stack is empty
func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}

//Push a new value onto the stack
func (st *Stack) Push(str string) {
	*st = append(*st, str) //Simply append the new value to the end of the stack
}

//Remove top element of stack. Return false if stack is empty.
func (st *Stack) Pop() bool {
	if st.IsEmpty() {
		return false
	} else {
		index := len(*st) - 1 // Get the index of top most element.
		*st = (*st)[:index]   // Remove it from the stack by slicing it off.
		return true
	}
}

//Return top element of stack. Return false if stack is empty.
func (st *Stack) Top() string {
	if st.IsEmpty() {
		return ""
	} else {
		index := len(*st) - 1   // Get the index of top most element.
		element := (*st)[index] // Index onto the slice and obtain the element.
		return element
	}
}

//Function to return precedence of operators
func prec(s string) int {
	if s == "^" {
		return 3
	} else if (s == "/") || (s == "*") {
		return 2
	} else if (s == "+") || (s == "-") {
		return 1
	} else {
		return -1
	}
}

func infixToPostfix(infix string) string {
	var sta Stack
	var postfix string
	for _, char := range infix {
		opchar := string(char)
		// if scanned character is operand, add it to output string
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			postfix = postfix + opchar
		} else if char == '(' {
			sta.Push(opchar)
		} else if char == ')' {
			for sta.Top() != "(" {
				postfix = postfix + sta.Top()
				sta.Pop()
			}
			sta.Pop()
		} else {
			for !sta.IsEmpty() && prec(opchar) <= prec(sta.Top()) {
				postfix = postfix + sta.Top()
				sta.Pop()
			}
			sta.Push(opchar)
		}
	}
	// Pop all the remaining elements from the stack
	for !sta.IsEmpty() {
		postfix = postfix + sta.Top()
		sta.Pop()
	}
	return postfix
}
func main() {
	//infix := "a+b"
	//infix := "a+b*c+d"
	//infix := "a+b*(c^d-e)^(f+g*h)-i" // abcd^e-fgh*+^*+i-
	//infix := "1+2+3*4+5/5-2"
	infix := "2+3*(2^3-5)^(2+1*2)-4" //abcd^e-fgh*+^*+i-
	postfix := infixToPostfix(infix)
	fmt.Printf("%s infix has %s postfix ", infix, postfix)

}
