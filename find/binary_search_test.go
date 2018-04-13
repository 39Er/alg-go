package find

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(BinarySearch(arr, 9))
}
