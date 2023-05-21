package algorithms

import "fmt"

type BinaryNode[T comparable] struct {
	Value T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}

func (node *BinaryNode[T]) PreOrderTraverse(path *[]T) *[]T {
	if node == nil {
		return path
	}
	*path = append(*path, node.Value)
	node.Left.PreOrderTraverse(path)
	node.Right.PreOrderTraverse(path)
	return path
}

func (node *BinaryNode[T]) InOrderTraverse(path *[]T) *[]T {
	if node == nil {
		return path
	}
	node.Left.InOrderTraverse(path)
	*path = append(*path, node.Value)
	node.Right.InOrderTraverse(path)
	return path
}

func (node *BinaryNode[T]) PostOrderTraverse(path *[]T) *[]T {
	if node == nil {
		return path
	}
	node.Left.PostOrderTraverse(path)
	node.Right.PostOrderTraverse(path)
	*path = append(*path, node.Value)
	return path
}

func (node *BinaryNode[T]) BreadthFirstSearch(needle T) bool {
	queue := Queue[BinaryNode[T]]{}
	queue.Enqueue(*node)
	for queue.Length > 0 {
		curr, _ := queue.Deque()
		if curr.Value == needle {
			return true
		}
		if curr.Left != nil {
			queue.Enqueue(*curr.Left)
		}
		if curr.Right != nil {
			queue.Enqueue(*curr.Right)
		}
	}
	return false
}

func (node *BinaryNode[T]) DepthFirstSearch(needle T) bool {
	fmt.Println(node.Value)

	result := false
	if node.Value == needle {
		result = true
	}
	if node.Left != nil && !result {
		result = node.Left.DepthFirstSearch(needle)
	}
	if node.Right != nil && !result {
		result = result || node.Right.DepthFirstSearch(needle)
	}
	return result
}
