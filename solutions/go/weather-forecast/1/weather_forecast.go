// Package weather contains all things that you need to do a weather forecast.
package weather

var (
	// CurrentCondition describes what is the weather like in a specific location.
	CurrentCondition string
	// CurrentLocation describes a location where the weather should be forecast.
	CurrentLocation string
)

// Forecast returns a string about the result of current forcast based on the input city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
