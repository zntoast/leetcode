package algorithm

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

