package _047_permutations_ii

import (
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test01",
			args{[]int{2, 1, 1}},
			[][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			}},
		{"test02",
			args{[]int{2, 2, 1, 1}},
			[][]int{
				{1, 1, 2, 2},
				{1, 2, 1, 2},
				{1, 2, 2, 1},
				{2, 1, 1, 2},
				{2, 1, 2, 1},
				{2, 2, 1, 1},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PermuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PermuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
