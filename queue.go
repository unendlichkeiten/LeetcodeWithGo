package leetCode

func CreateQueue() []*TreeNode {
	return make([]*TreeNode, 0)
}

func IsEmptyQueue(queue []*TreeNode) bool {
	return len(queue) == 0
}

func PushQueue(queue *[]*TreeNode, node *TreeNode) {
	*queue = append(*queue, node)
	return
}

func PopQueue(queue *[]*TreeNode) *TreeNode {
	node := (*queue)[0]
	*queue = (*queue)[1:]
	return node
}
