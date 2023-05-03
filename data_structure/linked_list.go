package data_structure

type Node[T any] struct {
	Value T
	Next  *Node[T]
	Back  *Node[T]
}

type ListNode[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func (l *ListNode[T]) Prepend(node *Node[T]) {
	second := l.head
	l.head = node
	l.head.Next = second
	l.length++
}
