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

func TestThreeSumClosest(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	fmt.Println(ThreeSumClosest(nums, 1))
}

func TestLetterCombinations(t *testing.T) {
	s := "29"
	fmt.Println(LetterCombinations(s))
}

func TestFourSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, -1, -2, -3, -4}
	fmt.Println(FourSum(nums, 0))
}
