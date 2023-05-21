package algorithms_test

import (
	"cybe/algorithms"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	heap := algorithms.MinHeap[int]{}
	assert.Equal(t, 0, heap.Length)

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)
	assert.Equal(t, 8, heap.Length)

	v, _ := heap.Delete()
	assert.Equal(t, 1, *v)
	v, _ = heap.Delete()
	assert.Equal(t, 3, *v)
	v, _ = heap.Delete()
	assert.Equal(t, 4, *v)
	v, _ = heap.Delete()
	assert.Equal(t, 5, *v)
	assert.Equal(t, 4, heap.Length)
	v, _ = heap.Delete()
	assert.Equal(t, 7, *v)
	v, _ = heap.Delete()
	assert.Equal(t, 8, *v)
	v, _ = heap.Delete()
	assert.Equal(t, 69, *v)
	v, _ = heap.Delete()
	assert.Equal(t, 420, *v)
	assert.Equal(t, 0, heap.Length)
}
