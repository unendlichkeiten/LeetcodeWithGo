package _934_shortet_bridage

import "testing"

func Test_shortestBridge(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test-01", args{[][]int{{0, 1}, {1, 0}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestBridge(tt.args.grid); got != tt.want {
				t.Errorf("shortestBridge() = %v, want %v", got, tt.want)
			}
		})
	}
}
