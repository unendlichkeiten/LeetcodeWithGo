package _678_valid_parenthesis_string

import "testing"

func TestCheckValidString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test-01", args{"()"}, true},
		{"test-02", args{"(*)"}, true},
		{"test-03", args{"(*))"}, true},
		{"test-04", args{"*(*))"}, true},
		{"test-05", args{"*((*))"}, true},
		{"test-06", args{"*)((*))"}, true},
		{"test-07", args{"))((*))"}, false},
		{"test-08", args{"**(()"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckValidString(tt.args.s); got != tt.want {
				t.Errorf("CheckValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
