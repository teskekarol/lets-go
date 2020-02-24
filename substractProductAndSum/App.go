package main

func subtractProductAndSum(n int) int {
	multiplication, sum := 1, 0

	for n != 0 {
		x := n % 10
		multiplication *= x
		sum += x
		n /= 10
	}

	return multiplication - sum
}
