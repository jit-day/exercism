package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	m := map[rune]int{}
	for _, ch := range strings.ToLower(word) {
		if !unicode.IsLetter(ch) {
			continue
		}
		m[ch] += 1
		if m[ch] > 1 {
			return false
		}
	}
	return true
}
