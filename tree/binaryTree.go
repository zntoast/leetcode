package tree

import "container/list"

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

// 递归········································
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

// 迭代算法
func PreorderTraversal1(head *TreeNode) []int {
	// 栈
	stack := list.New()
	// 数组
	list := []int{}
	root := head
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.PushBack(root)
			list = append(list, root.Val)
			root = root.Left
		}
		root = stack.Back().Value.(*TreeNode)
		stack.Remove(stack.Back())
		root = root.Right
	}
	return list
}
