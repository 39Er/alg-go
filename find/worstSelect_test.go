package find

import (
	"fmt"
	"testing"
)

func TestWorstSelect(t *testing.T) {
	arr := []int{6, 7, 2, 8, 10, 9, 11, 3, 13, 5, 12, 15, 5}
	fmt.Println(WorstSelect(arr, 4))
}
