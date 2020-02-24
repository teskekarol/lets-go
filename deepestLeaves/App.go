package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	seven := TreeNode{Val: 7}
	four := TreeNode{4, &seven, nil}
	five := TreeNode{Val: 5}
	two := TreeNode{2, &four, &five}
	eight := TreeNode{Val: 8}
	six := TreeNode{6, nil, &eight}
	three := TreeNode{3, nil, &six}
	one := TreeNode{1, &two, &three}

	result := deepestLeavesSum(&one)
	fmt.Printf("result is %d", result)
}

func deepestLeavesSum(root *TreeNode) int {
	currentLevel := 0
	results := make(map[int]int)
	helper(root, results, currentLevel)
	return results[len(results)-1]
}

func helper(node *TreeNode, results map[int]int, currentLevel int) {
	if node == nil {
		return
	}
	results[currentLevel] += node.Val
	currentLevel += 1
	helper(node.Left, results, currentLevel)
	helper(node.Right, results, currentLevel)
}
