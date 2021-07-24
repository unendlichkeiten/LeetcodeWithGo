package _004_median_of_two_sorted_array

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test-01",
			args{
				nums1: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 4},
				nums2: []int{1, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4},
			},
			3.0000,
		},
		{"test-02",
			args{
				nums1: []int{1, 1},
				nums2: []int{1, 2},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
