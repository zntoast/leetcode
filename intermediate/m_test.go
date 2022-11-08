package intermediate

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	array := ThreeSum(nums)
	fmt.Printf("array: %v\n", array)
}

func TestSetZeroes(t *testing.T) {
	nums := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	SetZeroes(nums)
	fmt.Printf("nums: %v\n", nums)
}

func TestGroupAnagrams(t *testing.T) {
	nums := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	m := GroupAnagrams(nums)
	fmt.Printf("m: %v\n", m)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	i := LengthOfLongestSubstring("abcabcbb")
	fmt.Printf("i: %v\n", i)
}

func TestLongestPalindrome(t *testing.T) {
	s := []string{"cbbd", "babad"}
	for _, v := range s {
		s1 := LongestPalindrome(v)
		fmt.Printf("s1: %v\n", s1)
	}
}

func TestIncreasingTriplet(t *testing.T) {
	s := [][]int{{1, 2, 3, 4, 5}, {2, 1, 5, 0, 4, 6}, {20, 100, 10, 12, 5, 13}, {5, 4, 3, 2, 1}}
	for k, v := range s {
		s1 := IncreasingTriplet(v)
		fmt.Printf("s%d: %v\n", k, s1)
	}
}

func TestAddTwoNumbers(t *testing.T) {
	li2 := &ListNode{3, nil}
	li1 := &ListNode{4, li2}
	li := &ListNode{2, li1}
	le2 := &ListNode{4, nil}
	le1 := &ListNode{6, le2}
	le := &ListNode{5, le1}
	head := AddTwoNumbers(li, le)
	fmt.Printf("head: %v\n", head)
}

func TestOddEvenList(t *testing.T) {
	li4 := &ListNode{5, nil}
	li3 := &ListNode{4, li4}
	li2 := &ListNode{3, li3}
	li1 := &ListNode{2, li2}
	li := &ListNode{1, li1}
	OddEvenList(li)
}

func TestZigzagLevelOrder(t *testing.T) {
	node7 := &TreeNode{7, nil, nil}
	node15 := &TreeNode{15, nil, nil}
	node20 := &TreeNode{20, node15, node7}
	node9 := &TreeNode{9, nil, nil}
	head := &TreeNode{3, node9, node20}
	res := ZigzagLevelOrder(head)
	fmt.Printf("res: %v\n", res)
}
