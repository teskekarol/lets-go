package main

func processQueries(queries []int, m int) []int {
	to := len(queries)
	results := make([]int, to)
	p := initP(m)

	for i := 0; i < to; i++ {
		divider := queries[i]
		indexOfDivider := indexOf(divider, p)
		results[i] = indexOfDivider
		if indexOfDivider == 0 {
			continue
		}
		p = append([]int{divider}, append(p[:indexOfDivider], p[indexOfDivider+1:]...)...)
	}
	return results
}

func indexOf(divider int, arr []int) int {
	for i, val := range arr {
		if val == divider {
			return i
		}
	}
	return -1
}

func initP(size int) []int {
	temp := make([]int, size)
	for i := 0; i<size; i++ {
		temp[i] = i+1
	}
	return temp
}