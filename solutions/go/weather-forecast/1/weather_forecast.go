// Package weather hat can forecast the current weather condition of various cities in Goblinocus.
package weather

var (
    //CurrentCondition string.
	CurrentCondition string
    //CurrentLocation  string.
	CurrentLocation  string
)

// Forecast return Forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
