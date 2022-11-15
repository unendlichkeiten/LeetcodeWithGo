package _021_remove_node_from_end

import leetcode "github.com/unendlichkeiten/LeetcodeWithGo"

func removeNthFromEnd(head *leetcode.ListNode, n int) *leetcode.ListNode {
	fast, slow := head, head
	// 倒数第 n 个节点，快指针先走 n 步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 快指针指向最后一个节点
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 删除头节点的情况end
	if fast == nil {
		return head.Next
	}

	// 删除中间节点
	slow.Next = slow.Next.Next

	return head
}
