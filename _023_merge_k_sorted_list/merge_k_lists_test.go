package _023_merge_k_sorted_list

import (
	"log"
	"testing"

	leetCode "github.com/unendlichkeiten/LeetcodeWithGo"
)

func Test_mergeKLists(t *testing.T) {
	lists := make([]*leetCode.ListNode, 3)
	for index := 0; index < 3; index++ {
		list, tmp := (*leetCode.ListNode)(nil), (*leetCode.ListNode)(nil)
		for i := 1; i < index+2; i++ {
			newNode := &leetCode.ListNode{
				Val:  2*i + index,
				Next: nil,
			}

			if list == nil {
				list = newNode
				tmp = newNode
			} else {
				tmp.Next = newNode
				tmp = tmp.Next
			}
		}
		lists[index] = list
	}

	newList := mergeKLists(lists)
	log.Printf("after merge k lists\n")
	for newList != nil {
		log.Printf("value is : %d \n", newList.Val)
		newList = newList.Next
	}
}
