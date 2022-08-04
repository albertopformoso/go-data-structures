// FIFO
package main

import "log"

// Queue : represents a queue that holds a slice
type Queue struct {
	items []int
}

// Enqueue : adds a value at the end
func (q *Queue) Enqueue(i ...int) {
	q.items = append(q.items, i...)
}

// Dequeue : removes a value at the begining and returns the removed value
func (q *Queue) Dequeue() (removedValue int) {
	removedValue = q.items[0]
	q.items = q.items[1:]
	return
}

func main() {
	myQueue := Queue{}
	log.Println(myQueue)
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3, 4, 5, 6)
	log.Println(myQueue)
	myQueue.Dequeue()
	log.Println(myQueue)
}
