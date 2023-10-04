package functions
import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)
func FormatArticles(txt string) string {
	// formatting from a -> an
	txt = regexp.MustCompile(`\b(a|A)+\s+[aeiouhAEIOUH]`).ReplaceAllStringFunc(txt, func(s string) string {
		if s[0] == 'a' {
			return "an" + s[1:]
		} else {
			return "An" + s[1:]
		}
	})
	// formatting from an -> a
	txt = regexp.MustCompile(`\b(an|AN|aN|An)+\s+[^aeiouhAEIOUH]`).ReplaceAllStringFunc(txt, func(s string) string {
		if strings.Contains(s, "an") ||
			strings.Contains(s, "aN") {
			return "a " + s[2:]
		}
		return "A " + s[2:]
	})
	// call recursion for recheking all articles
	if regexp.MustCompile(`\b(a|A)+\s+[aeiouhAEIOUH]`).MatchString(txt) || regexp.MustCompile(`\b(an|AN|aN|An)+\s+[^aeiouhAEIOUH]`).MatchString(txt) {
		txt = FormatArticles(txt)
	}
	return txt
}
func FormatCommands(text string) string {
	arg := ""
	noNumArgs := regexp.MustCompile(`\(\s*\w+\s*\)`)
	arrNoNumArgs := noNumArgs.FindAllStringSubmatch(text, -1)
	for _, args := range arrNoNumArgs {
		modifier := args[0]
		updatedModifier := modifier
		switch {
		case strings.Contains(modifier, "(cap"):
			updatedModifier = "(cap, 1)"
		case strings.Contains(modifier, "(up"):
			updatedModifier = "(up, 1)"
		case strings.Contains(modifier, "(low"):
			updatedModifier = "(low, 1)"
		case strings.Contains(modifier, "(hex"):
			updatedModifier = "(hex, 1)"
		case strings.Contains(modifier, "(bin"):
			updatedModifier = "(bin, 1)"
		}
		text = strings.Replace(text, modifier, updatedModifier, -1)
	}
	yesNumArgs := regexp.MustCompile(`\(\w+,\s*(\d+)\)`)
	arrYesNumArgs := yesNumArgs.FindAllStringSubmatch(text, -1)
	for _, args := range arrYesNumArgs {
		arg = args[0]
		number, err := strconv.Atoi(args[1])
		LogFatal(err)
		specificNumArg := regexp.MustCompile(fmt.Sprintf(`(\w+\s*([[:punct:]]*\s*)*){%d}%s`, number, regexp.QuoteMeta(arg)))
		applied := false
		text = specificNumArg.ReplaceAllStringFunc(text, func(match string) string {
			if applied {
				return match
			}
			updatedMatch := strings.Replace(match, arg, "", -1)
			switch {
			case strings.Contains(arg, "(up,"):
				updatedMatch = strings.ToUpper(updatedMatch)
				applied = true
			case strings.Contains(arg, "(cap,"):
				updatedMatch = strings.ToLower(updatedMatch)
				updatedMatch = strings.Title(updatedMatch)
				applied = true
			case strings.Contains(arg, "(low,"):
				updatedMatch = strings.ToLower(updatedMatch)
				applied = true
			case strings.Contains(arg, "(hex,"):
				updatedMatch = ConvertHexToDec(updatedMatch)
				applied = true
			case strings.Contains(arg, "(bin,"):
				updatedMatch = ConvertBinToDec(updatedMatch)
				applied = true
			}
			return updatedMatch
		})
	}
	remainNumberedArgs := regexp.MustCompile(`\((up|low|cap),\s*(\d+)\)`)
	arrOfRemainArgs := remainNumberedArgs.FindAllStringSubmatch(text, -1)
	if len(arrOfRemainArgs) != 0 {
		for _, args := range arrOfRemainArgs {
			arg = args[0]
			specificNumArg := regexp.MustCompile(fmt.Sprintf(`%s\n*`, regexp.QuoteMeta(arg)))
			text = specificNumArg.ReplaceAllString(text, "")
		}
	}
	return text
}
func FormatPunctuation(txt string) string {
	// formating punctuation to make Hello!World. -> Hello! World.
	punctuation := regexp.MustCompile(`(\w*)\s*(\.{3}|(!\?)|[.,!?:;])(\s*\n*)(\n*)`)
	txt = punctuation.ReplaceAllString(txt, "$1$2 $3$4")
	// formating !? to make Hello World !? -> Hello World!?
	excAndQuest := regexp.MustCompile(`(!\?)\s*(!\?)`)
	txt = excAndQuest.ReplaceAllString(txt, "$1")
	// working with quotes
	quotes := regexp.MustCompile(`(['"])\s*(.*?)\s*(['"])`)
	txt = quotes.ReplaceAllString(txt, "$1$2$3")
	return txt
}
func formatNewLines(txt string) string {
	count, line, res := 0, "", ""
	for _, ch := range txt {
		if ch == '\n' {
			count++
		}
	}
	splitted := strings.Split(txt, "\n")
	for i := 0; i <= count; i++ {
		word := strings.Fields(splitted[i])
		line = strings.Join(word, " ")
		res += line
		if i != len(splitted)-1 {
			res += "\n"
		}
	}
	// spacing
	finalSpacingPunct := regexp.MustCompile(`([.,!?:;])\s([.,!?:;])`)
	for finalSpacingPunct.MatchString(res) {
		res = finalSpacingPunct.ReplaceAllString(res, "$1$2")
	}
	return res
}