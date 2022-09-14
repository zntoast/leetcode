package leetcode75

import (
	"fmt"
	"testing"
)

func TestRunningSum(T *testing.T) {
	array := []int{1, 2, 3, 4}
	fmt.Println(RunningSum(array))
}

func TestPivotIndex(T *testing.T) {
	array := []int{-1, -1, 0, 1, 1, 0}
	fmt.Println(PivotIndex(array))
}

func TestIsIsomorphic(T *testing.T) {
	s := "badc"
	t := "baba"
	fmt.Println(IsIsomorphic(s, t))
}

func TestIsSubsequence(T *testing.T) {
	s := "ace"
	t := "abcde"
	fmt.Println(IsSubsequence(s, t))
}

func TestMergeTwoLists(T *testing.T) {
	s := "ace"
	t := "abcde"
	fmt.Println(IsSubsequence(s, t))
}

func TestMiddleNode(T *testing.T) {
	hd1 := ListNode{1, nil}
	hd2 := ListNode{2, &hd1}
	hd3 := ListNode{3, &hd2}
	hd4 := ListNode{4, &hd3}
	// hd5 := ListNode{5, &hd4}
	fmt.Println(MiddleNode(&hd4))
}

var (
	hd1 = ListNode{1, nil}
	hd2 = ListNode{2, &hd1}
	hd3 = ListNode{3, &hd2}
	hd4 = ListNode{4, &hd3}
)

func TestDetectCycle(T *testing.T) {
	hd1.Next = &hd3
	fmt.Println(DetectCycle(&hd4))
}

func TestLongestPalindrome(T *testing.T) {
	s := "dadwadhao"
	fmt.Println(LongestPalindrome(s))
}

func TestSearch(T *testing.T) {
	nums := []int{-1, 0, 3, 5, 9}
	fmt.Println(Search(nums, 4))
}

func TestLowestCommonAncestor(T *testing.T) {
	LowestCommonAncestor(Head, left5, nil)
}

func TestFloodFill(T *testing.T) {
	image := [][]int{{0, 0, 0}, {1, 0, 0}}
	FloodFill(image, 1, 0, 2)
	for k := range image {
		fmt.Println(image[k])
	}
}

func TestNumIslands(T *testing.T) {
	grid := [][]byte{{0, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}
	fmt.Println(NumIslands(grid))
}

func TestFindAnagrams(T *testing.T) {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(FindAnagrams(s, p))
}

func TestCharacterReplacement(T *testing.T) {
	s := "AABABBA"
	fmt.Println(CharacterReplacement(s, 1))
}

func TestBackspaceCompare(T *testing.T) {
	s := "a##c"
	t := "#a#c"
	fmt.Println(BackspaceCompare(s, t))
}
