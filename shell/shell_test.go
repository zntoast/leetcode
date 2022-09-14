package shell

import (
	"fmt"
	"testing"
)

func TestArrayRankTransform(T *testing.T) {
	arr := []int{40, 10, 20, 30}
	// arr1 := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	fmt.Println(ArrayRankTransform(arr))
}
