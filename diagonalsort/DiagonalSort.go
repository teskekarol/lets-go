package main

import (
	"fmt"
	"sort"
)

func main() {
	var input = [][]int{{3,3,1,1},{2,2,1,2},{1,1,1,2}}
	diagonalSort(input)
}

func diagonalSort(mat [][]int) [][]int {
	values := make(map[int][]int)
	fmt.Println(mat)
	n := len(mat)
	m := len(mat[0])
	for i, vert := range mat {
		for j, single := range vert {
			values[i-j] = append(values[i-j], single)
		}
	}

	for _, vals := range values {
		sort.Ints(vals)
	}

	for i := n-1; i >= 0; i-- {
		for j := m-1; j>=0; j--{
			diag := i-j
			a := values[diag]
			x, a := a[len(a)-1], a[:len(a)-1]
			values[diag] = a
			mat[i][j] = x
		}
	}

	return mat
}