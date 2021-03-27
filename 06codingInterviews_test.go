package leetCode

import (
	"testing"
)

func createList(l LinkList, array []int) {
	for _, value := range array {
		l.next = &LNode{
			data : value,
			next : nil,
		}

		l = l.next
	}
}
func Test_travelTail2Head(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{1, 2, 3, 4, 5}}},
		{"test02", args{[]int{}}},
	}

	list := &LNode{}
	for _, tt := range tests {
		createList(list, tt.args.array)
		t.Log(tt.name)
		travelTail2Head(list.next)
		list = &LNode{}
	}
}

func Test_travelTail2HeadByReverse(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{1, 2, 3, 4, 5}}},
		{"test02", args{[]int{}}},
	}

	list := &LNode{}
	for _, tt := range tests {
		createList(list, tt.args.array)
		t.Log(tt.name)
		travelTail2HeadByReverse(list.next)
		list = &LNode{}
	}
}