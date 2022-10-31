package _026_reorder_list

import (
	"log"
	"testing"

	leetCode "github.com/unendlichkeiten/LeetcodeWithGo"
)

func Test_reorderList(t *testing.T) {
	sample, tmp := (*leetCode.ListNode)(nil), (*leetCode.ListNode)(nil)
	for i := 1; i < 5; i++ {
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

	before, after := sample, sample
	log.Printf("before reorder lists\n")
	for before != nil {
		log.Printf("value is : %d \n", before.Val)
		before = before.Next
	}

	reorderList(sample)

	log.Printf("after reorder lists\n")
	for after != nil {
		log.Printf("value is : %d \n", after.Val)
		after = after.Next
	}
}
