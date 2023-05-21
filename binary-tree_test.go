package algorithms_test

import (
	"cybe/algorithms"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree = algorithms.BinaryNode[int]{
	20,
	&algorithms.BinaryNode[int]{
		50,
		&algorithms.BinaryNode[int]{100, nil, nil},
		&algorithms.BinaryNode[int]{30,
			&algorithms.BinaryNode[int]{45, nil, nil},
			&algorithms.BinaryNode[int]{29, nil, nil}},
	},
	&algorithms.BinaryNode[int]{
		10,
		&algorithms.BinaryNode[int]{15, nil, nil},
		&algorithms.BinaryNode[int]{5,
			&algorithms.BinaryNode[int]{7, nil, nil},
			nil},
	},
}

func TestBTPreOrder(t *testing.T) {
	result := *tree.PreOrder(&[]int{})
	expect := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	assert.EqualValues(t, expect, result)
}

func TestBTInOrder(t *testing.T) {
	result := *tree.InOrder(&[]int{})
	expect := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}
	assert.EqualValues(t, expect, result)
}

func TestBTPostOrder(t *testing.T) {
	result := *tree.PostOrder(&[]int{})
	expect := []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}
	assert.EqualValues(t, expect, result)
}
