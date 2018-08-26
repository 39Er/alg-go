package leetcode

import (
	"container/list"
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

//18 4sum
/**
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:

Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/
func FourSum(nums []int, target int) [][]int {
	var result [][]int
	length := len(nums)
	if length < 4 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < length-3; i++ {
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[length-3]+nums[length-2]+nums[length-1] < target {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			k, h := j+1, length-1
			for k < h {
				cursum := nums[i] + nums[j] + nums[k] + nums[h]
				if cursum == target {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[h]})
					k++
					h--
					for k < h && nums[k] == nums[k-1] {
						k++
					}
					for k < h && nums[h] == nums[h+1] {
						h--
					}
				} else if cursum < target {
					k++
				} else {
					h--
				}
			}
			for j < length-2 && nums[j+1] == nums[j] {
				j++
			}
		}
		for i < length-3 && nums[i+1] == nums[i] {
			i++
		}
	}
	return result
}

//19. Remove Nth Node From End of List
/**
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head
	for n > 0 && fast != nil {
		n--
		fast = fast.Next
	}
	if fast == nil {
		return slow.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

//20. Valid Parentheses
/**
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/
func IsValid(s string) bool {
	stack := list.New()
	for _, sub := range s {
		str := fmt.Sprintf("%c", sub)
		if str == "(" {
			stack.PushBack(")")
		} else if str == "[" {
			stack.PushBack("]")
		} else if str == "{" {
			stack.PushBack("}")
		} else {
			if stack.Len() == 0 {
				return false
			}
			e := stack.Back()
			if e.Value != str {
				return false
			}
			stack.Remove(e)
		}
	}
	if stack.Len() != 0 {
		return false
	}
	return true
}

//21. Merge Two Sorted Lists
/**
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

//递归
func MergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = MergeTwoLists1(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoLists1(l1, l2.Next)
		return l2
	}
}

//迭代
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := new(ListNode)
	dummy := cur
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	}
	if l2 == nil {
		cur.Next = l1
	}
	return dummy.Next
}

//22. Generate Parentheses
/**
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
func GenerateParenthesis(n int) []string {
	return backtrack("", 0, 0, n)
}
func backtrack(str string, open int, close int, n int) (list []string) {
	if len(str) == n*2 {
		return []string{str}
	}
	if open < n {
		list = append(list, backtrack(str+"(", open+1, close, n)...)
	}
	if close < open {
		list = append(list, backtrack(str+")", open, close+1, n)...)
	}
	return list
}

//23. Merge k Sorted Lists
/**
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/
func MergeKLists_bad(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	tmp := lists
	lists = nil
	for _, n := range tmp {
		if n != nil {
			lists = append(lists, n)
		}
	}
	if len(lists) == 0 {
		return nil
	}
	cur := new(ListNode)
	dummy := cur
	min := lists[0].Val
	index := 0
	for len(lists) > 1 {
		for i, list := range lists {
			if list != nil {
				if list.Val < min {
					min = list.Val
					index = i
				}
			} else {
				lists = append(lists[:i], lists[i+1:]...)
			}
		}
		cur.Next = lists[index]
		cur = cur.Next
		lists[index] = lists[index].Next
		if lists[index] == nil {
			lists = append(lists[:index], lists[index+1:]...)
			index = 0
		}
		min = lists[index].Val
	}
	cur.Next = lists[0]
	return dummy.Next
}

func MergeKLists(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return MergeTwoLists(lists[0], lists[1])
	default:
		n := len(lists) / 2
		l1 := MergeKLists(lists[:n])
		l2 := MergeKLists(lists[n:])
		return MergeTwoLists(l1, l2)
	}
}

//24. Swap Nodes in Pairs
/**
Given a linked list, swap every two adjacent nodes and return its head.

Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
Note:

Your algorithm should use only constant extra space.
You may not modify the values in the list's nodes, only nodes itself may be changed.

*/
func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head.Next
	i := 0
	var prev *ListNode
	var cur *ListNode
	var next *ListNode
	var nextNext *ListNode
	for head.Next != nil {
		if i%2 == 0 {
			cur = head
			next = head.Next
			nextNext = head.Next.Next
			head = head.Next
			head.Next = cur
			head.Next.Next = nextNext
			if prev != nil {
				prev.Next = next
			}
		} else {
			prev = head
		}
		head = head.Next
		i = i + 1
	}
	return dummy.Next
}
func PrintList(node *ListNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d\t", node.Val)
	for node.Next != nil {
		node = node.Next
		fmt.Printf("%d\t", node.Val)
	}
	fmt.Println()
}

//25. Reverse Nodes in k-Group
/**
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

Note:

Only constant extra memory is allowed.
You may not alter the values in the list's nodes, only nodes itself may be changed.
*/

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode)
	dummy.Next = head
	nodes := make([]*ListNode, k+1)
	var prev *ListNode
	for head != nil {
		nodes[0] = head
		for j := 1; j < k; j++ {
			if head.Next != nil {
				nodes[j] = head.Next
				head = head.Next
			} else {
				return dummy.Next
			}
		}
		nodes[k] = head.Next
		head = head.Next
		if prev != nil {
			prev.Next = nodes[k-1]
		} else {
			dummy.Next = nodes[k-1]
		}
		for n := 1; n < k; n++ {
			nodes[n].Next = nodes[n-1]
		}
		nodes[0].Next = nodes[k]
		prev = nodes[0]
	}
	return dummy.Next
}
