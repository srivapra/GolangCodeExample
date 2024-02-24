package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	elements []int
	lock     sync.Mutex
}

func (s *Stack) isEmpty() bool {
	if len(s.elements) == 0 {
		fmt.Println("The stack is empty ")
		return true
	}
	return false
}

func (s *Stack) Push(element int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.elements = append(s.elements, element)

}

func (s *Stack) Pop() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isEmpty() {
		return 0
	}
	lastElement := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return lastElement
}

func main() {

	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	fmt.Println("Stack elements after pushing the elements into the stack : ", stack.elements)

	stack.Pop()
	fmt.Println("Stack elements after pop operation : ", stack.elements)

	stack.Pop()
	fmt.Println("Stack elements after second pop operation : ", stack.elements)

	stack.Pop()
	fmt.Println("Stack elements after third pop operation : ", stack.elements)

	stack.Pop()
	fmt.Println("Stack elements after fourth pop operation : ", stack.elements)

	stack.Pop()
	fmt.Println("Stack elements after fifth pop operation : ", stack.elements)

	stack.Pop()
	fmt.Println("Stack elements after sixth pop operation : ", stack.elements)

}
