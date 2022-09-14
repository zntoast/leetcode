package leetcodehot100

import (
	"fmt"
	"testing"
)

func TestListNodeAdd(T *testing.T) {
	hd1 := ListNode{1, nil}
	hd2 := ListNode{2, &hd1}
	hd3 := ListNode{3, &hd2}
	hd4 := ListNode{4, &hd3}
	ListNode := AddTwoNumbers(&hd4, &hd4)
	fmt.Println(ListNode)
}

func TestLengthOfLongestSubstring(T *testing.T) {
	s := "abbabcbb"
	fmt.Println(LengthOfLongestSubstring(s))
}

func TestFindMedianSortedArrays(T *testing.T) {
	num1 := []int{3}
	num2 := []int{-2, -1}
	fmt.Println(FindMedianSortedArrays(num1, num2))
}

func TestIsPalindrome(T *testing.T) {
	x := 123321
	fmt.Println(IsPalindrome(x))
}

func TestIsMatch(T *testing.T) {
	s := "abb"
	p := ".*"
	fmt.Println(IsMatch(s, p))
}

func TestMaxArea(T *testing.T) {
	list := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(MaxArea(list))
}

func TestThreeSum(T *testing.T) {
	nums := []int{3, 0, -2, -1, 1, 2}
	fmt.Println(ThreeSum(nums))
}

func TestLetterCombinations(T *testing.T) {
	aa := "23"
	fmt.Println(LetterCombinations(aa))
}

func TestSearch(T *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(Search(nums, 0))
}

func TestAddTwoNumbers(T *testing.T) {
	l3 := &ListNode{3, nil}
	l2 := &ListNode{4, l3}
	l1 := &ListNode{2, l2}
	r3 := &ListNode{4, nil}
	r2 := &ListNode{6, r3}
	r1 := &ListNode{5, r2}
	AddTwoNumbers(l1, r1)
}

func TestRemoveElement(T *testing.T) {
	nums := []int{3, 2, 2, 3}
	RemoveElement(nums, 3)
}

func TestFindSubstring(T *testing.T) {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	fmt.Println(FindSubstring(s, words))
}
