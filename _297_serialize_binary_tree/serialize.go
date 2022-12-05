package _297_serialize_binary_tree

import (
	"fmt"
	leetcode "github.com/unendlichkeiten/LeetcodeWithGo"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *leetcode.TreeNode) string {
	q, p, res := createQueue(), root, ""
	if root == nil {
		return ""
	}

	pushQueue(&q, p)
	for !isEmptyQueue(q) {
		node := popQueue(&q)
		res = fmt.Sprintf("%s%d", res, node.Val)
		pushQueue(&q, node.Left)
		pushQueue(&q, node.Right)
	}

	return res
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *leetcode.TreeNode {
	return nil
}

func createQueue() []*leetcode.TreeNode {
	return make([]*leetcode.TreeNode, 0, 10000)
}

func isEmptyQueue(queue []*leetcode.TreeNode) bool {
	return len(queue) == 0
}

func pushQueue(queue *[]*leetcode.TreeNode, node *leetcode.TreeNode) {
	*queue = append(*queue, node)
	return
}

func popQueue(queue *[]*leetcode.TreeNode) *leetcode.TreeNode {
	node := (*queue)[0]
	*queue = (*queue)[1:]
	return node
}
