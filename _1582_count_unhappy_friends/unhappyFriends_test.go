package _1582_count_unhappy_friends

import "testing"

func Test_unhappyFriends(t *testing.T) {
	type args struct {
		n           int
		preferences [][]int
		pairs       [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-01",
			args{
				n:           4,
				preferences: [][]int{{1, 3, 2}, {2, 3, 0}, {1, 0, 3}, {1, 0, 2}},
				pairs:       [][]int{{2, 1}, {3, 0}},
			},
			0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unhappyFriends(tt.args.n, tt.args.preferences, tt.args.pairs); got != tt.want {
				t.Errorf("unhappyFriends() = %v, want %v", got, tt.want)
			}
		})
	}
}
