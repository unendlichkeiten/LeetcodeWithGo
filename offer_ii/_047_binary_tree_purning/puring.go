package _047_binary_tree_purning

import leetcode "github.com/unendlichkeiten/LeetcodeWithGo"

type nodeStack struct {
	Nodes []*leetcode.TreeNode
	Size  int
}

// pruneTree 递归解法
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

// pruneTree2 非递归解法：栈内的相邻元素一定是父子关系
func pruneTree2(root *leetcode.TreeNode) *leetcode.TreeNode {
	pre, p, s := (*leetcode.TreeNode)(nil), root, createStack()

	for p != nil || !isEmptyStack(s) {
		if p != nil {
			pushStack(s, p)
			p = p.Left
		} else {
			p = getTopNode(s)
			if p.Right != nil && p.Right != pre {
				p = p.Right
				pushStack(s, p)
				p = p.Left
			} else {
				p = popStack(s)
				// 减枝处理
				if p.Val == 0 && p.Left == nil && p.Right == nil {
					if isEmptyStack(s) {
						return nil
					} else {
						parent := getTopNode(s) // critical option
						if parent.Left == p {
							parent.Left = nil
						} else if parent.Right == p {
							parent.Right = nil
						}
					}
				}
				pre = p
				p = nil
			}
		}
	}

	return root
}

func isEmptyStack(stack *nodeStack) bool {
	return stack.Size == leetcode.Zero
}

func createStack() *nodeStack {
	return &nodeStack{
		Nodes: make([]*leetcode.TreeNode, 0),
		Size:  leetcode.Zero,
	}
}

func getTopNode(stack *nodeStack) *leetcode.TreeNode {
	return stack.Nodes[stack.Size-1]
}

func pushStack(stack *nodeStack, node *leetcode.TreeNode) {
	stack.Nodes = append(stack.Nodes, node)
	stack.Size++
}

func popStack(stack *nodeStack) *leetcode.TreeNode {
	node := stack.Nodes[stack.Size-1]
	stack.Size--
	stack.Nodes = stack.Nodes[:stack.Size]

	return node
}
