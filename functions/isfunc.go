package functions
import (
	"log"
	"regexp"
)
func IsValid(txt string) string {
	// checking for ascii symbols
	match, err := regexp.MatchString(`^[\x00-\x7F]+$`, txt)
	LogFatal(err)
	if !match {
		log.Fatal("text must be using only ascii symbols")
	}
	// checking for negative numbers
	matchNegativeNumber := regexp.MustCompile(`\(\w+,\s*\-\d+\)`).FindString(txt)
	if matchNegativeNumber != "" {
		log.Fatal("text must include only positive numbers")
	}
	return txt
}