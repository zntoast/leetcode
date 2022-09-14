package algorithm

import (
	"sort"
)

//插入排序
func InsertionSort(nums []int) []int {
	for j := 1; j < len(nums); j++ {
		key := nums[j]
		i := j - 1
		for i >= 0 && nums[i] > key {
			nums[i+1] = nums[i]
			i = i - 1
			nums[i+1] = key
		}
	}
	return nums
}

//分治排序
func Merge(nums []int) []int {
	mid := len(nums) / 2
	sort.Ints(nums[:mid])
	sort.Ints(nums[mid:])
	left := 0
	right := mid
	newNums := make([]int, len(nums))
	for k := range newNums {
		if nums[left] < nums[right] && left < mid {
			newNums[k] = nums[left]
			left++
		} else {
			newNums[k] = nums[right]
			right++
		}
	}
	return newNums
}
