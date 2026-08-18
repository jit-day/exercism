package collatzconjecture

import "fmt"

type CollatzConjectureError struct {
	value int
}

func (e CollatzConjectureError) Error() string {
	return fmt.Sprintf("Value %v is not valid for Collatz Conjecture", e.value)
}

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, CollatzConjectureError{n}
	}

	count := 0
	for ; n != 1; count += 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}

	return count, nil
}
