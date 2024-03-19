// Program to reverse the LL

package main

import "fmt"

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

func (l *linkedList) traverse(head *node) {
	current := head
	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}

}

func (l *linkedList) reverseLL() *node {
	temp := l.head
	var prev, front *node

	for temp != nil {
		front = temp.next
		temp.next = prev
		prev = temp
		temp = front
	}
	return prev

}

func main() {
	list := &linkedList{}
	list.addNode(10)
	list.addNode(20)
	list.addNode(30)
	list.addNode(40)
	list.addNode(50)

	fmt.Println("Original Linked List:")
	list.traverse(list.head)
	reversedHead := list.reverseLL()

	fmt.Println("Reversed Linked List:")

	list.traverse(reversedHead)
}
