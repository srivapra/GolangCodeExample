package main

import "fmt"

type Queue struct {
	elements []int
}

func (q *Queue) EnQueue(n int) {
	q.elements = append(q.elements, n)
}

func (q *Queue) DeQueue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is Empty")
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func main() {
	queue := &Queue{}
	queue.EnQueue(20)
	queue.EnQueue(10)
	queue.EnQueue(50)
	queue.EnQueue(30)
	queue.EnQueue(70)

	fmt.Println("Queue Elements ", queue)

	for !queue.IsEmpty() {
		element, err := queue.DeQueue()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Dequeued ", element)

	}

}
