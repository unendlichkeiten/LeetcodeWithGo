package _000_common_sort_algorithm

import (
	"reflect"
	"testing"
)

func Test_shellSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test-01", args{[]int{0, 3, 5, 2, 4, 9, 7, 6}}, []int{0, 2, 3, 4, 5, 6, 7, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shellSort2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shellSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
