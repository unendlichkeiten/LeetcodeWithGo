package _166_fraction_to_recurring_decimal

import "testing"

func Test_fractionToDecimal(t *testing.T) {
	type args struct {
		numerator   int
		denominator int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test-01", args{1, 2}, "0.5"},
		{"test-02", args{2, 1}, "2"},
		{"test-03", args{2, 3}, "0.(6)"},
		{"test-04", args{4, 333}, "0.(012)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractionToDecimal(tt.args.numerator, tt.args.denominator); got != tt.want {
				t.Errorf("fractionToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
