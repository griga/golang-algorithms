package algorithms_test

import (
	"cybe/algorithms"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	s := algorithms.Stack[int]{}
	fmt.Println(s, s.Length)

	s.Push(1)
	s.Push(2)
	assert.Equal(t, s.Length, 2)
	fmt.Println(s, s.Length)

	v, e := s.Pop()
	assert.Equal(t, s.Length, 1)
	assert.Equal(t, *v, 2)
	assert.Equal(t, e, nil)
	fmt.Println(s, s.Length)

	v, e = s.Pop()
	assert.Equal(t, s.Length, 0)
	assert.Equal(t, *v, 1)
	assert.Equal(t, e, nil)
	fmt.Println(s, s.Length)

	v, e = s.Pop()
	assert.Equal(t, s.Length, 0)
	assert.Nil(t, v)
	assert.NotNil(t, e)
	fmt.Println(s, e.Error())

	s.Push(3)
	s.Push(4)
	s.Push(5)
	assert.Equal(t, s.Length, 3)
	fmt.Println(s, s.Length)

}
