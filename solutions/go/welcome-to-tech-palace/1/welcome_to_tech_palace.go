package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return fmt.Sprintf("%s\n%s\n%s",
		strings.Repeat("*", numStarsPerLine),
		welcomeMsg,
		strings.Repeat("*", numStarsPerLine))
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.Trim(strings.ReplaceAll(oldMsg, "*", ""), " \n")
}
