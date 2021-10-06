package _414_third_maximum_number

import "testing"

func Test_thirdMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test-01", args{[]int{2, 2, 3, 1}}, 1},
		{"test-02", args{[]int{1, 1, 2, 3, 4, 5, 6, 7}}, 5},
		{"test-03", args{[]int{1, 2, 2, 5, 3, 5}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
