package raindrops

import "fmt"

// Convert returns a string representing raindrop sound based on given number
// number % 3 == 0 -> 'Pling'
// number % 5 == 0 -> 'Plang'
// number % 7 == 0 -> 'Plong'
// else -> number
func Convert(number int) string {
	result := ""
	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result = fmt.Sprint(number)
	}

	return result
}
