package _1441_array_stack_operation

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	type args struct {
		target []int
		n      int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test_01", args{[]int{1, 2, 3}, 3}, []string{"Push", "Push", "Push"}},
		{"test_02", args{[]int{1, 3}, 3}, []string{"Push", "Push", "Pop", "Push"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArray(tt.args.target, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
