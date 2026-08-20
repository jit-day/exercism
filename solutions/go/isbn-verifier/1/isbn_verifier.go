package isbnverifier

import "unicode"

func IsValidISBN(isbn string) bool {
	sum := 0
	multiplier := 10
	digitCount := 0

	for _, r := range isbn {
		if !unicode.IsDigit(r) && r != 'X' && r != '-' {
			return false
		}

		if r == '-' {
			continue
		}

		digitCount++
		digit := 0

		if r == 'X' && digitCount != 10 {
			return false
		}

		if r == 'X' {
			digit = 10
		} else {
			digit = int(r - '0')
		}

		sum += multiplier * int(digit)
		multiplier--
	}

	return digitCount == 10 && sum%11 == 0
}
