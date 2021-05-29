package _167_two_sum

import (
	"reflect"
	"testing"
)

func Test_twoSumByMap(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test01", args{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
		{"test02", args{[]int{2, 3, 4, 6}, 6}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumByMap(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumByMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSumByIndex(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test01", args{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
		{"test02", args{[]int{2, 3, 4, 6}, 6}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumByIndex(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
