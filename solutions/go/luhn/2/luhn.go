package luhn

import (
	"strings"
	"unicode"
)

// Valid returns whether the given input is valid as per Luhn formula
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	nonDigitIndex := strings.IndexFunc(input, func(r rune) bool {
		return !unicode.IsDigit(r)
	})

	if nonDigitIndex != -1 {
		return false
	}
	if len(input) <= 1 {
		return false
	}

	var sum int = 0
	for i := 0; i < len(input); i += 1 {
		idx := len(input) - i - 1
		digit := input[idx] - '0'

		if i%2 == 0 {
			sum += int(digit)
		} else {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
			sum += int(digit)
		}
	}

	return sum%10 == 0
}
