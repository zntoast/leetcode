package tx

import (
	"fmt"
	"testing"
)

func TestArrayRankTransform(T *testing.T) {
	s := "  +123a45"
	fmt.Println(MyAtoi(s))
}

func TestThreeSum(T *testing.T) {
	nums := []int{3, 0, -2, -1, 1, 2}
	fmt.Println(ThreeSum(nums))
}

func TestThreeSumClosest(T *testing.T) {
	nums := []int{4, 0, 5, -5, 3, 3, 0, -4, -5}
	fmt.Println(ThreeSumClosest(nums, -2))
}

func TestPermute(T *testing.T) {
	// nums := []int{5, 4, 6, 2}
	nums := []int{1, 2, 3}
	fmt.Println(Permute(nums))
}

func TestMaxSubArray(T *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(MaxSubArray(nums))
}
