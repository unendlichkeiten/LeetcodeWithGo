package _334_increasing_array

import "testing"

func Test_increasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test_01", args{[]int{1, 2, 3, 4, 5}}, true},
		{"test_02", args{[]int{5, 4, 3, 2, 1}}, false},
		{"test_03", args{[]int{2, 1, 5, 0, 4, 6}}, true},
		{"test_04", args{[]int{20, 100, 10, 12, 5, 13}}, true},
		{"test_05", args{[]int{2, 4, -2, -3}}, false},
		{"test_06", args{[]int{1, 0, 0, -1, 0, 100}}, true},
		{"test_07", args{[]int{1, 2, 1, 2, 1, 2, 3}}, true},
		{"test_08", args{[]int{1, 2, 1, 2, 1, 2, 3}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
