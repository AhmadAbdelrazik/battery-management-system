package main

import (
	"math"
)

// all resistors are in milli Ohms
// all Capacitors are in Kilo Farads
type Battery struct {
	R0 float64
	R1 float64
	R2 float64
	C1 float64
	C2 float64
}

type Kalman struct {
	Battery               *Battery
	Pk                    [][]float64
	Xk                    [][]float64
	VoltageSensorVariance [][]float64
	CurrentSensorVariance [][]float64
}

func main() {

}

func OCV(zk float64) float64 {
	return 0
}

func (k *Kalman) Init(b *Battery) {
	k.Xk = [][]float64{{1, 0, 0}}
	k.Pk = [][]float64{
		{0.005, 0, 0},
		{0, 0.0001, 0},
		{0, 0, 0.0001},
	}
	k.Battery = b
	k.VoltageSensorVariance = [][]float64{{math.Pow(0.035, 2)}}
}

// Hk
func (k *Kalman) GetHk() [][]float64 {
	var delta float64
	zk := k.Xk[0][0] // State of charge
	if zk != 0 {
		delta = (OCV(zk) - OCV(zk-0.5)) / 0.5
	} else {
		delta = (OCV(zk+0.5) - OCV(zk)) / 0.5
	}
	Hk := [][]float64{
		{delta, -k.Battery.R1, -k.Battery.R2},
	}

	return Hk
}

// Kalman Gain
func (k *Kalman) GetKk() [][]float64 {
	hk := k.GetHk()

	return MatDiv(MatMul(k.Pk, MatT(hk)), MatAdd(MatMul(MatMul(hk, k.Pk), MatT(hk)), k.VoltageSensorVariance))
}

func (k *Kalman) UpdateYk(i float64) [][]float64 {
	return [][]float64{{0}}
}
