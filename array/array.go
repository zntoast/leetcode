package array

// 删除排序数组中的重复项
// 例子: nums := []int{1, 1, 2, 2, 3, 3, 4, 5, 6, 7}
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}

// 买股票的最佳时机
// 例子: prices := []int{7, 1, 5, 3, 6, 4}
func MaxProfit(prices []int) int {
	price := 0
	if len(prices) == 0 {
		return 0
	}
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			price += prices[i+1] - prices[i]
		}
	}
	return price
}

// 旋转数组
// 例子: nums := []int{1, 2, 3, 4, 5, 6, 7}
func Rotate(nums []int, k int) {
	array := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i+k <= len(nums)-1 {
			array[i+k] = nums[i]
		} else {
			array[(i+k)%len(nums)] = nums[i]
		}
	}
	for k := range array {
		nums[k] = array[k]
	}
}

// 存在重复元素
// 例子: nums := []int{1, 2, 3, 1}
func ContainsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				// return true
			}
		}
	}
	return false
}

// 只出现一次的数字
// 例子: nums := []int{4, 1, 2, 1, 2}
func SingleNumber(nums []int) int {
	value := 0
	for _, v := range nums {
		value = value ^ v
	}
	return value
}

// 两个数组的交集
func Intersect(nums1 []int, nums2 []int) []int {
	Sort(nums1)
	Sort(nums2)
	top := 0
	down := 0
	array := make([]int, 0)
	for top <= len(nums1)-1 && down <= len(nums2)-1 {
		//如果相等时
		if nums1[top] == nums2[down] {
			array = append(array, nums1[top])
			top++
			down++
			continue
		}
		if nums1[top] > nums2[down] {
			down++
			continue
		}
		if nums1[top] < nums2[down] {
			top++
			continue
		}

	}
	return array
}

// 冒泡排序
func Sort(num []int) {
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			a := 0
			if num[i] > num[j] {
				a = num[i]
				num[i] = num[j]
				num[j] = a
			}
		}
	}
}

// 加一
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	array := make([]int, len(digits)+1)
	array[0] = 1
	return array
}

// 移动零
func MoveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i] = nums[j]
					nums[j] = 0
					break
				}
			}
		}
	}
}

// 两数之和
func TwoSum(nums []int, target int) []int {
	array := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				array[0] = i
				array[1] = j
				return array
			}
		}
	}
	return array
}

func NoRepetitionNums(nums1, nums2 []int64) []int64 {
	existed := map[int64]struct{}{}
	result := []int64{}
	for _, i := range nums1 {
		existed[i] = struct{}{}
	}
	for _, i := range nums2 {
		if _, ok := existed[i]; ok {
			continue
		}
		result = append(result, i)
	}
	// result := make([]int64, 0)
	// contains := func(n1 int64, ns2 []int64) bool {
	// 	found := false
	// 	for _, s := range ns2 {
	// 		if s == n1 {
	// 			found = true
	// 			break
	// 		}
	// 	}
	// 	return found
	// }
	// for i := 0; i < len(nums1); i++ {
	// 	if !contains(nums1[i], nums2) {
	// 		result = append(result, nums1[i])
	// 	}
	// }
	return result
}
