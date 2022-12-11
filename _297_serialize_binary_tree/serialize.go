package _297_serialize_binary_tree

import (
	"bytes"
	"strconv"
	"strings"

	leetcode "github.com/unendlichkeiten/LeetcodeWithGo"
)

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *leetcode.TreeNode) string {
	q, p, res := createQueue(), root, bytes.Buffer{}
	if root == nil {
		return ""
	}

	pushQueue(&q, p)
	queueLen, nullPos := len(q), 0
	for !isEmptyQueue(q) {
		queueLen, nullPos = len(q), 0
		for i := 0; i < queueLen; i++ {
			node := popQueue(&q)
			if node == nil {
				if nullPos == 0 {
					nullPos = res.Len()
				}
				res.WriteString("null")
			} else {
				res.WriteString(strconv.Itoa(node.Val))
				pushQueue(&q, node.Left)
				pushQueue(&q, node.Right)
			}
			res.WriteString(",")
		}
	}

	return res.String()[:nullPos-1]
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *leetcode.TreeNode {
	values, root, q := strings.Split(data, ","), new(leetcode.TreeNode), createQueue()
	if values[0] == "" {
		return nil
	}

	root.Val, _ = strconv.Atoi(values[0])
	pushQueue(&q, root)
	for i := 1; !isEmptyQueue(q) && i < len(values); i += 2 {
		node := popQueue(&q)

		// left
		if values[i] == "null" {
			node.Left = nil
		} else {
			value, _ := strconv.Atoi(values[i])
			node.Left = &leetcode.TreeNode{Val: value}
		}
		pushQueue(&q, node.Left)

		// right
		if i == len(values)-1 || values[i+1] == "null" {
			node.Right = nil
		} else {
			value, _ := strconv.Atoi(values[i+1])
			node.Right = &leetcode.TreeNode{Val: value}
		}
		pushQueue(&q, node.Right)
	}

	return root
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
