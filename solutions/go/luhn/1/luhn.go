package luhn

import (
	"strconv"
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

	var sum int64 = 0
	for i := 0; i < len(input); i += 1 {
		idx := len(input) - i - 1
		digit, ok := strconv.ParseInt(string(input[idx]), 10, 32)
		if ok != nil {
			return false
		}

		if i%2 == 0 {
			sum += digit
		} else {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
			sum += digit
		}
	}

	return sum%10 == 0
}
