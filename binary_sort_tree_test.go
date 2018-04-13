package alg_go

import (
	"fmt"
	"testing"
)

func TestCreateBST(t *testing.T) {
	arr := []int{3, 4, 52, 2, 7, 10, 1, 99, 11, 12, 15, 6, 5}
	tree := CreateBST(arr)
	ListBST(tree)
	fmt.Println()
	fmt.Println(DelBST(tree, 3))
	ListBST(tree)
}
