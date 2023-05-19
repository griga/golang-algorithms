package algorithms_test

import (
	"cybe/algorithms"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	var cases = map[int]bool{
		69:    true,
		70:    false,
		1336:  false,
		69420: true,
		69421: false,
		1:     true,
		0:     false,
	}

	for needle, expect := range cases {
		assert.Equal(t, expect, algorithms.BinarySearch(&foo, needle))
	}

}
