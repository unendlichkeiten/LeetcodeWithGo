package _2196_create_binary_tree

import (
	"fmt"
	leetcode "github.com/unendlichkeiten/LeetcodeWithGo"
	"testing"
)

func Test_createBinaryTree(t *testing.T) {
	type args struct {
		descriptions [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_01",
			args: args{
				descriptions: [][]int{
					{20, 15, 1},
					{20, 17, 0},
					{50, 20, 1},
					{50, 80, 0},
					{80, 19, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		root := createBinaryTree(tt.args.descriptions)
		// 使用先序遍历遍历二叉树
		preOrder(root)
	}
}

// preOrder 先序遍历之递归实现
func preOrder(p *leetcode.TreeNode) {
	if p != nil {
		fmt.Printf("node value is : %d\n", p.Val)
		preOrder(p.Left)
		preOrder(p.Right)
	}
}
