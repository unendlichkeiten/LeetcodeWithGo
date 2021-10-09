package _725_split_linked_list

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	ansList, tailList := make([]*ListNode, k), make([]*ListNode, k)

	node, index := head, 0
	for node != nil {
		if ansList[index%k] == nil {
			ansList[index%k] = node
			tailList[index%k] = ansList[index%k]
			tailList[index%k].Next = node
		} else {
			tailList[index%k].Next = node
			tailList[index%k] = tailList[index%k].Next
		}
		node = node.Next
	}

	return ansList
}
