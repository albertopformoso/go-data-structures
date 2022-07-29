/* https://levelup.gitconnected.com/go-singly-linked-lists-with-insertion-deletion-traversal-e9da5bb0dbe1 */
package main

import (
	"errors"
	"log"
)

// Node represents a node of a linked list
type Node struct {
	next *Node
	key interface{}
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
	len int
}

// Insert adds an element to the end of the list
func (l *LinkedList) Insert(key interface{}) {
	// create a new node
	n := Node{}
	n.key = key

	// If LinkedList is empty
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}

	currNode := l.head
	for i := 0; i < l.len; i++ {
		if currNode.next == nil {
			currNode.next = &n
			l.len++
			return
		}
		currNode = currNode.next
	}
}

// InsertAt adds an element on a specific index in the list
func (l *LinkedList) InsertAt(index int, key interface{}) {
	// create a new node
	newNode := Node{}
	newNode.key = key

	// validate the index
	if index < 0 {
		log.Printf("index cannot be negative")
		return
	}

	if index == 0 && l.len == 0 {
		l.head = &newNode
		l.len++
		return
	}

	if index == 0 && l.len > 0 {
		n := l.GetAt(index)
		newNode.next = n
		l.head = &newNode
		l.len++
		return
	}

	if index > l.len {
		log.Println("index out of range")
		return
	}

	n := l.GetAt(index)
	newNode.next = n
	prevNode := l.GetAt(index - 1)
	prevNode.next = &newNode
	l.len++
}

// GetAt returns a node at a given index from a linked list
func (l *LinkedList) GetAt(index int) *Node {
	node := l.head
	if index < 0 {
		log.Println("index cannot be negative")
		return node
	}
	if index > (l.len - 1) {
		return nil
	}
	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

// DeleteAt removes an element on a specific index in the list
func (l *LinkedList) DeleteAt(index int) error {
	// validate the index
	if index < 0 {
		return errors.New("index cannot be negative")
	}
	if index > l.len {
		return errors.New("index out of range")
	}
	if l.len == 0 {
		return errors.New("no nodes in list")
	}

	if index == 0 {
		nextNode := l.GetAt(index).next
		l.head = nextNode
		l.len--
		return nil
	}

	prevNode := l.GetAt(index - 1)
	if prevNode == nil {
		return errors.New("node not found")
	}

	prevNode.next = l.GetAt(index).next
	l.len--
	return nil
}

// DeleteValue deletes a node having a given value from a linked list
func (l *LinkedList) DeleteValue(val interface{}) error {
	node := l.head
	if l.len == 0 {
		return errors.New("empty list")
	}

	for i := 0; i < l.len; i++ {
		if node.key == val {
			if i > 0 {
				prevNode := l.GetAt(i - 1)
				prevNode.next = l.GetAt(i).next
			} else {
				l.head = node.next
			}
			l.len--
			return nil
		}
		node = node.next
	}

	return errors.New("node not found")
}

// Reverse reverses the linked list recursevely
func (l *LinkedList) Reverse() {
	var nxtInList *Node
	current := l.head

	for current != nil {
		next := current.next
		current.next = nxtInList
		nxtInList = current
		current = next
	}

	l.head = nxtInList
}

// Print displays the conents of the LinkedList
func (l *LinkedList) Print() {
	if l.len == 0 {
		log.Println("No nodes in the list")
	}
	list := l.head
	for list != nil {
		log.Printf("-> %+v", list.key)
		list = list.next
	}
}

func main() {
	ll := LinkedList{}
	ll.Insert(1)
	ll.Insert(3)
	ll.Insert(5)
	ll.InsertAt(0, 0)

	err := ll.DeleteAt(0)
	if err != nil {
		log.Println(err.Error())
	}

	err = ll.DeleteValue(3)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Ordered:")
	ll.Print()

	log.Println("Reversed:")
	ll.Reverse()
	ll.Print()

}