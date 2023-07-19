package main

import "fmt"

type Stack struct {
	elements []int
}

func (s *Stack) Push(n int) {
	s.elements = append(s.elements, n)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return " ", fmt.Errorf("Stack is empty")
	}
	lastindex := len(s.elements) - 1
	element := s.elements[lastindex]
	s.elements = s.elements[:lastindex]
	return element, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	stack := &Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(40)
	stack.Push(100)
	fmt.Println("Element pushed into stack : ", stack)

	// Pop elements from the stack
	for !stack.IsEmpty() {
		element, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Popped:", element)
	}
}
