package main

import (
	"fmt"
	"strings"
)

var lowestSOC, highestSOC float64

type SoCOCV map[float64]float64

func PopulateData(input string) SoCOCV {
	d := make(SoCOCV)

	for i, line := range strings.Split(input, "\n") {
		if i == 0 {
			continue
		}
		var soc, ocvA, ocvB float64

		fmt.Sscanf(line, "%v, %v, %v", &soc, &ocvA, &ocvB)
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
	var upper float64 = lowestSOC
	var lower float64 = highestSOC

	for soc := range d {
		if soc > lower && soc < z {
			lower = soc
		}
		if soc < upper && soc > z {
			upper = soc
		}
	}

	y0 := d[lower]
	y1 := d[upper]

	return (y0*(upper-z) + y1*(z-lower)) / (upper - lower)
}

func (d SoCOCV) derivative(z float64) float64 {
	var upper float64 = lowestSOC
	var lower float64 = highestSOC

	for soc := range d {
		if soc > lower && soc < z {
			lower = soc
		}
		if soc < upper && soc > z {
			upper = soc
		}
	}

	return (d[upper] - d[lower]) / (upper - lower)
}
