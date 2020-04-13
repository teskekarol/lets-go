package main

func main() {
	input := []int{17, 18, 5, 4, 6, 1}
	result := replaceElements(input)
	print(result)
}

func replaceElements(arr []int) []int {
	max := arr[len(arr)-1]
	for i := len(arr) - 1; i >= 0; i-- {
		probablyNextMax := arr[i]
		arr[i] = max

		if probablyNextMax > max {
			max = probablyNextMax
		}
	}
	arr[len(arr)-1] = -1
	return arr
}
