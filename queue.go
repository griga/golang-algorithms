package algorithms

import "fmt"

type QueueNode[T comparable] struct {
	value T
	next  *QueueNode[T]
}

type Queue[T comparable] struct {
	Length int
	head   *QueueNode[T]
	tail   *QueueNode[T]
}

func (q *Queue[T]) Enqueue(value T) {
	attach := QueueNode[T]{value: value}
	q.Length++
	if q.head == nil {
		q.head = &attach
	}
	if t := q.tail; t != nil {
		t.next = &attach
	}
	q.tail = &attach
}

func (q *Queue[T]) Deque() (*T, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("can't Dequeue because Queue is empty")
	}
	q.Length--
	detach := q.head
	q.head = detach.next
	detach.next = nil
	return &detach.value, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue[T]) Pick() (*T, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("can't Pick because Queue is empty")
	}
	return &q.head.value, nil
}

func (q Queue[T]) String() string {
	if q.head == nil {
		return "(Empty Queue)"
	}
	current := q.head
	result := ""
	for current.next != nil {
		result += fmt.Sprintf("(%v)<-", current.value)
		current = current.next
	}
	result += fmt.Sprintf("(%v)", current.value)
	return result
}
