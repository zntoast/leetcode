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
	list := PreorderTraversal2(head)
	fmt.Printf("list: %v\n", list)
}

func TestAddNode(T *testing.T) {
	head := &TreeNode{5, nil, nil}
	nums := []int{4, 6, 8, 7, 9}
	for _, v := range nums {
		head.TreeAdd(v)
	}
}

func TestupdateCounts(T *testing.T) {
	// Mocking the data
	data := &CustomerChoroplethMap{
		Adcode: "parent",
		Name:   "Parent Area",
		Sub: []*CustomerChoroplethMap{
			{
				Adcode: "child1",
				Name:   "Child Area 1",
				Sub: []*CustomerChoroplethMap{
					{
						Adcode: "grandchild1",
						Name:   "Grandchild Area 1",
						Sub:    nil,
					},
				},
			},
			{
				Adcode: "child2",
				Name:   "Child Area 2",
				Sub:    nil,
			},
		},
	}
	countMap := map[string]int64{
		"parent":      10,
		"child1":      5,
		"grandchild1": 2,
		"child2":      3,
	}
	updateCounts(data, countMap)

	fmt.Printf("Updated Counts:\n%+v\n", data)
}
