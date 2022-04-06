package main

import (
	"fmt"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int, bool) {

	l := len(s)

	if l == 0 {
		return s, -1, false
	}
	return s[:l-1], s[l-1], true
}

func main() {
	s := make(stack, 0)
	s = s.Push(1)
	s = s.Push(2)
	s = s.Push(3)

	s, p, flag := s.Pop()
	if flag {
		fmt.Println(p)
	}

	s, k, flag := s.Pop()
	if flag {
		fmt.Println(k)
	}

	s, z, flag := s.Pop()
	if flag {
		fmt.Println(z)
	}

	s, a, flag := s.Pop()
	if flag {
		fmt.Println(a)
	} else {
		fmt.Println("Stack is Empty")
	}
}
