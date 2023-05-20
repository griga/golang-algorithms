package algorithms_test

import (
	"cybe/algorithms"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {

	foo := []int{9, 3, 7, 4, 69, 420, 42}
	expected := []int{3, 4, 7, 9, 42, 69, 420}
	algorithms.QuickSort(&foo)
	assert.EqualValues(t, foo, expected)
}
