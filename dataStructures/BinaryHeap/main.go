package main

type minHeap struct {
	array []int
}

func parent(idx int) int {
	return (idx - 1) / 2
}
func leftChild(pIdx int) int {
	return (2 * pIdx) + 1
}
func rightChild(pIdx int) int {
	return (2 * pIdx) + 2
}
func (h *minHeap) swap(idx1, idx2 int) {
	h.array[idx1], h.array[idx2] = h.array[idx2], h.array[idx1]
}
func (h *minHeap) heapifyUp(idx int) {
	if h.array[parent(idx)]>h.array[idx]{
		
	}
}
func main() {

}
