package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fourth := ListNode{Val: 0}
	third := ListNode{1, &fourth}
	second := ListNode{0, &third}
	first := ListNode{1, &second}

	result := getDecimalValue(&first)

	fmt.Printf("%d", result)
}

func getDecimalValue(head *ListNode) int {
	var result uint32 = 0
	for nextNode := head; nextNode != nil; nextNode = nextNode.Next {
		if nextNode.Val == 1 {
			result = (result << 1) + 1
		} else {
			result = result << 1
		}
	}
	return int(result)
}
