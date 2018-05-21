package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 5, 7, 12, 4, 6, 9}
	fmt.Println(TwoSum(nums, 7))
}

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"adfds", "bsdafdsf", "casdfdsf"}
	fmt.Println(LongestCommonPrefix(strs))
}

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(ThreeSum(nums))
}
