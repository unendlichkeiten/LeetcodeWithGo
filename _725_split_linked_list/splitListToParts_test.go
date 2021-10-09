package _725_split_linked_list

import (
	"testing"
)

func Test_splitListToParts(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-01",
			args: args{
				head: tailInsert(3),
				k:    5,
			},
		},
	}
	for _, tt := range tests {
		splitListToParts(tt.args.head, tt.args.k)
	}
}

func tailInsert(k int) *ListNode {
	head, tail := (*ListNode)(nil), (*ListNode)(nil)

	for i := 1; i <= k; i++ {
		if head == nil {
			head = &ListNode{
				Val:  i,
				Next: nil,
			}

			tail = head
		} else {
			tail.Next = &ListNode{
				Val:  i,
				Next: nil,
			}

			tail = tail.Next
		}
	}

	return head
}
