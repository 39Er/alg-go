package leetcode

import (
	"testing"
)

func TestSumNumbers(t *testing.T) {
	left := &TreeNode{
		Val: 2,
	}

	right := &TreeNode{
		Val: 3,
	}
	root := &TreeNode{
		Val:   1,
		Left:  left,
		Right: right,
	}
	sum := sumNumbers(root)
	if sum != 25 {
		t.Errorf("sumNumbers resut is %d ,expected %d", sum, 25)
	} else {
		t.Log("success")
	}
}
