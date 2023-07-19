package main

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) AddNode(data int) {
	node := &Node{
		data: data,
		next: nil,
	}
	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current != nil {
			current = current.next
		}
		current.next = node
	}

}

func (l *LinkedList) RemoveNode(data int) {
	if l.head == nil {
		return
	}
	if l.head.data == data {

	}
}
