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
