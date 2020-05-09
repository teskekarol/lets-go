package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := maximum(candies)

	result := make([]bool, len(candies))

	for i, candy := range candies {
		maximumPossibleValue := candy + extraCandies
		if maximumPossibleValue >= max {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}

func maximum(candies []int) int {
	max := 0
	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}
	return max
}