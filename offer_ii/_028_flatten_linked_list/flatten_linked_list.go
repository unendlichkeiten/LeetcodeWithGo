package _028_flatten_linked_list

import leetcode "github.com/unendlichkeiten/LeetcodeWithGo"

type nodeStack struct {
	Nodes []*leetcode.DoubleListNode
	Size  int
}

// flattenV1 resolve problem with preOrder
func flatten(root *leetcode.DoubleListNode) *leetcode.DoubleListNode {
	if root == nil {
		return nil
	}

	// init stack
	stack := createStack()

	node := root
	newRoot, newRootTail := (*leetcode.DoubleListNode)(nil), (*leetcode.DoubleListNode)(nil)
	for node != nil || !isEmptyStack(stack) {
		if node != nil {
			tmpNode := new(leetcode.DoubleListNode)
			tmpNode.Val = node.Val
			if newRoot == nil {
				newRoot, newRootTail = tmpNode, tmpNode
			} else {
				newRootTail.Next = tmpNode
				tmpNode.Prev = newRootTail
				newRootTail = tmpNode
			}
			pushStack(stack, node)
			node = node.Child
		} else {
			node = popStack(stack)
			node = node.Next
		}
	}

	return newRoot
}

func isEmptyStack(stack *nodeStack) bool {
	return stack.Size == leetcode.Zero
}

func createStack() *nodeStack {
	return &nodeStack{
		Nodes: make([]*leetcode.DoubleListNode, 0),
		Size:  leetcode.Zero,
	}
}

func pushStack(stack *nodeStack, node *leetcode.DoubleListNode) {
	stack.Nodes = append(stack.Nodes, node)
	stack.Size++
}

func popStack(stack *nodeStack) *leetcode.DoubleListNode {
	node := stack.Nodes[stack.Size-1]
	stack.Size--
	stack.Nodes = stack.Nodes[:stack.Size]

	return node
}
