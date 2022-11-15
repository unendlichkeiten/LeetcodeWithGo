package _035_copy_random_list

import (
	"fmt"

	leetCode "github.com/unendlichkeiten/LeetcodeWithGo"
)

var (
	_nodeMap  = make(map[string]*leetCode.RandomNode)
	_nodeMap2 = make(map[*leetCode.RandomNode]*leetCode.RandomNode)
)

func copyRandomList2(head *leetCode.RandomNode) *leetCode.RandomNode {
	if head == nil {
		return nil
	}

	for node := head; node != nil; node = node.Next {
		_nodeMap2[node] = &leetCode.RandomNode{Val: node.Val}
	}

	for node := head; node != nil; node = node.Next {
		_nodeMap2[node].Next = _nodeMap2[node.Next]
		_nodeMap2[node].Random = _nodeMap2[node.Random]
	}

	return _nodeMap2[head]
}

// 采用递归函数 + hashMap
func copyRandomList(head *leetCode.RandomNode) *leetCode.RandomNode {

	if head == nil {
		return nil
	}

	key := fmt.Sprintf("%p", head)
	if n, ok := _nodeMap[key]; ok {
		return n
	}

	newNode := new(leetCode.RandomNode)
	newNode.Val = head.Val
	_nodeMap[key] = newNode
	newNode.Next = copyRandomList(head.Next)
	newNode.Random = copyRandomList(head.Random)

	return newNode
}
