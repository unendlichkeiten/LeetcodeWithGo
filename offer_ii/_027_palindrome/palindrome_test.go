package _027_palindrome

import (
	leetCode "github.com/unendlichkeiten/LeetcodeWithGo"
	"log"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	//sample := &leetCode.ListNode{
	//	Val: 1,
	//	Next: &leetCode.ListNode{
	//		Val: 3,
	//		Next: &leetCode.ListNode{
	//			Val: 3,
	//			Next: &leetCode.ListNode{
	//				Val:  1,
	//				Next: nil,
	//			},
	//		},
	//	},
	//}

	sample := &leetCode.ListNode{
		Val: 1,
		Next: &leetCode.ListNode{
			Val: 3,
			Next: &leetCode.ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	log.Printf("sample is palindrome linked list ? : %v", isPalindrome(sample))
}
