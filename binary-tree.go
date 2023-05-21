package algorithms

type BinaryNode[T comparable] struct {
	Value T
	Right *BinaryNode[T]
	Left  *BinaryNode[T]
}

func (curr *BinaryNode[T]) PreOrderTraverse(path *[]T) *[]T {
	if curr == nil {
		return path
	}
	*path = append(*path, curr.Value)
	curr.Left.PreOrderTraverse(path)
	curr.Right.PreOrderTraverse(path)
	return path
}

func (curr *BinaryNode[T]) InOrderTraverse(path *[]T) *[]T {
	if curr == nil {
		return path
	}
	curr.Left.InOrderTraverse(path)
	*path = append(*path, curr.Value)
	curr.Right.InOrderTraverse(path)
	return path
}

func (curr *BinaryNode[T]) PostOrderTraverse(path *[]T) *[]T {
	if curr == nil {
		return path
	}
	curr.Left.PostOrderTraverse(path)
	curr.Right.PostOrderTraverse(path)
	*path = append(*path, curr.Value)
	return path
}

func (head *BinaryNode[T]) BreadthFirstSearch(needle T) bool {
	queue := Queue[BinaryNode[T]]{}
	queue.Enqueue(*head)
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

func (head *BinaryNode[T]) DepthFirstSearch(needle T) bool {
	stack := Stack[BinaryNode[T]]{}
	stack.Push(*head)
	for stack.Length > 0 {
		curr, _ := stack.Pop()
		if curr.Value == needle {
			return true
		}
		if curr.Left != nil {
			stack.Push(*curr.Left)
		}
		if curr.Right != nil {
			stack.Push(*curr.Right)
		}
	}
	return false
}
