package leetCode

import "testing"

func TestReplaceBlank(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test01",args{[]byte("hello world")}},
		{"test02",args{[]byte("hello golang world")}},
	}
	for _, tt := range tests {
		ReplaceBlank(tt.args.src)
		t.Logf("%s", tt.args.src)
	}
}