package _025_reverse_nodes_k

import (
	"log"
	"testing"

	leetCode "github.com/unendlichkeiten/LeetcodeWithGo"
)

func Test_swapPairs(t *testing.T) {
	sample, tmp := (*leetCode.ListNode)(nil), (*leetCode.ListNode)(nil)
	for i := 1; i < 7; i++ {
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
	tmp = swapPairs(sample, 3)
	for tmp != nil {
		log.Printf("value is : %d \n", tmp.Val)
		tmp = tmp.Next
	}
}
