package tree

import (
	"fmt"
	"testing"
)

func TestNewIntSkipList(t *testing.T) {
	list := NewIntSkipList(16)
	list.Set(0, 0)
	fmt.Println(list.Len(), list.level())

	list.Set(1, 1)
	fmt.Println(list.Len(), list.level())

	list.Set(2, 2)
	fmt.Println(list.Len(), list.level())

	list.Set(3, 3)
	fmt.Println(list.Len(), list.level())

	list.Set(4, 4)
	fmt.Println(list.Len(), list.level())

	list.Set(5, 5)
	fmt.Println(list.Len(), list.level())

	list.Set(6, 6)
	fmt.Println(list.Len(), list.level())

	list.Set(7, 7)
	fmt.Println(list.Len(), list.level())
	fmt.Println(list.Get(4))
}
