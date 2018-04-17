package tree

import (
	"fmt"
	"testing"
)

func TestAvlInsert(t *testing.T) {
	arr := []int{3, 5, 2, 7, 9, 0, 10, 2}
	root := NewAvlNode(4)
	for _, key := range arr {
		root = AvlInsert(root, key)
	}
	fmt.Println(displayAvlAsc(root))
	root = AvlInsert(root, 8)
	fmt.Println(displayAvlDesc(root))
}
