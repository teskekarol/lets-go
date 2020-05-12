package main

func destCity(paths [][]string) string {
	results := make(map[string]int)
	for _, pair := range paths {
		left := pair[0]
		right := pair[1]

		results[left] += 1
		if _, ok2 := results[right]; ok2 {
			results[right] += 1
		} else {
			results[right] = 0
		}
	}

	for city, outRoutes := range results {
		if outRoutes == 0 {
			return city
		}
	}
	return "shall not happen"
}
