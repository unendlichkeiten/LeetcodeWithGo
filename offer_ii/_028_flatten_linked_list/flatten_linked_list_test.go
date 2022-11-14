package _028_flatten_linked_list

import (
	leCode "github.com/unendlichkeiten/LeetcodeWithGo"
	"testing"
)

func Test_flattenV1(t *testing.T) {
	sample := leCode.DoubleListNode{
		Val:  1,
		Prev: nil,
		Next: &leCode.DoubleListNode{
			Val:  2,
			Prev: nil,
		},
		Child: &leCode.DoubleListNode{
			Val:   3,
			Prev:  nil,
			Child: nil,
		},
	}

	sample.Next.Prev = &sample

	flattenV1(&sample)
}
