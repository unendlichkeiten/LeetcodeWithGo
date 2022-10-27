package _021_merge_two_sorted_lists

import (
	"log"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	list1, tmp1 := (*ListNode)(nil), (*ListNode)(nil)
	for i := 1; i < 6; i++ {
		newNode := &ListNode{
			Val:  2 * i,
			Next: nil,
		}

		if list1 == nil {
			list1 = newNode
			tmp1 = newNode
		} else {
			tmp1.Next = newNode
			tmp1 = tmp1.Next
		}
	}

	list2, tmp2 := (*ListNode)(nil), (*ListNode)(nil)
	for i := 0; i < 3; i++ {
		newNode := &ListNode{
			Val:  2*i + 1,
			Next: nil,
		}

		if list2 == nil {
			list2 = newNode
			tmp2 = newNode
		} else {
			tmp2.Next = newNode
			tmp2 = tmp2.Next
		}
	}

	newList := mergeTwoLists(list1, list2)
	log.Printf("after merge list\n")
	for newList != nil {
		log.Printf("value is : %d \n", newList.Val)
		newList = newList.Next
	}
}
