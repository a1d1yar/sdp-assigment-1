package functions
func insertStr(text string, index int, str string) string {
	return text[:index] + str + text[index:]
}