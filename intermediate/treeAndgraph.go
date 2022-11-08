package intermediate

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
