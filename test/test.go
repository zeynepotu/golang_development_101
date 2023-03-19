package test

func ReverseAlphabet(str string) string {
	var result string
	for _, ch := range str {
		if ch >= 'a' && ch <= 'z' {
			result += string('z' - (ch - 'a'))
		} else {
			result += string(ch)
		}
	}
	return result
}
