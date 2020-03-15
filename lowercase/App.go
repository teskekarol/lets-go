package main

func toLowerCase(str string) string {
	chars := make([]rune, len(str))
	for i, ch := range str {
		if ch >= 'A' && ch <= 'Z' {
			chars[i] = ch + 32
			continue
		}else {
			chars[i] = ch
		}
	}
	return string(chars)
}