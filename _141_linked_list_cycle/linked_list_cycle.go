package _141_linked_list_cycle

import (
	"fmt"

	leetCode "github.com/unendlichkeiten/LeetcodeWithGo"
)

// hasCycleV1 使用快慢指针的方法
func hasCycleV1(head *leetCode.ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// hasCycleV2 使用 map 的方法
func hasCycleV2(head *leetCode.ListNode) bool {
	pointMap := make(map[string]*leetCode.ListNode)
	for head != nil {
		address := fmt.Sprintf("%p", head)
		if _, ok := pointMap[address]; ok {
			return true
		}
		pointMap[address] = head
		head = head.Next
	}
	return false
}
