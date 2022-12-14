package interview

import (
	"fmt"
	"leetcode/array"
	"math/big"
	"strings"
	"time"
)

// 给定一个数组，数组全是整数。要求将奇数全部按顺序排在数组前面，偶数排在后面，不能新建其它数组
// 题解一、
func Array1(list []int) []int {
	length := 0
	for i := 0; i < len(list); i++ {
		if list[i]%2 == 1 {
			length++
		}
	}
	left := 0
	right := length
	for left < length && right < len(list) {
		if list[left]%2 == 0 && list[right]%2 == 1 {
			aa := list[left]
			list[left] = list[right]
			list[right] = aa
			left++
			right++
			continue
		}
		if list[left]%2 == 1 {
			left++
		}
		if list[right]%2 == 0 {
			right++
		}
	}
	array.Sort(list[:length])
	array.Sort(list[length:])
	return list
}

// 题解二、
func Array(list []int) []int {
	for i := 0; i < len(list); i++ {

	}
	return nil
}

// 给定一个数组1表示种有花,0表示没有。花与花之间不能紧挨着种。n表示可以种植的数量，可以则返回true，否则false
func Flower(list []int, n int) bool {
	left := 0
	right := 0
	m := 0
	for i := 0; i < len(list); i++ {
		if list[i] == 0 {
			left = i - 1
			right = i + 1
			if left < 0 {
				if list[right] == 0 {
					m++
					i++
				}
			} else if left >= 0 && right < len(list) {
				if list[left] == 0 && list[right] == 0 {
					m++
					i++
				}
			} else if right >= len(list) {
				if list[left] == 0 {
					m++
					i++
				}
			}
		}
	}
	fmt.Println(m)
	return m >= n
}

// 找出两个字符串的最大公共交集
func FindString(s1, s2 string) string {
	str := ""
	if len(s1) < len(s2) {
		str = do(s1, s2)
	} else {
		str = do(s2, s1)
	}
	return str
}
func do(min, max string) string {
	len1 := len(min)
	str := ""
	for i := 0; i < len1; i++ {
		s := string(min[i])
		index := i
		for strings.Contains(max, s) {
			if len(str) < len(s) {
				str = s
			}
			if index < len1-1 {
				s = s + string(min[index+1])
			} else {
				break
			}
			index++
		}
	}
	return str

}

/*
斐波那契数列
*/
func FibonacciSequence(num int) {
	start := time.Now()
	nums := make([]*big.Int, num)
	for i := 0; i < num; i++ {
		if i <= 1 {
			nums[i] = big.NewInt(1)
		} else {
			temp := new(big.Int)
			nums[i] = temp.Add(nums[i-1], nums[i-2])
		}
		fmt.Printf("数位第%v位: %v\n", i, nums[i])
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("执行所耗: %v\n", delta)
}

/*
快速排序
*/
func FastSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	//mid中间值(假定mid = values[0] 为中间值)	, i下标
	mid, i := nums[0], 1
	head, end := 0, len(nums)-1
	for head < end {
		// 把 小于mid 的值放到 head前面 反则亦然
		if nums[i] > mid {
			nums[i], nums[end] = nums[end], nums[i]
			end--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	// 中间值替换成mid
	nums[head] = mid
	FastSort(nums[:head])
	FastSort(nums[head+1:])
}

/*
冒泡排序
*/
func BubbleSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	time.Sleep(time.Nanosecond * 100)
}

/*
插入排序
*/
func InsertSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1], nums[j] = nums[j], key
			j--
		}
	}
}

/*
合并排序 将两个或多个有序数组进行排序
*/
func MergeSort(a []int, left, right int) {
	if left == right {
		return
	}
	if left < right {
		//拆分成n份进行排序
		mid := (left + right) / 2
		MergeSort(a, left, mid)
		MergeSort(a, mid+1, right)
		merge(a, left, right)
	}
}

//合并数组
func merge(a []int, left, right int) {
	m := right - left + 1
	//临时数组b
	b := make([]int, m)
	left0 := left
	// 数组下标
	i := 0
	//中间数组的中间值
	mid := (left + right) / 2
	k := mid + 1
	for left <= mid && k <= right {
		if a[left] < a[k] {
			b[i] = a[left]
			i++
			left++
		} else {
			b[i] = a[k]
			i++
			k++
		}
	}
	if left > mid {
		for k <= right {
			b[i] = a[k]
			i++
			k++
		}
	}
	if k > right {
		for left <= mid {
			b[i] = a[left]
			left++
			i++
		}
	}
	for j := 0; j < m; j++ {
		a[left0] = b[j]
		left0++
	}
}

/* func Cmp(a, b int) bool {
	// 判断ab是否同为奇数或偶数
	if a%2 != b%2 {
		// 判断a是否为偶数数
		return a%2 == 0
	} else {
		return a > b
	}
}
*/
/* func fast(values []int) {
	if len(values) == 0 {
		return
	}
	mid, i := values[0], 1
	head, end := 0, len(values)-1
	for head < end {
		if values[i] > mid {
			values[i], values[end] = values[end], values[i]
			end--
		} else {
			values[head], values[i] = values[i], values[head]
			i++
			head++
		}
	}
	values[head] = mid
	fast(values[:head])
	fast(values[head+1:])
}
*/
