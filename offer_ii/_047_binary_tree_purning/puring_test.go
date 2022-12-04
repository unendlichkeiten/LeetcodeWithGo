package _047_binary_tree_purning

import (
	leetcode "github.com/unendlichkeiten/LeetcodeWithGo"
	"testing"
)

func Test_pruneTree2(t *testing.T) {
	type args struct {
		root *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_01",
			args: args{
				&leetcode.TreeNode{
					Val:  1,
					Left: nil,
					Right: &leetcode.TreeNode{
						Val: 0,
						Left: &leetcode.TreeNode{
							Val:   0,
							Left:  nil,
							Right: nil,
						},
						Right: &leetcode.TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
		},
		{
			name: "test_02",
			args: args{
				&leetcode.TreeNode{
					Val: 1,
					Left: &leetcode.TreeNode{
						Val: 0,
						Left: &leetcode.TreeNode{
							Val:   0,
							Left:  nil,
							Right: nil,
						},
						Right: &leetcode.TreeNode{
							Val:   0,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &leetcode.TreeNode{
						Val: 2,
						Left: &leetcode.TreeNode{
							Val:   0,
							Left:  nil,
							Right: nil,
						},
						Right: &leetcode.TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		got := pruneTree2(tt.args.root)
		t.Logf("got root is %d", got.Val)
	}
}
