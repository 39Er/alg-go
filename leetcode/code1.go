package leetcode

import (
	"sort"
	"strings"
)

//1
func TwoSum(nums []int, target int) (int, int) {
	length := len(nums)
	m := make(map[int]int, length)
	for i := 0; i < length; i++ {
		num := nums[i]
		if index, ok := m[num]; ok {
			return index, i
		} else {
			m[target-num] = i
		}
	}
	return -1, -1
}

//14
// Input: ["flower","flow","flight"]
// Output: "fl"
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sub := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(sub) > 0 && !strings.HasPrefix(strs[i], sub) {
			sub = sub[0 : len(sub)-1]
			if len(sub) == 0 {
				return ""
			}
		}
	}
	return sub
}

//15 three sum
/**
Given array nums = [-1, 0, 1, 2, -1, -4],
A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
func ThreeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return result
}
