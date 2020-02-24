package main

import (
	"strings"
)

func main() {
	test := "Karol"
	letters := "aaa"

	var lettersArr []string

	for _,v := range letters {
		char := string(v)
		lettersArr = append(lettersArr, char)
	}

	result := 0
	for i := 0; i < len(lettersArr); i++{
		result += strings.Count(test, lettersArr[i])
	}
}

func numJewelsInStones(J string, S string) int {
	result := 0
	for _,v := range J {
		char := string(v)
		result += strings.Count(S, char)
	}
	return result
}
