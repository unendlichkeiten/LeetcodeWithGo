package _763_partition_labels

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name          string
		args          args
		wantPartition []int
	}{
		{"test01", args{"ababcbacadefegdehijhklij"}, []int{9, 7, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPartition := partitionLabels(tt.args.s); !reflect.DeepEqual(gotPartition, tt.wantPartition) {
				t.Errorf("partitionLabels() = %v, want %v", gotPartition, tt.wantPartition)
			}
		})
	}
}
