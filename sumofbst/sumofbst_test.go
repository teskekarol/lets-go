package main

import "testing"

func TestSumOfBst(t *testing.T) {
	//given
	node3 := TreeNode{Val: 3}
	node7 := TreeNode{Val: 7}
	node5 := TreeNode{5, &node3, &node7}
	node18 := TreeNode{Val: 18}
	node15 := TreeNode{15, nil, &node18}
	node10 := TreeNode{10, &node5, &node15}

	//when
	result := rangeSumBST(&node10, 7, 15)

	//then
	expected := 32
	if result != expected {
		t.Errorf("Result is %d, expected %d", result, expected)
	}

}