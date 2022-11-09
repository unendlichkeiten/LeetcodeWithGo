package _029_ordered_cycle_linked_list

import leetCode "github.com/unendlichkeiten/LeetcodeWithGo"

func insert(aNode *leetCode.ListNode, x int) *leetCode.ListNode {
	newNode := &leetCode.ListNode{
		Val:  x,
		Next: nil,
	}

	if aNode == nil {
		newNode.Next = newNode
		return newNode
	}

	curNode, nextNode := aNode, aNode.Next
	// x 比所有元素都大
	for nextNode != aNode {
		// cur <= x <= next
		if curNode.Val <= newNode.Val && newNode.Val <= nextNode.Val {
			curNode.Next = newNode
			newNode.Next = nextNode
			return aNode
		}

		// x >= 最大值 || x <= 最小值
		if curNode.Val > nextNode.Val && newNode.Val >= curNode.Val ||
			curNode.Val > nextNode.Val && newNode.Val <= nextNode.Val {
			curNode.Next = newNode
			newNode.Next = nextNode
			return aNode
		}

		// x <= cur = next
		// next <= x <= cur || next = cur <= x (next == cur)
		curNode = nextNode
		nextNode = nextNode.Next
	}

	// 遍历了所有节点都没有找到满足条件的位置，说明所有元素相等
	curNode.Next = newNode
	newNode.Next = nextNode

	return aNode
}
