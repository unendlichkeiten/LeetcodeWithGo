package _297_serialize_binary_tree

import (
	leetcode "github.com/unendlichkeiten/LeetcodeWithGo"
	"testing"
)

func TestCodec_deserialize(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test_01", args{"1,2,3,null,null,4,5"}},
		{"test_02", args{"1,2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Codec{}
			got := c.deserialize(tt.args.data)
			t.Logf("%#v", got)
		})
	}
}

func TestCodec_serialize(t *testing.T) {
	type args struct {
		root *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test_01",
			args{
				&leetcode.TreeNode{
					Val: 1,
					Left: &leetcode.TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &leetcode.TreeNode{
						Val: 3,
						Left: &leetcode.TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
						Right: &leetcode.TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
					},
				}}, "1,2,3,null,null,4,5",
		},
		{"test_02", args{(*leetcode.TreeNode)(nil)}, ""},
		{
			"test_01",
			args{
				&leetcode.TreeNode{
					Val: 1,
					Left: &leetcode.TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				}}, "1,2,null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Codec{}
			if got := c.serialize(tt.args.root); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
