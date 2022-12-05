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

func (s *Stack) Push(data interface{}) {
	fmt.Println("check s:", s)
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
	stack.Push("one ")
	stack.Push("two ")
	stack.Push("three ")
	stack.Push("four ")

	fmt.Println("check stack.size:", stack.size)
	for stack.size > 0 {
		fmt.Printf("%s", stack.Pop())
	}
	fmt.Println()

}
