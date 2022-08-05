package main

import "log"

// MinHeap struct has a slice that holds the array
type MinHeap struct {
	arr []int
}

// Insert adds an element to the heap
func (h *MinHeap) Insert(key int) {
	h.arr = append(h.arr, key)
	h.minHeapifyUp(len(h.arr) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MinHeap) Extract() int {
	// when the array is empty
	if len(h.arr) == 0 {
		log.Println("cannot extract from a slice with length of 0")
		return -1
	}
	
	extracted := h.arr[0]
	l := len(h.arr) - 1

	// take out the las index and put it in the root
	h.arr[0] = h.arr[l]
	h.arr = h.arr[:l]

	h.minHeapifyDown(0)

	return extracted
}

// maxHeapifyUp : will heapify from bottom to top
func (h *MinHeap) minHeapifyUp(index int) {
	for h.arr[parent(index)] > h.arr[index] {
		h.swap(index, parent(index))
		index = parent(index)
	}
}

// maxHeapifyDown will heapify top to bottom
func (h *MinHeap) minHeapifyDown(index int) {
	for {
		l, r := left(index), right(index)
		if l >= len(h.arr) || l < 0 { // < 0 int overflow
			break
		}

		if r < len(h.arr) && h.arr[l] >= h.arr[r] {
			l = r
		}
		if h.arr[l] >= h.arr[index] {
			break
		}
		h.swap(index, l)
		index = l
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
func (h *MinHeap) swap(i1, i2 int) {
	h.arr[i1], h.arr[i2] = h.arr[i2], h.arr[i1]
}

func main() {
	m := &MinHeap{}
	log.Println(m)
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