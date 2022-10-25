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

func TestMinPathSum(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {1, 2, 1}}
	i := MinPathSum(grid)
	fmt.Printf("i: %v\n", i)
}

func TestMinimumTotal(t *testing.T) {
	grid := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	i := MinimumTotal(grid)
	fmt.Printf("i: %v\n", i)
}

func TestMinFallingPathSum(t *testing.T) {
	grid := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	i := MinFallingPathSum(grid)
	fmt.Printf("i: %v\n", i)
}
