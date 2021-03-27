package leetCode

type Nodes struct {
	Data int
	Next *Nodes
}

type Queue struct {
	Front *Nodes
	Rear *Nodes
}

// 链表定义
type LinkList *LNode
type LNode struct{
	data 	int
	next  *LNode
}

// 二叉树节点定义
type BiTree *BiTNode
type BiTNode struct {
	data int
	lchild *BiTNode
	rchild *BiTNode
}