package _206_reverse_linked_list

import (
	"log"
	"testing"

	leetCode "github.com/unendlichkeiten/LeetcodeWithGo"
)

func Test_reverseList(t *testing.T) {
	sample, tmp := (*leetCode.ListNode)(nil), (*leetCode.ListNode)(nil)
	for i := 1; i < 6; i++ {
		newNode := &leetCode.ListNode{
			Val:  i,
			Next: nil,
		}

		if sample == nil {
			sample = newNode
			tmp = newNode
		} else {
			tmp.Next = newNode
			tmp = tmp.Next
		}
	}

	log.Printf("before reverse linked list\n")
	tmp = sample
	for tmp != nil {
		log.Printf("value is : %d \n", tmp.Val)
		tmp = tmp.Next
	}

	log.Printf("after reverse linked list\n")
	tmp = reverseListV2(sample)
	for tmp != nil {
		log.Printf("value is : %d \n", tmp.Val)
		tmp = tmp.Next
	}
}
