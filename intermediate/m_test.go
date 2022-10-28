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
