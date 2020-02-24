package main

import "fmt"

func main() {
	const name, age = "Karol", 18
	fmt.Printf("Name: %s, aged: %d", name, age)

}

func numberOfSteps(input int) int {
	if input == 0 {
		return 0
	}

	if input%2 == 0 {
		return 1 + numberOfSteps(input/2)
	} else {
		return 1 + numberOfSteps(input-1)
	}
}

func numberOfStepsBasic(input int) int {
	var result = 0
	var current = input

	for current != 0 {
		if current%2 == 0 {
			current = current / 2
		} else {
			current--
		}
		result++
	}

	return result
}
