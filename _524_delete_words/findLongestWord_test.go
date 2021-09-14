package _524_delete_words

import "testing"

func TestFindLongestWord(t *testing.T) {
	type args struct {
		s          string
		dictionary []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test-01", args{"abpcplea", []string{"ale", "apple", "monkey", "plea"}}, "apple"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLongestWord(tt.args.s, tt.args.dictionary); got != tt.want {
				t.Errorf("FindLongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
