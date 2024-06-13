package main

import "fmt"

// MaxHeap struct has a slice that holds the array

type MaxHeap struct {
	array []int
}

// Insert adds and element to the heap

func (h *MaxHeap) Insert(key int) {

	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap

func (h *MaxHeap) maxHeapifyUp(index int) {

	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}

}

// get parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child

// Parent index x 2 + 1 = left child index

func left(i int) int {

	return 2*i + 1
}

// Parent index x 2 + 2 = right child index

func right(i int) int {

	return 2*i + 2
}

// swap keys in the array

func (h *MaxHeap) swap(i1, i2 int) {

	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// Extract the largest key

func (h *MaxHeap) Extract() int {
	extracted := h.array[0]

	// error when the array is empty
	if len(h.array) == 0 {
		fmt.Println("cannot extract beacuse array length is 0")
		return -1
	}

	// taking out the last index and puting it as root
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]

	h.maxHeapifyDown(0)

	return extracted
}

// max HeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {

	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child

	for l <= lastIndex {
		if l == lastIndex { // when index is the only child

			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger

			childToCompare = l
		} else { // when right child is larger

			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}

	}
}

func main() {

	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 40, 4, 89}

	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println("adding value: ", v, " Heap: ", m)
	}

	fmt.Println(buildHeap)

	g := &MaxHeap{}

	buildHeap1 := []int{10, 34, 23, 3, 323, 33, 11, 9, 20, 40, 4, 89}

	for _, v := range buildHeap1 {
		g.Insert(v)
		fmt.Println("adding value: ", v, " Heap: ", m)
	}

	for i := 0; i < 5; i++ {
		g.Extract()
		fmt.Println(g)

	}

}

// O(log n)
