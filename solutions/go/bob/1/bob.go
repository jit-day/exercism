package bob

import (
	"strings"
	"unicode"
)

// Hey returns a 'Response' for the given 'remark'
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	askedQuestion := strings.HasSuffix(remark, "?")
	yelled := isUpper(remark)

	if remark == "Bob" || remark == "" {
		return "Fine. Be that way!"
	} else if askedQuestion && yelled {
		return "Calm down, I know what I'm doing!"
	} else if askedQuestion {
		return "Sure."
	} else if yelled {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}

func isUpper(text string) bool {
	hasLetter := false
	for _, ch := range text {
		if unicode.IsLetter(ch) {
			hasLetter = true
			if !unicode.IsUpper(ch) {
				return false
			}
		}
	}
	return hasLetter
}
