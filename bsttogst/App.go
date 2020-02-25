package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node0 := TreeNode{Val: 0}
	node3 := TreeNode{Val: 3}
	node2 := TreeNode{2, nil, &node3}
	node1 := TreeNode{1, &node0, &node2}
	node5 := TreeNode{Val: 5}
	node8 := TreeNode{Val: 8}
	node7 := TreeNode{7, nil, &node8}
	node6 := TreeNode{6, &node5, &node7}
	node4 := TreeNode{4, &node1, &node6}

	result := bstToGst(&node4)
	if result!= nil{
		fmt.Print("echo")
	}
}

func bstToGst(root *TreeNode) *TreeNode {
	result := 0
	helper(root, &result)
	return root
}

func helper(node *TreeNode, res * int) {

	if node.Left == nil && node.Right == nil {
		*res += node.Val
		node.Val = *res
		return
	}
	if node.Right != nil {
		helper(node.Right, res)
	}
	*res += node.Val
	node.Val = *res
	if node.Left != nil {
		helper(node.Left, res)
	}
}
