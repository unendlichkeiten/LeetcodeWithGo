package _055_binary_tree_depth

import (
	leetcode "github.com/unendlichkeiten/LeetcodeWithGo"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_01",
			args{
				&leetcode.TreeNode{
					Val: 3,
					Left: &leetcode.TreeNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &leetcode.TreeNode{
						Val: 20,
						Left: &leetcode.TreeNode{
							Val:   15,
							Left:  nil,
							Right: nil,
						},
						Right: &leetcode.TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				}}, 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
