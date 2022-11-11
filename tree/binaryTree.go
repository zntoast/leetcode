package tree

// 二叉树遍历常用的五种遍历方式
// 1、前序遍历
// 2、中序遍历
// 3、后序遍历
// 4、深度优先遍历 dfs
// 5、宽度优先遍历 bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//1、前序遍历
func PreorderTraversal(head *TreeNode) []int {
	list := make([]int, 0)
	pre := func(node *TreeNode) {}
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		list = append(list, node.Val)
		pre(node.Left)
		pre(node.Right)
	}
	pre(head)
	return list
}

//2、中序遍历
func InorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)
	pre := func(node *TreeNode) {}
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		pre(node.Left)
		list = append(list, node.Val)
		pre(node.Right)
	}
	pre(root)
	return list
}

//3、后序遍历
func PostorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)
	pre := func(node *TreeNode) {}
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		pre(node.Left)
		pre(node.Right)
		list = append(list, node.Val)
	}
	pre(root)
	return list
}
