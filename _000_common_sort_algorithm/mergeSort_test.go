package _000_common_sort_algorithm

import "testing"

func TestMergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test01", args{[]int{5, 4, 3, 2, 1}}},
		{"test01", args{[]int{1, 4, 2, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.nums)
		})
		for _, num := range tt.args.nums {
			t.Logf("%d\n", num)
		}
	}
}
