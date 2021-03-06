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

func TestRemoveNthFromEnd(t *testing.T) {
	tail := &ListNode{5, nil}
	node1 := &ListNode{4, tail}
	node2 := &ListNode{3, node1}
	node3 := &ListNode{2, node2}
	head := &ListNode{1, node3}
	PrintList(head)
	node := RemoveNthFromEnd(head, 2)
	PrintList(node)
}

func TestIsValid(t *testing.T) {
	s := "({([])}"
	fmt.Println(IsValid(s))
}

func TestMergeTwoLists(t *testing.T) {
	tail1 := &ListNode{10, nil}
	node11 := &ListNode{8, tail1}
	node21 := &ListNode{5, node11}
	node31 := &ListNode{2, node21}
	head11 := &ListNode{1, node31}

	tail2 := &ListNode{9, nil}
	node12 := &ListNode{6, tail2}
	node22 := &ListNode{5, node12}
	node32 := &ListNode{3, node22}
	head12 := &ListNode{0, node32}

	PrintList(MergeTwoLists(head11, head12))
}

func TestGenerateParenthesis(t *testing.T) {
	fmt.Println(GenerateParenthesis(3))
}

func TestMergeKLists(t *testing.T) {
	tail1 := &ListNode{10, nil}
	node11 := &ListNode{8, tail1}
	node21 := &ListNode{5, node11}
	node31 := &ListNode{2, node21}
	head11 := &ListNode{1, node31}

	tail2 := &ListNode{9, nil}
	node12 := &ListNode{6, tail2}
	node22 := &ListNode{5, node12}
	node32 := &ListNode{3, node22}
	head12 := &ListNode{0, node32}

	tail3 := &ListNode{9, nil}
	node23 := &ListNode{5, tail3}
	node33 := &ListNode{3, node23}
	head13 := &ListNode{0, node33}
	lists := []*ListNode{head11, head12, head13}
	PrintList(MergeKLists(lists))
}

func TestSwapPairs(t *testing.T) {
	tail1 := &ListNode{10, nil}
	node11 := &ListNode{8, tail1}
	node21 := &ListNode{5, node11}
	node31 := &ListNode{2, node21}
	head11 := &ListNode{1, node31}
	PrintList(SwapPairs(head11))
}

func TestReverseKGroup(t *testing.T) {
	tail1 := &ListNode{10, nil}
	node11 := &ListNode{8, tail1}
	node21 := &ListNode{5, node11}
	node31 := &ListNode{2, node21}
	head11 := &ListNode{1, node31}
	PrintList(ReverseKGroup(head11, 4))
}
