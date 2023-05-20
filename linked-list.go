package algorithms

import (
	"fmt"
)

type ListNode struct {
	Prev  *ListNode
	Next  *ListNode
	Value int
}

type LinkedList struct {
	Head   *ListNode
	Tail   *ListNode
	Length int
}

func (list *LinkedList) Prepend(value int) {
	node := &ListNode{Value: value}
	if list.Length == 0 {
		list.Head, list.Tail = node, node
	} else {
		next := list.Head
		next.Prev = node
		node.Next = next
		list.Head = node
	}
	list.Length++
}

func (list *LinkedList) InsertAt(idx int, v int) error {
	if idx > list.Length {
		return fmt.Errorf("can't InsertAt, idx %v is unavailable", idx)
	} else if idx == list.Length {
		list.Append(v)
		return nil
	} else if idx == 0 {
		list.Prepend(v)
		return nil
	}
	node := ListNode{Value: v}
	curr := list.Head
	for i := 0; i < idx; i++ {
		curr = curr.Next
	}
	prev := curr.Prev
	prev.Next = &node
	node.Prev = prev
	node.Next = curr
	curr.Prev = &node
	list.Length += 1
	return nil
}

func (list *LinkedList) Append(value int) {
	node := &ListNode{Value: value}
	if list.Length == 0 {
		list.Head, list.Tail = node, node
	} else {
		prev := list.Tail
		prev.Next = node
		node.Prev = prev
		list.Tail = node
	}
	list.Length += 1
}

func (list *LinkedList) Get(idx int) (*int, error) {
	if idx > list.Length-1 {
		return nil, fmt.Errorf("can't Get, idx %v is unavailable", idx)
	}
	current := list.Head
	for i := 0; i < idx; i++ {
		current = current.Next
	}
	return &current.Value, nil
}

func (list *LinkedList) remove(node *ListNode) *int {
	prev, next := node.Prev, node.Next
	if list.Length == 1 { // 1 node list
		list.Head, list.Tail = nil, nil
	} else if prev == nil { // head
		list.Head = next
	} else if next == nil {
		list.Tail = prev
	} else {
		prev.Next = next
		next.Prev = prev
	}
	list.Length--
	return &node.Value
}

func (list *LinkedList) RemoveAt(idx int) (*int, error) {
	if idx > list.Length-1 {
		return nil, fmt.Errorf("can't RemoveAt, idx %v is unavailable", idx)
	}
	curr := list.Head
	for i := 0; i < idx; i++ {
		curr = curr.Next
	}
	return list.remove(curr), nil
}

func (list LinkedList) String() string {
	if list.Head == nil {
		return "(Empty LinkedList)"
	}
	current := list.Head
	result := ""
	for current.Next != nil {
		result += fmt.Sprintf("(%v)-", current.Value)
		current = current.Next
	}
	result += fmt.Sprintf("(%v)", current.Value)
	return result
}
