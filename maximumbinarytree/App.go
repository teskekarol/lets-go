package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	arr := []int{3,2,1,6,0,5}
	result := constructMaximumBinaryTree(arr)

	if result != nil {
		fmt.Printf("success")
	}
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return nil
}
