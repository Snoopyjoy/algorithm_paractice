package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	r := myAtoi("  0000000000012345678")
	if r != 12345678 {
		t.Fatalf("err ret: %d", r)
	}
}
