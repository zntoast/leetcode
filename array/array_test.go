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
