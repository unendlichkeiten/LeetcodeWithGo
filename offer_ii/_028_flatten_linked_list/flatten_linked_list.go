package _028_flatten_linked_list

import leCode "github.com/unendlichkeiten/LeetcodeWithGo"

type nodeStack struct {
	Nodes []*leCode.DoubleListNode
	Size  int
}

// flattenV1 resolve problem with preOrder
func flattenV1(root *leCode.DoubleListNode) *leCode.DoubleListNode {
	if root == nil {
		return nil
	}

	// init stack
	stack := createStack()

	node := root
	newRoot, newRootTail := (*leCode.DoubleListNode)(nil), (*leCode.DoubleListNode)(nil)
	for node != nil || !isEmptyStack(stack) {
		if node != nil {
			if newRoot == nil {
				newRoot = node
				newRootTail = node //
			} else {
				newRootTail.Next = node
				node.Prev = newRootTail
				newRootTail = node
			}
			pushStack(stack, node)
			node = node.Child
			newRootTail.Child = nil
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
