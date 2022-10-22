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

// reverseListV2 使用递归的方法实现
func reverseListV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListV2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
