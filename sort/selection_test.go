package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{10, 3, 6, 7, 2, 5, 4}
	fmt.Println("排序前：", arr)
	fmt.Println("排序后：", SelectionSort(arr))
}
