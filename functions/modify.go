package functions
func Modify(txt string) string {
	txt = IsValid(txt)
	txt = FormatArticles(txt)
	txt = FormatCommands(txt)
	txt = FormatPunctuation(txt)
	txt = formatNewLines(txt)
	return txt
}