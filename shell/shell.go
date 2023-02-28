package shell

import "sort"

//数组序号转换
func ArrayRankTransform(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	sort.Ints(newArr)
	ranks := map[int]int{}
	for _, v := range newArr {
		if _, ok := ranks[v]; !ok {
			ranks[v] = len(ranks) + 1
		}
	}
	for k, v := range arr {
		arr[k] = ranks[v]
	}
	return arr
}
