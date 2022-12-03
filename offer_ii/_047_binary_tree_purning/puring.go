package _047_binary_tree_purning

import leetcode "github.com/unendlichkeiten/LeetcodeWithGo"

func pruneTree(root *leetcode.TreeNode) *leetcode.TreeNode {
	if root == nil {
		return root
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}

	return root
}
