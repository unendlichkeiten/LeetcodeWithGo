package _069_sqrtx

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{0}, 0},
		{"test01", args{1}, 1},
		{"test01", args{2}, 1},
		{"test01", args{3}, 1},
		{"test01", args{4}, 2},
		{"test01", args{8}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
