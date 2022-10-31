package _077_list_sorting

import (
	leetCode "github.com/unendlichkeiten/LeetcodeWithGo"
)

func SortList(head *leetCode.ListNode) *leetCode.ListNode {

	if head == nil {
		return head
	}

	nodeArr := make([]*leetCode.ListNode, 0)
	for head != nil {
		nodeArr = append(nodeArr, head)
		head = head.Next
	}

	// 使用快排
	QuickSort(nodeArr)

	newList, node := nodeArr[0], nodeArr[0]
	for i := 1; i < len(nodeArr); i++ {
		nodeArr[i].Next = nil
		node.Next = nodeArr[i]
		node = node.Next
	}

	return newList
}

func QuickSort(nodeArr []*leetCode.ListNode) {
	if len(nodeArr) <= 1 {
		return
	}

	quickSort(nodeArr, 0, len(nodeArr)-1)
	return
}

func quickSort(nodeArr []*leetCode.ListNode, l, r int) {
	pivotPos := 0
	if l < r {
		pivotPos = getPartition(nodeArr, l, r)
		quickSort(nodeArr, l, pivotPos-1)
		quickSort(nodeArr, pivotPos+1, r)
	}
}

func getPartition(nums []*leetCode.ListNode, l, r int) int {
	pivot := nums[l]
	for l < r {
		for l < r && pivot.Val <= nums[r].Val {
			r--
		}
		nums[l] = nums[r]

		for l < r && nums[l].Val <= pivot.Val {
			l++
		}
		nums[r] = nums[l]
	}

	nums[l] = pivot

	return l
}
