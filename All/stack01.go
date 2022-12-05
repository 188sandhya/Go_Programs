package main

import "fmt"

type Element struct {
	value interface{}
	next  *Element
}

type Stack struct {
	top  *Element
	size int
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(data interface{}) {
	s.top = &Element{
		value: data,
		next:  s.top,
	}
	s.size++
}

func (s *Stack) Pop() (data interface{}) {
	if s.size > 0 {
		data, s.top = s.top.value, s.top.next
		s.size--
		return
	}

	return nil
}

func main() {
	stack := new(Stack)
	stack.Push("firrst ")
	stack.Push("second ")
	stack.Push("third ")

	for stack.Len() > 0 {
		// We have to do a type assertion because we get back a variable of type
		// interface{} while the underlying type is a string.

		fmt.Printf("%s", stack.Pop().(string))
	}
	fmt.Println()
}
