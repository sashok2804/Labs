package str

func Reverse(s string) string {
	newStr := ""

	for i := len(s) - 1; i >= 0; i-- {
		newStr += string(s[i])
	}

	return newStr
}
