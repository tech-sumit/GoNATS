package lib

func ReverseString(value string) string {
	data := []rune(value)
	var result []rune
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	return string(result)
}
