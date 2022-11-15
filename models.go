package leetCode

const (
	Zero  = 0
	Empty = ""
)

// ListNode defines a normal linked list node
type ListNode struct {
	Val  int
	Next *ListNode
}

// RandomNode defines a linked list node which contains a random node
type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

// DoubleListNode defines a double linked list node
type DoubleListNode struct {
	Val   int
	Prev  *DoubleListNode
	Next  *DoubleListNode
	Child *DoubleListNode
}
