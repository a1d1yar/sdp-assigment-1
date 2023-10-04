package functions
import (
	"log"
	"strconv"
	"strings"
)
func ConvertHexToDec(text string) string {
	hex := ""
	for _, char := range text {
		if (char >= '0' && char <= '9') || (char >= 'A' && char <= 'F') || (char >= 'a' && char <= 'f') {
			hex += string(char)
		}
	}
	if hex == "" {
		log.Fatal("the number is not hex")
	}
	text = strings.Replace(text, hex, "", -1)
	num, err := strconv.ParseInt(hex, 16, 64)
	LogFatal(err)
	res := strconv.Itoa(int(num))
	text = res + text
	return text
}
func ConvertBinToDec(text string) string {
	bin := ""
	for _, char := range text {
		if char == '0' || char == '1' {
			bin += string(char)
		}
	}
	if bin == "" {
		log.Fatal("the number is not bin")
	}
	text = strings.Replace(text, bin, "", -1)
	num, err := strconv.ParseInt(bin, 2, 64)
	LogFatal(err)
	res := strconv.Itoa(int(num))
	text = res + text
	return text
}
