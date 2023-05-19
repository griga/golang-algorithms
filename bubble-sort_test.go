package algorithms_test

import (
	"cybe/algorithms"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	data := []int{9, 3, 7, 4, 69, 420, 42}
	algorithms.BubbleSort(&data)
	expect := []int{3, 4, 7, 9, 42, 69, 420}
	assert.ElementsMatch(t, data, expect)
	assert.Equal(t, 1, 1)
}
