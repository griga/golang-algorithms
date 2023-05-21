package algorithms

import "fmt"

type StackNode[T comparable] struct {
	value T
	next  *StackNode[T]
}

type Stack[T comparable] struct {
	head   *StackNode[T]
	Length int
}

func (s *Stack[T]) Push(v T) {
	s.Length++
	n := StackNode[T]{value: v}
	detach := s.head
	s.head = &n
	if detach != nil {
		s.head.next = detach
	}

}
func (s *Stack[T]) Pop() (*T, error) {
	if s.head == nil {
		return nil, fmt.Errorf("stack is empty, can't pop")
	}
	detach := s.head
	if detach.next != nil {
		s.head = detach.next
	} else {
		s.head = nil
	}
	s.Length--
	return &detach.value, nil
}

func (s Stack[T]) String() string {
	if s.head == nil {
		return "(Empty Stack)"
	}
	c := s.head
	r := "("
	for c.next != nil {
		r += fmt.Sprintf("[%v]", c.value)
		c = c.next
	}
	r += fmt.Sprintf("[%v])", c.value)
	return r
}
