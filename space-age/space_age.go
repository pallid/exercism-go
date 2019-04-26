package space

// Planet ...
type Planet string

const earthOrbitalPeriod float64 = 31557600

var periods = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age calculates how many earth years have passed
// in the transmitted number of seconds on another planet
func Age(seconds float64, planet Planet) float64 {

	res := seconds / earthOrbitalPeriod / periods[planet]

	return res
}
