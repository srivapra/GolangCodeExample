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

func (l *linkedList) deleteHead() {
	if l.head == nil {
		return
	}
	newHead := l.head.next
	l.head = newHead
}

func (l *linkedList) deleteTail() {
	if l.head == nil {
		return
	}

	temp := l.head
	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
}

func (l *linkedList) deleteKthElement(k int) {
	if k <= 0 || l.head == nil {
		return
	}

	if k == 1 {
		l.head = l.head.next
		return
	}

	temp := l.head
	var prev *node

	count := 0

	for temp != nil {
		count = count + 1
		if count == k {
			prev.next = prev.next.next
			break
		}
		prev = temp
		temp = temp.next
	}

}

func main() {
	list := &linkedList{}
	list.addNode(10)
	list.addNode(20)
	list.addNode(30)
	list.addNode(40)
	list.addNode(50)

	fmt.Println("Original Linked List : ")
	list.traverse()

	list.deleteHead()
	fmt.Println("Linked List after deleting the head : ")
	list.traverse()

	list.deleteTail()
	fmt.Println("Linked List after deleting the Tail : ")
	list.traverse()

	k := 2
	list.deleteKthElement(k)
	list.deleteKthElement(k)
	list.deleteKthElement(k)
	fmt.Println("Linked List after deleting the kth element : ", k)
	list.traverse()

}
