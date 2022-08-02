package main

import (
	"log"
	"math"
)

// ArraySize is the size of the hash table array
const ArraySize = 7

// HashTable structure
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket structure
type bucket struct {
	head *bucketNode
}

//bucketNode structure
type bucketNode struct {
	key string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array
func (ht *HashTable) Insert(key string) {
	index := hash(key)
	ht.array[index].insert(key)
}

// Search will take in a key and return true if that key is stored in the hash table
func (ht *HashTable) Search(key string) bool {
	index := hash(key)
	return ht.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (ht *HashTable) Delete(key string) {
	index := hash(key)
	ht.array[index].delete(key)
}

// insert will take in a key, create a node with the key and insert the node in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		log.Println("already exists")
	}
}

// search will take in a key and return true if the bucket has that key
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
// delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash
func hash(key string) int {
	var sum int
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func hash2(key string) int {
	var h int
	// char[0] * 31^n-1 + char[1] * 31^n-2 + ... + char[n-1]
	for i, char := range key {
		h += int(char) * int(math.Pow(31, float64(len(key)-i+1)))
	}
	return h
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	ht := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		ht.Insert(v)
	}
	ht.Delete("STAN")
	log.Println("STAN:", ht.Search("STAN"))
	log.Println("KENNY:", ht.Search("KENNY"))

	// b := &bucket{}
	// b.insert("RANDY")
	// b.delete("RANDY")
	// log.Println(b.search("RANDY"))
}