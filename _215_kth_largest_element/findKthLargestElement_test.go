package _215_kth_largest_element

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{[]int{3, 2, 1, 5, 6, 4}, 2}, 5},
		{"test02", args{[]int{1}, 1}, 1},
		{"test03", args{[]int{2, 1}, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
