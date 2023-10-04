package main
import (
	"fmt"
	"sdp-assigment-1/functions"
	"os"
	"regexp"
)
func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("The number of arguments must be 2")
		return
	}
	// file validation
	ext1 := regexp.MustCompile(`\.([^\s]+)$`).FindStringSubmatch(os.Args[1])
	ext2 := regexp.MustCompile(`\.([^\s]+)$`).FindStringSubmatch(os.Args[2])
	if ext1 == nil || ext2 == nil {
		fmt.Println("Error: one or both files don't have a valid extension.")
		return
	}
	if ext1[1] != "txt" || ext2[1] != "txt" {
		fmt.Println("Error: one of the output file extensions is not txt.")
		return
	}
	inputFile, err := os.ReadFile(os.Args[1])
	functions.LogFatal(err)
	txt := string(inputFile)
	txt = functions.Modify(txt)
	err = os.WriteFile(os.Args[2], []byte(txt), 0644)
	functions.LogFatal(err)
}
