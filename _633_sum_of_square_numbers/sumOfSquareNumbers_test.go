package _633_sum_of_square_numbers

import "testing"

func Test_judgeSquareSum(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test01", args{1}, true},
		{"test02", args{2}, true},
		{"test03", args{3}, false},
		{"test04", args{4}, true},
		{"test05", args{5}, true},
		{"test06", args{1000000000}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeSquareSum(tt.args.c); got != tt.want {
				t.Errorf("judgeSquareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
