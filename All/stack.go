package main

import "fmt"

type stack []int

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(data int) {
	*s = append(*s, data)
}

func (s *stack) pop() {
	if s.isEmpty() {
		fmt.Println("stack is empty")
	} else {
		ind := len(*s) - 1
		pop := (*s)[ind]
		fmt.Printf("%d popped ", pop)
		*s = (*s)[:ind]
	}
}

func main() {
	var s stack

	for i := 10; i > 0; i-- {
		s.push(i)
	}

	fmt.Println("ele in stack:", s)
	// fmt.Println("stack length", len(s))

	for len(s) > 0 {
		s.pop()
	}

	fmt.Println("\nafter pop:", s)
}
