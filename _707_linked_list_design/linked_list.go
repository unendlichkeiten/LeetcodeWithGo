package _707_linked_list_design

import leetCode "github.com/unendlichkeiten/LeetcodeWithGo"

type MyLinkedList struct {
	Head *leetCode.ListNode // head 是一个哨兵节点，有值节点从 Head.Next 开始
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Head: new(leetCode.ListNode),
		Size: leetCode.Zero,
	}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.Size {
		return -1
	}

	tmp := l.Head
	for i := 0; i <= index; i++ {
		tmp = tmp.Next
	}

	return tmp.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.Size, val)
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.Size {
		return
	}

	if index < 0 {
		index = 0
	}

	pre := l.Head
	for i := 0; i < index; i++ {
		pre = pre.Next
	}

	newNode := &leetCode.ListNode{
		Val:  val,
		Next: pre.Next,
	}

	pre.Next = newNode
	l.Size++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.Size {
		return
	}

	pre := l.Head
	for i := 0; i < index; i++ {
		pre = pre.Next
	}

	pre.Next = pre.Next.Next
	l.Size--
}
