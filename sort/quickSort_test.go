package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{49, 38, 65, 97, 76, 13, 27, 49, 78, 34, 12, 64, 1, 8}
	fmt.Println("before——————>")
	fmt.Println(arr)
	QuickSort(arr)
	fmt.Println("after——————>")
	fmt.Println(arr)
}
