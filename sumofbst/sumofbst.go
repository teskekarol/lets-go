package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	node3 := TreeNode{Val: 3}
	node7 := TreeNode{Val: 7}
	node5 := TreeNode{5, &node3, &node7}
	node18 := TreeNode{Val: 18}
	node15 := TreeNode{15, nil, &node18}
	node10 := TreeNode{10, &node5, &node15}

	result := rangeSumBST(&node10, 7, 15)

	fmt.Printf("result is %d", result)
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	moves := 0
	helper(root, L, R, &moves)
	return moves
}

func helper(root *TreeNode, l int, r int, i *int) {
	if root == nil {
		return
	}

	value := root.Val

	if value >= l && value <= r {
		*i += value
	}

	if value < r {
		helper(root.Right, l, r, i)
	}

	if value > l {
		helper(root.Left, l, r, i)
	}

}
