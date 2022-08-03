/* https://www.golangprograms.com/golang-program-for-implementation-of-linked-list.html */
package main

import (
	"errors"
	"log"
)

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

func (ll *LinkedList) Insert(key interface{}) {
	n := Node{}
	n.key = key

	if ll.head == nil {
		ll.head = &n
		ll.tail = &n
		ll.len++
		return
	}

	currNode := ll.head
	for currNode.next != nil {
		currNode = currNode.next
	}

	n.prev = currNode
	currNode.next = &n
	ll.tail = &n

	ll.len++
}

func (ll *LinkedList) InsertAt(index int, key interface{}) {
	newNode := Node{}
	newNode.key = key

	if index < 0 {
		log.Printf("index cannot be negative")
		return
	}

	if index > ll.len {
		log.Printf("index out of range")
		return
	}

	if index == 0 && ll.len > 0 {
		n := ll.GetAt(index)
		newNode.next = n
		ll.head = &newNode
		ll.len++
		return
	}

	if index == 0 && ll.len == 0 {
		ll.head = &newNode
		ll.tail = &newNode
		ll.len++
		return
	}

	if index == ll.len {
		prevNode := ll.GetAt(index - 1)
		newNode.prev = prevNode
		prevNode.next = &newNode
		ll.tail = &newNode
		ll.len++
		return
	}

	n := ll.GetAt(index)
	newNode.next = n
	n.prev = &newNode
	prevNode := ll.GetAt(index - 1)
	prevNode.next = &newNode
	newNode.prev = prevNode
	ll.len++
}

func (ll *LinkedList) GetAt(index int) *Node {
	node := ll.head
	if index < 0 {
		log.Printf("index cannot be negative")
		return node
	}
	if index > (ll.len - 1) {
		return nil
	}
	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

func (ll *LinkedList) DeleteAt(index int) error {
	if index < 0 {
		log.Printf("index cannot be negative")
		return errors.New("index cannot be negative")
	}

	if index > ll.len-1 {
		log.Printf("index out of range")
		return errors.New("index out of range")
	}

	if index == 0 && ll.len > 0 {
		nextNode := ll.GetAt(index).next
		ll.head = nextNode

		nxtNxtNode := nextNode.next
		nxtNxtNode.prev = nextNode
		ll.len--
		return nil
	}

	if index == ll.len {
		prevNode := ll.GetAt(index - 1)
		ll.tail = prevNode
		ll.len--
		return nil
	}

	prevNode := ll.GetAt(index).prev
	if prevNode == nil {
		return errors.New("node not found")
	}

	prevNode.next = ll.GetAt(index).next
	ll.len--
	return nil
}

func (ll *LinkedList) Reverse() {
	curr := ll.head
	var nxtInList *Node
	ll.tail, ll.head = ll.head, ll.tail

	for curr != nil {
		nxtInList = curr.next
		curr.next, curr.prev = curr.prev, curr.next
		curr = nxtInList
	}
}

func (ll *LinkedList) Print() {
	list := ll.head
	for list != nil {
		log.Printf("-> %+v", list.key)
		list = list.next
	}
}

func (ll *LinkedList) PrintReverse() {
	list := ll.tail
	for list != nil {
		log.Printf("-> %+v", list.key)
		list = list.prev
	}
}

func main() {
	ll := LinkedList{}
	ll.Insert(1)
	ll.Insert(3)
	ll.Insert(5)
	ll.InsertAt(3, 0)
	ll.DeleteAt(1)

	ll.Reverse()
	ll.PrintReverse()
}
