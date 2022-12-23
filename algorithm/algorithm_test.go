package algorithm

import (
	"fmt"
	"testing"
)

func TestInsertionSort(T *testing.T) {
	sums := []int{5, 3, 1, 4, 7}
	fmt.Println(InsertionSort(sums))
}

// func TestMerge(T *testing.T) {
// 	sums := []int{5, 3, 1, 4, 10, 2, 6, 2, 8, 9}
// 	fmt.Println(Merge(sums))
// }
