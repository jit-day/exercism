package twofer

import "fmt"

// ShareWith returns a string of the format "One for [name], one for me."
// If name is not given, then default value of 'you' is taken.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
