package _1668_maximum_repeating_substring

import "testing"

func Test_maxRepeating(t *testing.T) {
	type args struct {
		sequence string
		word     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test_01", args{"ababc", "ab"}, 2},
		{"test_02", args{"ababc", "ba"}, 1},
		{"test_03", args{"ababc", "ac"}, 0},
		{
			name: "test_04",
			args: args{"aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRepeating(tt.args.sequence, tt.args.word); got != tt.want {
				t.Errorf("maxRepeating() = %v, want %v", got, tt.want)
			}
		})
	}
}
