package main

import (
	"fmt"
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

	if upper == lowestSOC {
		return d[lower]
	} else if lower == highestSOC {
		return d[upper]
	} else {
		return (d[lower] + d[upper]) / 2
	}
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

	if upper == lowestSOC || lower == highestSOC {
		return 0
	} else {
		return (d[upper] - d[lower]) / (upper - lower)
	}
}
