package _027_palindrome

import (
	"math"

	leetCode "github.com/unendlichkeiten/LeetcodeWithGo"
)

func isPalindrome(head *leetCode.ListNode) bool {

	// 找到中点
	tmp, lenHead := head, leetCode.Zero
	for tmp != nil {
		lenHead++
		tmp = tmp.Next
	}

	midPos, mid := 0, head
	if lenHead%2 == 0 {
		midPos = (int)(math.Ceil(float64(lenHead+1) / 2))
	} else {
		midPos = (int)(math.Ceil(float64(lenHead) / 2))
	}

	for i := 1; i < midPos; i++ {
		mid = mid.Next
	}

	midReverse := reverse(mid)

	for midReverse != nil {
		if midReverse.Val != head.Val {
			return false
		}
		head = head.Next
		midReverse = midReverse.Next
	}

	return true
}

func reverse(head *leetCode.ListNode) *leetCode.ListNode {
	curNode, nextNode := (*leetCode.ListNode)(nil), head

	for nextNode != nil {
		tmp := nextNode
		nextNode = nextNode.Next
		tmp.Next = curNode
		curNode = tmp
	}

	return curNode
}
