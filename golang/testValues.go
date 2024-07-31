package main

import (
	"gonum.org/v1/gonum/stat/distuv"
)

func generateTestValues(mean, sigma float64) float64 {
	normalDist := distuv.Normal{
		Mu:    mean,
		Sigma: sigma,
	}

	return normalDist.Rand()
}
