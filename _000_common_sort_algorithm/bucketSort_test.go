package _000_common_sort_algorithm

import "testing"

func TestBucketSort(t *testing.T) {
	type args struct {
		theArray []int
		num      int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test01", args{[]int{10, 1, 18, 30, 23, 12, 7, 5, 18, 17}, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BucketSort(tt.args.theArray, tt.args.num)
			for _, num := range tt.args.theArray {
				t.Logf("%d ", num)
			}
			t.Logf("\n")
		})
	}
}
