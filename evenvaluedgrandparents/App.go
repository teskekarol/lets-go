package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	node9 := TreeNode{9, nil, nil}
	red2 := TreeNode{2, &node9, nil}
	green1 := TreeNode{1, nil, nil}
	green4 := TreeNode{Val: 4}
	red7 := TreeNode{7, &green1, &green4}
	green7 := TreeNode{7, &red2, &red7}
	red5 := TreeNode{Val: 5}
	red3 := TreeNode{3, nil, &red5}
	red1 := TreeNode{Val: 1}
	blue8 := TreeNode{8, &red1, &red3}
	blue6 := TreeNode{6, &green7, &blue8}

	result := sumEvenGrandparent(&blue6)

	fmt.Printf("result is %d", result)
}

func sumEvenGrandparent(root *TreeNode) int {
	var sum = 0
	helper(root, &sum)
	return sum
}

func helper(root *TreeNode, sum *int) {
	leftChild := root.Left
	rightChild := root.Right
	if root.Val%2 == 0 {

		if leftChild != nil {
			if leftChild.Left != nil {
				*sum += leftChild.Left.Val
			}

			if leftChild.Right != nil {
				*sum += leftChild.Right.Val
			}
		}

		if rightChild != nil {
			if rightChild.Left != nil {
				*sum += rightChild.Left.Val
			}

			if rightChild.Right != nil {
				*sum += rightChild.Right.Val
			}
		}

	}
	if leftChild != nil { helper(root.Left, sum) }
	if rightChild != nil { helper(root.Right, sum) }
}
