package _028_flatten_linked_list

import leCode "github.com/unendlichkeiten/LeetcodeWithGo"

type nodeStack struct {
	Nodes []*leCode.DoubleListNode
	Size  int
}

// flattenV1 resolve problem with preOrder
func flatten(root *leCode.DoubleListNode) *leCode.DoubleListNode {
	if root == nil {
		return nil
	}

	// init stack
	stack := createStack()

	node := root
	newRoot, newRootTail := (*leCode.DoubleListNode)(nil), (*leCode.DoubleListNode)(nil)
	for node != nil || !isEmptyStack(stack) {
		if node != nil {
			tmpNode := new(leCode.DoubleListNode)
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
	return stack.Size == leCode.Zero
}

func createStack() *nodeStack {
	return &nodeStack{
		Nodes: make([]*leCode.DoubleListNode, 0),
		Size:  leCode.Zero,
	}
}

func pushStack(stack *nodeStack, node *leCode.DoubleListNode) {
	stack.Nodes = append(stack.Nodes, node)
	stack.Size++
}

func popStack(stack *nodeStack) *leCode.DoubleListNode {
	node := stack.Nodes[stack.Size-1]
	stack.Size--
	stack.Nodes = stack.Nodes[:stack.Size]

	return node
}
