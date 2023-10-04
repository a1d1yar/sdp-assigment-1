package functions
func RemoveByIdx(text string, leftind int, rightind int) string {
	return text[:leftind] + text[rightind:]
}
