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

/*
统计n位数以内有多少个素数,不计入 0 ,1 ,埃筛法
*/
func Eratosthenes(n int) int {
	if n < 2 {
		fmt.Println("n 不能小于2")
		return n
	}
	count := 0
	// 素数存储
	isPrims := make([]bool, n)
	for i := 2; i < n; i++ {
		//判断是否为素数 ,是则进入
		if !isPrims[i] {
			count++
			for j := 2 * i; j < n; j += i {
				//将非素数改为true
				isPrims[j] = true
			}
		}
	}
	return count
}

/*
删除有序数组的重复项,每个元素只能出现一次,返回删除后数组的长度
考察双指针
*/
func RemoveDuplicates(nums []int) int {
	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}

/*
寻找数组的中心下标,使得左边数组的和跟右边数组的和相等
*/
func PivotIndex(nums []int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		if total == num {
			return i
		}
		num -= nums[i]
	}
	return -1
}

/*
 在不使用官方包情况, 算x的平方根 取整数
 二分查找
*/
func BinarySearch(x int) int {
	if x < 2 {
		return 1
	}
	index, left, right := -1, 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			index = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return index
}
