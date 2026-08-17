package hamming

type HammingDistanceError string

func (msg HammingDistanceError) Error() string {
	return string(msg)
}

// Distance returns the number of different character in the given strings
// For strings with unequal lengths, it returns an Error
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, HammingDistanceError("Hamming Distance is not defined for unequal strings")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
