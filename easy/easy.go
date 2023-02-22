package easy

import "sort"

/*
1005. K 次取反后最大化的数组和
*/
func LargestSumAfterKNegations(nums []int, k int) int {
	min_index := 0
	result := 0
	sort.Ints(nums)
	for i := range nums {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
		if nums[min_index] > nums[i] {
			min_index = i
		}
	}
	if k%2 == 1 {
		nums[min_index] *= -1
	}
	for _, v := range nums {
		result += v
	}
	return result
}
