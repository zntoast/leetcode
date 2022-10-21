package dp

import (
	"fmt"
	"testing"
)

// 最长公共子序列
func TestLongestCommonSubsequence(t *testing.T) {
	text1 := "accde"
	text2 := "ace"
	i := LongestCommonSubsequence(text1, text2)
	fmt.Printf("i: %v\n", i)
}

func TestUniquePaths(t *testing.T) {
	i := UniquePaths(3, 2)
	fmt.Printf("i: %v\n", i)
}

func TestUniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := make([][]int, 3)
	for k := range obstacleGrid {
		obstacleGrid[k] = make([]int, 3)
	}
	obstacleGrid[1][1] = 1
	i := UniquePathsWithObstacles(obstacleGrid)
	fmt.Printf("i: %v\n", i)
}
