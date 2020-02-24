package main

func main() {

}

type stack []string

func (s *stack) push(toAdd string) stack {
	*s = append(*s, toAdd)
	return *s
}

func (s *stack) pop() stack {
	i := len(*s)
	if i == 0 {
		return nil
	} else {
		aa := *s
		*s = aa[:i-1]
		return *s
	}
}

func (s stack) isEmpty() bool {
	return len(s) == 0
}

func (s stack) peek() string {
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1]
}

func balancedStringSplit(s string) int {
	result := 0
	previous := 0
	counter := 0
	for _, v := range s {
		num := int(v)
		if previous == 0 {
			previous = num
			counter++
			continue
		}
		if previous == num {
			counter++
		} else {
			counter--
			if counter == 0 {
				result++
				previous = 0
			} else {
				previous = num
			}
		}
	}
	return result
}
