// LIFO
package main

import "log"

// Stack : represents stack that hold a slice
type Stack struct {
	items []int
}

// Push : will add a value at the end
func (s *Stack) Push(i ...int) {
	s.items = append(s.items, i...)
}

// Pop : will remove a value at the end and returns the removed value
func (s *Stack) Pop() (removedValue int) {
	l := len(s.items) - 1
	removedValue = s.items[l]
	s.items = s.items[:l]
	return
}

func main() {
	myStack := Stack{}
	log.Println(myStack)
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3, 4, 5, 6)
	log.Println(myStack)
	myStack.Pop()
	log.Println(myStack)
}
