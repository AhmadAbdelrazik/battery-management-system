package main

import (
	"fmt"
	"math"
	"strings"
)

var lowestSOC, highestSOC float64

type SoCOCV map[float64]float64

func NewData(input string) SoCOCV {
	d := make(SoCOCV)

	for i, line := range strings.Split(input, "\n") {
		if i == 0 {
			continue
		}
		var soc, ocvA, ocvB float64

		fmt.Sscanf(strings.ReplaceAll(line, ",", " "), "%v %v, %v", &soc, &ocvA, &ocvB)
		d[soc] = ocvA
		if soc < lowestSOC {
			lowestSOC = soc
		}
		if soc > highestSOC {
			highestSOC = soc
		}
	}

	return d
}

func (d SoCOCV) GetVoltage(z float64) float64 {
	z1, z2 := math.Floor(z)+0.5, z

	if math.Ceil(z)-z < z-math.Floor(z) {
		z2 = math.Ceil(z)
	} else {
		z2 = math.Floor(z)
	}

	return (d[z1] + d[z2]) / 2
}

func (d SoCOCV) derivative(z float64) float64 {
	z1, z2 := math.Floor(z)+0.5, z

	if math.Ceil(z)-z < z-math.Floor(z) {
		z2 = math.Ceil(z)
	} else {
		z2 = math.Floor(z)
	}

	return (d[z1] - d[z2]) / (z1 - z2)
}
