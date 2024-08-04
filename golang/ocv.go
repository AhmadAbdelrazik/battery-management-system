package main

import (
	"fmt"
	"math"
	"strings"
)

type SoCOCV map[float64]float64

// NewData: Reads an .csv file that contains soc, ocv readings
// and map it to a map, soc values ranges from 0 to 100, with 0.5 step
// meaning a 201 readings mapping to the soc percentage.
func NewData(input string) SoCOCV {
	d := make(SoCOCV)

	for i, line := range strings.Split(input, "\n") {
		if i == 0 {
			continue
		}
		var soc, ocv float64

		fmt.Sscanf(strings.ReplaceAll(line, ",", " "), "%v %v", &soc, &ocv)
		d[soc] = ocv
	}

	return d
}

// GetVoltage: Get voltage based on SoC.
// z: SoC value, should be between 0 and 100.
func (d SoCOCV) GetVoltage(z float64) float64 {
	z1, z2 := math.Floor(z)+0.5, z

	if math.Ceil(z)-z < z-math.Floor(z) {
		z2 = math.Ceil(z)
	} else {
		z2 = math.Floor(z)
	}

	return (d[z1] + d[z2]) / 2
}

// Derivative: Gets the rate of change of the voltage to the SoC
// z: SoC value, should be between 0 and 100.
func (d SoCOCV) Derivative(z float64) float64 {
	z1, z2 := math.Floor(z)+0.5, z

	if math.Ceil(z)-z < z-math.Floor(z) {
		z2 = math.Ceil(z)
	} else {
		z2 = math.Floor(z)
	}

	return (d[z1] - d[z2]) / (z1 - z2)
}
