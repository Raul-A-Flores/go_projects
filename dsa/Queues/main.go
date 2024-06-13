package main

import "fmt"

type Queue struct {
	items []int
}

// Enqueue

func (q *Queue) Enqueue(i int) {

	q.items = append(q.items, i)

}

// Dequeue

func (q *Queue) Dequeue() int {

	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove

}

func main() {

	myQueue := Queue{}

	myQueue.Enqueue(20)
	myQueue.Enqueue(34)
	myQueue.Enqueue(23)
	myQueue.Enqueue(2320)
	myQueue.Enqueue(2023)
	myQueue.Enqueue(120)
	myQueue.Enqueue(20)
	myQueue.Enqueue(26)
	myQueue.Dequeue()
	myQueue.Dequeue()

	fmt.Println(myQueue)

}
