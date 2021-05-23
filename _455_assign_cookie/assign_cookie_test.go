package _455_assign_cookie

import "testing"

func TestFindContentChildren(t *testing.T) {
	type args struct {
		g []int
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{1, 2, 3}, []int{1, 1}}, 1},
		{"test02", args{[]int{1, 2, 3}, []int{1, 2}}, 2},
		{"test03", args{[]int{10, 9, 8, 7}, []int{5, 6, 7, 8}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContentChildren(tt.args.g, tt.args.s); got != tt.want {
				t.Errorf("findContentChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}
