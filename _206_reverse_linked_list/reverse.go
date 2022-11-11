package _206_reverse_linked_list

import leetCode "github.com/unendlichkeiten/LeetcodeWithGo"

func reverseList(head *leetCode.ListNode) *leetCode.ListNode {
	curNode, nextNode := (*leetCode.ListNode)(nil), head

	for nextNode != nil {
		tmp := nextNode
		nextNode = nextNode.Next
		tmp.Next = curNode
		curNode = tmp
	}

	return curNode
}

// reverseListV2 使用递归的方法实现
func reverseListV2(head *leetCode.ListNode) *leetCode.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListV2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
