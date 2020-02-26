package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	helper(root, val)
	return root
}

func helper(root *TreeNode, val int) {
	if val > root.Val {
		if root.Right != nil {
			helper(root.Right, val)
		}else {
			node := TreeNode{Val: val}
			root.Right = &node
		}
	} else {
		if root.Left != nil {
			helper(root.Left, val)
		}else {
			node := TreeNode{Val: val}
			root.Left = &node
		}
	}
}
