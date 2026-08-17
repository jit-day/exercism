package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	channels := make([]chan FreqMap, len(s))
	for idx, text := range s {
		go func(ch chan FreqMap, text string) {
			ch <- Frequency(text)
		}(channels[idx], text)
	}

	m := FreqMap{}
	for i := 0; i < len(channels); i += 1 {
		m2 := <-channels[i]
		for key, value := range m2 {
			m[key] += value
		}
	}

	return m
}
