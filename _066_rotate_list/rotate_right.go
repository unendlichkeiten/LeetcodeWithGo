package _066_rotate_list

import leetCode "github.com/unendlichkeiten/LeetcodeWithGo"

func rotateRight(head *leetCode.ListNode, k int) *leetCode.ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	tmpNode, headLens := head, 0
	for tmpNode != nil { // 找到尾节点
		tmpNode = tmpNode.Next
		headLens++
	}

	first, second := head, head
	for i := k % headLens; i > 0; i-- { // 优先指针起始位置
		first = first.Next
	}

	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	first.Next = head
	head = second.Next
	second.Next = nil

	return head
}
