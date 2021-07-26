package _008_string_to_integer

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test-01", args{"42"}, 42},
		{"test-02", args{"-42"}, -42},
		{"test-03", args{"4193 with words"}, 4193},
		{"test-04", args{""}, 0},
		{"test-05", args{"00000-42a1234"}, 0},
		{"test-06", args{"  0000000000012345678"}, 12345678},
		{"test-07", args{"21474836460"}, 2147483647},
		{"test-08", args{"9223372036854775808"}, 2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
