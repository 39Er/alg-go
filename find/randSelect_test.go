package find

import (
	"fmt"
	"testing"
)

func TestRandSelect(t *testing.T) {
	arr := []int{6, 7, 2, 8, 10, 9, 11, 3, 13, 5}
	//fmt.Println("result1", RandSelect1(arr, 4))
	fmt.Println("result2", RandSelect2(arr, 15))
}
