package leetCode

import "testing"

func TestReplicaValue_Hash(t *testing.T) {
	testArray := []int{2, 3, 1, 0, 2, 5, 3}

	if value, err := ReplicaValue_Hash(testArray); err == nil {
		t.Logf("exists replica value %d", value)
	}
}

func TestReplicaValue_Scan(t *testing.T) {
	testArray := []int{2, 3, 1, 0, 2, 5, 3}

	if value, err := ReplicaValue_Scan(testArray); err == nil {
		t.Logf("exists replica value %d", value)
	}
}