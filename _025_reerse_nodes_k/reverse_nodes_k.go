package _025_reverse_nodes_k

import leetCode "github.com/unendlichkeiten/LeetcodeWithGo"

// swapPairs 两两节点交换
func swapPairs(head *leetCode.ListNode, k int) *leetCode.ListNode {
	resList, resListTail, nextNode := (*leetCode.ListNode)(nil), (*leetCode.ListNode)(nil), head

	// 特殊情况处理
	if head == nil || head.Next == nil {
		return head
	}

	tmpNode, tmpNodeTail := head, head
	for n := 1; nextNode != nil; n++ {
		if n%k == 0 {
			// reverse 1->2
			tmpNodeTail = nextNode
			// 这里先将 next 指针指向下一个节点，否则可能导致后面元素无法便利
			nextNode = nextNode.Next
			tmpNodeTail.Next = nil
			tmpNode = reverseList(tmpNode)

			// push 2->1 list into resList tail
			if resList == nil {
				resList = tmpNode
				resListTail = tmpNodeTail
			} else {
				resListTail.Next = tmpNode
			}
			resListTail = getTailNode(tmpNode)

			tmpNode = nextNode
			tmpNodeTail = nextNode
			continue
		}

		nextNode = nextNode.Next
	}

	// 2->1->4->3->5 情况
	if tmpNode != nil {
		resListTail.Next = reverseList(tmpNode)
	}

	return resList
}

func getTailNode(head *leetCode.ListNode) *leetCode.ListNode {
	tmpNode := head

	for tmpNode != nil && tmpNode.Next != nil {
		tmpNode = tmpNode.Next
	}
	return tmpNode
}

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
