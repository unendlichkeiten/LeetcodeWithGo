package _029_ordered_cycle_linked_list

import (
	leetCode "github.com/unendlichkeiten/LeetcodeWithGo"
	"log"
	"testing"
)

func Test_insert(t *testing.T) {
	sample := &leetCode.ListNode{
		Val: 1,
		Next: &leetCode.ListNode{
			Val: 3,
			Next: &leetCode.ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	sample.Next.Next.Next = sample
	log.Printf("before insert linked list\n")
	for tmp := (*leetCode.ListNode)(nil); tmp != sample; {
		if tmp == nil {
			tmp = sample
		}
		log.Printf("value is : %d \n", tmp.Val)
		tmp = tmp.Next
	}

	log.Printf("after insert linked list\n")
	newNode := insert(sample, 1)
	for tmp := (*leetCode.ListNode)(nil); tmp != newNode; {
		if tmp == nil {
			tmp = newNode
		}
		log.Printf("value is : %d \n", newNode.Val)
		newNode = newNode.Next
	}
}
