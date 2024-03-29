package _023_merge_k_sorted_list

import leetCode "github.com/unendlichkeiten/LeetcodeWithGo"

func mergeKLists(lists []*leetCode.ListNode) *leetCode.ListNode {
	newList := (*leetCode.ListNode)(nil)
	for _, list := range lists {
		newList = mergeTwoLists(newList, list)
	}
	return newList
}

func mergeTwoLists(list1, list2 *leetCode.ListNode) *leetCode.ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	newList, newListTail := (*leetCode.ListNode)(nil), (*leetCode.ListNode)(nil)
	for list1 != nil && list2 != nil {
		// 确定头节点
		if newList == nil {
			if list1.Val <= list2.Val {
				newList, newListTail = list1, list1
				list1 = list1.Next
			} else {
				newList, newListTail = list2, list2
				list2 = list2.Next
			}
			continue
		} else {
			if list1.Val <= list2.Val {
				newListTail.Next = list1
				list1 = list1.Next
				newListTail = newListTail.Next
			} else {
				newListTail.Next = list2
				list2 = list2.Next
				newListTail = newListTail.Next
			}
			newListTail.Next = nil
		}
	}

	if list1 != nil {
		newListTail.Next = list1
	}

	if list2 != nil {
		newListTail.Next = list2
	}

	return newList
}
