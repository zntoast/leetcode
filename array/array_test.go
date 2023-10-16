package array

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2, 2, 3, 3, 4, 5, 6, 7}
	a := RemoveDuplicates(nums)
	fmt.Printf("nums: %v\n", nums[:a])
}

func TestNoRepetitionNums(t *testing.T) {
	datas := []struct {
		num1 []int64
		num2 []int64
	}{
		{[]int64{1, 2, 3, 4, 5}, []int64{2, 3, 4, 5, 6, 7, 8}},
		{[]int64{4, 5, 9}, []int64{2, 3, 4, 5, 6, 7, 8}},
		{[]int64{5, 9}, []int64{2, 3, 4, 5, 6, 7, 8}},
		{[]int64{4, 9}, []int64{2, 3, 4, 5, 6, 7, 8}},
		{[]int64{4, 5}, []int64{2, 3, 4, 5, 6, 7, 8}},
	}
	for k, data := range datas {
		r := NoRepetitionNums(data.num1, data.num2)
		fmt.Printf("第%d条数据: 结果 ：%v\n", k, r)
	}
}

func FuzzRemoveDuplicates(f *testing.F) {
	nums := []int{1, 1, 2, 2, 3, 3, 4, 5, 6, 7}
	for _, v := range nums {
		f.Add(v)
	}
	f.Fuzz(func(t *testing.T) {

	})

}
