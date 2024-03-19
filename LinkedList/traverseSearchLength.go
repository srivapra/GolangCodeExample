// Program that will create a Singly LL, Insert the element into singly LL
// Traverse the LL, find out the length of linked list, search the given value in linked list

package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func (l *linkedList) addNode(val int) {
	newNode := node{val, nil}
	if l.head == nil {
		l.head = &newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = &newNode
	}
}

func (l *linkedList) traverse() {
	current := l.head
	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}

}

func (l *linkedList) length() {
	current := l.head
	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}
	count := 0
	for current != nil {
		current = current.next
		count++
	}
	fmt.Println("Linked List Length is : ", count)
}

func (l *linkedList) search(val int) bool {
	temp := l.head
	if temp == nil {
		fmt.Println("Linked list is empty")
		return false
	}
	for temp != nil {
		if temp.data == val {
			return true
		}
		temp = temp.next
	}
	return false

}

func main() {
	list := &linkedList{}
	list.addNode(10)
	list.addNode(20)
	list.addNode(30)
	list.addNode(40)
	list.addNode(50)

	list.traverse()
	list.length()
	fmt.Println(list.search(2))
}
