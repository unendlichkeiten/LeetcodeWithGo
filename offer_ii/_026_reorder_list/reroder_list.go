package _026_reorder_list

import (
	"math"

	leetCode "github.com/unendlichkeiten/LeetcodeWithGo"
)

func reorderList(head *leetCode.ListNode) {

	// 空链表直接返回
	if head == nil {
		return
	}

	tmp, lenHead := head, leetCode.Zero
	for tmp != nil {
		lenHead++
		tmp = tmp.Next
	}

	// 找到中间节点
	mid, n := head, 0
	if lenHead%2 == 0 {
		n = (int)(math.Ceil(float64(lenHead+1) / 2))
	} else {
		n = (int)(math.Ceil(float64(lenHead) / 2))
	}
	for i := 1; i < n; i++ {
		mid = mid.Next
	}
	// 翻转后的中间链表
	revMidNode := reverse(mid)

	// list merge
	node := head
	for node != tmp && revMidNode.Next != nil {
		tmpNode := node.Next
		node.Next = revMidNode
		revMidNode = revMidNode.Next
		node.Next.Next = tmpNode
		node = node.Next.Next
	}

	return

}

func reverse(head *leetCode.ListNode) *leetCode.ListNode {

	newHead, pre, cur := (*leetCode.ListNode)(nil), (*leetCode.ListNode)(nil), head

	for cur != nil {
		pre = cur
		cur = cur.Next
		pre.Next = newHead
		newHead = pre
	}

	return newHead
}
