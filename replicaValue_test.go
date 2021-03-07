package leetCode

import "testing"

func TestReplicaValue_Hash(t *testing.T) {
	testArray := []int{0, 3, 1, 0, 2, 5, 3}

	if value, err := ReplicaValue_Hash(testArray); err == nil {
		t.Logf("exists replica value %d", value)
	}
}

func TestReplicaValue_Scan(t *testing.T) {
	testArray := []int{1, 3, 1, 0, 2, 5, 3}

	if value, err := ReplicaValue_Scan(testArray); err == nil {
		t.Logf("exists replica value %d", value)
	}
}

func TestQuickSort2(t *testing.T) {
	type args struct {
		num   []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test01",args{[]int{3,6,2,3,2,5,1,9},0,7}},
		{"test02",args{[]int{1,2,3,4,5,3,2,1},0,7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort2(tt.args.num, tt.args.left, tt.args.right)
			t.Log(tt.args.num)
		})
	}
}

func TestRepicaValue_Mem(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		// TODO: Add test cases.
		{"test01",args{[]int{3,6,2,3,2,5,1,9}}, 2, nil},
		{"test02",args{[]int{1,2,3,4,5,3,2,1}}, 1, nil},
	}
	for _, tt := range tests {
		got, err := RepicaValue_Mem(tt.args.array)
		if got != tt.want || err != nil {
			t.Errorf("RepicaValue_Mem() got = %v, want %v", got, tt.want)
		}
	}
}

func TestReplicaValue_Div(t *testing.T) {
	type args struct {
		array []int
		left  int
		right int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		// TODO: Add test cases.
		{"test01",args{[]int{0,3,6,2,3,2,5,1,9}, 1,8}, 2, nil},
		{"test02",args{[]int{0,1,2,3,4,5,3,2,1}, 1,8}, 1, nil},
		{"test03",args{[]int{0,1,1,1,2,2,2,2,2}, 1,8}, 1, nil},
	}
	for _, tt := range tests {
			got, err := ReplicaValue_Div(tt.args.array, tt.args.left, tt.args.right)
			if got != tt.want || err != nil {
				t.Errorf("ReplicaValue_Div() got = %v, want %v", got, tt.want)
			}
	}
}