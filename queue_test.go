package algorithms_test

import (
	"cybe/algorithms"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := algorithms.Queue[string]{}
	fmt.Println(q)
	assert.Equal(t, q.IsEmpty(), true)

	q.Enqueue("2")
	q.Enqueue("3")
	q.Enqueue("69")
	q.Enqueue("420")
	fmt.Println(q)
	val, _ := q.Deque()
	fmt.Println(*val)
	val, _ = q.Deque()
	fmt.Println(*val)
	fmt.Println(q)
	q.Enqueue("1323")
	q.Enqueue("2377")
	fmt.Println(q)
	val, _ = q.Deque()
	fmt.Println(*val)
	val, _ = q.Deque()
	fmt.Println(*val)
	val, _ = q.Deque()
	fmt.Println(*val)
	val, _ = q.Deque()
	fmt.Println(*val)
	fmt.Println(q)

	p, err := q.Pick()
	fmt.Println(p)
	fmt.Println(err)

}
