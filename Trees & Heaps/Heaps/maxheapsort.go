package main

import "log"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	arr []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.arr = append(h.arr, key)
	h.maxHeapifyUp(len(h.arr) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.arr[0]
	l := len(h.arr) - 1

	// when the array is empty
	if len(h.arr) == 0 {
		log.Println("cannot extract from a slice with length of 0")
		return -1
	}

	// take out the las index and put it in the root
	h.arr[0] = h.arr[l]
	h.arr = h.arr[:l]

	h.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp : will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.arr[parent(index)] < h.arr[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex :=  len(h.arr) - 1
	l, r := left(index), right(index)
	var childToCompare int

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.arr[l] > h.arr[r] { // when left child is greater
			childToCompare = l
		} else { // when right child is greater
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller
		if h.arr[index] < h.arr[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// parent get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// left get the left child index
func left(i int) int {
	return (2 * i) + 1
}

// right get the right child index
func right(i int) int {
	return (2 * i) + 2
}

// swap swaps keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.arr[i1], h.arr[i2] = h.arr[i2], h.arr[i1]
}

func main() {
	m := &MaxHeap{}
	log.Println(m)
	// buildHeap := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	buildHeap := []int{4, 7, 8, 2, 9, 3, 1, 6, 10}
	for _, v := range buildHeap {
		m.Insert(v)
	}
	log.Println(m)

	var sorted []int
	for i := 0; i < 5; i++ {
		v := m.Extract()
		sorted = append(sorted, v)
	}
	log.Println(m)
	log.Println(sorted)
}