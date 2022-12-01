package _2196_create_binary_tree

import leetcode "github.com/unendlichkeiten/LeetcodeWithGo"

func createBinaryTree(descriptions [][]int) *leetcode.TreeNode {

	// 将二维数组转换为二叉树节点 []int{parent, child, isLeft}
	nodeMap, root := make(map[int]*leetcode.TreeNode), new(leetcode.TreeNode)
	for _, v := range descriptions {
		node := &leetcode.TreeNode{
			Val: v[1],
		}

		nodeMap[v[1]] = node
	}

	// 找到根节点
	for _, v := range descriptions {
		if _, ok := nodeMap[v[0]]; !ok {
			root.Val = v[0]
			nodeMap[v[0]] = root
		}
	}

	// 连接各个节点
	for _, v := range descriptions {
		// 判断左右孩子 []int{parent, child, isLeft}
		if v[2] == 1 {
			// 左孩子
			nodeMap[v[0]].Left = nodeMap[v[1]]
		} else {
			// 右孩子
			nodeMap[v[0]].Right = nodeMap[v[1]]
		}
	}

	return root
}
