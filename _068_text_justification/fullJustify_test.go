package _068_text_justification

import (
	"reflect"
	"testing"
)

func TestFullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test-01",
			args{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16},
			[]string{
				"This    is    an",
				"example  of text",
				"justification.  "},
		},
		{"test-02",
			args{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16},
			[]string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        "},
		},
		{"test-03",
			args{[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to",
				"explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20},
			[]string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
