package intermediate

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	array := ThreeSum(nums)
	fmt.Printf("array: %v\n", array)
}
