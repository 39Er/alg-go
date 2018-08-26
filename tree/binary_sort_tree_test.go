package tree

import (
	"fmt"
	"testing"
)

func TestCreateBST(t *testing.T) {
	arr := []int{3, 4, 52, 2, 7, 10, 1, 99, 11, 12, 6, 5}
	tree := CreateBST(arr)
	fmt.Print("初始化的二叉排序树：")
	ListBST(tree)
	fmt.Println()
	tree = InsertBST(tree, 15)
	fmt.Print("增加后的二叉排序树：")
	ListBST(tree)
	fmt.Println()
	key := 3
	fmt.Println("删除", key, "是否成功：", DelBST(tree, key))
	fmt.Print("删除后的二叉排序树：")
	ListBST(tree)
	fmt.Println()
	LevelTraverse(tree)
}
