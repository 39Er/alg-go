package leetcode

import (
	"fmt"
	"regexp"
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

//16
/**
Given array nums = [-1, 2, 1, -4], and target = 1.
The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/
func ThreeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[n-1]
	dismin := dis(nums[0]+nums[1]+nums[n-1], target)
	for i := 0; i < n-2; i++ {
		tt := target - nums[i]
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		j, k := i+1, n-1
		vmin := dis(nums[j]+nums[k], tt)
		ttmin := nums[j] + nums[k]
		for j < k {
			tti := nums[j] + nums[k]
			if dis(tti, tt) < vmin {
				vmin = dis(tti, tt)
				ttmin = tti
			}
			if tti < tt {
				j++
			} else if tti > tt {
				k--
			} else {
				return target
			}
		}
		if dis(nums[i]+ttmin, target) < dismin {
			dismin = dis(nums[i]+ttmin, target)
			res = nums[i] + ttmin
		}
	}
	return res
}

func dis(v1 int, v2 int) int {
	return (v2 - v1) * (v2 - v1)
}

//17 Letter Combinations of a Phone Number
/**
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

*/
func LetterCombinations(digits string) []string {
	ok, err := regexp.MatchString(`^[2-9]+$`, digits)
	fmt.Println(ok)
	if !ok || err != nil {
		return nil
	}
	dicts := make(map[string][]string)
	dicts["2"] = []string{"a", "b", "c"}
	dicts["3"] = []string{"d", "e", "f"}
	dicts["4"] = []string{"g", "h", "i"}
	dicts["5"] = []string{"j", "k", "l"}
	dicts["6"] = []string{"m", "n", "o"}
	dicts["7"] = []string{"p", "q", "r", "s"}
	dicts["8"] = []string{"t", "u", "v"}
	dicts["9"] = []string{"w", "x", "y", "z"}

	return recursiveCombinations(0, digits, dicts)
}

func recursiveCombinations(startIndex int, digits string, dicts map[string][]string) []string {
	length := len(digits)
	if startIndex == length {
		return nil
	}
	ss := dicts[digits[startIndex:startIndex+1]]
	rest := recursiveCombinations(startIndex+1, digits, dicts)
	var output []string

	for i := 0; i < len(ss); i++ {
		if len(rest) > 0 {
			for j := 0; j < len(rest); j++ {
				output = append(output, ss[i]+rest[j])
			}
		} else {
			output = append(output, ss[i])
		}
	}
	return output
}
