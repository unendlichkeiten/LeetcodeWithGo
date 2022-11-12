package _142_linked_list_cycle_ii

import leetCode "github.com/unendlichkeiten/LeetcodeWithGo"

func detectCycle(head *leetCode.ListNode) *leetCode.ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
