package _382_linked_list_random_node

import (
	"math/rand"

	leetcode "github.com/unendlichkeiten/LeetcodeWithGo"
)

// 方法一：
//type Solution struct {
//	values []int
//}
//
//
//func Constructor(head *ListNode) Solution {
//	s := Solution{make([]int, 0)}
//	for head != nil {
//		s.values = append(s.values, head.Val)
//		head = head.Next
//	}
//	return s
//}
//
//
//func (this *Solution) GetRandom() int {
//	return this.values[rand.Intn(len(this.values))]
//}

// 方法二

type Solution struct {
	head *leetcode.ListNode
}

func Constructor(head *leetcode.ListNode) Solution {
	return Solution{head}
}

func (s *Solution) GetRandom() (ans int) {
	for node, i := s.head, 1; node != nil; node = node.Next {
		if rand.Intn(i) == 0 { // 1/i 的概率选中（替换为答案）
			ans = node.Val
		}
		i++
	}
	return
}
