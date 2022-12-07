package tree

import (
	"fmt"
	"testing"
)

func TestPreorderTraversal1(T *testing.T) {
	node2 := &TreeNode{3, nil, nil}
	node1 := &TreeNode{2, node2, nil}
	left := &TreeNode{4, nil, nil}
	head := &TreeNode{1, left, node1}
	list := PreorderTraversal1(head)
	fmt.Printf("list: %v\n", list)
}
