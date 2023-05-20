package algorithms_test

import (
	"cybe/algorithms"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {

	list := algorithms.LinkedList[int]{}

	list.Append(5)
	list.Append(7)
	list.Append(9)
	fmt.Println(list)

	v, _ := list.Get(2)
	assert.Equal(t, 9, *v)

	v, _ = list.RemoveAt(1)
	assert.Equal(t, 7, *v)
	assert.Equal(t, 2, list.Length)
	fmt.Println(list)

	list.Append(11)
	v, _ = list.RemoveAt(1)
	assert.Equal(t, 9, *v)
	v, _ = list.RemoveAt(0)
	assert.Equal(t, 5, *v)
	v, _ = list.RemoveAt(0)
	assert.Equal(t, 11, *v)
	assert.Equal(t, 0, list.Length)
	fmt.Println(list)

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)
	v, _ = list.Get(2)
	assert.Equal(t, 5, *v)
	v, _ = list.Get(0)
	assert.Equal(t, 9, *v)
	fmt.Println(list)

	list.Remove(9)
	assert.Equal(t, 2, list.Length)
	v, _ = list.Get(0)
	assert.Equal(t, 7, *v)
	fmt.Println(list)

}
