package _055_binary_tree_depth

import leetcode "github.com/unendlichkeiten/LeetcodeWithGo"

func maxDepth(root *leetcode.TreeNode) int {
	if root == nil {
		return 0
	}

	q, depth, tmpQueue := leetcode.CreateQueue(), 0, leetcode.CreateQueue()
	leetcode.PushQueue(&q, root)

	for !leetcode.IsEmptyQueue(q) {
		if tmp := leetcode.PopQueue(&q); tmp != nil {
			leetcode.PushQueue(&tmpQueue, tmp.Left)
			leetcode.PushQueue(&tmpQueue, tmp.Right)
		}

		if leetcode.IsEmptyQueue(q) {
			depth++
			q = tmpQueue
			tmpQueue = leetcode.CreateQueue()
		}
	}

	return depth - 1
}
