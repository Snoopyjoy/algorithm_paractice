package leetcode

import "testing"

func TestReversePairs(t *testing.T) {
	ret := reversePairs([]int{1, 3, 2, 3, 1})
	if ret != 2 {
		t.Fatalf("expected answer 2, got : %d", ret)
	}
}
