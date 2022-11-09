package intermediate

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
LC 二叉树的中序遍历
*/
func InorderTraversal(root *TreeNode) []int {
	list := []int{}
	LCR := func(node *TreeNode) {}
	LCR = func(node *TreeNode) {
		if node == nil {
			return
		}
		LCR(node.Left)
		list = append(list, node.Val)
		LCR(node.Right)
	}
	LCR(root)
	return list
}

/*
LC 二叉树的锯齿形层次遍历
*/
func ZigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	dfs := func(node *TreeNode, depth int, res [][]int) [][]int {
		return nil
	}
	dfs = func(node *TreeNode, depth int, res [][]int) [][]int {
		if node == nil {
			return res
		}
		//深度大于长度是追加一行
		if depth >= len(res) {
			res = append(res, []int{})
		}
		if depth%2 == 0 {
			res[depth] = append(res[depth], node.Val)
		} else {
			res[depth] = append([]int{node.Val}, res[depth]...)
		}
		res = dfs(node.Left, depth+1, res)
		res = dfs(node.Right, depth+1, res)
		return res
	}
	res = dfs(root, 0, res)
	return res
}

/*
LC 从前序与中序遍历序列构造二叉树
*/
func BuildTree(preorder []int, inorder []int) *TreeNode {
	in := 0
	pre := 0
	build := func(preorder []int, inorder []int, stop int) *TreeNode {
		return nil
	}
	build = func(preorder, inorder []int, stop int) *TreeNode {
		if pre >= len(preorder) {
			return nil
		}
		if inorder[in] == stop {
			in++
			return nil
		}
		node := &TreeNode{preorder[pre], nil, nil}
		pre++
		node.Left = build(preorder, inorder, node.Val)
		node.Right = build(preorder, inorder, stop)
		return node
	}
	return build(preorder, inorder, math.MinInt32)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
LC 填充每个节点的下一个右侧节点指针
*/
func Connect(root *Node) *Node {
	return nil
}

/*
LC 二叉搜索树中第K小的元素
*/
func KthSmallest(root *TreeNode, k int) int {
	return 0
}

/*
LC 岛屿的数量
*/
func NumIslands(grid [][]byte) int {
	nr := len(grid)
	nc := len(grid[0])
	// 将连接岛屿变为0
	dfs := func(r, c int) {}
	dfs = func(r, c int) {
		grid[r][c] = 48
		if r+1 < nr && grid[r+1][c] == 49 {
			dfs(r+1, c)
		}
		if r-1 >= 0 && grid[r-1][c] == 49 {
			dfs(r-1, c)
		}
		if c-1 >= 0 && grid[r][c-1] == 49 {
			dfs(r, c-1)
		}
		if c+1 < nc && grid[r][c+1] == 49 {
			dfs(r, c+1)
		}
	}
	count := 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == 49 {
				count++
				dfs(r, c)
			}
		}
	}
	return count
}
