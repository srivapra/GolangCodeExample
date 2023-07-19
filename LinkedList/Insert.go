package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Insert(value int) {
	n := &Node{value, nil}
	if l.Head != nil {
		l.Head = n
	}
}

func (l *LinkedList) Display() {
	head := l.Head
	for head != nil {
		fmt.Print(head.Data, " ")
		head = head.Next
	}
}

func main() {
	l := LinkedList{}
	l.Insert(5)
	l.Display()
}
