package _077_list_sorting

import (
	"log"
	"math/rand"
	"testing"

	leetCode "github.com/unendlichkeiten/LeetcodeWithGo"
)

func TestSortList(t *testing.T) {
	list1, tmp1 := (*leetCode.ListNode)(nil), (*leetCode.ListNode)(nil)
	for i := 1; i < 6; i++ {
		newNode := &leetCode.ListNode{
			Val:  rand.Intn(10),
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

	sortedList := SortList(list1)
	log.Printf("after sorting lists\n")
	for sortedList != nil {
		log.Printf("value is : %d \n", sortedList.Val)
		sortedList = sortedList.Next
	}
}
