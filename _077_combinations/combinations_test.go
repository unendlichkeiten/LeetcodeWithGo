package _077_combinations

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{"test-1",
			args{4, 2},
			[][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("combine() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
