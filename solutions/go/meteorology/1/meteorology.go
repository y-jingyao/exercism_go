package meteorology

import (
	"fmt"
	"strconv"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (t TemperatureUnit) String() string {
	if t == Celsius {
		return "°C"
	}
	if t == Fahrenheit {
		return "°F"
	}
	return ""
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	return strconv.Itoa(t.degree) + " " + t.unit.String()
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (s SpeedUnit) String() string {
	if s == KmPerHour {
		return "km/h"
	}
	if s == MilesPerHour {
		return "mph"
	}
	return ""
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	return strconv.Itoa(s.magnitude) + " " + s.unit.String()
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (m MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity",
		m.location, m.temperature, m.windDirection, m.windSpeed, m.humidity)
}


