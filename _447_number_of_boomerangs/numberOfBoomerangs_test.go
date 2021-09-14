package _447_number_of_boomerangs

import "testing"

func TestNumberOfBoomerangs(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test-01", args{[][]int{{0, 0}, {1, 0}, {2, 0}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfBoomerangs(tt.args.points); got != tt.want {
				t.Errorf("NumberOfBoomerangs() = %v, want %v", got, tt.want)
			}
		})
	}
}
