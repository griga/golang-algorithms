package algorithms_test

import (
	"cybe/algorithms"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {

	ll := algorithms.LinkedList{}

	ll.Append(102)
	ll.Append(103)
	ll.Prepend(101)
	ll.Prepend(100)
	fmt.Println(ll)
	assert.Equal(t, ll.Length, 4)

	node, _ := ll.Get(3)
	assert.Equal(t, node.Value, 103)

}
