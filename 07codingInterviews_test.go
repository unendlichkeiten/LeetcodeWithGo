package leetCode

import (
	"testing"
)

func TestCreateBiTreeByPreAndInOrder(t *testing.T) {
	type args struct {
		preOrder []int
		prel     int
		prer     int
		inOrder  []int
		inl      int
		inr      int
	}
	tests := []struct {
		name    string
		args    args
	}{
		// TODO: Add test cases.
		{"first",
			args{
			preOrder: []int{1, 2, 4, 7, 3, 5, 6, 8},
			prel: 0,
			prer: 7,
			inOrder: []int{4, 7, 2, 1, 5, 3, 8, 6},
			inl: 0,
			inr: 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biTree, err := CreateBiTreeByPreAndInOrder(tt.args.preOrder, tt.args.prel, tt.args.prer,
																									tt.args.inOrder, tt.args.inl, tt.args.inr)
			if err != nil {
				t.Log(err)
			} else {
				t.Log(biTree)
			}
		})
	}
}