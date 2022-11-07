package _066_rotate_list

import (
	"log"
	"testing"

	leetCode "github.com/unendlichkeiten/LeetcodeWithGo"
)

func Test_rotateRight(t *testing.T) {
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

	log.Printf("before rotate linked list\n")
	tmp = sample
	for tmp != nil {
		log.Printf("value is : %d \n", tmp.Val)
		tmp = tmp.Next
	}

	newHead := rotateRight(sample, 7)
	log.Printf("after rotate linked list\n")
	for newHead != nil {
		log.Printf("value is : %d \n", newHead.Val)
		newHead = newHead.Next
	}
}
