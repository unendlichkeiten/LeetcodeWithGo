package leetCode

import "testing"

func TestTwoDimensionalSearch(t *testing.T) {
	type args struct {
		array   [][]int
		rows    int
		columns int
		num     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test01",args{[][]int{{1,2,8,9},{2,4,9,12},{4,7,10,13},{6,8,11,15}},4,4,7},true},
		{"test02",args{[][]int{{1,2,8,9},{2,4,9,12},{4,7,10,13},{6,8,11,15}},4,4,5},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoDimensionalSearch(tt.args.array, tt.args.rows, tt.args.columns, tt.args.num); got != tt.want {
				t.Errorf("TwoDimensionalSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}