package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// This is a placeholder for the actual implementation.
	// The real implementation would go here.
	var testing string
	testing = "The Manager's Path: A Guide for Tech Leaders Navigating Growth and Change (Camille Fournier)"
	testing2 := "Get Better at Anything 12 Maxims for Mastery (Scott H. Young) (Z-Library)"
	aut, newStr := getAuthorAndFormatTitle(testing2)
	aut2, newStr2 := getAuthorAndFormatTitle(testing)

	fmt.Printf("author: %s, string formatted: %s\n", aut, newStr)
	fmt.Printf("author: %s, string formatted: %s\n", aut2, newStr2)

}
func getAuthorAndFormatTitle(str string) (author string, formattedTitle string) {
	// remove (Z-Library) if exists
	formattedTitle = strings.ReplaceAll(str, "(Z-Library)", "")

	// get the author
	re := regexp.MustCompile(`\(([^)]+)\)`)

	matches := re.FindStringSubmatch(formattedTitle)
	if len(matches) == 0 {
		return "", formattedTitle
	}
	formattedTitle = strings.ReplaceAll(formattedTitle, "("+matches[1]+")", "")
	formattedTitle = strings.TrimSpace(formattedTitle)
	return matches[1], formattedTitle
}
