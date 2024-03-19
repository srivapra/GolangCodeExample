// Program to find the middle of LL

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

func (l *linkedList) middleLL() *node {

	if l.head == nil {
		return nil
	}

	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow

}

func main() {

	list := &linkedList{}
	list.addNode(10)
	list.addNode(20)
	list.addNode(30)
	list.addNode(40)
	list.addNode(50)
	list.addNode(60)

	findMiddle := list.middleLL()
	fmt.Println(findMiddle.data)

}
