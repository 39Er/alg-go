package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return recursive(root, 0)
}

func recursive(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	} else {
		return recursive(root.Left, sum) + recursive(root.Right, sum)
	}
}
