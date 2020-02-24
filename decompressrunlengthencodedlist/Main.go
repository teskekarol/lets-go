package main


func decompressRLElist(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i += 2 {
		value := nums[i+1]
		times := nums[i]
		for ;times > 0; times-- {
			result = append(result, value)
		}
	}
	return result
}

