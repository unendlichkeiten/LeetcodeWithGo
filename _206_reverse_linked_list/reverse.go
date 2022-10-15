package _206_reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curNode, nextNode := (*ListNode)(nil), head

	for nextNode != nil {
		tmp := nextNode
		nextNode = nextNode.Next
		tmp.Next = curNode
		curNode = tmp
	}

	return curNode
}
