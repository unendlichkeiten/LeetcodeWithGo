package _028_sub_string_index

import (
	"reflect"
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{"hello world", "o wor"}, 4},
		{"test02", args{"leetcode", "leeto"}, -1},
		{"test03", args{"mississippi", "issi"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStrV2(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNext(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test_01", args{"abcac"}, []int{-1, 0, 0, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNext(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
