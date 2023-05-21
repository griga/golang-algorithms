package algorithms

import "fmt"

type MinHeap struct {
	Length int
	data   []int
}

func (heap *MinHeap) Insert(value int) {
	heap.data = append(heap.data, value)
	heap.hepifyUp(heap.Length)
	heap.Length++
}

func (heap *MinHeap) Delete() (int, error) {
	if heap.Length == 0 {
		return -1, fmt.Errorf("cant delete from empty heap")
	}

	out := heap.data[0]
	heap.Length--
	if heap.Length == 0 {
		heap.data = []int{}
		return out, nil
	}

	heap.data[0] = heap.data[heap.Length]
	heap.heapifyDown(0)
	return out, nil
}

func (heap *MinHeap) heapifyDown(idx int) {
	li := heap.left(idx)
	ri := heap.right(idx)
	if idx >= heap.Length || li >= heap.Length {
		return
	}
	cv := heap.data[idx]
	lv := heap.data[li]
	rv := heap.data[ri]
	if lv > rv && cv > rv {
		heap.data[idx], heap.data[ri] = rv, cv
		heap.heapifyDown(ri)
	} else if rv > lv && cv > lv {
		heap.data[idx], heap.data[li] = lv, cv
		heap.heapifyDown(li)
	}
}

func (heap *MinHeap) hepifyUp(idx int) {
	if idx == 0 {
		return
	}
	p := heap.parent(idx)
	pv := heap.data[p]
	v := heap.data[idx]

	if pv > v {
		heap.data[p], heap.data[idx] = heap.data[idx], heap.data[p]
		heap.hepifyUp(p)
	}
}

func (heap *MinHeap) parent(idx int) int {
	return (idx - 1) / 2
}

func (heap *MinHeap) left(idx int) int {
	return 2*idx + 1
}

func (heap *MinHeap) right(idx int) int {
	return 2*idx + 2
}
