package _007_build_binary_tree

import leetcode "github.com/unendlichkeiten/LeetcodeWithGo"

func buildTree(preorder []int, inorder []int) *leetcode.TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	inOrderRootPos, root := 0, &leetcode.TreeNode{Val: preorder[0]}
	for i, v := range inorder {
		if v == root.Val {
			inOrderRootPos = i
		}
	}

	root.Left = buildTree(preorder[1:inOrderRootPos+1], inorder[:inOrderRootPos])
	root.Right = buildTree(preorder[inOrderRootPos+1:], inorder[inOrderRootPos+1:])

	return root
}
