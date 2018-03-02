package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	a := []int{10, 12, 1, 4, 40, 2, 8, 28, 16, 38}
	HeapSort(a)
	fmt.Println(a)
}
