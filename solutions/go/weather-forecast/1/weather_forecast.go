// Package weather provides tools to describe the forecast of a city.
package weather

var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string
	// CurrentLocation represents the current location.
	CurrentLocation string
)

// Forecast returns a string describing the weather for the given city and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
