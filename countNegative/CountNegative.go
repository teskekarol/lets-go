package main

func main() {
	
}

func countNegatives(grid [][]int) int {
	result := 0
	c := make(chan int)
	for _, row := range grid {
		go binarySearch(row, c)
	}

	for i := 0; i < len(grid); i++ {
		result += <-c
	}
	return result
}

func binarySearch(row []int, c chan int) {
	searchedVal, begin, end := 0,0, len(row)

	for begin < end {
		mid := (begin+end) >> 1
		if row[mid] < searchedVal {
			end = mid
		} else if row[mid] > searchedVal {
			begin = mid + 1
		} else {
			if begin == mid {
				begin++
				break
			}
 			begin = mid
		}
	}
	c <- len(row) - begin
}
