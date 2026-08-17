package space

// Planet type representing the Planet name
type Planet string

const secondsInEarthYear float64 = 31557600

var m map[Planet]float64 = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns the orbital-age for the given seconds on a particulat Planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / secondsInEarthYear / m[planet]
}
