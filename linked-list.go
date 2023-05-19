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

func (ll LinkedList) String() string {
	if ll.Head == nil {
		return "(Empty LinkedList)"
	}
	current := ll.Head
	result := ""
	for current.Next != nil {
		result += fmt.Sprintf("(%v)-", current.Value)
		current = current.Next
	}
	result += fmt.Sprintf("(%v)", current.Value)
	return result
}

func (ll *LinkedList) Prepend(value int) {
	node := &ListNode{Value: value}
	if ll.Length == 0 {
		ll.Head = node
		ll.Tail = node
	} else {
		postHead := ll.Head
		postHead.Prev = node
		node.Next = postHead
		ll.Head = node
	}
	ll.Length += 1
}

func (ll *LinkedList) Append(value int) {
	node := &ListNode{Value: value}
	if ll.Length == 0 {
		ll.Head = node
		ll.Tail = node
	} else {
		preTail := ll.Tail
		preTail.Next = node
		node.Prev = preTail
		ll.Tail = node
	}
	ll.Length += 1
}

func (ll *LinkedList) Get(idx int) (*ListNode, error) {
	if idx > ll.Length-1 {
		return nil, fmt.Errorf("idx %v is unavailable", idx)
	}
	current := ll.Head
	for i := 0; i < idx; i++ {
		current = current.Next
	}
	return current, nil
}

func (ll *LinkedList) Remove(ln *ListNode) *ListNode {
	lnPrev, lnNext := ln.Prev, ln.Next

	if lnPrev != nil {
		lnPrev.Next = lnNext
		if lnPrev.Next == nil {
			ll.Tail = lnPrev
		}
	}

	if lnNext != nil {
		lnNext.Prev = lnPrev
		if lnNext.Prev == nil {
			ll.Head = lnPrev
		}
	}
	// if lnNext.Next == n

	return ln
}
func (ll *LinkedList) RemoveAt(idx int) (*ListNode, error) {
	ln, err := ll.Get(idx)
	if err != nil {
		return nil, err
	}
	return ll.Remove(ln), nil
}
