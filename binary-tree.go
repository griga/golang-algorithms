package algorithms

type BinaryNode[T comparable] struct {
	Value T
	Right *BinaryNode[T]
	Left  *BinaryNode[T]
}

func (curr *BinaryNode[T]) PreOrder(path *[]T) *[]T {
	if curr == nil {
		return path
	}
	*path = append(*path, curr.Value)
	curr.Left.PreOrder(path)
	curr.Right.PreOrder(path)
	return path
}

func (curr *BinaryNode[T]) InOrder(path *[]T) *[]T {
	if curr == nil {
		return path
	}
	curr.Left.InOrder(path)
	*path = append(*path, curr.Value)
	curr.Right.InOrder(path)
	return path
}

func (curr *BinaryNode[T]) PostOrder(path *[]T) *[]T {
	if curr == nil {
		return path
	}
	curr.Left.PostOrder(path)
	curr.Right.PostOrder(path)
	*path = append(*path, curr.Value)
	return path
}
