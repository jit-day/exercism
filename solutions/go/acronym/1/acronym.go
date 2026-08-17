package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns an abbreviation for given string by taking first letter of each word
func Abbreviate(s string) string {
	text := getAlphabets(s)
	words := strings.Fields(text)

	builder := strings.Builder{}
	for _, word := range words {
		if len(word) > 0 {
			builder.WriteRune(unicode.ToUpper([]rune(word)[0]))
		}
	}

	return builder.String()
}

func getAlphabets(text string) string {
	builder := strings.Builder{}
	for _, ch := range text {
		if unicode.IsLetter(ch) || ch == '\'' {
			builder.WriteRune(ch)
		} else {
			builder.WriteRune(' ')
		}
	}
	return builder.String()
}
